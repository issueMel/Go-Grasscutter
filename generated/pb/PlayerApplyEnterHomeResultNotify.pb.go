// https://github.com/SlushinPS/beach-simulator
// Copyright (C) 2023 Slushy Team
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.23.1
// source: proto/PlayerApplyEnterHomeResultNotify.proto

package pb

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

// Obf: HCLKFKDJEMO
type PlayerApplyEnterHomeResultNotify_Reason int32

const (
	PlayerApplyEnterHomeResultNotify_PLAYER_JUDGE                PlayerApplyEnterHomeResultNotify_Reason = 0
	PlayerApplyEnterHomeResultNotify_PLAYER_ENTER_OPTION_REFUSE  PlayerApplyEnterHomeResultNotify_Reason = 1
	PlayerApplyEnterHomeResultNotify_PLAYER_ENTER_OPTION_DIRECT  PlayerApplyEnterHomeResultNotify_Reason = 2
	PlayerApplyEnterHomeResultNotify_SYSTEM_JUDGE                PlayerApplyEnterHomeResultNotify_Reason = 3
	PlayerApplyEnterHomeResultNotify_HOST_IN_MATCH               PlayerApplyEnterHomeResultNotify_Reason = 4
	PlayerApplyEnterHomeResultNotify_PS_PLAYER_NOT_ACCEPT_OTHERS PlayerApplyEnterHomeResultNotify_Reason = 5
	PlayerApplyEnterHomeResultNotify_OPEN_STATE_NOT_OPEN         PlayerApplyEnterHomeResultNotify_Reason = 6
	PlayerApplyEnterHomeResultNotify_HOST_IN_EDIT_MODE           PlayerApplyEnterHomeResultNotify_Reason = 7
	PlayerApplyEnterHomeResultNotify_PRIOR_CHECK                 PlayerApplyEnterHomeResultNotify_Reason = 8
)

// Enum value maps for PlayerApplyEnterHomeResultNotify_Reason.
var (
	PlayerApplyEnterHomeResultNotify_Reason_name = map[int32]string{
		0: "PLAYER_JUDGE",
		1: "PLAYER_ENTER_OPTION_REFUSE",
		2: "PLAYER_ENTER_OPTION_DIRECT",
		3: "SYSTEM_JUDGE",
		4: "HOST_IN_MATCH",
		5: "PS_PLAYER_NOT_ACCEPT_OTHERS",
		6: "OPEN_STATE_NOT_OPEN",
		7: "HOST_IN_EDIT_MODE",
		8: "PRIOR_CHECK",
	}
	PlayerApplyEnterHomeResultNotify_Reason_value = map[string]int32{
		"PLAYER_JUDGE":                0,
		"PLAYER_ENTER_OPTION_REFUSE":  1,
		"PLAYER_ENTER_OPTION_DIRECT":  2,
		"SYSTEM_JUDGE":                3,
		"HOST_IN_MATCH":               4,
		"PS_PLAYER_NOT_ACCEPT_OTHERS": 5,
		"OPEN_STATE_NOT_OPEN":         6,
		"HOST_IN_EDIT_MODE":           7,
		"PRIOR_CHECK":                 8,
	}
)

func (x PlayerApplyEnterHomeResultNotify_Reason) Enum() *PlayerApplyEnterHomeResultNotify_Reason {
	p := new(PlayerApplyEnterHomeResultNotify_Reason)
	*p = x
	return p
}

func (x PlayerApplyEnterHomeResultNotify_Reason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PlayerApplyEnterHomeResultNotify_Reason) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_PlayerApplyEnterHomeResultNotify_proto_enumTypes[0].Descriptor()
}

func (PlayerApplyEnterHomeResultNotify_Reason) Type() protoreflect.EnumType {
	return &file_proto_PlayerApplyEnterHomeResultNotify_proto_enumTypes[0]
}

func (x PlayerApplyEnterHomeResultNotify_Reason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PlayerApplyEnterHomeResultNotify_Reason.Descriptor instead.
func (PlayerApplyEnterHomeResultNotify_Reason) EnumDescriptor() ([]byte, []int) {
	return file_proto_PlayerApplyEnterHomeResultNotify_proto_rawDescGZIP(), []int{0, 0}
}

// CmdId: 26518
// Obf: BDMPJIIHGNK
type PlayerApplyEnterHomeResultNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsAgreed       bool                                    `protobuf:"varint,1,opt,name=is_agreed,json=isAgreed,proto3" json:"is_agreed,omitempty"`
	TargetNickname string                                  `protobuf:"bytes,7,opt,name=target_nickname,json=targetNickname,proto3" json:"target_nickname,omitempty"`
	TargetUid      uint32                                  `protobuf:"varint,11,opt,name=target_uid,json=targetUid,proto3" json:"target_uid,omitempty"`
	Reason         PlayerApplyEnterHomeResultNotify_Reason `protobuf:"varint,4,opt,name=reason,proto3,enum=PlayerApplyEnterHomeResultNotify_Reason" json:"reason,omitempty"`
}

func (x *PlayerApplyEnterHomeResultNotify) Reset() {
	*x = PlayerApplyEnterHomeResultNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_PlayerApplyEnterHomeResultNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerApplyEnterHomeResultNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerApplyEnterHomeResultNotify) ProtoMessage() {}

func (x *PlayerApplyEnterHomeResultNotify) ProtoReflect() protoreflect.Message {
	mi := &file_proto_PlayerApplyEnterHomeResultNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerApplyEnterHomeResultNotify.ProtoReflect.Descriptor instead.
func (*PlayerApplyEnterHomeResultNotify) Descriptor() ([]byte, []int) {
	return file_proto_PlayerApplyEnterHomeResultNotify_proto_rawDescGZIP(), []int{0}
}

func (x *PlayerApplyEnterHomeResultNotify) GetIsAgreed() bool {
	if x != nil {
		return x.IsAgreed
	}
	return false
}

func (x *PlayerApplyEnterHomeResultNotify) GetTargetNickname() string {
	if x != nil {
		return x.TargetNickname
	}
	return ""
}

func (x *PlayerApplyEnterHomeResultNotify) GetTargetUid() uint32 {
	if x != nil {
		return x.TargetUid
	}
	return 0
}

func (x *PlayerApplyEnterHomeResultNotify) GetReason() PlayerApplyEnterHomeResultNotify_Reason {
	if x != nil {
		return x.Reason
	}
	return PlayerApplyEnterHomeResultNotify_PLAYER_JUDGE
}

var File_proto_PlayerApplyEnterHomeResultNotify_proto protoreflect.FileDescriptor

var file_proto_PlayerApplyEnterHomeResultNotify_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x41, 0x70,
	0x70, 0x6c, 0x79, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x48, 0x6f, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xad,
	0x03, 0x0a, 0x20, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x45, 0x6e,
	0x74, 0x65, 0x72, 0x48, 0x6f, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x4e, 0x6f, 0x74,
	0x69, 0x66, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x61, 0x67, 0x72, 0x65, 0x65, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x41, 0x67, 0x72, 0x65, 0x65, 0x64,
	0x12, 0x27, 0x0a, 0x0f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x6e, 0x69, 0x63, 0x6b, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x74, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x4e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x5f, 0x75, 0x69, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x74,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x55, 0x69, 0x64, 0x12, 0x40, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73,
	0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x28, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x48, 0x6f, 0x6d, 0x65, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x52, 0x65, 0x61, 0x73,
	0x6f, 0x6e, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x22, 0xe1, 0x01, 0x0a, 0x06, 0x52,
	0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x0c, 0x50, 0x4c, 0x41, 0x59, 0x45, 0x52, 0x5f,
	0x4a, 0x55, 0x44, 0x47, 0x45, 0x10, 0x00, 0x12, 0x1e, 0x0a, 0x1a, 0x50, 0x4c, 0x41, 0x59, 0x45,
	0x52, 0x5f, 0x45, 0x4e, 0x54, 0x45, 0x52, 0x5f, 0x4f, 0x50, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x52,
	0x45, 0x46, 0x55, 0x53, 0x45, 0x10, 0x01, 0x12, 0x1e, 0x0a, 0x1a, 0x50, 0x4c, 0x41, 0x59, 0x45,
	0x52, 0x5f, 0x45, 0x4e, 0x54, 0x45, 0x52, 0x5f, 0x4f, 0x50, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x44,
	0x49, 0x52, 0x45, 0x43, 0x54, 0x10, 0x02, 0x12, 0x10, 0x0a, 0x0c, 0x53, 0x59, 0x53, 0x54, 0x45,
	0x4d, 0x5f, 0x4a, 0x55, 0x44, 0x47, 0x45, 0x10, 0x03, 0x12, 0x11, 0x0a, 0x0d, 0x48, 0x4f, 0x53,
	0x54, 0x5f, 0x49, 0x4e, 0x5f, 0x4d, 0x41, 0x54, 0x43, 0x48, 0x10, 0x04, 0x12, 0x1f, 0x0a, 0x1b,
	0x50, 0x53, 0x5f, 0x50, 0x4c, 0x41, 0x59, 0x45, 0x52, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x41, 0x43,
	0x43, 0x45, 0x50, 0x54, 0x5f, 0x4f, 0x54, 0x48, 0x45, 0x52, 0x53, 0x10, 0x05, 0x12, 0x17, 0x0a,
	0x13, 0x4f, 0x50, 0x45, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x4e, 0x4f, 0x54, 0x5f,
	0x4f, 0x50, 0x45, 0x4e, 0x10, 0x06, 0x12, 0x15, 0x0a, 0x11, 0x48, 0x4f, 0x53, 0x54, 0x5f, 0x49,
	0x4e, 0x5f, 0x45, 0x44, 0x49, 0x54, 0x5f, 0x4d, 0x4f, 0x44, 0x45, 0x10, 0x07, 0x12, 0x0f, 0x0a,
	0x0b, 0x50, 0x52, 0x49, 0x4f, 0x52, 0x5f, 0x43, 0x48, 0x45, 0x43, 0x4b, 0x10, 0x08, 0x42, 0x1d,
	0x5a, 0x1b, 0x47, 0x6f, 0x2d, 0x47, 0x72, 0x61, 0x73, 0x73, 0x63, 0x75, 0x74, 0x74, 0x65, 0x72,
	0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_PlayerApplyEnterHomeResultNotify_proto_rawDescOnce sync.Once
	file_proto_PlayerApplyEnterHomeResultNotify_proto_rawDescData = file_proto_PlayerApplyEnterHomeResultNotify_proto_rawDesc
)

func file_proto_PlayerApplyEnterHomeResultNotify_proto_rawDescGZIP() []byte {
	file_proto_PlayerApplyEnterHomeResultNotify_proto_rawDescOnce.Do(func() {
		file_proto_PlayerApplyEnterHomeResultNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_PlayerApplyEnterHomeResultNotify_proto_rawDescData)
	})
	return file_proto_PlayerApplyEnterHomeResultNotify_proto_rawDescData
}

var file_proto_PlayerApplyEnterHomeResultNotify_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_PlayerApplyEnterHomeResultNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_PlayerApplyEnterHomeResultNotify_proto_goTypes = []interface{}{
	(PlayerApplyEnterHomeResultNotify_Reason)(0), // 0: PlayerApplyEnterHomeResultNotify.Reason
	(*PlayerApplyEnterHomeResultNotify)(nil),     // 1: PlayerApplyEnterHomeResultNotify
}
var file_proto_PlayerApplyEnterHomeResultNotify_proto_depIdxs = []int32{
	0, // 0: PlayerApplyEnterHomeResultNotify.reason:type_name -> PlayerApplyEnterHomeResultNotify.Reason
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_PlayerApplyEnterHomeResultNotify_proto_init() }
func file_proto_PlayerApplyEnterHomeResultNotify_proto_init() {
	if File_proto_PlayerApplyEnterHomeResultNotify_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_PlayerApplyEnterHomeResultNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerApplyEnterHomeResultNotify); i {
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
			RawDescriptor: file_proto_PlayerApplyEnterHomeResultNotify_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_PlayerApplyEnterHomeResultNotify_proto_goTypes,
		DependencyIndexes: file_proto_PlayerApplyEnterHomeResultNotify_proto_depIdxs,
		EnumInfos:         file_proto_PlayerApplyEnterHomeResultNotify_proto_enumTypes,
		MessageInfos:      file_proto_PlayerApplyEnterHomeResultNotify_proto_msgTypes,
	}.Build()
	File_proto_PlayerApplyEnterHomeResultNotify_proto = out.File
	file_proto_PlayerApplyEnterHomeResultNotify_proto_rawDesc = nil
	file_proto_PlayerApplyEnterHomeResultNotify_proto_goTypes = nil
	file_proto_PlayerApplyEnterHomeResultNotify_proto_depIdxs = nil
}
