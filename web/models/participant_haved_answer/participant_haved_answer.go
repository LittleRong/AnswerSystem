package participant_haved_answer

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"time"
)

type Participant_haved_answer struct {
	Refer_participant_id int `orm:"pk"` //参赛者id
	Refer_problem_id     int `orm:"pk"` //题id
	Refer_team_id        int            //关联的组id
	Answer_date          string         //用户答题日期
	User_answer          string         //用户答题结果
	True_or_false        bool           //用户答题是否正确
}

func UpdateUserAnswer(participant_id int, problem_id int, user_answer string, true_or_false bool) {
	//用户答题时间
	var answer_date string
	now_date := time.Now()
	unix_time := now_date.Unix()
	answer_date = time.Unix(unix_time, 0).Format("2006-01-02 15:04:05")

	//更新
	o := orm.NewOrm()
	rawSetter, err := o.Raw("UPDATE participant_haved_answer "+
		"SET answer_date=?, user_answer=?, true_or_false=? "+
		"WHERE refer_participant_id=? AND refer_problem_id=?", answer_date, user_answer, true_or_false, participant_id, problem_id).Exec();
	num, err := rawSetter.RowsAffected()
	if err != nil {
		beego.Info("======UpdateUserAnswer's err=====", err)
	} else {
		beego.Info("======UpdateUserAnswer's num=====", num)
	}

}

func AddProblems(problems Participant_haved_answer) {
	beego.Info("======AddProblems's problems=====", problems)
	o := orm.NewOrm()
	_, err := o.Raw("INSERT INTO participant_haved_answer "+
		"(refer_participant_id,refer_problem_id,refer_team_id,answer_date) "+
		"VALUES (?,?,?,?) ", problems.Refer_participant_id, problems.Refer_problem_id, problems.Refer_team_id, problems.Answer_date).Exec();
	if err != nil {
		beego.Info("======AddProblems's err!!!=====", err)
	}

}

func JudgeIfHaveAnswer(participant_id int) bool {
	var answer_date string
	now_date := time.Now()
	unix_time := now_date.Unix()
	answer_date = time.Unix(unix_time, 0).Format("2006-01-02")
	o := orm.NewOrm()
	var p []Participant_haved_answer
	num, err := o.Raw("SELECT * FROM participant_haved_answer WHERE refer_participant_id =? AND user_answer != ''  AND answer_date like ? ", participant_id, answer_date).QueryRows(&p);
	beego.Info("**************JudgeIfHaveAnswer*****************", err)

	if err == nil && num >0 {
		return true//已经提交了
	}
	return false
}