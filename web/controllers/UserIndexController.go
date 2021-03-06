package controllers

import (
	"github.com/astaxie/beego"

	participantProto "service/protoc/answerManage"
	eventProto "service/protoc/eventManage"
	userProto "service/protoc/userManage"
	"web/common"
)

type UserIndexController struct {
	beego.Controller
}

// @Title 获得用户信息首页
// @Description 获得用户信息首页
// @Success 200 {}
// @router / [get]
func (this *UserIndexController) UserIndexInit() {
	this.TplName = "index/user_index.html"
}

// @Title 获得用户信息
// @Description 获得用户信息
// @Success 200 {}
// @router /user_event [get]
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
		userManage, ctx := common.InitUserManage(this.CruSession)
		req := userProto.GetUserByIdReq{UserId: userId}
		var err error
		user_message, err = userManage.GetUserById(ctx, &req)
		if err == nil {
			beego.Info("-------err--------", err)
		}
	}

	//获取用户参与的事件，并获取事件信息
	var event_message_array []*eventProto.EventShowMesssage
	//call the participantManage method
	participantManage, ctx := common.InitParticipantManage(this.CruSession)
	req := participantProto.GetPListByUserIdReq{UserId: userId}
	var err error
	rsp, err := participantManage.GetParticipantListByUserId(ctx, &req)
	if (err != nil) {
		beego.Info("-------err--------", err)
	}

	for _, value := range rsp.PEList {
		//call the participantManage method
		eventManage, ctx := common.InitEventManage(this.CruSession)
		req := eventProto.EventIdReq{EventId: value.ReferEventId}
		var err error
		rsp, err := eventManage.GetEventByEventId(ctx, &req)
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
