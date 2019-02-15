package routers

import (
	"hello/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/index", &controllers.LoginController{}, `get:Index`)
	beego.Router("/check", &controllers.LoginController{}, `post:Check`)
	beego.Router("/index/user_manage", &controllers.UserManageController{}, `get:UserManage`)
	beego.Router("/index/user_index", &controllers.UserIndexController{}, `get:UserIndex`)
	beego.Router("/index/user_index_init", &controllers.UserIndexController{}, `get:UserIndexInit`)
	beego.Router("/answer/user_problem", &controllers.AnswerController{}, `get:ShowProblemsPage`)
	beego.Router("/answer/get_user_problem", &controllers.AnswerController{}, `get:GetUserProblems`)
	beego.Router("/answer/get_user_answer", &controllers.AnswerController{}, `post:GetUserAnswers`)
	beego.Router("/answer/event_message_init", &controllers.EventMessageController{}, `get:EventMessageInit`)
	beego.Router("/answer/get_event_message", &controllers.EventMessageController{}, `get:GetEventMessage`)
}
