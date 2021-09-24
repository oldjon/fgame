// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.11.4
// source: cmd.proto

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

//--------------------玩家指令 start--------------------
type UCmd int32

const (
	UCmd_UCmdNone       UCmd = 0
	UCmd_HeartBeat      UCmd = 1  //心跳
	UCmd_GatewayLogin   UCmd = 2  //登录
	UCmd_ServiceOnline  UCmd = 3  //上线
	UCmd_ServiceOffline UCmd = 4  //下线
	UCmd_Move           UCmd = 20 //移动
)

// Enum value maps for UCmd.
var (
	UCmd_name = map[int32]string{
		0:  "UCmdNone",
		1:  "HeartBeat",
		2:  "GatewayLogin",
		3:  "ServiceOnline",
		4:  "ServiceOffline",
		20: "Move",
	}
	UCmd_value = map[string]int32{
		"UCmdNone":       0,
		"HeartBeat":      1,
		"GatewayLogin":   2,
		"ServiceOnline":  3,
		"ServiceOffline": 4,
		"Move":           20,
	}
)

func (x UCmd) Enum() *UCmd {
	p := new(UCmd)
	*p = x
	return p
}

func (x UCmd) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UCmd) Descriptor() protoreflect.EnumDescriptor {
	return file_cmd_proto_enumTypes[0].Descriptor()
}

func (UCmd) Type() protoreflect.EnumType {
	return &file_cmd_proto_enumTypes[0]
}

func (x UCmd) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UCmd.Descriptor instead.
func (UCmd) EnumDescriptor() ([]byte, []int) {
	return file_cmd_proto_rawDescGZIP(), []int{0}
}

//--------------------服务器指令 start--------------------
type SCmd int32

const (
	SCmd_SCmdNone     SCmd = 0
	SCmd_ServerRegist SCmd = 1 //服务器注册
	SCmd_Player       SCmd = 2
)

// Enum value maps for SCmd.
var (
	SCmd_name = map[int32]string{
		0: "SCmdNone",
		1: "ServerRegist",
		2: "Player",
	}
	SCmd_value = map[string]int32{
		"SCmdNone":     0,
		"ServerRegist": 1,
		"Player":       2,
	}
)

func (x SCmd) Enum() *SCmd {
	p := new(SCmd)
	*p = x
	return p
}

func (x SCmd) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SCmd) Descriptor() protoreflect.EnumDescriptor {
	return file_cmd_proto_enumTypes[1].Descriptor()
}

func (SCmd) Type() protoreflect.EnumType {
	return &file_cmd_proto_enumTypes[1]
}

func (x SCmd) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SCmd.Descriptor instead.
func (SCmd) EnumDescriptor() ([]byte, []int) {
	return file_cmd_proto_rawDescGZIP(), []int{1}
}

type ReqGatewayLogin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key     string `protobuf:"bytes,1,opt,name=Key,proto3" json:"Key,omitempty"`
	Version string `protobuf:"bytes,2,opt,name=Version,proto3" json:"Version,omitempty"`
}

func (x *ReqGatewayLogin) Reset() {
	*x = ReqGatewayLogin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqGatewayLogin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqGatewayLogin) ProtoMessage() {}

func (x *ReqGatewayLogin) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqGatewayLogin.ProtoReflect.Descriptor instead.
func (*ReqGatewayLogin) Descriptor() ([]byte, []int) {
	return file_cmd_proto_rawDescGZIP(), []int{0}
}

func (x *ReqGatewayLogin) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *ReqGatewayLogin) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

var File_cmd_proto protoreflect.FileDescriptor

var file_cmd_proto_rawDesc = []byte{
	0x0a, 0x09, 0x63, 0x6d, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x66, 0x63, 0x6d,
	0x64, 0x22, 0x3d, 0x0a, 0x0f, 0x52, 0x65, 0x71, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x4c,
	0x6f, 0x67, 0x69, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x4b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x4b, 0x65, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x2a, 0x66, 0x0a, 0x04, 0x55, 0x43, 0x6d, 0x64, 0x12, 0x0c, 0x0a, 0x08, 0x55, 0x43, 0x6d, 0x64,
	0x4e, 0x6f, 0x6e, 0x65, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x48, 0x65, 0x61, 0x72, 0x74, 0x42,
	0x65, 0x61, 0x74, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x10, 0x02, 0x12, 0x11, 0x0a, 0x0d, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x10, 0x03, 0x12, 0x12, 0x0a, 0x0e, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x4f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65, 0x10, 0x04, 0x12, 0x08,
	0x0a, 0x04, 0x4d, 0x6f, 0x76, 0x65, 0x10, 0x14, 0x2a, 0x32, 0x0a, 0x04, 0x53, 0x43, 0x6d, 0x64,
	0x12, 0x0c, 0x0a, 0x08, 0x53, 0x43, 0x6d, 0x64, 0x4e, 0x6f, 0x6e, 0x65, 0x10, 0x00, 0x12, 0x10,
	0x0a, 0x0c, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x10, 0x01,
	0x12, 0x0a, 0x0a, 0x06, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x10, 0x02, 0x42, 0x08, 0x5a, 0x06,
	0x2e, 0x2f, 0x66, 0x63, 0x6d, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cmd_proto_rawDescOnce sync.Once
	file_cmd_proto_rawDescData = file_cmd_proto_rawDesc
)

func file_cmd_proto_rawDescGZIP() []byte {
	file_cmd_proto_rawDescOnce.Do(func() {
		file_cmd_proto_rawDescData = protoimpl.X.CompressGZIP(file_cmd_proto_rawDescData)
	})
	return file_cmd_proto_rawDescData
}

var file_cmd_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_cmd_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_cmd_proto_goTypes = []interface{}{
	(UCmd)(0),               // 0: fcmd.UCmd
	(SCmd)(0),               // 1: fcmd.SCmd
	(*ReqGatewayLogin)(nil), // 2: fcmd.ReqGatewayLogin
}
var file_cmd_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_cmd_proto_init() }
func file_cmd_proto_init() {
	if File_cmd_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cmd_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqGatewayLogin); i {
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
			RawDescriptor: file_cmd_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cmd_proto_goTypes,
		DependencyIndexes: file_cmd_proto_depIdxs,
		EnumInfos:         file_cmd_proto_enumTypes,
		MessageInfos:      file_cmd_proto_msgTypes,
	}.Build()
	File_cmd_proto = out.File
	file_cmd_proto_rawDesc = nil
	file_cmd_proto_goTypes = nil
	file_cmd_proto_depIdxs = nil
}
