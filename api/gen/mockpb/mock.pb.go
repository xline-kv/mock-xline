// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: mock.proto

package mockpb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type MockRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The serialized request
	Request  []byte      `protobuf:"bytes,1,opt,name=request,proto3" json:"request,omitempty"`
	Response []*Response `protobuf:"bytes,2,rep,name=response,proto3" json:"response,omitempty"`
}

func (x *MockRequest) Reset() {
	*x = MockRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mock_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MockRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MockRequest) ProtoMessage() {}

func (x *MockRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mock_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MockRequest.ProtoReflect.Descriptor instead.
func (*MockRequest) Descriptor() ([]byte, []int) {
	return file_mock_proto_rawDescGZIP(), []int{0}
}

func (x *MockRequest) GetRequest() []byte {
	if x != nil {
		return x.Request
	}
	return nil
}

func (x *MockRequest) GetResponse() []*Response {
	if x != nil {
		return x.Response
	}
	return nil
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Response:
	//
	//	*Response_Success
	//	*Response_Error
	Response isResponse_Response `protobuf_oneof:"response"`
	// The times of request when return the response
	Index uint64                 `protobuf:"varint,3,opt,name=index,proto3" json:"index,omitempty"`
	Delay *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=delay,proto3" json:"delay,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mock_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_mock_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_mock_proto_rawDescGZIP(), []int{1}
}

func (m *Response) GetResponse() isResponse_Response {
	if m != nil {
		return m.Response
	}
	return nil
}

func (x *Response) GetSuccess() []byte {
	if x, ok := x.GetResponse().(*Response_Success); ok {
		return x.Success
	}
	return nil
}

func (x *Response) GetError() *Status {
	if x, ok := x.GetResponse().(*Response_Error); ok {
		return x.Error
	}
	return nil
}

func (x *Response) GetIndex() uint64 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *Response) GetDelay() *timestamppb.Timestamp {
	if x != nil {
		return x.Delay
	}
	return nil
}

type isResponse_Response interface {
	isResponse_Response()
}

type Response_Success struct {
	// The serialized response
	Success []byte `protobuf:"bytes,1,opt,name=success,proto3,oneof"`
}

type Response_Error struct {
	Error *Status `protobuf:"bytes,2,opt,name=error,proto3,oneof"`
}

func (*Response_Success) isResponse_Response() {}

func (*Response_Error) isResponse_Response() {}

type Status struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32        `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string       `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Details []*anypb.Any `protobuf:"bytes,3,rep,name=details,proto3" json:"details,omitempty"`
}

func (x *Status) Reset() {
	*x = Status{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mock_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Status) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status) ProtoMessage() {}

func (x *Status) ProtoReflect() protoreflect.Message {
	mi := &file_mock_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Status.ProtoReflect.Descriptor instead.
func (*Status) Descriptor() ([]byte, []int) {
	return file_mock_proto_rawDescGZIP(), []int{2}
}

func (x *Status) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *Status) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Status) GetDetails() []*anypb.Any {
	if x != nil {
		return x.Details
	}
	return nil
}

type MockResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MockResponse) Reset() {
	*x = MockResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mock_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MockResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MockResponse) ProtoMessage() {}

func (x *MockResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mock_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MockResponse.ProtoReflect.Descriptor instead.
func (*MockResponse) Descriptor() ([]byte, []int) {
	return file_mock_proto_rawDescGZIP(), []int{3}
}

var File_mock_proto protoreflect.FileDescriptor

var file_mock_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x6d, 0x6f, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x6d, 0x6f,
	0x63, 0x6b, 0x70, 0x62, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x55, 0x0a, 0x0b, 0x4d, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x18, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x08, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6d, 0x6f,
	0x63, 0x6b, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x08, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xa2, 0x01, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x12, 0x26, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0e, 0x2e, 0x6d, 0x6f, 0x63, 0x6b, 0x70, 0x62, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x48,
	0x00, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65,
	0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x30,
	0x0a, 0x05, 0x64, 0x65, 0x6c, 0x61, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x05, 0x64, 0x65, 0x6c, 0x61, 0x79,
	0x42, 0x0a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x66, 0x0a, 0x06,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x2e, 0x0a, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x07, 0x64, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x73, 0x22, 0x0e, 0x0a, 0x0c, 0x4d, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x32, 0xbc, 0x01, 0x0a, 0x04, 0x4d, 0x6f, 0x63, 0x6b, 0x12, 0x3d, 0x0a,
	0x10, 0x4d, 0x6f, 0x63, 0x6b, 0x46, 0x65, 0x74, 0x63, 0x68, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x12, 0x13, 0x2e, 0x6d, 0x6f, 0x63, 0x6b, 0x70, 0x62, 0x2e, 0x4d, 0x6f, 0x63, 0x6b, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x6d, 0x6f, 0x63, 0x6b, 0x70, 0x62, 0x2e,
	0x4d, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x0b,
	0x4d, 0x6f, 0x63, 0x6b, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65, 0x12, 0x13, 0x2e, 0x6d, 0x6f,
	0x63, 0x6b, 0x70, 0x62, 0x2e, 0x4d, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x14, 0x2e, 0x6d, 0x6f, 0x63, 0x6b, 0x70, 0x62, 0x2e, 0x4d, 0x6f, 0x63, 0x6b, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b, 0x0a, 0x0e, 0x4d, 0x6f, 0x63, 0x6b, 0x57, 0x61,
	0x69, 0x74, 0x53, 0x79, 0x6e, 0x63, 0x65, 0x64, 0x12, 0x13, 0x2e, 0x6d, 0x6f, 0x63, 0x6b, 0x70,
	0x62, 0x2e, 0x4d, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e,
	0x6d, 0x6f, 0x63, 0x6b, 0x70, 0x62, 0x2e, 0x4d, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x3b, 0x6d, 0x6f, 0x63, 0x6b, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mock_proto_rawDescOnce sync.Once
	file_mock_proto_rawDescData = file_mock_proto_rawDesc
)

func file_mock_proto_rawDescGZIP() []byte {
	file_mock_proto_rawDescOnce.Do(func() {
		file_mock_proto_rawDescData = protoimpl.X.CompressGZIP(file_mock_proto_rawDescData)
	})
	return file_mock_proto_rawDescData
}

var file_mock_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_mock_proto_goTypes = []interface{}{
	(*MockRequest)(nil),           // 0: mockpb.MockRequest
	(*Response)(nil),              // 1: mockpb.Response
	(*Status)(nil),                // 2: mockpb.Status
	(*MockResponse)(nil),          // 3: mockpb.MockResponse
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
	(*anypb.Any)(nil),             // 5: google.protobuf.Any
}
var file_mock_proto_depIdxs = []int32{
	1, // 0: mockpb.MockRequest.response:type_name -> mockpb.Response
	2, // 1: mockpb.Response.error:type_name -> mockpb.Status
	4, // 2: mockpb.Response.delay:type_name -> google.protobuf.Timestamp
	5, // 3: mockpb.Status.details:type_name -> google.protobuf.Any
	0, // 4: mockpb.Mock.MockFetchCluster:input_type -> mockpb.MockRequest
	0, // 5: mockpb.Mock.MockPropose:input_type -> mockpb.MockRequest
	0, // 6: mockpb.Mock.MockWaitSynced:input_type -> mockpb.MockRequest
	3, // 7: mockpb.Mock.MockFetchCluster:output_type -> mockpb.MockResponse
	3, // 8: mockpb.Mock.MockPropose:output_type -> mockpb.MockResponse
	3, // 9: mockpb.Mock.MockWaitSynced:output_type -> mockpb.MockResponse
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_mock_proto_init() }
func file_mock_proto_init() {
	if File_mock_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mock_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MockRequest); i {
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
		file_mock_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
		file_mock_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Status); i {
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
		file_mock_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MockResponse); i {
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
	file_mock_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*Response_Success)(nil),
		(*Response_Error)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_mock_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_mock_proto_goTypes,
		DependencyIndexes: file_mock_proto_depIdxs,
		MessageInfos:      file_mock_proto_msgTypes,
	}.Build()
	File_mock_proto = out.File
	file_mock_proto_rawDesc = nil
	file_mock_proto_goTypes = nil
	file_mock_proto_depIdxs = nil
}
