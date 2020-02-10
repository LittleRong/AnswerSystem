package model

import (
	"encoding/json"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

// 完成User类型定义
type Participant struct {
	Participant_id int64 `orm:"pk"`
	Refer_event_id int64
	User_id        int64
	Team_id        int64
	Credit         float64
	Leader         bool
	Waited_answer  string
}

func init() {
	orm.RegisterModel(new(Participant))
}

func GetParticipantListByUserId(user_id int64) ([]Participant) {
	var p []Participant
	o := orm.NewOrm()
	o.QueryTable("participant").Filter("User_id", user_id).All(&p, "Participant_id", "Refer_event_id")
	return p
}

func GetCorrectAnswerByParticipantId(participant_id int64) (map[string]interface{}) {
	var waited_answer map[string]interface{}
	waited_answer = make(map[string]interface{})
	var p Participant
	o := orm.NewOrm()
	o.QueryTable("participant").Filter("Participant_id", participant_id).One(&p, "Waited_answer")
	if p.Waited_answer != "" {
		err := json.Unmarshal([]byte(p.Waited_answer), &waited_answer)
		if err != nil {
			logs.Error("GetCorrectAnswerByParticipantId's err:", err)
		}
	}
	return waited_answer
}

func GetAnswerTimeByParticipantId(participant_id int64) (string) {
	var waited_answer map[string]interface{}
	waited_answer = make(map[string]interface{})
	var p Participant
	var p_time string

	o := orm.NewOrm()
	o.QueryTable("participant").Filter("Participant_id", participant_id).One(&p, "Waited_answer")
	if p.Waited_answer != "" {
		err := json.Unmarshal([]byte(p.Waited_answer), &waited_answer)
		if err != nil {
			logs.Error("GetAnswerTimeByParticipantId's err:", err)
		}
		//获取时间
		p_time = waited_answer["participant_time"].(string)
		//if (p_time != "") {
		//	timeLayout := "2006-01-02 15:04:05"  //转化所需模板，go默认时间
		//	loc, _ := time.LoadLocation("Local") //获取本地时区
		//	participant_time, _ = time.ParseInLocation(timeLayout, p_time, loc)
		//}
	}
	return p_time
}

func GetParticipantById(user_id int64, event_id int64) Participant {
	var p Participant
	o := orm.NewOrm()
	o.QueryTable("participant").Filter("user_id", user_id).Filter("Refer_event_id", event_id).One(&p)
	return p
}

func UpdateParticipantCredit(participant_id int64, credit float64) float64 {
	new_credit := 0.0
	//更新
	participant := Participant{Participant_id: participant_id}
	o := orm.NewOrm()
	if o.Read(&participant) == nil {
		old_credit := participant.Credit
		new_credit = old_credit + credit
		participant.Credit = new_credit
		if num, err := o.Update(&participant, "Credit"); err == nil {
			logs.Debug("UpdateParticipantCredit's num:", num)
		} else if err != nil {
			logs.Error("UpdateParticipantCredit's err", err)
		}
	}
	return new_credit
}

func GetMemberCreditByTeamId(team_id int64, event_id int64) []Participant {
	var p []Participant
	o := orm.NewOrm()
	o.QueryTable("participant").Filter("team_id", team_id).Filter("Refer_event_id", event_id).All(&p)
	return p
}

func UpdateParticipantWaitedAnswer(participant_id int64, answer string) string {
	//待增加对answer格式的校验

	participant := Participant{Participant_id: participant_id}
	o := orm.NewOrm()
	if o.Read(&participant) == nil {
		participant.Waited_answer = answer
		if num, err := o.Update(&participant, "Waited_answer"); err == nil {
			logs.Debug("UpdateParticipantWaitedAnswer's num", num)
			return "success"
		} else if err != nil {
			logs.Error("UpdateParticipantWaitAnswer's err", err)
			return "faild"
		}
	}
	return "faild"

}

func AddParticipant(user_id int64, refer_event_id int64, team_id int64, leader bool) int64 {
	//login_name不能重复
	var p Participant
	o := orm.NewOrm()

	p.Team_id = team_id
	p.Refer_event_id = refer_event_id
	p.Credit = 0
	p.User_id = user_id
	p.Leader = leader
	p.Waited_answer = "null"

	id, err := o.Insert(&p)
	if err == nil {
		logs.Debug("AddParticipant's id:", id)
		return p.Participant_id
	} else {
		logs.Error("AddParticipant's err:", err)
		return -1
	}
}
