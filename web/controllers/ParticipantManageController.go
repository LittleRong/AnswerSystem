package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"context"
	micro "github.com/micro/go-micro"
	userProto "service/protoc/userManage"
	participantProto "service/protoc/participantManage"
	"strconv"
)

type ParticipantManageController struct {
	beego.Controller
}

func (this *ParticipantManageController) ParticipantInsertInit() {
	new_event_id := this.GetSession("new_event_id")
	if new_event_id == nil { //未设置
		this.Ctx.Redirect(302, "/manage/event_insert_init")
		return
	}
	this.TplName = "manage/participant_manage.html"
}

func (this *ParticipantManageController) ParticipantGetUser() {
	offset, _ := this.GetInt32("offset")
	limit, _ := this.GetInt32("limit")
	//获取用户信息
	userSession := this.GetSession("user_id")
	if userSession == nil { //未登陆
		this.Ctx.Redirect(304, "/index")
		return
	}
	userId := userSession.(int)

	//调用服务
	service := micro.NewService(micro.Name("UserManage.client"))
	service.Init()
	//create new client
	userManage := userProto.NewUserManageService("UserManage",service.Client())
	//call the userManage method
	req := userProto.GetUserListReq{Offset:offset,Limit:limit,ManageId:int64(userId)}
	rsp, err := userManage.GetUserListByOffstAndLimit(context.TODO(),&req)
	if err!=nil{
		beego.Info("======ParticipantGetUser=====", rsp.UserList,"-------err--------",err)
	}

	var result map[string]interface{}
	result = make(map[string]interface{})
	result["user_data"] = rsp.UserList
	result["page_num"] = offset
	this.Data["json"] = result
	this.ServeJSON()
	return
}

func (this *ParticipantManageController) EventParticipantInsert() {
	new_event_id := this.GetSession("new_event_id")
	if new_event_id == nil { //未设置
		this.Ctx.Redirect(302, "/manage/event_insert_init")
		return
	}
	event_id := new_event_id.(int)
	team_input := this.Ctx.Request.PostForm.Get("team_data")
	var f interface{}
	_ = json.Unmarshal([]byte(team_input), &f)
	team_array := f.([]interface{})

	//构建参数
	var memberList []*participantProto.ParticipantMember
	for _, value := range team_array {
		var member participantProto.ParticipantMember
		s := value.(map[string]interface{})
		leader_id, _ := strconv.Atoi(s["leader"].(string))
		member.LeaderId = int64(leader_id)

		member_array := s["member"].([]interface{})
		var m []int64
		for _, member := range member_array {
			member_id, _ := strconv.Atoi(member.(string))
			m = append(m,int64(member_id))

		}
		member.MemberId = m
		memberList = append(memberList,&member)

	}

	//调用服务
	service := micro.NewService(micro.Name("ParticipantManage.client"))
	service.Init()
	//create new client
	participantManage := participantProto.NewParticipantManageService("ParticipantManage",service.Client())
	//call the userManage method
	req := participantProto.EPInsertReq{EventId:int64(event_id),ParticipantMemberList:memberList}
	rsp, err := participantManage.EventParticipantInsert(context.TODO(),&req)
	if err!=nil{
		beego.Info("-------err--------",err)
	}

	var result map[string]interface{}
	result = make(map[string]interface{})
	result["result"] = rsp.Message
	this.Data["json"] = result
	this.ServeJSON()
	return
}
