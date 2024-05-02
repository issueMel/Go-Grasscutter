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
// source: proto/LAOBOIBJGOK.proto

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

type LAOBOIBJGOK struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InstancedAbilityId  uint32         `protobuf:"varint,5,opt,name=instanced_ability_id,json=instancedAbilityId,proto3" json:"instanced_ability_id,omitempty"`
	LocalId             int32          `protobuf:"varint,10,opt,name=local_id,json=localId,proto3" json:"local_id,omitempty"`
	InstancedModifierId uint32         `protobuf:"varint,7,opt,name=instanced_modifier_id,json=instancedModifierId,proto3" json:"instanced_modifier_id,omitempty"`
	ParentAbilityName   *AbilityString `protobuf:"bytes,6,opt,name=parent_ability_name,json=parentAbilityName,proto3" json:"parent_ability_name,omitempty"`
	CasterId            uint32         `protobuf:"varint,14,opt,name=caster_id,json=casterId,proto3" json:"caster_id,omitempty"`
	ModifierLocalId     int32          `protobuf:"varint,12,opt,name=modifier_local_id,json=modifierLocalId,proto3" json:"modifier_local_id,omitempty"`
}

func (x *LAOBOIBJGOK) Reset() {
	*x = LAOBOIBJGOK{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_LAOBOIBJGOK_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LAOBOIBJGOK) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LAOBOIBJGOK) ProtoMessage() {}

func (x *LAOBOIBJGOK) ProtoReflect() protoreflect.Message {
	mi := &file_proto_LAOBOIBJGOK_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LAOBOIBJGOK.ProtoReflect.Descriptor instead.
func (*LAOBOIBJGOK) Descriptor() ([]byte, []int) {
	return file_proto_LAOBOIBJGOK_proto_rawDescGZIP(), []int{0}
}

func (x *LAOBOIBJGOK) GetInstancedAbilityId() uint32 {
	if x != nil {
		return x.InstancedAbilityId
	}
	return 0
}

func (x *LAOBOIBJGOK) GetLocalId() int32 {
	if x != nil {
		return x.LocalId
	}
	return 0
}

func (x *LAOBOIBJGOK) GetInstancedModifierId() uint32 {
	if x != nil {
		return x.InstancedModifierId
	}
	return 0
}

func (x *LAOBOIBJGOK) GetParentAbilityName() *AbilityString {
	if x != nil {
		return x.ParentAbilityName
	}
	return nil
}

func (x *LAOBOIBJGOK) GetCasterId() uint32 {
	if x != nil {
		return x.CasterId
	}
	return 0
}

func (x *LAOBOIBJGOK) GetModifierLocalId() int32 {
	if x != nil {
		return x.ModifierLocalId
	}
	return 0
}

var File_proto_LAOBOIBJGOK_proto protoreflect.FileDescriptor

var file_proto_LAOBOIBJGOK_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x4c, 0x41, 0x4f, 0x42, 0x4f, 0x49, 0x42, 0x4a,
	0x47, 0x4f, 0x4b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x41, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x97, 0x02, 0x0a, 0x0b, 0x4c, 0x41, 0x4f, 0x42, 0x4f, 0x49, 0x42,
	0x4a, 0x47, 0x4f, 0x4b, 0x12, 0x30, 0x0a, 0x14, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65,
	0x64, 0x5f, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x12, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x64, 0x41, 0x62, 0x69,
	0x6c, 0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f,
	0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x49,
	0x64, 0x12, 0x32, 0x0a, 0x15, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x64, 0x5f, 0x6d,
	0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x13, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x64, 0x4d, 0x6f, 0x64, 0x69, 0x66,
	0x69, 0x65, 0x72, 0x49, 0x64, 0x12, 0x3e, 0x0a, 0x13, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f,
	0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x41, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x53, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x52, 0x11, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x41, 0x62, 0x69, 0x6c, 0x69, 0x74,
	0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x61, 0x73, 0x74, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x63, 0x61, 0x73, 0x74, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x2a, 0x0a, 0x11, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x72, 0x5f, 0x6c,
	0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x6d,
	0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x72, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x49, 0x64, 0x42, 0x20,
	0x5a, 0x1e, 0x47, 0x6f, 0x2d, 0x47, 0x72, 0x61, 0x73, 0x73, 0x63, 0x75, 0x74, 0x74, 0x65, 0x72,
	0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_LAOBOIBJGOK_proto_rawDescOnce sync.Once
	file_proto_LAOBOIBJGOK_proto_rawDescData = file_proto_LAOBOIBJGOK_proto_rawDesc
)

func file_proto_LAOBOIBJGOK_proto_rawDescGZIP() []byte {
	file_proto_LAOBOIBJGOK_proto_rawDescOnce.Do(func() {
		file_proto_LAOBOIBJGOK_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_LAOBOIBJGOK_proto_rawDescData)
	})
	return file_proto_LAOBOIBJGOK_proto_rawDescData
}

var file_proto_LAOBOIBJGOK_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_LAOBOIBJGOK_proto_goTypes = []interface{}{
	(*LAOBOIBJGOK)(nil),   // 0: LAOBOIBJGOK
	(*AbilityString)(nil), // 1: AbilityString
}
var file_proto_LAOBOIBJGOK_proto_depIdxs = []int32{
	1, // 0: LAOBOIBJGOK.parent_ability_name:type_name -> AbilityString
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_LAOBOIBJGOK_proto_init() }
func file_proto_LAOBOIBJGOK_proto_init() {
	if File_proto_LAOBOIBJGOK_proto != nil {
		return
	}
	file_proto_AbilityString_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_LAOBOIBJGOK_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LAOBOIBJGOK); i {
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
			RawDescriptor: file_proto_LAOBOIBJGOK_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_LAOBOIBJGOK_proto_goTypes,
		DependencyIndexes: file_proto_LAOBOIBJGOK_proto_depIdxs,
		MessageInfos:      file_proto_LAOBOIBJGOK_proto_msgTypes,
	}.Build()
	File_proto_LAOBOIBJGOK_proto = out.File
	file_proto_LAOBOIBJGOK_proto_rawDesc = nil
	file_proto_LAOBOIBJGOK_proto_goTypes = nil
	file_proto_LAOBOIBJGOK_proto_depIdxs = nil
}
