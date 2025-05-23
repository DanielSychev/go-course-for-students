// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: lesson9/homework/internal/ports/grpc/service.proto

package grpc

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type CreateAdRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Title         string                 `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Text          string                 `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
	UserId        int64                  `protobuf:"varint,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateAdRequest) Reset() {
	*x = CreateAdRequest{}
	mi := &file_lesson9_homework_internal_ports_grpc_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateAdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAdRequest) ProtoMessage() {}

func (x *CreateAdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_lesson9_homework_internal_ports_grpc_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAdRequest.ProtoReflect.Descriptor instead.
func (*CreateAdRequest) Descriptor() ([]byte, []int) {
	return file_lesson9_homework_internal_ports_grpc_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreateAdRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateAdRequest) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *CreateAdRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type ChangeAdStatusRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	AdId          int64                  `protobuf:"varint,1,opt,name=ad_id,json=adId,proto3" json:"ad_id,omitempty"`
	UserId        int64                  `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Published     bool                   `protobuf:"varint,3,opt,name=published,proto3" json:"published,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ChangeAdStatusRequest) Reset() {
	*x = ChangeAdStatusRequest{}
	mi := &file_lesson9_homework_internal_ports_grpc_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ChangeAdStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangeAdStatusRequest) ProtoMessage() {}

func (x *ChangeAdStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_lesson9_homework_internal_ports_grpc_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangeAdStatusRequest.ProtoReflect.Descriptor instead.
func (*ChangeAdStatusRequest) Descriptor() ([]byte, []int) {
	return file_lesson9_homework_internal_ports_grpc_service_proto_rawDescGZIP(), []int{1}
}

func (x *ChangeAdStatusRequest) GetAdId() int64 {
	if x != nil {
		return x.AdId
	}
	return 0
}

func (x *ChangeAdStatusRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *ChangeAdStatusRequest) GetPublished() bool {
	if x != nil {
		return x.Published
	}
	return false
}

type UpdateAdRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	AdId          int64                  `protobuf:"varint,1,opt,name=ad_id,json=adId,proto3" json:"ad_id,omitempty"`
	Title         string                 `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Text          string                 `protobuf:"bytes,3,opt,name=text,proto3" json:"text,omitempty"`
	UserId        int64                  `protobuf:"varint,4,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateAdRequest) Reset() {
	*x = UpdateAdRequest{}
	mi := &file_lesson9_homework_internal_ports_grpc_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateAdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAdRequest) ProtoMessage() {}

func (x *UpdateAdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_lesson9_homework_internal_ports_grpc_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAdRequest.ProtoReflect.Descriptor instead.
func (*UpdateAdRequest) Descriptor() ([]byte, []int) {
	return file_lesson9_homework_internal_ports_grpc_service_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateAdRequest) GetAdId() int64 {
	if x != nil {
		return x.AdId
	}
	return 0
}

func (x *UpdateAdRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *UpdateAdRequest) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *UpdateAdRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type AdResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title         string                 `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Text          string                 `protobuf:"bytes,3,opt,name=text,proto3" json:"text,omitempty"`
	AuthorId      int64                  `protobuf:"varint,4,opt,name=author_id,json=authorId,proto3" json:"author_id,omitempty"`
	Published     bool                   `protobuf:"varint,5,opt,name=published,proto3" json:"published,omitempty"`
	DateCreated   string                 `protobuf:"bytes,6,opt,name=date_created,json=dateCreated,proto3" json:"date_created,omitempty"`
	DateUpdated   string                 `protobuf:"bytes,7,opt,name=date_updated,json=dateUpdated,proto3" json:"date_updated,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AdResponse) Reset() {
	*x = AdResponse{}
	mi := &file_lesson9_homework_internal_ports_grpc_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdResponse) ProtoMessage() {}

func (x *AdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_lesson9_homework_internal_ports_grpc_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdResponse.ProtoReflect.Descriptor instead.
func (*AdResponse) Descriptor() ([]byte, []int) {
	return file_lesson9_homework_internal_ports_grpc_service_proto_rawDescGZIP(), []int{3}
}

func (x *AdResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *AdResponse) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *AdResponse) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *AdResponse) GetAuthorId() int64 {
	if x != nil {
		return x.AuthorId
	}
	return 0
}

func (x *AdResponse) GetPublished() bool {
	if x != nil {
		return x.Published
	}
	return false
}

func (x *AdResponse) GetDateCreated() string {
	if x != nil {
		return x.DateCreated
	}
	return ""
}

func (x *AdResponse) GetDateUpdated() string {
	if x != nil {
		return x.DateUpdated
	}
	return ""
}

type ListAdResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	List          []*AdResponse          `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListAdResponse) Reset() {
	*x = ListAdResponse{}
	mi := &file_lesson9_homework_internal_ports_grpc_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListAdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAdResponse) ProtoMessage() {}

func (x *ListAdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_lesson9_homework_internal_ports_grpc_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAdResponse.ProtoReflect.Descriptor instead.
func (*ListAdResponse) Descriptor() ([]byte, []int) {
	return file_lesson9_homework_internal_ports_grpc_service_proto_rawDescGZIP(), []int{4}
}

func (x *ListAdResponse) GetList() []*AdResponse {
	if x != nil {
		return x.List
	}
	return nil
}

type CreateUserRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateUserRequest) Reset() {
	*x = CreateUserRequest{}
	mi := &file_lesson9_homework_internal_ports_grpc_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserRequest) ProtoMessage() {}

func (x *CreateUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_lesson9_homework_internal_ports_grpc_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserRequest.ProtoReflect.Descriptor instead.
func (*CreateUserRequest) Descriptor() ([]byte, []int) {
	return file_lesson9_homework_internal_ports_grpc_service_proto_rawDescGZIP(), []int{5}
}

func (x *CreateUserRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type UserResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserResponse) Reset() {
	*x = UserResponse{}
	mi := &file_lesson9_homework_internal_ports_grpc_service_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserResponse) ProtoMessage() {}

func (x *UserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_lesson9_homework_internal_ports_grpc_service_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserResponse.ProtoReflect.Descriptor instead.
func (*UserResponse) Descriptor() ([]byte, []int) {
	return file_lesson9_homework_internal_ports_grpc_service_proto_rawDescGZIP(), []int{6}
}

func (x *UserResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UserResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GetUserRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetUserRequest) Reset() {
	*x = GetUserRequest{}
	mi := &file_lesson9_homework_internal_ports_grpc_service_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserRequest) ProtoMessage() {}

func (x *GetUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_lesson9_homework_internal_ports_grpc_service_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserRequest.ProtoReflect.Descriptor instead.
func (*GetUserRequest) Descriptor() ([]byte, []int) {
	return file_lesson9_homework_internal_ports_grpc_service_proto_rawDescGZIP(), []int{7}
}

func (x *GetUserRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteUserRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteUserRequest) Reset() {
	*x = DeleteUserRequest{}
	mi := &file_lesson9_homework_internal_ports_grpc_service_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteUserRequest) ProtoMessage() {}

func (x *DeleteUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_lesson9_homework_internal_ports_grpc_service_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteUserRequest.ProtoReflect.Descriptor instead.
func (*DeleteUserRequest) Descriptor() ([]byte, []int) {
	return file_lesson9_homework_internal_ports_grpc_service_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteUserRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteAdRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	AdId          int64                  `protobuf:"varint,1,opt,name=ad_id,json=adId,proto3" json:"ad_id,omitempty"`
	AuthorId      int64                  `protobuf:"varint,2,opt,name=author_id,json=authorId,proto3" json:"author_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteAdRequest) Reset() {
	*x = DeleteAdRequest{}
	mi := &file_lesson9_homework_internal_ports_grpc_service_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteAdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAdRequest) ProtoMessage() {}

func (x *DeleteAdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_lesson9_homework_internal_ports_grpc_service_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAdRequest.ProtoReflect.Descriptor instead.
func (*DeleteAdRequest) Descriptor() ([]byte, []int) {
	return file_lesson9_homework_internal_ports_grpc_service_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteAdRequest) GetAdId() int64 {
	if x != nil {
		return x.AdId
	}
	return 0
}

func (x *DeleteAdRequest) GetAuthorId() int64 {
	if x != nil {
		return x.AuthorId
	}
	return 0
}

var File_lesson9_homework_internal_ports_grpc_service_proto protoreflect.FileDescriptor

var file_lesson9_homework_internal_ports_grpc_service_proto_rawDesc = string([]byte{
	0x0a, 0x32, 0x6c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x39, 0x2f, 0x68, 0x6f, 0x6d, 0x65, 0x77, 0x6f,
	0x72, 0x6b, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x6f, 0x72, 0x74,
	0x73, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x61, 0x64, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x54, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65,
	0x78, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x63, 0x0a, 0x15, 0x43,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x41, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x13, 0x0a, 0x05, 0x61, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x04, 0x61, 0x64, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64,
	0x22, 0x69, 0x0a, 0x0f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x13, 0x0a, 0x05, 0x61, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x04, 0x61, 0x64, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65,
	0x78, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0xc7, 0x01, 0x0a, 0x0a,
	0x41, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x74, 0x65, 0x78, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x49,
	0x64, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x12,
	0x21, 0x0a, 0x0c, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x61, 0x74, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x61, 0x74, 0x65, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x22, 0x34, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x61, 0x64, 0x2e, 0x41, 0x64, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x22, 0x27, 0x0a, 0x11, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x22, 0x32, 0x0a, 0x0c, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x20, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x23, 0x0a, 0x11, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x43, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x13, 0x0a, 0x05, 0x61, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x04, 0x61, 0x64, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x61, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x49, 0x64, 0x32, 0xcf, 0x03, 0x0a, 0x09, 0x41, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x31, 0x0a, 0x08, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x64, 0x12, 0x13,
	0x2e, 0x61, 0x64, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x61, 0x64, 0x2e, 0x41, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x0e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x41,
	0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x19, 0x2e, 0x61, 0x64, 0x2e, 0x43, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x41, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x61, 0x64, 0x2e, 0x41, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x31, 0x0a, 0x08, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x64,
	0x12, 0x13, 0x2e, 0x61, 0x64, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x61, 0x64, 0x2e, 0x41, 0x64, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x07, 0x4c, 0x69, 0x73, 0x74, 0x41,
	0x64, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x12, 0x2e, 0x61, 0x64, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x41, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x37, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x15,
	0x2e, 0x61, 0x64, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x61, 0x64, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x31, 0x0a, 0x07, 0x47, 0x65, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x12, 0x12, 0x2e, 0x61, 0x64, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x61, 0x64, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x0a,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x15, 0x2e, 0x61, 0x64, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x08, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x64, 0x12, 0x13, 0x2e, 0x61, 0x64, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x41, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x26, 0x5a, 0x24, 0x6c, 0x65, 0x73, 0x73, 0x6f, 0x6e,
	0x39, 0x2f, 0x68, 0x6f, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_lesson9_homework_internal_ports_grpc_service_proto_rawDescOnce sync.Once
	file_lesson9_homework_internal_ports_grpc_service_proto_rawDescData []byte
)

func file_lesson9_homework_internal_ports_grpc_service_proto_rawDescGZIP() []byte {
	file_lesson9_homework_internal_ports_grpc_service_proto_rawDescOnce.Do(func() {
		file_lesson9_homework_internal_ports_grpc_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_lesson9_homework_internal_ports_grpc_service_proto_rawDesc), len(file_lesson9_homework_internal_ports_grpc_service_proto_rawDesc)))
	})
	return file_lesson9_homework_internal_ports_grpc_service_proto_rawDescData
}

var file_lesson9_homework_internal_ports_grpc_service_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_lesson9_homework_internal_ports_grpc_service_proto_goTypes = []any{
	(*CreateAdRequest)(nil),       // 0: ad.CreateAdRequest
	(*ChangeAdStatusRequest)(nil), // 1: ad.ChangeAdStatusRequest
	(*UpdateAdRequest)(nil),       // 2: ad.UpdateAdRequest
	(*AdResponse)(nil),            // 3: ad.AdResponse
	(*ListAdResponse)(nil),        // 4: ad.ListAdResponse
	(*CreateUserRequest)(nil),     // 5: ad.CreateUserRequest
	(*UserResponse)(nil),          // 6: ad.UserResponse
	(*GetUserRequest)(nil),        // 7: ad.GetUserRequest
	(*DeleteUserRequest)(nil),     // 8: ad.DeleteUserRequest
	(*DeleteAdRequest)(nil),       // 9: ad.DeleteAdRequest
	(*emptypb.Empty)(nil),         // 10: google.protobuf.Empty
}
var file_lesson9_homework_internal_ports_grpc_service_proto_depIdxs = []int32{
	3,  // 0: ad.ListAdResponse.list:type_name -> ad.AdResponse
	0,  // 1: ad.AdService.CreateAd:input_type -> ad.CreateAdRequest
	1,  // 2: ad.AdService.ChangeAdStatus:input_type -> ad.ChangeAdStatusRequest
	2,  // 3: ad.AdService.UpdateAd:input_type -> ad.UpdateAdRequest
	10, // 4: ad.AdService.ListAds:input_type -> google.protobuf.Empty
	5,  // 5: ad.AdService.CreateUser:input_type -> ad.CreateUserRequest
	7,  // 6: ad.AdService.GetUser:input_type -> ad.GetUserRequest
	8,  // 7: ad.AdService.DeleteUser:input_type -> ad.DeleteUserRequest
	9,  // 8: ad.AdService.DeleteAd:input_type -> ad.DeleteAdRequest
	3,  // 9: ad.AdService.CreateAd:output_type -> ad.AdResponse
	3,  // 10: ad.AdService.ChangeAdStatus:output_type -> ad.AdResponse
	3,  // 11: ad.AdService.UpdateAd:output_type -> ad.AdResponse
	4,  // 12: ad.AdService.ListAds:output_type -> ad.ListAdResponse
	6,  // 13: ad.AdService.CreateUser:output_type -> ad.UserResponse
	6,  // 14: ad.AdService.GetUser:output_type -> ad.UserResponse
	10, // 15: ad.AdService.DeleteUser:output_type -> google.protobuf.Empty
	10, // 16: ad.AdService.DeleteAd:output_type -> google.protobuf.Empty
	9,  // [9:17] is the sub-list for method output_type
	1,  // [1:9] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_lesson9_homework_internal_ports_grpc_service_proto_init() }
func file_lesson9_homework_internal_ports_grpc_service_proto_init() {
	if File_lesson9_homework_internal_ports_grpc_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_lesson9_homework_internal_ports_grpc_service_proto_rawDesc), len(file_lesson9_homework_internal_ports_grpc_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_lesson9_homework_internal_ports_grpc_service_proto_goTypes,
		DependencyIndexes: file_lesson9_homework_internal_ports_grpc_service_proto_depIdxs,
		MessageInfos:      file_lesson9_homework_internal_ports_grpc_service_proto_msgTypes,
	}.Build()
	File_lesson9_homework_internal_ports_grpc_service_proto = out.File
	file_lesson9_homework_internal_ports_grpc_service_proto_goTypes = nil
	file_lesson9_homework_internal_ports_grpc_service_proto_depIdxs = nil
}
