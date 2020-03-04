package main

import (
	"context"

	"github.com/astaxie/beego/logs"

	"service/answer/model"
	"service/common"
	proto "service/protoc/answerManage"
)

type ParticipantManage struct{}

func (this *ParticipantManage) EventParticipantInsert(ctx context.Context, req *proto.EPInsertReq, rsp *proto.EPInsertRsp) error {
	team_array := req.ParticipantMemberList
	eventId := req.EventId
	logs.Info("EventParticipantInsert",team_array)
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

func (this *ParticipantManage) GetParticipantListByUserId(ctx context.Context, req *proto.GetPListByUserIdReq, rsp *proto.PEMessageList) error {
	userId := req.UserId

	list := model.GetParticipantListByUserId(userId)
	//类型转换
	var pMessage []*proto.PEMessage
	for _, v := range list {
		u := proto.PEMessage{ParticipantId: int64(v.Participant_id), ReferEventId: int64(v.Refer_event_id)}
		pMessage = append(pMessage, &u)
	}
	rsp.PEList = pMessage

	return nil
}

func (this *ParticipantManage) GetMemberCreditByTeamId(ctx context.Context, req *proto.PTeamEventIdReq, rsp *proto.ParticipantMessageList) error {
	teamId := req.TeamId
	eventId := req.EventId

	list := model.GetMemberCreditByTeamId(teamId, eventId)
	//类型转换
	var pMessage []*proto.ParticipantMessage
	for _, v := range list {
		u := proto.ParticipantMessage{ParticipantId: v.Participant_id, EventId: v.Refer_event_id,
			TeamId: v.Team_id, UserId: v.User_id, Credit: v.Credit, Leader: v.Leader}
		pMessage = append(pMessage, &u)
	}
	rsp.PEList = pMessage

	return nil
}

func (this *ParticipantManage) GetParticipantByUserAndEvent(ctx context.Context, req *proto.PUserEventIdReq, rsp *proto.ParticipantMessage) error {
	userId := req.UserId
	eventId := req.EventId

	event := model.GetParticipantById(userId, eventId)

	rsp.EventId = event.Refer_event_id
	rsp.UserId = event.User_id
	rsp.Credit = event.Credit
	rsp.ParticipantId = event.Participant_id
	rsp.Leader = event.Leader
	rsp.TeamId = event.Team_id

	return nil
}

func (this *ParticipantManage) GetCorrectAnswerByParticipantId(ctx context.Context, req *proto.ParticipantIdReq, rsp *proto.WaitAnswerRsp) error {
	participantId := req.ParticipantId
	var correctAnswer map[string]interface{}
	correctAnswer = make(map[string]interface{})
	correctAnswer = model.GetCorrectAnswerByParticipantId(participantId)

	//单选题
	singleAnswer := correctAnswer["single"].(map[string]interface{})
	var singleAnswerList []*proto.NolAnswer
	for v := range singleAnswer {
		problemId := v
		answer := singleAnswer[v].(string)
		single := proto.NolAnswer{ProblemId: problemId, Answer: answer}
		singleAnswerList = append(singleAnswerList, &single)
	}

	//填空题
	fillAnswer := correctAnswer["fill"].(map[string]interface{})
	var fillAnswerList []*proto.NolAnswer
	for v := range fillAnswer {
		problemId := v
		answer := fillAnswer[v].(string)
		fill := proto.NolAnswer{ProblemId: problemId, Answer: answer}
		fillAnswerList = append(fillAnswerList, &fill)
	}

	//多选题
	multiAnswer := correctAnswer["multi"].(map[string]interface{})
	var multiAnswerList []*proto.MultiAnswer
	for v := range multiAnswer {
		problemId := v
		answer := multiAnswer[v].([]interface{})
		var answerArr []float64
		for _, s := range answer {
			answerArr = append(answerArr, s.(float64))
		}
		multi := proto.MultiAnswer{ProblemId: problemId, Answer: answerArr}
		multiAnswerList = append(multiAnswerList, &multi)
	}

	//判断题
	judgeAnswer := correctAnswer["judge"].(map[string]interface{})
	var judgeAnswerList []*proto.NolAnswer
	for v := range judgeAnswer {
		problemId := v
		answer := judgeAnswer[v].(string)
		judge := proto.NolAnswer{ProblemId: problemId, Answer: answer}
		judgeAnswerList = append(judgeAnswerList, &judge)
	}

	rsp.FillAnswerList = fillAnswerList
	rsp.MultiAnswerList = multiAnswerList
	rsp.JudgeAnswerList = judgeAnswerList
	rsp.SingleAnswerList = singleAnswerList
	rsp.ParticipantTime = correctAnswer["participant_time"].(string)

	return nil
}

func (this *ParticipantManage) GetAnswerTimeByParticipantId(ctx context.Context, req *proto.ParticipantIdReq, rsp *proto.AnswerTimeReq) error {
	participantId := req.ParticipantId

	answerTime := model.GetAnswerTimeByParticipantId(participantId)
	rsp.AnswerTime = answerTime

	return nil
}

func (this *ParticipantManage) UpdateUserAnswer(ctx context.Context, req *proto.UpdateUserAnswerReq, rsp *proto.UpdateUserAnswerRsp) error {
	participantId := req.ParticipantId
	problemId := req.ProblemId
	judge := req.TrueOrFalse
	userAnswer := req.UserAnswer

	result := model.UpdateUserAnswer(participantId, problemId, userAnswer, judge)
	rsp.Result = result

	return nil
}

func (this *ParticipantManage) UpdateParticipantWaitedAnswer(ctx context.Context, req *proto.UpdateWaitedAnswerReq, rsp *proto.UpdateWaitedAnswerResp) error {
	participantId := req.ParticipantId
	waitedAnswer := req.WaitedAnswer

	result := model.UpdateParticipantWaitedAnswer(participantId, waitedAnswer)
	rsp.Message = result

	return nil
}

func (this *ParticipantManage) JudgeIfHaveAnswer(ctx context.Context, req *proto.JudgeReq, rsp *proto.JudgeRsp) error {
	participantId := req.ParticipantId

	flag := model.JudgeIfHaveAnswer(participantId)
	rsp.AnswerFlag = flag

	return nil
}

func (this *ParticipantManage) AddProblemHavedAnswer(ctx context.Context, req *proto.AddProblemHavedAnswerReq, rsp *proto.AddProblemHavedAnswerRsp) error {
	participantId := req.ParticipantId
	problemId := req.ProblemId
	teamId := req.TeamId
	answerDate := req.AnswerDate

	result := model.AddProblems(model.Participant_haved_answer{Refer_participant_id: participantId, Refer_team_id: teamId, Refer_problem_id: problemId, Answer_date: answerDate})
	rsp.Message = result

	return nil
}

func main() {
	//初始化
	service,err := common.Init("ParticipantManage")
	if err != nil {
		panic(err)
	}

	//注册服务
	proto.RegisterParticipantManageHandler(service.Server(), new(ParticipantManage))

	//运行
	if err := service.Run(); err != nil {
		logs.Error("failed-to-do-somthing", err)
	}
}
