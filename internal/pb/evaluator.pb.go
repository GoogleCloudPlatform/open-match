// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/evaluator.proto

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

type EvaluateRequest struct {
	// List of Matches to evaluate.
	Match                []*Match `protobuf:"bytes,1,rep,name=match,proto3" json:"match,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EvaluateRequest) Reset()         { *m = EvaluateRequest{} }
func (m *EvaluateRequest) String() string { return proto.CompactTextString(m) }
func (*EvaluateRequest) ProtoMessage()    {}
func (*EvaluateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c58cb7dff9acb0f, []int{0}
}

func (m *EvaluateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EvaluateRequest.Unmarshal(m, b)
}
func (m *EvaluateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EvaluateRequest.Marshal(b, m, deterministic)
}
func (m *EvaluateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EvaluateRequest.Merge(m, src)
}
func (m *EvaluateRequest) XXX_Size() int {
	return xxx_messageInfo_EvaluateRequest.Size(m)
}
func (m *EvaluateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EvaluateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EvaluateRequest proto.InternalMessageInfo

func (m *EvaluateRequest) GetMatch() []*Match {
	if m != nil {
		return m.Match
	}
	return nil
}

type EvaluateResponse struct {
	// Accepted list of Matches.
	Match                []*Match `protobuf:"bytes,1,rep,name=match,proto3" json:"match,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EvaluateResponse) Reset()         { *m = EvaluateResponse{} }
func (m *EvaluateResponse) String() string { return proto.CompactTextString(m) }
func (*EvaluateResponse) ProtoMessage()    {}
func (*EvaluateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c58cb7dff9acb0f, []int{1}
}

func (m *EvaluateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EvaluateResponse.Unmarshal(m, b)
}
func (m *EvaluateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EvaluateResponse.Marshal(b, m, deterministic)
}
func (m *EvaluateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EvaluateResponse.Merge(m, src)
}
func (m *EvaluateResponse) XXX_Size() int {
	return xxx_messageInfo_EvaluateResponse.Size(m)
}
func (m *EvaluateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EvaluateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EvaluateResponse proto.InternalMessageInfo

func (m *EvaluateResponse) GetMatch() []*Match {
	if m != nil {
		return m.Match
	}
	return nil
}

func init() {
	proto.RegisterType((*EvaluateRequest)(nil), "api.EvaluateRequest")
	proto.RegisterType((*EvaluateResponse)(nil), "api.EvaluateResponse")
}

func init() { proto.RegisterFile("api/evaluator.proto", fileDescriptor_8c58cb7dff9acb0f) }

var fileDescriptor_8c58cb7dff9acb0f = []byte{
	// 453 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0x4d, 0x6f, 0xd4, 0x30,
	0x10, 0x55, 0xb2, 0x50, 0xc0, 0x3d, 0x50, 0x99, 0x0f, 0x55, 0x2b, 0x84, 0x4c, 0x7a, 0x81, 0x15,
	0x1b, 0xb7, 0xe9, 0x9e, 0x16, 0x21, 0xb5, 0xc0, 0x1e, 0x2a, 0x15, 0x90, 0x16, 0x89, 0x03, 0x37,
	0xc7, 0x19, 0x12, 0xa3, 0xc4, 0x63, 0x3c, 0xce, 0x96, 0x33, 0x3f, 0x01, 0x6e, 0xfc, 0x0d, 0x7e,
	0x0a, 0x37, 0xce, 0xdc, 0xf8, 0x13, 0x28, 0x09, 0xab, 0xad, 0x5a, 0x24, 0x4e, 0x96, 0xdf, 0x7b,
	0x33, 0x6f, 0xfc, 0x3c, 0xec, 0x96, 0x72, 0x46, 0xc2, 0x4a, 0xd5, 0xad, 0x0a, 0xe8, 0x53, 0xe7,
	0x31, 0x20, 0x1f, 0x29, 0x67, 0xc6, 0xbc, 0x63, 0x1a, 0x20, 0x52, 0x25, 0xd0, 0x40, 0x8c, 0xef,
	0x95, 0x88, 0x65, 0x0d, 0xb2, 0xa3, 0x94, 0xb5, 0x18, 0x54, 0x30, 0x68, 0xd7, 0xec, 0xe3, 0xfe,
	0xd0, 0xd3, 0x12, 0xec, 0x94, 0xce, 0x54, 0x59, 0x82, 0x97, 0xe8, 0x7a, 0xc5, 0x65, 0x75, 0x72,
	0xc8, 0x6e, 0x2e, 0x06, 0x5f, 0x58, 0xc2, 0xc7, 0x16, 0x28, 0x70, 0xc1, 0xae, 0x36, 0x2a, 0xe8,
	0x6a, 0x37, 0x12, 0xa3, 0x87, 0xdb, 0x19, 0x4b, 0x95, 0x33, 0xe9, 0xcb, 0x0e, 0x59, 0x0e, 0x44,
	0x32, 0x63, 0x3b, 0x9b, 0x22, 0x72, 0x68, 0x09, 0xfe, 0x5f, 0x95, 0x21, 0xbb, 0xb1, 0x58, 0x3f,
	0x91, 0xe7, 0xec, 0xfa, 0xba, 0x05, 0xbf, 0xdd, 0x6b, 0x2f, 0x8c, 0x31, 0xbe, 0x73, 0x01, 0x1d,
	0x7c, 0x92, 0x47, 0x9f, 0x7f, 0xfc, 0xfa, 0x1a, 0xef, 0x25, 0xf7, 0xe5, 0xea, 0x60, 0x13, 0x99,
	0xec, 0x2d, 0x80, 0xe6, 0x7f, 0x11, 0x98, 0x47, 0x93, 0x67, 0xbf, 0xe3, 0x2f, 0xc7, 0x3f, 0x63,
	0xfe, 0x3d, 0x3a, 0x67, 0x9c, 0x9c, 0x30, 0xf6, 0xda, 0x81, 0x15, 0xfd, 0x68, 0xfc, 0x6e, 0x15,
	0x82, 0xa3, 0xb9, 0x94, 0xe8, 0xc0, 0x4e, 0xfb, 0x26, 0x69, 0x01, 0xab, 0xf1, 0xde, 0xe6, 0x3e,
	0x2d, 0x0c, 0xe9, 0x96, 0xe8, 0x68, 0x48, 0xbd, 0xf4, 0xd8, 0x3a, 0x4a, 0x35, 0x36, 0x93, 0xb7,
	0x8c, 0x1f, 0x3b, 0xa5, 0x2b, 0x10, 0x59, 0xba, 0x2f, 0x4e, 0x8d, 0x86, 0x2e, 0x88, 0xa3, 0x75,
	0xcb, 0xd2, 0x84, 0xaa, 0xcd, 0x3b, 0xa5, 0x1c, 0x4a, 0xdf, 0xa3, 0x2f, 0x55, 0x03, 0x74, 0xce,
	0x4c, 0xe6, 0x35, 0xe6, 0xb2, 0x51, 0x14, 0xc0, 0xcb, 0xd3, 0x93, 0xe7, 0x8b, 0x57, 0x6f, 0x16,
	0xd9, 0xe8, 0x20, 0xdd, 0x9f, 0xc4, 0x51, 0x9c, 0xed, 0x28, 0xe7, 0x6a, 0xa3, 0xfb, 0x0f, 0x93,
	0x1f, 0x08, 0xed, 0xfc, 0x12, 0xb2, 0x7c, 0xc2, 0x46, 0xb3, 0xfd, 0x19, 0x9f, 0xb1, 0xc9, 0x12,
	0x42, 0xeb, 0x2d, 0x14, 0xe2, 0xac, 0x02, 0x2b, 0x42, 0x05, 0xc2, 0x03, 0x61, 0xeb, 0x35, 0x88,
	0x02, 0x81, 0x84, 0xc5, 0x20, 0xe0, 0x93, 0xa1, 0x90, 0xf2, 0x2d, 0x76, 0xe5, 0x5b, 0x1c, 0x5d,
	0xf3, 0x4f, 0xd9, 0xee, 0x26, 0x0c, 0xf1, 0x02, 0x75, 0xdb, 0x80, 0x1d, 0x16, 0x84, 0x3f, 0xf8,
	0x77, 0x34, 0x92, 0x4c, 0x00, 0x59, 0xa0, 0x26, 0xf9, 0x6e, 0xdb, 0xd8, 0x00, 0xde, 0xaa, 0x5a,
	0xba, 0x3c, 0xdf, 0xea, 0x17, 0xea, 0xf0, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x5d, 0xff, 0xc7,
	0x1f, 0xcc, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// EvaluatorClient is the client API for Evaluator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EvaluatorClient interface {
	// Evaluate accepts a list of proposed matches, evaluates them for quality,
	// collisions etc. and returns matches that should be accepted as results.
	Evaluate(ctx context.Context, in *EvaluateRequest, opts ...grpc.CallOption) (*EvaluateResponse, error)
}

type evaluatorClient struct {
	cc *grpc.ClientConn
}

func NewEvaluatorClient(cc *grpc.ClientConn) EvaluatorClient {
	return &evaluatorClient{cc}
}

func (c *evaluatorClient) Evaluate(ctx context.Context, in *EvaluateRequest, opts ...grpc.CallOption) (*EvaluateResponse, error) {
	out := new(EvaluateResponse)
	err := c.cc.Invoke(ctx, "/api.Evaluator/Evaluate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EvaluatorServer is the server API for Evaluator service.
type EvaluatorServer interface {
	// Evaluate accepts a list of proposed matches, evaluates them for quality,
	// collisions etc. and returns matches that should be accepted as results.
	Evaluate(context.Context, *EvaluateRequest) (*EvaluateResponse, error)
}

func RegisterEvaluatorServer(s *grpc.Server, srv EvaluatorServer) {
	s.RegisterService(&_Evaluator_serviceDesc, srv)
}

func _Evaluator_Evaluate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EvaluateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EvaluatorServer).Evaluate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Evaluator/Evaluate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EvaluatorServer).Evaluate(ctx, req.(*EvaluateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Evaluator_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.Evaluator",
	HandlerType: (*EvaluatorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Evaluate",
			Handler:    _Evaluator_Evaluate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/evaluator.proto",
}
