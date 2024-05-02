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
// source: proto/AbilityMixinWindSeedSpawner.proto

package proto

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

// Obf: GODBLNPFIIG
type AbilityMixinWindSeedSpawner struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Cmd:
	//
	//	*AbilityMixinWindSeedSpawner_AddSignal_
	//	*AbilityMixinWindSeedSpawner_RefreshSeed_
	//	*AbilityMixinWindSeedSpawner_CatchSeed_
	Cmd isAbilityMixinWindSeedSpawner_Cmd `protobuf_oneof:"cmd"`
}

func (x *AbilityMixinWindSeedSpawner) Reset() {
	*x = AbilityMixinWindSeedSpawner{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_AbilityMixinWindSeedSpawner_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AbilityMixinWindSeedSpawner) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AbilityMixinWindSeedSpawner) ProtoMessage() {}

func (x *AbilityMixinWindSeedSpawner) ProtoReflect() protoreflect.Message {
	mi := &file_proto_AbilityMixinWindSeedSpawner_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AbilityMixinWindSeedSpawner.ProtoReflect.Descriptor instead.
func (*AbilityMixinWindSeedSpawner) Descriptor() ([]byte, []int) {
	return file_proto_AbilityMixinWindSeedSpawner_proto_rawDescGZIP(), []int{0}
}

func (m *AbilityMixinWindSeedSpawner) GetCmd() isAbilityMixinWindSeedSpawner_Cmd {
	if m != nil {
		return m.Cmd
	}
	return nil
}

func (x *AbilityMixinWindSeedSpawner) GetAddSignal() *AbilityMixinWindSeedSpawner_AddSignal {
	if x, ok := x.GetCmd().(*AbilityMixinWindSeedSpawner_AddSignal_); ok {
		return x.AddSignal
	}
	return nil
}

func (x *AbilityMixinWindSeedSpawner) GetRefreshSeed() *AbilityMixinWindSeedSpawner_RefreshSeed {
	if x, ok := x.GetCmd().(*AbilityMixinWindSeedSpawner_RefreshSeed_); ok {
		return x.RefreshSeed
	}
	return nil
}

func (x *AbilityMixinWindSeedSpawner) GetCatchSeed() *AbilityMixinWindSeedSpawner_CatchSeed {
	if x, ok := x.GetCmd().(*AbilityMixinWindSeedSpawner_CatchSeed_); ok {
		return x.CatchSeed
	}
	return nil
}

type isAbilityMixinWindSeedSpawner_Cmd interface {
	isAbilityMixinWindSeedSpawner_Cmd()
}

type AbilityMixinWindSeedSpawner_AddSignal_ struct {
	AddSignal *AbilityMixinWindSeedSpawner_AddSignal `protobuf:"bytes,2,opt,name=add_signal,json=addSignal,proto3,oneof"`
}

type AbilityMixinWindSeedSpawner_RefreshSeed_ struct {
	RefreshSeed *AbilityMixinWindSeedSpawner_RefreshSeed `protobuf:"bytes,8,opt,name=refresh_seed,json=refreshSeed,proto3,oneof"`
}

type AbilityMixinWindSeedSpawner_CatchSeed_ struct {
	CatchSeed *AbilityMixinWindSeedSpawner_CatchSeed `protobuf:"bytes,15,opt,name=catch_seed,json=catchSeed,proto3,oneof"`
}

func (*AbilityMixinWindSeedSpawner_AddSignal_) isAbilityMixinWindSeedSpawner_Cmd() {}

func (*AbilityMixinWindSeedSpawner_RefreshSeed_) isAbilityMixinWindSeedSpawner_Cmd() {}

func (*AbilityMixinWindSeedSpawner_CatchSeed_) isAbilityMixinWindSeedSpawner_Cmd() {}

// Obf: KOLEALFBLKF
type AbilityMixinWindSeedSpawner_AddSignal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AbilityMixinWindSeedSpawner_AddSignal) Reset() {
	*x = AbilityMixinWindSeedSpawner_AddSignal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_AbilityMixinWindSeedSpawner_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AbilityMixinWindSeedSpawner_AddSignal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AbilityMixinWindSeedSpawner_AddSignal) ProtoMessage() {}

func (x *AbilityMixinWindSeedSpawner_AddSignal) ProtoReflect() protoreflect.Message {
	mi := &file_proto_AbilityMixinWindSeedSpawner_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AbilityMixinWindSeedSpawner_AddSignal.ProtoReflect.Descriptor instead.
func (*AbilityMixinWindSeedSpawner_AddSignal) Descriptor() ([]byte, []int) {
	return file_proto_AbilityMixinWindSeedSpawner_proto_rawDescGZIP(), []int{0, 0}
}

// Obf: PKIEIPCDKJN
type AbilityMixinWindSeedSpawner_RefreshSeed struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PosList []*Vector `protobuf:"bytes,9,rep,name=pos_list,json=posList,proto3" json:"pos_list,omitempty"`
}

func (x *AbilityMixinWindSeedSpawner_RefreshSeed) Reset() {
	*x = AbilityMixinWindSeedSpawner_RefreshSeed{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_AbilityMixinWindSeedSpawner_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AbilityMixinWindSeedSpawner_RefreshSeed) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AbilityMixinWindSeedSpawner_RefreshSeed) ProtoMessage() {}

func (x *AbilityMixinWindSeedSpawner_RefreshSeed) ProtoReflect() protoreflect.Message {
	mi := &file_proto_AbilityMixinWindSeedSpawner_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AbilityMixinWindSeedSpawner_RefreshSeed.ProtoReflect.Descriptor instead.
func (*AbilityMixinWindSeedSpawner_RefreshSeed) Descriptor() ([]byte, []int) {
	return file_proto_AbilityMixinWindSeedSpawner_proto_rawDescGZIP(), []int{0, 1}
}

func (x *AbilityMixinWindSeedSpawner_RefreshSeed) GetPosList() []*Vector {
	if x != nil {
		return x.PosList
	}
	return nil
}

// Obf: CMGICLJOGAM
type AbilityMixinWindSeedSpawner_CatchSeed struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EntityId uint32 `protobuf:"varint,9,opt,name=entity_id,json=entityId,proto3" json:"entity_id,omitempty"`
}

func (x *AbilityMixinWindSeedSpawner_CatchSeed) Reset() {
	*x = AbilityMixinWindSeedSpawner_CatchSeed{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_AbilityMixinWindSeedSpawner_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AbilityMixinWindSeedSpawner_CatchSeed) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AbilityMixinWindSeedSpawner_CatchSeed) ProtoMessage() {}

func (x *AbilityMixinWindSeedSpawner_CatchSeed) ProtoReflect() protoreflect.Message {
	mi := &file_proto_AbilityMixinWindSeedSpawner_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AbilityMixinWindSeedSpawner_CatchSeed.ProtoReflect.Descriptor instead.
func (*AbilityMixinWindSeedSpawner_CatchSeed) Descriptor() ([]byte, []int) {
	return file_proto_AbilityMixinWindSeedSpawner_proto_rawDescGZIP(), []int{0, 2}
}

func (x *AbilityMixinWindSeedSpawner_CatchSeed) GetEntityId() uint32 {
	if x != nil {
		return x.EntityId
	}
	return 0
}

var File_proto_AbilityMixinWindSeedSpawner_proto protoreflect.FileDescriptor

var file_proto_AbilityMixinWindSeedSpawner_proto_rawDesc = []byte{
	0x0a, 0x27, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x41, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x4d,
	0x69, 0x78, 0x69, 0x6e, 0x57, 0x69, 0x6e, 0x64, 0x53, 0x65, 0x65, 0x64, 0x53, 0x70, 0x61, 0x77,
	0x6e, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xef, 0x02,
	0x0a, 0x1b, 0x41, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x4d, 0x69, 0x78, 0x69, 0x6e, 0x57, 0x69,
	0x6e, 0x64, 0x53, 0x65, 0x65, 0x64, 0x53, 0x70, 0x61, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x47, 0x0a,
	0x0a, 0x61, 0x64, 0x64, 0x5f, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x26, 0x2e, 0x41, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x4d, 0x69, 0x78, 0x69, 0x6e,
	0x57, 0x69, 0x6e, 0x64, 0x53, 0x65, 0x65, 0x64, 0x53, 0x70, 0x61, 0x77, 0x6e, 0x65, 0x72, 0x2e,
	0x41, 0x64, 0x64, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x48, 0x00, 0x52, 0x09, 0x61, 0x64, 0x64,
	0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x12, 0x4d, 0x0a, 0x0c, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73,
	0x68, 0x5f, 0x73, 0x65, 0x65, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x41,
	0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x4d, 0x69, 0x78, 0x69, 0x6e, 0x57, 0x69, 0x6e, 0x64, 0x53,
	0x65, 0x65, 0x64, 0x53, 0x70, 0x61, 0x77, 0x6e, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x66, 0x72, 0x65,
	0x73, 0x68, 0x53, 0x65, 0x65, 0x64, 0x48, 0x00, 0x52, 0x0b, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73,
	0x68, 0x53, 0x65, 0x65, 0x64, 0x12, 0x47, 0x0a, 0x0a, 0x63, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x73,
	0x65, 0x65, 0x64, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x41, 0x62, 0x69, 0x6c,
	0x69, 0x74, 0x79, 0x4d, 0x69, 0x78, 0x69, 0x6e, 0x57, 0x69, 0x6e, 0x64, 0x53, 0x65, 0x65, 0x64,
	0x53, 0x70, 0x61, 0x77, 0x6e, 0x65, 0x72, 0x2e, 0x43, 0x61, 0x74, 0x63, 0x68, 0x53, 0x65, 0x65,
	0x64, 0x48, 0x00, 0x52, 0x09, 0x63, 0x61, 0x74, 0x63, 0x68, 0x53, 0x65, 0x65, 0x64, 0x1a, 0x0b,
	0x0a, 0x09, 0x41, 0x64, 0x64, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x1a, 0x31, 0x0a, 0x0b, 0x52,
	0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x53, 0x65, 0x65, 0x64, 0x12, 0x22, 0x0a, 0x08, 0x70, 0x6f,
	0x73, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x56,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x07, 0x70, 0x6f, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x1a, 0x28,
	0x0a, 0x09, 0x43, 0x61, 0x74, 0x63, 0x68, 0x53, 0x65, 0x65, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x64, 0x42, 0x05, 0x0a, 0x03, 0x63, 0x6d, 0x64, 0x42,
	0x20, 0x5a, 0x1e, 0x47, 0x6f, 0x2d, 0x47, 0x72, 0x61, 0x73, 0x73, 0x63, 0x75, 0x74, 0x74, 0x65,
	0x72, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_AbilityMixinWindSeedSpawner_proto_rawDescOnce sync.Once
	file_proto_AbilityMixinWindSeedSpawner_proto_rawDescData = file_proto_AbilityMixinWindSeedSpawner_proto_rawDesc
)

func file_proto_AbilityMixinWindSeedSpawner_proto_rawDescGZIP() []byte {
	file_proto_AbilityMixinWindSeedSpawner_proto_rawDescOnce.Do(func() {
		file_proto_AbilityMixinWindSeedSpawner_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_AbilityMixinWindSeedSpawner_proto_rawDescData)
	})
	return file_proto_AbilityMixinWindSeedSpawner_proto_rawDescData
}

var file_proto_AbilityMixinWindSeedSpawner_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_AbilityMixinWindSeedSpawner_proto_goTypes = []interface{}{
	(*AbilityMixinWindSeedSpawner)(nil),             // 0: AbilityMixinWindSeedSpawner
	(*AbilityMixinWindSeedSpawner_AddSignal)(nil),   // 1: AbilityMixinWindSeedSpawner.AddSignal
	(*AbilityMixinWindSeedSpawner_RefreshSeed)(nil), // 2: AbilityMixinWindSeedSpawner.RefreshSeed
	(*AbilityMixinWindSeedSpawner_CatchSeed)(nil),   // 3: AbilityMixinWindSeedSpawner.CatchSeed
	(*Vector)(nil), // 4: Vector
}
var file_proto_AbilityMixinWindSeedSpawner_proto_depIdxs = []int32{
	1, // 0: AbilityMixinWindSeedSpawner.add_signal:type_name -> AbilityMixinWindSeedSpawner.AddSignal
	2, // 1: AbilityMixinWindSeedSpawner.refresh_seed:type_name -> AbilityMixinWindSeedSpawner.RefreshSeed
	3, // 2: AbilityMixinWindSeedSpawner.catch_seed:type_name -> AbilityMixinWindSeedSpawner.CatchSeed
	4, // 3: AbilityMixinWindSeedSpawner.RefreshSeed.pos_list:type_name -> Vector
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_proto_AbilityMixinWindSeedSpawner_proto_init() }
func file_proto_AbilityMixinWindSeedSpawner_proto_init() {
	if File_proto_AbilityMixinWindSeedSpawner_proto != nil {
		return
	}
	file_proto_Vector_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_AbilityMixinWindSeedSpawner_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AbilityMixinWindSeedSpawner); i {
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
		file_proto_AbilityMixinWindSeedSpawner_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AbilityMixinWindSeedSpawner_AddSignal); i {
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
		file_proto_AbilityMixinWindSeedSpawner_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AbilityMixinWindSeedSpawner_RefreshSeed); i {
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
		file_proto_AbilityMixinWindSeedSpawner_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AbilityMixinWindSeedSpawner_CatchSeed); i {
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
	file_proto_AbilityMixinWindSeedSpawner_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*AbilityMixinWindSeedSpawner_AddSignal_)(nil),
		(*AbilityMixinWindSeedSpawner_RefreshSeed_)(nil),
		(*AbilityMixinWindSeedSpawner_CatchSeed_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_AbilityMixinWindSeedSpawner_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_AbilityMixinWindSeedSpawner_proto_goTypes,
		DependencyIndexes: file_proto_AbilityMixinWindSeedSpawner_proto_depIdxs,
		MessageInfos:      file_proto_AbilityMixinWindSeedSpawner_proto_msgTypes,
	}.Build()
	File_proto_AbilityMixinWindSeedSpawner_proto = out.File
	file_proto_AbilityMixinWindSeedSpawner_proto_rawDesc = nil
	file_proto_AbilityMixinWindSeedSpawner_proto_goTypes = nil
	file_proto_AbilityMixinWindSeedSpawner_proto_depIdxs = nil
}
