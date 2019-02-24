package controllers

import (
	"github.com/astaxie/beego"
	"web/models/user"
)

type UserManageController struct {
	beego.Controller
}

func (this *UserManageController) UserManageInit() {
	this.TplName = "manage/user_manage.html"
}

func (this *UserManageController) UserManage() {
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

func (this *UserManageController) ChangeUser() {
	change_id, _ := this.GetInt("change_id")
	user_name := this.GetString("user_name")
	login_name := this.GetString("login_name")
	user_phone_number := this.GetString("user_phone_number")
	user_job_number := this.GetString("user_job_number")
	user_gender, _ := this.GetInt("user_gender")

	r := user.UpdateUserById(change_id, user_name, login_name, user_phone_number, user_job_number, user_gender)
	var result map[string]interface{}
	result = make(map[string]interface{})
	result["result"] = r
	this.Data["json"] = result
	this.ServeJSON()
	return
}

func (this *UserManageController) AddUser() {
	user_name := this.GetString("user_name")
	login_name := this.GetString("login_name")
	user_phone_number := this.GetString("user_phone_number")
	user_job_number := this.GetString("user_job_number")
	user_gender, _ := this.GetInt("user_gender")

	r := user.AddUser(user_name, login_name, user_phone_number, user_job_number, user_gender)
	var result map[string]interface{}
	result = make(map[string]interface{})
	result["result"] = r
	this.Data["json"] = result
	this.ServeJSON()
	return

}

func (this *UserManageController) DeleteUser() {
	delete_id, _ := this.GetInt("delete_id")
	r := user.DeleteUserById(delete_id)
	var result map[string]interface{}
	result = make(map[string]interface{})
	result["result"] = r
	this.Data["json"] = result
	this.ServeJSON()
	return
}
