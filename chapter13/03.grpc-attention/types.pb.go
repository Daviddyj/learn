// Code generated by protoc-gen-go. DO NOT EDIT.
// source: types.proto

package main

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type PersonInformation struct {
	// @gotags: gorm:"primaryKey;column:id"
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// @gotags: gorm:"column:name"
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// @gotags: gorm:"column:sex"
	Sex string `protobuf:"bytes,3,opt,name=sex,proto3" json:"sex,omitempty"`
	// @gotags: gorm:"column:tall"
	Tall float32 `protobuf:"fixed32,4,opt,name=tall,proto3" json:"tall,omitempty"`
	// @gotags: gorm:"column:weight"
	Weight float32 `protobuf:"fixed32,5,opt,name=weight,proto3" json:"weight,omitempty"`
	// @gotags: gorm:"column:age"
	Age                  int64    `protobuf:"varint,6,opt,name=age,proto3" json:"age,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PersonInformation) Reset()         { *m = PersonInformation{} }
func (m *PersonInformation) String() string { return proto.CompactTextString(m) }
func (*PersonInformation) ProtoMessage()    {}
func (*PersonInformation) Descriptor() ([]byte, []int) {
	return fileDescriptor_d938547f84707355, []int{0}
}

func (m *PersonInformation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PersonInformation.Unmarshal(m, b)
}
func (m *PersonInformation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PersonInformation.Marshal(b, m, deterministic)
}
func (m *PersonInformation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PersonInformation.Merge(m, src)
}
func (m *PersonInformation) XXX_Size() int {
	return xxx_messageInfo_PersonInformation.Size(m)
}
func (m *PersonInformation) XXX_DiscardUnknown() {
	xxx_messageInfo_PersonInformation.DiscardUnknown(m)
}

var xxx_messageInfo_PersonInformation proto.InternalMessageInfo

func (m *PersonInformation) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *PersonInformation) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PersonInformation) GetSex() string {
	if m != nil {
		return m.Sex
	}
	return ""
}

func (m *PersonInformation) GetTall() float32 {
	if m != nil {
		return m.Tall
	}
	return 0
}

func (m *PersonInformation) GetWeight() float32 {
	if m != nil {
		return m.Weight
	}
	return 0
}

func (m *PersonInformation) GetAge() int64 {
	if m != nil {
		return m.Age
	}
	return 0
}

func init() {
	proto.RegisterType((*PersonInformation)(nil), "main.PersonInformation")
}

func init() { proto.RegisterFile("types.proto", fileDescriptor_d938547f84707355) }

var fileDescriptor_d938547f84707355 = []byte{
	// 157 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2e, 0xa9, 0x2c, 0x48,
	0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xc9, 0x4d, 0xcc, 0xcc, 0x53, 0x6a, 0x67,
	0xe4, 0x12, 0x0c, 0x48, 0x2d, 0x2a, 0xce, 0xcf, 0xf3, 0xcc, 0x4b, 0xcb, 0x2f, 0xca, 0x4d, 0x2c,
	0xc9, 0xcc, 0xcf, 0x13, 0xe2, 0xe3, 0x62, 0xca, 0x4c, 0x91, 0x60, 0x54, 0x60, 0xd4, 0x60, 0x0e,
	0x62, 0xca, 0x4c, 0x11, 0x12, 0xe2, 0x62, 0xc9, 0x4b, 0xcc, 0x4d, 0x95, 0x60, 0x52, 0x60, 0xd4,
	0xe0, 0x0c, 0x02, 0xb3, 0x85, 0x04, 0xb8, 0x98, 0x8b, 0x53, 0x2b, 0x24, 0x98, 0xc1, 0x42, 0x20,
	0x26, 0x48, 0x55, 0x49, 0x62, 0x4e, 0x8e, 0x04, 0x8b, 0x02, 0xa3, 0x06, 0x53, 0x10, 0x98, 0x2d,
	0x24, 0xc6, 0xc5, 0x56, 0x9e, 0x9a, 0x99, 0x9e, 0x51, 0x22, 0xc1, 0x0a, 0x16, 0x85, 0xf2, 0x40,
	0xba, 0x13, 0xd3, 0x53, 0x25, 0xd8, 0xc0, 0x56, 0x80, 0x98, 0x49, 0x6c, 0x60, 0x67, 0x19, 0x03,
	0x02, 0x00, 0x00, 0xff, 0xff, 0x87, 0x1f, 0xcd, 0x6b, 0xa5, 0x00, 0x00, 0x00,
}
