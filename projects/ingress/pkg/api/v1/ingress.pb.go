// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/ingress/api/v1/ingress.proto

package v1 // import "github.com/solo-io/gloo/projects/ingress/pkg/api/v1"

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import core "github.com/solo-io/solo-kit/pkg/api/v1/resources/core"

import bytes "bytes"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

//
// @solo-kit:resource.short_name=ig
// @solo-kit:resource.plural_name=ingresses
// @solo-kit:resource.resource_groups=api.ingress.solo.io
//
// A simple wrapper for a Kubernetes Ingress Object.
type Ingress struct {
	// a raw byte representation of the kubernetes ingress this resource wraps
	KubeIngressRaw []byte `protobuf:"bytes,1,opt,name=kube_ingress_raw,json=kubeIngressRaw,proto3" json:"kube_ingress_raw,omitempty"`
	// Status indicates the validation status of this resource.
	// Status is read-only by clients, and set by gloo during validation
	Status core.Status `protobuf:"bytes,6,opt,name=status" json:"status" testdiff:"ignore"`
	// Metadata contains the object metadata for this resource
	Metadata             core.Metadata `protobuf:"bytes,7,opt,name=metadata" json:"metadata"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Ingress) Reset()         { *m = Ingress{} }
func (m *Ingress) String() string { return proto.CompactTextString(m) }
func (*Ingress) ProtoMessage()    {}
func (*Ingress) Descriptor() ([]byte, []int) {
	return fileDescriptor_ingress_c4249ff5f02ab6e5, []int{0}
}
func (m *Ingress) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ingress.Unmarshal(m, b)
}
func (m *Ingress) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ingress.Marshal(b, m, deterministic)
}
func (dst *Ingress) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ingress.Merge(dst, src)
}
func (m *Ingress) XXX_Size() int {
	return xxx_messageInfo_Ingress.Size(m)
}
func (m *Ingress) XXX_DiscardUnknown() {
	xxx_messageInfo_Ingress.DiscardUnknown(m)
}

var xxx_messageInfo_Ingress proto.InternalMessageInfo

func (m *Ingress) GetKubeIngressRaw() []byte {
	if m != nil {
		return m.KubeIngressRaw
	}
	return nil
}

func (m *Ingress) GetStatus() core.Status {
	if m != nil {
		return m.Status
	}
	return core.Status{}
}

func (m *Ingress) GetMetadata() core.Metadata {
	if m != nil {
		return m.Metadata
	}
	return core.Metadata{}
}

func init() {
	proto.RegisterType((*Ingress)(nil), "ingress.solo.io.Ingress")
}
func (this *Ingress) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Ingress)
	if !ok {
		that2, ok := that.(Ingress)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !bytes.Equal(this.KubeIngressRaw, that1.KubeIngressRaw) {
		return false
	}
	if !this.Status.Equal(&that1.Status) {
		return false
	}
	if !this.Metadata.Equal(&that1.Metadata) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}

func init() {
	proto.RegisterFile("github.com/solo-io/gloo/projects/ingress/api/v1/ingress.proto", fileDescriptor_ingress_c4249ff5f02ab6e5)
}

var fileDescriptor_ingress_c4249ff5f02ab6e5 = []byte{
	// 273 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xb2, 0x4d, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x2f, 0xce, 0xcf, 0xc9, 0xd7, 0xcd, 0xcc, 0xd7, 0x4f,
	0xcf, 0xc9, 0xcf, 0xd7, 0x2f, 0x28, 0xca, 0xcf, 0x4a, 0x4d, 0x2e, 0x29, 0xd6, 0xcf, 0xcc, 0x4b,
	0x2f, 0x4a, 0x2d, 0x2e, 0xd6, 0x4f, 0x2c, 0xc8, 0xd4, 0x2f, 0x33, 0x84, 0x71, 0xf5, 0x0a, 0x8a,
	0xf2, 0x4b, 0xf2, 0x85, 0xf8, 0x61, 0x5c, 0x90, 0x5e, 0xbd, 0xcc, 0x7c, 0x29, 0x91, 0xf4, 0xfc,
	0xf4, 0x7c, 0xb0, 0x9c, 0x3e, 0x88, 0x05, 0x51, 0x26, 0x65, 0x88, 0xc5, 0x16, 0x30, 0x9d, 0x9d,
	0x59, 0x02, 0x33, 0x38, 0x37, 0xb5, 0x24, 0x31, 0x25, 0xb1, 0x24, 0x11, 0xaa, 0x45, 0x9f, 0x08,
	0x2d, 0xc5, 0x25, 0x89, 0x25, 0xa5, 0x50, 0xa7, 0x28, 0x6d, 0x63, 0xe4, 0x62, 0xf7, 0x84, 0xb8,
	0x46, 0x48, 0x83, 0x4b, 0x20, 0xbb, 0x34, 0x29, 0x35, 0x1e, 0xea, 0xba, 0xf8, 0xa2, 0xc4, 0x72,
	0x09, 0x46, 0x05, 0x46, 0x0d, 0x9e, 0x20, 0x3e, 0x90, 0x38, 0x54, 0x59, 0x50, 0x62, 0xb9, 0x90,
	0x3b, 0x17, 0x1b, 0xc4, 0x14, 0x09, 0x36, 0x05, 0x46, 0x0d, 0x6e, 0x23, 0x11, 0xbd, 0xe4, 0xfc,
	0xa2, 0x54, 0x98, 0x77, 0xf4, 0x82, 0xc1, 0x72, 0x4e, 0x92, 0x27, 0xee, 0xc9, 0x33, 0x7c, 0xba,
	0x27, 0x2f, 0x58, 0x92, 0x5a, 0x5c, 0x92, 0x92, 0x99, 0x96, 0x66, 0xa5, 0x94, 0x99, 0x9e, 0x97,
	0x5f, 0x94, 0xaa, 0x14, 0x04, 0xd5, 0x2e, 0x64, 0xc1, 0xc5, 0x01, 0xf3, 0x81, 0x04, 0x3b, 0xd8,
	0x28, 0x31, 0x54, 0xa3, 0x7c, 0xa1, 0xb2, 0x4e, 0x2c, 0x20, 0xc3, 0x82, 0xe0, 0xaa, 0x9d, 0x2c,
	0x57, 0x3c, 0x92, 0x63, 0x8c, 0x32, 0x26, 0x3a, 0x22, 0x0a, 0xb2, 0xd3, 0xa1, 0x01, 0x90, 0xc4,
	0x06, 0xf6, 0xba, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xcf, 0x51, 0x75, 0x8b, 0xc6, 0x01, 0x00,
	0x00,
}
