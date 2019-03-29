// Code generated by protoc-gen-go. DO NOT EDIT.
// source: esa2.proto

package esa2

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

type EhloMessage struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EhloMessage) Reset()         { *m = EhloMessage{} }
func (m *EhloMessage) String() string { return proto.CompactTextString(m) }
func (*EhloMessage) ProtoMessage()    {}
func (*EhloMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_523060ba8a6bed6e, []int{0}
}

func (m *EhloMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EhloMessage.Unmarshal(m, b)
}
func (m *EhloMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EhloMessage.Marshal(b, m, deterministic)
}
func (m *EhloMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EhloMessage.Merge(m, src)
}
func (m *EhloMessage) XXX_Size() int {
	return xxx_messageInfo_EhloMessage.Size(m)
}
func (m *EhloMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_EhloMessage.DiscardUnknown(m)
}

var xxx_messageInfo_EhloMessage proto.InternalMessageInfo

func (m *EhloMessage) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*EhloMessage)(nil), "sample.endsrvc.EhloMessage")
}

func init() { proto.RegisterFile("esa2.proto", fileDescriptor_523060ba8a6bed6e) }

var fileDescriptor_523060ba8a6bed6e = []byte{
	// 125 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0x2d, 0x4e, 0x34,
	0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x2b, 0x4e, 0xcc, 0x2d, 0xc8, 0x49, 0xd5, 0x4b,
	0xcd, 0x4b, 0x29, 0x2e, 0x2a, 0x4b, 0x56, 0x92, 0xe7, 0xe2, 0x76, 0xcd, 0xc8, 0xc9, 0xf7, 0x4d,
	0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0x15, 0x12, 0xe0, 0x62, 0xce, 0x2d, 0x4e, 0x97, 0x60, 0x54, 0x60,
	0xd4, 0xe0, 0x0c, 0x02, 0x31, 0x8d, 0xbc, 0xb9, 0x58, 0x5c, 0x83, 0x1d, 0x8d, 0x84, 0x9c, 0xb9,
	0x58, 0x3d, 0x52, 0x73, 0x72, 0xf2, 0x85, 0xa4, 0xf5, 0x50, 0x8d, 0xd0, 0x43, 0xd2, 0x2f, 0x85,
	0x4f, 0x52, 0x89, 0xc1, 0x89, 0x2d, 0x8a, 0x05, 0xe4, 0x96, 0x24, 0x36, 0xb0, 0x63, 0x8c, 0x01,
	0x01, 0x00, 0x00, 0xff, 0xff, 0xf8, 0x3b, 0xc4, 0xe2, 0x9a, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ESA2Client is the client API for ESA2 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ESA2Client interface {
	Hello(ctx context.Context, in *EhloMessage, opts ...grpc.CallOption) (*EhloMessage, error)
}

type eSA2Client struct {
	cc *grpc.ClientConn
}

func NewESA2Client(cc *grpc.ClientConn) ESA2Client {
	return &eSA2Client{cc}
}

func (c *eSA2Client) Hello(ctx context.Context, in *EhloMessage, opts ...grpc.CallOption) (*EhloMessage, error) {
	out := new(EhloMessage)
	err := c.cc.Invoke(ctx, "/sample.endsrvc.ESA2/Hello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ESA2Server is the server API for ESA2 service.
type ESA2Server interface {
	Hello(context.Context, *EhloMessage) (*EhloMessage, error)
}

func RegisterESA2Server(s *grpc.Server, srv ESA2Server) {
	s.RegisterService(&_ESA2_serviceDesc, srv)
}

func _ESA2_Hello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EhloMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ESA2Server).Hello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sample.endsrvc.ESA2/Hello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ESA2Server).Hello(ctx, req.(*EhloMessage))
	}
	return interceptor(ctx, in, info, handler)
}

var _ESA2_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sample.endsrvc.ESA2",
	HandlerType: (*ESA2Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Hello",
			Handler:    _ESA2_Hello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "esa2.proto",
}
