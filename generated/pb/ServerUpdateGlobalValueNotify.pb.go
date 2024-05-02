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
// source: proto/ServerUpdateGlobalValueNotify.proto

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

// Obf: AHKBKOMMMJD
type ServerUpdateGlobalValueNotify_UpdateType int32

const (
	ServerUpdateGlobalValueNotify_INVALUE ServerUpdateGlobalValueNotify_UpdateType = 0
	ServerUpdateGlobalValueNotify_ADD     ServerUpdateGlobalValueNotify_UpdateType = 1
	ServerUpdateGlobalValueNotify_SET     ServerUpdateGlobalValueNotify_UpdateType = 2
)

// Enum value maps for ServerUpdateGlobalValueNotify_UpdateType.
var (
	ServerUpdateGlobalValueNotify_UpdateType_name = map[int32]string{
		0: "INVALUE",
		1: "ADD",
		2: "SET",
	}
	ServerUpdateGlobalValueNotify_UpdateType_value = map[string]int32{
		"INVALUE": 0,
		"ADD":     1,
		"SET":     2,
	}
)

func (x ServerUpdateGlobalValueNotify_UpdateType) Enum() *ServerUpdateGlobalValueNotify_UpdateType {
	p := new(ServerUpdateGlobalValueNotify_UpdateType)
	*p = x
	return p
}

func (x ServerUpdateGlobalValueNotify_UpdateType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ServerUpdateGlobalValueNotify_UpdateType) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_ServerUpdateGlobalValueNotify_proto_enumTypes[0].Descriptor()
}

func (ServerUpdateGlobalValueNotify_UpdateType) Type() protoreflect.EnumType {
	return &file_proto_ServerUpdateGlobalValueNotify_proto_enumTypes[0]
}

func (x ServerUpdateGlobalValueNotify_UpdateType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ServerUpdateGlobalValueNotify_UpdateType.Descriptor instead.
func (ServerUpdateGlobalValueNotify_UpdateType) EnumDescriptor() ([]byte, []int) {
	return file_proto_ServerUpdateGlobalValueNotify_proto_rawDescGZIP(), []int{0, 0}
}

// CmdId: 1181
// Obf: BKDHPPNCFBO
type ServerUpdateGlobalValueNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Delta      float32                                  `protobuf:"fixed32,2,opt,name=delta,proto3" json:"delta,omitempty"`
	EntityId   uint32                                   `protobuf:"varint,8,opt,name=entity_id,json=entityId,proto3" json:"entity_id,omitempty"`
	KeyHash    uint32                                   `protobuf:"varint,5,opt,name=key_hash,json=keyHash,proto3" json:"key_hash,omitempty"`
	UpdateType ServerUpdateGlobalValueNotify_UpdateType `protobuf:"varint,3,opt,name=update_type,json=updateType,proto3,enum=ServerUpdateGlobalValueNotify_UpdateType" json:"update_type,omitempty"`
	Value      float32                                  `protobuf:"fixed32,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *ServerUpdateGlobalValueNotify) Reset() {
	*x = ServerUpdateGlobalValueNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ServerUpdateGlobalValueNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerUpdateGlobalValueNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerUpdateGlobalValueNotify) ProtoMessage() {}

func (x *ServerUpdateGlobalValueNotify) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ServerUpdateGlobalValueNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerUpdateGlobalValueNotify.ProtoReflect.Descriptor instead.
func (*ServerUpdateGlobalValueNotify) Descriptor() ([]byte, []int) {
	return file_proto_ServerUpdateGlobalValueNotify_proto_rawDescGZIP(), []int{0}
}

func (x *ServerUpdateGlobalValueNotify) GetDelta() float32 {
	if x != nil {
		return x.Delta
	}
	return 0
}

func (x *ServerUpdateGlobalValueNotify) GetEntityId() uint32 {
	if x != nil {
		return x.EntityId
	}
	return 0
}

func (x *ServerUpdateGlobalValueNotify) GetKeyHash() uint32 {
	if x != nil {
		return x.KeyHash
	}
	return 0
}

func (x *ServerUpdateGlobalValueNotify) GetUpdateType() ServerUpdateGlobalValueNotify_UpdateType {
	if x != nil {
		return x.UpdateType
	}
	return ServerUpdateGlobalValueNotify_INVALUE
}

func (x *ServerUpdateGlobalValueNotify) GetValue() float32 {
	if x != nil {
		return x.Value
	}
	return 0
}

var File_proto_ServerUpdateGlobalValueNotify_proto protoreflect.FileDescriptor

var file_proto_ServerUpdateGlobalValueNotify_proto_rawDesc = []byte{
	0x0a, 0x29, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x4e,
	0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfc, 0x01, 0x0a, 0x1d,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x6c, 0x6f, 0x62,
	0x61, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x64, 0x65, 0x6c, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x64, 0x65,
	0x6c, 0x74, 0x61, 0x12, 0x1b, 0x0a, 0x09, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x64,
	0x12, 0x19, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x07, 0x6b, 0x65, 0x79, 0x48, 0x61, 0x73, 0x68, 0x12, 0x4a, 0x0a, 0x0b, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x29, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47,
	0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x2b, 0x0a,
	0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x49,
	0x4e, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x41, 0x44, 0x44, 0x10,
	0x01, 0x12, 0x07, 0x0a, 0x03, 0x53, 0x45, 0x54, 0x10, 0x02, 0x42, 0x1d, 0x5a, 0x1b, 0x47, 0x6f,
	0x2d, 0x47, 0x72, 0x61, 0x73, 0x73, 0x63, 0x75, 0x74, 0x74, 0x65, 0x72, 0x2f, 0x67, 0x65, 0x6e,
	0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_proto_ServerUpdateGlobalValueNotify_proto_rawDescOnce sync.Once
	file_proto_ServerUpdateGlobalValueNotify_proto_rawDescData = file_proto_ServerUpdateGlobalValueNotify_proto_rawDesc
)

func file_proto_ServerUpdateGlobalValueNotify_proto_rawDescGZIP() []byte {
	file_proto_ServerUpdateGlobalValueNotify_proto_rawDescOnce.Do(func() {
		file_proto_ServerUpdateGlobalValueNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_ServerUpdateGlobalValueNotify_proto_rawDescData)
	})
	return file_proto_ServerUpdateGlobalValueNotify_proto_rawDescData
}

var file_proto_ServerUpdateGlobalValueNotify_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_ServerUpdateGlobalValueNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_ServerUpdateGlobalValueNotify_proto_goTypes = []interface{}{
	(ServerUpdateGlobalValueNotify_UpdateType)(0), // 0: ServerUpdateGlobalValueNotify.UpdateType
	(*ServerUpdateGlobalValueNotify)(nil),         // 1: ServerUpdateGlobalValueNotify
}
var file_proto_ServerUpdateGlobalValueNotify_proto_depIdxs = []int32{
	0, // 0: ServerUpdateGlobalValueNotify.update_type:type_name -> ServerUpdateGlobalValueNotify.UpdateType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_ServerUpdateGlobalValueNotify_proto_init() }
func file_proto_ServerUpdateGlobalValueNotify_proto_init() {
	if File_proto_ServerUpdateGlobalValueNotify_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_ServerUpdateGlobalValueNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerUpdateGlobalValueNotify); i {
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
			RawDescriptor: file_proto_ServerUpdateGlobalValueNotify_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_ServerUpdateGlobalValueNotify_proto_goTypes,
		DependencyIndexes: file_proto_ServerUpdateGlobalValueNotify_proto_depIdxs,
		EnumInfos:         file_proto_ServerUpdateGlobalValueNotify_proto_enumTypes,
		MessageInfos:      file_proto_ServerUpdateGlobalValueNotify_proto_msgTypes,
	}.Build()
	File_proto_ServerUpdateGlobalValueNotify_proto = out.File
	file_proto_ServerUpdateGlobalValueNotify_proto_rawDesc = nil
	file_proto_ServerUpdateGlobalValueNotify_proto_goTypes = nil
	file_proto_ServerUpdateGlobalValueNotify_proto_depIdxs = nil
}