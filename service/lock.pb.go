// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.1
// source: lock.proto

package service

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

type LockRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientId string `protobuf:"bytes,1,opt,name=clientId,proto3" json:"clientId,omitempty"`
	LockName string `protobuf:"bytes,2,opt,name=lockName,proto3" json:"lockName,omitempty"`
}

func (x *LockRequest) Reset() {
	*x = LockRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lock_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LockRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LockRequest) ProtoMessage() {}

func (x *LockRequest) ProtoReflect() protoreflect.Message {
	mi := &file_lock_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LockRequest.ProtoReflect.Descriptor instead.
func (*LockRequest) Descriptor() ([]byte, []int) {
	return file_lock_proto_rawDescGZIP(), []int{0}
}

func (x *LockRequest) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *LockRequest) GetLockName() string {
	if x != nil {
		return x.LockName
	}
	return ""
}

type LockReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result bool   `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	Msg    string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *LockReply) Reset() {
	*x = LockReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lock_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LockReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LockReply) ProtoMessage() {}

func (x *LockReply) ProtoReflect() protoreflect.Message {
	mi := &file_lock_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LockReply.ProtoReflect.Descriptor instead.
func (*LockReply) Descriptor() ([]byte, []int) {
	return file_lock_proto_rawDescGZIP(), []int{1}
}

func (x *LockReply) GetResult() bool {
	if x != nil {
		return x.Result
	}
	return false
}

func (x *LockReply) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type UnLockRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientId string `protobuf:"bytes,1,opt,name=clientId,proto3" json:"clientId,omitempty"`
	LockName string `protobuf:"bytes,2,opt,name=lockName,proto3" json:"lockName,omitempty"`
}

func (x *UnLockRequest) Reset() {
	*x = UnLockRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lock_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnLockRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnLockRequest) ProtoMessage() {}

func (x *UnLockRequest) ProtoReflect() protoreflect.Message {
	mi := &file_lock_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnLockRequest.ProtoReflect.Descriptor instead.
func (*UnLockRequest) Descriptor() ([]byte, []int) {
	return file_lock_proto_rawDescGZIP(), []int{2}
}

func (x *UnLockRequest) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *UnLockRequest) GetLockName() string {
	if x != nil {
		return x.LockName
	}
	return ""
}

type UnLockReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result bool   `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	Msg    string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *UnLockReply) Reset() {
	*x = UnLockReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lock_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnLockReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnLockReply) ProtoMessage() {}

func (x *UnLockReply) ProtoReflect() protoreflect.Message {
	mi := &file_lock_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnLockReply.ProtoReflect.Descriptor instead.
func (*UnLockReply) Descriptor() ([]byte, []int) {
	return file_lock_proto_rawDescGZIP(), []int{3}
}

func (x *UnLockReply) GetResult() bool {
	if x != nil {
		return x.Result
	}
	return false
}

func (x *UnLockReply) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type ForceLockRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientId string `protobuf:"bytes,1,opt,name=clientId,proto3" json:"clientId,omitempty"`
	LockName string `protobuf:"bytes,2,opt,name=lockName,proto3" json:"lockName,omitempty"`
}

func (x *ForceLockRequest) Reset() {
	*x = ForceLockRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lock_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ForceLockRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ForceLockRequest) ProtoMessage() {}

func (x *ForceLockRequest) ProtoReflect() protoreflect.Message {
	mi := &file_lock_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ForceLockRequest.ProtoReflect.Descriptor instead.
func (*ForceLockRequest) Descriptor() ([]byte, []int) {
	return file_lock_proto_rawDescGZIP(), []int{4}
}

func (x *ForceLockRequest) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *ForceLockRequest) GetLockName() string {
	if x != nil {
		return x.LockName
	}
	return ""
}

type ForceLockReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result bool   `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	Msg    string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *ForceLockReply) Reset() {
	*x = ForceLockReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lock_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ForceLockReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ForceLockReply) ProtoMessage() {}

func (x *ForceLockReply) ProtoReflect() protoreflect.Message {
	mi := &file_lock_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ForceLockReply.ProtoReflect.Descriptor instead.
func (*ForceLockReply) Descriptor() ([]byte, []int) {
	return file_lock_proto_rawDescGZIP(), []int{5}
}

func (x *ForceLockReply) GetResult() bool {
	if x != nil {
		return x.Result
	}
	return false
}

func (x *ForceLockReply) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type ForceUnLockRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LockName string `protobuf:"bytes,2,opt,name=lockName,proto3" json:"lockName,omitempty"`
}

func (x *ForceUnLockRequest) Reset() {
	*x = ForceUnLockRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lock_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ForceUnLockRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ForceUnLockRequest) ProtoMessage() {}

func (x *ForceUnLockRequest) ProtoReflect() protoreflect.Message {
	mi := &file_lock_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ForceUnLockRequest.ProtoReflect.Descriptor instead.
func (*ForceUnLockRequest) Descriptor() ([]byte, []int) {
	return file_lock_proto_rawDescGZIP(), []int{6}
}

func (x *ForceUnLockRequest) GetLockName() string {
	if x != nil {
		return x.LockName
	}
	return ""
}

type ForceUnLockReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result bool   `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	Msg    string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *ForceUnLockReply) Reset() {
	*x = ForceUnLockReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lock_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ForceUnLockReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ForceUnLockReply) ProtoMessage() {}

func (x *ForceUnLockReply) ProtoReflect() protoreflect.Message {
	mi := &file_lock_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ForceUnLockReply.ProtoReflect.Descriptor instead.
func (*ForceUnLockReply) Descriptor() ([]byte, []int) {
	return file_lock_proto_rawDescGZIP(), []int{7}
}

func (x *ForceUnLockReply) GetResult() bool {
	if x != nil {
		return x.Result
	}
	return false
}

func (x *ForceUnLockReply) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

var File_lock_proto protoreflect.FileDescriptor

var file_lock_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x45, 0x0a, 0x0b,
	0x4c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x6b, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x6b, 0x4e,
	0x61, 0x6d, 0x65, 0x22, 0x35, 0x0a, 0x09, 0x4c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x47, 0x0a, 0x0d, 0x55, 0x6e,
	0x4c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x6b, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x6b, 0x4e,
	0x61, 0x6d, 0x65, 0x22, 0x37, 0x0a, 0x0b, 0x55, 0x6e, 0x4c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73,
	0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x4a, 0x0a, 0x10,
	0x46, 0x6f, 0x72, 0x63, 0x65, 0x4c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08,
	0x6c, 0x6f, 0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x6c, 0x6f, 0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x3a, 0x0a, 0x0e, 0x46, 0x6f, 0x72, 0x63,
	0x65, 0x4c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6d, 0x73, 0x67, 0x22, 0x30, 0x0a, 0x12, 0x46, 0x6f, 0x72, 0x63, 0x65, 0x55, 0x6e, 0x4c,
	0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f,
	0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f,
	0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x3c, 0x0a, 0x10, 0x46, 0x6f, 0x72, 0x63, 0x65, 0x55,
	0x6e, 0x4c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6d, 0x73, 0x67, 0x32, 0xc7, 0x01, 0x0a, 0x0b, 0x4c, 0x6f, 0x63, 0x6b, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x22, 0x0a, 0x04, 0x4c, 0x6f, 0x63, 0x6b, 0x12, 0x0c, 0x2e, 0x4c,
	0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0a, 0x2e, 0x4c, 0x6f, 0x63,
	0x6b, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x28, 0x0a, 0x06, 0x55, 0x6e, 0x4c, 0x6f,
	0x63, 0x6b, 0x12, 0x0e, 0x2e, 0x55, 0x6e, 0x4c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x55, 0x6e, 0x4c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x00, 0x12, 0x31, 0x0a, 0x09, 0x46, 0x6f, 0x72, 0x63, 0x65, 0x4c, 0x6f, 0x63, 0x6b, 0x12,
	0x11, 0x2e, 0x46, 0x6f, 0x72, 0x63, 0x65, 0x4c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x46, 0x6f, 0x72, 0x63, 0x65, 0x4c, 0x6f, 0x63, 0x6b, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x0b, 0x46, 0x6f, 0x72, 0x63, 0x65, 0x55, 0x6e,
	0x4c, 0x6f, 0x63, 0x6b, 0x12, 0x13, 0x2e, 0x46, 0x6f, 0x72, 0x63, 0x65, 0x55, 0x6e, 0x4c, 0x6f,
	0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x46, 0x6f, 0x72, 0x63,
	0x65, 0x55, 0x6e, 0x4c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x0b,
	0x5a, 0x09, 0x2e, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_lock_proto_rawDescOnce sync.Once
	file_lock_proto_rawDescData = file_lock_proto_rawDesc
)

func file_lock_proto_rawDescGZIP() []byte {
	file_lock_proto_rawDescOnce.Do(func() {
		file_lock_proto_rawDescData = protoimpl.X.CompressGZIP(file_lock_proto_rawDescData)
	})
	return file_lock_proto_rawDescData
}

var file_lock_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_lock_proto_goTypes = []any{
	(*LockRequest)(nil),        // 0: LockRequest
	(*LockReply)(nil),          // 1: LockReply
	(*UnLockRequest)(nil),      // 2: UnLockRequest
	(*UnLockReply)(nil),        // 3: UnLockReply
	(*ForceLockRequest)(nil),   // 4: ForceLockRequest
	(*ForceLockReply)(nil),     // 5: ForceLockReply
	(*ForceUnLockRequest)(nil), // 6: ForceUnLockRequest
	(*ForceUnLockReply)(nil),   // 7: ForceUnLockReply
}
var file_lock_proto_depIdxs = []int32{
	0, // 0: LockService.Lock:input_type -> LockRequest
	2, // 1: LockService.UnLock:input_type -> UnLockRequest
	4, // 2: LockService.ForceLock:input_type -> ForceLockRequest
	6, // 3: LockService.ForceUnLock:input_type -> ForceUnLockRequest
	1, // 4: LockService.Lock:output_type -> LockReply
	3, // 5: LockService.UnLock:output_type -> UnLockReply
	5, // 6: LockService.ForceLock:output_type -> ForceLockReply
	7, // 7: LockService.ForceUnLock:output_type -> ForceUnLockReply
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_lock_proto_init() }
func file_lock_proto_init() {
	if File_lock_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_lock_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*LockRequest); i {
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
		file_lock_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*LockReply); i {
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
		file_lock_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*UnLockRequest); i {
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
		file_lock_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*UnLockReply); i {
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
		file_lock_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*ForceLockRequest); i {
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
		file_lock_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*ForceLockReply); i {
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
		file_lock_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*ForceUnLockRequest); i {
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
		file_lock_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*ForceUnLockReply); i {
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
			RawDescriptor: file_lock_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_lock_proto_goTypes,
		DependencyIndexes: file_lock_proto_depIdxs,
		MessageInfos:      file_lock_proto_msgTypes,
	}.Build()
	File_lock_proto = out.File
	file_lock_proto_rawDesc = nil
	file_lock_proto_goTypes = nil
	file_lock_proto_depIdxs = nil
}
