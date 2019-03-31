package controllers

import (
	"context"

	"github.com/astaxie/beego"

	userProto "service/protoc/userManage"
	"web/common"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Index() {
	this.TplName = "index.html"
}

func (this *LoginController) Check() {
	username := this.GetString("username") // login.html中传过来的数据，这个GetInt返回两个值
	password := this.GetString("password")

	var result map[string]interface{}
	userManage := common.InitUserManage()
	req := userProto.LoginReq{Username: username, Pwd: password}
	LoginRsp, err := userManage.Login(context.TODO(), &req)
	if err != nil {
		beego.Info("-------err--------", err)
	}

	if LoginRsp.LoginFlag == false { //登录失败
		result = map[string]interface{}{"result": "faild", "message": "登陆失败,用户名或密码错误"}
	} else {
		//设置session
		this.SetSession("user_id", LoginRsp.UserId)
		this.SetSession("permission", LoginRsp.Permission)

		//判断用户权限
		beego.Info("========Check======", LoginRsp.Permission)
		beego.Info("========Check user_id======", LoginRsp.UserId)
		if LoginRsp.Permission == 1 || LoginRsp.Permission == 2 { //管理员
			result = map[string]interface{}{"result": "admin"}
		} else { //普通用户
			result = map[string]interface{}{"result": "user"}
		}
	}

	this.Data["json"] = result
	this.ServeJSON()
	return
}

func (this *LoginController) ChangePwdInit() {
	this.TplName = "index/change_pwd.html"
}

func (this *LoginController) ChangePwd() {
	new_pwd := this.GetString("new_password")
	old_pwd := this.GetString("old_password")
	userSession := this.GetSession("user_id")
	if userSession == nil { //未登陆
		this.Ctx.Redirect(304, "/index")
		return
	}
	user_id := userSession.(int64)
	userManage := common.InitUserManage()
	req := userProto.UpdatePwdReq{UserId: user_id, OldPwd: old_pwd, NewPwd: new_pwd}
	rsp, err := userManage.UpdateUserPwd(context.TODO(), &req)
	if err == nil {
		beego.Info("-------err--------", err)
	}

	var result map[string]interface{}
	result = make(map[string]interface{})
	result["result"] = rsp.Message
	beego.Info("========result======", result)
	this.Data["json"] = result
	this.ServeJSON()
	return

}

func (this *LoginController) Logout() {
	this.DestroySession()
	this.Ctx.Redirect(302, "/index")
	return
}
