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
// source: proto/ShowEquip.proto

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

// Obf: OJMNHOEOFAN
type ShowEquip struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ItemId uint32 `protobuf:"varint,1,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
	// Types that are assignable to Detail:
	//
	//	*ShowEquip_Reliquary
	//	*ShowEquip_Weapon
	Detail isShowEquip_Detail `protobuf_oneof:"detail"`
}

func (x *ShowEquip) Reset() {
	*x = ShowEquip{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ShowEquip_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShowEquip) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShowEquip) ProtoMessage() {}

func (x *ShowEquip) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ShowEquip_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShowEquip.ProtoReflect.Descriptor instead.
func (*ShowEquip) Descriptor() ([]byte, []int) {
	return file_proto_ShowEquip_proto_rawDescGZIP(), []int{0}
}

func (x *ShowEquip) GetItemId() uint32 {
	if x != nil {
		return x.ItemId
	}
	return 0
}

func (m *ShowEquip) GetDetail() isShowEquip_Detail {
	if m != nil {
		return m.Detail
	}
	return nil
}

func (x *ShowEquip) GetReliquary() *Reliquary {
	if x, ok := x.GetDetail().(*ShowEquip_Reliquary); ok {
		return x.Reliquary
	}
	return nil
}

func (x *ShowEquip) GetWeapon() *Weapon {
	if x, ok := x.GetDetail().(*ShowEquip_Weapon); ok {
		return x.Weapon
	}
	return nil
}

type isShowEquip_Detail interface {
	isShowEquip_Detail()
}

type ShowEquip_Reliquary struct {
	Reliquary *Reliquary `protobuf:"bytes,2,opt,name=reliquary,proto3,oneof"`
}

type ShowEquip_Weapon struct {
	Weapon *Weapon `protobuf:"bytes,3,opt,name=weapon,proto3,oneof"`
}

func (*ShowEquip_Reliquary) isShowEquip_Detail() {}

func (*ShowEquip_Weapon) isShowEquip_Detail() {}

var File_proto_ShowEquip_proto protoreflect.FileDescriptor

var file_proto_ShowEquip_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x53, 0x68, 0x6f, 0x77, 0x45, 0x71, 0x75, 0x69,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x52,
	0x65, 0x6c, 0x69, 0x71, 0x75, 0x61, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x57, 0x65, 0x61, 0x70, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x7d, 0x0a, 0x09, 0x53, 0x68, 0x6f, 0x77, 0x45, 0x71, 0x75, 0x69, 0x70, 0x12,
	0x17, 0x0a, 0x07, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x06, 0x69, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x09, 0x72, 0x65, 0x6c, 0x69,
	0x71, 0x75, 0x61, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x52, 0x65,
	0x6c, 0x69, 0x71, 0x75, 0x61, 0x72, 0x79, 0x48, 0x00, 0x52, 0x09, 0x72, 0x65, 0x6c, 0x69, 0x71,
	0x75, 0x61, 0x72, 0x79, 0x12, 0x21, 0x0a, 0x06, 0x77, 0x65, 0x61, 0x70, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x57, 0x65, 0x61, 0x70, 0x6f, 0x6e, 0x48, 0x00, 0x52,
	0x06, 0x77, 0x65, 0x61, 0x70, 0x6f, 0x6e, 0x42, 0x08, 0x0a, 0x06, 0x64, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x42, 0x1d, 0x5a, 0x1b, 0x47, 0x6f, 0x2d, 0x47, 0x72, 0x61, 0x73, 0x73, 0x63, 0x75, 0x74,
	0x74, 0x65, 0x72, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_ShowEquip_proto_rawDescOnce sync.Once
	file_proto_ShowEquip_proto_rawDescData = file_proto_ShowEquip_proto_rawDesc
)

func file_proto_ShowEquip_proto_rawDescGZIP() []byte {
	file_proto_ShowEquip_proto_rawDescOnce.Do(func() {
		file_proto_ShowEquip_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_ShowEquip_proto_rawDescData)
	})
	return file_proto_ShowEquip_proto_rawDescData
}

var file_proto_ShowEquip_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_ShowEquip_proto_goTypes = []interface{}{
	(*ShowEquip)(nil), // 0: ShowEquip
	(*Reliquary)(nil), // 1: Reliquary
	(*Weapon)(nil),    // 2: Weapon
}
var file_proto_ShowEquip_proto_depIdxs = []int32{
	1, // 0: ShowEquip.reliquary:type_name -> Reliquary
	2, // 1: ShowEquip.weapon:type_name -> Weapon
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_ShowEquip_proto_init() }
func file_proto_ShowEquip_proto_init() {
	if File_proto_ShowEquip_proto != nil {
		return
	}
	file_proto_Reliquary_proto_init()
	file_proto_Weapon_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_ShowEquip_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShowEquip); i {
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
	file_proto_ShowEquip_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*ShowEquip_Reliquary)(nil),
		(*ShowEquip_Weapon)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_ShowEquip_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_ShowEquip_proto_goTypes,
		DependencyIndexes: file_proto_ShowEquip_proto_depIdxs,
		MessageInfos:      file_proto_ShowEquip_proto_msgTypes,
	}.Build()
	File_proto_ShowEquip_proto = out.File
	file_proto_ShowEquip_proto_rawDesc = nil
	file_proto_ShowEquip_proto_goTypes = nil
	file_proto_ShowEquip_proto_depIdxs = nil
}
