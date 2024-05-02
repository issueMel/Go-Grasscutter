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
// source: proto/KPKJMBEINMG.proto

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

// CmdId: 7502
type KPKJMBEINMG struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HPPCHAPOFGF uint32       `protobuf:"varint,10,opt,name=HPPCHAPOFGF,proto3" json:"HPPCHAPOFGF,omitempty"`
	Retcode     int32        `protobuf:"varint,11,opt,name=retcode,proto3" json:"retcode,omitempty"`
	ScheduleId  uint32       `protobuf:"varint,9,opt,name=schedule_id,json=scheduleId,proto3" json:"schedule_id,omitempty"`
	PLOKMIMHAKI *FLOFKPMGNLA `protobuf:"bytes,12,opt,name=PLOKMIMHAKI,proto3" json:"PLOKMIMHAKI,omitempty"`
}

func (x *KPKJMBEINMG) Reset() {
	*x = KPKJMBEINMG{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_KPKJMBEINMG_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KPKJMBEINMG) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KPKJMBEINMG) ProtoMessage() {}

func (x *KPKJMBEINMG) ProtoReflect() protoreflect.Message {
	mi := &file_proto_KPKJMBEINMG_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KPKJMBEINMG.ProtoReflect.Descriptor instead.
func (*KPKJMBEINMG) Descriptor() ([]byte, []int) {
	return file_proto_KPKJMBEINMG_proto_rawDescGZIP(), []int{0}
}

func (x *KPKJMBEINMG) GetHPPCHAPOFGF() uint32 {
	if x != nil {
		return x.HPPCHAPOFGF
	}
	return 0
}

func (x *KPKJMBEINMG) GetRetcode() int32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *KPKJMBEINMG) GetScheduleId() uint32 {
	if x != nil {
		return x.ScheduleId
	}
	return 0
}

func (x *KPKJMBEINMG) GetPLOKMIMHAKI() *FLOFKPMGNLA {
	if x != nil {
		return x.PLOKMIMHAKI
	}
	return nil
}

var File_proto_KPKJMBEINMG_proto protoreflect.FileDescriptor

var file_proto_KPKJMBEINMG_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x4b, 0x50, 0x4b, 0x4a, 0x4d, 0x42, 0x45, 0x49,
	0x4e, 0x4d, 0x47, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x46, 0x4c, 0x4f, 0x46, 0x4b, 0x50, 0x4d, 0x47, 0x4e, 0x4c, 0x41, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x9a, 0x01, 0x0a, 0x0b, 0x4b, 0x50, 0x4b, 0x4a, 0x4d, 0x42, 0x45, 0x49, 0x4e,
	0x4d, 0x47, 0x12, 0x20, 0x0a, 0x0b, 0x48, 0x50, 0x50, 0x43, 0x48, 0x41, 0x50, 0x4f, 0x46, 0x47,
	0x46, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x48, 0x50, 0x50, 0x43, 0x48, 0x41, 0x50,
	0x4f, 0x46, 0x47, 0x46, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x1f,
	0x0a, 0x0b, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0a, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x49, 0x64, 0x12,
	0x2e, 0x0a, 0x0b, 0x50, 0x4c, 0x4f, 0x4b, 0x4d, 0x49, 0x4d, 0x48, 0x41, 0x4b, 0x49, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x46, 0x4c, 0x4f, 0x46, 0x4b, 0x50, 0x4d, 0x47, 0x4e,
	0x4c, 0x41, 0x52, 0x0b, 0x50, 0x4c, 0x4f, 0x4b, 0x4d, 0x49, 0x4d, 0x48, 0x41, 0x4b, 0x49, 0x42,
	0x20, 0x5a, 0x1e, 0x47, 0x6f, 0x2d, 0x47, 0x72, 0x61, 0x73, 0x73, 0x63, 0x75, 0x74, 0x74, 0x65,
	0x72, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_KPKJMBEINMG_proto_rawDescOnce sync.Once
	file_proto_KPKJMBEINMG_proto_rawDescData = file_proto_KPKJMBEINMG_proto_rawDesc
)

func file_proto_KPKJMBEINMG_proto_rawDescGZIP() []byte {
	file_proto_KPKJMBEINMG_proto_rawDescOnce.Do(func() {
		file_proto_KPKJMBEINMG_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_KPKJMBEINMG_proto_rawDescData)
	})
	return file_proto_KPKJMBEINMG_proto_rawDescData
}

var file_proto_KPKJMBEINMG_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_KPKJMBEINMG_proto_goTypes = []interface{}{
	(*KPKJMBEINMG)(nil), // 0: KPKJMBEINMG
	(*FLOFKPMGNLA)(nil), // 1: FLOFKPMGNLA
}
var file_proto_KPKJMBEINMG_proto_depIdxs = []int32{
	1, // 0: KPKJMBEINMG.PLOKMIMHAKI:type_name -> FLOFKPMGNLA
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_KPKJMBEINMG_proto_init() }
func file_proto_KPKJMBEINMG_proto_init() {
	if File_proto_KPKJMBEINMG_proto != nil {
		return
	}
	file_proto_FLOFKPMGNLA_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_KPKJMBEINMG_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KPKJMBEINMG); i {
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
			RawDescriptor: file_proto_KPKJMBEINMG_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_KPKJMBEINMG_proto_goTypes,
		DependencyIndexes: file_proto_KPKJMBEINMG_proto_depIdxs,
		MessageInfos:      file_proto_KPKJMBEINMG_proto_msgTypes,
	}.Build()
	File_proto_KPKJMBEINMG_proto = out.File
	file_proto_KPKJMBEINMG_proto_rawDesc = nil
	file_proto_KPKJMBEINMG_proto_goTypes = nil
	file_proto_KPKJMBEINMG_proto_depIdxs = nil
}
