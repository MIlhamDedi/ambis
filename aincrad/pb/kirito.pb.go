// Code generated by protoc-gen-go. DO NOT EDIT.
// source: kirito.proto

package pb

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

type Status struct {
	Success              bool     `protobuf:"varint,1,opt,name=Success,proto3" json:"Success,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=Message,proto3" json:"Message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Status) Reset()         { *m = Status{} }
func (m *Status) String() string { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()    {}
func (*Status) Descriptor() ([]byte, []int) {
	return fileDescriptor_8dcb8eeaebe4b958, []int{0}
}

func (m *Status) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Status.Unmarshal(m, b)
}
func (m *Status) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Status.Marshal(b, m, deterministic)
}
func (m *Status) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Status.Merge(m, src)
}
func (m *Status) XXX_Size() int {
	return xxx_messageInfo_Status.Size(m)
}
func (m *Status) XXX_DiscardUnknown() {
	xxx_messageInfo_Status.DiscardUnknown(m)
}

var xxx_messageInfo_Status proto.InternalMessageInfo

func (m *Status) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *Status) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type User struct {
	Username             string   `protobuf:"bytes,1,opt,name=Username,proto3" json:"Username,omitempty"`
	FirstName            string   `protobuf:"bytes,2,opt,name=FirstName,proto3" json:"FirstName,omitempty"`
	LastName             string   `protobuf:"bytes,3,opt,name=LastName,proto3" json:"LastName,omitempty"`
	Email                string   `protobuf:"bytes,4,opt,name=Email,proto3" json:"Email,omitempty"`
	Password             string   `protobuf:"bytes,5,opt,name=Password,proto3" json:"Password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_8dcb8eeaebe4b958, []int{1}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *User) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *User) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type AuthRequest struct {
	User                 *User    `protobuf:"bytes,1,opt,name=User,proto3" json:"User,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthRequest) Reset()         { *m = AuthRequest{} }
func (m *AuthRequest) String() string { return proto.CompactTextString(m) }
func (*AuthRequest) ProtoMessage()    {}
func (*AuthRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8dcb8eeaebe4b958, []int{2}
}

func (m *AuthRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthRequest.Unmarshal(m, b)
}
func (m *AuthRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthRequest.Marshal(b, m, deterministic)
}
func (m *AuthRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthRequest.Merge(m, src)
}
func (m *AuthRequest) XXX_Size() int {
	return xxx_messageInfo_AuthRequest.Size(m)
}
func (m *AuthRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AuthRequest proto.InternalMessageInfo

func (m *AuthRequest) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type AuthResponse struct {
	Status               *Status  `protobuf:"bytes,1,opt,name=Status,proto3" json:"Status,omitempty"`
	User                 *User    `protobuf:"bytes,2,opt,name=User,proto3" json:"User,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthResponse) Reset()         { *m = AuthResponse{} }
func (m *AuthResponse) String() string { return proto.CompactTextString(m) }
func (*AuthResponse) ProtoMessage()    {}
func (*AuthResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8dcb8eeaebe4b958, []int{3}
}

func (m *AuthResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthResponse.Unmarshal(m, b)
}
func (m *AuthResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthResponse.Marshal(b, m, deterministic)
}
func (m *AuthResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthResponse.Merge(m, src)
}
func (m *AuthResponse) XXX_Size() int {
	return xxx_messageInfo_AuthResponse.Size(m)
}
func (m *AuthResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AuthResponse proto.InternalMessageInfo

func (m *AuthResponse) GetStatus() *Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *AuthResponse) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type Response struct {
	Status               *Status  `protobuf:"bytes,1,opt,name=Status,proto3" json:"Status,omitempty"`
	User                 *User    `protobuf:"bytes,2,opt,name=User,proto3" json:"User,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_8dcb8eeaebe4b958, []int{4}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetStatus() *Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *Response) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func init() {
	proto.RegisterType((*Status)(nil), "pb.Status")
	proto.RegisterType((*User)(nil), "pb.User")
	proto.RegisterType((*AuthRequest)(nil), "pb.AuthRequest")
	proto.RegisterType((*AuthResponse)(nil), "pb.AuthResponse")
	proto.RegisterType((*Response)(nil), "pb.Response")
}

func init() { proto.RegisterFile("kirito.proto", fileDescriptor_8dcb8eeaebe4b958) }

var fileDescriptor_8dcb8eeaebe4b958 = []byte{
	// 302 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x52, 0x4f, 0x4b, 0xf3, 0x30,
	0x18, 0xa7, 0x7b, 0xb7, 0xbe, 0xdd, 0xd3, 0x82, 0x12, 0x3c, 0x94, 0xb2, 0x43, 0xe9, 0x69, 0x20,
	0x14, 0xac, 0x57, 0x0f, 0x8a, 0xe8, 0x69, 0xca, 0x4c, 0xd1, 0x7b, 0x5a, 0x1f, 0x35, 0xe8, 0xda,
	0x9a, 0x27, 0x55, 0xfc, 0x12, 0x7e, 0x66, 0x49, 0xb3, 0x74, 0x1e, 0xf4, 0xe6, 0x29, 0xfd, 0xfd,
	0x6d, 0xf8, 0x11, 0x88, 0x9e, 0xa5, 0x92, 0xba, 0xcd, 0x3b, 0xd5, 0xea, 0x96, 0x4d, 0xba, 0x2a,
	0x3b, 0x01, 0xbf, 0xd4, 0x42, 0xf7, 0xc4, 0x62, 0xf8, 0x5f, 0xf6, 0x75, 0x8d, 0x44, 0xb1, 0x97,
	0x7a, 0xcb, 0x80, 0x3b, 0x68, 0x94, 0x2b, 0x24, 0x12, 0x8f, 0x18, 0x4f, 0x52, 0x6f, 0x39, 0xe7,
	0x0e, 0x66, 0x9f, 0x1e, 0x4c, 0x6f, 0x09, 0x15, 0x4b, 0x20, 0x30, 0x67, 0x23, 0x36, 0x38, 0xa4,
	0xe7, 0x7c, 0xc4, 0x6c, 0x01, 0xf3, 0x4b, 0xa9, 0x48, 0x5f, 0x1b, 0xd1, 0x16, 0xec, 0x08, 0x93,
	0x5c, 0x89, 0xad, 0xf8, 0xcf, 0x26, 0x1d, 0x66, 0x07, 0x30, 0xbb, 0xd8, 0x08, 0xf9, 0x12, 0x4f,
	0x07, 0xc1, 0x02, 0x93, 0x58, 0x0b, 0xa2, 0xf7, 0x56, 0xdd, 0xc7, 0x33, 0x9b, 0x70, 0x38, 0x3b,
	0x84, 0xf0, 0xac, 0xd7, 0x4f, 0x1c, 0x5f, 0x7b, 0x24, 0xcd, 0x16, 0xf6, 0x7a, 0xc3, 0x95, 0xc2,
	0x22, 0xc8, 0xbb, 0x2a, 0x37, 0x98, 0x0f, 0x6c, 0xb6, 0x86, 0xc8, 0x9a, 0xa9, 0x6b, 0x1b, 0x42,
	0x96, 0xb9, 0x2d, 0xb6, 0x7e, 0x30, 0x7e, 0xcb, 0x70, 0xb7, 0x92, 0x6b, 0x9c, 0xfc, 0xd8, 0xb8,
	0x82, 0xe0, 0xef, 0xda, 0x8a, 0x1b, 0x08, 0xcd, 0x59, 0xa2, 0x7a, 0x93, 0x35, 0xb2, 0x14, 0xfc,
	0x73, 0x85, 0x42, 0x23, 0x1b, 0x8d, 0x49, 0x64, 0xbe, 0xc6, 0x5f, 0xa6, 0xe0, 0xdf, 0xa1, 0x92,
	0x0f, 0x1f, 0xbf, 0x39, 0x8a, 0x53, 0xbb, 0x8f, 0xab, 0x3c, 0xb2, 0x0b, 0x60, 0xa3, 0x65, 0x6d,
	0x8a, 0xf7, 0x8c, 0xf9, 0xdb, 0x80, 0xc9, 0xfe, 0x8e, 0xb0, 0x0d, 0x95, 0x3f, 0xbc, 0x9d, 0xe3,
	0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x8e, 0x57, 0xc1, 0x7a, 0x4b, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserServiceClient interface {
	Create(ctx context.Context, in *User, opts ...grpc.CallOption) (*Response, error)
	Verify(ctx context.Context, in *User, opts ...grpc.CallOption) (*Response, error)
}

type userServiceClient struct {
	cc *grpc.ClientConn
}

func NewUserServiceClient(cc *grpc.ClientConn) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) Create(ctx context.Context, in *User, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/pb.UserService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Verify(ctx context.Context, in *User, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/pb.UserService/Verify", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
type UserServiceServer interface {
	Create(context.Context, *User) (*Response, error)
	Verify(context.Context, *User) (*Response, error)
}

func RegisterUserServiceServer(s *grpc.Server, srv UserServiceServer) {
	s.RegisterService(&_UserService_serviceDesc, srv)
}

func _UserService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Create(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Verify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Verify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserService/Verify",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Verify(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _UserService_Create_Handler,
		},
		{
			MethodName: "Verify",
			Handler:    _UserService_Verify_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "kirito.proto",
}

// AuthServiceClient is the client API for AuthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthServiceClient interface {
	Authenticate(ctx context.Context, in *AuthRequest, opts ...grpc.CallOption) (*AuthResponse, error)
}

type authServiceClient struct {
	cc *grpc.ClientConn
}

func NewAuthServiceClient(cc *grpc.ClientConn) AuthServiceClient {
	return &authServiceClient{cc}
}

func (c *authServiceClient) Authenticate(ctx context.Context, in *AuthRequest, opts ...grpc.CallOption) (*AuthResponse, error) {
	out := new(AuthResponse)
	err := c.cc.Invoke(ctx, "/pb.AuthService/Authenticate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServiceServer is the server API for AuthService service.
type AuthServiceServer interface {
	Authenticate(context.Context, *AuthRequest) (*AuthResponse, error)
}

func RegisterAuthServiceServer(s *grpc.Server, srv AuthServiceServer) {
	s.RegisterService(&_AuthService_serviceDesc, srv)
}

func _AuthService_Authenticate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).Authenticate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.AuthService/Authenticate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).Authenticate(ctx, req.(*AuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AuthService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.AuthService",
	HandlerType: (*AuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Authenticate",
			Handler:    _AuthService_Authenticate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "kirito.proto",
}
