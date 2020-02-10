package model

import (
	"context"
	"encoding/json"
	"github.com/astaxie/beego/logs"
	"strconv"
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/micro/go-micro"

	participantProto "service/protoc/answerManage"
)

type ProblemNum struct {
	Single   int32 `json:"single,string"`   //单选题数量
	Multiple int32 `json:"multiple,string"` //多选题数量
	Fill     int32 `json:"fill,string"`     //填空题数量
	Judge    int32 `json:"judge,string"`    //判断题数量
}

type Problem struct {
	Problem_id      int64 `orm:"pk"`
	Problem_content string
	Problem_option  string
	Problem_answer  string
	Problem_class   string
	Problem_type    int
}

func init() {
	orm.RegisterModel(new(Problem)) // 注册模型，建立User类型对象，注册模型时，需要引入包
}

func initParticipantManage() participantProto.ParticipantManageService {
	//调用服务
	service := micro.NewService(micro.Name("ParticipantManage.client"))
	service.Init()

	//create new client
	return participantProto.NewParticipantManageService("ParticipantManage", service.Client())
}

func GetProblemNoAnswer(user_id int64, event_id int64, team_id int64, participant_id int64, problemNum ProblemNum) ([]Problem, bool, bool) {
	var problems []Problem
	buildFlag := false  //是否已经生成过题目
	answerFlag := false //是否已经答题
	o := orm.NewOrm()
	now_time := time.Now()
	unix_time := now_time.Unix()
	now := time.Unix(unix_time, 0).Format("2006-01-02") //设置时间戳 使用模板格式化为日期字符串
	answer_date := time.Unix(unix_time, 0).Format("2006-01-02 15:04:05")

	//*****************************1.检查是否已完成答题*************************************************
	pManage := initParticipantManage()
	judgeReq := participantProto.JudgeReq{ParticipantId: participant_id}
	judgeRsp, judgeErr := pManage.JudgeIfHaveAnswer(context.TODO(), &judgeReq)
	if judgeErr != nil {
		logs.Error("judgeErr:", judgeErr)
	}
	answerFlag = judgeRsp.AnswerFlag
	logs.Debug("JudgeIfHaveAnswer:", answerFlag)
	if answerFlag == true {
		return nil, true, answerFlag
	}

	//*****************************2.检查是否已经生成题目，若已经生成，直接查询返回*************************
	_, err := o.Raw("SELECT problem.* "+
		"FROM problem, participant_haved_answer "+
		"WHERE problem.problem_id = participant_haved_answer.refer_problem_id "+
		"AND participant_haved_answer.refer_participant_id = ? "+
		"AND participant_haved_answer.answer_date = ? ", participant_id, now).QueryRows(&problems)
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
		pReq := participantProto.UpdateWaitedAnswerReq{ParticipantId: participant_id, WaitedAnswer: string(str)}
		_, pErr := pManage.UpdateParticipantWaitedAnswer(context.TODO(), &pReq)
		if pErr != nil {
			logs.Error("GetProblemNoAnswer:", pErr)
		}

		//*****************************5.将新题目拆入participant_haved_answer表*********
		for _, v := range problems {
			addProReq := participantProto.AddProblemHavedAnswerReq{ParticipantId: participant_id, ProblemId: v.Problem_id, TeamId: team_id, AnswerDate: answer_date}
			_, addProErr := pManage.AddProblemHavedAnswer(context.TODO(), &addProReq)
			if addProErr != nil {
				logs.Error("GetProblemNoAnswer:", addProErr)
			}

		}
	} else {
		buildFlag = true
	}

	return problems, buildFlag, answerFlag
}

func GeneratingProblems(event_id int64, participant_id int64, problem_type int, problem_num int32) []Problem {
	if (problem_num <= 0) {
		return nil
	}

	o := orm.NewOrm()
	var problems []Problem
	_, err := o.Raw("SELECT problem.* "+
		"FROM problem, event_problem "+
		"WHERE problem.problem_id = event_problem.problem_id "+
		"AND event_problem.refer_event_id = ? "+
		"AND problem.problem_type = ? "+
		"AND problem.problem_id NOT IN "+
		"(SELECT refer_problem_id FROM participant_haved_answer WHERE refer_participant_id = ?) LIMIT ?", event_id, problem_type, participant_id, problem_num).QueryRows(&problems)
	//增加随机！！
	if err == nil {
		return problems
	} else {
		logs.Error("GeneratingProblems's err:", err)
		return nil
	}

}

func GeneratingWaitedAnswer(problems []Problem, answer_date string) map[string]interface{} {
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
		problem_id := strconv.Itoa(int(v.Problem_id))
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

	return waited_answer
}
