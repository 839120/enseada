// Copyright 2019-2020 Enseada authors
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: auth/v1beta1/user.proto

package authv1beta1

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
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

// An OAuth user.
type User struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_64007389f379c5ba, []int{0}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func init() {
	proto.RegisterType((*User)(nil), "auth.v1beta1.User")
}

func init() { proto.RegisterFile("auth/v1beta1/user.proto", fileDescriptor_64007389f379c5ba) }

var fileDescriptor_64007389f379c5ba = []byte{
	// 143 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4f, 0x2c, 0x2d, 0xc9,
	0xd0, 0x2f, 0x33, 0x4c, 0x4a, 0x2d, 0x49, 0x34, 0xd4, 0x2f, 0x2d, 0x4e, 0x2d, 0xd2, 0x2b, 0x28,
	0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x01, 0x49, 0xe8, 0x41, 0x25, 0x94, 0x94, 0xb8, 0x58, 0x42, 0x8b,
	0x53, 0x8b, 0x84, 0xa4, 0xb8, 0x38, 0x40, 0x6a, 0xf2, 0x12, 0x73, 0x53, 0x25, 0x18, 0x15, 0x18,
	0x35, 0x38, 0x83, 0xe0, 0x7c, 0xa7, 0x70, 0x2e, 0xf1, 0xcc, 0x7c, 0xbd, 0xd4, 0xbc, 0xe2, 0xd4,
	0xc4, 0x94, 0x44, 0x3d, 0x64, 0xed, 0x4e, 0x9c, 0x20, 0xcd, 0x01, 0x20, 0x73, 0x03, 0x18, 0xa3,
	0xb8, 0x41, 0x52, 0x50, 0x99, 0x45, 0x4c, 0xcc, 0x8e, 0x11, 0x11, 0xab, 0x98, 0x78, 0x1c, 0x41,
	0xca, 0xc3, 0x0c, 0x9d, 0x40, 0x82, 0xa7, 0x20, 0xdc, 0x18, 0x28, 0x37, 0x89, 0x0d, 0xec, 0x22,
	0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x16, 0x8f, 0x66, 0xb5, 0xac, 0x00, 0x00, 0x00,
}