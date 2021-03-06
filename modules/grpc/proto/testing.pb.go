// Code generated by protoc-gen-go. DO NOT EDIT.
// source: testing.proto

package testing

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

type GetRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a0b3d6cfd712914, []int{0}
}

func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequest.Unmarshal(m, b)
}
func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
}
func (m *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(m, src)
}
func (m *GetRequest) XXX_Size() int {
	return xxx_messageInfo_GetRequest.Size(m)
}
func (m *GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequest proto.InternalMessageInfo

type AddRequest struct {
	Email                *string  `protobuf:"bytes,1,req,name=email" json:"email,omitempty"`
	Password             *string  `protobuf:"bytes,2,req,name=password" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddRequest) Reset()         { *m = AddRequest{} }
func (m *AddRequest) String() string { return proto.CompactTextString(m) }
func (*AddRequest) ProtoMessage()    {}
func (*AddRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a0b3d6cfd712914, []int{1}
}

func (m *AddRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddRequest.Unmarshal(m, b)
}
func (m *AddRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddRequest.Marshal(b, m, deterministic)
}
func (m *AddRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddRequest.Merge(m, src)
}
func (m *AddRequest) XXX_Size() int {
	return xxx_messageInfo_AddRequest.Size(m)
}
func (m *AddRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddRequest proto.InternalMessageInfo

func (m *AddRequest) GetEmail() string {
	if m != nil && m.Email != nil {
		return *m.Email
	}
	return ""
}

func (m *AddRequest) GetPassword() string {
	if m != nil && m.Password != nil {
		return *m.Password
	}
	return ""
}

type DeleteRequest struct {
	Id                   *string  `protobuf:"bytes,1,req,name=id" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteRequest) Reset()         { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()    {}
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a0b3d6cfd712914, []int{2}
}

func (m *DeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteRequest.Unmarshal(m, b)
}
func (m *DeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteRequest.Marshal(b, m, deterministic)
}
func (m *DeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteRequest.Merge(m, src)
}
func (m *DeleteRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteRequest.Size(m)
}
func (m *DeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteRequest proto.InternalMessageInfo

func (m *DeleteRequest) GetId() string {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return ""
}

type GetSingleReply struct {
	Email                *string  `protobuf:"bytes,1,req,name=email" json:"email,omitempty"`
	Password             *string  `protobuf:"bytes,2,req,name=password" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetSingleReply) Reset()         { *m = GetSingleReply{} }
func (m *GetSingleReply) String() string { return proto.CompactTextString(m) }
func (*GetSingleReply) ProtoMessage()    {}
func (*GetSingleReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a0b3d6cfd712914, []int{3}
}

func (m *GetSingleReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetSingleReply.Unmarshal(m, b)
}
func (m *GetSingleReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetSingleReply.Marshal(b, m, deterministic)
}
func (m *GetSingleReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetSingleReply.Merge(m, src)
}
func (m *GetSingleReply) XXX_Size() int {
	return xxx_messageInfo_GetSingleReply.Size(m)
}
func (m *GetSingleReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GetSingleReply.DiscardUnknown(m)
}

var xxx_messageInfo_GetSingleReply proto.InternalMessageInfo

func (m *GetSingleReply) GetEmail() string {
	if m != nil && m.Email != nil {
		return *m.Email
	}
	return ""
}

func (m *GetSingleReply) GetPassword() string {
	if m != nil && m.Password != nil {
		return *m.Password
	}
	return ""
}

type GetReply struct {
	User                 []*GetSingleReply `protobuf:"bytes,1,rep,name=User" json:"User,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *GetReply) Reset()         { *m = GetReply{} }
func (m *GetReply) String() string { return proto.CompactTextString(m) }
func (*GetReply) ProtoMessage()    {}
func (*GetReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a0b3d6cfd712914, []int{4}
}

func (m *GetReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetReply.Unmarshal(m, b)
}
func (m *GetReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetReply.Marshal(b, m, deterministic)
}
func (m *GetReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetReply.Merge(m, src)
}
func (m *GetReply) XXX_Size() int {
	return xxx_messageInfo_GetReply.Size(m)
}
func (m *GetReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GetReply.DiscardUnknown(m)
}

var xxx_messageInfo_GetReply proto.InternalMessageInfo

func (m *GetReply) GetUser() []*GetSingleReply {
	if m != nil {
		return m.User
	}
	return nil
}

type AddReply struct {
	Id                   *string  `protobuf:"bytes,1,req,name=id" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddReply) Reset()         { *m = AddReply{} }
func (m *AddReply) String() string { return proto.CompactTextString(m) }
func (*AddReply) ProtoMessage()    {}
func (*AddReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a0b3d6cfd712914, []int{5}
}

func (m *AddReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddReply.Unmarshal(m, b)
}
func (m *AddReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddReply.Marshal(b, m, deterministic)
}
func (m *AddReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddReply.Merge(m, src)
}
func (m *AddReply) XXX_Size() int {
	return xxx_messageInfo_AddReply.Size(m)
}
func (m *AddReply) XXX_DiscardUnknown() {
	xxx_messageInfo_AddReply.DiscardUnknown(m)
}

var xxx_messageInfo_AddReply proto.InternalMessageInfo

func (m *AddReply) GetId() string {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return ""
}

type DeleteReply struct {
	Id                   *string  `protobuf:"bytes,1,req,name=id" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteReply) Reset()         { *m = DeleteReply{} }
func (m *DeleteReply) String() string { return proto.CompactTextString(m) }
func (*DeleteReply) ProtoMessage()    {}
func (*DeleteReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a0b3d6cfd712914, []int{6}
}

func (m *DeleteReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteReply.Unmarshal(m, b)
}
func (m *DeleteReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteReply.Marshal(b, m, deterministic)
}
func (m *DeleteReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteReply.Merge(m, src)
}
func (m *DeleteReply) XXX_Size() int {
	return xxx_messageInfo_DeleteReply.Size(m)
}
func (m *DeleteReply) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteReply.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteReply proto.InternalMessageInfo

func (m *DeleteReply) GetId() string {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return ""
}

func init() {
	proto.RegisterType((*GetRequest)(nil), "GetRequest")
	proto.RegisterType((*AddRequest)(nil), "AddRequest")
	proto.RegisterType((*DeleteRequest)(nil), "DeleteRequest")
	proto.RegisterType((*GetSingleReply)(nil), "GetSingleReply")
	proto.RegisterType((*GetReply)(nil), "GetReply")
	proto.RegisterType((*AddReply)(nil), "AddReply")
	proto.RegisterType((*DeleteReply)(nil), "DeleteReply")
}

func init() {
	proto.RegisterFile("testing.proto", fileDescriptor_7a0b3d6cfd712914)
}

var fileDescriptor_7a0b3d6cfd712914 = []byte{
	// 250 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x90, 0xc1, 0x4a, 0xc3, 0x40,
	0x18, 0x84, 0x93, 0xa8, 0x34, 0x99, 0xb4, 0x11, 0x16, 0x0f, 0x61, 0x41, 0x2c, 0x5b, 0x0f, 0x3d,
	0xc8, 0x2f, 0xf4, 0x01, 0x84, 0x8a, 0xd0, 0x7b, 0xd4, 0x07, 0x28, 0xec, 0x4f, 0x59, 0x88, 0x4d,
	0xcc, 0xae, 0x88, 0x57, 0x9f, 0x5c, 0xb2, 0xe9, 0x5a, 0x23, 0x9e, 0x3c, 0xfe, 0xcc, 0xce, 0xcc,
	0xb7, 0x83, 0x99, 0x63, 0xeb, 0xcc, 0x7e, 0x47, 0x6d, 0xd7, 0xb8, 0x46, 0x4d, 0x81, 0x0d, 0xbb,
	0x8a, 0x5f, 0xdf, 0xd8, 0x3a, 0x75, 0x07, 0xac, 0xb5, 0x3e, 0x5c, 0xe2, 0x02, 0x67, 0xfc, 0xb2,
	0x35, 0x75, 0x19, 0xcf, 0x93, 0x65, 0x56, 0x0d, 0x87, 0x90, 0x48, 0xdb, 0xad, 0xb5, 0xef, 0x4d,
	0xa7, 0xcb, 0xc4, 0x0b, 0xdf, 0xb7, 0xba, 0xc2, 0xec, 0x81, 0x6b, 0x76, 0x1c, 0x22, 0x0a, 0x24,
	0x46, 0x1f, 0xfc, 0x89, 0xd1, 0xea, 0x1e, 0xc5, 0x86, 0xdd, 0xa3, 0xd9, 0xef, 0x6a, 0xae, 0xb8,
	0xad, 0x3f, 0xfe, 0x51, 0x72, 0x8b, 0xd4, 0x23, 0xf7, 0xee, 0x05, 0x4e, 0x9f, 0x2d, 0x77, 0x65,
	0x3c, 0x3f, 0x59, 0xe6, 0xab, 0x73, 0x1a, 0x87, 0x57, 0x5e, 0x54, 0x12, 0xa9, 0xff, 0x55, 0x6f,
	0xf8, 0x0d, 0x74, 0x89, 0x3c, 0x10, 0xff, 0x21, 0xaf, 0x3e, 0x63, 0x4c, 0x9e, 0x86, 0xc1, 0xc4,
	0xb5, 0xef, 0xed, 0x13, 0xad, 0xc8, 0xe9, 0xb8, 0x9a, 0xcc, 0x28, 0xf0, 0xa8, 0x48, 0x2c, 0x30,
	0x59, 0x6b, 0xdd, 0xbf, 0x12, 0x39, 0x1d, 0xc7, 0x94, 0x19, 0x05, 0x06, 0x15, 0x89, 0x1b, 0x60,
	0x68, 0xf5, 0xef, 0x0a, 0x1a, 0x8d, 0x26, 0xa7, 0xf4, 0x03, 0x49, 0x45, 0x5f, 0x01, 0x00, 0x00,
	0xff, 0xff, 0x83, 0x1f, 0xeb, 0x6a, 0xb3, 0x01, 0x00, 0x00,
}
