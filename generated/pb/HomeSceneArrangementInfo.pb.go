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
// source: proto/HomeSceneArrangementInfo.proto

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

// Obf: LEMEOONDDGM
type HomeSceneArrangementInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlockArrangementInfoList []*HomeBlockArrangementInfo `protobuf:"bytes,2,rep,name=block_arrangement_info_list,json=blockArrangementInfoList,proto3" json:"block_arrangement_info_list,omitempty"`
	HDOFAHGPJGK              []*HomeFurnitureData        `protobuf:"bytes,7,rep,name=HDOFAHGPJGK,proto3" json:"HDOFAHGPJGK,omitempty"`
	MLECLIOFALA              []*HomeFurnitureData        `protobuf:"bytes,12,rep,name=MLECLIOFALA,proto3" json:"MLECLIOFALA,omitempty"`
	BornRot                  *Vector                     `protobuf:"bytes,9,opt,name=born_rot,json=bornRot,proto3" json:"born_rot,omitempty"`
	SceneId                  uint32                      `protobuf:"varint,8,opt,name=scene_id,json=sceneId,proto3" json:"scene_id,omitempty"`
	BgmId                    uint32                      `protobuf:"varint,3,opt,name=bgm_id,json=bgmId,proto3" json:"bgm_id,omitempty"`
	ComfortValue             uint32                      `protobuf:"varint,5,opt,name=comfort_value,json=comfortValue,proto3" json:"comfort_value,omitempty"`
	MainHouse                *HomeFurnitureData          `protobuf:"bytes,6,opt,name=main_house,json=mainHouse,proto3" json:"main_house,omitempty"`
	TmpVersion               uint32                      `protobuf:"varint,13,opt,name=tmp_version,json=tmpVersion,proto3" json:"tmp_version,omitempty"`
	IsSetBornPos             bool                        `protobuf:"varint,1,opt,name=is_set_born_pos,json=isSetBornPos,proto3" json:"is_set_born_pos,omitempty"`
	DjinnPos                 *Vector                     `protobuf:"bytes,4,opt,name=djinn_pos,json=djinnPos,proto3" json:"djinn_pos,omitempty"`
	BornPos                  *Vector                     `protobuf:"bytes,10,opt,name=born_pos,json=bornPos,proto3" json:"born_pos,omitempty"`
}

func (x *HomeSceneArrangementInfo) Reset() {
	*x = HomeSceneArrangementInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_HomeSceneArrangementInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HomeSceneArrangementInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HomeSceneArrangementInfo) ProtoMessage() {}

func (x *HomeSceneArrangementInfo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_HomeSceneArrangementInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HomeSceneArrangementInfo.ProtoReflect.Descriptor instead.
func (*HomeSceneArrangementInfo) Descriptor() ([]byte, []int) {
	return file_proto_HomeSceneArrangementInfo_proto_rawDescGZIP(), []int{0}
}

func (x *HomeSceneArrangementInfo) GetBlockArrangementInfoList() []*HomeBlockArrangementInfo {
	if x != nil {
		return x.BlockArrangementInfoList
	}
	return nil
}

func (x *HomeSceneArrangementInfo) GetHDOFAHGPJGK() []*HomeFurnitureData {
	if x != nil {
		return x.HDOFAHGPJGK
	}
	return nil
}

func (x *HomeSceneArrangementInfo) GetMLECLIOFALA() []*HomeFurnitureData {
	if x != nil {
		return x.MLECLIOFALA
	}
	return nil
}

func (x *HomeSceneArrangementInfo) GetBornRot() *Vector {
	if x != nil {
		return x.BornRot
	}
	return nil
}

func (x *HomeSceneArrangementInfo) GetSceneId() uint32 {
	if x != nil {
		return x.SceneId
	}
	return 0
}

func (x *HomeSceneArrangementInfo) GetBgmId() uint32 {
	if x != nil {
		return x.BgmId
	}
	return 0
}

func (x *HomeSceneArrangementInfo) GetComfortValue() uint32 {
	if x != nil {
		return x.ComfortValue
	}
	return 0
}

func (x *HomeSceneArrangementInfo) GetMainHouse() *HomeFurnitureData {
	if x != nil {
		return x.MainHouse
	}
	return nil
}

func (x *HomeSceneArrangementInfo) GetTmpVersion() uint32 {
	if x != nil {
		return x.TmpVersion
	}
	return 0
}

func (x *HomeSceneArrangementInfo) GetIsSetBornPos() bool {
	if x != nil {
		return x.IsSetBornPos
	}
	return false
}

func (x *HomeSceneArrangementInfo) GetDjinnPos() *Vector {
	if x != nil {
		return x.DjinnPos
	}
	return nil
}

func (x *HomeSceneArrangementInfo) GetBornPos() *Vector {
	if x != nil {
		return x.BornPos
	}
	return nil
}

var File_proto_HomeSceneArrangementInfo_proto protoreflect.FileDescriptor

var file_proto_HomeSceneArrangementInfo_proto_rawDesc = []byte{
	0x0a, 0x24, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x48, 0x6f, 0x6d, 0x65, 0x53, 0x63, 0x65, 0x6e,
	0x65, 0x41, 0x72, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x48, 0x6f,
	0x6d, 0x65, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x41, 0x72, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x48, 0x6f, 0x6d, 0x65, 0x46, 0x75, 0x72, 0x6e, 0x69, 0x74, 0x75, 0x72,
	0x65, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xa0, 0x04, 0x0a, 0x18, 0x48, 0x6f, 0x6d, 0x65, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x41, 0x72, 0x72,
	0x61, 0x6e, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x58, 0x0a, 0x1b,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x61, 0x72, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x48, 0x6f, 0x6d, 0x65, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x41, 0x72, 0x72,
	0x61, 0x6e, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x18, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x41, 0x72, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x6e,
	0x66, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x34, 0x0a, 0x0b, 0x48, 0x44, 0x4f, 0x46, 0x41, 0x48,
	0x47, 0x50, 0x4a, 0x47, 0x4b, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x48, 0x6f,
	0x6d, 0x65, 0x46, 0x75, 0x72, 0x6e, 0x69, 0x74, 0x75, 0x72, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52,
	0x0b, 0x48, 0x44, 0x4f, 0x46, 0x41, 0x48, 0x47, 0x50, 0x4a, 0x47, 0x4b, 0x12, 0x34, 0x0a, 0x0b,
	0x4d, 0x4c, 0x45, 0x43, 0x4c, 0x49, 0x4f, 0x46, 0x41, 0x4c, 0x41, 0x18, 0x0c, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x12, 0x2e, 0x48, 0x6f, 0x6d, 0x65, 0x46, 0x75, 0x72, 0x6e, 0x69, 0x74, 0x75, 0x72,
	0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x0b, 0x4d, 0x4c, 0x45, 0x43, 0x4c, 0x49, 0x4f, 0x46, 0x41,
	0x4c, 0x41, 0x12, 0x22, 0x0a, 0x08, 0x62, 0x6f, 0x72, 0x6e, 0x5f, 0x72, 0x6f, 0x74, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x07, 0x62,
	0x6f, 0x72, 0x6e, 0x52, 0x6f, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x63, 0x65, 0x6e, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x73, 0x63, 0x65, 0x6e, 0x65, 0x49,
	0x64, 0x12, 0x15, 0x0a, 0x06, 0x62, 0x67, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x05, 0x62, 0x67, 0x6d, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x66,
	0x6f, 0x72, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0c, 0x63, 0x6f, 0x6d, 0x66, 0x6f, 0x72, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x31, 0x0a,
	0x0a, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x12, 0x2e, 0x48, 0x6f, 0x6d, 0x65, 0x46, 0x75, 0x72, 0x6e, 0x69, 0x74, 0x75, 0x72,
	0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x09, 0x6d, 0x61, 0x69, 0x6e, 0x48, 0x6f, 0x75, 0x73, 0x65,
	0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6d, 0x70, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18,
	0x0d, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x74, 0x6d, 0x70, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x25, 0x0a, 0x0f, 0x69, 0x73, 0x5f, 0x73, 0x65, 0x74, 0x5f, 0x62, 0x6f, 0x72, 0x6e,
	0x5f, 0x70, 0x6f, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x69, 0x73, 0x53, 0x65,
	0x74, 0x42, 0x6f, 0x72, 0x6e, 0x50, 0x6f, 0x73, 0x12, 0x24, 0x0a, 0x09, 0x64, 0x6a, 0x69, 0x6e,
	0x6e, 0x5f, 0x70, 0x6f, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x56, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x52, 0x08, 0x64, 0x6a, 0x69, 0x6e, 0x6e, 0x50, 0x6f, 0x73, 0x12, 0x22,
	0x0a, 0x08, 0x62, 0x6f, 0x72, 0x6e, 0x5f, 0x70, 0x6f, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x07, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x07, 0x62, 0x6f, 0x72, 0x6e, 0x50,
	0x6f, 0x73, 0x42, 0x1d, 0x5a, 0x1b, 0x47, 0x6f, 0x2d, 0x47, 0x72, 0x61, 0x73, 0x73, 0x63, 0x75,
	0x74, 0x74, 0x65, 0x72, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_HomeSceneArrangementInfo_proto_rawDescOnce sync.Once
	file_proto_HomeSceneArrangementInfo_proto_rawDescData = file_proto_HomeSceneArrangementInfo_proto_rawDesc
)

func file_proto_HomeSceneArrangementInfo_proto_rawDescGZIP() []byte {
	file_proto_HomeSceneArrangementInfo_proto_rawDescOnce.Do(func() {
		file_proto_HomeSceneArrangementInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_HomeSceneArrangementInfo_proto_rawDescData)
	})
	return file_proto_HomeSceneArrangementInfo_proto_rawDescData
}

var file_proto_HomeSceneArrangementInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_HomeSceneArrangementInfo_proto_goTypes = []interface{}{
	(*HomeSceneArrangementInfo)(nil), // 0: HomeSceneArrangementInfo
	(*HomeBlockArrangementInfo)(nil), // 1: HomeBlockArrangementInfo
	(*HomeFurnitureData)(nil),        // 2: HomeFurnitureData
	(*Vector)(nil),                   // 3: Vector
}
var file_proto_HomeSceneArrangementInfo_proto_depIdxs = []int32{
	1, // 0: HomeSceneArrangementInfo.block_arrangement_info_list:type_name -> HomeBlockArrangementInfo
	2, // 1: HomeSceneArrangementInfo.HDOFAHGPJGK:type_name -> HomeFurnitureData
	2, // 2: HomeSceneArrangementInfo.MLECLIOFALA:type_name -> HomeFurnitureData
	3, // 3: HomeSceneArrangementInfo.born_rot:type_name -> Vector
	2, // 4: HomeSceneArrangementInfo.main_house:type_name -> HomeFurnitureData
	3, // 5: HomeSceneArrangementInfo.djinn_pos:type_name -> Vector
	3, // 6: HomeSceneArrangementInfo.born_pos:type_name -> Vector
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_proto_HomeSceneArrangementInfo_proto_init() }
func file_proto_HomeSceneArrangementInfo_proto_init() {
	if File_proto_HomeSceneArrangementInfo_proto != nil {
		return
	}
	file_proto_HomeBlockArrangementInfo_proto_init()
	file_proto_HomeFurnitureData_proto_init()
	file_proto_Vector_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_HomeSceneArrangementInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HomeSceneArrangementInfo); i {
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
			RawDescriptor: file_proto_HomeSceneArrangementInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_HomeSceneArrangementInfo_proto_goTypes,
		DependencyIndexes: file_proto_HomeSceneArrangementInfo_proto_depIdxs,
		MessageInfos:      file_proto_HomeSceneArrangementInfo_proto_msgTypes,
	}.Build()
	File_proto_HomeSceneArrangementInfo_proto = out.File
	file_proto_HomeSceneArrangementInfo_proto_rawDesc = nil
	file_proto_HomeSceneArrangementInfo_proto_goTypes = nil
	file_proto_HomeSceneArrangementInfo_proto_depIdxs = nil
}
