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
// source: proto/ADHOJAAOLNO.proto

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

type ADHOJAAOLNO int32

const (
	ADHOJAAOLNO_ADHOJAAOLNO_DungeonCandidateAvatarTypeNone   ADHOJAAOLNO = 0
	ADHOJAAOLNO_ADHOJAAOLNO_DungeonCandidateAvatarTypeFormal ADHOJAAOLNO = 1
	ADHOJAAOLNO_ADHOJAAOLNO_DungeonCandidateAvatarTypeUgc    ADHOJAAOLNO = 2
)

// Enum value maps for ADHOJAAOLNO.
var (
	ADHOJAAOLNO_name = map[int32]string{
		0: "ADHOJAAOLNO_DungeonCandidateAvatarTypeNone",
		1: "ADHOJAAOLNO_DungeonCandidateAvatarTypeFormal",
		2: "ADHOJAAOLNO_DungeonCandidateAvatarTypeUgc",
	}
	ADHOJAAOLNO_value = map[string]int32{
		"ADHOJAAOLNO_DungeonCandidateAvatarTypeNone":   0,
		"ADHOJAAOLNO_DungeonCandidateAvatarTypeFormal": 1,
		"ADHOJAAOLNO_DungeonCandidateAvatarTypeUgc":    2,
	}
)

func (x ADHOJAAOLNO) Enum() *ADHOJAAOLNO {
	p := new(ADHOJAAOLNO)
	*p = x
	return p
}

func (x ADHOJAAOLNO) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ADHOJAAOLNO) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_ADHOJAAOLNO_proto_enumTypes[0].Descriptor()
}

func (ADHOJAAOLNO) Type() protoreflect.EnumType {
	return &file_proto_ADHOJAAOLNO_proto_enumTypes[0]
}

func (x ADHOJAAOLNO) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ADHOJAAOLNO.Descriptor instead.
func (ADHOJAAOLNO) EnumDescriptor() ([]byte, []int) {
	return file_proto_ADHOJAAOLNO_proto_rawDescGZIP(), []int{0}
}

var File_proto_ADHOJAAOLNO_proto protoreflect.FileDescriptor

var file_proto_ADHOJAAOLNO_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x41, 0x44, 0x48, 0x4f, 0x4a, 0x41, 0x41, 0x4f,
	0x4c, 0x4e, 0x4f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0x9e, 0x01, 0x0a, 0x0b, 0x41, 0x44,
	0x48, 0x4f, 0x4a, 0x41, 0x41, 0x4f, 0x4c, 0x4e, 0x4f, 0x12, 0x2e, 0x0a, 0x2a, 0x41, 0x44, 0x48,
	0x4f, 0x4a, 0x41, 0x41, 0x4f, 0x4c, 0x4e, 0x4f, 0x5f, 0x44, 0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e,
	0x43, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x54,
	0x79, 0x70, 0x65, 0x4e, 0x6f, 0x6e, 0x65, 0x10, 0x00, 0x12, 0x30, 0x0a, 0x2c, 0x41, 0x44, 0x48,
	0x4f, 0x4a, 0x41, 0x41, 0x4f, 0x4c, 0x4e, 0x4f, 0x5f, 0x44, 0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e,
	0x43, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x54,
	0x79, 0x70, 0x65, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x10, 0x01, 0x12, 0x2d, 0x0a, 0x29, 0x41,
	0x44, 0x48, 0x4f, 0x4a, 0x41, 0x41, 0x4f, 0x4c, 0x4e, 0x4f, 0x5f, 0x44, 0x75, 0x6e, 0x67, 0x65,
	0x6f, 0x6e, 0x43, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x41, 0x76, 0x61, 0x74, 0x61,
	0x72, 0x54, 0x79, 0x70, 0x65, 0x55, 0x67, 0x63, 0x10, 0x02, 0x42, 0x1d, 0x5a, 0x1b, 0x47, 0x6f,
	0x2d, 0x47, 0x72, 0x61, 0x73, 0x73, 0x63, 0x75, 0x74, 0x74, 0x65, 0x72, 0x2f, 0x67, 0x65, 0x6e,
	0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_proto_ADHOJAAOLNO_proto_rawDescOnce sync.Once
	file_proto_ADHOJAAOLNO_proto_rawDescData = file_proto_ADHOJAAOLNO_proto_rawDesc
)

func file_proto_ADHOJAAOLNO_proto_rawDescGZIP() []byte {
	file_proto_ADHOJAAOLNO_proto_rawDescOnce.Do(func() {
		file_proto_ADHOJAAOLNO_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_ADHOJAAOLNO_proto_rawDescData)
	})
	return file_proto_ADHOJAAOLNO_proto_rawDescData
}

var file_proto_ADHOJAAOLNO_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_ADHOJAAOLNO_proto_goTypes = []interface{}{
	(ADHOJAAOLNO)(0), // 0: ADHOJAAOLNO
}
var file_proto_ADHOJAAOLNO_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_ADHOJAAOLNO_proto_init() }
func file_proto_ADHOJAAOLNO_proto_init() {
	if File_proto_ADHOJAAOLNO_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_ADHOJAAOLNO_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_ADHOJAAOLNO_proto_goTypes,
		DependencyIndexes: file_proto_ADHOJAAOLNO_proto_depIdxs,
		EnumInfos:         file_proto_ADHOJAAOLNO_proto_enumTypes,
	}.Build()
	File_proto_ADHOJAAOLNO_proto = out.File
	file_proto_ADHOJAAOLNO_proto_rawDesc = nil
	file_proto_ADHOJAAOLNO_proto_goTypes = nil
	file_proto_ADHOJAAOLNO_proto_depIdxs = nil
}
