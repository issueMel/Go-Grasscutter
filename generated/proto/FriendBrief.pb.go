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
// source: proto/FriendBrief.proto

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

// Obf: JIFKJKFNBFF
type FriendBrief struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid                   uint32                  `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Nickname              string                  `protobuf:"bytes,2,opt,name=nickname,proto3" json:"nickname,omitempty"`
	Level                 uint32                  `protobuf:"varint,3,opt,name=level,proto3" json:"level,omitempty"`
	AvatarId              uint32                  `protobuf:"varint,4,opt,name=avatar_id,json=avatarId,proto3" json:"avatar_id,omitempty"`
	WorldLevel            uint32                  `protobuf:"varint,5,opt,name=world_level,json=worldLevel,proto3" json:"world_level,omitempty"`
	Signature             string                  `protobuf:"bytes,6,opt,name=signature,proto3" json:"signature,omitempty"`
	OnlineState           FriendOnlineState       `protobuf:"varint,7,opt,name=online_state,json=onlineState,proto3,enum=FriendOnlineState" json:"online_state,omitempty"`
	Param                 uint32                  `protobuf:"varint,8,opt,name=param,proto3" json:"param,omitempty"`
	IsMpModeAvailable     bool                    `protobuf:"varint,10,opt,name=is_mp_mode_available,json=isMpModeAvailable,proto3" json:"is_mp_mode_available,omitempty"`
	OnlineId              string                  `protobuf:"bytes,11,opt,name=online_id,json=onlineId,proto3" json:"online_id,omitempty"`
	LastActiveTime        uint32                  `protobuf:"varint,12,opt,name=last_active_time,json=lastActiveTime,proto3" json:"last_active_time,omitempty"`
	NameCardId            uint32                  `protobuf:"varint,13,opt,name=name_card_id,json=nameCardId,proto3" json:"name_card_id,omitempty"`
	MpPlayerNum           uint32                  `protobuf:"varint,14,opt,name=mp_player_num,json=mpPlayerNum,proto3" json:"mp_player_num,omitempty"`
	IsChatNoDisturb       bool                    `protobuf:"varint,15,opt,name=is_chat_no_disturb,json=isChatNoDisturb,proto3" json:"is_chat_no_disturb,omitempty"`
	ChatSequence          uint32                  `protobuf:"varint,16,opt,name=chat_sequence,json=chatSequence,proto3" json:"chat_sequence,omitempty"`
	RemarkName            string                  `protobuf:"bytes,17,opt,name=remark_name,json=remarkName,proto3" json:"remark_name,omitempty"`
	ShowAvatarInfoList    []*SocialShowAvatarInfo `protobuf:"bytes,22,rep,name=show_avatar_info_list,json=showAvatarInfoList,proto3" json:"show_avatar_info_list,omitempty"`
	FriendEnterHomeOption FriendEnterHomeOption   `protobuf:"varint,23,opt,name=friend_enter_home_option,json=friendEnterHomeOption,proto3,enum=FriendEnterHomeOption" json:"friend_enter_home_option,omitempty"`
	ProfilePicture        *ProfilePicture         `protobuf:"bytes,24,opt,name=profile_picture,json=profilePicture,proto3" json:"profile_picture,omitempty"`
	IsGameSource          bool                    `protobuf:"varint,25,opt,name=is_game_source,json=isGameSource,proto3" json:"is_game_source,omitempty"`
	IsPsnSource           bool                    `protobuf:"varint,26,opt,name=is_psn_source,json=isPsnSource,proto3" json:"is_psn_source,omitempty"`
	PlatformType          PlatformType            `protobuf:"varint,27,opt,name=platform_type,json=platformType,proto3,enum=PlatformType" json:"platform_type,omitempty"`
	IEAHDCLDOEJ           bool                    `protobuf:"varint,28,opt,name=IEAHDCLDOEJ,proto3" json:"IEAHDCLDOEJ,omitempty"`
	BJFJJMGENCH           bool                    `protobuf:"varint,29,opt,name=BJFJJMGENCH,proto3" json:"BJFJJMGENCH,omitempty"`
}

func (x *FriendBrief) Reset() {
	*x = FriendBrief{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_FriendBrief_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FriendBrief) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FriendBrief) ProtoMessage() {}

func (x *FriendBrief) ProtoReflect() protoreflect.Message {
	mi := &file_proto_FriendBrief_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FriendBrief.ProtoReflect.Descriptor instead.
func (*FriendBrief) Descriptor() ([]byte, []int) {
	return file_proto_FriendBrief_proto_rawDescGZIP(), []int{0}
}

func (x *FriendBrief) GetUid() uint32 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *FriendBrief) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *FriendBrief) GetLevel() uint32 {
	if x != nil {
		return x.Level
	}
	return 0
}

func (x *FriendBrief) GetAvatarId() uint32 {
	if x != nil {
		return x.AvatarId
	}
	return 0
}

func (x *FriendBrief) GetWorldLevel() uint32 {
	if x != nil {
		return x.WorldLevel
	}
	return 0
}

func (x *FriendBrief) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

func (x *FriendBrief) GetOnlineState() FriendOnlineState {
	if x != nil {
		return x.OnlineState
	}
	return FriendOnlineState_FRIEND_ONLINE_STATE_DISCONNECT
}

func (x *FriendBrief) GetParam() uint32 {
	if x != nil {
		return x.Param
	}
	return 0
}

func (x *FriendBrief) GetIsMpModeAvailable() bool {
	if x != nil {
		return x.IsMpModeAvailable
	}
	return false
}

func (x *FriendBrief) GetOnlineId() string {
	if x != nil {
		return x.OnlineId
	}
	return ""
}

func (x *FriendBrief) GetLastActiveTime() uint32 {
	if x != nil {
		return x.LastActiveTime
	}
	return 0
}

func (x *FriendBrief) GetNameCardId() uint32 {
	if x != nil {
		return x.NameCardId
	}
	return 0
}

func (x *FriendBrief) GetMpPlayerNum() uint32 {
	if x != nil {
		return x.MpPlayerNum
	}
	return 0
}

func (x *FriendBrief) GetIsChatNoDisturb() bool {
	if x != nil {
		return x.IsChatNoDisturb
	}
	return false
}

func (x *FriendBrief) GetChatSequence() uint32 {
	if x != nil {
		return x.ChatSequence
	}
	return 0
}

func (x *FriendBrief) GetRemarkName() string {
	if x != nil {
		return x.RemarkName
	}
	return ""
}

func (x *FriendBrief) GetShowAvatarInfoList() []*SocialShowAvatarInfo {
	if x != nil {
		return x.ShowAvatarInfoList
	}
	return nil
}

func (x *FriendBrief) GetFriendEnterHomeOption() FriendEnterHomeOption {
	if x != nil {
		return x.FriendEnterHomeOption
	}
	return FriendEnterHomeOption_FRIEND_ENTER_HOME_OPTION_NEED_CONFIRM
}

func (x *FriendBrief) GetProfilePicture() *ProfilePicture {
	if x != nil {
		return x.ProfilePicture
	}
	return nil
}

func (x *FriendBrief) GetIsGameSource() bool {
	if x != nil {
		return x.IsGameSource
	}
	return false
}

func (x *FriendBrief) GetIsPsnSource() bool {
	if x != nil {
		return x.IsPsnSource
	}
	return false
}

func (x *FriendBrief) GetPlatformType() PlatformType {
	if x != nil {
		return x.PlatformType
	}
	return PlatformType_PLATFORM_TYPE_EDITOR
}

func (x *FriendBrief) GetIEAHDCLDOEJ() bool {
	if x != nil {
		return x.IEAHDCLDOEJ
	}
	return false
}

func (x *FriendBrief) GetBJFJJMGENCH() bool {
	if x != nil {
		return x.BJFJJMGENCH
	}
	return false
}

var File_proto_FriendBrief_proto protoreflect.FileDescriptor

var file_proto_FriendBrief_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x42, 0x72,
	0x69, 0x65, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x53, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x53, 0x68, 0x6f, 0x77, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72,
	0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x48, 0x6f, 0x6d,
	0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x69, 0x63, 0x74,
	0x75, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xc2, 0x07, 0x0a, 0x0b, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x42, 0x72,
	0x69, 0x65, 0x66, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x76, 0x61, 0x74, 0x61,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x61, 0x76, 0x61, 0x74,
	0x61, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x5f, 0x6c, 0x65,
	0x76, 0x65, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x77, 0x6f, 0x72, 0x6c, 0x64,
	0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x12, 0x35, 0x0a, 0x0c, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x46, 0x72, 0x69, 0x65,
	0x6e, 0x64, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x0b, 0x6f,
	0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61,
	0x72, 0x61, 0x6d, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x70, 0x61, 0x72, 0x61, 0x6d,
	0x12, 0x2f, 0x0a, 0x14, 0x69, 0x73, 0x5f, 0x6d, 0x70, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x5f, 0x61,
	0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11,
	0x69, 0x73, 0x4d, 0x70, 0x4d, 0x6f, 0x64, 0x65, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c,
	0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x49, 0x64, 0x12, 0x28,
	0x0a, 0x10, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x6c, 0x61, 0x73, 0x74, 0x41, 0x63,
	0x74, 0x69, 0x76, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0c, 0x6e, 0x61, 0x6d, 0x65,
	0x5f, 0x63, 0x61, 0x72, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a,
	0x6e, 0x61, 0x6d, 0x65, 0x43, 0x61, 0x72, 0x64, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0d, 0x6d, 0x70,
	0x5f, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x0e, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0b, 0x6d, 0x70, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4e, 0x75, 0x6d, 0x12, 0x2b,
	0x0a, 0x12, 0x69, 0x73, 0x5f, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x6e, 0x6f, 0x5f, 0x64, 0x69, 0x73,
	0x74, 0x75, 0x72, 0x62, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x69, 0x73, 0x43, 0x68,
	0x61, 0x74, 0x4e, 0x6f, 0x44, 0x69, 0x73, 0x74, 0x75, 0x72, 0x62, 0x12, 0x23, 0x0a, 0x0d, 0x63,
	0x68, 0x61, 0x74, 0x5f, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x10, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0c, 0x63, 0x68, 0x61, 0x74, 0x53, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65,
	0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x48, 0x0a, 0x15, 0x73, 0x68, 0x6f, 0x77, 0x5f, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72,
	0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x16, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x15, 0x2e, 0x53, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x53, 0x68, 0x6f, 0x77, 0x41, 0x76, 0x61,
	0x74, 0x61, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x12, 0x73, 0x68, 0x6f, 0x77, 0x41, 0x76, 0x61,
	0x74, 0x61, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x4f, 0x0a, 0x18, 0x66,
	0x72, 0x69, 0x65, 0x6e, 0x64, 0x5f, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x5f, 0x68, 0x6f, 0x6d, 0x65,
	0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x17, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e,
	0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x48, 0x6f, 0x6d, 0x65, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x15, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x45, 0x6e, 0x74,
	0x65, 0x72, 0x48, 0x6f, 0x6d, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x38, 0x0a, 0x0f,
	0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x70, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x18,
	0x18, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x50,
	0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x52, 0x0e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x50,
	0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x12, 0x24, 0x0a, 0x0e, 0x69, 0x73, 0x5f, 0x67, 0x61, 0x6d,
	0x65, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x19, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c,
	0x69, 0x73, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x22, 0x0a, 0x0d,
	0x69, 0x73, 0x5f, 0x70, 0x73, 0x6e, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x1a, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0b, 0x69, 0x73, 0x50, 0x73, 0x6e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x12, 0x32, 0x0a, 0x0d, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x1b, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0c, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x49, 0x45, 0x41, 0x48, 0x44, 0x43, 0x4c, 0x44,
	0x4f, 0x45, 0x4a, 0x18, 0x1c, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x49, 0x45, 0x41, 0x48, 0x44,
	0x43, 0x4c, 0x44, 0x4f, 0x45, 0x4a, 0x12, 0x20, 0x0a, 0x0b, 0x42, 0x4a, 0x46, 0x4a, 0x4a, 0x4d,
	0x47, 0x45, 0x4e, 0x43, 0x48, 0x18, 0x1d, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x42, 0x4a, 0x46,
	0x4a, 0x4a, 0x4d, 0x47, 0x45, 0x4e, 0x43, 0x48, 0x42, 0x20, 0x5a, 0x1e, 0x47, 0x6f, 0x2d, 0x47,
	0x72, 0x61, 0x73, 0x73, 0x63, 0x75, 0x74, 0x74, 0x65, 0x72, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x74, 0x65, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_proto_FriendBrief_proto_rawDescOnce sync.Once
	file_proto_FriendBrief_proto_rawDescData = file_proto_FriendBrief_proto_rawDesc
)

func file_proto_FriendBrief_proto_rawDescGZIP() []byte {
	file_proto_FriendBrief_proto_rawDescOnce.Do(func() {
		file_proto_FriendBrief_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_FriendBrief_proto_rawDescData)
	})
	return file_proto_FriendBrief_proto_rawDescData
}

var file_proto_FriendBrief_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_FriendBrief_proto_goTypes = []interface{}{
	(*FriendBrief)(nil),          // 0: FriendBrief
	(FriendOnlineState)(0),       // 1: FriendOnlineState
	(*SocialShowAvatarInfo)(nil), // 2: SocialShowAvatarInfo
	(FriendEnterHomeOption)(0),   // 3: FriendEnterHomeOption
	(*ProfilePicture)(nil),       // 4: ProfilePicture
	(PlatformType)(0),            // 5: PlatformType
}
var file_proto_FriendBrief_proto_depIdxs = []int32{
	1, // 0: FriendBrief.online_state:type_name -> FriendOnlineState
	2, // 1: FriendBrief.show_avatar_info_list:type_name -> SocialShowAvatarInfo
	3, // 2: FriendBrief.friend_enter_home_option:type_name -> FriendEnterHomeOption
	4, // 3: FriendBrief.profile_picture:type_name -> ProfilePicture
	5, // 4: FriendBrief.platform_type:type_name -> PlatformType
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_proto_FriendBrief_proto_init() }
func file_proto_FriendBrief_proto_init() {
	if File_proto_FriendBrief_proto != nil {
		return
	}
	file_proto_FriendOnlineState_proto_init()
	file_proto_SocialShowAvatarInfo_proto_init()
	file_proto_FriendEnterHomeOption_proto_init()
	file_proto_ProfilePicture_proto_init()
	file_proto_PlatformType_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_FriendBrief_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FriendBrief); i {
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
			RawDescriptor: file_proto_FriendBrief_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_FriendBrief_proto_goTypes,
		DependencyIndexes: file_proto_FriendBrief_proto_depIdxs,
		MessageInfos:      file_proto_FriendBrief_proto_msgTypes,
	}.Build()
	File_proto_FriendBrief_proto = out.File
	file_proto_FriendBrief_proto_rawDesc = nil
	file_proto_FriendBrief_proto_goTypes = nil
	file_proto_FriendBrief_proto_depIdxs = nil
}
