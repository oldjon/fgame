// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.11.4
// source: player_msg.proto

package fcmd

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

type ReqLogin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Account    string `protobuf:"bytes,1,opt,name=Account,proto3" json:"Account,omitempty"`       //账号
	Password   string `protobuf:"bytes,2,opt,name=Password,proto3" json:"Password,omitempty"`     //密码
	Device     string `protobuf:"bytes,3,opt,name=Device,proto3" json:"Device,omitempty"`         //设备号
	Version    string `protobuf:"bytes,4,opt,name=Version,proto3" json:"Version,omitempty"`       //版本
	ResVersion string `protobuf:"bytes,5,opt,name=ResVersion,proto3" json:"ResVersion,omitempty"` //底包版本
	APKVersion string `protobuf:"bytes,6,opt,name=APKVersion,proto3" json:"APKVersion,omitempty"` //apk版本
	PhoneType  string `protobuf:"bytes,7,opt,name=PhoneType,proto3" json:"PhoneType,omitempty"`   //手机型号
	Net        uint32 `protobuf:"varint,8,opt,name=Net,proto3" json:"Net,omitempty"`              //网络
	Platform   uint32 `protobuf:"varint,9,opt,name=Platform,proto3" json:"Platform,omitempty"`    //平台 安卓1 苹果2
	SystemOS   string `protobuf:"bytes,10,opt,name=SystemOS,proto3" json:"SystemOS,omitempty"`    //系统
	Channel    string `protobuf:"bytes,11,opt,name=Channel,proto3" json:"Channel,omitempty"`      //渠道
	Tel        string `protobuf:"bytes,12,opt,name=Tel,proto3" json:"Tel,omitempty"`              //电话
	Captcha    string `protobuf:"bytes,13,opt,name=Captcha,proto3" json:"Captcha,omitempty"`      //验证码
}

func (x *ReqLogin) Reset() {
	*x = ReqLogin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_player_msg_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqLogin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqLogin) ProtoMessage() {}

func (x *ReqLogin) ProtoReflect() protoreflect.Message {
	mi := &file_player_msg_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqLogin.ProtoReflect.Descriptor instead.
func (*ReqLogin) Descriptor() ([]byte, []int) {
	return file_player_msg_proto_rawDescGZIP(), []int{0}
}

func (x *ReqLogin) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *ReqLogin) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *ReqLogin) GetDevice() string {
	if x != nil {
		return x.Device
	}
	return ""
}

func (x *ReqLogin) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *ReqLogin) GetResVersion() string {
	if x != nil {
		return x.ResVersion
	}
	return ""
}

func (x *ReqLogin) GetAPKVersion() string {
	if x != nil {
		return x.APKVersion
	}
	return ""
}

func (x *ReqLogin) GetPhoneType() string {
	if x != nil {
		return x.PhoneType
	}
	return ""
}

func (x *ReqLogin) GetNet() uint32 {
	if x != nil {
		return x.Net
	}
	return 0
}

func (x *ReqLogin) GetPlatform() uint32 {
	if x != nil {
		return x.Platform
	}
	return 0
}

func (x *ReqLogin) GetSystemOS() string {
	if x != nil {
		return x.SystemOS
	}
	return ""
}

func (x *ReqLogin) GetChannel() string {
	if x != nil {
		return x.Channel
	}
	return ""
}

func (x *ReqLogin) GetTel() string {
	if x != nil {
		return x.Tel
	}
	return ""
}

func (x *ReqLogin) GetCaptcha() string {
	if x != nil {
		return x.Captcha
	}
	return ""
}

type RetLogin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId       uint64 `protobuf:"varint,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
	IsBind       bool   `protobuf:"varint,2,opt,name=IsBind,proto3" json:"IsBind,omitempty"`
	Account      string `protobuf:"bytes,3,opt,name=Account,proto3" json:"Account,omitempty"`
	GatewayAddr  string `protobuf:"bytes,4,opt,name=GatewayAddr,proto3" json:"GatewayAddr,omitempty"`
	GatewayToken string `protobuf:"bytes,5,opt,name=GatewayToken,proto3" json:"GatewayToken,omitempty"`
}

func (x *RetLogin) Reset() {
	*x = RetLogin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_player_msg_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RetLogin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RetLogin) ProtoMessage() {}

func (x *RetLogin) ProtoReflect() protoreflect.Message {
	mi := &file_player_msg_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RetLogin.ProtoReflect.Descriptor instead.
func (*RetLogin) Descriptor() ([]byte, []int) {
	return file_player_msg_proto_rawDescGZIP(), []int{1}
}

func (x *RetLogin) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *RetLogin) GetIsBind() bool {
	if x != nil {
		return x.IsBind
	}
	return false
}

func (x *RetLogin) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *RetLogin) GetGatewayAddr() string {
	if x != nil {
		return x.GatewayAddr
	}
	return ""
}

func (x *RetLogin) GetGatewayToken() string {
	if x != nil {
		return x.GatewayToken
	}
	return ""
}

type ReqMove struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId uint64 `protobuf:"varint,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
	Power  uint32 `protobuf:"varint,2,opt,name=Power,proto3" json:"Power,omitempty"` //力度
	Way    uint32 `protobuf:"varint,3,opt,name=Way,proto3" json:"Way,omitempty"`     //方向
}

func (x *ReqMove) Reset() {
	*x = ReqMove{}
	if protoimpl.UnsafeEnabled {
		mi := &file_player_msg_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqMove) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqMove) ProtoMessage() {}

func (x *ReqMove) ProtoReflect() protoreflect.Message {
	mi := &file_player_msg_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqMove.ProtoReflect.Descriptor instead.
func (*ReqMove) Descriptor() ([]byte, []int) {
	return file_player_msg_proto_rawDescGZIP(), []int{2}
}

func (x *ReqMove) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *ReqMove) GetPower() uint32 {
	if x != nil {
		return x.Power
	}
	return 0
}

func (x *ReqMove) GetWay() uint32 {
	if x != nil {
		return x.Way
	}
	return 0
}

type ReqServiceOnline struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ReqServiceOnline) Reset() {
	*x = ReqServiceOnline{}
	if protoimpl.UnsafeEnabled {
		mi := &file_player_msg_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqServiceOnline) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqServiceOnline) ProtoMessage() {}

func (x *ReqServiceOnline) ProtoReflect() protoreflect.Message {
	mi := &file_player_msg_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqServiceOnline.ProtoReflect.Descriptor instead.
func (*ReqServiceOnline) Descriptor() ([]byte, []int) {
	return file_player_msg_proto_rawDescGZIP(), []int{3}
}

type RetServiceOnline struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RetServiceOnline) Reset() {
	*x = RetServiceOnline{}
	if protoimpl.UnsafeEnabled {
		mi := &file_player_msg_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RetServiceOnline) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RetServiceOnline) ProtoMessage() {}

func (x *RetServiceOnline) ProtoReflect() protoreflect.Message {
	mi := &file_player_msg_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RetServiceOnline.ProtoReflect.Descriptor instead.
func (*RetServiceOnline) Descriptor() ([]byte, []int) {
	return file_player_msg_proto_rawDescGZIP(), []int{4}
}

type ReqServiceOffline struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ReqServiceOffline) Reset() {
	*x = ReqServiceOffline{}
	if protoimpl.UnsafeEnabled {
		mi := &file_player_msg_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqServiceOffline) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqServiceOffline) ProtoMessage() {}

func (x *ReqServiceOffline) ProtoReflect() protoreflect.Message {
	mi := &file_player_msg_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqServiceOffline.ProtoReflect.Descriptor instead.
func (*ReqServiceOffline) Descriptor() ([]byte, []int) {
	return file_player_msg_proto_rawDescGZIP(), []int{5}
}

type RetServiceOffline struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RetServiceOffline) Reset() {
	*x = RetServiceOffline{}
	if protoimpl.UnsafeEnabled {
		mi := &file_player_msg_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RetServiceOffline) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RetServiceOffline) ProtoMessage() {}

func (x *RetServiceOffline) ProtoReflect() protoreflect.Message {
	mi := &file_player_msg_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RetServiceOffline.ProtoReflect.Descriptor instead.
func (*RetServiceOffline) Descriptor() ([]byte, []int) {
	return file_player_msg_proto_rawDescGZIP(), []int{6}
}

var File_player_msg_proto protoreflect.FileDescriptor

var file_player_msg_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x6d, 0x73, 0x67, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x04, 0x66, 0x63, 0x6d, 0x64, 0x22, 0xe0, 0x02, 0x0a, 0x08, 0x52, 0x65, 0x71,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x1a, 0x0a, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x44,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x44, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a,
	0x0a, 0x52, 0x65, 0x73, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x52, 0x65, 0x73, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a,
	0x0a, 0x41, 0x50, 0x4b, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x41, 0x50, 0x4b, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a,
	0x09, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x4e,
	0x65, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x4e, 0x65, 0x74, 0x12, 0x1a, 0x0a,
	0x08, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x08, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x1a, 0x0a, 0x08, 0x53, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x4f, 0x53, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x53, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x4f, 0x53, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12,
	0x10, 0x0a, 0x03, 0x54, 0x65, 0x6c, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x54, 0x65,
	0x6c, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x18, 0x0d, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x43, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x22, 0x9a, 0x01, 0x0a, 0x08,
	0x52, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x49, 0x73, 0x42, 0x69, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x06, 0x49, 0x73, 0x42, 0x69, 0x6e, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x41, 0x64, 0x64,
	0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x41, 0x64, 0x64, 0x72, 0x12, 0x22, 0x0a, 0x0c, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x47, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x49, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x4d,
	0x6f, 0x76, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x50,
	0x6f, 0x77, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x50, 0x6f, 0x77, 0x65,
	0x72, 0x12, 0x10, 0x0a, 0x03, 0x57, 0x61, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03,
	0x57, 0x61, 0x79, 0x22, 0x12, 0x0a, 0x10, 0x52, 0x65, 0x71, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x22, 0x12, 0x0a, 0x10, 0x52, 0x65, 0x74, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x22, 0x13, 0x0a, 0x11, 0x52,
	0x65, 0x71, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65,
	0x22, 0x13, 0x0a, 0x11, 0x52, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4f, 0x66,
	0x66, 0x6c, 0x69, 0x6e, 0x65, 0x42, 0x08, 0x5a, 0x06, 0x2e, 0x2f, 0x66, 0x63, 0x6d, 0x64, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_player_msg_proto_rawDescOnce sync.Once
	file_player_msg_proto_rawDescData = file_player_msg_proto_rawDesc
)

func file_player_msg_proto_rawDescGZIP() []byte {
	file_player_msg_proto_rawDescOnce.Do(func() {
		file_player_msg_proto_rawDescData = protoimpl.X.CompressGZIP(file_player_msg_proto_rawDescData)
	})
	return file_player_msg_proto_rawDescData
}

var file_player_msg_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_player_msg_proto_goTypes = []interface{}{
	(*ReqLogin)(nil),          // 0: fcmd.ReqLogin
	(*RetLogin)(nil),          // 1: fcmd.RetLogin
	(*ReqMove)(nil),           // 2: fcmd.ReqMove
	(*ReqServiceOnline)(nil),  // 3: fcmd.ReqServiceOnline
	(*RetServiceOnline)(nil),  // 4: fcmd.RetServiceOnline
	(*ReqServiceOffline)(nil), // 5: fcmd.ReqServiceOffline
	(*RetServiceOffline)(nil), // 6: fcmd.RetServiceOffline
}
var file_player_msg_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_player_msg_proto_init() }
func file_player_msg_proto_init() {
	if File_player_msg_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_player_msg_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqLogin); i {
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
		file_player_msg_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RetLogin); i {
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
		file_player_msg_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqMove); i {
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
		file_player_msg_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqServiceOnline); i {
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
		file_player_msg_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RetServiceOnline); i {
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
		file_player_msg_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqServiceOffline); i {
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
		file_player_msg_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RetServiceOffline); i {
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
			RawDescriptor: file_player_msg_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_player_msg_proto_goTypes,
		DependencyIndexes: file_player_msg_proto_depIdxs,
		MessageInfos:      file_player_msg_proto_msgTypes,
	}.Build()
	File_player_msg_proto = out.File
	file_player_msg_proto_rawDesc = nil
	file_player_msg_proto_goTypes = nil
	file_player_msg_proto_depIdxs = nil
}
