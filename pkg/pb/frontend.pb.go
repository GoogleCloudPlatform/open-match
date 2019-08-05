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
	// 616 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xdd, 0x6e, 0xd3, 0x30,
	0x14, 0x56, 0x5a, 0x34, 0x36, 0x6f, 0x82, 0xc9, 0xfb, 0xd1, 0x96, 0x21, 0x61, 0xb2, 0x9b, 0xa9,
	0x5a, 0xe3, 0xae, 0xf4, 0xaa, 0x13, 0xd2, 0xc6, 0x36, 0xa6, 0x4a, 0x03, 0xa4, 0x82, 0x90, 0xe0,
	0x06, 0xb9, 0xc9, 0x59, 0x6a, 0xd6, 0xd8, 0x26, 0x76, 0x36, 0x24, 0xc4, 0x05, 0x3c, 0x02, 0xdc,
	0xf1, 0x18, 0x3c, 0x00, 0x2f, 0xc1, 0x15, 0xf7, 0xbc, 0x07, 0xa8, 0x4e, 0xda, 0x66, 0x5d, 0xa4,
	0x8d, 0xab, 0x28, 0xe7, 0x3b, 0xdf, 0xf9, 0x72, 0xbe, 0xcf, 0x31, 0xc2, 0x4c, 0x71, 0x7a, 0x9a,
	0x48, 0x61, 0x40, 0x84, 0xbe, 0x4a, 0xa4, 0x91, 0xb8, 0xca, 0x14, 0x77, 0x2d, 0x10, 0x83, 0xd6,
	0x2c, 0x02, 0x9d, 0x01, 0xee, 0xbd, 0x48, 0xca, 0x68, 0x00, 0x74, 0x08, 0x31, 0x21, 0xa4, 0x61,
	0x86, 0x4b, 0x31, 0x42, 0xb7, 0xed, 0x23, 0xa8, 0x47, 0x20, 0xea, 0xfa, 0x82, 0x45, 0x11, 0x24,
	0x54, 0x2a, 0xdb, 0x71, 0xb5, 0xdb, 0x6b, 0xa3, 0xa5, 0x83, 0x04, 0x98, 0x81, 0x97, 0x3c, 0x38,
	0x03, 0xd3, 0x85, 0xf7, 0x29, 0x68, 0x83, 0x37, 0xd1, 0x8c, 0xb1, 0x85, 0x35, 0x87, 0x38, 0x5b,
	0xf3, 0xcd, 0x79, 0x9f, 0x29, 0xee, 0xe7, 0x3d, 0x39, 0xe4, 0xed, 0xa2, 0xe5, 0xcb, 0x5c, 0xad,
	0xa4, 0xd0, 0x70, 0x33, 0x72, 0x13, 0x2d, 0x1d, 0xc2, 0x00, 0xa6, 0x85, 0x37, 0xd0, 0x5c, 0xd6,
	0xf0, 0x96, 0x87, 0x96, 0x3e, 0xd7, 0x9d, 0xcd, 0x0a, 0x9d, 0xd0, 0x5b, 0x45, 0xcb, 0x97, 0x39,
	0x99, 0xa0, 0x47, 0xd1, 0xe2, 0x31, 0x98, 0xff, 0x18, 0xd4, 0x42, 0x2b, 0xc7, 0x60, 0xf6, 0xb5,
	0xe6, 0x91, 0x88, 0x41, 0x18, 0x7d, 0x23, 0x56, 0x07, 0xad, 0x4e, 0xb3, 0xf2, 0x8d, 0x29, 0x42,
	0x6c, 0x5c, 0xce, 0xb7, 0xbe, 0x6b, 0xb7, 0x9e, 0x74, 0x77, 0x0b, 0x2d, 0xcd, 0x9f, 0x55, 0x34,
	0xfb, 0x24, 0x8f, 0x1b, 0x87, 0x68, 0xa1, 0xe8, 0x23, 0x5e, 0xb3, 0xcc, 0x92, 0x58, 0xdc, 0xf5,
	0x12, 0x24, 0xf7, 0xe0, 0xfe, 0x97, 0x5f, 0x7f, 0xbe, 0x55, 0xd6, 0xbd, 0x65, 0x7a, 0xbe, 0x33,
	0x3e, 0x49, 0x34, 0xfb, 0x76, 0xdd, 0x76, 0x6a, 0x38, 0x46, 0x0b, 0x45, 0xf3, 0x72, 0x95, 0x92,
	0x0c, 0x72, 0x95, 0x52, 0xa7, 0xb7, 0xac, 0x8a, 0x57, 0x23, 0x65, 0x2a, 0xf4, 0xe3, 0xd8, 0xbb,
	0x4f, 0xf8, 0x35, 0x9a, 0x1b, 0x67, 0x82, 0x57, 0xec, 0xc4, 0xe9, 0x8c, 0xdc, 0xe2, 0xc1, 0x18,
	0x8d, 0xc6, 0xd7, 0x8f, 0xfe, 0xec, 0xa0, 0x3b, 0x97, 0x83, 0xc0, 0xee, 0x48, 0xe0, 0x6a, 0xa6,
	0xee, 0x46, 0x29, 0x96, 0x2f, 0xd4, 0xb2, 0xaa, 0x3e, 0xde, 0xbe, 0x4e, 0x95, 0x4e, 0xd2, 0xd3,
	0x0d, 0xe7, 0xf1, 0xdf, 0xca, 0xd7, 0xfd, 0xdf, 0x15, 0xfc, 0xc3, 0x99, 0xe4, 0xe8, 0x75, 0x10,
	0x7a, 0xae, 0x40, 0x90, 0xa7, 0xcc, 0x04, 0x7d, 0xbc, 0xda, 0x37, 0x46, 0xe9, 0x36, 0xa5, 0x52,
	0x81, 0xa8, 0xc7, 0xc3, 0x9a, 0x1f, 0xc2, 0xb9, 0xbb, 0x39, 0x79, 0xaf, 0x87, 0x5c, 0x07, 0xa9,
	0xd6, 0x7b, 0xd9, 0x0f, 0x1d, 0x25, 0x32, 0x55, 0xda, 0x0f, 0x64, 0x5c, 0x7b, 0x85, 0xf0, 0xbe,
	0x62, 0x41, 0x1f, 0x48, 0xd3, 0x6f, 0x90, 0x13, 0x1e, 0xc0, 0xf0, 0x98, 0xed, 0x8d, 0x46, 0x46,
	0xdc, 0xf4, 0xd3, 0xde, 0xb0, 0x93, 0x66, 0xd4, 0x53, 0x99, 0x44, 0x2c, 0x06, 0x5d, 0x10, 0xa3,
	0xbd, 0x81, 0xec, 0xd1, 0x98, 0x69, 0x03, 0x09, 0x3d, 0xe9, 0x1c, 0x1c, 0x3d, 0x7b, 0x71, 0xd4,
	0xac, 0xee, 0xf8, 0x8d, 0x5a, 0xc5, 0xa9, 0x34, 0x17, 0x99, 0x52, 0x03, 0x1e, 0xd8, 0xbb, 0x80,
	0xbe, 0xd3, 0x52, 0xb4, 0xaf, 0x54, 0xba, 0xbb, 0xa8, 0xda, 0x6a, 0xb4, 0x70, 0x0b, 0xd5, 0xba,
	0x60, 0xd2, 0x44, 0x40, 0x48, 0x2e, 0xfa, 0x20, 0x88, 0xe9, 0x03, 0x49, 0x40, 0xcb, 0x34, 0x09,
	0x80, 0x84, 0x12, 0x34, 0x11, 0xd2, 0x10, 0xf8, 0xc0, 0xb5, 0xf1, 0xf1, 0x0c, 0xba, 0xf5, 0xbd,
	0xe2, 0xdc, 0x4e, 0x1e, 0xa1, 0xb5, 0x89, 0x19, 0xe4, 0x50, 0x06, 0xe9, 0xd0, 0x3a, 0x3b, 0x1d,
	0x3f, 0x28, 0xb7, 0x86, 0x6a, 0x6e, 0x80, 0x86, 0x32, 0xd0, 0xf4, 0x0d, 0x99, 0x82, 0x0a, 0x7b,
	0xa9, 0xb3, 0x88, 0xaa, 0x5e, 0x6f, 0xc6, 0x5e, 0x60, 0x0f, 0xff, 0x05, 0x00, 0x00, 0xff, 0xff,
	0x7a, 0x76, 0xcc, 0x85, 0x3b, 0x05, 0x00, 0x00,
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
