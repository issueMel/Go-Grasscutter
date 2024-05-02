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
// source: proto/DealAddFriendReq.proto

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

// CmdId: 25212
// Obf: FAIFBJNFCAG
type DealAddFriendReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DealAddFriendResult DealAddFriendResultType `protobuf:"varint,11,opt,name=deal_add_friend_result,json=dealAddFriendResult,proto3,enum=DealAddFriendResultType" json:"deal_add_friend_result,omitempty"`
	TargetUid           uint32                  `protobuf:"varint,1,opt,name=target_uid,json=targetUid,proto3" json:"target_uid,omitempty"`
}

func (x *DealAddFriendReq) Reset() {
	*x = DealAddFriendReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_DealAddFriendReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DealAddFriendReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DealAddFriendReq) ProtoMessage() {}

func (x *DealAddFriendReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_DealAddFriendReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DealAddFriendReq.ProtoReflect.Descriptor instead.
func (*DealAddFriendReq) Descriptor() ([]byte, []int) {
	return file_proto_DealAddFriendReq_proto_rawDescGZIP(), []int{0}
}

func (x *DealAddFriendReq) GetDealAddFriendResult() DealAddFriendResultType {
	if x != nil {
		return x.DealAddFriendResult
	}
	return DealAddFriendResultType_DEAL_ADD_FRIEND_RESULT_TYPE_REJECT
}

func (x *DealAddFriendReq) GetTargetUid() uint32 {
	if x != nil {
		return x.TargetUid
	}
	return 0
}

var File_proto_DealAddFriendReq_proto protoreflect.FileDescriptor

var file_proto_DealAddFriendReq_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x44, 0x65, 0x61, 0x6c, 0x41, 0x64, 0x64, 0x46,
	0x72, 0x69, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x44, 0x65, 0x61, 0x6c, 0x41, 0x64, 0x64, 0x46, 0x72, 0x69,
	0x65, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x80, 0x01, 0x0a, 0x10, 0x44, 0x65, 0x61, 0x6c, 0x41, 0x64, 0x64, 0x46,
	0x72, 0x69, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x12, 0x4d, 0x0a, 0x16, 0x64, 0x65, 0x61, 0x6c,
	0x5f, 0x61, 0x64, 0x64, 0x5f, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x5f, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x44, 0x65, 0x61, 0x6c, 0x41,
	0x64, 0x64, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x13, 0x64, 0x65, 0x61, 0x6c, 0x41, 0x64, 0x64, 0x46, 0x72, 0x69, 0x65, 0x6e,
	0x64, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x5f, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x74, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x55, 0x69, 0x64, 0x42, 0x20, 0x5a, 0x1e, 0x47, 0x6f, 0x2d, 0x47, 0x72, 0x61,
	0x73, 0x73, 0x63, 0x75, 0x74, 0x74, 0x65, 0x72, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74,
	0x65, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_DealAddFriendReq_proto_rawDescOnce sync.Once
	file_proto_DealAddFriendReq_proto_rawDescData = file_proto_DealAddFriendReq_proto_rawDesc
)

func file_proto_DealAddFriendReq_proto_rawDescGZIP() []byte {
	file_proto_DealAddFriendReq_proto_rawDescOnce.Do(func() {
		file_proto_DealAddFriendReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_DealAddFriendReq_proto_rawDescData)
	})
	return file_proto_DealAddFriendReq_proto_rawDescData
}

var file_proto_DealAddFriendReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_DealAddFriendReq_proto_goTypes = []interface{}{
	(*DealAddFriendReq)(nil),     // 0: DealAddFriendReq
	(DealAddFriendResultType)(0), // 1: DealAddFriendResultType
}
var file_proto_DealAddFriendReq_proto_depIdxs = []int32{
	1, // 0: DealAddFriendReq.deal_add_friend_result:type_name -> DealAddFriendResultType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_DealAddFriendReq_proto_init() }
func file_proto_DealAddFriendReq_proto_init() {
	if File_proto_DealAddFriendReq_proto != nil {
		return
	}
	file_proto_DealAddFriendResultType_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_DealAddFriendReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DealAddFriendReq); i {
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
			RawDescriptor: file_proto_DealAddFriendReq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_DealAddFriendReq_proto_goTypes,
		DependencyIndexes: file_proto_DealAddFriendReq_proto_depIdxs,
		MessageInfos:      file_proto_DealAddFriendReq_proto_msgTypes,
	}.Build()
	File_proto_DealAddFriendReq_proto = out.File
	file_proto_DealAddFriendReq_proto_rawDesc = nil
	file_proto_DealAddFriendReq_proto_goTypes = nil
	file_proto_DealAddFriendReq_proto_depIdxs = nil
}
