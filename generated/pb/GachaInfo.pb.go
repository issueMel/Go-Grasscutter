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
// source: proto/GachaInfo.proto

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

// Obf: LBJPFCBOKFM
type GachaInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CostItemId                 uint32         `protobuf:"varint,13,opt,name=costItemId,proto3" json:"costItemId,omitempty"`
	GachaPrefabPath            string         `protobuf:"bytes,10,opt,name=gachaPrefabPath,proto3" json:"gachaPrefabPath,omitempty"`
	GachaType                  uint32         `protobuf:"varint,1,opt,name=gachaType,proto3" json:"gachaType,omitempty"`
	WishMaxProgress            uint32         `protobuf:"varint,960,opt,name=wishMaxProgress,proto3" json:"wishMaxProgress,omitempty"`
	CurScheduleDailyGachaTimes uint32         `protobuf:"varint,1456,opt,name=curScheduleDailyGachaTimes,proto3" json:"curScheduleDailyGachaTimes,omitempty"`
	LeftGachaTimes             uint32         `protobuf:"varint,15,opt,name=leftGachaTimes,proto3" json:"leftGachaTimes,omitempty"`
	IsNewWish                  bool           `protobuf:"varint,1537,opt,name=is_new_wish,json=isNewWish,proto3" json:"is_new_wish,omitempty"`
	TenCostItemNum             uint32         `protobuf:"varint,8,opt,name=tenCostItemNum,proto3" json:"tenCostItemNum,omitempty"`
	WishProgress               uint32         `protobuf:"varint,1511,opt,name=wishProgress,proto3" json:"wishProgress,omitempty"`
	ScheduleId                 uint32         `protobuf:"varint,5,opt,name=schedule_id,json=scheduleId,proto3" json:"schedule_id,omitempty"`
	TenCostItemId              uint32         `protobuf:"varint,6,opt,name=tenCostItemId,proto3" json:"tenCostItemId,omitempty"`
	GachaSortId                uint32         `protobuf:"varint,11,opt,name=gachaSortId,proto3" json:"gachaSortId,omitempty"`
	FFMOBAAKKBH                []uint32       `protobuf:"varint,1285,rep,packed,name=FFMOBAAKKBH,proto3" json:"FFMOBAAKKBH,omitempty"`
	GachaRecordUrlOversea      string         `protobuf:"bytes,1181,opt,name=gachaRecordUrlOversea,proto3" json:"gachaRecordUrlOversea,omitempty"`
	GachaProbUrlOversea        string         `protobuf:"bytes,1221,opt,name=gachaProbUrlOversea,proto3" json:"gachaProbUrlOversea,omitempty"`
	BeginTime                  uint32         `protobuf:"varint,3,opt,name=begin_time,json=beginTime,proto3" json:"begin_time,omitempty"`
	WishItemId                 uint32         `protobuf:"varint,870,opt,name=wishItemId,proto3" json:"wishItemId,omitempty"`
	EndTime                    uint32         `protobuf:"varint,7,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	GachaPreviewPrefabPath     string         `protobuf:"bytes,12,opt,name=gachaPreviewPrefabPath,proto3" json:"gachaPreviewPrefabPath,omitempty"`
	TitleTextmap               string         `protobuf:"bytes,847,opt,name=titleTextmap,proto3" json:"titleTextmap,omitempty"`
	OCIHCEACNNF                []uint32       `protobuf:"varint,1831,rep,packed,name=OCIHCEACNNF,proto3" json:"OCIHCEACNNF,omitempty"`
	GachaTimesLimit            uint32         `protobuf:"varint,14,opt,name=gachaTimesLimit,proto3" json:"gachaTimesLimit,omitempty"`
	GachaRecordUrl             string         `protobuf:"bytes,9,opt,name=gachaRecordUrl,proto3" json:"gachaRecordUrl,omitempty"`
	GachaUpInfoList            []*GachaUpInfo `protobuf:"bytes,1514,rep,name=gacha_up_info_list,json=gachaUpInfoList,proto3" json:"gacha_up_info_list,omitempty"`
	CostItemNum                uint32         `protobuf:"varint,2,opt,name=costItemNum,proto3" json:"costItemNum,omitempty"`
	GachaProbUrl               string         `protobuf:"bytes,4,opt,name=gachaProbUrl,proto3" json:"gachaProbUrl,omitempty"`
}

func (x *GachaInfo) Reset() {
	*x = GachaInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_GachaInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GachaInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GachaInfo) ProtoMessage() {}

func (x *GachaInfo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_GachaInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GachaInfo.ProtoReflect.Descriptor instead.
func (*GachaInfo) Descriptor() ([]byte, []int) {
	return file_proto_GachaInfo_proto_rawDescGZIP(), []int{0}
}

func (x *GachaInfo) GetCostItemId() uint32 {
	if x != nil {
		return x.CostItemId
	}
	return 0
}

func (x *GachaInfo) GetGachaPrefabPath() string {
	if x != nil {
		return x.GachaPrefabPath
	}
	return ""
}

func (x *GachaInfo) GetGachaType() uint32 {
	if x != nil {
		return x.GachaType
	}
	return 0
}

func (x *GachaInfo) GetWishMaxProgress() uint32 {
	if x != nil {
		return x.WishMaxProgress
	}
	return 0
}

func (x *GachaInfo) GetCurScheduleDailyGachaTimes() uint32 {
	if x != nil {
		return x.CurScheduleDailyGachaTimes
	}
	return 0
}

func (x *GachaInfo) GetLeftGachaTimes() uint32 {
	if x != nil {
		return x.LeftGachaTimes
	}
	return 0
}

func (x *GachaInfo) GetIsNewWish() bool {
	if x != nil {
		return x.IsNewWish
	}
	return false
}

func (x *GachaInfo) GetTenCostItemNum() uint32 {
	if x != nil {
		return x.TenCostItemNum
	}
	return 0
}

func (x *GachaInfo) GetWishProgress() uint32 {
	if x != nil {
		return x.WishProgress
	}
	return 0
}

func (x *GachaInfo) GetScheduleId() uint32 {
	if x != nil {
		return x.ScheduleId
	}
	return 0
}

func (x *GachaInfo) GetTenCostItemId() uint32 {
	if x != nil {
		return x.TenCostItemId
	}
	return 0
}

func (x *GachaInfo) GetGachaSortId() uint32 {
	if x != nil {
		return x.GachaSortId
	}
	return 0
}

func (x *GachaInfo) GetFFMOBAAKKBH() []uint32 {
	if x != nil {
		return x.FFMOBAAKKBH
	}
	return nil
}

func (x *GachaInfo) GetGachaRecordUrlOversea() string {
	if x != nil {
		return x.GachaRecordUrlOversea
	}
	return ""
}

func (x *GachaInfo) GetGachaProbUrlOversea() string {
	if x != nil {
		return x.GachaProbUrlOversea
	}
	return ""
}

func (x *GachaInfo) GetBeginTime() uint32 {
	if x != nil {
		return x.BeginTime
	}
	return 0
}

func (x *GachaInfo) GetWishItemId() uint32 {
	if x != nil {
		return x.WishItemId
	}
	return 0
}

func (x *GachaInfo) GetEndTime() uint32 {
	if x != nil {
		return x.EndTime
	}
	return 0
}

func (x *GachaInfo) GetGachaPreviewPrefabPath() string {
	if x != nil {
		return x.GachaPreviewPrefabPath
	}
	return ""
}

func (x *GachaInfo) GetTitleTextmap() string {
	if x != nil {
		return x.TitleTextmap
	}
	return ""
}

func (x *GachaInfo) GetOCIHCEACNNF() []uint32 {
	if x != nil {
		return x.OCIHCEACNNF
	}
	return nil
}

func (x *GachaInfo) GetGachaTimesLimit() uint32 {
	if x != nil {
		return x.GachaTimesLimit
	}
	return 0
}

func (x *GachaInfo) GetGachaRecordUrl() string {
	if x != nil {
		return x.GachaRecordUrl
	}
	return ""
}

func (x *GachaInfo) GetGachaUpInfoList() []*GachaUpInfo {
	if x != nil {
		return x.GachaUpInfoList
	}
	return nil
}

func (x *GachaInfo) GetCostItemNum() uint32 {
	if x != nil {
		return x.CostItemNum
	}
	return 0
}

func (x *GachaInfo) GetGachaProbUrl() string {
	if x != nil {
		return x.GachaProbUrl
	}
	return ""
}

var File_proto_GachaInfo_proto protoreflect.FileDescriptor

var file_proto_GachaInfo_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x47, 0x61, 0x63, 0x68, 0x61, 0x49, 0x6e, 0x66,
	0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x47,
	0x61, 0x63, 0x68, 0x61, 0x55, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x9a, 0x08, 0x0a, 0x09, 0x47, 0x61, 0x63, 0x68, 0x61, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1e,
	0x0a, 0x0a, 0x63, 0x6f, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x18, 0x0d, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0a, 0x63, 0x6f, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x28,
	0x0a, 0x0f, 0x67, 0x61, 0x63, 0x68, 0x61, 0x50, 0x72, 0x65, 0x66, 0x61, 0x62, 0x50, 0x61, 0x74,
	0x68, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x67, 0x61, 0x63, 0x68, 0x61, 0x50, 0x72,
	0x65, 0x66, 0x61, 0x62, 0x50, 0x61, 0x74, 0x68, 0x12, 0x1c, 0x0a, 0x09, 0x67, 0x61, 0x63, 0x68,
	0x61, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x67, 0x61, 0x63,
	0x68, 0x61, 0x54, 0x79, 0x70, 0x65, 0x12, 0x29, 0x0a, 0x0f, 0x77, 0x69, 0x73, 0x68, 0x4d, 0x61,
	0x78, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x18, 0xc0, 0x07, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0f, 0x77, 0x69, 0x73, 0x68, 0x4d, 0x61, 0x78, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73,
	0x73, 0x12, 0x3f, 0x0a, 0x1a, 0x63, 0x75, 0x72, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65,
	0x44, 0x61, 0x69, 0x6c, 0x79, 0x47, 0x61, 0x63, 0x68, 0x61, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x18,
	0xb0, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x1a, 0x63, 0x75, 0x72, 0x53, 0x63, 0x68, 0x65, 0x64,
	0x75, 0x6c, 0x65, 0x44, 0x61, 0x69, 0x6c, 0x79, 0x47, 0x61, 0x63, 0x68, 0x61, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x12, 0x26, 0x0a, 0x0e, 0x6c, 0x65, 0x66, 0x74, 0x47, 0x61, 0x63, 0x68, 0x61, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x6c, 0x65, 0x66, 0x74,
	0x47, 0x61, 0x63, 0x68, 0x61, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x73,
	0x5f, 0x6e, 0x65, 0x77, 0x5f, 0x77, 0x69, 0x73, 0x68, 0x18, 0x81, 0x0c, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x09, 0x69, 0x73, 0x4e, 0x65, 0x77, 0x57, 0x69, 0x73, 0x68, 0x12, 0x26, 0x0a, 0x0e, 0x74,
	0x65, 0x6e, 0x43, 0x6f, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x4e, 0x75, 0x6d, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0e, 0x74, 0x65, 0x6e, 0x43, 0x6f, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d,
	0x4e, 0x75, 0x6d, 0x12, 0x23, 0x0a, 0x0c, 0x77, 0x69, 0x73, 0x68, 0x50, 0x72, 0x6f, 0x67, 0x72,
	0x65, 0x73, 0x73, 0x18, 0xe7, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x77, 0x69, 0x73, 0x68,
	0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x63, 0x68, 0x65,
	0x64, 0x75, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x73,
	0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0d, 0x74, 0x65, 0x6e,
	0x43, 0x6f, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0d, 0x74, 0x65, 0x6e, 0x43, 0x6f, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x12,
	0x20, 0x0a, 0x0b, 0x67, 0x61, 0x63, 0x68, 0x61, 0x53, 0x6f, 0x72, 0x74, 0x49, 0x64, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x67, 0x61, 0x63, 0x68, 0x61, 0x53, 0x6f, 0x72, 0x74, 0x49,
	0x64, 0x12, 0x21, 0x0a, 0x0b, 0x46, 0x46, 0x4d, 0x4f, 0x42, 0x41, 0x41, 0x4b, 0x4b, 0x42, 0x48,
	0x18, 0x85, 0x0a, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0b, 0x46, 0x46, 0x4d, 0x4f, 0x42, 0x41, 0x41,
	0x4b, 0x4b, 0x42, 0x48, 0x12, 0x35, 0x0a, 0x15, 0x67, 0x61, 0x63, 0x68, 0x61, 0x52, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x55, 0x72, 0x6c, 0x4f, 0x76, 0x65, 0x72, 0x73, 0x65, 0x61, 0x18, 0x9d, 0x09,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x15, 0x67, 0x61, 0x63, 0x68, 0x61, 0x52, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x55, 0x72, 0x6c, 0x4f, 0x76, 0x65, 0x72, 0x73, 0x65, 0x61, 0x12, 0x31, 0x0a, 0x13, 0x67,
	0x61, 0x63, 0x68, 0x61, 0x50, 0x72, 0x6f, 0x62, 0x55, 0x72, 0x6c, 0x4f, 0x76, 0x65, 0x72, 0x73,
	0x65, 0x61, 0x18, 0xc5, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x67, 0x61, 0x63, 0x68, 0x61,
	0x50, 0x72, 0x6f, 0x62, 0x55, 0x72, 0x6c, 0x4f, 0x76, 0x65, 0x72, 0x73, 0x65, 0x61, 0x12, 0x1d,
	0x0a, 0x0a, 0x62, 0x65, 0x67, 0x69, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x09, 0x62, 0x65, 0x67, 0x69, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1f, 0x0a,
	0x0a, 0x77, 0x69, 0x73, 0x68, 0x49, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x18, 0xe6, 0x06, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0a, 0x77, 0x69, 0x73, 0x68, 0x49, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x19,
	0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x36, 0x0a, 0x16, 0x67, 0x61, 0x63,
	0x68, 0x61, 0x50, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x50, 0x72, 0x65, 0x66, 0x61, 0x62, 0x50,
	0x61, 0x74, 0x68, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x16, 0x67, 0x61, 0x63, 0x68, 0x61,
	0x50, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x50, 0x72, 0x65, 0x66, 0x61, 0x62, 0x50, 0x61, 0x74,
	0x68, 0x12, 0x23, 0x0a, 0x0c, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x54, 0x65, 0x78, 0x74, 0x6d, 0x61,
	0x70, 0x18, 0xcf, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x54,
	0x65, 0x78, 0x74, 0x6d, 0x61, 0x70, 0x12, 0x21, 0x0a, 0x0b, 0x4f, 0x43, 0x49, 0x48, 0x43, 0x45,
	0x41, 0x43, 0x4e, 0x4e, 0x46, 0x18, 0xa7, 0x0e, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0b, 0x4f, 0x43,
	0x49, 0x48, 0x43, 0x45, 0x41, 0x43, 0x4e, 0x4e, 0x46, 0x12, 0x28, 0x0a, 0x0f, 0x67, 0x61, 0x63,
	0x68, 0x61, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x0e, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0f, 0x67, 0x61, 0x63, 0x68, 0x61, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x4c, 0x69,
	0x6d, 0x69, 0x74, 0x12, 0x26, 0x0a, 0x0e, 0x67, 0x61, 0x63, 0x68, 0x61, 0x52, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x55, 0x72, 0x6c, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x67, 0x61, 0x63,
	0x68, 0x61, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x55, 0x72, 0x6c, 0x12, 0x3a, 0x0a, 0x12, 0x67,
	0x61, 0x63, 0x68, 0x61, 0x5f, 0x75, 0x70, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x5f, 0x6c, 0x69, 0x73,
	0x74, 0x18, 0xea, 0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x47, 0x61, 0x63, 0x68, 0x61,
	0x55, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0f, 0x67, 0x61, 0x63, 0x68, 0x61, 0x55, 0x70, 0x49,
	0x6e, 0x66, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x73, 0x74, 0x49,
	0x74, 0x65, 0x6d, 0x4e, 0x75, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x63, 0x6f,
	0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x4e, 0x75, 0x6d, 0x12, 0x22, 0x0a, 0x0c, 0x67, 0x61, 0x63,
	0x68, 0x61, 0x50, 0x72, 0x6f, 0x62, 0x55, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x67, 0x61, 0x63, 0x68, 0x61, 0x50, 0x72, 0x6f, 0x62, 0x55, 0x72, 0x6c, 0x42, 0x1d, 0x5a,
	0x1b, 0x47, 0x6f, 0x2d, 0x47, 0x72, 0x61, 0x73, 0x73, 0x63, 0x75, 0x74, 0x74, 0x65, 0x72, 0x2f,
	0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_GachaInfo_proto_rawDescOnce sync.Once
	file_proto_GachaInfo_proto_rawDescData = file_proto_GachaInfo_proto_rawDesc
)

func file_proto_GachaInfo_proto_rawDescGZIP() []byte {
	file_proto_GachaInfo_proto_rawDescOnce.Do(func() {
		file_proto_GachaInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_GachaInfo_proto_rawDescData)
	})
	return file_proto_GachaInfo_proto_rawDescData
}

var file_proto_GachaInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_GachaInfo_proto_goTypes = []interface{}{
	(*GachaInfo)(nil),   // 0: GachaInfo
	(*GachaUpInfo)(nil), // 1: GachaUpInfo
}
var file_proto_GachaInfo_proto_depIdxs = []int32{
	1, // 0: GachaInfo.gacha_up_info_list:type_name -> GachaUpInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_GachaInfo_proto_init() }
func file_proto_GachaInfo_proto_init() {
	if File_proto_GachaInfo_proto != nil {
		return
	}
	file_proto_GachaUpInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_GachaInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GachaInfo); i {
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
			RawDescriptor: file_proto_GachaInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_GachaInfo_proto_goTypes,
		DependencyIndexes: file_proto_GachaInfo_proto_depIdxs,
		MessageInfos:      file_proto_GachaInfo_proto_msgTypes,
	}.Build()
	File_proto_GachaInfo_proto = out.File
	file_proto_GachaInfo_proto_rawDesc = nil
	file_proto_GachaInfo_proto_goTypes = nil
	file_proto_GachaInfo_proto_depIdxs = nil
}