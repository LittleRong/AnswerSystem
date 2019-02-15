package controllers

import (
	"github.com/astaxie/beego"
	"hello/models/user"
)

type LoginController struct{
	beego.Controller
}

func (this *LoginController) Index(){
	/*
	user_id := this.GetSession("user_id")
	if user_id != nil {//用户已经登陆，直接进入用户中心
		user_id = user_id.(int)
		user,permission := models.IsManager(1)
		if permission == true {
			this.TplName = "index.html"
		} else {
			this.Data["json"] = &user
			this.ServeJSON()
			this.TplName = "welcome.html"
		}

	} else {//用户未登录，进入登陆界面
		this.TplName = "bad.html"
	}
	*/
	this.TplName = "index.html"
}

func (this *LoginController) Check(){
    //判断是否为POST方法

		username := this.GetString("username") // login.html中传过来的数据，这个GetInt返回两个值
		password := this.GetString("password")

		//校验
		//valid := validation.Validation{}
		//valid.Required(password, "password")
		// valid.MaxSize(id, 20, "id")
		//if valid.HasErrors() {
		//	fmt.Println(valid.Errors[0].Key + valid.Errors[0].Message)
		//	c.TplName = "bad.html"
		//	return
		//}
		var result map[string]interface{}
		user,loginFlag := user.Login(username,password)
		if loginFlag == false {	//登录失败
			result = map[string]interface{}{"result": "faild","message":"登陆失败,用户名或密码错误"}
		} else {
			user_id := this.GetSession("user_id")
			if user_id != nil { //已登陆
				result = map[string]interface{}{"result": "logged"}
			} else {
				//设置session
				user_id = user.Id
				this.SetSession("user_id", user_id)

				//判断用户权限
				if user.Permission == 1 || user.Permission == 2  {//管理员
					result = map[string]interface{}{"result": "admin"}
				} else {//普通用户
					result = map[string]interface{}{"result": "user"}
				}
			}
		}

	this.Data["json"] = result
	this.ServeJSON()
	return
}

func (this *LoginController) ChangePwdInit(){
	this.TplName = "index/change_pwd.html"
}

func (this *LoginController) ChangePwd(){
	new_pwd := this.GetString("new_password")
	old_pwd := this.GetString("old_password")
	userSession := this.GetSession("user_id")
	if userSession == nil { //未登陆
		this.Ctx.Redirect(304,"/index")
		return
	}
	user_id := userSession.(int)
	flag := user.UpdateUserPwd(user_id,old_pwd,new_pwd)
	var result map[string]interface{}
	result = make(map[string]interface{})
	if(flag){
		result["result"] = "success"
	}else{
		result["result"] = "修改失败，请联系管理员"
	}
	beego.Info("========result======",result)
	this.Data["json"] = result
	this.ServeJSON()
	return

}
