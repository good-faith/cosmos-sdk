// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: tendermint/p2p/types.proto

package p2p

import (
	_ "github.com/gogo/protobuf/gogoproto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ProtocolVersion struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	P2P   uint64 `protobuf:"varint,1,opt,name=p2p,proto3" json:"p2p,omitempty"`
	Block uint64 `protobuf:"varint,2,opt,name=block,proto3" json:"block,omitempty"`
	App   uint64 `protobuf:"varint,3,opt,name=app,proto3" json:"app,omitempty"`
}

func (x *ProtocolVersion) Reset() {
	*x = ProtocolVersion{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tendermint_p2p_types_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProtocolVersion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProtocolVersion) ProtoMessage() {}

func (x *ProtocolVersion) ProtoReflect() protoreflect.Message {
	mi := &file_tendermint_p2p_types_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProtocolVersion.ProtoReflect.Descriptor instead.
func (*ProtocolVersion) Descriptor() ([]byte, []int) {
	return file_tendermint_p2p_types_proto_rawDescGZIP(), []int{0}
}

func (x *ProtocolVersion) GetP2P() uint64 {
	if x != nil {
		return x.P2P
	}
	return 0
}

func (x *ProtocolVersion) GetBlock() uint64 {
	if x != nil {
		return x.Block
	}
	return 0
}

func (x *ProtocolVersion) GetApp() uint64 {
	if x != nil {
		return x.App
	}
	return 0
}

type NodeInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProtocolVersion *ProtocolVersion `protobuf:"bytes,1,opt,name=protocol_version,json=protocolVersion,proto3" json:"protocol_version,omitempty"`
	NodeId          string           `protobuf:"bytes,2,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	ListenAddr      string           `protobuf:"bytes,3,opt,name=listen_addr,json=listenAddr,proto3" json:"listen_addr,omitempty"`
	Network         string           `protobuf:"bytes,4,opt,name=network,proto3" json:"network,omitempty"`
	Version         string           `protobuf:"bytes,5,opt,name=version,proto3" json:"version,omitempty"`
	Channels        []byte           `protobuf:"bytes,6,opt,name=channels,proto3" json:"channels,omitempty"`
	Moniker         string           `protobuf:"bytes,7,opt,name=moniker,proto3" json:"moniker,omitempty"`
	Other           *NodeInfoOther   `protobuf:"bytes,8,opt,name=other,proto3" json:"other,omitempty"`
}

func (x *NodeInfo) Reset() {
	*x = NodeInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tendermint_p2p_types_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeInfo) ProtoMessage() {}

func (x *NodeInfo) ProtoReflect() protoreflect.Message {
	mi := &file_tendermint_p2p_types_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeInfo.ProtoReflect.Descriptor instead.
func (*NodeInfo) Descriptor() ([]byte, []int) {
	return file_tendermint_p2p_types_proto_rawDescGZIP(), []int{1}
}

func (x *NodeInfo) GetProtocolVersion() *ProtocolVersion {
	if x != nil {
		return x.ProtocolVersion
	}
	return nil
}

func (x *NodeInfo) GetNodeId() string {
	if x != nil {
		return x.NodeId
	}
	return ""
}

func (x *NodeInfo) GetListenAddr() string {
	if x != nil {
		return x.ListenAddr
	}
	return ""
}

func (x *NodeInfo) GetNetwork() string {
	if x != nil {
		return x.Network
	}
	return ""
}

func (x *NodeInfo) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *NodeInfo) GetChannels() []byte {
	if x != nil {
		return x.Channels
	}
	return nil
}

func (x *NodeInfo) GetMoniker() string {
	if x != nil {
		return x.Moniker
	}
	return ""
}

func (x *NodeInfo) GetOther() *NodeInfoOther {
	if x != nil {
		return x.Other
	}
	return nil
}

type NodeInfoOther struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TxIndex    string `protobuf:"bytes,1,opt,name=tx_index,json=txIndex,proto3" json:"tx_index,omitempty"`
	RpcAddress string `protobuf:"bytes,2,opt,name=rpc_address,json=rpcAddress,proto3" json:"rpc_address,omitempty"`
}

func (x *NodeInfoOther) Reset() {
	*x = NodeInfoOther{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tendermint_p2p_types_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeInfoOther) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeInfoOther) ProtoMessage() {}

func (x *NodeInfoOther) ProtoReflect() protoreflect.Message {
	mi := &file_tendermint_p2p_types_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeInfoOther.ProtoReflect.Descriptor instead.
func (*NodeInfoOther) Descriptor() ([]byte, []int) {
	return file_tendermint_p2p_types_proto_rawDescGZIP(), []int{2}
}

func (x *NodeInfoOther) GetTxIndex() string {
	if x != nil {
		return x.TxIndex
	}
	return ""
}

func (x *NodeInfoOther) GetRpcAddress() string {
	if x != nil {
		return x.RpcAddress
	}
	return ""
}

type PeerInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	AddressInfo   []*PeerAddressInfo     `protobuf:"bytes,2,rep,name=address_info,json=addressInfo,proto3" json:"address_info,omitempty"`
	LastConnected *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=last_connected,json=lastConnected,proto3" json:"last_connected,omitempty"`
}

func (x *PeerInfo) Reset() {
	*x = PeerInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tendermint_p2p_types_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PeerInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PeerInfo) ProtoMessage() {}

func (x *PeerInfo) ProtoReflect() protoreflect.Message {
	mi := &file_tendermint_p2p_types_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PeerInfo.ProtoReflect.Descriptor instead.
func (*PeerInfo) Descriptor() ([]byte, []int) {
	return file_tendermint_p2p_types_proto_rawDescGZIP(), []int{3}
}

func (x *PeerInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PeerInfo) GetAddressInfo() []*PeerAddressInfo {
	if x != nil {
		return x.AddressInfo
	}
	return nil
}

func (x *PeerInfo) GetLastConnected() *timestamppb.Timestamp {
	if x != nil {
		return x.LastConnected
	}
	return nil
}

type PeerAddressInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address         string                 `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	LastDialSuccess *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=last_dial_success,json=lastDialSuccess,proto3" json:"last_dial_success,omitempty"`
	LastDialFailure *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=last_dial_failure,json=lastDialFailure,proto3" json:"last_dial_failure,omitempty"`
	DialFailures    uint32                 `protobuf:"varint,4,opt,name=dial_failures,json=dialFailures,proto3" json:"dial_failures,omitempty"`
}

func (x *PeerAddressInfo) Reset() {
	*x = PeerAddressInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tendermint_p2p_types_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PeerAddressInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PeerAddressInfo) ProtoMessage() {}

func (x *PeerAddressInfo) ProtoReflect() protoreflect.Message {
	mi := &file_tendermint_p2p_types_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PeerAddressInfo.ProtoReflect.Descriptor instead.
func (*PeerAddressInfo) Descriptor() ([]byte, []int) {
	return file_tendermint_p2p_types_proto_rawDescGZIP(), []int{4}
}

func (x *PeerAddressInfo) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *PeerAddressInfo) GetLastDialSuccess() *timestamppb.Timestamp {
	if x != nil {
		return x.LastDialSuccess
	}
	return nil
}

func (x *PeerAddressInfo) GetLastDialFailure() *timestamppb.Timestamp {
	if x != nil {
		return x.LastDialFailure
	}
	return nil
}

func (x *PeerAddressInfo) GetDialFailures() uint32 {
	if x != nil {
		return x.DialFailures
	}
	return 0
}

var File_tendermint_p2p_types_proto protoreflect.FileDescriptor

var file_tendermint_p2p_types_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2f, 0x70, 0x32, 0x70,
	0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x74, 0x65,
	0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2e, 0x70, 0x32, 0x70, 0x1a, 0x14, 0x67, 0x6f,
	0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x54, 0x0a, 0x0f, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x56,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x19, 0x0a, 0x03, 0x70, 0x32, 0x70, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x42, 0x07, 0xe2, 0xde, 0x1f, 0x03, 0x50, 0x32, 0x50, 0x52, 0x03, 0x70, 0x32,
	0x70, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x05, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x70, 0x70, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x61, 0x70, 0x70, 0x22, 0xc7, 0x02, 0x0a, 0x08, 0x4e, 0x6f,
	0x64, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x50, 0x0a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1f, 0x2e, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2e, 0x70, 0x32,
	0x70, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x42, 0x04, 0xc8, 0xde, 0x1f, 0x00, 0x52, 0x0f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x23, 0x0a, 0x07, 0x6e, 0x6f, 0x64, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0xe2, 0xde, 0x1f, 0x06, 0x4e,
	0x6f, 0x64, 0x65, 0x49, 0x44, 0x52, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x12, 0x1f, 0x0a,
	0x0b, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x41, 0x64, 0x64, 0x72, 0x12, 0x18,
	0x0a, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x73, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x73, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x6f, 0x6e, 0x69, 0x6b, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x6f, 0x6e, 0x69, 0x6b, 0x65, 0x72, 0x12, 0x39, 0x0a, 0x05, 0x6f, 0x74, 0x68, 0x65,
	0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72,
	0x6d, 0x69, 0x6e, 0x74, 0x2e, 0x70, 0x32, 0x70, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x4f, 0x74, 0x68, 0x65, 0x72, 0x42, 0x04, 0xc8, 0xde, 0x1f, 0x00, 0x52, 0x05, 0x6f, 0x74,
	0x68, 0x65, 0x72, 0x22, 0x5b, 0x0a, 0x0d, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x4f,
	0x74, 0x68, 0x65, 0x72, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x78, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x78, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12,
	0x2f, 0x0a, 0x0b, 0x72, 0x70, 0x63, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x0e, 0xe2, 0xde, 0x1f, 0x0a, 0x52, 0x50, 0x43, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x52, 0x0a, 0x72, 0x70, 0x63, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x22, 0xaf, 0x01, 0x0a, 0x08, 0x50, 0x65, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x16, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xe2, 0xde, 0x1f, 0x02, 0x49,
	0x44, 0x52, 0x02, 0x69, 0x64, 0x12, 0x42, 0x0a, 0x0c, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x74, 0x65,
	0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2e, 0x70, 0x32, 0x70, 0x2e, 0x50, 0x65, 0x65,
	0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0b, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x47, 0x0a, 0x0e, 0x6c, 0x61, 0x73,
	0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x04, 0x90,
	0xdf, 0x1f, 0x01, 0x52, 0x0d, 0x6c, 0x61, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x65, 0x64, 0x22, 0xec, 0x01, 0x0a, 0x0f, 0x50, 0x65, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x12, 0x4c, 0x0a, 0x11, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x64, 0x69, 0x61, 0x6c, 0x5f, 0x73, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x04, 0x90, 0xdf, 0x1f, 0x01, 0x52, 0x0f, 0x6c,
	0x61, 0x73, 0x74, 0x44, 0x69, 0x61, 0x6c, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x4c,
	0x0a, 0x11, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x64, 0x69, 0x61, 0x6c, 0x5f, 0x66, 0x61, 0x69, 0x6c,
	0x75, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x04, 0x90, 0xdf, 0x1f, 0x01, 0x52, 0x0f, 0x6c, 0x61, 0x73,
	0x74, 0x44, 0x69, 0x61, 0x6c, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x12, 0x23, 0x0a, 0x0d,
	0x64, 0x69, 0x61, 0x6c, 0x5f, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x73, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0c, 0x64, 0x69, 0x61, 0x6c, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65,
	0x73, 0x42, 0xaa, 0x01, 0x0a, 0x12, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72,
	0x6d, 0x69, 0x6e, 0x74, 0x2e, 0x70, 0x32, 0x70, 0x42, 0x0a, 0x54, 0x79, 0x70, 0x65, 0x73, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73,
	0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d,
	0x69, 0x6e, 0x74, 0x2f, 0x70, 0x32, 0x70, 0xa2, 0x02, 0x03, 0x54, 0x50, 0x58, 0xaa, 0x02, 0x0e,
	0x54, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2e, 0x50, 0x32, 0x70, 0xca, 0x02,
	0x0e, 0x54, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x5c, 0x50, 0x32, 0x70, 0xe2,
	0x02, 0x1a, 0x54, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x5c, 0x50, 0x32, 0x70,
	0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0f, 0x54,
	0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x3a, 0x3a, 0x50, 0x32, 0x70, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tendermint_p2p_types_proto_rawDescOnce sync.Once
	file_tendermint_p2p_types_proto_rawDescData = file_tendermint_p2p_types_proto_rawDesc
)

func file_tendermint_p2p_types_proto_rawDescGZIP() []byte {
	file_tendermint_p2p_types_proto_rawDescOnce.Do(func() {
		file_tendermint_p2p_types_proto_rawDescData = protoimpl.X.CompressGZIP(file_tendermint_p2p_types_proto_rawDescData)
	})
	return file_tendermint_p2p_types_proto_rawDescData
}

var file_tendermint_p2p_types_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_tendermint_p2p_types_proto_goTypes = []interface{}{
	(*ProtocolVersion)(nil),       // 0: tendermint.p2p.ProtocolVersion
	(*NodeInfo)(nil),              // 1: tendermint.p2p.NodeInfo
	(*NodeInfoOther)(nil),         // 2: tendermint.p2p.NodeInfoOther
	(*PeerInfo)(nil),              // 3: tendermint.p2p.PeerInfo
	(*PeerAddressInfo)(nil),       // 4: tendermint.p2p.PeerAddressInfo
	(*timestamppb.Timestamp)(nil), // 5: google.protobuf.Timestamp
}
var file_tendermint_p2p_types_proto_depIdxs = []int32{
	0, // 0: tendermint.p2p.NodeInfo.protocol_version:type_name -> tendermint.p2p.ProtocolVersion
	2, // 1: tendermint.p2p.NodeInfo.other:type_name -> tendermint.p2p.NodeInfoOther
	4, // 2: tendermint.p2p.PeerInfo.address_info:type_name -> tendermint.p2p.PeerAddressInfo
	5, // 3: tendermint.p2p.PeerInfo.last_connected:type_name -> google.protobuf.Timestamp
	5, // 4: tendermint.p2p.PeerAddressInfo.last_dial_success:type_name -> google.protobuf.Timestamp
	5, // 5: tendermint.p2p.PeerAddressInfo.last_dial_failure:type_name -> google.protobuf.Timestamp
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_tendermint_p2p_types_proto_init() }
func file_tendermint_p2p_types_proto_init() {
	if File_tendermint_p2p_types_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tendermint_p2p_types_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProtocolVersion); i {
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
		file_tendermint_p2p_types_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeInfo); i {
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
		file_tendermint_p2p_types_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeInfoOther); i {
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
		file_tendermint_p2p_types_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PeerInfo); i {
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
		file_tendermint_p2p_types_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PeerAddressInfo); i {
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
			RawDescriptor: file_tendermint_p2p_types_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tendermint_p2p_types_proto_goTypes,
		DependencyIndexes: file_tendermint_p2p_types_proto_depIdxs,
		MessageInfos:      file_tendermint_p2p_types_proto_msgTypes,
	}.Build()
	File_tendermint_p2p_types_proto = out.File
	file_tendermint_p2p_types_proto_rawDesc = nil
	file_tendermint_p2p_types_proto_goTypes = nil
	file_tendermint_p2p_types_proto_depIdxs = nil
}