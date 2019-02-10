package event

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type Event struct{
	Event_id	       int `orm:"pk"`
	Manage_id          int
	Event_title        string
	Event_description  string
	Event_time         string
	Event_num          string
	Event_type         string
	Problem_random     bool
	Option_random      bool
	Answer_time        string
	Credit_rule        string
	Participant_num    int
}


type CreditRule struct{
	Single_score float64 `json:"single_score,string"` //单选题分数
	Multi_score float64 `json:"multi_score,string"`//多选题分数
	Fill_score float64 `json:"fill_score,string"`//填空题分数
	Judge_score float64 `json:"judge_score,string"`//判断题分数
	Person_score float64 `json:"person_score,string"`//当日本人全对额外加分
	Team_score float64 `json:"team_score,string"`//当日团队全对额外加分
	Team_score_up float64 `json:"team_score_up,string"`//团队总积分上限
	Person_score_up float64 `json:"person_score_up,string"`//个人总积分上限
}

func GetEventByEventId(event_id int) (event *Event){
	e := Event{Event_id:event_id}
	o := orm.NewOrm()
	err := o.Read(&e,"Event_id")
	if err != nil {
		return nil
	} else {
		return &e
	}
}

func GetCreditRuleByEventId(event_id int) CreditRule{
	beego.Info("========GetCreditRuleByEventId======")
	var params Event
	var creditRule CreditRule
	o := orm.NewOrm()
	o.QueryTable("event").Filter("Event_id", event_id).One(&params, "Credit_rule")
	if params.Credit_rule != ""{
		err := json.Unmarshal([]byte(params.Credit_rule), &creditRule)
		if err != nil {
			beego.Info("========err======",err)
		}
	}
	beego.Info("========creditRule======",creditRule)
	return creditRule
}