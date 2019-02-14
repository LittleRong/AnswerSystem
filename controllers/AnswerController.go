package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"hello/models/event"
	"hello/models/participant"
	"hello/models/participant_haved_answer"
	"hello/models/credit"
	"hello/models/team"
	"hello/models/union"
	"hello/models/user"
	"strconv"
	"time"
)

type AnswerController struct{
	beego.Controller
}

func (this *AnswerController) ShowProblemsPage(){
	this.TplName = "answer/user_problem.html"
}

func (this *AnswerController) GetUserProblems(){
	var result map[string]interface{}
	result = make(map[string]interface{})
	event_id,_ := this.GetInt("event_id")
	//user_id := this.GetSession("user_id")
	user_id := 2
	team_id := 1
	paticipant_id  := 1

	//*****************************1.获取用户题目*************************************************
	problemNum := event.GetProblemNumByEventId(event_id)
	problem,buildFlag := union.GetProblemNoAnswer(user_id,event_id,team_id,paticipant_id,problemNum)

	//*****************************2.获取剩余答题时间*************************************************
	var answer_time float64
	event := event.GetEventByEventId(event_id)
	answer_time,_ = strconv.ParseFloat(event.Answer_time,64)
	if (buildFlag) {
		//获取到生成题目的时间
		participant_time :=participant.GetAnswerTimeByParticipantId(paticipant_id)
		//计算出剩余时间
		now_time:= time.Now()
		left := now_time.Sub(participant_time)
		if(left.Hours() >= answer_time){//时间已经耗尽,left.Hours()获取小时格式
			answer_time = 0
		}else {
			answer_time = answer_time-left.Hours()
		}

	}

	//*****************************3.返回给前端*************************************************
	result["answer_time"] = answer_time
	result["data"] = problem
	this.Data["json"] = result
	this.ServeJSON()
	return

}

type userAnswerStruct struct {
	problem_id string
	q_id string
}

type userAnswer struct {
	single []map[string]string
	multi []map[string]string
	judge []map[string]string
	fill []map[string]string
}

func (this *AnswerController) GetUserAnswers(){
	event_id := 1
	team_id := 1
	paticipant_id  := 1

	//*****************************1.获取该事件评分标准*************************************************
	var creditRule event.CreditRule
	creditRule = event.GetCreditRuleByEventId(event_id)
	beego.Info("========creditRule======",creditRule)

	//*****************************2.获取正确答案*******************************************************
	var correct_answer map[string]interface{}
	correct_answer = make(map[string]interface{})
	correct_answer = participant.GetCorrectAnswerByParticipantId(paticipant_id)
	beego.Info("========correct_answer======",correct_answer)

	//获取用户输入的答案
	single_input := this.Ctx.Request.PostForm.Get("single")
	beego.Info("========input======",single_input)
	var f interface{}
	_ = json.Unmarshal([]byte(single_input), &f)
	single_array := f.([]interface{})
	beego.Info("========f======",f)

	//计算分数,并将用户答案写入participant_haved_answer表
	var user_score float64//用户今日答题总分
	single_right_num := 0//答对的单选题树木
	all_num := 0 //全部题目数
	single_correct_answer := correct_answer["single"].(map[string]interface{})
	for _,value := range single_array{
		//判断是否回答正确
		s := value.(map[string]interface {})
		problem_id := s["problem_id"].(string)
		problem_id_int,_ := strconv.Atoi(problem_id)
		user_answer := s["q_id"].(string)
		right_answer,ok := single_correct_answer[problem_id].(string)
		beego.Info("problem_id=",problem_id,"user_answer=",user_answer," right_answer=",right_answer)
		true_or_false := false
		if (ok && user_answer == right_answer){
			single_right_num++
			true_or_false = true
		}
		all_num++

		//将用户答案写入participant_haved_answer表
		participant_haved_answer.UpdateUserAnswer(paticipant_id,problem_id_int,user_answer,true_or_false)
	}
	user_score = user_score + float64(single_right_num)*creditRule.Single_score
	beego.Info("答对单选题：",single_right_num,"  每题：",creditRule.Single_score,"分 单选题总分:",user_score)


	//判断是否全对，user_all_right
	user_all_right := false
	if(single_right_num == all_num) {
		user_score += creditRule.Person_score
		user_all_right = true
		beego.Info("个人答对额外加分",creditRule.Person_score)
	}

	//更新积分
	//1.更新个人积分
	user_total_credit := participant.UpdateParticipantCredit(paticipant_id,user_score)
	now_time:= time.Now()
	UnixTime:=now_time.Unix()
	now := time.Unix(UnixTime, 0).Format("2006-01-02 15:04:05" )
	reason := ""
	user_score_log := user_score
	if (user_all_right) {
		reason = "当日全部答对额外加分"
		log := credit.Credit_log{Refer_event_id:event_id,Refer_participant_id:paticipant_id,
			Refer_team_id:team_id,Change_time:now,Change_value:creditRule.Person_score,Change_type:2,Change_reason:reason}
		credit.AddCreditLog(log)
		user_score_log = user_score - creditRule.Person_score
	}
	reason = "答题得分"
	log := credit.Credit_log{Refer_event_id:event_id,Refer_participant_id:paticipant_id,
		Refer_team_id:team_id,Change_time:now,Change_value:user_score_log,Change_type:1,Change_reason:reason}
	credit.AddCreditLog(log)

	//2.更新组积分
	team_score := user_score
	//判断是否当日全部答对，若组员全部答对额外加分
	now_date := time.Unix(UnixTime, 0).Format("2006-01-02" )
	event := event.GetEventByEventId(event_id)
	team_allright_flag := credit.WhetherMemberAllRight(team_id,now_date,event.Participant_num)
	if(team_allright_flag == true ){
		//写积分表
		reason = "当日全组全部答对额外加分"
		log := credit.Credit_log{Refer_event_id:event_id,Refer_participant_id:paticipant_id,
			Refer_team_id:team_id,Change_time:now,Change_value:creditRule.Team_score,Change_type:3,Change_reason:reason}
		credit.AddCreditLog(log)
		team_score += creditRule.Team_score
	}

	team.UpdateTeamCredit(team_id,team_score)

	//获取队友分数
	member := participant.GetMemberCreditByTeamId(team_id,event_id)
	var member_credit []map[string]string
	for _,v := range member {
		user_id := v.User_id
		u := user.GetUserById(user_id)
		var m map[string]string
		m = make(map[string]string)
		m["name"] = u.Name
		m["credit"] = strconv.FormatFloat(v.Credit,'f',-1,64)
		member_credit = append(member_credit, m)
	}


	//返回结果
	var result map[string]interface{}
	result = make(map[string]interface{})
	result["user_score"] = user_score //今日总分
	result["user_credit"] = user_total_credit//累计得分
	result["team_credit"] = team_score //团队累计得分
	if(user_all_right == true) {
		result["user_all_right"] = creditRule.Person_score //单人全部答对额外加分，没有全部答对则传空值
	}
	if(team_allright_flag == true) {
		result["team_all_right"] = creditRule.Team_score //团队全部答对额外加分，没有全部答对则传空值
	}

	result["team_mates"] = member_credit //队友得分[{"name":"A","credit":"1"},{"name":"B","credit":"2"},...]
	result["right_answer"] = correct_answer //正确答案,{"single":{"problem_id":"正确答案的q_id","problem_id":"正确答案的q_id",...}}
	beego.Info("========result======",result)


	this.Data["json"] = result
	this.ServeJSON()
	return


}