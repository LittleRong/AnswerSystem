package participant

import (
	"encoding/json"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
	"time"
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

func GetAnswerTimeByParticipantId(participant_id int) (time.Time) {
	var waited_answer map[string]interface{}
	waited_answer = make(map[string]interface{})
	var p Participant
	var participant_time time.Time
	participant_time = time.Time{}
	o := orm.NewOrm()
	o.QueryTable("participant").Filter("Participant_id", participant_id).One(&p, "Waited_answer")
	if p.Waited_answer != "" {
		err := json.Unmarshal([]byte(p.Waited_answer),&waited_answer)
		if err != nil {
			beego.Info("======err======",err)
		}
		//获取时间
		p_time := waited_answer["participant_time"].(string)
		if(p_time!=""){
			timeLayout := "2006-01-02 15:04:05"     //转化所需模板，go默认时间
			loc, _ := time.LoadLocation("Local")    //获取本地时区
			participant_time, _ = time.ParseInLocation(timeLayout, p_time,loc)
		}
	}
	beego.Info("======participant_time======",participant_time)
	return participant_time
}

func GetParticipantById(user_id int, event_id int) *Participant{
	e := Participant{Refer_event_id:event_id, User_id:user_id}
	o := orm.NewOrm()
	err := o.Read(&e,"Participant_id")
	if err != nil {
		return nil
	} else {
		return &e
	}
}

func UpdateParticipantCredit(participant_id int,credit float64) float64{
	new_credit := 0.0
	//更新
	participant := Participant{Participant_id:participant_id}
	o := orm.NewOrm()
	if o.Read(&participant) == nil {
		old_credit := participant.Credit
		new_credit = old_credit + credit
		participant.Credit = new_credit
		beego.Info("======UpdateParticipantCredit's old_credit=====",old_credit)
		beego.Info("======UpdateParticipantCredit's new_credit=====",new_credit)
		if num, err := o.Update(&participant,"Credit"); err == nil {
			beego.Info("======num=====",num)
		} else if err!=nil{
			beego.Info("======UpdateUserAnswer's err=====",err)
		}
	}
	return new_credit
}
