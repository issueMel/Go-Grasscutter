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
// source: proto/COEDDKCEEDK.proto

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

// CmdId: 3026
type COEDDKCEEDK struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DungeonGuid    uint64           `protobuf:"varint,7,opt,name=dungeon_guid,json=dungeonGuid,proto3" json:"dungeon_guid,omitempty"`
	RoomId         uint32           `protobuf:"varint,13,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
	AvatarInfoList []*UgcAvatarInfo `protobuf:"bytes,4,rep,name=avatar_info_list,json=avatarInfoList,proto3" json:"avatar_info_list,omitempty"`
	EnterType      FDGOOBGNJMP      `protobuf:"varint,8,opt,name=enter_type,json=enterType,proto3,enum=FDGOOBGNJMP" json:"enter_type,omitempty"`
	DungeonId      uint32           `protobuf:"varint,15,opt,name=dungeon_id,json=dungeonId,proto3" json:"dungeon_id,omitempty"`
}

func (x *COEDDKCEEDK) Reset() {
	*x = COEDDKCEEDK{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_COEDDKCEEDK_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *COEDDKCEEDK) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*COEDDKCEEDK) ProtoMessage() {}

func (x *COEDDKCEEDK) ProtoReflect() protoreflect.Message {
	mi := &file_proto_COEDDKCEEDK_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use COEDDKCEEDK.ProtoReflect.Descriptor instead.
func (*COEDDKCEEDK) Descriptor() ([]byte, []int) {
	return file_proto_COEDDKCEEDK_proto_rawDescGZIP(), []int{0}
}

func (x *COEDDKCEEDK) GetDungeonGuid() uint64 {
	if x != nil {
		return x.DungeonGuid
	}
	return 0
}

func (x *COEDDKCEEDK) GetRoomId() uint32 {
	if x != nil {
		return x.RoomId
	}
	return 0
}

func (x *COEDDKCEEDK) GetAvatarInfoList() []*UgcAvatarInfo {
	if x != nil {
		return x.AvatarInfoList
	}
	return nil
}

func (x *COEDDKCEEDK) GetEnterType() FDGOOBGNJMP {
	if x != nil {
		return x.EnterType
	}
	return FDGOOBGNJMP_FDGOOBGNJMP_EnterUgcDungeonNone
}

func (x *COEDDKCEEDK) GetDungeonId() uint32 {
	if x != nil {
		return x.DungeonId
	}
	return 0
}

var File_proto_COEDDKCEEDK_proto protoreflect.FileDescriptor

var file_proto_COEDDKCEEDK_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x43, 0x4f, 0x45, 0x44, 0x44, 0x4b, 0x43, 0x45,
	0x45, 0x44, 0x4b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x55, 0x67, 0x63, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x46, 0x44, 0x47, 0x4f,
	0x4f, 0x42, 0x47, 0x4e, 0x4a, 0x4d, 0x50, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xcf, 0x01,
	0x0a, 0x0b, 0x43, 0x4f, 0x45, 0x44, 0x44, 0x4b, 0x43, 0x45, 0x45, 0x44, 0x4b, 0x12, 0x21, 0x0a,
	0x0c, 0x64, 0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x5f, 0x67, 0x75, 0x69, 0x64, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x0b, 0x64, 0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x47, 0x75, 0x69, 0x64,
	0x12, 0x17, 0x0a, 0x07, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x12, 0x38, 0x0a, 0x10, 0x61, 0x76, 0x61,
	0x74, 0x61, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x04, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x55, 0x67, 0x63, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x0e, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x2b, 0x0a, 0x0a, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0c, 0x2e, 0x46, 0x44, 0x47, 0x4f, 0x4f, 0x42,
	0x47, 0x4e, 0x4a, 0x4d, 0x50, 0x52, 0x09, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x0f,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x64, 0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x49, 0x64, 0x42,
	0x1d, 0x5a, 0x1b, 0x47, 0x6f, 0x2d, 0x47, 0x72, 0x61, 0x73, 0x73, 0x63, 0x75, 0x74, 0x74, 0x65,
	0x72, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_COEDDKCEEDK_proto_rawDescOnce sync.Once
	file_proto_COEDDKCEEDK_proto_rawDescData = file_proto_COEDDKCEEDK_proto_rawDesc
)

func file_proto_COEDDKCEEDK_proto_rawDescGZIP() []byte {
	file_proto_COEDDKCEEDK_proto_rawDescOnce.Do(func() {
		file_proto_COEDDKCEEDK_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_COEDDKCEEDK_proto_rawDescData)
	})
	return file_proto_COEDDKCEEDK_proto_rawDescData
}

var file_proto_COEDDKCEEDK_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_COEDDKCEEDK_proto_goTypes = []interface{}{
	(*COEDDKCEEDK)(nil),   // 0: COEDDKCEEDK
	(*UgcAvatarInfo)(nil), // 1: UgcAvatarInfo
	(FDGOOBGNJMP)(0),      // 2: FDGOOBGNJMP
}
var file_proto_COEDDKCEEDK_proto_depIdxs = []int32{
	1, // 0: COEDDKCEEDK.avatar_info_list:type_name -> UgcAvatarInfo
	2, // 1: COEDDKCEEDK.enter_type:type_name -> FDGOOBGNJMP
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_COEDDKCEEDK_proto_init() }
func file_proto_COEDDKCEEDK_proto_init() {
	if File_proto_COEDDKCEEDK_proto != nil {
		return
	}
	file_proto_UgcAvatarInfo_proto_init()
	file_proto_FDGOOBGNJMP_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_COEDDKCEEDK_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*COEDDKCEEDK); i {
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
			RawDescriptor: file_proto_COEDDKCEEDK_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_COEDDKCEEDK_proto_goTypes,
		DependencyIndexes: file_proto_COEDDKCEEDK_proto_depIdxs,
		MessageInfos:      file_proto_COEDDKCEEDK_proto_msgTypes,
	}.Build()
	File_proto_COEDDKCEEDK_proto = out.File
	file_proto_COEDDKCEEDK_proto_rawDesc = nil
	file_proto_COEDDKCEEDK_proto_goTypes = nil
	file_proto_COEDDKCEEDK_proto_depIdxs = nil
}
