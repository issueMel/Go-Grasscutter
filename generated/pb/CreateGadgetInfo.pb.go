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
// source: proto/CreateGadgetInfo.proto

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

// Obf: NDJJNLJIAKO
type CreateGadgetInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BornType GadgetBornType          `protobuf:"varint,1,opt,name=born_type,json=bornType,proto3,enum=GadgetBornType" json:"born_type,omitempty"`
	Chest    *CreateGadgetInfo_Chest `protobuf:"bytes,2,opt,name=chest,proto3" json:"chest,omitempty"`
}

func (x *CreateGadgetInfo) Reset() {
	*x = CreateGadgetInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_CreateGadgetInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateGadgetInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateGadgetInfo) ProtoMessage() {}

func (x *CreateGadgetInfo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_CreateGadgetInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateGadgetInfo.ProtoReflect.Descriptor instead.
func (*CreateGadgetInfo) Descriptor() ([]byte, []int) {
	return file_proto_CreateGadgetInfo_proto_rawDescGZIP(), []int{0}
}

func (x *CreateGadgetInfo) GetBornType() GadgetBornType {
	if x != nil {
		return x.BornType
	}
	return GadgetBornType_GADGET_BORN_TYPE_NONE
}

func (x *CreateGadgetInfo) GetChest() *CreateGadgetInfo_Chest {
	if x != nil {
		return x.Chest
	}
	return nil
}

// Obf: FGAGBNDOGCB
type CreateGadgetInfo_Chest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChestDropId    uint32 `protobuf:"varint,1,opt,name=chest_drop_id,json=chestDropId,proto3" json:"chest_drop_id,omitempty"`
	IsShowCutscene bool   `protobuf:"varint,2,opt,name=is_show_cutscene,json=isShowCutscene,proto3" json:"is_show_cutscene,omitempty"`
}

func (x *CreateGadgetInfo_Chest) Reset() {
	*x = CreateGadgetInfo_Chest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_CreateGadgetInfo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateGadgetInfo_Chest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateGadgetInfo_Chest) ProtoMessage() {}

func (x *CreateGadgetInfo_Chest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_CreateGadgetInfo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateGadgetInfo_Chest.ProtoReflect.Descriptor instead.
func (*CreateGadgetInfo_Chest) Descriptor() ([]byte, []int) {
	return file_proto_CreateGadgetInfo_proto_rawDescGZIP(), []int{0, 0}
}

func (x *CreateGadgetInfo_Chest) GetChestDropId() uint32 {
	if x != nil {
		return x.ChestDropId
	}
	return 0
}

func (x *CreateGadgetInfo_Chest) GetIsShowCutscene() bool {
	if x != nil {
		return x.IsShowCutscene
	}
	return false
}

var File_proto_CreateGadgetInfo_proto protoreflect.FileDescriptor

var file_proto_CreateGadgetInfo_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x61,
	0x64, 0x67, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x47, 0x61, 0x64, 0x67, 0x65, 0x74, 0x42, 0x6f, 0x72, 0x6e,
	0x54, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc6, 0x01, 0x0a, 0x10, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x61, 0x64, 0x67, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x2c, 0x0a, 0x09, 0x62, 0x6f, 0x72, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x47, 0x61, 0x64, 0x67, 0x65, 0x74, 0x42, 0x6f, 0x72, 0x6e, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x08, 0x62, 0x6f, 0x72, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2d, 0x0a,
	0x05, 0x63, 0x68, 0x65, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x61, 0x64, 0x67, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x2e,
	0x43, 0x68, 0x65, 0x73, 0x74, 0x52, 0x05, 0x63, 0x68, 0x65, 0x73, 0x74, 0x1a, 0x55, 0x0a, 0x05,
	0x43, 0x68, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x0d, 0x63, 0x68, 0x65, 0x73, 0x74, 0x5f, 0x64,
	0x72, 0x6f, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x63, 0x68,
	0x65, 0x73, 0x74, 0x44, 0x72, 0x6f, 0x70, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x10, 0x69, 0x73, 0x5f,
	0x73, 0x68, 0x6f, 0x77, 0x5f, 0x63, 0x75, 0x74, 0x73, 0x63, 0x65, 0x6e, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0e, 0x69, 0x73, 0x53, 0x68, 0x6f, 0x77, 0x43, 0x75, 0x74, 0x73, 0x63,
	0x65, 0x6e, 0x65, 0x42, 0x1d, 0x5a, 0x1b, 0x47, 0x6f, 0x2d, 0x47, 0x72, 0x61, 0x73, 0x73, 0x63,
	0x75, 0x74, 0x74, 0x65, 0x72, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_CreateGadgetInfo_proto_rawDescOnce sync.Once
	file_proto_CreateGadgetInfo_proto_rawDescData = file_proto_CreateGadgetInfo_proto_rawDesc
)

func file_proto_CreateGadgetInfo_proto_rawDescGZIP() []byte {
	file_proto_CreateGadgetInfo_proto_rawDescOnce.Do(func() {
		file_proto_CreateGadgetInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_CreateGadgetInfo_proto_rawDescData)
	})
	return file_proto_CreateGadgetInfo_proto_rawDescData
}

var file_proto_CreateGadgetInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_CreateGadgetInfo_proto_goTypes = []interface{}{
	(*CreateGadgetInfo)(nil),       // 0: CreateGadgetInfo
	(*CreateGadgetInfo_Chest)(nil), // 1: CreateGadgetInfo.Chest
	(GadgetBornType)(0),            // 2: GadgetBornType
}
var file_proto_CreateGadgetInfo_proto_depIdxs = []int32{
	2, // 0: CreateGadgetInfo.born_type:type_name -> GadgetBornType
	1, // 1: CreateGadgetInfo.chest:type_name -> CreateGadgetInfo.Chest
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_CreateGadgetInfo_proto_init() }
func file_proto_CreateGadgetInfo_proto_init() {
	if File_proto_CreateGadgetInfo_proto != nil {
		return
	}
	file_proto_GadgetBornType_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_CreateGadgetInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateGadgetInfo); i {
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
		file_proto_CreateGadgetInfo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateGadgetInfo_Chest); i {
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
			RawDescriptor: file_proto_CreateGadgetInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_CreateGadgetInfo_proto_goTypes,
		DependencyIndexes: file_proto_CreateGadgetInfo_proto_depIdxs,
		MessageInfos:      file_proto_CreateGadgetInfo_proto_msgTypes,
	}.Build()
	File_proto_CreateGadgetInfo_proto = out.File
	file_proto_CreateGadgetInfo_proto_rawDesc = nil
	file_proto_CreateGadgetInfo_proto_goTypes = nil
	file_proto_CreateGadgetInfo_proto_depIdxs = nil
}
