// Code generated by protoc-gen-go. DO NOT EDIT.
// source: commpb.proto

package commpb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type VoteRsp_VtRes int32

const (
	VoteRsp_NO  VoteRsp_VtRes = 0
	VoteRsp_YES VoteRsp_VtRes = 1
)

var VoteRsp_VtRes_name = map[int32]string{
	0: "NO",
	1: "YES",
}

var VoteRsp_VtRes_value = map[string]int32{
	"NO":  0,
	"YES": 1,
}

func (x VoteRsp_VtRes) String() string {
	return proto.EnumName(VoteRsp_VtRes_name, int32(x))
}

func (VoteRsp_VtRes) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d2eb5c7490bed74a, []int{1, 0}
}

type VoteReq struct {
	Term                 int64    `protobuf:"varint,1,opt,name=term,proto3" json:"term,omitempty"`
	Id                   int64    `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VoteReq) Reset()         { *m = VoteReq{} }
func (m *VoteReq) String() string { return proto.CompactTextString(m) }
func (*VoteReq) ProtoMessage()    {}
func (*VoteReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_d2eb5c7490bed74a, []int{0}
}

func (m *VoteReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VoteReq.Unmarshal(m, b)
}
func (m *VoteReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VoteReq.Marshal(b, m, deterministic)
}
func (m *VoteReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VoteReq.Merge(m, src)
}
func (m *VoteReq) XXX_Size() int {
	return xxx_messageInfo_VoteReq.Size(m)
}
func (m *VoteReq) XXX_DiscardUnknown() {
	xxx_messageInfo_VoteReq.DiscardUnknown(m)
}

var xxx_messageInfo_VoteReq proto.InternalMessageInfo

func (m *VoteReq) GetTerm() int64 {
	if m != nil {
		return m.Term
	}
	return 0
}

func (m *VoteReq) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type VoteRsp struct {
	Vtres                VoteRsp_VtRes `protobuf:"varint,1,opt,name=vtres,proto3,enum=VoteRsp_VtRes" json:"vtres,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *VoteRsp) Reset()         { *m = VoteRsp{} }
func (m *VoteRsp) String() string { return proto.CompactTextString(m) }
func (*VoteRsp) ProtoMessage()    {}
func (*VoteRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_d2eb5c7490bed74a, []int{1}
}

func (m *VoteRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VoteRsp.Unmarshal(m, b)
}
func (m *VoteRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VoteRsp.Marshal(b, m, deterministic)
}
func (m *VoteRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VoteRsp.Merge(m, src)
}
func (m *VoteRsp) XXX_Size() int {
	return xxx_messageInfo_VoteRsp.Size(m)
}
func (m *VoteRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_VoteRsp.DiscardUnknown(m)
}

var xxx_messageInfo_VoteRsp proto.InternalMessageInfo

func (m *VoteRsp) GetVtres() VoteRsp_VtRes {
	if m != nil {
		return m.Vtres
	}
	return VoteRsp_NO
}

type HBReq struct {
	Term                 int64    `protobuf:"varint,1,opt,name=term,proto3" json:"term,omitempty"`
	Id                   int64    `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HBReq) Reset()         { *m = HBReq{} }
func (m *HBReq) String() string { return proto.CompactTextString(m) }
func (*HBReq) ProtoMessage()    {}
func (*HBReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_d2eb5c7490bed74a, []int{2}
}

func (m *HBReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HBReq.Unmarshal(m, b)
}
func (m *HBReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HBReq.Marshal(b, m, deterministic)
}
func (m *HBReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HBReq.Merge(m, src)
}
func (m *HBReq) XXX_Size() int {
	return xxx_messageInfo_HBReq.Size(m)
}
func (m *HBReq) XXX_DiscardUnknown() {
	xxx_messageInfo_HBReq.DiscardUnknown(m)
}

var xxx_messageInfo_HBReq proto.InternalMessageInfo

func (m *HBReq) GetTerm() int64 {
	if m != nil {
		return m.Term
	}
	return 0
}

func (m *HBReq) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type HBRsp struct {
	Term                 int64    `protobuf:"varint,1,opt,name=term,proto3" json:"term,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HBRsp) Reset()         { *m = HBRsp{} }
func (m *HBRsp) String() string { return proto.CompactTextString(m) }
func (*HBRsp) ProtoMessage()    {}
func (*HBRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_d2eb5c7490bed74a, []int{3}
}

func (m *HBRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HBRsp.Unmarshal(m, b)
}
func (m *HBRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HBRsp.Marshal(b, m, deterministic)
}
func (m *HBRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HBRsp.Merge(m, src)
}
func (m *HBRsp) XXX_Size() int {
	return xxx_messageInfo_HBRsp.Size(m)
}
func (m *HBRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_HBRsp.DiscardUnknown(m)
}

var xxx_messageInfo_HBRsp proto.InternalMessageInfo

func (m *HBRsp) GetTerm() int64 {
	if m != nil {
		return m.Term
	}
	return 0
}

type Msg struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Msg) Reset()         { *m = Msg{} }
func (m *Msg) String() string { return proto.CompactTextString(m) }
func (*Msg) ProtoMessage()    {}
func (*Msg) Descriptor() ([]byte, []int) {
	return fileDescriptor_d2eb5c7490bed74a, []int{4}
}

func (m *Msg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Msg.Unmarshal(m, b)
}
func (m *Msg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Msg.Marshal(b, m, deterministic)
}
func (m *Msg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Msg.Merge(m, src)
}
func (m *Msg) XXX_Size() int {
	return xxx_messageInfo_Msg.Size(m)
}
func (m *Msg) XXX_DiscardUnknown() {
	xxx_messageInfo_Msg.DiscardUnknown(m)
}

var xxx_messageInfo_Msg proto.InternalMessageInfo

type CP struct {
	Data                 []byte   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Status               int64    `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
	Nodeid               int64    `protobuf:"varint,3,opt,name=nodeid,proto3" json:"nodeid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CP) Reset()         { *m = CP{} }
func (m *CP) String() string { return proto.CompactTextString(m) }
func (*CP) ProtoMessage()    {}
func (*CP) Descriptor() ([]byte, []int) {
	return fileDescriptor_d2eb5c7490bed74a, []int{5}
}

func (m *CP) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CP.Unmarshal(m, b)
}
func (m *CP) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CP.Marshal(b, m, deterministic)
}
func (m *CP) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CP.Merge(m, src)
}
func (m *CP) XXX_Size() int {
	return xxx_messageInfo_CP.Size(m)
}
func (m *CP) XXX_DiscardUnknown() {
	xxx_messageInfo_CP.DiscardUnknown(m)
}

var xxx_messageInfo_CP proto.InternalMessageInfo

func (m *CP) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *CP) GetStatus() int64 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *CP) GetNodeid() int64 {
	if m != nil {
		return m.Nodeid
	}
	return 0
}

type Commit struct {
	Data                 []byte   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Term                 int64    `protobuf:"varint,2,opt,name=term,proto3" json:"term,omitempty"`
	Id                   int64    `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
	LastTerm             int64    `protobuf:"varint,4,opt,name=last_term,json=lastTerm,proto3" json:"last_term,omitempty"`
	LastId               int64    `protobuf:"varint,5,opt,name=last_id,json=lastId,proto3" json:"last_id,omitempty"`
	SrcId                int64    `protobuf:"varint,6,opt,name=src_id,json=srcId,proto3" json:"src_id,omitempty"`
	Type                 int64    `protobuf:"varint,7,opt,name=type,proto3" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Commit) Reset()         { *m = Commit{} }
func (m *Commit) String() string { return proto.CompactTextString(m) }
func (*Commit) ProtoMessage()    {}
func (*Commit) Descriptor() ([]byte, []int) {
	return fileDescriptor_d2eb5c7490bed74a, []int{6}
}

func (m *Commit) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Commit.Unmarshal(m, b)
}
func (m *Commit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Commit.Marshal(b, m, deterministic)
}
func (m *Commit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Commit.Merge(m, src)
}
func (m *Commit) XXX_Size() int {
	return xxx_messageInfo_Commit.Size(m)
}
func (m *Commit) XXX_DiscardUnknown() {
	xxx_messageInfo_Commit.DiscardUnknown(m)
}

var xxx_messageInfo_Commit proto.InternalMessageInfo

func (m *Commit) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Commit) GetTerm() int64 {
	if m != nil {
		return m.Term
	}
	return 0
}

func (m *Commit) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Commit) GetLastTerm() int64 {
	if m != nil {
		return m.LastTerm
	}
	return 0
}

func (m *Commit) GetLastId() int64 {
	if m != nil {
		return m.LastId
	}
	return 0
}

func (m *Commit) GetSrcId() int64 {
	if m != nil {
		return m.SrcId
	}
	return 0
}

func (m *Commit) GetType() int64 {
	if m != nil {
		return m.Type
	}
	return 0
}

func init() {
	proto.RegisterEnum("VoteRsp_VtRes", VoteRsp_VtRes_name, VoteRsp_VtRes_value)
	proto.RegisterType((*VoteReq)(nil), "VoteReq")
	proto.RegisterType((*VoteRsp)(nil), "VoteRsp")
	proto.RegisterType((*HBReq)(nil), "HBReq")
	proto.RegisterType((*HBRsp)(nil), "HBRsp")
	proto.RegisterType((*Msg)(nil), "Msg")
	proto.RegisterType((*CP)(nil), "CP")
	proto.RegisterType((*Commit)(nil), "Commit")
}

func init() {
	proto.RegisterFile("commpb.proto", fileDescriptor_d2eb5c7490bed74a)
}

var fileDescriptor_d2eb5c7490bed74a = []byte{
	// 432 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xd1, 0x6e, 0xd3, 0x30,
	0x14, 0x86, 0xdb, 0xa4, 0x49, 0xba, 0xa3, 0x52, 0x26, 0x4b, 0x40, 0x48, 0x41, 0x20, 0x8b, 0x0b,
	0x24, 0x84, 0x85, 0xc6, 0x13, 0xd0, 0x08, 0xd6, 0x4a, 0x0c, 0xa2, 0x6c, 0x9a, 0xc4, 0x15, 0x72,
	0xe3, 0x43, 0xb1, 0xd6, 0xc4, 0xc6, 0x76, 0x11, 0x7b, 0x1c, 0x1e, 0x8c, 0x77, 0x41, 0x76, 0xd2,
	0xab, 0xee, 0x62, 0x57, 0x3e, 0xe7, 0xff, 0x3f, 0xf9, 0xfc, 0xf2, 0x31, 0xcc, 0x1a, 0xd5, 0xb6,
	0x7a, 0xc3, 0xb4, 0x51, 0x4e, 0xd1, 0xb7, 0x90, 0x5d, 0x2b, 0x87, 0x35, 0xfe, 0x22, 0x04, 0x26,
	0x0e, 0x4d, 0x9b, 0x8f, 0x5f, 0x8e, 0x5f, 0xc7, 0x75, 0xa8, 0xc9, 0x1c, 0x22, 0x29, 0xf2, 0x28,
	0x28, 0x91, 0x14, 0x74, 0x3d, 0xe0, 0x56, 0x93, 0x57, 0x90, 0xfc, 0x76, 0x06, 0x6d, 0xe0, 0xe7,
	0x67, 0x73, 0x36, 0x18, 0xec, 0xda, 0xd5, 0x68, 0xeb, 0xde, 0xa4, 0x39, 0x24, 0xa1, 0x27, 0x29,
	0x44, 0x5f, 0xbe, 0x9e, 0x8e, 0x48, 0x06, 0xf1, 0xb7, 0x8f, 0x97, 0xa7, 0x63, 0xfa, 0x06, 0x92,
	0xd5, 0xf2, 0xbe, 0x73, 0x17, 0x01, 0xb6, 0xfa, 0x2e, 0x98, 0x26, 0x10, 0x5f, 0xd8, 0x2d, 0x5d,
	0x41, 0x54, 0x56, 0x1e, 0x10, 0xdc, 0xf1, 0x00, 0xcc, 0xea, 0x50, 0x93, 0xc7, 0x90, 0x5a, 0xc7,
	0xdd, 0xde, 0x0e, 0x37, 0x0e, 0x9d, 0xd7, 0x3b, 0x25, 0x50, 0x8a, 0x3c, 0xee, 0xf5, 0xbe, 0xa3,
	0x7f, 0xc7, 0x90, 0x96, 0xaa, 0x6d, 0xa5, 0xbb, 0xf3, 0xba, 0x43, 0x86, 0xe8, 0x28, 0x70, 0x7c,
	0x08, 0x4c, 0x16, 0x70, 0xb2, 0xe3, 0xd6, 0x7d, 0x0f, 0xe0, 0x24, 0xc8, 0x53, 0x2f, 0x5c, 0x79,
	0xf8, 0x09, 0x64, 0xc1, 0x94, 0x22, 0x4f, 0xfa, 0xc1, 0xbe, 0x5d, 0x0b, 0xf2, 0x08, 0x52, 0x6b,
	0x1a, 0xaf, 0xa7, 0x41, 0x4f, 0xac, 0x69, 0xd6, 0x22, 0x0c, 0xbc, 0xd5, 0x98, 0x67, 0xc3, 0xc0,
	0x5b, 0x8d, 0x67, 0xff, 0x86, 0x8c, 0x7a, 0x43, 0x28, 0xc0, 0x07, 0x7b, 0xf3, 0x49, 0x19, 0xbf,
	0x01, 0x32, 0x65, 0xc3, 0x42, 0x8b, 0xe9, 0x61, 0x25, 0x74, 0x44, 0x9e, 0xc3, 0xc9, 0x0a, 0xb9,
	0x71, 0x4b, 0xe4, 0x8e, 0xa4, 0x2c, 0xbc, 0x7c, 0x11, 0xce, 0x60, 0x3f, 0x83, 0x07, 0xe7, 0xe8,
	0xca, 0x9f, 0xd8, 0xdc, 0x54, 0x4a, 0x76, 0x8e, 0x4c, 0xd8, 0x85, 0xdd, 0x16, 0x31, 0x2b, 0x2b,
	0x3a, 0x22, 0x05, 0x64, 0x95, 0x41, 0xcd, 0x0d, 0x92, 0x8c, 0xf5, 0x0f, 0x53, 0x04, 0xa0, 0xf7,
	0x4a, 0xd5, 0xfd, 0x90, 0xa6, 0x3d, 0xf6, 0x9e, 0x42, 0x5a, 0xf2, 0xae, 0xc1, 0xdd, 0xb1, 0xf5,
	0x02, 0x66, 0x97, 0xd8, 0x89, 0x2b, 0xf5, 0x19, 0xb9, 0x40, 0x73, 0x04, 0x2c, 0xdf, 0xc1, 0x42,
	0x2a, 0xb6, 0x35, 0xba, 0x61, 0xf8, 0x87, 0xb7, 0x7a, 0x87, 0x96, 0x19, 0xb5, 0x77, 0xb8, 0xdd,
	0x4b, 0x81, 0xcb, 0x87, 0xb5, 0xaf, 0xcf, 0x7d, 0x5d, 0xf9, 0x8f, 0x5c, 0x8d, 0x37, 0x69, 0xf8,
	0xd1, 0xef, 0xff, 0x07, 0x00, 0x00, 0xff, 0xff, 0x71, 0x9f, 0xf7, 0x76, 0xe1, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CommpbClient is the client API for Commpb service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CommpbClient interface {
	AskForVote(ctx context.Context, in *VoteReq, opts ...grpc.CallOption) (*VoteRsp, error)
	HeartBeat(ctx context.Context, in *HBReq, opts ...grpc.CallOption) (*HBRsp, error)
	GetCheckPoint(ctx context.Context, in *Msg, opts ...grpc.CallOption) (*CP, error)
	Prepare(ctx context.Context, in *Commit, opts ...grpc.CallOption) (*Msg, error)
	Confirm(ctx context.Context, in *Commit, opts ...grpc.CallOption) (*Msg, error)
	Cancel(ctx context.Context, in *Commit, opts ...grpc.CallOption) (*Msg, error)
	SendToLeader(ctx context.Context, in *Commit, opts ...grpc.CallOption) (*Msg, error)
}

type commpbClient struct {
	cc grpc.ClientConnInterface
}

func NewCommpbClient(cc grpc.ClientConnInterface) CommpbClient {
	return &commpbClient{cc}
}

func (c *commpbClient) AskForVote(ctx context.Context, in *VoteReq, opts ...grpc.CallOption) (*VoteRsp, error) {
	out := new(VoteRsp)
	err := c.cc.Invoke(ctx, "/Commpb/AskForVote", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commpbClient) HeartBeat(ctx context.Context, in *HBReq, opts ...grpc.CallOption) (*HBRsp, error) {
	out := new(HBRsp)
	err := c.cc.Invoke(ctx, "/Commpb/HeartBeat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commpbClient) GetCheckPoint(ctx context.Context, in *Msg, opts ...grpc.CallOption) (*CP, error) {
	out := new(CP)
	err := c.cc.Invoke(ctx, "/Commpb/GetCheckPoint", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commpbClient) Prepare(ctx context.Context, in *Commit, opts ...grpc.CallOption) (*Msg, error) {
	out := new(Msg)
	err := c.cc.Invoke(ctx, "/Commpb/Prepare", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commpbClient) Confirm(ctx context.Context, in *Commit, opts ...grpc.CallOption) (*Msg, error) {
	out := new(Msg)
	err := c.cc.Invoke(ctx, "/Commpb/Confirm", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commpbClient) Cancel(ctx context.Context, in *Commit, opts ...grpc.CallOption) (*Msg, error) {
	out := new(Msg)
	err := c.cc.Invoke(ctx, "/Commpb/Cancel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commpbClient) SendToLeader(ctx context.Context, in *Commit, opts ...grpc.CallOption) (*Msg, error) {
	out := new(Msg)
	err := c.cc.Invoke(ctx, "/Commpb/SendToLeader", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CommpbServer is the server API for Commpb service.
type CommpbServer interface {
	AskForVote(context.Context, *VoteReq) (*VoteRsp, error)
	HeartBeat(context.Context, *HBReq) (*HBRsp, error)
	GetCheckPoint(context.Context, *Msg) (*CP, error)
	Prepare(context.Context, *Commit) (*Msg, error)
	Confirm(context.Context, *Commit) (*Msg, error)
	Cancel(context.Context, *Commit) (*Msg, error)
	SendToLeader(context.Context, *Commit) (*Msg, error)
}

// UnimplementedCommpbServer can be embedded to have forward compatible implementations.
type UnimplementedCommpbServer struct {
}

func (*UnimplementedCommpbServer) AskForVote(ctx context.Context, req *VoteReq) (*VoteRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AskForVote not implemented")
}
func (*UnimplementedCommpbServer) HeartBeat(ctx context.Context, req *HBReq) (*HBRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HeartBeat not implemented")
}
func (*UnimplementedCommpbServer) GetCheckPoint(ctx context.Context, req *Msg) (*CP, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCheckPoint not implemented")
}
func (*UnimplementedCommpbServer) Prepare(ctx context.Context, req *Commit) (*Msg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Prepare not implemented")
}
func (*UnimplementedCommpbServer) Confirm(ctx context.Context, req *Commit) (*Msg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Confirm not implemented")
}
func (*UnimplementedCommpbServer) Cancel(ctx context.Context, req *Commit) (*Msg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Cancel not implemented")
}
func (*UnimplementedCommpbServer) SendToLeader(ctx context.Context, req *Commit) (*Msg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendToLeader not implemented")
}

func RegisterCommpbServer(s *grpc.Server, srv CommpbServer) {
	s.RegisterService(&_Commpb_serviceDesc, srv)
}

func _Commpb_AskForVote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VoteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommpbServer).AskForVote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Commpb/AskForVote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommpbServer).AskForVote(ctx, req.(*VoteReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Commpb_HeartBeat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HBReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommpbServer).HeartBeat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Commpb/HeartBeat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommpbServer).HeartBeat(ctx, req.(*HBReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Commpb_GetCheckPoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Msg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommpbServer).GetCheckPoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Commpb/GetCheckPoint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommpbServer).GetCheckPoint(ctx, req.(*Msg))
	}
	return interceptor(ctx, in, info, handler)
}

func _Commpb_Prepare_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Commit)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommpbServer).Prepare(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Commpb/Prepare",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommpbServer).Prepare(ctx, req.(*Commit))
	}
	return interceptor(ctx, in, info, handler)
}

func _Commpb_Confirm_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Commit)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommpbServer).Confirm(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Commpb/Confirm",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommpbServer).Confirm(ctx, req.(*Commit))
	}
	return interceptor(ctx, in, info, handler)
}

func _Commpb_Cancel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Commit)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommpbServer).Cancel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Commpb/Cancel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommpbServer).Cancel(ctx, req.(*Commit))
	}
	return interceptor(ctx, in, info, handler)
}

func _Commpb_SendToLeader_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Commit)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommpbServer).SendToLeader(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Commpb/SendToLeader",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommpbServer).SendToLeader(ctx, req.(*Commit))
	}
	return interceptor(ctx, in, info, handler)
}

var _Commpb_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Commpb",
	HandlerType: (*CommpbServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AskForVote",
			Handler:    _Commpb_AskForVote_Handler,
		},
		{
			MethodName: "HeartBeat",
			Handler:    _Commpb_HeartBeat_Handler,
		},
		{
			MethodName: "GetCheckPoint",
			Handler:    _Commpb_GetCheckPoint_Handler,
		},
		{
			MethodName: "Prepare",
			Handler:    _Commpb_Prepare_Handler,
		},
		{
			MethodName: "Confirm",
			Handler:    _Commpb_Confirm_Handler,
		},
		{
			MethodName: "Cancel",
			Handler:    _Commpb_Cancel_Handler,
		},
		{
			MethodName: "SendToLeader",
			Handler:    _Commpb_SendToLeader_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "commpb.proto",
}
