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
// source: proto/DungeonEntryInfo.proto

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

// Obf: HFIHEBKHMPP
type DungeonEntryInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NextRefreshTime             uint32                       `protobuf:"varint,12,opt,name=next_refresh_time,json=nextRefreshTime,proto3" json:"next_refresh_time,omitempty"`
	CFAFALIPFEN                 uint32                       `protobuf:"varint,9,opt,name=CFAFALIPFEN,proto3" json:"CFAFALIPFEN,omitempty"`
	EOLGBLICEIA                 uint32                       `protobuf:"varint,14,opt,name=EOLGBLICEIA,proto3" json:"EOLGBLICEIA,omitempty"`
	MaxBossChestNum             uint32                       `protobuf:"varint,11,opt,name=max_boss_chest_num,json=maxBossChestNum,proto3" json:"max_boss_chest_num,omitempty"`
	DungeonId                   uint32                       `protobuf:"varint,13,opt,name=dungeon_id,json=dungeonId,proto3" json:"dungeon_id,omitempty"`
	EndTime                     uint32                       `protobuf:"varint,7,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	StartTime                   uint32                       `protobuf:"varint,2,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	IsPassed                    bool                         `protobuf:"varint,5,opt,name=is_passed,json=isPassed,proto3" json:"is_passed,omitempty"`
	WeeklyBossResinDiscountInfo *WeeklyBossResinDiscountInfo `protobuf:"bytes,6,opt,name=weekly_boss_resin_discount_info,json=weeklyBossResinDiscountInfo,proto3" json:"weekly_boss_resin_discount_info,omitempty"`
}

func (x *DungeonEntryInfo) Reset() {
	*x = DungeonEntryInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_DungeonEntryInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DungeonEntryInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DungeonEntryInfo) ProtoMessage() {}

func (x *DungeonEntryInfo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_DungeonEntryInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DungeonEntryInfo.ProtoReflect.Descriptor instead.
func (*DungeonEntryInfo) Descriptor() ([]byte, []int) {
	return file_proto_DungeonEntryInfo_proto_rawDescGZIP(), []int{0}
}

func (x *DungeonEntryInfo) GetNextRefreshTime() uint32 {
	if x != nil {
		return x.NextRefreshTime
	}
	return 0
}

func (x *DungeonEntryInfo) GetCFAFALIPFEN() uint32 {
	if x != nil {
		return x.CFAFALIPFEN
	}
	return 0
}

func (x *DungeonEntryInfo) GetEOLGBLICEIA() uint32 {
	if x != nil {
		return x.EOLGBLICEIA
	}
	return 0
}

func (x *DungeonEntryInfo) GetMaxBossChestNum() uint32 {
	if x != nil {
		return x.MaxBossChestNum
	}
	return 0
}

func (x *DungeonEntryInfo) GetDungeonId() uint32 {
	if x != nil {
		return x.DungeonId
	}
	return 0
}

func (x *DungeonEntryInfo) GetEndTime() uint32 {
	if x != nil {
		return x.EndTime
	}
	return 0
}

func (x *DungeonEntryInfo) GetStartTime() uint32 {
	if x != nil {
		return x.StartTime
	}
	return 0
}

func (x *DungeonEntryInfo) GetIsPassed() bool {
	if x != nil {
		return x.IsPassed
	}
	return false
}

func (x *DungeonEntryInfo) GetWeeklyBossResinDiscountInfo() *WeeklyBossResinDiscountInfo {
	if x != nil {
		return x.WeeklyBossResinDiscountInfo
	}
	return nil
}

var File_proto_DungeonEntryInfo_proto protoreflect.FileDescriptor

var file_proto_DungeonEntryInfo_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x44, 0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x27,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x57, 0x65, 0x65, 0x6b, 0x6c, 0x79, 0x42, 0x6f, 0x73, 0x73,
	0x52, 0x65, 0x73, 0x69, 0x6e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x6e, 0x66,
	0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x89, 0x03, 0x0a, 0x10, 0x44, 0x75, 0x6e, 0x67,
	0x65, 0x6f, 0x6e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x2a, 0x0a, 0x11,
	0x6e, 0x65, 0x78, 0x74, 0x5f, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x52, 0x65, 0x66,
	0x72, 0x65, 0x73, 0x68, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x43, 0x46, 0x41, 0x46,
	0x41, 0x4c, 0x49, 0x50, 0x46, 0x45, 0x4e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x43,
	0x46, 0x41, 0x46, 0x41, 0x4c, 0x49, 0x50, 0x46, 0x45, 0x4e, 0x12, 0x20, 0x0a, 0x0b, 0x45, 0x4f,
	0x4c, 0x47, 0x42, 0x4c, 0x49, 0x43, 0x45, 0x49, 0x41, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0b, 0x45, 0x4f, 0x4c, 0x47, 0x42, 0x4c, 0x49, 0x43, 0x45, 0x49, 0x41, 0x12, 0x2b, 0x0a, 0x12,
	0x6d, 0x61, 0x78, 0x5f, 0x62, 0x6f, 0x73, 0x73, 0x5f, 0x63, 0x68, 0x65, 0x73, 0x74, 0x5f, 0x6e,
	0x75, 0x6d, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x6d, 0x61, 0x78, 0x42, 0x6f, 0x73,
	0x73, 0x43, 0x68, 0x65, 0x73, 0x74, 0x4e, 0x75, 0x6d, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x75, 0x6e,
	0x67, 0x65, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x64,
	0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x65, 0x64, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x50, 0x61, 0x73, 0x73, 0x65, 0x64, 0x12,
	0x62, 0x0a, 0x1f, 0x77, 0x65, 0x65, 0x6b, 0x6c, 0x79, 0x5f, 0x62, 0x6f, 0x73, 0x73, 0x5f, 0x72,
	0x65, 0x73, 0x69, 0x6e, 0x5f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x6e,
	0x66, 0x6f, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x57, 0x65, 0x65, 0x6b, 0x6c,
	0x79, 0x42, 0x6f, 0x73, 0x73, 0x52, 0x65, 0x73, 0x69, 0x6e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x1b, 0x77, 0x65, 0x65, 0x6b, 0x6c, 0x79, 0x42, 0x6f,
	0x73, 0x73, 0x52, 0x65, 0x73, 0x69, 0x6e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49,
	0x6e, 0x66, 0x6f, 0x42, 0x1d, 0x5a, 0x1b, 0x47, 0x6f, 0x2d, 0x47, 0x72, 0x61, 0x73, 0x73, 0x63,
	0x75, 0x74, 0x74, 0x65, 0x72, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_DungeonEntryInfo_proto_rawDescOnce sync.Once
	file_proto_DungeonEntryInfo_proto_rawDescData = file_proto_DungeonEntryInfo_proto_rawDesc
)

func file_proto_DungeonEntryInfo_proto_rawDescGZIP() []byte {
	file_proto_DungeonEntryInfo_proto_rawDescOnce.Do(func() {
		file_proto_DungeonEntryInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_DungeonEntryInfo_proto_rawDescData)
	})
	return file_proto_DungeonEntryInfo_proto_rawDescData
}

var file_proto_DungeonEntryInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_DungeonEntryInfo_proto_goTypes = []interface{}{
	(*DungeonEntryInfo)(nil),            // 0: DungeonEntryInfo
	(*WeeklyBossResinDiscountInfo)(nil), // 1: WeeklyBossResinDiscountInfo
}
var file_proto_DungeonEntryInfo_proto_depIdxs = []int32{
	1, // 0: DungeonEntryInfo.weekly_boss_resin_discount_info:type_name -> WeeklyBossResinDiscountInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_DungeonEntryInfo_proto_init() }
func file_proto_DungeonEntryInfo_proto_init() {
	if File_proto_DungeonEntryInfo_proto != nil {
		return
	}
	file_proto_WeeklyBossResinDiscountInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_DungeonEntryInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DungeonEntryInfo); i {
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
			RawDescriptor: file_proto_DungeonEntryInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_DungeonEntryInfo_proto_goTypes,
		DependencyIndexes: file_proto_DungeonEntryInfo_proto_depIdxs,
		MessageInfos:      file_proto_DungeonEntryInfo_proto_msgTypes,
	}.Build()
	File_proto_DungeonEntryInfo_proto = out.File
	file_proto_DungeonEntryInfo_proto_rawDesc = nil
	file_proto_DungeonEntryInfo_proto_goTypes = nil
	file_proto_DungeonEntryInfo_proto_depIdxs = nil
}
