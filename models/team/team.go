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

