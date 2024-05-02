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
// source: proto/EffigyChallengeV2SettleInfo.proto

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

// Obf: GKEGGCBBJNB
type EffigyChallengeV2SettleInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CMCCNBNFAJE             uint32 `protobuf:"varint,11,opt,name=CMCCNBNFAJE,proto3" json:"CMCCNBNFAJE,omitempty"`
	LFNCGGFLHPL             bool   `protobuf:"varint,9,opt,name=LFNCGGFLHPL,proto3" json:"LFNCGGFLHPL,omitempty"`
	DCAKCNEDDEB             uint32 `protobuf:"varint,12,opt,name=DCAKCNEDDEB,proto3" json:"DCAKCNEDDEB,omitempty"`
	ChallengeModeDifficulty uint32 `protobuf:"varint,3,opt,name=challenge_mode_difficulty,json=challengeModeDifficulty,proto3" json:"challenge_mode_difficulty,omitempty"`
	FOFHONJNIHG             bool   `protobuf:"varint,4,opt,name=FOFHONJNIHG,proto3" json:"FOFHONJNIHG,omitempty"`
	PHNDBBLLHDI             uint32 `protobuf:"varint,1,opt,name=PHNDBBLLHDI,proto3" json:"PHNDBBLLHDI,omitempty"`
}

func (x *EffigyChallengeV2SettleInfo) Reset() {
	*x = EffigyChallengeV2SettleInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_EffigyChallengeV2SettleInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EffigyChallengeV2SettleInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EffigyChallengeV2SettleInfo) ProtoMessage() {}

func (x *EffigyChallengeV2SettleInfo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_EffigyChallengeV2SettleInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EffigyChallengeV2SettleInfo.ProtoReflect.Descriptor instead.
func (*EffigyChallengeV2SettleInfo) Descriptor() ([]byte, []int) {
	return file_proto_EffigyChallengeV2SettleInfo_proto_rawDescGZIP(), []int{0}
}

func (x *EffigyChallengeV2SettleInfo) GetCMCCNBNFAJE() uint32 {
	if x != nil {
		return x.CMCCNBNFAJE
	}
	return 0
}

func (x *EffigyChallengeV2SettleInfo) GetLFNCGGFLHPL() bool {
	if x != nil {
		return x.LFNCGGFLHPL
	}
	return false
}

func (x *EffigyChallengeV2SettleInfo) GetDCAKCNEDDEB() uint32 {
	if x != nil {
		return x.DCAKCNEDDEB
	}
	return 0
}

func (x *EffigyChallengeV2SettleInfo) GetChallengeModeDifficulty() uint32 {
	if x != nil {
		return x.ChallengeModeDifficulty
	}
	return 0
}

func (x *EffigyChallengeV2SettleInfo) GetFOFHONJNIHG() bool {
	if x != nil {
		return x.FOFHONJNIHG
	}
	return false
}

func (x *EffigyChallengeV2SettleInfo) GetPHNDBBLLHDI() uint32 {
	if x != nil {
		return x.PHNDBBLLHDI
	}
	return 0
}

var File_proto_EffigyChallengeV2SettleInfo_proto protoreflect.FileDescriptor

var file_proto_EffigyChallengeV2SettleInfo_proto_rawDesc = []byte{
	0x0a, 0x27, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x45, 0x66, 0x66, 0x69, 0x67, 0x79, 0x43, 0x68,
	0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x56, 0x32, 0x53, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x49,
	0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x83, 0x02, 0x0a, 0x1b, 0x45, 0x66,
	0x66, 0x69, 0x67, 0x79, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x56, 0x32, 0x53,
	0x65, 0x74, 0x74, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x20, 0x0a, 0x0b, 0x43, 0x4d, 0x43,
	0x43, 0x4e, 0x42, 0x4e, 0x46, 0x41, 0x4a, 0x45, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b,
	0x43, 0x4d, 0x43, 0x43, 0x4e, 0x42, 0x4e, 0x46, 0x41, 0x4a, 0x45, 0x12, 0x20, 0x0a, 0x0b, 0x4c,
	0x46, 0x4e, 0x43, 0x47, 0x47, 0x46, 0x4c, 0x48, 0x50, 0x4c, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0b, 0x4c, 0x46, 0x4e, 0x43, 0x47, 0x47, 0x46, 0x4c, 0x48, 0x50, 0x4c, 0x12, 0x20, 0x0a,
	0x0b, 0x44, 0x43, 0x41, 0x4b, 0x43, 0x4e, 0x45, 0x44, 0x44, 0x45, 0x42, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0b, 0x44, 0x43, 0x41, 0x4b, 0x43, 0x4e, 0x45, 0x44, 0x44, 0x45, 0x42, 0x12,
	0x3a, 0x0a, 0x19, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x5f, 0x6d, 0x6f, 0x64,
	0x65, 0x5f, 0x64, 0x69, 0x66, 0x66, 0x69, 0x63, 0x75, 0x6c, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x17, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x4d, 0x6f, 0x64,
	0x65, 0x44, 0x69, 0x66, 0x66, 0x69, 0x63, 0x75, 0x6c, 0x74, 0x79, 0x12, 0x20, 0x0a, 0x0b, 0x46,
	0x4f, 0x46, 0x48, 0x4f, 0x4e, 0x4a, 0x4e, 0x49, 0x48, 0x47, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0b, 0x46, 0x4f, 0x46, 0x48, 0x4f, 0x4e, 0x4a, 0x4e, 0x49, 0x48, 0x47, 0x12, 0x20, 0x0a,
	0x0b, 0x50, 0x48, 0x4e, 0x44, 0x42, 0x42, 0x4c, 0x4c, 0x48, 0x44, 0x49, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0b, 0x50, 0x48, 0x4e, 0x44, 0x42, 0x42, 0x4c, 0x4c, 0x48, 0x44, 0x49, 0x42,
	0x20, 0x5a, 0x1e, 0x47, 0x6f, 0x2d, 0x47, 0x72, 0x61, 0x73, 0x73, 0x63, 0x75, 0x74, 0x74, 0x65,
	0x72, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_EffigyChallengeV2SettleInfo_proto_rawDescOnce sync.Once
	file_proto_EffigyChallengeV2SettleInfo_proto_rawDescData = file_proto_EffigyChallengeV2SettleInfo_proto_rawDesc
)

func file_proto_EffigyChallengeV2SettleInfo_proto_rawDescGZIP() []byte {
	file_proto_EffigyChallengeV2SettleInfo_proto_rawDescOnce.Do(func() {
		file_proto_EffigyChallengeV2SettleInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_EffigyChallengeV2SettleInfo_proto_rawDescData)
	})
	return file_proto_EffigyChallengeV2SettleInfo_proto_rawDescData
}

var file_proto_EffigyChallengeV2SettleInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_EffigyChallengeV2SettleInfo_proto_goTypes = []interface{}{
	(*EffigyChallengeV2SettleInfo)(nil), // 0: EffigyChallengeV2SettleInfo
}
var file_proto_EffigyChallengeV2SettleInfo_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_EffigyChallengeV2SettleInfo_proto_init() }
func file_proto_EffigyChallengeV2SettleInfo_proto_init() {
	if File_proto_EffigyChallengeV2SettleInfo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_EffigyChallengeV2SettleInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EffigyChallengeV2SettleInfo); i {
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
			RawDescriptor: file_proto_EffigyChallengeV2SettleInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_EffigyChallengeV2SettleInfo_proto_goTypes,
		DependencyIndexes: file_proto_EffigyChallengeV2SettleInfo_proto_depIdxs,
		MessageInfos:      file_proto_EffigyChallengeV2SettleInfo_proto_msgTypes,
	}.Build()
	File_proto_EffigyChallengeV2SettleInfo_proto = out.File
	file_proto_EffigyChallengeV2SettleInfo_proto_rawDesc = nil
	file_proto_EffigyChallengeV2SettleInfo_proto_goTypes = nil
	file_proto_EffigyChallengeV2SettleInfo_proto_depIdxs = nil
}
