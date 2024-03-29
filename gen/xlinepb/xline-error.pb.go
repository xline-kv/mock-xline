// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: xline-error.proto

package xlineapi

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Command::Error
type ExecuteError struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Error:
	//
	//	*ExecuteError_KeyNotFound
	//	*ExecuteError_RevisionTooLarge
	//	*ExecuteError_RevisionCompacted
	//	*ExecuteError_LeaseNotFound
	//	*ExecuteError_LeaseExpired
	//	*ExecuteError_LeaseTtlTooLarge
	//	*ExecuteError_LeaseAlreadyExists
	//	*ExecuteError_AuthNotEnabled
	//	*ExecuteError_AuthFailed
	//	*ExecuteError_UserNotFound
	//	*ExecuteError_UserAlreadyExists
	//	*ExecuteError_UserAlreadyHasRole
	//	*ExecuteError_NoPasswordUser
	//	*ExecuteError_RoleNotFound
	//	*ExecuteError_RoleAlreadyExists
	//	*ExecuteError_RoleNotGranted
	//	*ExecuteError_RootRoleNotExist
	//	*ExecuteError_PermissionNotGranted
	//	*ExecuteError_PermissionNotGiven
	//	*ExecuteError_InvalidAuthManagement
	//	*ExecuteError_InvalidAuthToken
	//	*ExecuteError_TokenManagerNotInit
	//	*ExecuteError_TokenNotProvided
	//	*ExecuteError_TokenOldRevision
	//	*ExecuteError_DbError
	//	*ExecuteError_PermissionDenied
	Error isExecuteError_Error `protobuf_oneof:"error"`
}

func (x *ExecuteError) Reset() {
	*x = ExecuteError{}
	if protoimpl.UnsafeEnabled {
		mi := &file_xline_error_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExecuteError) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecuteError) ProtoMessage() {}

func (x *ExecuteError) ProtoReflect() protoreflect.Message {
	mi := &file_xline_error_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecuteError.ProtoReflect.Descriptor instead.
func (*ExecuteError) Descriptor() ([]byte, []int) {
	return file_xline_error_proto_rawDescGZIP(), []int{0}
}

func (m *ExecuteError) GetError() isExecuteError_Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (x *ExecuteError) GetKeyNotFound() *emptypb.Empty {
	if x, ok := x.GetError().(*ExecuteError_KeyNotFound); ok {
		return x.KeyNotFound
	}
	return nil
}

func (x *ExecuteError) GetRevisionTooLarge() *Revisions {
	if x, ok := x.GetError().(*ExecuteError_RevisionTooLarge); ok {
		return x.RevisionTooLarge
	}
	return nil
}

func (x *ExecuteError) GetRevisionCompacted() *Revisions {
	if x, ok := x.GetError().(*ExecuteError_RevisionCompacted); ok {
		return x.RevisionCompacted
	}
	return nil
}

func (x *ExecuteError) GetLeaseNotFound() int64 {
	if x, ok := x.GetError().(*ExecuteError_LeaseNotFound); ok {
		return x.LeaseNotFound
	}
	return 0
}

func (x *ExecuteError) GetLeaseExpired() int64 {
	if x, ok := x.GetError().(*ExecuteError_LeaseExpired); ok {
		return x.LeaseExpired
	}
	return 0
}

func (x *ExecuteError) GetLeaseTtlTooLarge() int64 {
	if x, ok := x.GetError().(*ExecuteError_LeaseTtlTooLarge); ok {
		return x.LeaseTtlTooLarge
	}
	return 0
}

func (x *ExecuteError) GetLeaseAlreadyExists() int64 {
	if x, ok := x.GetError().(*ExecuteError_LeaseAlreadyExists); ok {
		return x.LeaseAlreadyExists
	}
	return 0
}

func (x *ExecuteError) GetAuthNotEnabled() *emptypb.Empty {
	if x, ok := x.GetError().(*ExecuteError_AuthNotEnabled); ok {
		return x.AuthNotEnabled
	}
	return nil
}

func (x *ExecuteError) GetAuthFailed() *emptypb.Empty {
	if x, ok := x.GetError().(*ExecuteError_AuthFailed); ok {
		return x.AuthFailed
	}
	return nil
}

func (x *ExecuteError) GetUserNotFound() string {
	if x, ok := x.GetError().(*ExecuteError_UserNotFound); ok {
		return x.UserNotFound
	}
	return ""
}

func (x *ExecuteError) GetUserAlreadyExists() string {
	if x, ok := x.GetError().(*ExecuteError_UserAlreadyExists); ok {
		return x.UserAlreadyExists
	}
	return ""
}

func (x *ExecuteError) GetUserAlreadyHasRole() *UserRole {
	if x, ok := x.GetError().(*ExecuteError_UserAlreadyHasRole); ok {
		return x.UserAlreadyHasRole
	}
	return nil
}

func (x *ExecuteError) GetNoPasswordUser() *emptypb.Empty {
	if x, ok := x.GetError().(*ExecuteError_NoPasswordUser); ok {
		return x.NoPasswordUser
	}
	return nil
}

func (x *ExecuteError) GetRoleNotFound() string {
	if x, ok := x.GetError().(*ExecuteError_RoleNotFound); ok {
		return x.RoleNotFound
	}
	return ""
}

func (x *ExecuteError) GetRoleAlreadyExists() string {
	if x, ok := x.GetError().(*ExecuteError_RoleAlreadyExists); ok {
		return x.RoleAlreadyExists
	}
	return ""
}

func (x *ExecuteError) GetRoleNotGranted() string {
	if x, ok := x.GetError().(*ExecuteError_RoleNotGranted); ok {
		return x.RoleNotGranted
	}
	return ""
}

func (x *ExecuteError) GetRootRoleNotExist() *emptypb.Empty {
	if x, ok := x.GetError().(*ExecuteError_RootRoleNotExist); ok {
		return x.RootRoleNotExist
	}
	return nil
}

func (x *ExecuteError) GetPermissionNotGranted() *emptypb.Empty {
	if x, ok := x.GetError().(*ExecuteError_PermissionNotGranted); ok {
		return x.PermissionNotGranted
	}
	return nil
}

func (x *ExecuteError) GetPermissionNotGiven() *emptypb.Empty {
	if x, ok := x.GetError().(*ExecuteError_PermissionNotGiven); ok {
		return x.PermissionNotGiven
	}
	return nil
}

func (x *ExecuteError) GetInvalidAuthManagement() *emptypb.Empty {
	if x, ok := x.GetError().(*ExecuteError_InvalidAuthManagement); ok {
		return x.InvalidAuthManagement
	}
	return nil
}

func (x *ExecuteError) GetInvalidAuthToken() *emptypb.Empty {
	if x, ok := x.GetError().(*ExecuteError_InvalidAuthToken); ok {
		return x.InvalidAuthToken
	}
	return nil
}

func (x *ExecuteError) GetTokenManagerNotInit() *emptypb.Empty {
	if x, ok := x.GetError().(*ExecuteError_TokenManagerNotInit); ok {
		return x.TokenManagerNotInit
	}
	return nil
}

func (x *ExecuteError) GetTokenNotProvided() *emptypb.Empty {
	if x, ok := x.GetError().(*ExecuteError_TokenNotProvided); ok {
		return x.TokenNotProvided
	}
	return nil
}

func (x *ExecuteError) GetTokenOldRevision() *Revisions {
	if x, ok := x.GetError().(*ExecuteError_TokenOldRevision); ok {
		return x.TokenOldRevision
	}
	return nil
}

func (x *ExecuteError) GetDbError() string {
	if x, ok := x.GetError().(*ExecuteError_DbError); ok {
		return x.DbError
	}
	return ""
}

func (x *ExecuteError) GetPermissionDenied() *emptypb.Empty {
	if x, ok := x.GetError().(*ExecuteError_PermissionDenied); ok {
		return x.PermissionDenied
	}
	return nil
}

type isExecuteError_Error interface {
	isExecuteError_Error()
}

type ExecuteError_KeyNotFound struct {
	KeyNotFound *emptypb.Empty `protobuf:"bytes,1,opt,name=key_not_found,json=keyNotFound,proto3,oneof"`
}

type ExecuteError_RevisionTooLarge struct {
	RevisionTooLarge *Revisions `protobuf:"bytes,2,opt,name=revision_too_large,json=revisionTooLarge,proto3,oneof"`
}

type ExecuteError_RevisionCompacted struct {
	RevisionCompacted *Revisions `protobuf:"bytes,3,opt,name=revision_compacted,json=revisionCompacted,proto3,oneof"`
}

type ExecuteError_LeaseNotFound struct {
	LeaseNotFound int64 `protobuf:"varint,4,opt,name=lease_not_found,json=leaseNotFound,proto3,oneof"`
}

type ExecuteError_LeaseExpired struct {
	LeaseExpired int64 `protobuf:"varint,5,opt,name=lease_expired,json=leaseExpired,proto3,oneof"`
}

type ExecuteError_LeaseTtlTooLarge struct {
	LeaseTtlTooLarge int64 `protobuf:"varint,6,opt,name=lease_ttl_too_large,json=leaseTtlTooLarge,proto3,oneof"`
}

type ExecuteError_LeaseAlreadyExists struct {
	LeaseAlreadyExists int64 `protobuf:"varint,7,opt,name=lease_already_exists,json=leaseAlreadyExists,proto3,oneof"`
}

type ExecuteError_AuthNotEnabled struct {
	AuthNotEnabled *emptypb.Empty `protobuf:"bytes,8,opt,name=auth_not_enabled,json=authNotEnabled,proto3,oneof"`
}

type ExecuteError_AuthFailed struct {
	AuthFailed *emptypb.Empty `protobuf:"bytes,9,opt,name=auth_failed,json=authFailed,proto3,oneof"`
}

type ExecuteError_UserNotFound struct {
	UserNotFound string `protobuf:"bytes,10,opt,name=user_not_found,json=userNotFound,proto3,oneof"`
}

type ExecuteError_UserAlreadyExists struct {
	UserAlreadyExists string `protobuf:"bytes,11,opt,name=user_already_exists,json=userAlreadyExists,proto3,oneof"`
}

type ExecuteError_UserAlreadyHasRole struct {
	UserAlreadyHasRole *UserRole `protobuf:"bytes,12,opt,name=user_already_has_role,json=userAlreadyHasRole,proto3,oneof"`
}

type ExecuteError_NoPasswordUser struct {
	NoPasswordUser *emptypb.Empty `protobuf:"bytes,13,opt,name=no_password_user,json=noPasswordUser,proto3,oneof"`
}

type ExecuteError_RoleNotFound struct {
	RoleNotFound string `protobuf:"bytes,14,opt,name=role_not_found,json=roleNotFound,proto3,oneof"`
}

type ExecuteError_RoleAlreadyExists struct {
	RoleAlreadyExists string `protobuf:"bytes,15,opt,name=role_already_exists,json=roleAlreadyExists,proto3,oneof"`
}

type ExecuteError_RoleNotGranted struct {
	RoleNotGranted string `protobuf:"bytes,16,opt,name=role_not_granted,json=roleNotGranted,proto3,oneof"`
}

type ExecuteError_RootRoleNotExist struct {
	RootRoleNotExist *emptypb.Empty `protobuf:"bytes,17,opt,name=root_role_not_exist,json=rootRoleNotExist,proto3,oneof"`
}

type ExecuteError_PermissionNotGranted struct {
	PermissionNotGranted *emptypb.Empty `protobuf:"bytes,18,opt,name=permission_not_granted,json=permissionNotGranted,proto3,oneof"`
}

type ExecuteError_PermissionNotGiven struct {
	PermissionNotGiven *emptypb.Empty `protobuf:"bytes,19,opt,name=permission_not_given,json=permissionNotGiven,proto3,oneof"`
}

type ExecuteError_InvalidAuthManagement struct {
	InvalidAuthManagement *emptypb.Empty `protobuf:"bytes,20,opt,name=invalid_auth_management,json=invalidAuthManagement,proto3,oneof"`
}

type ExecuteError_InvalidAuthToken struct {
	InvalidAuthToken *emptypb.Empty `protobuf:"bytes,21,opt,name=invalid_auth_token,json=invalidAuthToken,proto3,oneof"`
}

type ExecuteError_TokenManagerNotInit struct {
	TokenManagerNotInit *emptypb.Empty `protobuf:"bytes,22,opt,name=token_manager_not_init,json=tokenManagerNotInit,proto3,oneof"`
}

type ExecuteError_TokenNotProvided struct {
	TokenNotProvided *emptypb.Empty `protobuf:"bytes,23,opt,name=token_not_provided,json=tokenNotProvided,proto3,oneof"`
}

type ExecuteError_TokenOldRevision struct {
	TokenOldRevision *Revisions `protobuf:"bytes,24,opt,name=token_old_revision,json=tokenOldRevision,proto3,oneof"`
}

type ExecuteError_DbError struct {
	DbError string `protobuf:"bytes,25,opt,name=db_error,json=dbError,proto3,oneof"`
}

type ExecuteError_PermissionDenied struct {
	PermissionDenied *emptypb.Empty `protobuf:"bytes,26,opt,name=permission_denied,json=permissionDenied,proto3,oneof"`
}

func (*ExecuteError_KeyNotFound) isExecuteError_Error() {}

func (*ExecuteError_RevisionTooLarge) isExecuteError_Error() {}

func (*ExecuteError_RevisionCompacted) isExecuteError_Error() {}

func (*ExecuteError_LeaseNotFound) isExecuteError_Error() {}

func (*ExecuteError_LeaseExpired) isExecuteError_Error() {}

func (*ExecuteError_LeaseTtlTooLarge) isExecuteError_Error() {}

func (*ExecuteError_LeaseAlreadyExists) isExecuteError_Error() {}

func (*ExecuteError_AuthNotEnabled) isExecuteError_Error() {}

func (*ExecuteError_AuthFailed) isExecuteError_Error() {}

func (*ExecuteError_UserNotFound) isExecuteError_Error() {}

func (*ExecuteError_UserAlreadyExists) isExecuteError_Error() {}

func (*ExecuteError_UserAlreadyHasRole) isExecuteError_Error() {}

func (*ExecuteError_NoPasswordUser) isExecuteError_Error() {}

func (*ExecuteError_RoleNotFound) isExecuteError_Error() {}

func (*ExecuteError_RoleAlreadyExists) isExecuteError_Error() {}

func (*ExecuteError_RoleNotGranted) isExecuteError_Error() {}

func (*ExecuteError_RootRoleNotExist) isExecuteError_Error() {}

func (*ExecuteError_PermissionNotGranted) isExecuteError_Error() {}

func (*ExecuteError_PermissionNotGiven) isExecuteError_Error() {}

func (*ExecuteError_InvalidAuthManagement) isExecuteError_Error() {}

func (*ExecuteError_InvalidAuthToken) isExecuteError_Error() {}

func (*ExecuteError_TokenManagerNotInit) isExecuteError_Error() {}

func (*ExecuteError_TokenNotProvided) isExecuteError_Error() {}

func (*ExecuteError_TokenOldRevision) isExecuteError_Error() {}

func (*ExecuteError_DbError) isExecuteError_Error() {}

func (*ExecuteError_PermissionDenied) isExecuteError_Error() {}

type Revisions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequiredRevision int64 `protobuf:"varint,1,opt,name=required_revision,json=requiredRevision,proto3" json:"required_revision,omitempty"`
	CurrentRevision  int64 `protobuf:"varint,2,opt,name=current_revision,json=currentRevision,proto3" json:"current_revision,omitempty"`
}

func (x *Revisions) Reset() {
	*x = Revisions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_xline_error_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Revisions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Revisions) ProtoMessage() {}

func (x *Revisions) ProtoReflect() protoreflect.Message {
	mi := &file_xline_error_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Revisions.ProtoReflect.Descriptor instead.
func (*Revisions) Descriptor() ([]byte, []int) {
	return file_xline_error_proto_rawDescGZIP(), []int{1}
}

func (x *Revisions) GetRequiredRevision() int64 {
	if x != nil {
		return x.RequiredRevision
	}
	return 0
}

func (x *Revisions) GetCurrentRevision() int64 {
	if x != nil {
		return x.CurrentRevision
	}
	return 0
}

type UserRole struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User string `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Role string `protobuf:"bytes,2,opt,name=role,proto3" json:"role,omitempty"`
}

func (x *UserRole) Reset() {
	*x = UserRole{}
	if protoimpl.UnsafeEnabled {
		mi := &file_xline_error_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserRole) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRole) ProtoMessage() {}

func (x *UserRole) ProtoReflect() protoreflect.Message {
	mi := &file_xline_error_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRole.ProtoReflect.Descriptor instead.
func (*UserRole) Descriptor() ([]byte, []int) {
	return file_xline_error_proto_rawDescGZIP(), []int{2}
}

func (x *UserRole) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *UserRole) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

var File_xline_error_proto protoreflect.FileDescriptor

var file_xline_error_proto_rawDesc = []byte{
	0x0a, 0x11, 0x78, 0x6c, 0x69, 0x6e, 0x65, 0x2d, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x07, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x70, 0x62, 0x1a, 0x1b, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d,
	0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbd, 0x0c, 0x0a, 0x0c, 0x45, 0x78,
	0x65, 0x63, 0x75, 0x74, 0x65, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x3c, 0x0a, 0x0d, 0x6b, 0x65,
	0x79, 0x5f, 0x6e, 0x6f, 0x74, 0x5f, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x48, 0x00, 0x52, 0x0b, 0x6b, 0x65, 0x79,
	0x4e, 0x6f, 0x74, 0x46, 0x6f, 0x75, 0x6e, 0x64, 0x12, 0x42, 0x0a, 0x12, 0x72, 0x65, 0x76, 0x69,
	0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x6f, 0x6f, 0x5f, 0x6c, 0x61, 0x72, 0x67, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x70, 0x62, 0x2e, 0x52,
	0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x48, 0x00, 0x52, 0x10, 0x72, 0x65, 0x76, 0x69,
	0x73, 0x69, 0x6f, 0x6e, 0x54, 0x6f, 0x6f, 0x4c, 0x61, 0x72, 0x67, 0x65, 0x12, 0x43, 0x0a, 0x12,
	0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x63, 0x74,
	0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x70, 0x62, 0x2e, 0x52, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x48, 0x00, 0x52, 0x11,
	0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x63, 0x74, 0x65,
	0x64, 0x12, 0x28, 0x0a, 0x0f, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x5f, 0x6e, 0x6f, 0x74, 0x5f, 0x66,
	0x6f, 0x75, 0x6e, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x0d, 0x6c, 0x65,
	0x61, 0x73, 0x65, 0x4e, 0x6f, 0x74, 0x46, 0x6f, 0x75, 0x6e, 0x64, 0x12, 0x25, 0x0a, 0x0d, 0x6c,
	0x65, 0x61, 0x73, 0x65, 0x5f, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x03, 0x48, 0x00, 0x52, 0x0c, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x45, 0x78, 0x70, 0x69, 0x72,
	0x65, 0x64, 0x12, 0x2f, 0x0a, 0x13, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x5f, 0x74, 0x74, 0x6c, 0x5f,
	0x74, 0x6f, 0x6f, 0x5f, 0x6c, 0x61, 0x72, 0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x48,
	0x00, 0x52, 0x10, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x54, 0x74, 0x6c, 0x54, 0x6f, 0x6f, 0x4c, 0x61,
	0x72, 0x67, 0x65, 0x12, 0x32, 0x0a, 0x14, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x5f, 0x61, 0x6c, 0x72,
	0x65, 0x61, 0x64, 0x79, 0x5f, 0x65, 0x78, 0x69, 0x73, 0x74, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x03, 0x48, 0x00, 0x52, 0x12, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x41, 0x6c, 0x72, 0x65, 0x61, 0x64,
	0x79, 0x45, 0x78, 0x69, 0x73, 0x74, 0x73, 0x12, 0x42, 0x0a, 0x10, 0x61, 0x75, 0x74, 0x68, 0x5f,
	0x6e, 0x6f, 0x74, 0x5f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x48, 0x00, 0x52, 0x0e, 0x61, 0x75, 0x74,
	0x68, 0x4e, 0x6f, 0x74, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x39, 0x0a, 0x0b, 0x61,
	0x75, 0x74, 0x68, 0x5f, 0x66, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x48, 0x00, 0x52, 0x0a, 0x61, 0x75, 0x74, 0x68,
	0x46, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x12, 0x26, 0x0a, 0x0e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6e,
	0x6f, 0x74, 0x5f, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00,
	0x52, 0x0c, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x6f, 0x74, 0x46, 0x6f, 0x75, 0x6e, 0x64, 0x12, 0x30,
	0x0a, 0x13, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x61, 0x6c, 0x72, 0x65, 0x61, 0x64, 0x79, 0x5f, 0x65,
	0x78, 0x69, 0x73, 0x74, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x11, 0x75,
	0x73, 0x65, 0x72, 0x41, 0x6c, 0x72, 0x65, 0x61, 0x64, 0x79, 0x45, 0x78, 0x69, 0x73, 0x74, 0x73,
	0x12, 0x46, 0x0a, 0x15, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x61, 0x6c, 0x72, 0x65, 0x61, 0x64, 0x79,
	0x5f, 0x68, 0x61, 0x73, 0x5f, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x11, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x70, 0x62, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x6f,
	0x6c, 0x65, 0x48, 0x00, 0x52, 0x12, 0x75, 0x73, 0x65, 0x72, 0x41, 0x6c, 0x72, 0x65, 0x61, 0x64,
	0x79, 0x48, 0x61, 0x73, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x42, 0x0a, 0x10, 0x6e, 0x6f, 0x5f, 0x70,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x18, 0x0d, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x48, 0x00, 0x52, 0x0e, 0x6e, 0x6f,
	0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x55, 0x73, 0x65, 0x72, 0x12, 0x26, 0x0a, 0x0e,
	0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x6e, 0x6f, 0x74, 0x5f, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x18, 0x0e,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0c, 0x72, 0x6f, 0x6c, 0x65, 0x4e, 0x6f, 0x74, 0x46,
	0x6f, 0x75, 0x6e, 0x64, 0x12, 0x30, 0x0a, 0x13, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x61, 0x6c, 0x72,
	0x65, 0x61, 0x64, 0x79, 0x5f, 0x65, 0x78, 0x69, 0x73, 0x74, 0x73, 0x18, 0x0f, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x00, 0x52, 0x11, 0x72, 0x6f, 0x6c, 0x65, 0x41, 0x6c, 0x72, 0x65, 0x61, 0x64, 0x79,
	0x45, 0x78, 0x69, 0x73, 0x74, 0x73, 0x12, 0x2a, 0x0a, 0x10, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x6e,
	0x6f, 0x74, 0x5f, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x65, 0x64, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x00, 0x52, 0x0e, 0x72, 0x6f, 0x6c, 0x65, 0x4e, 0x6f, 0x74, 0x47, 0x72, 0x61, 0x6e, 0x74,
	0x65, 0x64, 0x12, 0x47, 0x0a, 0x13, 0x72, 0x6f, 0x6f, 0x74, 0x5f, 0x72, 0x6f, 0x6c, 0x65, 0x5f,
	0x6e, 0x6f, 0x74, 0x5f, 0x65, 0x78, 0x69, 0x73, 0x74, 0x18, 0x11, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x48, 0x00, 0x52, 0x10, 0x72, 0x6f, 0x6f, 0x74, 0x52,
	0x6f, 0x6c, 0x65, 0x4e, 0x6f, 0x74, 0x45, 0x78, 0x69, 0x73, 0x74, 0x12, 0x4e, 0x0a, 0x16, 0x70,
	0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x6f, 0x74, 0x5f, 0x67, 0x72,
	0x61, 0x6e, 0x74, 0x65, 0x64, 0x18, 0x12, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x48, 0x00, 0x52, 0x14, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x4e, 0x6f, 0x74, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x65, 0x64, 0x12, 0x4a, 0x0a, 0x14, 0x70,
	0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x6f, 0x74, 0x5f, 0x67, 0x69,
	0x76, 0x65, 0x6e, 0x18, 0x13, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x48, 0x00, 0x52, 0x12, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x4e,
	0x6f, 0x74, 0x47, 0x69, 0x76, 0x65, 0x6e, 0x12, 0x50, 0x0a, 0x17, 0x69, 0x6e, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x48, 0x00, 0x52, 0x15, 0x69, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x41, 0x75, 0x74, 0x68, 0x4d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x46, 0x0a, 0x12, 0x69, 0x6e, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x15, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x48, 0x00, 0x52,
	0x10, 0x69, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x41, 0x75, 0x74, 0x68, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x12, 0x4d, 0x0a, 0x16, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x5f, 0x6e, 0x6f, 0x74, 0x5f, 0x69, 0x6e, 0x69, 0x74, 0x18, 0x16, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x48, 0x00, 0x52, 0x13, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x4e, 0x6f, 0x74, 0x49, 0x6e, 0x69, 0x74,
	0x12, 0x46, 0x0a, 0x12, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x6e, 0x6f, 0x74, 0x5f, 0x70, 0x72,
	0x6f, 0x76, 0x69, 0x64, 0x65, 0x64, 0x18, 0x17, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x48, 0x00, 0x52, 0x10, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x4e, 0x6f, 0x74,
	0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x64, 0x12, 0x42, 0x0a, 0x12, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x5f, 0x6f, 0x6c, 0x64, 0x5f, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x18,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x70, 0x62, 0x2e, 0x52,
	0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x48, 0x00, 0x52, 0x10, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x4f, 0x6c, 0x64, 0x52, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x08,
	0x64, 0x62, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x19, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00,
	0x52, 0x07, 0x64, 0x62, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x45, 0x0a, 0x11, 0x70, 0x65, 0x72,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x65, 0x6e, 0x69, 0x65, 0x64, 0x18, 0x1a,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x48, 0x00, 0x52, 0x10,
	0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x6e, 0x69, 0x65, 0x64,
	0x42, 0x07, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x63, 0x0a, 0x09, 0x52, 0x65, 0x76,
	0x69, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x2b, 0x0a, 0x11, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72,
	0x65, 0x64, 0x5f, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x10, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x52, 0x65, 0x76, 0x69, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x29, 0x0a, 0x10, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x72,
	0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x63,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x32,
	0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x12,
	0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x6f,
	0x6c, 0x65, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x3b, 0x78, 0x6c, 0x69, 0x6e, 0x65, 0x61, 0x70, 0x69,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_xline_error_proto_rawDescOnce sync.Once
	file_xline_error_proto_rawDescData = file_xline_error_proto_rawDesc
)

func file_xline_error_proto_rawDescGZIP() []byte {
	file_xline_error_proto_rawDescOnce.Do(func() {
		file_xline_error_proto_rawDescData = protoimpl.X.CompressGZIP(file_xline_error_proto_rawDescData)
	})
	return file_xline_error_proto_rawDescData
}

var file_xline_error_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_xline_error_proto_goTypes = []interface{}{
	(*ExecuteError)(nil),  // 0: errorpb.ExecuteError
	(*Revisions)(nil),     // 1: errorpb.Revisions
	(*UserRole)(nil),      // 2: errorpb.UserRole
	(*emptypb.Empty)(nil), // 3: google.protobuf.Empty
}
var file_xline_error_proto_depIdxs = []int32{
	3,  // 0: errorpb.ExecuteError.key_not_found:type_name -> google.protobuf.Empty
	1,  // 1: errorpb.ExecuteError.revision_too_large:type_name -> errorpb.Revisions
	1,  // 2: errorpb.ExecuteError.revision_compacted:type_name -> errorpb.Revisions
	3,  // 3: errorpb.ExecuteError.auth_not_enabled:type_name -> google.protobuf.Empty
	3,  // 4: errorpb.ExecuteError.auth_failed:type_name -> google.protobuf.Empty
	2,  // 5: errorpb.ExecuteError.user_already_has_role:type_name -> errorpb.UserRole
	3,  // 6: errorpb.ExecuteError.no_password_user:type_name -> google.protobuf.Empty
	3,  // 7: errorpb.ExecuteError.root_role_not_exist:type_name -> google.protobuf.Empty
	3,  // 8: errorpb.ExecuteError.permission_not_granted:type_name -> google.protobuf.Empty
	3,  // 9: errorpb.ExecuteError.permission_not_given:type_name -> google.protobuf.Empty
	3,  // 10: errorpb.ExecuteError.invalid_auth_management:type_name -> google.protobuf.Empty
	3,  // 11: errorpb.ExecuteError.invalid_auth_token:type_name -> google.protobuf.Empty
	3,  // 12: errorpb.ExecuteError.token_manager_not_init:type_name -> google.protobuf.Empty
	3,  // 13: errorpb.ExecuteError.token_not_provided:type_name -> google.protobuf.Empty
	1,  // 14: errorpb.ExecuteError.token_old_revision:type_name -> errorpb.Revisions
	3,  // 15: errorpb.ExecuteError.permission_denied:type_name -> google.protobuf.Empty
	16, // [16:16] is the sub-list for method output_type
	16, // [16:16] is the sub-list for method input_type
	16, // [16:16] is the sub-list for extension type_name
	16, // [16:16] is the sub-list for extension extendee
	0,  // [0:16] is the sub-list for field type_name
}

func init() { file_xline_error_proto_init() }
func file_xline_error_proto_init() {
	if File_xline_error_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_xline_error_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExecuteError); i {
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
		file_xline_error_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Revisions); i {
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
		file_xline_error_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserRole); i {
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
	file_xline_error_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*ExecuteError_KeyNotFound)(nil),
		(*ExecuteError_RevisionTooLarge)(nil),
		(*ExecuteError_RevisionCompacted)(nil),
		(*ExecuteError_LeaseNotFound)(nil),
		(*ExecuteError_LeaseExpired)(nil),
		(*ExecuteError_LeaseTtlTooLarge)(nil),
		(*ExecuteError_LeaseAlreadyExists)(nil),
		(*ExecuteError_AuthNotEnabled)(nil),
		(*ExecuteError_AuthFailed)(nil),
		(*ExecuteError_UserNotFound)(nil),
		(*ExecuteError_UserAlreadyExists)(nil),
		(*ExecuteError_UserAlreadyHasRole)(nil),
		(*ExecuteError_NoPasswordUser)(nil),
		(*ExecuteError_RoleNotFound)(nil),
		(*ExecuteError_RoleAlreadyExists)(nil),
		(*ExecuteError_RoleNotGranted)(nil),
		(*ExecuteError_RootRoleNotExist)(nil),
		(*ExecuteError_PermissionNotGranted)(nil),
		(*ExecuteError_PermissionNotGiven)(nil),
		(*ExecuteError_InvalidAuthManagement)(nil),
		(*ExecuteError_InvalidAuthToken)(nil),
		(*ExecuteError_TokenManagerNotInit)(nil),
		(*ExecuteError_TokenNotProvided)(nil),
		(*ExecuteError_TokenOldRevision)(nil),
		(*ExecuteError_DbError)(nil),
		(*ExecuteError_PermissionDenied)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_xline_error_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_xline_error_proto_goTypes,
		DependencyIndexes: file_xline_error_proto_depIdxs,
		MessageInfos:      file_xline_error_proto_msgTypes,
	}.Build()
	File_xline_error_proto = out.File
	file_xline_error_proto_rawDesc = nil
	file_xline_error_proto_goTypes = nil
	file_xline_error_proto_depIdxs = nil
}
