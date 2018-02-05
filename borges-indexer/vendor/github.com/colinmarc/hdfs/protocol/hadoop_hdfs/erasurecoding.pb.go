// Code generated by protoc-gen-go.
// source: erasurecoding.proto
// DO NOT EDIT!

package hadoop_hdfs

import proto "github.com/golang/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type SetErasureCodingPolicyRequestProto struct {
	Src              *string                   `protobuf:"bytes,1,req,name=src" json:"src,omitempty"`
	EcPolicy         *ErasureCodingPolicyProto `protobuf:"bytes,2,opt,name=ecPolicy" json:"ecPolicy,omitempty"`
	XXX_unrecognized []byte                    `json:"-"`
}

func (m *SetErasureCodingPolicyRequestProto) Reset()         { *m = SetErasureCodingPolicyRequestProto{} }
func (m *SetErasureCodingPolicyRequestProto) String() string { return proto.CompactTextString(m) }
func (*SetErasureCodingPolicyRequestProto) ProtoMessage()    {}

func (m *SetErasureCodingPolicyRequestProto) GetSrc() string {
	if m != nil && m.Src != nil {
		return *m.Src
	}
	return ""
}

func (m *SetErasureCodingPolicyRequestProto) GetEcPolicy() *ErasureCodingPolicyProto {
	if m != nil {
		return m.EcPolicy
	}
	return nil
}

type SetErasureCodingPolicyResponseProto struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *SetErasureCodingPolicyResponseProto) Reset()         { *m = SetErasureCodingPolicyResponseProto{} }
func (m *SetErasureCodingPolicyResponseProto) String() string { return proto.CompactTextString(m) }
func (*SetErasureCodingPolicyResponseProto) ProtoMessage()    {}

type GetErasureCodingPoliciesRequestProto struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *GetErasureCodingPoliciesRequestProto) Reset()         { *m = GetErasureCodingPoliciesRequestProto{} }
func (m *GetErasureCodingPoliciesRequestProto) String() string { return proto.CompactTextString(m) }
func (*GetErasureCodingPoliciesRequestProto) ProtoMessage()    {}

type GetErasureCodingPoliciesResponseProto struct {
	EcPolicies       []*ErasureCodingPolicyProto `protobuf:"bytes,1,rep,name=ecPolicies" json:"ecPolicies,omitempty"`
	XXX_unrecognized []byte                      `json:"-"`
}

func (m *GetErasureCodingPoliciesResponseProto) Reset()         { *m = GetErasureCodingPoliciesResponseProto{} }
func (m *GetErasureCodingPoliciesResponseProto) String() string { return proto.CompactTextString(m) }
func (*GetErasureCodingPoliciesResponseProto) ProtoMessage()    {}

func (m *GetErasureCodingPoliciesResponseProto) GetEcPolicies() []*ErasureCodingPolicyProto {
	if m != nil {
		return m.EcPolicies
	}
	return nil
}

type GetErasureCodingPolicyRequestProto struct {
	Src              *string `protobuf:"bytes,1,req,name=src" json:"src,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *GetErasureCodingPolicyRequestProto) Reset()         { *m = GetErasureCodingPolicyRequestProto{} }
func (m *GetErasureCodingPolicyRequestProto) String() string { return proto.CompactTextString(m) }
func (*GetErasureCodingPolicyRequestProto) ProtoMessage()    {}

func (m *GetErasureCodingPolicyRequestProto) GetSrc() string {
	if m != nil && m.Src != nil {
		return *m.Src
	}
	return ""
}

type GetErasureCodingPolicyResponseProto struct {
	EcPolicy         *ErasureCodingPolicyProto `protobuf:"bytes,1,opt,name=ecPolicy" json:"ecPolicy,omitempty"`
	XXX_unrecognized []byte                    `json:"-"`
}

func (m *GetErasureCodingPolicyResponseProto) Reset()         { *m = GetErasureCodingPolicyResponseProto{} }
func (m *GetErasureCodingPolicyResponseProto) String() string { return proto.CompactTextString(m) }
func (*GetErasureCodingPolicyResponseProto) ProtoMessage()    {}

func (m *GetErasureCodingPolicyResponseProto) GetEcPolicy() *ErasureCodingPolicyProto {
	if m != nil {
		return m.EcPolicy
	}
	return nil
}

// *
// Block erasure coding reconstruction info
type BlockECReconstructionInfoProto struct {
	Block              *ExtendedBlockProto       `protobuf:"bytes,1,req,name=block" json:"block,omitempty"`
	SourceDnInfos      *DatanodeInfosProto       `protobuf:"bytes,2,req,name=sourceDnInfos" json:"sourceDnInfos,omitempty"`
	TargetDnInfos      *DatanodeInfosProto       `protobuf:"bytes,3,req,name=targetDnInfos" json:"targetDnInfos,omitempty"`
	TargetStorageUuids *StorageUuidsProto        `protobuf:"bytes,4,req,name=targetStorageUuids" json:"targetStorageUuids,omitempty"`
	TargetStorageTypes *StorageTypesProto        `protobuf:"bytes,5,req,name=targetStorageTypes" json:"targetStorageTypes,omitempty"`
	LiveBlockIndices   []byte                    `protobuf:"bytes,6,req,name=liveBlockIndices" json:"liveBlockIndices,omitempty"`
	EcPolicy           *ErasureCodingPolicyProto `protobuf:"bytes,7,req,name=ecPolicy" json:"ecPolicy,omitempty"`
	XXX_unrecognized   []byte                    `json:"-"`
}

func (m *BlockECReconstructionInfoProto) Reset()         { *m = BlockECReconstructionInfoProto{} }
func (m *BlockECReconstructionInfoProto) String() string { return proto.CompactTextString(m) }
func (*BlockECReconstructionInfoProto) ProtoMessage()    {}

func (m *BlockECReconstructionInfoProto) GetBlock() *ExtendedBlockProto {
	if m != nil {
		return m.Block
	}
	return nil
}

func (m *BlockECReconstructionInfoProto) GetSourceDnInfos() *DatanodeInfosProto {
	if m != nil {
		return m.SourceDnInfos
	}
	return nil
}

func (m *BlockECReconstructionInfoProto) GetTargetDnInfos() *DatanodeInfosProto {
	if m != nil {
		return m.TargetDnInfos
	}
	return nil
}

func (m *BlockECReconstructionInfoProto) GetTargetStorageUuids() *StorageUuidsProto {
	if m != nil {
		return m.TargetStorageUuids
	}
	return nil
}

func (m *BlockECReconstructionInfoProto) GetTargetStorageTypes() *StorageTypesProto {
	if m != nil {
		return m.TargetStorageTypes
	}
	return nil
}

func (m *BlockECReconstructionInfoProto) GetLiveBlockIndices() []byte {
	if m != nil {
		return m.LiveBlockIndices
	}
	return nil
}

func (m *BlockECReconstructionInfoProto) GetEcPolicy() *ErasureCodingPolicyProto {
	if m != nil {
		return m.EcPolicy
	}
	return nil
}

func init() {
}
