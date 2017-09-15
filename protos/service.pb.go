// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

/*
Package trackem is a generated protocol buffer package.

It is generated from these files:
	service.proto

It has these top-level messages:
	EmailTracker
*/
package trackem

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type EmailTracker struct {
	Id       string            `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	From     string            `protobuf:"bytes,2,opt,name=from" json:"from,omitempty"`
	To       string            `protobuf:"bytes,3,opt,name=to" json:"to,omitempty"`
	Subject  string            `protobuf:"bytes,4,opt,name=subject" json:"subject,omitempty"`
	Metadata map[string]string `protobuf:"bytes,5,rep,name=metadata" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *EmailTracker) Reset()                    { *m = EmailTracker{} }
func (m *EmailTracker) String() string            { return proto.CompactTextString(m) }
func (*EmailTracker) ProtoMessage()               {}
func (*EmailTracker) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *EmailTracker) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *EmailTracker) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *EmailTracker) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

func (m *EmailTracker) GetSubject() string {
	if m != nil {
		return m.Subject
	}
	return ""
}

func (m *EmailTracker) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func init() {
	proto.RegisterType((*EmailTracker)(nil), "trackem.EmailTracker")
}

func init() { proto.RegisterFile("service.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 196 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x4e, 0x2d, 0x2a,
	0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2f, 0x29, 0x4a, 0x4c, 0xce,
	0x4e, 0xcd, 0x55, 0xba, 0xc5, 0xc8, 0xc5, 0xe3, 0x9a, 0x9b, 0x98, 0x99, 0x13, 0x02, 0x16, 0x28,
	0x12, 0xe2, 0xe3, 0x62, 0xca, 0x4c, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x62, 0xca, 0x4c,
	0x11, 0x12, 0xe2, 0x62, 0x49, 0x2b, 0xca, 0xcf, 0x95, 0x60, 0x02, 0x8b, 0x80, 0xd9, 0x20, 0x35,
	0x25, 0xf9, 0x12, 0xcc, 0x10, 0x35, 0x25, 0xf9, 0x42, 0x12, 0x5c, 0xec, 0xc5, 0xa5, 0x49, 0x59,
	0xa9, 0xc9, 0x25, 0x12, 0x2c, 0x60, 0x41, 0x18, 0x57, 0xc8, 0x9e, 0x8b, 0x23, 0x37, 0xb5, 0x24,
	0x31, 0x25, 0xb1, 0x24, 0x51, 0x82, 0x55, 0x81, 0x59, 0x83, 0xdb, 0x48, 0x59, 0x0f, 0x6a, 0xb5,
	0x1e, 0xb2, 0xb5, 0x7a, 0xbe, 0x50, 0x55, 0xae, 0x79, 0x25, 0x45, 0x95, 0x41, 0x70, 0x4d, 0x52,
	0xd6, 0x5c, 0xbc, 0x28, 0x52, 0x42, 0x02, 0x5c, 0xcc, 0xd9, 0xa9, 0x95, 0x50, 0x07, 0x82, 0x98,
	0x42, 0x22, 0x5c, 0xac, 0x65, 0x89, 0x39, 0xa5, 0xa9, 0x50, 0x27, 0x42, 0x38, 0x56, 0x4c, 0x16,
	0x8c, 0x49, 0x6c, 0x60, 0xcf, 0x1a, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x9e, 0x3c, 0xc7, 0x13,
	0xfd, 0x00, 0x00, 0x00,
}