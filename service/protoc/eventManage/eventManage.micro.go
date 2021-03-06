// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: eventManage.proto

/*
Package eventManage is a generated protocol buffer package.

It is generated from these files:
	eventManage.proto

It has these top-level messages:
	AddEventProblemReq
	AddEventProblemRsp
	EventIdReq
	EventDetailMesssage
	EventShowMesssage
	GetEventListReq
	EventMesssage
	EventListRsp
	AddEventReq
	AddEventRsp
	EventTime
	ProblemNum
	CreditRule
*/
package eventManage

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for EventManage service

type EventManageService interface {
	// 根据数量返回事件信息列表
	GetEventListByManageIdAndOffst(ctx context.Context, in *GetEventListReq, opts ...client.CallOption) (*EventListRsp, error)
	// 新增事件
	AddNewEvent(ctx context.Context, in *AddEventReq, opts ...client.CallOption) (*AddEventRsp, error)
	// 根据事件id返回事件概述信息
	GetEventByEventId(ctx context.Context, in *EventIdReq, opts ...client.CallOption) (*EventShowMesssage, error)
	// 根据事件id返回事件详细信息
	GetDetailEventByEventId(ctx context.Context, in *EventIdReq, opts ...client.CallOption) (*EventDetailMesssage, error)
	// 提供指定事件的积分规则
	GetCreditRuleByEventId(ctx context.Context, in *EventIdReq, opts ...client.CallOption) (*CreditRule, error)
	// 提供指定事件的题目数量
	GetProblemNumByEventId(ctx context.Context, in *EventIdReq, opts ...client.CallOption) (*ProblemNum, error)
	// 增加事件管理的题目
	AddEventProblem(ctx context.Context, in *AddEventProblemReq, opts ...client.CallOption) (*AddEventProblemRsp, error)
}

type eventManageService struct {
	c    client.Client
	name string
}

func NewEventManageService(name string, c client.Client) EventManageService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "eventmanage"
	}
	return &eventManageService{
		c:    c,
		name: name,
	}
}

func (c *eventManageService) GetEventListByManageIdAndOffst(ctx context.Context, in *GetEventListReq, opts ...client.CallOption) (*EventListRsp, error) {
	req := c.c.NewRequest(c.name, "EventManage.GetEventListByManageIdAndOffst", in)
	out := new(EventListRsp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventManageService) AddNewEvent(ctx context.Context, in *AddEventReq, opts ...client.CallOption) (*AddEventRsp, error) {
	req := c.c.NewRequest(c.name, "EventManage.AddNewEvent", in)
	out := new(AddEventRsp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventManageService) GetEventByEventId(ctx context.Context, in *EventIdReq, opts ...client.CallOption) (*EventShowMesssage, error) {
	req := c.c.NewRequest(c.name, "EventManage.GetEventByEventId", in)
	out := new(EventShowMesssage)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventManageService) GetDetailEventByEventId(ctx context.Context, in *EventIdReq, opts ...client.CallOption) (*EventDetailMesssage, error) {
	req := c.c.NewRequest(c.name, "EventManage.GetDetailEventByEventId", in)
	out := new(EventDetailMesssage)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventManageService) GetCreditRuleByEventId(ctx context.Context, in *EventIdReq, opts ...client.CallOption) (*CreditRule, error) {
	req := c.c.NewRequest(c.name, "EventManage.GetCreditRuleByEventId", in)
	out := new(CreditRule)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventManageService) GetProblemNumByEventId(ctx context.Context, in *EventIdReq, opts ...client.CallOption) (*ProblemNum, error) {
	req := c.c.NewRequest(c.name, "EventManage.GetProblemNumByEventId", in)
	out := new(ProblemNum)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventManageService) AddEventProblem(ctx context.Context, in *AddEventProblemReq, opts ...client.CallOption) (*AddEventProblemRsp, error) {
	req := c.c.NewRequest(c.name, "EventManage.AddEventProblem", in)
	out := new(AddEventProblemRsp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for EventManage service

type EventManageHandler interface {
	// 根据数量返回事件信息列表
	GetEventListByManageIdAndOffst(context.Context, *GetEventListReq, *EventListRsp) error
	// 新增事件
	AddNewEvent(context.Context, *AddEventReq, *AddEventRsp) error
	// 根据事件id返回事件概述信息
	GetEventByEventId(context.Context, *EventIdReq, *EventShowMesssage) error
	// 根据事件id返回事件详细信息
	GetDetailEventByEventId(context.Context, *EventIdReq, *EventDetailMesssage) error
	// 提供指定事件的积分规则
	GetCreditRuleByEventId(context.Context, *EventIdReq, *CreditRule) error
	// 提供指定事件的题目数量
	GetProblemNumByEventId(context.Context, *EventIdReq, *ProblemNum) error
	// 增加事件管理的题目
	AddEventProblem(context.Context, *AddEventProblemReq, *AddEventProblemRsp) error
}

func RegisterEventManageHandler(s server.Server, hdlr EventManageHandler, opts ...server.HandlerOption) error {
	type eventManage interface {
		GetEventListByManageIdAndOffst(ctx context.Context, in *GetEventListReq, out *EventListRsp) error
		AddNewEvent(ctx context.Context, in *AddEventReq, out *AddEventRsp) error
		GetEventByEventId(ctx context.Context, in *EventIdReq, out *EventShowMesssage) error
		GetDetailEventByEventId(ctx context.Context, in *EventIdReq, out *EventDetailMesssage) error
		GetCreditRuleByEventId(ctx context.Context, in *EventIdReq, out *CreditRule) error
		GetProblemNumByEventId(ctx context.Context, in *EventIdReq, out *ProblemNum) error
		AddEventProblem(ctx context.Context, in *AddEventProblemReq, out *AddEventProblemRsp) error
	}
	type EventManage struct {
		eventManage
	}
	h := &eventManageHandler{hdlr}
	return s.Handle(s.NewHandler(&EventManage{h}, opts...))
}

type eventManageHandler struct {
	EventManageHandler
}

func (h *eventManageHandler) GetEventListByManageIdAndOffst(ctx context.Context, in *GetEventListReq, out *EventListRsp) error {
	return h.EventManageHandler.GetEventListByManageIdAndOffst(ctx, in, out)
}

func (h *eventManageHandler) AddNewEvent(ctx context.Context, in *AddEventReq, out *AddEventRsp) error {
	return h.EventManageHandler.AddNewEvent(ctx, in, out)
}

func (h *eventManageHandler) GetEventByEventId(ctx context.Context, in *EventIdReq, out *EventShowMesssage) error {
	return h.EventManageHandler.GetEventByEventId(ctx, in, out)
}

func (h *eventManageHandler) GetDetailEventByEventId(ctx context.Context, in *EventIdReq, out *EventDetailMesssage) error {
	return h.EventManageHandler.GetDetailEventByEventId(ctx, in, out)
}

func (h *eventManageHandler) GetCreditRuleByEventId(ctx context.Context, in *EventIdReq, out *CreditRule) error {
	return h.EventManageHandler.GetCreditRuleByEventId(ctx, in, out)
}

func (h *eventManageHandler) GetProblemNumByEventId(ctx context.Context, in *EventIdReq, out *ProblemNum) error {
	return h.EventManageHandler.GetProblemNumByEventId(ctx, in, out)
}

func (h *eventManageHandler) AddEventProblem(ctx context.Context, in *AddEventProblemReq, out *AddEventProblemRsp) error {
	return h.EventManageHandler.AddEventProblem(ctx, in, out)
}
