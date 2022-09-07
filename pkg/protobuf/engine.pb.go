// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: engine.proto

package protobuf

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

type HeartBeatsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeviceId string `protobuf:"bytes,1,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	Os       string `protobuf:"bytes,2,opt,name=os,proto3" json:"os,omitempty"`
	CpuInfo  string `protobuf:"bytes,3,opt,name=cpu_info,json=cpuInfo,proto3" json:"cpu_info,omitempty"`
	MemInfo  string `protobuf:"bytes,4,opt,name=mem_info,json=memInfo,proto3" json:"mem_info,omitempty"`
	Ip       string `protobuf:"bytes,5,opt,name=ip,proto3" json:"ip,omitempty"`
}

func (x *HeartBeatsRequest) Reset() {
	*x = HeartBeatsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_engine_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeartBeatsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeartBeatsRequest) ProtoMessage() {}

func (x *HeartBeatsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_engine_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeartBeatsRequest.ProtoReflect.Descriptor instead.
func (*HeartBeatsRequest) Descriptor() ([]byte, []int) {
	return file_engine_proto_rawDescGZIP(), []int{0}
}

func (x *HeartBeatsRequest) GetDeviceId() string {
	if x != nil {
		return x.DeviceId
	}
	return ""
}

func (x *HeartBeatsRequest) GetOs() string {
	if x != nil {
		return x.Os
	}
	return ""
}

func (x *HeartBeatsRequest) GetCpuInfo() string {
	if x != nil {
		return x.CpuInfo
	}
	return ""
}

func (x *HeartBeatsRequest) GetMemInfo() string {
	if x != nil {
		return x.MemInfo
	}
	return ""
}

func (x *HeartBeatsRequest) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

type HeartBeatsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *HeartBeatsResponse) Reset() {
	*x = HeartBeatsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_engine_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeartBeatsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeartBeatsResponse) ProtoMessage() {}

func (x *HeartBeatsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_engine_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeartBeatsResponse.ProtoReflect.Descriptor instead.
func (*HeartBeatsResponse) Descriptor() ([]byte, []int) {
	return file_engine_proto_rawDescGZIP(), []int{1}
}

var File_engine_proto protoreflect.FileDescriptor

var file_engine_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x1a, 0x09, 0x6a, 0x6f, 0x62, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x86, 0x01, 0x0a, 0x11, 0x48, 0x65, 0x61, 0x72, 0x74, 0x42, 0x65, 0x61,
	0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x6f, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x63, 0x70, 0x75, 0x5f, 0x69, 0x6e,
	0x66, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x70, 0x75, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x19, 0x0a, 0x08, 0x6d, 0x65, 0x6d, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x70, 0x22, 0x14, 0x0a, 0x12,
	0x48, 0x65, 0x61, 0x72, 0x74, 0x42, 0x65, 0x61, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x32, 0xe3, 0x01, 0x0a, 0x06, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x12, 0x47, 0x0a,
	0x0a, 0x48, 0x65, 0x61, 0x72, 0x74, 0x42, 0x65, 0x61, 0x74, 0x73, 0x12, 0x1b, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x48, 0x65, 0x61, 0x72, 0x74, 0x42, 0x65, 0x61, 0x74,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x48, 0x65, 0x61, 0x72, 0x74, 0x42, 0x65, 0x61, 0x74, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x4e, 0x65, 0x77,
	0x4a, 0x6f, 0x62, 0x12, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4a,
	0x6f, 0x62, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4a, 0x6f, 0x62, 0x47, 0x65, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x50, 0x0a, 0x11, 0x50, 0x6f, 0x70, 0x75, 0x6c, 0x61,
	0x74, 0x65, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x1c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4a, 0x6f, 0x62, 0x50, 0x6f, 0x70, 0x75, 0x6c, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4a, 0x6f, 0x62, 0x50, 0x6f, 0x70, 0x75, 0x6c, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x29, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x61, 0x74, 0x68, 0x2d, 0x72, 0x75, 0x6e, 0x2f,
	0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_engine_proto_rawDescOnce sync.Once
	file_engine_proto_rawDescData = file_engine_proto_rawDesc
)

func file_engine_proto_rawDescGZIP() []byte {
	file_engine_proto_rawDescOnce.Do(func() {
		file_engine_proto_rawDescData = protoimpl.X.CompressGZIP(file_engine_proto_rawDescData)
	})
	return file_engine_proto_rawDescData
}

var file_engine_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_engine_proto_goTypes = []interface{}{
	(*HeartBeatsRequest)(nil),   // 0: protobuf.HeartBeatsRequest
	(*HeartBeatsResponse)(nil),  // 1: protobuf.HeartBeatsResponse
	(*JobGetRequest)(nil),       // 2: protobuf.JobGetRequest
	(*JobPopulateRequest)(nil),  // 3: protobuf.JobPopulateRequest
	(*JobGetResponse)(nil),      // 4: protobuf.JobGetResponse
	(*JobPopulateResponse)(nil), // 5: protobuf.JobPopulateResponse
}
var file_engine_proto_depIdxs = []int32{
	0, // 0: protobuf.engine.HeartBeats:input_type -> protobuf.HeartBeatsRequest
	2, // 1: protobuf.engine.GetNewJob:input_type -> protobuf.JobGetRequest
	3, // 2: protobuf.engine.PopulateJobResult:input_type -> protobuf.JobPopulateRequest
	1, // 3: protobuf.engine.HeartBeats:output_type -> protobuf.HeartBeatsResponse
	4, // 4: protobuf.engine.GetNewJob:output_type -> protobuf.JobGetResponse
	5, // 5: protobuf.engine.PopulateJobResult:output_type -> protobuf.JobPopulateResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_engine_proto_init() }
func file_engine_proto_init() {
	if File_engine_proto != nil {
		return
	}
	file_job_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_engine_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HeartBeatsRequest); i {
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
		file_engine_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HeartBeatsResponse); i {
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
			RawDescriptor: file_engine_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_engine_proto_goTypes,
		DependencyIndexes: file_engine_proto_depIdxs,
		MessageInfos:      file_engine_proto_msgTypes,
	}.Build()
	File_engine_proto = out.File
	file_engine_proto_rawDesc = nil
	file_engine_proto_goTypes = nil
	file_engine_proto_depIdxs = nil
}
