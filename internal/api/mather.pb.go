// Code generated by protoc-gen-go. DO NOT EDIT.
// source: mather.proto

package api

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

type MathRequest struct {
	A                    int64    `protobuf:"varint,1,opt,name=a,proto3" json:"a,omitempty"`
	B                    int64    `protobuf:"varint,2,opt,name=b,proto3" json:"b,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MathRequest) Reset()         { *m = MathRequest{} }
func (m *MathRequest) String() string { return proto.CompactTextString(m) }
func (*MathRequest) ProtoMessage()    {}
func (*MathRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f08a2222f3f8c84a, []int{0}
}

func (m *MathRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MathRequest.Unmarshal(m, b)
}
func (m *MathRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MathRequest.Marshal(b, m, deterministic)
}
func (m *MathRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MathRequest.Merge(m, src)
}
func (m *MathRequest) XXX_Size() int {
	return xxx_messageInfo_MathRequest.Size(m)
}
func (m *MathRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MathRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MathRequest proto.InternalMessageInfo

func (m *MathRequest) GetA() int64 {
	if m != nil {
		return m.A
	}
	return 0
}

func (m *MathRequest) GetB() int64 {
	if m != nil {
		return m.B
	}
	return 0
}

type NameRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NameRequest) Reset()         { *m = NameRequest{} }
func (m *NameRequest) String() string { return proto.CompactTextString(m) }
func (*NameRequest) ProtoMessage()    {}
func (*NameRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f08a2222f3f8c84a, []int{1}
}

func (m *NameRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NameRequest.Unmarshal(m, b)
}
func (m *NameRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NameRequest.Marshal(b, m, deterministic)
}
func (m *NameRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NameRequest.Merge(m, src)
}
func (m *NameRequest) XXX_Size() int {
	return xxx_messageInfo_NameRequest.Size(m)
}
func (m *NameRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NameRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NameRequest proto.InternalMessageInfo

func (m *NameRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

//TODO: These could be a single message with a oneof
type MathIntResponse struct {
	Result               int64    `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MathIntResponse) Reset()         { *m = MathIntResponse{} }
func (m *MathIntResponse) String() string { return proto.CompactTextString(m) }
func (*MathIntResponse) ProtoMessage()    {}
func (*MathIntResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f08a2222f3f8c84a, []int{2}
}

func (m *MathIntResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MathIntResponse.Unmarshal(m, b)
}
func (m *MathIntResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MathIntResponse.Marshal(b, m, deterministic)
}
func (m *MathIntResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MathIntResponse.Merge(m, src)
}
func (m *MathIntResponse) XXX_Size() int {
	return xxx_messageInfo_MathIntResponse.Size(m)
}
func (m *MathIntResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MathIntResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MathIntResponse proto.InternalMessageInfo

func (m *MathIntResponse) GetResult() int64 {
	if m != nil {
		return m.Result
	}
	return 0
}

type MathFloatResponse struct {
	Result               float64  `protobuf:"fixed64,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MathFloatResponse) Reset()         { *m = MathFloatResponse{} }
func (m *MathFloatResponse) String() string { return proto.CompactTextString(m) }
func (*MathFloatResponse) ProtoMessage()    {}
func (*MathFloatResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f08a2222f3f8c84a, []int{3}
}

func (m *MathFloatResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MathFloatResponse.Unmarshal(m, b)
}
func (m *MathFloatResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MathFloatResponse.Marshal(b, m, deterministic)
}
func (m *MathFloatResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MathFloatResponse.Merge(m, src)
}
func (m *MathFloatResponse) XXX_Size() int {
	return xxx_messageInfo_MathFloatResponse.Size(m)
}
func (m *MathFloatResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MathFloatResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MathFloatResponse proto.InternalMessageInfo

func (m *MathFloatResponse) GetResult() float64 {
	if m != nil {
		return m.Result
	}
	return 0
}

type HelloResponse struct {
	Greeting             string   `protobuf:"bytes,1,opt,name=greeting,proto3" json:"greeting,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloResponse) Reset()         { *m = HelloResponse{} }
func (m *HelloResponse) String() string { return proto.CompactTextString(m) }
func (*HelloResponse) ProtoMessage()    {}
func (*HelloResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f08a2222f3f8c84a, []int{4}
}

func (m *HelloResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloResponse.Unmarshal(m, b)
}
func (m *HelloResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloResponse.Marshal(b, m, deterministic)
}
func (m *HelloResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloResponse.Merge(m, src)
}
func (m *HelloResponse) XXX_Size() int {
	return xxx_messageInfo_HelloResponse.Size(m)
}
func (m *HelloResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HelloResponse proto.InternalMessageInfo

func (m *HelloResponse) GetGreeting() string {
	if m != nil {
		return m.Greeting
	}
	return ""
}

func init() {
	proto.RegisterType((*MathRequest)(nil), "api.MathRequest")
	proto.RegisterType((*NameRequest)(nil), "api.NameRequest")
	proto.RegisterType((*MathIntResponse)(nil), "api.MathIntResponse")
	proto.RegisterType((*MathFloatResponse)(nil), "api.MathFloatResponse")
	proto.RegisterType((*HelloResponse)(nil), "api.HelloResponse")
}

func init() { proto.RegisterFile("mather.proto", fileDescriptor_f08a2222f3f8c84a) }

var fileDescriptor_f08a2222f3f8c84a = []byte{
	// 265 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0x93, 0x46, 0x43, 0x9d, 0x56, 0xd4, 0x45, 0x4a, 0xc9, 0x49, 0xf7, 0x64, 0x11, 0x52,
	0x50, 0x5f, 0x40, 0x10, 0xd1, 0x43, 0x3c, 0xc4, 0x27, 0x98, 0x98, 0xa1, 0x5d, 0xd8, 0x64, 0xd7,
	0xcd, 0x44, 0xf0, 0x1d, 0x7c, 0x68, 0xc9, 0x1a, 0x43, 0x0a, 0x7a, 0xe8, 0x6d, 0xff, 0x9f, 0xef,
	0x5f, 0x66, 0xfe, 0x81, 0x79, 0x85, 0xbc, 0x25, 0x97, 0x5a, 0x67, 0xd8, 0x88, 0x08, 0xad, 0x92,
	0x2b, 0x98, 0x65, 0xc8, 0xdb, 0x9c, 0xde, 0x5b, 0x6a, 0x58, 0xcc, 0x21, 0xc4, 0x65, 0x78, 0x11,
	0x5e, 0x45, 0x79, 0x88, 0x9d, 0x2a, 0x96, 0x93, 0x1f, 0x55, 0xc8, 0x4b, 0x98, 0xbd, 0x60, 0x45,
	0xbf, 0xa8, 0x80, 0x83, 0x1a, 0x2b, 0xf2, 0xf4, 0x51, 0xee, 0xdf, 0x72, 0x05, 0x27, 0xdd, 0x6f,
	0xcf, 0x35, 0xe7, 0xd4, 0x58, 0x53, 0x37, 0x24, 0x16, 0x10, 0x3b, 0x6a, 0x5a, 0xcd, 0xfd, 0xb7,
	0xbd, 0x92, 0xd7, 0x70, 0xd6, 0xa1, 0x8f, 0xda, 0xe0, 0x7f, 0x70, 0x38, 0x82, 0x8f, 0x9f, 0x48,
	0x6b, 0x33, 0x80, 0x09, 0x4c, 0x37, 0x8e, 0x88, 0x55, 0xbd, 0xe9, 0x07, 0x18, 0xf4, 0xcd, 0xd7,
	0x04, 0xe2, 0xcc, 0x2f, 0x2a, 0xd6, 0x10, 0xdd, 0x97, 0xa5, 0x38, 0x4d, 0xd1, 0xaa, 0x74, 0xb4,
	0x67, 0x72, 0x3e, 0x38, 0xa3, 0x59, 0x65, 0x20, 0xee, 0x60, 0xfa, 0xda, 0x16, 0xec, 0xf0, 0x8d,
	0xf7, 0x4b, 0x65, 0xad, 0x66, 0x65, 0xf5, 0xe7, 0x5e, 0xa9, 0xf8, 0x41, 0x7d, 0xa8, 0x92, 0xfe,
	0xc8, 0x2c, 0x06, 0x67, 0xa7, 0x20, 0x19, 0x88, 0x35, 0x1c, 0xfa, 0x2a, 0xfa, 0xd0, 0xe8, 0x22,
	0x89, 0xf0, 0xce, 0x4e, 0x51, 0x32, 0x28, 0x62, 0x7f, 0xed, 0xdb, 0xef, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x3e, 0xbf, 0xe1, 0x19, 0xfd, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MatherClient is the client API for Mather service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MatherClient interface {
	Add(ctx context.Context, in *MathRequest, opts ...grpc.CallOption) (*MathIntResponse, error)
	Subtract(ctx context.Context, in *MathRequest, opts ...grpc.CallOption) (*MathIntResponse, error)
	Multiply(ctx context.Context, in *MathRequest, opts ...grpc.CallOption) (*MathIntResponse, error)
	Divide(ctx context.Context, in *MathRequest, opts ...grpc.CallOption) (*MathFloatResponse, error)
	Hello(ctx context.Context, in *NameRequest, opts ...grpc.CallOption) (*HelloResponse, error)
}

type matherClient struct {
	cc *grpc.ClientConn
}

func NewMatherClient(cc *grpc.ClientConn) MatherClient {
	return &matherClient{cc}
}

func (c *matherClient) Add(ctx context.Context, in *MathRequest, opts ...grpc.CallOption) (*MathIntResponse, error) {
	out := new(MathIntResponse)
	err := c.cc.Invoke(ctx, "/api.Mather/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *matherClient) Subtract(ctx context.Context, in *MathRequest, opts ...grpc.CallOption) (*MathIntResponse, error) {
	out := new(MathIntResponse)
	err := c.cc.Invoke(ctx, "/api.Mather/Subtract", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *matherClient) Multiply(ctx context.Context, in *MathRequest, opts ...grpc.CallOption) (*MathIntResponse, error) {
	out := new(MathIntResponse)
	err := c.cc.Invoke(ctx, "/api.Mather/Multiply", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *matherClient) Divide(ctx context.Context, in *MathRequest, opts ...grpc.CallOption) (*MathFloatResponse, error) {
	out := new(MathFloatResponse)
	err := c.cc.Invoke(ctx, "/api.Mather/Divide", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *matherClient) Hello(ctx context.Context, in *NameRequest, opts ...grpc.CallOption) (*HelloResponse, error) {
	out := new(HelloResponse)
	err := c.cc.Invoke(ctx, "/api.Mather/Hello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MatherServer is the server API for Mather service.
type MatherServer interface {
	Add(context.Context, *MathRequest) (*MathIntResponse, error)
	Subtract(context.Context, *MathRequest) (*MathIntResponse, error)
	Multiply(context.Context, *MathRequest) (*MathIntResponse, error)
	Divide(context.Context, *MathRequest) (*MathFloatResponse, error)
	Hello(context.Context, *NameRequest) (*HelloResponse, error)
}

func RegisterMatherServer(s *grpc.Server, srv MatherServer) {
	s.RegisterService(&_Mather_serviceDesc, srv)
}

func _Mather_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MathRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MatherServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Mather/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MatherServer).Add(ctx, req.(*MathRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mather_Subtract_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MathRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MatherServer).Subtract(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Mather/Subtract",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MatherServer).Subtract(ctx, req.(*MathRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mather_Multiply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MathRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MatherServer).Multiply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Mather/Multiply",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MatherServer).Multiply(ctx, req.(*MathRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mather_Divide_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MathRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MatherServer).Divide(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Mather/Divide",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MatherServer).Divide(ctx, req.(*MathRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mather_Hello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MatherServer).Hello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Mather/Hello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MatherServer).Hello(ctx, req.(*NameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Mather_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.Mather",
	HandlerType: (*MatherServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _Mather_Add_Handler,
		},
		{
			MethodName: "Subtract",
			Handler:    _Mather_Subtract_Handler,
		},
		{
			MethodName: "Multiply",
			Handler:    _Mather_Multiply_Handler,
		},
		{
			MethodName: "Divide",
			Handler:    _Mather_Divide_Handler,
		},
		{
			MethodName: "Hello",
			Handler:    _Mather_Hello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mather.proto",
}