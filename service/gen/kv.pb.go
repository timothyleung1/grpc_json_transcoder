// Code generated by protoc-gen-go. DO NOT EDIT.
// source: kv.proto

package kv

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type GetRequest struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2216fe83c9c12408, []int{0}
}

func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequest.Unmarshal(m, b)
}
func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
}
func (m *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(m, src)
}
func (m *GetRequest) XXX_Size() int {
	return xxx_messageInfo_GetRequest.Size(m)
}
func (m *GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequest proto.InternalMessageInfo

func (m *GetRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type GetResponse struct {
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetResponse) Reset()         { *m = GetResponse{} }
func (m *GetResponse) String() string { return proto.CompactTextString(m) }
func (*GetResponse) ProtoMessage()    {}
func (*GetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2216fe83c9c12408, []int{1}
}

func (m *GetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetResponse.Unmarshal(m, b)
}
func (m *GetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetResponse.Marshal(b, m, deterministic)
}
func (m *GetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetResponse.Merge(m, src)
}
func (m *GetResponse) XXX_Size() int {
	return xxx_messageInfo_GetResponse.Size(m)
}
func (m *GetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetResponse proto.InternalMessageInfo

func (m *GetResponse) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type SetRequest struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetRequest) Reset()         { *m = SetRequest{} }
func (m *SetRequest) String() string { return proto.CompactTextString(m) }
func (*SetRequest) ProtoMessage()    {}
func (*SetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2216fe83c9c12408, []int{2}
}

func (m *SetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetRequest.Unmarshal(m, b)
}
func (m *SetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetRequest.Marshal(b, m, deterministic)
}
func (m *SetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetRequest.Merge(m, src)
}
func (m *SetRequest) XXX_Size() int {
	return xxx_messageInfo_SetRequest.Size(m)
}
func (m *SetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetRequest proto.InternalMessageInfo

func (m *SetRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *SetRequest) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type SetResponse struct {
	Ok                   bool     `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetResponse) Reset()         { *m = SetResponse{} }
func (m *SetResponse) String() string { return proto.CompactTextString(m) }
func (*SetResponse) ProtoMessage()    {}
func (*SetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2216fe83c9c12408, []int{3}
}

func (m *SetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetResponse.Unmarshal(m, b)
}
func (m *SetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetResponse.Marshal(b, m, deterministic)
}
func (m *SetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetResponse.Merge(m, src)
}
func (m *SetResponse) XXX_Size() int {
	return xxx_messageInfo_SetResponse.Size(m)
}
func (m *SetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SetResponse proto.InternalMessageInfo

func (m *SetResponse) GetOk() bool {
	if m != nil {
		return m.Ok
	}
	return false
}

type CountRequest struct {
	Count                uint32   `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CountRequest) Reset()         { *m = CountRequest{} }
func (m *CountRequest) String() string { return proto.CompactTextString(m) }
func (*CountRequest) ProtoMessage()    {}
func (*CountRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2216fe83c9c12408, []int{4}
}

func (m *CountRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CountRequest.Unmarshal(m, b)
}
func (m *CountRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CountRequest.Marshal(b, m, deterministic)
}
func (m *CountRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CountRequest.Merge(m, src)
}
func (m *CountRequest) XXX_Size() int {
	return xxx_messageInfo_CountRequest.Size(m)
}
func (m *CountRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CountRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CountRequest proto.InternalMessageInfo

func (m *CountRequest) GetCount() uint32 {
	if m != nil {
		return m.Count
	}
	return 0
}

type CountResponse struct {
	Count                uint32   `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CountResponse) Reset()         { *m = CountResponse{} }
func (m *CountResponse) String() string { return proto.CompactTextString(m) }
func (*CountResponse) ProtoMessage()    {}
func (*CountResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2216fe83c9c12408, []int{5}
}

func (m *CountResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CountResponse.Unmarshal(m, b)
}
func (m *CountResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CountResponse.Marshal(b, m, deterministic)
}
func (m *CountResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CountResponse.Merge(m, src)
}
func (m *CountResponse) XXX_Size() int {
	return xxx_messageInfo_CountResponse.Size(m)
}
func (m *CountResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CountResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CountResponse proto.InternalMessageInfo

func (m *CountResponse) GetCount() uint32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func init() {
	proto.RegisterType((*GetRequest)(nil), "kv.GetRequest")
	proto.RegisterType((*GetResponse)(nil), "kv.GetResponse")
	proto.RegisterType((*SetRequest)(nil), "kv.SetRequest")
	proto.RegisterType((*SetResponse)(nil), "kv.SetResponse")
	proto.RegisterType((*CountRequest)(nil), "kv.CountRequest")
	proto.RegisterType((*CountResponse)(nil), "kv.CountResponse")
}

func init() { proto.RegisterFile("kv.proto", fileDescriptor_2216fe83c9c12408) }

var fileDescriptor_2216fe83c9c12408 = []byte{
	// 288 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xc8, 0x2e, 0xd3, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xca, 0x2e, 0x93, 0x92, 0x49, 0xcf, 0xcf, 0x4f, 0xcf, 0x49,
	0xd5, 0x4f, 0x2c, 0xc8, 0xd4, 0x4f, 0xcc, 0xcb, 0xcb, 0x2f, 0x49, 0x2c, 0xc9, 0xcc, 0xcf, 0x2b,
	0x86, 0xa8, 0x50, 0x92, 0xe3, 0xe2, 0x72, 0x4f, 0x2d, 0x09, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e,
	0x11, 0x12, 0xe0, 0x62, 0xce, 0x4e, 0xad, 0x94, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0x31,
	0x95, 0x94, 0xb9, 0xb8, 0xc1, 0xf2, 0xc5, 0x05, 0xf9, 0x79, 0xc5, 0xa9, 0x42, 0x22, 0x5c, 0xac,
	0x65, 0x89, 0x39, 0xa5, 0xa9, 0x50, 0x25, 0x10, 0x8e, 0x92, 0x09, 0x17, 0x57, 0x30, 0x1e, 0x43,
	0x10, 0xba, 0x98, 0x90, 0x75, 0xc9, 0x72, 0x71, 0x07, 0x23, 0x19, 0xcd, 0xc7, 0xc5, 0x94, 0x9f,
	0x0d, 0xd6, 0xc5, 0x11, 0xc4, 0x94, 0x9f, 0xad, 0xa4, 0xc2, 0xc5, 0xe3, 0x9c, 0x5f, 0x9a, 0x07,
	0x37, 0x56, 0x84, 0x8b, 0x35, 0x19, 0xc4, 0x07, 0x2b, 0xe1, 0x0d, 0x82, 0x70, 0x94, 0x54, 0xb9,
	0x78, 0xa1, 0xaa, 0x10, 0x2e, 0xc4, 0x54, 0x66, 0xb4, 0x97, 0x91, 0x8b, 0xc9, 0x3b, 0x4c, 0xc8,
	0x92, 0x8b, 0xd9, 0x3d, 0xb5, 0x44, 0x88, 0x4f, 0x2f, 0xbb, 0x4c, 0x0f, 0xe1, 0x6d, 0x29, 0x7e,
	0x38, 0x1f, 0x62, 0x88, 0x92, 0x60, 0xd3, 0xe5, 0x27, 0x93, 0x99, 0xb8, 0x85, 0x38, 0xf5, 0xb3,
	0xcb, 0xf4, 0xab, 0xb3, 0x53, 0x2b, 0x6b, 0x85, 0x6c, 0xb8, 0x98, 0x83, 0x61, 0x5a, 0x83, 0xd1,
	0xb4, 0x22, 0x79, 0x43, 0x49, 0x04, 0xac, 0x95, 0x4f, 0x0a, 0xa1, 0xd5, 0x8a, 0x51, 0x4b, 0xc8,
	0x8e, 0x8b, 0x15, 0xec, 0x4c, 0x21, 0x01, 0x90, 0x7a, 0x64, 0x7f, 0x49, 0x09, 0x22, 0x89, 0x40,
	0xcd, 0xe0, 0x03, 0x9b, 0xc1, 0x21, 0xc4, 0xa6, 0x0f, 0x76, 0xbd, 0x01, 0x63, 0x12, 0x1b, 0x38,
	0xb6, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xe6, 0x32, 0x27, 0xc9, 0xdb, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// KVClient is the client API for KV service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type KVClient interface {
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	Set(ctx context.Context, in *SetRequest, opts ...grpc.CallOption) (*SetResponse, error)
	Count(ctx context.Context, in *CountRequest, opts ...grpc.CallOption) (KV_CountClient, error)
}

type kVClient struct {
	cc grpc.ClientConnInterface
}

func NewKVClient(cc grpc.ClientConnInterface) KVClient {
	return &kVClient{cc}
}

func (c *kVClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/kv.KV/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kVClient) Set(ctx context.Context, in *SetRequest, opts ...grpc.CallOption) (*SetResponse, error) {
	out := new(SetResponse)
	err := c.cc.Invoke(ctx, "/kv.KV/Set", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kVClient) Count(ctx context.Context, in *CountRequest, opts ...grpc.CallOption) (KV_CountClient, error) {
	stream, err := c.cc.NewStream(ctx, &_KV_serviceDesc.Streams[0], "/kv.KV/Count", opts...)
	if err != nil {
		return nil, err
	}
	x := &kVCountClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type KV_CountClient interface {
	Recv() (*CountResponse, error)
	grpc.ClientStream
}

type kVCountClient struct {
	grpc.ClientStream
}

func (x *kVCountClient) Recv() (*CountResponse, error) {
	m := new(CountResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// KVServer is the server API for KV service.
type KVServer interface {
	Get(context.Context, *GetRequest) (*GetResponse, error)
	Set(context.Context, *SetRequest) (*SetResponse, error)
	Count(*CountRequest, KV_CountServer) error
}

// UnimplementedKVServer can be embedded to have forward compatible implementations.
type UnimplementedKVServer struct {
}

func (*UnimplementedKVServer) Get(ctx context.Context, req *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedKVServer) Set(ctx context.Context, req *SetRequest) (*SetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Set not implemented")
}
func (*UnimplementedKVServer) Count(req *CountRequest, srv KV_CountServer) error {
	return status.Errorf(codes.Unimplemented, "method Count not implemented")
}

func RegisterKVServer(s *grpc.Server, srv KVServer) {
	s.RegisterService(&_KV_serviceDesc, srv)
}

func _KV_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KVServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kv.KV/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KVServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KV_Set_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KVServer).Set(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kv.KV/Set",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KVServer).Set(ctx, req.(*SetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KV_Count_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(CountRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(KVServer).Count(m, &kVCountServer{stream})
}

type KV_CountServer interface {
	Send(*CountResponse) error
	grpc.ServerStream
}

type kVCountServer struct {
	grpc.ServerStream
}

func (x *kVCountServer) Send(m *CountResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _KV_serviceDesc = grpc.ServiceDesc{
	ServiceName: "kv.KV",
	HandlerType: (*KVServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _KV_Get_Handler,
		},
		{
			MethodName: "Set",
			Handler:    _KV_Set_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Count",
			Handler:       _KV_Count_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "kv.proto",
}