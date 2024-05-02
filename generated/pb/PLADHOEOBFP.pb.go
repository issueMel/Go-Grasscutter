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
// source: proto/PLADHOEOBFP.proto

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

// CmdId: 2676
type PLADHOEOBFP struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SceneId     uint32         `protobuf:"varint,3,opt,name=scene_id,json=sceneId,proto3" json:"scene_id,omitempty"`
	OBIKGEGABND []*KDIMOOGPNIF `protobuf:"bytes,14,rep,name=OBIKGEGABND,proto3" json:"OBIKGEGABND,omitempty"`
}

func (x *PLADHOEOBFP) Reset() {
	*x = PLADHOEOBFP{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_PLADHOEOBFP_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PLADHOEOBFP) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PLADHOEOBFP) ProtoMessage() {}

func (x *PLADHOEOBFP) ProtoReflect() protoreflect.Message {
	mi := &file_proto_PLADHOEOBFP_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PLADHOEOBFP.ProtoReflect.Descriptor instead.
func (*PLADHOEOBFP) Descriptor() ([]byte, []int) {
	return file_proto_PLADHOEOBFP_proto_rawDescGZIP(), []int{0}
}

func (x *PLADHOEOBFP) GetSceneId() uint32 {
	if x != nil {
		return x.SceneId
	}
	return 0
}

func (x *PLADHOEOBFP) GetOBIKGEGABND() []*KDIMOOGPNIF {
	if x != nil {
		return x.OBIKGEGABND
	}
	return nil
}

var File_proto_PLADHOEOBFP_proto protoreflect.FileDescriptor

var file_proto_PLADHOEOBFP_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x50, 0x4c, 0x41, 0x44, 0x48, 0x4f, 0x45, 0x4f,
	0x42, 0x46, 0x50, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x4b, 0x44, 0x49, 0x4d, 0x4f, 0x4f, 0x47, 0x50, 0x4e, 0x49, 0x46, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x58, 0x0a, 0x0b, 0x50, 0x4c, 0x41, 0x44, 0x48, 0x4f, 0x45, 0x4f, 0x42, 0x46,
	0x50, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x63, 0x65, 0x6e, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x07, 0x73, 0x63, 0x65, 0x6e, 0x65, 0x49, 0x64, 0x12, 0x2e, 0x0a, 0x0b,
	0x4f, 0x42, 0x49, 0x4b, 0x47, 0x45, 0x47, 0x41, 0x42, 0x4e, 0x44, 0x18, 0x0e, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0c, 0x2e, 0x4b, 0x44, 0x49, 0x4d, 0x4f, 0x4f, 0x47, 0x50, 0x4e, 0x49, 0x46, 0x52,
	0x0b, 0x4f, 0x42, 0x49, 0x4b, 0x47, 0x45, 0x47, 0x41, 0x42, 0x4e, 0x44, 0x42, 0x1d, 0x5a, 0x1b,
	0x47, 0x6f, 0x2d, 0x47, 0x72, 0x61, 0x73, 0x73, 0x63, 0x75, 0x74, 0x74, 0x65, 0x72, 0x2f, 0x67,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_proto_PLADHOEOBFP_proto_rawDescOnce sync.Once
	file_proto_PLADHOEOBFP_proto_rawDescData = file_proto_PLADHOEOBFP_proto_rawDesc
)

func file_proto_PLADHOEOBFP_proto_rawDescGZIP() []byte {
	file_proto_PLADHOEOBFP_proto_rawDescOnce.Do(func() {
		file_proto_PLADHOEOBFP_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_PLADHOEOBFP_proto_rawDescData)
	})
	return file_proto_PLADHOEOBFP_proto_rawDescData
}

var file_proto_PLADHOEOBFP_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_PLADHOEOBFP_proto_goTypes = []interface{}{
	(*PLADHOEOBFP)(nil), // 0: PLADHOEOBFP
	(*KDIMOOGPNIF)(nil), // 1: KDIMOOGPNIF
}
var file_proto_PLADHOEOBFP_proto_depIdxs = []int32{
	1, // 0: PLADHOEOBFP.OBIKGEGABND:type_name -> KDIMOOGPNIF
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_PLADHOEOBFP_proto_init() }
func file_proto_PLADHOEOBFP_proto_init() {
	if File_proto_PLADHOEOBFP_proto != nil {
		return
	}
	file_proto_KDIMOOGPNIF_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_PLADHOEOBFP_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PLADHOEOBFP); i {
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
			RawDescriptor: file_proto_PLADHOEOBFP_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_PLADHOEOBFP_proto_goTypes,
		DependencyIndexes: file_proto_PLADHOEOBFP_proto_depIdxs,
		MessageInfos:      file_proto_PLADHOEOBFP_proto_msgTypes,
	}.Build()
	File_proto_PLADHOEOBFP_proto = out.File
	file_proto_PLADHOEOBFP_proto_rawDesc = nil
	file_proto_PLADHOEOBFP_proto_goTypes = nil
	file_proto_PLADHOEOBFP_proto_depIdxs = nil
}
