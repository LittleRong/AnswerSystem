package event

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type EventProblem struct {
	Refer_event_id int
	Problem_id     int
}

func AddEventProblem(p EventProblem) {
	beego.Info("======AddProblems's problems=====", p)
	o := orm.NewOrm()
	_, err := o.Raw("INSERT INTO event_problem "+
		"(refer_event_id,problem_id) "+
		"VALUES (?,?) ", p.Refer_event_id, p.Problem_id).Exec();
	if err != nil {
		beego.Info("======AddProblems's err!!!=====", err)
	}
}
