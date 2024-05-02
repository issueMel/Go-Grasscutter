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
// source: proto/PlayerHomeCompInfo.proto

package proto

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

// Obf: GJLOEDNEKNN
type PlayerHomeCompInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PDAIEFBCLNI           []uint32              `protobuf:"varint,3,rep,packed,name=PDAIEFBCLNI,proto3" json:"PDAIEFBCLNI,omitempty"`
	GPEAAFAIGAF           []uint32              `protobuf:"varint,10,rep,packed,name=GPEAAFAIGAF,proto3" json:"GPEAAFAIGAF,omitempty"`
	SeenModuleIdList      []uint32              `protobuf:"varint,8,rep,packed,name=seen_module_id_list,json=seenModuleIdList,proto3" json:"seen_module_id_list,omitempty"`
	FriendEnterHomeOption FriendEnterHomeOption `protobuf:"varint,15,opt,name=friend_enter_home_option,json=friendEnterHomeOption,proto3,enum=FriendEnterHomeOption" json:"friend_enter_home_option,omitempty"`
}

func (x *PlayerHomeCompInfo) Reset() {
	*x = PlayerHomeCompInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_PlayerHomeCompInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerHomeCompInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerHomeCompInfo) ProtoMessage() {}

func (x *PlayerHomeCompInfo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_PlayerHomeCompInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerHomeCompInfo.ProtoReflect.Descriptor instead.
func (*PlayerHomeCompInfo) Descriptor() ([]byte, []int) {
	return file_proto_PlayerHomeCompInfo_proto_rawDescGZIP(), []int{0}
}

func (x *PlayerHomeCompInfo) GetPDAIEFBCLNI() []uint32 {
	if x != nil {
		return x.PDAIEFBCLNI
	}
	return nil
}

func (x *PlayerHomeCompInfo) GetGPEAAFAIGAF() []uint32 {
	if x != nil {
		return x.GPEAAFAIGAF
	}
	return nil
}

func (x *PlayerHomeCompInfo) GetSeenModuleIdList() []uint32 {
	if x != nil {
		return x.SeenModuleIdList
	}
	return nil
}

func (x *PlayerHomeCompInfo) GetFriendEnterHomeOption() FriendEnterHomeOption {
	if x != nil {
		return x.FriendEnterHomeOption
	}
	return FriendEnterHomeOption_FRIEND_ENTER_HOME_OPTION_NEED_CONFIRM
}

var File_proto_PlayerHomeCompInfo_proto protoreflect.FileDescriptor

var file_proto_PlayerHomeCompInfo_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x48, 0x6f,
	0x6d, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x21, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x45, 0x6e,
	0x74, 0x65, 0x72, 0x48, 0x6f, 0x6d, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xd8, 0x01, 0x0a, 0x12, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x48, 0x6f,
	0x6d, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x20, 0x0a, 0x0b, 0x50, 0x44,
	0x41, 0x49, 0x45, 0x46, 0x42, 0x43, 0x4c, 0x4e, 0x49, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0d, 0x52,
	0x0b, 0x50, 0x44, 0x41, 0x49, 0x45, 0x46, 0x42, 0x43, 0x4c, 0x4e, 0x49, 0x12, 0x20, 0x0a, 0x0b,
	0x47, 0x50, 0x45, 0x41, 0x41, 0x46, 0x41, 0x49, 0x47, 0x41, 0x46, 0x18, 0x0a, 0x20, 0x03, 0x28,
	0x0d, 0x52, 0x0b, 0x47, 0x50, 0x45, 0x41, 0x41, 0x46, 0x41, 0x49, 0x47, 0x41, 0x46, 0x12, 0x2d,
	0x0a, 0x13, 0x73, 0x65, 0x65, 0x6e, 0x5f, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x69, 0x64,
	0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x10, 0x73, 0x65, 0x65,
	0x6e, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x49, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x4f, 0x0a,
	0x18, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x5f, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x5f, 0x68, 0x6f,
	0x6d, 0x65, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x16, 0x2e, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x48, 0x6f, 0x6d,
	0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x15, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x45,
	0x6e, 0x74, 0x65, 0x72, 0x48, 0x6f, 0x6d, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x20,
	0x5a, 0x1e, 0x47, 0x6f, 0x2d, 0x47, 0x72, 0x61, 0x73, 0x73, 0x63, 0x75, 0x74, 0x74, 0x65, 0x72,
	0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_PlayerHomeCompInfo_proto_rawDescOnce sync.Once
	file_proto_PlayerHomeCompInfo_proto_rawDescData = file_proto_PlayerHomeCompInfo_proto_rawDesc
)

func file_proto_PlayerHomeCompInfo_proto_rawDescGZIP() []byte {
	file_proto_PlayerHomeCompInfo_proto_rawDescOnce.Do(func() {
		file_proto_PlayerHomeCompInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_PlayerHomeCompInfo_proto_rawDescData)
	})
	return file_proto_PlayerHomeCompInfo_proto_rawDescData
}

var file_proto_PlayerHomeCompInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_PlayerHomeCompInfo_proto_goTypes = []interface{}{
	(*PlayerHomeCompInfo)(nil), // 0: PlayerHomeCompInfo
	(FriendEnterHomeOption)(0), // 1: FriendEnterHomeOption
}
var file_proto_PlayerHomeCompInfo_proto_depIdxs = []int32{
	1, // 0: PlayerHomeCompInfo.friend_enter_home_option:type_name -> FriendEnterHomeOption
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_PlayerHomeCompInfo_proto_init() }
func file_proto_PlayerHomeCompInfo_proto_init() {
	if File_proto_PlayerHomeCompInfo_proto != nil {
		return
	}
	file_proto_FriendEnterHomeOption_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_PlayerHomeCompInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerHomeCompInfo); i {
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
			RawDescriptor: file_proto_PlayerHomeCompInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_PlayerHomeCompInfo_proto_goTypes,
		DependencyIndexes: file_proto_PlayerHomeCompInfo_proto_depIdxs,
		MessageInfos:      file_proto_PlayerHomeCompInfo_proto_msgTypes,
	}.Build()
	File_proto_PlayerHomeCompInfo_proto = out.File
	file_proto_PlayerHomeCompInfo_proto_rawDesc = nil
	file_proto_PlayerHomeCompInfo_proto_goTypes = nil
	file_proto_PlayerHomeCompInfo_proto_depIdxs = nil
}
