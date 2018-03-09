// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/loomnetwork/loom/vm/vm.proto

/*
Package vm is a generated protocol buffer package.

It is generated from these files:
	github.com/loomnetwork/loom/vm/vm.proto

It has these top-level messages:
	DeployTx
	SendTx
*/
package vm

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import loom "github.com/loomnetwork/loom"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type DeployTx struct {
	To   *loom.Address `protobuf:"bytes,1,opt,name=to" json:"to,omitempty"`
	Code []byte        `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
}

func (m *DeployTx) Reset()                    { *m = DeployTx{} }
func (m *DeployTx) String() string            { return proto.CompactTextString(m) }
func (*DeployTx) ProtoMessage()               {}
func (*DeployTx) Descriptor() ([]byte, []int) { return fileDescriptorVm, []int{0} }

func (m *DeployTx) GetTo() *loom.Address {
	if m != nil {
		return m.To
	}
	return nil
}

func (m *DeployTx) GetCode() []byte {
	if m != nil {
		return m.Code
	}
	return nil
}

type SendTx struct {
	To *loom.Address `protobuf:"bytes,1,opt,name=to" json:"to,omitempty"`
}

func (m *SendTx) Reset()                    { *m = SendTx{} }
func (m *SendTx) String() string            { return proto.CompactTextString(m) }
func (*SendTx) ProtoMessage()               {}
func (*SendTx) Descriptor() ([]byte, []int) { return fileDescriptorVm, []int{1} }

func (m *SendTx) GetTo() *loom.Address {
	if m != nil {
		return m.To
	}
	return nil
}

func init() {
	proto.RegisterType((*DeployTx)(nil), "DeployTx")
	proto.RegisterType((*SendTx)(nil), "SendTx")
}

func init() { proto.RegisterFile("github.com/loomnetwork/loom/vm/vm.proto", fileDescriptorVm) }

var fileDescriptorVm = []byte{
	// 137 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4f, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0xcf, 0xc9, 0xcf, 0xcf, 0xcd, 0x4b, 0x2d, 0x29, 0xcf,
	0x2f, 0xca, 0x06, 0xb3, 0xf5, 0xcb, 0x40, 0x48, 0xaf, 0xa0, 0x28, 0xbf, 0x24, 0x5f, 0x4a, 0x0d,
	0x9f, 0x42, 0x10, 0x01, 0x51, 0xa7, 0x64, 0xc1, 0xc5, 0xe1, 0x92, 0x5a, 0x90, 0x93, 0x5f, 0x19,
	0x52, 0x21, 0x24, 0xc1, 0xc5, 0x54, 0x92, 0x2f, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x6d, 0xc4, 0xa1,
	0xe7, 0x98, 0x92, 0x52, 0x94, 0x5a, 0x5c, 0x1c, 0xc4, 0x54, 0x92, 0x2f, 0x24, 0xc4, 0xc5, 0x92,
	0x9c, 0x9f, 0x92, 0x2a, 0xc1, 0xa4, 0xc0, 0xa8, 0xc1, 0x13, 0x04, 0x66, 0x2b, 0x29, 0x71, 0xb1,
	0x05, 0xa7, 0xe6, 0xa5, 0xe0, 0xd3, 0x97, 0xc4, 0x06, 0xb6, 0xc4, 0x18, 0x10, 0x00, 0x00, 0xff,
	0xff, 0xda, 0x50, 0x96, 0x1f, 0xb7, 0x00, 0x00, 0x00,
}
