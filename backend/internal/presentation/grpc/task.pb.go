// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.4
// source: task.proto

package grpc

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Task_Info_Priority int32

const (
	Task_Info_PRIORITY_UNSPECIFIED Task_Info_Priority = 0
	Task_Info_PRIORITY_LOW         Task_Info_Priority = 1
	Task_Info_PRIORITY_MIDDLE      Task_Info_Priority = 2
	Task_Info_PRIORITY_HIGH        Task_Info_Priority = 3
)

// Enum value maps for Task_Info_Priority.
var (
	Task_Info_Priority_name = map[int32]string{
		0: "PRIORITY_UNSPECIFIED",
		1: "PRIORITY_LOW",
		2: "PRIORITY_MIDDLE",
		3: "PRIORITY_HIGH",
	}
	Task_Info_Priority_value = map[string]int32{
		"PRIORITY_UNSPECIFIED": 0,
		"PRIORITY_LOW":         1,
		"PRIORITY_MIDDLE":      2,
		"PRIORITY_HIGH":        3,
	}
)

func (x Task_Info_Priority) Enum() *Task_Info_Priority {
	p := new(Task_Info_Priority)
	*p = x
	return p
}

func (x Task_Info_Priority) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Task_Info_Priority) Descriptor() protoreflect.EnumDescriptor {
	return file_task_proto_enumTypes[0].Descriptor()
}

func (Task_Info_Priority) Type() protoreflect.EnumType {
	return &file_task_proto_enumTypes[0]
}

func (x Task_Info_Priority) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Task_Info_Priority.Descriptor instead.
func (Task_Info_Priority) EnumDescriptor() ([]byte, []int) {
	return file_task_proto_rawDescGZIP(), []int{5, 0, 0}
}

type Task_Info_Status int32

const (
	Task_Info_STATUS_UNSPECIFIED Task_Info_Status = 0
	Task_Info_STATUS_TODO        Task_Info_Status = 1
	Task_Info_STATUS_DOING       Task_Info_Status = 2
	Task_Info_STATUS_DONE        Task_Info_Status = 3
)

// Enum value maps for Task_Info_Status.
var (
	Task_Info_Status_name = map[int32]string{
		0: "STATUS_UNSPECIFIED",
		1: "STATUS_TODO",
		2: "STATUS_DOING",
		3: "STATUS_DONE",
	}
	Task_Info_Status_value = map[string]int32{
		"STATUS_UNSPECIFIED": 0,
		"STATUS_TODO":        1,
		"STATUS_DOING":       2,
		"STATUS_DONE":        3,
	}
)

func (x Task_Info_Status) Enum() *Task_Info_Status {
	p := new(Task_Info_Status)
	*p = x
	return p
}

func (x Task_Info_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Task_Info_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_task_proto_enumTypes[1].Descriptor()
}

func (Task_Info_Status) Type() protoreflect.EnumType {
	return &file_task_proto_enumTypes[1]
}

func (x Task_Info_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Task_Info_Status.Descriptor instead.
func (Task_Info_Status) EnumDescriptor() ([]byte, []int) {
	return file_task_proto_rawDescGZIP(), []int{5, 0, 1}
}

type TaskListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *TaskListRequest) Reset() {
	*x = TaskListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_task_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskListRequest) ProtoMessage() {}

func (x *TaskListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_task_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskListRequest.ProtoReflect.Descriptor instead.
func (*TaskListRequest) Descriptor() ([]byte, []int) {
	return file_task_proto_rawDescGZIP(), []int{0}
}

type TaskListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tasks []*Task `protobuf:"bytes,1,rep,name=tasks,proto3" json:"tasks,omitempty"`
}

func (x *TaskListResponse) Reset() {
	*x = TaskListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_task_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskListResponse) ProtoMessage() {}

func (x *TaskListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_task_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskListResponse.ProtoReflect.Descriptor instead.
func (*TaskListResponse) Descriptor() ([]byte, []int) {
	return file_task_proto_rawDescGZIP(), []int{1}
}

func (x *TaskListResponse) GetTasks() []*Task {
	if x != nil {
		return x.Tasks
	}
	return nil
}

type TaskGetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *TaskGetRequest) Reset() {
	*x = TaskGetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_task_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskGetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskGetRequest) ProtoMessage() {}

func (x *TaskGetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_task_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskGetRequest.ProtoReflect.Descriptor instead.
func (*TaskGetRequest) Descriptor() ([]byte, []int) {
	return file_task_proto_rawDescGZIP(), []int{2}
}

func (x *TaskGetRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type TaskCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *Task_Info `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *TaskCreateRequest) Reset() {
	*x = TaskCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_task_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskCreateRequest) ProtoMessage() {}

func (x *TaskCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_task_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskCreateRequest.ProtoReflect.Descriptor instead.
func (*TaskCreateRequest) Descriptor() ([]byte, []int) {
	return file_task_proto_rawDescGZIP(), []int{3}
}

func (x *TaskCreateRequest) GetInfo() *Task_Info {
	if x != nil {
		return x.Info
	}
	return nil
}

type TaskDeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *TaskDeleteRequest) Reset() {
	*x = TaskDeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_task_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskDeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskDeleteRequest) ProtoMessage() {}

func (x *TaskDeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_task_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskDeleteRequest.ProtoReflect.Descriptor instead.
func (*TaskDeleteRequest) Descriptor() ([]byte, []int) {
	return file_task_proto_rawDescGZIP(), []int{4}
}

func (x *TaskDeleteRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type Task struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Info *Task_Info `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *Task) Reset() {
	*x = Task{}
	if protoimpl.UnsafeEnabled {
		mi := &file_task_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Task) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Task) ProtoMessage() {}

func (x *Task) ProtoReflect() protoreflect.Message {
	mi := &file_task_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Task.ProtoReflect.Descriptor instead.
func (*Task) Descriptor() ([]byte, []int) {
	return file_task_proto_rawDescGZIP(), []int{5}
}

func (x *Task) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Task) GetInfo() *Task_Info {
	if x != nil {
		return x.Info
	}
	return nil
}

type Task_Info struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title       string                 `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description string                 `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Limit       *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=limit,proto3" json:"limit,omitempty"`
	Priority    Task_Info_Priority     `protobuf:"varint,4,opt,name=priority,proto3,enum=api.Task_Info_Priority" json:"priority,omitempty"`
	Status      Task_Info_Status       `protobuf:"varint,5,opt,name=status,proto3,enum=api.Task_Info_Status" json:"status,omitempty"`
	Labels      []string               `protobuf:"bytes,6,rep,name=labels,proto3" json:"labels,omitempty"`
}

func (x *Task_Info) Reset() {
	*x = Task_Info{}
	if protoimpl.UnsafeEnabled {
		mi := &file_task_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Task_Info) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Task_Info) ProtoMessage() {}

func (x *Task_Info) ProtoReflect() protoreflect.Message {
	mi := &file_task_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Task_Info.ProtoReflect.Descriptor instead.
func (*Task_Info) Descriptor() ([]byte, []int) {
	return file_task_proto_rawDescGZIP(), []int{5, 0}
}

func (x *Task_Info) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Task_Info) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Task_Info) GetLimit() *timestamppb.Timestamp {
	if x != nil {
		return x.Limit
	}
	return nil
}

func (x *Task_Info) GetPriority() Task_Info_Priority {
	if x != nil {
		return x.Priority
	}
	return Task_Info_PRIORITY_UNSPECIFIED
}

func (x *Task_Info) GetStatus() Task_Info_Status {
	if x != nil {
		return x.Status
	}
	return Task_Info_STATUS_UNSPECIFIED
}

func (x *Task_Info) GetLabels() []string {
	if x != nil {
		return x.Labels
	}
	return nil
}

var File_task_proto protoreflect.FileDescriptor

var file_task_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x61, 0x70,
	0x69, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x11, 0x0a, 0x0f, 0x54, 0x61, 0x73, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x33, 0x0a, 0x10, 0x54, 0x61, 0x73, 0x6b, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x05, 0x74, 0x61, 0x73,
	0x6b, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x54,
	0x61, 0x73, 0x6b, 0x52, 0x05, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x22, 0x20, 0x0a, 0x0e, 0x54, 0x61,
	0x73, 0x6b, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x37, 0x0a, 0x11,
	0x54, 0x61, 0x73, 0x6b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x22, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x04, 0x69, 0x6e, 0x66, 0x6f, 0x22, 0x23, 0x0a, 0x11, 0x54, 0x61, 0x73, 0x6b, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0xdf, 0x03, 0x0a, 0x04, 0x54,
	0x61, 0x73, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x22, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x2e, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x1a, 0xa2, 0x03, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x30, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x33, 0x0a, 0x08, 0x70, 0x72,
	0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x50, 0x72, 0x69,
	0x6f, 0x72, 0x69, 0x74, 0x79, 0x52, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x12,
	0x2d, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x15, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x2e,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16,
	0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06,
	0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x22, 0x5e, 0x0a, 0x08, 0x50, 0x72, 0x69, 0x6f, 0x72, 0x69,
	0x74, 0x79, 0x12, 0x18, 0x0a, 0x14, 0x50, 0x52, 0x49, 0x4f, 0x52, 0x49, 0x54, 0x59, 0x5f, 0x55,
	0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c,
	0x50, 0x52, 0x49, 0x4f, 0x52, 0x49, 0x54, 0x59, 0x5f, 0x4c, 0x4f, 0x57, 0x10, 0x01, 0x12, 0x13,
	0x0a, 0x0f, 0x50, 0x52, 0x49, 0x4f, 0x52, 0x49, 0x54, 0x59, 0x5f, 0x4d, 0x49, 0x44, 0x44, 0x4c,
	0x45, 0x10, 0x02, 0x12, 0x11, 0x0a, 0x0d, 0x50, 0x52, 0x49, 0x4f, 0x52, 0x49, 0x54, 0x59, 0x5f,
	0x48, 0x49, 0x47, 0x48, 0x10, 0x03, 0x22, 0x54, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x16, 0x0a, 0x12, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x53, 0x54, 0x41, 0x54,
	0x55, 0x53, 0x5f, 0x54, 0x4f, 0x44, 0x4f, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x53, 0x54, 0x41,
	0x54, 0x55, 0x53, 0x5f, 0x44, 0x4f, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x12, 0x0f, 0x0a, 0x0b, 0x53,
	0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x44, 0x4f, 0x4e, 0x45, 0x10, 0x03, 0x32, 0x87, 0x02, 0x0a,
	0x0b, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x33, 0x0a, 0x04,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x14, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x54, 0x61, 0x73, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x25, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x13, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x54,
	0x61, 0x73, 0x6b, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x09, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x37, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x12, 0x16, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x54, 0x61, 0x73, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x2a, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x09, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x1a, 0x15, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x54, 0x61, 0x73,
	0x6b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a,
	0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x16, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x54, 0x61,
	0x73, 0x6b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x15, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x26, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x6f, 0x69, 0x53, 0x61, 0x74, 0x6f, 0x6e, 0x61, 0x6b, 0x61,
	0x2f, 0x74, 0x6f, 0x64, 0x6f, 0x2d, 0x61, 0x70, 0x70, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_task_proto_rawDescOnce sync.Once
	file_task_proto_rawDescData = file_task_proto_rawDesc
)

func file_task_proto_rawDescGZIP() []byte {
	file_task_proto_rawDescOnce.Do(func() {
		file_task_proto_rawDescData = protoimpl.X.CompressGZIP(file_task_proto_rawDescData)
	})
	return file_task_proto_rawDescData
}

var file_task_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_task_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_task_proto_goTypes = []interface{}{
	(Task_Info_Priority)(0),       // 0: api.Task.Info.Priority
	(Task_Info_Status)(0),         // 1: api.Task.Info.Status
	(*TaskListRequest)(nil),       // 2: api.TaskListRequest
	(*TaskListResponse)(nil),      // 3: api.TaskListResponse
	(*TaskGetRequest)(nil),        // 4: api.TaskGetRequest
	(*TaskCreateRequest)(nil),     // 5: api.TaskCreateRequest
	(*TaskDeleteRequest)(nil),     // 6: api.TaskDeleteRequest
	(*Task)(nil),                  // 7: api.Task
	(*Task_Info)(nil),             // 8: api.Task.Info
	(*timestamppb.Timestamp)(nil), // 9: google.protobuf.Timestamp
}
var file_task_proto_depIdxs = []int32{
	7,  // 0: api.TaskListResponse.tasks:type_name -> api.Task
	8,  // 1: api.TaskCreateRequest.info:type_name -> api.Task.Info
	8,  // 2: api.Task.info:type_name -> api.Task.Info
	9,  // 3: api.Task.Info.limit:type_name -> google.protobuf.Timestamp
	0,  // 4: api.Task.Info.priority:type_name -> api.Task.Info.Priority
	1,  // 5: api.Task.Info.status:type_name -> api.Task.Info.Status
	2,  // 6: api.TaskService.List:input_type -> api.TaskListRequest
	4,  // 7: api.TaskService.Get:input_type -> api.TaskGetRequest
	5,  // 8: api.TaskService.Create:input_type -> api.TaskCreateRequest
	7,  // 9: api.TaskService.Update:input_type -> api.Task
	6,  // 10: api.TaskService.Delete:input_type -> api.TaskDeleteRequest
	3,  // 11: api.TaskService.List:output_type -> api.TaskListResponse
	7,  // 12: api.TaskService.Get:output_type -> api.Task
	3,  // 13: api.TaskService.Create:output_type -> api.TaskListResponse
	3,  // 14: api.TaskService.Update:output_type -> api.TaskListResponse
	3,  // 15: api.TaskService.Delete:output_type -> api.TaskListResponse
	11, // [11:16] is the sub-list for method output_type
	6,  // [6:11] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_task_proto_init() }
func file_task_proto_init() {
	if File_task_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_task_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskListRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_task_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskListResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_task_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskGetRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_task_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskCreateRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_task_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskDeleteRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_task_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Task); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_task_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Task_Info); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_task_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_task_proto_goTypes,
		DependencyIndexes: file_task_proto_depIdxs,
		EnumInfos:         file_task_proto_enumTypes,
		MessageInfos:      file_task_proto_msgTypes,
	}.Build()
	File_task_proto = out.File
	file_task_proto_rawDesc = nil
	file_task_proto_goTypes = nil
	file_task_proto_depIdxs = nil
}
