// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/backend.proto

package pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	status "google.golang.org/genproto/googleapis/rpc/status"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status1 "google.golang.org/grpc/status"
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

type FunctionConfig_Type int32

const (
	FunctionConfig_GRPC FunctionConfig_Type = 0
	FunctionConfig_REST FunctionConfig_Type = 1
)

var FunctionConfig_Type_name = map[int32]string{
	0: "GRPC",
	1: "REST",
}

var FunctionConfig_Type_value = map[string]int32{
	"GRPC": 0,
	"REST": 1,
}

func (x FunctionConfig_Type) String() string {
	return proto.EnumName(FunctionConfig_Type_name, int32(x))
}

func (FunctionConfig_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8dab762378f455cd, []int{0, 0}
}

// FunctionConfig specifies a MMF address and client type for Backend to establish connections with the MMF
type FunctionConfig struct {
	Host                 string              `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	Port                 int32               `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	Type                 FunctionConfig_Type `protobuf:"varint,3,opt,name=type,proto3,enum=openmatch.FunctionConfig_Type" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *FunctionConfig) Reset()         { *m = FunctionConfig{} }
func (m *FunctionConfig) String() string { return proto.CompactTextString(m) }
func (*FunctionConfig) ProtoMessage()    {}
func (*FunctionConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_8dab762378f455cd, []int{0}
}

func (m *FunctionConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FunctionConfig.Unmarshal(m, b)
}
func (m *FunctionConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FunctionConfig.Marshal(b, m, deterministic)
}
func (m *FunctionConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FunctionConfig.Merge(m, src)
}
func (m *FunctionConfig) XXX_Size() int {
	return xxx_messageInfo_FunctionConfig.Size(m)
}
func (m *FunctionConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_FunctionConfig.DiscardUnknown(m)
}

var xxx_messageInfo_FunctionConfig proto.InternalMessageInfo

func (m *FunctionConfig) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *FunctionConfig) GetPort() int32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *FunctionConfig) GetType() FunctionConfig_Type {
	if m != nil {
		return m.Type
	}
	return FunctionConfig_GRPC
}

type FetchMatchesRequest struct {
	// FunctionConfig specifies a MMF address and client type for Backend to establish connections with the MMF
	Config *FunctionConfig `protobuf:"bytes,1,opt,name=config,proto3" json:"config,omitempty"`
	// MatchProfiles that will be sent to thhe MMF specified in the FunctionConfig.
	Profile              *MatchProfile `protobuf:"bytes,2,opt,name=profile,proto3" json:"profile,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *FetchMatchesRequest) Reset()         { *m = FetchMatchesRequest{} }
func (m *FetchMatchesRequest) String() string { return proto.CompactTextString(m) }
func (*FetchMatchesRequest) ProtoMessage()    {}
func (*FetchMatchesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8dab762378f455cd, []int{1}
}

func (m *FetchMatchesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FetchMatchesRequest.Unmarshal(m, b)
}
func (m *FetchMatchesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FetchMatchesRequest.Marshal(b, m, deterministic)
}
func (m *FetchMatchesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FetchMatchesRequest.Merge(m, src)
}
func (m *FetchMatchesRequest) XXX_Size() int {
	return xxx_messageInfo_FetchMatchesRequest.Size(m)
}
func (m *FetchMatchesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FetchMatchesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FetchMatchesRequest proto.InternalMessageInfo

func (m *FetchMatchesRequest) GetConfig() *FunctionConfig {
	if m != nil {
		return m.Config
	}
	return nil
}

func (m *FetchMatchesRequest) GetProfile() *MatchProfile {
	if m != nil {
		return m.Profile
	}
	return nil
}

type FetchMatchesResponse struct {
	// Fetch matches will return a stream of responses containing matches,
	// followed by a single FetchSummary before closing the response stream.
	//
	// Types that are valid to be assigned to Response:
	//	*FetchMatchesResponse_Match
	//	*FetchMatchesResponse_FetchMatchesSummary
	Response             isFetchMatchesResponse_Response `protobuf_oneof:"response"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *FetchMatchesResponse) Reset()         { *m = FetchMatchesResponse{} }
func (m *FetchMatchesResponse) String() string { return proto.CompactTextString(m) }
func (*FetchMatchesResponse) ProtoMessage()    {}
func (*FetchMatchesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8dab762378f455cd, []int{2}
}

func (m *FetchMatchesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FetchMatchesResponse.Unmarshal(m, b)
}
func (m *FetchMatchesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FetchMatchesResponse.Marshal(b, m, deterministic)
}
func (m *FetchMatchesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FetchMatchesResponse.Merge(m, src)
}
func (m *FetchMatchesResponse) XXX_Size() int {
	return xxx_messageInfo_FetchMatchesResponse.Size(m)
}
func (m *FetchMatchesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FetchMatchesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FetchMatchesResponse proto.InternalMessageInfo

type isFetchMatchesResponse_Response interface {
	isFetchMatchesResponse_Response()
}

type FetchMatchesResponse_Match struct {
	Match *Match `protobuf:"bytes,1,opt,name=match,proto3,oneof"`
}

type FetchMatchesResponse_FetchMatchesSummary struct {
	FetchMatchesSummary *FetchMatchesSummary `protobuf:"bytes,2,opt,name=fetch_matches_summary,json=fetchMatchesSummary,proto3,oneof"`
}

func (*FetchMatchesResponse_Match) isFetchMatchesResponse_Response() {}

func (*FetchMatchesResponse_FetchMatchesSummary) isFetchMatchesResponse_Response() {}

func (m *FetchMatchesResponse) GetResponse() isFetchMatchesResponse_Response {
	if m != nil {
		return m.Response
	}
	return nil
}

func (m *FetchMatchesResponse) GetMatch() *Match {
	if x, ok := m.GetResponse().(*FetchMatchesResponse_Match); ok {
		return x.Match
	}
	return nil
}

func (m *FetchMatchesResponse) GetFetchMatchesSummary() *FetchMatchesSummary {
	if x, ok := m.GetResponse().(*FetchMatchesResponse_FetchMatchesSummary); ok {
		return x.FetchMatchesSummary
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*FetchMatchesResponse) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*FetchMatchesResponse_Match)(nil),
		(*FetchMatchesResponse_FetchMatchesSummary)(nil),
	}
}

type FetchMatchesSummary struct {
	// Response status from running the matchfunction.  Will be present and OK in the event
	// of no errors.
	MmfStatus *status.Status `protobuf:"bytes,1,opt,name=mmf_status,json=mmfStatus,proto3" json:"mmf_status,omitempty"`
	// Response status from running the evaluator.  Will be present and OK in the event
	// of no errors.
	EvaluatorStatus *status.Status `protobuf:"bytes,2,opt,name=evaluator_status,json=evaluatorStatus,proto3" json:"evaluator_status,omitempty"`
	// Response status from internal Open Match system.  Will contain any errors
	// not originating from the mmf or evaluator.
	SystemStatus         *status.Status `protobuf:"bytes,3,opt,name=system_status,json=systemStatus,proto3" json:"system_status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *FetchMatchesSummary) Reset()         { *m = FetchMatchesSummary{} }
func (m *FetchMatchesSummary) String() string { return proto.CompactTextString(m) }
func (*FetchMatchesSummary) ProtoMessage()    {}
func (*FetchMatchesSummary) Descriptor() ([]byte, []int) {
	return fileDescriptor_8dab762378f455cd, []int{3}
}

func (m *FetchMatchesSummary) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FetchMatchesSummary.Unmarshal(m, b)
}
func (m *FetchMatchesSummary) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FetchMatchesSummary.Marshal(b, m, deterministic)
}
func (m *FetchMatchesSummary) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FetchMatchesSummary.Merge(m, src)
}
func (m *FetchMatchesSummary) XXX_Size() int {
	return xxx_messageInfo_FetchMatchesSummary.Size(m)
}
func (m *FetchMatchesSummary) XXX_DiscardUnknown() {
	xxx_messageInfo_FetchMatchesSummary.DiscardUnknown(m)
}

var xxx_messageInfo_FetchMatchesSummary proto.InternalMessageInfo

func (m *FetchMatchesSummary) GetMmfStatus() *status.Status {
	if m != nil {
		return m.MmfStatus
	}
	return nil
}

func (m *FetchMatchesSummary) GetEvaluatorStatus() *status.Status {
	if m != nil {
		return m.EvaluatorStatus
	}
	return nil
}

func (m *FetchMatchesSummary) GetSystemStatus() *status.Status {
	if m != nil {
		return m.SystemStatus
	}
	return nil
}

type AssignTicketsRequest struct {
	// TicketIds is a list of strings representing Open Match generated Ids which apply to an Assignment.
	TicketIds []string `protobuf:"bytes,1,rep,name=ticket_ids,json=ticketIds,proto3" json:"ticket_ids,omitempty"`
	// An Assignment specifies game connection related information to be associated with the TicketIds.
	Assignment           *Assignment `protobuf:"bytes,2,opt,name=assignment,proto3" json:"assignment,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *AssignTicketsRequest) Reset()         { *m = AssignTicketsRequest{} }
func (m *AssignTicketsRequest) String() string { return proto.CompactTextString(m) }
func (*AssignTicketsRequest) ProtoMessage()    {}
func (*AssignTicketsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8dab762378f455cd, []int{4}
}

func (m *AssignTicketsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AssignTicketsRequest.Unmarshal(m, b)
}
func (m *AssignTicketsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AssignTicketsRequest.Marshal(b, m, deterministic)
}
func (m *AssignTicketsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AssignTicketsRequest.Merge(m, src)
}
func (m *AssignTicketsRequest) XXX_Size() int {
	return xxx_messageInfo_AssignTicketsRequest.Size(m)
}
func (m *AssignTicketsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AssignTicketsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AssignTicketsRequest proto.InternalMessageInfo

func (m *AssignTicketsRequest) GetTicketIds() []string {
	if m != nil {
		return m.TicketIds
	}
	return nil
}

func (m *AssignTicketsRequest) GetAssignment() *Assignment {
	if m != nil {
		return m.Assignment
	}
	return nil
}

type AssignTicketsResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AssignTicketsResponse) Reset()         { *m = AssignTicketsResponse{} }
func (m *AssignTicketsResponse) String() string { return proto.CompactTextString(m) }
func (*AssignTicketsResponse) ProtoMessage()    {}
func (*AssignTicketsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8dab762378f455cd, []int{5}
}

func (m *AssignTicketsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AssignTicketsResponse.Unmarshal(m, b)
}
func (m *AssignTicketsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AssignTicketsResponse.Marshal(b, m, deterministic)
}
func (m *AssignTicketsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AssignTicketsResponse.Merge(m, src)
}
func (m *AssignTicketsResponse) XXX_Size() int {
	return xxx_messageInfo_AssignTicketsResponse.Size(m)
}
func (m *AssignTicketsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AssignTicketsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AssignTicketsResponse proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("openmatch.FunctionConfig_Type", FunctionConfig_Type_name, FunctionConfig_Type_value)
	proto.RegisterType((*FunctionConfig)(nil), "openmatch.FunctionConfig")
	proto.RegisterType((*FetchMatchesRequest)(nil), "openmatch.FetchMatchesRequest")
	proto.RegisterType((*FetchMatchesResponse)(nil), "openmatch.FetchMatchesResponse")
	proto.RegisterType((*FetchMatchesSummary)(nil), "openmatch.FetchMatchesSummary")
	proto.RegisterType((*AssignTicketsRequest)(nil), "openmatch.AssignTicketsRequest")
	proto.RegisterType((*AssignTicketsResponse)(nil), "openmatch.AssignTicketsResponse")
}

func init() { proto.RegisterFile("api/backend.proto", fileDescriptor_8dab762378f455cd) }

var fileDescriptor_8dab762378f455cd = []byte{
	// 824 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x54, 0x41, 0x6f, 0x1b, 0x45,
	0x14, 0xce, 0xae, 0xdd, 0xa4, 0x7e, 0x6d, 0x83, 0x99, 0x36, 0x24, 0x58, 0x40, 0xb7, 0x8b, 0x40,
	0x96, 0x69, 0x76, 0x62, 0x13, 0x84, 0x64, 0x54, 0xa9, 0x49, 0x48, 0x69, 0xa4, 0x02, 0xd5, 0x3a,
	0xe2, 0xc0, 0x25, 0x5a, 0x8f, 0x9f, 0xd7, 0x4b, 0xbc, 0x33, 0xc3, 0xcc, 0x6c, 0x4a, 0x84, 0x84,
	0x10, 0x27, 0xc4, 0x91, 0xde, 0x38, 0x72, 0xe4, 0xc6, 0x3f, 0xe0, 0x3f, 0x70, 0xe1, 0x07, 0xf0,
	0x43, 0xaa, 0x9d, 0x59, 0xa7, 0x76, 0x12, 0x9f, 0x3c, 0xfb, 0xbe, 0xef, 0x7d, 0xef, 0x9b, 0x6f,
	0x66, 0x0c, 0x6f, 0x26, 0x32, 0xa3, 0xc3, 0x84, 0x9d, 0x22, 0x1f, 0x45, 0x52, 0x09, 0x23, 0x48,
	0x43, 0x48, 0xe4, 0x79, 0x62, 0xd8, 0xa4, 0x45, 0x4a, 0x34, 0x47, 0xad, 0x93, 0x14, 0xb5, 0x83,
	0x5b, 0xef, 0xa4, 0x42, 0xa4, 0x53, 0xa4, 0x25, 0x94, 0x70, 0x2e, 0x4c, 0x62, 0x32, 0xc1, 0x67,
	0xe8, 0x66, 0x85, 0x2a, 0xc9, 0xa8, 0x36, 0x89, 0x29, 0x66, 0xc0, 0x43, 0xfb, 0xc3, 0xb6, 0x53,
	0xe4, 0xdb, 0xfa, 0x45, 0x92, 0xa6, 0xa8, 0xa8, 0x90, 0xb6, 0xf5, 0xaa, 0x4c, 0xf8, 0xab, 0x07,
	0xeb, 0x4f, 0x0a, 0xce, 0xca, 0xda, 0x81, 0xe0, 0xe3, 0x2c, 0x25, 0x04, 0xea, 0x13, 0xa1, 0xcd,
	0x96, 0x17, 0x78, 0xed, 0x46, 0x6c, 0xd7, 0x65, 0x4d, 0x0a, 0x65, 0xb6, 0xfc, 0xc0, 0x6b, 0xdf,
	0x88, 0xed, 0x9a, 0xf4, 0xa0, 0x6e, 0xce, 0x25, 0x6e, 0xd5, 0x02, 0xaf, 0xbd, 0xde, 0x7b, 0x2f,
	0xba, 0xd8, 0x4d, 0xb4, 0x28, 0x18, 0x1d, 0x9f, 0x4b, 0x8c, 0x2d, 0x37, 0x6c, 0x41, 0xbd, 0xfc,
	0x22, 0x37, 0xa1, 0xfe, 0x45, 0xfc, 0xfc, 0xa0, 0xb9, 0x52, 0xae, 0xe2, 0xc3, 0xc1, 0x71, 0xd3,
	0x0b, 0x7f, 0x84, 0xbb, 0x4f, 0xd0, 0xb0, 0xc9, 0x97, 0xa5, 0x06, 0xea, 0x18, 0xbf, 0x2f, 0x50,
	0x1b, 0xd2, 0x85, 0x55, 0x66, 0x75, 0xac, 0xa1, 0x5b, 0xbd, 0xb7, 0x97, 0x0e, 0x8a, 0x2b, 0x22,
	0xe9, 0xc2, 0x9a, 0x54, 0x62, 0x9c, 0x4d, 0xd1, 0x1a, 0xbe, 0xd5, 0xdb, 0x9c, 0xeb, 0xb1, 0xf2,
	0xcf, 0x1d, 0x1c, 0xcf, 0x78, 0xe1, 0x9f, 0x1e, 0xdc, 0x5b, 0x9c, 0xae, 0xa5, 0xe0, 0x1a, 0x49,
	0x1b, 0x6e, 0xd8, 0xbe, 0x6a, 0x7a, 0xf3, 0xb2, 0xd2, 0xd3, 0x95, 0xd8, 0x11, 0xc8, 0x31, 0x6c,
	0x8c, 0x4b, 0x85, 0x93, 0xdc, 0x49, 0x9c, 0xe8, 0x22, 0xcf, 0x13, 0x75, 0x5e, 0x79, 0x58, 0x08,
	0x68, 0x6e, 0xd2, 0xc0, 0xb1, 0x9e, 0xae, 0xc4, 0x77, 0xc7, 0x57, 0xcb, 0xfb, 0x00, 0x37, 0x55,
	0xe5, 0x25, 0xfc, 0xc7, 0x5b, 0x8c, 0xa8, 0xe2, 0x90, 0x2e, 0x40, 0x9e, 0x8f, 0x4f, 0xdc, 0x35,
	0xa8, 0x8c, 0x92, 0xc8, 0x5d, 0x90, 0x48, 0x49, 0x16, 0x0d, 0x2c, 0x12, 0x37, 0xf2, 0x7c, 0xec,
	0x96, 0xe4, 0x11, 0x34, 0xf1, 0x2c, 0x99, 0x16, 0x89, 0x11, 0x6a, 0xd6, 0xe8, 0x2f, 0x6d, 0x7c,
	0xe3, 0x82, 0x5b, 0xb5, 0x7f, 0x0a, 0x77, 0xf4, 0xb9, 0x36, 0x98, 0xcf, 0x7a, 0x6b, 0x4b, 0x7b,
	0x6f, 0x3b, 0xa2, 0xfb, 0x0a, 0xa7, 0x70, 0x6f, 0x4f, 0xeb, 0x2c, 0xe5, 0xc7, 0x19, 0x3b, 0x45,
	0x73, 0x71, 0xca, 0xef, 0x02, 0x18, 0x5b, 0x39, 0xc9, 0x46, 0xe5, 0x16, 0x6a, 0xed, 0x46, 0xdc,
	0x70, 0x95, 0xa3, 0x91, 0x26, 0x9f, 0x00, 0x24, 0xb6, 0x2d, 0x47, 0x6e, 0x2a, 0xa3, 0x1b, 0x73,
	0x81, 0xee, 0x5d, 0x80, 0xf1, 0x1c, 0x31, 0xdc, 0x84, 0x8d, 0x4b, 0xd3, 0x5c, 0x92, 0xbd, 0x97,
	0x3e, 0xac, 0xef, 0xbb, 0xc7, 0x38, 0x40, 0x75, 0x96, 0x31, 0x24, 0x3f, 0xc1, 0xed, 0xf9, 0x6c,
	0xc9, 0xb2, 0xf3, 0xaa, 0x1c, 0xb7, 0xee, 0x2f, 0xc5, 0xab, 0xd3, 0xfa, 0xe8, 0x97, 0x7f, 0xff,
	0x7f, 0xe9, 0x7f, 0x10, 0x06, 0xf4, 0xac, 0x3b, 0x7b, 0xf9, 0xda, 0x0d, 0xa3, 0xd5, 0x15, 0xe9,
	0xdb, 0x13, 0xef, 0x7b, 0x9d, 0x1d, 0x8f, 0xfc, 0xec, 0xc1, 0x9d, 0x05, 0xb3, 0xe4, 0xfe, 0x95,
	0x0d, 0x2e, 0x86, 0xd6, 0x0a, 0x96, 0x13, 0x2a, 0x0f, 0x0f, 0xad, 0x87, 0x0f, 0xc3, 0x07, 0xd7,
	0x78, 0x70, 0xe9, 0xea, 0xbe, 0xcb, 0xab, 0xef, 0x75, 0xf6, 0x7f, 0xab, 0xfd, 0xbe, 0xf7, 0x9f,
	0x4f, 0xfe, 0xf6, 0x60, 0xad, 0x0a, 0x27, 0x3c, 0x02, 0xf8, 0x5a, 0x22, 0x0f, 0xec, 0xe6, 0xc8,
	0x5b, 0x13, 0x63, 0xa4, 0xee, 0x53, 0x5a, 0xce, 0xdd, 0x76, 0x83, 0x47, 0x78, 0xd6, 0x7a, 0xff,
	0xf5, 0xf7, 0xf6, 0x28, 0xd3, 0xac, 0xd0, 0xfa, 0xb1, 0xbb, 0x09, 0xa9, 0x12, 0x85, 0xd4, 0x11,
	0x13, 0x79, 0xe7, 0x1b, 0x20, 0x7b, 0x32, 0x61, 0x13, 0x0c, 0x7a, 0xd1, 0x4e, 0xf0, 0x2c, 0x63,
	0x58, 0xbe, 0xaf, 0xc7, 0x33, 0xc9, 0x34, 0x33, 0x93, 0x62, 0x58, 0x32, 0xa9, 0x6b, 0x1d, 0x0b,
	0x95, 0x26, 0x39, 0xea, 0xb9, 0x61, 0x74, 0x38, 0x15, 0x43, 0x9a, 0x27, 0xda, 0xa0, 0xa2, 0xcf,
	0x8e, 0x0e, 0x0e, 0xbf, 0x1a, 0x1c, 0xf6, 0x6a, 0xdd, 0x68, 0xa7, 0xe3, 0x7b, 0x7e, 0xaf, 0x99,
	0x48, 0x39, 0xcd, 0x98, 0xfd, 0x7f, 0xa3, 0xdf, 0x69, 0xc1, 0xfb, 0x57, 0x2a, 0xf1, 0x67, 0x50,
	0xdb, 0xdd, 0xd9, 0x25, 0xbb, 0xd0, 0x89, 0xd1, 0x14, 0x8a, 0xe3, 0x28, 0x78, 0x31, 0x41, 0x1e,
	0x98, 0x09, 0x06, 0x0a, 0xb5, 0x28, 0x14, 0xc3, 0x60, 0x24, 0x50, 0x07, 0x5c, 0x98, 0x00, 0x7f,
	0xc8, 0xb4, 0x89, 0xc8, 0x2a, 0xd4, 0xff, 0xf0, 0xbd, 0x35, 0xf5, 0x08, 0xb6, 0x5e, 0x87, 0x11,
	0x7c, 0x2e, 0x58, 0x51, 0xde, 0x32, 0xab, 0x4e, 0x1e, 0x5c, 0x1f, 0x0d, 0xd5, 0x99, 0x41, 0x3a,
	0x12, 0x4c, 0xd3, 0x6f, 0x83, 0x4b, 0xd0, 0xdc, 0xbe, 0xe4, 0x69, 0x4a, 0xe5, 0xf0, 0x2f, 0xbf,
	0x51, 0xea, 0x5b, 0xf9, 0xe1, 0xaa, 0xfd, 0x83, 0xfe, 0xf8, 0x55, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x13, 0xba, 0x34, 0xbb, 0x39, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BackendServiceClient is the client API for BackendService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BackendServiceClient interface {
	// FetchMatches triggers a MatchFunction with the specified MatchProfile, which
	// returns a set of match proposals.  Those proposals are sent to the evaluator to
	// deduplicate tickets in matches.  Then matches are streamed back to the fetch
	// matches call.  The stream ends with the FetchSummary.
	//
	// WARNING: FetchMatches DOES NOT return an error if the match function or evaluator
	// return an error.  This is done so that if they return partial results, those results can be
	// processed.  Make sure to inspect the FetchSummary's status for both the match function
	// and evaluator to log and otherwise handle errors.
	FetchMatches(ctx context.Context, in *FetchMatchesRequest, opts ...grpc.CallOption) (BackendService_FetchMatchesClient, error)
	// AssignTickets overwrites the Assignment field of the input TicketIds.
	AssignTickets(ctx context.Context, in *AssignTicketsRequest, opts ...grpc.CallOption) (*AssignTicketsResponse, error)
}

type backendServiceClient struct {
	cc *grpc.ClientConn
}

func NewBackendServiceClient(cc *grpc.ClientConn) BackendServiceClient {
	return &backendServiceClient{cc}
}

func (c *backendServiceClient) FetchMatches(ctx context.Context, in *FetchMatchesRequest, opts ...grpc.CallOption) (BackendService_FetchMatchesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_BackendService_serviceDesc.Streams[0], "/openmatch.BackendService/FetchMatches", opts...)
	if err != nil {
		return nil, err
	}
	x := &backendServiceFetchMatchesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type BackendService_FetchMatchesClient interface {
	Recv() (*FetchMatchesResponse, error)
	grpc.ClientStream
}

type backendServiceFetchMatchesClient struct {
	grpc.ClientStream
}

func (x *backendServiceFetchMatchesClient) Recv() (*FetchMatchesResponse, error) {
	m := new(FetchMatchesResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *backendServiceClient) AssignTickets(ctx context.Context, in *AssignTicketsRequest, opts ...grpc.CallOption) (*AssignTicketsResponse, error) {
	out := new(AssignTicketsResponse)
	err := c.cc.Invoke(ctx, "/openmatch.BackendService/AssignTickets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BackendServiceServer is the server API for BackendService service.
type BackendServiceServer interface {
	// FetchMatches triggers a MatchFunction with the specified MatchProfile, which
	// returns a set of match proposals.  Those proposals are sent to the evaluator to
	// deduplicate tickets in matches.  Then matches are streamed back to the fetch
	// matches call.  The stream ends with the FetchSummary.
	//
	// WARNING: FetchMatches DOES NOT return an error if the match function or evaluator
	// return an error.  This is done so that if they return partial results, those results can be
	// processed.  Make sure to inspect the FetchSummary's status for both the match function
	// and evaluator to log and otherwise handle errors.
	FetchMatches(*FetchMatchesRequest, BackendService_FetchMatchesServer) error
	// AssignTickets overwrites the Assignment field of the input TicketIds.
	AssignTickets(context.Context, *AssignTicketsRequest) (*AssignTicketsResponse, error)
}

// UnimplementedBackendServiceServer can be embedded to have forward compatible implementations.
type UnimplementedBackendServiceServer struct {
}

func (*UnimplementedBackendServiceServer) FetchMatches(req *FetchMatchesRequest, srv BackendService_FetchMatchesServer) error {
	return status1.Errorf(codes.Unimplemented, "method FetchMatches not implemented")
}
func (*UnimplementedBackendServiceServer) AssignTickets(ctx context.Context, req *AssignTicketsRequest) (*AssignTicketsResponse, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method AssignTickets not implemented")
}

func RegisterBackendServiceServer(s *grpc.Server, srv BackendServiceServer) {
	s.RegisterService(&_BackendService_serviceDesc, srv)
}

func _BackendService_FetchMatches_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(FetchMatchesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BackendServiceServer).FetchMatches(m, &backendServiceFetchMatchesServer{stream})
}

type BackendService_FetchMatchesServer interface {
	Send(*FetchMatchesResponse) error
	grpc.ServerStream
}

type backendServiceFetchMatchesServer struct {
	grpc.ServerStream
}

func (x *backendServiceFetchMatchesServer) Send(m *FetchMatchesResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _BackendService_AssignTickets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AssignTicketsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackendServiceServer).AssignTickets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openmatch.BackendService/AssignTickets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackendServiceServer).AssignTickets(ctx, req.(*AssignTicketsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _BackendService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "openmatch.BackendService",
	HandlerType: (*BackendServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AssignTickets",
			Handler:    _BackendService_AssignTickets_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "FetchMatches",
			Handler:       _BackendService_FetchMatches_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "api/backend.proto",
}
