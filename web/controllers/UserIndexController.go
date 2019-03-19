package controllers

import (
	"github.com/astaxie/beego"
	"github.com/micro/go-micro"
	userProto "service/protoc/userManage" //proto文件放置路径
	"context"
	participantProto "service/protoc/answerManage" //proto文件放置路径
	eventProto "service/protoc/eventManage"
)

type UserIndexController struct {
	beego.Controller
}

func (this *UserIndexController) UserIndexInit() {
	this.TplName = "index/user_index.html"
}

func (this *UserIndexController) initUserManage() userProto.UserManageService{
	//调用服务
	service := micro.NewService(micro.Name("UserManage.client"))
	service.Init()

	//create new client
	return userProto.NewUserManageService("UserManage",service.Client())
}

func (this *UserIndexController) initParticipantManage() participantProto.ParticipantManageService{
	//调用服务
	service := micro.NewService(micro.Name("ParticipantManage.client"))
	service.Init()

	//create new client
	return participantProto.NewParticipantManageService("ParticipantManage",service.Client())
}

func (this *UserIndexController) initEventManage() eventProto.EventManageService{
	//调用服务
	service := micro.NewService(micro.Name("EventManage.client"))
	service.Init()

	//create new client
	return eventProto.NewEventManageService("EventManage",service.Client())
}


func (this *UserIndexController) UserIndex() {
	var result map[string]interface{}
	result = make(map[string]interface{})
	//获取用户信息
	var user_message *userProto.UserMesssage
	userSession := this.GetSession("user_id")
	var userId int
	if userSession == nil { //未登陆
		this.Ctx.Redirect(304, "/index")
		return
	} else {
		userId = userSession.(int)
		//call the userManage method
		userManage := this.initUserManage()
		req := userProto.GetUserByIdReq{UserId:int64(userId)}
		var err error
		user_message,err = userManage.GetUserById(context.TODO(),&req)
		if err==nil{
			beego.Info("-------err--------",err)
		}
	}

	//获取用户参与的事件，并获取事件信息
	var event_message_array []*eventProto.EventShowMesssage
	//call the participantManage method
	participantManage := this.initParticipantManage()
	req := participantProto.GetPListByUserIdReq{UserId:int64(userId)}
	var err error
	rsp,err := participantManage.GetParticipantListByUserId(context.TODO(),&req)
	if err!=nil{
		beego.Info("======UserIndex user_event_list=====", rsp.PEList,"-------err--------",err)
	}

	for _, value := range rsp.PEList {
		//call the participantManage method
		eventManage := this.initEventManage()
		req := eventProto.EventIdReq{EventId:value.ReferEventId}
		var err error
		rsp,err := eventManage.GetEventByEventId(context.TODO(),&req)
		if err!=nil{
			beego.Info("-------err--------",err)
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
