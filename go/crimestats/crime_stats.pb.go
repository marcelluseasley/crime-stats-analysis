// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: proto/crime_stats.proto

package crimestats

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

type CrimeStatsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Street  string `protobuf:"bytes,1,opt,name=street,proto3" json:"street,omitempty"`
	City    string `protobuf:"bytes,2,opt,name=city,proto3" json:"city,omitempty"`
	State   string `protobuf:"bytes,3,opt,name=state,proto3" json:"state,omitempty"`
	Zipcode string `protobuf:"bytes,4,opt,name=zipcode,proto3" json:"zipcode,omitempty"`
}

func (x *CrimeStatsRequest) Reset() {
	*x = CrimeStatsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_crime_stats_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CrimeStatsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CrimeStatsRequest) ProtoMessage() {}

func (x *CrimeStatsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_crime_stats_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CrimeStatsRequest.ProtoReflect.Descriptor instead.
func (*CrimeStatsRequest) Descriptor() ([]byte, []int) {
	return file_proto_crime_stats_proto_rawDescGZIP(), []int{0}
}

func (x *CrimeStatsRequest) GetStreet() string {
	if x != nil {
		return x.Street
	}
	return ""
}

func (x *CrimeStatsRequest) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *CrimeStatsRequest) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *CrimeStatsRequest) GetZipcode() string {
	if x != nil {
		return x.Zipcode
	}
	return ""
}

type CrimeStatsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*CrimeStatsResponse_Crime `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *CrimeStatsResponse) Reset() {
	*x = CrimeStatsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_crime_stats_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CrimeStatsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CrimeStatsResponse) ProtoMessage() {}

func (x *CrimeStatsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_crime_stats_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CrimeStatsResponse.ProtoReflect.Descriptor instead.
func (*CrimeStatsResponse) Descriptor() ([]byte, []int) {
	return file_proto_crime_stats_proto_rawDescGZIP(), []int{1}
}

func (x *CrimeStatsResponse) GetItems() []*CrimeStatsResponse_Crime {
	if x != nil {
		return x.Items
	}
	return nil
}

type CrimeStatsResponse_Crime struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Description string    `protobuf:"bytes,1,opt,name=description,proto3" json:"description,omitempty"`
	Datetime    string    `protobuf:"bytes,2,opt,name=datetime,proto3" json:"datetime,omitempty"`
	Location    []float64 `protobuf:"fixed64,3,rep,packed,name=location,proto3" json:"location,omitempty"`
}

func (x *CrimeStatsResponse_Crime) Reset() {
	*x = CrimeStatsResponse_Crime{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_crime_stats_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CrimeStatsResponse_Crime) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CrimeStatsResponse_Crime) ProtoMessage() {}

func (x *CrimeStatsResponse_Crime) ProtoReflect() protoreflect.Message {
	mi := &file_proto_crime_stats_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CrimeStatsResponse_Crime.ProtoReflect.Descriptor instead.
func (*CrimeStatsResponse_Crime) Descriptor() ([]byte, []int) {
	return file_proto_crime_stats_proto_rawDescGZIP(), []int{1, 0}
}

func (x *CrimeStatsResponse_Crime) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CrimeStatsResponse_Crime) GetDatetime() string {
	if x != nil {
		return x.Datetime
	}
	return ""
}

func (x *CrimeStatsResponse_Crime) GetLocation() []float64 {
	if x != nil {
		return x.Location
	}
	return nil
}

var File_proto_crime_stats_proto protoreflect.FileDescriptor

var file_proto_crime_stats_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x72, 0x69, 0x6d, 0x65, 0x5f, 0x73, 0x74,
	0x61, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x63, 0x72, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x22, 0x6f, 0x0a, 0x11, 0x43, 0x72, 0x69, 0x6d,
	0x65, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x74, 0x72, 0x65, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x7a, 0x69, 0x70, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x7a, 0x69, 0x70, 0x63, 0x6f, 0x64, 0x65, 0x22, 0xb6, 0x01, 0x0a, 0x12, 0x43, 0x72,
	0x69, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x3d, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x27, 0x2e, 0x63, 0x72, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x72, 0x69, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x2e, 0x43, 0x72, 0x69, 0x6d, 0x65, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x1a,
	0x61, 0x0a, 0x05, 0x43, 0x72, 0x69, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x61,
	0x74, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x61,
	0x74, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x03, 0x20, 0x03, 0x28, 0x01, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x32, 0x66, 0x0a, 0x11, 0x43, 0x72, 0x69, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x74, 0x73,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x51, 0x0a, 0x0a, 0x43, 0x72, 0x69, 0x6d, 0x65,
	0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x20, 0x2e, 0x63, 0x72, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x69, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x74, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x63, 0x72, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x69, 0x6d, 0x65, 0x53, 0x74, 0x61,
	0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x3c, 0x5a, 0x3a, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x61, 0x72, 0x63, 0x65, 0x6c, 0x6c,
	0x75, 0x73, 0x65, 0x61, 0x73, 0x6c, 0x65, 0x79, 0x2f, 0x63, 0x72, 0x69, 0x6d, 0x65, 0x2d, 0x73,
	0x74, 0x61, 0x74, 0x73, 0x2d, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x2f, 0x63, 0x72,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x74, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_crime_stats_proto_rawDescOnce sync.Once
	file_proto_crime_stats_proto_rawDescData = file_proto_crime_stats_proto_rawDesc
)

func file_proto_crime_stats_proto_rawDescGZIP() []byte {
	file_proto_crime_stats_proto_rawDescOnce.Do(func() {
		file_proto_crime_stats_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_crime_stats_proto_rawDescData)
	})
	return file_proto_crime_stats_proto_rawDescData
}

var file_proto_crime_stats_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_crime_stats_proto_goTypes = []interface{}{
	(*CrimeStatsRequest)(nil),        // 0: crimestats.v1.CrimeStatsRequest
	(*CrimeStatsResponse)(nil),       // 1: crimestats.v1.CrimeStatsResponse
	(*CrimeStatsResponse_Crime)(nil), // 2: crimestats.v1.CrimeStatsResponse.Crime
}
var file_proto_crime_stats_proto_depIdxs = []int32{
	2, // 0: crimestats.v1.CrimeStatsResponse.items:type_name -> crimestats.v1.CrimeStatsResponse.Crime
	0, // 1: crimestats.v1.CrimeStatsService.CrimeStats:input_type -> crimestats.v1.CrimeStatsRequest
	1, // 2: crimestats.v1.CrimeStatsService.CrimeStats:output_type -> crimestats.v1.CrimeStatsResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_crime_stats_proto_init() }
func file_proto_crime_stats_proto_init() {
	if File_proto_crime_stats_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_crime_stats_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CrimeStatsRequest); i {
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
		file_proto_crime_stats_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CrimeStatsResponse); i {
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
		file_proto_crime_stats_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CrimeStatsResponse_Crime); i {
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
			RawDescriptor: file_proto_crime_stats_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_crime_stats_proto_goTypes,
		DependencyIndexes: file_proto_crime_stats_proto_depIdxs,
		MessageInfos:      file_proto_crime_stats_proto_msgTypes,
	}.Build()
	File_proto_crime_stats_proto = out.File
	file_proto_crime_stats_proto_rawDesc = nil
	file_proto_crime_stats_proto_goTypes = nil
	file_proto_crime_stats_proto_depIdxs = nil
}
