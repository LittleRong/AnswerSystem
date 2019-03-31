package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"service/problem/model"
	"github.com/micro/go-micro/registry/consul"
	_ "github.com/go-sql-driver/mysql"
	"context"
	micro "github.com/micro/go-micro"
	proto "service/protoc/problemManage" //proto文件放置路径
)

type ProblemManage struct{}

func (this *ProblemManage) GetProblemListByOffstAndLimit(ctx context.Context, req *proto.GetProblemListReq, rsp *proto.ProblemListRsp) error{
	offset := 0
	limit :=10
	problemList := model.GetProblemListByOffstAndLimit(offset,limit)
	beego.Info("========GetProblemListByOffstAndLimit000===========",problemList)
	//类型转换
	var problemMessage []*proto.ProblemMesssage
	for _,v := range problemList {
		u := proto.ProblemMesssage{ProblemId:int64(v.Problem_id),ProblemContent:v.Problem_content,ProblemOption:v.Problem_option,ProblemAnswer:v.Problem_answer,ProblemClass:v.Problem_class,ProblemType:int32(v.Problem_type)}
		problemMessage = append(problemMessage,&u)
	}
	rsp.ProblemList = problemMessage

	return nil
}

func main(){

	// 开启 orm 调试模式：开发过程中建议打开，release时需要关闭
	orm.Debug = true
	// 自动建表
	orm.RunSyncdb("default", false, true)

	//create service
	service := micro.NewService(micro.Name("ProblemManage"),micro.Registry(consul.NewRegistry()))

	//init
	service.Init()

	//register handler
	proto.RegisterProblemManageHandler(service.Server(), new(ProblemManage))

	//run the server
	if err:=service.Run();err != nil {
		beego.Info("========ProblemManage's err===========",err)
	}
}

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:password123@tcp(localhost:3306)/problem?charset=utf8")
}