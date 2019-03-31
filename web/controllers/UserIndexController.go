package controllers

import (
	"context"

	"github.com/astaxie/beego"

	participantProto "service/protoc/answerManage"
	eventProto "service/protoc/eventManage"
	userProto "service/protoc/userManage"
	"web/common"
)

type UserIndexController struct {
	beego.Controller
}

func (this *UserIndexController) UserIndexInit() {
	this.TplName = "index/user_index.html"
}

func (this *UserIndexController) UserIndex() {
	var result map[string]interface{}
	result = make(map[string]interface{})
	//获取用户信息
	var user_message *userProto.UserMesssage
	userSession := this.GetSession("user_id")
	var userId int64
	if userSession == nil { //未登陆
		this.Ctx.Redirect(304, "/index")
		return
	} else {
		userId = userSession.(int64)
		//call the userManage method
		userManage := common.InitUserManage()
		req := userProto.GetUserByIdReq{UserId: userId}
		var err error
		user_message, err = userManage.GetUserById(context.TODO(), &req)
		if err == nil {
			beego.Info("-------err--------", err)
		}
	}

	//获取用户参与的事件，并获取事件信息
	var event_message_array []*eventProto.EventShowMesssage
	//call the participantManage method
	participantManage := common.InitParticipantManage()
	req := participantProto.GetPListByUserIdReq{UserId: userId}
	var err error
	rsp, err := participantManage.GetParticipantListByUserId(context.TODO(), &req)

	beego.Info("======UserIndex user_event_list=====", userId, rsp.PEList, "-------err--------", err)

	for _, value := range rsp.PEList {
		//call the participantManage method
		eventManage := common.InitEventManage()
		req := eventProto.EventIdReq{EventId: value.ReferEventId}
		var err error
		rsp, err := eventManage.GetEventByEventId(context.TODO(), &req)
		if err != nil {
			beego.Info("-------err--------", err)
		}
		//增加
		event_message_array = append(event_message_array, rsp)

	}

	//返回
	result["user_message"] = user_message
	result["event_message"] = event_message_array
	this.Data["json"] = result
	this.ServeJSON()
	return

}
