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
// source: proto/BeginCameraSceneLookNotify.proto

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

// Obf: ONABGJCLONK
type BeginCameraSceneLookNotify_KeepRotType int32

const (
	BeginCameraSceneLookNotify_KEEP_ROT_X  BeginCameraSceneLookNotify_KeepRotType = 0
	BeginCameraSceneLookNotify_KEEP_ROT_XY BeginCameraSceneLookNotify_KeepRotType = 1
)

// Enum value maps for BeginCameraSceneLookNotify_KeepRotType.
var (
	BeginCameraSceneLookNotify_KeepRotType_name = map[int32]string{
		0: "KEEP_ROT_X",
		1: "KEEP_ROT_XY",
	}
	BeginCameraSceneLookNotify_KeepRotType_value = map[string]int32{
		"KEEP_ROT_X":  0,
		"KEEP_ROT_XY": 1,
	}
)

func (x BeginCameraSceneLookNotify_KeepRotType) Enum() *BeginCameraSceneLookNotify_KeepRotType {
	p := new(BeginCameraSceneLookNotify_KeepRotType)
	*p = x
	return p
}

func (x BeginCameraSceneLookNotify_KeepRotType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BeginCameraSceneLookNotify_KeepRotType) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_BeginCameraSceneLookNotify_proto_enumTypes[0].Descriptor()
}

func (BeginCameraSceneLookNotify_KeepRotType) Type() protoreflect.EnumType {
	return &file_proto_BeginCameraSceneLookNotify_proto_enumTypes[0]
}

func (x BeginCameraSceneLookNotify_KeepRotType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BeginCameraSceneLookNotify_KeepRotType.Descriptor instead.
func (BeginCameraSceneLookNotify_KeepRotType) EnumDescriptor() ([]byte, []int) {
	return file_proto_BeginCameraSceneLookNotify_proto_rawDescGZIP(), []int{0, 0}
}

// CmdId: 27569
// Obf: KAGMEJHECDK
type BeginCameraSceneLookNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EntityId    uint32                                 `protobuf:"varint,609,opt,name=entity_id,json=entityId,proto3" json:"entity_id,omitempty"`
	JBCEDEGCGFJ bool                                   `protobuf:"varint,8,opt,name=JBCEDEGCGFJ,proto3" json:"JBCEDEGCGFJ,omitempty"`
	ADPLABBOBKF bool                                   `protobuf:"varint,15,opt,name=ADPLABBOBKF,proto3" json:"ADPLABBOBKF,omitempty"`
	DOCBFPIPMPK uint32                                 `protobuf:"varint,1006,opt,name=DOCBFPIPMPK,proto3" json:"DOCBFPIPMPK,omitempty"`
	LKGGBPLCEJI *Vector                                `protobuf:"bytes,12,opt,name=LKGGBPLCEJI,proto3" json:"LKGGBPLCEJI,omitempty"`
	KJHLOCKMIMB bool                                   `protobuf:"varint,1477,opt,name=KJHLOCKMIMB,proto3" json:"KJHLOCKMIMB,omitempty"`
	GJNLEMGJCJI float32                                `protobuf:"fixed32,10,opt,name=GJNLEMGJCJI,proto3" json:"GJNLEMGJCJI,omitempty"`
	IPKFLPFLNNI bool                                   `protobuf:"varint,9,opt,name=IPKFLPFLNNI,proto3" json:"IPKFLPFLNNI,omitempty"`
	PDOCOOCFAIH bool                                   `protobuf:"varint,5,opt,name=PDOCOOCFAIH,proto3" json:"PDOCOOCFAIH,omitempty"`
	MMOMOKPCOJK *Vector                                `protobuf:"bytes,14,opt,name=MMOMOKPCOJK,proto3" json:"MMOMOKPCOJK,omitempty"`
	FPDPEHICGAC bool                                   `protobuf:"varint,7,opt,name=FPDPEHICGAC,proto3" json:"FPDPEHICGAC,omitempty"`
	MIKBDFOBIBB float32                                `protobuf:"fixed32,3,opt,name=MIKBDFOBIBB,proto3" json:"MIKBDFOBIBB,omitempty"`
	CLLOFLJICAF float32                                `protobuf:"fixed32,11,opt,name=CLLOFLJICAF,proto3" json:"CLLOFLJICAF,omitempty"`
	OtherParams []string                               `protobuf:"bytes,2,rep,name=other_params,json=otherParams,proto3" json:"other_params,omitempty"`
	KeepRotType BeginCameraSceneLookNotify_KeepRotType `protobuf:"varint,4,opt,name=keep_rot_type,json=keepRotType,proto3,enum=BeginCameraSceneLookNotify_KeepRotType" json:"keep_rot_type,omitempty"`
	FFHJPJBDLJC bool                                   `protobuf:"varint,1,opt,name=FFHJPJBDLJC,proto3" json:"FFHJPJBDLJC,omitempty"`
	Duration    float32                                `protobuf:"fixed32,6,opt,name=duration,proto3" json:"duration,omitempty"`
	GEBMAIIKMBJ uint32                                 `protobuf:"varint,54,opt,name=GEBMAIIKMBJ,proto3" json:"GEBMAIIKMBJ,omitempty"`
	MCFGAPGPNLN bool                                   `protobuf:"varint,13,opt,name=MCFGAPGPNLN,proto3" json:"MCFGAPGPNLN,omitempty"`
	OJKGLFJCMID float32                                `protobuf:"fixed32,801,opt,name=OJKGLFJCMID,proto3" json:"OJKGLFJCMID,omitempty"`
}

func (x *BeginCameraSceneLookNotify) Reset() {
	*x = BeginCameraSceneLookNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_BeginCameraSceneLookNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BeginCameraSceneLookNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BeginCameraSceneLookNotify) ProtoMessage() {}

func (x *BeginCameraSceneLookNotify) ProtoReflect() protoreflect.Message {
	mi := &file_proto_BeginCameraSceneLookNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BeginCameraSceneLookNotify.ProtoReflect.Descriptor instead.
func (*BeginCameraSceneLookNotify) Descriptor() ([]byte, []int) {
	return file_proto_BeginCameraSceneLookNotify_proto_rawDescGZIP(), []int{0}
}

func (x *BeginCameraSceneLookNotify) GetEntityId() uint32 {
	if x != nil {
		return x.EntityId
	}
	return 0
}

func (x *BeginCameraSceneLookNotify) GetJBCEDEGCGFJ() bool {
	if x != nil {
		return x.JBCEDEGCGFJ
	}
	return false
}

func (x *BeginCameraSceneLookNotify) GetADPLABBOBKF() bool {
	if x != nil {
		return x.ADPLABBOBKF
	}
	return false
}

func (x *BeginCameraSceneLookNotify) GetDOCBFPIPMPK() uint32 {
	if x != nil {
		return x.DOCBFPIPMPK
	}
	return 0
}

func (x *BeginCameraSceneLookNotify) GetLKGGBPLCEJI() *Vector {
	if x != nil {
		return x.LKGGBPLCEJI
	}
	return nil
}

func (x *BeginCameraSceneLookNotify) GetKJHLOCKMIMB() bool {
	if x != nil {
		return x.KJHLOCKMIMB
	}
	return false
}

func (x *BeginCameraSceneLookNotify) GetGJNLEMGJCJI() float32 {
	if x != nil {
		return x.GJNLEMGJCJI
	}
	return 0
}

func (x *BeginCameraSceneLookNotify) GetIPKFLPFLNNI() bool {
	if x != nil {
		return x.IPKFLPFLNNI
	}
	return false
}

func (x *BeginCameraSceneLookNotify) GetPDOCOOCFAIH() bool {
	if x != nil {
		return x.PDOCOOCFAIH
	}
	return false
}

func (x *BeginCameraSceneLookNotify) GetMMOMOKPCOJK() *Vector {
	if x != nil {
		return x.MMOMOKPCOJK
	}
	return nil
}

func (x *BeginCameraSceneLookNotify) GetFPDPEHICGAC() bool {
	if x != nil {
		return x.FPDPEHICGAC
	}
	return false
}

func (x *BeginCameraSceneLookNotify) GetMIKBDFOBIBB() float32 {
	if x != nil {
		return x.MIKBDFOBIBB
	}
	return 0
}

func (x *BeginCameraSceneLookNotify) GetCLLOFLJICAF() float32 {
	if x != nil {
		return x.CLLOFLJICAF
	}
	return 0
}

func (x *BeginCameraSceneLookNotify) GetOtherParams() []string {
	if x != nil {
		return x.OtherParams
	}
	return nil
}

func (x *BeginCameraSceneLookNotify) GetKeepRotType() BeginCameraSceneLookNotify_KeepRotType {
	if x != nil {
		return x.KeepRotType
	}
	return BeginCameraSceneLookNotify_KEEP_ROT_X
}

func (x *BeginCameraSceneLookNotify) GetFFHJPJBDLJC() bool {
	if x != nil {
		return x.FFHJPJBDLJC
	}
	return false
}

func (x *BeginCameraSceneLookNotify) GetDuration() float32 {
	if x != nil {
		return x.Duration
	}
	return 0
}

func (x *BeginCameraSceneLookNotify) GetGEBMAIIKMBJ() uint32 {
	if x != nil {
		return x.GEBMAIIKMBJ
	}
	return 0
}

func (x *BeginCameraSceneLookNotify) GetMCFGAPGPNLN() bool {
	if x != nil {
		return x.MCFGAPGPNLN
	}
	return false
}

func (x *BeginCameraSceneLookNotify) GetOJKGLFJCMID() float32 {
	if x != nil {
		return x.OJKGLFJCMID
	}
	return 0
}

var File_proto_BeginCameraSceneLookNotify_proto protoreflect.FileDescriptor

var file_proto_BeginCameraSceneLookNotify_proto_rawDesc = []byte{
	0x0a, 0x26, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x42, 0x65, 0x67, 0x69, 0x6e, 0x43, 0x61, 0x6d,
	0x65, 0x72, 0x61, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x4c, 0x6f, 0x6f, 0x6b, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xab, 0x06, 0x0a,
	0x1a, 0x42, 0x65, 0x67, 0x69, 0x6e, 0x43, 0x61, 0x6d, 0x65, 0x72, 0x61, 0x53, 0x63, 0x65, 0x6e,
	0x65, 0x4c, 0x6f, 0x6f, 0x6b, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0xe1, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x08, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x4a, 0x42, 0x43,
	0x45, 0x44, 0x45, 0x47, 0x43, 0x47, 0x46, 0x4a, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b,
	0x4a, 0x42, 0x43, 0x45, 0x44, 0x45, 0x47, 0x43, 0x47, 0x46, 0x4a, 0x12, 0x20, 0x0a, 0x0b, 0x41,
	0x44, 0x50, 0x4c, 0x41, 0x42, 0x42, 0x4f, 0x42, 0x4b, 0x46, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0b, 0x41, 0x44, 0x50, 0x4c, 0x41, 0x42, 0x42, 0x4f, 0x42, 0x4b, 0x46, 0x12, 0x21, 0x0a,
	0x0b, 0x44, 0x4f, 0x43, 0x42, 0x46, 0x50, 0x49, 0x50, 0x4d, 0x50, 0x4b, 0x18, 0xee, 0x07, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0b, 0x44, 0x4f, 0x43, 0x42, 0x46, 0x50, 0x49, 0x50, 0x4d, 0x50, 0x4b,
	0x12, 0x29, 0x0a, 0x0b, 0x4c, 0x4b, 0x47, 0x47, 0x42, 0x50, 0x4c, 0x43, 0x45, 0x4a, 0x49, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x0b,
	0x4c, 0x4b, 0x47, 0x47, 0x42, 0x50, 0x4c, 0x43, 0x45, 0x4a, 0x49, 0x12, 0x21, 0x0a, 0x0b, 0x4b,
	0x4a, 0x48, 0x4c, 0x4f, 0x43, 0x4b, 0x4d, 0x49, 0x4d, 0x42, 0x18, 0xc5, 0x0b, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0b, 0x4b, 0x4a, 0x48, 0x4c, 0x4f, 0x43, 0x4b, 0x4d, 0x49, 0x4d, 0x42, 0x12, 0x20,
	0x0a, 0x0b, 0x47, 0x4a, 0x4e, 0x4c, 0x45, 0x4d, 0x47, 0x4a, 0x43, 0x4a, 0x49, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x0b, 0x47, 0x4a, 0x4e, 0x4c, 0x45, 0x4d, 0x47, 0x4a, 0x43, 0x4a, 0x49,
	0x12, 0x20, 0x0a, 0x0b, 0x49, 0x50, 0x4b, 0x46, 0x4c, 0x50, 0x46, 0x4c, 0x4e, 0x4e, 0x49, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x49, 0x50, 0x4b, 0x46, 0x4c, 0x50, 0x46, 0x4c, 0x4e,
	0x4e, 0x49, 0x12, 0x20, 0x0a, 0x0b, 0x50, 0x44, 0x4f, 0x43, 0x4f, 0x4f, 0x43, 0x46, 0x41, 0x49,
	0x48, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x50, 0x44, 0x4f, 0x43, 0x4f, 0x4f, 0x43,
	0x46, 0x41, 0x49, 0x48, 0x12, 0x29, 0x0a, 0x0b, 0x4d, 0x4d, 0x4f, 0x4d, 0x4f, 0x4b, 0x50, 0x43,
	0x4f, 0x4a, 0x4b, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x56, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x52, 0x0b, 0x4d, 0x4d, 0x4f, 0x4d, 0x4f, 0x4b, 0x50, 0x43, 0x4f, 0x4a, 0x4b, 0x12,
	0x20, 0x0a, 0x0b, 0x46, 0x50, 0x44, 0x50, 0x45, 0x48, 0x49, 0x43, 0x47, 0x41, 0x43, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x46, 0x50, 0x44, 0x50, 0x45, 0x48, 0x49, 0x43, 0x47, 0x41,
	0x43, 0x12, 0x20, 0x0a, 0x0b, 0x4d, 0x49, 0x4b, 0x42, 0x44, 0x46, 0x4f, 0x42, 0x49, 0x42, 0x42,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b, 0x4d, 0x49, 0x4b, 0x42, 0x44, 0x46, 0x4f, 0x42,
	0x49, 0x42, 0x42, 0x12, 0x20, 0x0a, 0x0b, 0x43, 0x4c, 0x4c, 0x4f, 0x46, 0x4c, 0x4a, 0x49, 0x43,
	0x41, 0x46, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b, 0x43, 0x4c, 0x4c, 0x4f, 0x46, 0x4c,
	0x4a, 0x49, 0x43, 0x41, 0x46, 0x12, 0x21, 0x0a, 0x0c, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x5f, 0x70,
	0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x6f, 0x74, 0x68,
	0x65, 0x72, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x4b, 0x0a, 0x0d, 0x6b, 0x65, 0x65, 0x70,
	0x5f, 0x72, 0x6f, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x27, 0x2e, 0x42, 0x65, 0x67, 0x69, 0x6e, 0x43, 0x61, 0x6d, 0x65, 0x72, 0x61, 0x53, 0x63, 0x65,
	0x6e, 0x65, 0x4c, 0x6f, 0x6f, 0x6b, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x4b, 0x65, 0x65,
	0x70, 0x52, 0x6f, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0b, 0x6b, 0x65, 0x65, 0x70, 0x52, 0x6f,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x46, 0x46, 0x48, 0x4a, 0x50, 0x4a, 0x42,
	0x44, 0x4c, 0x4a, 0x43, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x46, 0x46, 0x48, 0x4a,
	0x50, 0x4a, 0x42, 0x44, 0x4c, 0x4a, 0x43, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x47, 0x45, 0x42, 0x4d, 0x41, 0x49, 0x49, 0x4b, 0x4d,
	0x42, 0x4a, 0x18, 0x36, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x47, 0x45, 0x42, 0x4d, 0x41, 0x49,
	0x49, 0x4b, 0x4d, 0x42, 0x4a, 0x12, 0x20, 0x0a, 0x0b, 0x4d, 0x43, 0x46, 0x47, 0x41, 0x50, 0x47,
	0x50, 0x4e, 0x4c, 0x4e, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x4d, 0x43, 0x46, 0x47,
	0x41, 0x50, 0x47, 0x50, 0x4e, 0x4c, 0x4e, 0x12, 0x21, 0x0a, 0x0b, 0x4f, 0x4a, 0x4b, 0x47, 0x4c,
	0x46, 0x4a, 0x43, 0x4d, 0x49, 0x44, 0x18, 0xa1, 0x06, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b, 0x4f,
	0x4a, 0x4b, 0x47, 0x4c, 0x46, 0x4a, 0x43, 0x4d, 0x49, 0x44, 0x22, 0x2e, 0x0a, 0x0b, 0x4b, 0x65,
	0x65, 0x70, 0x52, 0x6f, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0e, 0x0a, 0x0a, 0x4b, 0x45, 0x45,
	0x50, 0x5f, 0x52, 0x4f, 0x54, 0x5f, 0x58, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x4b, 0x45, 0x45,
	0x50, 0x5f, 0x52, 0x4f, 0x54, 0x5f, 0x58, 0x59, 0x10, 0x01, 0x42, 0x1d, 0x5a, 0x1b, 0x47, 0x6f,
	0x2d, 0x47, 0x72, 0x61, 0x73, 0x73, 0x63, 0x75, 0x74, 0x74, 0x65, 0x72, 0x2f, 0x67, 0x65, 0x6e,
	0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_proto_BeginCameraSceneLookNotify_proto_rawDescOnce sync.Once
	file_proto_BeginCameraSceneLookNotify_proto_rawDescData = file_proto_BeginCameraSceneLookNotify_proto_rawDesc
)

func file_proto_BeginCameraSceneLookNotify_proto_rawDescGZIP() []byte {
	file_proto_BeginCameraSceneLookNotify_proto_rawDescOnce.Do(func() {
		file_proto_BeginCameraSceneLookNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_BeginCameraSceneLookNotify_proto_rawDescData)
	})
	return file_proto_BeginCameraSceneLookNotify_proto_rawDescData
}

var file_proto_BeginCameraSceneLookNotify_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_BeginCameraSceneLookNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_BeginCameraSceneLookNotify_proto_goTypes = []interface{}{
	(BeginCameraSceneLookNotify_KeepRotType)(0), // 0: BeginCameraSceneLookNotify.KeepRotType
	(*BeginCameraSceneLookNotify)(nil),          // 1: BeginCameraSceneLookNotify
	(*Vector)(nil),                              // 2: Vector
}
var file_proto_BeginCameraSceneLookNotify_proto_depIdxs = []int32{
	2, // 0: BeginCameraSceneLookNotify.LKGGBPLCEJI:type_name -> Vector
	2, // 1: BeginCameraSceneLookNotify.MMOMOKPCOJK:type_name -> Vector
	0, // 2: BeginCameraSceneLookNotify.keep_rot_type:type_name -> BeginCameraSceneLookNotify.KeepRotType
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_proto_BeginCameraSceneLookNotify_proto_init() }
func file_proto_BeginCameraSceneLookNotify_proto_init() {
	if File_proto_BeginCameraSceneLookNotify_proto != nil {
		return
	}
	file_proto_Vector_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_BeginCameraSceneLookNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BeginCameraSceneLookNotify); i {
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
			RawDescriptor: file_proto_BeginCameraSceneLookNotify_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_BeginCameraSceneLookNotify_proto_goTypes,
		DependencyIndexes: file_proto_BeginCameraSceneLookNotify_proto_depIdxs,
		EnumInfos:         file_proto_BeginCameraSceneLookNotify_proto_enumTypes,
		MessageInfos:      file_proto_BeginCameraSceneLookNotify_proto_msgTypes,
	}.Build()
	File_proto_BeginCameraSceneLookNotify_proto = out.File
	file_proto_BeginCameraSceneLookNotify_proto_rawDesc = nil
	file_proto_BeginCameraSceneLookNotify_proto_goTypes = nil
	file_proto_BeginCameraSceneLookNotify_proto_depIdxs = nil
}
