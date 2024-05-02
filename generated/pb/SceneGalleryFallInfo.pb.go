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
// source: proto/SceneGalleryFallInfo.proto

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

// Obf: IIBDBKMHMBI
type SceneGalleryFallInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ScenePlayerFallInfoMap map[uint32]*FallPlayerInfo `protobuf:"bytes,5,rep,name=scene_player_fall_info_map,json=scenePlayerFallInfoMap,proto3" json:"scene_player_fall_info_map,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	EndTime                uint32                     `protobuf:"varint,7,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
}

func (x *SceneGalleryFallInfo) Reset() {
	*x = SceneGalleryFallInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_SceneGalleryFallInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SceneGalleryFallInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SceneGalleryFallInfo) ProtoMessage() {}

func (x *SceneGalleryFallInfo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_SceneGalleryFallInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SceneGalleryFallInfo.ProtoReflect.Descriptor instead.
func (*SceneGalleryFallInfo) Descriptor() ([]byte, []int) {
	return file_proto_SceneGalleryFallInfo_proto_rawDescGZIP(), []int{0}
}

func (x *SceneGalleryFallInfo) GetScenePlayerFallInfoMap() map[uint32]*FallPlayerInfo {
	if x != nil {
		return x.ScenePlayerFallInfoMap
	}
	return nil
}

func (x *SceneGalleryFallInfo) GetEndTime() uint32 {
	if x != nil {
		return x.EndTime
	}
	return 0
}

var File_proto_SceneGalleryFallInfo_proto protoreflect.FileDescriptor

var file_proto_SceneGalleryFallInfo_proto_rawDesc = []byte{
	0x0a, 0x20, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x47, 0x61, 0x6c,
	0x6c, 0x65, 0x72, 0x79, 0x46, 0x61, 0x6c, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x46, 0x61, 0x6c, 0x6c, 0x50, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfc,
	0x01, 0x0a, 0x14, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x47, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x79, 0x46,
	0x61, 0x6c, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x6d, 0x0a, 0x1a, 0x73, 0x63, 0x65, 0x6e, 0x65,
	0x5f, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x66, 0x61, 0x6c, 0x6c, 0x5f, 0x69, 0x6e, 0x66,
	0x6f, 0x5f, 0x6d, 0x61, 0x70, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x53, 0x63,
	0x65, 0x6e, 0x65, 0x47, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x79, 0x46, 0x61, 0x6c, 0x6c, 0x49, 0x6e,
	0x66, 0x6f, 0x2e, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x46, 0x61,
	0x6c, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x16,
	0x73, 0x63, 0x65, 0x6e, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x46, 0x61, 0x6c, 0x6c, 0x49,
	0x6e, 0x66, 0x6f, 0x4d, 0x61, 0x70, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d,
	0x65, 0x1a, 0x5a, 0x0a, 0x1b, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x46, 0x61, 0x6c, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x25, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0f, 0x2e, 0x46, 0x61, 0x6c, 0x6c, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x1d, 0x5a,
	0x1b, 0x47, 0x6f, 0x2d, 0x47, 0x72, 0x61, 0x73, 0x73, 0x63, 0x75, 0x74, 0x74, 0x65, 0x72, 0x2f,
	0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_SceneGalleryFallInfo_proto_rawDescOnce sync.Once
	file_proto_SceneGalleryFallInfo_proto_rawDescData = file_proto_SceneGalleryFallInfo_proto_rawDesc
)

func file_proto_SceneGalleryFallInfo_proto_rawDescGZIP() []byte {
	file_proto_SceneGalleryFallInfo_proto_rawDescOnce.Do(func() {
		file_proto_SceneGalleryFallInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_SceneGalleryFallInfo_proto_rawDescData)
	})
	return file_proto_SceneGalleryFallInfo_proto_rawDescData
}

var file_proto_SceneGalleryFallInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_SceneGalleryFallInfo_proto_goTypes = []interface{}{
	(*SceneGalleryFallInfo)(nil), // 0: SceneGalleryFallInfo
	nil,                          // 1: SceneGalleryFallInfo.ScenePlayerFallInfoMapEntry
	(*FallPlayerInfo)(nil),       // 2: FallPlayerInfo
}
var file_proto_SceneGalleryFallInfo_proto_depIdxs = []int32{
	1, // 0: SceneGalleryFallInfo.scene_player_fall_info_map:type_name -> SceneGalleryFallInfo.ScenePlayerFallInfoMapEntry
	2, // 1: SceneGalleryFallInfo.ScenePlayerFallInfoMapEntry.value:type_name -> FallPlayerInfo
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_SceneGalleryFallInfo_proto_init() }
func file_proto_SceneGalleryFallInfo_proto_init() {
	if File_proto_SceneGalleryFallInfo_proto != nil {
		return
	}
	file_proto_FallPlayerInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_SceneGalleryFallInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SceneGalleryFallInfo); i {
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
			RawDescriptor: file_proto_SceneGalleryFallInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_SceneGalleryFallInfo_proto_goTypes,
		DependencyIndexes: file_proto_SceneGalleryFallInfo_proto_depIdxs,
		MessageInfos:      file_proto_SceneGalleryFallInfo_proto_msgTypes,
	}.Build()
	File_proto_SceneGalleryFallInfo_proto = out.File
	file_proto_SceneGalleryFallInfo_proto_rawDesc = nil
	file_proto_SceneGalleryFallInfo_proto_goTypes = nil
	file_proto_SceneGalleryFallInfo_proto_depIdxs = nil
}
