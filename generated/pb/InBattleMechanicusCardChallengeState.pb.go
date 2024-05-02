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
// source: proto/InBattleMechanicusCardChallengeState.proto

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

// Obf: NCNMHENBLME
type InBattleMechanicusCardChallengeState int32

const (
	InBattleMechanicusCardChallengeState_IN_BATTLE_MECHANICUS_CARD_CHALLENGE_NONE     InBattleMechanicusCardChallengeState = 0
	InBattleMechanicusCardChallengeState_IN_BATTLE_MECHANICUS_CARD_CHALLENGE_ON_GOING InBattleMechanicusCardChallengeState = 1
	InBattleMechanicusCardChallengeState_IN_BATTLE_MECHANICUS_CARD_CHALLENGE_FAIL     InBattleMechanicusCardChallengeState = 2
	InBattleMechanicusCardChallengeState_IN_BATTLE_MECHANICUS_CARD_CHALLENGE_SUCCESS  InBattleMechanicusCardChallengeState = 3
)

// Enum value maps for InBattleMechanicusCardChallengeState.
var (
	InBattleMechanicusCardChallengeState_name = map[int32]string{
		0: "IN_BATTLE_MECHANICUS_CARD_CHALLENGE_NONE",
		1: "IN_BATTLE_MECHANICUS_CARD_CHALLENGE_ON_GOING",
		2: "IN_BATTLE_MECHANICUS_CARD_CHALLENGE_FAIL",
		3: "IN_BATTLE_MECHANICUS_CARD_CHALLENGE_SUCCESS",
	}
	InBattleMechanicusCardChallengeState_value = map[string]int32{
		"IN_BATTLE_MECHANICUS_CARD_CHALLENGE_NONE":     0,
		"IN_BATTLE_MECHANICUS_CARD_CHALLENGE_ON_GOING": 1,
		"IN_BATTLE_MECHANICUS_CARD_CHALLENGE_FAIL":     2,
		"IN_BATTLE_MECHANICUS_CARD_CHALLENGE_SUCCESS":  3,
	}
)

func (x InBattleMechanicusCardChallengeState) Enum() *InBattleMechanicusCardChallengeState {
	p := new(InBattleMechanicusCardChallengeState)
	*p = x
	return p
}

func (x InBattleMechanicusCardChallengeState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (InBattleMechanicusCardChallengeState) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_InBattleMechanicusCardChallengeState_proto_enumTypes[0].Descriptor()
}

func (InBattleMechanicusCardChallengeState) Type() protoreflect.EnumType {
	return &file_proto_InBattleMechanicusCardChallengeState_proto_enumTypes[0]
}

func (x InBattleMechanicusCardChallengeState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use InBattleMechanicusCardChallengeState.Descriptor instead.
func (InBattleMechanicusCardChallengeState) EnumDescriptor() ([]byte, []int) {
	return file_proto_InBattleMechanicusCardChallengeState_proto_rawDescGZIP(), []int{0}
}

var File_proto_InBattleMechanicusCardChallengeState_proto protoreflect.FileDescriptor

var file_proto_InBattleMechanicusCardChallengeState_proto_rawDesc = []byte{
	0x0a, 0x30, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x49, 0x6e, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65,
	0x4d, 0x65, 0x63, 0x68, 0x61, 0x6e, 0x69, 0x63, 0x75, 0x73, 0x43, 0x61, 0x72, 0x64, 0x43, 0x68,
	0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2a, 0xe5, 0x01, 0x0a, 0x24, 0x49, 0x6e, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x4d,
	0x65, 0x63, 0x68, 0x61, 0x6e, 0x69, 0x63, 0x75, 0x73, 0x43, 0x61, 0x72, 0x64, 0x43, 0x68, 0x61,
	0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x2c, 0x0a, 0x28, 0x49,
	0x4e, 0x5f, 0x42, 0x41, 0x54, 0x54, 0x4c, 0x45, 0x5f, 0x4d, 0x45, 0x43, 0x48, 0x41, 0x4e, 0x49,
	0x43, 0x55, 0x53, 0x5f, 0x43, 0x41, 0x52, 0x44, 0x5f, 0x43, 0x48, 0x41, 0x4c, 0x4c, 0x45, 0x4e,
	0x47, 0x45, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x30, 0x0a, 0x2c, 0x49, 0x4e, 0x5f,
	0x42, 0x41, 0x54, 0x54, 0x4c, 0x45, 0x5f, 0x4d, 0x45, 0x43, 0x48, 0x41, 0x4e, 0x49, 0x43, 0x55,
	0x53, 0x5f, 0x43, 0x41, 0x52, 0x44, 0x5f, 0x43, 0x48, 0x41, 0x4c, 0x4c, 0x45, 0x4e, 0x47, 0x45,
	0x5f, 0x4f, 0x4e, 0x5f, 0x47, 0x4f, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x2c, 0x0a, 0x28, 0x49,
	0x4e, 0x5f, 0x42, 0x41, 0x54, 0x54, 0x4c, 0x45, 0x5f, 0x4d, 0x45, 0x43, 0x48, 0x41, 0x4e, 0x49,
	0x43, 0x55, 0x53, 0x5f, 0x43, 0x41, 0x52, 0x44, 0x5f, 0x43, 0x48, 0x41, 0x4c, 0x4c, 0x45, 0x4e,
	0x47, 0x45, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x10, 0x02, 0x12, 0x2f, 0x0a, 0x2b, 0x49, 0x4e, 0x5f,
	0x42, 0x41, 0x54, 0x54, 0x4c, 0x45, 0x5f, 0x4d, 0x45, 0x43, 0x48, 0x41, 0x4e, 0x49, 0x43, 0x55,
	0x53, 0x5f, 0x43, 0x41, 0x52, 0x44, 0x5f, 0x43, 0x48, 0x41, 0x4c, 0x4c, 0x45, 0x4e, 0x47, 0x45,
	0x5f, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x03, 0x42, 0x1d, 0x5a, 0x1b, 0x47, 0x6f,
	0x2d, 0x47, 0x72, 0x61, 0x73, 0x73, 0x63, 0x75, 0x74, 0x74, 0x65, 0x72, 0x2f, 0x67, 0x65, 0x6e,
	0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_proto_InBattleMechanicusCardChallengeState_proto_rawDescOnce sync.Once
	file_proto_InBattleMechanicusCardChallengeState_proto_rawDescData = file_proto_InBattleMechanicusCardChallengeState_proto_rawDesc
)

func file_proto_InBattleMechanicusCardChallengeState_proto_rawDescGZIP() []byte {
	file_proto_InBattleMechanicusCardChallengeState_proto_rawDescOnce.Do(func() {
		file_proto_InBattleMechanicusCardChallengeState_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_InBattleMechanicusCardChallengeState_proto_rawDescData)
	})
	return file_proto_InBattleMechanicusCardChallengeState_proto_rawDescData
}

var file_proto_InBattleMechanicusCardChallengeState_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_InBattleMechanicusCardChallengeState_proto_goTypes = []interface{}{
	(InBattleMechanicusCardChallengeState)(0), // 0: InBattleMechanicusCardChallengeState
}
var file_proto_InBattleMechanicusCardChallengeState_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_InBattleMechanicusCardChallengeState_proto_init() }
func file_proto_InBattleMechanicusCardChallengeState_proto_init() {
	if File_proto_InBattleMechanicusCardChallengeState_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_InBattleMechanicusCardChallengeState_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_InBattleMechanicusCardChallengeState_proto_goTypes,
		DependencyIndexes: file_proto_InBattleMechanicusCardChallengeState_proto_depIdxs,
		EnumInfos:         file_proto_InBattleMechanicusCardChallengeState_proto_enumTypes,
	}.Build()
	File_proto_InBattleMechanicusCardChallengeState_proto = out.File
	file_proto_InBattleMechanicusCardChallengeState_proto_rawDesc = nil
	file_proto_InBattleMechanicusCardChallengeState_proto_goTypes = nil
	file_proto_InBattleMechanicusCardChallengeState_proto_depIdxs = nil
}
