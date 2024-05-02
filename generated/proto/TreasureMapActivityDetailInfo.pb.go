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
// source: proto/TreasureMapActivityDetailInfo.proto

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

// Obf: EPFHAMJMACE
type TreasureMapActivityDetailInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TreasureCloseTime    uint32                           `protobuf:"varint,10,opt,name=treasure_close_time,json=treasureCloseTime,proto3" json:"treasure_close_time,omitempty"`
	IsMpChallengeTouched bool                             `protobuf:"varint,12,opt,name=is_mp_challenge_touched,json=isMpChallengeTouched,proto3" json:"is_mp_challenge_touched,omitempty"`
	CurrencyNum          uint32                           `protobuf:"varint,5,opt,name=currency_num,json=currencyNum,proto3" json:"currency_num,omitempty"`
	LNGOLLDBIEC          uint32                           `protobuf:"varint,7,opt,name=LNGOLLDBIEC,proto3" json:"LNGOLLDBIEC,omitempty"`
	RegionInfoList       []*TreasureMapRegionInfo         `protobuf:"bytes,6,rep,name=region_info_list,json=regionInfoList,proto3" json:"region_info_list,omitempty"`
	ActiveRegionIndex    uint32                           `protobuf:"varint,14,opt,name=active_region_index,json=activeRegionIndex,proto3" json:"active_region_index,omitempty"`
	BonusChallengeList   []*TreasureMapBonusChallengeInfo `protobuf:"bytes,11,rep,name=bonus_challenge_list,json=bonusChallengeList,proto3" json:"bonus_challenge_list,omitempty"`
	MPJCDHBBHMM          uint32                           `protobuf:"varint,9,opt,name=MPJCDHBBHMM,proto3" json:"MPJCDHBBHMM,omitempty"`
	OLCIPBELKIB          uint32                           `protobuf:"varint,8,opt,name=OLCIPBELKIB,proto3" json:"OLCIPBELKIB,omitempty"`
}

func (x *TreasureMapActivityDetailInfo) Reset() {
	*x = TreasureMapActivityDetailInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_TreasureMapActivityDetailInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TreasureMapActivityDetailInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TreasureMapActivityDetailInfo) ProtoMessage() {}

func (x *TreasureMapActivityDetailInfo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_TreasureMapActivityDetailInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TreasureMapActivityDetailInfo.ProtoReflect.Descriptor instead.
func (*TreasureMapActivityDetailInfo) Descriptor() ([]byte, []int) {
	return file_proto_TreasureMapActivityDetailInfo_proto_rawDescGZIP(), []int{0}
}

func (x *TreasureMapActivityDetailInfo) GetTreasureCloseTime() uint32 {
	if x != nil {
		return x.TreasureCloseTime
	}
	return 0
}

func (x *TreasureMapActivityDetailInfo) GetIsMpChallengeTouched() bool {
	if x != nil {
		return x.IsMpChallengeTouched
	}
	return false
}

func (x *TreasureMapActivityDetailInfo) GetCurrencyNum() uint32 {
	if x != nil {
		return x.CurrencyNum
	}
	return 0
}

func (x *TreasureMapActivityDetailInfo) GetLNGOLLDBIEC() uint32 {
	if x != nil {
		return x.LNGOLLDBIEC
	}
	return 0
}

func (x *TreasureMapActivityDetailInfo) GetRegionInfoList() []*TreasureMapRegionInfo {
	if x != nil {
		return x.RegionInfoList
	}
	return nil
}

func (x *TreasureMapActivityDetailInfo) GetActiveRegionIndex() uint32 {
	if x != nil {
		return x.ActiveRegionIndex
	}
	return 0
}

func (x *TreasureMapActivityDetailInfo) GetBonusChallengeList() []*TreasureMapBonusChallengeInfo {
	if x != nil {
		return x.BonusChallengeList
	}
	return nil
}

func (x *TreasureMapActivityDetailInfo) GetMPJCDHBBHMM() uint32 {
	if x != nil {
		return x.MPJCDHBBHMM
	}
	return 0
}

func (x *TreasureMapActivityDetailInfo) GetOLCIPBELKIB() uint32 {
	if x != nil {
		return x.OLCIPBELKIB
	}
	return 0
}

var File_proto_TreasureMapActivityDetailInfo_proto protoreflect.FileDescriptor

var file_proto_TreasureMapActivityDetailInfo_proto_rawDesc = []byte{
	0x0a, 0x29, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x54, 0x72, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65,
	0x4d, 0x61, 0x70, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x54, 0x72, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x4d, 0x61, 0x70, 0x52, 0x65,
	0x67, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x29,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x54, 0x72, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x4d, 0x61,
	0x70, 0x42, 0x6f, 0x6e, 0x75, 0x73, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x49,
	0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd3, 0x03, 0x0a, 0x1d, 0x54, 0x72,
	0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x4d, 0x61, 0x70, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74,
	0x79, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x2e, 0x0a, 0x13, 0x74,
	0x72, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x5f, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x11, 0x74, 0x72, 0x65, 0x61, 0x73, 0x75,
	0x72, 0x65, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x35, 0x0a, 0x17, 0x69,
	0x73, 0x5f, 0x6d, 0x70, 0x5f, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x5f, 0x74,
	0x6f, 0x75, 0x63, 0x68, 0x65, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x08, 0x52, 0x14, 0x69, 0x73,
	0x4d, 0x70, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x54, 0x6f, 0x75, 0x63, 0x68,
	0x65, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x6e,
	0x75, 0x6d, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x63, 0x79, 0x4e, 0x75, 0x6d, 0x12, 0x20, 0x0a, 0x0b, 0x4c, 0x4e, 0x47, 0x4f, 0x4c, 0x4c, 0x44,
	0x42, 0x49, 0x45, 0x43, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x4c, 0x4e, 0x47, 0x4f,
	0x4c, 0x4c, 0x44, 0x42, 0x49, 0x45, 0x43, 0x12, 0x40, 0x0a, 0x10, 0x72, 0x65, 0x67, 0x69, 0x6f,
	0x6e, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x06, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x16, 0x2e, 0x54, 0x72, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x4d, 0x61, 0x70, 0x52,
	0x65, 0x67, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0e, 0x72, 0x65, 0x67, 0x69, 0x6f,
	0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x13, 0x61, 0x63, 0x74,
	0x69, 0x76, 0x65, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78,
	0x18, 0x0e, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x11, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x52, 0x65,
	0x67, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x50, 0x0a, 0x14, 0x62, 0x6f, 0x6e,
	0x75, 0x73, 0x5f, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x5f, 0x6c, 0x69, 0x73,
	0x74, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x54, 0x72, 0x65, 0x61, 0x73, 0x75,
	0x72, 0x65, 0x4d, 0x61, 0x70, 0x42, 0x6f, 0x6e, 0x75, 0x73, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65,
	0x6e, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x12, 0x62, 0x6f, 0x6e, 0x75, 0x73, 0x43, 0x68,
	0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x4d,
	0x50, 0x4a, 0x43, 0x44, 0x48, 0x42, 0x42, 0x48, 0x4d, 0x4d, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0b, 0x4d, 0x50, 0x4a, 0x43, 0x44, 0x48, 0x42, 0x42, 0x48, 0x4d, 0x4d, 0x12, 0x20, 0x0a,
	0x0b, 0x4f, 0x4c, 0x43, 0x49, 0x50, 0x42, 0x45, 0x4c, 0x4b, 0x49, 0x42, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0b, 0x4f, 0x4c, 0x43, 0x49, 0x50, 0x42, 0x45, 0x4c, 0x4b, 0x49, 0x42, 0x42,
	0x20, 0x5a, 0x1e, 0x47, 0x6f, 0x2d, 0x47, 0x72, 0x61, 0x73, 0x73, 0x63, 0x75, 0x74, 0x74, 0x65,
	0x72, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_TreasureMapActivityDetailInfo_proto_rawDescOnce sync.Once
	file_proto_TreasureMapActivityDetailInfo_proto_rawDescData = file_proto_TreasureMapActivityDetailInfo_proto_rawDesc
)

func file_proto_TreasureMapActivityDetailInfo_proto_rawDescGZIP() []byte {
	file_proto_TreasureMapActivityDetailInfo_proto_rawDescOnce.Do(func() {
		file_proto_TreasureMapActivityDetailInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_TreasureMapActivityDetailInfo_proto_rawDescData)
	})
	return file_proto_TreasureMapActivityDetailInfo_proto_rawDescData
}

var file_proto_TreasureMapActivityDetailInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_TreasureMapActivityDetailInfo_proto_goTypes = []interface{}{
	(*TreasureMapActivityDetailInfo)(nil), // 0: TreasureMapActivityDetailInfo
	(*TreasureMapRegionInfo)(nil),         // 1: TreasureMapRegionInfo
	(*TreasureMapBonusChallengeInfo)(nil), // 2: TreasureMapBonusChallengeInfo
}
var file_proto_TreasureMapActivityDetailInfo_proto_depIdxs = []int32{
	1, // 0: TreasureMapActivityDetailInfo.region_info_list:type_name -> TreasureMapRegionInfo
	2, // 1: TreasureMapActivityDetailInfo.bonus_challenge_list:type_name -> TreasureMapBonusChallengeInfo
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_TreasureMapActivityDetailInfo_proto_init() }
func file_proto_TreasureMapActivityDetailInfo_proto_init() {
	if File_proto_TreasureMapActivityDetailInfo_proto != nil {
		return
	}
	file_proto_TreasureMapRegionInfo_proto_init()
	file_proto_TreasureMapBonusChallengeInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_TreasureMapActivityDetailInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TreasureMapActivityDetailInfo); i {
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
			RawDescriptor: file_proto_TreasureMapActivityDetailInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_TreasureMapActivityDetailInfo_proto_goTypes,
		DependencyIndexes: file_proto_TreasureMapActivityDetailInfo_proto_depIdxs,
		MessageInfos:      file_proto_TreasureMapActivityDetailInfo_proto_msgTypes,
	}.Build()
	File_proto_TreasureMapActivityDetailInfo_proto = out.File
	file_proto_TreasureMapActivityDetailInfo_proto_rawDesc = nil
	file_proto_TreasureMapActivityDetailInfo_proto_goTypes = nil
	file_proto_TreasureMapActivityDetailInfo_proto_depIdxs = nil
}
