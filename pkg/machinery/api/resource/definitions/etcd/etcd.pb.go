// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v6.30.1
// source: resource/definitions/etcd/etcd.proto

package etcd

import (
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"

	common "github.com/siderolabs/talos/pkg/machinery/api/common"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// ConfigSpec describes (some) configuration settings of etcd.
type ConfigSpec struct {
	state                   protoimpl.MessageState `protogen:"open.v1"`
	AdvertiseValidSubnets   []string               `protobuf:"bytes,1,rep,name=advertise_valid_subnets,json=advertiseValidSubnets,proto3" json:"advertise_valid_subnets,omitempty"`
	AdvertiseExcludeSubnets []string               `protobuf:"bytes,2,rep,name=advertise_exclude_subnets,json=advertiseExcludeSubnets,proto3" json:"advertise_exclude_subnets,omitempty"`
	Image                   string                 `protobuf:"bytes,3,opt,name=image,proto3" json:"image,omitempty"`
	ExtraArgs               map[string]string      `protobuf:"bytes,4,rep,name=extra_args,json=extraArgs,proto3" json:"extra_args,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	ListenValidSubnets      []string               `protobuf:"bytes,5,rep,name=listen_valid_subnets,json=listenValidSubnets,proto3" json:"listen_valid_subnets,omitempty"`
	ListenExcludeSubnets    []string               `protobuf:"bytes,6,rep,name=listen_exclude_subnets,json=listenExcludeSubnets,proto3" json:"listen_exclude_subnets,omitempty"`
	unknownFields           protoimpl.UnknownFields
	sizeCache               protoimpl.SizeCache
}

func (x *ConfigSpec) Reset() {
	*x = ConfigSpec{}
	mi := &file_resource_definitions_etcd_etcd_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ConfigSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigSpec) ProtoMessage() {}

func (x *ConfigSpec) ProtoReflect() protoreflect.Message {
	mi := &file_resource_definitions_etcd_etcd_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigSpec.ProtoReflect.Descriptor instead.
func (*ConfigSpec) Descriptor() ([]byte, []int) {
	return file_resource_definitions_etcd_etcd_proto_rawDescGZIP(), []int{0}
}

func (x *ConfigSpec) GetAdvertiseValidSubnets() []string {
	if x != nil {
		return x.AdvertiseValidSubnets
	}
	return nil
}

func (x *ConfigSpec) GetAdvertiseExcludeSubnets() []string {
	if x != nil {
		return x.AdvertiseExcludeSubnets
	}
	return nil
}

func (x *ConfigSpec) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *ConfigSpec) GetExtraArgs() map[string]string {
	if x != nil {
		return x.ExtraArgs
	}
	return nil
}

func (x *ConfigSpec) GetListenValidSubnets() []string {
	if x != nil {
		return x.ListenValidSubnets
	}
	return nil
}

func (x *ConfigSpec) GetListenExcludeSubnets() []string {
	if x != nil {
		return x.ListenExcludeSubnets
	}
	return nil
}

// MemberSpec holds information about an etcd member.
type MemberSpec struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	MemberId      string                 `protobuf:"bytes,1,opt,name=member_id,json=memberId,proto3" json:"member_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MemberSpec) Reset() {
	*x = MemberSpec{}
	mi := &file_resource_definitions_etcd_etcd_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MemberSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemberSpec) ProtoMessage() {}

func (x *MemberSpec) ProtoReflect() protoreflect.Message {
	mi := &file_resource_definitions_etcd_etcd_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemberSpec.ProtoReflect.Descriptor instead.
func (*MemberSpec) Descriptor() ([]byte, []int) {
	return file_resource_definitions_etcd_etcd_proto_rawDescGZIP(), []int{1}
}

func (x *MemberSpec) GetMemberId() string {
	if x != nil {
		return x.MemberId
	}
	return ""
}

// PKIStatusSpec describes status of rendered secrets.
type PKIStatusSpec struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Ready         bool                   `protobuf:"varint,1,opt,name=ready,proto3" json:"ready,omitempty"`
	Version       string                 `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PKIStatusSpec) Reset() {
	*x = PKIStatusSpec{}
	mi := &file_resource_definitions_etcd_etcd_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PKIStatusSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PKIStatusSpec) ProtoMessage() {}

func (x *PKIStatusSpec) ProtoReflect() protoreflect.Message {
	mi := &file_resource_definitions_etcd_etcd_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PKIStatusSpec.ProtoReflect.Descriptor instead.
func (*PKIStatusSpec) Descriptor() ([]byte, []int) {
	return file_resource_definitions_etcd_etcd_proto_rawDescGZIP(), []int{2}
}

func (x *PKIStatusSpec) GetReady() bool {
	if x != nil {
		return x.Ready
	}
	return false
}

func (x *PKIStatusSpec) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

// SpecSpec describes (some) Specuration settings of etcd.
type SpecSpec struct {
	state                 protoimpl.MessageState `protogen:"open.v1"`
	Name                  string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	AdvertisedAddresses   []*common.NetIP        `protobuf:"bytes,2,rep,name=advertised_addresses,json=advertisedAddresses,proto3" json:"advertised_addresses,omitempty"`
	Image                 string                 `protobuf:"bytes,3,opt,name=image,proto3" json:"image,omitempty"`
	ExtraArgs             map[string]string      `protobuf:"bytes,4,rep,name=extra_args,json=extraArgs,proto3" json:"extra_args,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	ListenPeerAddresses   []*common.NetIP        `protobuf:"bytes,5,rep,name=listen_peer_addresses,json=listenPeerAddresses,proto3" json:"listen_peer_addresses,omitempty"`
	ListenClientAddresses []*common.NetIP        `protobuf:"bytes,6,rep,name=listen_client_addresses,json=listenClientAddresses,proto3" json:"listen_client_addresses,omitempty"`
	unknownFields         protoimpl.UnknownFields
	sizeCache             protoimpl.SizeCache
}

func (x *SpecSpec) Reset() {
	*x = SpecSpec{}
	mi := &file_resource_definitions_etcd_etcd_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SpecSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SpecSpec) ProtoMessage() {}

func (x *SpecSpec) ProtoReflect() protoreflect.Message {
	mi := &file_resource_definitions_etcd_etcd_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SpecSpec.ProtoReflect.Descriptor instead.
func (*SpecSpec) Descriptor() ([]byte, []int) {
	return file_resource_definitions_etcd_etcd_proto_rawDescGZIP(), []int{3}
}

func (x *SpecSpec) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SpecSpec) GetAdvertisedAddresses() []*common.NetIP {
	if x != nil {
		return x.AdvertisedAddresses
	}
	return nil
}

func (x *SpecSpec) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *SpecSpec) GetExtraArgs() map[string]string {
	if x != nil {
		return x.ExtraArgs
	}
	return nil
}

func (x *SpecSpec) GetListenPeerAddresses() []*common.NetIP {
	if x != nil {
		return x.ListenPeerAddresses
	}
	return nil
}

func (x *SpecSpec) GetListenClientAddresses() []*common.NetIP {
	if x != nil {
		return x.ListenClientAddresses
	}
	return nil
}

var File_resource_definitions_etcd_etcd_proto protoreflect.FileDescriptor

var file_resource_definitions_etcd_etcd_proto_rawDesc = string([]byte{
	0x0a, 0x24, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x64, 0x65, 0x66, 0x69, 0x6e,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x65, 0x74, 0x63, 0x64, 0x2f, 0x65, 0x74, 0x63, 0x64,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x74, 0x61, 0x6c, 0x6f, 0x73, 0x2e, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x65, 0x74, 0x63, 0x64, 0x1a, 0x13, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x97, 0x03, 0x0a,
	0x0a, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x70, 0x65, 0x63, 0x12, 0x36, 0x0a, 0x17, 0x61,
	0x64, 0x76, 0x65, 0x72, 0x74, 0x69, 0x73, 0x65, 0x5f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x5f, 0x73,
	0x75, 0x62, 0x6e, 0x65, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x15, 0x61, 0x64,
	0x76, 0x65, 0x72, 0x74, 0x69, 0x73, 0x65, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x53, 0x75, 0x62, 0x6e,
	0x65, 0x74, 0x73, 0x12, 0x3a, 0x0a, 0x19, 0x61, 0x64, 0x76, 0x65, 0x72, 0x74, 0x69, 0x73, 0x65,
	0x5f, 0x65, 0x78, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x5f, 0x73, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x17, 0x61, 0x64, 0x76, 0x65, 0x72, 0x74, 0x69, 0x73,
	0x65, 0x45, 0x78, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x53, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x73, 0x12,
	0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x59, 0x0a, 0x0a, 0x65, 0x78, 0x74, 0x72, 0x61, 0x5f, 0x61,
	0x72, 0x67, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3a, 0x2e, 0x74, 0x61, 0x6c, 0x6f,
	0x73, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x64, 0x65, 0x66, 0x69, 0x6e,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x65, 0x74, 0x63, 0x64, 0x2e, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x53, 0x70, 0x65, 0x63, 0x2e, 0x45, 0x78, 0x74, 0x72, 0x61, 0x41, 0x72, 0x67, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x09, 0x65, 0x78, 0x74, 0x72, 0x61, 0x41, 0x72, 0x67, 0x73,
	0x12, 0x30, 0x0a, 0x14, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x5f, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x5f, 0x73, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x12,
	0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x53, 0x75, 0x62, 0x6e, 0x65,
	0x74, 0x73, 0x12, 0x34, 0x0a, 0x16, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x5f, 0x65, 0x78, 0x63,
	0x6c, 0x75, 0x64, 0x65, 0x5f, 0x73, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x73, 0x18, 0x06, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x14, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x45, 0x78, 0x63, 0x6c, 0x75, 0x64,
	0x65, 0x53, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x73, 0x1a, 0x3c, 0x0a, 0x0e, 0x45, 0x78, 0x74, 0x72,
	0x61, 0x41, 0x72, 0x67, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x29, 0x0a, 0x0a, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x53, 0x70, 0x65, 0x63, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49,
	0x64, 0x22, 0x3f, 0x0a, 0x0d, 0x50, 0x4b, 0x49, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x53, 0x70,
	0x65, 0x63, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x65, 0x61, 0x64, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x05, 0x72, 0x65, 0x61, 0x64, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x22, 0x97, 0x03, 0x0a, 0x08, 0x53, 0x70, 0x65, 0x63, 0x53, 0x70, 0x65, 0x63, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x40, 0x0a, 0x14, 0x61, 0x64, 0x76, 0x65, 0x72, 0x74, 0x69, 0x73, 0x65,
	0x64, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4e, 0x65, 0x74, 0x49, 0x50,
	0x52, 0x13, 0x61, 0x64, 0x76, 0x65, 0x72, 0x74, 0x69, 0x73, 0x65, 0x64, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x57, 0x0a, 0x0a, 0x65,
	0x78, 0x74, 0x72, 0x61, 0x5f, 0x61, 0x72, 0x67, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x38, 0x2e, 0x74, 0x61, 0x6c, 0x6f, 0x73, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x2e, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x65, 0x74, 0x63,
	0x64, 0x2e, 0x53, 0x70, 0x65, 0x63, 0x53, 0x70, 0x65, 0x63, 0x2e, 0x45, 0x78, 0x74, 0x72, 0x61,
	0x41, 0x72, 0x67, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x09, 0x65, 0x78, 0x74, 0x72, 0x61,
	0x41, 0x72, 0x67, 0x73, 0x12, 0x41, 0x0a, 0x15, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x5f, 0x70,
	0x65, 0x65, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x18, 0x05, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4e, 0x65, 0x74,
	0x49, 0x50, 0x52, 0x13, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x50, 0x65, 0x65, 0x72, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x12, 0x45, 0x0a, 0x17, 0x6c, 0x69, 0x73, 0x74, 0x65,
	0x6e, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x65, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x4e, 0x65, 0x74, 0x49, 0x50, 0x52, 0x15, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x43,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x1a, 0x3c,
	0x0a, 0x0e, 0x45, 0x78, 0x74, 0x72, 0x61, 0x41, 0x72, 0x67, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x72, 0x0a, 0x27,
	0x64, 0x65, 0x76, 0x2e, 0x74, 0x61, 0x6c, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x65, 0x74, 0x63, 0x64, 0x5a, 0x47, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x69, 0x64, 0x65, 0x72, 0x6f, 0x6c, 0x61, 0x62, 0x73, 0x2f, 0x74,
	0x61, 0x6c, 0x6f, 0x73, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65,
	0x72, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f,
	0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x65, 0x74, 0x63, 0x64,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_resource_definitions_etcd_etcd_proto_rawDescOnce sync.Once
	file_resource_definitions_etcd_etcd_proto_rawDescData []byte
)

func file_resource_definitions_etcd_etcd_proto_rawDescGZIP() []byte {
	file_resource_definitions_etcd_etcd_proto_rawDescOnce.Do(func() {
		file_resource_definitions_etcd_etcd_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_resource_definitions_etcd_etcd_proto_rawDesc), len(file_resource_definitions_etcd_etcd_proto_rawDesc)))
	})
	return file_resource_definitions_etcd_etcd_proto_rawDescData
}

var file_resource_definitions_etcd_etcd_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_resource_definitions_etcd_etcd_proto_goTypes = []any{
	(*ConfigSpec)(nil),    // 0: talos.resource.definitions.etcd.ConfigSpec
	(*MemberSpec)(nil),    // 1: talos.resource.definitions.etcd.MemberSpec
	(*PKIStatusSpec)(nil), // 2: talos.resource.definitions.etcd.PKIStatusSpec
	(*SpecSpec)(nil),      // 3: talos.resource.definitions.etcd.SpecSpec
	nil,                   // 4: talos.resource.definitions.etcd.ConfigSpec.ExtraArgsEntry
	nil,                   // 5: talos.resource.definitions.etcd.SpecSpec.ExtraArgsEntry
	(*common.NetIP)(nil),  // 6: common.NetIP
}
var file_resource_definitions_etcd_etcd_proto_depIdxs = []int32{
	4, // 0: talos.resource.definitions.etcd.ConfigSpec.extra_args:type_name -> talos.resource.definitions.etcd.ConfigSpec.ExtraArgsEntry
	6, // 1: talos.resource.definitions.etcd.SpecSpec.advertised_addresses:type_name -> common.NetIP
	5, // 2: talos.resource.definitions.etcd.SpecSpec.extra_args:type_name -> talos.resource.definitions.etcd.SpecSpec.ExtraArgsEntry
	6, // 3: talos.resource.definitions.etcd.SpecSpec.listen_peer_addresses:type_name -> common.NetIP
	6, // 4: talos.resource.definitions.etcd.SpecSpec.listen_client_addresses:type_name -> common.NetIP
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_resource_definitions_etcd_etcd_proto_init() }
func file_resource_definitions_etcd_etcd_proto_init() {
	if File_resource_definitions_etcd_etcd_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_resource_definitions_etcd_etcd_proto_rawDesc), len(file_resource_definitions_etcd_etcd_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_resource_definitions_etcd_etcd_proto_goTypes,
		DependencyIndexes: file_resource_definitions_etcd_etcd_proto_depIdxs,
		MessageInfos:      file_resource_definitions_etcd_etcd_proto_msgTypes,
	}.Build()
	File_resource_definitions_etcd_etcd_proto = out.File
	file_resource_definitions_etcd_etcd_proto_goTypes = nil
	file_resource_definitions_etcd_etcd_proto_depIdxs = nil
}
