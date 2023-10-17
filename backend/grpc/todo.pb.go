// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.4
// source: todo.proto

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

type ToDo_Info_Priority int32

const (
	ToDo_Info_PRIORITY_UNSPECIFIED   ToDo_Info_Priority = 0
	ToDo_Info_PRIORITY_LOW           ToDo_Info_Priority = 1
	ToDo_Info_PRIORITY_NORMAL_MIDDLE ToDo_Info_Priority = 2
	ToDo_Info_PRIORITY_HIGH          ToDo_Info_Priority = 3
)

// Enum value maps for ToDo_Info_Priority.
var (
	ToDo_Info_Priority_name = map[int32]string{
		0: "PRIORITY_UNSPECIFIED",
		1: "PRIORITY_LOW",
		2: "PRIORITY_NORMAL_MIDDLE",
		3: "PRIORITY_HIGH",
	}
	ToDo_Info_Priority_value = map[string]int32{
		"PRIORITY_UNSPECIFIED":   0,
		"PRIORITY_LOW":           1,
		"PRIORITY_NORMAL_MIDDLE": 2,
		"PRIORITY_HIGH":          3,
	}
)

func (x ToDo_Info_Priority) Enum() *ToDo_Info_Priority {
	p := new(ToDo_Info_Priority)
	*p = x
	return p
}

func (x ToDo_Info_Priority) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ToDo_Info_Priority) Descriptor() protoreflect.EnumDescriptor {
	return file_todo_proto_enumTypes[0].Descriptor()
}

func (ToDo_Info_Priority) Type() protoreflect.EnumType {
	return &file_todo_proto_enumTypes[0]
}

func (x ToDo_Info_Priority) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ToDo_Info_Priority.Descriptor instead.
func (ToDo_Info_Priority) EnumDescriptor() ([]byte, []int) {
	return file_todo_proto_rawDescGZIP(), []int{5, 0, 0}
}

type ToDo_Info_Status int32

const (
	ToDo_Info_STATUS_UNSPECIFIED ToDo_Info_Status = 0
	ToDo_Info_STATUS_NOT_YET     ToDo_Info_Status = 1
	ToDo_Info_STATUS_DOING       ToDo_Info_Status = 2
	ToDo_Info_STATUS_DONE        ToDo_Info_Status = 3
)

// Enum value maps for ToDo_Info_Status.
var (
	ToDo_Info_Status_name = map[int32]string{
		0: "STATUS_UNSPECIFIED",
		1: "STATUS_NOT_YET",
		2: "STATUS_DOING",
		3: "STATUS_DONE",
	}
	ToDo_Info_Status_value = map[string]int32{
		"STATUS_UNSPECIFIED": 0,
		"STATUS_NOT_YET":     1,
		"STATUS_DOING":       2,
		"STATUS_DONE":        3,
	}
)

func (x ToDo_Info_Status) Enum() *ToDo_Info_Status {
	p := new(ToDo_Info_Status)
	*p = x
	return p
}

func (x ToDo_Info_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ToDo_Info_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_todo_proto_enumTypes[1].Descriptor()
}

func (ToDo_Info_Status) Type() protoreflect.EnumType {
	return &file_todo_proto_enumTypes[1]
}

func (x ToDo_Info_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ToDo_Info_Status.Descriptor instead.
func (ToDo_Info_Status) EnumDescriptor() ([]byte, []int) {
	return file_todo_proto_rawDescGZIP(), []int{5, 0, 1}
}

type ToDoListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ToDoListRequest) Reset() {
	*x = ToDoListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ToDoListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ToDoListRequest) ProtoMessage() {}

func (x *ToDoListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_todo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ToDoListRequest.ProtoReflect.Descriptor instead.
func (*ToDoListRequest) Descriptor() ([]byte, []int) {
	return file_todo_proto_rawDescGZIP(), []int{0}
}

type ToDoListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Todos []*ToDo `protobuf:"bytes,1,rep,name=todos,proto3" json:"todos,omitempty"`
}

func (x *ToDoListResponse) Reset() {
	*x = ToDoListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ToDoListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ToDoListResponse) ProtoMessage() {}

func (x *ToDoListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_todo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ToDoListResponse.ProtoReflect.Descriptor instead.
func (*ToDoListResponse) Descriptor() ([]byte, []int) {
	return file_todo_proto_rawDescGZIP(), []int{1}
}

func (x *ToDoListResponse) GetTodos() []*ToDo {
	if x != nil {
		return x.Todos
	}
	return nil
}

type ToDoGetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ToDoGetRequest) Reset() {
	*x = ToDoGetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todo_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ToDoGetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ToDoGetRequest) ProtoMessage() {}

func (x *ToDoGetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_todo_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ToDoGetRequest.ProtoReflect.Descriptor instead.
func (*ToDoGetRequest) Descriptor() ([]byte, []int) {
	return file_todo_proto_rawDescGZIP(), []int{2}
}

func (x *ToDoGetRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type ToDoAddRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *ToDo_Info `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *ToDoAddRequest) Reset() {
	*x = ToDoAddRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todo_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ToDoAddRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ToDoAddRequest) ProtoMessage() {}

func (x *ToDoAddRequest) ProtoReflect() protoreflect.Message {
	mi := &file_todo_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ToDoAddRequest.ProtoReflect.Descriptor instead.
func (*ToDoAddRequest) Descriptor() ([]byte, []int) {
	return file_todo_proto_rawDescGZIP(), []int{3}
}

func (x *ToDoAddRequest) GetInfo() *ToDo_Info {
	if x != nil {
		return x.Info
	}
	return nil
}

type ToDoDeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ToDoDeleteRequest) Reset() {
	*x = ToDoDeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todo_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ToDoDeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ToDoDeleteRequest) ProtoMessage() {}

func (x *ToDoDeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_todo_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ToDoDeleteRequest.ProtoReflect.Descriptor instead.
func (*ToDoDeleteRequest) Descriptor() ([]byte, []int) {
	return file_todo_proto_rawDescGZIP(), []int{4}
}

func (x *ToDoDeleteRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type ToDo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int32      `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Info *ToDo_Info `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *ToDo) Reset() {
	*x = ToDo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todo_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ToDo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ToDo) ProtoMessage() {}

func (x *ToDo) ProtoReflect() protoreflect.Message {
	mi := &file_todo_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ToDo.ProtoReflect.Descriptor instead.
func (*ToDo) Descriptor() ([]byte, []int) {
	return file_todo_proto_rawDescGZIP(), []int{5}
}

func (x *ToDo) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ToDo) GetInfo() *ToDo_Info {
	if x != nil {
		return x.Info
	}
	return nil
}

type ToDo_Info struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title       string                 `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description string                 `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Limit       *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=limit,proto3" json:"limit,omitempty"`
	Priority    ToDo_Info_Priority     `protobuf:"varint,4,opt,name=priority,proto3,enum=api.todo.ToDo_Info_Priority" json:"priority,omitempty"`
	Status      ToDo_Info_Status       `protobuf:"varint,5,opt,name=status,proto3,enum=api.todo.ToDo_Info_Status" json:"status,omitempty"`
	Labels      []string               `protobuf:"bytes,6,rep,name=labels,proto3" json:"labels,omitempty"`
}

func (x *ToDo_Info) Reset() {
	*x = ToDo_Info{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todo_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ToDo_Info) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ToDo_Info) ProtoMessage() {}

func (x *ToDo_Info) ProtoReflect() protoreflect.Message {
	mi := &file_todo_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ToDo_Info.ProtoReflect.Descriptor instead.
func (*ToDo_Info) Descriptor() ([]byte, []int) {
	return file_todo_proto_rawDescGZIP(), []int{5, 0}
}

func (x *ToDo_Info) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *ToDo_Info) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ToDo_Info) GetLimit() *timestamppb.Timestamp {
	if x != nil {
		return x.Limit
	}
	return nil
}

func (x *ToDo_Info) GetPriority() ToDo_Info_Priority {
	if x != nil {
		return x.Priority
	}
	return ToDo_Info_PRIORITY_UNSPECIFIED
}

func (x *ToDo_Info) GetStatus() ToDo_Info_Status {
	if x != nil {
		return x.Status
	}
	return ToDo_Info_STATUS_UNSPECIFIED
}

func (x *ToDo_Info) GetLabels() []string {
	if x != nil {
		return x.Labels
	}
	return nil
}

var File_todo_proto protoreflect.FileDescriptor

var file_todo_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x61, 0x70,
	0x69, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x11, 0x0a, 0x0f, 0x54, 0x6f, 0x44, 0x6f, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x38, 0x0a, 0x10, 0x54, 0x6f,
	0x44, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24,
	0x0a, 0x05, 0x74, 0x6f, 0x64, 0x6f, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x54, 0x6f, 0x44, 0x6f, 0x52, 0x05, 0x74,
	0x6f, 0x64, 0x6f, 0x73, 0x22, 0x20, 0x0a, 0x0e, 0x54, 0x6f, 0x44, 0x6f, 0x47, 0x65, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x39, 0x0a, 0x0e, 0x54, 0x6f, 0x44, 0x6f, 0x41, 0x64,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x6f, 0x64,
	0x6f, 0x2e, 0x54, 0x6f, 0x44, 0x6f, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66,
	0x6f, 0x22, 0x23, 0x0a, 0x11, 0x54, 0x6f, 0x44, 0x6f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0xf8, 0x03, 0x0a, 0x04, 0x54, 0x6f, 0x44, 0x6f, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x27, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x54, 0x6f, 0x44, 0x6f, 0x2e, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x1a, 0xb6, 0x03, 0x0a, 0x04, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x30, 0x0a, 0x05, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x38, 0x0a, 0x08, 0x70,
	0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x54, 0x6f, 0x44, 0x6f, 0x2e, 0x49, 0x6e,
	0x66, 0x6f, 0x2e, 0x50, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x52, 0x08, 0x70, 0x72, 0x69,
	0x6f, 0x72, 0x69, 0x74, 0x79, 0x12, 0x32, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x6f, 0x64, 0x6f,
	0x2e, 0x54, 0x6f, 0x44, 0x6f, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x61, 0x62,
	0x65, 0x6c, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c,
	0x73, 0x22, 0x65, 0x0a, 0x08, 0x50, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x12, 0x18, 0x0a,
	0x14, 0x50, 0x52, 0x49, 0x4f, 0x52, 0x49, 0x54, 0x59, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x50, 0x52, 0x49, 0x4f, 0x52,
	0x49, 0x54, 0x59, 0x5f, 0x4c, 0x4f, 0x57, 0x10, 0x01, 0x12, 0x1a, 0x0a, 0x16, 0x50, 0x52, 0x49,
	0x4f, 0x52, 0x49, 0x54, 0x59, 0x5f, 0x4e, 0x4f, 0x52, 0x4d, 0x41, 0x4c, 0x5f, 0x4d, 0x49, 0x44,
	0x44, 0x4c, 0x45, 0x10, 0x02, 0x12, 0x11, 0x0a, 0x0d, 0x50, 0x52, 0x49, 0x4f, 0x52, 0x49, 0x54,
	0x59, 0x5f, 0x48, 0x49, 0x47, 0x48, 0x10, 0x03, 0x22, 0x57, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x16, 0x0a, 0x12, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53,
	0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x53, 0x54,
	0x41, 0x54, 0x55, 0x53, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x59, 0x45, 0x54, 0x10, 0x01, 0x12, 0x10,
	0x0a, 0x0c, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x44, 0x4f, 0x49, 0x4e, 0x47, 0x10, 0x02,
	0x12, 0x0f, 0x0a, 0x0b, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x44, 0x4f, 0x4e, 0x45, 0x10,
	0x03, 0x32, 0xb3, 0x02, 0x0a, 0x0b, 0x54, 0x6f, 0x44, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x3d, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x54, 0x6f, 0x44, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e,
	0x54, 0x6f, 0x44, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x2f, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x18, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x6f,
	0x64, 0x6f, 0x2e, 0x54, 0x6f, 0x44, 0x6f, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x0e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x54, 0x6f, 0x44,
	0x6f, 0x12, 0x3b, 0x0a, 0x03, 0x41, 0x64, 0x64, 0x12, 0x18, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74,
	0x6f, 0x64, 0x6f, 0x2e, 0x54, 0x6f, 0x44, 0x6f, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x54, 0x6f,
	0x44, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34,
	0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x0e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74,
	0x6f, 0x64, 0x6f, 0x2e, 0x54, 0x6f, 0x44, 0x6f, 0x1a, 0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74,
	0x6f, 0x64, 0x6f, 0x2e, 0x54, 0x6f, 0x44, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x1b,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x54, 0x6f, 0x44, 0x6f, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x54, 0x6f, 0x44, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x26, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x6f, 0x69, 0x53, 0x61, 0x74, 0x6f, 0x6e, 0x61, 0x6b,
	0x61, 0x2f, 0x74, 0x6f, 0x64, 0x6f, 0x2d, 0x61, 0x70, 0x70, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_todo_proto_rawDescOnce sync.Once
	file_todo_proto_rawDescData = file_todo_proto_rawDesc
)

func file_todo_proto_rawDescGZIP() []byte {
	file_todo_proto_rawDescOnce.Do(func() {
		file_todo_proto_rawDescData = protoimpl.X.CompressGZIP(file_todo_proto_rawDescData)
	})
	return file_todo_proto_rawDescData
}

var file_todo_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_todo_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_todo_proto_goTypes = []interface{}{
	(ToDo_Info_Priority)(0),       // 0: api.todo.ToDo.Info.Priority
	(ToDo_Info_Status)(0),         // 1: api.todo.ToDo.Info.Status
	(*ToDoListRequest)(nil),       // 2: api.todo.ToDoListRequest
	(*ToDoListResponse)(nil),      // 3: api.todo.ToDoListResponse
	(*ToDoGetRequest)(nil),        // 4: api.todo.ToDoGetRequest
	(*ToDoAddRequest)(nil),        // 5: api.todo.ToDoAddRequest
	(*ToDoDeleteRequest)(nil),     // 6: api.todo.ToDoDeleteRequest
	(*ToDo)(nil),                  // 7: api.todo.ToDo
	(*ToDo_Info)(nil),             // 8: api.todo.ToDo.Info
	(*timestamppb.Timestamp)(nil), // 9: google.protobuf.Timestamp
}
var file_todo_proto_depIdxs = []int32{
	7,  // 0: api.todo.ToDoListResponse.todos:type_name -> api.todo.ToDo
	8,  // 1: api.todo.ToDoAddRequest.info:type_name -> api.todo.ToDo.Info
	8,  // 2: api.todo.ToDo.info:type_name -> api.todo.ToDo.Info
	9,  // 3: api.todo.ToDo.Info.limit:type_name -> google.protobuf.Timestamp
	0,  // 4: api.todo.ToDo.Info.priority:type_name -> api.todo.ToDo.Info.Priority
	1,  // 5: api.todo.ToDo.Info.status:type_name -> api.todo.ToDo.Info.Status
	2,  // 6: api.todo.ToDoService.List:input_type -> api.todo.ToDoListRequest
	4,  // 7: api.todo.ToDoService.Get:input_type -> api.todo.ToDoGetRequest
	5,  // 8: api.todo.ToDoService.Add:input_type -> api.todo.ToDoAddRequest
	7,  // 9: api.todo.ToDoService.Update:input_type -> api.todo.ToDo
	6,  // 10: api.todo.ToDoService.Delete:input_type -> api.todo.ToDoDeleteRequest
	3,  // 11: api.todo.ToDoService.List:output_type -> api.todo.ToDoListResponse
	7,  // 12: api.todo.ToDoService.Get:output_type -> api.todo.ToDo
	3,  // 13: api.todo.ToDoService.Add:output_type -> api.todo.ToDoListResponse
	3,  // 14: api.todo.ToDoService.Update:output_type -> api.todo.ToDoListResponse
	3,  // 15: api.todo.ToDoService.Delete:output_type -> api.todo.ToDoListResponse
	11, // [11:16] is the sub-list for method output_type
	6,  // [6:11] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_todo_proto_init() }
func file_todo_proto_init() {
	if File_todo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_todo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ToDoListRequest); i {
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
		file_todo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ToDoListResponse); i {
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
		file_todo_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ToDoGetRequest); i {
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
		file_todo_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ToDoAddRequest); i {
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
		file_todo_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ToDoDeleteRequest); i {
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
		file_todo_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ToDo); i {
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
		file_todo_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ToDo_Info); i {
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
			RawDescriptor: file_todo_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_todo_proto_goTypes,
		DependencyIndexes: file_todo_proto_depIdxs,
		EnumInfos:         file_todo_proto_enumTypes,
		MessageInfos:      file_todo_proto_msgTypes,
	}.Build()
	File_todo_proto = out.File
	file_todo_proto_rawDesc = nil
	file_todo_proto_goTypes = nil
	file_todo_proto_depIdxs = nil
}
