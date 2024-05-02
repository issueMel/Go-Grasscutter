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
// source: proto/HCOLEIFENKF.proto

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

// CmdId: 24858
type HCOLEIFENKF struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Guid            uint64         `protobuf:"varint,14,opt,name=guid,proto3" json:"guid,omitempty"`
	PABNPBHBCGH     *LPIFKAIBNHJ   `protobuf:"bytes,6,opt,name=PABNPBHBCGH,proto3" json:"PABNPBHBCGH,omitempty"`
	IsUpdateSetting bool           `protobuf:"varint,3,opt,name=is_update_setting,json=isUpdateSetting,proto3" json:"is_update_setting,omitempty"`
	RoomList        []*JOGGPMEEOEL `protobuf:"bytes,9,rep,name=room_list,json=roomList,proto3" json:"room_list,omitempty"`
	PFNPNODIOFE     bool           `protobuf:"varint,11,opt,name=PFNPNODIOFE,proto3" json:"PFNPNODIOFE,omitempty"`
}

func (x *HCOLEIFENKF) Reset() {
	*x = HCOLEIFENKF{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_HCOLEIFENKF_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HCOLEIFENKF) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HCOLEIFENKF) ProtoMessage() {}

func (x *HCOLEIFENKF) ProtoReflect() protoreflect.Message {
	mi := &file_proto_HCOLEIFENKF_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HCOLEIFENKF.ProtoReflect.Descriptor instead.
func (*HCOLEIFENKF) Descriptor() ([]byte, []int) {
	return file_proto_HCOLEIFENKF_proto_rawDescGZIP(), []int{0}
}

func (x *HCOLEIFENKF) GetGuid() uint64 {
	if x != nil {
		return x.Guid
	}
	return 0
}

func (x *HCOLEIFENKF) GetPABNPBHBCGH() *LPIFKAIBNHJ {
	if x != nil {
		return x.PABNPBHBCGH
	}
	return nil
}

func (x *HCOLEIFENKF) GetIsUpdateSetting() bool {
	if x != nil {
		return x.IsUpdateSetting
	}
	return false
}

func (x *HCOLEIFENKF) GetRoomList() []*JOGGPMEEOEL {
	if x != nil {
		return x.RoomList
	}
	return nil
}

func (x *HCOLEIFENKF) GetPFNPNODIOFE() bool {
	if x != nil {
		return x.PFNPNODIOFE
	}
	return false
}

var File_proto_HCOLEIFENKF_proto protoreflect.FileDescriptor

var file_proto_HCOLEIFENKF_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x48, 0x43, 0x4f, 0x4c, 0x45, 0x49, 0x46, 0x45,
	0x4e, 0x4b, 0x46, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x4c, 0x50, 0x49, 0x46, 0x4b, 0x41, 0x49, 0x42, 0x4e, 0x48, 0x4a, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x4a, 0x4f, 0x47, 0x47, 0x50, 0x4d,
	0x45, 0x45, 0x4f, 0x45, 0x4c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xca, 0x01, 0x0a, 0x0b,
	0x48, 0x43, 0x4f, 0x4c, 0x45, 0x49, 0x46, 0x45, 0x4e, 0x4b, 0x46, 0x12, 0x12, 0x0a, 0x04, 0x67,
	0x75, 0x69, 0x64, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x67, 0x75, 0x69, 0x64, 0x12,
	0x2e, 0x0a, 0x0b, 0x50, 0x41, 0x42, 0x4e, 0x50, 0x42, 0x48, 0x42, 0x43, 0x47, 0x48, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x4c, 0x50, 0x49, 0x46, 0x4b, 0x41, 0x49, 0x42, 0x4e,
	0x48, 0x4a, 0x52, 0x0b, 0x50, 0x41, 0x42, 0x4e, 0x50, 0x42, 0x48, 0x42, 0x43, 0x47, 0x48, 0x12,
	0x2a, 0x0a, 0x11, 0x69, 0x73, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x73, 0x65, 0x74,
	0x74, 0x69, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x69, 0x73, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x29, 0x0a, 0x09, 0x72,
	0x6f, 0x6f, 0x6d, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c,
	0x2e, 0x4a, 0x4f, 0x47, 0x47, 0x50, 0x4d, 0x45, 0x45, 0x4f, 0x45, 0x4c, 0x52, 0x08, 0x72, 0x6f,
	0x6f, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x50, 0x46, 0x4e, 0x50, 0x4e, 0x4f,
	0x44, 0x49, 0x4f, 0x46, 0x45, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x50, 0x46, 0x4e,
	0x50, 0x4e, 0x4f, 0x44, 0x49, 0x4f, 0x46, 0x45, 0x42, 0x1d, 0x5a, 0x1b, 0x47, 0x6f, 0x2d, 0x47,
	0x72, 0x61, 0x73, 0x73, 0x63, 0x75, 0x74, 0x74, 0x65, 0x72, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x74, 0x65, 0x64, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_HCOLEIFENKF_proto_rawDescOnce sync.Once
	file_proto_HCOLEIFENKF_proto_rawDescData = file_proto_HCOLEIFENKF_proto_rawDesc
)

func file_proto_HCOLEIFENKF_proto_rawDescGZIP() []byte {
	file_proto_HCOLEIFENKF_proto_rawDescOnce.Do(func() {
		file_proto_HCOLEIFENKF_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_HCOLEIFENKF_proto_rawDescData)
	})
	return file_proto_HCOLEIFENKF_proto_rawDescData
}

var file_proto_HCOLEIFENKF_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_HCOLEIFENKF_proto_goTypes = []interface{}{
	(*HCOLEIFENKF)(nil), // 0: HCOLEIFENKF
	(*LPIFKAIBNHJ)(nil), // 1: LPIFKAIBNHJ
	(*JOGGPMEEOEL)(nil), // 2: JOGGPMEEOEL
}
var file_proto_HCOLEIFENKF_proto_depIdxs = []int32{
	1, // 0: HCOLEIFENKF.PABNPBHBCGH:type_name -> LPIFKAIBNHJ
	2, // 1: HCOLEIFENKF.room_list:type_name -> JOGGPMEEOEL
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_HCOLEIFENKF_proto_init() }
func file_proto_HCOLEIFENKF_proto_init() {
	if File_proto_HCOLEIFENKF_proto != nil {
		return
	}
	file_proto_LPIFKAIBNHJ_proto_init()
	file_proto_JOGGPMEEOEL_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_HCOLEIFENKF_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HCOLEIFENKF); i {
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
			RawDescriptor: file_proto_HCOLEIFENKF_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_HCOLEIFENKF_proto_goTypes,
		DependencyIndexes: file_proto_HCOLEIFENKF_proto_depIdxs,
		MessageInfos:      file_proto_HCOLEIFENKF_proto_msgTypes,
	}.Build()
	File_proto_HCOLEIFENKF_proto = out.File
	file_proto_HCOLEIFENKF_proto_rawDesc = nil
	file_proto_HCOLEIFENKF_proto_goTypes = nil
	file_proto_HCOLEIFENKF_proto_depIdxs = nil
}
