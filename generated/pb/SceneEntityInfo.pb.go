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
// source: proto/SceneEntityInfo.proto

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

// Obf: DGPPPPBJILO
type SceneEntityInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EntityType       ProtEntityType                    `protobuf:"varint,1,opt,name=entity_type,json=entityType,proto3,enum=ProtEntityType" json:"entity_type,omitempty"`
	EntityId         uint32                            `protobuf:"varint,2,opt,name=entity_id,json=entityId,proto3" json:"entity_id,omitempty"`
	Name             string                            `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	MotionInfo       *MotionInfo                       `protobuf:"bytes,4,opt,name=motion_info,json=motionInfo,proto3" json:"motion_info,omitempty"`
	PropList         []*PropPair                       `protobuf:"bytes,5,rep,name=prop_list,json=propList,proto3" json:"prop_list,omitempty"`
	FightPropList    []*FightPropPair                  `protobuf:"bytes,6,rep,name=fight_prop_list,json=fightPropList,proto3" json:"fight_prop_list,omitempty"`
	LifeState        uint32                            `protobuf:"varint,7,opt,name=life_state,json=lifeState,proto3" json:"life_state,omitempty"`
	AnimatorParaList []*AnimatorParameterValueInfoPair `protobuf:"bytes,9,rep,name=animator_para_list,json=animatorParaList,proto3" json:"animator_para_list,omitempty"`
	// Types that are assignable to Entity:
	//
	//	*SceneEntityInfo_Avatar
	//	*SceneEntityInfo_Monster
	//	*SceneEntityInfo_Npc
	//	*SceneEntityInfo_Gadget
	Entity                    isSceneEntityInfo_Entity `protobuf_oneof:"entity"`
	LastMoveSceneTimeMs       uint32                   `protobuf:"varint,17,opt,name=last_move_scene_time_ms,json=lastMoveSceneTimeMs,proto3" json:"last_move_scene_time_ms,omitempty"`
	LastMoveReliableSeq       uint32                   `protobuf:"varint,18,opt,name=last_move_reliable_seq,json=lastMoveReliableSeq,proto3" json:"last_move_reliable_seq,omitempty"`
	EntityClientData          *EntityClientData        `protobuf:"bytes,19,opt,name=entity_client_data,json=entityClientData,proto3" json:"entity_client_data,omitempty"`
	EntityEnvironmentInfoList []*EntityEnvironmentInfo `protobuf:"bytes,20,rep,name=entity_environment_info_list,json=entityEnvironmentInfoList,proto3" json:"entity_environment_info_list,omitempty"`
	EntityAuthorityInfo       *EntityAuthorityInfo     `protobuf:"bytes,21,opt,name=entity_authority_info,json=entityAuthorityInfo,proto3" json:"entity_authority_info,omitempty"`
	TagList                   []string                 `protobuf:"bytes,22,rep,name=tag_list,json=tagList,proto3" json:"tag_list,omitempty"`
	ServerBuffList            []*ServerBuff            `protobuf:"bytes,23,rep,name=server_buff_list,json=serverBuffList,proto3" json:"server_buff_list,omitempty"`
}

func (x *SceneEntityInfo) Reset() {
	*x = SceneEntityInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_SceneEntityInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SceneEntityInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SceneEntityInfo) ProtoMessage() {}

func (x *SceneEntityInfo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_SceneEntityInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SceneEntityInfo.ProtoReflect.Descriptor instead.
func (*SceneEntityInfo) Descriptor() ([]byte, []int) {
	return file_proto_SceneEntityInfo_proto_rawDescGZIP(), []int{0}
}

func (x *SceneEntityInfo) GetEntityType() ProtEntityType {
	if x != nil {
		return x.EntityType
	}
	return ProtEntityType_PROT_ENTITY_TYPE_NONE
}

func (x *SceneEntityInfo) GetEntityId() uint32 {
	if x != nil {
		return x.EntityId
	}
	return 0
}

func (x *SceneEntityInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SceneEntityInfo) GetMotionInfo() *MotionInfo {
	if x != nil {
		return x.MotionInfo
	}
	return nil
}

func (x *SceneEntityInfo) GetPropList() []*PropPair {
	if x != nil {
		return x.PropList
	}
	return nil
}

func (x *SceneEntityInfo) GetFightPropList() []*FightPropPair {
	if x != nil {
		return x.FightPropList
	}
	return nil
}

func (x *SceneEntityInfo) GetLifeState() uint32 {
	if x != nil {
		return x.LifeState
	}
	return 0
}

func (x *SceneEntityInfo) GetAnimatorParaList() []*AnimatorParameterValueInfoPair {
	if x != nil {
		return x.AnimatorParaList
	}
	return nil
}

func (m *SceneEntityInfo) GetEntity() isSceneEntityInfo_Entity {
	if m != nil {
		return m.Entity
	}
	return nil
}

func (x *SceneEntityInfo) GetAvatar() *SceneAvatarInfo {
	if x, ok := x.GetEntity().(*SceneEntityInfo_Avatar); ok {
		return x.Avatar
	}
	return nil
}

func (x *SceneEntityInfo) GetMonster() *SceneMonsterInfo {
	if x, ok := x.GetEntity().(*SceneEntityInfo_Monster); ok {
		return x.Monster
	}
	return nil
}

func (x *SceneEntityInfo) GetNpc() *SceneNpcInfo {
	if x, ok := x.GetEntity().(*SceneEntityInfo_Npc); ok {
		return x.Npc
	}
	return nil
}

func (x *SceneEntityInfo) GetGadget() *SceneGadgetInfo {
	if x, ok := x.GetEntity().(*SceneEntityInfo_Gadget); ok {
		return x.Gadget
	}
	return nil
}

func (x *SceneEntityInfo) GetLastMoveSceneTimeMs() uint32 {
	if x != nil {
		return x.LastMoveSceneTimeMs
	}
	return 0
}

func (x *SceneEntityInfo) GetLastMoveReliableSeq() uint32 {
	if x != nil {
		return x.LastMoveReliableSeq
	}
	return 0
}

func (x *SceneEntityInfo) GetEntityClientData() *EntityClientData {
	if x != nil {
		return x.EntityClientData
	}
	return nil
}

func (x *SceneEntityInfo) GetEntityEnvironmentInfoList() []*EntityEnvironmentInfo {
	if x != nil {
		return x.EntityEnvironmentInfoList
	}
	return nil
}

func (x *SceneEntityInfo) GetEntityAuthorityInfo() *EntityAuthorityInfo {
	if x != nil {
		return x.EntityAuthorityInfo
	}
	return nil
}

func (x *SceneEntityInfo) GetTagList() []string {
	if x != nil {
		return x.TagList
	}
	return nil
}

func (x *SceneEntityInfo) GetServerBuffList() []*ServerBuff {
	if x != nil {
		return x.ServerBuffList
	}
	return nil
}

type isSceneEntityInfo_Entity interface {
	isSceneEntityInfo_Entity()
}

type SceneEntityInfo_Avatar struct {
	Avatar *SceneAvatarInfo `protobuf:"bytes,10,opt,name=avatar,proto3,oneof"`
}

type SceneEntityInfo_Monster struct {
	Monster *SceneMonsterInfo `protobuf:"bytes,11,opt,name=monster,proto3,oneof"`
}

type SceneEntityInfo_Npc struct {
	Npc *SceneNpcInfo `protobuf:"bytes,12,opt,name=npc,proto3,oneof"`
}

type SceneEntityInfo_Gadget struct {
	Gadget *SceneGadgetInfo `protobuf:"bytes,13,opt,name=gadget,proto3,oneof"`
}

func (*SceneEntityInfo_Avatar) isSceneEntityInfo_Entity() {}

func (*SceneEntityInfo_Monster) isSceneEntityInfo_Entity() {}

func (*SceneEntityInfo_Npc) isSceneEntityInfo_Entity() {}

func (*SceneEntityInfo_Gadget) isSceneEntityInfo_Entity() {}

var File_proto_SceneEntityInfo_proto protoreflect.FileDescriptor

var file_proto_SceneEntityInfo_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x45, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x50, 0x72, 0x6f, 0x74, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x54,
	0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x4d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x50, 0x72, 0x6f, 0x70, 0x50, 0x61, 0x69,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x46,
	0x69, 0x67, 0x68, 0x74, 0x50, 0x72, 0x6f, 0x70, 0x50, 0x61, 0x69, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x2a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x41, 0x6e, 0x69, 0x6d, 0x61, 0x74,
	0x6f, 0x72, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x49, 0x6e, 0x66, 0x6f, 0x50, 0x61, 0x69, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x41, 0x76, 0x61, 0x74, 0x61,
	0x72, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x4d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x4e, 0x70, 0x63, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x53, 0x63, 0x65, 0x6e, 0x65,
	0x47, 0x61, 0x64, 0x67, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x43, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x45, 0x6e, 0x76, 0x69,
	0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x41,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x42, 0x75, 0x66, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc5, 0x07, 0x0a, 0x0f, 0x53,
	0x63, 0x65, 0x6e, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x30,
	0x0a, 0x0b, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x1b, 0x0a, 0x09, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x08, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x2c, 0x0a, 0x0b, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x6e, 0x66, 0x6f,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x4d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x0a, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x26, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x70, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x05, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x09, 0x2e, 0x50, 0x72, 0x6f, 0x70, 0x50, 0x61, 0x69, 0x72, 0x52, 0x08, 0x70,
	0x72, 0x6f, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x36, 0x0a, 0x0f, 0x66, 0x69, 0x67, 0x68, 0x74,
	0x5f, 0x70, 0x72, 0x6f, 0x70, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0e, 0x2e, 0x46, 0x69, 0x67, 0x68, 0x74, 0x50, 0x72, 0x6f, 0x70, 0x50, 0x61, 0x69, 0x72,
	0x52, 0x0d, 0x66, 0x69, 0x67, 0x68, 0x74, 0x50, 0x72, 0x6f, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x1d, 0x0a, 0x0a, 0x6c, 0x69, 0x66, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x09, 0x6c, 0x69, 0x66, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x4d,
	0x0a, 0x12, 0x61, 0x6e, 0x69, 0x6d, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x5f,
	0x6c, 0x69, 0x73, 0x74, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x41, 0x6e, 0x69,
	0x6d, 0x61, 0x74, 0x6f, 0x72, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x50, 0x61, 0x69, 0x72, 0x52, 0x10, 0x61, 0x6e, 0x69,
	0x6d, 0x61, 0x74, 0x6f, 0x72, 0x50, 0x61, 0x72, 0x61, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2a, 0x0a,
	0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e,
	0x53, 0x63, 0x65, 0x6e, 0x65, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x48,
	0x00, 0x52, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x12, 0x2d, 0x0a, 0x07, 0x6d, 0x6f, 0x6e,
	0x73, 0x74, 0x65, 0x72, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x53, 0x63, 0x65,
	0x6e, 0x65, 0x4d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x48, 0x00, 0x52,
	0x07, 0x6d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x12, 0x21, 0x0a, 0x03, 0x6e, 0x70, 0x63, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x4e, 0x70, 0x63,
	0x49, 0x6e, 0x66, 0x6f, 0x48, 0x00, 0x52, 0x03, 0x6e, 0x70, 0x63, 0x12, 0x2a, 0x0a, 0x06, 0x67,
	0x61, 0x64, 0x67, 0x65, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x53, 0x63,
	0x65, 0x6e, 0x65, 0x47, 0x61, 0x64, 0x67, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x48, 0x00, 0x52,
	0x06, 0x67, 0x61, 0x64, 0x67, 0x65, 0x74, 0x12, 0x34, 0x0a, 0x17, 0x6c, 0x61, 0x73, 0x74, 0x5f,
	0x6d, 0x6f, 0x76, 0x65, 0x5f, 0x73, 0x63, 0x65, 0x6e, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f,
	0x6d, 0x73, 0x18, 0x11, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x13, 0x6c, 0x61, 0x73, 0x74, 0x4d, 0x6f,
	0x76, 0x65, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x4d, 0x73, 0x12, 0x33, 0x0a,
	0x16, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6d, 0x6f, 0x76, 0x65, 0x5f, 0x72, 0x65, 0x6c, 0x69, 0x61,
	0x62, 0x6c, 0x65, 0x5f, 0x73, 0x65, 0x71, 0x18, 0x12, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x13, 0x6c,
	0x61, 0x73, 0x74, 0x4d, 0x6f, 0x76, 0x65, 0x52, 0x65, 0x6c, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x53,
	0x65, 0x71, 0x12, 0x3f, 0x0a, 0x12, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x13, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11,
	0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x44, 0x61, 0x74,
	0x61, 0x52, 0x10, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x44,
	0x61, 0x74, 0x61, 0x12, 0x57, 0x0a, 0x1c, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x65, 0x6e,
	0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x5f, 0x6c,
	0x69, 0x73, 0x74, 0x18, 0x14, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x45, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x19, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e,
	0x6d, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x48, 0x0a, 0x15,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79,
	0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x15, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x45, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x13, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69,
	0x74, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x61, 0x67, 0x5f, 0x6c, 0x69,
	0x73, 0x74, 0x18, 0x16, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x74, 0x61, 0x67, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x35, 0x0a, 0x10, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x62, 0x75, 0x66, 0x66,
	0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x17, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x42, 0x75, 0x66, 0x66, 0x52, 0x0e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x42, 0x75, 0x66, 0x66, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x08, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x42, 0x1d, 0x5a, 0x1b, 0x47, 0x6f, 0x2d, 0x47, 0x72, 0x61, 0x73, 0x73, 0x63, 0x75,
	0x74, 0x74, 0x65, 0x72, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_SceneEntityInfo_proto_rawDescOnce sync.Once
	file_proto_SceneEntityInfo_proto_rawDescData = file_proto_SceneEntityInfo_proto_rawDesc
)

func file_proto_SceneEntityInfo_proto_rawDescGZIP() []byte {
	file_proto_SceneEntityInfo_proto_rawDescOnce.Do(func() {
		file_proto_SceneEntityInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_SceneEntityInfo_proto_rawDescData)
	})
	return file_proto_SceneEntityInfo_proto_rawDescData
}

var file_proto_SceneEntityInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_SceneEntityInfo_proto_goTypes = []interface{}{
	(*SceneEntityInfo)(nil),                // 0: SceneEntityInfo
	(ProtEntityType)(0),                    // 1: ProtEntityType
	(*MotionInfo)(nil),                     // 2: MotionInfo
	(*PropPair)(nil),                       // 3: PropPair
	(*FightPropPair)(nil),                  // 4: FightPropPair
	(*AnimatorParameterValueInfoPair)(nil), // 5: AnimatorParameterValueInfoPair
	(*SceneAvatarInfo)(nil),                // 6: SceneAvatarInfo
	(*SceneMonsterInfo)(nil),               // 7: SceneMonsterInfo
	(*SceneNpcInfo)(nil),                   // 8: SceneNpcInfo
	(*SceneGadgetInfo)(nil),                // 9: SceneGadgetInfo
	(*EntityClientData)(nil),               // 10: EntityClientData
	(*EntityEnvironmentInfo)(nil),          // 11: EntityEnvironmentInfo
	(*EntityAuthorityInfo)(nil),            // 12: EntityAuthorityInfo
	(*ServerBuff)(nil),                     // 13: ServerBuff
}
var file_proto_SceneEntityInfo_proto_depIdxs = []int32{
	1,  // 0: SceneEntityInfo.entity_type:type_name -> ProtEntityType
	2,  // 1: SceneEntityInfo.motion_info:type_name -> MotionInfo
	3,  // 2: SceneEntityInfo.prop_list:type_name -> PropPair
	4,  // 3: SceneEntityInfo.fight_prop_list:type_name -> FightPropPair
	5,  // 4: SceneEntityInfo.animator_para_list:type_name -> AnimatorParameterValueInfoPair
	6,  // 5: SceneEntityInfo.avatar:type_name -> SceneAvatarInfo
	7,  // 6: SceneEntityInfo.monster:type_name -> SceneMonsterInfo
	8,  // 7: SceneEntityInfo.npc:type_name -> SceneNpcInfo
	9,  // 8: SceneEntityInfo.gadget:type_name -> SceneGadgetInfo
	10, // 9: SceneEntityInfo.entity_client_data:type_name -> EntityClientData
	11, // 10: SceneEntityInfo.entity_environment_info_list:type_name -> EntityEnvironmentInfo
	12, // 11: SceneEntityInfo.entity_authority_info:type_name -> EntityAuthorityInfo
	13, // 12: SceneEntityInfo.server_buff_list:type_name -> ServerBuff
	13, // [13:13] is the sub-list for method output_type
	13, // [13:13] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_proto_SceneEntityInfo_proto_init() }
func file_proto_SceneEntityInfo_proto_init() {
	if File_proto_SceneEntityInfo_proto != nil {
		return
	}
	file_proto_ProtEntityType_proto_init()
	file_proto_MotionInfo_proto_init()
	file_proto_PropPair_proto_init()
	file_proto_FightPropPair_proto_init()
	file_proto_AnimatorParameterValueInfoPair_proto_init()
	file_proto_SceneAvatarInfo_proto_init()
	file_proto_SceneMonsterInfo_proto_init()
	file_proto_SceneNpcInfo_proto_init()
	file_proto_SceneGadgetInfo_proto_init()
	file_proto_EntityClientData_proto_init()
	file_proto_EntityEnvironmentInfo_proto_init()
	file_proto_EntityAuthorityInfo_proto_init()
	file_proto_ServerBuff_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_SceneEntityInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SceneEntityInfo); i {
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
	file_proto_SceneEntityInfo_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*SceneEntityInfo_Avatar)(nil),
		(*SceneEntityInfo_Monster)(nil),
		(*SceneEntityInfo_Npc)(nil),
		(*SceneEntityInfo_Gadget)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_SceneEntityInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_SceneEntityInfo_proto_goTypes,
		DependencyIndexes: file_proto_SceneEntityInfo_proto_depIdxs,
		MessageInfos:      file_proto_SceneEntityInfo_proto_msgTypes,
	}.Build()
	File_proto_SceneEntityInfo_proto = out.File
	file_proto_SceneEntityInfo_proto_rawDesc = nil
	file_proto_SceneEntityInfo_proto_goTypes = nil
	file_proto_SceneEntityInfo_proto_depIdxs = nil
}