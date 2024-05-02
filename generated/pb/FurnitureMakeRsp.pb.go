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
// source: proto/FurnitureMakeRsp.proto

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

// CmdId: 7783
// Obf: GPDPDPCIOPN
type FurnitureMakeRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MakeInfoList      []*FurnitureMakeMakeInfo     `protobuf:"bytes,5,rep,name=make_info_list,json=makeInfoList,proto3" json:"make_info_list,omitempty"`
	Retcode           int32                        `protobuf:"varint,14,opt,name=retcode,proto3" json:"retcode,omitempty"`
	HelpDataList      []*FurnitureMakeHelpData     `protobuf:"bytes,6,rep,name=help_data_list,json=helpDataList,proto3" json:"help_data_list,omitempty"`
	HelpedDataList    []*FurnitureMakeBeHelpedData `protobuf:"bytes,4,rep,name=helped_data_list,json=helpedDataList,proto3" json:"helped_data_list,omitempty"`
	FurnitureMakeSlot *FurnitureMakeSlot           `protobuf:"bytes,8,opt,name=furniture_make_slot,json=furnitureMakeSlot,proto3" json:"furniture_make_slot,omitempty"`
}

func (x *FurnitureMakeRsp) Reset() {
	*x = FurnitureMakeRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_FurnitureMakeRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FurnitureMakeRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FurnitureMakeRsp) ProtoMessage() {}

func (x *FurnitureMakeRsp) ProtoReflect() protoreflect.Message {
	mi := &file_proto_FurnitureMakeRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FurnitureMakeRsp.ProtoReflect.Descriptor instead.
func (*FurnitureMakeRsp) Descriptor() ([]byte, []int) {
	return file_proto_FurnitureMakeRsp_proto_rawDescGZIP(), []int{0}
}

func (x *FurnitureMakeRsp) GetMakeInfoList() []*FurnitureMakeMakeInfo {
	if x != nil {
		return x.MakeInfoList
	}
	return nil
}

func (x *FurnitureMakeRsp) GetRetcode() int32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *FurnitureMakeRsp) GetHelpDataList() []*FurnitureMakeHelpData {
	if x != nil {
		return x.HelpDataList
	}
	return nil
}

func (x *FurnitureMakeRsp) GetHelpedDataList() []*FurnitureMakeBeHelpedData {
	if x != nil {
		return x.HelpedDataList
	}
	return nil
}

func (x *FurnitureMakeRsp) GetFurnitureMakeSlot() *FurnitureMakeSlot {
	if x != nil {
		return x.FurnitureMakeSlot
	}
	return nil
}

var File_proto_FurnitureMakeRsp_proto protoreflect.FileDescriptor

var file_proto_FurnitureMakeRsp_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x46, 0x75, 0x72, 0x6e, 0x69, 0x74, 0x75, 0x72,
	0x65, 0x4d, 0x61, 0x6b, 0x65, 0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x46, 0x75, 0x72, 0x6e, 0x69, 0x74, 0x75, 0x72, 0x65, 0x4d,
	0x61, 0x6b, 0x65, 0x4d, 0x61, 0x6b, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x21, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x46, 0x75, 0x72, 0x6e, 0x69, 0x74, 0x75,
	0x72, 0x65, 0x4d, 0x61, 0x6b, 0x65, 0x48, 0x65, 0x6c, 0x70, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x25, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x46, 0x75, 0x72, 0x6e,
	0x69, 0x74, 0x75, 0x72, 0x65, 0x4d, 0x61, 0x6b, 0x65, 0x42, 0x65, 0x48, 0x65, 0x6c, 0x70, 0x65,
	0x64, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x46, 0x75, 0x72, 0x6e, 0x69, 0x74, 0x75, 0x72, 0x65, 0x4d, 0x61, 0x6b, 0x65,
	0x53, 0x6c, 0x6f, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb2, 0x02, 0x0a, 0x10, 0x46,
	0x75, 0x72, 0x6e, 0x69, 0x74, 0x75, 0x72, 0x65, 0x4d, 0x61, 0x6b, 0x65, 0x52, 0x73, 0x70, 0x12,
	0x3c, 0x0a, 0x0e, 0x6d, 0x61, 0x6b, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x5f, 0x6c, 0x69, 0x73,
	0x74, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x46, 0x75, 0x72, 0x6e, 0x69, 0x74,
	0x75, 0x72, 0x65, 0x4d, 0x61, 0x6b, 0x65, 0x4d, 0x61, 0x6b, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x0c, 0x6d, 0x61, 0x6b, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x18, 0x0a,
	0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07,
	0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x3c, 0x0a, 0x0e, 0x68, 0x65, 0x6c, 0x70, 0x5f,
	0x64, 0x61, 0x74, 0x61, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x16, 0x2e, 0x46, 0x75, 0x72, 0x6e, 0x69, 0x74, 0x75, 0x72, 0x65, 0x4d, 0x61, 0x6b, 0x65, 0x48,
	0x65, 0x6c, 0x70, 0x44, 0x61, 0x74, 0x61, 0x52, 0x0c, 0x68, 0x65, 0x6c, 0x70, 0x44, 0x61, 0x74,
	0x61, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x44, 0x0a, 0x10, 0x68, 0x65, 0x6c, 0x70, 0x65, 0x64, 0x5f,
	0x64, 0x61, 0x74, 0x61, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x46, 0x75, 0x72, 0x6e, 0x69, 0x74, 0x75, 0x72, 0x65, 0x4d, 0x61, 0x6b, 0x65, 0x42,
	0x65, 0x48, 0x65, 0x6c, 0x70, 0x65, 0x64, 0x44, 0x61, 0x74, 0x61, 0x52, 0x0e, 0x68, 0x65, 0x6c,
	0x70, 0x65, 0x64, 0x44, 0x61, 0x74, 0x61, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x42, 0x0a, 0x13, 0x66,
	0x75, 0x72, 0x6e, 0x69, 0x74, 0x75, 0x72, 0x65, 0x5f, 0x6d, 0x61, 0x6b, 0x65, 0x5f, 0x73, 0x6c,
	0x6f, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x46, 0x75, 0x72, 0x6e, 0x69,
	0x74, 0x75, 0x72, 0x65, 0x4d, 0x61, 0x6b, 0x65, 0x53, 0x6c, 0x6f, 0x74, 0x52, 0x11, 0x66, 0x75,
	0x72, 0x6e, 0x69, 0x74, 0x75, 0x72, 0x65, 0x4d, 0x61, 0x6b, 0x65, 0x53, 0x6c, 0x6f, 0x74, 0x42,
	0x1d, 0x5a, 0x1b, 0x47, 0x6f, 0x2d, 0x47, 0x72, 0x61, 0x73, 0x73, 0x63, 0x75, 0x74, 0x74, 0x65,
	0x72, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_FurnitureMakeRsp_proto_rawDescOnce sync.Once
	file_proto_FurnitureMakeRsp_proto_rawDescData = file_proto_FurnitureMakeRsp_proto_rawDesc
)

func file_proto_FurnitureMakeRsp_proto_rawDescGZIP() []byte {
	file_proto_FurnitureMakeRsp_proto_rawDescOnce.Do(func() {
		file_proto_FurnitureMakeRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_FurnitureMakeRsp_proto_rawDescData)
	})
	return file_proto_FurnitureMakeRsp_proto_rawDescData
}

var file_proto_FurnitureMakeRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_FurnitureMakeRsp_proto_goTypes = []interface{}{
	(*FurnitureMakeRsp)(nil),          // 0: FurnitureMakeRsp
	(*FurnitureMakeMakeInfo)(nil),     // 1: FurnitureMakeMakeInfo
	(*FurnitureMakeHelpData)(nil),     // 2: FurnitureMakeHelpData
	(*FurnitureMakeBeHelpedData)(nil), // 3: FurnitureMakeBeHelpedData
	(*FurnitureMakeSlot)(nil),         // 4: FurnitureMakeSlot
}
var file_proto_FurnitureMakeRsp_proto_depIdxs = []int32{
	1, // 0: FurnitureMakeRsp.make_info_list:type_name -> FurnitureMakeMakeInfo
	2, // 1: FurnitureMakeRsp.help_data_list:type_name -> FurnitureMakeHelpData
	3, // 2: FurnitureMakeRsp.helped_data_list:type_name -> FurnitureMakeBeHelpedData
	4, // 3: FurnitureMakeRsp.furniture_make_slot:type_name -> FurnitureMakeSlot
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_proto_FurnitureMakeRsp_proto_init() }
func file_proto_FurnitureMakeRsp_proto_init() {
	if File_proto_FurnitureMakeRsp_proto != nil {
		return
	}
	file_proto_FurnitureMakeMakeInfo_proto_init()
	file_proto_FurnitureMakeHelpData_proto_init()
	file_proto_FurnitureMakeBeHelpedData_proto_init()
	file_proto_FurnitureMakeSlot_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_FurnitureMakeRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FurnitureMakeRsp); i {
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
			RawDescriptor: file_proto_FurnitureMakeRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_FurnitureMakeRsp_proto_goTypes,
		DependencyIndexes: file_proto_FurnitureMakeRsp_proto_depIdxs,
		MessageInfos:      file_proto_FurnitureMakeRsp_proto_msgTypes,
	}.Build()
	File_proto_FurnitureMakeRsp_proto = out.File
	file_proto_FurnitureMakeRsp_proto_rawDesc = nil
	file_proto_FurnitureMakeRsp_proto_goTypes = nil
	file_proto_FurnitureMakeRsp_proto_depIdxs = nil
}
