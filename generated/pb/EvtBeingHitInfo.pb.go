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
// source: proto/EvtBeingHitInfo.proto

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

// Obf: LGMKEKBEELG
type EvtBeingHitInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PeerId       uint32        `protobuf:"varint,4,opt,name=peer_id,json=peerId,proto3" json:"peer_id,omitempty"`
	FrameNum     uint32        `protobuf:"varint,15,opt,name=frame_num,json=frameNum,proto3" json:"frame_num,omitempty"`
	AttackResult *AttackResult `protobuf:"bytes,10,opt,name=attack_result,json=attackResult,proto3" json:"attack_result,omitempty"`
}

func (x *EvtBeingHitInfo) Reset() {
	*x = EvtBeingHitInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_EvtBeingHitInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EvtBeingHitInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EvtBeingHitInfo) ProtoMessage() {}

func (x *EvtBeingHitInfo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_EvtBeingHitInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EvtBeingHitInfo.ProtoReflect.Descriptor instead.
func (*EvtBeingHitInfo) Descriptor() ([]byte, []int) {
	return file_proto_EvtBeingHitInfo_proto_rawDescGZIP(), []int{0}
}

func (x *EvtBeingHitInfo) GetPeerId() uint32 {
	if x != nil {
		return x.PeerId
	}
	return 0
}

func (x *EvtBeingHitInfo) GetFrameNum() uint32 {
	if x != nil {
		return x.FrameNum
	}
	return 0
}

func (x *EvtBeingHitInfo) GetAttackResult() *AttackResult {
	if x != nil {
		return x.AttackResult
	}
	return nil
}

var File_proto_EvtBeingHitInfo_proto protoreflect.FileDescriptor

var file_proto_EvtBeingHitInfo_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x45, 0x76, 0x74, 0x42, 0x65, 0x69, 0x6e, 0x67,
	0x48, 0x69, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x41, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7b, 0x0a, 0x0f, 0x45, 0x76, 0x74, 0x42, 0x65,
	0x69, 0x6e, 0x67, 0x48, 0x69, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x65,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x70, 0x65, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x5f, 0x6e, 0x75, 0x6d,
	0x18, 0x0f, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x4e, 0x75, 0x6d,
	0x12, 0x32, 0x0a, 0x0d, 0x61, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x41, 0x74, 0x74, 0x61, 0x63, 0x6b,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x0c, 0x61, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x42, 0x1d, 0x5a, 0x1b, 0x47, 0x6f, 0x2d, 0x47, 0x72, 0x61, 0x73, 0x73,
	0x63, 0x75, 0x74, 0x74, 0x65, 0x72, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64,
	0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_EvtBeingHitInfo_proto_rawDescOnce sync.Once
	file_proto_EvtBeingHitInfo_proto_rawDescData = file_proto_EvtBeingHitInfo_proto_rawDesc
)

func file_proto_EvtBeingHitInfo_proto_rawDescGZIP() []byte {
	file_proto_EvtBeingHitInfo_proto_rawDescOnce.Do(func() {
		file_proto_EvtBeingHitInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_EvtBeingHitInfo_proto_rawDescData)
	})
	return file_proto_EvtBeingHitInfo_proto_rawDescData
}

var file_proto_EvtBeingHitInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_EvtBeingHitInfo_proto_goTypes = []interface{}{
	(*EvtBeingHitInfo)(nil), // 0: EvtBeingHitInfo
	(*AttackResult)(nil),    // 1: AttackResult
}
var file_proto_EvtBeingHitInfo_proto_depIdxs = []int32{
	1, // 0: EvtBeingHitInfo.attack_result:type_name -> AttackResult
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_EvtBeingHitInfo_proto_init() }
func file_proto_EvtBeingHitInfo_proto_init() {
	if File_proto_EvtBeingHitInfo_proto != nil {
		return
	}
	file_proto_AttackResult_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_EvtBeingHitInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EvtBeingHitInfo); i {
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
			RawDescriptor: file_proto_EvtBeingHitInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_EvtBeingHitInfo_proto_goTypes,
		DependencyIndexes: file_proto_EvtBeingHitInfo_proto_depIdxs,
		MessageInfos:      file_proto_EvtBeingHitInfo_proto_msgTypes,
	}.Build()
	File_proto_EvtBeingHitInfo_proto = out.File
	file_proto_EvtBeingHitInfo_proto_rawDesc = nil
	file_proto_EvtBeingHitInfo_proto_goTypes = nil
	file_proto_EvtBeingHitInfo_proto_depIdxs = nil
}