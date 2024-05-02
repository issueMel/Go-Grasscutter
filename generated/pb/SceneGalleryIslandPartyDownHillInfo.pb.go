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
// source: proto/SceneGalleryIslandPartyDownHillInfo.proto

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

// Obf: OIGHHDCGKOJ
type SceneGalleryIslandPartyDownHillInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TotalKillMonsterCount uint32             `protobuf:"varint,7,opt,name=total_kill_monster_count,json=totalKillMonsterCount,proto3" json:"total_kill_monster_count,omitempty"`
	Coin                  uint32             `protobuf:"varint,13,opt,name=coin,proto3" json:"coin,omitempty"`
	StartSource           GalleryStartSource `protobuf:"varint,3,opt,name=start_source,json=startSource,proto3,enum=GalleryStartSource" json:"start_source,omitempty"`
	MaxKillMonsterCount   uint32             `protobuf:"varint,14,opt,name=max_kill_monster_count,json=maxKillMonsterCount,proto3" json:"max_kill_monster_count,omitempty"`
}

func (x *SceneGalleryIslandPartyDownHillInfo) Reset() {
	*x = SceneGalleryIslandPartyDownHillInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_SceneGalleryIslandPartyDownHillInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SceneGalleryIslandPartyDownHillInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SceneGalleryIslandPartyDownHillInfo) ProtoMessage() {}

func (x *SceneGalleryIslandPartyDownHillInfo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_SceneGalleryIslandPartyDownHillInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SceneGalleryIslandPartyDownHillInfo.ProtoReflect.Descriptor instead.
func (*SceneGalleryIslandPartyDownHillInfo) Descriptor() ([]byte, []int) {
	return file_proto_SceneGalleryIslandPartyDownHillInfo_proto_rawDescGZIP(), []int{0}
}

func (x *SceneGalleryIslandPartyDownHillInfo) GetTotalKillMonsterCount() uint32 {
	if x != nil {
		return x.TotalKillMonsterCount
	}
	return 0
}

func (x *SceneGalleryIslandPartyDownHillInfo) GetCoin() uint32 {
	if x != nil {
		return x.Coin
	}
	return 0
}

func (x *SceneGalleryIslandPartyDownHillInfo) GetStartSource() GalleryStartSource {
	if x != nil {
		return x.StartSource
	}
	return GalleryStartSource_GALLERY_START_BY_NONE
}

func (x *SceneGalleryIslandPartyDownHillInfo) GetMaxKillMonsterCount() uint32 {
	if x != nil {
		return x.MaxKillMonsterCount
	}
	return 0
}

var File_proto_SceneGalleryIslandPartyDownHillInfo_proto protoreflect.FileDescriptor

var file_proto_SceneGalleryIslandPartyDownHillInfo_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x47, 0x61, 0x6c,
	0x6c, 0x65, 0x72, 0x79, 0x49, 0x73, 0x6c, 0x61, 0x6e, 0x64, 0x50, 0x61, 0x72, 0x74, 0x79, 0x44,
	0x6f, 0x77, 0x6e, 0x48, 0x69, 0x6c, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x47, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x79,
	0x53, 0x74, 0x61, 0x72, 0x74, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xdf, 0x01, 0x0a, 0x23, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x47, 0x61, 0x6c, 0x6c, 0x65,
	0x72, 0x79, 0x49, 0x73, 0x6c, 0x61, 0x6e, 0x64, 0x50, 0x61, 0x72, 0x74, 0x79, 0x44, 0x6f, 0x77,
	0x6e, 0x48, 0x69, 0x6c, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x37, 0x0a, 0x18, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x5f, 0x6b, 0x69, 0x6c, 0x6c, 0x5f, 0x6d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x5f,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x15, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x4b, 0x69, 0x6c, 0x6c, 0x4d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x69, 0x6e, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x04, 0x63, 0x6f, 0x69, 0x6e, 0x12, 0x36, 0x0a, 0x0c, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x47,
	0x61, 0x6c, 0x6c, 0x65, 0x72, 0x79, 0x53, 0x74, 0x61, 0x72, 0x74, 0x53, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x52, 0x0b, 0x73, 0x74, 0x61, 0x72, 0x74, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x33,
	0x0a, 0x16, 0x6d, 0x61, 0x78, 0x5f, 0x6b, 0x69, 0x6c, 0x6c, 0x5f, 0x6d, 0x6f, 0x6e, 0x73, 0x74,
	0x65, 0x72, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x13,
	0x6d, 0x61, 0x78, 0x4b, 0x69, 0x6c, 0x6c, 0x4d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x42, 0x1d, 0x5a, 0x1b, 0x47, 0x6f, 0x2d, 0x47, 0x72, 0x61, 0x73, 0x73, 0x63,
	0x75, 0x74, 0x74, 0x65, 0x72, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_SceneGalleryIslandPartyDownHillInfo_proto_rawDescOnce sync.Once
	file_proto_SceneGalleryIslandPartyDownHillInfo_proto_rawDescData = file_proto_SceneGalleryIslandPartyDownHillInfo_proto_rawDesc
)

func file_proto_SceneGalleryIslandPartyDownHillInfo_proto_rawDescGZIP() []byte {
	file_proto_SceneGalleryIslandPartyDownHillInfo_proto_rawDescOnce.Do(func() {
		file_proto_SceneGalleryIslandPartyDownHillInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_SceneGalleryIslandPartyDownHillInfo_proto_rawDescData)
	})
	return file_proto_SceneGalleryIslandPartyDownHillInfo_proto_rawDescData
}

var file_proto_SceneGalleryIslandPartyDownHillInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_SceneGalleryIslandPartyDownHillInfo_proto_goTypes = []interface{}{
	(*SceneGalleryIslandPartyDownHillInfo)(nil), // 0: SceneGalleryIslandPartyDownHillInfo
	(GalleryStartSource)(0),                     // 1: GalleryStartSource
}
var file_proto_SceneGalleryIslandPartyDownHillInfo_proto_depIdxs = []int32{
	1, // 0: SceneGalleryIslandPartyDownHillInfo.start_source:type_name -> GalleryStartSource
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_SceneGalleryIslandPartyDownHillInfo_proto_init() }
func file_proto_SceneGalleryIslandPartyDownHillInfo_proto_init() {
	if File_proto_SceneGalleryIslandPartyDownHillInfo_proto != nil {
		return
	}
	file_proto_GalleryStartSource_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_SceneGalleryIslandPartyDownHillInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SceneGalleryIslandPartyDownHillInfo); i {
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
			RawDescriptor: file_proto_SceneGalleryIslandPartyDownHillInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_SceneGalleryIslandPartyDownHillInfo_proto_goTypes,
		DependencyIndexes: file_proto_SceneGalleryIslandPartyDownHillInfo_proto_depIdxs,
		MessageInfos:      file_proto_SceneGalleryIslandPartyDownHillInfo_proto_msgTypes,
	}.Build()
	File_proto_SceneGalleryIslandPartyDownHillInfo_proto = out.File
	file_proto_SceneGalleryIslandPartyDownHillInfo_proto_rawDesc = nil
	file_proto_SceneGalleryIslandPartyDownHillInfo_proto_goTypes = nil
	file_proto_SceneGalleryIslandPartyDownHillInfo_proto_depIdxs = nil
}
