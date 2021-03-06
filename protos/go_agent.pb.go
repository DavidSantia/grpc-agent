// Wrapper for New Relic Go Agent

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0-devel
// 	protoc        v3.11.4
// source: go_agent.proto

package protos

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	License string `protobuf:"bytes,2,opt,name=license,proto3" json:"license,omitempty"`
	Host    string `protobuf:"bytes,3,opt,name=host,proto3" json:"host,omitempty"`
}

func (x *Config) Reset() {
	*x = Config{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_agent_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_go_agent_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Config.ProtoReflect.Descriptor instead.
func (*Config) Descriptor() ([]byte, []int) {
	return file_go_agent_proto_rawDescGZIP(), []int{0}
}

func (x *Config) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Config) GetLicense() string {
	if x != nil {
		return x.License
	}
	return ""
}

func (x *Config) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

type Index struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Idx uint64 `protobuf:"varint,1,opt,name=idx,proto3" json:"idx,omitempty"`
}

func (x *Index) Reset() {
	*x = Index{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_agent_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Index) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Index) ProtoMessage() {}

func (x *Index) ProtoReflect() protoreflect.Message {
	mi := &file_go_agent_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Index.ProtoReflect.Descriptor instead.
func (*Index) Descriptor() ([]byte, []int) {
	return file_go_agent_proto_rawDescGZIP(), []int{1}
}

func (x *Index) GetIdx() uint64 {
	if x != nil {
		return x.Idx
	}
	return 0
}

type NameIndex struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Idx  uint64 `protobuf:"varint,1,opt,name=idx,proto3" json:"idx,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *NameIndex) Reset() {
	*x = NameIndex{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_agent_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NameIndex) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NameIndex) ProtoMessage() {}

func (x *NameIndex) ProtoReflect() protoreflect.Message {
	mi := &file_go_agent_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NameIndex.ProtoReflect.Descriptor instead.
func (*NameIndex) Descriptor() ([]byte, []int) {
	return file_go_agent_proto_rawDescGZIP(), []int{2}
}

func (x *NameIndex) GetIdx() uint64 {
	if x != nil {
		return x.Idx
	}
	return 0
}

func (x *NameIndex) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type LogIndex struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Idx     uint64 `protobuf:"varint,1,opt,name=idx,proto3" json:"idx,omitempty"`
	Level   int32  `protobuf:"varint,2,opt,name=level,proto3" json:"level,omitempty"`
	Message string `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *LogIndex) Reset() {
	*x = LogIndex{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_agent_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogIndex) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogIndex) ProtoMessage() {}

func (x *LogIndex) ProtoReflect() protoreflect.Message {
	mi := &file_go_agent_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogIndex.ProtoReflect.Descriptor instead.
func (*LogIndex) Descriptor() ([]byte, []int) {
	return file_go_agent_proto_rawDescGZIP(), []int{3}
}

func (x *LogIndex) GetIdx() uint64 {
	if x != nil {
		return x.Idx
	}
	return 0
}

func (x *LogIndex) GetLevel() int32 {
	if x != nil {
		return x.Level
	}
	return 0
}

func (x *LogIndex) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_go_agent_proto protoreflect.FileDescriptor

var file_go_agent_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x67, 0x6f, 0x5f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x22, 0x4a, 0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x68, 0x6f, 0x73, 0x74, 0x22, 0x19, 0x0a, 0x05, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x10, 0x0a,
	0x03, 0x69, 0x64, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x69, 0x64, 0x78, 0x22,
	0x31, 0x0a, 0x09, 0x4e, 0x61, 0x6d, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x10, 0x0a, 0x03,
	0x69, 0x64, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x69, 0x64, 0x78, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x22, 0x4c, 0x0a, 0x08, 0x4c, 0x6f, 0x67, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x10,
	0x0a, 0x03, 0x69, 0x64, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x69, 0x64, 0x78,
	0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x32, 0x94, 0x02, 0x0a, 0x07, 0x47, 0x6f, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x12, 0x2c, 0x0a, 0x09,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x70, 0x70, 0x12, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x1a, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x22, 0x00, 0x12, 0x2c, 0x0a, 0x06, 0x4e, 0x65,
	0x77, 0x54, 0x78, 0x6e, 0x12, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x4e, 0x61,
	0x6d, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x1a, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x22, 0x00, 0x12, 0x2c, 0x0a, 0x06, 0x4e, 0x65, 0x77, 0x53,
	0x65, 0x67, 0x12, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x4e, 0x61, 0x6d, 0x65,
	0x49, 0x6e, 0x64, 0x65, 0x78, 0x1a, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x49,
	0x6e, 0x64, 0x65, 0x78, 0x22, 0x00, 0x12, 0x28, 0x0a, 0x06, 0x45, 0x6e, 0x64, 0x53, 0x65, 0x67,
	0x12, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x1a,
	0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x22, 0x00,
	0x12, 0x28, 0x0a, 0x06, 0x45, 0x6e, 0x64, 0x54, 0x78, 0x6e, 0x12, 0x0d, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x1a, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x22, 0x00, 0x12, 0x2b, 0x0a, 0x06, 0x4c, 0x6f,
	0x67, 0x54, 0x78, 0x6e, 0x12, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x4c, 0x6f,
	0x67, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x1a, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e,
	0x49, 0x6e, 0x64, 0x65, 0x78, 0x22, 0x00, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x3b, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_go_agent_proto_rawDescOnce sync.Once
	file_go_agent_proto_rawDescData = file_go_agent_proto_rawDesc
)

func file_go_agent_proto_rawDescGZIP() []byte {
	file_go_agent_proto_rawDescOnce.Do(func() {
		file_go_agent_proto_rawDescData = protoimpl.X.CompressGZIP(file_go_agent_proto_rawDescData)
	})
	return file_go_agent_proto_rawDescData
}

var file_go_agent_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_go_agent_proto_goTypes = []interface{}{
	(*Config)(nil),    // 0: protos.Config
	(*Index)(nil),     // 1: protos.Index
	(*NameIndex)(nil), // 2: protos.NameIndex
	(*LogIndex)(nil),  // 3: protos.LogIndex
}
var file_go_agent_proto_depIdxs = []int32{
	0, // 0: protos.GoAgent.CreateApp:input_type -> protos.Config
	2, // 1: protos.GoAgent.NewTxn:input_type -> protos.NameIndex
	2, // 2: protos.GoAgent.NewSeg:input_type -> protos.NameIndex
	1, // 3: protos.GoAgent.EndSeg:input_type -> protos.Index
	1, // 4: protos.GoAgent.EndTxn:input_type -> protos.Index
	3, // 5: protos.GoAgent.LogTxn:input_type -> protos.LogIndex
	1, // 6: protos.GoAgent.CreateApp:output_type -> protos.Index
	1, // 7: protos.GoAgent.NewTxn:output_type -> protos.Index
	1, // 8: protos.GoAgent.NewSeg:output_type -> protos.Index
	1, // 9: protos.GoAgent.EndSeg:output_type -> protos.Index
	1, // 10: protos.GoAgent.EndTxn:output_type -> protos.Index
	1, // 11: protos.GoAgent.LogTxn:output_type -> protos.Index
	6, // [6:12] is the sub-list for method output_type
	0, // [0:6] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_go_agent_proto_init() }
func file_go_agent_proto_init() {
	if File_go_agent_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_go_agent_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Config); i {
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
		file_go_agent_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Index); i {
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
		file_go_agent_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NameIndex); i {
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
		file_go_agent_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogIndex); i {
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
			RawDescriptor: file_go_agent_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_go_agent_proto_goTypes,
		DependencyIndexes: file_go_agent_proto_depIdxs,
		MessageInfos:      file_go_agent_proto_msgTypes,
	}.Build()
	File_go_agent_proto = out.File
	file_go_agent_proto_rawDesc = nil
	file_go_agent_proto_goTypes = nil
	file_go_agent_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// GoAgentClient is the client API for GoAgent service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GoAgentClient interface {
	CreateApp(ctx context.Context, in *Config, opts ...grpc.CallOption) (*Index, error)
	NewTxn(ctx context.Context, in *NameIndex, opts ...grpc.CallOption) (*Index, error)
	NewSeg(ctx context.Context, in *NameIndex, opts ...grpc.CallOption) (*Index, error)
	EndSeg(ctx context.Context, in *Index, opts ...grpc.CallOption) (*Index, error)
	EndTxn(ctx context.Context, in *Index, opts ...grpc.CallOption) (*Index, error)
	LogTxn(ctx context.Context, in *LogIndex, opts ...grpc.CallOption) (*Index, error)
}

type goAgentClient struct {
	cc grpc.ClientConnInterface
}

func NewGoAgentClient(cc grpc.ClientConnInterface) GoAgentClient {
	return &goAgentClient{cc}
}

func (c *goAgentClient) CreateApp(ctx context.Context, in *Config, opts ...grpc.CallOption) (*Index, error) {
	out := new(Index)
	err := c.cc.Invoke(ctx, "/protos.GoAgent/CreateApp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goAgentClient) NewTxn(ctx context.Context, in *NameIndex, opts ...grpc.CallOption) (*Index, error) {
	out := new(Index)
	err := c.cc.Invoke(ctx, "/protos.GoAgent/NewTxn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goAgentClient) NewSeg(ctx context.Context, in *NameIndex, opts ...grpc.CallOption) (*Index, error) {
	out := new(Index)
	err := c.cc.Invoke(ctx, "/protos.GoAgent/NewSeg", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goAgentClient) EndSeg(ctx context.Context, in *Index, opts ...grpc.CallOption) (*Index, error) {
	out := new(Index)
	err := c.cc.Invoke(ctx, "/protos.GoAgent/EndSeg", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goAgentClient) EndTxn(ctx context.Context, in *Index, opts ...grpc.CallOption) (*Index, error) {
	out := new(Index)
	err := c.cc.Invoke(ctx, "/protos.GoAgent/EndTxn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goAgentClient) LogTxn(ctx context.Context, in *LogIndex, opts ...grpc.CallOption) (*Index, error) {
	out := new(Index)
	err := c.cc.Invoke(ctx, "/protos.GoAgent/LogTxn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GoAgentServer is the server API for GoAgent service.
type GoAgentServer interface {
	CreateApp(context.Context, *Config) (*Index, error)
	NewTxn(context.Context, *NameIndex) (*Index, error)
	NewSeg(context.Context, *NameIndex) (*Index, error)
	EndSeg(context.Context, *Index) (*Index, error)
	EndTxn(context.Context, *Index) (*Index, error)
	LogTxn(context.Context, *LogIndex) (*Index, error)
}

// UnimplementedGoAgentServer can be embedded to have forward compatible implementations.
type UnimplementedGoAgentServer struct {
}

func (*UnimplementedGoAgentServer) CreateApp(context.Context, *Config) (*Index, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateApp not implemented")
}
func (*UnimplementedGoAgentServer) NewTxn(context.Context, *NameIndex) (*Index, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewTxn not implemented")
}
func (*UnimplementedGoAgentServer) NewSeg(context.Context, *NameIndex) (*Index, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewSeg not implemented")
}
func (*UnimplementedGoAgentServer) EndSeg(context.Context, *Index) (*Index, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EndSeg not implemented")
}
func (*UnimplementedGoAgentServer) EndTxn(context.Context, *Index) (*Index, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EndTxn not implemented")
}
func (*UnimplementedGoAgentServer) LogTxn(context.Context, *LogIndex) (*Index, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LogTxn not implemented")
}

func RegisterGoAgentServer(s *grpc.Server, srv GoAgentServer) {
	s.RegisterService(&_GoAgent_serviceDesc, srv)
}

func _GoAgent_CreateApp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Config)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoAgentServer).CreateApp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.GoAgent/CreateApp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoAgentServer).CreateApp(ctx, req.(*Config))
	}
	return interceptor(ctx, in, info, handler)
}

func _GoAgent_NewTxn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NameIndex)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoAgentServer).NewTxn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.GoAgent/NewTxn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoAgentServer).NewTxn(ctx, req.(*NameIndex))
	}
	return interceptor(ctx, in, info, handler)
}

func _GoAgent_NewSeg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NameIndex)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoAgentServer).NewSeg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.GoAgent/NewSeg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoAgentServer).NewSeg(ctx, req.(*NameIndex))
	}
	return interceptor(ctx, in, info, handler)
}

func _GoAgent_EndSeg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Index)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoAgentServer).EndSeg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.GoAgent/EndSeg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoAgentServer).EndSeg(ctx, req.(*Index))
	}
	return interceptor(ctx, in, info, handler)
}

func _GoAgent_EndTxn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Index)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoAgentServer).EndTxn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.GoAgent/EndTxn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoAgentServer).EndTxn(ctx, req.(*Index))
	}
	return interceptor(ctx, in, info, handler)
}

func _GoAgent_LogTxn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogIndex)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoAgentServer).LogTxn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.GoAgent/LogTxn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoAgentServer).LogTxn(ctx, req.(*LogIndex))
	}
	return interceptor(ctx, in, info, handler)
}

var _GoAgent_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protos.GoAgent",
	HandlerType: (*GoAgentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateApp",
			Handler:    _GoAgent_CreateApp_Handler,
		},
		{
			MethodName: "NewTxn",
			Handler:    _GoAgent_NewTxn_Handler,
		},
		{
			MethodName: "NewSeg",
			Handler:    _GoAgent_NewSeg_Handler,
		},
		{
			MethodName: "EndSeg",
			Handler:    _GoAgent_EndSeg_Handler,
		},
		{
			MethodName: "EndTxn",
			Handler:    _GoAgent_EndTxn_Handler,
		},
		{
			MethodName: "LogTxn",
			Handler:    _GoAgent_LogTxn_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "go_agent.proto",
}
