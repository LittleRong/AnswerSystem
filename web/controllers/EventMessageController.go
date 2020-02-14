package controllers

import (
	"github.com/astaxie/beego"

	creditProto "service/protoc/answerManage"
	participantProto "service/protoc/answerManage"
	eventProto "service/protoc/eventManage"
	"web/common"
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
	user_id := userSession.(int64)
	pManage,ctx := common.InitParticipantManage(this.CruSession)
	pReq := participantProto.PUserEventIdReq{EventId: int64(event_id), UserId: user_id}
	participant, pErr := pManage.GetParticipantByUserAndEvent(ctx, &pReq)
	if pErr != nil {
		beego.Info("-------pErr--------", pErr)
	}
	team_id := participant.TeamId

	//*****************************1.获取事件信息event_message*************************************************
	eventManage,ctx := common.InitEventManage(this.CruSession)
	req := eventProto.EventIdReq{EventId: int64(event_id)}
	var err error
	event_message, err := eventManage.GetDetailEventByEventId(ctx, &req)
	if err != nil {
		beego.Info("-------err--------", err)
	}

	//*****************************2.获取积分信息credit_message************************************************
	creditManage,ctx := common.InitCreditManage(this.CruSession)
	userCreditReq := creditProto.UserEventIdReq{EventId: int64(event_id), UserId: int64(user_id)}
	personCredit, personErr := creditManage.GetPersonCredit(ctx, &userCreditReq)
	if personErr != nil {
		beego.Info("-------personErr--------", personErr)
	}

	teamCreditReq := creditProto.TeamEventIdReq{EventId: int64(event_id), TeamId: int64(team_id)}
	teamCredit, teamErr := creditManage.GetTeamCredit(ctx, &teamCreditReq)
	if teamErr != nil {
		beego.Info("-------teamErr--------", teamErr)
	}

	var credit_message map[string]interface{}
	credit_message = make(map[string]interface{})
	credit_message["person_credit"] = personCredit.Credit
	credit_message["team_credit"] = teamCredit.Credit

	//*****************************3.获取积分日志************************************************
	creditLogReq := creditProto.TeamIdReq{TeamId: int64(team_id)}
	creditLogCredit, creditLogErr := creditManage.GetCreditLogByTeamId(ctx, &creditLogReq)
	if creditLogErr != nil {
		beego.Info("-------creditLogErr--------", creditLogErr)
	}
	credit_message["detail_credit"] = creditLogCredit.CreditLogList

	//*****************************4.返回结果******************************************************************
	var result map[string]interface{}
	result = make(map[string]interface{})
	result["event_message"] = event_message
	result["credit_message"] = credit_message
	this.Data["json"] = result
	this.ServeJSON()
	return
}
