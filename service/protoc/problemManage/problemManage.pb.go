// Code generated by protoc-gen-go. DO NOT EDIT.
// source: problemManage.proto

package problemManage

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type GetEndProblemIdReq struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetEndProblemIdReq) Reset()         { *m = GetEndProblemIdReq{} }
func (m *GetEndProblemIdReq) String() string { return proto.CompactTextString(m) }
func (*GetEndProblemIdReq) ProtoMessage()    {}
func (*GetEndProblemIdReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d2c35f53333d071, []int{0}
}

func (m *GetEndProblemIdReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetEndProblemIdReq.Unmarshal(m, b)
}
func (m *GetEndProblemIdReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetEndProblemIdReq.Marshal(b, m, deterministic)
}
func (m *GetEndProblemIdReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetEndProblemIdReq.Merge(m, src)
}
func (m *GetEndProblemIdReq) XXX_Size() int {
	return xxx_messageInfo_GetEndProblemIdReq.Size(m)
}
func (m *GetEndProblemIdReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetEndProblemIdReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetEndProblemIdReq proto.InternalMessageInfo

type GetEndProblemIdRsp struct {
	EndId                int64    `protobuf:"varint,1,opt,name=endId,proto3" json:"endId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetEndProblemIdRsp) Reset()         { *m = GetEndProblemIdRsp{} }
func (m *GetEndProblemIdRsp) String() string { return proto.CompactTextString(m) }
func (*GetEndProblemIdRsp) ProtoMessage()    {}
func (*GetEndProblemIdRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d2c35f53333d071, []int{1}
}

func (m *GetEndProblemIdRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetEndProblemIdRsp.Unmarshal(m, b)
}
func (m *GetEndProblemIdRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetEndProblemIdRsp.Marshal(b, m, deterministic)
}
func (m *GetEndProblemIdRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetEndProblemIdRsp.Merge(m, src)
}
func (m *GetEndProblemIdRsp) XXX_Size() int {
	return xxx_messageInfo_GetEndProblemIdRsp.Size(m)
}
func (m *GetEndProblemIdRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_GetEndProblemIdRsp.DiscardUnknown(m)
}

var xxx_messageInfo_GetEndProblemIdRsp proto.InternalMessageInfo

func (m *GetEndProblemIdRsp) GetEndId() int64 {
	if m != nil {
		return m.EndId
	}
	return 0
}

type GetProblemListReq struct {
	Offset               int32    `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit                int32    `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	ManageId             int64    `protobuf:"varint,3,opt,name=manageId,proto3" json:"manageId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetProblemListReq) Reset()         { *m = GetProblemListReq{} }
func (m *GetProblemListReq) String() string { return proto.CompactTextString(m) }
func (*GetProblemListReq) ProtoMessage()    {}
func (*GetProblemListReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d2c35f53333d071, []int{2}
}

func (m *GetProblemListReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetProblemListReq.Unmarshal(m, b)
}
func (m *GetProblemListReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetProblemListReq.Marshal(b, m, deterministic)
}
func (m *GetProblemListReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetProblemListReq.Merge(m, src)
}
func (m *GetProblemListReq) XXX_Size() int {
	return xxx_messageInfo_GetProblemListReq.Size(m)
}
func (m *GetProblemListReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetProblemListReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetProblemListReq proto.InternalMessageInfo

func (m *GetProblemListReq) GetOffset() int32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *GetProblemListReq) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *GetProblemListReq) GetManageId() int64 {
	if m != nil {
		return m.ManageId
	}
	return 0
}

type ProblemMesssage struct {
	ProblemId            int64    `protobuf:"varint,1,opt,name=problemId,proto3" json:"problemId,omitempty"`
	ProblemContent       string   `protobuf:"bytes,2,opt,name=problemContent,proto3" json:"problemContent,omitempty"`
	ProblemOption        string   `protobuf:"bytes,3,opt,name=problemOption,proto3" json:"problemOption,omitempty"`
	ProblemAnswer        string   `protobuf:"bytes,4,opt,name=problemAnswer,proto3" json:"problemAnswer,omitempty"`
	ProblemClass         string   `protobuf:"bytes,5,opt,name=problemClass,proto3" json:"problemClass,omitempty"`
	ProblemType          int32    `protobuf:"varint,6,opt,name=problemType,proto3" json:"problemType,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProblemMesssage) Reset()         { *m = ProblemMesssage{} }
func (m *ProblemMesssage) String() string { return proto.CompactTextString(m) }
func (*ProblemMesssage) ProtoMessage()    {}
func (*ProblemMesssage) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d2c35f53333d071, []int{3}
}

func (m *ProblemMesssage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProblemMesssage.Unmarshal(m, b)
}
func (m *ProblemMesssage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProblemMesssage.Marshal(b, m, deterministic)
}
func (m *ProblemMesssage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProblemMesssage.Merge(m, src)
}
func (m *ProblemMesssage) XXX_Size() int {
	return xxx_messageInfo_ProblemMesssage.Size(m)
}
func (m *ProblemMesssage) XXX_DiscardUnknown() {
	xxx_messageInfo_ProblemMesssage.DiscardUnknown(m)
}

var xxx_messageInfo_ProblemMesssage proto.InternalMessageInfo

func (m *ProblemMesssage) GetProblemId() int64 {
	if m != nil {
		return m.ProblemId
	}
	return 0
}

func (m *ProblemMesssage) GetProblemContent() string {
	if m != nil {
		return m.ProblemContent
	}
	return ""
}

func (m *ProblemMesssage) GetProblemOption() string {
	if m != nil {
		return m.ProblemOption
	}
	return ""
}

func (m *ProblemMesssage) GetProblemAnswer() string {
	if m != nil {
		return m.ProblemAnswer
	}
	return ""
}

func (m *ProblemMesssage) GetProblemClass() string {
	if m != nil {
		return m.ProblemClass
	}
	return ""
}

func (m *ProblemMesssage) GetProblemType() int32 {
	if m != nil {
		return m.ProblemType
	}
	return 0
}

type ProblemListRsp struct {
	ProblemList          []*ProblemMesssage `protobuf:"bytes,1,rep,name=problemList,proto3" json:"problemList,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *ProblemListRsp) Reset()         { *m = ProblemListRsp{} }
func (m *ProblemListRsp) String() string { return proto.CompactTextString(m) }
func (*ProblemListRsp) ProtoMessage()    {}
func (*ProblemListRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d2c35f53333d071, []int{4}
}

func (m *ProblemListRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProblemListRsp.Unmarshal(m, b)
}
func (m *ProblemListRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProblemListRsp.Marshal(b, m, deterministic)
}
func (m *ProblemListRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProblemListRsp.Merge(m, src)
}
func (m *ProblemListRsp) XXX_Size() int {
	return xxx_messageInfo_ProblemListRsp.Size(m)
}
func (m *ProblemListRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_ProblemListRsp.DiscardUnknown(m)
}

var xxx_messageInfo_ProblemListRsp proto.InternalMessageInfo

func (m *ProblemListRsp) GetProblemList() []*ProblemMesssage {
	if m != nil {
		return m.ProblemList
	}
	return nil
}

type GetNewProblemByTypeReq struct {
	FirstProblemId       int64    `protobuf:"varint,1,opt,name=firstProblemId,proto3" json:"firstProblemId,omitempty"`
	ProblemType          int32    `protobuf:"varint,2,opt,name=problemType,proto3" json:"problemType,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetNewProblemByTypeReq) Reset()         { *m = GetNewProblemByTypeReq{} }
func (m *GetNewProblemByTypeReq) String() string { return proto.CompactTextString(m) }
func (*GetNewProblemByTypeReq) ProtoMessage()    {}
func (*GetNewProblemByTypeReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d2c35f53333d071, []int{5}
}

func (m *GetNewProblemByTypeReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetNewProblemByTypeReq.Unmarshal(m, b)
}
func (m *GetNewProblemByTypeReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetNewProblemByTypeReq.Marshal(b, m, deterministic)
}
func (m *GetNewProblemByTypeReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetNewProblemByTypeReq.Merge(m, src)
}
func (m *GetNewProblemByTypeReq) XXX_Size() int {
	return xxx_messageInfo_GetNewProblemByTypeReq.Size(m)
}
func (m *GetNewProblemByTypeReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetNewProblemByTypeReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetNewProblemByTypeReq proto.InternalMessageInfo

func (m *GetNewProblemByTypeReq) GetFirstProblemId() int64 {
	if m != nil {
		return m.FirstProblemId
	}
	return 0
}

func (m *GetNewProblemByTypeReq) GetProblemType() int32 {
	if m != nil {
		return m.ProblemType
	}
	return 0
}

type AddProblemRsp struct {
	ProblemId            int64    `protobuf:"varint,1,opt,name=problemId,proto3" json:"problemId,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddProblemRsp) Reset()         { *m = AddProblemRsp{} }
func (m *AddProblemRsp) String() string { return proto.CompactTextString(m) }
func (*AddProblemRsp) ProtoMessage()    {}
func (*AddProblemRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d2c35f53333d071, []int{6}
}

func (m *AddProblemRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddProblemRsp.Unmarshal(m, b)
}
func (m *AddProblemRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddProblemRsp.Marshal(b, m, deterministic)
}
func (m *AddProblemRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddProblemRsp.Merge(m, src)
}
func (m *AddProblemRsp) XXX_Size() int {
	return xxx_messageInfo_AddProblemRsp.Size(m)
}
func (m *AddProblemRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_AddProblemRsp.DiscardUnknown(m)
}

var xxx_messageInfo_AddProblemRsp proto.InternalMessageInfo

func (m *AddProblemRsp) GetProblemId() int64 {
	if m != nil {
		return m.ProblemId
	}
	return 0
}

func (m *AddProblemRsp) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*GetEndProblemIdReq)(nil), "GetEndProblemIdReq")
	proto.RegisterType((*GetEndProblemIdRsp)(nil), "GetEndProblemIdRsp")
	proto.RegisterType((*GetProblemListReq)(nil), "GetProblemListReq")
	proto.RegisterType((*ProblemMesssage)(nil), "ProblemMesssage")
	proto.RegisterType((*ProblemListRsp)(nil), "ProblemListRsp")
	proto.RegisterType((*GetNewProblemByTypeReq)(nil), "GetNewProblemByTypeReq")
	proto.RegisterType((*AddProblemRsp)(nil), "AddProblemRsp")
}

func init() { proto.RegisterFile("problemManage.proto", fileDescriptor_4d2c35f53333d071) }

var fileDescriptor_4d2c35f53333d071 = []byte{
	// 414 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x53, 0x4d, 0x6b, 0xdb, 0x40,
	0x10, 0x8d, 0xec, 0xda, 0xad, 0x27, 0xb5, 0xdc, 0x8e, 0x43, 0xba, 0x88, 0x16, 0xcc, 0x52, 0x42,
	0xe8, 0x41, 0x14, 0xf7, 0xdc, 0x83, 0xd2, 0x0f, 0x61, 0x48, 0x1a, 0x23, 0x7a, 0xed, 0x41, 0x46,
	0xa3, 0x20, 0xb0, 0x56, 0xb2, 0x76, 0xc1, 0xf8, 0x37, 0xf4, 0x97, 0xf6, 0x5f, 0x14, 0xad, 0x56,
	0xb6, 0x3e, 0x4c, 0x8e, 0xef, 0xe9, 0xed, 0xcc, 0xd3, 0x9b, 0x19, 0x98, 0xe7, 0x45, 0xb6, 0xd9,
	0x52, 0xfa, 0x10, 0x8a, 0xf0, 0x89, 0xdc, 0xbc, 0xc8, 0x54, 0xc6, 0xaf, 0x00, 0x7d, 0x52, 0x3f,
	0x44, 0xb4, 0xae, 0x3e, 0xae, 0xa2, 0x80, 0x76, 0xfc, 0x53, 0x9f, 0x95, 0x39, 0x5e, 0xc1, 0x88,
	0x44, 0xb4, 0x8a, 0x98, 0xb5, 0xb0, 0x6e, 0x87, 0x41, 0x05, 0xf8, 0x1f, 0x78, 0xeb, 0x93, 0x32,
	0xc2, 0xfb, 0x44, 0xaa, 0x80, 0x76, 0x78, 0x0d, 0xe3, 0x2c, 0x8e, 0x25, 0x29, 0xad, 0x1d, 0x05,
	0x06, 0x95, 0x25, 0xb6, 0x49, 0x9a, 0x28, 0x36, 0xd0, 0x74, 0x05, 0xd0, 0x81, 0x57, 0xa9, 0x36,
	0xb5, 0x8a, 0xd8, 0x50, 0xd7, 0x3e, 0x62, 0xfe, 0xcf, 0x82, 0x99, 0x29, 0xfe, 0x40, 0x52, 0xca,
	0xf0, 0x89, 0xf0, 0x3d, 0x4c, 0xf2, 0xda, 0x98, 0x31, 0x73, 0x22, 0xf0, 0x06, 0x6c, 0x03, 0xbe,
	0x65, 0x42, 0x91, 0xa8, 0x9a, 0x4d, 0x82, 0x0e, 0x8b, 0x1f, 0x61, 0x6a, 0x98, 0xc7, 0x5c, 0x25,
	0x99, 0xd0, 0xad, 0x27, 0x41, 0x9b, 0x6c, 0xa8, 0x3c, 0x21, 0xf7, 0x54, 0xb0, 0x17, 0x2d, 0x55,
	0x45, 0x22, 0x87, 0xd7, 0x75, 0xf5, 0x6d, 0x28, 0x25, 0x1b, 0x69, 0x51, 0x8b, 0xc3, 0x05, 0x5c,
	0x1a, 0xfc, 0xfb, 0x90, 0x13, 0x1b, 0xeb, 0x04, 0x9a, 0x14, 0xff, 0x0e, 0x76, 0x33, 0x47, 0x99,
	0xe3, 0xf2, 0xf8, 0xa6, 0x64, 0x98, 0xb5, 0x18, 0xde, 0x5e, 0x2e, 0xdf, 0xb8, 0x9d, 0x40, 0x82,
	0xa6, 0x88, 0x6f, 0xe0, 0xda, 0x27, 0xf5, 0x8b, 0xf6, 0x46, 0x75, 0x77, 0x28, 0x8b, 0x97, 0x53,
	0xb9, 0x01, 0x3b, 0x4e, 0x0a, 0xa9, 0xd6, 0x9d, 0xf0, 0x3a, 0x6c, 0xd7, 0xe9, 0xa0, 0xef, 0xd4,
	0x87, 0xa9, 0x17, 0xd5, 0xdb, 0x51, 0x1a, 0x7d, 0x7e, 0x24, 0x0c, 0x5e, 0xa6, 0xa4, 0xad, 0x9a,
	0x59, 0xd4, 0x70, 0xf9, 0x77, 0x00, 0xd3, 0x75, 0x73, 0x2f, 0xf1, 0x27, 0x7c, 0x68, 0xef, 0xd3,
	0xdd, 0xe1, 0x31, 0x8e, 0xa5, 0xf2, 0x44, 0x74, 0xaf, 0xb7, 0x05, 0xdd, 0xde, 0xbe, 0x39, 0x33,
	0xb7, 0x1d, 0x1c, 0xbf, 0xc0, 0xcf, 0x00, 0x27, 0x8b, 0xd8, 0xcb, 0xcc, 0xb1, 0xdd, 0xd6, 0x1f,
	0xf0, 0x0b, 0xf4, 0x60, 0x7e, 0x26, 0x38, 0x7c, 0xe7, 0x9e, 0x8f, 0xf3, 0x5c, 0xd3, 0xaf, 0x30,
	0xeb, 0x1c, 0x0e, 0xce, 0xdd, 0xfe, 0x81, 0x39, 0x7d, 0xb2, 0x7c, 0xbe, 0x19, 0xeb, 0xa3, 0xfc,
	0xf2, 0x3f, 0x00, 0x00, 0xff, 0xff, 0x6d, 0x54, 0x25, 0xa3, 0xab, 0x03, 0x00, 0x00,
}
