package union

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"hello/models/participant"
	"strconv"
)

// 完成User类型定义
type Problem struct {
	Problem_id int `orm:"pk"`
	Problem_content string
	Problem_option string
	Problem_answer string
	Problem_class string
	Problem_type int
}

func GetProblemNoAnswer(user_id int,event_id int,now string) (map[string]interface{},bool){
	var problems []Problem
	var participant_id int
	buildFlag := false//是否已经生成过题目
	o := orm.NewOrm()

	//获取用户的participant_id
	u := participant.Participant{Refer_event_id:event_id ,User_id:user_id}
	err := o.Read(&u,"Refer_event_id","User_id")
	if err == nil {
		participant_id = u.Participant_id
	} else {

	}
	beego.Info("event_id=", event_id,"participant_id=", participant_id)

	//检查是否已经生成题目，若已经生成，直接查询返回
	//AND participant_haved_answer.answer_date = ?
	_, err = o.Raw("SELECT problem.* " +
		"FROM problem, participant_haved_answer " +
		"WHERE problem.problem_id = participant_haved_answer.refer_problem_id " +
		"AND participant_haved_answer.refer_participant_id = ? " +
		"AND participant_haved_answer.answer_date = ? ", participant_id, now).QueryRows(&problems)
	beego.Info("problems", problems)
	if problems == nil && err ==nil {
		buildFlag = false
		_, err := o.Raw("SELECT * " +
			"FROM problem, event_problem" +
			"WHERE problem.problem_id = event_problem.problem_id " +
			"AND event_problem.refer_event_id = ?" +
			"AND problem.problem_id NOT IN" +
			"(SELECT refer_problem_id FROM participant_haved_answer WHERE refer_participant_id = ?)",event_id,participant_id).QueryRows(&problems)
		if err == nil {
			//将新题目拆入participant_haved_answer表

		}
	} else {
		buildFlag = true
	}

	var single []map[string]string
	for _,v := range problems {
		var a map[string]string
		a = make(map[string]string)
		a["problem_id"] = strconv.Itoa(v.Problem_id)
		a["problem"] = v.Problem_content
		beego.Info("v=",v)
		//乱序

		//设置题目选项,数组[{"q_id":"1","content":"选项A"},{"q_id":"2","content":"选项B"},{"q_id":"3","content":"选项C"}]
		var problem_option =v.Problem_option

		var shuffled_option map[string]interface{}
		shuffled_option = make(map[string]interface{})
		if problem_option != "" {
			beego.Info("problem_option=",problem_option)
			var f interface{}
			_ = json.Unmarshal([]byte(problem_option), &f)

			option,err3 := f.([]interface{})
			option_num := len(option)
			beego.Info("err3=",err3)
			beego.Info("f=",f)
			beego.Info("option=",option)

			for i:=0; i<option_num; i++ {
				var tmp = strconv.Itoa(65+i)
				shuffled_option[tmp] = option[i]
			}
		}
		beego.Info("shuffled_option=",shuffled_option)
		str, err2 := json.Marshal(shuffled_option)
		if err2 != nil {
			fmt.Println(err2)
		}
		beego.Info("str=",str)
		a["option"] = string(str)
		beego.Info("a=",a)
		single = append(single, a)
	}


	//总题目
	var result map[string]interface{}
	result = make(map[string]interface{})
	result["single"] = single
	beego.Info("result", result)

	return result,buildFlag
}
