// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.12
// source: waCommon/WACommon.proto

package waCommon

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"

	_ "embed"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type FutureProofBehavior int32

const (
	FutureProofBehavior_PLACEHOLDER    FutureProofBehavior = 0
	FutureProofBehavior_NO_PLACEHOLDER FutureProofBehavior = 1
	FutureProofBehavior_IGNORE         FutureProofBehavior = 2
)

// Enum value maps for FutureProofBehavior.
var (
	FutureProofBehavior_name = map[int32]string{
		0: "PLACEHOLDER",
		1: "NO_PLACEHOLDER",
		2: "IGNORE",
	}
	FutureProofBehavior_value = map[string]int32{
		"PLACEHOLDER":    0,
		"NO_PLACEHOLDER": 1,
		"IGNORE":         2,
	}
)

func (x FutureProofBehavior) Enum() *FutureProofBehavior {
	p := new(FutureProofBehavior)
	*p = x
	return p
}

func (x FutureProofBehavior) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FutureProofBehavior) Descriptor() protoreflect.EnumDescriptor {
	return file_waCommon_WACommon_proto_enumTypes[0].Descriptor()
}

func (FutureProofBehavior) Type() protoreflect.EnumType {
	return &file_waCommon_WACommon_proto_enumTypes[0]
}

func (x FutureProofBehavior) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FutureProofBehavior.Descriptor instead.
func (FutureProofBehavior) EnumDescriptor() ([]byte, []int) {
	return file_waCommon_WACommon_proto_rawDescGZIP(), []int{0}
}

type Command_CommandType int32

const (
	Command_COMMANDTYPE_UNKNOWN Command_CommandType = 0
	Command_EVERYONE            Command_CommandType = 1
	Command_SILENT              Command_CommandType = 2
	Command_AI                  Command_CommandType = 3
)

// Enum value maps for Command_CommandType.
var (
	Command_CommandType_name = map[int32]string{
		0: "COMMANDTYPE_UNKNOWN",
		1: "EVERYONE",
		2: "SILENT",
		3: "AI",
	}
	Command_CommandType_value = map[string]int32{
		"COMMANDTYPE_UNKNOWN": 0,
		"EVERYONE":            1,
		"SILENT":              2,
		"AI":                  3,
	}
)

func (x Command_CommandType) Enum() *Command_CommandType {
	p := new(Command_CommandType)
	*p = x
	return p
}

func (x Command_CommandType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Command_CommandType) Descriptor() protoreflect.EnumDescriptor {
	return file_waCommon_WACommon_proto_enumTypes[1].Descriptor()
}

func (Command_CommandType) Type() protoreflect.EnumType {
	return &file_waCommon_WACommon_proto_enumTypes[1]
}

func (x Command_CommandType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Command_CommandType.Descriptor instead.
func (Command_CommandType) EnumDescriptor() ([]byte, []int) {
	return file_waCommon_WACommon_proto_rawDescGZIP(), []int{1, 0}
}

type MessageKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RemoteJID   string `protobuf:"bytes,1,opt,name=remoteJID,proto3" json:"remoteJID,omitempty"`
	FromMe      bool   `protobuf:"varint,2,opt,name=fromMe,proto3" json:"fromMe,omitempty"`
	ID          string `protobuf:"bytes,3,opt,name=ID,proto3" json:"ID,omitempty"`
	Participant string `protobuf:"bytes,4,opt,name=participant,proto3" json:"participant,omitempty"`
}

func (x *MessageKey) Reset() {
	*x = MessageKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_waCommon_WACommon_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageKey) ProtoMessage() {}

func (x *MessageKey) ProtoReflect() protoreflect.Message {
	mi := &file_waCommon_WACommon_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageKey.ProtoReflect.Descriptor instead.
func (*MessageKey) Descriptor() ([]byte, []int) {
	return file_waCommon_WACommon_proto_rawDescGZIP(), []int{0}
}

func (x *MessageKey) GetRemoteJID() string {
	if x != nil {
		return x.RemoteJID
	}
	return ""
}

func (x *MessageKey) GetFromMe() bool {
	if x != nil {
		return x.FromMe
	}
	return false
}

func (x *MessageKey) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *MessageKey) GetParticipant() string {
	if x != nil {
		return x.Participant
	}
	return ""
}

type Command struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommandType     Command_CommandType `protobuf:"varint,1,opt,name=commandType,proto3,enum=WACommon.Command_CommandType" json:"commandType,omitempty"`
	Offset          uint32              `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	Length          uint32              `protobuf:"varint,3,opt,name=length,proto3" json:"length,omitempty"`
	ValidationToken string              `protobuf:"bytes,4,opt,name=validationToken,proto3" json:"validationToken,omitempty"`
}

func (x *Command) Reset() {
	*x = Command{}
	if protoimpl.UnsafeEnabled {
		mi := &file_waCommon_WACommon_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Command) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Command) ProtoMessage() {}

func (x *Command) ProtoReflect() protoreflect.Message {
	mi := &file_waCommon_WACommon_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Command.ProtoReflect.Descriptor instead.
func (*Command) Descriptor() ([]byte, []int) {
	return file_waCommon_WACommon_proto_rawDescGZIP(), []int{1}
}

func (x *Command) GetCommandType() Command_CommandType {
	if x != nil {
		return x.CommandType
	}
	return Command_COMMANDTYPE_UNKNOWN
}

func (x *Command) GetOffset() uint32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *Command) GetLength() uint32 {
	if x != nil {
		return x.Length
	}
	return 0
}

func (x *Command) GetValidationToken() string {
	if x != nil {
		return x.ValidationToken
	}
	return ""
}

type MessageText struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text         string     `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	MentionedJID []string   `protobuf:"bytes,2,rep,name=mentionedJID,proto3" json:"mentionedJID,omitempty"`
	Commands     []*Command `protobuf:"bytes,3,rep,name=commands,proto3" json:"commands,omitempty"`
}

func (x *MessageText) Reset() {
	*x = MessageText{}
	if protoimpl.UnsafeEnabled {
		mi := &file_waCommon_WACommon_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageText) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageText) ProtoMessage() {}

func (x *MessageText) ProtoReflect() protoreflect.Message {
	mi := &file_waCommon_WACommon_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageText.ProtoReflect.Descriptor instead.
func (*MessageText) Descriptor() ([]byte, []int) {
	return file_waCommon_WACommon_proto_rawDescGZIP(), []int{2}
}

func (x *MessageText) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *MessageText) GetMentionedJID() []string {
	if x != nil {
		return x.MentionedJID
	}
	return nil
}

func (x *MessageText) GetCommands() []*Command {
	if x != nil {
		return x.Commands
	}
	return nil
}

type SubProtocol struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Payload []byte `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
	Version int32  `protobuf:"varint,2,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *SubProtocol) Reset() {
	*x = SubProtocol{}
	if protoimpl.UnsafeEnabled {
		mi := &file_waCommon_WACommon_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubProtocol) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubProtocol) ProtoMessage() {}

func (x *SubProtocol) ProtoReflect() protoreflect.Message {
	mi := &file_waCommon_WACommon_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubProtocol.ProtoReflect.Descriptor instead.
func (*SubProtocol) Descriptor() ([]byte, []int) {
	return file_waCommon_WACommon_proto_rawDescGZIP(), []int{3}
}

func (x *SubProtocol) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

func (x *SubProtocol) GetVersion() int32 {
	if x != nil {
		return x.Version
	}
	return 0
}

var File_waCommon_WACommon_proto protoreflect.FileDescriptor

//go:embed WACommon.pb.raw
var file_waCommon_WACommon_proto_rawDesc []byte

var (
	file_waCommon_WACommon_proto_rawDescOnce sync.Once
	file_waCommon_WACommon_proto_rawDescData = file_waCommon_WACommon_proto_rawDesc
)

func file_waCommon_WACommon_proto_rawDescGZIP() []byte {
	file_waCommon_WACommon_proto_rawDescOnce.Do(func() {
		file_waCommon_WACommon_proto_rawDescData = protoimpl.X.CompressGZIP(file_waCommon_WACommon_proto_rawDescData)
	})
	return file_waCommon_WACommon_proto_rawDescData
}

var file_waCommon_WACommon_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_waCommon_WACommon_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_waCommon_WACommon_proto_goTypes = []interface{}{
	(FutureProofBehavior)(0), // 0: WACommon.FutureProofBehavior
	(Command_CommandType)(0), // 1: WACommon.Command.CommandType
	(*MessageKey)(nil),       // 2: WACommon.MessageKey
	(*Command)(nil),          // 3: WACommon.Command
	(*MessageText)(nil),      // 4: WACommon.MessageText
	(*SubProtocol)(nil),      // 5: WACommon.SubProtocol
}
var file_waCommon_WACommon_proto_depIdxs = []int32{
	1, // 0: WACommon.Command.commandType:type_name -> WACommon.Command.CommandType
	3, // 1: WACommon.MessageText.commands:type_name -> WACommon.Command
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_waCommon_WACommon_proto_init() }
func file_waCommon_WACommon_proto_init() {
	if File_waCommon_WACommon_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_waCommon_WACommon_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageKey); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_waCommon_WACommon_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Command); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_waCommon_WACommon_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageText); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_waCommon_WACommon_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubProtocol); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_waCommon_WACommon_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_waCommon_WACommon_proto_goTypes,
		DependencyIndexes: file_waCommon_WACommon_proto_depIdxs,
		EnumInfos:         file_waCommon_WACommon_proto_enumTypes,
		MessageInfos:      file_waCommon_WACommon_proto_msgTypes,
	}.Build()
	File_waCommon_WACommon_proto = out.File
	file_waCommon_WACommon_proto_rawDesc = nil
	file_waCommon_WACommon_proto_goTypes = nil
	file_waCommon_WACommon_proto_depIdxs = nil
}
