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
// source: proto/AvatarEnterSceneInfo.proto

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

// Obf: AJDGFFHCOCL
type AvatarEnterSceneInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AvatarAbilityInfo *AbilitySyncStateInfo `protobuf:"bytes,8,opt,name=avatar_ability_info,json=avatarAbilityInfo,proto3" json:"avatar_ability_info,omitempty"`
	WeaponGuid        uint64                `protobuf:"varint,1,opt,name=weapon_guid,json=weaponGuid,proto3" json:"weapon_guid,omitempty"`
	AvatarEntityId    uint32                `protobuf:"varint,5,opt,name=avatar_entity_id,json=avatarEntityId,proto3" json:"avatar_entity_id,omitempty"`
	BuffIdList        []uint32              `protobuf:"varint,4,rep,packed,name=buff_id_list,json=buffIdList,proto3" json:"buff_id_list,omitempty"`
	WeaponEntityId    uint32                `protobuf:"varint,3,opt,name=weapon_entity_id,json=weaponEntityId,proto3" json:"weapon_entity_id,omitempty"`
	AvatarGuid        uint64                `protobuf:"varint,9,opt,name=avatar_guid,json=avatarGuid,proto3" json:"avatar_guid,omitempty"`
	ServerBuffList    []*ServerBuff         `protobuf:"bytes,10,rep,name=server_buff_list,json=serverBuffList,proto3" json:"server_buff_list,omitempty"`
	WeaponAbilityInfo *AbilitySyncStateInfo `protobuf:"bytes,15,opt,name=weapon_ability_info,json=weaponAbilityInfo,proto3" json:"weapon_ability_info,omitempty"`
}

func (x *AvatarEnterSceneInfo) Reset() {
	*x = AvatarEnterSceneInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_AvatarEnterSceneInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AvatarEnterSceneInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AvatarEnterSceneInfo) ProtoMessage() {}

func (x *AvatarEnterSceneInfo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_AvatarEnterSceneInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AvatarEnterSceneInfo.ProtoReflect.Descriptor instead.
func (*AvatarEnterSceneInfo) Descriptor() ([]byte, []int) {
	return file_proto_AvatarEnterSceneInfo_proto_rawDescGZIP(), []int{0}
}

func (x *AvatarEnterSceneInfo) GetAvatarAbilityInfo() *AbilitySyncStateInfo {
	if x != nil {
		return x.AvatarAbilityInfo
	}
	return nil
}

func (x *AvatarEnterSceneInfo) GetWeaponGuid() uint64 {
	if x != nil {
		return x.WeaponGuid
	}
	return 0
}

func (x *AvatarEnterSceneInfo) GetAvatarEntityId() uint32 {
	if x != nil {
		return x.AvatarEntityId
	}
	return 0
}

func (x *AvatarEnterSceneInfo) GetBuffIdList() []uint32 {
	if x != nil {
		return x.BuffIdList
	}
	return nil
}

func (x *AvatarEnterSceneInfo) GetWeaponEntityId() uint32 {
	if x != nil {
		return x.WeaponEntityId
	}
	return 0
}

func (x *AvatarEnterSceneInfo) GetAvatarGuid() uint64 {
	if x != nil {
		return x.AvatarGuid
	}
	return 0
}

func (x *AvatarEnterSceneInfo) GetServerBuffList() []*ServerBuff {
	if x != nil {
		return x.ServerBuffList
	}
	return nil
}

func (x *AvatarEnterSceneInfo) GetWeaponAbilityInfo() *AbilitySyncStateInfo {
	if x != nil {
		return x.WeaponAbilityInfo
	}
	return nil
}

var File_proto_AvatarEnterSceneInfo_proto protoreflect.FileDescriptor

var file_proto_AvatarEnterSceneInfo_proto_rawDesc = []byte{
	0x0a, 0x20, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x45, 0x6e,
	0x74, 0x65, 0x72, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x20, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x41, 0x62, 0x69, 0x6c, 0x69, 0x74,
	0x79, 0x53, 0x79, 0x6e, 0x63, 0x53, 0x74, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x42, 0x75, 0x66, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x93, 0x03, 0x0a,
	0x14, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x53, 0x63, 0x65, 0x6e,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x45, 0x0a, 0x13, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f,
	0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x15, 0x2e, 0x41, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x53, 0x79, 0x6e, 0x63,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x11, 0x61, 0x76, 0x61, 0x74, 0x61,
	0x72, 0x41, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1f, 0x0a, 0x0b,
	0x77, 0x65, 0x61, 0x70, 0x6f, 0x6e, 0x5f, 0x67, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x0a, 0x77, 0x65, 0x61, 0x70, 0x6f, 0x6e, 0x47, 0x75, 0x69, 0x64, 0x12, 0x28, 0x0a,
	0x10, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x69,
	0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x45,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0c, 0x62, 0x75, 0x66, 0x66, 0x5f,
	0x69, 0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0a, 0x62,
	0x75, 0x66, 0x66, 0x49, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x10, 0x77, 0x65, 0x61,
	0x70, 0x6f, 0x6e, 0x5f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0e, 0x77, 0x65, 0x61, 0x70, 0x6f, 0x6e, 0x45, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x67, 0x75,
	0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72,
	0x47, 0x75, 0x69, 0x64, 0x12, 0x35, 0x0a, 0x10, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x62,
	0x75, 0x66, 0x66, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b,
	0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x42, 0x75, 0x66, 0x66, 0x52, 0x0e, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x42, 0x75, 0x66, 0x66, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x45, 0x0a, 0x13, 0x77,
	0x65, 0x61, 0x70, 0x6f, 0x6e, 0x5f, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x6e,
	0x66, 0x6f, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x41, 0x62, 0x69, 0x6c, 0x69,
	0x74, 0x79, 0x53, 0x79, 0x6e, 0x63, 0x53, 0x74, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x11, 0x77, 0x65, 0x61, 0x70, 0x6f, 0x6e, 0x41, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x49, 0x6e,
	0x66, 0x6f, 0x42, 0x20, 0x5a, 0x1e, 0x47, 0x6f, 0x2d, 0x47, 0x72, 0x61, 0x73, 0x73, 0x63, 0x75,
	0x74, 0x74, 0x65, 0x72, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_AvatarEnterSceneInfo_proto_rawDescOnce sync.Once
	file_proto_AvatarEnterSceneInfo_proto_rawDescData = file_proto_AvatarEnterSceneInfo_proto_rawDesc
)

func file_proto_AvatarEnterSceneInfo_proto_rawDescGZIP() []byte {
	file_proto_AvatarEnterSceneInfo_proto_rawDescOnce.Do(func() {
		file_proto_AvatarEnterSceneInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_AvatarEnterSceneInfo_proto_rawDescData)
	})
	return file_proto_AvatarEnterSceneInfo_proto_rawDescData
}

var file_proto_AvatarEnterSceneInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_AvatarEnterSceneInfo_proto_goTypes = []interface{}{
	(*AvatarEnterSceneInfo)(nil), // 0: AvatarEnterSceneInfo
	(*AbilitySyncStateInfo)(nil), // 1: AbilitySyncStateInfo
	(*ServerBuff)(nil),           // 2: ServerBuff
}
var file_proto_AvatarEnterSceneInfo_proto_depIdxs = []int32{
	1, // 0: AvatarEnterSceneInfo.avatar_ability_info:type_name -> AbilitySyncStateInfo
	2, // 1: AvatarEnterSceneInfo.server_buff_list:type_name -> ServerBuff
	1, // 2: AvatarEnterSceneInfo.weapon_ability_info:type_name -> AbilitySyncStateInfo
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_proto_AvatarEnterSceneInfo_proto_init() }
func file_proto_AvatarEnterSceneInfo_proto_init() {
	if File_proto_AvatarEnterSceneInfo_proto != nil {
		return
	}
	file_proto_AbilitySyncStateInfo_proto_init()
	file_proto_ServerBuff_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_AvatarEnterSceneInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AvatarEnterSceneInfo); i {
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
			RawDescriptor: file_proto_AvatarEnterSceneInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_AvatarEnterSceneInfo_proto_goTypes,
		DependencyIndexes: file_proto_AvatarEnterSceneInfo_proto_depIdxs,
		MessageInfos:      file_proto_AvatarEnterSceneInfo_proto_msgTypes,
	}.Build()
	File_proto_AvatarEnterSceneInfo_proto = out.File
	file_proto_AvatarEnterSceneInfo_proto_rawDesc = nil
	file_proto_AvatarEnterSceneInfo_proto_goTypes = nil
	file_proto_AvatarEnterSceneInfo_proto_depIdxs = nil
}
