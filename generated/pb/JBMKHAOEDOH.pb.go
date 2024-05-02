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
// source: proto/JBMKHAOEDOH.proto

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

type JBMKHAOEDOH_NBAEOAGGMEM int32

const (
	JBMKHAOEDOH_HMNKHDMCEKN_StatusFail    JBMKHAOEDOH_NBAEOAGGMEM = 0
	JBMKHAOEDOH_HMNKHDMCEKN_StatusSucc    JBMKHAOEDOH_NBAEOAGGMEM = 1
	JBMKHAOEDOH_HMNKHDMCEKN_StatusPartial JBMKHAOEDOH_NBAEOAGGMEM = 2
)

// Enum value maps for JBMKHAOEDOH_NBAEOAGGMEM.
var (
	JBMKHAOEDOH_NBAEOAGGMEM_name = map[int32]string{
		0: "HMNKHDMCEKN_StatusFail",
		1: "HMNKHDMCEKN_StatusSucc",
		2: "HMNKHDMCEKN_StatusPartial",
	}
	JBMKHAOEDOH_NBAEOAGGMEM_value = map[string]int32{
		"HMNKHDMCEKN_StatusFail":    0,
		"HMNKHDMCEKN_StatusSucc":    1,
		"HMNKHDMCEKN_StatusPartial": 2,
	}
)

func (x JBMKHAOEDOH_NBAEOAGGMEM) Enum() *JBMKHAOEDOH_NBAEOAGGMEM {
	p := new(JBMKHAOEDOH_NBAEOAGGMEM)
	*p = x
	return p
}

func (x JBMKHAOEDOH_NBAEOAGGMEM) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (JBMKHAOEDOH_NBAEOAGGMEM) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_JBMKHAOEDOH_proto_enumTypes[0].Descriptor()
}

func (JBMKHAOEDOH_NBAEOAGGMEM) Type() protoreflect.EnumType {
	return &file_proto_JBMKHAOEDOH_proto_enumTypes[0]
}

func (x JBMKHAOEDOH_NBAEOAGGMEM) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use JBMKHAOEDOH_NBAEOAGGMEM.Descriptor instead.
func (JBMKHAOEDOH_NBAEOAGGMEM) EnumDescriptor() ([]byte, []int) {
	return file_proto_JBMKHAOEDOH_proto_rawDescGZIP(), []int{0, 0}
}

// CmdId: 2345
type JBMKHAOEDOH struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	QueryId     int32                   `protobuf:"varint,10,opt,name=query_id,json=queryId,proto3" json:"query_id,omitempty"`
	FJOJPPADDCM []*DCIGLMFIBKP          `protobuf:"bytes,13,rep,name=FJOJPPADDCM,proto3" json:"FJOJPPADDCM,omitempty"`
	Retcode     int32                   `protobuf:"varint,3,opt,name=retcode,proto3" json:"retcode,omitempty"`
	PDMHNHLGOAP []*Vector               `protobuf:"bytes,4,rep,name=PDMHNHLGOAP,proto3" json:"PDMHNHLGOAP,omitempty"`
	QueryStatus JBMKHAOEDOH_NBAEOAGGMEM `protobuf:"varint,6,opt,name=query_status,json=queryStatus,proto3,enum=JBMKHAOEDOH_NBAEOAGGMEM" json:"query_status,omitempty"`
	Corners     []*Vector               `protobuf:"bytes,2,rep,name=corners,proto3" json:"corners,omitempty"`
	DCPAICGCAHG []*DCIGLMFIBKP          `protobuf:"bytes,9,rep,name=DCPAICGCAHG,proto3" json:"DCPAICGCAHG,omitempty"`
}

func (x *JBMKHAOEDOH) Reset() {
	*x = JBMKHAOEDOH{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_JBMKHAOEDOH_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JBMKHAOEDOH) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JBMKHAOEDOH) ProtoMessage() {}

func (x *JBMKHAOEDOH) ProtoReflect() protoreflect.Message {
	mi := &file_proto_JBMKHAOEDOH_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JBMKHAOEDOH.ProtoReflect.Descriptor instead.
func (*JBMKHAOEDOH) Descriptor() ([]byte, []int) {
	return file_proto_JBMKHAOEDOH_proto_rawDescGZIP(), []int{0}
}

func (x *JBMKHAOEDOH) GetQueryId() int32 {
	if x != nil {
		return x.QueryId
	}
	return 0
}

func (x *JBMKHAOEDOH) GetFJOJPPADDCM() []*DCIGLMFIBKP {
	if x != nil {
		return x.FJOJPPADDCM
	}
	return nil
}

func (x *JBMKHAOEDOH) GetRetcode() int32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *JBMKHAOEDOH) GetPDMHNHLGOAP() []*Vector {
	if x != nil {
		return x.PDMHNHLGOAP
	}
	return nil
}

func (x *JBMKHAOEDOH) GetQueryStatus() JBMKHAOEDOH_NBAEOAGGMEM {
	if x != nil {
		return x.QueryStatus
	}
	return JBMKHAOEDOH_HMNKHDMCEKN_StatusFail
}

func (x *JBMKHAOEDOH) GetCorners() []*Vector {
	if x != nil {
		return x.Corners
	}
	return nil
}

func (x *JBMKHAOEDOH) GetDCPAICGCAHG() []*DCIGLMFIBKP {
	if x != nil {
		return x.DCPAICGCAHG
	}
	return nil
}

var File_proto_JBMKHAOEDOH_proto protoreflect.FileDescriptor

var file_proto_JBMKHAOEDOH_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x4a, 0x42, 0x4d, 0x4b, 0x48, 0x41, 0x4f, 0x45,
	0x44, 0x4f, 0x48, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x44, 0x43, 0x49, 0x47, 0x4c, 0x4d, 0x46, 0x49, 0x42, 0x4b, 0x50, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x93, 0x03, 0x0a, 0x0b, 0x4a, 0x42, 0x4d, 0x4b, 0x48,
	0x41, 0x4f, 0x45, 0x44, 0x4f, 0x48, 0x12, 0x19, 0x0a, 0x08, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f,
	0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x71, 0x75, 0x65, 0x72, 0x79, 0x49,
	0x64, 0x12, 0x2e, 0x0a, 0x0b, 0x46, 0x4a, 0x4f, 0x4a, 0x50, 0x50, 0x41, 0x44, 0x44, 0x43, 0x4d,
	0x18, 0x0d, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x44, 0x43, 0x49, 0x47, 0x4c, 0x4d, 0x46,
	0x49, 0x42, 0x4b, 0x50, 0x52, 0x0b, 0x46, 0x4a, 0x4f, 0x4a, 0x50, 0x50, 0x41, 0x44, 0x44, 0x43,
	0x4d, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x29, 0x0a, 0x0b, 0x50,
	0x44, 0x4d, 0x48, 0x4e, 0x48, 0x4c, 0x47, 0x4f, 0x41, 0x50, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x07, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x0b, 0x50, 0x44, 0x4d, 0x48, 0x4e,
	0x48, 0x4c, 0x47, 0x4f, 0x41, 0x50, 0x12, 0x3b, 0x0a, 0x0c, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x4a,
	0x42, 0x4d, 0x4b, 0x48, 0x41, 0x4f, 0x45, 0x44, 0x4f, 0x48, 0x2e, 0x4e, 0x42, 0x41, 0x45, 0x4f,
	0x41, 0x47, 0x47, 0x4d, 0x45, 0x4d, 0x52, 0x0b, 0x71, 0x75, 0x65, 0x72, 0x79, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x21, 0x0a, 0x07, 0x63, 0x6f, 0x72, 0x6e, 0x65, 0x72, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x07, 0x63,
	0x6f, 0x72, 0x6e, 0x65, 0x72, 0x73, 0x12, 0x2e, 0x0a, 0x0b, 0x44, 0x43, 0x50, 0x41, 0x49, 0x43,
	0x47, 0x43, 0x41, 0x48, 0x47, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x44, 0x43,
	0x49, 0x47, 0x4c, 0x4d, 0x46, 0x49, 0x42, 0x4b, 0x50, 0x52, 0x0b, 0x44, 0x43, 0x50, 0x41, 0x49,
	0x43, 0x47, 0x43, 0x41, 0x48, 0x47, 0x22, 0x64, 0x0a, 0x0b, 0x4e, 0x42, 0x41, 0x45, 0x4f, 0x41,
	0x47, 0x47, 0x4d, 0x45, 0x4d, 0x12, 0x1a, 0x0a, 0x16, 0x48, 0x4d, 0x4e, 0x4b, 0x48, 0x44, 0x4d,
	0x43, 0x45, 0x4b, 0x4e, 0x5f, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x46, 0x61, 0x69, 0x6c, 0x10,
	0x00, 0x12, 0x1a, 0x0a, 0x16, 0x48, 0x4d, 0x4e, 0x4b, 0x48, 0x44, 0x4d, 0x43, 0x45, 0x4b, 0x4e,
	0x5f, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x53, 0x75, 0x63, 0x63, 0x10, 0x01, 0x12, 0x1d, 0x0a,
	0x19, 0x48, 0x4d, 0x4e, 0x4b, 0x48, 0x44, 0x4d, 0x43, 0x45, 0x4b, 0x4e, 0x5f, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x50, 0x61, 0x72, 0x74, 0x69, 0x61, 0x6c, 0x10, 0x02, 0x42, 0x1d, 0x5a, 0x1b,
	0x47, 0x6f, 0x2d, 0x47, 0x72, 0x61, 0x73, 0x73, 0x63, 0x75, 0x74, 0x74, 0x65, 0x72, 0x2f, 0x67,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_proto_JBMKHAOEDOH_proto_rawDescOnce sync.Once
	file_proto_JBMKHAOEDOH_proto_rawDescData = file_proto_JBMKHAOEDOH_proto_rawDesc
)

func file_proto_JBMKHAOEDOH_proto_rawDescGZIP() []byte {
	file_proto_JBMKHAOEDOH_proto_rawDescOnce.Do(func() {
		file_proto_JBMKHAOEDOH_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_JBMKHAOEDOH_proto_rawDescData)
	})
	return file_proto_JBMKHAOEDOH_proto_rawDescData
}

var file_proto_JBMKHAOEDOH_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_JBMKHAOEDOH_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_JBMKHAOEDOH_proto_goTypes = []interface{}{
	(JBMKHAOEDOH_NBAEOAGGMEM)(0), // 0: JBMKHAOEDOH.NBAEOAGGMEM
	(*JBMKHAOEDOH)(nil),          // 1: JBMKHAOEDOH
	(*DCIGLMFIBKP)(nil),          // 2: DCIGLMFIBKP
	(*Vector)(nil),               // 3: Vector
}
var file_proto_JBMKHAOEDOH_proto_depIdxs = []int32{
	2, // 0: JBMKHAOEDOH.FJOJPPADDCM:type_name -> DCIGLMFIBKP
	3, // 1: JBMKHAOEDOH.PDMHNHLGOAP:type_name -> Vector
	0, // 2: JBMKHAOEDOH.query_status:type_name -> JBMKHAOEDOH.NBAEOAGGMEM
	3, // 3: JBMKHAOEDOH.corners:type_name -> Vector
	2, // 4: JBMKHAOEDOH.DCPAICGCAHG:type_name -> DCIGLMFIBKP
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_proto_JBMKHAOEDOH_proto_init() }
func file_proto_JBMKHAOEDOH_proto_init() {
	if File_proto_JBMKHAOEDOH_proto != nil {
		return
	}
	file_proto_DCIGLMFIBKP_proto_init()
	file_proto_Vector_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_JBMKHAOEDOH_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JBMKHAOEDOH); i {
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
			RawDescriptor: file_proto_JBMKHAOEDOH_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_JBMKHAOEDOH_proto_goTypes,
		DependencyIndexes: file_proto_JBMKHAOEDOH_proto_depIdxs,
		EnumInfos:         file_proto_JBMKHAOEDOH_proto_enumTypes,
		MessageInfos:      file_proto_JBMKHAOEDOH_proto_msgTypes,
	}.Build()
	File_proto_JBMKHAOEDOH_proto = out.File
	file_proto_JBMKHAOEDOH_proto_rawDesc = nil
	file_proto_JBMKHAOEDOH_proto_goTypes = nil
	file_proto_JBMKHAOEDOH_proto_depIdxs = nil
}
