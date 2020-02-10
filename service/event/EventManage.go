package main

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/astaxie/beego/logs"
	_ "github.com/go-sql-driver/mysql"
	"service/common"
	"service/event/model"
	proto "service/protoc/eventManage"
)

type EventManage struct{}

func (this *EventManage) GetEventListByManageIdAndOffst(ctx context.Context, req *proto.GetEventListReq, rsp *proto.EventListRsp) error {
	manageId := req.ManageId
	offset := int(req.Offset)
	limit := int(req.Limit)
	eventList := model.GetEventListByManageIdAndOffst(manageId, offset, limit)
	//类型转换
	var eventMessage []*proto.EventMesssage
	for _, v := range eventList {
		u := proto.EventMesssage{EventId: int64(v.Event_id), EventTitle: v.Event_title, EventDescription: v.Event_description, EventType: v.Event_type}
		eventMessage = append(eventMessage, &u)
	}
	rsp.EventList = eventMessage

	return nil
}

func (this *EventManage) GetEventByEventId(ctx context.Context, req *proto.EventIdReq, rsp *proto.EventShowMesssage) error {
	eventId := req.EventId

	event := model.GetEventByEventId(eventId)
	rsp.EventId = event.Event_id
	rsp.EventTitle = event.Event_title
	rsp.EventDescription = event.Event_description
	rsp.ParticipantNum = int32(event.Participant_num)

	event_time := event.Event_time
	var event_time_map map[string]interface{}
	//使用 json.Unmarshal(data []byte, v interface{})进行转换,返回 error 信息
	if err := json.Unmarshal([]byte(event_time), &event_time_map); err != nil {
		return err
	}
	rsp.StartTime = event_time_map["start_time"].(string)
	rsp.EndTime = event_time_map["end_time"].(string)

	event_num := event.Event_num
	var event_num_map map[string]interface{}
	//使用 json.Unmarshal(data []byte, v interface{})进行转换,返回 error 信息
	if err := json.Unmarshal([]byte(event_num), &event_num_map); err != nil {
		return err
	}
	rsp.Single = event_num_map["single"].(string)
	rsp.Fill = event_num_map["fill"].(string)
	rsp.Judge = event_num_map["judge"].(string)
	rsp.Multiple = event_num_map["multiple"].(string)
	rsp.AnswerTime, _ = strconv.ParseFloat(event.Answer_time, 64)

	return nil
}

func (this *EventManage) GetCreditRuleByEventId(ctx context.Context, req *proto.EventIdReq, rsp *proto.CreditRule) error {
	eventId := req.EventId

	event := model.GetEventByEventId(eventId)
	var credit_rule model.CreditRule
	if err := json.Unmarshal([]byte(event.Credit_rule), &credit_rule); err != nil {
		return err
	}

	rsp.SingleScore = credit_rule.Single_score
	rsp.MultipleScore = credit_rule.Multi_score
	rsp.JudgeScore = credit_rule.Judge_score
	rsp.FillScore = credit_rule.Fill_score
	rsp.TeamScore = credit_rule.Team_score
	rsp.TeamScoreUp = credit_rule.Team_score_up
	rsp.PersonScore = credit_rule.Person_score
	rsp.PersonScoreUp = credit_rule.Person_score_up

	return nil
}

func (this *EventManage) GetProblemNumByEventId(ctx context.Context, req *proto.EventIdReq, rsp *proto.ProblemNum) error {
	eventId := req.EventId

	event := model.GetEventByEventId(eventId)
	event_num := event.Event_num
	var event_num_map map[string]interface{}
	//使用 json.Unmarshal(data []byte, v interface{})进行转换,返回 error 信息
	if err := json.Unmarshal([]byte(event_num), &event_num_map); err != nil {
		return err
	}
	var val int32
	StrToInt(event_num_map["single"].(string), &val)
	rsp.Single = val
	StrToInt(event_num_map["fill"].(string), &val)
	rsp.Fill = val
	StrToInt(event_num_map["judge"].(string), &val)
	rsp.Judge = val
	StrToInt(event_num_map["multiple"].(string), &val)
	rsp.Multiple = val

	return nil
}

func StrToInt(strNumber string, value interface{}) (err error) {
	var number interface{}
	number, err = strconv.ParseInt(strNumber, 10, 64)
	switch v := number.(type) {
	case int64:
		switch d := value.(type) {
		case *int64:
			*d = v
		case *int:
			*d = int(v)
		case *int16:
			*d = int16(v)
		case *int32:
			*d = int32(v)
		case *int8:
			*d = int8(v)
		}
	}
	return
}

func (this *EventManage) GetDetailEventByEventId(ctx context.Context, req *proto.EventIdReq, rsp *proto.EventDetailMesssage) error {
	eventId := req.EventId

	event := model.GetEventByEventId(eventId)
	rsp.EventId = event.Event_id
	rsp.EventTitle = event.Event_title
	rsp.EventDescription = event.Event_description
	rsp.ParticipantNum = int32(event.Participant_num)

	event_time := event.Event_time
	var event_time_map map[string]interface{}
	//使用 json.Unmarshal(data []byte, v interface{})进行转换,返回 error 信息
	if err := json.Unmarshal([]byte(event_time), &event_time_map); err != nil {
		return err
	}
	rsp.StartTime = event_time_map["start_time"].(string)
	rsp.EndTime = event_time_map["end_time"].(string)
	rsp.AnswerDay = event_time_map["answer_day"].(string)

	var credit_rule model.CreditRule
	if err := json.Unmarshal([]byte(event.Credit_rule), &credit_rule); err != nil {
		return err
	}

	rsp.SingleScore = credit_rule.Single_score
	rsp.MultipleScore = credit_rule.Multi_score
	rsp.JudgeScore = credit_rule.Judge_score
	rsp.FillScore = credit_rule.Fill_score
	rsp.TeamScore = credit_rule.Team_score
	rsp.TeamScoreUp = credit_rule.Team_score_up
	rsp.PersonScore = credit_rule.Person_score
	rsp.PersonScoreUp = credit_rule.Person_score_up

	event_num := event.Event_num
	var event_num_map map[string]interface{}
	//使用 json.Unmarshal(data []byte, v interface{})进行转换,返回 error 信息
	if err := json.Unmarshal([]byte(event_num), &event_num_map); err != nil {
		return err
	}
	rsp.Single = event_num_map["single"].(string)
	rsp.Fill = event_num_map["fill"].(string)
	rsp.Judge = event_num_map["judge"].(string)
	rsp.Multiple = event_num_map["multiple"].(string)

	return nil
}

func (this *EventManage) AddNewEvent(ctx context.Context, req *proto.AddEventReq, rsp *proto.AddEventRsp) error {
	var e model.Event
	e.Manage_id = int(req.ManageId)
	e.Event_title = req.EventTitle
	e.Event_description = req.EventDescription
	e.Event_time = req.EventTime
	e.Event_num = req.EventNum
	e.Event_type = req.EventType
	e.Problem_random = req.ProblemRandom
	e.Option_random = req.OptionRandom
	e.Answer_time = req.AnswerTime
	e.Credit_rule = req.CreditRule
	e.Participant_num = int(req.ParticipantNum)

	result, id := model.AddNewEvent(e)

	rsp.Message = result
	rsp.EventId = id

	return nil
}

func (this *EventManage) AddEventProblem(ctx context.Context, req *proto.AddEventProblemReq, rsp *proto.AddEventProblemRsp) error {
	var e model.EventProblem
	e.Refer_event_id = req.EventId
	e.Problem_id = req.ProblemId

	result := model.AddEventProblem(e)

	rsp.Message = result

	return nil
}

func main() {
	//初始化
	service,err := common.Init("EventManage")
	if err != nil {
		panic(err)
	}

	//注册服务
	proto.RegisterEventManageHandler(service.Server(), new(EventManage))

	//运行
	if err := service.Run(); err != nil {
		logs.Error("failed-to-do-somthing", err)
	}
}
