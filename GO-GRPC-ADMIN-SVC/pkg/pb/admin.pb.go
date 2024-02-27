// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.3
// source: pkg/pb/admin.proto

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

type AdminSignupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Firstname string `protobuf:"bytes,1,opt,name=firstname,proto3" json:"firstname,omitempty"`
	Lastname  string `protobuf:"bytes,2,opt,name=lastname,proto3" json:"lastname,omitempty"`
	Email     string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Password  string `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *AdminSignupRequest) Reset() {
	*x = AdminSignupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_admin_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdminSignupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdminSignupRequest) ProtoMessage() {}

func (x *AdminSignupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_admin_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdminSignupRequest.ProtoReflect.Descriptor instead.
func (*AdminSignupRequest) Descriptor() ([]byte, []int) {
	return file_pkg_pb_admin_proto_rawDescGZIP(), []int{0}
}

func (x *AdminSignupRequest) GetFirstname() string {
	if x != nil {
		return x.Firstname
	}
	return ""
}

func (x *AdminSignupRequest) GetLastname() string {
	if x != nil {
		return x.Lastname
	}
	return ""
}

func (x *AdminSignupRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *AdminSignupRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type AdminSignupResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status       int64         `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	AdminDetails *AdminDetails `protobuf:"bytes,2,opt,name=adminDetails,proto3" json:"adminDetails,omitempty"`
	Token        string        `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *AdminSignupResponse) Reset() {
	*x = AdminSignupResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_admin_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdminSignupResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdminSignupResponse) ProtoMessage() {}

func (x *AdminSignupResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_admin_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdminSignupResponse.ProtoReflect.Descriptor instead.
func (*AdminSignupResponse) Descriptor() ([]byte, []int) {
	return file_pkg_pb_admin_proto_rawDescGZIP(), []int{1}
}

func (x *AdminSignupResponse) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *AdminSignupResponse) GetAdminDetails() *AdminDetails {
	if x != nil {
		return x.AdminDetails
	}
	return nil
}

func (x *AdminSignupResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type AdminDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Firstname string `protobuf:"bytes,2,opt,name=firstname,proto3" json:"firstname,omitempty"`
	Lastname  string `protobuf:"bytes,3,opt,name=lastname,proto3" json:"lastname,omitempty"`
	Email     string `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *AdminDetails) Reset() {
	*x = AdminDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_admin_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdminDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdminDetails) ProtoMessage() {}

func (x *AdminDetails) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_admin_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdminDetails.ProtoReflect.Descriptor instead.
func (*AdminDetails) Descriptor() ([]byte, []int) {
	return file_pkg_pb_admin_proto_rawDescGZIP(), []int{2}
}

func (x *AdminDetails) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *AdminDetails) GetFirstname() string {
	if x != nil {
		return x.Firstname
	}
	return ""
}

func (x *AdminDetails) GetLastname() string {
	if x != nil {
		return x.Lastname
	}
	return ""
}

func (x *AdminDetails) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type AdminLoginInRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email    string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *AdminLoginInRequest) Reset() {
	*x = AdminLoginInRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_admin_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdminLoginInRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdminLoginInRequest) ProtoMessage() {}

func (x *AdminLoginInRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_admin_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdminLoginInRequest.ProtoReflect.Descriptor instead.
func (*AdminLoginInRequest) Descriptor() ([]byte, []int) {
	return file_pkg_pb_admin_proto_rawDescGZIP(), []int{3}
}

func (x *AdminLoginInRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *AdminLoginInRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type AdminLoginResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status       int64         `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	AdminDetails *AdminDetails `protobuf:"bytes,2,opt,name=adminDetails,proto3" json:"adminDetails,omitempty"`
	Token        string        `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *AdminLoginResponse) Reset() {
	*x = AdminLoginResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_admin_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdminLoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdminLoginResponse) ProtoMessage() {}

func (x *AdminLoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_admin_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdminLoginResponse.ProtoReflect.Descriptor instead.
func (*AdminLoginResponse) Descriptor() ([]byte, []int) {
	return file_pkg_pb_admin_proto_rawDescGZIP(), []int{4}
}

func (x *AdminLoginResponse) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *AdminLoginResponse) GetAdminDetails() *AdminDetails {
	if x != nil {
		return x.AdminDetails
	}
	return nil
}

func (x *AdminLoginResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

var File_pkg_pb_admin_proto protoreflect.FileDescriptor

var file_pkg_pb_admin_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x22, 0x80, 0x01, 0x0a, 0x12,
	0x41, 0x64, 0x6d, 0x69, 0x6e, 0x53, 0x69, 0x67, 0x6e, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x7c,
	0x0a, 0x13, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x53, 0x69, 0x67, 0x6e, 0x75, 0x70, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x37, 0x0a,
	0x0c, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x41, 0x64, 0x6d, 0x69,
	0x6e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x0c, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x6e, 0x0a, 0x0c,
	0x41, 0x64, 0x6d, 0x69, 0x6e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09,
	0x66, 0x69, 0x72, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61,
	0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61,
	0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x47, 0x0a, 0x13,
	0x41, 0x64, 0x6d, 0x69, 0x6e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x49, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x7b, 0x0a, 0x12, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x37, 0x0a, 0x0c, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x0c,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x32, 0x96, 0x01, 0x0a, 0x05, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x46, 0x0a, 0x0b,
	0x41, 0x64, 0x6d, 0x69, 0x6e, 0x53, 0x69, 0x67, 0x6e, 0x75, 0x70, 0x12, 0x19, 0x2e, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x53, 0x69, 0x67, 0x6e, 0x75, 0x70, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x41,
	0x64, 0x6d, 0x69, 0x6e, 0x53, 0x69, 0x67, 0x6e, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x45, 0x0a, 0x0a, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x12, 0x1a, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x49, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19,
	0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x4c, 0x6f, 0x67, 0x69,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0a, 0x5a, 0x08, 0x2e,
	0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_pb_admin_proto_rawDescOnce sync.Once
	file_pkg_pb_admin_proto_rawDescData = file_pkg_pb_admin_proto_rawDesc
)

func file_pkg_pb_admin_proto_rawDescGZIP() []byte {
	file_pkg_pb_admin_proto_rawDescOnce.Do(func() {
		file_pkg_pb_admin_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_pb_admin_proto_rawDescData)
	})
	return file_pkg_pb_admin_proto_rawDescData
}

var file_pkg_pb_admin_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_pkg_pb_admin_proto_goTypes = []interface{}{
	(*AdminSignupRequest)(nil),  // 0: admin.AdminSignupRequest
	(*AdminSignupResponse)(nil), // 1: admin.AdminSignupResponse
	(*AdminDetails)(nil),        // 2: admin.AdminDetails
	(*AdminLoginInRequest)(nil), // 3: admin.AdminLoginInRequest
	(*AdminLoginResponse)(nil),  // 4: admin.AdminLoginResponse
}
var file_pkg_pb_admin_proto_depIdxs = []int32{
	2, // 0: admin.AdminSignupResponse.adminDetails:type_name -> admin.AdminDetails
	2, // 1: admin.AdminLoginResponse.adminDetails:type_name -> admin.AdminDetails
	0, // 2: admin.Admin.AdminSignup:input_type -> admin.AdminSignupRequest
	3, // 3: admin.Admin.AdminLogin:input_type -> admin.AdminLoginInRequest
	1, // 4: admin.Admin.AdminSignup:output_type -> admin.AdminSignupResponse
	4, // 5: admin.Admin.AdminLogin:output_type -> admin.AdminLoginResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_pkg_pb_admin_proto_init() }
func file_pkg_pb_admin_proto_init() {
	if File_pkg_pb_admin_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_pb_admin_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdminSignupRequest); i {
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
		file_pkg_pb_admin_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdminSignupResponse); i {
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
		file_pkg_pb_admin_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdminDetails); i {
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
		file_pkg_pb_admin_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdminLoginInRequest); i {
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
		file_pkg_pb_admin_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdminLoginResponse); i {
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
			RawDescriptor: file_pkg_pb_admin_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_pb_admin_proto_goTypes,
		DependencyIndexes: file_pkg_pb_admin_proto_depIdxs,
		MessageInfos:      file_pkg_pb_admin_proto_msgTypes,
	}.Build()
	File_pkg_pb_admin_proto = out.File
	file_pkg_pb_admin_proto_rawDesc = nil
	file_pkg_pb_admin_proto_goTypes = nil
	file_pkg_pb_admin_proto_depIdxs = nil
}
