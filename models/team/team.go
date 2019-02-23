package team

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type Team struct {
	Team_id int `orm:"pk"`//组id
	Team_name string //组名
	Refer_event_id int //参见的事件id,关联的事件的id
	Team_credit float64 //本组在事件中的积分
}

func UpdateTeamCredit(team_id int, team_credit float64) float64 {
	new_credit := 0.0
	//更新
	team := Team{Team_id:team_id}
	o := orm.NewOrm()
	if o.Read(&team) == nil {
		old_credit := team.Team_credit
		new_credit = old_credit + team_credit
		team.Team_credit = new_credit
		beego.Info("======UpdateTeamCredit's old_credit=====",old_credit)
		beego.Info("======UpdateTeamCredit's new_credit=====",new_credit)
		if num, err := o.Update(&team,"Team_credit"); err == nil {
			beego.Info("======UpdateTeamCredit's num=====",num)
		} else if err!=nil{
			beego.Info("======UpdateTeamCredit's err=====",err)
		}
	}
	return new_credit
}

func GetTeamById(team_id int, event_id int) Team{
	var t Team
	o := orm.NewOrm()
	o.QueryTable("team").Filter("team_id", team_id).Filter("Refer_event_id", event_id).One(&t)
	beego.Info("======GetTeamById=====",t)
	return t
}

func AddTeam (team_name string,refer_event_id int) int {
	//login_name不能重复
	var t Team
	o := orm.NewOrm()

	t.Refer_event_id = refer_event_id
	t.Team_name = team_name
	t.Team_credit = 0
	id,err := o.Insert(&t)
	if err == nil {
		beego.Info("======AddTeam's id=====",id)
		return int(id)
	} else {
		beego.Info("======AddTeam's err=====",err)
		return -1
	}
}