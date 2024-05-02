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
// source: proto/FishBattleEndRsp.proto

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

// Obf: KNCHAIIMIDM
type FishBattleEndRsp_FishNoRewardReason int32

const (
	FishBattleEndRsp_FISH_NO_REWARD_NONE           FishBattleEndRsp_FishNoRewardReason = 0
	FishBattleEndRsp_FISH_NO_REWARD_ACTIVITY_LIMIT FishBattleEndRsp_FishNoRewardReason = 1
	FishBattleEndRsp_FISH_NO_REWARD_BAG_LIMIT      FishBattleEndRsp_FishNoRewardReason = 2
	FishBattleEndRsp_FISH_NO_REWARD_POOL_LIMIT     FishBattleEndRsp_FishNoRewardReason = 3
)

// Enum value maps for FishBattleEndRsp_FishNoRewardReason.
var (
	FishBattleEndRsp_FishNoRewardReason_name = map[int32]string{
		0: "FISH_NO_REWARD_NONE",
		1: "FISH_NO_REWARD_ACTIVITY_LIMIT",
		2: "FISH_NO_REWARD_BAG_LIMIT",
		3: "FISH_NO_REWARD_POOL_LIMIT",
	}
	FishBattleEndRsp_FishNoRewardReason_value = map[string]int32{
		"FISH_NO_REWARD_NONE":           0,
		"FISH_NO_REWARD_ACTIVITY_LIMIT": 1,
		"FISH_NO_REWARD_BAG_LIMIT":      2,
		"FISH_NO_REWARD_POOL_LIMIT":     3,
	}
)

func (x FishBattleEndRsp_FishNoRewardReason) Enum() *FishBattleEndRsp_FishNoRewardReason {
	p := new(FishBattleEndRsp_FishNoRewardReason)
	*p = x
	return p
}

func (x FishBattleEndRsp_FishNoRewardReason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FishBattleEndRsp_FishNoRewardReason) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_FishBattleEndRsp_proto_enumTypes[0].Descriptor()
}

func (FishBattleEndRsp_FishNoRewardReason) Type() protoreflect.EnumType {
	return &file_proto_FishBattleEndRsp_proto_enumTypes[0]
}

func (x FishBattleEndRsp_FishNoRewardReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FishBattleEndRsp_FishNoRewardReason.Descriptor instead.
func (FishBattleEndRsp_FishNoRewardReason) EnumDescriptor() ([]byte, []int) {
	return file_proto_FishBattleEndRsp_proto_rawDescGZIP(), []int{0, 0}
}

// CmdId: 2479
// Obf: IDIIDOMLDIA
type FishBattleEndRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Retcode        int32                               `protobuf:"varint,9,opt,name=retcode,proto3" json:"retcode,omitempty"`
	BattleResult   FishBattleResult                    `protobuf:"varint,7,opt,name=battle_result,json=battleResult,proto3,enum=FishBattleResult" json:"battle_result,omitempty"`
	RewardItemList []*ItemParam                        `protobuf:"bytes,11,rep,name=reward_item_list,json=rewardItemList,proto3" json:"reward_item_list,omitempty"`
	NoRewardReason FishBattleEndRsp_FishNoRewardReason `protobuf:"varint,12,opt,name=no_reward_reason,json=noRewardReason,proto3,enum=FishBattleEndRsp_FishNoRewardReason" json:"no_reward_reason,omitempty"`
	FILLNPDLDGI    []*ItemParam                        `protobuf:"bytes,8,rep,name=FILLNPDLDGI,proto3" json:"FILLNPDLDGI,omitempty"`
	IsGotReward    bool                                `protobuf:"varint,13,opt,name=is_got_reward,json=isGotReward,proto3" json:"is_got_reward,omitempty"`
	AGADCCFPDEI    []*ItemParam                        `protobuf:"bytes,6,rep,name=AGADCCFPDEI,proto3" json:"AGADCCFPDEI,omitempty"`
}

func (x *FishBattleEndRsp) Reset() {
	*x = FishBattleEndRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_FishBattleEndRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FishBattleEndRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FishBattleEndRsp) ProtoMessage() {}

func (x *FishBattleEndRsp) ProtoReflect() protoreflect.Message {
	mi := &file_proto_FishBattleEndRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FishBattleEndRsp.ProtoReflect.Descriptor instead.
func (*FishBattleEndRsp) Descriptor() ([]byte, []int) {
	return file_proto_FishBattleEndRsp_proto_rawDescGZIP(), []int{0}
}

func (x *FishBattleEndRsp) GetRetcode() int32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *FishBattleEndRsp) GetBattleResult() FishBattleResult {
	if x != nil {
		return x.BattleResult
	}
	return FishBattleResult_FISH_BATTLE_RESULT_NONE
}

func (x *FishBattleEndRsp) GetRewardItemList() []*ItemParam {
	if x != nil {
		return x.RewardItemList
	}
	return nil
}

func (x *FishBattleEndRsp) GetNoRewardReason() FishBattleEndRsp_FishNoRewardReason {
	if x != nil {
		return x.NoRewardReason
	}
	return FishBattleEndRsp_FISH_NO_REWARD_NONE
}

func (x *FishBattleEndRsp) GetFILLNPDLDGI() []*ItemParam {
	if x != nil {
		return x.FILLNPDLDGI
	}
	return nil
}

func (x *FishBattleEndRsp) GetIsGotReward() bool {
	if x != nil {
		return x.IsGotReward
	}
	return false
}

func (x *FishBattleEndRsp) GetAGADCCFPDEI() []*ItemParam {
	if x != nil {
		return x.AGADCCFPDEI
	}
	return nil
}

var File_proto_FishBattleEndRsp_proto protoreflect.FileDescriptor

var file_proto_FishBattleEndRsp_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x46, 0x69, 0x73, 0x68, 0x42, 0x61, 0x74, 0x74,
	0x6c, 0x65, 0x45, 0x6e, 0x64, 0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x46, 0x69, 0x73, 0x68, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x49, 0x74, 0x65, 0x6d, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xfa, 0x03, 0x0a, 0x10, 0x46, 0x69, 0x73, 0x68, 0x42, 0x61, 0x74, 0x74,
	0x6c, 0x65, 0x45, 0x6e, 0x64, 0x52, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f,
	0x64, 0x65, 0x12, 0x36, 0x0a, 0x0d, 0x62, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x5f, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x46, 0x69, 0x73, 0x68,
	0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x0c, 0x62, 0x61,
	0x74, 0x74, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x34, 0x0a, 0x10, 0x72, 0x65,
	0x77, 0x61, 0x72, 0x64, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0b,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x52, 0x0e, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x4e, 0x0a, 0x10, 0x6e, 0x6f, 0x5f, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x72, 0x65,
	0x61, 0x73, 0x6f, 0x6e, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x24, 0x2e, 0x46, 0x69, 0x73,
	0x68, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x45, 0x6e, 0x64, 0x52, 0x73, 0x70, 0x2e, 0x46, 0x69,
	0x73, 0x68, 0x4e, 0x6f, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e,
	0x52, 0x0e, 0x6e, 0x6f, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e,
	0x12, 0x2c, 0x0a, 0x0b, 0x46, 0x49, 0x4c, 0x4c, 0x4e, 0x50, 0x44, 0x4c, 0x44, 0x47, 0x49, 0x18,
	0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x52, 0x0b, 0x46, 0x49, 0x4c, 0x4c, 0x4e, 0x50, 0x44, 0x4c, 0x44, 0x47, 0x49, 0x12, 0x22,
	0x0a, 0x0d, 0x69, 0x73, 0x5f, 0x67, 0x6f, 0x74, 0x5f, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x18,
	0x0d, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x69, 0x73, 0x47, 0x6f, 0x74, 0x52, 0x65, 0x77, 0x61,
	0x72, 0x64, 0x12, 0x2c, 0x0a, 0x0b, 0x41, 0x47, 0x41, 0x44, 0x43, 0x43, 0x46, 0x50, 0x44, 0x45,
	0x49, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x50, 0x61,
	0x72, 0x61, 0x6d, 0x52, 0x0b, 0x41, 0x47, 0x41, 0x44, 0x43, 0x43, 0x46, 0x50, 0x44, 0x45, 0x49,
	0x22, 0x8d, 0x01, 0x0a, 0x12, 0x46, 0x69, 0x73, 0x68, 0x4e, 0x6f, 0x52, 0x65, 0x77, 0x61, 0x72,
	0x64, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x17, 0x0a, 0x13, 0x46, 0x49, 0x53, 0x48, 0x5f,
	0x4e, 0x4f, 0x5f, 0x52, 0x45, 0x57, 0x41, 0x52, 0x44, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00,
	0x12, 0x21, 0x0a, 0x1d, 0x46, 0x49, 0x53, 0x48, 0x5f, 0x4e, 0x4f, 0x5f, 0x52, 0x45, 0x57, 0x41,
	0x52, 0x44, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x56, 0x49, 0x54, 0x59, 0x5f, 0x4c, 0x49, 0x4d, 0x49,
	0x54, 0x10, 0x01, 0x12, 0x1c, 0x0a, 0x18, 0x46, 0x49, 0x53, 0x48, 0x5f, 0x4e, 0x4f, 0x5f, 0x52,
	0x45, 0x57, 0x41, 0x52, 0x44, 0x5f, 0x42, 0x41, 0x47, 0x5f, 0x4c, 0x49, 0x4d, 0x49, 0x54, 0x10,
	0x02, 0x12, 0x1d, 0x0a, 0x19, 0x46, 0x49, 0x53, 0x48, 0x5f, 0x4e, 0x4f, 0x5f, 0x52, 0x45, 0x57,
	0x41, 0x52, 0x44, 0x5f, 0x50, 0x4f, 0x4f, 0x4c, 0x5f, 0x4c, 0x49, 0x4d, 0x49, 0x54, 0x10, 0x03,
	0x42, 0x20, 0x5a, 0x1e, 0x47, 0x6f, 0x2d, 0x47, 0x72, 0x61, 0x73, 0x73, 0x63, 0x75, 0x74, 0x74,
	0x65, 0x72, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_FishBattleEndRsp_proto_rawDescOnce sync.Once
	file_proto_FishBattleEndRsp_proto_rawDescData = file_proto_FishBattleEndRsp_proto_rawDesc
)

func file_proto_FishBattleEndRsp_proto_rawDescGZIP() []byte {
	file_proto_FishBattleEndRsp_proto_rawDescOnce.Do(func() {
		file_proto_FishBattleEndRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_FishBattleEndRsp_proto_rawDescData)
	})
	return file_proto_FishBattleEndRsp_proto_rawDescData
}

var file_proto_FishBattleEndRsp_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_FishBattleEndRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_FishBattleEndRsp_proto_goTypes = []interface{}{
	(FishBattleEndRsp_FishNoRewardReason)(0), // 0: FishBattleEndRsp.FishNoRewardReason
	(*FishBattleEndRsp)(nil),                 // 1: FishBattleEndRsp
	(FishBattleResult)(0),                    // 2: FishBattleResult
	(*ItemParam)(nil),                        // 3: ItemParam
}
var file_proto_FishBattleEndRsp_proto_depIdxs = []int32{
	2, // 0: FishBattleEndRsp.battle_result:type_name -> FishBattleResult
	3, // 1: FishBattleEndRsp.reward_item_list:type_name -> ItemParam
	0, // 2: FishBattleEndRsp.no_reward_reason:type_name -> FishBattleEndRsp.FishNoRewardReason
	3, // 3: FishBattleEndRsp.FILLNPDLDGI:type_name -> ItemParam
	3, // 4: FishBattleEndRsp.AGADCCFPDEI:type_name -> ItemParam
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_proto_FishBattleEndRsp_proto_init() }
func file_proto_FishBattleEndRsp_proto_init() {
	if File_proto_FishBattleEndRsp_proto != nil {
		return
	}
	file_proto_FishBattleResult_proto_init()
	file_proto_ItemParam_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_FishBattleEndRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FishBattleEndRsp); i {
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
			RawDescriptor: file_proto_FishBattleEndRsp_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_FishBattleEndRsp_proto_goTypes,
		DependencyIndexes: file_proto_FishBattleEndRsp_proto_depIdxs,
		EnumInfos:         file_proto_FishBattleEndRsp_proto_enumTypes,
		MessageInfos:      file_proto_FishBattleEndRsp_proto_msgTypes,
	}.Build()
	File_proto_FishBattleEndRsp_proto = out.File
	file_proto_FishBattleEndRsp_proto_rawDesc = nil
	file_proto_FishBattleEndRsp_proto_goTypes = nil
	file_proto_FishBattleEndRsp_proto_depIdxs = nil
}
