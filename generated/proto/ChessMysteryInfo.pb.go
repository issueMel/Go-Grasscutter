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
// source: proto/ChessMysteryInfo.proto

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

// Obf: ALPNHHGCIKK
type ChessMysteryInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EntrancePointMap map[uint32]uint32                   `protobuf:"bytes,8,rep,name=entrance_point_map,json=entrancePointMap,proto3" json:"entrance_point_map,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	DetailInfoMap    map[uint32]*ChessEntranceDetailInfo `protobuf:"bytes,14,rep,name=detail_info_map,json=detailInfoMap,proto3" json:"detail_info_map,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	ExitPointIdList  []uint32                            `protobuf:"varint,6,rep,packed,name=exit_point_id_list,json=exitPointIdList,proto3" json:"exit_point_id_list,omitempty"`
}

func (x *ChessMysteryInfo) Reset() {
	*x = ChessMysteryInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ChessMysteryInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChessMysteryInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChessMysteryInfo) ProtoMessage() {}

func (x *ChessMysteryInfo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ChessMysteryInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChessMysteryInfo.ProtoReflect.Descriptor instead.
func (*ChessMysteryInfo) Descriptor() ([]byte, []int) {
	return file_proto_ChessMysteryInfo_proto_rawDescGZIP(), []int{0}
}

func (x *ChessMysteryInfo) GetEntrancePointMap() map[uint32]uint32 {
	if x != nil {
		return x.EntrancePointMap
	}
	return nil
}

func (x *ChessMysteryInfo) GetDetailInfoMap() map[uint32]*ChessEntranceDetailInfo {
	if x != nil {
		return x.DetailInfoMap
	}
	return nil
}

func (x *ChessMysteryInfo) GetExitPointIdList() []uint32 {
	if x != nil {
		return x.ExitPointIdList
	}
	return nil
}

var File_proto_ChessMysteryInfo_proto protoreflect.FileDescriptor

var file_proto_ChessMysteryInfo_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x43, 0x68, 0x65, 0x73, 0x73, 0x4d, 0x79, 0x73,
	0x74, 0x65, 0x72, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x43, 0x68, 0x65, 0x73, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x61,
	0x6e, 0x63, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x85, 0x03, 0x0a, 0x10, 0x43, 0x68, 0x65, 0x73, 0x73, 0x4d, 0x79, 0x73,
	0x74, 0x65, 0x72, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x55, 0x0a, 0x12, 0x65, 0x6e, 0x74, 0x72,
	0x61, 0x6e, 0x63, 0x65, 0x5f, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x6d, 0x61, 0x70, 0x18, 0x08,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x43, 0x68, 0x65, 0x73, 0x73, 0x4d, 0x79, 0x73, 0x74,
	0x65, 0x72, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x45, 0x6e, 0x74, 0x72, 0x61, 0x6e, 0x63, 0x65,
	0x50, 0x6f, 0x69, 0x6e, 0x74, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x10, 0x65,
	0x6e, 0x74, 0x72, 0x61, 0x6e, 0x63, 0x65, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x4d, 0x61, 0x70, 0x12,
	0x4c, 0x0a, 0x0f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x5f, 0x6d,
	0x61, 0x70, 0x18, 0x0e, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x43, 0x68, 0x65, 0x73, 0x73,
	0x4d, 0x79, 0x73, 0x74, 0x65, 0x72, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0d,
	0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x4d, 0x61, 0x70, 0x12, 0x2b, 0x0a,
	0x12, 0x65, 0x78, 0x69, 0x74, 0x5f, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x5f, 0x6c,
	0x69, 0x73, 0x74, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0f, 0x65, 0x78, 0x69, 0x74, 0x50,
	0x6f, 0x69, 0x6e, 0x74, 0x49, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x1a, 0x43, 0x0a, 0x15, 0x45, 0x6e,
	0x74, 0x72, 0x61, 0x6e, 0x63, 0x65, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x4d, 0x61, 0x70, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a,
	0x5a, 0x0a, 0x12, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x4d, 0x61, 0x70,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2e, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x43, 0x68, 0x65, 0x73, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x61, 0x6e, 0x63, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x20, 0x5a, 0x1e, 0x47,
	0x6f, 0x2d, 0x47, 0x72, 0x61, 0x73, 0x73, 0x63, 0x75, 0x74, 0x74, 0x65, 0x72, 0x2f, 0x67, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_ChessMysteryInfo_proto_rawDescOnce sync.Once
	file_proto_ChessMysteryInfo_proto_rawDescData = file_proto_ChessMysteryInfo_proto_rawDesc
)

func file_proto_ChessMysteryInfo_proto_rawDescGZIP() []byte {
	file_proto_ChessMysteryInfo_proto_rawDescOnce.Do(func() {
		file_proto_ChessMysteryInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_ChessMysteryInfo_proto_rawDescData)
	})
	return file_proto_ChessMysteryInfo_proto_rawDescData
}

var file_proto_ChessMysteryInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_ChessMysteryInfo_proto_goTypes = []interface{}{
	(*ChessMysteryInfo)(nil),        // 0: ChessMysteryInfo
	nil,                             // 1: ChessMysteryInfo.EntrancePointMapEntry
	nil,                             // 2: ChessMysteryInfo.DetailInfoMapEntry
	(*ChessEntranceDetailInfo)(nil), // 3: ChessEntranceDetailInfo
}
var file_proto_ChessMysteryInfo_proto_depIdxs = []int32{
	1, // 0: ChessMysteryInfo.entrance_point_map:type_name -> ChessMysteryInfo.EntrancePointMapEntry
	2, // 1: ChessMysteryInfo.detail_info_map:type_name -> ChessMysteryInfo.DetailInfoMapEntry
	3, // 2: ChessMysteryInfo.DetailInfoMapEntry.value:type_name -> ChessEntranceDetailInfo
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_proto_ChessMysteryInfo_proto_init() }
func file_proto_ChessMysteryInfo_proto_init() {
	if File_proto_ChessMysteryInfo_proto != nil {
		return
	}
	file_proto_ChessEntranceDetailInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_ChessMysteryInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChessMysteryInfo); i {
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
			RawDescriptor: file_proto_ChessMysteryInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_ChessMysteryInfo_proto_goTypes,
		DependencyIndexes: file_proto_ChessMysteryInfo_proto_depIdxs,
		MessageInfos:      file_proto_ChessMysteryInfo_proto_msgTypes,
	}.Build()
	File_proto_ChessMysteryInfo_proto = out.File
	file_proto_ChessMysteryInfo_proto_rawDesc = nil
	file_proto_ChessMysteryInfo_proto_goTypes = nil
	file_proto_ChessMysteryInfo_proto_depIdxs = nil
}
