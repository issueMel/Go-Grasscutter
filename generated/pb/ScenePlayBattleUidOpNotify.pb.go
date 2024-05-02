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
// source: proto/ScenePlayBattleUidOpNotify.proto

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

// CmdId: 602
// Obf: FOEFOFMLBAI
type ScenePlayBattleUidOpNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayType        uint32   `protobuf:"varint,8,opt,name=play_type,json=playType,proto3" json:"play_type,omitempty"`
	ParamStr        string   `protobuf:"bytes,15,opt,name=param_str,json=paramStr,proto3" json:"param_str,omitempty"`
	ParamDuration   uint32   `protobuf:"varint,9,opt,name=param_duration,json=paramDuration,proto3" json:"param_duration,omitempty"`
	Op              uint32   `protobuf:"varint,5,opt,name=op,proto3" json:"op,omitempty"`
	EntityId        uint32   `protobuf:"varint,2,opt,name=entity_id,json=entityId,proto3" json:"entity_id,omitempty"`
	ParamIndex      uint32   `protobuf:"varint,4,opt,name=param_index,json=paramIndex,proto3" json:"param_index,omitempty"`
	UidList         []uint32 `protobuf:"varint,14,rep,packed,name=uid_list,json=uidList,proto3" json:"uid_list,omitempty"`
	ParamList       []uint32 `protobuf:"varint,7,rep,packed,name=param_list,json=paramList,proto3" json:"param_list,omitempty"`
	ParamTargetList []uint32 `protobuf:"varint,12,rep,packed,name=param_target_list,json=paramTargetList,proto3" json:"param_target_list,omitempty"`
	PlayId          uint32   `protobuf:"varint,3,opt,name=play_id,json=playId,proto3" json:"play_id,omitempty"`
}

func (x *ScenePlayBattleUidOpNotify) Reset() {
	*x = ScenePlayBattleUidOpNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ScenePlayBattleUidOpNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScenePlayBattleUidOpNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScenePlayBattleUidOpNotify) ProtoMessage() {}

func (x *ScenePlayBattleUidOpNotify) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ScenePlayBattleUidOpNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScenePlayBattleUidOpNotify.ProtoReflect.Descriptor instead.
func (*ScenePlayBattleUidOpNotify) Descriptor() ([]byte, []int) {
	return file_proto_ScenePlayBattleUidOpNotify_proto_rawDescGZIP(), []int{0}
}

func (x *ScenePlayBattleUidOpNotify) GetPlayType() uint32 {
	if x != nil {
		return x.PlayType
	}
	return 0
}

func (x *ScenePlayBattleUidOpNotify) GetParamStr() string {
	if x != nil {
		return x.ParamStr
	}
	return ""
}

func (x *ScenePlayBattleUidOpNotify) GetParamDuration() uint32 {
	if x != nil {
		return x.ParamDuration
	}
	return 0
}

func (x *ScenePlayBattleUidOpNotify) GetOp() uint32 {
	if x != nil {
		return x.Op
	}
	return 0
}

func (x *ScenePlayBattleUidOpNotify) GetEntityId() uint32 {
	if x != nil {
		return x.EntityId
	}
	return 0
}

func (x *ScenePlayBattleUidOpNotify) GetParamIndex() uint32 {
	if x != nil {
		return x.ParamIndex
	}
	return 0
}

func (x *ScenePlayBattleUidOpNotify) GetUidList() []uint32 {
	if x != nil {
		return x.UidList
	}
	return nil
}

func (x *ScenePlayBattleUidOpNotify) GetParamList() []uint32 {
	if x != nil {
		return x.ParamList
	}
	return nil
}

func (x *ScenePlayBattleUidOpNotify) GetParamTargetList() []uint32 {
	if x != nil {
		return x.ParamTargetList
	}
	return nil
}

func (x *ScenePlayBattleUidOpNotify) GetPlayId() uint32 {
	if x != nil {
		return x.PlayId
	}
	return 0
}

var File_proto_ScenePlayBattleUidOpNotify_proto protoreflect.FileDescriptor

var file_proto_ScenePlayBattleUidOpNotify_proto_rawDesc = []byte{
	0x0a, 0x26, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x50, 0x6c, 0x61,
	0x79, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x55, 0x69, 0x64, 0x4f, 0x70, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xca, 0x02, 0x0a, 0x1a, 0x53, 0x63, 0x65,
	0x6e, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x55, 0x69, 0x64, 0x4f,
	0x70, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x6c, 0x61, 0x79, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x79,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x5f, 0x73, 0x74,
	0x72, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x53, 0x74,
	0x72, 0x12, 0x25, 0x0a, 0x0e, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x5f, 0x64, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x70, 0x61, 0x72, 0x61, 0x6d,
	0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x70, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x6f, 0x70, 0x12, 0x1b, 0x0a, 0x09, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x5f, 0x69,
	0x6e, 0x64, 0x65, 0x78, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x70, 0x61, 0x72, 0x61,
	0x6d, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x19, 0x0a, 0x08, 0x75, 0x69, 0x64, 0x5f, 0x6c, 0x69,
	0x73, 0x74, 0x18, 0x0e, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x07, 0x75, 0x69, 0x64, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18,
	0x07, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x09, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x2a, 0x0a, 0x11, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x5f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0f, 0x70, 0x61, 0x72,
	0x61, 0x6d, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07,
	0x70, 0x6c, 0x61, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x70,
	0x6c, 0x61, 0x79, 0x49, 0x64, 0x42, 0x1d, 0x5a, 0x1b, 0x47, 0x6f, 0x2d, 0x47, 0x72, 0x61, 0x73,
	0x73, 0x63, 0x75, 0x74, 0x74, 0x65, 0x72, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65,
	0x64, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_ScenePlayBattleUidOpNotify_proto_rawDescOnce sync.Once
	file_proto_ScenePlayBattleUidOpNotify_proto_rawDescData = file_proto_ScenePlayBattleUidOpNotify_proto_rawDesc
)

func file_proto_ScenePlayBattleUidOpNotify_proto_rawDescGZIP() []byte {
	file_proto_ScenePlayBattleUidOpNotify_proto_rawDescOnce.Do(func() {
		file_proto_ScenePlayBattleUidOpNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_ScenePlayBattleUidOpNotify_proto_rawDescData)
	})
	return file_proto_ScenePlayBattleUidOpNotify_proto_rawDescData
}

var file_proto_ScenePlayBattleUidOpNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_ScenePlayBattleUidOpNotify_proto_goTypes = []interface{}{
	(*ScenePlayBattleUidOpNotify)(nil), // 0: ScenePlayBattleUidOpNotify
}
var file_proto_ScenePlayBattleUidOpNotify_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_ScenePlayBattleUidOpNotify_proto_init() }
func file_proto_ScenePlayBattleUidOpNotify_proto_init() {
	if File_proto_ScenePlayBattleUidOpNotify_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_ScenePlayBattleUidOpNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScenePlayBattleUidOpNotify); i {
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
			RawDescriptor: file_proto_ScenePlayBattleUidOpNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_ScenePlayBattleUidOpNotify_proto_goTypes,
		DependencyIndexes: file_proto_ScenePlayBattleUidOpNotify_proto_depIdxs,
		MessageInfos:      file_proto_ScenePlayBattleUidOpNotify_proto_msgTypes,
	}.Build()
	File_proto_ScenePlayBattleUidOpNotify_proto = out.File
	file_proto_ScenePlayBattleUidOpNotify_proto_rawDesc = nil
	file_proto_ScenePlayBattleUidOpNotify_proto_goTypes = nil
	file_proto_ScenePlayBattleUidOpNotify_proto_depIdxs = nil
}
