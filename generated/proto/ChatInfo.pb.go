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
// source: proto/ChatInfo.proto

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

// Obf: KFPJAAEHBNB
type ChatInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ToUid uint32 `protobuf:"varint,10,opt,name=to_uid,json=toUid,proto3" json:"to_uid,omitempty"`
	Time  uint32 `protobuf:"varint,13,opt,name=time,proto3" json:"time,omitempty"`
	Uid   uint32 `protobuf:"varint,4,opt,name=uid,proto3" json:"uid,omitempty"`
	// Types that are assignable to Content:
	//
	//	*ChatInfo_Text
	//	*ChatInfo_Icon
	//	*ChatInfo_SystemHint_
	Content  isChatInfo_Content `protobuf_oneof:"content"`
	IsRead   bool               `protobuf:"varint,7,opt,name=is_read,json=isRead,proto3" json:"is_read,omitempty"`
	Sequence uint32             `protobuf:"varint,5,opt,name=sequence,proto3" json:"sequence,omitempty"`
}

func (x *ChatInfo) Reset() {
	*x = ChatInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ChatInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatInfo) ProtoMessage() {}

func (x *ChatInfo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ChatInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatInfo.ProtoReflect.Descriptor instead.
func (*ChatInfo) Descriptor() ([]byte, []int) {
	return file_proto_ChatInfo_proto_rawDescGZIP(), []int{0}
}

func (x *ChatInfo) GetToUid() uint32 {
	if x != nil {
		return x.ToUid
	}
	return 0
}

func (x *ChatInfo) GetTime() uint32 {
	if x != nil {
		return x.Time
	}
	return 0
}

func (x *ChatInfo) GetUid() uint32 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (m *ChatInfo) GetContent() isChatInfo_Content {
	if m != nil {
		return m.Content
	}
	return nil
}

func (x *ChatInfo) GetText() string {
	if x, ok := x.GetContent().(*ChatInfo_Text); ok {
		return x.Text
	}
	return ""
}

func (x *ChatInfo) GetIcon() uint32 {
	if x, ok := x.GetContent().(*ChatInfo_Icon); ok {
		return x.Icon
	}
	return 0
}

func (x *ChatInfo) GetSystemHint() *ChatInfo_SystemHint {
	if x, ok := x.GetContent().(*ChatInfo_SystemHint_); ok {
		return x.SystemHint
	}
	return nil
}

func (x *ChatInfo) GetIsRead() bool {
	if x != nil {
		return x.IsRead
	}
	return false
}

func (x *ChatInfo) GetSequence() uint32 {
	if x != nil {
		return x.Sequence
	}
	return 0
}

type isChatInfo_Content interface {
	isChatInfo_Content()
}

type ChatInfo_Text struct {
	Text string `protobuf:"bytes,962,opt,name=text,proto3,oneof"`
}

type ChatInfo_Icon struct {
	Icon uint32 `protobuf:"varint,1562,opt,name=icon,proto3,oneof"`
}

type ChatInfo_SystemHint_ struct {
	SystemHint *ChatInfo_SystemHint `protobuf:"bytes,841,opt,name=system_hint,json=systemHint,proto3,oneof"`
}

func (*ChatInfo_Text) isChatInfo_Content() {}

func (*ChatInfo_Icon) isChatInfo_Content() {}

func (*ChatInfo_SystemHint_) isChatInfo_Content() {}

// Obf: LGDLFPKODMA
type ChatInfo_SystemHint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type uint32 `protobuf:"varint,15,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *ChatInfo_SystemHint) Reset() {
	*x = ChatInfo_SystemHint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ChatInfo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatInfo_SystemHint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatInfo_SystemHint) ProtoMessage() {}

func (x *ChatInfo_SystemHint) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ChatInfo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatInfo_SystemHint.ProtoReflect.Descriptor instead.
func (*ChatInfo_SystemHint) Descriptor() ([]byte, []int) {
	return file_proto_ChatInfo_proto_rawDescGZIP(), []int{0, 0}
}

func (x *ChatInfo_SystemHint) GetType() uint32 {
	if x != nil {
		return x.Type
	}
	return 0
}

var File_proto_ChatInfo_proto protoreflect.FileDescriptor

var file_proto_ChatInfo_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x43, 0x68, 0x61, 0x74, 0x49, 0x6e, 0x66, 0x6f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x91, 0x02, 0x0a, 0x08, 0x43, 0x68, 0x61, 0x74, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x15, 0x0a, 0x06, 0x74, 0x6f, 0x5f, 0x75, 0x69, 0x64, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x05, 0x74, 0x6f, 0x55, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x75, 0x69, 0x64,
	0x12, 0x15, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0xc2, 0x07, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x00, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x15, 0x0a, 0x04, 0x69, 0x63, 0x6f, 0x6e, 0x18,
	0x9a, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x00, 0x52, 0x04, 0x69, 0x63, 0x6f, 0x6e, 0x12, 0x38,
	0x0a, 0x0b, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x5f, 0x68, 0x69, 0x6e, 0x74, 0x18, 0xc9, 0x06,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x2e,
	0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x48, 0x69, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x0a, 0x73, 0x79,
	0x73, 0x74, 0x65, 0x6d, 0x48, 0x69, 0x6e, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x73, 0x5f, 0x72,
	0x65, 0x61, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x69, 0x73, 0x52, 0x65, 0x61,
	0x64, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x08, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x1a, 0x20, 0x0a,
	0x0a, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x48, 0x69, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x42,
	0x09, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x42, 0x20, 0x5a, 0x1e, 0x47, 0x6f,
	0x2d, 0x47, 0x72, 0x61, 0x73, 0x73, 0x63, 0x75, 0x74, 0x74, 0x65, 0x72, 0x2f, 0x67, 0x65, 0x6e,
	0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_ChatInfo_proto_rawDescOnce sync.Once
	file_proto_ChatInfo_proto_rawDescData = file_proto_ChatInfo_proto_rawDesc
)

func file_proto_ChatInfo_proto_rawDescGZIP() []byte {
	file_proto_ChatInfo_proto_rawDescOnce.Do(func() {
		file_proto_ChatInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_ChatInfo_proto_rawDescData)
	})
	return file_proto_ChatInfo_proto_rawDescData
}

var file_proto_ChatInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_ChatInfo_proto_goTypes = []interface{}{
	(*ChatInfo)(nil),            // 0: ChatInfo
	(*ChatInfo_SystemHint)(nil), // 1: ChatInfo.SystemHint
}
var file_proto_ChatInfo_proto_depIdxs = []int32{
	1, // 0: ChatInfo.system_hint:type_name -> ChatInfo.SystemHint
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_ChatInfo_proto_init() }
func file_proto_ChatInfo_proto_init() {
	if File_proto_ChatInfo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_ChatInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatInfo); i {
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
		file_proto_ChatInfo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatInfo_SystemHint); i {
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
	file_proto_ChatInfo_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*ChatInfo_Text)(nil),
		(*ChatInfo_Icon)(nil),
		(*ChatInfo_SystemHint_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_ChatInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_ChatInfo_proto_goTypes,
		DependencyIndexes: file_proto_ChatInfo_proto_depIdxs,
		MessageInfos:      file_proto_ChatInfo_proto_msgTypes,
	}.Build()
	File_proto_ChatInfo_proto = out.File
	file_proto_ChatInfo_proto_rawDesc = nil
	file_proto_ChatInfo_proto_goTypes = nil
	file_proto_ChatInfo_proto_depIdxs = nil
}
