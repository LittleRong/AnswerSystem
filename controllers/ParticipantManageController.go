package controllers

import (
	"github.com/astaxie/beego"
	"hello/models/user"
)

type ParticipantManageController struct{
	beego.Controller
}

func (this *ParticipantManageController) ParticipantInsertInit(){
	this.TplName = "manage/participant_manage.html"
}

func (this *ParticipantManageController) ParticipantGetUser() {
	offset,_ := this.GetInt("offset")
	limit,_ := this.GetInt("limit")

	user_list := user.GetUserListByOffstAndLimit(offset,limit)

	//user_data,page_num
	beego.Info("======user_list=====",user_list)
	var result map[string]interface{}
	result = make(map[string]interface{})
	result["user_data"] = user_list
	result["page_num"] = offset
	this.Data["json"] = result
	this.ServeJSON()
	return
}

func (this *ParticipantManageController) EventParticipantInsert() {

}