package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"web/models/event"
	"strconv"
)

type EventManageController struct {
	beego.Controller
}

func (this *EventManageController) EventManageInit() {
	this.TplName = "manage/event_manage.html"
}

func (this *EventManageController) EventManage() {
	////还需要加上偏移
	//offset,_ := this.GetInt("offset")
	//limit,_ := this.GetInt("limit")
	beego.Info("========EventManage======")

	userSession := this.GetSession("user_id")
	if userSession == nil { //未登陆
		this.Ctx.Redirect(304, "/index")
		beego.Info("========未登录======")
		return
	}
	user_id := userSession.(int)
	e := event.GetEventByManageId(user_id)
	var event_data []map[string]string
	for _, v := range e {
		var t map[string]string
		t = make(map[string]string)
		t["event_id"] = strconv.Itoa(v.Event_id)
		t["event_title"] = v.Event_title
		t["event_description"] = v.Event_description
		t["event_type"] = v.Event_type
		event_data = append(event_data, t)
	}

	var result map[string]interface{}
	result = make(map[string]interface{})
	result["event_data"] = event_data
	this.Data["json"] = result
	this.ServeJSON()
	return

}

func (this *EventManageController) EventInsertInit() {
	this.TplName = "manage/event_insert.html"
}

func (this *EventManageController) EventInsert() {
	userSession := this.GetSession("user_id")
	if userSession == nil { //未登陆
		this.Ctx.Redirect(304, "/index")
		beego.Info("========未登录======")
		return
	}
	manage_id := userSession.(int)
	etitle := this.GetString("etitle")
	message := this.GetString("message")
	ekind := this.GetString("ekind")
	pro_random, _ := this.GetBool("pro_random")
	opt_random, _ := this.GetBool("opt_random")
	answer_time := this.GetString("answer_time")
	participant_num, _ := this.GetInt("participant_num")

	//event_time
	start_time := this.GetString("start_time")
	end_time := this.GetString("end_time")
	answer_day := this.GetString("answer_day")
	//answer_day转换成字符串
	etime := event.EventTime{Start_time: start_time, End_time: end_time, Answer_day: answer_day}
	event_time, _ := json.Marshal(etime)

	//event_num
	single, _ := this.GetInt("single")
	multiple, _ := this.GetInt("multiple")
	fill, _ := this.GetInt("fill")
	judge, _ := this.GetInt("judge")
	enum := event.ProblemNum{Single: single, Multiple: multiple, Fill: fill, Judge: judge}
	event_num, _ := json.Marshal(enum)

	//credit_rule
	single_score, _ := this.GetFloat("single_score")
	multiple_score, _ := this.GetFloat("multiple_score")
	fill_score, _ := this.GetFloat("fill_score")
	judge_score, _ := this.GetFloat("judge_score")
	person_score, _ := this.GetFloat("person_score")
	team_score, _ := this.GetFloat("team_score")
	person_score_up, _ := this.GetFloat("person_score_up")
	team_score_up, _ := this.GetFloat("team_score_up")
	crule := event.CreditRule{Single_score: single_score, Multi_score: multiple_score, Fill_score: fill_score, Judge_score: judge_score, Person_score: person_score, Person_score_up: person_score_up, Team_score: team_score, Team_score_up: team_score_up}
	credit_rule, _ := json.Marshal(crule)

	e := event.Event{Manage_id: manage_id,
		Event_title:       etitle,
		Event_description: message,
		Event_time:        string(event_time),
		Event_num:         string(event_num),
		Event_type:        ekind,
		Problem_random:    pro_random,
		Option_random:     opt_random,
		Answer_time:       answer_time,
		Credit_rule:       string(credit_rule),
		Participant_num:   participant_num}

	flag, event_id := event.AddNewEvent(e)

	var result map[string]interface{}
	result = make(map[string]interface{})
	if (flag) {
		result["result"] = "success"
		result["event_id"] = event_id
	} else {
		result["result"] = "faild"
	}

	this.SetSession("new_event_id", event_id)

	this.Data["json"] = result
	this.ServeJSON()
	return
}
