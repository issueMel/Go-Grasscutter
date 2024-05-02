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
// source: proto/WindFieldGalleryChallengeInfoNotify.proto

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

// CmdId: 23657
// Obf: OEBEJDAIKIP
type WindFieldGalleryChallengeInfoNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsStart     bool   `protobuf:"varint,7,opt,name=is_start,json=isStart,proto3" json:"is_start,omitempty"`
	GCMNEADPFBH uint32 `protobuf:"varint,2,opt,name=GCMNEADPFBH,proto3" json:"GCMNEADPFBH,omitempty"`
	EHLLLIJFOJG uint32 `protobuf:"varint,3,opt,name=EHLLLIJFOJG,proto3" json:"EHLLLIJFOJG,omitempty"`
	OKJOJDABKGC uint32 `protobuf:"varint,9,opt,name=OKJOJDABKGC,proto3" json:"OKJOJDABKGC,omitempty"`
	DPOBEKFGJOP uint32 `protobuf:"varint,11,opt,name=DPOBEKFGJOP,proto3" json:"DPOBEKFGJOP,omitempty"`
	IsSuccess   bool   `protobuf:"varint,14,opt,name=is_success,json=isSuccess,proto3" json:"is_success,omitempty"`
	KJCLNIGPNDF uint32 `protobuf:"varint,5,opt,name=KJCLNIGPNDF,proto3" json:"KJCLNIGPNDF,omitempty"`
}

func (x *WindFieldGalleryChallengeInfoNotify) Reset() {
	*x = WindFieldGalleryChallengeInfoNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_WindFieldGalleryChallengeInfoNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WindFieldGalleryChallengeInfoNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WindFieldGalleryChallengeInfoNotify) ProtoMessage() {}

func (x *WindFieldGalleryChallengeInfoNotify) ProtoReflect() protoreflect.Message {
	mi := &file_proto_WindFieldGalleryChallengeInfoNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WindFieldGalleryChallengeInfoNotify.ProtoReflect.Descriptor instead.
func (*WindFieldGalleryChallengeInfoNotify) Descriptor() ([]byte, []int) {
	return file_proto_WindFieldGalleryChallengeInfoNotify_proto_rawDescGZIP(), []int{0}
}

func (x *WindFieldGalleryChallengeInfoNotify) GetIsStart() bool {
	if x != nil {
		return x.IsStart
	}
	return false
}

func (x *WindFieldGalleryChallengeInfoNotify) GetGCMNEADPFBH() uint32 {
	if x != nil {
		return x.GCMNEADPFBH
	}
	return 0
}

func (x *WindFieldGalleryChallengeInfoNotify) GetEHLLLIJFOJG() uint32 {
	if x != nil {
		return x.EHLLLIJFOJG
	}
	return 0
}

func (x *WindFieldGalleryChallengeInfoNotify) GetOKJOJDABKGC() uint32 {
	if x != nil {
		return x.OKJOJDABKGC
	}
	return 0
}

func (x *WindFieldGalleryChallengeInfoNotify) GetDPOBEKFGJOP() uint32 {
	if x != nil {
		return x.DPOBEKFGJOP
	}
	return 0
}

func (x *WindFieldGalleryChallengeInfoNotify) GetIsSuccess() bool {
	if x != nil {
		return x.IsSuccess
	}
	return false
}

func (x *WindFieldGalleryChallengeInfoNotify) GetKJCLNIGPNDF() uint32 {
	if x != nil {
		return x.KJCLNIGPNDF
	}
	return 0
}

var File_proto_WindFieldGalleryChallengeInfoNotify_proto protoreflect.FileDescriptor

var file_proto_WindFieldGalleryChallengeInfoNotify_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x57, 0x69, 0x6e, 0x64, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x47, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x79, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x89, 0x02, 0x0a, 0x23, 0x57, 0x69, 0x6e, 0x64, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x47,
	0x61, 0x6c, 0x6c, 0x65, 0x72, 0x79, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x49,
	0x6e, 0x66, 0x6f, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x73, 0x5f,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x73, 0x53,
	0x74, 0x61, 0x72, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x47, 0x43, 0x4d, 0x4e, 0x45, 0x41, 0x44, 0x50,
	0x46, 0x42, 0x48, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x47, 0x43, 0x4d, 0x4e, 0x45,
	0x41, 0x44, 0x50, 0x46, 0x42, 0x48, 0x12, 0x20, 0x0a, 0x0b, 0x45, 0x48, 0x4c, 0x4c, 0x4c, 0x49,
	0x4a, 0x46, 0x4f, 0x4a, 0x47, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x45, 0x48, 0x4c,
	0x4c, 0x4c, 0x49, 0x4a, 0x46, 0x4f, 0x4a, 0x47, 0x12, 0x20, 0x0a, 0x0b, 0x4f, 0x4b, 0x4a, 0x4f,
	0x4a, 0x44, 0x41, 0x42, 0x4b, 0x47, 0x43, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x4f,
	0x4b, 0x4a, 0x4f, 0x4a, 0x44, 0x41, 0x42, 0x4b, 0x47, 0x43, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x50,
	0x4f, 0x42, 0x45, 0x4b, 0x46, 0x47, 0x4a, 0x4f, 0x50, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0b, 0x44, 0x50, 0x4f, 0x42, 0x45, 0x4b, 0x46, 0x47, 0x4a, 0x4f, 0x50, 0x12, 0x1d, 0x0a, 0x0a,
	0x69, 0x73, 0x5f, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x09, 0x69, 0x73, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x4b,
	0x4a, 0x43, 0x4c, 0x4e, 0x49, 0x47, 0x50, 0x4e, 0x44, 0x46, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0b, 0x4b, 0x4a, 0x43, 0x4c, 0x4e, 0x49, 0x47, 0x50, 0x4e, 0x44, 0x46, 0x42, 0x20, 0x5a,
	0x1e, 0x47, 0x6f, 0x2d, 0x47, 0x72, 0x61, 0x73, 0x73, 0x63, 0x75, 0x74, 0x74, 0x65, 0x72, 0x2f,
	0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_WindFieldGalleryChallengeInfoNotify_proto_rawDescOnce sync.Once
	file_proto_WindFieldGalleryChallengeInfoNotify_proto_rawDescData = file_proto_WindFieldGalleryChallengeInfoNotify_proto_rawDesc
)

func file_proto_WindFieldGalleryChallengeInfoNotify_proto_rawDescGZIP() []byte {
	file_proto_WindFieldGalleryChallengeInfoNotify_proto_rawDescOnce.Do(func() {
		file_proto_WindFieldGalleryChallengeInfoNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_WindFieldGalleryChallengeInfoNotify_proto_rawDescData)
	})
	return file_proto_WindFieldGalleryChallengeInfoNotify_proto_rawDescData
}

var file_proto_WindFieldGalleryChallengeInfoNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_WindFieldGalleryChallengeInfoNotify_proto_goTypes = []interface{}{
	(*WindFieldGalleryChallengeInfoNotify)(nil), // 0: WindFieldGalleryChallengeInfoNotify
}
var file_proto_WindFieldGalleryChallengeInfoNotify_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_WindFieldGalleryChallengeInfoNotify_proto_init() }
func file_proto_WindFieldGalleryChallengeInfoNotify_proto_init() {
	if File_proto_WindFieldGalleryChallengeInfoNotify_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_WindFieldGalleryChallengeInfoNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WindFieldGalleryChallengeInfoNotify); i {
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
			RawDescriptor: file_proto_WindFieldGalleryChallengeInfoNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_WindFieldGalleryChallengeInfoNotify_proto_goTypes,
		DependencyIndexes: file_proto_WindFieldGalleryChallengeInfoNotify_proto_depIdxs,
		MessageInfos:      file_proto_WindFieldGalleryChallengeInfoNotify_proto_msgTypes,
	}.Build()
	File_proto_WindFieldGalleryChallengeInfoNotify_proto = out.File
	file_proto_WindFieldGalleryChallengeInfoNotify_proto_rawDesc = nil
	file_proto_WindFieldGalleryChallengeInfoNotify_proto_goTypes = nil
	file_proto_WindFieldGalleryChallengeInfoNotify_proto_depIdxs = nil
}
