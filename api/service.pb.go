// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

package api

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

type Voteable struct {
	Uuid                 string   `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Question             string   `protobuf:"bytes,2,opt,name=question,proto3" json:"question,omitempty"`
	Answers              []string `protobuf:"bytes,3,rep,name=answers,proto3" json:"answers,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Voteable) Reset()         { *m = Voteable{} }
func (m *Voteable) String() string { return proto.CompactTextString(m) }
func (*Voteable) ProtoMessage()    {}
func (*Voteable) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{0}
}

func (m *Voteable) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Voteable.Unmarshal(m, b)
}
func (m *Voteable) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Voteable.Marshal(b, m, deterministic)
}
func (m *Voteable) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Voteable.Merge(m, src)
}
func (m *Voteable) XXX_Size() int {
	return xxx_messageInfo_Voteable.Size(m)
}
func (m *Voteable) XXX_DiscardUnknown() {
	xxx_messageInfo_Voteable.DiscardUnknown(m)
}

var xxx_messageInfo_Voteable proto.InternalMessageInfo

func (m *Voteable) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *Voteable) GetQuestion() string {
	if m != nil {
		return m.Question
	}
	return ""
}

func (m *Voteable) GetAnswers() []string {
	if m != nil {
		return m.Answers
	}
	return nil
}

type CreateVoteableRequest struct {
	Question             string   `protobuf:"bytes,1,opt,name=question,proto3" json:"question,omitempty"`
	Answers              []string `protobuf:"bytes,2,rep,name=answers,proto3" json:"answers,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateVoteableRequest) Reset()         { *m = CreateVoteableRequest{} }
func (m *CreateVoteableRequest) String() string { return proto.CompactTextString(m) }
func (*CreateVoteableRequest) ProtoMessage()    {}
func (*CreateVoteableRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{1}
}

func (m *CreateVoteableRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateVoteableRequest.Unmarshal(m, b)
}
func (m *CreateVoteableRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateVoteableRequest.Marshal(b, m, deterministic)
}
func (m *CreateVoteableRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateVoteableRequest.Merge(m, src)
}
func (m *CreateVoteableRequest) XXX_Size() int {
	return xxx_messageInfo_CreateVoteableRequest.Size(m)
}
func (m *CreateVoteableRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateVoteableRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateVoteableRequest proto.InternalMessageInfo

func (m *CreateVoteableRequest) GetQuestion() string {
	if m != nil {
		return m.Question
	}
	return ""
}

func (m *CreateVoteableRequest) GetAnswers() []string {
	if m != nil {
		return m.Answers
	}
	return nil
}

type CreateVoteableResponse struct {
	Uuid                 string   `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateVoteableResponse) Reset()         { *m = CreateVoteableResponse{} }
func (m *CreateVoteableResponse) String() string { return proto.CompactTextString(m) }
func (*CreateVoteableResponse) ProtoMessage()    {}
func (*CreateVoteableResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{2}
}

func (m *CreateVoteableResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateVoteableResponse.Unmarshal(m, b)
}
func (m *CreateVoteableResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateVoteableResponse.Marshal(b, m, deterministic)
}
func (m *CreateVoteableResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateVoteableResponse.Merge(m, src)
}
func (m *CreateVoteableResponse) XXX_Size() int {
	return xxx_messageInfo_CreateVoteableResponse.Size(m)
}
func (m *CreateVoteableResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateVoteableResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateVoteableResponse proto.InternalMessageInfo

func (m *CreateVoteableResponse) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

type ListVoteableRequest struct {
	Limit                int32    `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	LastIndex            int64    `protobuf:"varint,2,opt,name=last_index,json=lastIndex,proto3" json:"last_index,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListVoteableRequest) Reset()         { *m = ListVoteableRequest{} }
func (m *ListVoteableRequest) String() string { return proto.CompactTextString(m) }
func (*ListVoteableRequest) ProtoMessage()    {}
func (*ListVoteableRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{3}
}

func (m *ListVoteableRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListVoteableRequest.Unmarshal(m, b)
}
func (m *ListVoteableRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListVoteableRequest.Marshal(b, m, deterministic)
}
func (m *ListVoteableRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListVoteableRequest.Merge(m, src)
}
func (m *ListVoteableRequest) XXX_Size() int {
	return xxx_messageInfo_ListVoteableRequest.Size(m)
}
func (m *ListVoteableRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListVoteableRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListVoteableRequest proto.InternalMessageInfo

func (m *ListVoteableRequest) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *ListVoteableRequest) GetLastIndex() int64 {
	if m != nil {
		return m.LastIndex
	}
	return 0
}

type ListVoteableResponse struct {
	Votables             []*Voteable `protobuf:"bytes,1,rep,name=votables,proto3" json:"votables,omitempty"`
	LastIndex            int64       `protobuf:"varint,2,opt,name=last_index,json=lastIndex,proto3" json:"last_index,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ListVoteableResponse) Reset()         { *m = ListVoteableResponse{} }
func (m *ListVoteableResponse) String() string { return proto.CompactTextString(m) }
func (*ListVoteableResponse) ProtoMessage()    {}
func (*ListVoteableResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{4}
}

func (m *ListVoteableResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListVoteableResponse.Unmarshal(m, b)
}
func (m *ListVoteableResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListVoteableResponse.Marshal(b, m, deterministic)
}
func (m *ListVoteableResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListVoteableResponse.Merge(m, src)
}
func (m *ListVoteableResponse) XXX_Size() int {
	return xxx_messageInfo_ListVoteableResponse.Size(m)
}
func (m *ListVoteableResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListVoteableResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListVoteableResponse proto.InternalMessageInfo

func (m *ListVoteableResponse) GetVotables() []*Voteable {
	if m != nil {
		return m.Votables
	}
	return nil
}

func (m *ListVoteableResponse) GetLastIndex() int64 {
	if m != nil {
		return m.LastIndex
	}
	return 0
}

type CastVoteRequest struct {
	Uuid                 string   `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	AnswerIndex          int64    `protobuf:"varint,2,opt,name=answer_index,json=answerIndex,proto3" json:"answer_index,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CastVoteRequest) Reset()         { *m = CastVoteRequest{} }
func (m *CastVoteRequest) String() string { return proto.CompactTextString(m) }
func (*CastVoteRequest) ProtoMessage()    {}
func (*CastVoteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{5}
}

func (m *CastVoteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CastVoteRequest.Unmarshal(m, b)
}
func (m *CastVoteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CastVoteRequest.Marshal(b, m, deterministic)
}
func (m *CastVoteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CastVoteRequest.Merge(m, src)
}
func (m *CastVoteRequest) XXX_Size() int {
	return xxx_messageInfo_CastVoteRequest.Size(m)
}
func (m *CastVoteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CastVoteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CastVoteRequest proto.InternalMessageInfo

func (m *CastVoteRequest) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *CastVoteRequest) GetAnswerIndex() int64 {
	if m != nil {
		return m.AnswerIndex
	}
	return 0
}

type CastVoteResponse struct {
	Status               string   `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CastVoteResponse) Reset()         { *m = CastVoteResponse{} }
func (m *CastVoteResponse) String() string { return proto.CompactTextString(m) }
func (*CastVoteResponse) ProtoMessage()    {}
func (*CastVoteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{6}
}

func (m *CastVoteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CastVoteResponse.Unmarshal(m, b)
}
func (m *CastVoteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CastVoteResponse.Marshal(b, m, deterministic)
}
func (m *CastVoteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CastVoteResponse.Merge(m, src)
}
func (m *CastVoteResponse) XXX_Size() int {
	return xxx_messageInfo_CastVoteResponse.Size(m)
}
func (m *CastVoteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CastVoteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CastVoteResponse proto.InternalMessageInfo

func (m *CastVoteResponse) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func init() {
	proto.RegisterType((*Voteable)(nil), "api.Voteable")
	proto.RegisterType((*CreateVoteableRequest)(nil), "api.CreateVoteableRequest")
	proto.RegisterType((*CreateVoteableResponse)(nil), "api.CreateVoteableResponse")
	proto.RegisterType((*ListVoteableRequest)(nil), "api.ListVoteableRequest")
	proto.RegisterType((*ListVoteableResponse)(nil), "api.ListVoteableResponse")
	proto.RegisterType((*CastVoteRequest)(nil), "api.CastVoteRequest")
	proto.RegisterType((*CastVoteResponse)(nil), "api.CastVoteResponse")
}

func init() { proto.RegisterFile("service.proto", fileDescriptor_a0b84a42fa06f626) }

var fileDescriptor_a0b84a42fa06f626 = []byte{
	// 339 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0x4d, 0x4f, 0xc2, 0x40,
	0x10, 0x86, 0x2d, 0x15, 0x2c, 0x83, 0x55, 0xb3, 0x02, 0x59, 0x6b, 0x4c, 0xb0, 0x27, 0x34, 0x86,
	0x03, 0x9e, 0x3c, 0x93, 0x18, 0x3f, 0x2f, 0xd5, 0x70, 0xd5, 0x45, 0x36, 0x66, 0x13, 0xec, 0xd6,
	0xce, 0x16, 0xfd, 0xb7, 0xfe, 0x15, 0xc3, 0x6e, 0x5b, 0x5d, 0xb2, 0xc6, 0x5b, 0xe7, 0x9d, 0xe9,
	0xf3, 0xce, 0xc7, 0x42, 0x88, 0x3c, 0x5f, 0x8a, 0x17, 0x3e, 0xca, 0x72, 0xa9, 0x24, 0xf1, 0x59,
	0x26, 0xe2, 0x47, 0x08, 0xa6, 0x52, 0x71, 0x36, 0x5b, 0x70, 0x42, 0x60, 0xb3, 0x28, 0xc4, 0x9c,
	0x7a, 0x03, 0x6f, 0xd8, 0x4e, 0xf4, 0x37, 0x89, 0x20, 0x78, 0x2f, 0x38, 0x2a, 0x21, 0x53, 0xda,
	0xd0, 0x7a, 0x1d, 0x13, 0x0a, 0x5b, 0x2c, 0xc5, 0x0f, 0x9e, 0x23, 0xf5, 0x07, 0xfe, 0xb0, 0x9d,
	0x54, 0x61, 0x7c, 0x0f, 0xbd, 0x49, 0xce, 0x99, 0xe2, 0x15, 0x3b, 0xe1, 0xfa, 0x2f, 0x0b, 0xe7,
	0xfd, 0x8d, 0x6b, 0xd8, 0xb8, 0x33, 0xe8, 0xaf, 0xe3, 0x30, 0x93, 0x29, 0x3a, 0x5b, 0x8e, 0x6f,
	0x60, 0xff, 0x4e, 0xa0, 0x5a, 0xb7, 0xee, 0x42, 0x73, 0x21, 0xde, 0x84, 0xd2, 0xb5, 0xcd, 0xc4,
	0x04, 0xe4, 0x08, 0x60, 0xc1, 0x50, 0x3d, 0x89, 0x74, 0xce, 0x3f, 0xf5, 0x84, 0x7e, 0xd2, 0x5e,
	0x29, 0xd7, 0x2b, 0x21, 0x7e, 0x86, 0xae, 0xcd, 0x2a, 0x7d, 0x4f, 0x20, 0x58, 0x4a, 0xb5, 0x92,
	0x90, 0x7a, 0x03, 0x7f, 0xd8, 0x19, 0x87, 0x23, 0x96, 0x89, 0x51, 0x5d, 0x58, 0xa7, 0xff, 0x73,
	0xb8, 0x82, 0xdd, 0x09, 0x33, 0x0e, 0x55, 0xa7, 0xae, 0x3b, 0x1c, 0xc3, 0xb6, 0xd9, 0x86, 0xc5,
	0xe9, 0x18, 0xcd, 0x90, 0x4e, 0x61, 0xef, 0x87, 0x54, 0xf6, 0xd9, 0x87, 0x16, 0x2a, 0xa6, 0x0a,
	0x2c, 0x61, 0x65, 0x34, 0xfe, 0xf2, 0x20, 0x9c, 0x4a, 0x25, 0xd2, 0xd7, 0x07, 0xf3, 0x26, 0xc8,
	0x2d, 0xec, 0xd8, 0x3b, 0x26, 0x91, 0x9e, 0xc8, 0x79, 0xc7, 0xe8, 0xd0, 0x99, 0x33, 0xa6, 0xf1,
	0x06, 0xb9, 0x84, 0xf0, 0xf7, 0xda, 0x90, 0x50, 0x5d, 0xef, 0x38, 0x4b, 0x74, 0xe0, 0xc8, 0xd4,
	0x9c, 0x0b, 0x08, 0xaa, 0x91, 0x48, 0xd7, 0x58, 0xda, 0xbb, 0x8a, 0x7a, 0x6b, 0x6a, 0xf5, 0xeb,
	0xac, 0xa5, 0x1f, 0xf9, 0xf9, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x61, 0x11, 0x2f, 0x37, 0xf5,
	0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// VotingServiceClient is the client API for VotingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type VotingServiceClient interface {
	CreateVoteable(ctx context.Context, in *CreateVoteableRequest, opts ...grpc.CallOption) (*CreateVoteableResponse, error)
	ListVoteables(ctx context.Context, in *ListVoteableRequest, opts ...grpc.CallOption) (*ListVoteableResponse, error)
	CastVote(ctx context.Context, in *CastVoteRequest, opts ...grpc.CallOption) (*CastVoteResponse, error)
}

type votingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewVotingServiceClient(cc grpc.ClientConnInterface) VotingServiceClient {
	return &votingServiceClient{cc}
}

func (c *votingServiceClient) CreateVoteable(ctx context.Context, in *CreateVoteableRequest, opts ...grpc.CallOption) (*CreateVoteableResponse, error) {
	out := new(CreateVoteableResponse)
	err := c.cc.Invoke(ctx, "/api.VotingService/CreateVoteable", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *votingServiceClient) ListVoteables(ctx context.Context, in *ListVoteableRequest, opts ...grpc.CallOption) (*ListVoteableResponse, error) {
	out := new(ListVoteableResponse)
	err := c.cc.Invoke(ctx, "/api.VotingService/ListVoteables", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *votingServiceClient) CastVote(ctx context.Context, in *CastVoteRequest, opts ...grpc.CallOption) (*CastVoteResponse, error) {
	out := new(CastVoteResponse)
	err := c.cc.Invoke(ctx, "/api.VotingService/CastVote", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VotingServiceServer is the server API for VotingService service.
type VotingServiceServer interface {
	CreateVoteable(context.Context, *CreateVoteableRequest) (*CreateVoteableResponse, error)
	ListVoteables(context.Context, *ListVoteableRequest) (*ListVoteableResponse, error)
	CastVote(context.Context, *CastVoteRequest) (*CastVoteResponse, error)
}

// UnimplementedVotingServiceServer can be embedded to have forward compatible implementations.
type UnimplementedVotingServiceServer struct {
}

func (*UnimplementedVotingServiceServer) CreateVoteable(ctx context.Context, req *CreateVoteableRequest) (*CreateVoteableResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateVoteable not implemented")
}
func (*UnimplementedVotingServiceServer) ListVoteables(ctx context.Context, req *ListVoteableRequest) (*ListVoteableResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListVoteables not implemented")
}
func (*UnimplementedVotingServiceServer) CastVote(ctx context.Context, req *CastVoteRequest) (*CastVoteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CastVote not implemented")
}

func RegisterVotingServiceServer(s *grpc.Server, srv VotingServiceServer) {
	s.RegisterService(&_VotingService_serviceDesc, srv)
}

func _VotingService_CreateVoteable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateVoteableRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VotingServiceServer).CreateVoteable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.VotingService/CreateVoteable",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VotingServiceServer).CreateVoteable(ctx, req.(*CreateVoteableRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VotingService_ListVoteables_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListVoteableRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VotingServiceServer).ListVoteables(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.VotingService/ListVoteables",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VotingServiceServer).ListVoteables(ctx, req.(*ListVoteableRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VotingService_CastVote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CastVoteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VotingServiceServer).CastVote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.VotingService/CastVote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VotingServiceServer).CastVote(ctx, req.(*CastVoteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _VotingService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.VotingService",
	HandlerType: (*VotingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateVoteable",
			Handler:    _VotingService_CreateVoteable_Handler,
		},
		{
			MethodName: "ListVoteables",
			Handler:    _VotingService_ListVoteables_Handler,
		},
		{
			MethodName: "CastVote",
			Handler:    _VotingService_CastVote_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}
