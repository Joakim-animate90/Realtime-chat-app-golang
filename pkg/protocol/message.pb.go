// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v3.21.12
// source: protocol/message.proto

package protocol

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Avatar       string `protobuf:"bytes,1,opt,name=avatar,proto3" json:"avatar,omitempty"`             // avatar
	FromUsername string `protobuf:"bytes,2,opt,name=fromUsername,proto3" json:"fromUsername,omitempty"` // Send the user name of the message user
	From         string `protobuf:"bytes,3,opt,name=from,proto3" json:"from,omitempty"`                 // Send message user UUID
	To           string `protobuf:"bytes,4,opt,name=to,proto3" json:"to,omitempty"`                     // Send to UUID to the end user
	Content      string `protobuf:"bytes,5,opt,name=content,proto3" json:"content,omitempty"`           // Text message content
	ContentType  int32  `protobuf:"varint,6,opt,name=contentType,proto3" json:"contentType,omitempty"`  // Message Content Type: 1. Text 2. Ordinary File 3. Picture 4. Audio 5. Video 6. Voice Chat 7. Video chat
	Type         string `protobuf:"bytes,7,opt,name=type,proto3" json:"type,omitempty"`                 // Message Transmission Type: If it is a heartbeat message, this content is headbeat, online video or audio is webrtc
	MessageType  int32  `protobuf:"varint,8,opt,name=messageType,proto3" json:"messageType,omitempty"`  // message type, 1. Single chat 2. Group chat
	Url          string `protobuf:"bytes,9,opt,name=url,proto3" json:"url,omitempty"`                   // picture, video, voice path
	FileSuffix   string `protobuf:"bytes,10,opt,name=fileSuffix,proto3" json:"fileSuffix,omitempty"`    // File suffix, if you cannot analyze the file suffix through the binary head, use the suffix
	File         []byte `protobuf:"bytes,11,opt,name=file,proto3" json:"file,omitempty"`                // If it is a binary of pictures, files, videos, etc.
}

func (x *Message) Reset() {
	*x = Message{}
	mi := &file_protocol_message_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_message_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_protocol_message_proto_rawDescGZIP(), []int{0}
}

func (x *Message) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *Message) GetFromUsername() string {
	if x != nil {
		return x.FromUsername
	}
	return ""
}

func (x *Message) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *Message) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *Message) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Message) GetContentType() int32 {
	if x != nil {
		return x.ContentType
	}
	return 0
}

func (x *Message) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Message) GetMessageType() int32 {
	if x != nil {
		return x.MessageType
	}
	return 0
}

func (x *Message) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *Message) GetFileSuffix() string {
	if x != nil {
		return x.FileSuffix
	}
	return ""
}

func (x *Message) GetFile() []byte {
	if x != nil {
		return x.File
	}
	return nil
}

var File_protocol_message_proto protoreflect.FileDescriptor

var file_protocol_message_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x22, 0xa1, 0x02, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x12, 0x22, 0x0a, 0x0c, 0x66, 0x72, 0x6f, 0x6d, 0x55, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x66, 0x72,
	0x6f, 0x6d, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72,
	0x6f, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x0e,
	0x0a, 0x02, 0x74, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x18,
	0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x20,
	0x0a, 0x0b, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0b, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75,
	0x72, 0x6c, 0x12, 0x1e, 0x0a, 0x0a, 0x66, 0x69, 0x6c, 0x65, 0x53, 0x75, 0x66, 0x66, 0x69, 0x78,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x66, 0x69, 0x6c, 0x65, 0x53, 0x75, 0x66, 0x66,
	0x69, 0x78, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protocol_message_proto_rawDescOnce sync.Once
	file_protocol_message_proto_rawDescData = file_protocol_message_proto_rawDesc
)

func file_protocol_message_proto_rawDescGZIP() []byte {
	file_protocol_message_proto_rawDescOnce.Do(func() {
		file_protocol_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_protocol_message_proto_rawDescData)
	})
	return file_protocol_message_proto_rawDescData
}

var file_protocol_message_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_protocol_message_proto_goTypes = []any{
	(*Message)(nil), // 0: protocol.Message
}
var file_protocol_message_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protocol_message_proto_init() }
func file_protocol_message_proto_init() {
	if File_protocol_message_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protocol_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protocol_message_proto_goTypes,
		DependencyIndexes: file_protocol_message_proto_depIdxs,
		MessageInfos:      file_protocol_message_proto_msgTypes,
	}.Build()
	File_protocol_message_proto = out.File
	file_protocol_message_proto_rawDesc = nil
	file_protocol_message_proto_goTypes = nil
	file_protocol_message_proto_depIdxs = nil
}
