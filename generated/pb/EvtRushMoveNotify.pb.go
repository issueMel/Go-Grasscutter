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
// source: proto/EvtRushMoveNotify.proto

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

// CmdId: 25987
// Obf: HLMODAGGEGO
type EvtRushMoveNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ForwardType     ForwardType      `protobuf:"varint,6,opt,name=forward_type,json=forwardType,proto3,enum=ForwardType" json:"forward_type,omitempty"`
	EvtRushMoveInfo *EvtRushMoveInfo `protobuf:"bytes,15,opt,name=evt_rush_move_info,json=evtRushMoveInfo,proto3" json:"evt_rush_move_info,omitempty"`
}

func (x *EvtRushMoveNotify) Reset() {
	*x = EvtRushMoveNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_EvtRushMoveNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EvtRushMoveNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EvtRushMoveNotify) ProtoMessage() {}

func (x *EvtRushMoveNotify) ProtoReflect() protoreflect.Message {
	mi := &file_proto_EvtRushMoveNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EvtRushMoveNotify.ProtoReflect.Descriptor instead.
func (*EvtRushMoveNotify) Descriptor() ([]byte, []int) {
	return file_proto_EvtRushMoveNotify_proto_rawDescGZIP(), []int{0}
}

func (x *EvtRushMoveNotify) GetForwardType() ForwardType {
	if x != nil {
		return x.ForwardType
	}
	return ForwardType_FORWARD_TYPE_LOCAL
}

func (x *EvtRushMoveNotify) GetEvtRushMoveInfo() *EvtRushMoveInfo {
	if x != nil {
		return x.EvtRushMoveInfo
	}
	return nil
}

var File_proto_EvtRushMoveNotify_proto protoreflect.FileDescriptor

var file_proto_EvtRushMoveNotify_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x45, 0x76, 0x74, 0x52, 0x75, 0x73, 0x68, 0x4d,
	0x6f, 0x76, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x46, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x54, 0x79,
	0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x45, 0x76, 0x74, 0x52, 0x75, 0x73, 0x68, 0x4d, 0x6f, 0x76, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x83, 0x01, 0x0a, 0x11, 0x45, 0x76, 0x74, 0x52, 0x75, 0x73,
	0x68, 0x4d, 0x6f, 0x76, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x2f, 0x0a, 0x0c, 0x66,
	0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x0c, 0x2e, 0x46, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x0b, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x3d, 0x0a, 0x12,
	0x65, 0x76, 0x74, 0x5f, 0x72, 0x75, 0x73, 0x68, 0x5f, 0x6d, 0x6f, 0x76, 0x65, 0x5f, 0x69, 0x6e,
	0x66, 0x6f, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x45, 0x76, 0x74, 0x52, 0x75,
	0x73, 0x68, 0x4d, 0x6f, 0x76, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0f, 0x65, 0x76, 0x74, 0x52,
	0x75, 0x73, 0x68, 0x4d, 0x6f, 0x76, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x1d, 0x5a, 0x1b, 0x47,
	0x6f, 0x2d, 0x47, 0x72, 0x61, 0x73, 0x73, 0x63, 0x75, 0x74, 0x74, 0x65, 0x72, 0x2f, 0x67, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_proto_EvtRushMoveNotify_proto_rawDescOnce sync.Once
	file_proto_EvtRushMoveNotify_proto_rawDescData = file_proto_EvtRushMoveNotify_proto_rawDesc
)

func file_proto_EvtRushMoveNotify_proto_rawDescGZIP() []byte {
	file_proto_EvtRushMoveNotify_proto_rawDescOnce.Do(func() {
		file_proto_EvtRushMoveNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_EvtRushMoveNotify_proto_rawDescData)
	})
	return file_proto_EvtRushMoveNotify_proto_rawDescData
}

var file_proto_EvtRushMoveNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_EvtRushMoveNotify_proto_goTypes = []interface{}{
	(*EvtRushMoveNotify)(nil), // 0: EvtRushMoveNotify
	(ForwardType)(0),          // 1: ForwardType
	(*EvtRushMoveInfo)(nil),   // 2: EvtRushMoveInfo
}
var file_proto_EvtRushMoveNotify_proto_depIdxs = []int32{
	1, // 0: EvtRushMoveNotify.forward_type:type_name -> ForwardType
	2, // 1: EvtRushMoveNotify.evt_rush_move_info:type_name -> EvtRushMoveInfo
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_EvtRushMoveNotify_proto_init() }
func file_proto_EvtRushMoveNotify_proto_init() {
	if File_proto_EvtRushMoveNotify_proto != nil {
		return
	}
	file_proto_ForwardType_proto_init()
	file_proto_EvtRushMoveInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_EvtRushMoveNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EvtRushMoveNotify); i {
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
			RawDescriptor: file_proto_EvtRushMoveNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_EvtRushMoveNotify_proto_goTypes,
		DependencyIndexes: file_proto_EvtRushMoveNotify_proto_depIdxs,
		MessageInfos:      file_proto_EvtRushMoveNotify_proto_msgTypes,
	}.Build()
	File_proto_EvtRushMoveNotify_proto = out.File
	file_proto_EvtRushMoveNotify_proto_rawDesc = nil
	file_proto_EvtRushMoveNotify_proto_goTypes = nil
	file_proto_EvtRushMoveNotify_proto_depIdxs = nil
}
