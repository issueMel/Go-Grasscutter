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
// source: proto/DraftGuestReplyInviteRsp.proto

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

// CmdId: 5277
// Obf: LAGBECGIPIP
type DraftGuestReplyInviteRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DraftId uint32 `protobuf:"varint,14,opt,name=draft_id,json=draftId,proto3" json:"draft_id,omitempty"`
	IsAgree bool   `protobuf:"varint,3,opt,name=is_agree,json=isAgree,proto3" json:"is_agree,omitempty"`
	Retcode int32  `protobuf:"varint,2,opt,name=retcode,proto3" json:"retcode,omitempty"`
}

func (x *DraftGuestReplyInviteRsp) Reset() {
	*x = DraftGuestReplyInviteRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_DraftGuestReplyInviteRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DraftGuestReplyInviteRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DraftGuestReplyInviteRsp) ProtoMessage() {}

func (x *DraftGuestReplyInviteRsp) ProtoReflect() protoreflect.Message {
	mi := &file_proto_DraftGuestReplyInviteRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DraftGuestReplyInviteRsp.ProtoReflect.Descriptor instead.
func (*DraftGuestReplyInviteRsp) Descriptor() ([]byte, []int) {
	return file_proto_DraftGuestReplyInviteRsp_proto_rawDescGZIP(), []int{0}
}

func (x *DraftGuestReplyInviteRsp) GetDraftId() uint32 {
	if x != nil {
		return x.DraftId
	}
	return 0
}

func (x *DraftGuestReplyInviteRsp) GetIsAgree() bool {
	if x != nil {
		return x.IsAgree
	}
	return false
}

func (x *DraftGuestReplyInviteRsp) GetRetcode() int32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

var File_proto_DraftGuestReplyInviteRsp_proto protoreflect.FileDescriptor

var file_proto_DraftGuestReplyInviteRsp_proto_rawDesc = []byte{
	0x0a, 0x24, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x44, 0x72, 0x61, 0x66, 0x74, 0x47, 0x75, 0x65,
	0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x52, 0x73, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6a, 0x0a, 0x18, 0x44, 0x72, 0x61, 0x66, 0x74, 0x47,
	0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x52,
	0x73, 0x70, 0x12, 0x19, 0x0a, 0x08, 0x64, 0x72, 0x61, 0x66, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x0e,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x64, 0x72, 0x61, 0x66, 0x74, 0x49, 0x64, 0x12, 0x19, 0x0a,
	0x08, 0x69, 0x73, 0x5f, 0x61, 0x67, 0x72, 0x65, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x07, 0x69, 0x73, 0x41, 0x67, 0x72, 0x65, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f,
	0x64, 0x65, 0x42, 0x20, 0x5a, 0x1e, 0x47, 0x6f, 0x2d, 0x47, 0x72, 0x61, 0x73, 0x73, 0x63, 0x75,
	0x74, 0x74, 0x65, 0x72, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_DraftGuestReplyInviteRsp_proto_rawDescOnce sync.Once
	file_proto_DraftGuestReplyInviteRsp_proto_rawDescData = file_proto_DraftGuestReplyInviteRsp_proto_rawDesc
)

func file_proto_DraftGuestReplyInviteRsp_proto_rawDescGZIP() []byte {
	file_proto_DraftGuestReplyInviteRsp_proto_rawDescOnce.Do(func() {
		file_proto_DraftGuestReplyInviteRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_DraftGuestReplyInviteRsp_proto_rawDescData)
	})
	return file_proto_DraftGuestReplyInviteRsp_proto_rawDescData
}

var file_proto_DraftGuestReplyInviteRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_DraftGuestReplyInviteRsp_proto_goTypes = []interface{}{
	(*DraftGuestReplyInviteRsp)(nil), // 0: DraftGuestReplyInviteRsp
}
var file_proto_DraftGuestReplyInviteRsp_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_DraftGuestReplyInviteRsp_proto_init() }
func file_proto_DraftGuestReplyInviteRsp_proto_init() {
	if File_proto_DraftGuestReplyInviteRsp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_DraftGuestReplyInviteRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DraftGuestReplyInviteRsp); i {
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
			RawDescriptor: file_proto_DraftGuestReplyInviteRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_DraftGuestReplyInviteRsp_proto_goTypes,
		DependencyIndexes: file_proto_DraftGuestReplyInviteRsp_proto_depIdxs,
		MessageInfos:      file_proto_DraftGuestReplyInviteRsp_proto_msgTypes,
	}.Build()
	File_proto_DraftGuestReplyInviteRsp_proto = out.File
	file_proto_DraftGuestReplyInviteRsp_proto_rawDesc = nil
	file_proto_DraftGuestReplyInviteRsp_proto_goTypes = nil
	file_proto_DraftGuestReplyInviteRsp_proto_depIdxs = nil
}
