package main

import (
	"context"

	"github.com/astaxie/beego/logs"

	"service/answer/model"
	"service/common"
	proto "service/protoc/answerManage"
)

type CreditManage struct{}

func (this *CreditManage) GetTeamCredit(ctx context.Context, req *proto.TeamEventIdReq, rsp *proto.CreditRsp) error {
	teamId := req.TeamId
	eventId := req.EventId

	t := model.GetTeamById(teamId, eventId)
	rsp.Credit = t.Team_credit

	return nil
}

func (this *CreditManage) GetPersonCredit(ctx context.Context, req *proto.UserEventIdReq, rsp *proto.CreditRsp) error {
	userId := req.UserId
	eventId := req.EventId

	t := model.GetParticipantById(userId, eventId)
	rsp.Credit = t.Credit

	return nil
}

func (this *CreditManage) GetCreditLogByTeamId(ctx context.Context, req *proto.TeamIdReq, rsp *proto.CreditLogListRsp) error {
	teamId := req.TeamId

	list := model.GetCreditLogByTeamId(teamId)
	//类型转换
	var pMessage []*proto.CreditLog
	for _, v := range list {
		log := proto.CreditLog{CreditLogId: v.Credit_log_id, TeamId: v.Refer_team_id, ParticipantId: v.Refer_participant_id,
			EventId: v.Refer_event_id, ChangeReason: v.Change_reason, ChangeTime: v.Change_time, ChangeValue: v.Change_value, ChangeType: v.Change_type}
		pMessage = append(pMessage, &log)
	}
	rsp.CreditLogList = pMessage

	return nil
}

func (this *CreditManage) WhetherMemberAllRight(ctx context.Context, req *proto.AllRightReq, rsp *proto.AllRightRsp) error {
	teamId := req.TeamId
	participantNum := req.ParticipantNum
	nowDate := req.NowDate

	t := model.WhetherMemberAllRight(teamId, nowDate, int(participantNum))
	rsp.AllRightFlag = t

	return nil
}

func (this *CreditManage) UpdateTeamCredit(ctx context.Context, req *proto.UpdateTeamCreditReq, rsp *proto.CreditRsp) error {
	teamId := req.TeamId
	changeCredit := req.ChangeCredit

	t := model.UpdateTeamCredit(teamId, changeCredit)
	rsp.Credit = t

	return nil
}

func (this *CreditManage) UpdateParticipantCredit(ctx context.Context, req *proto.UpdatePCreditReq, rsp *proto.CreditRsp) error {
	paticipantId := req.PaticipantId
	changeCredit := req.ChangeCredit

	t := model.UpdateParticipantCredit(paticipantId, changeCredit)
	rsp.Credit = t

	return nil
}

func (this *CreditManage) AddCreditLog(ctx context.Context, req *proto.CreditLog, rsp *proto.AddCreditLogRsp) error {
	log := model.Credit_log{Refer_event_id: req.EventId, Refer_participant_id: req.ParticipantId,
		Refer_team_id: req.TeamId, Change_time: req.ChangeTime, Change_value: req.ChangeValue, Change_type: req.ChangeType, Change_reason: req.ChangeReason}

	result, id := model.AddCreditLog(log)

	rsp.Message = result
	rsp.CreditLogId = int64(id)
	return nil
}

func main() {
	//初始化
	service,err := common.Init("CreditManage")
	if err != nil {
		panic(err)
	}

	//注册服务
	proto.RegisterCreditManageHandler(service.Server(), new(CreditManage))

	//运行
	if err := service.Run(); err != nil {
		logs.Error("failed-to-do-somthing", err)
	}
}
