package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"

	proto "service/protoc/eventManage"
	"web/common"
)

type EventManageController struct {
	beego.Controller
}

// @Title 获得事件管理页面
// @Description 获得事件管理页面
// @Success 200 {}
// @router / [get]
func (this *EventManageController) EventManageInit() {
	this.TplName = "manage/event_manage.html"
}

// @Title 获取事件列表
// @Description 获取事件列表
// @Success 200 {}
// @Param   offset   query   string  true       "页码"
// @Param   limit query   string  true       "一页展示数量"
// @router /all [get]
func (this *EventManageController) EventManage() {
	offset, _ := this.GetInt32("offset")
	limit, _ := this.GetInt32("limit")
	//获取用户信息
	userSession := this.GetSession("user_id")
	if userSession == nil { //未登陆
		this.Ctx.Redirect(304, "/index")
		return
	}
	userId := userSession.(int64)

	//call the userManage method
	eventManage, ctx := common.InitEventManage(this.CruSession)
	req := proto.GetEventListReq{Offset: offset, Limit: limit, ManageId: userId}
	rsp, err := eventManage.GetEventListByManageIdAndOffst(ctx, &req)
	if err != nil {
		beego.Info("======ProblemManage=====", rsp.EventList, "-------err--------", err)
	}

	var result map[string]interface{}
	result = make(map[string]interface{})
	result["event_data"] = rsp.EventList
	this.Data["json"] = result
	this.ServeJSON()
	return

}

// @Title 获得新增事件页面
// @Description 获得新增事件页面
// @Success 200 {}
// @router /newevent [get]
func (this *EventManageController) EventInsertInit() {
	this.TplName = "manage/event_insert.html"
}

// @Title 新增事件
// @Description 新增事件
// @Success 200 {}
// @Param   etitle   formData   string  true	"事件名称"
// @Param   message	formData   string  true	"事件描述"
// @Param   ekind   formData   string  true	"事件种类，如务知识类竞赛、党建知识类"
// @Param   pro_random	formData   bool  true	"是否控制题目随机顺序"
// @Param   opt_random   formData   bool  true	"是否控制选项随机顺序"
// @Param   answer_time	formData   string  true	"答题时长"
// @Param   participant_num   formData   int  true	"参赛人数"
// @Param   single	formData   int32  true	"单选题每天答题数量"
// @Param   multiple	formData   int32  true	"多选题每天答题数量"
// @Param   fill	formData   int32  true	"填空题题每天答题数量"
// @Param   judge	formData   int32  true	"判断题每天答题数量"
// @Param   start_time	formData   string  true	"事件开始日期"
// @Param   end_time	formData   string  true	"事件结束日期"
// @Param   answer_day	formData   string  true	"可以答题的日志"
// @Param   single_score	formData   number  true	"单选题答对分值"
// @Param   multiple_score	formData   number  true	"多选题答对分值"
// @Param   fill_score	formData   number  true	"填空题答对分值"
// @Param   judge_score	formData   number  true	"判断题答对分值"
// @Param   person_score	formData   number  true	"当日本人全对额外加分"
// @Param   team_score	formData   number  true	"当日团队全对额外加分"
// @Param   person_score_up	formData   number  true	"团队总积分上限"
// @Param   team_score_up	formData   number  true	"个人总积分上限"
// @router /newevent [post]
func (this *EventManageController) EventInsert() {
	userSession := this.GetSession("user_id")
	if userSession == nil { //未登陆
		this.Ctx.Redirect(304, "/index")
		beego.Info("========未登录======")
		return
	}
	manage_id := userSession.(int64)

	//获取传入参数
	etitle := this.GetString("etitle")
	message := this.GetString("message")
	ekind := this.GetString("ekind")
	pro_random, _ := this.GetBool("pro_random")
	opt_random, _ := this.GetBool("opt_random")
	answer_time := this.GetString("answer_time")
	participant_num, _ := this.GetInt("participant_num")

	//event_time
	start_time := this.GetString("start_time")
	end_time := this.GetString("end_time")
	answer_day := this.GetString("answer_day")
	//answer_day转换成字符串
	etime := proto.EventTime{StartTime: start_time, EndTime: end_time, AnswerDay: answer_day}
	event_time, _ := json.Marshal(etime)

	//event_num
	single, _ := this.GetInt32("single")
	multiple, _ := this.GetInt32("multiple")
	fill, _ := this.GetInt32("fill")
	judge, _ := this.GetInt32("judge")
	enum := proto.ProblemNum{Single: single, Multiple: multiple, Fill: fill, Judge: judge}
	event_num, _ := json.Marshal(enum)

	//credit_rule
	single_score, _ := this.GetFloat("single_score")
	multiple_score, _ := this.GetFloat("multiple_score")
	fill_score, _ := this.GetFloat("fill_score")
	judge_score, _ := this.GetFloat("judge_score")
	person_score, _ := this.GetFloat("person_score")
	team_score, _ := this.GetFloat("team_score")
	person_score_up, _ := this.GetFloat("person_score_up")
	team_score_up, _ := this.GetFloat("team_score_up")
	crule := proto.CreditRule{SingleScore: single_score, MultipleScore: multiple_score, FillScore: fill_score, JudgeScore: judge_score, PersonScore: person_score, PersonScoreUp: person_score_up, TeamScore: team_score, TeamScoreUp: team_score_up}
	credit_rule, _ := json.Marshal(crule)

	//call the userManage method
	eventManage, ctx := common.InitEventManage(this.CruSession)
	req := proto.AddEventReq{ManageId: int64(manage_id),
		EventTitle:       etitle,
		EventDescription: message,
		EventTime:        string(event_time),
		EventNum:         string(event_num),
		EventType:        ekind,
		ProblemRandom:    pro_random,
		OptionRandom:     opt_random,
		AnswerTime:       answer_time,
		CreditRule:       string(credit_rule),
		ParticipantNum:   int32(participant_num)}
	rsp, err := eventManage.AddNewEvent(ctx, &req)
	if err != nil {
		beego.Info("-------err--------", err)
	}

	var result map[string]interface{}
	result = make(map[string]interface{})
	if (rsp.EventId != -1) {
		result["result"] = rsp.Message
		result["event_id"] = rsp.EventId
	} else {
		result["result"] = rsp.Message
	}

	this.SetSession("new_event_id", int(rsp.EventId))

	this.Data["json"] = result
	this.ServeJSON()
	return
}
