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
// source: proto/GCGControllerShowInfo.proto

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

// Obf: LKDLHNOBDFB
type GCGControllerShowInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NickName       string          `protobuf:"bytes,12,opt,name=nick_name,json=nickName,proto3" json:"nick_name,omitempty"`
	ProfilePicture *ProfilePicture `protobuf:"bytes,4,opt,name=profile_picture,json=profilePicture,proto3" json:"profile_picture,omitempty"`
	OnlineId       string          `protobuf:"bytes,15,opt,name=online_id,json=onlineId,proto3" json:"online_id,omitempty"`
	ControllerId   uint32          `protobuf:"varint,14,opt,name=controller_id,json=controllerId,proto3" json:"controller_id,omitempty"`
	PsnId          string          `protobuf:"bytes,3,opt,name=psn_id,json=psnId,proto3" json:"psn_id,omitempty"`
}

func (x *GCGControllerShowInfo) Reset() {
	*x = GCGControllerShowInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_GCGControllerShowInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GCGControllerShowInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GCGControllerShowInfo) ProtoMessage() {}

func (x *GCGControllerShowInfo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_GCGControllerShowInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GCGControllerShowInfo.ProtoReflect.Descriptor instead.
func (*GCGControllerShowInfo) Descriptor() ([]byte, []int) {
	return file_proto_GCGControllerShowInfo_proto_rawDescGZIP(), []int{0}
}

func (x *GCGControllerShowInfo) GetNickName() string {
	if x != nil {
		return x.NickName
	}
	return ""
}

func (x *GCGControllerShowInfo) GetProfilePicture() *ProfilePicture {
	if x != nil {
		return x.ProfilePicture
	}
	return nil
}

func (x *GCGControllerShowInfo) GetOnlineId() string {
	if x != nil {
		return x.OnlineId
	}
	return ""
}

func (x *GCGControllerShowInfo) GetControllerId() uint32 {
	if x != nil {
		return x.ControllerId
	}
	return 0
}

func (x *GCGControllerShowInfo) GetPsnId() string {
	if x != nil {
		return x.PsnId
	}
	return ""
}

var File_proto_GCGControllerShowInfo_proto protoreflect.FileDescriptor

var file_proto_GCGControllerShowInfo_proto_rawDesc = []byte{
	0x0a, 0x21, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x47, 0x43, 0x47, 0x43, 0x6f, 0x6e, 0x74, 0x72,
	0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x53, 0x68, 0x6f, 0x77, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x50, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x50, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xc7, 0x01, 0x0a, 0x15, 0x47, 0x43, 0x47, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65,
	0x72, 0x53, 0x68, 0x6f, 0x77, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x69, 0x63,
	0x6b, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69,
	0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x38, 0x0a, 0x0f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x5f, 0x70, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0f, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65,
	0x52, 0x0e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65,
	0x12, 0x1b, 0x0a, 0x09, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x0f, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x49, 0x64, 0x12, 0x23, 0x0a,
	0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x0e,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x70, 0x73, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x70, 0x73, 0x6e, 0x49, 0x64, 0x42, 0x1d, 0x5a, 0x1b, 0x47, 0x6f, 0x2d,
	0x47, 0x72, 0x61, 0x73, 0x73, 0x63, 0x75, 0x74, 0x74, 0x65, 0x72, 0x2f, 0x67, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_GCGControllerShowInfo_proto_rawDescOnce sync.Once
	file_proto_GCGControllerShowInfo_proto_rawDescData = file_proto_GCGControllerShowInfo_proto_rawDesc
)

func file_proto_GCGControllerShowInfo_proto_rawDescGZIP() []byte {
	file_proto_GCGControllerShowInfo_proto_rawDescOnce.Do(func() {
		file_proto_GCGControllerShowInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_GCGControllerShowInfo_proto_rawDescData)
	})
	return file_proto_GCGControllerShowInfo_proto_rawDescData
}

var file_proto_GCGControllerShowInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_GCGControllerShowInfo_proto_goTypes = []interface{}{
	(*GCGControllerShowInfo)(nil), // 0: GCGControllerShowInfo
	(*ProfilePicture)(nil),        // 1: ProfilePicture
}
var file_proto_GCGControllerShowInfo_proto_depIdxs = []int32{
	1, // 0: GCGControllerShowInfo.profile_picture:type_name -> ProfilePicture
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_GCGControllerShowInfo_proto_init() }
func file_proto_GCGControllerShowInfo_proto_init() {
	if File_proto_GCGControllerShowInfo_proto != nil {
		return
	}
	file_proto_ProfilePicture_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_GCGControllerShowInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GCGControllerShowInfo); i {
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
			RawDescriptor: file_proto_GCGControllerShowInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_GCGControllerShowInfo_proto_goTypes,
		DependencyIndexes: file_proto_GCGControllerShowInfo_proto_depIdxs,
		MessageInfos:      file_proto_GCGControllerShowInfo_proto_msgTypes,
	}.Build()
	File_proto_GCGControllerShowInfo_proto = out.File
	file_proto_GCGControllerShowInfo_proto_rawDesc = nil
	file_proto_GCGControllerShowInfo_proto_goTypes = nil
	file_proto_GCGControllerShowInfo_proto_depIdxs = nil
}
