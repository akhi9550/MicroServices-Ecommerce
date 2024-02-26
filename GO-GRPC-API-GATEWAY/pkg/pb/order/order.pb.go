// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.3
// source: pkg/pb/order/order.proto

package order

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

type OrderItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AddressID int64 `protobuf:"varint,1,opt,name=AddressID,proto3" json:"AddressID,omitempty"`
	PaymentID int64 `protobuf:"varint,2,opt,name=PaymentID,proto3" json:"PaymentID,omitempty"`
}

func (x *OrderItem) Reset() {
	*x = OrderItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_order_order_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderItem) ProtoMessage() {}

func (x *OrderItem) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_order_order_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderItem.ProtoReflect.Descriptor instead.
func (*OrderItem) Descriptor() ([]byte, []int) {
	return file_pkg_pb_order_order_proto_rawDescGZIP(), []int{0}
}

func (x *OrderItem) GetAddressID() int64 {
	if x != nil {
		return x.AddressID
	}
	return 0
}

func (x *OrderItem) GetPaymentID() int64 {
	if x != nil {
		return x.PaymentID
	}
	return 0
}

type OrderItemsFromCartRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderFromCart *OrderItem `protobuf:"bytes,1,opt,name=OrderFromCart,proto3" json:"OrderFromCart,omitempty"`
	UserID        int64      `protobuf:"varint,2,opt,name=UserID,proto3" json:"UserID,omitempty"`
}

func (x *OrderItemsFromCartRequest) Reset() {
	*x = OrderItemsFromCartRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_order_order_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderItemsFromCartRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderItemsFromCartRequest) ProtoMessage() {}

func (x *OrderItemsFromCartRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_order_order_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderItemsFromCartRequest.ProtoReflect.Descriptor instead.
func (*OrderItemsFromCartRequest) Descriptor() ([]byte, []int) {
	return file_pkg_pb_order_order_proto_rawDescGZIP(), []int{1}
}

func (x *OrderItemsFromCartRequest) GetOrderFromCart() *OrderItem {
	if x != nil {
		return x.OrderFromCart
	}
	return nil
}

func (x *OrderItemsFromCartRequest) GetUserID() int64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

type OrderItemsFromCartResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderID        int64  `protobuf:"varint,1,opt,name=OrderID,proto3" json:"OrderID,omitempty"`
	Shipmentstatus string `protobuf:"bytes,2,opt,name=Shipmentstatus,proto3" json:"Shipmentstatus,omitempty"`
	Error          string `protobuf:"bytes,3,opt,name=Error,proto3" json:"Error,omitempty"`
}

func (x *OrderItemsFromCartResponse) Reset() {
	*x = OrderItemsFromCartResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_order_order_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderItemsFromCartResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderItemsFromCartResponse) ProtoMessage() {}

func (x *OrderItemsFromCartResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_order_order_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderItemsFromCartResponse.ProtoReflect.Descriptor instead.
func (*OrderItemsFromCartResponse) Descriptor() ([]byte, []int) {
	return file_pkg_pb_order_order_proto_rawDescGZIP(), []int{2}
}

func (x *OrderItemsFromCartResponse) GetOrderID() int64 {
	if x != nil {
		return x.OrderID
	}
	return 0
}

func (x *OrderItemsFromCartResponse) GetShipmentstatus() string {
	if x != nil {
		return x.Shipmentstatus
	}
	return ""
}

func (x *OrderItemsFromCartResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type GetOrderDetailsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID int64 `protobuf:"varint,1,opt,name=UserID,proto3" json:"UserID,omitempty"`
	Page   int64 `protobuf:"varint,2,opt,name=Page,proto3" json:"Page,omitempty"`
	Count  int64 `protobuf:"varint,3,opt,name=Count,proto3" json:"Count,omitempty"`
}

func (x *GetOrderDetailsRequest) Reset() {
	*x = GetOrderDetailsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_order_order_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOrderDetailsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrderDetailsRequest) ProtoMessage() {}

func (x *GetOrderDetailsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_order_order_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOrderDetailsRequest.ProtoReflect.Descriptor instead.
func (*GetOrderDetailsRequest) Descriptor() ([]byte, []int) {
	return file_pkg_pb_order_order_proto_rawDescGZIP(), []int{3}
}

func (x *GetOrderDetailsRequest) GetUserID() int64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *GetOrderDetailsRequest) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetOrderDetailsRequest) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

type OrderDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderID        int64   `protobuf:"varint,1,opt,name=OrderID,proto3" json:"OrderID,omitempty"`
	Price          float32 `protobuf:"fixed32,2,opt,name=Price,proto3" json:"Price,omitempty"`
	Shipmentstatus string  `protobuf:"bytes,3,opt,name=Shipmentstatus,proto3" json:"Shipmentstatus,omitempty"`
	Paymentstatus  string  `protobuf:"bytes,4,opt,name=Paymentstatus,proto3" json:"Paymentstatus,omitempty"`
}

func (x *OrderDetails) Reset() {
	*x = OrderDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_order_order_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderDetails) ProtoMessage() {}

func (x *OrderDetails) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_order_order_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderDetails.ProtoReflect.Descriptor instead.
func (*OrderDetails) Descriptor() ([]byte, []int) {
	return file_pkg_pb_order_order_proto_rawDescGZIP(), []int{4}
}

func (x *OrderDetails) GetOrderID() int64 {
	if x != nil {
		return x.OrderID
	}
	return 0
}

func (x *OrderDetails) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *OrderDetails) GetShipmentstatus() string {
	if x != nil {
		return x.Shipmentstatus
	}
	return ""
}

func (x *OrderDetails) GetPaymentstatus() string {
	if x != nil {
		return x.Paymentstatus
	}
	return ""
}

type OrderProductDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductID   int64   `protobuf:"varint,1,opt,name=ProductID,proto3" json:"ProductID,omitempty"`
	ProductName string  `protobuf:"bytes,2,opt,name=ProductName,proto3" json:"ProductName,omitempty"`
	Quantity    int64   `protobuf:"varint,3,opt,name=Quantity,proto3" json:"Quantity,omitempty"`
	Price       float32 `protobuf:"fixed32,4,opt,name=Price,proto3" json:"Price,omitempty"`
}

func (x *OrderProductDetails) Reset() {
	*x = OrderProductDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_order_order_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderProductDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderProductDetails) ProtoMessage() {}

func (x *OrderProductDetails) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_order_order_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderProductDetails.ProtoReflect.Descriptor instead.
func (*OrderProductDetails) Descriptor() ([]byte, []int) {
	return file_pkg_pb_order_order_proto_rawDescGZIP(), []int{5}
}

func (x *OrderProductDetails) GetProductID() int64 {
	if x != nil {
		return x.ProductID
	}
	return 0
}

func (x *OrderProductDetails) GetProductName() string {
	if x != nil {
		return x.ProductName
	}
	return ""
}

func (x *OrderProductDetails) GetQuantity() int64 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *OrderProductDetails) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

type FullOrderDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Orderdetails        *OrderDetails          `protobuf:"bytes,1,opt,name=orderdetails,proto3" json:"orderdetails,omitempty"`
	OrderProductDetails []*OrderProductDetails `protobuf:"bytes,2,rep,name=OrderProductDetails,proto3" json:"OrderProductDetails,omitempty"`
}

func (x *FullOrderDetails) Reset() {
	*x = FullOrderDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_order_order_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FullOrderDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FullOrderDetails) ProtoMessage() {}

func (x *FullOrderDetails) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_order_order_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FullOrderDetails.ProtoReflect.Descriptor instead.
func (*FullOrderDetails) Descriptor() ([]byte, []int) {
	return file_pkg_pb_order_order_proto_rawDescGZIP(), []int{6}
}

func (x *FullOrderDetails) GetOrderdetails() *OrderDetails {
	if x != nil {
		return x.Orderdetails
	}
	return nil
}

func (x *FullOrderDetails) GetOrderProductDetails() []*OrderProductDetails {
	if x != nil {
		return x.OrderProductDetails
	}
	return nil
}

type GetOrderDetailsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Details []*FullOrderDetails `protobuf:"bytes,1,rep,name=Details,proto3" json:"Details,omitempty"`
	Error   string              `protobuf:"bytes,2,opt,name=Error,proto3" json:"Error,omitempty"`
}

func (x *GetOrderDetailsResponse) Reset() {
	*x = GetOrderDetailsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_order_order_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOrderDetailsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrderDetailsResponse) ProtoMessage() {}

func (x *GetOrderDetailsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_order_order_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOrderDetailsResponse.ProtoReflect.Descriptor instead.
func (*GetOrderDetailsResponse) Descriptor() ([]byte, []int) {
	return file_pkg_pb_order_order_proto_rawDescGZIP(), []int{7}
}

func (x *GetOrderDetailsResponse) GetDetails() []*FullOrderDetails {
	if x != nil {
		return x.Details
	}
	return nil
}

func (x *GetOrderDetailsResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

var File_pkg_pb_order_order_proto protoreflect.FileDescriptor

var file_pkg_pb_order_order_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2f, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x22, 0x47, 0x0a, 0x09, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x1c,
	0x0a, 0x09, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x09, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09,
	0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x22, 0x6b, 0x0a, 0x19, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x46, 0x72, 0x6f, 0x6d, 0x43, 0x61, 0x72, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x36, 0x0a, 0x0d, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x46, 0x72, 0x6f, 0x6d, 0x43, 0x61, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10,
	0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d,
	0x52, 0x0d, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x46, 0x72, 0x6f, 0x6d, 0x43, 0x61, 0x72, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x22, 0x74, 0x0a, 0x1a, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x49, 0x74, 0x65, 0x6d, 0x73, 0x46, 0x72, 0x6f, 0x6d, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x44, 0x12,
	0x26, 0x0a, 0x0e, 0x53, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x53, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x5a, 0x0a,
	0x16, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12,
	0x12, 0x0a, 0x04, 0x50, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x50,
	0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x05, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x8c, 0x01, 0x0a, 0x0c, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x05, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x53, 0x68,
	0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0e, 0x53, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x24, 0x0a, 0x0d, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x50, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x87, 0x01, 0x0a, 0x13, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73,
	0x12, 0x1c, 0x0a, 0x09, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x09, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x44, 0x12, 0x20,
	0x0a, 0x0b, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x08, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x50, 0x72, 0x69,
	0x63, 0x65, 0x22, 0x99, 0x01, 0x0a, 0x10, 0x46, 0x75, 0x6c, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x37, 0x0a, 0x0c, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x73, 0x52, 0x0c, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73,
	0x12, 0x4c, 0x0a, 0x13, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x13, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x22, 0x62,
	0x0a, 0x17, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x07, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x2e, 0x46, 0x75, 0x6c, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x73, 0x52, 0x07, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x14, 0x0a, 0x05,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x32, 0xb8, 0x01, 0x0a, 0x05, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x5b, 0x0a, 0x12,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x46, 0x72, 0x6f, 0x6d, 0x43, 0x61,
	0x72, 0x74, 0x12, 0x20, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x49, 0x74, 0x65, 0x6d, 0x73, 0x46, 0x72, 0x6f, 0x6d, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x46, 0x72, 0x6f, 0x6d, 0x43, 0x61, 0x72, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x52, 0x0a, 0x0f, 0x47, 0x65, 0x74,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x1d, 0x2e, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x10, 0x5a,
	0x0e, 0x2e, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_pb_order_order_proto_rawDescOnce sync.Once
	file_pkg_pb_order_order_proto_rawDescData = file_pkg_pb_order_order_proto_rawDesc
)

func file_pkg_pb_order_order_proto_rawDescGZIP() []byte {
	file_pkg_pb_order_order_proto_rawDescOnce.Do(func() {
		file_pkg_pb_order_order_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_pb_order_order_proto_rawDescData)
	})
	return file_pkg_pb_order_order_proto_rawDescData
}

var file_pkg_pb_order_order_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_pkg_pb_order_order_proto_goTypes = []interface{}{
	(*OrderItem)(nil),                  // 0: order.OrderItem
	(*OrderItemsFromCartRequest)(nil),  // 1: order.OrderItemsFromCartRequest
	(*OrderItemsFromCartResponse)(nil), // 2: order.OrderItemsFromCartResponse
	(*GetOrderDetailsRequest)(nil),     // 3: order.GetOrderDetailsRequest
	(*OrderDetails)(nil),               // 4: order.OrderDetails
	(*OrderProductDetails)(nil),        // 5: order.OrderProductDetails
	(*FullOrderDetails)(nil),           // 6: order.FullOrderDetails
	(*GetOrderDetailsResponse)(nil),    // 7: order.GetOrderDetailsResponse
}
var file_pkg_pb_order_order_proto_depIdxs = []int32{
	0, // 0: order.OrderItemsFromCartRequest.OrderFromCart:type_name -> order.OrderItem
	4, // 1: order.FullOrderDetails.orderdetails:type_name -> order.OrderDetails
	5, // 2: order.FullOrderDetails.OrderProductDetails:type_name -> order.OrderProductDetails
	6, // 3: order.GetOrderDetailsResponse.Details:type_name -> order.FullOrderDetails
	1, // 4: order.Order.OrderItemsFromCart:input_type -> order.OrderItemsFromCartRequest
	3, // 5: order.Order.GetOrderDetails:input_type -> order.GetOrderDetailsRequest
	2, // 6: order.Order.OrderItemsFromCart:output_type -> order.OrderItemsFromCartResponse
	7, // 7: order.Order.GetOrderDetails:output_type -> order.GetOrderDetailsResponse
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_pkg_pb_order_order_proto_init() }
func file_pkg_pb_order_order_proto_init() {
	if File_pkg_pb_order_order_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_pb_order_order_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderItem); i {
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
		file_pkg_pb_order_order_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderItemsFromCartRequest); i {
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
		file_pkg_pb_order_order_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderItemsFromCartResponse); i {
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
		file_pkg_pb_order_order_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOrderDetailsRequest); i {
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
		file_pkg_pb_order_order_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderDetails); i {
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
		file_pkg_pb_order_order_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderProductDetails); i {
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
		file_pkg_pb_order_order_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FullOrderDetails); i {
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
		file_pkg_pb_order_order_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOrderDetailsResponse); i {
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
			RawDescriptor: file_pkg_pb_order_order_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_pb_order_order_proto_goTypes,
		DependencyIndexes: file_pkg_pb_order_order_proto_depIdxs,
		MessageInfos:      file_pkg_pb_order_order_proto_msgTypes,
	}.Build()
	File_pkg_pb_order_order_proto = out.File
	file_pkg_pb_order_order_proto_rawDesc = nil
	file_pkg_pb_order_order_proto_goTypes = nil
	file_pkg_pb_order_order_proto_depIdxs = nil
}
