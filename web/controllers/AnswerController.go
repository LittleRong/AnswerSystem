package controllers

import (
	"context"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/micro/go-micro"
	creditProto "service/protoc/answerManage"
	participantProto "service/protoc/answerManage"
	eventProto "service/protoc/eventManage"
	userProto "service/protoc/userManage"
	"strconv"
	"time"
	"web/models/event"
	"web/models/participant"
	"web/models/participant_haved_answer"
	"web/models/union"
)

type AnswerController struct {
	beego.Controller
}

type userAnswerStruct struct {
	problem_id string
	q_id       string
}

type userAnswer struct {
	single []map[string]string
	multi  []map[string]string
	judge  []map[string]string
	fill   []map[string]string
}

func (this *AnswerController) ShowProblemsPage() {
	this.TplName = "answer/user_problem.html"
	event_id, _ := this.GetInt("event_id")
	this.SetSession("event_id", event_id)

}

func (this *AnswerController) initEventManage() eventProto.EventManageService{
	//调用服务
	service := micro.NewService(micro.Name("EventManage.client"))
	service.Init()

	//create new client
	return eventProto.NewEventManageService("EventManage",service.Client())
}

func (this *AnswerController) initUserManage() userProto.UserManageService{
	//调用服务
	service := micro.NewService(micro.Name("UserManage.client"))
	service.Init()

	//create new client
	return userProto.NewUserManageService("UserManage",service.Client())
}


func (this *AnswerController) initParticipantManage() participantProto.ParticipantManageService{
	//调用服务
	service := micro.NewService(micro.Name("ParticipantManage.client"))
	service.Init()

	//create new client
	return participantProto.NewParticipantManageService("ParticipantManage",service.Client())
}

func (this *AnswerController) initCreditManage() creditProto.CreditManageService{
	//调用服务
	service := micro.NewService(micro.Name("CreditManage.client"))
	service.Init()

	//create new client
	return creditProto.NewCreditManageService("CreditManage",service.Client())
}

func (this *AnswerController) GetUserProblems() {
	eventSession := this.GetSession("event_id")
	if eventSession == nil { //未登陆
		this.Ctx.Redirect(304, "/index")
		return
	}
	event_id := eventSession.(int)
	userSession := this.GetSession("user_id")
	if userSession == nil { //未登陆
		this.Ctx.Redirect(304, "/index")
		return
	}
	user_id := userSession.(int)
	p := participant.GetParticipantById(user_id, event_id)
	paticipant_id := p.Participant_id
	team_id := p.Team_id

	//*****************************1.获取用户题目*************************************************
	problemNum := event.GetProblemNumByEventId(event_id)
	problem, buildFlag, answerFlag := union.GetProblemNoAnswer(user_id, event_id, team_id, paticipant_id, problemNum)
	if answerFlag == true {
		var result map[string]interface{}
		result = make(map[string]interface{})
		result["result"] = "已经完成答题，不能再答题了！"
		this.Data["json"] = result
		this.ServeJSON()
		return
	}

	//*****************************2.获取剩余答题时间*************************************************
	var answer_time float64
	event := event.GetEventByEventId(event_id)
	answer_time, _ = strconv.ParseFloat(event.Answer_time, 64)
	if (buildFlag) {
		//获取到生成题目的时间
		participant_time := participant.GetAnswerTimeByParticipantId(paticipant_id)
		//计算出剩余时间
		now_time := time.Now()
		left := now_time.Sub(participant_time)
		if (left.Hours() >= answer_time) { //时间已经耗尽,left.Hours()获取小时格式
			answer_time = 0
		} else {
			answer_time = answer_time - left.Hours()
		}

	}

	//*****************************3.返回给前端*************************************************
	var result map[string]interface{}
	result = make(map[string]interface{})
	result["answer_time"] = answer_time
	result["data"] = problem
	result["event_id"] = event_id
	this.Data["json"] = result
	this.ServeJSON()
	return

}

func (this *AnswerController) GetUserAnswers() {
	eventSession := this.GetSession("event_id")
	if eventSession == nil { //未登陆
		this.Ctx.Redirect(304, "/index")
		return
	}
	event_id := eventSession.(int)
	userSession := this.GetSession("user_id")
	if userSession == nil { //未登陆
		this.Ctx.Redirect(304, "/index")
		return
	}
	user_id := userSession.(int)
	p := participant.GetParticipantById(user_id, event_id)
	paticipant_id := p.Participant_id
	team_id := p.Team_id

	//*****************************1.获取该事件评分标准*************************************************
	eventManage := this.initEventManage()
	eventReq := eventProto.EventIdReq{EventId:int64(event_id)}
	creditRule,creditRuleErr := eventManage.GetCreditRuleByEventId(context.TODO(),&eventReq)
	if creditRuleErr!=nil{
		beego.Info("-------creditRuleErr--------",creditRuleErr)
	}

	//*****************************2.获取正确答案*******************************************************
	var correct_answer map[string]interface{}
	correct_answer = make(map[string]interface{})
	correct_answer = participant.GetCorrectAnswerByParticipantId(paticipant_id)
	beego.Info("========correct_answer======", correct_answer)

	//*****************************3.获取用户输入的答案*******************************************************
	single_input := this.Ctx.Request.PostForm.Get("single")
	var f interface{}
	_ = json.Unmarshal([]byte(single_input), &f)
	single_array := f.([]interface{})

	multi_input := this.Ctx.Request.PostForm.Get("multi")
	_ = json.Unmarshal([]byte(multi_input), &f)
	multi_array := f.([]interface{})

	judge_input := this.Ctx.Request.PostForm.Get("judge")
	_ = json.Unmarshal([]byte(judge_input), &f)
	judge_array := f.([]interface{})

	fill_input := this.Ctx.Request.PostForm.Get("fill")
	_ = json.Unmarshal([]byte(fill_input), &f)
	fill_array := f.([]interface{})

	//*****************************4.计算分数,并将用户答案写入participant_haved_answer表***********************
	single_user_score, single_right_num := JudgeUserInputAnswer(single_array, correct_answer["single"].(map[string]interface{}), paticipant_id, creditRule.SingleScore, 1)
	judge_score, judge_right_num := JudgeUserInputAnswer(judge_array, correct_answer["judge"].(map[string]interface{}), paticipant_id, creditRule.JudgeScore, 3)
	fill_score, fill_right_num := JudgeUserInputAnswer(fill_array, correct_answer["fill"].(map[string]interface{}), paticipant_id, creditRule.FillScore, 0)
	multi_score, multi_right_num := JudgeUserInputAnswer(multi_array, correct_answer["multi"].(map[string]interface{}), paticipant_id, creditRule.MultipleScore, 2)

	user_score := single_user_score + judge_score + fill_score + multi_score
	right_num := single_right_num + judge_right_num + fill_right_num + multi_right_num

	//*****************************5.判断是否全对*****************************************
	user_all_right := false
	problemNum := event.GetProblemNumByEventId(event_id)
	all_num := problemNum.Fill + problemNum.Multiple + problemNum.Single + problemNum.Judge
	if (right_num == all_num) {
		user_score += creditRule.PersonScore
		user_all_right = true
		beego.Info("个人答对额外加分", creditRule.PersonScore)
	}

	//*****************************6.更新积分*****************************************
	//1.更新个人积分
	creditManage := this.initCreditManage()
	pCreditReq := creditProto.UpdatePCreditReq{PaticipantId:int64(paticipant_id),ChangeCredit:user_score}
	pCreditRsp,_ := creditManage.UpdateParticipantCredit(context.TODO(),&pCreditReq)
	user_total_credit := pCreditRsp.Credit
	now_time := time.Now()
	UnixTime := now_time.Unix()
	now := time.Unix(UnixTime, 0).Format("2006-01-02 15:04:05")
	reason := ""
	user_score_log := user_score

	if (user_all_right) {
		reason = "当日全部答对额外加分"
		log := creditProto.CreditLog{EventId: int64(event_id), ParticipantId: int64(paticipant_id),
			TeamId: int64(team_id), ChangeTime: now, ChangeValue: float32(creditRule.PersonScore), ChangeType: 2, ChangeReason: reason}
		creditManage.AddCreditLog(context.TODO(),&log)
		user_score_log = user_score - creditRule.PersonScore
	}
	reason = "答题得分"
	log := creditProto.CreditLog{EventId: int64(event_id), ParticipantId: int64(paticipant_id),
		TeamId: int64(team_id), ChangeTime: now, ChangeValue: float32(user_score_log), ChangeType: 1, ChangeReason: reason}
	creditManage.AddCreditLog(context.TODO(),&log)

	//2.更新组积分
	event := event.GetEventByEventId(event_id)

	team_score := user_score
	//判断是否当日全部答对，若组员全部答对额外加分
	now_date := time.Unix(UnixTime, 0).Format("2006-01-02")
	allRightReq := creditProto.AllRightReq{TeamId:int64(team_id),NowDate:now_date,ParticipantNum:int32(event.Participant_num)}
	allRightRsp,_ := creditManage.WhetherMemberAllRight(context.TODO(),&allRightReq)
	team_allright_flag := allRightRsp.AllRightFlag
	if (team_allright_flag == true) {
		//写积分表
		reason = "当日全组全部答对额外加分"
		log := creditProto.CreditLog{EventId: int64(event_id), ParticipantId: int64(paticipant_id),
			TeamId: int64(team_id), ChangeTime: now, ChangeValue: float32(creditRule.TeamScore), ChangeType: 3, ChangeReason: reason}
		creditManage.AddCreditLog(context.TODO(),&log)
		team_score += creditRule.TeamScore
	}

	teamCreditReq := creditProto.UpdateTeamCreditReq{TeamId:int64(team_id),ChangeCredit:team_score}
	teamCreditRsp,_ := creditManage.UpdateTeamCredit(context.TODO(),&teamCreditReq)
	team_score = teamCreditRsp.Credit

	//*****************************7.获取队友分数*****************************************
	member := participant.GetMemberCreditByTeamId(team_id, event_id)
	var member_credit []map[string]string
	for _, v := range member {
		userId := v.User_id
		userManage := this.initUserManage()
		req := userProto.GetUserByIdReq{UserId:int64(userId)}
		user_message,err := userManage.GetUserById(context.TODO(),&req)
		if err==nil{
			beego.Info("-------err--------",err)
		}
		var m map[string]string
		m = make(map[string]string)
		m["name"] = user_message.Name
		m["credit"] = strconv.FormatFloat(v.Credit, 'f', -1, 64)
		member_credit = append(member_credit, m)
	}

	//*****************************8.返回结果****************************************
	var result map[string]interface{}
	result = make(map[string]interface{})
	result["user_score"] = user_score         //今日总分
	result["user_credit"] = user_total_credit //累计得分
	result["team_credit"] = team_score        //团队累计得分
	if (user_all_right == true) {
		result["user_all_right"] = creditRule.PersonScore //单人全部答对额外加分，没有全部答对则传空值
	}
	if (team_allright_flag == true) {
		result["team_all_right"] = creditRule.TeamScore //团队全部答对额外加分，没有全部答对则传空值
	}
	result["team_mates"] = member_credit    //队友得分[{"name":"A","credit":"1"},{"name":"B","credit":"2"},...]
	result["right_answer"] = correct_answer //正确答案,{"single":{"problem_id":"正确答案的q_id","problem_id":"正确答案的q_id",...}}
	beego.Info("========result======", result)
	this.Data["json"] = result
	this.ServeJSON()
	return
}

//user_score用户答题总分,single_right_num答对的数目
func JudgeUserInputAnswer(input_array []interface{}, correct_answer map[string]interface{}, paticipant_id int, score float64, problem_type int) (float64, int) {
	var right_num int
	var user_score float64
	beego.Info("========input_array======", input_array)
	beego.Info("========correct_answer======", correct_answer)

	if (input_array == nil || correct_answer == nil) {
		return 0, 0
	}

	for _, value := range input_array {
		//判断是否回答正确
		s := value.(map[string]interface{})
		problem_id := s["problem_id"].(string)
		problem_id_int, _ := strconv.Atoi(problem_id)
		user_answer := ""
		right_answer := ""
		true_or_false := false
		if (problem_type == 2) {
			user_answer_i := s["answer"].([]interface{})
			str, _ := json.Marshal(user_answer_i)
			user_answer = string(str)
			right_answer_i := correct_answer[problem_id].([]interface{})
			str, _ = json.Marshal(right_answer_i)
			right_answer = string(str)
		} else {
			if (s["answer"] != nil && s["answer"] != "") {
				user_answer = s["answer"].(string)
			}
			right_answer = correct_answer[problem_id].(string)
		}

		if (user_answer == right_answer) {
			right_num++
			true_or_false = true
		}
		beego.Info("problem_id=", problem_id, "user_answer=", user_answer, " right_answer=", right_answer)

		//将用户答案写入participant_haved_answer表
		participant_haved_answer.UpdateUserAnswer(paticipant_id, problem_id_int, user_answer, true_or_false)
	}
	user_score = user_score + float64(right_num)*score
	beego.Info("答对题目类型", problem_type, "：", right_num, "  ,每题：", score, "分 ,总分:", user_score)
	return user_score, right_num
}
