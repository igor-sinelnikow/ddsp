// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

type GetRequest struct {
	Key                  uint32   `protobuf:"varint,1,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_pb_fd9dd68e056fbc00, []int{0}
}
func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequest.Unmarshal(m, b)
}
func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
}
func (dst *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(dst, src)
}
func (m *GetRequest) XXX_Size() int {
	return xxx_messageInfo_GetRequest.Size(m)
}
func (m *GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequest proto.InternalMessageInfo

func (m *GetRequest) GetKey() uint32 {
	if m != nil {
		return m.Key
	}
	return 0
}

type GetReply struct {
	Status               int32    `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Error                string   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Data                 []byte   `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetReply) Reset()         { *m = GetReply{} }
func (m *GetReply) String() string { return proto.CompactTextString(m) }
func (*GetReply) ProtoMessage()    {}
func (*GetReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_pb_fd9dd68e056fbc00, []int{1}
}
func (m *GetReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetReply.Unmarshal(m, b)
}
func (m *GetReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetReply.Marshal(b, m, deterministic)
}
func (dst *GetReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetReply.Merge(dst, src)
}
func (m *GetReply) XXX_Size() int {
	return xxx_messageInfo_GetReply.Size(m)
}
func (m *GetReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GetReply.DiscardUnknown(m)
}

var xxx_messageInfo_GetReply proto.InternalMessageInfo

func (m *GetReply) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *GetReply) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *GetReply) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type PutRequest struct {
	Key                  uint32   `protobuf:"varint,1,opt,name=key,proto3" json:"key,omitempty"`
	Data                 []byte   `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PutRequest) Reset()         { *m = PutRequest{} }
func (m *PutRequest) String() string { return proto.CompactTextString(m) }
func (*PutRequest) ProtoMessage()    {}
func (*PutRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_pb_fd9dd68e056fbc00, []int{2}
}
func (m *PutRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PutRequest.Unmarshal(m, b)
}
func (m *PutRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PutRequest.Marshal(b, m, deterministic)
}
func (dst *PutRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PutRequest.Merge(dst, src)
}
func (m *PutRequest) XXX_Size() int {
	return xxx_messageInfo_PutRequest.Size(m)
}
func (m *PutRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PutRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PutRequest proto.InternalMessageInfo

func (m *PutRequest) GetKey() uint32 {
	if m != nil {
		return m.Key
	}
	return 0
}

func (m *PutRequest) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type PutReply struct {
	Status               int32    `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Error                string   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PutReply) Reset()         { *m = PutReply{} }
func (m *PutReply) String() string { return proto.CompactTextString(m) }
func (*PutReply) ProtoMessage()    {}
func (*PutReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_pb_fd9dd68e056fbc00, []int{3}
}
func (m *PutReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PutReply.Unmarshal(m, b)
}
func (m *PutReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PutReply.Marshal(b, m, deterministic)
}
func (dst *PutReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PutReply.Merge(dst, src)
}
func (m *PutReply) XXX_Size() int {
	return xxx_messageInfo_PutReply.Size(m)
}
func (m *PutReply) XXX_DiscardUnknown() {
	xxx_messageInfo_PutReply.DiscardUnknown(m)
}

var xxx_messageInfo_PutReply proto.InternalMessageInfo

func (m *PutReply) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *PutReply) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type DelRequest struct {
	Key                  uint32   `protobuf:"varint,1,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DelRequest) Reset()         { *m = DelRequest{} }
func (m *DelRequest) String() string { return proto.CompactTextString(m) }
func (*DelRequest) ProtoMessage()    {}
func (*DelRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_pb_fd9dd68e056fbc00, []int{4}
}
func (m *DelRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DelRequest.Unmarshal(m, b)
}
func (m *DelRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DelRequest.Marshal(b, m, deterministic)
}
func (dst *DelRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DelRequest.Merge(dst, src)
}
func (m *DelRequest) XXX_Size() int {
	return xxx_messageInfo_DelRequest.Size(m)
}
func (m *DelRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DelRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DelRequest proto.InternalMessageInfo

func (m *DelRequest) GetKey() uint32 {
	if m != nil {
		return m.Key
	}
	return 0
}

type DelReply struct {
	Status               int32    `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Error                string   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DelReply) Reset()         { *m = DelReply{} }
func (m *DelReply) String() string { return proto.CompactTextString(m) }
func (*DelReply) ProtoMessage()    {}
func (*DelReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_pb_fd9dd68e056fbc00, []int{5}
}
func (m *DelReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DelReply.Unmarshal(m, b)
}
func (m *DelReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DelReply.Marshal(b, m, deterministic)
}
func (dst *DelReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DelReply.Merge(dst, src)
}
func (m *DelReply) XXX_Size() int {
	return xxx_messageInfo_DelReply.Size(m)
}
func (m *DelReply) XXX_DiscardUnknown() {
	xxx_messageInfo_DelReply.DiscardUnknown(m)
}

var xxx_messageInfo_DelReply proto.InternalMessageInfo

func (m *DelReply) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *DelReply) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func init() {
	proto.RegisterType((*GetRequest)(nil), "GetRequest")
	proto.RegisterType((*GetReply)(nil), "GetReply")
	proto.RegisterType((*PutRequest)(nil), "PutRequest")
	proto.RegisterType((*PutReply)(nil), "PutReply")
	proto.RegisterType((*DelRequest)(nil), "DelRequest")
	proto.RegisterType((*DelReply)(nil), "DelReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// StorageClient is the client API for Storage service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StorageClient interface {
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetReply, error)
	Put(ctx context.Context, in *PutRequest, opts ...grpc.CallOption) (*PutReply, error)
	Del(ctx context.Context, in *DelRequest, opts ...grpc.CallOption) (*DelReply, error)
}

type storageClient struct {
	cc *grpc.ClientConn
}

func NewStorageClient(cc *grpc.ClientConn) StorageClient {
	return &storageClient{cc}
}

func (c *storageClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetReply, error) {
	out := new(GetReply)
	err := c.cc.Invoke(ctx, "/Storage/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageClient) Put(ctx context.Context, in *PutRequest, opts ...grpc.CallOption) (*PutReply, error) {
	out := new(PutReply)
	err := c.cc.Invoke(ctx, "/Storage/Put", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageClient) Del(ctx context.Context, in *DelRequest, opts ...grpc.CallOption) (*DelReply, error) {
	out := new(DelReply)
	err := c.cc.Invoke(ctx, "/Storage/Del", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StorageServer is the server API for Storage service.
type StorageServer interface {
	Get(context.Context, *GetRequest) (*GetReply, error)
	Put(context.Context, *PutRequest) (*PutReply, error)
	Del(context.Context, *DelRequest) (*DelReply, error)
}

func RegisterStorageServer(s *grpc.Server, srv StorageServer) {
	s.RegisterService(&_Storage_serviceDesc, srv)
}

func _Storage_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Storage/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Storage_Put_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServer).Put(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Storage/Put",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServer).Put(ctx, req.(*PutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Storage_Del_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServer).Del(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Storage/Del",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServer).Del(ctx, req.(*DelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Storage_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Storage",
	HandlerType: (*StorageServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _Storage_Get_Handler,
		},
		{
			MethodName: "Put",
			Handler:    _Storage_Put_Handler,
		},
		{
			MethodName: "Del",
			Handler:    _Storage_Del_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb.proto",
}

func init() { proto.RegisterFile("pb.proto", fileDescriptor_pb_fd9dd68e056fbc00) }

var fileDescriptor_pb_fd9dd68e056fbc00 = []byte{
	// 226 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x28, 0x48, 0xd2, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x57, 0x92, 0xe3, 0xe2, 0x72, 0x4f, 0x2d, 0x09, 0x4a, 0x2d, 0x2c, 0x4d,
	0x2d, 0x2e, 0x11, 0x12, 0xe0, 0x62, 0xce, 0x4e, 0xad, 0x94, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0d,
	0x02, 0x31, 0x95, 0x7c, 0xb8, 0x38, 0xc0, 0xf2, 0x05, 0x39, 0x95, 0x42, 0x62, 0x5c, 0x6c, 0xc5,
	0x25, 0x89, 0x25, 0xa5, 0xc5, 0x60, 0x05, 0xac, 0x41, 0x50, 0x9e, 0x90, 0x08, 0x17, 0x6b, 0x6a,
	0x51, 0x51, 0x7e, 0x91, 0x04, 0x93, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x84, 0x23, 0x24, 0xc4, 0xc5,
	0x92, 0x92, 0x58, 0x92, 0x28, 0xc1, 0xac, 0xc0, 0xa8, 0xc1, 0x13, 0x04, 0x66, 0x2b, 0x19, 0x71,
	0x71, 0x05, 0x94, 0xe2, 0xb6, 0x0d, 0xae, 0x87, 0x09, 0x49, 0x8f, 0x05, 0x17, 0x07, 0x58, 0x0f,
	0xc9, 0x2e, 0x00, 0xf9, 0xcd, 0x25, 0x35, 0x07, 0xb7, 0xdf, 0x2c, 0xb8, 0x38, 0xc0, 0xf2, 0x24,
	0x9b, 0x6c, 0x94, 0xc3, 0xc5, 0x1e, 0x5c, 0x92, 0x5f, 0x94, 0x98, 0x9e, 0x2a, 0x24, 0xcf, 0xc5,
	0xec, 0x9e, 0x5a, 0x22, 0xc4, 0xad, 0x87, 0x08, 0x46, 0x29, 0x4e, 0x3d, 0x58, 0x98, 0x29, 0x31,
	0x80, 0x14, 0x04, 0x94, 0x82, 0x14, 0x20, 0x7c, 0x2e, 0xc5, 0xa9, 0x07, 0xf3, 0x12, 0x44, 0x81,
	0x4b, 0x6a, 0x8e, 0x10, 0xb7, 0x1e, 0xc2, 0xb1, 0x52, 0x9c, 0x7a, 0x30, 0x97, 0x29, 0x31, 0x24,
	0xb1, 0x81, 0xa3, 0xca, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x38, 0xe5, 0xcc, 0x66, 0xb6, 0x01,
	0x00, 0x00,
}
