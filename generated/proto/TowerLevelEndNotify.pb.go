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
// source: proto/TowerLevelEndNotify.proto

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

// CmdId: 20886
// Obf: BHDDDPNGPNF
type TowerLevelEndNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FinishedStarCondList []uint32     `protobuf:"varint,14,rep,packed,name=finished_star_cond_list,json=finishedStarCondList,proto3" json:"finished_star_cond_list,omitempty"`
	IsSuccess            bool         `protobuf:"varint,7,opt,name=is_success,json=isSuccess,proto3" json:"is_success,omitempty"`
	ContinueState        uint32       `protobuf:"varint,4,opt,name=continue_state,json=continueState,proto3" json:"continue_state,omitempty"`
	RewardItemList       []*ItemParam `protobuf:"bytes,13,rep,name=reward_item_list,json=rewardItemList,proto3" json:"reward_item_list,omitempty"`
	NextFloorId          uint32       `protobuf:"varint,5,opt,name=next_floor_id,json=nextFloorId,proto3" json:"next_floor_id,omitempty"`
}

func (x *TowerLevelEndNotify) Reset() {
	*x = TowerLevelEndNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_TowerLevelEndNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TowerLevelEndNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TowerLevelEndNotify) ProtoMessage() {}

func (x *TowerLevelEndNotify) ProtoReflect() protoreflect.Message {
	mi := &file_proto_TowerLevelEndNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TowerLevelEndNotify.ProtoReflect.Descriptor instead.
func (*TowerLevelEndNotify) Descriptor() ([]byte, []int) {
	return file_proto_TowerLevelEndNotify_proto_rawDescGZIP(), []int{0}
}

func (x *TowerLevelEndNotify) GetFinishedStarCondList() []uint32 {
	if x != nil {
		return x.FinishedStarCondList
	}
	return nil
}

func (x *TowerLevelEndNotify) GetIsSuccess() bool {
	if x != nil {
		return x.IsSuccess
	}
	return false
}

func (x *TowerLevelEndNotify) GetContinueState() uint32 {
	if x != nil {
		return x.ContinueState
	}
	return 0
}

func (x *TowerLevelEndNotify) GetRewardItemList() []*ItemParam {
	if x != nil {
		return x.RewardItemList
	}
	return nil
}

func (x *TowerLevelEndNotify) GetNextFloorId() uint32 {
	if x != nil {
		return x.NextFloorId
	}
	return 0
}

var File_proto_TowerLevelEndNotify_proto protoreflect.FileDescriptor

var file_proto_TowerLevelEndNotify_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x54, 0x6f, 0x77, 0x65, 0x72, 0x4c, 0x65, 0x76,
	0x65, 0x6c, 0x45, 0x6e, 0x64, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x49, 0x74, 0x65, 0x6d, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xec, 0x01, 0x0a, 0x13, 0x54, 0x6f, 0x77,
	0x65, 0x72, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x45, 0x6e, 0x64, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79,
	0x12, 0x35, 0x0a, 0x17, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x5f, 0x73, 0x74, 0x61,
	0x72, 0x5f, 0x63, 0x6f, 0x6e, 0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0e, 0x20, 0x03, 0x28,
	0x0d, 0x52, 0x14, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x53, 0x74, 0x61, 0x72, 0x43,
	0x6f, 0x6e, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x73, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x53,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x74, 0x69, 0x6e,
	0x75, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d,
	0x63, 0x6f, 0x6e, 0x74, 0x69, 0x6e, 0x75, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x34, 0x0a,
	0x10, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x6c, 0x69, 0x73,
	0x74, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x50, 0x61,
	0x72, 0x61, 0x6d, 0x52, 0x0e, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x66, 0x6c, 0x6f, 0x6f,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x6e, 0x65, 0x78, 0x74,
	0x46, 0x6c, 0x6f, 0x6f, 0x72, 0x49, 0x64, 0x42, 0x20, 0x5a, 0x1e, 0x47, 0x6f, 0x2d, 0x47, 0x72,
	0x61, 0x73, 0x73, 0x63, 0x75, 0x74, 0x74, 0x65, 0x72, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x74, 0x65, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_proto_TowerLevelEndNotify_proto_rawDescOnce sync.Once
	file_proto_TowerLevelEndNotify_proto_rawDescData = file_proto_TowerLevelEndNotify_proto_rawDesc
)

func file_proto_TowerLevelEndNotify_proto_rawDescGZIP() []byte {
	file_proto_TowerLevelEndNotify_proto_rawDescOnce.Do(func() {
		file_proto_TowerLevelEndNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_TowerLevelEndNotify_proto_rawDescData)
	})
	return file_proto_TowerLevelEndNotify_proto_rawDescData
}

var file_proto_TowerLevelEndNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_TowerLevelEndNotify_proto_goTypes = []interface{}{
	(*TowerLevelEndNotify)(nil), // 0: TowerLevelEndNotify
	(*ItemParam)(nil),           // 1: ItemParam
}
var file_proto_TowerLevelEndNotify_proto_depIdxs = []int32{
	1, // 0: TowerLevelEndNotify.reward_item_list:type_name -> ItemParam
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_TowerLevelEndNotify_proto_init() }
func file_proto_TowerLevelEndNotify_proto_init() {
	if File_proto_TowerLevelEndNotify_proto != nil {
		return
	}
	file_proto_ItemParam_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_TowerLevelEndNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TowerLevelEndNotify); i {
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
			RawDescriptor: file_proto_TowerLevelEndNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_TowerLevelEndNotify_proto_goTypes,
		DependencyIndexes: file_proto_TowerLevelEndNotify_proto_depIdxs,
		MessageInfos:      file_proto_TowerLevelEndNotify_proto_msgTypes,
	}.Build()
	File_proto_TowerLevelEndNotify_proto = out.File
	file_proto_TowerLevelEndNotify_proto_rawDesc = nil
	file_proto_TowerLevelEndNotify_proto_goTypes = nil
	file_proto_TowerLevelEndNotify_proto_depIdxs = nil
}
