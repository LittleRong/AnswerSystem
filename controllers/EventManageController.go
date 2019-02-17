package controllers

import (
	"github.com/astaxie/beego"
	"hello/models/event"
	"strconv"
)

type EventManageController struct{
	beego.Controller
}

func (this *EventManageController) EventManageInit(){
	this.TplName = "manage/event_manage.html"
}

func (this *EventManageController) EventManage(){
	////还需要加上偏移
	//offset,_ := this.GetInt("offset")
	//limit,_ := this.GetInt("limit")
	beego.Info("========EventManage======")

	userSession := this.GetSession("user_id")
	if userSession == nil { //未登陆
		this.Ctx.Redirect(304,"/index")
		beego.Info("========未登录======")
		return
	}
	user_id := userSession.(int)
	e:=event.GetEventByManageId(user_id)
	var event_data []map[string]string
	for _,v := range e{
		var t map[string]string
		t = make(map[string]string)
		t["event_id"] = strconv.Itoa(v.Event_id)
		t["event_title"] = v.Event_title
		t["event_description"] = v.Event_description
		t["event_type"] = v.Event_type
		event_data = append(event_data, t)
	}

	var result map[string]interface{}
	result = make(map[string]interface{})
	result["event_data"] = event_data
	this.Data["json"] = result
	this.ServeJSON()
	return

}