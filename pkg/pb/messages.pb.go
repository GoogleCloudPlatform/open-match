// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/messages.proto

package pb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	_struct "github.com/golang/protobuf/ptypes/struct"
	status "google.golang.org/genproto/googleapis/rpc/status"
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

// A Ticket is a basic matchmaking entity in Open Match. In order to enter
// matchmaking using Open Match, the client should generate a Ticket, passing in
// the properties to be associated with this Ticket. Open Match will generate an
// ID for a Ticket during creation. A Ticket could be used to represent an
// individual 'Player' or a 'Group' of players. Open Match will not interpret
// what the Ticket represents but just treat it as a matchmaking unit with a set
// of properties. Open Match stores the Ticket in state storage and enables an
// Assignment to be associated with this Ticket.
type Ticket struct {
	// The Ticket ID generated by Open Match.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Properties contains custom info about the ticket.  Top level values can be
	// used in indexing and filtering to find tickets.
	// TODO: Deprecate and remove this field.
	Properties *_struct.Struct `protobuf:"bytes,2,opt,name=properties,proto3" json:"properties,omitempty"`
	// Assignment associated with the Ticket.
	Assignment *Assignment `protobuf:"bytes,3,opt,name=assignment,proto3" json:"assignment,omitempty"`
	// Values visible to Open Match which can be used when querying for tickets
	// with specific values.
	SearchFields *SearchFields `protobuf:"bytes,4,opt,name=search_fields,json=searchFields,proto3" json:"search_fields,omitempty"`
	// Customized information to be used by the Match Making Function.  Optional,
	// depending on the requirements of the MMF.
	Extension            *any.Any `protobuf:"bytes,5,opt,name=extension,proto3" json:"extension,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ticket) Reset()         { *m = Ticket{} }
func (m *Ticket) String() string { return proto.CompactTextString(m) }
func (*Ticket) ProtoMessage()    {}
func (*Ticket) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb9fb1f207fd5b8c, []int{0}
}

func (m *Ticket) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ticket.Unmarshal(m, b)
}
func (m *Ticket) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ticket.Marshal(b, m, deterministic)
}
func (m *Ticket) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ticket.Merge(m, src)
}
func (m *Ticket) XXX_Size() int {
	return xxx_messageInfo_Ticket.Size(m)
}
func (m *Ticket) XXX_DiscardUnknown() {
	xxx_messageInfo_Ticket.DiscardUnknown(m)
}

var xxx_messageInfo_Ticket proto.InternalMessageInfo

func (m *Ticket) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Ticket) GetProperties() *_struct.Struct {
	if m != nil {
		return m.Properties
	}
	return nil
}

func (m *Ticket) GetAssignment() *Assignment {
	if m != nil {
		return m.Assignment
	}
	return nil
}

func (m *Ticket) GetSearchFields() *SearchFields {
	if m != nil {
		return m.SearchFields
	}
	return nil
}

func (m *Ticket) GetExtension() *any.Any {
	if m != nil {
		return m.Extension
	}
	return nil
}

// Search fields are the fields which Open Match is aware of, and can be used
// when specifying filters.
type SearchFields struct {
	// Float arguments.  Filterable on ranges.
	DoubleArgs map[string]float64 `protobuf:"bytes,1,rep,name=double_args,json=doubleArgs,proto3" json:"double_args,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"fixed64,2,opt,name=value,proto3"`
	// String arguments.  Filterable on equality.
	StringArgs map[string]string `protobuf:"bytes,2,rep,name=string_args,json=stringArgs,proto3" json:"string_args,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Filterable on presence or absence of given value.
	Tags                 []string `protobuf:"bytes,3,rep,name=tags,proto3" json:"tags,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchFields) Reset()         { *m = SearchFields{} }
func (m *SearchFields) String() string { return proto.CompactTextString(m) }
func (*SearchFields) ProtoMessage()    {}
func (*SearchFields) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb9fb1f207fd5b8c, []int{1}
}

func (m *SearchFields) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchFields.Unmarshal(m, b)
}
func (m *SearchFields) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchFields.Marshal(b, m, deterministic)
}
func (m *SearchFields) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchFields.Merge(m, src)
}
func (m *SearchFields) XXX_Size() int {
	return xxx_messageInfo_SearchFields.Size(m)
}
func (m *SearchFields) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchFields.DiscardUnknown(m)
}

var xxx_messageInfo_SearchFields proto.InternalMessageInfo

func (m *SearchFields) GetDoubleArgs() map[string]float64 {
	if m != nil {
		return m.DoubleArgs
	}
	return nil
}

func (m *SearchFields) GetStringArgs() map[string]string {
	if m != nil {
		return m.StringArgs
	}
	return nil
}

func (m *SearchFields) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

// An Assignment object represents the assignment associated with a Ticket. Open
// match does not require or inspect any fields on assignment.
type Assignment struct {
	// Connection information for this Assignment.
	Connection string `protobuf:"bytes,1,opt,name=connection,proto3" json:"connection,omitempty"`
	// Other details to be sent to the players.
	// TODO: Deprecate and remove this field
	Properties *_struct.Struct `protobuf:"bytes,2,opt,name=properties,proto3" json:"properties,omitempty"`
	// Error when finding an Assignment for this Ticket.
	Error *status.Status `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
	// Customized information to be sent to the clients.  Optional, depending on
	// what callers are expecting.
	Extension            *any.Any `protobuf:"bytes,4,opt,name=extension,proto3" json:"extension,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Assignment) Reset()         { *m = Assignment{} }
func (m *Assignment) String() string { return proto.CompactTextString(m) }
func (*Assignment) ProtoMessage()    {}
func (*Assignment) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb9fb1f207fd5b8c, []int{2}
}

func (m *Assignment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Assignment.Unmarshal(m, b)
}
func (m *Assignment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Assignment.Marshal(b, m, deterministic)
}
func (m *Assignment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Assignment.Merge(m, src)
}
func (m *Assignment) XXX_Size() int {
	return xxx_messageInfo_Assignment.Size(m)
}
func (m *Assignment) XXX_DiscardUnknown() {
	xxx_messageInfo_Assignment.DiscardUnknown(m)
}

var xxx_messageInfo_Assignment proto.InternalMessageInfo

func (m *Assignment) GetConnection() string {
	if m != nil {
		return m.Connection
	}
	return ""
}

func (m *Assignment) GetProperties() *_struct.Struct {
	if m != nil {
		return m.Properties
	}
	return nil
}

func (m *Assignment) GetError() *status.Status {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *Assignment) GetExtension() *any.Any {
	if m != nil {
		return m.Extension
	}
	return nil
}

// Filters numerical values to only those within a range.
//   attribute: "foo"
//   max: 10
//   min: 5
// matches:
//   {"foo": 5}
//   {"foo": 7.5}
//   {"foo": 10}
// does not match:
//   {"foo": 4}
//   {"foo": 10.01}
//   {"foo": "7.5"}
//   {}
type FloatRangeFilter struct {
	// Name of the ticket attribute this Filter operates on.
	// TODO: rename double_args
	Attribute string `protobuf:"bytes,1,opt,name=attribute,proto3" json:"attribute,omitempty"`
	// Maximum value. Defaults to positive infinity (any value above minv).
	Max float64 `protobuf:"fixed64,2,opt,name=max,proto3" json:"max,omitempty"`
	// Minimum value. Defaults to 0.
	Min                  float64  `protobuf:"fixed64,3,opt,name=min,proto3" json:"min,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FloatRangeFilter) Reset()         { *m = FloatRangeFilter{} }
func (m *FloatRangeFilter) String() string { return proto.CompactTextString(m) }
func (*FloatRangeFilter) ProtoMessage()    {}
func (*FloatRangeFilter) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb9fb1f207fd5b8c, []int{3}
}

func (m *FloatRangeFilter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FloatRangeFilter.Unmarshal(m, b)
}
func (m *FloatRangeFilter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FloatRangeFilter.Marshal(b, m, deterministic)
}
func (m *FloatRangeFilter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FloatRangeFilter.Merge(m, src)
}
func (m *FloatRangeFilter) XXX_Size() int {
	return xxx_messageInfo_FloatRangeFilter.Size(m)
}
func (m *FloatRangeFilter) XXX_DiscardUnknown() {
	xxx_messageInfo_FloatRangeFilter.DiscardUnknown(m)
}

var xxx_messageInfo_FloatRangeFilter proto.InternalMessageInfo

func (m *FloatRangeFilter) GetAttribute() string {
	if m != nil {
		return m.Attribute
	}
	return ""
}

func (m *FloatRangeFilter) GetMax() float64 {
	if m != nil {
		return m.Max
	}
	return 0
}

func (m *FloatRangeFilter) GetMin() float64 {
	if m != nil {
		return m.Min
	}
	return 0
}

// Filters strings exactly equaling a value.
//   attribute: "foo"
//   value: "bar"
// matches:
//   {"foo": "bar"}
// does not match:
//   {"foo": "baz"}
//   {"bar": "foo"}
//   {}
type StringEqualsFilter struct {
	// TODO: rename string_args
	Attribute            string   `protobuf:"bytes,1,opt,name=attribute,proto3" json:"attribute,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StringEqualsFilter) Reset()         { *m = StringEqualsFilter{} }
func (m *StringEqualsFilter) String() string { return proto.CompactTextString(m) }
func (*StringEqualsFilter) ProtoMessage()    {}
func (*StringEqualsFilter) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb9fb1f207fd5b8c, []int{4}
}

func (m *StringEqualsFilter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StringEqualsFilter.Unmarshal(m, b)
}
func (m *StringEqualsFilter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StringEqualsFilter.Marshal(b, m, deterministic)
}
func (m *StringEqualsFilter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StringEqualsFilter.Merge(m, src)
}
func (m *StringEqualsFilter) XXX_Size() int {
	return xxx_messageInfo_StringEqualsFilter.Size(m)
}
func (m *StringEqualsFilter) XXX_DiscardUnknown() {
	xxx_messageInfo_StringEqualsFilter.DiscardUnknown(m)
}

var xxx_messageInfo_StringEqualsFilter proto.InternalMessageInfo

func (m *StringEqualsFilter) GetAttribute() string {
	if m != nil {
		return m.Attribute
	}
	return ""
}

func (m *StringEqualsFilter) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

// Filters to the tag being present on the search_fields.
//   tag: "foo"
// matches:
//   ["foo"]
//   ["bar","foo"]
// does not match:
//   ["bar"]
//   []
type TagPresentFilter struct {
	Tag                  string   `protobuf:"bytes,1,opt,name=tag,proto3" json:"tag,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TagPresentFilter) Reset()         { *m = TagPresentFilter{} }
func (m *TagPresentFilter) String() string { return proto.CompactTextString(m) }
func (*TagPresentFilter) ProtoMessage()    {}
func (*TagPresentFilter) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb9fb1f207fd5b8c, []int{5}
}

func (m *TagPresentFilter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TagPresentFilter.Unmarshal(m, b)
}
func (m *TagPresentFilter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TagPresentFilter.Marshal(b, m, deterministic)
}
func (m *TagPresentFilter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TagPresentFilter.Merge(m, src)
}
func (m *TagPresentFilter) XXX_Size() int {
	return xxx_messageInfo_TagPresentFilter.Size(m)
}
func (m *TagPresentFilter) XXX_DiscardUnknown() {
	xxx_messageInfo_TagPresentFilter.DiscardUnknown(m)
}

var xxx_messageInfo_TagPresentFilter proto.InternalMessageInfo

func (m *TagPresentFilter) GetTag() string {
	if m != nil {
		return m.Tag
	}
	return ""
}

type Pool struct {
	// A developer-chosen human-readable name for this Pool.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Set of Filters indicating the filtering criteria. Selected players must
	// match every Filter.
	FloatRangeFilters    []*FloatRangeFilter   `protobuf:"bytes,2,rep,name=float_range_filters,json=floatRangeFilters,proto3" json:"float_range_filters,omitempty"`
	StringEqualsFilters  []*StringEqualsFilter `protobuf:"bytes,4,rep,name=string_equals_filters,json=stringEqualsFilters,proto3" json:"string_equals_filters,omitempty"`
	TagPresentFilters    []*TagPresentFilter   `protobuf:"bytes,5,rep,name=tag_present_filters,json=tagPresentFilters,proto3" json:"tag_present_filters,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Pool) Reset()         { *m = Pool{} }
func (m *Pool) String() string { return proto.CompactTextString(m) }
func (*Pool) ProtoMessage()    {}
func (*Pool) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb9fb1f207fd5b8c, []int{6}
}

func (m *Pool) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Pool.Unmarshal(m, b)
}
func (m *Pool) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Pool.Marshal(b, m, deterministic)
}
func (m *Pool) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pool.Merge(m, src)
}
func (m *Pool) XXX_Size() int {
	return xxx_messageInfo_Pool.Size(m)
}
func (m *Pool) XXX_DiscardUnknown() {
	xxx_messageInfo_Pool.DiscardUnknown(m)
}

var xxx_messageInfo_Pool proto.InternalMessageInfo

func (m *Pool) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Pool) GetFloatRangeFilters() []*FloatRangeFilter {
	if m != nil {
		return m.FloatRangeFilters
	}
	return nil
}

func (m *Pool) GetStringEqualsFilters() []*StringEqualsFilter {
	if m != nil {
		return m.StringEqualsFilters
	}
	return nil
}

func (m *Pool) GetTagPresentFilters() []*TagPresentFilter {
	if m != nil {
		return m.TagPresentFilters
	}
	return nil
}

// A Roster is a named collection of Ticket IDs. It exists so that a Tickets
// associated with a Match can be labelled to belong to a team, sub-team etc. It
// can also be used to represent the current state of a Match in scenarios such
// as backfill, join-in-progress etc.
type Roster struct {
	// A developer-chosen human-readable name for this Roster.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Tickets belonging to this Roster.
	TicketIds            []string `protobuf:"bytes,2,rep,name=ticket_ids,json=ticketIds,proto3" json:"ticket_ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Roster) Reset()         { *m = Roster{} }
func (m *Roster) String() string { return proto.CompactTextString(m) }
func (*Roster) ProtoMessage()    {}
func (*Roster) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb9fb1f207fd5b8c, []int{7}
}

func (m *Roster) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Roster.Unmarshal(m, b)
}
func (m *Roster) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Roster.Marshal(b, m, deterministic)
}
func (m *Roster) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Roster.Merge(m, src)
}
func (m *Roster) XXX_Size() int {
	return xxx_messageInfo_Roster.Size(m)
}
func (m *Roster) XXX_DiscardUnknown() {
	xxx_messageInfo_Roster.DiscardUnknown(m)
}

var xxx_messageInfo_Roster proto.InternalMessageInfo

func (m *Roster) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Roster) GetTicketIds() []string {
	if m != nil {
		return m.TicketIds
	}
	return nil
}

// A MatchProfile is Open Match's representation of a Match specification. It is
// used to indicate the criteria for selecting players for a match. A
// MatchProfile is the input to the API to get matches and is passed to the
// MatchFunction. It contains all the information required by the MatchFunction
// to generate match proposals.
type MatchProfile struct {
	// Name of this match profile.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Set of properties associated with this MatchProfile. (Optional)
	// Open Match does not interpret these properties but passes them through to
	// the MatchFunction.
	// TODO: Deprecate and remove.
	Properties *_struct.Struct `protobuf:"bytes,2,opt,name=properties,proto3" json:"properties,omitempty"`
	// Set of pools to be queried when generating a match for this MatchProfile.
	// The pool names can be used in empty Rosters to specify composition of a
	// match.
	Pools []*Pool `protobuf:"bytes,3,rep,name=pools,proto3" json:"pools,omitempty"`
	// Set of Rosters for this match request. Could be empty Rosters used to
	// indicate the composition of the generated Match or they could be partially
	// pre-populated Ticket list to be used in scenarios such as backfill / join
	// in progress.
	Rosters []*Roster `protobuf:"bytes,4,rep,name=rosters,proto3" json:"rosters,omitempty"`
	// Customized information on how the match function should run.  Optional,
	// depending on the requirements of the match function.
	Extension            *any.Any `protobuf:"bytes,5,opt,name=extension,proto3" json:"extension,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MatchProfile) Reset()         { *m = MatchProfile{} }
func (m *MatchProfile) String() string { return proto.CompactTextString(m) }
func (*MatchProfile) ProtoMessage()    {}
func (*MatchProfile) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb9fb1f207fd5b8c, []int{8}
}

func (m *MatchProfile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MatchProfile.Unmarshal(m, b)
}
func (m *MatchProfile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MatchProfile.Marshal(b, m, deterministic)
}
func (m *MatchProfile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MatchProfile.Merge(m, src)
}
func (m *MatchProfile) XXX_Size() int {
	return xxx_messageInfo_MatchProfile.Size(m)
}
func (m *MatchProfile) XXX_DiscardUnknown() {
	xxx_messageInfo_MatchProfile.DiscardUnknown(m)
}

var xxx_messageInfo_MatchProfile proto.InternalMessageInfo

func (m *MatchProfile) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *MatchProfile) GetProperties() *_struct.Struct {
	if m != nil {
		return m.Properties
	}
	return nil
}

func (m *MatchProfile) GetPools() []*Pool {
	if m != nil {
		return m.Pools
	}
	return nil
}

func (m *MatchProfile) GetRosters() []*Roster {
	if m != nil {
		return m.Rosters
	}
	return nil
}

func (m *MatchProfile) GetExtension() *any.Any {
	if m != nil {
		return m.Extension
	}
	return nil
}

// A Match is used to represent a completed match object. It can be generated by
// a MatchFunction as a proposal or can be returned by OpenMatch as a result in
// response to the FetchMatches call.
// When a match is returned by the FetchMatches call, it should contain at least
// one ticket to be considered as valid.
type Match struct {
	// A Match ID that should be passed through the stack for tracing.
	MatchId string `protobuf:"bytes,1,opt,name=match_id,json=matchId,proto3" json:"match_id,omitempty"`
	// Name of the match profile that generated this Match.
	MatchProfile string `protobuf:"bytes,2,opt,name=match_profile,json=matchProfile,proto3" json:"match_profile,omitempty"`
	// Name of the match function that generated this Match.
	MatchFunction string `protobuf:"bytes,3,opt,name=match_function,json=matchFunction,proto3" json:"match_function,omitempty"`
	// Tickets belonging to this match.
	Tickets []*Ticket `protobuf:"bytes,4,rep,name=tickets,proto3" json:"tickets,omitempty"`
	// Set of Rosters that comprise this Match
	Rosters []*Roster `protobuf:"bytes,5,rep,name=rosters,proto3" json:"rosters,omitempty"`
	// Match properties for this Match. Open Match does not interpret this field.
	// TODO: Deprecate and remove.
	Properties *_struct.Struct `protobuf:"bytes,6,opt,name=properties,proto3" json:"properties,omitempty"`
	// Customized information for the evaluator.  Optional, depending on the
	// requirements of the configured evaluator.
	EvaluationInput *any.Any `protobuf:"bytes,7,opt,name=evaluation_input,json=evaluationInput,proto3" json:"evaluation_input,omitempty"`
	// Customized information for how the caller of FetchMatches should handle
	// this match.  Optional, depending on the requirements of the FetchMatches
	// caller.
	Extension            *any.Any `protobuf:"bytes,8,opt,name=extension,proto3" json:"extension,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Match) Reset()         { *m = Match{} }
func (m *Match) String() string { return proto.CompactTextString(m) }
func (*Match) ProtoMessage()    {}
func (*Match) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb9fb1f207fd5b8c, []int{9}
}

func (m *Match) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Match.Unmarshal(m, b)
}
func (m *Match) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Match.Marshal(b, m, deterministic)
}
func (m *Match) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Match.Merge(m, src)
}
func (m *Match) XXX_Size() int {
	return xxx_messageInfo_Match.Size(m)
}
func (m *Match) XXX_DiscardUnknown() {
	xxx_messageInfo_Match.DiscardUnknown(m)
}

var xxx_messageInfo_Match proto.InternalMessageInfo

func (m *Match) GetMatchId() string {
	if m != nil {
		return m.MatchId
	}
	return ""
}

func (m *Match) GetMatchProfile() string {
	if m != nil {
		return m.MatchProfile
	}
	return ""
}

func (m *Match) GetMatchFunction() string {
	if m != nil {
		return m.MatchFunction
	}
	return ""
}

func (m *Match) GetTickets() []*Ticket {
	if m != nil {
		return m.Tickets
	}
	return nil
}

func (m *Match) GetRosters() []*Roster {
	if m != nil {
		return m.Rosters
	}
	return nil
}

func (m *Match) GetProperties() *_struct.Struct {
	if m != nil {
		return m.Properties
	}
	return nil
}

func (m *Match) GetEvaluationInput() *any.Any {
	if m != nil {
		return m.EvaluationInput
	}
	return nil
}

func (m *Match) GetExtension() *any.Any {
	if m != nil {
		return m.Extension
	}
	return nil
}

func init() {
	proto.RegisterType((*Ticket)(nil), "openmatch.Ticket")
	proto.RegisterType((*SearchFields)(nil), "openmatch.SearchFields")
	proto.RegisterMapType((map[string]float64)(nil), "openmatch.SearchFields.DoubleArgsEntry")
	proto.RegisterMapType((map[string]string)(nil), "openmatch.SearchFields.StringArgsEntry")
	proto.RegisterType((*Assignment)(nil), "openmatch.Assignment")
	proto.RegisterType((*FloatRangeFilter)(nil), "openmatch.FloatRangeFilter")
	proto.RegisterType((*StringEqualsFilter)(nil), "openmatch.StringEqualsFilter")
	proto.RegisterType((*TagPresentFilter)(nil), "openmatch.TagPresentFilter")
	proto.RegisterType((*Pool)(nil), "openmatch.Pool")
	proto.RegisterType((*Roster)(nil), "openmatch.Roster")
	proto.RegisterType((*MatchProfile)(nil), "openmatch.MatchProfile")
	proto.RegisterType((*Match)(nil), "openmatch.Match")
}

func init() { proto.RegisterFile("api/messages.proto", fileDescriptor_cb9fb1f207fd5b8c) }

var fileDescriptor_cb9fb1f207fd5b8c = []byte{
	// 809 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x55, 0xdd, 0x6e, 0xdb, 0x36,
	0x14, 0x86, 0x24, 0x3b, 0x89, 0x4e, 0xd2, 0xc6, 0x61, 0x5b, 0x54, 0xcd, 0xda, 0xc1, 0xd0, 0x56,
	0xcc, 0xc0, 0x30, 0x19, 0xc8, 0x30, 0x6c, 0xd8, 0x0f, 0x86, 0x0c, 0x6b, 0xd0, 0x6c, 0x18, 0x96,
	0xd1, 0xb9, 0xda, 0x8d, 0x40, 0x4b, 0xb4, 0x4a, 0x44, 0x26, 0x35, 0x92, 0x2a, 0xea, 0x87, 0xd8,
	0x8b, 0xec, 0x7a, 0x6f, 0xb0, 0x47, 0xd9, 0x3b, 0xec, 0x62, 0x57, 0x03, 0x49, 0xc9, 0x56, 0x6c,
	0xaf, 0x35, 0x72, 0x47, 0x9e, 0xf3, 0x9d, 0x8f, 0xfc, 0xbe, 0x73, 0x44, 0x01, 0x22, 0x15, 0x1b,
	0xcf, 0xa9, 0x52, 0xa4, 0xa0, 0x2a, 0xa9, 0xa4, 0xd0, 0x02, 0x85, 0xa2, 0xa2, 0x7c, 0x4e, 0x74,
	0xf6, 0xea, 0xf4, 0x71, 0x21, 0x44, 0x51, 0xd2, 0xb1, 0xac, 0xb2, 0xb1, 0xd2, 0x44, 0xd7, 0x0d,
	0xe6, 0xf4, 0x69, 0x93, 0xb0, 0xbb, 0x69, 0x3d, 0x1b, 0x2b, 0x2d, 0xeb, 0x4c, 0x37, 0xd9, 0x27,
	0xeb, 0x59, 0xc2, 0x17, 0x2e, 0x15, 0xff, 0xeb, 0xc1, 0xde, 0x35, 0xcb, 0x6e, 0xa8, 0x46, 0xf7,
	0xc1, 0x67, 0x79, 0xe4, 0x0d, 0xbd, 0x51, 0x88, 0x7d, 0x96, 0xa3, 0xcf, 0x01, 0x2a, 0x29, 0x2a,
	0x2a, 0x35, 0xa3, 0x2a, 0xf2, 0x87, 0xde, 0xe8, 0xf0, 0xec, 0x71, 0xe2, 0xa8, 0x92, 0x96, 0x2a,
	0x99, 0xd8, 0x83, 0x70, 0x07, 0x8a, 0x3e, 0x03, 0x20, 0x4a, 0xb1, 0x82, 0xcf, 0x29, 0xd7, 0x51,
	0x60, 0x0b, 0x1f, 0x25, 0x4b, 0x15, 0xc9, 0xf9, 0x32, 0x89, 0x3b, 0x40, 0xf4, 0x35, 0xdc, 0x53,
	0x94, 0xc8, 0xec, 0x55, 0x3a, 0x63, 0xb4, 0xcc, 0x55, 0xd4, 0x6b, 0x8e, 0x5c, 0x55, 0x4e, 0x6c,
	0xfe, 0xc2, 0xa6, 0xf1, 0x91, 0xea, 0xec, 0xd0, 0x19, 0x84, 0xf4, 0x8d, 0xa6, 0x5c, 0x31, 0xc1,
	0xa3, 0xbe, 0xad, 0x7c, 0xb8, 0x71, 0xd9, 0x73, 0xbe, 0xc0, 0x2b, 0x58, 0xfc, 0xa7, 0x0f, 0x47,
	0x5d, 0x4a, 0xf4, 0x12, 0x0e, 0x73, 0x51, 0x4f, 0x4b, 0x9a, 0x12, 0x59, 0xa8, 0xc8, 0x1b, 0x06,
	0xa3, 0xc3, 0xb3, 0x8f, 0xfe, 0xe7, 0x02, 0xc9, 0xf7, 0x16, 0x7a, 0x2e, 0x0b, 0xf5, 0x82, 0x6b,
	0xb9, 0xc0, 0x90, 0x2f, 0x03, 0x86, 0x49, 0x69, 0xc9, 0x78, 0xe1, 0x98, 0xfc, 0xb7, 0x33, 0x4d,
	0x2c, 0xb4, 0xc3, 0xa4, 0x96, 0x01, 0x84, 0xa0, 0xa7, 0x49, 0xa1, 0xa2, 0x60, 0x18, 0x8c, 0x42,
	0x6c, 0xd7, 0xa7, 0xdf, 0xc0, 0xf1, 0xda, 0xe1, 0x68, 0x00, 0xc1, 0x0d, 0x5d, 0x34, 0xed, 0x33,
	0x4b, 0xf4, 0x10, 0xfa, 0xaf, 0x49, 0x59, 0x53, 0xdb, 0x3a, 0x0f, 0xbb, 0xcd, 0x97, 0xfe, 0x17,
	0x9e, 0x29, 0x5f, 0x3b, 0xf1, 0x5d, 0xe5, 0x61, 0xa7, 0x3c, 0xfe, 0xcb, 0x03, 0x58, 0xf5, 0x10,
	0xbd, 0x0f, 0x90, 0x09, 0xce, 0x69, 0xa6, 0x8d, 0xf5, 0x8e, 0xa1, 0x13, 0xb9, 0xfb, 0x1c, 0x8d,
	0xa0, 0x4f, 0xa5, 0x14, 0xb2, 0x19, 0x21, 0xd4, 0xd6, 0xc8, 0x2a, 0x4b, 0x26, 0x76, 0xfa, 0xb1,
	0x03, 0xdc, 0x6e, 0x7e, 0x6f, 0xb7, 0xe6, 0x5f, 0xc3, 0xe0, 0xa2, 0x14, 0x44, 0x63, 0xc2, 0x0b,
	0x7a, 0xc1, 0x4a, 0x4d, 0x25, 0x7a, 0x0a, 0x21, 0xd1, 0x5a, 0xb2, 0x69, 0xad, 0x69, 0xa3, 0x64,
	0x15, 0x30, 0x1e, 0xcd, 0xc9, 0x9b, 0xc6, 0x4e, 0xb3, 0xb4, 0x11, 0xc6, 0xed, 0xfd, 0x4c, 0x84,
	0xf1, 0xf8, 0x25, 0x20, 0x67, 0xed, 0x8b, 0xdf, 0x6a, 0x52, 0xaa, 0x9d, 0x78, 0xb7, 0x3a, 0x1d,
	0x7f, 0x08, 0x83, 0x6b, 0x52, 0x5c, 0x49, 0xaa, 0x28, 0xd7, 0x0d, 0xcf, 0x00, 0x02, 0x4d, 0x8a,
	0xb6, 0x4b, 0x9a, 0x14, 0xf1, 0xef, 0x3e, 0xf4, 0xae, 0x84, 0x28, 0xcd, 0x98, 0x70, 0x32, 0x6f,
	0xd9, 0xed, 0x1a, 0xfd, 0x08, 0x0f, 0x66, 0x46, 0x62, 0x2a, 0x8d, 0xc6, 0x74, 0x66, 0x49, 0xda,
	0x61, 0x7c, 0xaf, 0x33, 0x8c, 0xeb, 0x46, 0xe0, 0x93, 0xd9, 0x5a, 0x44, 0xa1, 0x5f, 0xe0, 0x51,
	0x33, 0xd1, 0xd4, 0x4a, 0x5b, 0xd2, 0xf5, 0x2c, 0xdd, 0xb3, 0xee, 0x6c, 0x6f, 0x38, 0x80, 0x1f,
	0xa8, 0x8d, 0x98, 0x32, 0xf7, 0xd3, 0xa4, 0x48, 0x2b, 0xa7, 0x71, 0x49, 0xd8, 0xdf, 0xb8, 0xdf,
	0xba, 0x11, 0xf8, 0x44, 0xaf, 0x45, 0xd4, 0x0f, 0xbd, 0x83, 0x60, 0xd0, 0x8b, 0xbf, 0x82, 0x3d,
	0x2c, 0x94, 0xf1, 0x6a, 0x9b, 0x21, 0xcf, 0x00, 0xb4, 0x7d, 0xec, 0x52, 0x96, 0x3b, 0x1f, 0x42,
	0x1c, 0xba, 0xc8, 0x65, 0xae, 0xe2, 0xbf, 0x3d, 0x38, 0xfa, 0xc9, 0x1c, 0x78, 0x25, 0xc5, 0x8c,
	0x95, 0x74, 0x2b, 0xc7, 0x9d, 0xc7, 0xf9, 0x39, 0xf4, 0x2b, 0x21, 0x4a, 0xf7, 0x25, 0x1f, 0x9e,
	0x1d, 0x77, 0xf4, 0x99, 0x0e, 0x62, 0x97, 0x45, 0x1f, 0xc3, 0xbe, 0xb4, 0x0a, 0x5a, 0x67, 0x4f,
	0x3a, 0x40, 0xa7, 0x0d, 0xb7, 0x88, 0x3b, 0xbd, 0x7a, 0xff, 0xf8, 0xd0, 0xb7, 0x2a, 0xd1, 0x13,
	0x38, 0xb0, 0xb4, 0xe9, 0xf2, 0xdd, 0xdf, 0xb7, 0xfb, 0xcb, 0x1c, 0x7d, 0x00, 0xf7, 0x5c, 0xaa,
	0x72, 0x56, 0x34, 0xb3, 0x79, 0x34, 0xef, 0xda, 0xf3, 0x1c, 0xee, 0x3b, 0xd0, 0xac, 0xe6, 0xee,
	0xeb, 0x0f, 0x2c, 0xca, 0x95, 0x5e, 0x34, 0x41, 0xa3, 0xc8, 0x79, 0xbc, 0x4d, 0x91, 0xfb, 0xf9,
	0xe0, 0x16, 0xd1, 0x95, 0xdf, 0x7f, 0xa7, 0xfc, 0xdb, 0xbd, 0xd8, 0xdb, 0xbd, 0x17, 0xdf, 0xc2,
	0x80, 0x9a, 0xcf, 0x8c, 0x98, 0x0b, 0xa6, 0x8c, 0x57, 0xb5, 0x8e, 0xf6, 0xdf, 0x62, 0xdf, 0xf1,
	0x0a, 0x7d, 0x69, 0xc0, 0xb7, 0x8d, 0x3f, 0xd8, 0xc9, 0xf8, 0xef, 0x92, 0x5f, 0x87, 0x46, 0xca,
	0x27, 0x4e, 0x4b, 0x4e, 0x5f, 0x8f, 0x57, 0xdb, 0x71, 0x75, 0x53, 0x8c, 0xab, 0xe9, 0x1f, 0x7e,
	0xf8, 0x73, 0x45, 0xb9, 0x6d, 0xcf, 0x74, 0xcf, 0x12, 0x7d, 0xfa, 0x5f, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x88, 0xd3, 0x4d, 0x7f, 0x15, 0x08, 0x00, 0x00,
}
