package controllers

import (
	"encoding/json"
	"github.com/spf13/viper"
	"strconv"
	"time"

	"github.com/astaxie/beego"

	creditProto "service/protoc/answerManage"
	participantProto "service/protoc/answerManage"
	eventProto "service/protoc/eventManage"
	unionProto "service/protoc/unionManage"
	userProto "service/protoc/userManage"
	"web/common"
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
	user_id := userSession.(int64)
	pManage,ctx := common.InitParticipantManage(this.CruSession)
	pReq := participantProto.PUserEventIdReq{EventId: int64(event_id), UserId: user_id}
	p, pErr := pManage.GetParticipantByUserAndEvent(ctx, &pReq)
	if pErr != nil {
		beego.Info("-------pErr--------", pErr)
	}
	paticipant_id := p.ParticipantId
	team_id := p.TeamId

	//*****************************1.获取用户题目*************************************************
	eventManage,ctx := common.InitEventManage(this.CruSession)
	req := eventProto.EventIdReq{EventId: int64(event_id)}
	problemNum, problemNumErr := eventManage.GetProblemNumByEventId(ctx, &req)
	if problemNumErr != nil {
		beego.Info("-------err--------", problemNumErr)
	}

	uniontManage,ctx := common.InitUniontManage(this.CruSession)
	problemNumsReq := unionProto.ProblemNum{Single: problemNum.Single, Multiple: problemNum.Multiple, Fill: problemNum.Fill, Judge: problemNum.Judge}
	unionReq := unionProto.GetProblemNoAnswerReq{EventId: int64(event_id), UserId: int64(user_id), TeamId: int64(team_id), PaticipantId: paticipant_id, ProblemNum: &problemNumsReq}
	unionRsp, unionErr := uniontManage.GetProblemNoAnswer(ctx, &unionReq)
	if unionErr != nil {
		beego.Info("-------unionErr--------", unionErr)
	}

	buildFlag := unionRsp.BuildFlag
	answerFlag := unionRsp.AnswerFlag
	var problem map[string]interface{}
	problem = make(map[string]interface{})
	problem["single"] = unionRsp.Single
	problem["multiple"] = unionRsp.Multiple
	problem["fill"] = unionRsp.Fill
	problem["judge"] = unionRsp.Judge

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
	var err error
	event_message, err := eventManage.GetEventByEventId(ctx, &req)
	if err != nil {
		beego.Info("-------err--------", err)
	}

	answer_time = event_message.AnswerTime
	if (buildFlag) {
		//获取到生成题目的时间
		pAnswerTimeReq := participantProto.ParticipantIdReq{ParticipantId: paticipant_id}
		pAnswerTimeRsp, pAnswerTimeErr := pManage.GetAnswerTimeByParticipantId(ctx, &pAnswerTimeReq)
		if pAnswerTimeErr != nil {
			beego.Info("-------pErr--------", pAnswerTimeErr)
		}

		timeLayout := "2006-01-02 15:04:05"  //转化所需模板，go默认时间
		loc, _ := time.LoadLocation("Local") //获取本地时区
		participant_time, _ := time.ParseInLocation(timeLayout, pAnswerTimeRsp.AnswerTime, loc)

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
	user_id := userSession.(int64)
	pManage,ctx := common.InitParticipantManage(this.CruSession)
	pReq := participantProto.PUserEventIdReq{EventId: int64(event_id), UserId: user_id}
	p, pErr := pManage.GetParticipantByUserAndEvent(ctx, &pReq)
	if pErr != nil {
		beego.Info("-------pErr--------", pErr)
	}
	paticipant_id := p.ParticipantId
	team_id := p.TeamId

	//*****************************1.获取该事件评分标准*************************************************
	eventManage,ctx := common.InitEventManage(this.CruSession)
	eventReq := eventProto.EventIdReq{EventId: int64(event_id)}
	creditRule, creditRuleErr := eventManage.GetCreditRuleByEventId(ctx, &eventReq)
	if creditRuleErr != nil {
		beego.Info("-------creditRuleErr--------", creditRuleErr)
	}

	//*****************************2.获取正确答案*******************************************************
	pIdReq := participantProto.ParticipantIdReq{ParticipantId: int64(paticipant_id)}
	correct_answer, pIdErr := pManage.GetCorrectAnswerByParticipantId(ctx, &pIdReq)
	if pIdErr != nil {
		beego.Info("-------pIdErr--------", pIdErr)
	}
	beego.Info("========correct_answer11111======", correct_answer)

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
	single_user_score, single_right_num, singleFront := this.JudgeUserInputAnswer(single_array, correct_answer.SingleAnswerList, paticipant_id, creditRule.SingleScore, viper.GetInt("enum.problemType.singleType"))
	judge_score, judge_right_num, judgeFront := this.JudgeUserInputAnswer(judge_array, correct_answer.JudgeAnswerList, paticipant_id, creditRule.JudgeScore, viper.GetInt("enum.problemType.judgeType"))
	fill_score, fill_right_num, fillFront := this.JudgeUserInputAnswer(fill_array, correct_answer.FillAnswerList, paticipant_id, creditRule.FillScore, viper.GetInt("enum.problemType.fillType"))
	multi_score, multi_right_num, multiFront := this.JudgeUserMultiInputAnswer(multi_array, correct_answer.MultiAnswerList, paticipant_id, creditRule.MultipleScore, viper.GetInt("enum.problemType.multipleType"))

	user_score := single_user_score + judge_score + fill_score + multi_score
	right_num := single_right_num + judge_right_num + fill_right_num + multi_right_num

	//*****************************5.判断是否全对*****************************************
	user_all_right := false
	problemNum, problemNumErr := eventManage.GetProblemNumByEventId(ctx, &eventReq)
	if problemNumErr != nil {
		beego.Info("-------problemNumErr--------", problemNumErr)
	}

	all_num := int(problemNum.Fill + problemNum.Multiple + problemNum.Single + problemNum.Judge)
	if (right_num == all_num) {
		user_score += creditRule.PersonScore
		user_all_right = true
		beego.Info("个人答对额外加分", creditRule.PersonScore)
	}

	//*****************************6.更新积分*****************************************
	//1.更新个人积分
	creditManage,ctx := common.InitCreditManage(this.CruSession)
	pCreditReq := creditProto.UpdatePCreditReq{PaticipantId: int64(paticipant_id), ChangeCredit: user_score}
	pCreditRsp, _ := creditManage.UpdateParticipantCredit(ctx, &pCreditReq)
	user_total_credit := pCreditRsp.Credit
	now_time := time.Now()
	UnixTime := now_time.Unix()
	now := time.Unix(UnixTime, 0).Format("2006-01-02 15:04:05")
	reason := ""
	user_score_log := user_score

	if (user_all_right) {
		reason = "当日全部答对额外加分"
		log := creditProto.CreditLog{EventId: int64(event_id), ParticipantId: int64(paticipant_id),
			TeamId: int64(team_id), ChangeTime: now, ChangeValue: float32(creditRule.PersonScore), ChangeType: viper.GetInt32("enum.creditChangeType.personAllRight"), ChangeReason: reason}
		creditManage.AddCreditLog(ctx, &log)
		user_score_log = user_score - creditRule.PersonScore
	}
	reason = "答题得分"
	log := creditProto.CreditLog{EventId: int64(event_id), ParticipantId: int64(paticipant_id),
		TeamId: int64(team_id), ChangeTime: now, ChangeValue: float32(user_score_log), ChangeType: viper.GetInt32("enum.creditChangeType.dailyAnswer"), ChangeReason: reason}
	creditManage.AddCreditLog(ctx, &log)

	//2.更新组积分
	event, err := eventManage.GetEventByEventId(ctx, &eventReq)
	if err != nil {
		beego.Info("-------err--------", err)
	}

	team_score := user_score
	//判断是否当日全部答对，若组员全部答对额外加分
	now_date := time.Unix(UnixTime, 0).Format("2006-01-02")
	allRightReq := creditProto.AllRightReq{TeamId: int64(team_id), NowDate: now_date, ParticipantNum: int32(event.ParticipantNum)}
	allRightRsp, _ := creditManage.WhetherMemberAllRight(ctx, &allRightReq)
	team_allright_flag := allRightRsp.AllRightFlag
	if (team_allright_flag == true) {
		//写积分表
		reason = "当日全组全部答对额外加分"
		log := creditProto.CreditLog{EventId: int64(event_id), ParticipantId: int64(paticipant_id),
			TeamId: int64(team_id), ChangeTime: now, ChangeValue: float32(creditRule.TeamScore), ChangeType: viper.GetInt32("enum.creditChangeType.teamAllRight"), ChangeReason: reason}
		creditManage.AddCreditLog(ctx, &log)
		team_score += creditRule.TeamScore
	}

	teamCreditReq := creditProto.UpdateTeamCreditReq{TeamId: int64(team_id), ChangeCredit: team_score}
	teamCreditRsp, _ := creditManage.UpdateTeamCredit(ctx, &teamCreditReq)
	team_score = teamCreditRsp.Credit

	//*****************************7.获取队友分数*****************************************
	memberReq := participantProto.PTeamEventIdReq{EventId: int64(event_id), TeamId: int64(team_id)}
	memberRsp, memberErr := pManage.GetMemberCreditByTeamId(ctx, &memberReq)
	if memberErr != nil {
		beego.Info("-------memberErr--------", memberErr)
	}
	member := memberRsp.PEList

	var member_credit []map[string]string
	for _, v := range member {
		userId := v.UserId
		userManage,ctx := common.InitUserManage(this.CruSession)
		req := userProto.GetUserByIdReq{UserId: int64(userId)}
		user_message, err := userManage.GetUserById(ctx, &req)
		if err == nil {
			beego.Info("-------err--------", err)
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
	result["team_mates"] = member_credit //队友得分[{"name":"A","credit":"1"},{"name":"B","credit":"2"},...]
	var frontAnswer map[string]interface{}
	frontAnswer = make(map[string]interface{})
	frontAnswer["single"] = singleFront
	frontAnswer["judge"] = judgeFront
	frontAnswer["fill"] = fillFront
	frontAnswer["multi"] = multiFront
	result["right_answer"] = frontAnswer //正确答案,{"single":{"problem_id":"正确答案的q_id","problem_id":"正确答案的q_id",...},"multi":{"problem_id":[正确答案的q_id1,正确答案的q_id2]}}
	beego.Info("========result======", result)
	this.Data["json"] = result
	this.ServeJSON()
	return
}

//user_score用户答题总分,single_right_num答对的数目
func (this *AnswerController)JudgeUserInputAnswer(input_array []interface{}, correct_answer []*participantProto.NolAnswer, paticipant_id int64, score float64, problem_type int) (float64, int, map[string]string) {
	var right_num int
	var user_score float64
	var frontAnswer map[string]string
	frontAnswer = make(map[string]string)
	beego.Info("========input_array======", input_array)
	beego.Info("========correct_answer======", correct_answer)

	if (input_array == nil || correct_answer == nil) {
		return 0, 0, nil
	}

	var answer_arr map[string]string
	answer_arr = make(map[string]string)
	for _, v := range correct_answer {
		answer_arr[v.ProblemId] = v.Answer
	}

	for _, value := range input_array {
		//判断是否回答正确
		s := value.(map[string]interface{})
		problem_id := s["problem_id"].(string)
		user_answer := ""
		right_answer := ""
		true_or_false := false
		if (s["answer"] != nil && s["answer"] != "") {
			user_answer = s["answer"].(string)
		}
		right_answer = answer_arr[problem_id]

		if (user_answer == right_answer) {
			right_num++
			true_or_false = true
		}
		beego.Info("problem_id=", problem_id, "user_answer=", user_answer, " right_answer=", right_answer)

		//将用户答案写入participant_haved_answer表
		pManage,ctx := common.InitParticipantManage(this.CruSession)
		problemIdStr, _ := strconv.ParseInt(problem_id, 10, 64)
		pReq := participantProto.UpdateUserAnswerReq{ParticipantId: paticipant_id, ProblemId: problemIdStr, UserAnswer: user_answer, TrueOrFalse: true_or_false}
		_, pErr := pManage.UpdateUserAnswer(ctx, &pReq)
		if pErr != nil {
			beego.Info("-------pErr--------", pErr)
		}

		//前端显示
		frontAnswer[problem_id] = answer_arr[problem_id]
	}

	user_score = user_score + float64(right_num)*score
	beego.Info("答对题目类型", problem_type, "：", right_num, "  ,每题：", score, "分 ,总分:", user_score)
	return user_score, right_num, frontAnswer
}

//user_score用户答题总分,single_right_num答对的数目
func (this *AnswerController)JudgeUserMultiInputAnswer(input_array []interface{}, correct_answer []*participantProto.MultiAnswer, paticipant_id int64, score float64, problem_type int) (float64, int, map[string][]float64) {
	var right_num int
	var user_score float64
	var frontAnswer map[string][]float64
	frontAnswer = make(map[string][]float64)
	beego.Info("========input_array======", input_array)
	beego.Info("========correct_answer======", correct_answer)

	if (input_array == nil || correct_answer == nil) {
		return 0, 0, nil
	}

	var answer_arr map[string][]float64
	answer_arr = make(map[string][]float64)
	for _, v := range correct_answer {
		answer_arr[v.ProblemId] = v.Answer
	}
	for _, value := range input_array {
		//判断是否回答正确
		s := value.(map[string]interface{})
		problem_id := s["problem_id"].(string)
		user_answer := ""
		right_answer := ""
		true_or_false := false
		if (problem_type == viper.GetInt("enum.problemType.multipleType")) {
			user_answer_i := s["answer"].([]interface{})
			str, _ := json.Marshal(user_answer_i)
			user_answer = string(str)

			right_answer_i := answer_arr[problem_id]
			str, _ = json.Marshal(right_answer_i)
			right_answer = string(str)
		}

		if (user_answer == right_answer) {
			right_num++
			true_or_false = true
		}
		beego.Info("problem_id=", problem_id, "user_answer=", user_answer, " right_answer=", right_answer)

		//将用户答案写入participant_haved_answer表
		pManage,ctx := common.InitParticipantManage(this.CruSession)
		problemIdStr, _ := strconv.ParseInt(problem_id, 10, 64)
		pReq := participantProto.UpdateUserAnswerReq{ParticipantId: paticipant_id, ProblemId: problemIdStr, UserAnswer: user_answer, TrueOrFalse: true_or_false}
		_, pErr := pManage.UpdateUserAnswer(ctx, &pReq)
		if pErr != nil {
			beego.Info("-------pErr--------", pErr)
		}

		//前端显示
		frontAnswer[problem_id] = answer_arr[problem_id]
	}

	user_score = user_score + float64(right_num)*score
	beego.Info("答对题目类型", problem_type, "：", right_num, "  ,每题：", score, "分 ,总分:", user_score)
	return user_score, right_num, frontAnswer
}
