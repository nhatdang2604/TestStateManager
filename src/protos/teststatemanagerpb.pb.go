// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.12.4
// source: protos/teststatemanagerpb.proto

package protos

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type StartTestRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int32 `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	TestId int32 `protobuf:"varint,2,opt,name=testId,proto3" json:"testId,omitempty"`
}

func (x *StartTestRequest) Reset() {
	*x = StartTestRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_teststatemanagerpb_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartTestRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartTestRequest) ProtoMessage() {}

func (x *StartTestRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_teststatemanagerpb_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartTestRequest.ProtoReflect.Descriptor instead.
func (*StartTestRequest) Descriptor() ([]byte, []int) {
	return file_protos_teststatemanagerpb_proto_rawDescGZIP(), []int{0}
}

func (x *StartTestRequest) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *StartTestRequest) GetTestId() int32 {
	if x != nil {
		return x.TestId
	}
	return 0
}

type StartTestResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TestAttemptId int32 `protobuf:"varint,1,opt,name=testAttemptId,proto3" json:"testAttemptId,omitempty"`
}

func (x *StartTestResponse) Reset() {
	*x = StartTestResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_teststatemanagerpb_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartTestResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartTestResponse) ProtoMessage() {}

func (x *StartTestResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_teststatemanagerpb_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartTestResponse.ProtoReflect.Descriptor instead.
func (*StartTestResponse) Descriptor() ([]byte, []int) {
	return file_protos_teststatemanagerpb_proto_rawDescGZIP(), []int{1}
}

func (x *StartTestResponse) GetTestAttemptId() int32 {
	if x != nil {
		return x.TestAttemptId
	}
	return 0
}

type GetInprogressTestRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TestAttemptId int32 `protobuf:"varint,1,opt,name=testAttemptId,proto3" json:"testAttemptId,omitempty"`
}

func (x *GetInprogressTestRequest) Reset() {
	*x = GetInprogressTestRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_teststatemanagerpb_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetInprogressTestRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInprogressTestRequest) ProtoMessage() {}

func (x *GetInprogressTestRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_teststatemanagerpb_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetInprogressTestRequest.ProtoReflect.Descriptor instead.
func (*GetInprogressTestRequest) Descriptor() ([]byte, []int) {
	return file_protos_teststatemanagerpb_proto_rawDescGZIP(), []int{2}
}

func (x *GetInprogressTestRequest) GetTestAttemptId() int32 {
	if x != nil {
		return x.TestAttemptId
	}
	return 0
}

type GetInprogressTestResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TimeRemainInSecond int32 `protobuf:"varint,1,opt,name=timeRemainInSecond,proto3" json:"timeRemainInSecond,omitempty"`
}

func (x *GetInprogressTestResponse) Reset() {
	*x = GetInprogressTestResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_teststatemanagerpb_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetInprogressTestResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInprogressTestResponse) ProtoMessage() {}

func (x *GetInprogressTestResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_teststatemanagerpb_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetInprogressTestResponse.ProtoReflect.Descriptor instead.
func (*GetInprogressTestResponse) Descriptor() ([]byte, []int) {
	return file_protos_teststatemanagerpb_proto_rawDescGZIP(), []int{3}
}

func (x *GetInprogressTestResponse) GetTimeRemainInSecond() int32 {
	if x != nil {
		return x.TimeRemainInSecond
	}
	return 0
}

type EndTestRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TestAttemptId int32 `protobuf:"varint,1,opt,name=testAttemptId,proto3" json:"testAttemptId,omitempty"`
}

func (x *EndTestRequest) Reset() {
	*x = EndTestRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_teststatemanagerpb_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EndTestRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EndTestRequest) ProtoMessage() {}

func (x *EndTestRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_teststatemanagerpb_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EndTestRequest.ProtoReflect.Descriptor instead.
func (*EndTestRequest) Descriptor() ([]byte, []int) {
	return file_protos_teststatemanagerpb_proto_rawDescGZIP(), []int{4}
}

func (x *EndTestRequest) GetTestAttemptId() int32 {
	if x != nil {
		return x.TestAttemptId
	}
	return 0
}

type EndTestResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ErrorCode int32  `protobuf:"varint,1,opt,name=errorCode,proto3" json:"errorCode,omitempty"`
	Message   string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *EndTestResponse) Reset() {
	*x = EndTestResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_teststatemanagerpb_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EndTestResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EndTestResponse) ProtoMessage() {}

func (x *EndTestResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_teststatemanagerpb_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EndTestResponse.ProtoReflect.Descriptor instead.
func (*EndTestResponse) Descriptor() ([]byte, []int) {
	return file_protos_teststatemanagerpb_proto_rawDescGZIP(), []int{5}
}

func (x *EndTestResponse) GetErrorCode() int32 {
	if x != nil {
		return x.ErrorCode
	}
	return 0
}

func (x *EndTestResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_protos_teststatemanagerpb_proto protoreflect.FileDescriptor

var file_protos_teststatemanagerpb_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x70, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x10, 0x74, 0x65, 0x73, 0x74, 0x73, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x22, 0x42, 0x0a, 0x10, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x65, 0x73, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x16, 0x0a, 0x06, 0x74, 0x65, 0x73, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x74, 0x65, 0x73, 0x74, 0x49, 0x64, 0x22, 0x39, 0x0a, 0x11, 0x53, 0x74, 0x61, 0x72, 0x74,
	0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x0d,
	0x74, 0x65, 0x73, 0x74, 0x41, 0x74, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0d, 0x74, 0x65, 0x73, 0x74, 0x41, 0x74, 0x74, 0x65, 0x6d, 0x70, 0x74,
	0x49, 0x64, 0x22, 0x40, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x70, 0x72, 0x6f, 0x67, 0x72,
	0x65, 0x73, 0x73, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24,
	0x0a, 0x0d, 0x74, 0x65, 0x73, 0x74, 0x41, 0x74, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x74, 0x65, 0x73, 0x74, 0x41, 0x74, 0x74, 0x65, 0x6d,
	0x70, 0x74, 0x49, 0x64, 0x22, 0x4b, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x70, 0x72, 0x6f,
	0x67, 0x72, 0x65, 0x73, 0x73, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x2e, 0x0a, 0x12, 0x74, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x6d, 0x61, 0x69, 0x6e, 0x49,
	0x6e, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x12, 0x74,
	0x69, 0x6d, 0x65, 0x52, 0x65, 0x6d, 0x61, 0x69, 0x6e, 0x49, 0x6e, 0x53, 0x65, 0x63, 0x6f, 0x6e,
	0x64, 0x22, 0x36, 0x0a, 0x0e, 0x45, 0x6e, 0x64, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x74, 0x65, 0x73, 0x74, 0x41, 0x74, 0x74, 0x65, 0x6d,
	0x70, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x74, 0x65, 0x73, 0x74,
	0x41, 0x74, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x49, 0x64, 0x22, 0x49, 0x0a, 0x0f, 0x45, 0x6e, 0x64,
	0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x32, 0xaf, 0x02, 0x0a, 0x13, 0x54, 0x65, 0x73, 0x74, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x6e, 0x0a, 0x11,
	0x47, 0x65, 0x74, 0x49, 0x6e, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x54, 0x65, 0x73,
	0x74, 0x12, 0x2a, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x73, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65,
	0x73, 0x73, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e,
	0x74, 0x65, 0x73, 0x74, 0x73, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x2e, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x54, 0x65,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x56, 0x0a, 0x09,
	0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x65, 0x73, 0x74, 0x12, 0x22, 0x2e, 0x74, 0x65, 0x73, 0x74,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x61,
	0x72, 0x74, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e,
	0x74, 0x65, 0x73, 0x74, 0x73, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x2e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x50, 0x0a, 0x07, 0x45, 0x6e, 0x64, 0x54, 0x65, 0x73, 0x74, 0x12,
	0x20, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x73, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x2e, 0x45, 0x6e, 0x64, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x21, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x73, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x2e, 0x45, 0x6e, 0x64, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_teststatemanagerpb_proto_rawDescOnce sync.Once
	file_protos_teststatemanagerpb_proto_rawDescData = file_protos_teststatemanagerpb_proto_rawDesc
)

func file_protos_teststatemanagerpb_proto_rawDescGZIP() []byte {
	file_protos_teststatemanagerpb_proto_rawDescOnce.Do(func() {
		file_protos_teststatemanagerpb_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_teststatemanagerpb_proto_rawDescData)
	})
	return file_protos_teststatemanagerpb_proto_rawDescData
}

var file_protos_teststatemanagerpb_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_protos_teststatemanagerpb_proto_goTypes = []interface{}{
	(*StartTestRequest)(nil),          // 0: teststatemanager.StartTestRequest
	(*StartTestResponse)(nil),         // 1: teststatemanager.StartTestResponse
	(*GetInprogressTestRequest)(nil),  // 2: teststatemanager.GetInprogressTestRequest
	(*GetInprogressTestResponse)(nil), // 3: teststatemanager.GetInprogressTestResponse
	(*EndTestRequest)(nil),            // 4: teststatemanager.EndTestRequest
	(*EndTestResponse)(nil),           // 5: teststatemanager.EndTestResponse
}
var file_protos_teststatemanagerpb_proto_depIdxs = []int32{
	2, // 0: teststatemanager.TestStateManagement.GetInprogressTest:input_type -> teststatemanager.GetInprogressTestRequest
	0, // 1: teststatemanager.TestStateManagement.StartTest:input_type -> teststatemanager.StartTestRequest
	4, // 2: teststatemanager.TestStateManagement.EndTest:input_type -> teststatemanager.EndTestRequest
	3, // 3: teststatemanager.TestStateManagement.GetInprogressTest:output_type -> teststatemanager.GetInprogressTestResponse
	1, // 4: teststatemanager.TestStateManagement.StartTest:output_type -> teststatemanager.StartTestResponse
	5, // 5: teststatemanager.TestStateManagement.EndTest:output_type -> teststatemanager.EndTestResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protos_teststatemanagerpb_proto_init() }
func file_protos_teststatemanagerpb_proto_init() {
	if File_protos_teststatemanagerpb_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_teststatemanagerpb_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartTestRequest); i {
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
		file_protos_teststatemanagerpb_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartTestResponse); i {
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
		file_protos_teststatemanagerpb_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetInprogressTestRequest); i {
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
		file_protos_teststatemanagerpb_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetInprogressTestResponse); i {
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
		file_protos_teststatemanagerpb_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EndTestRequest); i {
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
		file_protos_teststatemanagerpb_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EndTestResponse); i {
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
			RawDescriptor: file_protos_teststatemanagerpb_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_teststatemanagerpb_proto_goTypes,
		DependencyIndexes: file_protos_teststatemanagerpb_proto_depIdxs,
		MessageInfos:      file_protos_teststatemanagerpb_proto_msgTypes,
	}.Build()
	File_protos_teststatemanagerpb_proto = out.File
	file_protos_teststatemanagerpb_proto_rawDesc = nil
	file_protos_teststatemanagerpb_proto_goTypes = nil
	file_protos_teststatemanagerpb_proto_depIdxs = nil
}
