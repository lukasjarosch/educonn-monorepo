// Code generated by protoc-gen-go. DO NOT EDIT.
// source: account/account.proto

package educonn_account

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
import _ "google.golang.org/genproto/googleapis/api/annotations"

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

type AccountInfo struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Email                string   `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Password             string   `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccountInfo) Reset()         { *m = AccountInfo{} }
func (m *AccountInfo) String() string { return proto.CompactTextString(m) }
func (*AccountInfo) ProtoMessage()    {}
func (*AccountInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_account_d559eb6722d9f087, []int{0}
}
func (m *AccountInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountInfo.Unmarshal(m, b)
}
func (m *AccountInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountInfo.Marshal(b, m, deterministic)
}
func (dst *AccountInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountInfo.Merge(dst, src)
}
func (m *AccountInfo) XXX_Size() int {
	return xxx_messageInfo_AccountInfo.Size(m)
}
func (m *AccountInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountInfo.DiscardUnknown(m)
}

var xxx_messageInfo_AccountInfo proto.InternalMessageInfo

func (m *AccountInfo) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *AccountInfo) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *AccountInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AccountInfo) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type CreateAccountRequest struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Password             string   `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateAccountRequest) Reset()         { *m = CreateAccountRequest{} }
func (m *CreateAccountRequest) String() string { return proto.CompactTextString(m) }
func (*CreateAccountRequest) ProtoMessage()    {}
func (*CreateAccountRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_account_d559eb6722d9f087, []int{1}
}
func (m *CreateAccountRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateAccountRequest.Unmarshal(m, b)
}
func (m *CreateAccountRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateAccountRequest.Marshal(b, m, deterministic)
}
func (dst *CreateAccountRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateAccountRequest.Merge(dst, src)
}
func (m *CreateAccountRequest) XXX_Size() int {
	return xxx_messageInfo_CreateAccountRequest.Size(m)
}
func (m *CreateAccountRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateAccountRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateAccountRequest proto.InternalMessageInfo

func (m *CreateAccountRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *CreateAccountRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateAccountRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type CreateAccountResponse struct {
	Account              *AccountInfo `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *CreateAccountResponse) Reset()         { *m = CreateAccountResponse{} }
func (m *CreateAccountResponse) String() string { return proto.CompactTextString(m) }
func (*CreateAccountResponse) ProtoMessage()    {}
func (*CreateAccountResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_account_d559eb6722d9f087, []int{2}
}
func (m *CreateAccountResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateAccountResponse.Unmarshal(m, b)
}
func (m *CreateAccountResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateAccountResponse.Marshal(b, m, deterministic)
}
func (dst *CreateAccountResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateAccountResponse.Merge(dst, src)
}
func (m *CreateAccountResponse) XXX_Size() int {
	return xxx_messageInfo_CreateAccountResponse.Size(m)
}
func (m *CreateAccountResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateAccountResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateAccountResponse proto.InternalMessageInfo

func (m *CreateAccountResponse) GetAccount() *AccountInfo {
	if m != nil {
		return m.Account
	}
	return nil
}

func init() {
	proto.RegisterType((*AccountInfo)(nil), "educonn.account.AccountInfo")
	proto.RegisterType((*CreateAccountRequest)(nil), "educonn.account.CreateAccountRequest")
	proto.RegisterType((*CreateAccountResponse)(nil), "educonn.account.CreateAccountResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AccountClient is the client API for Account service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AccountClient interface {
	// Create is used to create a new educonn account
	Create(ctx context.Context, in *CreateAccountRequest, opts ...grpc.CallOption) (*CreateAccountResponse, error)
}

type accountClient struct {
	cc *grpc.ClientConn
}

func NewAccountClient(cc *grpc.ClientConn) AccountClient {
	return &accountClient{cc}
}

func (c *accountClient) Create(ctx context.Context, in *CreateAccountRequest, opts ...grpc.CallOption) (*CreateAccountResponse, error) {
	out := new(CreateAccountResponse)
	err := c.cc.Invoke(ctx, "/educonn.account.Account/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccountServer is the server API for Account service.
type AccountServer interface {
	// Create is used to create a new educonn account
	Create(context.Context, *CreateAccountRequest) (*CreateAccountResponse, error)
}

func RegisterAccountServer(s *grpc.Server, srv AccountServer) {
	s.RegisterService(&_Account_serviceDesc, srv)
}

func _Account_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/educonn.account.Account/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).Create(ctx, req.(*CreateAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Account_serviceDesc = grpc.ServiceDesc{
	ServiceName: "educonn.account.Account",
	HandlerType: (*AccountServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Account_Create_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "account/account.proto",
}

func init() { proto.RegisterFile("account/account.proto", fileDescriptor_account_d559eb6722d9f087) }

var fileDescriptor_account_d559eb6722d9f087 = []byte{
	// 286 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0xcf, 0x4a, 0xc3, 0x40,
	0x10, 0xc6, 0x49, 0x5a, 0x5b, 0x9d, 0x80, 0xca, 0xd2, 0x4a, 0x08, 0x3d, 0x48, 0x40, 0x11, 0xb1,
	0x09, 0x56, 0xf0, 0xe0, 0x4d, 0x3c, 0x79, 0x12, 0x72, 0xf6, 0xb2, 0x26, 0x63, 0x08, 0xb4, 0x33,
	0x31, 0xbb, 0xb1, 0xe0, 0xd1, 0x57, 0xf0, 0xd1, 0x7c, 0x05, 0x1f, 0x44, 0xd8, 0xdd, 0xaa, 0xa4,
	0x95, 0x9e, 0x66, 0xe7, 0xe3, 0xdb, 0xdf, 0xfc, 0x83, 0xb1, 0xcc, 0x73, 0x6e, 0x49, 0xa7, 0x2e,
	0x26, 0x75, 0xc3, 0x9a, 0xc5, 0x01, 0x16, 0x6d, 0xce, 0x44, 0x89, 0x93, 0xa3, 0x49, 0xc9, 0x5c,
	0xce, 0x31, 0x95, 0x75, 0x95, 0x4a, 0x22, 0xd6, 0x52, 0x57, 0x4c, 0xca, 0xda, 0xa3, 0x0b, 0x13,
	0xf2, 0x69, 0x89, 0x34, 0x55, 0x4b, 0x59, 0x96, 0xd8, 0xa4, 0x5c, 0x1b, 0xc7, 0xba, 0x3b, 0xce,
	0x21, 0xb8, 0xb5, 0xd8, 0x7b, 0x7a, 0x66, 0xb1, 0x0f, 0x7e, 0x55, 0x84, 0xde, 0xb1, 0x77, 0xb6,
	0x97, 0xf9, 0x55, 0x21, 0x46, 0xb0, 0x83, 0x0b, 0x59, 0xcd, 0x43, 0xdf, 0x48, 0x36, 0x11, 0x02,
	0xfa, 0x24, 0x17, 0x18, 0xf6, 0x8c, 0x68, 0xde, 0x22, 0x82, 0xdd, 0x5a, 0x2a, 0xb5, 0xe4, 0xa6,
	0x08, 0xfb, 0x46, 0xff, 0xc9, 0xe3, 0x47, 0x18, 0xdd, 0x35, 0x28, 0x35, 0xba, 0x52, 0x19, 0xbe,
	0xb4, 0xa8, 0xf4, 0x2f, 0xdd, 0xdb, 0x44, 0xf7, 0xff, 0xa1, 0xf7, 0x3a, 0xf4, 0x07, 0x18, 0x77,
	0xe8, 0xaa, 0x66, 0x52, 0x28, 0xae, 0x61, 0xe8, 0x56, 0x66, 0x0a, 0x04, 0xb3, 0x49, 0xd2, 0x59,
	0x65, 0xf2, 0x67, 0xf6, 0x6c, 0x65, 0x9e, 0xbd, 0xc1, 0xd0, 0xe9, 0x82, 0x61, 0x60, 0xd9, 0xe2,
	0x64, 0xed, 0xef, 0xa6, 0x91, 0xa2, 0xd3, 0x6d, 0x36, 0xdb, 0x5b, 0x7c, 0xf4, 0xfe, 0xf9, 0xf5,
	0xe1, 0x1f, 0xc6, 0x41, 0xfa, 0x7a, 0xb9, 0xba, 0xf7, 0x8d, 0x77, 0xfe, 0x34, 0x30, 0x67, 0xb9,
	0xfa, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x26, 0xcf, 0xf4, 0x45, 0x0c, 0x02, 0x00, 0x00,
}
