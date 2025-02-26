// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v3.11.3
// source: api/turtle/v1/turtle.proto

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

type Turtle_State int32

const (
	// 未知状态
	Turtle_INVALID Turtle_State = 0
	// 审核中
	Turtle_AUDITING Turtle_State = 1
	// 已上线
	Turtle_ONLINE Turtle_State = 2
	// 已下线
	Turtle_OFFLINE Turtle_State = 3
)

// Enum value maps for Turtle_State.
var (
	Turtle_State_name = map[int32]string{
		0: "INVALID",
		1: "AUDITING",
		2: "ONLINE",
		3: "OFFLINE",
	}
	Turtle_State_value = map[string]int32{
		"INVALID":  0,
		"AUDITING": 1,
		"ONLINE":   2,
		"OFFLINE":  3,
	}
)

func (x Turtle_State) Enum() *Turtle_State {
	p := new(Turtle_State)
	*p = x
	return p
}

func (x Turtle_State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Turtle_State) Descriptor() protoreflect.EnumDescriptor {
	return file_api_turtle_v1_turtle_proto_enumTypes[0].Descriptor()
}

func (Turtle_State) Type() protoreflect.EnumType {
	return &file_api_turtle_v1_turtle_proto_enumTypes[0]
}

func (x Turtle_State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Turtle_State.Descriptor instead.
func (Turtle_State) EnumDescriptor() ([]byte, []int) {
	return file_api_turtle_v1_turtle_proto_rawDescGZIP(), []int{3, 0}
}

type Turtle_Category int32

const (
	// 未知分类
	Turtle_UNKNOWN Turtle_Category = 0
	// 恐怖
	Turtle_HORROR Turtle_Category = 1
	// 推理
	Turtle_DETECTIVE Turtle_Category = 2
	// 搞笑
	Turtle_FUNNY Turtle_Category = 3
	// 其他
	Turtle_CUSTOM Turtle_Category = 100
)

// Enum value maps for Turtle_Category.
var (
	Turtle_Category_name = map[int32]string{
		0:   "UNKNOWN",
		1:   "HORROR",
		2:   "DETECTIVE",
		3:   "FUNNY",
		100: "CUSTOM",
	}
	Turtle_Category_value = map[string]int32{
		"UNKNOWN":   0,
		"HORROR":    1,
		"DETECTIVE": 2,
		"FUNNY":     3,
		"CUSTOM":    100,
	}
)

func (x Turtle_Category) Enum() *Turtle_Category {
	p := new(Turtle_Category)
	*p = x
	return p
}

func (x Turtle_Category) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Turtle_Category) Descriptor() protoreflect.EnumDescriptor {
	return file_api_turtle_v1_turtle_proto_enumTypes[1].Descriptor()
}

func (Turtle_Category) Type() protoreflect.EnumType {
	return &file_api_turtle_v1_turtle_proto_enumTypes[1]
}

func (x Turtle_Category) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Turtle_Category.Descriptor instead.
func (Turtle_Category) EnumDescriptor() ([]byte, []int) {
	return file_api_turtle_v1_turtle_proto_rawDescGZIP(), []int{3, 1}
}

type GetTurtleListReq struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// 占位符
	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	// 分页参数
	Page     int32 `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	PageSize int32 `protobuf:"varint,3,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// 状态
	State Turtle_State `protobuf:"varint,4,opt,name=state,proto3,enum=turtle.v1.Turtle_State" json:"state,omitempty"`
	// 分类
	Category      Turtle_Category `protobuf:"varint,5,opt,name=category,proto3,enum=turtle.v1.Turtle_Category" json:"category,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetTurtleListReq) Reset() {
	*x = GetTurtleListReq{}
	mi := &file_api_turtle_v1_turtle_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetTurtleListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTurtleListReq) ProtoMessage() {}

func (x *GetTurtleListReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_turtle_v1_turtle_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTurtleListReq.ProtoReflect.Descriptor instead.
func (*GetTurtleListReq) Descriptor() ([]byte, []int) {
	return file_api_turtle_v1_turtle_proto_rawDescGZIP(), []int{0}
}

func (x *GetTurtleListReq) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *GetTurtleListReq) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetTurtleListReq) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *GetTurtleListReq) GetState() Turtle_State {
	if x != nil {
		return x.State
	}
	return Turtle_INVALID
}

func (x *GetTurtleListReq) GetCategory() Turtle_Category {
	if x != nil {
		return x.Category
	}
	return Turtle_UNKNOWN
}

type GetTurtleListResp struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Turtles       []*Turtle              `protobuf:"bytes,1,rep,name=turtles,proto3" json:"turtles,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetTurtleListResp) Reset() {
	*x = GetTurtleListResp{}
	mi := &file_api_turtle_v1_turtle_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetTurtleListResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTurtleListResp) ProtoMessage() {}

func (x *GetTurtleListResp) ProtoReflect() protoreflect.Message {
	mi := &file_api_turtle_v1_turtle_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTurtleListResp.ProtoReflect.Descriptor instead.
func (*GetTurtleListResp) Descriptor() ([]byte, []int) {
	return file_api_turtle_v1_turtle_proto_rawDescGZIP(), []int{1}
}

func (x *GetTurtleListResp) GetTurtles() []*Turtle {
	if x != nil {
		return x.Turtles
	}
	return nil
}

// The request message containing the user's name.
type SetTurtleBatchReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Turtles       []*Turtle              `protobuf:"bytes,1,rep,name=turtles,proto3" json:"turtles,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SetTurtleBatchReq) Reset() {
	*x = SetTurtleBatchReq{}
	mi := &file_api_turtle_v1_turtle_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SetTurtleBatchReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetTurtleBatchReq) ProtoMessage() {}

func (x *SetTurtleBatchReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_turtle_v1_turtle_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetTurtleBatchReq.ProtoReflect.Descriptor instead.
func (*SetTurtleBatchReq) Descriptor() ([]byte, []int) {
	return file_api_turtle_v1_turtle_proto_rawDescGZIP(), []int{2}
}

func (x *SetTurtleBatchReq) GetTurtles() []*Turtle {
	if x != nil {
		return x.Turtles
	}
	return nil
}

type Turtle struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// 唯一标识（例：HT_001）
	Qid string `protobuf:"bytes,1,opt,name=qid,proto3" json:"qid,omitempty"`
	// 谜题标题（例：“深夜弹珠声”）
	Title string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	// 谜题完整描述
	Content string `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	// 标准答案
	Answer string `protobuf:"bytes,4,opt,name=answer,proto3" json:"answer,omitempty"`
	// 分类（恐怖/推理/搞笑/其他）
	Category Turtle_Category `protobuf:"varint,5,opt,name=category,proto3,enum=turtle.v1.Turtle_Category" json:"category,omitempty"`
	// 创建者（主播UID或观众UID）
	Creator string `protobuf:"bytes,6,opt,name=creator,proto3" json:"creator,omitempty"`
	// 状态（审核中/已上线/已删除）
	State         Turtle_State `protobuf:"varint,7,opt,name=state,proto3,enum=turtle.v1.Turtle_State" json:"state,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Turtle) Reset() {
	*x = Turtle{}
	mi := &file_api_turtle_v1_turtle_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Turtle) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Turtle) ProtoMessage() {}

func (x *Turtle) ProtoReflect() protoreflect.Message {
	mi := &file_api_turtle_v1_turtle_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Turtle.ProtoReflect.Descriptor instead.
func (*Turtle) Descriptor() ([]byte, []int) {
	return file_api_turtle_v1_turtle_proto_rawDescGZIP(), []int{3}
}

func (x *Turtle) GetQid() string {
	if x != nil {
		return x.Qid
	}
	return ""
}

func (x *Turtle) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Turtle) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Turtle) GetAnswer() string {
	if x != nil {
		return x.Answer
	}
	return ""
}

func (x *Turtle) GetCategory() Turtle_Category {
	if x != nil {
		return x.Category
	}
	return Turtle_UNKNOWN
}

func (x *Turtle) GetCreator() string {
	if x != nil {
		return x.Creator
	}
	return ""
}

func (x *Turtle) GetState() Turtle_State {
	if x != nil {
		return x.State
	}
	return Turtle_INVALID
}

var File_api_turtle_v1_turtle_proto protoreflect.FileDescriptor

var file_api_turtle_v1_turtle_proto_rawDesc = string([]byte{
	0x0a, 0x1a, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x75, 0x72, 0x74, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x2f,
	0x74, 0x75, 0x72, 0x74, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x74, 0x75,
	0x72, 0x74, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc0, 0x01,
	0x0a, 0x10, 0x47, 0x65, 0x74, 0x54, 0x75, 0x72, 0x74, 0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1b, 0x0a, 0x09,
	0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x2d, 0x0a, 0x05, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x74, 0x75, 0x72, 0x74, 0x6c,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x75, 0x72, 0x74, 0x6c, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x36, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x74, 0x75, 0x72,
	0x74, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x75, 0x72, 0x74, 0x6c, 0x65, 0x2e, 0x43, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x22, 0x40, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x54, 0x75, 0x72, 0x74, 0x6c, 0x65, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x2b, 0x0a, 0x07, 0x74, 0x75, 0x72, 0x74, 0x6c, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x74, 0x75, 0x72, 0x74, 0x6c, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x54, 0x75, 0x72, 0x74, 0x6c, 0x65, 0x52, 0x07, 0x74, 0x75, 0x72, 0x74, 0x6c,
	0x65, 0x73, 0x22, 0x40, 0x0a, 0x11, 0x53, 0x65, 0x74, 0x54, 0x75, 0x72, 0x74, 0x6c, 0x65, 0x42,
	0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x12, 0x2b, 0x0a, 0x07, 0x74, 0x75, 0x72, 0x74, 0x6c,
	0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x74, 0x75, 0x72, 0x74, 0x6c,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x75, 0x72, 0x74, 0x6c, 0x65, 0x52, 0x07, 0x74, 0x75, 0x72,
	0x74, 0x6c, 0x65, 0x73, 0x22, 0x99, 0x03, 0x0a, 0x06, 0x54, 0x75, 0x72, 0x74, 0x6c, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x71, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x71, 0x69,
	0x64, 0x12, 0x1d, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x12, 0x21, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x06, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x06, 0x61, 0x6e,
	0x73, 0x77, 0x65, 0x72, 0x12, 0x40, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x74, 0x75, 0x72, 0x74, 0x6c, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x54, 0x75, 0x72, 0x74, 0x6c, 0x65, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x82, 0x01, 0x02, 0x10, 0x01, 0x52, 0x08, 0x63, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x21, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f,
	0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01,
	0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x2d, 0x0a, 0x05, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x74, 0x75, 0x72, 0x74, 0x6c,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x75, 0x72, 0x74, 0x6c, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x22, 0x3b, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x12, 0x0b, 0x0a, 0x07, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x00, 0x12, 0x0c,
	0x0a, 0x08, 0x41, 0x55, 0x44, 0x49, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06,
	0x4f, 0x4e, 0x4c, 0x49, 0x4e, 0x45, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x4f, 0x46, 0x46, 0x4c,
	0x49, 0x4e, 0x45, 0x10, 0x03, 0x22, 0x49, 0x0a, 0x08, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0a,
	0x0a, 0x06, 0x48, 0x4f, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x44, 0x45,
	0x54, 0x45, 0x43, 0x54, 0x49, 0x56, 0x45, 0x10, 0x02, 0x12, 0x09, 0x0a, 0x05, 0x46, 0x55, 0x4e,
	0x4e, 0x59, 0x10, 0x03, 0x12, 0x0a, 0x0a, 0x06, 0x43, 0x55, 0x53, 0x54, 0x4f, 0x4d, 0x10, 0x64,
	0x32, 0xdc, 0x01, 0x0a, 0x06, 0x74, 0x75, 0x72, 0x74, 0x6c, 0x65, 0x12, 0x64, 0x0a, 0x0e, 0x53,
	0x65, 0x74, 0x54, 0x75, 0x72, 0x74, 0x6c, 0x65, 0x42, 0x61, 0x74, 0x63, 0x68, 0x12, 0x1c, 0x2e,
	0x74, 0x75, 0x72, 0x74, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x74, 0x54, 0x75, 0x72,
	0x74, 0x6c, 0x65, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x3a, 0x01, 0x2a, 0x22, 0x11,
	0x2f, 0x73, 0x65, 0x74, 0x2f, 0x74, 0x75, 0x72, 0x74, 0x6c, 0x65, 0x2f, 0x62, 0x61, 0x74, 0x63,
	0x68, 0x12, 0x6c, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x54, 0x75, 0x72, 0x74, 0x6c, 0x65, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x1b, 0x2e, 0x74, 0x75, 0x72, 0x74, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x54, 0x75, 0x72, 0x74, 0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x1a,
	0x1c, 0x2e, 0x74, 0x75, 0x72, 0x74, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x54,
	0x75, 0x72, 0x74, 0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x22, 0x20, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x12, 0x18, 0x2f, 0x67, 0x65, 0x74, 0x2f, 0x74, 0x75, 0x72, 0x74,
	0x6c, 0x65, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x2f, 0x7b, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x7d, 0x42,
	0x46, 0x0a, 0x18, 0x64, 0x65, 0x76, 0x2e, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x74, 0x75, 0x72, 0x74, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x0d, 0x54, 0x75, 0x72,
	0x74, 0x6c, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x56, 0x31, 0x50, 0x01, 0x5a, 0x19, 0x78, 0x69,
	0x75, 0x79, 0x69, 0x50, 0x72, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x75, 0x72, 0x74, 0x6c,
	0x65, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_api_turtle_v1_turtle_proto_rawDescOnce sync.Once
	file_api_turtle_v1_turtle_proto_rawDescData []byte
)

func file_api_turtle_v1_turtle_proto_rawDescGZIP() []byte {
	file_api_turtle_v1_turtle_proto_rawDescOnce.Do(func() {
		file_api_turtle_v1_turtle_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_api_turtle_v1_turtle_proto_rawDesc), len(file_api_turtle_v1_turtle_proto_rawDesc)))
	})
	return file_api_turtle_v1_turtle_proto_rawDescData
}

var file_api_turtle_v1_turtle_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_api_turtle_v1_turtle_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_api_turtle_v1_turtle_proto_goTypes = []any{
	(Turtle_State)(0),         // 0: turtle.v1.Turtle.State
	(Turtle_Category)(0),      // 1: turtle.v1.Turtle.Category
	(*GetTurtleListReq)(nil),  // 2: turtle.v1.GetTurtleListReq
	(*GetTurtleListResp)(nil), // 3: turtle.v1.GetTurtleListResp
	(*SetTurtleBatchReq)(nil), // 4: turtle.v1.SetTurtleBatchReq
	(*Turtle)(nil),            // 5: turtle.v1.Turtle
	(*emptypb.Empty)(nil),     // 6: google.protobuf.Empty
}
var file_api_turtle_v1_turtle_proto_depIdxs = []int32{
	0, // 0: turtle.v1.GetTurtleListReq.state:type_name -> turtle.v1.Turtle.State
	1, // 1: turtle.v1.GetTurtleListReq.category:type_name -> turtle.v1.Turtle.Category
	5, // 2: turtle.v1.GetTurtleListResp.turtles:type_name -> turtle.v1.Turtle
	5, // 3: turtle.v1.SetTurtleBatchReq.turtles:type_name -> turtle.v1.Turtle
	1, // 4: turtle.v1.Turtle.category:type_name -> turtle.v1.Turtle.Category
	0, // 5: turtle.v1.Turtle.state:type_name -> turtle.v1.Turtle.State
	4, // 6: turtle.v1.turtle.SetTurtleBatch:input_type -> turtle.v1.SetTurtleBatchReq
	2, // 7: turtle.v1.turtle.GetTurtleList:input_type -> turtle.v1.GetTurtleListReq
	6, // 8: turtle.v1.turtle.SetTurtleBatch:output_type -> google.protobuf.Empty
	3, // 9: turtle.v1.turtle.GetTurtleList:output_type -> turtle.v1.GetTurtleListResp
	8, // [8:10] is the sub-list for method output_type
	6, // [6:8] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_api_turtle_v1_turtle_proto_init() }
func file_api_turtle_v1_turtle_proto_init() {
	if File_api_turtle_v1_turtle_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_api_turtle_v1_turtle_proto_rawDesc), len(file_api_turtle_v1_turtle_proto_rawDesc)),
			NumEnums:      2,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_turtle_v1_turtle_proto_goTypes,
		DependencyIndexes: file_api_turtle_v1_turtle_proto_depIdxs,
		EnumInfos:         file_api_turtle_v1_turtle_proto_enumTypes,
		MessageInfos:      file_api_turtle_v1_turtle_proto_msgTypes,
	}.Build()
	File_api_turtle_v1_turtle_proto = out.File
	file_api_turtle_v1_turtle_proto_goTypes = nil
	file_api_turtle_v1_turtle_proto_depIdxs = nil
}
