// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v3.11.3
// source: idiom/v1/idiom.proto

package v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/descriptorpb"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Idiom_State int32

const (
	// 未知状态
	Idiom_INVALID Idiom_State = 0
	// 审核中
	Idiom_AUDITING Idiom_State = 1
	// 已上线
	Idiom_ONLINE Idiom_State = 2
	// 已下线
	Idiom_OFFLINE Idiom_State = 3
)

// Enum value maps for Idiom_State.
var (
	Idiom_State_name = map[int32]string{
		0: "INVALID",
		1: "AUDITING",
		2: "ONLINE",
		3: "OFFLINE",
	}
	Idiom_State_value = map[string]int32{
		"INVALID":  0,
		"AUDITING": 1,
		"ONLINE":   2,
		"OFFLINE":  3,
	}
)

func (x Idiom_State) Enum() *Idiom_State {
	p := new(Idiom_State)
	*p = x
	return p
}

func (x Idiom_State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Idiom_State) Descriptor() protoreflect.EnumDescriptor {
	return file_idiom_v1_idiom_proto_enumTypes[0].Descriptor()
}

func (Idiom_State) Type() protoreflect.EnumType {
	return &file_idiom_v1_idiom_proto_enumTypes[0]
}

func (x Idiom_State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Idiom_State.Descriptor instead.
func (Idiom_State) EnumDescriptor() ([]byte, []int) {
	return file_idiom_v1_idiom_proto_rawDescGZIP(), []int{3, 0}
}

type Idiom_Degree int32

const (
	// 未知难度
	Idiom_UNDECIDED Idiom_Degree = 0
	// 简单
	Idiom_EASY Idiom_Degree = 1
	// 中等
	Idiom_MEDIUM Idiom_Degree = 2
	// 困难
	Idiom_HARD Idiom_Degree = 3
)

// Enum value maps for Idiom_Degree.
var (
	Idiom_Degree_name = map[int32]string{
		0: "UNDECIDED",
		1: "EASY",
		2: "MEDIUM",
		3: "HARD",
	}
	Idiom_Degree_value = map[string]int32{
		"UNDECIDED": 0,
		"EASY":      1,
		"MEDIUM":    2,
		"HARD":      3,
	}
)

func (x Idiom_Degree) Enum() *Idiom_Degree {
	p := new(Idiom_Degree)
	*p = x
	return p
}

func (x Idiom_Degree) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Idiom_Degree) Descriptor() protoreflect.EnumDescriptor {
	return file_idiom_v1_idiom_proto_enumTypes[1].Descriptor()
}

func (Idiom_Degree) Type() protoreflect.EnumType {
	return &file_idiom_v1_idiom_proto_enumTypes[1]
}

func (x Idiom_Degree) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Idiom_Degree.Descriptor instead.
func (Idiom_Degree) EnumDescriptor() ([]byte, []int) {
	return file_idiom_v1_idiom_proto_rawDescGZIP(), []int{3, 1}
}

type SetIdiomBatchReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Idioms        []*Idiom               `protobuf:"bytes,1,rep,name=idioms,proto3" json:"idioms,omitempty"`
	Total         int32                  `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SetIdiomBatchReq) Reset() {
	*x = SetIdiomBatchReq{}
	mi := &file_idiom_v1_idiom_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SetIdiomBatchReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetIdiomBatchReq) ProtoMessage() {}

func (x *SetIdiomBatchReq) ProtoReflect() protoreflect.Message {
	mi := &file_idiom_v1_idiom_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetIdiomBatchReq.ProtoReflect.Descriptor instead.
func (*SetIdiomBatchReq) Descriptor() ([]byte, []int) {
	return file_idiom_v1_idiom_proto_rawDescGZIP(), []int{0}
}

func (x *SetIdiomBatchReq) GetIdioms() []*Idiom {
	if x != nil {
		return x.Idioms
	}
	return nil
}

func (x *SetIdiomBatchReq) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

type GetIdiomReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" form:"id"` // @inject_tag: form:"id"
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetIdiomReq) Reset() {
	*x = GetIdiomReq{}
	mi := &file_idiom_v1_idiom_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetIdiomReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetIdiomReq) ProtoMessage() {}

func (x *GetIdiomReq) ProtoReflect() protoreflect.Message {
	mi := &file_idiom_v1_idiom_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetIdiomReq.ProtoReflect.Descriptor instead.
func (*GetIdiomReq) Descriptor() ([]byte, []int) {
	return file_idiom_v1_idiom_proto_rawDescGZIP(), []int{1}
}

func (x *GetIdiomReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetIdiomResp struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Idiom         *Idiom                 `protobuf:"bytes,1,opt,name=idiom,proto3" json:"idiom,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetIdiomResp) Reset() {
	*x = GetIdiomResp{}
	mi := &file_idiom_v1_idiom_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetIdiomResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetIdiomResp) ProtoMessage() {}

func (x *GetIdiomResp) ProtoReflect() protoreflect.Message {
	mi := &file_idiom_v1_idiom_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetIdiomResp.ProtoReflect.Descriptor instead.
func (*GetIdiomResp) Descriptor() ([]byte, []int) {
	return file_idiom_v1_idiom_proto_rawDescGZIP(), []int{2}
}

func (x *GetIdiomResp) GetIdiom() *Idiom {
	if x != nil {
		return x.Idiom
	}
	return nil
}

type Idiom struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	Id    int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// 唯一标识（例：HT_001）
	Iid string `protobuf:"bytes,2,opt,name=iid,proto3" json:"iid,omitempty"`
	// 成语
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// 图片地址
	Image string `protobuf:"bytes,4,opt,name=image,proto3" json:"image,omitempty"`
	// 难度
	Difficulty Idiom_Degree `protobuf:"varint,5,opt,name=difficulty,proto3,enum=idiom.v1.Idiom_Degree" json:"difficulty,omitempty"`
	// 创建者
	Creator string `protobuf:"bytes,6,opt,name=creator,proto3" json:"creator,omitempty"`
	// 状态（审核中/已上线/已删除）
	State         Idiom_State `protobuf:"varint,7,opt,name=state,proto3,enum=idiom.v1.Idiom_State" json:"state,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Idiom) Reset() {
	*x = Idiom{}
	mi := &file_idiom_v1_idiom_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Idiom) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Idiom) ProtoMessage() {}

func (x *Idiom) ProtoReflect() protoreflect.Message {
	mi := &file_idiom_v1_idiom_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Idiom.ProtoReflect.Descriptor instead.
func (*Idiom) Descriptor() ([]byte, []int) {
	return file_idiom_v1_idiom_proto_rawDescGZIP(), []int{3}
}

func (x *Idiom) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Idiom) GetIid() string {
	if x != nil {
		return x.Iid
	}
	return ""
}

func (x *Idiom) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Idiom) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *Idiom) GetDifficulty() Idiom_Degree {
	if x != nil {
		return x.Difficulty
	}
	return Idiom_UNDECIDED
}

func (x *Idiom) GetCreator() string {
	if x != nil {
		return x.Creator
	}
	return ""
}

func (x *Idiom) GetState() Idiom_State {
	if x != nil {
		return x.State
	}
	return Idiom_INVALID
}

var File_idiom_v1_idiom_proto protoreflect.FileDescriptor

var file_idiom_v1_idiom_proto_rawDesc = string([]byte{
	0x0a, 0x14, 0x69, 0x64, 0x69, 0x6f, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x64, 0x69, 0x6f, 0x6d,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x69, 0x64, 0x69, 0x6f, 0x6d, 0x2e, 0x76, 0x31,
	0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x51, 0x0a, 0x10, 0x53, 0x65, 0x74, 0x49, 0x64, 0x69,
	0x6f, 0x6d, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x12, 0x27, 0x0a, 0x06, 0x69, 0x64,
	0x69, 0x6f, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x69, 0x64, 0x69,
	0x6f, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x64, 0x69, 0x6f, 0x6d, 0x52, 0x06, 0x69, 0x64, 0x69,
	0x6f, 0x6d, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x1d, 0x0a, 0x0b, 0x47, 0x65, 0x74,
	0x49, 0x64, 0x69, 0x6f, 0x6d, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x35, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x49,
	0x64, 0x69, 0x6f, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x12, 0x25, 0x0a, 0x05, 0x69, 0x64, 0x69, 0x6f,
	0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x69, 0x64, 0x69, 0x6f, 0x6d, 0x2e,
	0x76, 0x31, 0x2e, 0x49, 0x64, 0x69, 0x6f, 0x6d, 0x52, 0x05, 0x69, 0x64, 0x69, 0x6f, 0x6d, 0x22,
	0xd1, 0x02, 0x0a, 0x05, 0x49, 0x64, 0x69, 0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x69, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x36, 0x0a, 0x0a, 0x64, 0x69, 0x66, 0x66, 0x69, 0x63, 0x75,
	0x6c, 0x74, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x69, 0x64, 0x69, 0x6f,
	0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x64, 0x69, 0x6f, 0x6d, 0x2e, 0x44, 0x65, 0x67, 0x72, 0x65,
	0x65, 0x52, 0x0a, 0x64, 0x69, 0x66, 0x66, 0x69, 0x63, 0x75, 0x6c, 0x74, 0x79, 0x12, 0x21, 0x0a,
	0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07,
	0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72,
	0x12, 0x2b, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x15, 0x2e, 0x69, 0x64, 0x69, 0x6f, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x64, 0x69, 0x6f, 0x6d,
	0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x22, 0x3b, 0x0a,
	0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49,
	0x44, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x41, 0x55, 0x44, 0x49, 0x54, 0x49, 0x4e, 0x47, 0x10,
	0x01, 0x12, 0x0a, 0x0a, 0x06, 0x4f, 0x4e, 0x4c, 0x49, 0x4e, 0x45, 0x10, 0x02, 0x12, 0x0b, 0x0a,
	0x07, 0x4f, 0x46, 0x46, 0x4c, 0x49, 0x4e, 0x45, 0x10, 0x03, 0x22, 0x37, 0x0a, 0x06, 0x44, 0x65,
	0x67, 0x72, 0x65, 0x65, 0x12, 0x0d, 0x0a, 0x09, 0x55, 0x4e, 0x44, 0x45, 0x43, 0x49, 0x44, 0x45,
	0x44, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x45, 0x41, 0x53, 0x59, 0x10, 0x01, 0x12, 0x0a, 0x0a,
	0x06, 0x4d, 0x45, 0x44, 0x49, 0x55, 0x4d, 0x10, 0x02, 0x12, 0x08, 0x0a, 0x04, 0x48, 0x41, 0x52,
	0x44, 0x10, 0x03, 0x32, 0x8b, 0x01, 0x0a, 0x05, 0x69, 0x64, 0x69, 0x6f, 0x6d, 0x12, 0x45, 0x0a,
	0x0d, 0x53, 0x65, 0x74, 0x49, 0x64, 0x69, 0x6f, 0x6d, 0x42, 0x61, 0x74, 0x63, 0x68, 0x12, 0x1a,
	0x2e, 0x69, 0x64, 0x69, 0x6f, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x74, 0x49, 0x64, 0x69,
	0x6f, 0x6d, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x49, 0x64, 0x69, 0x6f, 0x6d,
	0x12, 0x15, 0x2e, 0x69, 0x64, 0x69, 0x6f, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x49,
	0x64, 0x69, 0x6f, 0x6d, 0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x69, 0x64, 0x69, 0x6f, 0x6d, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x64, 0x69, 0x6f, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x22,
	0x00, 0x42, 0x43, 0x0a, 0x17, 0x64, 0x65, 0x76, 0x2e, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x69, 0x64, 0x69, 0x6f, 0x6d, 0x2e, 0x76, 0x31, 0x42, 0x0c, 0x49, 0x64,
	0x69, 0x6f, 0x6d, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x56, 0x31, 0x50, 0x01, 0x5a, 0x18, 0x78, 0x69,
	0x75, 0x79, 0x69, 0x50, 0x72, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x69, 0x64, 0x69, 0x6f, 0x6d,
	0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_idiom_v1_idiom_proto_rawDescOnce sync.Once
	file_idiom_v1_idiom_proto_rawDescData []byte
)

func file_idiom_v1_idiom_proto_rawDescGZIP() []byte {
	file_idiom_v1_idiom_proto_rawDescOnce.Do(func() {
		file_idiom_v1_idiom_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_idiom_v1_idiom_proto_rawDesc), len(file_idiom_v1_idiom_proto_rawDesc)))
	})
	return file_idiom_v1_idiom_proto_rawDescData
}

var file_idiom_v1_idiom_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_idiom_v1_idiom_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_idiom_v1_idiom_proto_goTypes = []any{
	(Idiom_State)(0),         // 0: idiom.v1.Idiom.State
	(Idiom_Degree)(0),        // 1: idiom.v1.Idiom.Degree
	(*SetIdiomBatchReq)(nil), // 2: idiom.v1.SetIdiomBatchReq
	(*GetIdiomReq)(nil),      // 3: idiom.v1.GetIdiomReq
	(*GetIdiomResp)(nil),     // 4: idiom.v1.GetIdiomResp
	(*Idiom)(nil),            // 5: idiom.v1.Idiom
	(*emptypb.Empty)(nil),    // 6: google.protobuf.Empty
}
var file_idiom_v1_idiom_proto_depIdxs = []int32{
	5, // 0: idiom.v1.SetIdiomBatchReq.idioms:type_name -> idiom.v1.Idiom
	5, // 1: idiom.v1.GetIdiomResp.idiom:type_name -> idiom.v1.Idiom
	1, // 2: idiom.v1.Idiom.difficulty:type_name -> idiom.v1.Idiom.Degree
	0, // 3: idiom.v1.Idiom.state:type_name -> idiom.v1.Idiom.State
	2, // 4: idiom.v1.idiom.SetIdiomBatch:input_type -> idiom.v1.SetIdiomBatchReq
	3, // 5: idiom.v1.idiom.GetIdiom:input_type -> idiom.v1.GetIdiomReq
	6, // 6: idiom.v1.idiom.SetIdiomBatch:output_type -> google.protobuf.Empty
	4, // 7: idiom.v1.idiom.GetIdiom:output_type -> idiom.v1.GetIdiomResp
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_idiom_v1_idiom_proto_init() }
func file_idiom_v1_idiom_proto_init() {
	if File_idiom_v1_idiom_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_idiom_v1_idiom_proto_rawDesc), len(file_idiom_v1_idiom_proto_rawDesc)),
			NumEnums:      2,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_idiom_v1_idiom_proto_goTypes,
		DependencyIndexes: file_idiom_v1_idiom_proto_depIdxs,
		EnumInfos:         file_idiom_v1_idiom_proto_enumTypes,
		MessageInfos:      file_idiom_v1_idiom_proto_msgTypes,
	}.Build()
	File_idiom_v1_idiom_proto = out.File
	file_idiom_v1_idiom_proto_goTypes = nil
	file_idiom_v1_idiom_proto_depIdxs = nil
}
