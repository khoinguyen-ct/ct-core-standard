// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: api.proto

package ct_grpc_protobuf

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type GetAdInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ListId int64 `protobuf:"varint,1,opt,name=list_id,json=listId,proto3" json:"list_id,omitempty"`
}

func (x *GetAdInfoRequest) Reset() {
	*x = GetAdInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAdInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAdInfoRequest) ProtoMessage() {}

func (x *GetAdInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAdInfoRequest.ProtoReflect.Descriptor instead.
func (*GetAdInfoRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{0}
}

func (x *GetAdInfoRequest) GetListId() int64 {
	if x != nil {
		return x.ListId
	}
	return 0
}

type GetAdInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ad         *AdInfo       `protobuf:"bytes,1,opt,name=ad,proto3" json:"ad,omitempty"`
	Parameters []*Parameters `protobuf:"bytes,2,rep,name=parameters,proto3" json:"parameters,omitempty"`
}

func (x *GetAdInfoResponse) Reset() {
	*x = GetAdInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAdInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAdInfoResponse) ProtoMessage() {}

func (x *GetAdInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAdInfoResponse.ProtoReflect.Descriptor instead.
func (*GetAdInfoResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{1}
}

func (x *GetAdInfoResponse) GetAd() *AdInfo {
	if x != nil {
		return x.Ad
	}
	return nil
}

func (x *GetAdInfoResponse) GetParameters() []*Parameters {
	if x != nil {
		return x.Parameters
	}
	return nil
}

type AdInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AdId    int64    `protobuf:"varint,1,opt,name=ad_id,json=adId,proto3" json:"ad_id,omitempty"`
	ListId  int64    `protobuf:"varint,2,opt,name=list_id,json=listId,proto3" json:"list_id,omitempty"`
	Subject string   `protobuf:"bytes,3,opt,name=subject,proto3" json:"subject,omitempty"`
	Body    string   `protobuf:"bytes,4,opt,name=body,proto3" json:"body,omitempty"`
	Images  []string `protobuf:"bytes,5,rep,name=images,proto3" json:"images,omitempty"`
}

func (x *AdInfo) Reset() {
	*x = AdInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdInfo) ProtoMessage() {}

func (x *AdInfo) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdInfo.ProtoReflect.Descriptor instead.
func (*AdInfo) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{2}
}

func (x *AdInfo) GetAdId() int64 {
	if x != nil {
		return x.AdId
	}
	return 0
}

func (x *AdInfo) GetListId() int64 {
	if x != nil {
		return x.ListId
	}
	return 0
}

func (x *AdInfo) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

func (x *AdInfo) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

func (x *AdInfo) GetImages() []string {
	if x != nil {
		return x.Images
	}
	return nil
}

type Parameters struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Label string `protobuf:"bytes,3,opt,name=label,proto3" json:"label,omitempty"`
}

func (x *Parameters) Reset() {
	*x = Parameters{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Parameters) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Parameters) ProtoMessage() {}

func (x *Parameters) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Parameters.ProtoReflect.Descriptor instead.
func (*Parameters) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{3}
}

func (x *Parameters) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Parameters) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *Parameters) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

var File_api_proto protoreflect.FileDescriptor

var file_api_proto_rawDesc = []byte{
	0x0a, 0x09, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x63, 0x68, 0x6f,
	0x74, 0x6f, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x64, 0x5f, 0x6c, 0x69,
	0x73, 0x74, 0x69, 0x6e, 0x67, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x2b, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x41, 0x64, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x6c, 0x69, 0x73, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6c, 0x69, 0x73, 0x74, 0x49, 0x64,
	0x22, 0x8b, 0x01, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x41, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x02, 0x61, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x20, 0x2e, 0x63, 0x68, 0x6f, 0x74, 0x6f, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x31, 0x2e, 0x61, 0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x41, 0x64,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x02, 0x61, 0x64, 0x12, 0x44, 0x0a, 0x0a, 0x70, 0x61, 0x72, 0x61,
	0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x63,
	0x68, 0x6f, 0x74, 0x6f, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x64, 0x5f,
	0x6c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65,
	0x72, 0x73, 0x52, 0x0a, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x22, 0x7c,
	0x0a, 0x06, 0x41, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x13, 0x0a, 0x05, 0x61, 0x64, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x61, 0x64, 0x49, 0x64, 0x12, 0x17, 0x0a,
	0x07, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x6c, 0x69, 0x73, 0x74, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x62, 0x6f, 0x64, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x18, 0x05,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x22, 0x48, 0x0a, 0x0a,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x32, 0xa2, 0x01, 0x0a, 0x10, 0x41, 0x64, 0x4c, 0x69, 0x73,
	0x74, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x8d, 0x01, 0x0a, 0x09,
	0x47, 0x65, 0x74, 0x41, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x2a, 0x2e, 0x63, 0x68, 0x6f, 0x74,
	0x6f, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x64, 0x5f, 0x6c, 0x69, 0x73,
	0x74, 0x69, 0x6e, 0x67, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x63, 0x68, 0x6f, 0x74, 0x6f, 0x74, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67,
	0x2e, 0x47, 0x65, 0x74, 0x41, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x27, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x21, 0x12, 0x1f, 0x2f, 0x76, 0x31, 0x2f,
	0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2f, 0x61, 0x64, 0x2d, 0x6c, 0x69, 0x73, 0x74, 0x69, 0x6e,
	0x67, 0x2f, 0x7b, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x7d, 0x42, 0x27, 0x5a, 0x25, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x61, 0x72, 0x6f, 0x75, 0x73,
	0x65, 0x6c, 0x6c, 0x2f, 0x63, 0x74, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_proto_rawDescOnce sync.Once
	file_api_proto_rawDescData = file_api_proto_rawDesc
)

func file_api_proto_rawDescGZIP() []byte {
	file_api_proto_rawDescOnce.Do(func() {
		file_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_rawDescData)
	})
	return file_api_proto_rawDescData
}

var file_api_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_api_proto_goTypes = []interface{}{
	(*GetAdInfoRequest)(nil),  // 0: chotot.api.v1.ad_listing.GetAdInfoRequest
	(*GetAdInfoResponse)(nil), // 1: chotot.api.v1.ad_listing.GetAdInfoResponse
	(*AdInfo)(nil),            // 2: chotot.api.v1.ad_listing.AdInfo
	(*Parameters)(nil),        // 3: chotot.api.v1.ad_listing.Parameters
}
var file_api_proto_depIdxs = []int32{
	2, // 0: chotot.api.v1.ad_listing.GetAdInfoResponse.ad:type_name -> chotot.api.v1.ad_listing.AdInfo
	3, // 1: chotot.api.v1.ad_listing.GetAdInfoResponse.parameters:type_name -> chotot.api.v1.ad_listing.Parameters
	0, // 2: chotot.api.v1.ad_listing.AdListingService.GetAdInfo:input_type -> chotot.api.v1.ad_listing.GetAdInfoRequest
	1, // 3: chotot.api.v1.ad_listing.AdListingService.GetAdInfo:output_type -> chotot.api.v1.ad_listing.GetAdInfoResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_api_proto_init() }
func file_api_proto_init() {
	if File_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAdInfoRequest); i {
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
		file_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAdInfoResponse); i {
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
		file_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdInfo); i {
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
		file_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Parameters); i {
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
			RawDescriptor: file_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_proto_goTypes,
		DependencyIndexes: file_api_proto_depIdxs,
		MessageInfos:      file_api_proto_msgTypes,
	}.Build()
	File_api_proto = out.File
	file_api_proto_rawDesc = nil
	file_api_proto_goTypes = nil
	file_api_proto_depIdxs = nil
}
