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
// source: proto/AvatarExpeditionInfo.proto

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

// Obf: MBBNGLKDKFD
type AvatarExpeditionInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	State        AvatarExpeditionState `protobuf:"varint,1,opt,name=state,proto3,enum=AvatarExpeditionState" json:"state,omitempty"`
	ExpId        uint32                `protobuf:"varint,2,opt,name=exp_id,json=expId,proto3" json:"exp_id,omitempty"`
	HourTime     uint32                `protobuf:"varint,3,opt,name=hour_time,json=hourTime,proto3" json:"hour_time,omitempty"`
	StartTime    uint32                `protobuf:"varint,4,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	ShortenRatio float32               `protobuf:"fixed32,5,opt,name=shorten_ratio,json=shortenRatio,proto3" json:"shorten_ratio,omitempty"`
}

func (x *AvatarExpeditionInfo) Reset() {
	*x = AvatarExpeditionInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_AvatarExpeditionInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AvatarExpeditionInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AvatarExpeditionInfo) ProtoMessage() {}

func (x *AvatarExpeditionInfo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_AvatarExpeditionInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AvatarExpeditionInfo.ProtoReflect.Descriptor instead.
func (*AvatarExpeditionInfo) Descriptor() ([]byte, []int) {
	return file_proto_AvatarExpeditionInfo_proto_rawDescGZIP(), []int{0}
}

func (x *AvatarExpeditionInfo) GetState() AvatarExpeditionState {
	if x != nil {
		return x.State
	}
	return AvatarExpeditionState_AVATAR_EXPEDITION_NONE
}

func (x *AvatarExpeditionInfo) GetExpId() uint32 {
	if x != nil {
		return x.ExpId
	}
	return 0
}

func (x *AvatarExpeditionInfo) GetHourTime() uint32 {
	if x != nil {
		return x.HourTime
	}
	return 0
}

func (x *AvatarExpeditionInfo) GetStartTime() uint32 {
	if x != nil {
		return x.StartTime
	}
	return 0
}

func (x *AvatarExpeditionInfo) GetShortenRatio() float32 {
	if x != nil {
		return x.ShortenRatio
	}
	return 0
}

var File_proto_AvatarExpeditionInfo_proto protoreflect.FileDescriptor

var file_proto_AvatarExpeditionInfo_proto_rawDesc = []byte{
	0x0a, 0x20, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x45, 0x78,
	0x70, 0x65, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x21, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72,
	0x45, 0x78, 0x70, 0x65, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbc, 0x01, 0x0a, 0x14, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72,
	0x45, 0x78, 0x70, 0x65, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x2c,
	0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e,
	0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x45, 0x78, 0x70, 0x65, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x15, 0x0a, 0x06,
	0x65, 0x78, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x65, 0x78,
	0x70, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x68, 0x6f, 0x75, 0x72, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x68, 0x6f, 0x75, 0x72, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x23, 0x0a, 0x0d, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x5f, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0c, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x52,
	0x61, 0x74, 0x69, 0x6f, 0x42, 0x20, 0x5a, 0x1e, 0x47, 0x6f, 0x2d, 0x47, 0x72, 0x61, 0x73, 0x73,
	0x63, 0x75, 0x74, 0x74, 0x65, 0x72, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_AvatarExpeditionInfo_proto_rawDescOnce sync.Once
	file_proto_AvatarExpeditionInfo_proto_rawDescData = file_proto_AvatarExpeditionInfo_proto_rawDesc
)

func file_proto_AvatarExpeditionInfo_proto_rawDescGZIP() []byte {
	file_proto_AvatarExpeditionInfo_proto_rawDescOnce.Do(func() {
		file_proto_AvatarExpeditionInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_AvatarExpeditionInfo_proto_rawDescData)
	})
	return file_proto_AvatarExpeditionInfo_proto_rawDescData
}

var file_proto_AvatarExpeditionInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_AvatarExpeditionInfo_proto_goTypes = []interface{}{
	(*AvatarExpeditionInfo)(nil), // 0: AvatarExpeditionInfo
	(AvatarExpeditionState)(0),   // 1: AvatarExpeditionState
}
var file_proto_AvatarExpeditionInfo_proto_depIdxs = []int32{
	1, // 0: AvatarExpeditionInfo.state:type_name -> AvatarExpeditionState
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_AvatarExpeditionInfo_proto_init() }
func file_proto_AvatarExpeditionInfo_proto_init() {
	if File_proto_AvatarExpeditionInfo_proto != nil {
		return
	}
	file_proto_AvatarExpeditionState_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_AvatarExpeditionInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AvatarExpeditionInfo); i {
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
			RawDescriptor: file_proto_AvatarExpeditionInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_AvatarExpeditionInfo_proto_goTypes,
		DependencyIndexes: file_proto_AvatarExpeditionInfo_proto_depIdxs,
		MessageInfos:      file_proto_AvatarExpeditionInfo_proto_msgTypes,
	}.Build()
	File_proto_AvatarExpeditionInfo_proto = out.File
	file_proto_AvatarExpeditionInfo_proto_rawDesc = nil
	file_proto_AvatarExpeditionInfo_proto_goTypes = nil
	file_proto_AvatarExpeditionInfo_proto_depIdxs = nil
}
