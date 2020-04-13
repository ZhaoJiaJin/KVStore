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
	Id                   int32    `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
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

func (m *VoteReq) GetId() int32 {
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
	Id                   int32    `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
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

func (m *HBReq) GetId() int32 {
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
	Status               int32    `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
	Nodeid               int32    `protobuf:"varint,3,opt,name=nodeid,proto3" json:"nodeid,omitempty"`
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

func (m *CP) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *CP) GetNodeid() int32 {
	if m != nil {
		return m.Nodeid
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
}

func init() {
	proto.RegisterFile("commpb.proto", fileDescriptor_d2eb5c7490bed74a)
}

var fileDescriptor_d2eb5c7490bed74a = []byte{
	// 330 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xc1, 0x6a, 0xe3, 0x30,
	0x10, 0x86, 0x63, 0x3b, 0x76, 0xb2, 0x43, 0xd6, 0x1b, 0x74, 0x58, 0x8c, 0xb3, 0x0b, 0x41, 0xf4,
	0x10, 0x28, 0x15, 0x25, 0x7d, 0x82, 0xda, 0xb4, 0x49, 0x0f, 0x69, 0x8d, 0x02, 0x81, 0x1e, 0x9d,
	0x68, 0x70, 0x4d, 0xea, 0xc8, 0x95, 0x94, 0xd2, 0x47, 0xe8, 0x63, 0x17, 0x29, 0x2e, 0xf4, 0xd0,
	0x43, 0x4f, 0xfa, 0x67, 0xfe, 0x4f, 0x12, 0xff, 0x0c, 0x8c, 0x76, 0xb2, 0x69, 0xda, 0x2d, 0x6b,
	0x95, 0x34, 0x92, 0x5e, 0xc0, 0x60, 0x23, 0x0d, 0x72, 0x7c, 0x21, 0x04, 0xfa, 0x06, 0x55, 0x93,
	0x78, 0x53, 0x6f, 0x16, 0x70, 0xa7, 0x49, 0x0c, 0x7e, 0x2d, 0x12, 0x7f, 0xea, 0xcd, 0x42, 0xee,
	0xd7, 0x82, 0xde, 0x75, 0xb8, 0x6e, 0xc9, 0x19, 0x84, 0xaf, 0x46, 0xa1, 0x76, 0x7c, 0x3c, 0x8f,
	0x59, 0x67, 0xb0, 0x8d, 0xe1, 0xa8, 0xf9, 0xc9, 0xa4, 0x09, 0x84, 0xae, 0x26, 0x11, 0xf8, 0xf7,
	0x0f, 0xe3, 0x1e, 0x19, 0x40, 0xf0, 0x78, 0xb3, 0x1e, 0x7b, 0xf4, 0x1c, 0xc2, 0x65, 0xf6, 0xd3,
	0x7f, 0x27, 0x0e, 0xd6, 0xed, 0x77, 0x30, 0x0d, 0x21, 0x58, 0xe9, 0x8a, 0x2e, 0xc1, 0xcf, 0x0b,
	0x0b, 0x88, 0xd2, 0x94, 0x0e, 0x18, 0x71, 0xa7, 0xc9, 0x5f, 0x88, 0xb4, 0x29, 0xcd, 0x51, 0x77,
	0x2f, 0x76, 0x95, 0xed, 0x1f, 0xa4, 0xc0, 0x5a, 0x24, 0xc1, 0xa9, 0x7f, 0xaa, 0xe6, 0xef, 0x1e,
	0x44, 0xb9, 0x9b, 0x12, 0xa1, 0x00, 0xd7, 0x7a, 0x7f, 0x2b, 0x95, 0x4d, 0x47, 0x86, 0xac, 0x1b,
	0x56, 0x3a, 0xfc, 0x8c, 0x4b, 0x7b, 0xe4, 0x3f, 0xfc, 0x5a, 0x62, 0xa9, 0x4c, 0x86, 0xa5, 0x21,
	0x11, 0x73, 0xa9, 0x52, 0x77, 0x3a, 0xfb, 0x1f, 0xfc, 0x5e, 0xa0, 0xc9, 0x9f, 0x70, 0xb7, 0x2f,
	0x64, 0x7d, 0x30, 0xa4, 0xcf, 0x56, 0xba, 0x4a, 0x03, 0x96, 0x17, 0xee, 0x72, 0xbc, 0xc6, 0x83,
	0xf8, 0x62, 0x5b, 0x23, 0x75, 0x0c, 0xed, 0x65, 0x97, 0x30, 0xa9, 0x25, 0xab, 0x54, 0xbb, 0x63,
	0xf8, 0x56, 0x36, 0xed, 0x33, 0x6a, 0xa6, 0xe4, 0xd1, 0x60, 0x75, 0xac, 0x05, 0x66, 0x7f, 0xb8,
	0xd5, 0x0b, 0xab, 0x0b, 0xbb, 0xcf, 0xc2, 0xdb, 0x46, 0x6e, 0xb1, 0x57, 0x1f, 0x01, 0x00, 0x00,
	0xff, 0xff, 0x87, 0xf3, 0xe0, 0x4c, 0xe8, 0x01, 0x00, 0x00,
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
	SendCheckPoint(ctx context.Context, in *CP, opts ...grpc.CallOption) (*Msg, error)
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

func (c *commpbClient) SendCheckPoint(ctx context.Context, in *CP, opts ...grpc.CallOption) (*Msg, error) {
	out := new(Msg)
	err := c.cc.Invoke(ctx, "/Commpb/SendCheckPoint", in, out, opts...)
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
	SendCheckPoint(context.Context, *CP) (*Msg, error)
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
func (*UnimplementedCommpbServer) SendCheckPoint(ctx context.Context, req *CP) (*Msg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendCheckPoint not implemented")
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

func _Commpb_SendCheckPoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CP)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommpbServer).SendCheckPoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Commpb/SendCheckPoint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommpbServer).SendCheckPoint(ctx, req.(*CP))
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
			MethodName: "SendCheckPoint",
			Handler:    _Commpb_SendCheckPoint_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "commpb.proto",
}
