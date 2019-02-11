package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"hello/models/event"
	"hello/models/participant"
	"hello/models/union"
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
	now_time:= time.Now()
	UnixTime:=now_time.Unix()
	now := time.Unix(UnixTime, 0).Format("2006-01-02" ) //设置时间戳 使用模板格式化为日期字符串
	paticipant_id  := 1

	beego.Info("========now======",now)

	//获取用户题目
	problem,buildFlag := union.GetProblemNoAnswer(user_id,event_id,now)

	//获取答题时间
	var answer_time float64
	event := event.GetEventByEventId(event_id)
	answer_time,_ = strconv.ParseFloat(event.Answer_time,64)

	if (buildFlag) {
		//获取到生成题目的时间
		participant_time :=participant.GetAnswerTimeByParticipantId(paticipant_id)
		//计算出剩余时间
		left := now_time.Sub(participant_time)
		if(left.Hours() >= answer_time){//时间已经耗尽,left.Hours()获取小时格式
			answer_time = 0
		}else {
			answer_time = answer_time-left.Hours()
		}

	}
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
	//获取该事件评分标准
	event_id := 1
	var creditRule event.CreditRule
	creditRule = event.GetCreditRuleByEventId(event_id)
	beego.Info("========creditRule======",creditRule)

	//获取正确答案
	paticipant_id  := 1
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

	//计算分数
	var user_score float64
	single_right_num := 0
	all_num := 0 //全部题目数
	single_correct_answer := correct_answer["single"].(map[string]interface{})
	for _,value := range single_array{
		s := value.(map[string]interface {})
		problem_id := s["problem_id"].(string)
		answer := s["q_id"]
		right_answer,ok := single_correct_answer[problem_id].(string)
		beego.Info("problem_id=",problem_id,"answer=",answer," right_answer=",right_answer)
		if (ok && answer == right_answer){
			single_right_num++
		}
		all_num++
	}
	user_score = user_score + float64(single_right_num)*creditRule.Single_score
	beego.Info("========single_score======",user_score)


	//判断是否全对，user_all_right
	user_all_right := false
	if(single_right_num == all_num) {
		user_score += creditRule.Person_score
		user_all_right = true
	}

	beego.Info("========user_score======",user_score)
	var result map[string]interface{}
	result = make(map[string]interface{})
	result["user_score"] = user_score //今日总分
	result["user_credit"] = "1" //累计得分
	result["team_credit"] = "1" //团队累计得分
	if(user_all_right == true) {
		result["user_all_right"] = creditRule.Person_score //单人全部答对额外加分，没有全部答对则传空值
	}

	result["team_all_right"] = "1" //团队全部答对额外加分，没有全部答对则传空值
	result["team_mates"] = "1" //队友得分[{"name":"A","credit":"1"},{"name":"B","credit":"2"},...]
	result["right_answer"] = correct_answer //正确答案,{"single":{"problem_id":"正确答案的q_id","problem_id":"正确答案的q_id",...}}
	beego.Info("========result======",result)


	this.Data["json"] = result
	this.ServeJSON()
	return


}