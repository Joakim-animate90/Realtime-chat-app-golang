// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pkg/protocol/message.proto

package protocol

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Message struct {
	Avatar               string   `protobuf:"bytes,1,opt,name=avatar,proto3" json:"avatar,omitempty"`
	FromUsername         string   `protobuf:"bytes,2,opt,name=fromUsername,proto3" json:"fromUsername,omitempty"`
	From                 string   `protobuf:"bytes,3,opt,name=from,proto3" json:"from,omitempty"`
	To                   string   `protobuf:"bytes,4,opt,name=to,proto3" json:"to,omitempty"`
	Content              string   `protobuf:"bytes,5,opt,name=content,proto3" json:"content,omitempty"`
	ContentType          int32    `protobuf:"varint,6,opt,name=contentType,proto3" json:"contentType,omitempty"`
	Type                 string   `protobuf:"bytes,7,opt,name=type,proto3" json:"type,omitempty"`
	MessageType          int32    `protobuf:"varint,8,opt,name=messageType,proto3" json:"messageType,omitempty"`
	Url                  string   `protobuf:"bytes,9,opt,name=url,proto3" json:"url,omitempty"`
	FileSuffix           string   `protobuf:"bytes,10,opt,name=fileSuffix,proto3" json:"fileSuffix,omitempty"`
	File                 []byte   `protobuf:"bytes,11,opt,name=file,proto3" json:"file,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc7b5b78212a3e3f, []int{0}
}
func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetAvatar() string {
	if m != nil {
		return m.Avatar
	}
	return ""
}

func (m *Message) GetFromUsername() string {
	if m != nil {
		return m.FromUsername
	}
	return ""
}

func (m *Message) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *Message) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

func (m *Message) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *Message) GetContentType() int32 {
	if m != nil {
		return m.ContentType
	}
	return 0
}

func (m *Message) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Message) GetMessageType() int32 {
	if m != nil {
		return m.MessageType
	}
	return 0
}

func (m *Message) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *Message) GetFileSuffix() string {
	if m != nil {
		return m.FileSuffix
	}
	return ""
}

func (m *Message) GetFile() []byte {
	if m != nil {
		return m.File
	}
	return nil
}

func init() {
	proto.RegisterType((*Message)(nil), "protocol.Message")
}

func init() { proto.RegisterFile("pkg/protocol/message.proto", fileDescriptor_fc7b5b78212a3e3f) }

var fileDescriptor_fc7b5b78212a3e3f = []byte{
	// 231 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x86, 0x15, 0xb7, 0x4d, 0xda, 0x6b, 0x84, 0xd0, 0x0d, 0xe8, 0xc4, 0x80, 0xac, 0x4e, 0x99,
	0xda, 0x81, 0x37, 0x60, 0x67, 0x09, 0xb0, 0xb0, 0x99, 0xea, 0x5c, 0x45, 0x38, 0x71, 0xe4, 0xb8,
	0x08, 0x1e, 0x87, 0x37, 0x45, 0xbe, 0xa4, 0x28, 0xdd, 0xbe, 0xff, 0xfb, 0xf5, 0x5b, 0x3a, 0xc3,
	0x7d, 0xff, 0x79, 0x3a, 0xf4, 0xc1, 0x47, 0x7f, 0xf4, 0xee, 0xd0, 0xf2, 0x30, 0x98, 0x13, 0xef,
	0x45, 0xe0, 0xfa, 0xe2, 0x77, 0xbf, 0x0a, 0x8a, 0xe7, 0xb1, 0xc3, 0x3b, 0xc8, 0xcd, 0x97, 0x89,
	0x26, 0x50, 0xa6, 0xb3, 0x6a, 0x53, 0x4f, 0x09, 0x77, 0x50, 0xda, 0xe0, 0xdb, 0xb7, 0x81, 0x43,
	0x67, 0x5a, 0x26, 0x25, 0xed, 0x95, 0x43, 0x84, 0x65, 0xca, 0xb4, 0x90, 0x4e, 0x18, 0x6f, 0x40,
	0x45, 0x4f, 0x4b, 0x31, 0x2a, 0x7a, 0x24, 0x28, 0x8e, 0xbe, 0x8b, 0xdc, 0x45, 0x5a, 0x89, 0xbc,
	0x44, 0xd4, 0xb0, 0x9d, 0xf0, 0xf5, 0xa7, 0x67, 0xca, 0x75, 0x56, 0xad, 0xea, 0xb9, 0x4a, 0xef,
	0xc7, 0x54, 0x15, 0xe3, 0xfb, 0x89, 0xd3, 0x6a, 0x3a, 0x4b, 0x56, 0xeb, 0x71, 0x35, 0x53, 0x78,
	0x0b, 0x8b, 0x73, 0x70, 0xb4, 0x91, 0x51, 0x42, 0x7c, 0x00, 0xb0, 0x8d, 0xe3, 0x97, 0xb3, 0xb5,
	0xcd, 0x37, 0x81, 0x14, 0x33, 0x23, 0x77, 0x34, 0x8e, 0x69, 0xab, 0xb3, 0xaa, 0xac, 0x85, 0x9f,
	0xca, 0x77, 0xd8, 0xff, 0xff, 0xe4, 0x47, 0x2e, 0xf4, 0xf8, 0x17, 0x00, 0x00, 0xff, 0xff, 0x89,
	0x8f, 0x3c, 0x3c, 0x60, 0x01, 0x00, 0x00,
}
