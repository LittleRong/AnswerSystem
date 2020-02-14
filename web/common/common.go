package common

import (
	"fmt"
	"context"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/session"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"
	"github.com/micro/go-micro/metadata"
	creditProto "service/protoc/answerManage"
	participantProto "service/protoc/answerManage"
	eventProto "service/protoc/eventManage"
	problemProto "service/protoc/problemManage"
	unionProto "service/protoc/unionManage"
	userProto "service/protoc/userManage"
)

type CommonController struct {
	beego.Controller
}

type tokenWrapper struct {
	client.Client
}

func (l *tokenWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
	fmt.Printf("[wrapper] client request to service: %s method: %s\n", req.Service(), req.Endpoint())
	return l.Client.Call(ctx, req, rsp)
}

// 实现client.Wrapper，token包装器
func tokenWrap(c client.Client) client.Client {
	return &tokenWrapper{c}
}

func ServiceRegistryInit(s session.Store,serviceName string) (micro.Service,context.Context){

	//create service
	service := micro.NewService(micro.Name(serviceName),
		micro.Metadata(map[string]string{"type": "hello world"}),
		micro.Registry(consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			beego.AppConfig.String("consulhost")+":"+beego.AppConfig.String("consulport"),
		}
	})))

	//init
	service.Init()

	//设置JWT token
	token := ""
	//设置JWT token
	if s != nil {
		tokenSession := s.Get("token")
		if tokenSession != nil { //还没有token
			token = tokenSession.(string)
		}
	}
	md := metadata.Metadata{
		"Authorization": token,
	}
	ctx := metadata.NewContext(context.TODO(), md)


	return service,ctx
}

func InitUserManage(s session.Store) (userProto.UserManageService,context.Context) {
	service,ctx := ServiceRegistryInit(s,"UserManage.client")
	return userProto.NewUserManageService("UserManage", service.Client()),ctx
}

func InitEventManage(s session.Store) (eventProto.EventManageService,context.Context) {
	service,ctx := ServiceRegistryInit(s,"EventManage.client")
	return eventProto.NewEventManageService("EventManage", service.Client()),ctx
}

func InitUniontManage(s session.Store) (unionProto.UnionManageService,context.Context) {
	service,ctx := ServiceRegistryInit(s,"UnionManage.client")
	return unionProto.NewUnionManageService("UnionManage", service.Client()),ctx
}

func InitParticipantManage(s session.Store) (participantProto.ParticipantManageService,context.Context) {
	service,ctx := ServiceRegistryInit(s,"ParticipantManage.client")
	return participantProto.NewParticipantManageService("ParticipantManage", service.Client()),ctx
}

func InitCreditManage(s session.Store) (creditProto.CreditManageService,context.Context) {
	service,ctx := ServiceRegistryInit(s,"CreditManage.client")
	return creditProto.NewCreditManageService("CreditManage", service.Client()),ctx
}

func InitProblemManage(s session.Store) (problemProto.ProblemManageService,context.Context) {
	service,ctx := ServiceRegistryInit(s,"ProblemManage.client")
	return problemProto.NewProblemManageService("ProblemManage", service.Client()),ctx
}
