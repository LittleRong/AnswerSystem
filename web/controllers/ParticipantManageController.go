package controllers

import (
	"encoding/json"
	"strconv"

	"github.com/astaxie/beego"

	participantProto "service/protoc/answerManage"
	userProto "service/protoc/userManage"
	"web/common"
)

type ParticipantManageController struct {
	beego.Controller
}

// @Title 获得添加参赛者页面
// @Description 获得添加参赛者页面
// @Success 200 {}
// @router / [get]
func (this *ParticipantManageController) ParticipantInsertInit() {
	new_event_id := this.GetSession("new_event_id")
	if new_event_id == nil { //未设置
		this.Ctx.Redirect(302, "/manage/event_insert_init")
		return
	}
	this.TplName = "manage/participant_manage.html"
}
// @Title 获取参赛者信息列表
// @Description 获取参赛者信息列表
// @Success 200 {}
// @Param   offset   query   string  true       "页码"
// @Param   limit query   string  true       "一页展示数量"
// @router /all [get]
func (this *ParticipantManageController) ParticipantGetUser() {
	offset, _ := this.GetInt32("offset")
	limit, _ := this.GetInt32("limit")
	//获取用户信息
	userSession := this.GetSession("user_id")
	if userSession == nil { //未登陆
		this.Ctx.Redirect(304, "/index")
		return
	}
	userId := userSession.(int64)

	//调用服务
	userManage, ctx := common.InitUserManage(this.CruSession)
	req := userProto.GetUserListReq{Offset: offset, Limit: limit, ManageId: userId}
	rsp, err := userManage.GetUserListByOffstAndLimit(ctx, &req)
	beego.Info("======ParticipantGetUser=====", rsp.UserList)
	if err != nil {
		beego.Info("-------err--------", err)
	}

	var result map[string]interface{}
	result = make(map[string]interface{})
	result["user_data"] = rsp.UserList
	result["page_num"] = offset
	this.Data["json"] = result
	this.ServeJSON()
	return
}

// @Title 批量新增参赛者
// @Description 批量新增参赛者
// @Param	team_data	formData	string	false	"新增参赛者信息"
// @Success 200 {string} result
// @router /batch [post]
func (this *ParticipantManageController) EventParticipantInsert() {
	new_event_id := this.GetSession("new_event_id")
	if new_event_id == nil { //未设置
		this.Ctx.Redirect(302, "/manage/event_insert_init")
		return
	}
	event_id := new_event_id.(int)
	team_input := this.Ctx.Request.PostForm.Get("team_data")
	beego.Info("-----------EventParticipantInsert-----------",team_input)
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
			m = append(m, int64(member_id))

		}
		member.MemberId = m
		memberList = append(memberList, &member)

	}

	//调用服务
	participantManage, ctx := common.InitParticipantManage(this.CruSession)
	req := participantProto.EPInsertReq{EventId: int64(event_id), ParticipantMemberList: memberList}
	rsp, err := participantManage.EventParticipantInsert(ctx, &req)
	if err != nil {
		beego.Info("-------err--------", err)
	}

	var result map[string]interface{}
	result = make(map[string]interface{})
	result["result"] = rsp.Message
	this.Data["json"] = result
	this.ServeJSON()
	return
}
