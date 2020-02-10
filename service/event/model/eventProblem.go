package model

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

type EventProblem struct {
	Refer_event_id int64
	Problem_id     int64
}

func AddEventProblem(p EventProblem) (string){
	o := orm.NewOrm()
	_, err := o.Raw("INSERT INTO event_problem "+
		"(refer_event_id,problem_id) "+
		"VALUES (?,?) ", p.Refer_event_id, p.Problem_id).Exec();
	if err != nil {
		logs.Error("AddProblems's ", err)
		return "faild"
	}
	return "success"
}
