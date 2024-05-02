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
// source: proto/ReunionMissionInfo.proto

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

// Obf: NOGCOGMPAGJ
type ReunionMissionInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsFinished        bool                  `protobuf:"varint,7,opt,name=is_finished,json=isFinished,proto3" json:"is_finished,omitempty"`
	MissionId         uint32                `protobuf:"varint,15,opt,name=mission_id,json=missionId,proto3" json:"mission_id,omitempty"`
	IsTakenReward     bool                  `protobuf:"varint,8,opt,name=is_taken_reward,json=isTakenReward,proto3" json:"is_taken_reward,omitempty"`
	NextRefreshTime   uint32                `protobuf:"varint,4,opt,name=next_refresh_time,json=nextRefreshTime,proto3" json:"next_refresh_time,omitempty"`
	WatcherList       []*ReunionWatcherInfo `protobuf:"bytes,10,rep,name=watcher_list,json=watcherList,proto3" json:"watcher_list,omitempty"`
	IsTakenRewardList []bool                `protobuf:"varint,6,rep,packed,name=is_taken_reward_list,json=isTakenRewardList,proto3" json:"is_taken_reward_list,omitempty"`
	CurDayWatcherList []*ReunionWatcherInfo `protobuf:"bytes,1,rep,name=cur_day_watcher_list,json=curDayWatcherList,proto3" json:"cur_day_watcher_list,omitempty"`
	CurScore          uint32                `protobuf:"varint,12,opt,name=cur_score,json=curScore,proto3" json:"cur_score,omitempty"`
}

func (x *ReunionMissionInfo) Reset() {
	*x = ReunionMissionInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ReunionMissionInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReunionMissionInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReunionMissionInfo) ProtoMessage() {}

func (x *ReunionMissionInfo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ReunionMissionInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReunionMissionInfo.ProtoReflect.Descriptor instead.
func (*ReunionMissionInfo) Descriptor() ([]byte, []int) {
	return file_proto_ReunionMissionInfo_proto_rawDescGZIP(), []int{0}
}

func (x *ReunionMissionInfo) GetIsFinished() bool {
	if x != nil {
		return x.IsFinished
	}
	return false
}

func (x *ReunionMissionInfo) GetMissionId() uint32 {
	if x != nil {
		return x.MissionId
	}
	return 0
}

func (x *ReunionMissionInfo) GetIsTakenReward() bool {
	if x != nil {
		return x.IsTakenReward
	}
	return false
}

func (x *ReunionMissionInfo) GetNextRefreshTime() uint32 {
	if x != nil {
		return x.NextRefreshTime
	}
	return 0
}

func (x *ReunionMissionInfo) GetWatcherList() []*ReunionWatcherInfo {
	if x != nil {
		return x.WatcherList
	}
	return nil
}

func (x *ReunionMissionInfo) GetIsTakenRewardList() []bool {
	if x != nil {
		return x.IsTakenRewardList
	}
	return nil
}

func (x *ReunionMissionInfo) GetCurDayWatcherList() []*ReunionWatcherInfo {
	if x != nil {
		return x.CurDayWatcherList
	}
	return nil
}

func (x *ReunionMissionInfo) GetCurScore() uint32 {
	if x != nil {
		return x.CurScore
	}
	return 0
}

var File_proto_ReunionMissionInfo_proto protoreflect.FileDescriptor

var file_proto_ReunionMissionInfo_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x52, 0x65, 0x75, 0x6e, 0x69, 0x6f, 0x6e, 0x4d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x52, 0x65, 0x75, 0x6e, 0x69, 0x6f, 0x6e, 0x57,
	0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xf4, 0x02, 0x0a, 0x12, 0x52, 0x65, 0x75, 0x6e, 0x69, 0x6f, 0x6e, 0x4d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x73, 0x5f, 0x66, 0x69,
	0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x69, 0x73,
	0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x0f, 0x69, 0x73, 0x5f, 0x74, 0x61,
	0x6b, 0x65, 0x6e, 0x5f, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0d, 0x69, 0x73, 0x54, 0x61, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x12,
	0x2a, 0x0a, 0x11, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x6e, 0x65, 0x78, 0x74,
	0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x36, 0x0a, 0x0c, 0x77,
	0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0a, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x13, 0x2e, 0x52, 0x65, 0x75, 0x6e, 0x69, 0x6f, 0x6e, 0x57, 0x61, 0x74, 0x63, 0x68,
	0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0b, 0x77, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x2f, 0x0a, 0x14, 0x69, 0x73, 0x5f, 0x74, 0x61, 0x6b, 0x65, 0x6e, 0x5f,
	0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x06, 0x20, 0x03, 0x28,
	0x08, 0x52, 0x11, 0x69, 0x73, 0x54, 0x61, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x44, 0x0a, 0x14, 0x63, 0x75, 0x72, 0x5f, 0x64, 0x61, 0x79, 0x5f,
	0x77, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x13, 0x2e, 0x52, 0x65, 0x75, 0x6e, 0x69, 0x6f, 0x6e, 0x57, 0x61, 0x74, 0x63,
	0x68, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x11, 0x63, 0x75, 0x72, 0x44, 0x61, 0x79, 0x57,
	0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x75,
	0x72, 0x5f, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x63,
	0x75, 0x72, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x42, 0x1d, 0x5a, 0x1b, 0x47, 0x6f, 0x2d, 0x47, 0x72,
	0x61, 0x73, 0x73, 0x63, 0x75, 0x74, 0x74, 0x65, 0x72, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x74, 0x65, 0x64, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_ReunionMissionInfo_proto_rawDescOnce sync.Once
	file_proto_ReunionMissionInfo_proto_rawDescData = file_proto_ReunionMissionInfo_proto_rawDesc
)

func file_proto_ReunionMissionInfo_proto_rawDescGZIP() []byte {
	file_proto_ReunionMissionInfo_proto_rawDescOnce.Do(func() {
		file_proto_ReunionMissionInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_ReunionMissionInfo_proto_rawDescData)
	})
	return file_proto_ReunionMissionInfo_proto_rawDescData
}

var file_proto_ReunionMissionInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_ReunionMissionInfo_proto_goTypes = []interface{}{
	(*ReunionMissionInfo)(nil), // 0: ReunionMissionInfo
	(*ReunionWatcherInfo)(nil), // 1: ReunionWatcherInfo
}
var file_proto_ReunionMissionInfo_proto_depIdxs = []int32{
	1, // 0: ReunionMissionInfo.watcher_list:type_name -> ReunionWatcherInfo
	1, // 1: ReunionMissionInfo.cur_day_watcher_list:type_name -> ReunionWatcherInfo
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_ReunionMissionInfo_proto_init() }
func file_proto_ReunionMissionInfo_proto_init() {
	if File_proto_ReunionMissionInfo_proto != nil {
		return
	}
	file_proto_ReunionWatcherInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_ReunionMissionInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReunionMissionInfo); i {
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
			RawDescriptor: file_proto_ReunionMissionInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_ReunionMissionInfo_proto_goTypes,
		DependencyIndexes: file_proto_ReunionMissionInfo_proto_depIdxs,
		MessageInfos:      file_proto_ReunionMissionInfo_proto_msgTypes,
	}.Build()
	File_proto_ReunionMissionInfo_proto = out.File
	file_proto_ReunionMissionInfo_proto_rawDesc = nil
	file_proto_ReunionMissionInfo_proto_goTypes = nil
	file_proto_ReunionMissionInfo_proto_depIdxs = nil
}
