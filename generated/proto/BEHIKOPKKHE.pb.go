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
// source: proto/BEHIKOPKKHE.proto

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

type BEHIKOPKKHE int32

const (
	BEHIKOPKKHE_BEHIKOPKKHE_GcgGameNone                BEHIKOPKKHE = 0
	BEHIKOPKKHE_BEHIKOPKKHE_GcgGameGm                  BEHIKOPKKHE = 1
	BEHIKOPKKHE_BEHIKOPKKHE_GcgGameMatch               BEHIKOPKKHE = 2
	BEHIKOPKKHE_BEHIKOPKKHE_GcgGamePvp                 BEHIKOPKKHE = 3
	BEHIKOPKKHE_BEHIKOPKKHE_GcgGameTavernChallenge     BEHIKOPKKHE = 4
	BEHIKOPKKHE_BEHIKOPKKHE_GcgGameConstChallenge      BEHIKOPKKHE = 5
	BEHIKOPKKHE_BEHIKOPKKHE_GcgGameWorldChallenge      BEHIKOPKKHE = 6
	BEHIKOPKKHE_BEHIKOPKKHE_GcgGameBossChallenge       BEHIKOPKKHE = 7
	BEHIKOPKKHE_BEHIKOPKKHE_GcgGameWeekChallenge       BEHIKOPKKHE = 8
	BEHIKOPKKHE_BEHIKOPKKHE_GcgGameBreakChallenge      BEHIKOPKKHE = 9
	BEHIKOPKKHE_BEHIKOPKKHE_GcgGameQuest               BEHIKOPKKHE = 10
	BEHIKOPKKHE_BEHIKOPKKHE_GcgGameGuideGroup          BEHIKOPKKHE = 11
	BEHIKOPKKHE_BEHIKOPKKHE_GcgGameOther               BEHIKOPKKHE = 12
	BEHIKOPKKHE_BEHIKOPKKHE_GcgGameActivityJourney     BEHIKOPKKHE = 13
	BEHIKOPKKHE_BEHIKOPKKHE_GcgGamePveActivity         BEHIKOPKKHE = 14
	BEHIKOPKKHE_BEHIKOPKKHE_GcgGameArena               BEHIKOPKKHE = 15
	BEHIKOPKKHE_BEHIKOPKKHE_GcgGameActivityPveInfinite BEHIKOPKKHE = 16
	BEHIKOPKKHE_BEHIKOPKKHE_GcgGameActivityPvePuzzle   BEHIKOPKKHE = 17
)

// Enum value maps for BEHIKOPKKHE.
var (
	BEHIKOPKKHE_name = map[int32]string{
		0:  "BEHIKOPKKHE_GcgGameNone",
		1:  "BEHIKOPKKHE_GcgGameGm",
		2:  "BEHIKOPKKHE_GcgGameMatch",
		3:  "BEHIKOPKKHE_GcgGamePvp",
		4:  "BEHIKOPKKHE_GcgGameTavernChallenge",
		5:  "BEHIKOPKKHE_GcgGameConstChallenge",
		6:  "BEHIKOPKKHE_GcgGameWorldChallenge",
		7:  "BEHIKOPKKHE_GcgGameBossChallenge",
		8:  "BEHIKOPKKHE_GcgGameWeekChallenge",
		9:  "BEHIKOPKKHE_GcgGameBreakChallenge",
		10: "BEHIKOPKKHE_GcgGameQuest",
		11: "BEHIKOPKKHE_GcgGameGuideGroup",
		12: "BEHIKOPKKHE_GcgGameOther",
		13: "BEHIKOPKKHE_GcgGameActivityJourney",
		14: "BEHIKOPKKHE_GcgGamePveActivity",
		15: "BEHIKOPKKHE_GcgGameArena",
		16: "BEHIKOPKKHE_GcgGameActivityPveInfinite",
		17: "BEHIKOPKKHE_GcgGameActivityPvePuzzle",
	}
	BEHIKOPKKHE_value = map[string]int32{
		"BEHIKOPKKHE_GcgGameNone":                0,
		"BEHIKOPKKHE_GcgGameGm":                  1,
		"BEHIKOPKKHE_GcgGameMatch":               2,
		"BEHIKOPKKHE_GcgGamePvp":                 3,
		"BEHIKOPKKHE_GcgGameTavernChallenge":     4,
		"BEHIKOPKKHE_GcgGameConstChallenge":      5,
		"BEHIKOPKKHE_GcgGameWorldChallenge":      6,
		"BEHIKOPKKHE_GcgGameBossChallenge":       7,
		"BEHIKOPKKHE_GcgGameWeekChallenge":       8,
		"BEHIKOPKKHE_GcgGameBreakChallenge":      9,
		"BEHIKOPKKHE_GcgGameQuest":               10,
		"BEHIKOPKKHE_GcgGameGuideGroup":          11,
		"BEHIKOPKKHE_GcgGameOther":               12,
		"BEHIKOPKKHE_GcgGameActivityJourney":     13,
		"BEHIKOPKKHE_GcgGamePveActivity":         14,
		"BEHIKOPKKHE_GcgGameArena":               15,
		"BEHIKOPKKHE_GcgGameActivityPveInfinite": 16,
		"BEHIKOPKKHE_GcgGameActivityPvePuzzle":   17,
	}
)

func (x BEHIKOPKKHE) Enum() *BEHIKOPKKHE {
	p := new(BEHIKOPKKHE)
	*p = x
	return p
}

func (x BEHIKOPKKHE) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BEHIKOPKKHE) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_BEHIKOPKKHE_proto_enumTypes[0].Descriptor()
}

func (BEHIKOPKKHE) Type() protoreflect.EnumType {
	return &file_proto_BEHIKOPKKHE_proto_enumTypes[0]
}

func (x BEHIKOPKKHE) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BEHIKOPKKHE.Descriptor instead.
func (BEHIKOPKKHE) EnumDescriptor() ([]byte, []int) {
	return file_proto_BEHIKOPKKHE_proto_rawDescGZIP(), []int{0}
}

var File_proto_BEHIKOPKKHE_proto protoreflect.FileDescriptor

var file_proto_BEHIKOPKKHE_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x42, 0x45, 0x48, 0x49, 0x4b, 0x4f, 0x50, 0x4b,
	0x4b, 0x48, 0x45, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0x87, 0x05, 0x0a, 0x0b, 0x42, 0x45,
	0x48, 0x49, 0x4b, 0x4f, 0x50, 0x4b, 0x4b, 0x48, 0x45, 0x12, 0x1b, 0x0a, 0x17, 0x42, 0x45, 0x48,
	0x49, 0x4b, 0x4f, 0x50, 0x4b, 0x4b, 0x48, 0x45, 0x5f, 0x47, 0x63, 0x67, 0x47, 0x61, 0x6d, 0x65,
	0x4e, 0x6f, 0x6e, 0x65, 0x10, 0x00, 0x12, 0x19, 0x0a, 0x15, 0x42, 0x45, 0x48, 0x49, 0x4b, 0x4f,
	0x50, 0x4b, 0x4b, 0x48, 0x45, 0x5f, 0x47, 0x63, 0x67, 0x47, 0x61, 0x6d, 0x65, 0x47, 0x6d, 0x10,
	0x01, 0x12, 0x1c, 0x0a, 0x18, 0x42, 0x45, 0x48, 0x49, 0x4b, 0x4f, 0x50, 0x4b, 0x4b, 0x48, 0x45,
	0x5f, 0x47, 0x63, 0x67, 0x47, 0x61, 0x6d, 0x65, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x10, 0x02, 0x12,
	0x1a, 0x0a, 0x16, 0x42, 0x45, 0x48, 0x49, 0x4b, 0x4f, 0x50, 0x4b, 0x4b, 0x48, 0x45, 0x5f, 0x47,
	0x63, 0x67, 0x47, 0x61, 0x6d, 0x65, 0x50, 0x76, 0x70, 0x10, 0x03, 0x12, 0x26, 0x0a, 0x22, 0x42,
	0x45, 0x48, 0x49, 0x4b, 0x4f, 0x50, 0x4b, 0x4b, 0x48, 0x45, 0x5f, 0x47, 0x63, 0x67, 0x47, 0x61,
	0x6d, 0x65, 0x54, 0x61, 0x76, 0x65, 0x72, 0x6e, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67,
	0x65, 0x10, 0x04, 0x12, 0x25, 0x0a, 0x21, 0x42, 0x45, 0x48, 0x49, 0x4b, 0x4f, 0x50, 0x4b, 0x4b,
	0x48, 0x45, 0x5f, 0x47, 0x63, 0x67, 0x47, 0x61, 0x6d, 0x65, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x43,
	0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x10, 0x05, 0x12, 0x25, 0x0a, 0x21, 0x42, 0x45,
	0x48, 0x49, 0x4b, 0x4f, 0x50, 0x4b, 0x4b, 0x48, 0x45, 0x5f, 0x47, 0x63, 0x67, 0x47, 0x61, 0x6d,
	0x65, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x10,
	0x06, 0x12, 0x24, 0x0a, 0x20, 0x42, 0x45, 0x48, 0x49, 0x4b, 0x4f, 0x50, 0x4b, 0x4b, 0x48, 0x45,
	0x5f, 0x47, 0x63, 0x67, 0x47, 0x61, 0x6d, 0x65, 0x42, 0x6f, 0x73, 0x73, 0x43, 0x68, 0x61, 0x6c,
	0x6c, 0x65, 0x6e, 0x67, 0x65, 0x10, 0x07, 0x12, 0x24, 0x0a, 0x20, 0x42, 0x45, 0x48, 0x49, 0x4b,
	0x4f, 0x50, 0x4b, 0x4b, 0x48, 0x45, 0x5f, 0x47, 0x63, 0x67, 0x47, 0x61, 0x6d, 0x65, 0x57, 0x65,
	0x65, 0x6b, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x10, 0x08, 0x12, 0x25, 0x0a,
	0x21, 0x42, 0x45, 0x48, 0x49, 0x4b, 0x4f, 0x50, 0x4b, 0x4b, 0x48, 0x45, 0x5f, 0x47, 0x63, 0x67,
	0x47, 0x61, 0x6d, 0x65, 0x42, 0x72, 0x65, 0x61, 0x6b, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e,
	0x67, 0x65, 0x10, 0x09, 0x12, 0x1c, 0x0a, 0x18, 0x42, 0x45, 0x48, 0x49, 0x4b, 0x4f, 0x50, 0x4b,
	0x4b, 0x48, 0x45, 0x5f, 0x47, 0x63, 0x67, 0x47, 0x61, 0x6d, 0x65, 0x51, 0x75, 0x65, 0x73, 0x74,
	0x10, 0x0a, 0x12, 0x21, 0x0a, 0x1d, 0x42, 0x45, 0x48, 0x49, 0x4b, 0x4f, 0x50, 0x4b, 0x4b, 0x48,
	0x45, 0x5f, 0x47, 0x63, 0x67, 0x47, 0x61, 0x6d, 0x65, 0x47, 0x75, 0x69, 0x64, 0x65, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x10, 0x0b, 0x12, 0x1c, 0x0a, 0x18, 0x42, 0x45, 0x48, 0x49, 0x4b, 0x4f, 0x50,
	0x4b, 0x4b, 0x48, 0x45, 0x5f, 0x47, 0x63, 0x67, 0x47, 0x61, 0x6d, 0x65, 0x4f, 0x74, 0x68, 0x65,
	0x72, 0x10, 0x0c, 0x12, 0x26, 0x0a, 0x22, 0x42, 0x45, 0x48, 0x49, 0x4b, 0x4f, 0x50, 0x4b, 0x4b,
	0x48, 0x45, 0x5f, 0x47, 0x63, 0x67, 0x47, 0x61, 0x6d, 0x65, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69,
	0x74, 0x79, 0x4a, 0x6f, 0x75, 0x72, 0x6e, 0x65, 0x79, 0x10, 0x0d, 0x12, 0x22, 0x0a, 0x1e, 0x42,
	0x45, 0x48, 0x49, 0x4b, 0x4f, 0x50, 0x4b, 0x4b, 0x48, 0x45, 0x5f, 0x47, 0x63, 0x67, 0x47, 0x61,
	0x6d, 0x65, 0x50, 0x76, 0x65, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x10, 0x0e, 0x12,
	0x1c, 0x0a, 0x18, 0x42, 0x45, 0x48, 0x49, 0x4b, 0x4f, 0x50, 0x4b, 0x4b, 0x48, 0x45, 0x5f, 0x47,
	0x63, 0x67, 0x47, 0x61, 0x6d, 0x65, 0x41, 0x72, 0x65, 0x6e, 0x61, 0x10, 0x0f, 0x12, 0x2a, 0x0a,
	0x26, 0x42, 0x45, 0x48, 0x49, 0x4b, 0x4f, 0x50, 0x4b, 0x4b, 0x48, 0x45, 0x5f, 0x47, 0x63, 0x67,
	0x47, 0x61, 0x6d, 0x65, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x50, 0x76, 0x65, 0x49,
	0x6e, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x65, 0x10, 0x10, 0x12, 0x28, 0x0a, 0x24, 0x42, 0x45, 0x48,
	0x49, 0x4b, 0x4f, 0x50, 0x4b, 0x4b, 0x48, 0x45, 0x5f, 0x47, 0x63, 0x67, 0x47, 0x61, 0x6d, 0x65,
	0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x50, 0x76, 0x65, 0x50, 0x75, 0x7a, 0x7a, 0x6c,
	0x65, 0x10, 0x11, 0x42, 0x20, 0x5a, 0x1e, 0x47, 0x6f, 0x2d, 0x47, 0x72, 0x61, 0x73, 0x73, 0x63,
	0x75, 0x74, 0x74, 0x65, 0x72, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_BEHIKOPKKHE_proto_rawDescOnce sync.Once
	file_proto_BEHIKOPKKHE_proto_rawDescData = file_proto_BEHIKOPKKHE_proto_rawDesc
)

func file_proto_BEHIKOPKKHE_proto_rawDescGZIP() []byte {
	file_proto_BEHIKOPKKHE_proto_rawDescOnce.Do(func() {
		file_proto_BEHIKOPKKHE_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_BEHIKOPKKHE_proto_rawDescData)
	})
	return file_proto_BEHIKOPKKHE_proto_rawDescData
}

var file_proto_BEHIKOPKKHE_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_BEHIKOPKKHE_proto_goTypes = []interface{}{
	(BEHIKOPKKHE)(0), // 0: BEHIKOPKKHE
}
var file_proto_BEHIKOPKKHE_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_BEHIKOPKKHE_proto_init() }
func file_proto_BEHIKOPKKHE_proto_init() {
	if File_proto_BEHIKOPKKHE_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_BEHIKOPKKHE_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_BEHIKOPKKHE_proto_goTypes,
		DependencyIndexes: file_proto_BEHIKOPKKHE_proto_depIdxs,
		EnumInfos:         file_proto_BEHIKOPKKHE_proto_enumTypes,
	}.Build()
	File_proto_BEHIKOPKKHE_proto = out.File
	file_proto_BEHIKOPKKHE_proto_rawDesc = nil
	file_proto_BEHIKOPKKHE_proto_goTypes = nil
	file_proto_BEHIKOPKKHE_proto_depIdxs = nil
}
