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

func init() {
	proto.RegisterEnum("VoteRsp_VtRes", VoteRsp_VtRes_name, VoteRsp_VtRes_value)
	proto.RegisterType((*VoteReq)(nil), "VoteReq")
	proto.RegisterType((*VoteRsp)(nil), "VoteRsp")
	proto.RegisterType((*HBReq)(nil), "HBReq")
	proto.RegisterType((*HBRsp)(nil), "HBRsp")
}

func init() {
	proto.RegisterFile("commpb.proto", fileDescriptor_d2eb5c7490bed74a)
}

var fileDescriptor_d2eb5c7490bed74a = []byte{
	// 248 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0x51, 0x4b, 0xc3, 0x30,
	0x14, 0x85, 0x97, 0xce, 0x76, 0xf3, 0x22, 0x75, 0xdc, 0xa7, 0xb2, 0x21, 0x8c, 0xe0, 0xc3, 0x40,
	0x0c, 0x32, 0x7f, 0x81, 0x15, 0x75, 0x22, 0xe8, 0x88, 0x30, 0xf0, 0xb1, 0x5b, 0x2f, 0xa3, 0x68,
	0x49, 0x4c, 0x52, 0xf1, 0xe7, 0x4b, 0xae, 0xf5, 0xcd, 0x87, 0x3d, 0xe5, 0x84, 0xf3, 0xc1, 0x3d,
	0xe7, 0xc0, 0xc9, 0xce, 0xb4, 0xad, 0xdd, 0x2a, 0xeb, 0x4c, 0x30, 0xf2, 0x12, 0x46, 0x1b, 0x13,
	0x48, 0xd3, 0x27, 0x22, 0x1c, 0x05, 0x72, 0x6d, 0x21, 0xe6, 0x62, 0x31, 0xd4, 0xac, 0x31, 0x87,
	0xa4, 0xa9, 0x8b, 0x64, 0x2e, 0x16, 0xa9, 0x4e, 0x9a, 0x5a, 0x3e, 0xf6, 0xb8, 0xb7, 0x78, 0x0e,
	0xe9, 0x57, 0x70, 0xe4, 0x99, 0xcf, 0x97, 0xb9, 0xea, 0x0d, 0xb5, 0x09, 0x9a, 0xbc, 0xfe, 0x35,
	0x65, 0x01, 0x29, 0xff, 0x31, 0x83, 0xe4, 0xf9, 0x65, 0x32, 0xc0, 0x11, 0x0c, 0xdf, 0xee, 0x5e,
	0x27, 0x42, 0x5e, 0x40, 0xba, 0x2a, 0x0f, 0xbd, 0x3b, 0x63, 0xd8, 0xdb, 0xff, 0xe0, 0xe5, 0x13,
	0x64, 0xb7, 0xdc, 0x09, 0x25, 0xc0, 0x8d, 0x7f, 0xbf, 0x37, 0x2e, 0x66, 0xc1, 0xb1, 0xea, 0xab,
	0x4d, 0xc7, 0x7f, 0xe1, 0xe4, 0x00, 0xcf, 0xe0, 0x78, 0x45, 0x95, 0x0b, 0x25, 0x55, 0x01, 0x33,
	0xc5, 0x19, 0xa6, 0xfc, 0x46, 0xbb, 0xbc, 0x82, 0x59, 0x63, 0xd4, 0xde, 0xd9, 0x9d, 0xa2, 0xef,
	0xaa, 0xb5, 0x1f, 0xe4, 0x95, 0x33, 0x5d, 0xa0, 0x7d, 0xd7, 0xd4, 0x54, 0x9e, 0xea, 0xa8, 0x1f,
	0xa2, 0x5e, 0xc7, 0x01, 0xd7, 0x62, 0x9b, 0xf1, 0x92, 0xd7, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff,
	0xc5, 0x77, 0xab, 0xe2, 0x59, 0x01, 0x00, 0x00,
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

// CommpbServer is the server API for Commpb service.
type CommpbServer interface {
	AskForVote(context.Context, *VoteReq) (*VoteRsp, error)
	HeartBeat(context.Context, *HBReq) (*HBRsp, error)
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
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "commpb.proto",
}
