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
// source: proto/AbilityInvokeEntry.proto

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

// Obf: BNCFAGLEKAM
type AbilityInvokeEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ArgumentType  AbilityInvokeArgument   `protobuf:"varint,7,opt,name=argument_type,json=argumentType,proto3,enum=AbilityInvokeArgument" json:"argument_type,omitempty"`
	AbilityData   []byte                  `protobuf:"bytes,2,opt,name=ability_data,json=abilityData,proto3" json:"ability_data,omitempty"`
	EntityId      uint32                  `protobuf:"varint,9,opt,name=entity_id,json=entityId,proto3" json:"entity_id,omitempty"`
	ForwardType   ForwardType             `protobuf:"varint,14,opt,name=forward_type,json=forwardType,proto3,enum=ForwardType" json:"forward_type,omitempty"`
	TotalTickTime float64                 `protobuf:"fixed64,5,opt,name=total_tick_time,json=totalTickTime,proto3" json:"total_tick_time,omitempty"`
	ForwardPeer   uint32                  `protobuf:"varint,12,opt,name=forward_peer,json=forwardPeer,proto3" json:"forward_peer,omitempty"`
	EventId       uint32                  `protobuf:"varint,1,opt,name=event_id,json=eventId,proto3" json:"event_id,omitempty"`
	IsIgnoreAuth  bool                    `protobuf:"varint,13,opt,name=is_ignore_auth,json=isIgnoreAuth,proto3" json:"is_ignore_auth,omitempty"`
	Head          *AbilityInvokeEntryHead `protobuf:"bytes,8,opt,name=head,proto3" json:"head,omitempty"`
}

func (x *AbilityInvokeEntry) Reset() {
	*x = AbilityInvokeEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_AbilityInvokeEntry_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AbilityInvokeEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AbilityInvokeEntry) ProtoMessage() {}

func (x *AbilityInvokeEntry) ProtoReflect() protoreflect.Message {
	mi := &file_proto_AbilityInvokeEntry_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AbilityInvokeEntry.ProtoReflect.Descriptor instead.
func (*AbilityInvokeEntry) Descriptor() ([]byte, []int) {
	return file_proto_AbilityInvokeEntry_proto_rawDescGZIP(), []int{0}
}

func (x *AbilityInvokeEntry) GetArgumentType() AbilityInvokeArgument {
	if x != nil {
		return x.ArgumentType
	}
	return AbilityInvokeArgument_ABILITY_INVOKE_ARGUMENT_NONE
}

func (x *AbilityInvokeEntry) GetAbilityData() []byte {
	if x != nil {
		return x.AbilityData
	}
	return nil
}

func (x *AbilityInvokeEntry) GetEntityId() uint32 {
	if x != nil {
		return x.EntityId
	}
	return 0
}

func (x *AbilityInvokeEntry) GetForwardType() ForwardType {
	if x != nil {
		return x.ForwardType
	}
	return ForwardType_FORWARD_TYPE_LOCAL
}

func (x *AbilityInvokeEntry) GetTotalTickTime() float64 {
	if x != nil {
		return x.TotalTickTime
	}
	return 0
}

func (x *AbilityInvokeEntry) GetForwardPeer() uint32 {
	if x != nil {
		return x.ForwardPeer
	}
	return 0
}

func (x *AbilityInvokeEntry) GetEventId() uint32 {
	if x != nil {
		return x.EventId
	}
	return 0
}

func (x *AbilityInvokeEntry) GetIsIgnoreAuth() bool {
	if x != nil {
		return x.IsIgnoreAuth
	}
	return false
}

func (x *AbilityInvokeEntry) GetHead() *AbilityInvokeEntryHead {
	if x != nil {
		return x.Head
	}
	return nil
}

var File_proto_AbilityInvokeEntry_proto protoreflect.FileDescriptor

var file_proto_AbilityInvokeEntry_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x41, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x49,
	0x6e, 0x76, 0x6f, 0x6b, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x21, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x41, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x49,
	0x6e, 0x76, 0x6f, 0x6b, 0x65, 0x41, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x46, 0x6f, 0x72, 0x77, 0x61,
	0x72, 0x64, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x41, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x49, 0x6e, 0x76, 0x6f, 0x6b,
	0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x48, 0x65, 0x61, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xfb, 0x02, 0x0a, 0x12, 0x41, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x49, 0x6e, 0x76, 0x6f,
	0x6b, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x3b, 0x0a, 0x0d, 0x61, 0x72, 0x67, 0x75, 0x6d,
	0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16,
	0x2e, 0x41, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x49, 0x6e, 0x76, 0x6f, 0x6b, 0x65, 0x41, 0x72,
	0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0c, 0x61, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x5f,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b, 0x61, 0x62, 0x69, 0x6c,
	0x69, 0x74, 0x79, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1b, 0x0a, 0x09, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x49, 0x64, 0x12, 0x2f, 0x0a, 0x0c, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0c, 0x2e, 0x46, 0x6f, 0x72,
	0x77, 0x61, 0x72, 0x64, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0b, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72,
	0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x26, 0x0a, 0x0f, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x74,
	0x69, 0x63, 0x6b, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0d,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x54, 0x69, 0x63, 0x6b, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x21, 0x0a,
	0x0c, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x70, 0x65, 0x65, 0x72, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0b, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x50, 0x65, 0x65, 0x72,
	0x12, 0x19, 0x0a, 0x08, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x07, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0e, 0x69,
	0x73, 0x5f, 0x69, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x18, 0x0d, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0c, 0x69, 0x73, 0x49, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x41, 0x75, 0x74,
	0x68, 0x12, 0x2b, 0x0a, 0x04, 0x68, 0x65, 0x61, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x17, 0x2e, 0x41, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x49, 0x6e, 0x76, 0x6f, 0x6b, 0x65, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x48, 0x65, 0x61, 0x64, 0x52, 0x04, 0x68, 0x65, 0x61, 0x64, 0x42, 0x20,
	0x5a, 0x1e, 0x47, 0x6f, 0x2d, 0x47, 0x72, 0x61, 0x73, 0x73, 0x63, 0x75, 0x74, 0x74, 0x65, 0x72,
	0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_AbilityInvokeEntry_proto_rawDescOnce sync.Once
	file_proto_AbilityInvokeEntry_proto_rawDescData = file_proto_AbilityInvokeEntry_proto_rawDesc
)

func file_proto_AbilityInvokeEntry_proto_rawDescGZIP() []byte {
	file_proto_AbilityInvokeEntry_proto_rawDescOnce.Do(func() {
		file_proto_AbilityInvokeEntry_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_AbilityInvokeEntry_proto_rawDescData)
	})
	return file_proto_AbilityInvokeEntry_proto_rawDescData
}

var file_proto_AbilityInvokeEntry_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_AbilityInvokeEntry_proto_goTypes = []interface{}{
	(*AbilityInvokeEntry)(nil),     // 0: AbilityInvokeEntry
	(AbilityInvokeArgument)(0),     // 1: AbilityInvokeArgument
	(ForwardType)(0),               // 2: ForwardType
	(*AbilityInvokeEntryHead)(nil), // 3: AbilityInvokeEntryHead
}
var file_proto_AbilityInvokeEntry_proto_depIdxs = []int32{
	1, // 0: AbilityInvokeEntry.argument_type:type_name -> AbilityInvokeArgument
	2, // 1: AbilityInvokeEntry.forward_type:type_name -> ForwardType
	3, // 2: AbilityInvokeEntry.head:type_name -> AbilityInvokeEntryHead
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_proto_AbilityInvokeEntry_proto_init() }
func file_proto_AbilityInvokeEntry_proto_init() {
	if File_proto_AbilityInvokeEntry_proto != nil {
		return
	}
	file_proto_AbilityInvokeArgument_proto_init()
	file_proto_ForwardType_proto_init()
	file_proto_AbilityInvokeEntryHead_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_AbilityInvokeEntry_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AbilityInvokeEntry); i {
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
			RawDescriptor: file_proto_AbilityInvokeEntry_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_AbilityInvokeEntry_proto_goTypes,
		DependencyIndexes: file_proto_AbilityInvokeEntry_proto_depIdxs,
		MessageInfos:      file_proto_AbilityInvokeEntry_proto_msgTypes,
	}.Build()
	File_proto_AbilityInvokeEntry_proto = out.File
	file_proto_AbilityInvokeEntry_proto_rawDesc = nil
	file_proto_AbilityInvokeEntry_proto_goTypes = nil
	file_proto_AbilityInvokeEntry_proto_depIdxs = nil
}
