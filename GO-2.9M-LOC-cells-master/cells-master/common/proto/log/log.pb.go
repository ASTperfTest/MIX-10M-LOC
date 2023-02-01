// Code generated by protoc-gen-go. DO NOT EDIT.
// source: log.proto

/*
Package log is a generated protocol buffer package.

It is generated from these files:
	log.proto

It has these top-level messages:
	RecorderPutResponse
	Log
	LogMessage
	ListLogRequest
	ListLogResponse
	DeleteLogsResponse
	TimeRangeResponse
	TimeRangeResult
	TimeRangeRequest
	TimeRangeCursor
*/
package log

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Relative links types.
// Note that First is time.Now() and last time.Unix(0).
// We added an unused NONE enum with value 0 to workaround 0 issues between JSON and proto3.
type RelType int32

const (
	RelType_NONE  RelType = 0
	RelType_FIRST RelType = 1
	RelType_PREV  RelType = 2
	RelType_NEXT  RelType = 3
	RelType_LAST  RelType = 4
)

var RelType_name = map[int32]string{
	0: "NONE",
	1: "FIRST",
	2: "PREV",
	3: "NEXT",
	4: "LAST",
}
var RelType_value = map[string]int32{
	"NONE":  0,
	"FIRST": 1,
	"PREV":  2,
	"NEXT":  3,
	"LAST":  4,
}

func (x RelType) String() string {
	return proto.EnumName(RelType_name, int32(x))
}
func (RelType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// Output Format
type ListLogRequest_LogFormat int32

const (
	ListLogRequest_JSON ListLogRequest_LogFormat = 0
	ListLogRequest_CSV  ListLogRequest_LogFormat = 1
	ListLogRequest_XLSX ListLogRequest_LogFormat = 2
)

var ListLogRequest_LogFormat_name = map[int32]string{
	0: "JSON",
	1: "CSV",
	2: "XLSX",
}
var ListLogRequest_LogFormat_value = map[string]int32{
	"JSON": 0,
	"CSV":  1,
	"XLSX": 2,
}

func (x ListLogRequest_LogFormat) String() string {
	return proto.EnumName(ListLogRequest_LogFormat_name, int32(x))
}
func (ListLogRequest_LogFormat) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{3, 0} }

type RecorderPutResponse struct {
}

func (m *RecorderPutResponse) Reset()                    { *m = RecorderPutResponse{} }
func (m *RecorderPutResponse) String() string            { return proto.CompactTextString(m) }
func (*RecorderPutResponse) ProtoMessage()               {}
func (*RecorderPutResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// Log is a generic message format used by the sync service
// to publish log messages to the various log repositories (typically, bleve).
type Log struct {
	Message map[string]string `protobuf:"bytes,1,rep,name=message" json:"message,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Log) Reset()                    { *m = Log{} }
func (m *Log) String() string            { return proto.CompactTextString(m) }
func (*Log) ProtoMessage()               {}
func (*Log) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Log) GetMessage() map[string]string {
	if m != nil {
		return m.Message
	}
	return nil
}

// LogMessage is the format used to transmit log messages to clients via the REST API.
type LogMessage struct {
	// Generic zap fields
	Ts     int32  `protobuf:"varint,1,opt,name=Ts" json:"Ts,omitempty"`
	Level  string `protobuf:"bytes,2,opt,name=Level" json:"Level,omitempty"`
	Logger string `protobuf:"bytes,3,opt,name=Logger" json:"Logger,omitempty"`
	Msg    string `protobuf:"bytes,4,opt,name=Msg" json:"Msg,omitempty"`
	// Pydio specific
	MsgId string `protobuf:"bytes,5,opt,name=MsgId" json:"MsgId,omitempty"`
	// User Info
	UserName  string   `protobuf:"bytes,6,opt,name=UserName" json:"UserName,omitempty"`
	UserUuid  string   `protobuf:"bytes,7,opt,name=UserUuid" json:"UserUuid,omitempty"`
	GroupPath string   `protobuf:"bytes,8,opt,name=GroupPath" json:"GroupPath,omitempty"`
	Profile   string   `protobuf:"bytes,16,opt,name=Profile" json:"Profile,omitempty"`
	RoleUuids []string `protobuf:"bytes,9,rep,name=RoleUuids" json:"RoleUuids,omitempty"`
	// Client info
	RemoteAddress string `protobuf:"bytes,10,opt,name=RemoteAddress" json:"RemoteAddress,omitempty"`
	UserAgent     string `protobuf:"bytes,11,opt,name=UserAgent" json:"UserAgent,omitempty"`
	HttpProtocol  string `protobuf:"bytes,12,opt,name=HttpProtocol" json:"HttpProtocol,omitempty"`
	// Tree Info
	NodeUuid string `protobuf:"bytes,13,opt,name=NodeUuid" json:"NodeUuid,omitempty"`
	NodePath string `protobuf:"bytes,14,opt,name=NodePath" json:"NodePath,omitempty"`
	WsUuid   string `protobuf:"bytes,15,opt,name=WsUuid" json:"WsUuid,omitempty"`
	WsScope  string `protobuf:"bytes,17,opt,name=WsScope" json:"WsScope,omitempty"`
	// Span Info
	SpanUuid       string `protobuf:"bytes,18,opt,name=SpanUuid" json:"SpanUuid,omitempty"`
	SpanParentUuid string `protobuf:"bytes,19,opt,name=SpanParentUuid" json:"SpanParentUuid,omitempty"`
	SpanRootUuid   string `protobuf:"bytes,20,opt,name=SpanRootUuid" json:"SpanRootUuid,omitempty"`
	// High Level Operation Info
	OperationUuid           string `protobuf:"bytes,21,opt,name=OperationUuid" json:"OperationUuid,omitempty"`
	OperationLabel          string `protobuf:"bytes,22,opt,name=OperationLabel" json:"OperationLabel,omitempty"`
	SchedulerJobUuid        string `protobuf:"bytes,23,opt,name=SchedulerJobUuid" json:"SchedulerJobUuid,omitempty"`
	SchedulerTaskUuid       string `protobuf:"bytes,24,opt,name=SchedulerTaskUuid" json:"SchedulerTaskUuid,omitempty"`
	SchedulerTaskActionPath string `protobuf:"bytes,25,opt,name=SchedulerTaskActionPath" json:"SchedulerTaskActionPath,omitempty"`
}

func (m *LogMessage) Reset()                    { *m = LogMessage{} }
func (m *LogMessage) String() string            { return proto.CompactTextString(m) }
func (*LogMessage) ProtoMessage()               {}
func (*LogMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *LogMessage) GetTs() int32 {
	if m != nil {
		return m.Ts
	}
	return 0
}

func (m *LogMessage) GetLevel() string {
	if m != nil {
		return m.Level
	}
	return ""
}

func (m *LogMessage) GetLogger() string {
	if m != nil {
		return m.Logger
	}
	return ""
}

func (m *LogMessage) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *LogMessage) GetMsgId() string {
	if m != nil {
		return m.MsgId
	}
	return ""
}

func (m *LogMessage) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *LogMessage) GetUserUuid() string {
	if m != nil {
		return m.UserUuid
	}
	return ""
}

func (m *LogMessage) GetGroupPath() string {
	if m != nil {
		return m.GroupPath
	}
	return ""
}

func (m *LogMessage) GetProfile() string {
	if m != nil {
		return m.Profile
	}
	return ""
}

func (m *LogMessage) GetRoleUuids() []string {
	if m != nil {
		return m.RoleUuids
	}
	return nil
}

func (m *LogMessage) GetRemoteAddress() string {
	if m != nil {
		return m.RemoteAddress
	}
	return ""
}

func (m *LogMessage) GetUserAgent() string {
	if m != nil {
		return m.UserAgent
	}
	return ""
}

func (m *LogMessage) GetHttpProtocol() string {
	if m != nil {
		return m.HttpProtocol
	}
	return ""
}

func (m *LogMessage) GetNodeUuid() string {
	if m != nil {
		return m.NodeUuid
	}
	return ""
}

func (m *LogMessage) GetNodePath() string {
	if m != nil {
		return m.NodePath
	}
	return ""
}

func (m *LogMessage) GetWsUuid() string {
	if m != nil {
		return m.WsUuid
	}
	return ""
}

func (m *LogMessage) GetWsScope() string {
	if m != nil {
		return m.WsScope
	}
	return ""
}

func (m *LogMessage) GetSpanUuid() string {
	if m != nil {
		return m.SpanUuid
	}
	return ""
}

func (m *LogMessage) GetSpanParentUuid() string {
	if m != nil {
		return m.SpanParentUuid
	}
	return ""
}

func (m *LogMessage) GetSpanRootUuid() string {
	if m != nil {
		return m.SpanRootUuid
	}
	return ""
}

func (m *LogMessage) GetOperationUuid() string {
	if m != nil {
		return m.OperationUuid
	}
	return ""
}

func (m *LogMessage) GetOperationLabel() string {
	if m != nil {
		return m.OperationLabel
	}
	return ""
}

func (m *LogMessage) GetSchedulerJobUuid() string {
	if m != nil {
		return m.SchedulerJobUuid
	}
	return ""
}

func (m *LogMessage) GetSchedulerTaskUuid() string {
	if m != nil {
		return m.SchedulerTaskUuid
	}
	return ""
}

func (m *LogMessage) GetSchedulerTaskActionPath() string {
	if m != nil {
		return m.SchedulerTaskActionPath
	}
	return ""
}

// ListLogRequest launches a parameterised query in the log repository and streams the results.
type ListLogRequest struct {
	// Bleve-type Query stsring
	Query string `protobuf:"bytes,1,opt,name=Query" json:"Query,omitempty"`
	// Start at page
	Page int32 `protobuf:"varint,2,opt,name=Page" json:"Page,omitempty"`
	// Number of results
	Size   int32                    `protobuf:"varint,3,opt,name=Size" json:"Size,omitempty"`
	Format ListLogRequest_LogFormat `protobuf:"varint,4,opt,name=Format,enum=log.ListLogRequest_LogFormat" json:"Format,omitempty"`
}

func (m *ListLogRequest) Reset()                    { *m = ListLogRequest{} }
func (m *ListLogRequest) String() string            { return proto.CompactTextString(m) }
func (*ListLogRequest) ProtoMessage()               {}
func (*ListLogRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ListLogRequest) GetQuery() string {
	if m != nil {
		return m.Query
	}
	return ""
}

func (m *ListLogRequest) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *ListLogRequest) GetSize() int32 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *ListLogRequest) GetFormat() ListLogRequest_LogFormat {
	if m != nil {
		return m.Format
	}
	return ListLogRequest_JSON
}

type ListLogResponse struct {
	LogMessage *LogMessage `protobuf:"bytes,1,opt,name=LogMessage" json:"LogMessage,omitempty"`
}

func (m *ListLogResponse) Reset()                    { *m = ListLogResponse{} }
func (m *ListLogResponse) String() string            { return proto.CompactTextString(m) }
func (*ListLogResponse) ProtoMessage()               {}
func (*ListLogResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ListLogResponse) GetLogMessage() *LogMessage {
	if m != nil {
		return m.LogMessage
	}
	return nil
}

type DeleteLogsResponse struct {
	Deleted int64 `protobuf:"varint,1,opt,name=Deleted" json:"Deleted,omitempty"`
}

func (m *DeleteLogsResponse) Reset()                    { *m = DeleteLogsResponse{} }
func (m *DeleteLogsResponse) String() string            { return proto.CompactTextString(m) }
func (*DeleteLogsResponse) ProtoMessage()               {}
func (*DeleteLogsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *DeleteLogsResponse) GetDeleted() int64 {
	if m != nil {
		return m.Deleted
	}
	return 0
}

// TimeRangeResponse contains either one aggregated result of a facetted request
// OR a time range cursor.
type TimeRangeResponse struct {
	TimeRangeResult *TimeRangeResult `protobuf:"bytes,1,opt,name=TimeRangeResult" json:"TimeRangeResult,omitempty"`
	TimeRangeCursor *TimeRangeCursor `protobuf:"bytes,2,opt,name=TimeRangeCursor" json:"TimeRangeCursor,omitempty"`
}

func (m *TimeRangeResponse) Reset()                    { *m = TimeRangeResponse{} }
func (m *TimeRangeResponse) String() string            { return proto.CompactTextString(m) }
func (*TimeRangeResponse) ProtoMessage()               {}
func (*TimeRangeResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *TimeRangeResponse) GetTimeRangeResult() *TimeRangeResult {
	if m != nil {
		return m.TimeRangeResult
	}
	return nil
}

func (m *TimeRangeResponse) GetTimeRangeCursor() *TimeRangeCursor {
	if m != nil {
		return m.TimeRangeCursor
	}
	return nil
}

// TimeRangeResult represents one point of a graph.
type TimeRangeResult struct {
	// a label for this time range
	Name string `protobuf:"bytes,1,opt,name=Name" json:"Name,omitempty"`
	// begin timestamp
	Start int32 `protobuf:"varint,2,opt,name=Start" json:"Start,omitempty"`
	// end timestamp
	End int32 `protobuf:"varint,3,opt,name=End" json:"End,omitempty"`
	// nb of occurrences found within this range
	Count int32 `protobuf:"varint,4,opt,name=Count" json:"Count,omitempty"`
	// a score between 1 and 100 that gives the relevance of this result:
	// if End > now, we ponderate the returned count with the duration of the last time range
	// for instance for a hour range if now is 6PM, last count will be
	// multiplied by 4/3 and have a relevance of 75.
	// Relevance will be almost always equals to 100
	Relevance int32 `protobuf:"varint,5,opt,name=Relevance" json:"Relevance,omitempty"`
}

func (m *TimeRangeResult) Reset()                    { *m = TimeRangeResult{} }
func (m *TimeRangeResult) String() string            { return proto.CompactTextString(m) }
func (*TimeRangeResult) ProtoMessage()               {}
func (*TimeRangeResult) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *TimeRangeResult) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TimeRangeResult) GetStart() int32 {
	if m != nil {
		return m.Start
	}
	return 0
}

func (m *TimeRangeResult) GetEnd() int32 {
	if m != nil {
		return m.End
	}
	return 0
}

func (m *TimeRangeResult) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *TimeRangeResult) GetRelevance() int32 {
	if m != nil {
		return m.Relevance
	}
	return 0
}

// TimeRangeRequest contains the parameter to configure the query to
// retrieve the number of audit events of this type for a given time range
// defined by last timestamp and a range type.
type TimeRangeRequest struct {
	// Type of event we are auditing
	MsgId string `protobuf:"bytes,1,opt,name=MsgId" json:"MsgId,omitempty"`
	// Known types: H, D, W, M or Y
	TimeRangeType string `protobuf:"bytes,2,opt,name=TimeRangeType" json:"TimeRangeType,omitempty"`
	// Upper bound for our request
	RefTime int32 `protobuf:"varint,3,opt,name=RefTime" json:"RefTime,omitempty"`
}

func (m *TimeRangeRequest) Reset()                    { *m = TimeRangeRequest{} }
func (m *TimeRangeRequest) String() string            { return proto.CompactTextString(m) }
func (*TimeRangeRequest) ProtoMessage()               {}
func (*TimeRangeRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *TimeRangeRequest) GetMsgId() string {
	if m != nil {
		return m.MsgId
	}
	return ""
}

func (m *TimeRangeRequest) GetTimeRangeType() string {
	if m != nil {
		return m.TimeRangeType
	}
	return ""
}

func (m *TimeRangeRequest) GetRefTime() int32 {
	if m != nil {
		return m.RefTime
	}
	return 0
}

// Ease implementation of data navigation for a chart.
type TimeRangeCursor struct {
	Rel     RelType `protobuf:"varint,1,opt,name=Rel,enum=log.RelType" json:"Rel,omitempty"`
	RefTime int32   `protobuf:"varint,2,opt,name=RefTime" json:"RefTime,omitempty"`
	Count   int32   `protobuf:"varint,3,opt,name=Count" json:"Count,omitempty"`
}

func (m *TimeRangeCursor) Reset()                    { *m = TimeRangeCursor{} }
func (m *TimeRangeCursor) String() string            { return proto.CompactTextString(m) }
func (*TimeRangeCursor) ProtoMessage()               {}
func (*TimeRangeCursor) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *TimeRangeCursor) GetRel() RelType {
	if m != nil {
		return m.Rel
	}
	return RelType_NONE
}

func (m *TimeRangeCursor) GetRefTime() int32 {
	if m != nil {
		return m.RefTime
	}
	return 0
}

func (m *TimeRangeCursor) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func init() {
	proto.RegisterType((*RecorderPutResponse)(nil), "log.RecorderPutResponse")
	proto.RegisterType((*Log)(nil), "log.Log")
	proto.RegisterType((*LogMessage)(nil), "log.LogMessage")
	proto.RegisterType((*ListLogRequest)(nil), "log.ListLogRequest")
	proto.RegisterType((*ListLogResponse)(nil), "log.ListLogResponse")
	proto.RegisterType((*DeleteLogsResponse)(nil), "log.DeleteLogsResponse")
	proto.RegisterType((*TimeRangeResponse)(nil), "log.TimeRangeResponse")
	proto.RegisterType((*TimeRangeResult)(nil), "log.TimeRangeResult")
	proto.RegisterType((*TimeRangeRequest)(nil), "log.TimeRangeRequest")
	proto.RegisterType((*TimeRangeCursor)(nil), "log.TimeRangeCursor")
	proto.RegisterEnum("log.RelType", RelType_name, RelType_value)
	proto.RegisterEnum("log.ListLogRequest_LogFormat", ListLogRequest_LogFormat_name, ListLogRequest_LogFormat_value)
}

func init() { proto.RegisterFile("log.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 936 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x56, 0x5f, 0x6f, 0xdb, 0x36,
	0x10, 0x8f, 0xac, 0xd8, 0x89, 0x2f, 0x89, 0xa3, 0x30, 0xff, 0xb8, 0x60, 0x1b, 0x02, 0xa1, 0x18,
	0x82, 0x62, 0x48, 0x8b, 0x0c, 0x03, 0xba, 0xa2, 0x18, 0x90, 0x65, 0xe9, 0xd6, 0x42, 0x49, 0x3c,
	0xca, 0x6d, 0xf3, 0xaa, 0x58, 0x57, 0xc5, 0xa8, 0x6c, 0x7a, 0xa2, 0x14, 0x20, 0x7b, 0xdb, 0x67,
	0xd8, 0xeb, 0xbe, 0xc3, 0x3e, 0xe1, 0x80, 0xe1, 0x8e, 0x92, 0x6c, 0xd9, 0xdd, 0x1b, 0xef, 0x77,
	0xbf, 0xdf, 0x1d, 0x8f, 0x3c, 0x9e, 0x04, 0xdd, 0x54, 0x27, 0xa7, 0xd3, 0x4c, 0xe7, 0x5a, 0xb8,
	0xa9, 0x4e, 0xfc, 0x7d, 0xd8, 0x55, 0x38, 0xd4, 0x59, 0x8c, 0x59, 0xbf, 0xc8, 0x15, 0x9a, 0xa9,
	0x9e, 0x18, 0xf4, 0x33, 0x70, 0x03, 0x9d, 0x88, 0x67, 0xb0, 0x36, 0x46, 0x63, 0xa2, 0x04, 0xa5,
	0x73, 0xec, 0x9e, 0x6c, 0x9c, 0xed, 0x9f, 0x92, 0x3e, 0xd0, 0xc9, 0xe9, 0x95, 0xc5, 0x2f, 0x27,
	0x79, 0xf6, 0xa8, 0x2a, 0xd6, 0xd1, 0x4b, 0xd8, 0x9c, 0x77, 0x08, 0x0f, 0xdc, 0x4f, 0xf8, 0x28,
	0x9d, 0x63, 0xe7, 0xa4, 0xab, 0x68, 0x29, 0xf6, 0xa0, 0xfd, 0x10, 0xa5, 0x05, 0xca, 0x16, 0x63,
	0xd6, 0x78, 0xd9, 0x7a, 0xe1, 0xf8, 0x7f, 0x77, 0x00, 0x02, 0x9d, 0x94, 0x7a, 0xd1, 0x83, 0xd6,
	0xc0, 0xb0, 0xb2, 0xad, 0x5a, 0x03, 0x43, 0xc2, 0x00, 0x1f, 0x30, 0xad, 0x84, 0x6c, 0x88, 0x03,
	0xe8, 0x04, 0x3a, 0x49, 0x30, 0x93, 0x2e, 0xc3, 0xa5, 0x45, 0x89, 0xaf, 0x4c, 0x22, 0x57, 0x6d,
	0xe2, 0x2b, 0x93, 0x90, 0xfe, 0xca, 0x24, 0x6f, 0x62, 0xd9, 0xb6, 0x7a, 0x36, 0xc4, 0x11, 0xac,
	0xbf, 0x33, 0x98, 0x5d, 0x47, 0x63, 0x94, 0x1d, 0x76, 0xd4, 0x76, 0xe5, 0x7b, 0x57, 0x8c, 0x62,
	0xb9, 0x36, 0xf3, 0x91, 0x2d, 0xbe, 0x84, 0xee, 0x2f, 0x99, 0x2e, 0xa6, 0xfd, 0x28, 0xbf, 0x97,
	0xeb, 0xec, 0x9c, 0x01, 0x42, 0xc2, 0x5a, 0x3f, 0xd3, 0x1f, 0x47, 0x29, 0x4a, 0x8f, 0x7d, 0x95,
	0x49, 0x3a, 0xa5, 0x53, 0xa4, 0x18, 0x46, 0x76, 0x8f, 0x5d, 0xd2, 0xd5, 0x80, 0x78, 0x02, 0x5b,
	0x0a, 0xc7, 0x3a, 0xc7, 0xf3, 0x38, 0xce, 0xd0, 0x18, 0x09, 0xac, 0x6e, 0x82, 0x14, 0x83, 0xf6,
	0x71, 0x9e, 0xe0, 0x24, 0x97, 0x1b, 0x36, 0x77, 0x0d, 0x08, 0x1f, 0x36, 0x7f, 0xcd, 0xf3, 0x69,
	0x9f, 0xee, 0x78, 0xa8, 0x53, 0xb9, 0xc9, 0x84, 0x06, 0x46, 0x95, 0x5d, 0xeb, 0x98, 0x93, 0xca,
	0x2d, 0x5b, 0x59, 0x65, 0x57, 0x3e, 0x2e, 0xac, 0x37, 0xf3, 0x71, 0x5d, 0x07, 0xd0, 0xf9, 0x60,
	0x58, 0xb5, 0x6d, 0x4f, 0xdb, 0x5a, 0x54, 0xef, 0x07, 0x13, 0x0e, 0xf5, 0x14, 0xe5, 0x8e, 0xad,
	0xb7, 0x34, 0x29, 0x5a, 0x38, 0x8d, 0x26, 0xac, 0x11, 0x36, 0x5a, 0x65, 0x8b, 0x6f, 0xa0, 0x47,
	0xeb, 0x7e, 0x94, 0xe1, 0x24, 0x67, 0xc6, 0x2e, 0x33, 0x16, 0x50, 0xaa, 0x88, 0x10, 0xa5, 0xb5,
	0x65, 0xed, 0xd9, 0x8a, 0xe6, 0x31, 0x3a, 0xb9, 0x9b, 0x29, 0x66, 0x51, 0x3e, 0xd2, 0x36, 0xd9,
	0xbe, 0x3d, 0xb9, 0x06, 0x48, 0x19, 0x6b, 0x20, 0x88, 0xee, 0x30, 0x95, 0x07, 0x36, 0x63, 0x13,
	0x15, 0x4f, 0xc1, 0x0b, 0x87, 0xf7, 0x18, 0x17, 0x29, 0x66, 0x6f, 0xf5, 0x1d, 0x07, 0x3c, 0x64,
	0xe6, 0x12, 0x2e, 0xbe, 0x85, 0x9d, 0x1a, 0x1b, 0x44, 0xe6, 0x13, 0x93, 0x25, 0x93, 0x97, 0x1d,
	0xe2, 0x05, 0x1c, 0x36, 0xc0, 0xf3, 0x21, 0x65, 0xe5, 0xc3, 0xfe, 0x82, 0x35, 0xff, 0xe7, 0xf6,
	0xff, 0x71, 0xa0, 0x17, 0x8c, 0x4c, 0x1e, 0xe8, 0x44, 0xe1, 0xef, 0x05, 0x9a, 0x9c, 0x5a, 0xfa,
	0xb7, 0x02, 0xb3, 0xea, 0x7d, 0x59, 0x43, 0x08, 0x58, 0xed, 0xd3, 0x8b, 0x6d, 0xf1, 0xd3, 0xe1,
	0x35, 0x61, 0xe1, 0xe8, 0x0f, 0xe4, 0x47, 0xd2, 0x56, 0xbc, 0x16, 0xdf, 0x43, 0xe7, 0xb5, 0xce,
	0xc6, 0x51, 0xce, 0xaf, 0xa4, 0x77, 0xf6, 0x95, 0x7d, 0xdb, 0x8d, 0x14, 0xf4, 0xd4, 0x2d, 0x49,
	0x95, 0x64, 0xff, 0x04, 0xba, 0x35, 0x28, 0xd6, 0x61, 0xf5, 0x6d, 0x78, 0x73, 0xed, 0xad, 0x88,
	0x35, 0x70, 0x2f, 0xc2, 0xf7, 0x9e, 0x43, 0xd0, 0x6d, 0x10, 0xde, 0x7a, 0x2d, 0xff, 0x27, 0xd8,
	0xae, 0xa3, 0xd9, 0xb9, 0x22, 0x9e, 0xcd, 0x3f, 0x71, 0xde, 0xf6, 0xc6, 0xd9, 0x76, 0x35, 0x53,
	0x4a, 0x58, 0xcd, 0x51, 0xfc, 0x53, 0x10, 0x3f, 0x63, 0x8a, 0x39, 0x06, 0x3a, 0x31, 0x75, 0x18,
	0x09, 0x6b, 0x16, 0x8d, 0x39, 0x86, 0xab, 0x2a, 0xd3, 0xff, 0xcb, 0x81, 0x9d, 0xc1, 0x68, 0x8c,
	0x2a, 0x9a, 0x24, 0x58, 0xf3, 0x7f, 0x84, 0xed, 0x79, 0xb0, 0x48, 0xf3, 0x32, 0xf7, 0x1e, 0xe7,
	0x5e, 0xf0, 0xa9, 0x45, 0x72, 0x43, 0x7f, 0x51, 0x64, 0x46, 0x67, 0x7c, 0xba, 0x4b, 0x7a, 0xeb,
	0x53, 0x8b, 0x64, 0xff, 0x4f, 0x67, 0x69, 0x03, 0x74, 0x25, 0x3c, 0x75, 0xec, 0xdd, 0xf1, 0x9a,
	0x2e, 0x34, 0xcc, 0xa3, 0x2c, 0x2f, 0xef, 0xce, 0x1a, 0x34, 0xcb, 0x2e, 0x27, 0x71, 0x79, 0x77,
	0xb4, 0x24, 0xde, 0x85, 0x2e, 0x26, 0xf6, 0xe6, 0xda, 0xca, 0x1a, 0x3c, 0x5b, 0x30, 0xc5, 0x87,
	0x68, 0x32, 0x44, 0x9e, 0x72, 0x6d, 0x35, 0x03, 0xfc, 0x7b, 0xf0, 0xe6, 0xb6, 0x50, 0x37, 0x90,
	0x9d, 0x89, 0xce, 0xfc, 0x4c, 0x7c, 0x02, 0x5b, 0x35, 0x73, 0xf0, 0x38, 0xad, 0x46, 0x75, 0x13,
	0xa4, 0x3b, 0x50, 0xf8, 0x91, 0xb0, 0x72, 0x67, 0x95, 0xe9, 0x47, 0x4b, 0xa7, 0x25, 0xbe, 0x06,
	0x57, 0x61, 0xca, 0x69, 0x7a, 0x67, 0x9b, 0x7c, 0x68, 0x0a, 0x53, 0x8a, 0xa3, 0xc8, 0x31, 0x1f,
	0xac, 0xd5, 0x08, 0x36, 0x2b, 0xd5, 0x9d, 0x2b, 0xf5, 0xe9, 0x2b, 0xe2, 0xb3, 0x9e, 0xfa, 0xed,
	0xfa, 0xe6, 0xfa, 0xd2, 0x5b, 0x11, 0x5d, 0x68, 0xbf, 0x7e, 0xa3, 0xc2, 0x81, 0x6d, 0xc2, 0xbe,
	0xba, 0x7c, 0xef, 0xb5, 0xd8, 0x7d, 0x79, 0x3b, 0xf0, 0x5c, 0x5a, 0x05, 0xe7, 0xe1, 0xc0, 0x5b,
	0x3d, 0xfb, 0xd7, 0x81, 0x0d, 0xee, 0x4a, 0xfb, 0xe1, 0x13, 0xcf, 0xa1, 0xd3, 0x2f, 0xa8, 0x4f,
	0xc5, 0x7a, 0xd5, 0x8b, 0x47, 0xb2, 0xdc, 0xe4, 0xf2, 0xb7, 0x71, 0xe5, 0xc4, 0x11, 0x3f, 0xc0,
	0x7a, 0xd9, 0xda, 0x46, 0xec, 0x7e, 0xe6, 0xdd, 0x1c, 0xed, 0x35, 0xc1, 0x4a, 0xfa, 0xdc, 0x11,
	0xaf, 0x00, 0x66, 0x1d, 0xfd, 0x79, 0xf1, 0x21, 0x83, 0xcb, 0x7d, 0xef, 0xaf, 0x88, 0x0b, 0xe8,
	0x9d, 0x27, 0x49, 0x86, 0x49, 0x94, 0x63, 0xcc, 0x11, 0xf6, 0x17, 0x5b, 0xd8, 0xc6, 0x38, 0x58,
	0xea, 0xec, 0x7a, 0x0b, 0x77, 0x1d, 0xfe, 0x01, 0xf8, 0xee, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xea, 0x1e, 0x84, 0x56, 0x0d, 0x08, 0x00, 0x00,
}
