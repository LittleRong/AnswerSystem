package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["web/controllers:AnswerController"] = append(beego.GlobalControllerRouter["web/controllers:AnswerController"],
        beego.ControllerComments{
            Method: "ShowProblemsPage",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["web/controllers:AnswerController"] = append(beego.GlobalControllerRouter["web/controllers:AnswerController"],
        beego.ControllerComments{
            Method: "GetUserAnswers",
            Router: `/user_answer`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["web/controllers:AnswerController"] = append(beego.GlobalControllerRouter["web/controllers:AnswerController"],
        beego.ControllerComments{
            Method: "GetUserProblems",
            Router: `/user_problems`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["web/controllers:EventManageController"] = append(beego.GlobalControllerRouter["web/controllers:EventManageController"],
        beego.ControllerComments{
            Method: "EventManageInit",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["web/controllers:EventManageController"] = append(beego.GlobalControllerRouter["web/controllers:EventManageController"],
        beego.ControllerComments{
            Method: "EventManage",
            Router: `/all`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["web/controllers:EventManageController"] = append(beego.GlobalControllerRouter["web/controllers:EventManageController"],
        beego.ControllerComments{
            Method: "EventInsertInit",
            Router: `/newevent`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["web/controllers:EventManageController"] = append(beego.GlobalControllerRouter["web/controllers:EventManageController"],
        beego.ControllerComments{
            Method: "EventInsert",
            Router: `/newevent`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["web/controllers:EventMessageController"] = append(beego.GlobalControllerRouter["web/controllers:EventMessageController"],
        beego.ControllerComments{
            Method: "EventMessageInit",
            Router: `/detail`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["web/controllers:EventMessageController"] = append(beego.GlobalControllerRouter["web/controllers:EventMessageController"],
        beego.ControllerComments{
            Method: "GetEventMessage",
            Router: `/detail/:event_id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["web/controllers:LoginController"] = append(beego.GlobalControllerRouter["web/controllers:LoginController"],
        beego.ControllerComments{
            Method: "Index",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["web/controllers:LoginController"] = append(beego.GlobalControllerRouter["web/controllers:LoginController"],
        beego.ControllerComments{
            Method: "ChangePwdInit",
            Router: `/change_pwd_init`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["web/controllers:LoginController"] = append(beego.GlobalControllerRouter["web/controllers:LoginController"],
        beego.ControllerComments{
            Method: "Check",
            Router: `/check`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["web/controllers:LoginController"] = append(beego.GlobalControllerRouter["web/controllers:LoginController"],
        beego.ControllerComments{
            Method: "Logout",
            Router: `/logout`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["web/controllers:LoginController"] = append(beego.GlobalControllerRouter["web/controllers:LoginController"],
        beego.ControllerComments{
            Method: "ChangePwd",
            Router: `/password`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["web/controllers:ParticipantManageController"] = append(beego.GlobalControllerRouter["web/controllers:ParticipantManageController"],
        beego.ControllerComments{
            Method: "ParticipantInsertInit",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["web/controllers:ParticipantManageController"] = append(beego.GlobalControllerRouter["web/controllers:ParticipantManageController"],
        beego.ControllerComments{
            Method: "ParticipantGetUser",
            Router: `/all`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["web/controllers:ParticipantManageController"] = append(beego.GlobalControllerRouter["web/controllers:ParticipantManageController"],
        beego.ControllerComments{
            Method: "EventParticipantInsert",
            Router: `/batch`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["web/controllers:ProblemManageController"] = append(beego.GlobalControllerRouter["web/controllers:ProblemManageController"],
        beego.ControllerComments{
            Method: "ProblemManageInit",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["web/controllers:ProblemManageController"] = append(beego.GlobalControllerRouter["web/controllers:ProblemManageController"],
        beego.ControllerComments{
            Method: "ProblemManage",
            Router: `/all`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["web/controllers:ProblemManageController"] = append(beego.GlobalControllerRouter["web/controllers:ProblemManageController"],
        beego.ControllerComments{
            Method: "ProblemUploadInit",
            Router: `/upload`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["web/controllers:ProblemManageController"] = append(beego.GlobalControllerRouter["web/controllers:ProblemManageController"],
        beego.ControllerComments{
            Method: "ProblemFileInsert",
            Router: `/upload`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["web/controllers:UserIndexController"] = append(beego.GlobalControllerRouter["web/controllers:UserIndexController"],
        beego.ControllerComments{
            Method: "UserIndexInit",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["web/controllers:UserIndexController"] = append(beego.GlobalControllerRouter["web/controllers:UserIndexController"],
        beego.ControllerComments{
            Method: "UserIndex",
            Router: `/user_event`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["web/controllers:UserManageController"] = append(beego.GlobalControllerRouter["web/controllers:UserManageController"],
        beego.ControllerComments{
            Method: "UserManageInit",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["web/controllers:UserManageController"] = append(beego.GlobalControllerRouter["web/controllers:UserManageController"],
        beego.ControllerComments{
            Method: "ChangeUser",
            Router: `/`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["web/controllers:UserManageController"] = append(beego.GlobalControllerRouter["web/controllers:UserManageController"],
        beego.ControllerComments{
            Method: "AddUser",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["web/controllers:UserManageController"] = append(beego.GlobalControllerRouter["web/controllers:UserManageController"],
        beego.ControllerComments{
            Method: "DeleteUserById",
            Router: `/`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["web/controllers:UserManageController"] = append(beego.GlobalControllerRouter["web/controllers:UserManageController"],
        beego.ControllerComments{
            Method: "UserManage",
            Router: `/all`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
