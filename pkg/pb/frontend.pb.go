// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/frontend.proto

package pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type CreateTicketRequest struct {
	// Ticket object with the properties of the Ticket to be created.
	Ticket               *Ticket  `protobuf:"bytes,1,opt,name=ticket,proto3" json:"ticket,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateTicketRequest) Reset()         { *m = CreateTicketRequest{} }
func (m *CreateTicketRequest) String() string { return proto.CompactTextString(m) }
func (*CreateTicketRequest) ProtoMessage()    {}
func (*CreateTicketRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_06c902cf58d2ae57, []int{0}
}

func (m *CreateTicketRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateTicketRequest.Unmarshal(m, b)
}
func (m *CreateTicketRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateTicketRequest.Marshal(b, m, deterministic)
}
func (m *CreateTicketRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateTicketRequest.Merge(m, src)
}
func (m *CreateTicketRequest) XXX_Size() int {
	return xxx_messageInfo_CreateTicketRequest.Size(m)
}
func (m *CreateTicketRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateTicketRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateTicketRequest proto.InternalMessageInfo

func (m *CreateTicketRequest) GetTicket() *Ticket {
	if m != nil {
		return m.Ticket
	}
	return nil
}

type CreateTicketResponse struct {
	// Ticket object for the created Ticket - with the ticket ID populated.
	Ticket               *Ticket  `protobuf:"bytes,1,opt,name=ticket,proto3" json:"ticket,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateTicketResponse) Reset()         { *m = CreateTicketResponse{} }
func (m *CreateTicketResponse) String() string { return proto.CompactTextString(m) }
func (*CreateTicketResponse) ProtoMessage()    {}
func (*CreateTicketResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_06c902cf58d2ae57, []int{1}
}

func (m *CreateTicketResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateTicketResponse.Unmarshal(m, b)
}
func (m *CreateTicketResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateTicketResponse.Marshal(b, m, deterministic)
}
func (m *CreateTicketResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateTicketResponse.Merge(m, src)
}
func (m *CreateTicketResponse) XXX_Size() int {
	return xxx_messageInfo_CreateTicketResponse.Size(m)
}
func (m *CreateTicketResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateTicketResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateTicketResponse proto.InternalMessageInfo

func (m *CreateTicketResponse) GetTicket() *Ticket {
	if m != nil {
		return m.Ticket
	}
	return nil
}

type DeleteTicketRequest struct {
	// Ticket ID of the Ticket to be deleted.
	TicketId             string   `protobuf:"bytes,1,opt,name=ticket_id,json=ticketId,proto3" json:"ticket_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteTicketRequest) Reset()         { *m = DeleteTicketRequest{} }
func (m *DeleteTicketRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteTicketRequest) ProtoMessage()    {}
func (*DeleteTicketRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_06c902cf58d2ae57, []int{2}
}

func (m *DeleteTicketRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteTicketRequest.Unmarshal(m, b)
}
func (m *DeleteTicketRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteTicketRequest.Marshal(b, m, deterministic)
}
func (m *DeleteTicketRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteTicketRequest.Merge(m, src)
}
func (m *DeleteTicketRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteTicketRequest.Size(m)
}
func (m *DeleteTicketRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteTicketRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteTicketRequest proto.InternalMessageInfo

func (m *DeleteTicketRequest) GetTicketId() string {
	if m != nil {
		return m.TicketId
	}
	return ""
}

type DeleteTicketResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteTicketResponse) Reset()         { *m = DeleteTicketResponse{} }
func (m *DeleteTicketResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteTicketResponse) ProtoMessage()    {}
func (*DeleteTicketResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_06c902cf58d2ae57, []int{3}
}

func (m *DeleteTicketResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteTicketResponse.Unmarshal(m, b)
}
func (m *DeleteTicketResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteTicketResponse.Marshal(b, m, deterministic)
}
func (m *DeleteTicketResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteTicketResponse.Merge(m, src)
}
func (m *DeleteTicketResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteTicketResponse.Size(m)
}
func (m *DeleteTicketResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteTicketResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteTicketResponse proto.InternalMessageInfo

type GetTicketRequest struct {
	// Ticket ID of the Ticket to fetch.
	TicketId             string   `protobuf:"bytes,1,opt,name=ticket_id,json=ticketId,proto3" json:"ticket_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTicketRequest) Reset()         { *m = GetTicketRequest{} }
func (m *GetTicketRequest) String() string { return proto.CompactTextString(m) }
func (*GetTicketRequest) ProtoMessage()    {}
func (*GetTicketRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_06c902cf58d2ae57, []int{4}
}

func (m *GetTicketRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTicketRequest.Unmarshal(m, b)
}
func (m *GetTicketRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTicketRequest.Marshal(b, m, deterministic)
}
func (m *GetTicketRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTicketRequest.Merge(m, src)
}
func (m *GetTicketRequest) XXX_Size() int {
	return xxx_messageInfo_GetTicketRequest.Size(m)
}
func (m *GetTicketRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTicketRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetTicketRequest proto.InternalMessageInfo

func (m *GetTicketRequest) GetTicketId() string {
	if m != nil {
		return m.TicketId
	}
	return ""
}

type GetAssignmentsRequest struct {
	// Ticket ID of the Ticket to get updates on.
	TicketId             string   `protobuf:"bytes,1,opt,name=ticket_id,json=ticketId,proto3" json:"ticket_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAssignmentsRequest) Reset()         { *m = GetAssignmentsRequest{} }
func (m *GetAssignmentsRequest) String() string { return proto.CompactTextString(m) }
func (*GetAssignmentsRequest) ProtoMessage()    {}
func (*GetAssignmentsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_06c902cf58d2ae57, []int{5}
}

func (m *GetAssignmentsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAssignmentsRequest.Unmarshal(m, b)
}
func (m *GetAssignmentsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAssignmentsRequest.Marshal(b, m, deterministic)
}
func (m *GetAssignmentsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAssignmentsRequest.Merge(m, src)
}
func (m *GetAssignmentsRequest) XXX_Size() int {
	return xxx_messageInfo_GetAssignmentsRequest.Size(m)
}
func (m *GetAssignmentsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAssignmentsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAssignmentsRequest proto.InternalMessageInfo

func (m *GetAssignmentsRequest) GetTicketId() string {
	if m != nil {
		return m.TicketId
	}
	return ""
}

type GetAssignmentsResponse struct {
	// The updated Ticket object.
	Assignment           *Assignment `protobuf:"bytes,1,opt,name=assignment,proto3" json:"assignment,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *GetAssignmentsResponse) Reset()         { *m = GetAssignmentsResponse{} }
func (m *GetAssignmentsResponse) String() string { return proto.CompactTextString(m) }
func (*GetAssignmentsResponse) ProtoMessage()    {}
func (*GetAssignmentsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_06c902cf58d2ae57, []int{6}
}

func (m *GetAssignmentsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAssignmentsResponse.Unmarshal(m, b)
}
func (m *GetAssignmentsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAssignmentsResponse.Marshal(b, m, deterministic)
}
func (m *GetAssignmentsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAssignmentsResponse.Merge(m, src)
}
func (m *GetAssignmentsResponse) XXX_Size() int {
	return xxx_messageInfo_GetAssignmentsResponse.Size(m)
}
func (m *GetAssignmentsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAssignmentsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetAssignmentsResponse proto.InternalMessageInfo

func (m *GetAssignmentsResponse) GetAssignment() *Assignment {
	if m != nil {
		return m.Assignment
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateTicketRequest)(nil), "api.CreateTicketRequest")
	proto.RegisterType((*CreateTicketResponse)(nil), "api.CreateTicketResponse")
	proto.RegisterType((*DeleteTicketRequest)(nil), "api.DeleteTicketRequest")
	proto.RegisterType((*DeleteTicketResponse)(nil), "api.DeleteTicketResponse")
	proto.RegisterType((*GetTicketRequest)(nil), "api.GetTicketRequest")
	proto.RegisterType((*GetAssignmentsRequest)(nil), "api.GetAssignmentsRequest")
	proto.RegisterType((*GetAssignmentsResponse)(nil), "api.GetAssignmentsResponse")
}

func init() { proto.RegisterFile("api/frontend.proto", fileDescriptor_06c902cf58d2ae57) }

var fileDescriptor_06c902cf58d2ae57 = []byte{
	// 612 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xcb, 0x6e, 0xd3, 0x40,
	0x14, 0x95, 0x13, 0x14, 0xda, 0x69, 0x05, 0xd5, 0xf4, 0xa1, 0x36, 0x45, 0xc2, 0xb8, 0x9b, 0x2a,
	0x6a, 0x3c, 0x69, 0xc8, 0x2a, 0x15, 0x52, 0x4b, 0x5b, 0xaa, 0x48, 0x05, 0xa4, 0x80, 0x90, 0x60,
	0x83, 0x26, 0xf6, 0xad, 0x33, 0x34, 0x9e, 0x19, 0x7c, 0xc7, 0x2d, 0x12, 0x62, 0x01, 0x9f, 0x00,
	0x3b, 0x3e, 0x83, 0x0f, 0xe0, 0x27, 0x58, 0xb1, 0x87, 0xff, 0x40, 0x19, 0x3b, 0x8f, 0xb6, 0x96,
	0x5a, 0x56, 0x96, 0xef, 0x3d, 0xe7, 0x1e, 0xdf, 0x73, 0xc6, 0x43, 0x28, 0xd7, 0x82, 0x9d, 0x24,
	0x4a, 0x1a, 0x90, 0xa1, 0xaf, 0x13, 0x65, 0x14, 0x2d, 0x73, 0x2d, 0xaa, 0xb6, 0x11, 0x03, 0x22,
	0x8f, 0x00, 0xb3, 0x46, 0xf5, 0x5e, 0xa4, 0x54, 0x34, 0x00, 0x36, 0x6c, 0x71, 0x29, 0x95, 0xe1,
	0x46, 0x28, 0x39, 0xea, 0x6e, 0xd9, 0x47, 0x50, 0x8f, 0x40, 0xd6, 0xf1, 0x9c, 0x47, 0x11, 0x24,
	0x4c, 0x69, 0x8b, 0xb8, 0x8a, 0xf6, 0xda, 0x64, 0x71, 0x3f, 0x01, 0x6e, 0xe0, 0xa5, 0x08, 0x4e,
	0xc1, 0x74, 0xe1, 0x7d, 0x0a, 0x68, 0xe8, 0x06, 0xa9, 0x18, 0x5b, 0x58, 0x75, 0x5c, 0x67, 0x73,
	0xae, 0x39, 0xe7, 0x73, 0x2d, 0xfc, 0x1c, 0x93, 0xb7, 0xbc, 0x1d, 0xb2, 0x74, 0x91, 0x8b, 0x5a,
	0x49, 0x84, 0x9b, 0x91, 0x9b, 0x64, 0xf1, 0x00, 0x06, 0x70, 0x59, 0x78, 0x9d, 0xcc, 0x66, 0x80,
	0xb7, 0x22, 0xb4, 0xf4, 0xd9, 0xee, 0x4c, 0x56, 0xe8, 0x84, 0xde, 0x0a, 0x59, 0xba, 0xc8, 0xc9,
	0x04, 0x3d, 0x46, 0x16, 0x8e, 0xc0, 0xfc, 0xc7, 0xa0, 0x16, 0x59, 0x3e, 0x02, 0xb3, 0x87, 0x28,
	0x22, 0x19, 0x83, 0x34, 0x78, 0x23, 0x56, 0x87, 0xac, 0x5c, 0x66, 0xe5, 0x1b, 0x33, 0x42, 0xf8,
	0xb8, 0x9c, 0x6f, 0x7d, 0xd7, 0x6e, 0x3d, 0x41, 0x77, 0xa7, 0x20, 0xcd, 0x9f, 0x65, 0x32, 0xf3,
	0x24, 0x8f, 0x9b, 0x86, 0x64, 0x7e, 0xda, 0x47, 0xba, 0x6a, 0x99, 0x05, 0xb1, 0x54, 0xd7, 0x0a,
	0x3a, 0xb9, 0x07, 0xf7, 0xbf, 0xfc, 0xfa, 0xf3, 0xad, 0xb4, 0xe6, 0x2d, 0xb1, 0xb3, 0xed, 0xf1,
	0x49, 0x62, 0xd9, 0xb7, 0x63, 0xdb, 0xa9, 0xd1, 0x98, 0xcc, 0x4f, 0x9b, 0x97, 0xab, 0x14, 0x64,
	0x90, 0xab, 0x14, 0x3a, 0xbd, 0x69, 0x55, 0xbc, 0x9a, 0x5b, 0xa4, 0xc2, 0x3e, 0x8e, 0xbd, 0xfb,
	0x44, 0x5f, 0x93, 0xd9, 0x71, 0x26, 0x74, 0xd9, 0x4e, 0xbc, 0x9c, 0x51, 0x75, 0xfa, 0x60, 0x8c,
	0x46, 0xd3, 0xeb, 0x47, 0x7f, 0x76, 0xc8, 0x9d, 0x8b, 0x41, 0xd0, 0xea, 0x48, 0xe0, 0x6a, 0xa6,
	0xd5, 0xf5, 0xc2, 0x5e, 0xbe, 0x50, 0xcb, 0xaa, 0xfa, 0x74, 0xeb, 0x3a, 0x55, 0x36, 0x49, 0x0f,
	0x1b, 0xce, 0xe3, 0xbf, 0xa5, 0xaf, 0x7b, 0xbf, 0x4b, 0xf4, 0x87, 0x33, 0xc9, 0xd1, 0xeb, 0x10,
	0xf2, 0x5c, 0x83, 0x74, 0x9f, 0x72, 0x13, 0xf4, 0xe9, 0x4a, 0xdf, 0x18, 0x8d, 0x6d, 0xc6, 0x94,
	0x06, 0x59, 0x8f, 0x87, 0x35, 0x3f, 0x84, 0xb3, 0xea, 0xc6, 0xe4, 0xbd, 0x1e, 0x0a, 0x0c, 0x52,
	0xc4, 0xdd, 0xec, 0x87, 0x8e, 0x12, 0x95, 0x6a, 0xf4, 0x03, 0x15, 0xd7, 0x5e, 0x11, 0xba, 0xa7,
	0x79, 0xd0, 0x07, 0xb7, 0xe9, 0x37, 0xdc, 0x63, 0x11, 0xc0, 0xf0, 0x98, 0xed, 0x8e, 0x46, 0x46,
	0xc2, 0xf4, 0xd3, 0xde, 0x10, 0xc9, 0x32, 0xea, 0x89, 0x4a, 0x22, 0x1e, 0x03, 0x4e, 0x89, 0xb1,
	0xde, 0x40, 0xf5, 0x58, 0xcc, 0xd1, 0x40, 0xc2, 0x8e, 0x3b, 0xfb, 0x87, 0xcf, 0x5e, 0x1c, 0x36,
	0xcb, 0xdb, 0x7e, 0xa3, 0x56, 0x72, 0x4a, 0xcd, 0x05, 0xae, 0xf5, 0x40, 0x04, 0xf6, 0x2e, 0x60,
	0xef, 0x50, 0xc9, 0xf6, 0x95, 0x4a, 0x77, 0x87, 0x94, 0x5b, 0x8d, 0x16, 0x6d, 0x91, 0x5a, 0x17,
	0x4c, 0x9a, 0x48, 0x08, 0xdd, 0xf3, 0x3e, 0x48, 0xd7, 0xf4, 0xc1, 0x4d, 0x00, 0x55, 0x9a, 0x04,
	0xe0, 0x86, 0x0a, 0xd0, 0x95, 0xca, 0xb8, 0xf0, 0x41, 0xa0, 0xf1, 0x69, 0x85, 0xdc, 0xfa, 0x5e,
	0x72, 0x6e, 0x27, 0x8f, 0xc8, 0xea, 0xc4, 0x0c, 0xf7, 0x40, 0x05, 0xe9, 0xd0, 0x3a, 0x3b, 0x9d,
	0x3e, 0x28, 0xb6, 0x86, 0xa1, 0x30, 0xc0, 0x42, 0x15, 0x20, 0x7b, 0x53, 0xd1, 0xa7, 0x11, 0xd3,
	0xbd, 0x5e, 0xc5, 0x5e, 0x53, 0x0f, 0xff, 0x05, 0x00, 0x00, 0xff, 0xff, 0x57, 0x7c, 0x24, 0x88,
	0x21, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// FrontendClient is the client API for Frontend service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FrontendClient interface {
	// CreateTicket will create a new ticket, assign a Ticket ID to it and put the
	// Ticket in state storage. It will then look through the 'properties' field
	// for the attributes defined as indices the matchmakaking config. If the
	// attributes exist and are valid integers, they will be indexed. Creating a
	// ticket adds the Ticket to the pool of Tickets considered for matchmaking.
	CreateTicket(ctx context.Context, in *CreateTicketRequest, opts ...grpc.CallOption) (*CreateTicketResponse, error)
	// DeleteTicket removes the Ticket from state storage and from corresponding
	// configured indices and lazily removes the ticket from state storage.
	// Deleting a ticket immediately stops the ticket from being
	// considered for future matchmaking requests, yet when the ticket itself will be deleted
	// is undeterministic. Users may still be able to assign/get a ticket after calling DeleteTicket on it.
	DeleteTicket(ctx context.Context, in *DeleteTicketRequest, opts ...grpc.CallOption) (*DeleteTicketResponse, error)
	// GetTicket fetches the ticket associated with the specified Ticket ID.
	GetTicket(ctx context.Context, in *GetTicketRequest, opts ...grpc.CallOption) (*Ticket, error)
	// GetAssignments streams matchmaking results from Open Match for the
	// provided Ticket ID.
	GetAssignments(ctx context.Context, in *GetAssignmentsRequest, opts ...grpc.CallOption) (Frontend_GetAssignmentsClient, error)
}

type frontendClient struct {
	cc *grpc.ClientConn
}

func NewFrontendClient(cc *grpc.ClientConn) FrontendClient {
	return &frontendClient{cc}
}

func (c *frontendClient) CreateTicket(ctx context.Context, in *CreateTicketRequest, opts ...grpc.CallOption) (*CreateTicketResponse, error) {
	out := new(CreateTicketResponse)
	err := c.cc.Invoke(ctx, "/api.Frontend/CreateTicket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *frontendClient) DeleteTicket(ctx context.Context, in *DeleteTicketRequest, opts ...grpc.CallOption) (*DeleteTicketResponse, error) {
	out := new(DeleteTicketResponse)
	err := c.cc.Invoke(ctx, "/api.Frontend/DeleteTicket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *frontendClient) GetTicket(ctx context.Context, in *GetTicketRequest, opts ...grpc.CallOption) (*Ticket, error) {
	out := new(Ticket)
	err := c.cc.Invoke(ctx, "/api.Frontend/GetTicket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *frontendClient) GetAssignments(ctx context.Context, in *GetAssignmentsRequest, opts ...grpc.CallOption) (Frontend_GetAssignmentsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Frontend_serviceDesc.Streams[0], "/api.Frontend/GetAssignments", opts...)
	if err != nil {
		return nil, err
	}
	x := &frontendGetAssignmentsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Frontend_GetAssignmentsClient interface {
	Recv() (*GetAssignmentsResponse, error)
	grpc.ClientStream
}

type frontendGetAssignmentsClient struct {
	grpc.ClientStream
}

func (x *frontendGetAssignmentsClient) Recv() (*GetAssignmentsResponse, error) {
	m := new(GetAssignmentsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// FrontendServer is the server API for Frontend service.
type FrontendServer interface {
	// CreateTicket will create a new ticket, assign a Ticket ID to it and put the
	// Ticket in state storage. It will then look through the 'properties' field
	// for the attributes defined as indices the matchmakaking config. If the
	// attributes exist and are valid integers, they will be indexed. Creating a
	// ticket adds the Ticket to the pool of Tickets considered for matchmaking.
	CreateTicket(context.Context, *CreateTicketRequest) (*CreateTicketResponse, error)
	// DeleteTicket removes the Ticket from state storage and from corresponding
	// configured indices and lazily removes the ticket from state storage.
	// Deleting a ticket immediately stops the ticket from being
	// considered for future matchmaking requests, yet when the ticket itself will be deleted
	// is undeterministic. Users may still be able to assign/get a ticket after calling DeleteTicket on it.
	DeleteTicket(context.Context, *DeleteTicketRequest) (*DeleteTicketResponse, error)
	// GetTicket fetches the ticket associated with the specified Ticket ID.
	GetTicket(context.Context, *GetTicketRequest) (*Ticket, error)
	// GetAssignments streams matchmaking results from Open Match for the
	// provided Ticket ID.
	GetAssignments(*GetAssignmentsRequest, Frontend_GetAssignmentsServer) error
}

func RegisterFrontendServer(s *grpc.Server, srv FrontendServer) {
	s.RegisterService(&_Frontend_serviceDesc, srv)
}

func _Frontend_CreateTicket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTicketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FrontendServer).CreateTicket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Frontend/CreateTicket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FrontendServer).CreateTicket(ctx, req.(*CreateTicketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Frontend_DeleteTicket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTicketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FrontendServer).DeleteTicket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Frontend/DeleteTicket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FrontendServer).DeleteTicket(ctx, req.(*DeleteTicketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Frontend_GetTicket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTicketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FrontendServer).GetTicket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Frontend/GetTicket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FrontendServer).GetTicket(ctx, req.(*GetTicketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Frontend_GetAssignments_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetAssignmentsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FrontendServer).GetAssignments(m, &frontendGetAssignmentsServer{stream})
}

type Frontend_GetAssignmentsServer interface {
	Send(*GetAssignmentsResponse) error
	grpc.ServerStream
}

type frontendGetAssignmentsServer struct {
	grpc.ServerStream
}

func (x *frontendGetAssignmentsServer) Send(m *GetAssignmentsResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _Frontend_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.Frontend",
	HandlerType: (*FrontendServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTicket",
			Handler:    _Frontend_CreateTicket_Handler,
		},
		{
			MethodName: "DeleteTicket",
			Handler:    _Frontend_DeleteTicket_Handler,
		},
		{
			MethodName: "GetTicket",
			Handler:    _Frontend_GetTicket_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetAssignments",
			Handler:       _Frontend_GetAssignments_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "api/frontend.proto",
}
