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
// source: proto/KDHIAEFLGFM.proto

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

type KDHIAEFLGFM int32

const (
	KDHIAEFLGFM_KDHIAEFLGFM_GalleryStopNone                KDHIAEFLGFM = 0
	KDHIAEFLGFM_KDHIAEFLGFM_GalleryStopTimeup              KDHIAEFLGFM = 1
	KDHIAEFLGFM_KDHIAEFLGFM_GalleryStopClientInterrupt     KDHIAEFLGFM = 2
	KDHIAEFLGFM_KDHIAEFLGFM_GalleryStopLuaInterruptSuccess KDHIAEFLGFM = 3
	KDHIAEFLGFM_KDHIAEFLGFM_GalleryStopLuaInterruptFail    KDHIAEFLGFM = 4
	KDHIAEFLGFM_KDHIAEFLGFM_GalleryStopOwnerLeaveScene     KDHIAEFLGFM = 5
	KDHIAEFLGFM_KDHIAEFLGFM_GalleryStopPlayInitFailed      KDHIAEFLGFM = 6
	KDHIAEFLGFM_KDHIAEFLGFM_GalleryStopOtherPlayerEnter    KDHIAEFLGFM = 7
	KDHIAEFLGFM_KDHIAEFLGFM_GalleryStopAvatarDie           KDHIAEFLGFM = 8
	KDHIAEFLGFM_KDHIAEFLGFM_GalleryStopFinished            KDHIAEFLGFM = 9
	KDHIAEFLGFM_KDHIAEFLGFM_GalleryStopFungusAllDie        KDHIAEFLGFM = 10
	KDHIAEFLGFM_KDHIAEFLGFM_GalleryStopLifeCountZero       KDHIAEFLGFM = 11
	KDHIAEFLGFM_KDHIAEFLGFM_GalleryStopActivityClosed      KDHIAEFLGFM = 12
	KDHIAEFLGFM_KDHIAEFLGFM_GalleryStopLeaveRegionFail     KDHIAEFLGFM = 13
	KDHIAEFLGFM_KDHIAEFLGFM_GalleryStopHeartBloodUsedOut   KDHIAEFLGFM = 14
	KDHIAEFLGFM_KDHIAEFLGFM_GalleryStopGuardianStoneDie    KDHIAEFLGFM = 15
)

// Enum value maps for KDHIAEFLGFM.
var (
	KDHIAEFLGFM_name = map[int32]string{
		0:  "KDHIAEFLGFM_GalleryStopNone",
		1:  "KDHIAEFLGFM_GalleryStopTimeup",
		2:  "KDHIAEFLGFM_GalleryStopClientInterrupt",
		3:  "KDHIAEFLGFM_GalleryStopLuaInterruptSuccess",
		4:  "KDHIAEFLGFM_GalleryStopLuaInterruptFail",
		5:  "KDHIAEFLGFM_GalleryStopOwnerLeaveScene",
		6:  "KDHIAEFLGFM_GalleryStopPlayInitFailed",
		7:  "KDHIAEFLGFM_GalleryStopOtherPlayerEnter",
		8:  "KDHIAEFLGFM_GalleryStopAvatarDie",
		9:  "KDHIAEFLGFM_GalleryStopFinished",
		10: "KDHIAEFLGFM_GalleryStopFungusAllDie",
		11: "KDHIAEFLGFM_GalleryStopLifeCountZero",
		12: "KDHIAEFLGFM_GalleryStopActivityClosed",
		13: "KDHIAEFLGFM_GalleryStopLeaveRegionFail",
		14: "KDHIAEFLGFM_GalleryStopHeartBloodUsedOut",
		15: "KDHIAEFLGFM_GalleryStopGuardianStoneDie",
	}
	KDHIAEFLGFM_value = map[string]int32{
		"KDHIAEFLGFM_GalleryStopNone":                0,
		"KDHIAEFLGFM_GalleryStopTimeup":              1,
		"KDHIAEFLGFM_GalleryStopClientInterrupt":     2,
		"KDHIAEFLGFM_GalleryStopLuaInterruptSuccess": 3,
		"KDHIAEFLGFM_GalleryStopLuaInterruptFail":    4,
		"KDHIAEFLGFM_GalleryStopOwnerLeaveScene":     5,
		"KDHIAEFLGFM_GalleryStopPlayInitFailed":      6,
		"KDHIAEFLGFM_GalleryStopOtherPlayerEnter":    7,
		"KDHIAEFLGFM_GalleryStopAvatarDie":           8,
		"KDHIAEFLGFM_GalleryStopFinished":            9,
		"KDHIAEFLGFM_GalleryStopFungusAllDie":        10,
		"KDHIAEFLGFM_GalleryStopLifeCountZero":       11,
		"KDHIAEFLGFM_GalleryStopActivityClosed":      12,
		"KDHIAEFLGFM_GalleryStopLeaveRegionFail":     13,
		"KDHIAEFLGFM_GalleryStopHeartBloodUsedOut":   14,
		"KDHIAEFLGFM_GalleryStopGuardianStoneDie":    15,
	}
)

func (x KDHIAEFLGFM) Enum() *KDHIAEFLGFM {
	p := new(KDHIAEFLGFM)
	*p = x
	return p
}

func (x KDHIAEFLGFM) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (KDHIAEFLGFM) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_KDHIAEFLGFM_proto_enumTypes[0].Descriptor()
}

func (KDHIAEFLGFM) Type() protoreflect.EnumType {
	return &file_proto_KDHIAEFLGFM_proto_enumTypes[0]
}

func (x KDHIAEFLGFM) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use KDHIAEFLGFM.Descriptor instead.
func (KDHIAEFLGFM) EnumDescriptor() ([]byte, []int) {
	return file_proto_KDHIAEFLGFM_proto_rawDescGZIP(), []int{0}
}

var File_proto_KDHIAEFLGFM_proto protoreflect.FileDescriptor

var file_proto_KDHIAEFLGFM_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x4b, 0x44, 0x48, 0x49, 0x41, 0x45, 0x46, 0x4c,
	0x47, 0x46, 0x4d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0xae, 0x05, 0x0a, 0x0b, 0x4b, 0x44,
	0x48, 0x49, 0x41, 0x45, 0x46, 0x4c, 0x47, 0x46, 0x4d, 0x12, 0x1f, 0x0a, 0x1b, 0x4b, 0x44, 0x48,
	0x49, 0x41, 0x45, 0x46, 0x4c, 0x47, 0x46, 0x4d, 0x5f, 0x47, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x79,
	0x53, 0x74, 0x6f, 0x70, 0x4e, 0x6f, 0x6e, 0x65, 0x10, 0x00, 0x12, 0x21, 0x0a, 0x1d, 0x4b, 0x44,
	0x48, 0x49, 0x41, 0x45, 0x46, 0x4c, 0x47, 0x46, 0x4d, 0x5f, 0x47, 0x61, 0x6c, 0x6c, 0x65, 0x72,
	0x79, 0x53, 0x74, 0x6f, 0x70, 0x54, 0x69, 0x6d, 0x65, 0x75, 0x70, 0x10, 0x01, 0x12, 0x2a, 0x0a,
	0x26, 0x4b, 0x44, 0x48, 0x49, 0x41, 0x45, 0x46, 0x4c, 0x47, 0x46, 0x4d, 0x5f, 0x47, 0x61, 0x6c,
	0x6c, 0x65, 0x72, 0x79, 0x53, 0x74, 0x6f, 0x70, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x6e,
	0x74, 0x65, 0x72, 0x72, 0x75, 0x70, 0x74, 0x10, 0x02, 0x12, 0x2e, 0x0a, 0x2a, 0x4b, 0x44, 0x48,
	0x49, 0x41, 0x45, 0x46, 0x4c, 0x47, 0x46, 0x4d, 0x5f, 0x47, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x79,
	0x53, 0x74, 0x6f, 0x70, 0x4c, 0x75, 0x61, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x72, 0x75, 0x70, 0x74,
	0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x10, 0x03, 0x12, 0x2b, 0x0a, 0x27, 0x4b, 0x44, 0x48,
	0x49, 0x41, 0x45, 0x46, 0x4c, 0x47, 0x46, 0x4d, 0x5f, 0x47, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x79,
	0x53, 0x74, 0x6f, 0x70, 0x4c, 0x75, 0x61, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x72, 0x75, 0x70, 0x74,
	0x46, 0x61, 0x69, 0x6c, 0x10, 0x04, 0x12, 0x2a, 0x0a, 0x26, 0x4b, 0x44, 0x48, 0x49, 0x41, 0x45,
	0x46, 0x4c, 0x47, 0x46, 0x4d, 0x5f, 0x47, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x79, 0x53, 0x74, 0x6f,
	0x70, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x53, 0x63, 0x65, 0x6e, 0x65,
	0x10, 0x05, 0x12, 0x29, 0x0a, 0x25, 0x4b, 0x44, 0x48, 0x49, 0x41, 0x45, 0x46, 0x4c, 0x47, 0x46,
	0x4d, 0x5f, 0x47, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x79, 0x53, 0x74, 0x6f, 0x70, 0x50, 0x6c, 0x61,
	0x79, 0x49, 0x6e, 0x69, 0x74, 0x46, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x10, 0x06, 0x12, 0x2b, 0x0a,
	0x27, 0x4b, 0x44, 0x48, 0x49, 0x41, 0x45, 0x46, 0x4c, 0x47, 0x46, 0x4d, 0x5f, 0x47, 0x61, 0x6c,
	0x6c, 0x65, 0x72, 0x79, 0x53, 0x74, 0x6f, 0x70, 0x4f, 0x74, 0x68, 0x65, 0x72, 0x50, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x10, 0x07, 0x12, 0x24, 0x0a, 0x20, 0x4b, 0x44,
	0x48, 0x49, 0x41, 0x45, 0x46, 0x4c, 0x47, 0x46, 0x4d, 0x5f, 0x47, 0x61, 0x6c, 0x6c, 0x65, 0x72,
	0x79, 0x53, 0x74, 0x6f, 0x70, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x44, 0x69, 0x65, 0x10, 0x08,
	0x12, 0x23, 0x0a, 0x1f, 0x4b, 0x44, 0x48, 0x49, 0x41, 0x45, 0x46, 0x4c, 0x47, 0x46, 0x4d, 0x5f,
	0x47, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x79, 0x53, 0x74, 0x6f, 0x70, 0x46, 0x69, 0x6e, 0x69, 0x73,
	0x68, 0x65, 0x64, 0x10, 0x09, 0x12, 0x27, 0x0a, 0x23, 0x4b, 0x44, 0x48, 0x49, 0x41, 0x45, 0x46,
	0x4c, 0x47, 0x46, 0x4d, 0x5f, 0x47, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x79, 0x53, 0x74, 0x6f, 0x70,
	0x46, 0x75, 0x6e, 0x67, 0x75, 0x73, 0x41, 0x6c, 0x6c, 0x44, 0x69, 0x65, 0x10, 0x0a, 0x12, 0x28,
	0x0a, 0x24, 0x4b, 0x44, 0x48, 0x49, 0x41, 0x45, 0x46, 0x4c, 0x47, 0x46, 0x4d, 0x5f, 0x47, 0x61,
	0x6c, 0x6c, 0x65, 0x72, 0x79, 0x53, 0x74, 0x6f, 0x70, 0x4c, 0x69, 0x66, 0x65, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x5a, 0x65, 0x72, 0x6f, 0x10, 0x0b, 0x12, 0x29, 0x0a, 0x25, 0x4b, 0x44, 0x48, 0x49,
	0x41, 0x45, 0x46, 0x4c, 0x47, 0x46, 0x4d, 0x5f, 0x47, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x79, 0x53,
	0x74, 0x6f, 0x70, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x43, 0x6c, 0x6f, 0x73, 0x65,
	0x64, 0x10, 0x0c, 0x12, 0x2a, 0x0a, 0x26, 0x4b, 0x44, 0x48, 0x49, 0x41, 0x45, 0x46, 0x4c, 0x47,
	0x46, 0x4d, 0x5f, 0x47, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x79, 0x53, 0x74, 0x6f, 0x70, 0x4c, 0x65,
	0x61, 0x76, 0x65, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x46, 0x61, 0x69, 0x6c, 0x10, 0x0d, 0x12,
	0x2c, 0x0a, 0x28, 0x4b, 0x44, 0x48, 0x49, 0x41, 0x45, 0x46, 0x4c, 0x47, 0x46, 0x4d, 0x5f, 0x47,
	0x61, 0x6c, 0x6c, 0x65, 0x72, 0x79, 0x53, 0x74, 0x6f, 0x70, 0x48, 0x65, 0x61, 0x72, 0x74, 0x42,
	0x6c, 0x6f, 0x6f, 0x64, 0x55, 0x73, 0x65, 0x64, 0x4f, 0x75, 0x74, 0x10, 0x0e, 0x12, 0x2b, 0x0a,
	0x27, 0x4b, 0x44, 0x48, 0x49, 0x41, 0x45, 0x46, 0x4c, 0x47, 0x46, 0x4d, 0x5f, 0x47, 0x61, 0x6c,
	0x6c, 0x65, 0x72, 0x79, 0x53, 0x74, 0x6f, 0x70, 0x47, 0x75, 0x61, 0x72, 0x64, 0x69, 0x61, 0x6e,
	0x53, 0x74, 0x6f, 0x6e, 0x65, 0x44, 0x69, 0x65, 0x10, 0x0f, 0x42, 0x1d, 0x5a, 0x1b, 0x47, 0x6f,
	0x2d, 0x47, 0x72, 0x61, 0x73, 0x73, 0x63, 0x75, 0x74, 0x74, 0x65, 0x72, 0x2f, 0x67, 0x65, 0x6e,
	0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_proto_KDHIAEFLGFM_proto_rawDescOnce sync.Once
	file_proto_KDHIAEFLGFM_proto_rawDescData = file_proto_KDHIAEFLGFM_proto_rawDesc
)

func file_proto_KDHIAEFLGFM_proto_rawDescGZIP() []byte {
	file_proto_KDHIAEFLGFM_proto_rawDescOnce.Do(func() {
		file_proto_KDHIAEFLGFM_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_KDHIAEFLGFM_proto_rawDescData)
	})
	return file_proto_KDHIAEFLGFM_proto_rawDescData
}

var file_proto_KDHIAEFLGFM_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_KDHIAEFLGFM_proto_goTypes = []interface{}{
	(KDHIAEFLGFM)(0), // 0: KDHIAEFLGFM
}
var file_proto_KDHIAEFLGFM_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_KDHIAEFLGFM_proto_init() }
func file_proto_KDHIAEFLGFM_proto_init() {
	if File_proto_KDHIAEFLGFM_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_KDHIAEFLGFM_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_KDHIAEFLGFM_proto_goTypes,
		DependencyIndexes: file_proto_KDHIAEFLGFM_proto_depIdxs,
		EnumInfos:         file_proto_KDHIAEFLGFM_proto_enumTypes,
	}.Build()
	File_proto_KDHIAEFLGFM_proto = out.File
	file_proto_KDHIAEFLGFM_proto_rawDesc = nil
	file_proto_KDHIAEFLGFM_proto_goTypes = nil
	file_proto_KDHIAEFLGFM_proto_depIdxs = nil
}