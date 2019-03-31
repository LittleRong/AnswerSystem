package common

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry/consul"

	userProto "service/protoc/userManage"
	creditProto "service/protoc/answerManage"
	participantProto "service/protoc/answerManage"
	eventProto "service/protoc/eventManage"
	unionProto "service/protoc/unionManage"
	problemProto "service/protoc/problemManage"
)

func InitUserManage() userProto.UserManageService{
	//调用服务
	service := micro.NewService(micro.Name("UserManage.client"),micro.Registry(consul.NewRegistry()))
	service.Init()

	//create new client
	return userProto.NewUserManageService("UserManage",service.Client())
}

func InitEventManage() eventProto.EventManageService{
	//调用服务
	service := micro.NewService(micro.Name("EventManage.client"),micro.Registry(consul.NewRegistry()))
	service.Init()

	//create new client
	return eventProto.NewEventManageService("EventManage",service.Client())
}

func InitUniontManage() unionProto.UnionManageService{
	//调用服务
	service := micro.NewService(micro.Name("UnionManage.client"),micro.Registry(consul.NewRegistry()))
	service.Init()

	//create new client
	return unionProto.NewUnionManageService("UnionManage",service.Client())
}

func InitParticipantManage() participantProto.ParticipantManageService{
	//调用服务
	service := micro.NewService(micro.Name("ParticipantManage.client"),micro.Registry(consul.NewRegistry()))
	service.Init()

	//create new client
	return participantProto.NewParticipantManageService("ParticipantManage",service.Client())
}

func InitCreditManage() creditProto.CreditManageService{
	//调用服务
	service := micro.NewService(micro.Name("CreditManage.client"),micro.Registry(consul.NewRegistry()))
	service.Init()

	//create new client
	return creditProto.NewCreditManageService("CreditManage",service.Client())
}

func InitProblemManage() problemProto.ProblemManageService{
	//调用服务
	service := micro.NewService(micro.Name("ProblemManage.client"),micro.Registry(consul.NewRegistry()))
	service.Init()

	//create new client
	return problemProto.NewProblemManageService("ProblemManage",service.Client())
}