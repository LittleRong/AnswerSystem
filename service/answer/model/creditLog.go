package model

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type Credit_log struct {
	Credit_log_id        int64 `orm:"pk"` //日志id
	Refer_event_id       int64            //关联的事件的id
	Refer_participant_id int64            //参赛者id
	Refer_team_id        int64            //进行操作的组id
	Change_time          string           //操作时间
	Change_value         float32          //操作值,更改的值,正为加分，负为减分
	Change_type          int32            //分数操作类型：1答题加分，2当日全部答对额外加分，3当日小组全部答对额外加分
	Change_reason        string           //更改原因
}

func init() {
	orm.RegisterModel(new(Credit_log))
}

func AddCreditLog(log Credit_log) (string, int64) {
	o := orm.NewOrm()
	id, err := o.Insert(&log)
	if err != nil {
		beego.Info("======AddCreditLog's err=====", err)
		return "faild", -1
	} else {
		beego.Info("======AddCreditLog's id=====", id)
		return "success", id
	}
}

func WhetherMemberAllRight(team_id int64, date string, team_num int) bool {
	var log []Credit_log
	o := orm.NewOrm()
	like_date := "%" + date + "%"
	_, err := o.Raw("SELECT DISTINCT(refer_participant_id) FROM credit_log WHERE refer_event_id = ? AND change_time LIKE ? AND change_type = ? ", team_id, like_date, 2).QueryRows(&log)
	//_, err := o.Raw("SELECT refer_participant_id FROM credit_log").QueryRows(&log)
	if (err == nil) {
		right_num := len(log)
		if (right_num == team_num) {
			return true
		}
		beego.Info("======WhetherMemberAllRight's right_num=====", right_num)
	} else {
		beego.Info("======!!!@@WhetherMemberAllRight's err=====", err)
	}
	return false
}

func GetCreditLogByTeamId(team_id int64) []Credit_log {
	var log []Credit_log
	o := orm.NewOrm()
	o.QueryTable("credit_log").Filter("refer_team_id", team_id).All(&log)
	return log
}
