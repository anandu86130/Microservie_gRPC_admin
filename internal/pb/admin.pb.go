// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.3
// source: pb/admin.proto

package __

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

type AdminNoParam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AdminNoParam) Reset() {
	*x = AdminNoParam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_admin_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdminNoParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdminNoParam) ProtoMessage() {}

func (x *AdminNoParam) ProtoReflect() protoreflect.Message {
	mi := &file_pb_admin_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdminNoParam.ProtoReflect.Descriptor instead.
func (*AdminNoParam) Descriptor() ([]byte, []int) {
	return file_pb_admin_proto_rawDescGZIP(), []int{0}
}

type AdminRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Role     string `protobuf:"bytes,3,opt,name=role,proto3" json:"role,omitempty"`
}

func (x *AdminRequest) Reset() {
	*x = AdminRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_admin_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdminRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdminRequest) ProtoMessage() {}

func (x *AdminRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_admin_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdminRequest.ProtoReflect.Descriptor instead.
func (*AdminRequest) Descriptor() ([]byte, []int) {
	return file_pb_admin_proto_rawDescGZIP(), []int{1}
}

func (x *AdminRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *AdminRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *AdminRequest) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

type AdminResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Error   string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Message string `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *AdminResponse) Reset() {
	*x = AdminResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_admin_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdminResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdminResponse) ProtoMessage() {}

func (x *AdminResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pb_admin_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdminResponse.ProtoReflect.Descriptor instead.
func (*AdminResponse) Descriptor() ([]byte, []int) {
	return file_pb_admin_proto_rawDescGZIP(), []int{2}
}

func (x *AdminResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *AdminResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *AdminResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type AProductDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          uint32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Category    string  `protobuf:"bytes,2,opt,name=category,proto3" json:"category,omitempty"`
	Name        string  `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Price       float64 `protobuf:"fixed64,4,opt,name=price,proto3" json:"price,omitempty"`
	Imagepath   string  `protobuf:"bytes,5,opt,name=imagepath,proto3" json:"imagepath,omitempty"`
	Description string  `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty"`
	Size        string  `protobuf:"bytes,7,opt,name=size,proto3" json:"size,omitempty"`
	Quantity    uint32  `protobuf:"varint,8,opt,name=quantity,proto3" json:"quantity,omitempty"`
}

func (x *AProductDetails) Reset() {
	*x = AProductDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_admin_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AProductDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AProductDetails) ProtoMessage() {}

func (x *AProductDetails) ProtoReflect() protoreflect.Message {
	mi := &file_pb_admin_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AProductDetails.ProtoReflect.Descriptor instead.
func (*AProductDetails) Descriptor() ([]byte, []int) {
	return file_pb_admin_proto_rawDescGZIP(), []int{3}
}

func (x *AProductDetails) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *AProductDetails) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *AProductDetails) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AProductDetails) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *AProductDetails) GetImagepath() string {
	if x != nil {
		return x.Imagepath
	}
	return ""
}

func (x *AProductDetails) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *AProductDetails) GetSize() string {
	if x != nil {
		return x.Size
	}
	return ""
}

func (x *AProductDetails) GetQuantity() uint32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

type AProductById struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *AProductById) Reset() {
	*x = AProductById{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_admin_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AProductById) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AProductById) ProtoMessage() {}

func (x *AProductById) ProtoReflect() protoreflect.Message {
	mi := &file_pb_admin_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AProductById.ProtoReflect.Descriptor instead.
func (*AProductById) Descriptor() ([]byte, []int) {
	return file_pb_admin_proto_rawDescGZIP(), []int{4}
}

func (x *AProductById) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type AProductByName struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *AProductByName) Reset() {
	*x = AProductByName{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_admin_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AProductByName) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AProductByName) ProtoMessage() {}

func (x *AProductByName) ProtoReflect() protoreflect.Message {
	mi := &file_pb_admin_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AProductByName.ProtoReflect.Descriptor instead.
func (*AProductByName) Descriptor() ([]byte, []int) {
	return file_pb_admin_proto_rawDescGZIP(), []int{5}
}

func (x *AProductByName) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type AProductList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Products []*AProductDetails `protobuf:"bytes,1,rep,name=products,proto3" json:"products,omitempty"`
}

func (x *AProductList) Reset() {
	*x = AProductList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_admin_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AProductList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AProductList) ProtoMessage() {}

func (x *AProductList) ProtoReflect() protoreflect.Message {
	mi := &file_pb_admin_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AProductList.ProtoReflect.Descriptor instead.
func (*AProductList) Descriptor() ([]byte, []int) {
	return file_pb_admin_proto_rawDescGZIP(), []int{6}
}

func (x *AProductList) GetProducts() []*AProductDetails {
	if x != nil {
		return x.Products
	}
	return nil
}

var File_pb_admin_proto protoreflect.FileDescriptor

var file_pb_admin_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x70, 0x62, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x03, 0x50, 0x62, 0x41, 0x22, 0x0e, 0x0a, 0x0c, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x4e, 0x6f,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x22, 0x5a, 0x0a, 0x0c, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x6f, 0x6c,
	0x65, 0x22, 0x57, 0x0a, 0x0d, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0xd7, 0x01, 0x0a, 0x0f, 0x41,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a,
	0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x70,
	0x72, 0x69, 0x63, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x70, 0x61, 0x74,
	0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x70, 0x61,
	0x74, 0x68, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x22, 0x1e, 0x0a, 0x0c, 0x41, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x42, 0x79, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x24, 0x0a, 0x0e, 0x41, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x40, 0x0a, 0x0c, 0x41, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x30, 0x0a, 0x08, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x50,
	0x62, 0x41, 0x2e, 0x41, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x73, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x32, 0xac, 0x02, 0x0a,
	0x0c, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x33, 0x0a,
	0x0a, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x11, 0x2e, 0x50, 0x62,
	0x41, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12,
	0x2e, 0x50, 0x62, 0x41, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x39, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x12, 0x14, 0x2e, 0x50, 0x62, 0x41, 0x2e, 0x41, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x1a, 0x12, 0x2e, 0x50, 0x62, 0x41, 0x2e,
	0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b, 0x0a,
	0x10, 0x46, 0x65, 0x74, 0x63, 0x68, 0x42, 0x79, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49,
	0x44, 0x12, 0x11, 0x2e, 0x50, 0x62, 0x41, 0x2e, 0x41, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x42, 0x79, 0x49, 0x64, 0x1a, 0x14, 0x2e, 0x50, 0x62, 0x41, 0x2e, 0x41, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x38, 0x0a, 0x0b, 0x46, 0x65,
	0x74, 0x63, 0x68, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x13, 0x2e, 0x50, 0x62, 0x41, 0x2e,
	0x41, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x1a, 0x14,
	0x2e, 0x50, 0x62, 0x41, 0x2e, 0x41, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x73, 0x12, 0x35, 0x0a, 0x0d, 0x46, 0x65, 0x74, 0x63, 0x68, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x73, 0x12, 0x11, 0x2e, 0x50, 0x62, 0x41, 0x2e, 0x41, 0x64, 0x6d, 0x69,
	0x6e, 0x4e, 0x6f, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x1a, 0x11, 0x2e, 0x50, 0x62, 0x41, 0x2e, 0x41,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x04, 0x5a, 0x02, 0x2e,
	0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_admin_proto_rawDescOnce sync.Once
	file_pb_admin_proto_rawDescData = file_pb_admin_proto_rawDesc
)

func file_pb_admin_proto_rawDescGZIP() []byte {
	file_pb_admin_proto_rawDescOnce.Do(func() {
		file_pb_admin_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_admin_proto_rawDescData)
	})
	return file_pb_admin_proto_rawDescData
}

var file_pb_admin_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_pb_admin_proto_goTypes = []any{
	(*AdminNoParam)(nil),    // 0: PbA.AdminNoParam
	(*AdminRequest)(nil),    // 1: PbA.AdminRequest
	(*AdminResponse)(nil),   // 2: PbA.AdminResponse
	(*AProductDetails)(nil), // 3: PbA.AProductDetails
	(*AProductById)(nil),    // 4: PbA.AProductById
	(*AProductByName)(nil),  // 5: PbA.AProductByName
	(*AProductList)(nil),    // 6: PbA.AProductList
}
var file_pb_admin_proto_depIdxs = []int32{
	3, // 0: PbA.AProductList.products:type_name -> PbA.AProductDetails
	1, // 1: PbA.AdminService.AdminLogin:input_type -> PbA.AdminRequest
	3, // 2: PbA.AdminService.CreateProduct:input_type -> PbA.AProductDetails
	4, // 3: PbA.AdminService.FetchByProductID:input_type -> PbA.AProductById
	5, // 4: PbA.AdminService.FetchByName:input_type -> PbA.AProductByName
	0, // 5: PbA.AdminService.FetchProducts:input_type -> PbA.AdminNoParam
	2, // 6: PbA.AdminService.AdminLogin:output_type -> PbA.AdminResponse
	2, // 7: PbA.AdminService.CreateProduct:output_type -> PbA.AdminResponse
	3, // 8: PbA.AdminService.FetchByProductID:output_type -> PbA.AProductDetails
	3, // 9: PbA.AdminService.FetchByName:output_type -> PbA.AProductDetails
	6, // 10: PbA.AdminService.FetchProducts:output_type -> PbA.AProductList
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pb_admin_proto_init() }
func file_pb_admin_proto_init() {
	if File_pb_admin_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_admin_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*AdminNoParam); i {
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
		file_pb_admin_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*AdminRequest); i {
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
		file_pb_admin_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*AdminResponse); i {
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
		file_pb_admin_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*AProductDetails); i {
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
		file_pb_admin_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*AProductById); i {
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
		file_pb_admin_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*AProductByName); i {
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
		file_pb_admin_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*AProductList); i {
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
			RawDescriptor: file_pb_admin_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pb_admin_proto_goTypes,
		DependencyIndexes: file_pb_admin_proto_depIdxs,
		MessageInfos:      file_pb_admin_proto_msgTypes,
	}.Build()
	File_pb_admin_proto = out.File
	file_pb_admin_proto_rawDesc = nil
	file_pb_admin_proto_goTypes = nil
	file_pb_admin_proto_depIdxs = nil
}
