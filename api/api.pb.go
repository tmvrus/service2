// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.11.2
// source: api.proto

package api

import (
	proto "github.com/golang/protobuf/proto"
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

type OfferEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Types that are assignable to Event:
	//	*OfferEvent_OfferCreate
	//	*OfferEvent_OfferDelete
	Event isOfferEvent_Event `protobuf_oneof:"Event"`
}

func (x *OfferEvent) Reset() {
	*x = OfferEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OfferEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OfferEvent) ProtoMessage() {}

func (x *OfferEvent) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use OfferEvent.ProtoReflect.Descriptor instead.
func (*OfferEvent) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{0}
}

func (x *OfferEvent) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (m *OfferEvent) GetEvent() isOfferEvent_Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func (x *OfferEvent) GetOfferCreate() *OfferCreate {
	if x, ok := x.GetEvent().(*OfferEvent_OfferCreate); ok {
		return x.OfferCreate
	}
	return nil
}

func (x *OfferEvent) GetOfferDelete() *OfferDelete {
	if x, ok := x.GetEvent().(*OfferEvent_OfferDelete); ok {
		return x.OfferDelete
	}
	return nil
}

type isOfferEvent_Event interface {
	isOfferEvent_Event()
}

type OfferEvent_OfferCreate struct {
	OfferCreate *OfferCreate `protobuf:"bytes,2,opt,name=offer_create,json=offerCreate,proto3,oneof"`
}

type OfferEvent_OfferDelete struct {
	OfferDelete *OfferDelete `protobuf:"bytes,3,opt,name=offer_delete,json=offerDelete,proto3,oneof"`
}

func (*OfferEvent_OfferCreate) isOfferEvent_Event() {}

func (*OfferEvent_OfferDelete) isOfferEvent_Event() {}

type OfferCreate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title        string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	AdvertiserId string `protobuf:"bytes,3,opt,name=advertiser_id,json=advertiserId,proto3" json:"advertiser_id,omitempty"`
	Date         string `protobuf:"bytes,4,opt,name=date,proto3" json:"date,omitempty"`
}

func (x *OfferCreate) Reset() {
	*x = OfferCreate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OfferCreate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OfferCreate) ProtoMessage() {}

func (x *OfferCreate) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use OfferCreate.ProtoReflect.Descriptor instead.
func (*OfferCreate) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{1}
}

func (x *OfferCreate) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *OfferCreate) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *OfferCreate) GetAdvertiserId() string {
	if x != nil {
		return x.AdvertiserId
	}
	return ""
}

func (x *OfferCreate) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

type OfferDelete struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Date string `protobuf:"bytes,2,opt,name=date,proto3" json:"date,omitempty"`
}

func (x *OfferDelete) Reset() {
	*x = OfferDelete{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OfferDelete) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OfferDelete) ProtoMessage() {}

func (x *OfferDelete) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use OfferDelete.ProtoReflect.Descriptor instead.
func (*OfferDelete) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{2}
}

func (x *OfferDelete) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *OfferDelete) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

var File_api_proto protoreflect.FileDescriptor

var file_api_proto_rawDesc = []byte{
	0x0a, 0x09, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x63, 0x6f, 0x6d,
	0x2e, 0x61, 0x66, 0x66, 0x69, 0x73, 0x65, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x73, 0x22, 0xad, 0x01,
	0x0a, 0x0a, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x42, 0x0a, 0x0c,
	0x6f, 0x66, 0x66, 0x65, 0x72, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x66, 0x66, 0x69, 0x73, 0x65, 0x2e,
	0x73, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x48, 0x00, 0x52, 0x0b, 0x6f, 0x66, 0x66, 0x65, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x12, 0x42, 0x0a, 0x0c, 0x6f, 0x66, 0x66, 0x65, 0x72, 0x5f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x66, 0x66,
	0x69, 0x73, 0x65, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x48, 0x00, 0x52, 0x0b, 0x6f, 0x66, 0x66, 0x65, 0x72, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x42, 0x07, 0x0a, 0x05, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x22, 0x6c, 0x0a,
	0x0b, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x61, 0x64, 0x76, 0x65, 0x72, 0x74, 0x69, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x61, 0x64, 0x76, 0x65, 0x72,
	0x74, 0x69, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x22, 0x31, 0x0a, 0x0b, 0x4f,
	0x66, 0x66, 0x65, 0x72, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x42, 0x20,
	0x5a, 0x1e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x6d, 0x76,
	0x72, 0x75, 0x73, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x32, 0x2f, 0x61, 0x70, 0x69,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_api_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_api_proto_goTypes = []interface{}{
	(*OfferEvent)(nil),  // 0: com.affise.stats.OfferEvent
	(*OfferCreate)(nil), // 1: com.affise.stats.OfferCreate
	(*OfferDelete)(nil), // 2: com.affise.stats.OfferDelete
}
var file_api_proto_depIdxs = []int32{
	1, // 0: com.affise.stats.OfferEvent.offer_create:type_name -> com.affise.stats.OfferCreate
	2, // 1: com.affise.stats.OfferEvent.offer_delete:type_name -> com.affise.stats.OfferDelete
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
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
			switch v := v.(*OfferEvent); i {
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
			switch v := v.(*OfferCreate); i {
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
			switch v := v.(*OfferDelete); i {
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
	file_api_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*OfferEvent_OfferCreate)(nil),
		(*OfferEvent_OfferDelete)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
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
