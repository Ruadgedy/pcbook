// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.15.5
// source: laptop_service.proto

package pb

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

type CreateLaptopRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Laptop *Laptop `protobuf:"bytes,1,opt,name=laptop,proto3" json:"laptop,omitempty"`
}

func (x *CreateLaptopRequest) Reset() {
	*x = CreateLaptopRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_laptop_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateLaptopRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateLaptopRequest) ProtoMessage() {}

func (x *CreateLaptopRequest) ProtoReflect() protoreflect.Message {
	mi := &file_laptop_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateLaptopRequest.ProtoReflect.Descriptor instead.
func (*CreateLaptopRequest) Descriptor() ([]byte, []int) {
	return file_laptop_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreateLaptopRequest) GetLaptop() *Laptop {
	if x != nil {
		return x.Laptop
	}
	return nil
}

type CreateLaptopResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CreateLaptopResponse) Reset() {
	*x = CreateLaptopResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_laptop_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateLaptopResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateLaptopResponse) ProtoMessage() {}

func (x *CreateLaptopResponse) ProtoReflect() protoreflect.Message {
	mi := &file_laptop_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateLaptopResponse.ProtoReflect.Descriptor instead.
func (*CreateLaptopResponse) Descriptor() ([]byte, []int) {
	return file_laptop_service_proto_rawDescGZIP(), []int{1}
}

func (x *CreateLaptopResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type SearchLaptopRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filter *Filter `protobuf:"bytes,1,opt,name=filter,proto3" json:"filter,omitempty"`
}

func (x *SearchLaptopRequest) Reset() {
	*x = SearchLaptopRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_laptop_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchLaptopRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchLaptopRequest) ProtoMessage() {}

func (x *SearchLaptopRequest) ProtoReflect() protoreflect.Message {
	mi := &file_laptop_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchLaptopRequest.ProtoReflect.Descriptor instead.
func (*SearchLaptopRequest) Descriptor() ([]byte, []int) {
	return file_laptop_service_proto_rawDescGZIP(), []int{2}
}

func (x *SearchLaptopRequest) GetFilter() *Filter {
	if x != nil {
		return x.Filter
	}
	return nil
}

type SearchLaptopResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Laptop *Laptop `protobuf:"bytes,1,opt,name=laptop,proto3" json:"laptop,omitempty"`
}

func (x *SearchLaptopResponse) Reset() {
	*x = SearchLaptopResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_laptop_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchLaptopResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchLaptopResponse) ProtoMessage() {}

func (x *SearchLaptopResponse) ProtoReflect() protoreflect.Message {
	mi := &file_laptop_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchLaptopResponse.ProtoReflect.Descriptor instead.
func (*SearchLaptopResponse) Descriptor() ([]byte, []int) {
	return file_laptop_service_proto_rawDescGZIP(), []int{3}
}

func (x *SearchLaptopResponse) GetLaptop() *Laptop {
	if x != nil {
		return x.Laptop
	}
	return nil
}

// 将Image切分为一块块数据，每一个Request当中包含一块数据
type UploadImageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Data:
	//	*UploadImageRequest_Info
	//	*UploadImageRequest_ChunkData
	Data isUploadImageRequest_Data `protobuf_oneof:"data"`
}

func (x *UploadImageRequest) Reset() {
	*x = UploadImageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_laptop_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadImageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadImageRequest) ProtoMessage() {}

func (x *UploadImageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_laptop_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadImageRequest.ProtoReflect.Descriptor instead.
func (*UploadImageRequest) Descriptor() ([]byte, []int) {
	return file_laptop_service_proto_rawDescGZIP(), []int{4}
}

func (m *UploadImageRequest) GetData() isUploadImageRequest_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (x *UploadImageRequest) GetInfo() *ImageInfo {
	if x, ok := x.GetData().(*UploadImageRequest_Info); ok {
		return x.Info
	}
	return nil
}

func (x *UploadImageRequest) GetChunkData() []byte {
	if x, ok := x.GetData().(*UploadImageRequest_ChunkData); ok {
		return x.ChunkData
	}
	return nil
}

type isUploadImageRequest_Data interface {
	isUploadImageRequest_Data()
}

type UploadImageRequest_Info struct {
	Info *ImageInfo `protobuf:"bytes,1,opt,name=info,proto3,oneof"` // First request will only contain the metadata
}

type UploadImageRequest_ChunkData struct {
	ChunkData []byte `protobuf:"bytes,2,opt,name=chunk_data,json=chunkData,proto3,oneof"` // And the following request will contain the image data chunks
}

func (*UploadImageRequest_Info) isUploadImageRequest_Data() {}

func (*UploadImageRequest_ChunkData) isUploadImageRequest_Data() {}

type ImageInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LaptopId  string `protobuf:"bytes,1,opt,name=laptop_id,json=laptopId,proto3" json:"laptop_id,omitempty"`
	ImageType string `protobuf:"bytes,2,opt,name=image_type,json=imageType,proto3" json:"image_type,omitempty"`
}

func (x *ImageInfo) Reset() {
	*x = ImageInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_laptop_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImageInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImageInfo) ProtoMessage() {}

func (x *ImageInfo) ProtoReflect() protoreflect.Message {
	mi := &file_laptop_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImageInfo.ProtoReflect.Descriptor instead.
func (*ImageInfo) Descriptor() ([]byte, []int) {
	return file_laptop_service_proto_rawDescGZIP(), []int{5}
}

func (x *ImageInfo) GetLaptopId() string {
	if x != nil {
		return x.LaptopId
	}
	return ""
}

func (x *ImageInfo) GetImageType() string {
	if x != nil {
		return x.ImageType
	}
	return ""
}

type UploadImageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Size uint32 `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"` // the total size of the uploaded image in bytes
}

func (x *UploadImageResponse) Reset() {
	*x = UploadImageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_laptop_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadImageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadImageResponse) ProtoMessage() {}

func (x *UploadImageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_laptop_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadImageResponse.ProtoReflect.Descriptor instead.
func (*UploadImageResponse) Descriptor() ([]byte, []int) {
	return file_laptop_service_proto_rawDescGZIP(), []int{6}
}

func (x *UploadImageResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UploadImageResponse) GetSize() uint32 {
	if x != nil {
		return x.Size
	}
	return 0
}

var File_laptop_service_proto protoreflect.FileDescriptor

var file_laptop_service_proto_rawDesc = []byte{
	0x0a, 0x14, 0x6c, 0x61, 0x70, 0x74, 0x6f, 0x70, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x74, 0x65, 0x63, 0x68, 0x73, 0x63, 0x68, 0x6f,
	0x6f, 0x6c, 0x2e, 0x70, 0x63, 0x62, 0x6f, 0x6f, 0x6b, 0x1a, 0x14, 0x6c, 0x61, 0x70, 0x74, 0x6f,
	0x70, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x14, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x48, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c,
	0x61, 0x70, 0x74, 0x6f, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x31, 0x0a, 0x06,
	0x6c, 0x61, 0x70, 0x74, 0x6f, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x74,
	0x65, 0x63, 0x68, 0x73, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x2e, 0x70, 0x63, 0x62, 0x6f, 0x6f, 0x6b,
	0x2e, 0x4c, 0x61, 0x70, 0x74, 0x6f, 0x70, 0x52, 0x06, 0x6c, 0x61, 0x70, 0x74, 0x6f, 0x70, 0x22,
	0x26, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x61, 0x70, 0x74, 0x6f, 0x70, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x48, 0x0a, 0x13, 0x53, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x4c, 0x61, 0x70, 0x74, 0x6f, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x31,
	0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x74, 0x65, 0x63, 0x68, 0x73, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x2e, 0x70, 0x63, 0x62, 0x6f,
	0x6f, 0x6b, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x22, 0x49, 0x0a, 0x14, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4c, 0x61, 0x70, 0x74, 0x6f,
	0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x06, 0x6c, 0x61, 0x70,
	0x74, 0x6f, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x74, 0x65, 0x63, 0x68,
	0x73, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x2e, 0x70, 0x63, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x4c, 0x61,
	0x70, 0x74, 0x6f, 0x70, 0x52, 0x06, 0x6c, 0x61, 0x70, 0x74, 0x6f, 0x70, 0x22, 0x71, 0x0a, 0x12,
	0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x32, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x74, 0x65, 0x63, 0x68, 0x73, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x2e, 0x70, 0x63,
	0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x48, 0x00,
	0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x1f, 0x0a, 0x0a, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x5f,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x09, 0x63, 0x68,
	0x75, 0x6e, 0x6b, 0x44, 0x61, 0x74, 0x61, 0x42, 0x06, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22,
	0x47, 0x0a, 0x09, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1b, 0x0a, 0x09,
	0x6c, 0x61, 0x70, 0x74, 0x6f, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x6c, 0x61, 0x70, 0x74, 0x6f, 0x70, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69,
	0x6d, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x22, 0x39, 0x0a, 0x13, 0x55, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x73,
	0x69, 0x7a, 0x65, 0x32, 0xb9, 0x02, 0x0a, 0x0d, 0x4c, 0x61, 0x70, 0x74, 0x6f, 0x70, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x61, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c,
	0x61, 0x70, 0x74, 0x6f, 0x70, 0x12, 0x26, 0x2e, 0x74, 0x65, 0x63, 0x68, 0x73, 0x63, 0x68, 0x6f,
	0x6f, 0x6c, 0x2e, 0x70, 0x63, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x4c, 0x61, 0x70, 0x74, 0x6f, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e,
	0x74, 0x65, 0x63, 0x68, 0x73, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x2e, 0x70, 0x63, 0x62, 0x6f, 0x6f,
	0x6b, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x61, 0x70, 0x74, 0x6f, 0x70, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x63, 0x0a, 0x0c, 0x53, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x4c, 0x61, 0x70, 0x74, 0x6f, 0x70, 0x12, 0x26, 0x2e, 0x74, 0x65, 0x63, 0x68, 0x73,
	0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x2e, 0x70, 0x63, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x4c, 0x61, 0x70, 0x74, 0x6f, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x27, 0x2e, 0x74, 0x65, 0x63, 0x68, 0x73, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x2e, 0x70, 0x63,
	0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4c, 0x61, 0x70, 0x74, 0x6f,
	0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x60, 0x0a,
	0x0b, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x25, 0x2e, 0x74,
	0x65, 0x63, 0x68, 0x73, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x2e, 0x70, 0x63, 0x62, 0x6f, 0x6f, 0x6b,
	0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x74, 0x65, 0x63, 0x68, 0x73, 0x63, 0x68, 0x6f, 0x6f, 0x6c,
	0x2e, 0x70, 0x63, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x49, 0x6d,
	0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x42,
	0x29, 0x0a, 0x1f, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x74, 0x65,
	0x63, 0x68, 0x73, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x2e, 0x70, 0x63, 0x62, 0x6f, 0x6f, 0x6b, 0x2e,
	0x70, 0x62, 0x50, 0x01, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_laptop_service_proto_rawDescOnce sync.Once
	file_laptop_service_proto_rawDescData = file_laptop_service_proto_rawDesc
)

func file_laptop_service_proto_rawDescGZIP() []byte {
	file_laptop_service_proto_rawDescOnce.Do(func() {
		file_laptop_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_laptop_service_proto_rawDescData)
	})
	return file_laptop_service_proto_rawDescData
}

var file_laptop_service_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_laptop_service_proto_goTypes = []interface{}{
	(*CreateLaptopRequest)(nil),  // 0: techschool.pcbook.CreateLaptopRequest
	(*CreateLaptopResponse)(nil), // 1: techschool.pcbook.CreateLaptopResponse
	(*SearchLaptopRequest)(nil),  // 2: techschool.pcbook.SearchLaptopRequest
	(*SearchLaptopResponse)(nil), // 3: techschool.pcbook.SearchLaptopResponse
	(*UploadImageRequest)(nil),   // 4: techschool.pcbook.UploadImageRequest
	(*ImageInfo)(nil),            // 5: techschool.pcbook.ImageInfo
	(*UploadImageResponse)(nil),  // 6: techschool.pcbook.UploadImageResponse
	(*Laptop)(nil),               // 7: techschool.pcbook.Laptop
	(*Filter)(nil),               // 8: techschool.pcbook.Filter
}
var file_laptop_service_proto_depIdxs = []int32{
	7, // 0: techschool.pcbook.CreateLaptopRequest.laptop:type_name -> techschool.pcbook.Laptop
	8, // 1: techschool.pcbook.SearchLaptopRequest.filter:type_name -> techschool.pcbook.Filter
	7, // 2: techschool.pcbook.SearchLaptopResponse.laptop:type_name -> techschool.pcbook.Laptop
	5, // 3: techschool.pcbook.UploadImageRequest.info:type_name -> techschool.pcbook.ImageInfo
	0, // 4: techschool.pcbook.LaptopService.CreateLaptop:input_type -> techschool.pcbook.CreateLaptopRequest
	2, // 5: techschool.pcbook.LaptopService.SearchLaptop:input_type -> techschool.pcbook.SearchLaptopRequest
	4, // 6: techschool.pcbook.LaptopService.UploadImage:input_type -> techschool.pcbook.UploadImageRequest
	1, // 7: techschool.pcbook.LaptopService.CreateLaptop:output_type -> techschool.pcbook.CreateLaptopResponse
	3, // 8: techschool.pcbook.LaptopService.SearchLaptop:output_type -> techschool.pcbook.SearchLaptopResponse
	6, // 9: techschool.pcbook.LaptopService.UploadImage:output_type -> techschool.pcbook.UploadImageResponse
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_laptop_service_proto_init() }
func file_laptop_service_proto_init() {
	if File_laptop_service_proto != nil {
		return
	}
	file_laptop_message_proto_init()
	file_filter_message_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_laptop_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateLaptopRequest); i {
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
		file_laptop_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateLaptopResponse); i {
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
		file_laptop_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchLaptopRequest); i {
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
		file_laptop_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchLaptopResponse); i {
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
		file_laptop_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadImageRequest); i {
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
		file_laptop_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImageInfo); i {
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
		file_laptop_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadImageResponse); i {
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
	file_laptop_service_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*UploadImageRequest_Info)(nil),
		(*UploadImageRequest_ChunkData)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_laptop_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_laptop_service_proto_goTypes,
		DependencyIndexes: file_laptop_service_proto_depIdxs,
		MessageInfos:      file_laptop_service_proto_msgTypes,
	}.Build()
	File_laptop_service_proto = out.File
	file_laptop_service_proto_rawDesc = nil
	file_laptop_service_proto_goTypes = nil
	file_laptop_service_proto_depIdxs = nil
}
