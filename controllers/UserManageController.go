package controllers

import "github.com/astaxie/beego"

type UserManageController struct{
	beego.Controller
}

func (this *UserManageController) UserManage(){
	this.TplName = "index/user_manage.html"
}