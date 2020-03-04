// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: problemManage.proto

/*
Package problemManage is a generated protocol buffer package.

It is generated from these files:
	problemManage.proto

It has these top-level messages:
	GetEndProblemIdReq
	GetEndProblemIdRsp
	GetProblemListReq
	ProblemMesssage
	ProblemListRsp
	GetNewProblemByTypeReq
	AddProblemRsp
*/
package problemManage

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

// Client API for ProblemManage service

type ProblemManageService interface {
	// 根据数量返回题目列表
	GetProblemListByOffstAndLimit(ctx context.Context, in *GetProblemListReq, opts ...client.CallOption) (*ProblemListRsp, error)
	// 新增题目
	AddProblem(ctx context.Context, in *ProblemMesssage, opts ...client.CallOption) (*AddProblemRsp, error)
	// 根据题目类型获取题目
	GetNewProblemByType(ctx context.Context, in *GetNewProblemByTypeReq, opts ...client.CallOption) (*ProblemListRsp, error)
	// 获取最大题目ID
	GetEndProblemId(ctx context.Context, in *GetEndProblemIdReq, opts ...client.CallOption) (*GetEndProblemIdRsp, error)
}

type problemManageService struct {
	c    client.Client
	name string
}

func NewProblemManageService(name string, c client.Client) ProblemManageService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "problemmanage"
	}
	return &problemManageService{
		c:    c,
		name: name,
	}
}

func (c *problemManageService) GetProblemListByOffstAndLimit(ctx context.Context, in *GetProblemListReq, opts ...client.CallOption) (*ProblemListRsp, error) {
	req := c.c.NewRequest(c.name, "ProblemManage.GetProblemListByOffstAndLimit", in)
	out := new(ProblemListRsp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *problemManageService) AddProblem(ctx context.Context, in *ProblemMesssage, opts ...client.CallOption) (*AddProblemRsp, error) {
	req := c.c.NewRequest(c.name, "ProblemManage.AddProblem", in)
	out := new(AddProblemRsp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *problemManageService) GetNewProblemByType(ctx context.Context, in *GetNewProblemByTypeReq, opts ...client.CallOption) (*ProblemListRsp, error) {
	req := c.c.NewRequest(c.name, "ProblemManage.GetNewProblemByType", in)
	out := new(ProblemListRsp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *problemManageService) GetEndProblemId(ctx context.Context, in *GetEndProblemIdReq, opts ...client.CallOption) (*GetEndProblemIdRsp, error) {
	req := c.c.NewRequest(c.name, "ProblemManage.GetEndProblemId", in)
	out := new(GetEndProblemIdRsp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ProblemManage service

type ProblemManageHandler interface {
	// 根据数量返回题目列表
	GetProblemListByOffstAndLimit(context.Context, *GetProblemListReq, *ProblemListRsp) error
	// 新增题目
	AddProblem(context.Context, *ProblemMesssage, *AddProblemRsp) error
	// 根据题目类型获取题目
	GetNewProblemByType(context.Context, *GetNewProblemByTypeReq, *ProblemListRsp) error
	// 获取最大题目ID
	GetEndProblemId(context.Context, *GetEndProblemIdReq, *GetEndProblemIdRsp) error
}

func RegisterProblemManageHandler(s server.Server, hdlr ProblemManageHandler, opts ...server.HandlerOption) error {
	type problemManage interface {
		GetProblemListByOffstAndLimit(ctx context.Context, in *GetProblemListReq, out *ProblemListRsp) error
		AddProblem(ctx context.Context, in *ProblemMesssage, out *AddProblemRsp) error
		GetNewProblemByType(ctx context.Context, in *GetNewProblemByTypeReq, out *ProblemListRsp) error
		GetEndProblemId(ctx context.Context, in *GetEndProblemIdReq, out *GetEndProblemIdRsp) error
	}
	type ProblemManage struct {
		problemManage
	}
	h := &problemManageHandler{hdlr}
	return s.Handle(s.NewHandler(&ProblemManage{h}, opts...))
}

type problemManageHandler struct {
	ProblemManageHandler
}

func (h *problemManageHandler) GetProblemListByOffstAndLimit(ctx context.Context, in *GetProblemListReq, out *ProblemListRsp) error {
	return h.ProblemManageHandler.GetProblemListByOffstAndLimit(ctx, in, out)
}

func (h *problemManageHandler) AddProblem(ctx context.Context, in *ProblemMesssage, out *AddProblemRsp) error {
	return h.ProblemManageHandler.AddProblem(ctx, in, out)
}

func (h *problemManageHandler) GetNewProblemByType(ctx context.Context, in *GetNewProblemByTypeReq, out *ProblemListRsp) error {
	return h.ProblemManageHandler.GetNewProblemByType(ctx, in, out)
}

func (h *problemManageHandler) GetEndProblemId(ctx context.Context, in *GetEndProblemIdReq, out *GetEndProblemIdRsp) error {
	return h.ProblemManageHandler.GetEndProblemId(ctx, in, out)
}
