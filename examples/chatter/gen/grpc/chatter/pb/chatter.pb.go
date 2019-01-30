// Code generated by protoc-gen-go. DO NOT EDIT.
// source: chatter.proto

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

type LoginRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginRequest) Reset()         { *m = LoginRequest{} }
func (m *LoginRequest) String() string { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()    {}
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_chatter_51f9545dcc6a8a48, []int{0}
}
func (m *LoginRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginRequest.Unmarshal(m, b)
}
func (m *LoginRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginRequest.Marshal(b, m, deterministic)
}
func (dst *LoginRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginRequest.Merge(dst, src)
}
func (m *LoginRequest) XXX_Size() int {
	return xxx_messageInfo_LoginRequest.Size(m)
}
func (m *LoginRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LoginRequest proto.InternalMessageInfo

type LoginResponse struct {
	Field                string   `protobuf:"bytes,1,opt,name=field,proto3" json:"field,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginResponse) Reset()         { *m = LoginResponse{} }
func (m *LoginResponse) String() string { return proto.CompactTextString(m) }
func (*LoginResponse) ProtoMessage()    {}
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_chatter_51f9545dcc6a8a48, []int{1}
}
func (m *LoginResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginResponse.Unmarshal(m, b)
}
func (m *LoginResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginResponse.Marshal(b, m, deterministic)
}
func (dst *LoginResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginResponse.Merge(dst, src)
}
func (m *LoginResponse) XXX_Size() int {
	return xxx_messageInfo_LoginResponse.Size(m)
}
func (m *LoginResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LoginResponse proto.InternalMessageInfo

func (m *LoginResponse) GetField() string {
	if m != nil {
		return m.Field
	}
	return ""
}

type EchoerStreamingRequest struct {
	Field                string   `protobuf:"bytes,1,opt,name=field,proto3" json:"field,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EchoerStreamingRequest) Reset()         { *m = EchoerStreamingRequest{} }
func (m *EchoerStreamingRequest) String() string { return proto.CompactTextString(m) }
func (*EchoerStreamingRequest) ProtoMessage()    {}
func (*EchoerStreamingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_chatter_51f9545dcc6a8a48, []int{2}
}
func (m *EchoerStreamingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EchoerStreamingRequest.Unmarshal(m, b)
}
func (m *EchoerStreamingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EchoerStreamingRequest.Marshal(b, m, deterministic)
}
func (dst *EchoerStreamingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EchoerStreamingRequest.Merge(dst, src)
}
func (m *EchoerStreamingRequest) XXX_Size() int {
	return xxx_messageInfo_EchoerStreamingRequest.Size(m)
}
func (m *EchoerStreamingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EchoerStreamingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EchoerStreamingRequest proto.InternalMessageInfo

func (m *EchoerStreamingRequest) GetField() string {
	if m != nil {
		return m.Field
	}
	return ""
}

type EchoerResponse struct {
	Field                string   `protobuf:"bytes,1,opt,name=field,proto3" json:"field,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EchoerResponse) Reset()         { *m = EchoerResponse{} }
func (m *EchoerResponse) String() string { return proto.CompactTextString(m) }
func (*EchoerResponse) ProtoMessage()    {}
func (*EchoerResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_chatter_51f9545dcc6a8a48, []int{3}
}
func (m *EchoerResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EchoerResponse.Unmarshal(m, b)
}
func (m *EchoerResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EchoerResponse.Marshal(b, m, deterministic)
}
func (dst *EchoerResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EchoerResponse.Merge(dst, src)
}
func (m *EchoerResponse) XXX_Size() int {
	return xxx_messageInfo_EchoerResponse.Size(m)
}
func (m *EchoerResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EchoerResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EchoerResponse proto.InternalMessageInfo

func (m *EchoerResponse) GetField() string {
	if m != nil {
		return m.Field
	}
	return ""
}

type ListenerStreamingRequest struct {
	Field                string   `protobuf:"bytes,1,opt,name=field,proto3" json:"field,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListenerStreamingRequest) Reset()         { *m = ListenerStreamingRequest{} }
func (m *ListenerStreamingRequest) String() string { return proto.CompactTextString(m) }
func (*ListenerStreamingRequest) ProtoMessage()    {}
func (*ListenerStreamingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_chatter_51f9545dcc6a8a48, []int{4}
}
func (m *ListenerStreamingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListenerStreamingRequest.Unmarshal(m, b)
}
func (m *ListenerStreamingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListenerStreamingRequest.Marshal(b, m, deterministic)
}
func (dst *ListenerStreamingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListenerStreamingRequest.Merge(dst, src)
}
func (m *ListenerStreamingRequest) XXX_Size() int {
	return xxx_messageInfo_ListenerStreamingRequest.Size(m)
}
func (m *ListenerStreamingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListenerStreamingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListenerStreamingRequest proto.InternalMessageInfo

func (m *ListenerStreamingRequest) GetField() string {
	if m != nil {
		return m.Field
	}
	return ""
}

type ListenerResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListenerResponse) Reset()         { *m = ListenerResponse{} }
func (m *ListenerResponse) String() string { return proto.CompactTextString(m) }
func (*ListenerResponse) ProtoMessage()    {}
func (*ListenerResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_chatter_51f9545dcc6a8a48, []int{5}
}
func (m *ListenerResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListenerResponse.Unmarshal(m, b)
}
func (m *ListenerResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListenerResponse.Marshal(b, m, deterministic)
}
func (dst *ListenerResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListenerResponse.Merge(dst, src)
}
func (m *ListenerResponse) XXX_Size() int {
	return xxx_messageInfo_ListenerResponse.Size(m)
}
func (m *ListenerResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListenerResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListenerResponse proto.InternalMessageInfo

type SummaryStreamingRequest struct {
	Field                string   `protobuf:"bytes,1,opt,name=field,proto3" json:"field,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SummaryStreamingRequest) Reset()         { *m = SummaryStreamingRequest{} }
func (m *SummaryStreamingRequest) String() string { return proto.CompactTextString(m) }
func (*SummaryStreamingRequest) ProtoMessage()    {}
func (*SummaryStreamingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_chatter_51f9545dcc6a8a48, []int{6}
}
func (m *SummaryStreamingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SummaryStreamingRequest.Unmarshal(m, b)
}
func (m *SummaryStreamingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SummaryStreamingRequest.Marshal(b, m, deterministic)
}
func (dst *SummaryStreamingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SummaryStreamingRequest.Merge(dst, src)
}
func (m *SummaryStreamingRequest) XXX_Size() int {
	return xxx_messageInfo_SummaryStreamingRequest.Size(m)
}
func (m *SummaryStreamingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SummaryStreamingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SummaryStreamingRequest proto.InternalMessageInfo

func (m *SummaryStreamingRequest) GetField() string {
	if m != nil {
		return m.Field
	}
	return ""
}

type ChatSummaryCollection struct {
	Field                []*ChatSummary `protobuf:"bytes,1,rep,name=field,proto3" json:"field,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ChatSummaryCollection) Reset()         { *m = ChatSummaryCollection{} }
func (m *ChatSummaryCollection) String() string { return proto.CompactTextString(m) }
func (*ChatSummaryCollection) ProtoMessage()    {}
func (*ChatSummaryCollection) Descriptor() ([]byte, []int) {
	return fileDescriptor_chatter_51f9545dcc6a8a48, []int{7}
}
func (m *ChatSummaryCollection) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChatSummaryCollection.Unmarshal(m, b)
}
func (m *ChatSummaryCollection) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChatSummaryCollection.Marshal(b, m, deterministic)
}
func (dst *ChatSummaryCollection) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChatSummaryCollection.Merge(dst, src)
}
func (m *ChatSummaryCollection) XXX_Size() int {
	return xxx_messageInfo_ChatSummaryCollection.Size(m)
}
func (m *ChatSummaryCollection) XXX_DiscardUnknown() {
	xxx_messageInfo_ChatSummaryCollection.DiscardUnknown(m)
}

var xxx_messageInfo_ChatSummaryCollection proto.InternalMessageInfo

func (m *ChatSummaryCollection) GetField() []*ChatSummary {
	if m != nil {
		return m.Field
	}
	return nil
}

type ChatSummary struct {
	// Message sent to the server
	Message_ string `protobuf:"bytes,1,opt,name=message_,json=message,proto3" json:"message_,omitempty"`
	// Length of the message sent
	Length int32 `protobuf:"zigzag32,2,opt,name=length,proto3" json:"length,omitempty"`
	// Time at which the message was sent
	SentAt               string   `protobuf:"bytes,3,opt,name=sent_at,json=sentAt,proto3" json:"sent_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChatSummary) Reset()         { *m = ChatSummary{} }
func (m *ChatSummary) String() string { return proto.CompactTextString(m) }
func (*ChatSummary) ProtoMessage()    {}
func (*ChatSummary) Descriptor() ([]byte, []int) {
	return fileDescriptor_chatter_51f9545dcc6a8a48, []int{8}
}
func (m *ChatSummary) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChatSummary.Unmarshal(m, b)
}
func (m *ChatSummary) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChatSummary.Marshal(b, m, deterministic)
}
func (dst *ChatSummary) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChatSummary.Merge(dst, src)
}
func (m *ChatSummary) XXX_Size() int {
	return xxx_messageInfo_ChatSummary.Size(m)
}
func (m *ChatSummary) XXX_DiscardUnknown() {
	xxx_messageInfo_ChatSummary.DiscardUnknown(m)
}

var xxx_messageInfo_ChatSummary proto.InternalMessageInfo

func (m *ChatSummary) GetMessage_() string {
	if m != nil {
		return m.Message_
	}
	return ""
}

func (m *ChatSummary) GetLength() int32 {
	if m != nil {
		return m.Length
	}
	return 0
}

func (m *ChatSummary) GetSentAt() string {
	if m != nil {
		return m.SentAt
	}
	return ""
}

type HistoryRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HistoryRequest) Reset()         { *m = HistoryRequest{} }
func (m *HistoryRequest) String() string { return proto.CompactTextString(m) }
func (*HistoryRequest) ProtoMessage()    {}
func (*HistoryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_chatter_51f9545dcc6a8a48, []int{9}
}
func (m *HistoryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HistoryRequest.Unmarshal(m, b)
}
func (m *HistoryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HistoryRequest.Marshal(b, m, deterministic)
}
func (dst *HistoryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HistoryRequest.Merge(dst, src)
}
func (m *HistoryRequest) XXX_Size() int {
	return xxx_messageInfo_HistoryRequest.Size(m)
}
func (m *HistoryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HistoryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HistoryRequest proto.InternalMessageInfo

type HistoryResponse struct {
	// Message sent to the server
	Message_ string `protobuf:"bytes,1,opt,name=message_,json=message,proto3" json:"message_,omitempty"`
	// Length of the message sent
	Length int32 `protobuf:"zigzag32,2,opt,name=length,proto3" json:"length,omitempty"`
	// Time at which the message was sent
	SentAt               string   `protobuf:"bytes,3,opt,name=sent_at,json=sentAt,proto3" json:"sent_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HistoryResponse) Reset()         { *m = HistoryResponse{} }
func (m *HistoryResponse) String() string { return proto.CompactTextString(m) }
func (*HistoryResponse) ProtoMessage()    {}
func (*HistoryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_chatter_51f9545dcc6a8a48, []int{10}
}
func (m *HistoryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HistoryResponse.Unmarshal(m, b)
}
func (m *HistoryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HistoryResponse.Marshal(b, m, deterministic)
}
func (dst *HistoryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HistoryResponse.Merge(dst, src)
}
func (m *HistoryResponse) XXX_Size() int {
	return xxx_messageInfo_HistoryResponse.Size(m)
}
func (m *HistoryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HistoryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HistoryResponse proto.InternalMessageInfo

func (m *HistoryResponse) GetMessage_() string {
	if m != nil {
		return m.Message_
	}
	return ""
}

func (m *HistoryResponse) GetLength() int32 {
	if m != nil {
		return m.Length
	}
	return 0
}

func (m *HistoryResponse) GetSentAt() string {
	if m != nil {
		return m.SentAt
	}
	return ""
}

func init() {
	proto.RegisterType((*LoginRequest)(nil), "pb.LoginRequest")
	proto.RegisterType((*LoginResponse)(nil), "pb.LoginResponse")
	proto.RegisterType((*EchoerStreamingRequest)(nil), "pb.EchoerStreamingRequest")
	proto.RegisterType((*EchoerResponse)(nil), "pb.EchoerResponse")
	proto.RegisterType((*ListenerStreamingRequest)(nil), "pb.ListenerStreamingRequest")
	proto.RegisterType((*ListenerResponse)(nil), "pb.ListenerResponse")
	proto.RegisterType((*SummaryStreamingRequest)(nil), "pb.SummaryStreamingRequest")
	proto.RegisterType((*ChatSummaryCollection)(nil), "pb.ChatSummaryCollection")
	proto.RegisterType((*ChatSummary)(nil), "pb.ChatSummary")
	proto.RegisterType((*HistoryRequest)(nil), "pb.HistoryRequest")
	proto.RegisterType((*HistoryResponse)(nil), "pb.HistoryResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ChatterClient is the client API for Chatter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ChatterClient interface {
	// Creates a valid JWT token for auth to chat.
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	// Echoes the message sent by the client.
	Echoer(ctx context.Context, opts ...grpc.CallOption) (Chatter_EchoerClient, error)
	// Listens to the messages sent by the client.
	Listener(ctx context.Context, opts ...grpc.CallOption) (Chatter_ListenerClient, error)
	// Summarizes the chat messages sent by the client.
	Summary(ctx context.Context, opts ...grpc.CallOption) (Chatter_SummaryClient, error)
	// Returns the chat messages sent to the server.
	History(ctx context.Context, in *HistoryRequest, opts ...grpc.CallOption) (Chatter_HistoryClient, error)
}

type chatterClient struct {
	cc *grpc.ClientConn
}

func NewChatterClient(cc *grpc.ClientConn) ChatterClient {
	return &chatterClient{cc}
}

func (c *chatterClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/pb.Chatter/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatterClient) Echoer(ctx context.Context, opts ...grpc.CallOption) (Chatter_EchoerClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Chatter_serviceDesc.Streams[0], "/pb.Chatter/Echoer", opts...)
	if err != nil {
		return nil, err
	}
	x := &chatterEchoerClient{stream}
	return x, nil
}

type Chatter_EchoerClient interface {
	Send(*EchoerStreamingRequest) error
	Recv() (*EchoerResponse, error)
	grpc.ClientStream
}

type chatterEchoerClient struct {
	grpc.ClientStream
}

func (x *chatterEchoerClient) Send(m *EchoerStreamingRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *chatterEchoerClient) Recv() (*EchoerResponse, error) {
	m := new(EchoerResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *chatterClient) Listener(ctx context.Context, opts ...grpc.CallOption) (Chatter_ListenerClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Chatter_serviceDesc.Streams[1], "/pb.Chatter/Listener", opts...)
	if err != nil {
		return nil, err
	}
	x := &chatterListenerClient{stream}
	return x, nil
}

type Chatter_ListenerClient interface {
	Send(*ListenerStreamingRequest) error
	CloseAndRecv() (*ListenerResponse, error)
	grpc.ClientStream
}

type chatterListenerClient struct {
	grpc.ClientStream
}

func (x *chatterListenerClient) Send(m *ListenerStreamingRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *chatterListenerClient) CloseAndRecv() (*ListenerResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(ListenerResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *chatterClient) Summary(ctx context.Context, opts ...grpc.CallOption) (Chatter_SummaryClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Chatter_serviceDesc.Streams[2], "/pb.Chatter/Summary", opts...)
	if err != nil {
		return nil, err
	}
	x := &chatterSummaryClient{stream}
	return x, nil
}

type Chatter_SummaryClient interface {
	Send(*SummaryStreamingRequest) error
	CloseAndRecv() (*ChatSummaryCollection, error)
	grpc.ClientStream
}

type chatterSummaryClient struct {
	grpc.ClientStream
}

func (x *chatterSummaryClient) Send(m *SummaryStreamingRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *chatterSummaryClient) CloseAndRecv() (*ChatSummaryCollection, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(ChatSummaryCollection)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *chatterClient) History(ctx context.Context, in *HistoryRequest, opts ...grpc.CallOption) (Chatter_HistoryClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Chatter_serviceDesc.Streams[3], "/pb.Chatter/History", opts...)
	if err != nil {
		return nil, err
	}
	x := &chatterHistoryClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Chatter_HistoryClient interface {
	Recv() (*HistoryResponse, error)
	grpc.ClientStream
}

type chatterHistoryClient struct {
	grpc.ClientStream
}

func (x *chatterHistoryClient) Recv() (*HistoryResponse, error) {
	m := new(HistoryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ChatterServer is the server API for Chatter service.
type ChatterServer interface {
	// Creates a valid JWT token for auth to chat.
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	// Echoes the message sent by the client.
	Echoer(Chatter_EchoerServer) error
	// Listens to the messages sent by the client.
	Listener(Chatter_ListenerServer) error
	// Summarizes the chat messages sent by the client.
	Summary(Chatter_SummaryServer) error
	// Returns the chat messages sent to the server.
	History(*HistoryRequest, Chatter_HistoryServer) error
}

func RegisterChatterServer(s *grpc.Server, srv ChatterServer) {
	s.RegisterService(&_Chatter_serviceDesc, srv)
}

func _Chatter_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatterServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Chatter/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatterServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chatter_Echoer_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ChatterServer).Echoer(&chatterEchoerServer{stream})
}

type Chatter_EchoerServer interface {
	Send(*EchoerResponse) error
	Recv() (*EchoerStreamingRequest, error)
	grpc.ServerStream
}

type chatterEchoerServer struct {
	grpc.ServerStream
}

func (x *chatterEchoerServer) Send(m *EchoerResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *chatterEchoerServer) Recv() (*EchoerStreamingRequest, error) {
	m := new(EchoerStreamingRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Chatter_Listener_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ChatterServer).Listener(&chatterListenerServer{stream})
}

type Chatter_ListenerServer interface {
	SendAndClose(*ListenerResponse) error
	Recv() (*ListenerStreamingRequest, error)
	grpc.ServerStream
}

type chatterListenerServer struct {
	grpc.ServerStream
}

func (x *chatterListenerServer) SendAndClose(m *ListenerResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *chatterListenerServer) Recv() (*ListenerStreamingRequest, error) {
	m := new(ListenerStreamingRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Chatter_Summary_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ChatterServer).Summary(&chatterSummaryServer{stream})
}

type Chatter_SummaryServer interface {
	SendAndClose(*ChatSummaryCollection) error
	Recv() (*SummaryStreamingRequest, error)
	grpc.ServerStream
}

type chatterSummaryServer struct {
	grpc.ServerStream
}

func (x *chatterSummaryServer) SendAndClose(m *ChatSummaryCollection) error {
	return x.ServerStream.SendMsg(m)
}

func (x *chatterSummaryServer) Recv() (*SummaryStreamingRequest, error) {
	m := new(SummaryStreamingRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Chatter_History_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(HistoryRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ChatterServer).History(m, &chatterHistoryServer{stream})
}

type Chatter_HistoryServer interface {
	Send(*HistoryResponse) error
	grpc.ServerStream
}

type chatterHistoryServer struct {
	grpc.ServerStream
}

func (x *chatterHistoryServer) Send(m *HistoryResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _Chatter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Chatter",
	HandlerType: (*ChatterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _Chatter_Login_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Echoer",
			Handler:       _Chatter_Echoer_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "Listener",
			Handler:       _Chatter_Listener_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "Summary",
			Handler:       _Chatter_Summary_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "History",
			Handler:       _Chatter_History_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "chatter.proto",
}

func init() { proto.RegisterFile("chatter.proto", fileDescriptor_chatter_51f9545dcc6a8a48) }

var fileDescriptor_chatter_51f9545dcc6a8a48 = []byte{
	// 372 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x53, 0x4d, 0x6f, 0xaa, 0x40,
	0x14, 0x0d, 0x1a, 0xc1, 0x77, 0x7d, 0x7e, 0xdd, 0xe7, 0x53, 0xe4, 0xbd, 0x85, 0x21, 0xb1, 0x61,
	0xd1, 0x50, 0x63, 0xbb, 0x6c, 0x9a, 0x36, 0xa4, 0x49, 0x17, 0xae, 0x74, 0xd5, 0x45, 0x63, 0xd0,
	0x4e, 0x81, 0x04, 0x18, 0xca, 0x8c, 0x0b, 0xff, 0x56, 0x7f, 0x61, 0x03, 0xcc, 0x14, 0x6c, 0x6a,
	0xea, 0xa2, 0xcb, 0x7b, 0xef, 0x39, 0xe7, 0x0e, 0xf7, 0x1c, 0xa0, 0xbd, 0xf5, 0x5d, 0xce, 0x49,
	0x6a, 0x27, 0x29, 0xe5, 0x14, 0x6b, 0xc9, 0xc6, 0xec, 0xc0, 0xef, 0x05, 0xf5, 0x82, 0x78, 0x49,
	0x5e, 0x77, 0x84, 0x71, 0x73, 0x0a, 0x6d, 0x51, 0xb3, 0x84, 0xc6, 0x8c, 0xe0, 0x00, 0x1a, 0x2f,
	0x01, 0x09, 0x9f, 0x75, 0x65, 0xa2, 0x58, 0xbf, 0x96, 0x45, 0x61, 0xda, 0x30, 0xbc, 0xdf, 0xfa,
	0x94, 0xa4, 0x2b, 0x9e, 0x12, 0x37, 0x0a, 0x62, 0x4f, 0x08, 0x1c, 0xc1, 0x9f, 0x41, 0xa7, 0xc0,
	0x7f, 0xa3, 0x3b, 0x03, 0x7d, 0x11, 0x30, 0x4e, 0xe2, 0x93, 0x95, 0x11, 0x7a, 0x92, 0x21, 0xb5,
	0xcd, 0x0b, 0x18, 0xad, 0x76, 0x51, 0xe4, 0xa6, 0xfb, 0x13, 0x45, 0x6e, 0xe0, 0xaf, 0xe3, 0xbb,
	0x5c, 0x90, 0x1c, 0x1a, 0x86, 0x64, 0xcb, 0x03, 0x1a, 0xe3, 0xb4, 0x84, 0xd7, 0xad, 0xd6, 0xbc,
	0x6b, 0x27, 0x1b, 0xbb, 0x82, 0x94, 0xfc, 0x47, 0x68, 0x55, 0xba, 0x38, 0x86, 0x66, 0x44, 0x18,
	0x73, 0x3d, 0xb2, 0x16, 0x7b, 0x34, 0x51, 0xe3, 0x10, 0xd4, 0x90, 0xc4, 0x1e, 0xf7, 0xf5, 0xda,
	0x44, 0xb1, 0xfa, 0x4b, 0x51, 0xe1, 0x08, 0x34, 0x46, 0x62, 0xbe, 0x76, 0xb9, 0x5e, 0xcf, 0x19,
	0x6a, 0x56, 0xde, 0x71, 0xb3, 0x07, 0x9d, 0x87, 0x80, 0x71, 0x9a, 0xee, 0xa5, 0x45, 0x4f, 0xd0,
	0xfd, 0xe8, 0x88, 0x63, 0xfe, 0xe0, 0xc2, 0xf9, 0x5b, 0x0d, 0x34, 0xa7, 0xc8, 0x09, 0x9e, 0x43,
	0x23, 0x4f, 0x03, 0xf6, 0xb2, 0x0f, 0xaf, 0x06, 0xc5, 0xe8, 0x57, 0x3a, 0xe2, 0x15, 0xd7, 0xa0,
	0x16, 0x26, 0xa3, 0x91, 0x0d, 0xbf, 0x0e, 0x88, 0x81, 0xe5, 0x4c, 0x32, 0x2d, 0x65, 0xa6, 0xe0,
	0x2d, 0x34, 0xa5, 0x91, 0xf8, 0x3f, 0x17, 0x3f, 0x12, 0x04, 0x63, 0x50, 0x9d, 0x96, 0x1a, 0xe8,
	0x80, 0x26, 0x1d, 0xf8, 0x97, 0x41, 0x8e, 0x64, 0xc0, 0x18, 0x7f, 0x72, 0xb1, 0xf4, 0xdb, 0x52,
	0xf0, 0x0a, 0x34, 0x71, 0x5d, 0xcc, 0x5f, 0x7a, 0x78, 0x7c, 0xe3, 0xcf, 0x41, 0xaf, 0x58, 0x3d,
	0x53, 0x36, 0x6a, 0xfe, 0x47, 0x5d, 0xbe, 0x07, 0x00, 0x00, 0xff, 0xff, 0xb2, 0xd3, 0xad, 0x54,
	0x62, 0x03, 0x00, 0x00,
}