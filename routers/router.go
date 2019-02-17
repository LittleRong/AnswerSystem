package routers

import (
	"hello/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/index", &controllers.LoginController{}, `get:Index`)
	beego.Router("/check", &controllers.LoginController{}, `post:Check`)
	beego.Router("/index/change_pwd_init", &controllers.LoginController{}, `get:ChangePwdInit`)
	beego.Router("/index/change_pwd", &controllers.LoginController{}, `post:ChangePwd`)
	beego.Router("/manage/user_manage_init", &controllers.UserManageController{}, `get:UserManageInit`)
	beego.Router("/manage/user_manage", &controllers.UserManageController{}, `get:UserManage`)
	beego.Router("/manage/change_user", &controllers.UserManageController{}, `post:ChangeUser`)
	beego.Router("/manage/add_user", &controllers.UserManageController{}, `post:AddUser`)
	beego.Router("/manage/delete_user", &controllers.UserManageController{}, `post:DeleteUser`)
	beego.Router("/index/user_index", &controllers.UserIndexController{}, `get:UserIndex`)
	beego.Router("/index/user_index_init", &controllers.UserIndexController{}, `get:UserIndexInit`)
	beego.Router("/answer/user_problem", &controllers.AnswerController{}, `get:ShowProblemsPage`)
	beego.Router("/answer/get_user_problem", &controllers.AnswerController{}, `get:GetUserProblems`)
	beego.Router("/answer/get_user_answer", &controllers.AnswerController{}, `post:GetUserAnswers`)
	beego.Router("/answer/event_message_init", &controllers.EventMessageController{}, `get:EventMessageInit`)
	beego.Router("/answer/get_event_message", &controllers.EventMessageController{}, `get:GetEventMessage`)
}
