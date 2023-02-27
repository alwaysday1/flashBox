// # Copyright 2023 LightTree <alwaysday1@qq.com>. All rights reserved.
// # Use of this source code is governed by a MIT style
// # license that can be found in the LICENSE file. The original repo for
// # this file is https://github.com/alwaysday1/flashBox.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.9
// source: flashBox/v1/flashBox.proto

package v1

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

type UserInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username  string                 `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Nickname  string                 `protobuf:"bytes,2,opt,name=nickname,proto3" json:"nickname,omitempty"`
	Email     string                 `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Phone     string                 `protobuf:"bytes,4,opt,name=phone,proto3" json:"phone,omitempty"`
	PostCount int64                  `protobuf:"varint,5,opt,name=postCount,proto3" json:"postCount,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
}

func (x *UserInfo) Reset() {
	*x = UserInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flashBox_v1_flashBox_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfo) ProtoMessage() {}

func (x *UserInfo) ProtoReflect() protoreflect.Message {
	mi := &file_flashBox_v1_flashBox_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfo.ProtoReflect.Descriptor instead.
func (*UserInfo) Descriptor() ([]byte, []int) {
	return file_flashBox_v1_flashBox_proto_rawDescGZIP(), []int{0}
}

func (x *UserInfo) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *UserInfo) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *UserInfo) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UserInfo) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *UserInfo) GetPostCount() int64 {
	if x != nil {
		return x.PostCount
	}
	return 0
}

func (x *UserInfo) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *UserInfo) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

// ListUserRequest 指定了 `ListUser` 接口的请求参数，相当于 HTTP Request 并对每个属性都定义数据类型.
// 需要为每个属性分配一个唯一编号，称为标记。此标记由 protobuf 用于表示属性，而不是使用属性名称.
// 因此，在 JSON 中我们每次都需要传递属性名称 name，而 protobuf 将使用数字 1 来表示 name.
type ListUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit  int64 `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset int64 `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *ListUserRequest) Reset() {
	*x = ListUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flashBox_v1_flashBox_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListUserRequest) ProtoMessage() {}

func (x *ListUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_flashBox_v1_flashBox_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListUserRequest.ProtoReflect.Descriptor instead.
func (*ListUserRequest) Descriptor() ([]byte, []int) {
	return file_flashBox_v1_flashBox_proto_rawDescGZIP(), []int{1}
}

func (x *ListUserRequest) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ListUserRequest) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

// ListUserResponse 指定了 `GET /v1/users` 接口的返回参数，相当于 HTTP Response.
type ListUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TotalCount int64       `protobuf:"varint,1,opt,name=totalCount,proto3" json:"totalCount,omitempty"`
	Users      []*UserInfo `protobuf:"bytes,2,rep,name=Users,proto3" json:"Users,omitempty"`
}

func (x *ListUserResponse) Reset() {
	*x = ListUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flashBox_v1_flashBox_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListUserResponse) ProtoMessage() {}

func (x *ListUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_flashBox_v1_flashBox_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListUserResponse.ProtoReflect.Descriptor instead.
func (*ListUserResponse) Descriptor() ([]byte, []int) {
	return file_flashBox_v1_flashBox_proto_rawDescGZIP(), []int{2}
}

func (x *ListUserResponse) GetTotalCount() int64 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

func (x *ListUserResponse) GetUsers() []*UserInfo {
	if x != nil {
		return x.Users
	}
	return nil
}

// 示例 message 定义，用来展示 protobuf 修饰符，编译后的效果
type ModifierExample struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username  string                 `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Nickname  *string                `protobuf:"bytes,3,opt,name=nickname,proto3,oneof" json:"nickname,omitempty"` // 可选值
	HasEmail  bool                   `protobuf:"varint,4,opt,name=hasEmail,proto3" json:"hasEmail,omitempty"`
	Phones    []string               `protobuf:"bytes,5,rep,name=phones,proto3" json:"phones,omitempty"` // 任意数量（包括0），可以看成是一个数组.
	Content   []byte                 `protobuf:"bytes,6,opt,name=content,proto3" json:"content,omitempty"`
	Money     int32                  `protobuf:"varint,7,opt,name=money,proto3" json:"money,omitempty"`
	Count     int64                  `protobuf:"varint,8,opt,name=count,proto3" json:"count,omitempty"`
	Address   map[string]int64       `protobuf:"bytes,9,rep,name=address,proto3" json:"address,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"` // 编译后：`Address   map[string]string`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,10,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
}

func (x *ModifierExample) Reset() {
	*x = ModifierExample{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flashBox_v1_flashBox_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModifierExample) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModifierExample) ProtoMessage() {}

func (x *ModifierExample) ProtoReflect() protoreflect.Message {
	mi := &file_flashBox_v1_flashBox_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModifierExample.ProtoReflect.Descriptor instead.
func (*ModifierExample) Descriptor() ([]byte, []int) {
	return file_flashBox_v1_flashBox_proto_rawDescGZIP(), []int{3}
}

func (x *ModifierExample) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *ModifierExample) GetNickname() string {
	if x != nil && x.Nickname != nil {
		return *x.Nickname
	}
	return ""
}

func (x *ModifierExample) GetHasEmail() bool {
	if x != nil {
		return x.HasEmail
	}
	return false
}

func (x *ModifierExample) GetPhones() []string {
	if x != nil {
		return x.Phones
	}
	return nil
}

func (x *ModifierExample) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

func (x *ModifierExample) GetMoney() int32 {
	if x != nil {
		return x.Money
	}
	return 0
}

func (x *ModifierExample) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *ModifierExample) GetAddress() map[string]int64 {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *ModifierExample) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

var File_flashBox_v1_flashBox_proto protoreflect.FileDescriptor

var file_flashBox_v1_flashBox_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x66, 0x6c, 0x61, 0x73, 0x68, 0x42, 0x6f, 0x78, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x6c,
	0x61, 0x73, 0x68, 0x42, 0x6f, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x76, 0x31,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x80, 0x02, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1a,
	0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69,
	0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69,
	0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x14, 0x0a, 0x05,
	0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f,
	0x6e, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x6f, 0x73, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x70, 0x6f, 0x73, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x38, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x38, 0x0a, 0x09, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x22, 0x3f, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6f,
	0x66, 0x66, 0x73, 0x65, 0x74, 0x22, 0x56, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x22, 0x0a, 0x05, 0x55, 0x73, 0x65,
	0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x55, 0x73, 0x65, 0x72, 0x73, 0x22, 0x93, 0x03,
	0x0a, 0x0f, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x72, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a,
	0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x00, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1a,
	0x0a, 0x08, 0x68, 0x61, 0x73, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x08, 0x68, 0x61, 0x73, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x68,
	0x6f, 0x6e, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x70, 0x68, 0x6f, 0x6e,
	0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x6d, 0x6f, 0x6e, 0x65, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6d, 0x6f, 0x6e,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x3a, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x76, 0x31, 0x2e, 0x4d,
	0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x72, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x12, 0x38, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x1a, 0x3a,
	0x0a, 0x0c, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x6e,
	0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x4a, 0x04, 0x08, 0x02, 0x10, 0x03, 0x4a, 0x04, 0x08,
	0x0f, 0x10, 0x1a, 0x32, 0x43, 0x0a, 0x08, 0x46, 0x6c, 0x61, 0x73, 0x68, 0x42, 0x6f, 0x78, 0x12,
	0x37, 0x0a, 0x08, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x12, 0x13, 0x2e, 0x76, 0x31,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x14, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x28, 0x5a, 0x26, 0x66, 0x6c, 0x61, 0x73,
	0x68, 0x5f, 0x62, 0x6f, 0x78, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x6b, 0x67,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x66, 0x6c, 0x61, 0x73, 0x68, 0x42, 0x6f, 0x78, 0x2f,
	0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_flashBox_v1_flashBox_proto_rawDescOnce sync.Once
	file_flashBox_v1_flashBox_proto_rawDescData = file_flashBox_v1_flashBox_proto_rawDesc
)

func file_flashBox_v1_flashBox_proto_rawDescGZIP() []byte {
	file_flashBox_v1_flashBox_proto_rawDescOnce.Do(func() {
		file_flashBox_v1_flashBox_proto_rawDescData = protoimpl.X.CompressGZIP(file_flashBox_v1_flashBox_proto_rawDescData)
	})
	return file_flashBox_v1_flashBox_proto_rawDescData
}

var file_flashBox_v1_flashBox_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_flashBox_v1_flashBox_proto_goTypes = []interface{}{
	(*UserInfo)(nil),              // 0: v1.UserInfo
	(*ListUserRequest)(nil),       // 1: v1.ListUserRequest
	(*ListUserResponse)(nil),      // 2: v1.ListUserResponse
	(*ModifierExample)(nil),       // 3: v1.ModifierExample
	nil,                           // 4: v1.ModifierExample.AddressEntry
	(*timestamppb.Timestamp)(nil), // 5: google.protobuf.Timestamp
}
var file_flashBox_v1_flashBox_proto_depIdxs = []int32{
	5, // 0: v1.UserInfo.createdAt:type_name -> google.protobuf.Timestamp
	5, // 1: v1.UserInfo.updatedAt:type_name -> google.protobuf.Timestamp
	0, // 2: v1.ListUserResponse.Users:type_name -> v1.UserInfo
	4, // 3: v1.ModifierExample.address:type_name -> v1.ModifierExample.AddressEntry
	5, // 4: v1.ModifierExample.createdAt:type_name -> google.protobuf.Timestamp
	1, // 5: v1.FlashBox.ListUser:input_type -> v1.ListUserRequest
	2, // 6: v1.FlashBox.ListUser:output_type -> v1.ListUserResponse
	6, // [6:7] is the sub-list for method output_type
	5, // [5:6] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_flashBox_v1_flashBox_proto_init() }
func file_flashBox_v1_flashBox_proto_init() {
	if File_flashBox_v1_flashBox_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_flashBox_v1_flashBox_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserInfo); i {
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
		file_flashBox_v1_flashBox_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListUserRequest); i {
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
		file_flashBox_v1_flashBox_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListUserResponse); i {
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
		file_flashBox_v1_flashBox_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModifierExample); i {
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
	file_flashBox_v1_flashBox_proto_msgTypes[3].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_flashBox_v1_flashBox_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_flashBox_v1_flashBox_proto_goTypes,
		DependencyIndexes: file_flashBox_v1_flashBox_proto_depIdxs,
		MessageInfos:      file_flashBox_v1_flashBox_proto_msgTypes,
	}.Build()
	File_flashBox_v1_flashBox_proto = out.File
	file_flashBox_v1_flashBox_proto_rawDesc = nil
	file_flashBox_v1_flashBox_proto_goTypes = nil
	file_flashBox_v1_flashBox_proto_depIdxs = nil
}
