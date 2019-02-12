package credit

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type Credit_log struct {
	Credit_log_id int `orm:"pk"` //日志id
	Refer_event_id int //关联的事件的id
	Refer_participant_id int //参赛者id
	Refer_team_id int //进行操作的组id
	Change_time string //操作时间
	Change_value float64 //操作值,更改的值,正为加分，负为减分
	Change_reason string //更改原因
}

func AddCreditLog(log Credit_log)  {
	//生成新credit_log_id
	o := orm.NewOrm()
	id,err := o.Insert(&log)
	if err != nil {
		beego.Info("======AddCreditLog's err=====",err)
	} else {
		beego.Info("======AddCreditLog's id=====",id)
	}
}