package common

import (
	"fmt"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"
	creditProto "service/protoc/answerManage"
	participantProto "service/protoc/answerManage"
	eventProto "service/protoc/eventManage"
	problemProto "service/protoc/problemManage"
	unionProto "service/protoc/unionManage"
	userProto "service/protoc/userManage"
)

func InitUserManage() userProto.UserManageService {
	// 修改consul地址，如果是本机，这段代码和后面的那行使用代码都是可以不用的
	reg := consul.NewRegistry(func(op *registry.Options){
		//add := os.Getenv("CONSUL_PORT_8500_TCP_ADDR")+":8500"
		add := "127.0.0.1:8500"
		op.Addrs = []string{
			add,
		}
		fmt.Println("sjdsghjadgsahjdg"+add)
	})
	service := micro.NewService(micro.Name("UserManage.client"), micro.Registry(reg))
	service.Init()
	return userProto.NewUserManageService("UserManage", service.Client())
}

func InitEventManage() eventProto.EventManageService {
	service := micro.NewService(micro.Name("EventManage.client"), micro.Registry(consul.NewRegistry()))
	service.Init()
	return eventProto.NewEventManageService("EventManage", service.Client())
}

func InitUniontManage() unionProto.UnionManageService {
	service := micro.NewService(micro.Name("UnionManage.client"), micro.Registry(consul.NewRegistry()))
	service.Init()
	return unionProto.NewUnionManageService("UnionManage", service.Client())
}

func InitParticipantManage() participantProto.ParticipantManageService {
	service := micro.NewService(micro.Name("ParticipantManage.client"), micro.Registry(consul.NewRegistry()))
	service.Init()
	return participantProto.NewParticipantManageService("ParticipantManage", service.Client())
}

func InitCreditManage() creditProto.CreditManageService {
	service := micro.NewService(micro.Name("CreditManage.client"), micro.Registry(consul.NewRegistry()))
	service.Init()
	return creditProto.NewCreditManageService("CreditManage", service.Client())
}

func InitProblemManage() problemProto.ProblemManageService {
	service := micro.NewService(micro.Name("ProblemManage.client"), micro.Registry(consul.NewRegistry()))
	service.Init()
	return problemProto.NewProblemManageService("ProblemManage", service.Client())
}
