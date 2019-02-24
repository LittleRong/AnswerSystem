package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"web/models/participant"
	"web/models/team"
	"web/models/user"
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
	offset, _ := this.GetInt("offset")
	limit, _ := this.GetInt("limit")

	user_list := user.GetUserListByOffstAndLimit(offset, limit)

	//user_data,page_num
	beego.Info("======user_list=====", user_list)
	var result map[string]interface{}
	result = make(map[string]interface{})
	result["user_data"] = user_list
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

	for _, value := range team_array {
		beego.Info(value)
		s := value.(map[string]interface{})
		//插入新team
		team_id := team.AddTeam("", event_id)
		if (team_id != -1) {
			//插入participant
			leader_id, _ := strconv.Atoi(s["leader"].(string))
			participant.AddParticipant(leader_id, event_id, team_id, true)

			member_array := s["member"].([]interface{})
			for _, member := range member_array {
				member_id, _ := strconv.Atoi(member.(string))
				participant.AddParticipant(member_id, event_id, team_id, false)
			}
		}
	}

	//待完善，一个插入失败怎么办
	var result map[string]interface{}
	result = make(map[string]interface{})
	result["result"] = "success"
	this.Data["json"] = result
	this.ServeJSON()
	return
}
