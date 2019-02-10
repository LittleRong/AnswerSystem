package participant

import (
	"encoding/json"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
)

// 完成User类型定义
type Participant struct {
	Participant_id	int `orm:"pk"`
	Refer_event_id  int
	User_id         int
	Team_id         int
	Credit          float64
	Leader          bool
	Waited_answer   string
}

func GetEventListByUserId(user_id int) ([]Participant) {
	var p []Participant
	o := orm.NewOrm()
	o.QueryTable("participant").Filter("User_id", user_id).All(&p, "Participant_id", "Refer_event_id")
	return p
}

func GetCorrectAnswerByParticipantId(participant_id int) (map[string]interface{}) {
	var waited_answer map[string]interface{}
	waited_answer = make(map[string]interface{})
	var p Participant
	o := orm.NewOrm()
	o.QueryTable("participant").Filter("Participant_id", participant_id).One(&p, "Waited_answer")
	if p.Waited_answer != "" {
		err := json.Unmarshal([]byte(p.Waited_answer),&waited_answer)
		if err != nil {
			beego.Info("======err======",err)
		}
	}
	return waited_answer
}
