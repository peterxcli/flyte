// Code generated by protoc-gen-go. DO NOT EDIT.
// source: flyteidl/service/external_plugin_service.proto

package service

import (
	context "context"
	fmt "fmt"
	core "github.com/flyteorg/flyte/flyteidl/gen/pb-go/flyteidl/core"
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

// The state of the execution is used to control its visibility in the UI/CLI.
type State int32 // Deprecated: Do not use.
const (
	State_RETRYABLE_FAILURE State = 0
	State_PERMANENT_FAILURE State = 1
	State_PENDING           State = 2
	State_RUNNING           State = 3
	State_SUCCEEDED         State = 4
)

var State_name = map[int32]string{
	0: "RETRYABLE_FAILURE",
	1: "PERMANENT_FAILURE",
	2: "PENDING",
	3: "RUNNING",
	4: "SUCCEEDED",
}

var State_value = map[string]int32{
	"RETRYABLE_FAILURE": 0,
	"PERMANENT_FAILURE": 1,
	"PENDING":           2,
	"RUNNING":           3,
	"SUCCEEDED":         4,
}

func (x State) String() string {
	return proto.EnumName(State_name, int32(x))
}

func (State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_74cbdb08eef5b1d1, []int{0}
}

// Represents a request structure to create task.
//
// Deprecated: Do not use.
type TaskCreateRequest struct {
	// The inputs required to start the execution. All required inputs must be
	// included in this map. If not required and not provided, defaults apply.
	// +optional
	Inputs *core.LiteralMap `protobuf:"bytes,1,opt,name=inputs,proto3" json:"inputs,omitempty"`
	// Template of the task that encapsulates all the metadata of the task.
	Template *core.TaskTemplate `protobuf:"bytes,2,opt,name=template,proto3" json:"template,omitempty"`
	// Prefix for where task output data will be written. (e.g. s3://my-bucket/randomstring)
	OutputPrefix         string   `protobuf:"bytes,3,opt,name=output_prefix,json=outputPrefix,proto3" json:"output_prefix,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TaskCreateRequest) Reset()         { *m = TaskCreateRequest{} }
func (m *TaskCreateRequest) String() string { return proto.CompactTextString(m) }
func (*TaskCreateRequest) ProtoMessage()    {}
func (*TaskCreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_74cbdb08eef5b1d1, []int{0}
}

func (m *TaskCreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskCreateRequest.Unmarshal(m, b)
}
func (m *TaskCreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskCreateRequest.Marshal(b, m, deterministic)
}
func (m *TaskCreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskCreateRequest.Merge(m, src)
}
func (m *TaskCreateRequest) XXX_Size() int {
	return xxx_messageInfo_TaskCreateRequest.Size(m)
}
func (m *TaskCreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskCreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TaskCreateRequest proto.InternalMessageInfo

func (m *TaskCreateRequest) GetInputs() *core.LiteralMap {
	if m != nil {
		return m.Inputs
	}
	return nil
}

func (m *TaskCreateRequest) GetTemplate() *core.TaskTemplate {
	if m != nil {
		return m.Template
	}
	return nil
}

func (m *TaskCreateRequest) GetOutputPrefix() string {
	if m != nil {
		return m.OutputPrefix
	}
	return ""
}

// Represents a create response structure.
//
// Deprecated: Do not use.
type TaskCreateResponse struct {
	JobId                string   `protobuf:"bytes,1,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TaskCreateResponse) Reset()         { *m = TaskCreateResponse{} }
func (m *TaskCreateResponse) String() string { return proto.CompactTextString(m) }
func (*TaskCreateResponse) ProtoMessage()    {}
func (*TaskCreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_74cbdb08eef5b1d1, []int{1}
}

func (m *TaskCreateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskCreateResponse.Unmarshal(m, b)
}
func (m *TaskCreateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskCreateResponse.Marshal(b, m, deterministic)
}
func (m *TaskCreateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskCreateResponse.Merge(m, src)
}
func (m *TaskCreateResponse) XXX_Size() int {
	return xxx_messageInfo_TaskCreateResponse.Size(m)
}
func (m *TaskCreateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskCreateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TaskCreateResponse proto.InternalMessageInfo

func (m *TaskCreateResponse) GetJobId() string {
	if m != nil {
		return m.JobId
	}
	return ""
}

// A message used to fetch a job state from backend plugin server.
//
// Deprecated: Do not use.
type TaskGetRequest struct {
	// A predefined yet extensible Task type identifier.
	TaskType string `protobuf:"bytes,1,opt,name=task_type,json=taskType,proto3" json:"task_type,omitempty"`
	// The unique id identifying the job.
	JobId                string   `protobuf:"bytes,2,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TaskGetRequest) Reset()         { *m = TaskGetRequest{} }
func (m *TaskGetRequest) String() string { return proto.CompactTextString(m) }
func (*TaskGetRequest) ProtoMessage()    {}
func (*TaskGetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_74cbdb08eef5b1d1, []int{2}
}

func (m *TaskGetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskGetRequest.Unmarshal(m, b)
}
func (m *TaskGetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskGetRequest.Marshal(b, m, deterministic)
}
func (m *TaskGetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskGetRequest.Merge(m, src)
}
func (m *TaskGetRequest) XXX_Size() int {
	return xxx_messageInfo_TaskGetRequest.Size(m)
}
func (m *TaskGetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskGetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TaskGetRequest proto.InternalMessageInfo

func (m *TaskGetRequest) GetTaskType() string {
	if m != nil {
		return m.TaskType
	}
	return ""
}

func (m *TaskGetRequest) GetJobId() string {
	if m != nil {
		return m.JobId
	}
	return ""
}

// Response to get an individual task state.
//
// Deprecated: Do not use.
type TaskGetResponse struct {
	// The state of the execution is used to control its visibility in the UI/CLI.
	State State `protobuf:"varint,1,opt,name=state,proto3,enum=flyteidl.service.State" json:"state,omitempty"`
	// The outputs of the execution. It's typically used by sql task. Flyteplugins service will create a
	// Structured dataset pointing to the query result table.
	// +optional
	Outputs              *core.LiteralMap `protobuf:"bytes,2,opt,name=outputs,proto3" json:"outputs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *TaskGetResponse) Reset()         { *m = TaskGetResponse{} }
func (m *TaskGetResponse) String() string { return proto.CompactTextString(m) }
func (*TaskGetResponse) ProtoMessage()    {}
func (*TaskGetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_74cbdb08eef5b1d1, []int{3}
}

func (m *TaskGetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskGetResponse.Unmarshal(m, b)
}
func (m *TaskGetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskGetResponse.Marshal(b, m, deterministic)
}
func (m *TaskGetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskGetResponse.Merge(m, src)
}
func (m *TaskGetResponse) XXX_Size() int {
	return xxx_messageInfo_TaskGetResponse.Size(m)
}
func (m *TaskGetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskGetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TaskGetResponse proto.InternalMessageInfo

func (m *TaskGetResponse) GetState() State {
	if m != nil {
		return m.State
	}
	return State_RETRYABLE_FAILURE
}

func (m *TaskGetResponse) GetOutputs() *core.LiteralMap {
	if m != nil {
		return m.Outputs
	}
	return nil
}

// A message used to delete a task.
//
// Deprecated: Do not use.
type TaskDeleteRequest struct {
	// A predefined yet extensible Task type identifier.
	TaskType string `protobuf:"bytes,1,opt,name=task_type,json=taskType,proto3" json:"task_type,omitempty"`
	// The unique id identifying the job.
	JobId                string   `protobuf:"bytes,2,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TaskDeleteRequest) Reset()         { *m = TaskDeleteRequest{} }
func (m *TaskDeleteRequest) String() string { return proto.CompactTextString(m) }
func (*TaskDeleteRequest) ProtoMessage()    {}
func (*TaskDeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_74cbdb08eef5b1d1, []int{4}
}

func (m *TaskDeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskDeleteRequest.Unmarshal(m, b)
}
func (m *TaskDeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskDeleteRequest.Marshal(b, m, deterministic)
}
func (m *TaskDeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskDeleteRequest.Merge(m, src)
}
func (m *TaskDeleteRequest) XXX_Size() int {
	return xxx_messageInfo_TaskDeleteRequest.Size(m)
}
func (m *TaskDeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskDeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TaskDeleteRequest proto.InternalMessageInfo

func (m *TaskDeleteRequest) GetTaskType() string {
	if m != nil {
		return m.TaskType
	}
	return ""
}

func (m *TaskDeleteRequest) GetJobId() string {
	if m != nil {
		return m.JobId
	}
	return ""
}

// Response to delete a task.
//
// Deprecated: Do not use.
type TaskDeleteResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TaskDeleteResponse) Reset()         { *m = TaskDeleteResponse{} }
func (m *TaskDeleteResponse) String() string { return proto.CompactTextString(m) }
func (*TaskDeleteResponse) ProtoMessage()    {}
func (*TaskDeleteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_74cbdb08eef5b1d1, []int{5}
}

func (m *TaskDeleteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskDeleteResponse.Unmarshal(m, b)
}
func (m *TaskDeleteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskDeleteResponse.Marshal(b, m, deterministic)
}
func (m *TaskDeleteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskDeleteResponse.Merge(m, src)
}
func (m *TaskDeleteResponse) XXX_Size() int {
	return xxx_messageInfo_TaskDeleteResponse.Size(m)
}
func (m *TaskDeleteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskDeleteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TaskDeleteResponse proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("flyteidl.service.State", State_name, State_value)
	proto.RegisterType((*TaskCreateRequest)(nil), "flyteidl.service.TaskCreateRequest")
	proto.RegisterType((*TaskCreateResponse)(nil), "flyteidl.service.TaskCreateResponse")
	proto.RegisterType((*TaskGetRequest)(nil), "flyteidl.service.TaskGetRequest")
	proto.RegisterType((*TaskGetResponse)(nil), "flyteidl.service.TaskGetResponse")
	proto.RegisterType((*TaskDeleteRequest)(nil), "flyteidl.service.TaskDeleteRequest")
	proto.RegisterType((*TaskDeleteResponse)(nil), "flyteidl.service.TaskDeleteResponse")
}

func init() {
	proto.RegisterFile("flyteidl/service/external_plugin_service.proto", fileDescriptor_74cbdb08eef5b1d1)
}

var fileDescriptor_74cbdb08eef5b1d1 = []byte{
	// 544 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0x51, 0x6f, 0xd2, 0x50,
	0x14, 0xb6, 0x9d, 0x1b, 0xe3, 0xcc, 0x4d, 0x76, 0x13, 0x62, 0xc7, 0x34, 0x41, 0xe6, 0xc3, 0x62,
	0xb2, 0x36, 0xb2, 0x07, 0x13, 0x13, 0x63, 0x18, 0x54, 0x24, 0xb2, 0x86, 0x14, 0x78, 0xd0, 0x2c,
	0x69, 0x5a, 0x38, 0x60, 0xb7, 0xae, 0xbd, 0xb6, 0xb7, 0x66, 0xfc, 0x03, 0x7f, 0xca, 0x7e, 0xa6,
	0xb9, 0xf7, 0xb6, 0x1d, 0x30, 0xe5, 0xc1, 0x37, 0x7a, 0xce, 0x77, 0xbe, 0x73, 0xbe, 0xef, 0x9c,
	0x0b, 0xe8, 0xb3, 0x60, 0xc1, 0xd0, 0x9f, 0x06, 0x46, 0x82, 0xf1, 0x2f, 0x7f, 0x82, 0x06, 0xde,
	0x31, 0x8c, 0x43, 0x37, 0x70, 0x68, 0x90, 0xce, 0xfd, 0xd0, 0xc9, 0xe2, 0x3a, 0x8d, 0x23, 0x16,
	0x91, 0x4a, 0x8e, 0xd7, 0xb3, 0x78, 0xed, 0x65, 0xc1, 0x30, 0x89, 0x62, 0x34, 0x02, 0x9f, 0x61,
	0xec, 0x06, 0x89, 0xc4, 0xd7, 0x8e, 0x56, 0xb3, 0xcc, 0x4d, 0x6e, 0xf2, 0xd4, 0xab, 0xd5, 0x94,
	0x1f, 0x32, 0x8c, 0x67, 0x6e, 0xde, 0xa9, 0x71, 0xaf, 0xc0, 0xe1, 0xc8, 0x4d, 0x6e, 0xda, 0x31,
	0xba, 0x0c, 0x6d, 0xfc, 0x99, 0x62, 0xc2, 0xc8, 0x3b, 0xd8, 0xf1, 0x43, 0x9a, 0xb2, 0x44, 0x53,
	0xea, 0xca, 0xe9, 0x5e, 0xf3, 0xa8, 0x10, 0xa0, 0x73, 0x16, 0xbd, 0x2f, 0xdb, 0x5f, 0xba, 0xd4,
	0xce, 0x80, 0xe4, 0x3d, 0xec, 0x32, 0xbc, 0xa5, 0x81, 0xcb, 0x50, 0x53, 0x45, 0xd1, 0xf1, 0x5a,
	0x11, 0x6f, 0x33, 0xca, 0x20, 0x76, 0x01, 0x26, 0x27, 0xb0, 0x1f, 0xa5, 0x8c, 0xa6, 0xcc, 0xa1,
	0x31, 0xce, 0xfc, 0x3b, 0x6d, 0xab, 0xae, 0x9c, 0x96, 0xed, 0x67, 0x32, 0x38, 0x10, 0xb1, 0x0f,
	0xaa, 0xa6, 0x34, 0x0c, 0x20, 0xcb, 0x93, 0x26, 0x34, 0x0a, 0x13, 0x24, 0x55, 0xd8, 0xb9, 0x8e,
	0x3c, 0xc7, 0x9f, 0x8a, 0x51, 0xcb, 0xf6, 0xf6, 0x75, 0xe4, 0xf5, 0xa6, 0xa2, 0xe0, 0x0b, 0x1c,
	0xf0, 0x82, 0x2e, 0xb2, 0x5c, 0xd7, 0x31, 0x94, 0xb9, 0x37, 0x0e, 0x5b, 0x50, 0xcc, 0xf0, 0xbb,
	0x3c, 0x30, 0x5a, 0xd0, 0x65, 0x26, 0x75, 0x9d, 0x69, 0x01, 0xcf, 0x0b, 0xa6, 0xac, 0xef, 0x19,
	0x6c, 0x27, 0x8c, 0x8b, 0xe5, 0x34, 0x07, 0xcd, 0x17, 0xfa, 0xfa, 0xca, 0xf4, 0x21, 0x4f, 0xdb,
	0x12, 0x45, 0xce, 0xa1, 0x24, 0x05, 0x25, 0x99, 0x3b, 0x1b, 0x2c, 0xcd, 0x91, 0xa2, 0xf5, 0x57,
	0xb9, 0x9f, 0x0e, 0x06, 0xf8, 0xb0, 0x9f, 0xff, 0xd5, 0xa1, 0x49, 0x0b, 0x73, 0x32, 0x29, 0x85,
	0x67, 0xde, 0x7a, 0xb0, 0x2d, 0xe6, 0x25, 0x55, 0x38, 0xb4, 0xcd, 0x91, 0xfd, 0xad, 0x75, 0xd1,
	0x37, 0x9d, 0xcf, 0xad, 0x5e, 0x7f, 0x6c, 0x9b, 0x95, 0x27, 0x3c, 0x3c, 0x30, 0xed, 0xcb, 0x96,
	0x65, 0x5a, 0xa3, 0x22, 0xac, 0x90, 0x3d, 0x28, 0x0d, 0x4c, 0xab, 0xd3, 0xb3, 0xba, 0x15, 0x95,
	0x7f, 0xd8, 0x63, 0xcb, 0xe2, 0x1f, 0x5b, 0x64, 0x1f, 0xca, 0xc3, 0x71, 0xbb, 0x6d, 0x9a, 0x1d,
	0xb3, 0x53, 0x79, 0x5a, 0x53, 0x35, 0xa5, 0x79, 0xaf, 0x42, 0xd5, 0xcc, 0xee, 0x7e, 0x20, 0xce,
	0x7e, 0x28, 0xad, 0x22, 0x57, 0x00, 0x72, 0xad, 0x7c, 0x3a, 0x72, 0xf2, 0xd8, 0xcb, 0x47, 0x27,
	0x5a, 0x7b, 0xb3, 0x19, 0x24, 0xa5, 0x35, 0xb6, 0x7e, 0xab, 0x0a, 0x19, 0x42, 0xa9, 0x8b, 0x4c,
	0x50, 0xd7, 0xff, 0x5e, 0xf5, 0x70, 0x22, 0xb5, 0xd7, 0x1b, 0x10, 0xcb, 0xa4, 0x57, 0x00, 0xd2,
	0xc6, 0x4d, 0x23, 0xaf, 0x6c, 0xed, 0x5f, 0x23, 0xaf, 0x6e, 0x43, 0xb0, 0x5f, 0x7c, 0xfa, 0xfe,
	0x71, 0xee, 0xb3, 0x1f, 0xa9, 0xa7, 0x4f, 0xa2, 0x5b, 0x43, 0x94, 0x45, 0xf1, 0x5c, 0xfe, 0x30,
	0x8a, 0x17, 0x3d, 0xc7, 0xd0, 0xa0, 0xde, 0xd9, 0x3c, 0x32, 0xd6, 0xff, 0x5f, 0xbc, 0x1d, 0xf1,
	0xbc, 0xcf, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0xda, 0xdb, 0x38, 0x09, 0x7a, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ExternalPluginServiceClient is the client API for ExternalPluginService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ExternalPluginServiceClient interface {
	// Send a task create request to the backend plugin server.
	CreateTask(ctx context.Context, in *TaskCreateRequest, opts ...grpc.CallOption) (*TaskCreateResponse, error)
	// Get job status.
	GetTask(ctx context.Context, in *TaskGetRequest, opts ...grpc.CallOption) (*TaskGetResponse, error)
	// Delete the task resource.
	DeleteTask(ctx context.Context, in *TaskDeleteRequest, opts ...grpc.CallOption) (*TaskDeleteResponse, error)
}

type externalPluginServiceClient struct {
	cc *grpc.ClientConn
}

func NewExternalPluginServiceClient(cc *grpc.ClientConn) ExternalPluginServiceClient {
	return &externalPluginServiceClient{cc}
}

// Deprecated: Do not use.
func (c *externalPluginServiceClient) CreateTask(ctx context.Context, in *TaskCreateRequest, opts ...grpc.CallOption) (*TaskCreateResponse, error) {
	out := new(TaskCreateResponse)
	err := c.cc.Invoke(ctx, "/flyteidl.service.ExternalPluginService/CreateTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Deprecated: Do not use.
func (c *externalPluginServiceClient) GetTask(ctx context.Context, in *TaskGetRequest, opts ...grpc.CallOption) (*TaskGetResponse, error) {
	out := new(TaskGetResponse)
	err := c.cc.Invoke(ctx, "/flyteidl.service.ExternalPluginService/GetTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Deprecated: Do not use.
func (c *externalPluginServiceClient) DeleteTask(ctx context.Context, in *TaskDeleteRequest, opts ...grpc.CallOption) (*TaskDeleteResponse, error) {
	out := new(TaskDeleteResponse)
	err := c.cc.Invoke(ctx, "/flyteidl.service.ExternalPluginService/DeleteTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExternalPluginServiceServer is the server API for ExternalPluginService service.
type ExternalPluginServiceServer interface {
	// Send a task create request to the backend plugin server.
	CreateTask(context.Context, *TaskCreateRequest) (*TaskCreateResponse, error)
	// Get job status.
	GetTask(context.Context, *TaskGetRequest) (*TaskGetResponse, error)
	// Delete the task resource.
	DeleteTask(context.Context, *TaskDeleteRequest) (*TaskDeleteResponse, error)
}

// UnimplementedExternalPluginServiceServer can be embedded to have forward compatible implementations.
type UnimplementedExternalPluginServiceServer struct {
}

func (*UnimplementedExternalPluginServiceServer) CreateTask(ctx context.Context, req *TaskCreateRequest) (*TaskCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTask not implemented")
}
func (*UnimplementedExternalPluginServiceServer) GetTask(ctx context.Context, req *TaskGetRequest) (*TaskGetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTask not implemented")
}
func (*UnimplementedExternalPluginServiceServer) DeleteTask(ctx context.Context, req *TaskDeleteRequest) (*TaskDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTask not implemented")
}

func RegisterExternalPluginServiceServer(s *grpc.Server, srv ExternalPluginServiceServer) {
	s.RegisterService(&_ExternalPluginService_serviceDesc, srv)
}

func _ExternalPluginService_CreateTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExternalPluginServiceServer).CreateTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flyteidl.service.ExternalPluginService/CreateTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExternalPluginServiceServer).CreateTask(ctx, req.(*TaskCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExternalPluginService_GetTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExternalPluginServiceServer).GetTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flyteidl.service.ExternalPluginService/GetTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExternalPluginServiceServer).GetTask(ctx, req.(*TaskGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExternalPluginService_DeleteTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExternalPluginServiceServer).DeleteTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flyteidl.service.ExternalPluginService/DeleteTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExternalPluginServiceServer).DeleteTask(ctx, req.(*TaskDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ExternalPluginService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "flyteidl.service.ExternalPluginService",
	HandlerType: (*ExternalPluginServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTask",
			Handler:    _ExternalPluginService_CreateTask_Handler,
		},
		{
			MethodName: "GetTask",
			Handler:    _ExternalPluginService_GetTask_Handler,
		},
		{
			MethodName: "DeleteTask",
			Handler:    _ExternalPluginService_DeleteTask_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "flyteidl/service/external_plugin_service.proto",
}