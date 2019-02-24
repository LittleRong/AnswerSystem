package problem

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

// 完成User类型定义
type Problem struct {
	Problem_id      int `orm:"pk"`
	Problem_content string
	Problem_option  string
	Problem_answer  string
	Problem_class   string
	Problem_type    int
}

func AddProblem(p Problem) int {
	o := orm.NewOrm()
	id, err := o.Insert(&p)
	if err != nil {
		beego.Info("======AddProblem's err=====", err)
	} else {
		beego.Info("======AddProblem's id=====", id)
	}
	return int(id)
}

func GetEndProblemId() int {
	var p Problem
	o := orm.NewOrm()
	o.QueryTable("problem").OrderBy("Problem_id").Limit(1).One(&p, "Problem_id")
	return p.Problem_id
}

func GetNewProblemByType(start int, peoblem_type int) []Problem {
	var p []Problem
	o := orm.NewOrm()
	o.QueryTable("problem").Filter("problem_id__gte", start).Filter("problem_type", peoblem_type).All(&p)
	return p
}
