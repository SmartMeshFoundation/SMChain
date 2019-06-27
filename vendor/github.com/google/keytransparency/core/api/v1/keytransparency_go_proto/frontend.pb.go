// Code generated by protoc-gen-go. DO NOT EDIT.
// source: v1/frontend.proto

// Key Transparency Frontend
//
// The Key Transparency API consists of a map of user names to public
// keys. Each user name also has a history of public keys that have been
// associated with it.

package keytransparency_go_proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
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

// QueueKeyUpdateRequest enqueues an update to a user's identity keys.
type QueueKeyUpdateRequest struct {
	// directory_id identifies the directory in which the user lives.
	DirectoryId string `protobuf:"bytes,1,opt,name=directory_id,json=directoryId,proto3" json:"directory_id,omitempty"`
	// user_id specifies the id for the user whose keys are being updated.
	UserId string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	// key_data is the key data to store.
	KeyData              []byte   `protobuf:"bytes,3,opt,name=key_data,json=keyData,proto3" json:"key_data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueueKeyUpdateRequest) Reset()         { *m = QueueKeyUpdateRequest{} }
func (m *QueueKeyUpdateRequest) String() string { return proto.CompactTextString(m) }
func (*QueueKeyUpdateRequest) ProtoMessage()    {}
func (*QueueKeyUpdateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bbdf9efa2240600c, []int{0}
}

func (m *QueueKeyUpdateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueueKeyUpdateRequest.Unmarshal(m, b)
}
func (m *QueueKeyUpdateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueueKeyUpdateRequest.Marshal(b, m, deterministic)
}
func (m *QueueKeyUpdateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueueKeyUpdateRequest.Merge(m, src)
}
func (m *QueueKeyUpdateRequest) XXX_Size() int {
	return xxx_messageInfo_QueueKeyUpdateRequest.Size(m)
}
func (m *QueueKeyUpdateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueueKeyUpdateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueueKeyUpdateRequest proto.InternalMessageInfo

func (m *QueueKeyUpdateRequest) GetDirectoryId() string {
	if m != nil {
		return m.DirectoryId
	}
	return ""
}

func (m *QueueKeyUpdateRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *QueueKeyUpdateRequest) GetKeyData() []byte {
	if m != nil {
		return m.KeyData
	}
	return nil
}

func init() {
	proto.RegisterType((*QueueKeyUpdateRequest)(nil), "google.keytransparency.v1.QueueKeyUpdateRequest")
}

func init() { proto.RegisterFile("v1/frontend.proto", fileDescriptor_bbdf9efa2240600c) }

var fileDescriptor_bbdf9efa2240600c = []byte{
	// 317 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xcf, 0x4a, 0xc3, 0x30,
	0x18, 0xc0, 0xe9, 0x84, 0x4d, 0xe3, 0x10, 0x2c, 0xe8, 0xfe, 0xe8, 0x61, 0xee, 0x34, 0x3c, 0x24,
	0xd6, 0x9d, 0xf4, 0x28, 0x3a, 0x1c, 0x3b, 0x39, 0xf4, 0xe2, 0xa5, 0x64, 0xcd, 0xb7, 0x5a, 0xe6,
	0x92, 0x9a, 0x7e, 0x29, 0x84, 0xb1, 0x8b, 0xaf, 0x20, 0x3e, 0x88, 0xcf, 0xe2, 0x2b, 0xf8, 0x20,
	0xd2, 0xb5, 0x95, 0x39, 0xf4, 0x14, 0xbe, 0xfc, 0xf8, 0xe0, 0xf7, 0x4b, 0xc8, 0x7e, 0xea, 0xb1,
	0xa9, 0x56, 0x12, 0x41, 0x0a, 0x1a, 0x6b, 0x85, 0xca, 0x6d, 0x85, 0x4a, 0x85, 0xcf, 0x40, 0x67,
	0x60, 0x51, 0x73, 0x99, 0xc4, 0x5c, 0x83, 0x0c, 0x2c, 0x4d, 0xbd, 0xf6, 0x71, 0x8e, 0x18, 0x8f,
	0x23, 0xc6, 0xa5, 0x54, 0xc8, 0x31, 0x52, 0x32, 0xc9, 0x17, 0xdb, 0x47, 0x05, 0x5d, 0x4d, 0x13,
	0x33, 0x65, 0x30, 0x8f, 0xd1, 0xe6, 0xb0, 0x2b, 0xc9, 0xc1, 0x9d, 0x01, 0x03, 0x23, 0xb0, 0x0f,
	0xb1, 0xe0, 0x08, 0x63, 0x78, 0x31, 0x90, 0xa0, 0x7b, 0x42, 0xea, 0x22, 0xd2, 0x10, 0xa0, 0xd2,
	0xd6, 0x8f, 0x44, 0xd3, 0xe9, 0x38, 0xbd, 0x9d, 0xf1, 0xee, 0xcf, 0xdd, 0x50, 0xb8, 0x0d, 0x52,
	0x33, 0x09, 0xe8, 0x8c, 0x56, 0x56, 0xb4, 0x9a, 0x8d, 0x43, 0xe1, 0xb6, 0xc8, 0xf6, 0x0c, 0xac,
	0x2f, 0x38, 0xf2, 0xe6, 0x56, 0xc7, 0xe9, 0xd5, 0xc7, 0xb5, 0x19, 0xd8, 0x6b, 0x8e, 0xfc, 0xfc,
	0xc3, 0x21, 0x8d, 0x11, 0xd8, 0xfb, 0xb5, 0x82, 0x41, 0xd1, 0xe9, 0xbe, 0x3b, 0x64, 0xef, 0xb7,
	0x8c, 0x7b, 0x46, 0xff, 0xad, 0xa6, 0x7f, 0x7a, 0xb7, 0x0f, 0xcb, 0x8d, 0x32, 0x97, 0xde, 0x64,
	0xb9, 0xdd, 0x8b, 0xd7, 0xcf, 0xaf, 0xb7, 0x4a, 0xbf, 0x4b, 0x59, 0xea, 0xb1, 0xb2, 0x22, 0x82,
	0x84, 0x2d, 0xd6, 0x33, 0x97, 0x2c, 0x4b, 0x48, 0xd8, 0xa2, 0x08, 0x5b, 0x5e, 0x3a, 0xa7, 0x57,
	0xb7, 0x8f, 0x83, 0x30, 0xc2, 0x27, 0x33, 0xa1, 0x81, 0x9a, 0xb3, 0xe2, 0x35, 0x37, 0x84, 0x58,
	0xa0, 0x74, 0xfe, 0x01, 0xa9, 0xb7, 0xc9, 0xfc, 0x50, 0xf9, 0xb9, 0x4d, 0x75, 0x75, 0xf4, 0xbf,
	0x03, 0x00, 0x00, 0xff, 0xff, 0xb8, 0xd0, 0x75, 0x47, 0xdf, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// KeyTransparencyFrontendClient is the client API for KeyTransparencyFrontend service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type KeyTransparencyFrontendClient interface {
	// Enqueues an update to a user's identity keys.
	QueueKeyUpdate(ctx context.Context, in *QueueKeyUpdateRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type keyTransparencyFrontendClient struct {
	cc *grpc.ClientConn
}

func NewKeyTransparencyFrontendClient(cc *grpc.ClientConn) KeyTransparencyFrontendClient {
	return &keyTransparencyFrontendClient{cc}
}

func (c *keyTransparencyFrontendClient) QueueKeyUpdate(ctx context.Context, in *QueueKeyUpdateRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/google.keytransparency.v1.KeyTransparencyFrontend/QueueKeyUpdate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KeyTransparencyFrontendServer is the server API for KeyTransparencyFrontend service.
type KeyTransparencyFrontendServer interface {
	// Enqueues an update to a user's identity keys.
	QueueKeyUpdate(context.Context, *QueueKeyUpdateRequest) (*empty.Empty, error)
}

// UnimplementedKeyTransparencyFrontendServer can be embedded to have forward compatible implementations.
type UnimplementedKeyTransparencyFrontendServer struct {
}

func (*UnimplementedKeyTransparencyFrontendServer) QueueKeyUpdate(ctx context.Context, req *QueueKeyUpdateRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueueKeyUpdate not implemented")
}

func RegisterKeyTransparencyFrontendServer(s *grpc.Server, srv KeyTransparencyFrontendServer) {
	s.RegisterService(&_KeyTransparencyFrontend_serviceDesc, srv)
}

func _KeyTransparencyFrontend_QueueKeyUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueueKeyUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyTransparencyFrontendServer).QueueKeyUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.keytransparency.v1.KeyTransparencyFrontend/QueueKeyUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyTransparencyFrontendServer).QueueKeyUpdate(ctx, req.(*QueueKeyUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _KeyTransparencyFrontend_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.keytransparency.v1.KeyTransparencyFrontend",
	HandlerType: (*KeyTransparencyFrontendServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "QueueKeyUpdate",
			Handler:    _KeyTransparencyFrontend_QueueKeyUpdate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/frontend.proto",
}