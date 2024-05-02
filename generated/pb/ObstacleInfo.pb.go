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
// source: proto/ObstacleInfo.proto

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

// Obf: HEJHGAICHLM
type ObstacleInfo_ShapeType int32

const (
	ObstacleInfo_OBSTACLE_SHAPE_CAPSULE ObstacleInfo_ShapeType = 0
	ObstacleInfo_OBSTACLE_SHAPE_BOX     ObstacleInfo_ShapeType = 1
)

// Enum value maps for ObstacleInfo_ShapeType.
var (
	ObstacleInfo_ShapeType_name = map[int32]string{
		0: "OBSTACLE_SHAPE_CAPSULE",
		1: "OBSTACLE_SHAPE_BOX",
	}
	ObstacleInfo_ShapeType_value = map[string]int32{
		"OBSTACLE_SHAPE_CAPSULE": 0,
		"OBSTACLE_SHAPE_BOX":     1,
	}
)

func (x ObstacleInfo_ShapeType) Enum() *ObstacleInfo_ShapeType {
	p := new(ObstacleInfo_ShapeType)
	*p = x
	return p
}

func (x ObstacleInfo_ShapeType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ObstacleInfo_ShapeType) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_ObstacleInfo_proto_enumTypes[0].Descriptor()
}

func (ObstacleInfo_ShapeType) Type() protoreflect.EnumType {
	return &file_proto_ObstacleInfo_proto_enumTypes[0]
}

func (x ObstacleInfo_ShapeType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ObstacleInfo_ShapeType.Descriptor instead.
func (ObstacleInfo_ShapeType) EnumDescriptor() ([]byte, []int) {
	return file_proto_ObstacleInfo_proto_rawDescGZIP(), []int{0, 0}
}

// Obf: HDJAECHANAJ
type ObstacleInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rotation   *MathQuaternion        `protobuf:"bytes,4,opt,name=rotation,proto3" json:"rotation,omitempty"`
	Shape      ObstacleInfo_ShapeType `protobuf:"varint,5,opt,name=shape,proto3,enum=ObstacleInfo_ShapeType" json:"shape,omitempty"`
	ObstacleId int32                  `protobuf:"varint,6,opt,name=obstacle_id,json=obstacleId,proto3" json:"obstacle_id,omitempty"`
	Extents    *Vector3Int            `protobuf:"bytes,15,opt,name=extents,proto3" json:"extents,omitempty"`
	Center     *Vector                `protobuf:"bytes,9,opt,name=center,proto3" json:"center,omitempty"`
}

func (x *ObstacleInfo) Reset() {
	*x = ObstacleInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ObstacleInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ObstacleInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ObstacleInfo) ProtoMessage() {}

func (x *ObstacleInfo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ObstacleInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ObstacleInfo.ProtoReflect.Descriptor instead.
func (*ObstacleInfo) Descriptor() ([]byte, []int) {
	return file_proto_ObstacleInfo_proto_rawDescGZIP(), []int{0}
}

func (x *ObstacleInfo) GetRotation() *MathQuaternion {
	if x != nil {
		return x.Rotation
	}
	return nil
}

func (x *ObstacleInfo) GetShape() ObstacleInfo_ShapeType {
	if x != nil {
		return x.Shape
	}
	return ObstacleInfo_OBSTACLE_SHAPE_CAPSULE
}

func (x *ObstacleInfo) GetObstacleId() int32 {
	if x != nil {
		return x.ObstacleId
	}
	return 0
}

func (x *ObstacleInfo) GetExtents() *Vector3Int {
	if x != nil {
		return x.Extents
	}
	return nil
}

func (x *ObstacleInfo) GetCenter() *Vector {
	if x != nil {
		return x.Center
	}
	return nil
}

var File_proto_ObstacleInfo_proto protoreflect.FileDescriptor

var file_proto_ObstacleInfo_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x4f, 0x62, 0x73, 0x74, 0x61, 0x63, 0x6c, 0x65,
	0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x4d, 0x61, 0x74, 0x68, 0x51, 0x75, 0x61, 0x74, 0x65, 0x72, 0x6e, 0x69, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x56, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x33, 0x49, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x94, 0x02, 0x0a, 0x0c, 0x4f, 0x62, 0x73, 0x74, 0x61, 0x63, 0x6c, 0x65, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x2b, 0x0a, 0x08, 0x72, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x4d, 0x61, 0x74, 0x68, 0x51, 0x75, 0x61, 0x74,
	0x65, 0x72, 0x6e, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x72, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x2d, 0x0a, 0x05, 0x73, 0x68, 0x61, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x17, 0x2e, 0x4f, 0x62, 0x73, 0x74, 0x61, 0x63, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x53,
	0x68, 0x61, 0x70, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x05, 0x73, 0x68, 0x61, 0x70, 0x65, 0x12,
	0x1f, 0x0a, 0x0b, 0x6f, 0x62, 0x73, 0x74, 0x61, 0x63, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x6f, 0x62, 0x73, 0x74, 0x61, 0x63, 0x6c, 0x65, 0x49, 0x64,
	0x12, 0x25, 0x0a, 0x07, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x0f, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0b, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x33, 0x49, 0x6e, 0x74, 0x52, 0x07,
	0x65, 0x78, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x1f, 0x0a, 0x06, 0x63, 0x65, 0x6e, 0x74, 0x65,
	0x72, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x52, 0x06, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x22, 0x3f, 0x0a, 0x09, 0x53, 0x68, 0x61, 0x70,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x16, 0x4f, 0x42, 0x53, 0x54, 0x41, 0x43, 0x4c,
	0x45, 0x5f, 0x53, 0x48, 0x41, 0x50, 0x45, 0x5f, 0x43, 0x41, 0x50, 0x53, 0x55, 0x4c, 0x45, 0x10,
	0x00, 0x12, 0x16, 0x0a, 0x12, 0x4f, 0x42, 0x53, 0x54, 0x41, 0x43, 0x4c, 0x45, 0x5f, 0x53, 0x48,
	0x41, 0x50, 0x45, 0x5f, 0x42, 0x4f, 0x58, 0x10, 0x01, 0x42, 0x1d, 0x5a, 0x1b, 0x47, 0x6f, 0x2d,
	0x47, 0x72, 0x61, 0x73, 0x73, 0x63, 0x75, 0x74, 0x74, 0x65, 0x72, 0x2f, 0x67, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_ObstacleInfo_proto_rawDescOnce sync.Once
	file_proto_ObstacleInfo_proto_rawDescData = file_proto_ObstacleInfo_proto_rawDesc
)

func file_proto_ObstacleInfo_proto_rawDescGZIP() []byte {
	file_proto_ObstacleInfo_proto_rawDescOnce.Do(func() {
		file_proto_ObstacleInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_ObstacleInfo_proto_rawDescData)
	})
	return file_proto_ObstacleInfo_proto_rawDescData
}

var file_proto_ObstacleInfo_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_ObstacleInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_ObstacleInfo_proto_goTypes = []interface{}{
	(ObstacleInfo_ShapeType)(0), // 0: ObstacleInfo.ShapeType
	(*ObstacleInfo)(nil),        // 1: ObstacleInfo
	(*MathQuaternion)(nil),      // 2: MathQuaternion
	(*Vector3Int)(nil),          // 3: Vector3Int
	(*Vector)(nil),              // 4: Vector
}
var file_proto_ObstacleInfo_proto_depIdxs = []int32{
	2, // 0: ObstacleInfo.rotation:type_name -> MathQuaternion
	0, // 1: ObstacleInfo.shape:type_name -> ObstacleInfo.ShapeType
	3, // 2: ObstacleInfo.extents:type_name -> Vector3Int
	4, // 3: ObstacleInfo.center:type_name -> Vector
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_proto_ObstacleInfo_proto_init() }
func file_proto_ObstacleInfo_proto_init() {
	if File_proto_ObstacleInfo_proto != nil {
		return
	}
	file_proto_MathQuaternion_proto_init()
	file_proto_Vector3Int_proto_init()
	file_proto_Vector_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_ObstacleInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ObstacleInfo); i {
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
			RawDescriptor: file_proto_ObstacleInfo_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_ObstacleInfo_proto_goTypes,
		DependencyIndexes: file_proto_ObstacleInfo_proto_depIdxs,
		EnumInfos:         file_proto_ObstacleInfo_proto_enumTypes,
		MessageInfos:      file_proto_ObstacleInfo_proto_msgTypes,
	}.Build()
	File_proto_ObstacleInfo_proto = out.File
	file_proto_ObstacleInfo_proto_rawDesc = nil
	file_proto_ObstacleInfo_proto_goTypes = nil
	file_proto_ObstacleInfo_proto_depIdxs = nil
}