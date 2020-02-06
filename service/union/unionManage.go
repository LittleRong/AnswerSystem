package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"service/common"
	proto "service/protoc/unionManage"
	"service/union/model"
)

type UnionManage struct{}

func (this *UnionManage) GetProblemNoAnswer(ctx context.Context, req *proto.GetProblemNoAnswerReq, rsp *proto.GetProblemNoAnswerRsp) error {
	eventId := req.EventId
	teamId := req.TeamId
	paticipantId := req.PaticipantId
	problemNumRsp := req.ProblemNum
	userId := req.UserId

	problemNum := model.ProblemNum{Single: problemNumRsp.Single, Multiple: problemNumRsp.Multiple, Judge: problemNumRsp.Judge, Fill: problemNumRsp.Fill}
	result, buildFlag, answerFlag := model.GetProblemNoAnswer(userId, eventId, teamId, paticipantId, problemNum)

	rsp.AnswerFlag = answerFlag
	rsp.BuildFlag = buildFlag
	GeneratingFrontProblems(result, rsp)

	return nil
}

func GeneratingFrontProblems(problems []model.Problem, rsp *proto.GetProblemNoAnswerRsp) {
	var single []*proto.ProblemItem
	var mutiple []*proto.ProblemItem
	var fill []*proto.ProblemItem
	var judge []*proto.ProblemItem

	for _, v := range problems {
		var a proto.ProblemItem
		a.ProblemId = v.Problem_id
		a.Problem = v.Problem_content
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
			a.Option = string(str)
		}

		//将题目加入对应数组中
		switch v.Problem_type {
		case 0:
			fill = append(fill, &a)
		case 1:
			single = append(single, &a)
		case 2:
			mutiple = append(mutiple, &a)
		case 3:
			judge = append(judge, &a)
		}
		beego.Info("=========GeneratingFrontProblems mutiple=========", mutiple)
	}
	//总题目
	rsp.Fill = fill
	rsp.Single = single
	rsp.Multiple = mutiple
	rsp.Judge = judge
	beego.Info("=========GeneratingFrontProblems result=========", rsp)

}

func main() {
	//数据库初始化
	common.DatabaseInit()

	// 开启 orm 调试模式：开发过程中建议打开，release时需要关闭
	orm.Debug = true
	// 自动建表
	orm.RunSyncdb("default", false, true)

	//consul初始化
	service := common.ServiceRegistryInit("UnionManage")

	//register handler
	proto.RegisterUnionManageHandler(service.Server(), new(UnionManage))

	//run the server
	if err := service.Run(); err != nil {
		beego.Info("========ProblemManage's err===========", err)
	}
}
