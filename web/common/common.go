package common

import (
	"fmt"
	"github.com/astaxie/beego"
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


func ServiceRegistryInit(serviceName string) micro.Service{

	//create service
	service := micro.NewService(micro.Name(serviceName), micro.Registry(consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			beego.AppConfig.String("consulhost")+":"+beego.AppConfig.String("consulport"),
		}
	})))

	//init
	service.Init()
	return service
}

func InitUserManage() userProto.UserManageService {
	service := ServiceRegistryInit("UserManage.client")
	fmt.Println(service)
	fmt.Println("aaaaaaaaaaa")
	return userProto.NewUserManageService("UserManage", service.Client())
}

func InitEventManage() eventProto.EventManageService {
	service := ServiceRegistryInit("EventManage.client")
	return eventProto.NewEventManageService("EventManage", service.Client())
}

func InitUniontManage() unionProto.UnionManageService {
	service := ServiceRegistryInit("UnionManage.client")
	return unionProto.NewUnionManageService("UnionManage", service.Client())
}

func InitParticipantManage() participantProto.ParticipantManageService {
	service := ServiceRegistryInit("ParticipantManage.client")
	return participantProto.NewParticipantManageService("ParticipantManage", service.Client())
}

func InitCreditManage() creditProto.CreditManageService {
	service := ServiceRegistryInit("CreditManage.client")
	return creditProto.NewCreditManageService("CreditManage", service.Client())
}

func InitProblemManage() problemProto.ProblemManageService {
	service := ServiceRegistryInit("ProblemManage.client")
	return problemProto.NewProblemManageService("ProblemManage", service.Client())
}
