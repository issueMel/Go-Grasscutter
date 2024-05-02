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
// source: proto/AvatarInfo.proto

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

// Obf: FHPLJBCPFFF
type AvatarInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AvatarId                 uint32                      `protobuf:"varint,1,opt,name=avatar_id,json=avatarId,proto3" json:"avatar_id,omitempty"`
	Guid                     uint64                      `protobuf:"varint,2,opt,name=guid,proto3" json:"guid,omitempty"`
	PropMap                  map[uint32]*PropValue       `protobuf:"bytes,3,rep,name=prop_map,json=propMap,proto3" json:"prop_map,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	LifeState                uint32                      `protobuf:"varint,4,opt,name=life_state,json=lifeState,proto3" json:"life_state,omitempty"`
	EquipGuidList            []uint64                    `protobuf:"varint,5,rep,packed,name=equip_guid_list,json=equipGuidList,proto3" json:"equip_guid_list,omitempty"`
	TalentIdList             []uint32                    `protobuf:"varint,6,rep,packed,name=talent_id_list,json=talentIdList,proto3" json:"talent_id_list,omitempty"`
	FightPropMap             map[uint32]float32          `protobuf:"bytes,7,rep,name=fight_prop_map,json=fightPropMap,proto3" json:"fight_prop_map,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"fixed32,2,opt,name=value,proto3"`
	TrialAvatarInfo          *TrialAvatarInfo            `protobuf:"bytes,9,opt,name=trial_avatar_info,json=trialAvatarInfo,proto3" json:"trial_avatar_info,omitempty"`
	SkillMap                 map[uint32]*AvatarSkillInfo `protobuf:"bytes,10,rep,name=skill_map,json=skillMap,proto3" json:"skill_map,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	SkillDepotId             uint32                      `protobuf:"varint,11,opt,name=skill_depot_id,json=skillDepotId,proto3" json:"skill_depot_id,omitempty"`
	FetterInfo               *AvatarFetterInfo           `protobuf:"bytes,12,opt,name=fetter_info,json=fetterInfo,proto3" json:"fetter_info,omitempty"`
	CoreProudSkillLevel      uint32                      `protobuf:"varint,13,opt,name=core_proud_skill_level,json=coreProudSkillLevel,proto3" json:"core_proud_skill_level,omitempty"`
	InherentProudSkillList   []uint32                    `protobuf:"varint,14,rep,packed,name=inherent_proud_skill_list,json=inherentProudSkillList,proto3" json:"inherent_proud_skill_list,omitempty"`
	SkillLevelMap            map[uint32]uint32           `protobuf:"bytes,15,rep,name=skill_level_map,json=skillLevelMap,proto3" json:"skill_level_map,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	ExpeditionState          AvatarExpeditionState       `protobuf:"varint,16,opt,name=expedition_state,json=expeditionState,proto3,enum=AvatarExpeditionState" json:"expedition_state,omitempty"`
	ProudSkillExtraLevelMap  map[uint32]uint32           `protobuf:"bytes,17,rep,name=proud_skill_extra_level_map,json=proudSkillExtraLevelMap,proto3" json:"proud_skill_extra_level_map,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	IsFocus                  bool                        `protobuf:"varint,18,opt,name=is_focus,json=isFocus,proto3" json:"is_focus,omitempty"`
	AvatarType               uint32                      `protobuf:"varint,19,opt,name=avatar_type,json=avatarType,proto3" json:"avatar_type,omitempty"`
	TeamResonanceList        []uint32                    `protobuf:"varint,20,rep,packed,name=team_resonance_list,json=teamResonanceList,proto3" json:"team_resonance_list,omitempty"`
	WearingFlycloakId        uint32                      `protobuf:"varint,21,opt,name=wearing_flycloak_id,json=wearingFlycloakId,proto3" json:"wearing_flycloak_id,omitempty"`
	EquipAffixList           []*AvatarEquipAffixInfo     `protobuf:"bytes,22,rep,name=equip_affix_list,json=equipAffixList,proto3" json:"equip_affix_list,omitempty"`
	BornTime                 uint32                      `protobuf:"varint,23,opt,name=born_time,json=bornTime,proto3" json:"born_time,omitempty"`
	PendingPromoteRewardList []uint32                    `protobuf:"varint,24,rep,packed,name=pending_promote_reward_list,json=pendingPromoteRewardList,proto3" json:"pending_promote_reward_list,omitempty"`
	CostumeId                uint32                      `protobuf:"varint,25,opt,name=costume_id,json=costumeId,proto3" json:"costume_id,omitempty"`
	ExcelInfo                *AvatarExcelInfo            `protobuf:"bytes,26,opt,name=excel_info,json=excelInfo,proto3" json:"excel_info,omitempty"`
	AnimHash                 uint32                      `protobuf:"varint,27,opt,name=anim_hash,json=animHash,proto3" json:"anim_hash,omitempty"`
	PDGKJIIEPIO              *JCDPOCOOGCI                `protobuf:"bytes,28,opt,name=PDGKJIIEPIO,proto3" json:"PDGKJIIEPIO,omitempty"`
	GJFKDCJENJE              uint32                      `protobuf:"varint,29,opt,name=GJFKDCJENJE,proto3" json:"GJFKDCJENJE,omitempty"`
}

func (x *AvatarInfo) Reset() {
	*x = AvatarInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_AvatarInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AvatarInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AvatarInfo) ProtoMessage() {}

func (x *AvatarInfo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_AvatarInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AvatarInfo.ProtoReflect.Descriptor instead.
func (*AvatarInfo) Descriptor() ([]byte, []int) {
	return file_proto_AvatarInfo_proto_rawDescGZIP(), []int{0}
}

func (x *AvatarInfo) GetAvatarId() uint32 {
	if x != nil {
		return x.AvatarId
	}
	return 0
}

func (x *AvatarInfo) GetGuid() uint64 {
	if x != nil {
		return x.Guid
	}
	return 0
}

func (x *AvatarInfo) GetPropMap() map[uint32]*PropValue {
	if x != nil {
		return x.PropMap
	}
	return nil
}

func (x *AvatarInfo) GetLifeState() uint32 {
	if x != nil {
		return x.LifeState
	}
	return 0
}

func (x *AvatarInfo) GetEquipGuidList() []uint64 {
	if x != nil {
		return x.EquipGuidList
	}
	return nil
}

func (x *AvatarInfo) GetTalentIdList() []uint32 {
	if x != nil {
		return x.TalentIdList
	}
	return nil
}

func (x *AvatarInfo) GetFightPropMap() map[uint32]float32 {
	if x != nil {
		return x.FightPropMap
	}
	return nil
}

func (x *AvatarInfo) GetTrialAvatarInfo() *TrialAvatarInfo {
	if x != nil {
		return x.TrialAvatarInfo
	}
	return nil
}

func (x *AvatarInfo) GetSkillMap() map[uint32]*AvatarSkillInfo {
	if x != nil {
		return x.SkillMap
	}
	return nil
}

func (x *AvatarInfo) GetSkillDepotId() uint32 {
	if x != nil {
		return x.SkillDepotId
	}
	return 0
}

func (x *AvatarInfo) GetFetterInfo() *AvatarFetterInfo {
	if x != nil {
		return x.FetterInfo
	}
	return nil
}

func (x *AvatarInfo) GetCoreProudSkillLevel() uint32 {
	if x != nil {
		return x.CoreProudSkillLevel
	}
	return 0
}

func (x *AvatarInfo) GetInherentProudSkillList() []uint32 {
	if x != nil {
		return x.InherentProudSkillList
	}
	return nil
}

func (x *AvatarInfo) GetSkillLevelMap() map[uint32]uint32 {
	if x != nil {
		return x.SkillLevelMap
	}
	return nil
}

func (x *AvatarInfo) GetExpeditionState() AvatarExpeditionState {
	if x != nil {
		return x.ExpeditionState
	}
	return AvatarExpeditionState_AVATAR_EXPEDITION_NONE
}

func (x *AvatarInfo) GetProudSkillExtraLevelMap() map[uint32]uint32 {
	if x != nil {
		return x.ProudSkillExtraLevelMap
	}
	return nil
}

func (x *AvatarInfo) GetIsFocus() bool {
	if x != nil {
		return x.IsFocus
	}
	return false
}

func (x *AvatarInfo) GetAvatarType() uint32 {
	if x != nil {
		return x.AvatarType
	}
	return 0
}

func (x *AvatarInfo) GetTeamResonanceList() []uint32 {
	if x != nil {
		return x.TeamResonanceList
	}
	return nil
}

func (x *AvatarInfo) GetWearingFlycloakId() uint32 {
	if x != nil {
		return x.WearingFlycloakId
	}
	return 0
}

func (x *AvatarInfo) GetEquipAffixList() []*AvatarEquipAffixInfo {
	if x != nil {
		return x.EquipAffixList
	}
	return nil
}

func (x *AvatarInfo) GetBornTime() uint32 {
	if x != nil {
		return x.BornTime
	}
	return 0
}

func (x *AvatarInfo) GetPendingPromoteRewardList() []uint32 {
	if x != nil {
		return x.PendingPromoteRewardList
	}
	return nil
}

func (x *AvatarInfo) GetCostumeId() uint32 {
	if x != nil {
		return x.CostumeId
	}
	return 0
}

func (x *AvatarInfo) GetExcelInfo() *AvatarExcelInfo {
	if x != nil {
		return x.ExcelInfo
	}
	return nil
}

func (x *AvatarInfo) GetAnimHash() uint32 {
	if x != nil {
		return x.AnimHash
	}
	return 0
}

func (x *AvatarInfo) GetPDGKJIIEPIO() *JCDPOCOOGCI {
	if x != nil {
		return x.PDGKJIIEPIO
	}
	return nil
}

func (x *AvatarInfo) GetGJFKDCJENJE() uint32 {
	if x != nil {
		return x.GJFKDCJENJE
	}
	return 0
}

var File_proto_AvatarInfo_proto protoreflect.FileDescriptor

var file_proto_AvatarInfo_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x50, 0x72, 0x6f, 0x70, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x54, 0x72, 0x69, 0x61, 0x6c, 0x41, 0x76, 0x61, 0x74,
	0x61, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x49,
	0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x46, 0x65, 0x74, 0x74, 0x65, 0x72, 0x49, 0x6e, 0x66,
	0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x41,
	0x76, 0x61, 0x74, 0x61, 0x72, 0x45, 0x78, 0x70, 0x65, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x45, 0x71, 0x75, 0x69, 0x70, 0x41, 0x66, 0x66,
	0x69, 0x78, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x45, 0x78, 0x63, 0x65, 0x6c, 0x49,
	0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x4a, 0x43, 0x44, 0x50, 0x4f, 0x43, 0x4f, 0x4f, 0x47, 0x43, 0x49, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xb5, 0x0d, 0x0a, 0x0a, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x49, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x67, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x67, 0x75,
	0x69, 0x64, 0x12, 0x33, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x70, 0x5f, 0x6d, 0x61, 0x70, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x49, 0x6e, 0x66,
	0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x70, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07,
	0x70, 0x72, 0x6f, 0x70, 0x4d, 0x61, 0x70, 0x12, 0x1d, 0x0a, 0x0a, 0x6c, 0x69, 0x66, 0x65, 0x5f,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x6c, 0x69, 0x66,
	0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x26, 0x0a, 0x0f, 0x65, 0x71, 0x75, 0x69, 0x70, 0x5f,
	0x67, 0x75, 0x69, 0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x05, 0x20, 0x03, 0x28, 0x04, 0x52,
	0x0d, 0x65, 0x71, 0x75, 0x69, 0x70, 0x47, 0x75, 0x69, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x24,
	0x0a, 0x0e, 0x74, 0x61, 0x6c, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74,
	0x18, 0x06, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0c, 0x74, 0x61, 0x6c, 0x65, 0x6e, 0x74, 0x49, 0x64,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x43, 0x0a, 0x0e, 0x66, 0x69, 0x67, 0x68, 0x74, 0x5f, 0x70, 0x72,
	0x6f, 0x70, 0x5f, 0x6d, 0x61, 0x70, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x41,
	0x76, 0x61, 0x74, 0x61, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x46, 0x69, 0x67, 0x68, 0x74, 0x50,
	0x72, 0x6f, 0x70, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0c, 0x66, 0x69, 0x67,
	0x68, 0x74, 0x50, 0x72, 0x6f, 0x70, 0x4d, 0x61, 0x70, 0x12, 0x3c, 0x0a, 0x11, 0x74, 0x72, 0x69,
	0x61, 0x6c, 0x5f, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x54, 0x72, 0x69, 0x61, 0x6c, 0x41, 0x76, 0x61, 0x74,
	0x61, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0f, 0x74, 0x72, 0x69, 0x61, 0x6c, 0x41, 0x76, 0x61,
	0x74, 0x61, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x36, 0x0a, 0x09, 0x73, 0x6b, 0x69, 0x6c, 0x6c,
	0x5f, 0x6d, 0x61, 0x70, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x41, 0x76, 0x61,
	0x74, 0x61, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x4d, 0x61, 0x70,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x4d, 0x61, 0x70, 0x12,
	0x24, 0x0a, 0x0e, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x5f, 0x64, 0x65, 0x70, 0x6f, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x44, 0x65,
	0x70, 0x6f, 0x74, 0x49, 0x64, 0x12, 0x32, 0x0a, 0x0b, 0x66, 0x65, 0x74, 0x74, 0x65, 0x72, 0x5f,
	0x69, 0x6e, 0x66, 0x6f, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x41, 0x76, 0x61,
	0x74, 0x61, 0x72, 0x46, 0x65, 0x74, 0x74, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0a, 0x66,
	0x65, 0x74, 0x74, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x33, 0x0a, 0x16, 0x63, 0x6f, 0x72,
	0x65, 0x5f, 0x70, 0x72, 0x6f, 0x75, 0x64, 0x5f, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x5f, 0x6c, 0x65,
	0x76, 0x65, 0x6c, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x13, 0x63, 0x6f, 0x72, 0x65, 0x50,
	0x72, 0x6f, 0x75, 0x64, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x39,
	0x0a, 0x19, 0x69, 0x6e, 0x68, 0x65, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x75, 0x64,
	0x5f, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0e, 0x20, 0x03, 0x28,
	0x0d, 0x52, 0x16, 0x69, 0x6e, 0x68, 0x65, 0x72, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x75, 0x64,
	0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x46, 0x0a, 0x0f, 0x73, 0x6b, 0x69,
	0x6c, 0x6c, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x5f, 0x6d, 0x61, 0x70, 0x18, 0x0f, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x2e,
	0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x0d, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x4d, 0x61,
	0x70, 0x12, 0x41, 0x0a, 0x10, 0x65, 0x78, 0x70, 0x65, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x41, 0x76,
	0x61, 0x74, 0x61, 0x72, 0x45, 0x78, 0x70, 0x65, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x52, 0x0f, 0x65, 0x78, 0x70, 0x65, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x12, 0x66, 0x0a, 0x1b, 0x70, 0x72, 0x6f, 0x75, 0x64, 0x5f, 0x73, 0x6b,
	0x69, 0x6c, 0x6c, 0x5f, 0x65, 0x78, 0x74, 0x72, 0x61, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x5f,
	0x6d, 0x61, 0x70, 0x18, 0x11, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x41, 0x76, 0x61, 0x74,
	0x61, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x75, 0x64, 0x53, 0x6b, 0x69, 0x6c,
	0x6c, 0x45, 0x78, 0x74, 0x72, 0x61, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x4d, 0x61, 0x70, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x17, 0x70, 0x72, 0x6f, 0x75, 0x64, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x45,
	0x78, 0x74, 0x72, 0x61, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x4d, 0x61, 0x70, 0x12, 0x19, 0x0a, 0x08,
	0x69, 0x73, 0x5f, 0x66, 0x6f, 0x63, 0x75, 0x73, 0x18, 0x12, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07,
	0x69, 0x73, 0x46, 0x6f, 0x63, 0x75, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x76, 0x61, 0x74, 0x61,
	0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x13, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x61, 0x76,
	0x61, 0x74, 0x61, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2e, 0x0a, 0x13, 0x74, 0x65, 0x61, 0x6d,
	0x5f, 0x72, 0x65, 0x73, 0x6f, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18,
	0x14, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x11, 0x74, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x6f, 0x6e,
	0x61, 0x6e, 0x63, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x13, 0x77, 0x65, 0x61, 0x72,
	0x69, 0x6e, 0x67, 0x5f, 0x66, 0x6c, 0x79, 0x63, 0x6c, 0x6f, 0x61, 0x6b, 0x5f, 0x69, 0x64, 0x18,
	0x15, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x11, 0x77, 0x65, 0x61, 0x72, 0x69, 0x6e, 0x67, 0x46, 0x6c,
	0x79, 0x63, 0x6c, 0x6f, 0x61, 0x6b, 0x49, 0x64, 0x12, 0x3f, 0x0a, 0x10, 0x65, 0x71, 0x75, 0x69,
	0x70, 0x5f, 0x61, 0x66, 0x66, 0x69, 0x78, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x16, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x15, 0x2e, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x45, 0x71, 0x75, 0x69, 0x70,
	0x41, 0x66, 0x66, 0x69, 0x78, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0e, 0x65, 0x71, 0x75, 0x69, 0x70,
	0x41, 0x66, 0x66, 0x69, 0x78, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x62, 0x6f, 0x72,
	0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x17, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x62, 0x6f,
	0x72, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x3d, 0x0a, 0x1b, 0x70, 0x65, 0x6e, 0x64, 0x69, 0x6e,
	0x67, 0x5f, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x65, 0x5f, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64,
	0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x18, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x18, 0x70, 0x65, 0x6e,
	0x64, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x77, 0x61, 0x72,
	0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x73, 0x74, 0x75, 0x6d, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x19, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x63, 0x6f, 0x73, 0x74, 0x75,
	0x6d, 0x65, 0x49, 0x64, 0x12, 0x2f, 0x0a, 0x0a, 0x65, 0x78, 0x63, 0x65, 0x6c, 0x5f, 0x69, 0x6e,
	0x66, 0x6f, 0x18, 0x1a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x41, 0x76, 0x61, 0x74, 0x61,
	0x72, 0x45, 0x78, 0x63, 0x65, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x09, 0x65, 0x78, 0x63, 0x65,
	0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x6e, 0x69, 0x6d, 0x5f, 0x68, 0x61,
	0x73, 0x68, 0x18, 0x1b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x61, 0x6e, 0x69, 0x6d, 0x48, 0x61,
	0x73, 0x68, 0x12, 0x2e, 0x0a, 0x0b, 0x50, 0x44, 0x47, 0x4b, 0x4a, 0x49, 0x49, 0x45, 0x50, 0x49,
	0x4f, 0x18, 0x1c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x4a, 0x43, 0x44, 0x50, 0x4f, 0x43,
	0x4f, 0x4f, 0x47, 0x43, 0x49, 0x52, 0x0b, 0x50, 0x44, 0x47, 0x4b, 0x4a, 0x49, 0x49, 0x45, 0x50,
	0x49, 0x4f, 0x12, 0x20, 0x0a, 0x0b, 0x47, 0x4a, 0x46, 0x4b, 0x44, 0x43, 0x4a, 0x45, 0x4e, 0x4a,
	0x45, 0x18, 0x1d, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x47, 0x4a, 0x46, 0x4b, 0x44, 0x43, 0x4a,
	0x45, 0x4e, 0x4a, 0x45, 0x1a, 0x46, 0x0a, 0x0c, 0x50, 0x72, 0x6f, 0x70, 0x4d, 0x61, 0x70, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x20, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x50, 0x72, 0x6f, 0x70, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x3f, 0x0a, 0x11,
	0x46, 0x69, 0x67, 0x68, 0x74, 0x50, 0x72, 0x6f, 0x70, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x4d, 0x0a,
	0x0d, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x26, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x10, 0x2e, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x40, 0x0a, 0x12,
	0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x4a,
	0x0a, 0x1c, 0x50, 0x72, 0x6f, 0x75, 0x64, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x45, 0x78, 0x74, 0x72,
	0x61, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x1d, 0x5a, 0x1b, 0x47, 0x6f,
	0x2d, 0x47, 0x72, 0x61, 0x73, 0x73, 0x63, 0x75, 0x74, 0x74, 0x65, 0x72, 0x2f, 0x67, 0x65, 0x6e,
	0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_proto_AvatarInfo_proto_rawDescOnce sync.Once
	file_proto_AvatarInfo_proto_rawDescData = file_proto_AvatarInfo_proto_rawDesc
)

func file_proto_AvatarInfo_proto_rawDescGZIP() []byte {
	file_proto_AvatarInfo_proto_rawDescOnce.Do(func() {
		file_proto_AvatarInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_AvatarInfo_proto_rawDescData)
	})
	return file_proto_AvatarInfo_proto_rawDescData
}

var file_proto_AvatarInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_proto_AvatarInfo_proto_goTypes = []interface{}{
	(*AvatarInfo)(nil),           // 0: AvatarInfo
	nil,                          // 1: AvatarInfo.PropMapEntry
	nil,                          // 2: AvatarInfo.FightPropMapEntry
	nil,                          // 3: AvatarInfo.SkillMapEntry
	nil,                          // 4: AvatarInfo.SkillLevelMapEntry
	nil,                          // 5: AvatarInfo.ProudSkillExtraLevelMapEntry
	(*TrialAvatarInfo)(nil),      // 6: TrialAvatarInfo
	(*AvatarFetterInfo)(nil),     // 7: AvatarFetterInfo
	(AvatarExpeditionState)(0),   // 8: AvatarExpeditionState
	(*AvatarEquipAffixInfo)(nil), // 9: AvatarEquipAffixInfo
	(*AvatarExcelInfo)(nil),      // 10: AvatarExcelInfo
	(*JCDPOCOOGCI)(nil),          // 11: JCDPOCOOGCI
	(*PropValue)(nil),            // 12: PropValue
	(*AvatarSkillInfo)(nil),      // 13: AvatarSkillInfo
}
var file_proto_AvatarInfo_proto_depIdxs = []int32{
	1,  // 0: AvatarInfo.prop_map:type_name -> AvatarInfo.PropMapEntry
	2,  // 1: AvatarInfo.fight_prop_map:type_name -> AvatarInfo.FightPropMapEntry
	6,  // 2: AvatarInfo.trial_avatar_info:type_name -> TrialAvatarInfo
	3,  // 3: AvatarInfo.skill_map:type_name -> AvatarInfo.SkillMapEntry
	7,  // 4: AvatarInfo.fetter_info:type_name -> AvatarFetterInfo
	4,  // 5: AvatarInfo.skill_level_map:type_name -> AvatarInfo.SkillLevelMapEntry
	8,  // 6: AvatarInfo.expedition_state:type_name -> AvatarExpeditionState
	5,  // 7: AvatarInfo.proud_skill_extra_level_map:type_name -> AvatarInfo.ProudSkillExtraLevelMapEntry
	9,  // 8: AvatarInfo.equip_affix_list:type_name -> AvatarEquipAffixInfo
	10, // 9: AvatarInfo.excel_info:type_name -> AvatarExcelInfo
	11, // 10: AvatarInfo.PDGKJIIEPIO:type_name -> JCDPOCOOGCI
	12, // 11: AvatarInfo.PropMapEntry.value:type_name -> PropValue
	13, // 12: AvatarInfo.SkillMapEntry.value:type_name -> AvatarSkillInfo
	13, // [13:13] is the sub-list for method output_type
	13, // [13:13] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_proto_AvatarInfo_proto_init() }
func file_proto_AvatarInfo_proto_init() {
	if File_proto_AvatarInfo_proto != nil {
		return
	}
	file_proto_PropValue_proto_init()
	file_proto_TrialAvatarInfo_proto_init()
	file_proto_AvatarSkillInfo_proto_init()
	file_proto_AvatarFetterInfo_proto_init()
	file_proto_AvatarExpeditionState_proto_init()
	file_proto_AvatarEquipAffixInfo_proto_init()
	file_proto_AvatarExcelInfo_proto_init()
	file_proto_JCDPOCOOGCI_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_AvatarInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AvatarInfo); i {
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
			RawDescriptor: file_proto_AvatarInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_AvatarInfo_proto_goTypes,
		DependencyIndexes: file_proto_AvatarInfo_proto_depIdxs,
		MessageInfos:      file_proto_AvatarInfo_proto_msgTypes,
	}.Build()
	File_proto_AvatarInfo_proto = out.File
	file_proto_AvatarInfo_proto_rawDesc = nil
	file_proto_AvatarInfo_proto_goTypes = nil
	file_proto_AvatarInfo_proto_depIdxs = nil
}
