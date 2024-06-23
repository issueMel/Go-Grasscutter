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
// source: GetPlayerTokenRsp.proto

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

// CmdId: 2407
// Obf: BKFNIGMLAOA
type GetPlayerTokenRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PELBMOHDKHJ            bool            `protobuf:"varint,1292,opt,name=PELBMOHDKHJ,proto3" json:"PELBMOHDKHJ,omitempty"`
	ILGBPPMFHIF            bool            `protobuf:"varint,1962,opt,name=ILGBPPMFHIF,proto3" json:"ILGBPPMFHIF,omitempty"`
	SecretKeySeed          uint64          `protobuf:"varint,13,opt,name=secret_key_seed,json=secretKeySeed,proto3" json:"secret_key_seed,omitempty"`
	ClientVersionRandomKey string          `protobuf:"bytes,678,opt,name=client_version_random_key,json=clientVersionRandomKey,proto3" json:"client_version_random_key,omitempty"`
	StopServer             *StopServerInfo `protobuf:"bytes,875,opt,name=stop_server,json=stopServer,proto3" json:"stop_server,omitempty"`
	ChannelId              uint32          `protobuf:"varint,259,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
	PlatformType           uint32          `protobuf:"varint,5,opt,name=platform_type,json=platformType,proto3" json:"platform_type,omitempty"`
	RegPlatform            uint32          `protobuf:"varint,1241,opt,name=reg_platform,json=regPlatform,proto3" json:"reg_platform,omitempty"`
	AccountUid             string          `protobuf:"bytes,3,opt,name=account_uid,json=accountUid,proto3" json:"account_uid,omitempty"`
	Uid                    uint32          `protobuf:"varint,7,opt,name=uid,proto3" json:"uid,omitempty"`
	Tag                    uint32          `protobuf:"varint,1572,opt,name=tag,proto3" json:"tag,omitempty"`
	Birthday               string          `protobuf:"bytes,1437,opt,name=birthday,proto3" json:"birthday,omitempty"`
	KeyId                  uint32          `protobuf:"varint,1676,opt,name=key_id,json=keyId,proto3" json:"key_id,omitempty"`
	Retcode                int32           `protobuf:"varint,4,opt,name=retcode,proto3" json:"retcode,omitempty"`
	IsGuest                bool            `protobuf:"varint,2,opt,name=is_guest,json=isGuest,proto3" json:"is_guest,omitempty"`
	FinishCollectionIdList []uint32        `protobuf:"varint,451,rep,packed,name=finish_collection_id_list,json=finishCollectionIdList,proto3" json:"finish_collection_id_list,omitempty"`
	ServerRandKey          string          `protobuf:"bytes,517,opt,name=server_rand_key,json=serverRandKey,proto3" json:"server_rand_key,omitempty"`
	SecurityCmdBuffer      []byte          `protobuf:"bytes,15,opt,name=security_cmd_buffer,json=securityCmdBuffer,proto3" json:"security_cmd_buffer,omitempty"`
	Msg                    string          `protobuf:"bytes,12,opt,name=msg,proto3" json:"msg,omitempty"`
	IsProficientPlayer     bool            `protobuf:"varint,8,opt,name=is_proficient_player,json=isProficientPlayer,proto3" json:"is_proficient_player,omitempty"`
	Token                  string          `protobuf:"bytes,11,opt,name=token,proto3" json:"token,omitempty"`
	Sign                   string          `protobuf:"bytes,394,opt,name=sign,proto3" json:"sign,omitempty"`
	AccountType            uint32          `protobuf:"varint,1585,opt,name=account_type,json=accountType,proto3" json:"account_type,omitempty"`
	ExtraBinData           []byte          `protobuf:"bytes,1,opt,name=extra_bin_data,json=extraBinData,proto3" json:"extra_bin_data,omitempty"`
	PsnId                  string          `protobuf:"bytes,1790,opt,name=psn_id,json=psnId,proto3" json:"psn_id,omitempty"`
	SubChannelId           uint32          `protobuf:"varint,14,opt,name=sub_channel_id,json=subChannelId,proto3" json:"sub_channel_id,omitempty"`
	SecretKey              string          `protobuf:"bytes,9,opt,name=secret_key,json=secretKey,proto3" json:"secret_key,omitempty"`
	ClientIpStr            string          `protobuf:"bytes,1921,opt,name=client_ip_str,json=clientIpStr,proto3" json:"client_ip_str,omitempty"`
	BlackUidEndTime        uint32          `protobuf:"varint,10,opt,name=blackUidEndTime,proto3" json:"blackUidEndTime,omitempty"`
	KCFIGJAPNIB            uint32          `protobuf:"varint,6,opt,name=KCFIGJAPNIB,proto3" json:"KCFIGJAPNIB,omitempty"`
	CountryCode            string          `protobuf:"bytes,1412,opt,name=country_code,json=countryCode,proto3" json:"country_code,omitempty"`
	CloudClientIp          uint32          `protobuf:"varint,508,opt,name=cloudClientIp,proto3" json:"cloudClientIp,omitempty"`
}

func (x *GetPlayerTokenRsp) Reset() {
	*x = GetPlayerTokenRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GetPlayerTokenRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPlayerTokenRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPlayerTokenRsp) ProtoMessage() {}

func (x *GetPlayerTokenRsp) ProtoReflect() protoreflect.Message {
	mi := &file_GetPlayerTokenRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPlayerTokenRsp.ProtoReflect.Descriptor instead.
func (*GetPlayerTokenRsp) Descriptor() ([]byte, []int) {
	return file_GetPlayerTokenRsp_proto_rawDescGZIP(), []int{0}
}

func (x *GetPlayerTokenRsp) GetPELBMOHDKHJ() bool {
	if x != nil {
		return x.PELBMOHDKHJ
	}
	return false
}

func (x *GetPlayerTokenRsp) GetILGBPPMFHIF() bool {
	if x != nil {
		return x.ILGBPPMFHIF
	}
	return false
}

func (x *GetPlayerTokenRsp) GetSecretKeySeed() uint64 {
	if x != nil {
		return x.SecretKeySeed
	}
	return 0
}

func (x *GetPlayerTokenRsp) GetClientVersionRandomKey() string {
	if x != nil {
		return x.ClientVersionRandomKey
	}
	return ""
}

func (x *GetPlayerTokenRsp) GetStopServer() *StopServerInfo {
	if x != nil {
		return x.StopServer
	}
	return nil
}

func (x *GetPlayerTokenRsp) GetChannelId() uint32 {
	if x != nil {
		return x.ChannelId
	}
	return 0
}

func (x *GetPlayerTokenRsp) GetPlatformType() uint32 {
	if x != nil {
		return x.PlatformType
	}
	return 0
}

func (x *GetPlayerTokenRsp) GetRegPlatform() uint32 {
	if x != nil {
		return x.RegPlatform
	}
	return 0
}

func (x *GetPlayerTokenRsp) GetAccountUid() string {
	if x != nil {
		return x.AccountUid
	}
	return ""
}

func (x *GetPlayerTokenRsp) GetUid() uint32 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *GetPlayerTokenRsp) GetTag() uint32 {
	if x != nil {
		return x.Tag
	}
	return 0
}

func (x *GetPlayerTokenRsp) GetBirthday() string {
	if x != nil {
		return x.Birthday
	}
	return ""
}

func (x *GetPlayerTokenRsp) GetKeyId() uint32 {
	if x != nil {
		return x.KeyId
	}
	return 0
}

func (x *GetPlayerTokenRsp) GetRetcode() int32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *GetPlayerTokenRsp) GetIsGuest() bool {
	if x != nil {
		return x.IsGuest
	}
	return false
}

func (x *GetPlayerTokenRsp) GetFinishCollectionIdList() []uint32 {
	if x != nil {
		return x.FinishCollectionIdList
	}
	return nil
}

func (x *GetPlayerTokenRsp) GetServerRandKey() string {
	if x != nil {
		return x.ServerRandKey
	}
	return ""
}

func (x *GetPlayerTokenRsp) GetSecurityCmdBuffer() []byte {
	if x != nil {
		return x.SecurityCmdBuffer
	}
	return nil
}

func (x *GetPlayerTokenRsp) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *GetPlayerTokenRsp) GetIsProficientPlayer() bool {
	if x != nil {
		return x.IsProficientPlayer
	}
	return false
}

func (x *GetPlayerTokenRsp) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *GetPlayerTokenRsp) GetSign() string {
	if x != nil {
		return x.Sign
	}
	return ""
}

func (x *GetPlayerTokenRsp) GetAccountType() uint32 {
	if x != nil {
		return x.AccountType
	}
	return 0
}

func (x *GetPlayerTokenRsp) GetExtraBinData() []byte {
	if x != nil {
		return x.ExtraBinData
	}
	return nil
}

func (x *GetPlayerTokenRsp) GetPsnId() string {
	if x != nil {
		return x.PsnId
	}
	return ""
}

func (x *GetPlayerTokenRsp) GetSubChannelId() uint32 {
	if x != nil {
		return x.SubChannelId
	}
	return 0
}

func (x *GetPlayerTokenRsp) GetSecretKey() string {
	if x != nil {
		return x.SecretKey
	}
	return ""
}

func (x *GetPlayerTokenRsp) GetClientIpStr() string {
	if x != nil {
		return x.ClientIpStr
	}
	return ""
}

func (x *GetPlayerTokenRsp) GetBlackUidEndTime() uint32 {
	if x != nil {
		return x.BlackUidEndTime
	}
	return 0
}

func (x *GetPlayerTokenRsp) GetKCFIGJAPNIB() uint32 {
	if x != nil {
		return x.KCFIGJAPNIB
	}
	return 0
}

func (x *GetPlayerTokenRsp) GetCountryCode() string {
	if x != nil {
		return x.CountryCode
	}
	return ""
}

func (x *GetPlayerTokenRsp) GetCloudClientIp() uint32 {
	if x != nil {
		return x.CloudClientIp
	}
	return 0
}

var File_GetPlayerTokenRsp_proto protoreflect.FileDescriptor

var file_GetPlayerTokenRsp_proto_rawDesc = []byte{
	0x0a, 0x17, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x53, 0x74, 0x6f, 0x70, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xf0, 0x08, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x52, 0x73, 0x70, 0x12, 0x21, 0x0a, 0x0b, 0x50, 0x45, 0x4c, 0x42, 0x4d, 0x4f, 0x48,
	0x44, 0x4b, 0x48, 0x4a, 0x18, 0x8c, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x50, 0x45, 0x4c,
	0x42, 0x4d, 0x4f, 0x48, 0x44, 0x4b, 0x48, 0x4a, 0x12, 0x21, 0x0a, 0x0b, 0x49, 0x4c, 0x47, 0x42,
	0x50, 0x50, 0x4d, 0x46, 0x48, 0x49, 0x46, 0x18, 0xaa, 0x0f, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b,
	0x49, 0x4c, 0x47, 0x42, 0x50, 0x50, 0x4d, 0x46, 0x48, 0x49, 0x46, 0x12, 0x26, 0x0a, 0x0f, 0x73,
	0x65, 0x63, 0x72, 0x65, 0x74, 0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x73, 0x65, 0x65, 0x64, 0x18, 0x0d,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x0d, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x4b, 0x65, 0x79, 0x53,
	0x65, 0x65, 0x64, 0x12, 0x3a, 0x0a, 0x19, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x5f, 0x6b, 0x65, 0x79,
	0x18, 0xa6, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x16, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x56,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x4b, 0x65, 0x79, 0x12,
	0x31, 0x0a, 0x0b, 0x73, 0x74, 0x6f, 0x70, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x18, 0xeb,
	0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x53, 0x74, 0x6f, 0x70, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0a, 0x73, 0x74, 0x6f, 0x70, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x69, 0x64,
	0x18, 0x83, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x70, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x72, 0x65, 0x67, 0x5f, 0x70,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x18, 0xd9, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b,
	0x72, 0x65, 0x67, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x1f, 0x0a, 0x0b, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x75, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x55, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03,
	0x75, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x11,
	0x0a, 0x03, 0x74, 0x61, 0x67, 0x18, 0xa4, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x74, 0x61,
	0x67, 0x12, 0x1b, 0x0a, 0x08, 0x62, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x79, 0x18, 0x9d, 0x0b,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x79, 0x12, 0x16,
	0x0a, 0x06, 0x6b, 0x65, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x8c, 0x0d, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x05, 0x6b, 0x65, 0x79, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65,
	0x12, 0x19, 0x0a, 0x08, 0x69, 0x73, 0x5f, 0x67, 0x75, 0x65, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x69, 0x73, 0x47, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3a, 0x0a, 0x19, 0x66,
	0x69, 0x6e, 0x69, 0x73, 0x68, 0x5f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x69, 0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0xc3, 0x03, 0x20, 0x03, 0x28, 0x0d, 0x52,
	0x16, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x5f, 0x72, 0x61, 0x6e, 0x64, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x85, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x61, 0x6e, 0x64, 0x4b, 0x65, 0x79,
	0x12, 0x2e, 0x0a, 0x13, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x5f, 0x63, 0x6d, 0x64,
	0x5f, 0x62, 0x75, 0x66, 0x66, 0x65, 0x72, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x11, 0x73,
	0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x43, 0x6d, 0x64, 0x42, 0x75, 0x66, 0x66, 0x65, 0x72,
	0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d,
	0x73, 0x67, 0x12, 0x30, 0x0a, 0x14, 0x69, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x63, 0x69,
	0x65, 0x6e, 0x74, 0x5f, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x12, 0x69, 0x73, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x63, 0x69, 0x65, 0x6e, 0x74, 0x50, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x13, 0x0a, 0x04, 0x73, 0x69,
	0x67, 0x6e, 0x18, 0x8a, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x69, 0x67, 0x6e, 0x12,
	0x22, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0xb1, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x24, 0x0a, 0x0e, 0x65, 0x78, 0x74, 0x72, 0x61, 0x5f, 0x62, 0x69, 0x6e,
	0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c, 0x65, 0x78, 0x74,
	0x72, 0x61, 0x42, 0x69, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x73, 0x6e,
	0x5f, 0x69, 0x64, 0x18, 0xfe, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x73, 0x6e, 0x49,
	0x64, 0x12, 0x24, 0x0a, 0x0e, 0x73, 0x75, 0x62, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x5f, 0x69, 0x64, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x73, 0x75, 0x62, 0x43, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x63, 0x72, 0x65,
	0x74, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x63,
	0x72, 0x65, 0x74, 0x4b, 0x65, 0x79, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x5f, 0x69, 0x70, 0x5f, 0x73, 0x74, 0x72, 0x18, 0x81, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x70, 0x53, 0x74, 0x72, 0x12, 0x28, 0x0a, 0x0f, 0x62,
	0x6c, 0x61, 0x63, 0x6b, 0x55, 0x69, 0x64, 0x45, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x62, 0x6c, 0x61, 0x63, 0x6b, 0x55, 0x69, 0x64, 0x45, 0x6e,
	0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x4b, 0x43, 0x46, 0x49, 0x47, 0x4a, 0x41,
	0x50, 0x4e, 0x49, 0x42, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x4b, 0x43, 0x46, 0x49,
	0x47, 0x4a, 0x41, 0x50, 0x4e, 0x49, 0x42, 0x12, 0x22, 0x0a, 0x0c, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x72, 0x79, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x84, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x25, 0x0a, 0x0d, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x70, 0x18, 0xfc, 0x03, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x49, 0x70, 0x42, 0x1d, 0x5a, 0x1b, 0x47, 0x6f, 0x2d, 0x47, 0x72, 0x61, 0x73, 0x73, 0x63, 0x75,
	0x74, 0x74, 0x65, 0x72, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_GetPlayerTokenRsp_proto_rawDescOnce sync.Once
	file_GetPlayerTokenRsp_proto_rawDescData = file_GetPlayerTokenRsp_proto_rawDesc
)

func file_GetPlayerTokenRsp_proto_rawDescGZIP() []byte {
	file_GetPlayerTokenRsp_proto_rawDescOnce.Do(func() {
		file_GetPlayerTokenRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_GetPlayerTokenRsp_proto_rawDescData)
	})
	return file_GetPlayerTokenRsp_proto_rawDescData
}

var file_GetPlayerTokenRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_GetPlayerTokenRsp_proto_goTypes = []interface{}{
	(*GetPlayerTokenRsp)(nil), // 0: GetPlayerTokenRsp
	(*StopServerInfo)(nil),    // 1: StopServerInfo
}
var file_GetPlayerTokenRsp_proto_depIdxs = []int32{
	1, // 0: GetPlayerTokenRsp.stop_server:type_name -> StopServerInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_GetPlayerTokenRsp_proto_init() }
func file_GetPlayerTokenRsp_proto_init() {
	if File_GetPlayerTokenRsp_proto != nil {
		return
	}
	file_proto_StopServerInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_GetPlayerTokenRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPlayerTokenRsp); i {
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
			RawDescriptor: file_GetPlayerTokenRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GetPlayerTokenRsp_proto_goTypes,
		DependencyIndexes: file_GetPlayerTokenRsp_proto_depIdxs,
		MessageInfos:      file_GetPlayerTokenRsp_proto_msgTypes,
	}.Build()
	File_GetPlayerTokenRsp_proto = out.File
	file_GetPlayerTokenRsp_proto_rawDesc = nil
	file_GetPlayerTokenRsp_proto_goTypes = nil
	file_GetPlayerTokenRsp_proto_depIdxs = nil
}
