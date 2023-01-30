// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.9
// source: proto/data.proto

package goproto

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

type Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query string `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	PosX  int32  `protobuf:"varint,2,opt,name=pos_x,json=posX,proto3" json:"pos_x,omitempty"`
	PosY  int32  `protobuf:"varint,3,opt,name=pos_y,json=posY,proto3" json:"pos_y,omitempty"`
}

func (x *Data) Reset() {
	*x = Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_data_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data) ProtoMessage() {}

func (x *Data) ProtoReflect() protoreflect.Message {
	mi := &file_proto_data_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Data.ProtoReflect.Descriptor instead.
func (*Data) Descriptor() ([]byte, []int) {
	return file_proto_data_proto_rawDescGZIP(), []int{0}
}

func (x *Data) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

func (x *Data) GetPosX() int32 {
	if x != nil {
		return x.PosX
	}
	return 0
}

func (x *Data) GetPosY() int32 {
	if x != nil {
		return x.PosY
	}
	return 0
}

var File_proto_data_proto protoreflect.FileDescriptor

var file_proto_data_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x46, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x12, 0x14, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x13, 0x0a, 0x05, 0x70, 0x6f, 0x73, 0x5f, 0x78,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x6f, 0x73, 0x58, 0x12, 0x13, 0x0a, 0x05,
	0x70, 0x6f, 0x73, 0x5f, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x6f, 0x73,
	0x59, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x2f, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_data_proto_rawDescOnce sync.Once
	file_proto_data_proto_rawDescData = file_proto_data_proto_rawDesc
)

func file_proto_data_proto_rawDescGZIP() []byte {
	file_proto_data_proto_rawDescOnce.Do(func() {
		file_proto_data_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_data_proto_rawDescData)
	})
	return file_proto_data_proto_rawDescData
}

var file_proto_data_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_data_proto_goTypes = []interface{}{
	(*Data)(nil), // 0: proto.data
}
var file_proto_data_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_data_proto_init() }
func file_proto_data_proto_init() {
	if File_proto_data_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_data_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Data); i {
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
			RawDescriptor: file_proto_data_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_data_proto_goTypes,
		DependencyIndexes: file_proto_data_proto_depIdxs,
		MessageInfos:      file_proto_data_proto_msgTypes,
	}.Build()
	File_proto_data_proto = out.File
	file_proto_data_proto_rawDesc = nil
	file_proto_data_proto_goTypes = nil
	file_proto_data_proto_depIdxs = nil
}
