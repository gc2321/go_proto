// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.29.0
// source: proto/basic/user_group.proto

package basic

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

type UserGroup struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GroupId     uint32   `protobuf:"varint,1,opt,name=group_id,proto3" json:"group_id,omitempty"`
	GroupName   string   `protobuf:"bytes,2,opt,name=group_name,proto3" json:"group_name,omitempty"`
	Roles       []string `protobuf:"bytes,3,rep,name=roles,proto3" json:"roles,omitempty"`
	Users       []*User  `protobuf:"bytes,4,rep,name=users,proto3" json:"users,omitempty"`
	Description string   `protobuf:"bytes,16,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *UserGroup) Reset() {
	*x = UserGroup{}
	mi := &file_proto_basic_user_group_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserGroup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserGroup) ProtoMessage() {}

func (x *UserGroup) ProtoReflect() protoreflect.Message {
	mi := &file_proto_basic_user_group_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserGroup.ProtoReflect.Descriptor instead.
func (*UserGroup) Descriptor() ([]byte, []int) {
	return file_proto_basic_user_group_proto_rawDescGZIP(), []int{0}
}

func (x *UserGroup) GetGroupId() uint32 {
	if x != nil {
		return x.GroupId
	}
	return 0
}

func (x *UserGroup) GetGroupName() string {
	if x != nil {
		return x.GroupName
	}
	return ""
}

func (x *UserGroup) GetRoles() []string {
	if x != nil {
		return x.Roles
	}
	return nil
}

func (x *UserGroup) GetUsers() []*User {
	if x != nil {
		return x.Users
	}
	return nil
}

func (x *UserGroup) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

var File_proto_basic_user_group_proto protoreflect.FileDescriptor

var file_proto_basic_user_group_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2f, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2f, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9c, 0x01, 0x0a, 0x09, 0x55, 0x73, 0x65, 0x72, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x12, 0x1a, 0x0a, 0x08, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64,
	0x12, 0x1e, 0x0a, 0x0a, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x12, 0x1b, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x05, 0x75, 0x73,
	0x65, 0x72, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x1c, 0x5a, 0x1a, 0x6d, 0x79, 0x2d, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x61,
	0x73, 0x69, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_basic_user_group_proto_rawDescOnce sync.Once
	file_proto_basic_user_group_proto_rawDescData = file_proto_basic_user_group_proto_rawDesc
)

func file_proto_basic_user_group_proto_rawDescGZIP() []byte {
	file_proto_basic_user_group_proto_rawDescOnce.Do(func() {
		file_proto_basic_user_group_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_basic_user_group_proto_rawDescData)
	})
	return file_proto_basic_user_group_proto_rawDescData
}

var file_proto_basic_user_group_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_basic_user_group_proto_goTypes = []any{
	(*UserGroup)(nil), // 0: UserGroup
	(*User)(nil),      // 1: User
}
var file_proto_basic_user_group_proto_depIdxs = []int32{
	1, // 0: UserGroup.users:type_name -> User
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_basic_user_group_proto_init() }
func file_proto_basic_user_group_proto_init() {
	if File_proto_basic_user_group_proto != nil {
		return
	}
	file_proto_basic_user_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_basic_user_group_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_basic_user_group_proto_goTypes,
		DependencyIndexes: file_proto_basic_user_group_proto_depIdxs,
		MessageInfos:      file_proto_basic_user_group_proto_msgTypes,
	}.Build()
	File_proto_basic_user_group_proto = out.File
	file_proto_basic_user_group_proto_rawDesc = nil
	file_proto_basic_user_group_proto_goTypes = nil
	file_proto_basic_user_group_proto_depIdxs = nil
}
