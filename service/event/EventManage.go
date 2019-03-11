package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"service/event/model"

	_ "github.com/go-sql-driver/mysql"
	"context"
	micro "github.com/micro/go-micro"
	proto "service/protoc/eventManage" //proto文件放置路径
)

type EventManage struct{}

func (this *EventManage) GetEventListByManageIdAndOffst(ctx context.Context, req *proto.GetEventListReq, rsp *proto.EventListRsp) error{
	manageId := req.ManageId
	offset := int(req.Offset)
	limit := int(req.Limit)
	eventList := model.GetEventListByManageIdAndOffst(manageId,offset,limit)
	beego.Info("========GetEventListByOffstAndLimit000===========",eventList)
	//类型转换
	var eventMessage []*proto.EventMesssage
	for _,v := range eventList {
		u := proto.EventMesssage{EventId:int64(v.Event_id),EventTitle:v.Event_title,EventDescription:v.Event_description,EventType:v.Event_type}
		eventMessage = append(eventMessage,&u)
	}
	rsp.EventList = eventMessage

	return nil
}

func main(){

	// 开启 orm 调试模式：开发过程中建议打开，release时需要关闭
	orm.Debug = true
	// 自动建表
	orm.RunSyncdb("default", false, true)

	//create service
	service := micro.NewService(micro.Name("EventManage"))

	//init
	service.Init()

	//register handler
	proto.RegisterEventManageHandler(service.Server(), new(EventManage))

	//run the server
	if err:=service.Run();err != nil {
		beego.Info("========EventManage's err===========",err)
	}
}

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:ganxiaorong0703@tcp(localhost:3306)/problem?charset=utf8")
}