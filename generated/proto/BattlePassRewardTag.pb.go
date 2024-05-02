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
// source: proto/BattlePassRewardTag.proto

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

// Obf: MMDANGGKGCL
type BattlePassRewardTag struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Level        uint32                 `protobuf:"varint,12,opt,name=level,proto3" json:"level,omitempty"`
	RewardId     uint32                 `protobuf:"varint,8,opt,name=reward_id,json=rewardId,proto3" json:"reward_id,omitempty"`
	UnlockStatus BattlePassUnlockStatus `protobuf:"varint,1,opt,name=unlock_status,json=unlockStatus,proto3,enum=BattlePassUnlockStatus" json:"unlock_status,omitempty"`
}

func (x *BattlePassRewardTag) Reset() {
	*x = BattlePassRewardTag{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_BattlePassRewardTag_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BattlePassRewardTag) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BattlePassRewardTag) ProtoMessage() {}

func (x *BattlePassRewardTag) ProtoReflect() protoreflect.Message {
	mi := &file_proto_BattlePassRewardTag_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BattlePassRewardTag.ProtoReflect.Descriptor instead.
func (*BattlePassRewardTag) Descriptor() ([]byte, []int) {
	return file_proto_BattlePassRewardTag_proto_rawDescGZIP(), []int{0}
}

func (x *BattlePassRewardTag) GetLevel() uint32 {
	if x != nil {
		return x.Level
	}
	return 0
}

func (x *BattlePassRewardTag) GetRewardId() uint32 {
	if x != nil {
		return x.RewardId
	}
	return 0
}

func (x *BattlePassRewardTag) GetUnlockStatus() BattlePassUnlockStatus {
	if x != nil {
		return x.UnlockStatus
	}
	return BattlePassUnlockStatus_BATTLE_PASS_UNLOCK_STATUS_INVALID
}

var File_proto_BattlePassRewardTag_proto protoreflect.FileDescriptor

var file_proto_BattlePassRewardTag_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x50, 0x61,
	0x73, 0x73, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x54, 0x61, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x22, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x50,
	0x61, 0x73, 0x73, 0x55, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x86, 0x01, 0x0a, 0x13, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65,
	0x50, 0x61, 0x73, 0x73, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x54, 0x61, 0x67, 0x12, 0x14, 0x0a,
	0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6c, 0x65,
	0x76, 0x65, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x69, 0x64,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x49, 0x64,
	0x12, 0x3c, 0x0a, 0x0d, 0x75, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65,
	0x50, 0x61, 0x73, 0x73, 0x55, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x0c, 0x75, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x20,
	0x5a, 0x1e, 0x47, 0x6f, 0x2d, 0x47, 0x72, 0x61, 0x73, 0x73, 0x63, 0x75, 0x74, 0x74, 0x65, 0x72,
	0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_BattlePassRewardTag_proto_rawDescOnce sync.Once
	file_proto_BattlePassRewardTag_proto_rawDescData = file_proto_BattlePassRewardTag_proto_rawDesc
)

func file_proto_BattlePassRewardTag_proto_rawDescGZIP() []byte {
	file_proto_BattlePassRewardTag_proto_rawDescOnce.Do(func() {
		file_proto_BattlePassRewardTag_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_BattlePassRewardTag_proto_rawDescData)
	})
	return file_proto_BattlePassRewardTag_proto_rawDescData
}

var file_proto_BattlePassRewardTag_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_BattlePassRewardTag_proto_goTypes = []interface{}{
	(*BattlePassRewardTag)(nil), // 0: BattlePassRewardTag
	(BattlePassUnlockStatus)(0), // 1: BattlePassUnlockStatus
}
var file_proto_BattlePassRewardTag_proto_depIdxs = []int32{
	1, // 0: BattlePassRewardTag.unlock_status:type_name -> BattlePassUnlockStatus
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_BattlePassRewardTag_proto_init() }
func file_proto_BattlePassRewardTag_proto_init() {
	if File_proto_BattlePassRewardTag_proto != nil {
		return
	}
	file_proto_BattlePassUnlockStatus_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_BattlePassRewardTag_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BattlePassRewardTag); i {
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
			RawDescriptor: file_proto_BattlePassRewardTag_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_BattlePassRewardTag_proto_goTypes,
		DependencyIndexes: file_proto_BattlePassRewardTag_proto_depIdxs,
		MessageInfos:      file_proto_BattlePassRewardTag_proto_msgTypes,
	}.Build()
	File_proto_BattlePassRewardTag_proto = out.File
	file_proto_BattlePassRewardTag_proto_rawDesc = nil
	file_proto_BattlePassRewardTag_proto_goTypes = nil
	file_proto_BattlePassRewardTag_proto_depIdxs = nil
}
