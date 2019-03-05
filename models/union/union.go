package union

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"web/models/event"
	"web/models/participant"
	"web/models/participant_haved_answer"
	"web/models/problem"
	"strconv"
	"time"
)

func GetProblemNoAnswer(user_id int, event_id int, team_id int, participant_id int, problemNum event.ProblemNum) (map[string]interface{}, bool,bool) {
	var problems []problem.Problem
	buildFlag := false //是否已经生成过题目
	answerFlag := false//是否已经答题
	o := orm.NewOrm()
	now_time := time.Now()
	unix_time := now_time.Unix()
	now := time.Unix(unix_time, 0).Format("2006-01-02") //设置时间戳 使用模板格式化为日期字符串
	answer_date := time.Unix(unix_time, 0).Format("2006-01-02 15:04:05")

	//*****************************1.检查是否已完成答题*************************************************
	answerFlag = participant_haved_answer.JudgeIfHaveAnswer(participant_id)
	beego.Info("**************JudgeIfHaveAnswer*****************", answerFlag)
	if answerFlag == true {
		return nil, true, answerFlag
	}

	//*****************************2.检查是否已经生成题目，若已经生成，直接查询返回*************************
	_, err := o.Raw("SELECT problem.* "+
		"FROM problem, participant_haved_answer "+
		"WHERE problem.problem_id = participant_haved_answer.refer_problem_id "+
		"AND participant_haved_answer.refer_participant_id = ? "+
		"AND participant_haved_answer.answer_date = ? ", participant_id, now).QueryRows(&problems)
	beego.Info("problems", problems)
	if problems == nil && err == nil {
		//*****************************3.未生成题目，则新生成题目*************************
		buildFlag = false
		//1.生成单选题
		single_problem := GeneratingProblems(event_id, participant_id, 1, problemNum.Single)
		problems = append(problems, single_problem...)

		//2.生成多选题
		multiple_problem := GeneratingProblems(event_id, participant_id, 2, problemNum.Multiple)
		problems = append(problems, multiple_problem...)

		//3.生成填空题
		fill_problem := GeneratingProblems(event_id, participant_id, 0, problemNum.Fill)
		problems = append(problems, fill_problem...)

		//4.生成判断题
		judge_problem := GeneratingProblems(event_id, participant_id, 3, problemNum.Fill)
		problems = append(problems, judge_problem...)

		//*****************************4.将新题目的答案写入participant表wait_answer字段*****
		var waited_answer map[string]interface{}
		waited_answer = GeneratingWaitedAnswer(problems, answer_date)
		str, _ := json.Marshal(waited_answer)
		participant.UpdateParticipantWaitedAnswer(participant_id, string(str))

		//*****************************5.将新题目拆入participant_haved_answer表*********
		for _, v := range problems {
			n := participant_haved_answer.Participant_haved_answer{Refer_participant_id: participant_id,
				Refer_problem_id: v.Problem_id,
				Refer_team_id:    team_id,
				Answer_date:      answer_date}
			participant_haved_answer.AddProblems(n)
		}
	} else {
		buildFlag = true
	}

	//*****************************5.设置传给前端的参数*****************************************
	result := GeneratingFrontProblems(problems)

	return result, buildFlag, answerFlag
}

func GeneratingProblems(event_id int, participant_id int, problem_type int, problem_num int) []problem.Problem {
	if (problem_num <= 0) {
		return nil
	}

	o := orm.NewOrm()
	var problems []problem.Problem
	_, err := o.Raw("SELECT problem.* "+
		"FROM problem, event_problem "+
		"WHERE problem.problem_id = event_problem.problem_id "+
		"AND event_problem.refer_event_id = ? "+
		"AND problem.problem_type = ? "+
		"AND problem.problem_id NOT IN "+
		"(SELECT refer_problem_id FROM participant_haved_answer WHERE refer_participant_id = ?) LIMIT ?", event_id, problem_type, participant_id, problem_num).QueryRows(&problems)
	//增加随机！！
	if err == nil {
		beego.Info("========GeneratingProblems's problems======", problems)
		return problems
	} else {
		beego.Info("========GeneratingProblems's err======", err)
		return nil
	}

}

func GeneratingFrontProblems(problems []problem.Problem) map[string]interface{} {
	var single []map[string]string
	var mutiple []map[string]string
	var fill []map[string]string
	var judge []map[string]string

	for _, v := range problems {
		var a map[string]string
		a = make(map[string]string)
		a["problem_id"] = strconv.Itoa(v.Problem_id)
		a["problem"] = v.Problem_content
		beego.Info("v=", v)
		//生成option
		if (v.Problem_type == 1 || v.Problem_type == 2) {
			//乱序

			//设置题目选项,数组[{"q_id":"1","content":"选项A"},{"q_id":"2","content":"选项B"},{"q_id":"3","content":"选项C"}]
			var problem_option = v.Problem_option
			var shuffled_option map[string]interface{}
			shuffled_option = make(map[string]interface{})
			if problem_option != "" {
				var f interface{}
				_ = json.Unmarshal([]byte(problem_option), &f)
				option, _ := f.([]interface{})
				option_num := len(option)
				for i := 0; i < option_num; i++ {
					var tmp = string(65 + i)
					shuffled_option[tmp] = option[i]
				}
			}
			str, err2 := json.Marshal(shuffled_option)
			if err2 != nil {
				fmt.Println(err2)
			}
			a["option"] = string(str)
		}

		//将题目加入对应数组中
		switch v.Problem_type {
		case 0:
			fill = append(fill, a)
		case 1:
			single = append(single, a)
		case 2:
			mutiple = append(mutiple, a)
		case 3:
			judge = append(judge, a)
		}
		beego.Info("=========GeneratingFrontProblems mutiple=========", mutiple)
	}
	//总题目
	var result map[string]interface{}
	result = make(map[string]interface{})
	result["fill"] = fill
	result["single"] = single
	result["multiple"] = mutiple
	result["judge"] = judge
	beego.Info("=========GeneratingFrontProblems result=========", result)

	return result
}

func GeneratingWaitedAnswer(problems []problem.Problem, answer_date string) map[string]interface{} {
	var waited_answer map[string]interface{}
	waited_answer = make(map[string]interface{})

	var single_answer map[string]string
	single_answer = make(map[string]string)
	var multiple_answer map[string]interface{}
	multiple_answer = make(map[string]interface{})
	var fill_answer map[string]string
	fill_answer = make(map[string]string)
	var judge_answer map[string]string
	judge_answer = make(map[string]string)
	for _, v := range problems {
		problem_id := strconv.Itoa(v.Problem_id)
		problem_answer := v.Problem_answer
		if (v.Problem_type == 1) {
			//解析json，取出q_id
			var f map[string]string
			_ = json.Unmarshal([]byte(v.Problem_answer), &f)
			single_answer[problem_id] = f["q_id"]
		} else if (v.Problem_type == 2) {
			//返回q_id数组
			var f interface{}
			_ = json.Unmarshal([]byte(v.Problem_answer), &f)
			a := f.([]interface{})
			var b []int
			for _, v := range a {
				c := v.(map[string]interface{})
				s := c["q_id"].(string)
				i, _ := strconv.Atoi(s)
				b = append(b, i)
			}
			multiple_answer[problem_id] = b
		} else if (v.Problem_type == 3) {
			judge_answer[problem_id] = problem_answer
		} else if (v.Problem_type == 0) {
			fill_answer[problem_id] = problem_answer
		}
	}
	waited_answer["single"] = single_answer
	waited_answer["multi"] = multiple_answer
	waited_answer["judge"] = judge_answer
	waited_answer["fill"] = fill_answer
	waited_answer["participant_time"] = answer_date
	beego.Info("========(((((((((((waited_answer))))))))======", waited_answer)

	return waited_answer
}
