package controllers

import (
	"context"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/micro/go-micro"
	proto "service/protoc/eventManage" //proto文件放置路径
	"web/models/event"
)

type EventManageController struct {
	beego.Controller
}

func (this *EventManageController) initEventManage() proto.EventManageService{
	//调用服务
	service := micro.NewService(micro.Name("EventManage.client"))
	service.Init()

	//create new client
	return proto.NewEventManageService("EventManage",service.Client())
}

func (this *EventManageController) EventManageInit() {
	this.TplName = "manage/event_manage.html"
}

func (this *EventManageController) EventManage() {
	offset,_ := this.GetInt32("offset")
	limit,_ := this.GetInt32("limit")
	//获取用户信息
	userSession := this.GetSession("user_id")
	if userSession == nil { //未登陆
		this.Ctx.Redirect(304, "/index")
		return
	}
	userId := userSession.(int)

	//call the userManage method
	eventManage := this.initEventManage()
	req := proto.GetEventListReq{Offset:offset,Limit:limit,ManageId:int64(userId)}
	rsp, err := eventManage.GetEventListByManageIdAndOffst(context.TODO(),&req)
	if err!=nil{
		beego.Info("======ProblemManage=====", rsp.EventList,"-------err--------",err)
	}

	var result map[string]interface{}
	result = make(map[string]interface{})
	result["event_data"] = rsp.EventList
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

	//call the userManage method
	eventManage := this.initEventManage()
	req := proto.AddEventReq{ManageId: int64(manage_id),
		EventTitle:       etitle,
		EventDescription: message,
		EventTime:        string(event_time),
		EventNum:         string(event_num),
		EventType:        ekind,
		ProblemRandom:    pro_random,
		OptionRandom:     opt_random,
		AnswerTime:       answer_time,
		CreditRule:       string(credit_rule),
		ParticipantNum:   int32(participant_num)}
	rsp, err := eventManage.AddNewEvent(context.TODO(),&req)
	if err!=nil{
		beego.Info("-------err--------",err)
	}

	var result map[string]interface{}
	result = make(map[string]interface{})
	if (rsp.EventId != -1) {
		result["result"] = rsp.Message
		result["event_id"] = rsp.EventId
	} else {
		result["result"] = rsp.Message
	}

	this.SetSession("new_event_id", int(rsp.EventId))

	this.Data["json"] = result
	this.ServeJSON()
	return
}
