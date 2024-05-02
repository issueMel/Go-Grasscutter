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
// source: proto/RegionSearchChangeRegionNotify.proto

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

// Obf: AOEBCMIOGBL
type RegionSearchChangeRegionNotify_RegionEvent int32

const (
	RegionSearchChangeRegionNotify_REGION_EVENT_NONE  RegionSearchChangeRegionNotify_RegionEvent = 0
	RegionSearchChangeRegionNotify_REGION_EVENT_ENTER RegionSearchChangeRegionNotify_RegionEvent = 1
	RegionSearchChangeRegionNotify_REGION_EVENT_LEAVE RegionSearchChangeRegionNotify_RegionEvent = 2
)

// Enum value maps for RegionSearchChangeRegionNotify_RegionEvent.
var (
	RegionSearchChangeRegionNotify_RegionEvent_name = map[int32]string{
		0: "REGION_EVENT_NONE",
		1: "REGION_EVENT_ENTER",
		2: "REGION_EVENT_LEAVE",
	}
	RegionSearchChangeRegionNotify_RegionEvent_value = map[string]int32{
		"REGION_EVENT_NONE":  0,
		"REGION_EVENT_ENTER": 1,
		"REGION_EVENT_LEAVE": 2,
	}
)

func (x RegionSearchChangeRegionNotify_RegionEvent) Enum() *RegionSearchChangeRegionNotify_RegionEvent {
	p := new(RegionSearchChangeRegionNotify_RegionEvent)
	*p = x
	return p
}

func (x RegionSearchChangeRegionNotify_RegionEvent) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RegionSearchChangeRegionNotify_RegionEvent) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_RegionSearchChangeRegionNotify_proto_enumTypes[0].Descriptor()
}

func (RegionSearchChangeRegionNotify_RegionEvent) Type() protoreflect.EnumType {
	return &file_proto_RegionSearchChangeRegionNotify_proto_enumTypes[0]
}

func (x RegionSearchChangeRegionNotify_RegionEvent) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RegionSearchChangeRegionNotify_RegionEvent.Descriptor instead.
func (RegionSearchChangeRegionNotify_RegionEvent) EnumDescriptor() ([]byte, []int) {
	return file_proto_RegionSearchChangeRegionNotify_proto_rawDescGZIP(), []int{0, 0}
}

// CmdId: 28838
// Obf: MFLFKDDNGCP
type RegionSearchChangeRegionNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Event    RegionSearchChangeRegionNotify_RegionEvent `protobuf:"varint,11,opt,name=event,proto3,enum=RegionSearchChangeRegionNotify_RegionEvent" json:"event,omitempty"`
	RegionId uint32                                     `protobuf:"varint,15,opt,name=region_id,json=regionId,proto3" json:"region_id,omitempty"`
}

func (x *RegionSearchChangeRegionNotify) Reset() {
	*x = RegionSearchChangeRegionNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_RegionSearchChangeRegionNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegionSearchChangeRegionNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegionSearchChangeRegionNotify) ProtoMessage() {}

func (x *RegionSearchChangeRegionNotify) ProtoReflect() protoreflect.Message {
	mi := &file_proto_RegionSearchChangeRegionNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegionSearchChangeRegionNotify.ProtoReflect.Descriptor instead.
func (*RegionSearchChangeRegionNotify) Descriptor() ([]byte, []int) {
	return file_proto_RegionSearchChangeRegionNotify_proto_rawDescGZIP(), []int{0}
}

func (x *RegionSearchChangeRegionNotify) GetEvent() RegionSearchChangeRegionNotify_RegionEvent {
	if x != nil {
		return x.Event
	}
	return RegionSearchChangeRegionNotify_REGION_EVENT_NONE
}

func (x *RegionSearchChangeRegionNotify) GetRegionId() uint32 {
	if x != nil {
		return x.RegionId
	}
	return 0
}

var File_proto_RegionSearchChangeRegionNotify_proto protoreflect.FileDescriptor

var file_proto_RegionSearchChangeRegionNotify_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x53, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd6, 0x01, 0x0a,
	0x1e, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x43, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12,
	0x41, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2b,
	0x2e, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x43, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e,
	0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x05, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18,
	0x0f, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22,
	0x54, 0x0a, 0x0b, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x15,
	0x0a, 0x11, 0x52, 0x45, 0x47, 0x49, 0x4f, 0x4e, 0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x4e,
	0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x16, 0x0a, 0x12, 0x52, 0x45, 0x47, 0x49, 0x4f, 0x4e, 0x5f,
	0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x45, 0x4e, 0x54, 0x45, 0x52, 0x10, 0x01, 0x12, 0x16, 0x0a,
	0x12, 0x52, 0x45, 0x47, 0x49, 0x4f, 0x4e, 0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x4c, 0x45,
	0x41, 0x56, 0x45, 0x10, 0x02, 0x42, 0x20, 0x5a, 0x1e, 0x47, 0x6f, 0x2d, 0x47, 0x72, 0x61, 0x73,
	0x73, 0x63, 0x75, 0x74, 0x74, 0x65, 0x72, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65,
	0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_RegionSearchChangeRegionNotify_proto_rawDescOnce sync.Once
	file_proto_RegionSearchChangeRegionNotify_proto_rawDescData = file_proto_RegionSearchChangeRegionNotify_proto_rawDesc
)

func file_proto_RegionSearchChangeRegionNotify_proto_rawDescGZIP() []byte {
	file_proto_RegionSearchChangeRegionNotify_proto_rawDescOnce.Do(func() {
		file_proto_RegionSearchChangeRegionNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_RegionSearchChangeRegionNotify_proto_rawDescData)
	})
	return file_proto_RegionSearchChangeRegionNotify_proto_rawDescData
}

var file_proto_RegionSearchChangeRegionNotify_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_RegionSearchChangeRegionNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_RegionSearchChangeRegionNotify_proto_goTypes = []interface{}{
	(RegionSearchChangeRegionNotify_RegionEvent)(0), // 0: RegionSearchChangeRegionNotify.RegionEvent
	(*RegionSearchChangeRegionNotify)(nil),          // 1: RegionSearchChangeRegionNotify
}
var file_proto_RegionSearchChangeRegionNotify_proto_depIdxs = []int32{
	0, // 0: RegionSearchChangeRegionNotify.event:type_name -> RegionSearchChangeRegionNotify.RegionEvent
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_RegionSearchChangeRegionNotify_proto_init() }
func file_proto_RegionSearchChangeRegionNotify_proto_init() {
	if File_proto_RegionSearchChangeRegionNotify_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_RegionSearchChangeRegionNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegionSearchChangeRegionNotify); i {
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
			RawDescriptor: file_proto_RegionSearchChangeRegionNotify_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_RegionSearchChangeRegionNotify_proto_goTypes,
		DependencyIndexes: file_proto_RegionSearchChangeRegionNotify_proto_depIdxs,
		EnumInfos:         file_proto_RegionSearchChangeRegionNotify_proto_enumTypes,
		MessageInfos:      file_proto_RegionSearchChangeRegionNotify_proto_msgTypes,
	}.Build()
	File_proto_RegionSearchChangeRegionNotify_proto = out.File
	file_proto_RegionSearchChangeRegionNotify_proto_rawDesc = nil
	file_proto_RegionSearchChangeRegionNotify_proto_goTypes = nil
	file_proto_RegionSearchChangeRegionNotify_proto_depIdxs = nil
}
