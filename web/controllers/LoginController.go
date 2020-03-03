package controllers

import (
	"github.com/astaxie/beego"
	"github.com/spf13/viper"
	userProto "service/protoc/userManage"
	"web/common"
)

type LoginController struct {
	beego.Controller
}

// @Title 获得登陆首页
// @Description 获得登陆首页
// @Success 200 {}
// @router / [get]
func (this *LoginController) Index() {
	this.TplName = "index.html"
}

// @Title Check user login校验用户登陆信息
// @Description Check user login by username and password通过username和password校验用户登陆信息
// @Success 200 {string} result
// @Param   username   query   string  false       "用户的登陆名username"
// @Param   password query   string  false       "用户的密码password"
// @router /check [post]
func (this *LoginController) Check() {
	username := this.GetString("username") // login.html中传过来的数据，这个GetInt返回两个值
	password := this.GetString("password")

	var result map[string]interface{}
	userManage, ctx := common.InitUserManage(this.CruSession)
	req := userProto.LoginReq{Username: username, Pwd: password}

	LoginRsp, err := userManage.Login(ctx, &req)
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
		if LoginRsp.Permission == viper.GetInt32("enum.permission.admin") || LoginRsp.Permission == viper.GetInt32("enum.permission.administrators") { //管理员
			result = map[string]interface{}{"result": "admin"}
		} else { //普通用户
			result = map[string]interface{}{"result": "user"}
		}

		//存放token
		if (LoginRsp.Token != "") {
			this.SetSession("token", LoginRsp.Token)
			beego.Info("========token======", LoginRsp.Token)
		} else {
			beego.Error("获取token失败")
		}

	}

	this.Data["json"] = result
	this.ServeJSON()
	return
}

// @Title 获得修改密码页面
// @Description 获得修改密码页面
// @Success 200 {}
// @router /change_pwd_init [get]
func (this *LoginController) ChangePwdInit() {
	this.TplName = "index/change_pwd.html"
}

// @Title Change user's password修改用户密码
// @Description Change user's password修改用户密码
// @Success 200 {string} result
// @Param   user_id   formData   int  false       "用户id"
// @Param   new_password	formData   string  false       "用户的旧密码"
// @Param   old_password	formData   string  false       "用户的新密码"
// @router /password [post]
func (this *LoginController) ChangePwd() {
	new_pwd := this.GetString("new_password")
	old_pwd := this.GetString("old_password")
	userSession := this.GetSession("user_id")
	if userSession == nil { //未登陆
		this.Ctx.Redirect(304, "/index")
		return
	}
	user_id := userSession.(int64)
	userManage, ctx := common.InitUserManage(this.CruSession)
	req := userProto.UpdatePwdReq{UserId: user_id, OldPwd: old_pwd, NewPwd: new_pwd}
	rsp, err := userManage.UpdateUserPwd(ctx, &req)
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

// @Title logout登出请求
// @Description logout登出请求
// @Success 200 {}
// @router /logout [get]
func (this *LoginController) Logout() {
	this.DestroySession()
	this.Ctx.Redirect(302, "/index")
	return
}
