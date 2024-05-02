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
// source: proto/GCGDSDeckData.proto

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

// Obf: NNODKGHGGBA
type GCGDSDeckData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FieldId     uint32   `protobuf:"varint,9,opt,name=field_id,json=fieldId,proto3" json:"field_id,omitempty"`
	CreateTime  uint32   `protobuf:"fixed32,10,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	Name        string   `protobuf:"bytes,7,opt,name=name,proto3" json:"name,omitempty"`
	Id          uint32   `protobuf:"varint,5,opt,name=id,proto3" json:"id,omitempty"`
	CardBackId  uint32   `protobuf:"varint,11,opt,name=card_back_id,json=cardBackId,proto3" json:"card_back_id,omitempty"`
	APCFHCPFONE []uint32 `protobuf:"varint,4,rep,packed,name=APCFHCPFONE,proto3" json:"APCFHCPFONE,omitempty"`
	IsValid     bool     `protobuf:"varint,1,opt,name=is_valid,json=isValid,proto3" json:"is_valid,omitempty"`
	CardList    []uint32 `protobuf:"varint,13,rep,packed,name=card_list,json=cardList,proto3" json:"card_list,omitempty"`
}

func (x *GCGDSDeckData) Reset() {
	*x = GCGDSDeckData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_GCGDSDeckData_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GCGDSDeckData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GCGDSDeckData) ProtoMessage() {}

func (x *GCGDSDeckData) ProtoReflect() protoreflect.Message {
	mi := &file_proto_GCGDSDeckData_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GCGDSDeckData.ProtoReflect.Descriptor instead.
func (*GCGDSDeckData) Descriptor() ([]byte, []int) {
	return file_proto_GCGDSDeckData_proto_rawDescGZIP(), []int{0}
}

func (x *GCGDSDeckData) GetFieldId() uint32 {
	if x != nil {
		return x.FieldId
	}
	return 0
}

func (x *GCGDSDeckData) GetCreateTime() uint32 {
	if x != nil {
		return x.CreateTime
	}
	return 0
}

func (x *GCGDSDeckData) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GCGDSDeckData) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GCGDSDeckData) GetCardBackId() uint32 {
	if x != nil {
		return x.CardBackId
	}
	return 0
}

func (x *GCGDSDeckData) GetAPCFHCPFONE() []uint32 {
	if x != nil {
		return x.APCFHCPFONE
	}
	return nil
}

func (x *GCGDSDeckData) GetIsValid() bool {
	if x != nil {
		return x.IsValid
	}
	return false
}

func (x *GCGDSDeckData) GetCardList() []uint32 {
	if x != nil {
		return x.CardList
	}
	return nil
}

var File_proto_GCGDSDeckData_proto protoreflect.FileDescriptor

var file_proto_GCGDSDeckData_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x47, 0x43, 0x47, 0x44, 0x53, 0x44, 0x65, 0x63,
	0x6b, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xeb, 0x01, 0x0a, 0x0d,
	0x47, 0x43, 0x47, 0x44, 0x53, 0x44, 0x65, 0x63, 0x6b, 0x44, 0x61, 0x74, 0x61, 0x12, 0x19, 0x0a,
	0x08, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x07, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x07, 0x52, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x20, 0x0a,
	0x0c, 0x63, 0x61, 0x72, 0x64, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0a, 0x63, 0x61, 0x72, 0x64, 0x42, 0x61, 0x63, 0x6b, 0x49, 0x64, 0x12,
	0x20, 0x0a, 0x0b, 0x41, 0x50, 0x43, 0x46, 0x48, 0x43, 0x50, 0x46, 0x4f, 0x4e, 0x45, 0x18, 0x04,
	0x20, 0x03, 0x28, 0x0d, 0x52, 0x0b, 0x41, 0x50, 0x43, 0x46, 0x48, 0x43, 0x50, 0x46, 0x4f, 0x4e,
	0x45, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x73, 0x5f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x73, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09,
	0x63, 0x61, 0x72, 0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x0d, 0x52,
	0x08, 0x63, 0x61, 0x72, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x1d, 0x5a, 0x1b, 0x47, 0x6f, 0x2d,
	0x47, 0x72, 0x61, 0x73, 0x73, 0x63, 0x75, 0x74, 0x74, 0x65, 0x72, 0x2f, 0x67, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_GCGDSDeckData_proto_rawDescOnce sync.Once
	file_proto_GCGDSDeckData_proto_rawDescData = file_proto_GCGDSDeckData_proto_rawDesc
)

func file_proto_GCGDSDeckData_proto_rawDescGZIP() []byte {
	file_proto_GCGDSDeckData_proto_rawDescOnce.Do(func() {
		file_proto_GCGDSDeckData_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_GCGDSDeckData_proto_rawDescData)
	})
	return file_proto_GCGDSDeckData_proto_rawDescData
}

var file_proto_GCGDSDeckData_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_GCGDSDeckData_proto_goTypes = []interface{}{
	(*GCGDSDeckData)(nil), // 0: GCGDSDeckData
}
var file_proto_GCGDSDeckData_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_GCGDSDeckData_proto_init() }
func file_proto_GCGDSDeckData_proto_init() {
	if File_proto_GCGDSDeckData_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_GCGDSDeckData_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GCGDSDeckData); i {
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
			RawDescriptor: file_proto_GCGDSDeckData_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_GCGDSDeckData_proto_goTypes,
		DependencyIndexes: file_proto_GCGDSDeckData_proto_depIdxs,
		MessageInfos:      file_proto_GCGDSDeckData_proto_msgTypes,
	}.Build()
	File_proto_GCGDSDeckData_proto = out.File
	file_proto_GCGDSDeckData_proto_rawDesc = nil
	file_proto_GCGDSDeckData_proto_goTypes = nil
	file_proto_GCGDSDeckData_proto_depIdxs = nil
}
