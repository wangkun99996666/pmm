// Code generated by protoc-gen-go. DO NOT EDIT.
// source: inventory/host.proto

package inventory

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/mwitkow/go-proto-validators"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// HostNodeInfo describes the way Service or Agent runs on Node.
type HostNodeInfo struct {
	// Node identifier where Service or Agent runs.
	NodeId string `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	// Docker container ID.
	ContainerId string `protobuf:"bytes,2,opt,name=container_id,json=containerId,proto3" json:"container_id,omitempty"`
	// Docker container name.
	ContainerName string `protobuf:"bytes,3,opt,name=container_name,json=containerName,proto3" json:"container_name,omitempty"`
	// Kubernetes pod UID.
	KubernetesPodUid string `protobuf:"bytes,4,opt,name=kubernetes_pod_uid,json=kubernetesPodUid,proto3" json:"kubernetes_pod_uid,omitempty"`
	// Kubernetes pod name.
	KubernetesPodName    string   `protobuf:"bytes,5,opt,name=kubernetes_pod_name,json=kubernetesPodName,proto3" json:"kubernetes_pod_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HostNodeInfo) Reset()         { *m = HostNodeInfo{} }
func (m *HostNodeInfo) String() string { return proto.CompactTextString(m) }
func (*HostNodeInfo) ProtoMessage()    {}
func (*HostNodeInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_host_913ff863c11c834a, []int{0}
}
func (m *HostNodeInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HostNodeInfo.Unmarshal(m, b)
}
func (m *HostNodeInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HostNodeInfo.Marshal(b, m, deterministic)
}
func (dst *HostNodeInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HostNodeInfo.Merge(dst, src)
}
func (m *HostNodeInfo) XXX_Size() int {
	return xxx_messageInfo_HostNodeInfo.Size(m)
}
func (m *HostNodeInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_HostNodeInfo.DiscardUnknown(m)
}

var xxx_messageInfo_HostNodeInfo proto.InternalMessageInfo

func (m *HostNodeInfo) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

func (m *HostNodeInfo) GetContainerId() string {
	if m != nil {
		return m.ContainerId
	}
	return ""
}

func (m *HostNodeInfo) GetContainerName() string {
	if m != nil {
		return m.ContainerName
	}
	return ""
}

func (m *HostNodeInfo) GetKubernetesPodUid() string {
	if m != nil {
		return m.KubernetesPodUid
	}
	return ""
}

func (m *HostNodeInfo) GetKubernetesPodName() string {
	if m != nil {
		return m.KubernetesPodName
	}
	return ""
}

func init() {
	proto.RegisterType((*HostNodeInfo)(nil), "inventory.HostNodeInfo")
}

func init() { proto.RegisterFile("inventory/host.proto", fileDescriptor_host_913ff863c11c834a) }

var fileDescriptor_host_913ff863c11c834a = []byte{
	// 244 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0xd0, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0xc6, 0x71, 0x52, 0x35, 0xd2, 0x6d, 0x15, 0x5d, 0x3d, 0x04, 0x2f, 0x55, 0x41, 0xf0, 0x60,
	0x93, 0x83, 0xe0, 0x03, 0xf4, 0x64, 0x2e, 0x45, 0x0a, 0x82, 0x78, 0x09, 0x9b, 0xce, 0xd8, 0x2e,
	0x35, 0x33, 0x65, 0x33, 0x69, 0xf0, 0x49, 0x3d, 0x0a, 0x3e, 0x89, 0x64, 0x85, 0x0d, 0x7a, 0x5b,
	0xfe, 0xfb, 0xe3, 0x83, 0x5d, 0x75, 0x6e, 0x69, 0x87, 0x24, 0xec, 0x3e, 0xb2, 0x35, 0xd7, 0x92,
	0x6e, 0x1d, 0x0b, 0xeb, 0x61, 0xa8, 0x17, 0x0f, 0x2b, 0x2b, 0xeb, 0xa6, 0x4c, 0x97, 0x5c, 0x65,
	0x55, 0x6b, 0x65, 0xc3, 0x6d, 0xb6, 0xe2, 0xa9, 0x77, 0xd3, 0x9d, 0x79, 0xb7, 0x60, 0x84, 0x5d,
	0x9d, 0x85, 0xe3, 0xef, 0xc4, 0xf5, 0x67, 0xa4, 0xc6, 0x8f, 0x5c, 0xcb, 0x9c, 0x01, 0x73, 0x7a,
	0x63, 0x3d, 0x51, 0x87, 0xc4, 0x80, 0x85, 0x85, 0x24, 0xba, 0x8c, 0x6e, 0x87, 0xb3, 0xf8, 0xfb,
	0x6b, 0x32, 0x78, 0x89, 0x16, 0x71, 0x97, 0x73, 0xd0, 0x57, 0x6a, 0xbc, 0x64, 0x12, 0x63, 0x09,
	0x5d, 0xa7, 0x06, 0x9d, 0x5a, 0x8c, 0x42, 0xcb, 0x41, 0xdf, 0xa8, 0xe3, 0x9e, 0x90, 0xa9, 0x30,
	0xd9, 0xf3, 0xe8, 0x28, 0xd4, 0xb9, 0xa9, 0x50, 0xdf, 0x29, 0xbd, 0x69, 0x4a, 0x74, 0x84, 0x82,
	0x75, 0xb1, 0x65, 0x28, 0x1a, 0x0b, 0xc9, 0xbe, 0xa7, 0x27, 0xfd, 0xcd, 0x13, 0xc3, 0xb3, 0x05,
	0x9d, 0xaa, 0xb3, 0x7f, 0xda, 0x2f, 0x1f, 0x78, 0x7e, 0xfa, 0x87, 0x77, 0xeb, 0xb3, 0xd1, 0x6b,
	0xff, 0x3d, 0x65, 0xec, 0x5f, 0x7b, 0xff, 0x13, 0x00, 0x00, 0xff, 0xff, 0x41, 0x95, 0x41, 0xa8,
	0x48, 0x01, 0x00, 0x00,
}
