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
// source: proto/MotionInfo.proto

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

// Obf: LJEPKMPIOKJ
type MotionInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pos              *Vector     `protobuf:"bytes,1,opt,name=pos,proto3" json:"pos,omitempty"`
	Rot              *Vector     `protobuf:"bytes,2,opt,name=rot,proto3" json:"rot,omitempty"`
	Speed            *Vector     `protobuf:"bytes,3,opt,name=speed,proto3" json:"speed,omitempty"`
	State            MotionState `protobuf:"varint,4,opt,name=state,proto3,enum=MotionState" json:"state,omitempty"`
	Params           []*Vector   `protobuf:"bytes,5,rep,name=params,proto3" json:"params,omitempty"`
	RefPos           *Vector     `protobuf:"bytes,6,opt,name=ref_pos,json=refPos,proto3" json:"ref_pos,omitempty"`
	RefId            uint32      `protobuf:"varint,7,opt,name=ref_id,json=refId,proto3" json:"ref_id,omitempty"`
	SceneTime        uint32      `protobuf:"varint,8,opt,name=scene_time,json=sceneTime,proto3" json:"scene_time,omitempty"`
	IntervalVelocity uint64      `protobuf:"varint,9,opt,name=interval_velocity,json=intervalVelocity,proto3" json:"interval_velocity,omitempty"`
}

func (x *MotionInfo) Reset() {
	*x = MotionInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_MotionInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MotionInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MotionInfo) ProtoMessage() {}

func (x *MotionInfo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_MotionInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MotionInfo.ProtoReflect.Descriptor instead.
func (*MotionInfo) Descriptor() ([]byte, []int) {
	return file_proto_MotionInfo_proto_rawDescGZIP(), []int{0}
}

func (x *MotionInfo) GetPos() *Vector {
	if x != nil {
		return x.Pos
	}
	return nil
}

func (x *MotionInfo) GetRot() *Vector {
	if x != nil {
		return x.Rot
	}
	return nil
}

func (x *MotionInfo) GetSpeed() *Vector {
	if x != nil {
		return x.Speed
	}
	return nil
}

func (x *MotionInfo) GetState() MotionState {
	if x != nil {
		return x.State
	}
	return MotionState_MOTION_STATE_NONE
}

func (x *MotionInfo) GetParams() []*Vector {
	if x != nil {
		return x.Params
	}
	return nil
}

func (x *MotionInfo) GetRefPos() *Vector {
	if x != nil {
		return x.RefPos
	}
	return nil
}

func (x *MotionInfo) GetRefId() uint32 {
	if x != nil {
		return x.RefId
	}
	return 0
}

func (x *MotionInfo) GetSceneTime() uint32 {
	if x != nil {
		return x.SceneTime
	}
	return 0
}

func (x *MotionInfo) GetIntervalVelocity() uint64 {
	if x != nil {
		return x.IntervalVelocity
	}
	return 0
}

var File_proto_MotionInfo_proto protoreflect.FileDescriptor

var file_proto_MotionInfo_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x4d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e,
	0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x4d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xab, 0x02, 0x0a, 0x0a, 0x4d, 0x6f, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x19, 0x0a, 0x03, 0x70, 0x6f, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x07, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x03, 0x70, 0x6f, 0x73, 0x12,
	0x19, 0x0a, 0x03, 0x72, 0x6f, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x56,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x03, 0x72, 0x6f, 0x74, 0x12, 0x1d, 0x0a, 0x05, 0x73, 0x70,
	0x65, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x56, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x52, 0x05, 0x73, 0x70, 0x65, 0x65, 0x64, 0x12, 0x22, 0x0a, 0x05, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0c, 0x2e, 0x4d, 0x6f, 0x74, 0x69, 0x6f,
	0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1f, 0x0a,
	0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x07, 0x2e,
	0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x20,
	0x0a, 0x07, 0x72, 0x65, 0x66, 0x5f, 0x70, 0x6f, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x07, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x06, 0x72, 0x65, 0x66, 0x50, 0x6f, 0x73,
	0x12, 0x15, 0x0a, 0x06, 0x72, 0x65, 0x66, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x05, 0x72, 0x65, 0x66, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x63, 0x65, 0x6e, 0x65,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x73, 0x63, 0x65,
	0x6e, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x2b, 0x0a, 0x11, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76,
	0x61, 0x6c, 0x5f, 0x76, 0x65, 0x6c, 0x6f, 0x63, 0x69, 0x74, 0x79, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x10, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x56, 0x65, 0x6c, 0x6f, 0x63,
	0x69, 0x74, 0x79, 0x42, 0x1d, 0x5a, 0x1b, 0x47, 0x6f, 0x2d, 0x47, 0x72, 0x61, 0x73, 0x73, 0x63,
	0x75, 0x74, 0x74, 0x65, 0x72, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_MotionInfo_proto_rawDescOnce sync.Once
	file_proto_MotionInfo_proto_rawDescData = file_proto_MotionInfo_proto_rawDesc
)

func file_proto_MotionInfo_proto_rawDescGZIP() []byte {
	file_proto_MotionInfo_proto_rawDescOnce.Do(func() {
		file_proto_MotionInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_MotionInfo_proto_rawDescData)
	})
	return file_proto_MotionInfo_proto_rawDescData
}

var file_proto_MotionInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_MotionInfo_proto_goTypes = []interface{}{
	(*MotionInfo)(nil), // 0: MotionInfo
	(*Vector)(nil),     // 1: Vector
	(MotionState)(0),   // 2: MotionState
}
var file_proto_MotionInfo_proto_depIdxs = []int32{
	1, // 0: MotionInfo.pos:type_name -> Vector
	1, // 1: MotionInfo.rot:type_name -> Vector
	1, // 2: MotionInfo.speed:type_name -> Vector
	2, // 3: MotionInfo.state:type_name -> MotionState
	1, // 4: MotionInfo.params:type_name -> Vector
	1, // 5: MotionInfo.ref_pos:type_name -> Vector
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_proto_MotionInfo_proto_init() }
func file_proto_MotionInfo_proto_init() {
	if File_proto_MotionInfo_proto != nil {
		return
	}
	file_proto_Vector_proto_init()
	file_proto_MotionState_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_MotionInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MotionInfo); i {
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
			RawDescriptor: file_proto_MotionInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_MotionInfo_proto_goTypes,
		DependencyIndexes: file_proto_MotionInfo_proto_depIdxs,
		MessageInfos:      file_proto_MotionInfo_proto_msgTypes,
	}.Build()
	File_proto_MotionInfo_proto = out.File
	file_proto_MotionInfo_proto_rawDesc = nil
	file_proto_MotionInfo_proto_goTypes = nil
	file_proto_MotionInfo_proto_depIdxs = nil
}
