package routers

import (
	"github.com/astaxie/beego"

	"web/controllers"
)

func init() {
	beego.Router("/index", &controllers.LoginController{}, `get:Index`)
	beego.Router("/check", &controllers.LoginController{}, `post:Check`)
	beego.Router("/logout", &controllers.LoginController{}, `get:Logout`)
	beego.Router("/index/change_pwd_init", &controllers.LoginController{}, `get:ChangePwdInit`)
	beego.Router("/index/change_pwd", &controllers.LoginController{}, `post:ChangePwd`)
	beego.Router("/manage/user_manage_init", &controllers.UserManageController{}, `get:UserManageInit`)
	beego.Router("/manage/user_manage", &controllers.UserManageController{}, `get:UserManage`)
	beego.Router("/manage/change_user", &controllers.UserManageController{}, `post:ChangeUser`)
	beego.Router("/manage/add_user", &controllers.UserManageController{}, `post:AddUser`)
	beego.Router("/manage/delete_user", &controllers.UserManageController{}, `post:DeleteUserById`)

	beego.Router("/manage/problem_manage_init", &controllers.ProblemManageController{}, `get:ProblemManageInit`)
	beego.Router("/manage/problem_manage", &controllers.ProblemManageController{}, `get:ProblemManage`)
	beego.Router("/manage/change_problem", &controllers.ProblemManageController{}, `post:ChangeProblem`)
	beego.Router("/manage/add_problem", &controllers.ProblemManageController{}, `post:AddProblem`)
	beego.Router("/manage/delete_problem", &controllers.ProblemManageController{}, `post:DeleteProblem`)
	beego.Router("/manage/problem_upload_init", &controllers.ProblemManageController{}, `get:ProblemUploadInit`)
	beego.Router("/manage/problem_file_insert", &controllers.ProblemManageController{}, `post:ProblemFileInsert`)

	beego.Router("/manage/event_manage_init", &controllers.EventManageController{}, `get:EventManageInit`)
	beego.Router("/manage/event_manage", &controllers.EventManageController{}, `get:EventManage`)
	beego.Router("/manage/event_insert_init", &controllers.EventManageController{}, `get:EventInsertInit`)
	beego.Router("/manage/event_insert", &controllers.EventManageController{}, `post:EventInsert`)

	beego.Router("/manage/participant_insert_init", &controllers.ParticipantManageController{}, `get:ParticipantInsertInit`)
	beego.Router("/manage/participant_get_user", &controllers.ParticipantManageController{}, `get:ParticipantGetUser`)
	beego.Router("/manage/event_participant_insert", &controllers.ParticipantManageController{}, `post:EventParticipantInsert`)

	beego.Router("/index/user_index", &controllers.UserIndexController{}, `get:UserIndex`)
	beego.Router("/index/user_index_init", &controllers.UserIndexController{}, `get:UserIndexInit`)
	beego.Router("/answer/user_problem", &controllers.AnswerController{}, `get:ShowProblemsPage`)
	beego.Router("/answer/get_user_problem", &controllers.AnswerController{}, `get:GetUserProblems`)
	beego.Router("/answer/get_user_answer", &controllers.AnswerController{}, `post:GetUserAnswers`)
	beego.Router("/answer/event_message_init", &controllers.EventMessageController{}, `get:EventMessageInit`)
	beego.Router("/answer/get_event_message", &controllers.EventMessageController{}, `get:GetEventMessage`)
}
