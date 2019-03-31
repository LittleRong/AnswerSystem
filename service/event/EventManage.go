package main

import (
	"context"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/micro/go-micro"
	"service/event/model"
	proto "service/protoc/eventManage" //proto文件放置路径
	"strconv"
	"github.com/micro/go-micro/registry/consul"
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

func (this *EventManage) GetEventByEventId(ctx context.Context, req *proto.EventIdReq, rsp *proto.EventShowMesssage) error{
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
	rsp.Fill= event_num_map["fill"].(string)
	rsp.Judge = event_num_map["judge"].(string)
	rsp.Multiple = event_num_map["multiple"].(string)
	rsp.AnswerTime,_ = strconv.ParseFloat(event.Answer_time,64)

	beego.Info("======UserIndex rsp=====", rsp)

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

	beego.Info("======GetCreditRuleByEventId rsp=====", rsp)

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
	rsp.Single =  val
	StrToInt(event_num_map["fill"].(string), &val)
	rsp.Fill =  val
	StrToInt(event_num_map["judge"].(string), &val)
	rsp.Judge =  val
	StrToInt(event_num_map["multiple"].(string), &val)
	rsp.Multiple =  val

	beego.Info("======GetProblemNumByEventId rsp=====", rsp)

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

func (this *EventManage) GetDetailEventByEventId(ctx context.Context, req *proto.EventIdReq, rsp *proto.EventDetailMesssage) error{
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
	rsp.Fill= event_num_map["fill"].(string)
	rsp.Judge = event_num_map["judge"].(string)
	rsp.Multiple = event_num_map["multiple"].(string)

	beego.Info("======UserIndex rsp=====", rsp)

	return nil
}

func (this *EventManage) AddNewEvent(ctx context.Context, req *proto.AddEventReq, rsp *proto.AddEventRsp) error{
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

	result,id := model.AddNewEvent(e)

	rsp.Message = result
	rsp.EventId = int64(id)

	return nil
}

func main(){

	// 开启 orm 调试模式：开发过程中建议打开，release时需要关闭
	orm.Debug = true
	// 自动建表
	orm.RunSyncdb("default", false, true)

	//create service
	service := micro.NewService(micro.Name("EventManage"),micro.Registry(consul.NewRegistry()))

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
	orm.RegisterDataBase("default", "mysql", "root:password123@tcp(localhost:3306)/problem?charset=utf8")
}