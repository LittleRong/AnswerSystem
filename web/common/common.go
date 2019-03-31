package common

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry/consul"

	creditProto "service/protoc/answerManage"
	participantProto "service/protoc/answerManage"
	eventProto "service/protoc/eventManage"
	problemProto "service/protoc/problemManage"
	unionProto "service/protoc/unionManage"
	userProto "service/protoc/userManage"
)

func InitUserManage() userProto.UserManageService {
	service := micro.NewService(micro.Name("UserManage.client"), micro.Registry(consul.NewRegistry()))
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
