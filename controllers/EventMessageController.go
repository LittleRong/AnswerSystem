package controllers

import "github.com/astaxie/beego"

type EventMessageController struct{
	beego.Controller
}

func (this *EventMessageController) EventMessageInit(){
	this.TplName = "index/user_index.html"
}
