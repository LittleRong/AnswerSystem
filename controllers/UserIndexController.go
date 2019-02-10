package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"hello/models/event"
	"hello/models/participant"
	"hello/models/user"
	"strconv"
)

type UserIndexController struct{
	beego.Controller
}

func (this *UserIndexController) UserIndex(){
	this.TplName = "index/user_index.html"
}

func (this *UserIndexController) UserIndexInit(){
	var result map[string]interface{}
	result = make(map[string]interface{})
	//获取用户信息
	var user_message user.User
	user_id := this.GetSession("user_id")
	if user_id == nil { //未登陆
		this.Ctx.Redirect(304,"/index")
		return
	} else {
		user_message = user.GetUserById(user_id.(int))
	}

	//获取用户参与的事件，并获取事件信息
	var event_message map[string]string
	event_message = make(map[string]string)
	var event_message_array []map[string]string
	user_event_list := participant.GetEventListByUserId(user_id.(int))
	for _,valus := range user_event_list {
		event := event.GetEventByEventId(valus.Refer_event_id)
		event_message["event_id"] = strconv.Itoa(event.Event_id)
		event_message["event_title"] = event.Event_title
		event_message["event_description"] = event.Event_description
		event_message["participant_num"] = strconv.Itoa(event.Participant_num)

		event_time := event.Event_time
		var event_time_map map[string]interface{}
		//使用 json.Unmarshal(data []byte, v interface{})进行转换,返回 error 信息
		if err := json.Unmarshal([]byte(event_time), &event_time_map); err != nil {
			return
		}
		event_message["start_time"] = event_time_map["start_time"].(string)
		event_message["end_time"] = event_time_map["end_time"].(string)

		event_num := event.Event_num
		var event_num_map map[string]interface{}
		//使用 json.Unmarshal(data []byte, v interface{})进行转换,返回 error 信息
		if err := json.Unmarshal([]byte(event_num), &event_num_map); err != nil {
			return
		}
		event_message["single"] = event_num_map["single"].(string)
		event_message["fill"] = event_num_map["fill"].(string)
		event_message["judge"] = event_num_map["judge"].(string)
		event_message["multiple"] = event_num_map["multiple"].(string)

		//增加
		event_message_array = append(event_message_array,event_message)

	}
	beego.Info(event_message_array)

	//返回
	result["user_message"] = user_message
	result["event_message"] = event_message_array
	this.Data["json"] = result
	this.ServeJSON()
	return

}