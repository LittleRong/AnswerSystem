package model

import (
	"time"

	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

type Participant_haved_answer struct {
	Refer_participant_id int64 `orm:"pk"` //参赛者id
	Refer_problem_id     int64 `orm:"pk"` //题id
	Refer_team_id        int64            //关联的组id
	Answer_date          string           //用户答题日期
	User_answer          string           //用户答题结果
	True_or_false        bool             //用户答题是否正确
}

func UpdateUserAnswer(participant_id int64, problem_id int64, user_answer string, true_or_false bool) (string) {
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
		logs.Error("UpdateUserAnswer's err:", err)
		return "faild"
	} else {
		logs.Debug("UpdateUserAnswer's num:", num)
		return "success"
	}

}

func AddProblems(problems Participant_haved_answer) string {
	o := orm.NewOrm()
	_, err := o.Raw("INSERT INTO participant_haved_answer "+
		"(refer_participant_id,refer_problem_id,refer_team_id,answer_date) "+
		"VALUES (?,?,?,?) ", problems.Refer_participant_id, problems.Refer_problem_id, problems.Refer_team_id, problems.Answer_date).Exec();
	if err != nil {
		logs.Error("AddProblems's err:", err)
		return "falid"
	}
	return "success"

}

func JudgeIfHaveAnswer(participant_id int64) bool {
	var answer_date string
	now_date := time.Now()
	unix_time := now_date.Unix()
	answer_date = time.Unix(unix_time, 0).Format("2006-01-02")
	o := orm.NewOrm()
	var p []Participant_haved_answer
	num, err := o.Raw("SELECT * FROM participant_haved_answer WHERE refer_participant_id =? AND user_answer != ''  AND answer_date like ? ", participant_id, answer_date).QueryRows(&p);


	if err == nil && num > 0 {
		return true //已经提交了
	}else{
		logs.Error("JudgeIfHaveAnswer", err)
	}
	return false
}
