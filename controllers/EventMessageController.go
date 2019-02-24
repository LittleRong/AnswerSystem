package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"hello/models/credit"
	"hello/models/event"
	"hello/models/participant"
	"hello/models/team"
	"strconv"
)

type EventMessageController struct {
	beego.Controller
}

func (this *EventMessageController) EventMessageInit() {
	this.TplName = "answer/event_message.html"
}

func (this *EventMessageController) GetEventMessage() {
	event_id, _ := this.GetInt("event_id")
	userSession := this.GetSession("user_id")
	if userSession == nil { //未登陆
		this.Ctx.Redirect(304, "/index")
		return
	}
	user_id := userSession.(int)
	p := participant.GetParticipantById(user_id, event_id)
	team_id := p.Team_id
	t := team.GetTeamById(team_id, event_id)

	//*****************************1.获取事件信息event_message*************************************************
	var event_message map[string]interface{}
	event_message = make(map[string]interface{})
	e := event.GetEventByEventId(event_id)
	event_message["event_title"] = e.Event_title
	event_message["event_type"] = e.Event_type
	event_message["event_description"] = e.Event_description
	//时间
	var event_time event.EventTime
	json.Unmarshal([]byte(e.Event_time), &event_time)
	event_message["start_time"] = event_time.Start_time
	event_message["end_time"] = event_time.End_time
	event_message["answer_day"] = event_time.Answer_day
	//积分规则
	var credit_rule event.CreditRule
	json.Unmarshal([]byte(e.Credit_rule), &credit_rule)
	event_message["single_score"] = credit_rule.Single_score
	event_message["multiple_score"] = credit_rule.Multi_score
	event_message["judge_score"] = credit_rule.Judge_score
	event_message["fill_score"] = credit_rule.Fill_score
	event_message["team_score"] = credit_rule.Team_score
	event_message["team_score_up"] = credit_rule.Team_score_up
	event_message["person_score"] = credit_rule.Person_score
	event_message["person_score_up"] = credit_rule.Person_score_up
	//数量
	var problem_num event.ProblemNum
	json.Unmarshal([]byte(e.Event_num), &problem_num)
	event_message["single"] = problem_num.Single
	event_message["multiple"] = problem_num.Multiple
	event_message["judge"] = problem_num.Judge
	event_message["fill"] = problem_num.Fill
	beego.Info("========event_message======", event_message)

	//*****************************2.获取积分信息credit_message************************************************
	var credit_message map[string]interface{}
	credit_message = make(map[string]interface{})
	credit_message["person_credit"] = p.Credit
	credit_message["team_credit"] = t.Team_credit
	credit_log := credit.GetCreditLogByTeamId(team_id)
	var detail_credit []map[string]interface{}
	for _, v := range credit_log {
		var d map[string]interface{}
		d = make(map[string]interface{})
		d["team_id"] = v.Refer_team_id
		change_value := strconv.FormatFloat(v.Change_value, 'f', -1, 64)
		d["change_reason"] = v.Change_reason + "," + change_value + "分"
		d["change_time"] = v.Change_time
		detail_credit = append(detail_credit, d)
	}
	credit_message["detail_credit"] = detail_credit

	//*****************************3.返回结果******************************************************************
	var result map[string]interface{}
	result = make(map[string]interface{})
	result["event_message"] = event_message
	result["credit_message"] = credit_message
	this.Data["json"] = result
	this.ServeJSON()
	return
}
