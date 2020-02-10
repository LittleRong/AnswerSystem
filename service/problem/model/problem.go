package model

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

type Problem struct {
	Problem_id      int64 `orm:"pk"`
	Problem_content string
	Problem_option  string
	Problem_answer  string
	Problem_class   string
	Problem_type    int32
}

func init() {
	orm.RegisterModel(new(Problem)) // 注册模型，建立User类型对象，注册模型时，需要引入包
}

func AddProblem(p Problem) (int64,string) {
	o := orm.NewOrm()
	id, err := o.Insert(&p)
	if err != nil {
		logs.Error("AddProblem's err:", err)
		return -1,"falid"
	} else {
		logs.Debug("AddProblem's id:", id)
		return id,"success"
	}

}

func GetEndProblemId() int64 {
	var p Problem
	o := orm.NewOrm()
	o.QueryTable("problem").OrderBy("Problem_id").Limit(1).One(&p, "Problem_id")
	return p.Problem_id
}

func GetNewProblemByType(start int64, problem_type int32) []Problem {
	var p []Problem
	o := orm.NewOrm()
	o.QueryTable("problem").Filter("problem_id__gte", start).Filter("problem_type", problem_type).All(&p)
	return p
}

func GetProblemListByOffstAndLimit(offset int, limit int) []Problem {
	var u []Problem
	o := orm.NewOrm()
	offset = offset - 1
	o.QueryTable("problem").Offset(offset * limit).Limit(limit).All(&u)
	return u
}
