package controllers

import (
	"github.com/astaxie/beego"
	"web/models/user"
	"context"
	micro "github.com/micro/go-micro"
	proto "service/protoc" //proto文件放置路径
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
	//获取用户信息
	userSession := this.GetSession("user_id")
	if userSession == nil { //未登陆
		this.Ctx.Redirect(304, "/index")
		return
	}
	userId := userSession.(int)

	//user_list := user.GetUserListByOffstAndLimit(offset, limit)
	//调用服务
	service := micro.NewService(micro.Name("UserManage.client"))
	service.Init()

	//create new client
	userManage := proto.NewUserManageService("UserManage",service.Client())

	//call the userManage method
	req := proto.GetUserListReq{Offset:int32(offset),Limit:int32(limit),ManageId:int64(userId)}
	user_list, err := userManage.GetUserListByOffstAndLimit(context.TODO(),&req)

	//user_data,page_num
	beego.Info("======user_list=====", user_list.UserList,"-----------------",err)
	var result map[string]interface{}
	result = make(map[string]interface{})
	result["user_data"] = user_list.UserList
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
