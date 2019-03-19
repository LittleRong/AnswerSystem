package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"service/answer/model"

	_ "github.com/go-sql-driver/mysql"
	"context"
	micro "github.com/micro/go-micro"
	proto "service/protoc/answerManage" //proto文件放置路径
)

type ParticipantManage struct{}

func (this *ParticipantManage) EventParticipantInsert(ctx context.Context, req *proto.EPInsertReq, rsp *proto.EPInsertRsp) error{
	team_array:= req.ParticipantMemberList
	eventId := req.EventId

	for _, value := range team_array {
		//插入新team
		teamId := model.AddTeam("", eventId)
		if (teamId != -1) {
			//插入participant
			model.AddParticipant(value.LeaderId, eventId, teamId, true)

			for _, member := range value.MemberId {
				model.AddParticipant(member, eventId, teamId, false)
			}
		}
	}

	//待完善，一个插入失败怎么办
	rsp.Message = "success"

	return nil
}

func (this *ParticipantManage) GetParticipantListByUserId (ctx context.Context, req *proto.GetPListByUserIdReq, rsp *proto.PEMessageList) error {
	userId := req.UserId

	list := model.GetParticipantListByUserId(userId)
	//类型转换
	var pMessage []*proto.PEMessage
	for _,v := range list {
		u := proto.PEMessage{ParticipantId:int64(v.Participant_id),ReferEventId:int64(v.Refer_event_id)}
		pMessage = append(pMessage,&u)
	}
	rsp.PEList = pMessage

	return nil
}

func (this *ParticipantManage) GetParticipantByUserAndEvent (ctx context.Context, req *proto.PUserEventIdReq, rsp *proto.ParticipantMessage) error {
	userId := req.UserId
	eventId := req.EventId

	event := model.GetParticipantById(userId,eventId)

	rsp.EventId = event.Refer_event_id
	rsp.UserId = event.User_id
	rsp.Credit = event.Credit
	rsp.ParticipantId = event.Participant_id
	rsp.Leader  = event.Leader
	rsp.TeamId = event.Team_id

	return nil
}

func main(){

	// 开启 orm 调试模式：开发过程中建议打开，release时需要关闭
	orm.Debug = true
	// 自动建表
	orm.RunSyncdb("default", false, true)

	//create service
	service := micro.NewService(micro.Name("ParticipantManage"))

	//init
	service.Init()

	//register handler
	proto.RegisterParticipantManageHandler(service.Server(), new(ParticipantManage))

	//run the server
	if err:=service.Run();err != nil {
		beego.Info("========ParticipantManage's err===========",err)
	}
}

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:ganxiaorong0703@tcp(localhost:3306)/problem?charset=utf8")
}