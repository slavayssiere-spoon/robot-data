// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.15.8
// source: ActionWithInteractionResponse.proto

package robot_data

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

type ActionWithInteractionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	/////////////////////////////////////// add by robot-data service ///////////////////////
	// @inject_tag: json:"timestamp"
	Publishtime string `protobuf:"bytes,1,opt,name=publishtime,proto3" json:"timestamp"`
	// @inject_tag: json:"-"
	Identity string `protobuf:"bytes,2,opt,name=identity,proto3" json:"-"`
	// @inject_tag: json:"-"
	Groups []string `protobuf:"bytes,3,rep,name=groups,proto3" json:"-"`
	// @inject_tag: json:"-"
	Paths []string `protobuf:"bytes,4,rep,name=paths,proto3" json:"-"`
	/////////////////////////////////////// add SpoonExportableData ///////////////////////
	// @inject_tag: json:"DataType"
	DataType string `protobuf:"bytes,5,opt,name=DataType,proto3" json:"DataType"`
	// @inject_tag: json:"ID"
	ID string `protobuf:"bytes,6,opt,name=ID,proto3" json:"ID"`
	// @inject_tag: json:"robotInfo"
	RobotInfo *RobotInfo `protobuf:"bytes,7,opt,name=RobotInfo,proto3" json:"robotInfo"`
	// @inject_tag: json:"name"
	Name string `protobuf:"bytes,8,opt,name=Name,proto3" json:"name"`
	// @inject_tag: json:"interactionStatus"
	InteractionStatus string `protobuf:"bytes,9,opt,name=InteractionStatus,proto3" json:"interactionStatus"`
	// @inject_tag: json:"userResponse"
	UserResponse *UserInteractionResponse `protobuf:"bytes,10,opt,name=UserResponse,proto3" json:"userResponse"`
	// @inject_tag: json:"userInput"
	UserInput *UserInteractionInput `protobuf:"bytes,11,opt,name=UserInput,proto3" json:"userInput"`
	// @inject_tag: json:"totalDisplayedTime"
	TotalDisplayedTime float32 `protobuf:"fixed32,12,opt,name=TotalDisplayedTime,proto3" json:"totalDisplayedTime"`
}

func (x *ActionWithInteractionResponse) Reset() {
	*x = ActionWithInteractionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ActionWithInteractionResponse_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActionWithInteractionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActionWithInteractionResponse) ProtoMessage() {}

func (x *ActionWithInteractionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ActionWithInteractionResponse_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActionWithInteractionResponse.ProtoReflect.Descriptor instead.
func (*ActionWithInteractionResponse) Descriptor() ([]byte, []int) {
	return file_ActionWithInteractionResponse_proto_rawDescGZIP(), []int{0}
}

func (x *ActionWithInteractionResponse) GetPublishtime() string {
	if x != nil {
		return x.Publishtime
	}
	return ""
}

func (x *ActionWithInteractionResponse) GetIdentity() string {
	if x != nil {
		return x.Identity
	}
	return ""
}

func (x *ActionWithInteractionResponse) GetGroups() []string {
	if x != nil {
		return x.Groups
	}
	return nil
}

func (x *ActionWithInteractionResponse) GetPaths() []string {
	if x != nil {
		return x.Paths
	}
	return nil
}

func (x *ActionWithInteractionResponse) GetDataType() string {
	if x != nil {
		return x.DataType
	}
	return ""
}

func (x *ActionWithInteractionResponse) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *ActionWithInteractionResponse) GetRobotInfo() *RobotInfo {
	if x != nil {
		return x.RobotInfo
	}
	return nil
}

func (x *ActionWithInteractionResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ActionWithInteractionResponse) GetInteractionStatus() string {
	if x != nil {
		return x.InteractionStatus
	}
	return ""
}

func (x *ActionWithInteractionResponse) GetUserResponse() *UserInteractionResponse {
	if x != nil {
		return x.UserResponse
	}
	return nil
}

func (x *ActionWithInteractionResponse) GetUserInput() *UserInteractionInput {
	if x != nil {
		return x.UserInput
	}
	return nil
}

func (x *ActionWithInteractionResponse) GetTotalDisplayedTime() float32 {
	if x != nil {
		return x.TotalDisplayedTime
	}
	return 0
}

type UserInteractionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"DataType"
	DataType string `protobuf:"bytes,1,opt,name=DataType,proto3" json:"DataType"`
	// @inject_tag: json:"ID"
	ID string `protobuf:"bytes,2,opt,name=ID,proto3" json:"ID"`
	// @inject_tag: json:"response_type"
	Type string `protobuf:"bytes,3,opt,name=Type,proto3" json:"response_type"`
	// @inject_tag: json:"concept"
	Concept *LocalizableConcept `protobuf:"bytes,4,opt,name=Concept,proto3" json:"concept"`
}

func (x *UserInteractionResponse) Reset() {
	*x = UserInteractionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ActionWithInteractionResponse_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserInteractionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInteractionResponse) ProtoMessage() {}

func (x *UserInteractionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ActionWithInteractionResponse_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInteractionResponse.ProtoReflect.Descriptor instead.
func (*UserInteractionResponse) Descriptor() ([]byte, []int) {
	return file_ActionWithInteractionResponse_proto_rawDescGZIP(), []int{1}
}

func (x *UserInteractionResponse) GetDataType() string {
	if x != nil {
		return x.DataType
	}
	return ""
}

func (x *UserInteractionResponse) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *UserInteractionResponse) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *UserInteractionResponse) GetConcept() *LocalizableConcept {
	if x != nil {
		return x.Concept
	}
	return nil
}

type LocalizableConcept struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"DataType"
	DataType string `protobuf:"bytes,1,opt,name=DataType,proto3" json:"DataType"`
	// @inject_tag: json:"ID"
	ID string `protobuf:"bytes,2,opt,name=ID,proto3" json:"ID"`
	// @inject_tag: json:"name"
	Name *LocalizableString `protobuf:"bytes,3,opt,name=Name,proto3" json:"name"`
}

func (x *LocalizableConcept) Reset() {
	*x = LocalizableConcept{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ActionWithInteractionResponse_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LocalizableConcept) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LocalizableConcept) ProtoMessage() {}

func (x *LocalizableConcept) ProtoReflect() protoreflect.Message {
	mi := &file_ActionWithInteractionResponse_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LocalizableConcept.ProtoReflect.Descriptor instead.
func (*LocalizableConcept) Descriptor() ([]byte, []int) {
	return file_ActionWithInteractionResponse_proto_rawDescGZIP(), []int{2}
}

func (x *LocalizableConcept) GetDataType() string {
	if x != nil {
		return x.DataType
	}
	return ""
}

func (x *LocalizableConcept) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *LocalizableConcept) GetName() *LocalizableString {
	if x != nil {
		return x.Name
	}
	return nil
}

type LocalizableString struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"DataType"
	DataType string `protobuf:"bytes,1,opt,name=DataType,proto3" json:"DataType"`
	// @inject_tag: json:"ID"
	ID string `protobuf:"bytes,2,opt,name=ID,proto3" json:"ID"`
	// @inject_tag: json:"content"
	Content string `protobuf:"bytes,3,opt,name=Content,proto3" json:"content"`
}

func (x *LocalizableString) Reset() {
	*x = LocalizableString{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ActionWithInteractionResponse_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LocalizableString) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LocalizableString) ProtoMessage() {}

func (x *LocalizableString) ProtoReflect() protoreflect.Message {
	mi := &file_ActionWithInteractionResponse_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LocalizableString.ProtoReflect.Descriptor instead.
func (*LocalizableString) Descriptor() ([]byte, []int) {
	return file_ActionWithInteractionResponse_proto_rawDescGZIP(), []int{3}
}

func (x *LocalizableString) GetDataType() string {
	if x != nil {
		return x.DataType
	}
	return ""
}

func (x *LocalizableString) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *LocalizableString) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type UserInteractionInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"DataType"
	DataType string `protobuf:"bytes,1,opt,name=DataType,proto3" json:"DataType"`
	// @inject_tag: json:"ID"
	ID string `protobuf:"bytes,2,opt,name=ID,proto3" json:"ID"`
	// @inject_tag: json:"input_source"
	Input string `protobuf:"bytes,3,opt,name=Input,proto3" json:"input_source"`
	// @inject_tag: json:"content"
	Content *LocalizableString `protobuf:"bytes,4,opt,name=Content,proto3" json:"content"`
	// @inject_tag: json:"speech"
	Speech *Speech `protobuf:"bytes,5,opt,name=Speech,proto3" json:"speech"`
	// @inject_tag: json:"matched"
	Matched bool `protobuf:"varint,6,opt,name=Matched,proto3" json:"matched"`
}

func (x *UserInteractionInput) Reset() {
	*x = UserInteractionInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ActionWithInteractionResponse_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserInteractionInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInteractionInput) ProtoMessage() {}

func (x *UserInteractionInput) ProtoReflect() protoreflect.Message {
	mi := &file_ActionWithInteractionResponse_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInteractionInput.ProtoReflect.Descriptor instead.
func (*UserInteractionInput) Descriptor() ([]byte, []int) {
	return file_ActionWithInteractionResponse_proto_rawDescGZIP(), []int{4}
}

func (x *UserInteractionInput) GetDataType() string {
	if x != nil {
		return x.DataType
	}
	return ""
}

func (x *UserInteractionInput) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *UserInteractionInput) GetInput() string {
	if x != nil {
		return x.Input
	}
	return ""
}

func (x *UserInteractionInput) GetContent() *LocalizableString {
	if x != nil {
		return x.Content
	}
	return nil
}

func (x *UserInteractionInput) GetSpeech() *Speech {
	if x != nil {
		return x.Speech
	}
	return nil
}

func (x *UserInteractionInput) GetMatched() bool {
	if x != nil {
		return x.Matched
	}
	return false
}

var File_ActionWithInteractionResponse_proto protoreflect.FileDescriptor

var file_ActionWithInteractionResponse_proto_rawDesc = []byte{
	0x0a, 0x23, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x57, 0x69, 0x74, 0x68, 0x49, 0x6e, 0x74, 0x65,
	0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x5f, 0x64, 0x61, 0x74,
	0x61, 0x1a, 0x0e, 0x62, 0x71, 0x5f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x0e, 0x62, 0x71, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x0a, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb2, 0x04,
	0x0a, 0x1d, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x57, 0x69, 0x74, 0x68, 0x49, 0x6e, 0x74, 0x65,
	0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x32, 0x0a, 0x0b, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x10, 0xea, 0x3f, 0x0d, 0x08, 0x01, 0x12, 0x09, 0x54, 0x49, 0x4d,
	0x45, 0x53, 0x54, 0x41, 0x4d, 0x50, 0x52, 0x0b, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x74,
	0x69, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x05, 0xea, 0x3f, 0x02, 0x08, 0x01, 0x52, 0x08, 0x69, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x12, 0x14,
	0x0a, 0x05, 0x70, 0x61, 0x74, 0x68, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x70,
	0x61, 0x74, 0x68, 0x73, 0x12, 0x21, 0x0a, 0x08, 0x44, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x05, 0xea, 0x3f, 0x02, 0x08, 0x01, 0x52, 0x08, 0x44,
	0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x12, 0x15, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x05, 0xea, 0x3f, 0x02, 0x08, 0x01, 0x52, 0x02, 0x49, 0x44, 0x12, 0x33,
	0x0a, 0x09, 0x52, 0x6f, 0x62, 0x6f, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x15, 0x2e, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x52,
	0x6f, 0x62, 0x6f, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x09, 0x52, 0x6f, 0x62, 0x6f, 0x74, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2c, 0x0a, 0x11, 0x49, 0x6e, 0x74, 0x65, 0x72,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x11, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x47, 0x0a, 0x0c, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x72, 0x6f,
	0x62, 0x6f, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x74,
	0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x52, 0x0c, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e,
	0x0a, 0x09, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x20, 0x2e, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e,
	0x70, 0x75, 0x74, 0x52, 0x09, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x2e,
	0x0a, 0x12, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x64,
	0x54, 0x69, 0x6d, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x02, 0x52, 0x12, 0x54, 0x6f, 0x74, 0x61,
	0x6c, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x3a, 0x22,
	0xea, 0x3f, 0x1f, 0x0a, 0x1d, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x57, 0x69, 0x74, 0x68, 0x49,
	0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x93, 0x01, 0x0a, 0x17, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x74, 0x65, 0x72,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x44, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x44, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x79,
	0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x38,
	0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x63, 0x65, 0x70, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1e, 0x2e, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x4c, 0x6f, 0x63,
	0x61, 0x6c, 0x69, 0x7a, 0x61, 0x62, 0x6c, 0x65, 0x43, 0x6f, 0x6e, 0x63, 0x65, 0x70, 0x74, 0x52,
	0x07, 0x43, 0x6f, 0x6e, 0x63, 0x65, 0x70, 0x74, 0x22, 0x73, 0x0a, 0x12, 0x4c, 0x6f, 0x63, 0x61,
	0x6c, 0x69, 0x7a, 0x61, 0x62, 0x6c, 0x65, 0x43, 0x6f, 0x6e, 0x63, 0x65, 0x70, 0x74, 0x12, 0x1a,
	0x0a, 0x08, 0x44, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x44, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x31, 0x0a, 0x04, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x72, 0x6f, 0x62, 0x6f, 0x74,
	0x5f, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x62, 0x6c,
	0x65, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x59, 0x0a,
	0x11, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x12, 0x1a, 0x0a, 0x08, 0x44, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x44, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x18,
	0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0xd7, 0x01, 0x0a, 0x14, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x70, 0x75,
	0x74, 0x12, 0x1a, 0x0a, 0x08, 0x44, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x44, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x14, 0x0a,
	0x05, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x49, 0x6e,
	0x70, 0x75, 0x74, 0x12, 0x37, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x5f, 0x64, 0x61, 0x74,
	0x61, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x52, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x2a, 0x0a, 0x06,
	0x53, 0x70, 0x65, 0x65, 0x63, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x72,
	0x6f, 0x62, 0x6f, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x53, 0x70, 0x65, 0x65, 0x63, 0x68,
	0x52, 0x06, 0x53, 0x70, 0x65, 0x65, 0x63, 0x68, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x61, 0x74, 0x63,
	0x68, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x4d, 0x61, 0x74, 0x63, 0x68,
	0x65, 0x64, 0x42, 0x2a, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x73, 0x6c, 0x61, 0x76, 0x61, 0x79, 0x73, 0x73, 0x69, 0x65, 0x72, 0x65, 0x2d, 0x73, 0x70,
	0x6f, 0x6f, 0x6e, 0x2f, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x2d, 0x64, 0x61, 0x74, 0x61, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ActionWithInteractionResponse_proto_rawDescOnce sync.Once
	file_ActionWithInteractionResponse_proto_rawDescData = file_ActionWithInteractionResponse_proto_rawDesc
)

func file_ActionWithInteractionResponse_proto_rawDescGZIP() []byte {
	file_ActionWithInteractionResponse_proto_rawDescOnce.Do(func() {
		file_ActionWithInteractionResponse_proto_rawDescData = protoimpl.X.CompressGZIP(file_ActionWithInteractionResponse_proto_rawDescData)
	})
	return file_ActionWithInteractionResponse_proto_rawDescData
}

var file_ActionWithInteractionResponse_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_ActionWithInteractionResponse_proto_goTypes = []interface{}{
	(*ActionWithInteractionResponse)(nil), // 0: robot_data.ActionWithInteractionResponse
	(*UserInteractionResponse)(nil),       // 1: robot_data.UserInteractionResponse
	(*LocalizableConcept)(nil),            // 2: robot_data.LocalizableConcept
	(*LocalizableString)(nil),             // 3: robot_data.LocalizableString
	(*UserInteractionInput)(nil),          // 4: robot_data.UserInteractionInput
	(*RobotInfo)(nil),                     // 5: robot_data.RobotInfo
	(*Speech)(nil),                        // 6: robot_data.Speech
}
var file_ActionWithInteractionResponse_proto_depIdxs = []int32{
	5, // 0: robot_data.ActionWithInteractionResponse.RobotInfo:type_name -> robot_data.RobotInfo
	1, // 1: robot_data.ActionWithInteractionResponse.UserResponse:type_name -> robot_data.UserInteractionResponse
	4, // 2: robot_data.ActionWithInteractionResponse.UserInput:type_name -> robot_data.UserInteractionInput
	2, // 3: robot_data.UserInteractionResponse.Concept:type_name -> robot_data.LocalizableConcept
	3, // 4: robot_data.LocalizableConcept.Name:type_name -> robot_data.LocalizableString
	3, // 5: robot_data.UserInteractionInput.Content:type_name -> robot_data.LocalizableString
	6, // 6: robot_data.UserInteractionInput.Speech:type_name -> robot_data.Speech
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_ActionWithInteractionResponse_proto_init() }
func file_ActionWithInteractionResponse_proto_init() {
	if File_ActionWithInteractionResponse_proto != nil {
		return
	}
	file_bq_table_proto_init()
	file_bq_field_proto_init()
	file_data_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_ActionWithInteractionResponse_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActionWithInteractionResponse); i {
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
		file_ActionWithInteractionResponse_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserInteractionResponse); i {
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
		file_ActionWithInteractionResponse_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LocalizableConcept); i {
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
		file_ActionWithInteractionResponse_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LocalizableString); i {
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
		file_ActionWithInteractionResponse_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserInteractionInput); i {
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
			RawDescriptor: file_ActionWithInteractionResponse_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ActionWithInteractionResponse_proto_goTypes,
		DependencyIndexes: file_ActionWithInteractionResponse_proto_depIdxs,
		MessageInfos:      file_ActionWithInteractionResponse_proto_msgTypes,
	}.Build()
	File_ActionWithInteractionResponse_proto = out.File
	file_ActionWithInteractionResponse_proto_rawDesc = nil
	file_ActionWithInteractionResponse_proto_goTypes = nil
	file_ActionWithInteractionResponse_proto_depIdxs = nil
}
