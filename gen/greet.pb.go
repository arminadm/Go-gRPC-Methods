// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/greet.proto

package gen

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

type NoParam struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NoParam) Reset()         { *m = NoParam{} }
func (m *NoParam) String() string { return proto.CompactTextString(m) }
func (*NoParam) ProtoMessage()    {}
func (*NoParam) Descriptor() ([]byte, []int) {
	return fileDescriptor_95ca2ad3f55a58e3, []int{0}
}

func (m *NoParam) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NoParam.Unmarshal(m, b)
}
func (m *NoParam) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NoParam.Marshal(b, m, deterministic)
}
func (m *NoParam) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NoParam.Merge(m, src)
}
func (m *NoParam) XXX_Size() int {
	return xxx_messageInfo_NoParam.Size(m)
}
func (m *NoParam) XXX_DiscardUnknown() {
	xxx_messageInfo_NoParam.DiscardUnknown(m)
}

var xxx_messageInfo_NoParam proto.InternalMessageInfo

type HelloRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloRequest) Reset()         { *m = HelloRequest{} }
func (m *HelloRequest) String() string { return proto.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()    {}
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_95ca2ad3f55a58e3, []int{1}
}

func (m *HelloRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloRequest.Unmarshal(m, b)
}
func (m *HelloRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloRequest.Marshal(b, m, deterministic)
}
func (m *HelloRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloRequest.Merge(m, src)
}
func (m *HelloRequest) XXX_Size() int {
	return xxx_messageInfo_HelloRequest.Size(m)
}
func (m *HelloRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HelloRequest proto.InternalMessageInfo

func (m *HelloRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type HelloResponse struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloResponse) Reset()         { *m = HelloResponse{} }
func (m *HelloResponse) String() string { return proto.CompactTextString(m) }
func (*HelloResponse) ProtoMessage()    {}
func (*HelloResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_95ca2ad3f55a58e3, []int{2}
}

func (m *HelloResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloResponse.Unmarshal(m, b)
}
func (m *HelloResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloResponse.Marshal(b, m, deterministic)
}
func (m *HelloResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloResponse.Merge(m, src)
}
func (m *HelloResponse) XXX_Size() int {
	return xxx_messageInfo_HelloResponse.Size(m)
}
func (m *HelloResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HelloResponse proto.InternalMessageInfo

func (m *HelloResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type NameList struct {
	Names                []string `protobuf:"bytes,1,rep,name=names,proto3" json:"names,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NameList) Reset()         { *m = NameList{} }
func (m *NameList) String() string { return proto.CompactTextString(m) }
func (*NameList) ProtoMessage()    {}
func (*NameList) Descriptor() ([]byte, []int) {
	return fileDescriptor_95ca2ad3f55a58e3, []int{3}
}

func (m *NameList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NameList.Unmarshal(m, b)
}
func (m *NameList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NameList.Marshal(b, m, deterministic)
}
func (m *NameList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NameList.Merge(m, src)
}
func (m *NameList) XXX_Size() int {
	return xxx_messageInfo_NameList.Size(m)
}
func (m *NameList) XXX_DiscardUnknown() {
	xxx_messageInfo_NameList.DiscardUnknown(m)
}

var xxx_messageInfo_NameList proto.InternalMessageInfo

func (m *NameList) GetNames() []string {
	if m != nil {
		return m.Names
	}
	return nil
}

type MessageList struct {
	Messages             []string `protobuf:"bytes,1,rep,name=messages,proto3" json:"messages,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MessageList) Reset()         { *m = MessageList{} }
func (m *MessageList) String() string { return proto.CompactTextString(m) }
func (*MessageList) ProtoMessage()    {}
func (*MessageList) Descriptor() ([]byte, []int) {
	return fileDescriptor_95ca2ad3f55a58e3, []int{4}
}

func (m *MessageList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessageList.Unmarshal(m, b)
}
func (m *MessageList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessageList.Marshal(b, m, deterministic)
}
func (m *MessageList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageList.Merge(m, src)
}
func (m *MessageList) XXX_Size() int {
	return xxx_messageInfo_MessageList.Size(m)
}
func (m *MessageList) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageList.DiscardUnknown(m)
}

var xxx_messageInfo_MessageList proto.InternalMessageInfo

func (m *MessageList) GetMessages() []string {
	if m != nil {
		return m.Messages
	}
	return nil
}

func init() {
	proto.RegisterType((*NoParam)(nil), "greet_proto.NoParam")
	proto.RegisterType((*HelloRequest)(nil), "greet_proto.HelloRequest")
	proto.RegisterType((*HelloResponse)(nil), "greet_proto.HelloResponse")
	proto.RegisterType((*NameList)(nil), "greet_proto.NameList")
	proto.RegisterType((*MessageList)(nil), "greet_proto.MessageList")
}

func init() {
	proto.RegisterFile("proto/greet.proto", fileDescriptor_95ca2ad3f55a58e3)
}

var fileDescriptor_95ca2ad3f55a58e3 = []byte{
	// 280 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x90, 0x31, 0x4f, 0xc3, 0x30,
	0x10, 0x85, 0x65, 0xa0, 0xb4, 0xb9, 0x96, 0x81, 0x53, 0x87, 0x34, 0x53, 0xe5, 0x29, 0x5d, 0xd2,
	0x0a, 0x66, 0x16, 0x3a, 0x80, 0x04, 0xad, 0xaa, 0x54, 0x62, 0x60, 0x41, 0x6e, 0x39, 0x45, 0x96,
	0x62, 0x3b, 0xd8, 0x86, 0xff, 0xc7, 0x3f, 0x43, 0x71, 0x13, 0xd4, 0x48, 0x08, 0x31, 0xb0, 0xdd,
	0xb3, 0xbf, 0xf7, 0xf4, 0xee, 0xe0, 0xb2, 0xb2, 0xc6, 0x9b, 0x79, 0x61, 0x89, 0x7c, 0x16, 0x66,
	0x1c, 0x06, 0xf1, 0x12, 0x04, 0x8f, 0xa0, 0xbf, 0x36, 0x1b, 0x61, 0x85, 0xe2, 0x1c, 0x46, 0xf7,
	0x54, 0x96, 0x26, 0xa7, 0xb7, 0x77, 0x72, 0x1e, 0x11, 0xce, 0xb4, 0x50, 0x14, 0xb3, 0x29, 0x4b,
	0xa3, 0x3c, 0xcc, 0x7c, 0x06, 0x17, 0x0d, 0xe3, 0x2a, 0xa3, 0x1d, 0x61, 0x0c, 0x7d, 0x45, 0xce,
	0x89, 0xa2, 0xe5, 0x5a, 0xc9, 0xa7, 0x30, 0x58, 0x0b, 0x45, 0x8f, 0xd2, 0x79, 0x1c, 0x43, 0xaf,
	0xb6, 0xbb, 0x98, 0x4d, 0x4f, 0xd3, 0x28, 0x3f, 0x08, 0x3e, 0x83, 0xe1, 0xea, 0x00, 0x07, 0x28,
	0x81, 0x41, 0xe3, 0x6d, 0xb9, 0x6f, 0x7d, 0xf5, 0x79, 0x02, 0xa3, 0xbb, 0xba, 0xf6, 0x96, 0xec,
	0x87, 0xdc, 0x13, 0xde, 0x40, 0xb4, 0x95, 0xaa, 0x2a, 0x29, 0xdf, 0x2c, 0x71, 0x9c, 0x1d, 0xad,
	0x94, 0x35, 0xfb, 0x24, 0x49, 0xe7, 0xb5, 0x5b, 0xfb, 0x01, 0x70, 0x59, 0x4a, 0xd2, 0x7e, 0xeb,
	0x2d, 0x09, 0x25, 0x75, 0x51, 0xe7, 0x4c, 0x7e, 0x72, 0x84, 0x63, 0x24, 0x71, 0xe7, 0xeb, 0xa8,
	0x76, 0xca, 0x70, 0x05, 0x58, 0xd7, 0x22, 0xfb, 0xd7, 0xb0, 0x5f, 0x9a, 0x2d, 0x18, 0x3e, 0xc1,
	0x64, 0x27, 0x5f, 0xa5, 0xa5, 0xbd, 0x97, 0x46, 0x8b, 0xf2, 0x1f, 0x52, 0x53, 0xb6, 0x60, 0xb7,
	0xfd, 0xe7, 0x5e, 0x36, 0x2f, 0x48, 0xef, 0xce, 0x03, 0x71, 0xfd, 0x15, 0x00, 0x00, 0xff, 0xff,
	0x79, 0x72, 0x4b, 0xbf, 0x1c, 0x02, 0x00, 0x00,
}
