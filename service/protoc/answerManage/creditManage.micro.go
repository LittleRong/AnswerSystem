// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: creditManage.proto

/*
Package creditManage is a generated protocol buffer package.

It is generated from these files:
	creditManage.proto

It has these top-level messages:
	UpdateTeamCreditReq
	UpdatePCreditReq
	AllRightReq
	AllRightRsp
	TeamIdReq
	TeamEventIdReq
	UserEventIdReq
	CreditRsp
	CreditLog
	CreditLogListRsp
	AddCreditLogRsp
*/
package answerManage

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

// Client API for CreditManage service

type CreditManageService interface {
	GetTeamCredit(ctx context.Context, in *TeamEventIdReq, opts ...client.CallOption) (*CreditRsp, error)
	GetPersonCredit(ctx context.Context, in *UserEventIdReq, opts ...client.CallOption) (*CreditRsp, error)
	GetCreditLogByTeamId(ctx context.Context, in *TeamIdReq, opts ...client.CallOption) (*CreditLogListRsp, error)
	AddCreditLog(ctx context.Context, in *CreditLog, opts ...client.CallOption) (*AddCreditLogRsp, error)
	WhetherMemberAllRight(ctx context.Context, in *AllRightReq, opts ...client.CallOption) (*AllRightRsp, error)
	UpdateTeamCredit(ctx context.Context, in *UpdateTeamCreditReq, opts ...client.CallOption) (*CreditRsp, error)
	UpdateParticipantCredit(ctx context.Context, in *UpdatePCreditReq, opts ...client.CallOption) (*CreditRsp, error)
}

type creditManageService struct {
	c    client.Client
	name string
}

func NewCreditManageService(name string, c client.Client) CreditManageService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "creditmanage"
	}
	return &creditManageService{
		c:    c,
		name: name,
	}
}

func (c *creditManageService) GetTeamCredit(ctx context.Context, in *TeamEventIdReq, opts ...client.CallOption) (*CreditRsp, error) {
	req := c.c.NewRequest(c.name, "CreditManage.GetTeamCredit", in)
	out := new(CreditRsp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *creditManageService) GetPersonCredit(ctx context.Context, in *UserEventIdReq, opts ...client.CallOption) (*CreditRsp, error) {
	req := c.c.NewRequest(c.name, "CreditManage.GetPersonCredit", in)
	out := new(CreditRsp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *creditManageService) GetCreditLogByTeamId(ctx context.Context, in *TeamIdReq, opts ...client.CallOption) (*CreditLogListRsp, error) {
	req := c.c.NewRequest(c.name, "CreditManage.GetCreditLogByTeamId", in)
	out := new(CreditLogListRsp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *creditManageService) AddCreditLog(ctx context.Context, in *CreditLog, opts ...client.CallOption) (*AddCreditLogRsp, error) {
	req := c.c.NewRequest(c.name, "CreditManage.AddCreditLog", in)
	out := new(AddCreditLogRsp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *creditManageService) WhetherMemberAllRight(ctx context.Context, in *AllRightReq, opts ...client.CallOption) (*AllRightRsp, error) {
	req := c.c.NewRequest(c.name, "CreditManage.WhetherMemberAllRight", in)
	out := new(AllRightRsp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *creditManageService) UpdateTeamCredit(ctx context.Context, in *UpdateTeamCreditReq, opts ...client.CallOption) (*CreditRsp, error) {
	req := c.c.NewRequest(c.name, "CreditManage.UpdateTeamCredit", in)
	out := new(CreditRsp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *creditManageService) UpdateParticipantCredit(ctx context.Context, in *UpdatePCreditReq, opts ...client.CallOption) (*CreditRsp, error) {
	req := c.c.NewRequest(c.name, "CreditManage.UpdateParticipantCredit", in)
	out := new(CreditRsp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CreditManage service

type CreditManageHandler interface {
	GetTeamCredit(context.Context, *TeamEventIdReq, *CreditRsp) error
	GetPersonCredit(context.Context, *UserEventIdReq, *CreditRsp) error
	GetCreditLogByTeamId(context.Context, *TeamIdReq, *CreditLogListRsp) error
	AddCreditLog(context.Context, *CreditLog, *AddCreditLogRsp) error
	WhetherMemberAllRight(context.Context, *AllRightReq, *AllRightRsp) error
	UpdateTeamCredit(context.Context, *UpdateTeamCreditReq, *CreditRsp) error
	UpdateParticipantCredit(context.Context, *UpdatePCreditReq, *CreditRsp) error
}

func RegisterCreditManageHandler(s server.Server, hdlr CreditManageHandler, opts ...server.HandlerOption) error {
	type creditManage interface {
		GetTeamCredit(ctx context.Context, in *TeamEventIdReq, out *CreditRsp) error
		GetPersonCredit(ctx context.Context, in *UserEventIdReq, out *CreditRsp) error
		GetCreditLogByTeamId(ctx context.Context, in *TeamIdReq, out *CreditLogListRsp) error
		AddCreditLog(ctx context.Context, in *CreditLog, out *AddCreditLogRsp) error
		WhetherMemberAllRight(ctx context.Context, in *AllRightReq, out *AllRightRsp) error
		UpdateTeamCredit(ctx context.Context, in *UpdateTeamCreditReq, out *CreditRsp) error
		UpdateParticipantCredit(ctx context.Context, in *UpdatePCreditReq, out *CreditRsp) error
	}
	type CreditManage struct {
		creditManage
	}
	h := &creditManageHandler{hdlr}
	return s.Handle(s.NewHandler(&CreditManage{h}, opts...))
}

type creditManageHandler struct {
	CreditManageHandler
}

func (h *creditManageHandler) GetTeamCredit(ctx context.Context, in *TeamEventIdReq, out *CreditRsp) error {
	return h.CreditManageHandler.GetTeamCredit(ctx, in, out)
}

func (h *creditManageHandler) GetPersonCredit(ctx context.Context, in *UserEventIdReq, out *CreditRsp) error {
	return h.CreditManageHandler.GetPersonCredit(ctx, in, out)
}

func (h *creditManageHandler) GetCreditLogByTeamId(ctx context.Context, in *TeamIdReq, out *CreditLogListRsp) error {
	return h.CreditManageHandler.GetCreditLogByTeamId(ctx, in, out)
}

func (h *creditManageHandler) AddCreditLog(ctx context.Context, in *CreditLog, out *AddCreditLogRsp) error {
	return h.CreditManageHandler.AddCreditLog(ctx, in, out)
}

func (h *creditManageHandler) WhetherMemberAllRight(ctx context.Context, in *AllRightReq, out *AllRightRsp) error {
	return h.CreditManageHandler.WhetherMemberAllRight(ctx, in, out)
}

func (h *creditManageHandler) UpdateTeamCredit(ctx context.Context, in *UpdateTeamCreditReq, out *CreditRsp) error {
	return h.CreditManageHandler.UpdateTeamCredit(ctx, in, out)
}

func (h *creditManageHandler) UpdateParticipantCredit(ctx context.Context, in *UpdatePCreditReq, out *CreditRsp) error {
	return h.CreditManageHandler.UpdateParticipantCredit(ctx, in, out)
}
