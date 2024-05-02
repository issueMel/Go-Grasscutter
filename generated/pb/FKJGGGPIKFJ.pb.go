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
// source: proto/FKJGGGPIKFJ.proto

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

type FKJGGGPIKFJ struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	INHKDNJEOAP uint32 `protobuf:"varint,11,opt,name=INHKDNJEOAP,proto3" json:"INHKDNJEOAP,omitempty"`
	LLPGPDBBLDI uint32 `protobuf:"varint,10,opt,name=LLPGPDBBLDI,proto3" json:"LLPGPDBBLDI,omitempty"`
	CurScore    uint32 `protobuf:"varint,1,opt,name=cur_score,json=curScore,proto3" json:"cur_score,omitempty"`
	KillNum     uint32 `protobuf:"varint,15,opt,name=kill_num,json=killNum,proto3" json:"kill_num,omitempty"`
	DGIGCHJADFA int64  `protobuf:"fixed64,4,opt,name=DGIGCHJADFA,proto3" json:"DGIGCHJADFA,omitempty"`
	JJELBNOKOFP int64  `protobuf:"fixed64,3,opt,name=JJELBNOKOFP,proto3" json:"JJELBNOKOFP,omitempty"`
	Uid         uint32 `protobuf:"varint,7,opt,name=uid,proto3" json:"uid,omitempty"`
	HOJGCOLAPCM int64  `protobuf:"fixed64,12,opt,name=HOJGCOLAPCM,proto3" json:"HOJGCOLAPCM,omitempty"`
}

func (x *FKJGGGPIKFJ) Reset() {
	*x = FKJGGGPIKFJ{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_FKJGGGPIKFJ_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FKJGGGPIKFJ) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FKJGGGPIKFJ) ProtoMessage() {}

func (x *FKJGGGPIKFJ) ProtoReflect() protoreflect.Message {
	mi := &file_proto_FKJGGGPIKFJ_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FKJGGGPIKFJ.ProtoReflect.Descriptor instead.
func (*FKJGGGPIKFJ) Descriptor() ([]byte, []int) {
	return file_proto_FKJGGGPIKFJ_proto_rawDescGZIP(), []int{0}
}

func (x *FKJGGGPIKFJ) GetINHKDNJEOAP() uint32 {
	if x != nil {
		return x.INHKDNJEOAP
	}
	return 0
}

func (x *FKJGGGPIKFJ) GetLLPGPDBBLDI() uint32 {
	if x != nil {
		return x.LLPGPDBBLDI
	}
	return 0
}

func (x *FKJGGGPIKFJ) GetCurScore() uint32 {
	if x != nil {
		return x.CurScore
	}
	return 0
}

func (x *FKJGGGPIKFJ) GetKillNum() uint32 {
	if x != nil {
		return x.KillNum
	}
	return 0
}

func (x *FKJGGGPIKFJ) GetDGIGCHJADFA() int64 {
	if x != nil {
		return x.DGIGCHJADFA
	}
	return 0
}

func (x *FKJGGGPIKFJ) GetJJELBNOKOFP() int64 {
	if x != nil {
		return x.JJELBNOKOFP
	}
	return 0
}

func (x *FKJGGGPIKFJ) GetUid() uint32 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *FKJGGGPIKFJ) GetHOJGCOLAPCM() int64 {
	if x != nil {
		return x.HOJGCOLAPCM
	}
	return 0
}

var File_proto_FKJGGGPIKFJ_proto protoreflect.FileDescriptor

var file_proto_FKJGGGPIKFJ_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x46, 0x4b, 0x4a, 0x47, 0x47, 0x47, 0x50, 0x49,
	0x4b, 0x46, 0x4a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x81, 0x02, 0x0a, 0x0b, 0x46, 0x4b,
	0x4a, 0x47, 0x47, 0x47, 0x50, 0x49, 0x4b, 0x46, 0x4a, 0x12, 0x20, 0x0a, 0x0b, 0x49, 0x4e, 0x48,
	0x4b, 0x44, 0x4e, 0x4a, 0x45, 0x4f, 0x41, 0x50, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b,
	0x49, 0x4e, 0x48, 0x4b, 0x44, 0x4e, 0x4a, 0x45, 0x4f, 0x41, 0x50, 0x12, 0x20, 0x0a, 0x0b, 0x4c,
	0x4c, 0x50, 0x47, 0x50, 0x44, 0x42, 0x42, 0x4c, 0x44, 0x49, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0b, 0x4c, 0x4c, 0x50, 0x47, 0x50, 0x44, 0x42, 0x42, 0x4c, 0x44, 0x49, 0x12, 0x1b, 0x0a,
	0x09, 0x63, 0x75, 0x72, 0x5f, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x08, 0x63, 0x75, 0x72, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6b, 0x69,
	0x6c, 0x6c, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x6b, 0x69,
	0x6c, 0x6c, 0x4e, 0x75, 0x6d, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x47, 0x49, 0x47, 0x43, 0x48, 0x4a,
	0x41, 0x44, 0x46, 0x41, 0x18, 0x04, 0x20, 0x01, 0x28, 0x10, 0x52, 0x0b, 0x44, 0x47, 0x49, 0x47,
	0x43, 0x48, 0x4a, 0x41, 0x44, 0x46, 0x41, 0x12, 0x20, 0x0a, 0x0b, 0x4a, 0x4a, 0x45, 0x4c, 0x42,
	0x4e, 0x4f, 0x4b, 0x4f, 0x46, 0x50, 0x18, 0x03, 0x20, 0x01, 0x28, 0x10, 0x52, 0x0b, 0x4a, 0x4a,
	0x45, 0x4c, 0x42, 0x4e, 0x4f, 0x4b, 0x4f, 0x46, 0x50, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x48,
	0x4f, 0x4a, 0x47, 0x43, 0x4f, 0x4c, 0x41, 0x50, 0x43, 0x4d, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x10,
	0x52, 0x0b, 0x48, 0x4f, 0x4a, 0x47, 0x43, 0x4f, 0x4c, 0x41, 0x50, 0x43, 0x4d, 0x42, 0x1d, 0x5a,
	0x1b, 0x47, 0x6f, 0x2d, 0x47, 0x72, 0x61, 0x73, 0x73, 0x63, 0x75, 0x74, 0x74, 0x65, 0x72, 0x2f,
	0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_FKJGGGPIKFJ_proto_rawDescOnce sync.Once
	file_proto_FKJGGGPIKFJ_proto_rawDescData = file_proto_FKJGGGPIKFJ_proto_rawDesc
)

func file_proto_FKJGGGPIKFJ_proto_rawDescGZIP() []byte {
	file_proto_FKJGGGPIKFJ_proto_rawDescOnce.Do(func() {
		file_proto_FKJGGGPIKFJ_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_FKJGGGPIKFJ_proto_rawDescData)
	})
	return file_proto_FKJGGGPIKFJ_proto_rawDescData
}

var file_proto_FKJGGGPIKFJ_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_FKJGGGPIKFJ_proto_goTypes = []interface{}{
	(*FKJGGGPIKFJ)(nil), // 0: FKJGGGPIKFJ
}
var file_proto_FKJGGGPIKFJ_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_FKJGGGPIKFJ_proto_init() }
func file_proto_FKJGGGPIKFJ_proto_init() {
	if File_proto_FKJGGGPIKFJ_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_FKJGGGPIKFJ_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FKJGGGPIKFJ); i {
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
			RawDescriptor: file_proto_FKJGGGPIKFJ_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_FKJGGGPIKFJ_proto_goTypes,
		DependencyIndexes: file_proto_FKJGGGPIKFJ_proto_depIdxs,
		MessageInfos:      file_proto_FKJGGGPIKFJ_proto_msgTypes,
	}.Build()
	File_proto_FKJGGGPIKFJ_proto = out.File
	file_proto_FKJGGGPIKFJ_proto_rawDesc = nil
	file_proto_FKJGGGPIKFJ_proto_goTypes = nil
	file_proto_FKJGGGPIKFJ_proto_depIdxs = nil
}