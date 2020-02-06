package main

import (
	"context"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"service/common"
	"service/problem/model"
	proto "service/protoc/problemManage"
)

type ProblemManage struct{}

func (this *ProblemManage) GetProblemListByOffstAndLimit(ctx context.Context, req *proto.GetProblemListReq, rsp *proto.ProblemListRsp) error {
	offset := 0
	limit := 10
	problemList := model.GetProblemListByOffstAndLimit(offset, limit)
	beego.Info("========GetProblemListByOffstAndLimit000===========", problemList)
	//类型转换
	var problemMessage []*proto.ProblemMesssage
	for _, v := range problemList {
		u := proto.ProblemMesssage{ProblemId: int64(v.Problem_id), ProblemContent: v.Problem_content, ProblemOption: v.Problem_option, ProblemAnswer: v.Problem_answer, ProblemClass: v.Problem_class, ProblemType: int32(v.Problem_type)}
		problemMessage = append(problemMessage, &u)
	}
	rsp.ProblemList = problemMessage

	return nil
}

func (this *ProblemManage) AddProblem(ctx context.Context, req *proto.ProblemMesssage, rsp *proto.AddProblemRsp) error {
	var p model.Problem
	p.Problem_option = req.ProblemOption
	p.Problem_type = req.ProblemType
	p.Problem_class = req.ProblemClass
	p.Problem_answer = req.ProblemAnswer
	p.Problem_content = req.ProblemContent

	id, result := model.AddProblem(p)
	rsp.Message = result
	rsp.ProblemId = id

	return nil
}

func (this *ProblemManage) GetNewProblemByType(ctx context.Context, req *proto.GetNewProblemByTypeReq, rsp *proto.ProblemListRsp) error {
	firstProblemId := req.FirstProblemId
	problemType := req.ProblemType

	problemList := model.GetNewProblemByType(firstProblemId, problemType)
	beego.Info("========GetProblemListByOffstAndLimit000===========", problemList)
	//类型转换
	var problemMessage []*proto.ProblemMesssage
	for _, v := range problemList {
		u := proto.ProblemMesssage{ProblemId: int64(v.Problem_id), ProblemContent: v.Problem_content, ProblemOption: v.Problem_option, ProblemAnswer: v.Problem_answer, ProblemClass: v.Problem_class, ProblemType: int32(v.Problem_type)}
		problemMessage = append(problemMessage, &u)
	}
	rsp.ProblemList = problemMessage
	return nil
}

func (this *ProblemManage) GetEndProblemId(ctx context.Context, req *proto.GetEndProblemIdReq, rsp *proto.GetEndProblemIdRsp) error {
	endId := model.GetEndProblemId()

	rsp.EndId = endId
	return nil
}

func main() {
	//数据库初始化
	common.DatabaseInit()

	// 开启 orm 调试模式：开发过程中建议打开，release时需要关闭
	orm.Debug = true
	// 自动建表
	orm.RunSyncdb("default", false, true)

	//consul初始化
	service := common.ServiceRegistryInit("ProblemManage")

	//register handler
	proto.RegisterProblemManageHandler(service.Server(), new(ProblemManage))

	//run the server
	if err := service.Run(); err != nil {
		beego.Info("========ProblemManage's err===========", err)
	}
}
