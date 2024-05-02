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
// source: proto/FungusFighterDetailInfo.proto

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

// Obf: HIDHOOJCCJF
type FungusFighterDetailInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HOHBALCCDMB                       []uint32                        `protobuf:"varint,13,rep,packed,name=HOHBALCCDMB,proto3" json:"HOHBALCCDMB,omitempty"`
	FOGDNLKADKI                       []uint32                        `protobuf:"varint,1,rep,packed,name=FOGDNLKADKI,proto3" json:"FOGDNLKADKI,omitempty"`
	TrainingDungeonProgressDetailList []*FungusTrainingProgressDetail `protobuf:"bytes,8,rep,name=training_dungeon_progress_detail_list,json=trainingDungeonProgressDetailList,proto3" json:"training_dungeon_progress_detail_list,omitempty"`
	TrainingDungeonDetailList         []*FungusTrainingDungeonDetail  `protobuf:"bytes,9,rep,name=training_dungeon_detail_list,json=trainingDungeonDetailList,proto3" json:"training_dungeon_detail_list,omitempty"`
	FungusDetailList                  []*FungusDetail                 `protobuf:"bytes,5,rep,name=fungus_detail_list,json=fungusDetailList,proto3" json:"fungus_detail_list,omitempty"`
	IGAAEONOMGM                       []uint32                        `protobuf:"varint,7,rep,packed,name=IGAAEONOMGM,proto3" json:"IGAAEONOMGM,omitempty"`
	PlotStageDetailList               []*FungusPlotStageDetail        `protobuf:"bytes,15,rep,name=plot_stage_detail_list,json=plotStageDetailList,proto3" json:"plot_stage_detail_list,omitempty"`
}

func (x *FungusFighterDetailInfo) Reset() {
	*x = FungusFighterDetailInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_FungusFighterDetailInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FungusFighterDetailInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FungusFighterDetailInfo) ProtoMessage() {}

func (x *FungusFighterDetailInfo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_FungusFighterDetailInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FungusFighterDetailInfo.ProtoReflect.Descriptor instead.
func (*FungusFighterDetailInfo) Descriptor() ([]byte, []int) {
	return file_proto_FungusFighterDetailInfo_proto_rawDescGZIP(), []int{0}
}

func (x *FungusFighterDetailInfo) GetHOHBALCCDMB() []uint32 {
	if x != nil {
		return x.HOHBALCCDMB
	}
	return nil
}

func (x *FungusFighterDetailInfo) GetFOGDNLKADKI() []uint32 {
	if x != nil {
		return x.FOGDNLKADKI
	}
	return nil
}

func (x *FungusFighterDetailInfo) GetTrainingDungeonProgressDetailList() []*FungusTrainingProgressDetail {
	if x != nil {
		return x.TrainingDungeonProgressDetailList
	}
	return nil
}

func (x *FungusFighterDetailInfo) GetTrainingDungeonDetailList() []*FungusTrainingDungeonDetail {
	if x != nil {
		return x.TrainingDungeonDetailList
	}
	return nil
}

func (x *FungusFighterDetailInfo) GetFungusDetailList() []*FungusDetail {
	if x != nil {
		return x.FungusDetailList
	}
	return nil
}

func (x *FungusFighterDetailInfo) GetIGAAEONOMGM() []uint32 {
	if x != nil {
		return x.IGAAEONOMGM
	}
	return nil
}

func (x *FungusFighterDetailInfo) GetPlotStageDetailList() []*FungusPlotStageDetail {
	if x != nil {
		return x.PlotStageDetailList
	}
	return nil
}

var File_proto_FungusFighterDetailInfo_proto protoreflect.FileDescriptor

var file_proto_FungusFighterDetailInfo_proto_rawDesc = []byte{
	0x0a, 0x23, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x46, 0x75, 0x6e, 0x67, 0x75, 0x73, 0x46, 0x69,
	0x67, 0x68, 0x74, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x28, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x46, 0x75, 0x6e,
	0x67, 0x75, 0x73, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x6f, 0x67, 0x72,
	0x65, 0x73, 0x73, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x27, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x46, 0x75, 0x6e, 0x67, 0x75, 0x73, 0x54, 0x72, 0x61,
	0x69, 0x6e, 0x69, 0x6e, 0x67, 0x44, 0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x46, 0x75, 0x6e, 0x67, 0x75, 0x73, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x21, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x46, 0x75, 0x6e, 0x67, 0x75, 0x73,
	0x50, 0x6c, 0x6f, 0x74, 0x53, 0x74, 0x61, 0x67, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd9, 0x03, 0x0a, 0x17, 0x46, 0x75, 0x6e, 0x67, 0x75, 0x73,
	0x46, 0x69, 0x67, 0x68, 0x74, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x20, 0x0a, 0x0b, 0x48, 0x4f, 0x48, 0x42, 0x41, 0x4c, 0x43, 0x43, 0x44, 0x4d, 0x42,
	0x18, 0x0d, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0b, 0x48, 0x4f, 0x48, 0x42, 0x41, 0x4c, 0x43, 0x43,
	0x44, 0x4d, 0x42, 0x12, 0x20, 0x0a, 0x0b, 0x46, 0x4f, 0x47, 0x44, 0x4e, 0x4c, 0x4b, 0x41, 0x44,
	0x4b, 0x49, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0b, 0x46, 0x4f, 0x47, 0x44, 0x4e, 0x4c,
	0x4b, 0x41, 0x44, 0x4b, 0x49, 0x12, 0x6f, 0x0a, 0x25, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e,
	0x67, 0x5f, 0x64, 0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65,
	0x73, 0x73, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x08,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x46, 0x75, 0x6e, 0x67, 0x75, 0x73, 0x54, 0x72, 0x61,
	0x69, 0x6e, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x52, 0x21, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x44, 0x75, 0x6e,
	0x67, 0x65, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x5d, 0x0a, 0x1c, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x69,
	0x6e, 0x67, 0x5f, 0x64, 0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x46,
	0x75, 0x6e, 0x67, 0x75, 0x73, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x44, 0x75, 0x6e,
	0x67, 0x65, 0x6f, 0x6e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x19, 0x74, 0x72, 0x61, 0x69,
	0x6e, 0x69, 0x6e, 0x67, 0x44, 0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x3b, 0x0a, 0x12, 0x66, 0x75, 0x6e, 0x67, 0x75, 0x73, 0x5f,
	0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x05, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0d, 0x2e, 0x46, 0x75, 0x6e, 0x67, 0x75, 0x73, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x52, 0x10, 0x66, 0x75, 0x6e, 0x67, 0x75, 0x73, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x49, 0x47, 0x41, 0x41, 0x45, 0x4f, 0x4e, 0x4f, 0x4d, 0x47,
	0x4d, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0b, 0x49, 0x47, 0x41, 0x41, 0x45, 0x4f, 0x4e,
	0x4f, 0x4d, 0x47, 0x4d, 0x12, 0x4b, 0x0a, 0x16, 0x70, 0x6c, 0x6f, 0x74, 0x5f, 0x73, 0x74, 0x61,
	0x67, 0x65, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0f,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x46, 0x75, 0x6e, 0x67, 0x75, 0x73, 0x50, 0x6c, 0x6f,
	0x74, 0x53, 0x74, 0x61, 0x67, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x13, 0x70, 0x6c,
	0x6f, 0x74, 0x53, 0x74, 0x61, 0x67, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x4c, 0x69, 0x73,
	0x74, 0x42, 0x20, 0x5a, 0x1e, 0x47, 0x6f, 0x2d, 0x47, 0x72, 0x61, 0x73, 0x73, 0x63, 0x75, 0x74,
	0x74, 0x65, 0x72, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_FungusFighterDetailInfo_proto_rawDescOnce sync.Once
	file_proto_FungusFighterDetailInfo_proto_rawDescData = file_proto_FungusFighterDetailInfo_proto_rawDesc
)

func file_proto_FungusFighterDetailInfo_proto_rawDescGZIP() []byte {
	file_proto_FungusFighterDetailInfo_proto_rawDescOnce.Do(func() {
		file_proto_FungusFighterDetailInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_FungusFighterDetailInfo_proto_rawDescData)
	})
	return file_proto_FungusFighterDetailInfo_proto_rawDescData
}

var file_proto_FungusFighterDetailInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_FungusFighterDetailInfo_proto_goTypes = []interface{}{
	(*FungusFighterDetailInfo)(nil),      // 0: FungusFighterDetailInfo
	(*FungusTrainingProgressDetail)(nil), // 1: FungusTrainingProgressDetail
	(*FungusTrainingDungeonDetail)(nil),  // 2: FungusTrainingDungeonDetail
	(*FungusDetail)(nil),                 // 3: FungusDetail
	(*FungusPlotStageDetail)(nil),        // 4: FungusPlotStageDetail
}
var file_proto_FungusFighterDetailInfo_proto_depIdxs = []int32{
	1, // 0: FungusFighterDetailInfo.training_dungeon_progress_detail_list:type_name -> FungusTrainingProgressDetail
	2, // 1: FungusFighterDetailInfo.training_dungeon_detail_list:type_name -> FungusTrainingDungeonDetail
	3, // 2: FungusFighterDetailInfo.fungus_detail_list:type_name -> FungusDetail
	4, // 3: FungusFighterDetailInfo.plot_stage_detail_list:type_name -> FungusPlotStageDetail
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_proto_FungusFighterDetailInfo_proto_init() }
func file_proto_FungusFighterDetailInfo_proto_init() {
	if File_proto_FungusFighterDetailInfo_proto != nil {
		return
	}
	file_proto_FungusTrainingProgressDetail_proto_init()
	file_proto_FungusTrainingDungeonDetail_proto_init()
	file_proto_FungusDetail_proto_init()
	file_proto_FungusPlotStageDetail_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_FungusFighterDetailInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FungusFighterDetailInfo); i {
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
			RawDescriptor: file_proto_FungusFighterDetailInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_FungusFighterDetailInfo_proto_goTypes,
		DependencyIndexes: file_proto_FungusFighterDetailInfo_proto_depIdxs,
		MessageInfos:      file_proto_FungusFighterDetailInfo_proto_msgTypes,
	}.Build()
	File_proto_FungusFighterDetailInfo_proto = out.File
	file_proto_FungusFighterDetailInfo_proto_rawDesc = nil
	file_proto_FungusFighterDetailInfo_proto_goTypes = nil
	file_proto_FungusFighterDetailInfo_proto_depIdxs = nil
}
