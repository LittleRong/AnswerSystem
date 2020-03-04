package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	proto "service/protoc/userManage"
	"strconv"
	"web/common"
)

type UserManageController struct {
	beego.Controller
}

// @Title 获得用户管理页面
// @Description 获得用户管理页面
// @Success 200 {}
// @router / [get]
func (this *UserManageController) UserManageInit() {
	this.TplName = "manage/user_manage.html"
}

// @Title 获取用户列表
// @Description 获取用户列表
// @Success 200 {}
// @Param   offset   query   string  true       "页码"
// @Param   limit query   string  true       "一页展示数量"
// @router /all [get]
func (this *UserManageController) GetUser() {
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
	//call the userManage method
	req := proto.GetUserListReq{Offset: offset, Limit: limit, ManageId: int64(userId)}
	rsp, err := userManage.GetUserListByOffstAndLimit(ctx, &req)

	//user_data,page_num
	beego.Info("======user_list=====", rsp.UserList, "-----------------", err)
	var result map[string]interface{}
	result = make(map[string]interface{})
	result["user_data"] = rsp.UserList
	result["page_num"] = offset
	this.Data["json"] = result
	this.ServeJSON()
	return
}

// @Title 修改用户信息
// @Description 修改用户信息
// @Param	user_name	formData	string	false	"用户名"
// @Param	login_name	formData	string	false	"用户登陆名"
// @Param	user_phone_number	formData	string	false	"用户手机号码"
// @Param	user_job_number	formData	string	false	"用户工号"
// @Param	user_gender	formData	int	false	"用户性别"
// @Success 200 {string} success
// @Failure 400 user doesn't exit
// @Failure 500 server's wrong
// @router / [put]
func (this *UserManageController) ChangeUser() {
	changeId, _ := this.GetInt64("change_id")
	userName := this.GetString("user_name")
	loginName := this.GetString("login_name")
	userPhoneNumber := this.GetString("user_phone_number")
	userJobNumber := this.GetString("user_job_number")
	userGender, _ := this.GetInt32("user_gender")

	//call the userManage method
	this.StartSession()
	userManage, ctx := common.InitUserManage(this.CruSession)
	req := proto.ChangeUserReq{ChangeId: changeId, Name: userName, LoginName: loginName, PhoneNumber: userPhoneNumber, JobNumber: userJobNumber, Gender: userGender}
	rsp, err := userManage.UpdateUserById(ctx, &req)
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

// @Title 新增用户
// @Description 新增用户
// @Param	user_name	formData	string	false	"用户名"
// @Param	login_name	formData	string	false	"用户登陆名"
// @Param	user_phone_number	formData	string	false	"用户手机号码"
// @Param	user_job_number	formData	string	false	"用户工号"
// @Param	user_gender	formData	int	false	"用户性别"
// @Success 200 {string} result
// @router / [post]
func (this *UserManageController) AddUser() {
	userName := this.GetString("user_name")
	loginName := this.GetString("login_name")
	userPhoneNumber := this.GetString("user_phone_number")
	userJobNumber := this.GetString("user_job_number")
	userGender, _ := this.GetInt32("user_gender")

	//call the userManage method
	this.StartSession()
	userManage, ctx := common.InitUserManage(this.CruSession)
	req := proto.AddUserReq{Name: userName, LoginName: loginName, PhoneNumber: userPhoneNumber, JobNumber: userJobNumber, Gender: userGender}
	rsp, err := userManage.AddUser(ctx, &req)
	if err != nil {
		beego.Info("======AddUser=====", rsp.UserId, "-------err--------", err)
	}

	var result map[string]interface{}
	result = make(map[string]interface{})
	result["result"] = rsp.Message
	this.Data["json"] = result
	this.ServeJSON()
	return

}

// @Title 删除用户
// @Description 删除用户
// @Param	user_name	body	int64	false	"用户id"
// @Success 200 {string} success
// @Failure 400 no enough input
// @Failure 500 server's wrong
// @router / [delete]
func (this *UserManageController) DeleteUserById() {
	type deleteInput struct {
		Delete_id string `json:"delete_id"`
	}
	inparam := deleteInput{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &inparam)
	if (err != nil) {
		beego.Error(err)
		return
	}

	deleteId, err := strconv.ParseInt(inparam.Delete_id, 10, 64)

	if (deleteId == 0) {
		beego.Info("-------deleteId error-------- deleteId=", deleteId)
		return
	}

	//call the userManage method
	this.StartSession()
	userManage, ctx := common.InitUserManage(this.CruSession)
	req := proto.DeleteUserReq{DeleteId: deleteId}
	rsp, err := userManage.DeleteUserById(ctx, &req)
	if err != nil {
		beego.Info("======DeleteUserById=====", rsp.UserId, "-------err--------", err)
	}

	var result map[string]interface{}
	result = make(map[string]interface{})
	result["result"] = rsp.Message
	this.Data["json"] = result
	this.ServeJSON()
	return
}
