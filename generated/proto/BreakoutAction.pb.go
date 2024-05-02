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
// source: proto/BreakoutAction.proto

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

// Obf: BGLHJHMLHAI
type BreakoutAction_BreakoutActionType int32

const (
	BreakoutAction_ACTION_TYPE_NONE           BreakoutAction_BreakoutActionType = 0
	BreakoutAction_ACTION_TYPE_LAUNCH_BALL    BreakoutAction_BreakoutActionType = 1
	BreakoutAction_ACTION_TYPE_DESTROY_BALL   BreakoutAction_BreakoutActionType = 2
	BreakoutAction_ACTION_TYPE_FALLING_OBJECT BreakoutAction_BreakoutActionType = 3
	BreakoutAction_ACTION_TYPE_MISSILE        BreakoutAction_BreakoutActionType = 4
)

// Enum value maps for BreakoutAction_BreakoutActionType.
var (
	BreakoutAction_BreakoutActionType_name = map[int32]string{
		0: "ACTION_TYPE_NONE",
		1: "ACTION_TYPE_LAUNCH_BALL",
		2: "ACTION_TYPE_DESTROY_BALL",
		3: "ACTION_TYPE_FALLING_OBJECT",
		4: "ACTION_TYPE_MISSILE",
	}
	BreakoutAction_BreakoutActionType_value = map[string]int32{
		"ACTION_TYPE_NONE":           0,
		"ACTION_TYPE_LAUNCH_BALL":    1,
		"ACTION_TYPE_DESTROY_BALL":   2,
		"ACTION_TYPE_FALLING_OBJECT": 3,
		"ACTION_TYPE_MISSILE":        4,
	}
)

func (x BreakoutAction_BreakoutActionType) Enum() *BreakoutAction_BreakoutActionType {
	p := new(BreakoutAction_BreakoutActionType)
	*p = x
	return p
}

func (x BreakoutAction_BreakoutActionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BreakoutAction_BreakoutActionType) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_BreakoutAction_proto_enumTypes[0].Descriptor()
}

func (BreakoutAction_BreakoutActionType) Type() protoreflect.EnumType {
	return &file_proto_BreakoutAction_proto_enumTypes[0]
}

func (x BreakoutAction_BreakoutActionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BreakoutAction_BreakoutActionType.Descriptor instead.
func (BreakoutAction_BreakoutActionType) EnumDescriptor() ([]byte, []int) {
	return file_proto_BreakoutAction_proto_rawDescGZIP(), []int{0, 0}
}

// Obf: JBHHKOPAEFB
type BreakoutAction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ActionType          BreakoutAction_BreakoutActionType `protobuf:"varint,1,opt,name=action_type,json=actionType,proto3,enum=BreakoutAction_BreakoutActionType" json:"action_type,omitempty"`
	ClientGameTime      uint64                            `protobuf:"varint,2,opt,name=client_game_time,json=clientGameTime,proto3" json:"client_game_time,omitempty"`
	ServerGameTime      uint64                            `protobuf:"varint,3,opt,name=server_game_time,json=serverGameTime,proto3" json:"server_game_time,omitempty"`
	IsFailed            bool                              `protobuf:"varint,4,opt,name=is_failed,json=isFailed,proto3" json:"is_failed,omitempty"`
	PreIndex            uint32                            `protobuf:"varint,5,opt,name=pre_index,json=preIndex,proto3" json:"pre_index,omitempty"`
	NewIndex            uint32                            `protobuf:"varint,6,opt,name=new_index,json=newIndex,proto3" json:"new_index,omitempty"`
	Pos                 *BreakoutVector2                  `protobuf:"bytes,7,opt,name=pos,proto3" json:"pos,omitempty"`
	MoveDir             *BreakoutVector2                  `protobuf:"bytes,8,opt,name=move_dir,json=moveDir,proto3" json:"move_dir,omitempty"`
	Speed               int32                             `protobuf:"varint,9,opt,name=speed,proto3" json:"speed,omitempty"`
	PeerId              uint32                            `protobuf:"varint,10,opt,name=peer_id,json=peerId,proto3" json:"peer_id,omitempty"`
	ElementType         uint32                            `protobuf:"varint,11,opt,name=element_type,json=elementType,proto3" json:"element_type,omitempty"`
	ElementReactionBuff uint32                            `protobuf:"varint,12,opt,name=element_reaction_buff,json=elementReactionBuff,proto3" json:"element_reaction_buff,omitempty"`
	SpeedIncreaseCount  uint32                            `protobuf:"varint,13,opt,name=speed_increase_count,json=speedIncreaseCount,proto3" json:"speed_increase_count,omitempty"`
	HasExtraBall        bool                              `protobuf:"varint,14,opt,name=has_extra_ball,json=hasExtraBall,proto3" json:"has_extra_ball,omitempty"`
	ExtraBallDir        *BreakoutVector2                  `protobuf:"bytes,15,opt,name=extra_ball_dir,json=extraBallDir,proto3" json:"extra_ball_dir,omitempty"`
	ExtraBallIndex      uint32                            `protobuf:"varint,16,opt,name=extra_ball_index,json=extraBallIndex,proto3" json:"extra_ball_index,omitempty"`
	Offset              int32                             `protobuf:"varint,17,opt,name=offset,proto3" json:"offset,omitempty"`
	FCCNGNCIFAI         uint64                            `protobuf:"varint,18,opt,name=FCCNGNCIFAI,proto3" json:"FCCNGNCIFAI,omitempty"`
}

func (x *BreakoutAction) Reset() {
	*x = BreakoutAction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_BreakoutAction_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BreakoutAction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BreakoutAction) ProtoMessage() {}

func (x *BreakoutAction) ProtoReflect() protoreflect.Message {
	mi := &file_proto_BreakoutAction_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BreakoutAction.ProtoReflect.Descriptor instead.
func (*BreakoutAction) Descriptor() ([]byte, []int) {
	return file_proto_BreakoutAction_proto_rawDescGZIP(), []int{0}
}

func (x *BreakoutAction) GetActionType() BreakoutAction_BreakoutActionType {
	if x != nil {
		return x.ActionType
	}
	return BreakoutAction_ACTION_TYPE_NONE
}

func (x *BreakoutAction) GetClientGameTime() uint64 {
	if x != nil {
		return x.ClientGameTime
	}
	return 0
}

func (x *BreakoutAction) GetServerGameTime() uint64 {
	if x != nil {
		return x.ServerGameTime
	}
	return 0
}

func (x *BreakoutAction) GetIsFailed() bool {
	if x != nil {
		return x.IsFailed
	}
	return false
}

func (x *BreakoutAction) GetPreIndex() uint32 {
	if x != nil {
		return x.PreIndex
	}
	return 0
}

func (x *BreakoutAction) GetNewIndex() uint32 {
	if x != nil {
		return x.NewIndex
	}
	return 0
}

func (x *BreakoutAction) GetPos() *BreakoutVector2 {
	if x != nil {
		return x.Pos
	}
	return nil
}

func (x *BreakoutAction) GetMoveDir() *BreakoutVector2 {
	if x != nil {
		return x.MoveDir
	}
	return nil
}

func (x *BreakoutAction) GetSpeed() int32 {
	if x != nil {
		return x.Speed
	}
	return 0
}

func (x *BreakoutAction) GetPeerId() uint32 {
	if x != nil {
		return x.PeerId
	}
	return 0
}

func (x *BreakoutAction) GetElementType() uint32 {
	if x != nil {
		return x.ElementType
	}
	return 0
}

func (x *BreakoutAction) GetElementReactionBuff() uint32 {
	if x != nil {
		return x.ElementReactionBuff
	}
	return 0
}

func (x *BreakoutAction) GetSpeedIncreaseCount() uint32 {
	if x != nil {
		return x.SpeedIncreaseCount
	}
	return 0
}

func (x *BreakoutAction) GetHasExtraBall() bool {
	if x != nil {
		return x.HasExtraBall
	}
	return false
}

func (x *BreakoutAction) GetExtraBallDir() *BreakoutVector2 {
	if x != nil {
		return x.ExtraBallDir
	}
	return nil
}

func (x *BreakoutAction) GetExtraBallIndex() uint32 {
	if x != nil {
		return x.ExtraBallIndex
	}
	return 0
}

func (x *BreakoutAction) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *BreakoutAction) GetFCCNGNCIFAI() uint64 {
	if x != nil {
		return x.FCCNGNCIFAI
	}
	return 0
}

var File_proto_BreakoutAction_proto protoreflect.FileDescriptor

var file_proto_BreakoutAction_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x42, 0x72, 0x65, 0x61, 0x6b, 0x6f, 0x75, 0x74,
	0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x42, 0x72, 0x65, 0x61, 0x6b, 0x6f, 0x75, 0x74, 0x56, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x32, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xec, 0x06, 0x0a, 0x0e, 0x42, 0x72,
	0x65, 0x61, 0x6b, 0x6f, 0x75, 0x74, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x43, 0x0a, 0x0b,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x22, 0x2e, 0x42, 0x72, 0x65, 0x61, 0x6b, 0x6f, 0x75, 0x74, 0x41, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x42, 0x72, 0x65, 0x61, 0x6b, 0x6f, 0x75, 0x74, 0x41, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x28, 0x0a, 0x10, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x67, 0x61, 0x6d, 0x65,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0e, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x47, 0x61, 0x6d, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x28, 0x0a, 0x10, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x47, 0x61, 0x6d,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x66, 0x61, 0x69, 0x6c,
	0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x46, 0x61, 0x69, 0x6c,
	0x65, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x72, 0x65, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x70, 0x72, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12,
	0x1b, 0x0a, 0x09, 0x6e, 0x65, 0x77, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x08, 0x6e, 0x65, 0x77, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x22, 0x0a, 0x03,
	0x70, 0x6f, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x42, 0x72, 0x65, 0x61,
	0x6b, 0x6f, 0x75, 0x74, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x32, 0x52, 0x03, 0x70, 0x6f, 0x73,
	0x12, 0x2b, 0x0a, 0x08, 0x6d, 0x6f, 0x76, 0x65, 0x5f, 0x64, 0x69, 0x72, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x10, 0x2e, 0x42, 0x72, 0x65, 0x61, 0x6b, 0x6f, 0x75, 0x74, 0x56, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x32, 0x52, 0x07, 0x6d, 0x6f, 0x76, 0x65, 0x44, 0x69, 0x72, 0x12, 0x14, 0x0a,
	0x05, 0x73, 0x70, 0x65, 0x65, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x73, 0x70,
	0x65, 0x65, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x65, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x70, 0x65, 0x65, 0x72, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c,
	0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0b, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x32, 0x0a, 0x15, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x72, 0x65, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x62, 0x75, 0x66, 0x66, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x13,
	0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42,
	0x75, 0x66, 0x66, 0x12, 0x30, 0x0a, 0x14, 0x73, 0x70, 0x65, 0x65, 0x64, 0x5f, 0x69, 0x6e, 0x63,
	0x72, 0x65, 0x61, 0x73, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x12, 0x73, 0x70, 0x65, 0x65, 0x64, 0x49, 0x6e, 0x63, 0x72, 0x65, 0x61, 0x73, 0x65,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x24, 0x0a, 0x0e, 0x68, 0x61, 0x73, 0x5f, 0x65, 0x78, 0x74,
	0x72, 0x61, 0x5f, 0x62, 0x61, 0x6c, 0x6c, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x68,
	0x61, 0x73, 0x45, 0x78, 0x74, 0x72, 0x61, 0x42, 0x61, 0x6c, 0x6c, 0x12, 0x36, 0x0a, 0x0e, 0x65,
	0x78, 0x74, 0x72, 0x61, 0x5f, 0x62, 0x61, 0x6c, 0x6c, 0x5f, 0x64, 0x69, 0x72, 0x18, 0x0f, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x42, 0x72, 0x65, 0x61, 0x6b, 0x6f, 0x75, 0x74, 0x56, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x32, 0x52, 0x0c, 0x65, 0x78, 0x74, 0x72, 0x61, 0x42, 0x61, 0x6c, 0x6c,
	0x44, 0x69, 0x72, 0x12, 0x28, 0x0a, 0x10, 0x65, 0x78, 0x74, 0x72, 0x61, 0x5f, 0x62, 0x61, 0x6c,
	0x6c, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x65,
	0x78, 0x74, 0x72, 0x61, 0x42, 0x61, 0x6c, 0x6c, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x16, 0x0a,
	0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x11, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f,
	0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x46, 0x43, 0x43, 0x4e, 0x47, 0x4e, 0x43,
	0x49, 0x46, 0x41, 0x49, 0x18, 0x12, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x46, 0x43, 0x43, 0x4e,
	0x47, 0x4e, 0x43, 0x49, 0x46, 0x41, 0x49, 0x22, 0x9e, 0x01, 0x0a, 0x12, 0x42, 0x72, 0x65, 0x61,
	0x6b, 0x6f, 0x75, 0x74, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14,
	0x0a, 0x10, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4e, 0x4f,
	0x4e, 0x45, 0x10, 0x00, 0x12, 0x1b, 0x0a, 0x17, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x4c, 0x41, 0x55, 0x4e, 0x43, 0x48, 0x5f, 0x42, 0x41, 0x4c, 0x4c, 0x10,
	0x01, 0x12, 0x1c, 0x0a, 0x18, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x44, 0x45, 0x53, 0x54, 0x52, 0x4f, 0x59, 0x5f, 0x42, 0x41, 0x4c, 0x4c, 0x10, 0x02, 0x12,
	0x1e, 0x0a, 0x1a, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x46,
	0x41, 0x4c, 0x4c, 0x49, 0x4e, 0x47, 0x5f, 0x4f, 0x42, 0x4a, 0x45, 0x43, 0x54, 0x10, 0x03, 0x12,
	0x17, 0x0a, 0x13, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4d,
	0x49, 0x53, 0x53, 0x49, 0x4c, 0x45, 0x10, 0x04, 0x42, 0x20, 0x5a, 0x1e, 0x47, 0x6f, 0x2d, 0x47,
	0x72, 0x61, 0x73, 0x73, 0x63, 0x75, 0x74, 0x74, 0x65, 0x72, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x74, 0x65, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_proto_BreakoutAction_proto_rawDescOnce sync.Once
	file_proto_BreakoutAction_proto_rawDescData = file_proto_BreakoutAction_proto_rawDesc
)

func file_proto_BreakoutAction_proto_rawDescGZIP() []byte {
	file_proto_BreakoutAction_proto_rawDescOnce.Do(func() {
		file_proto_BreakoutAction_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_BreakoutAction_proto_rawDescData)
	})
	return file_proto_BreakoutAction_proto_rawDescData
}

var file_proto_BreakoutAction_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_BreakoutAction_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_BreakoutAction_proto_goTypes = []interface{}{
	(BreakoutAction_BreakoutActionType)(0), // 0: BreakoutAction.BreakoutActionType
	(*BreakoutAction)(nil),                 // 1: BreakoutAction
	(*BreakoutVector2)(nil),                // 2: BreakoutVector2
}
var file_proto_BreakoutAction_proto_depIdxs = []int32{
	0, // 0: BreakoutAction.action_type:type_name -> BreakoutAction.BreakoutActionType
	2, // 1: BreakoutAction.pos:type_name -> BreakoutVector2
	2, // 2: BreakoutAction.move_dir:type_name -> BreakoutVector2
	2, // 3: BreakoutAction.extra_ball_dir:type_name -> BreakoutVector2
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_proto_BreakoutAction_proto_init() }
func file_proto_BreakoutAction_proto_init() {
	if File_proto_BreakoutAction_proto != nil {
		return
	}
	file_proto_BreakoutVector2_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_BreakoutAction_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BreakoutAction); i {
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
			RawDescriptor: file_proto_BreakoutAction_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_BreakoutAction_proto_goTypes,
		DependencyIndexes: file_proto_BreakoutAction_proto_depIdxs,
		EnumInfos:         file_proto_BreakoutAction_proto_enumTypes,
		MessageInfos:      file_proto_BreakoutAction_proto_msgTypes,
	}.Build()
	File_proto_BreakoutAction_proto = out.File
	file_proto_BreakoutAction_proto_rawDesc = nil
	file_proto_BreakoutAction_proto_goTypes = nil
	file_proto_BreakoutAction_proto_depIdxs = nil
}
