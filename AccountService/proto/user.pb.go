// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.23.0
// source: user.proto

package __

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID   string `protobuf:"bytes,1,opt,name=UserID,proto3" json:"UserID,omitempty"`
	Email    string `protobuf:"bytes,2,opt,name=Email,proto3" json:"Email,omitempty"`
	Password string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	Username string `protobuf:"bytes,4,opt,name=Username,proto3" json:"Username,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *User) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *User) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type AuthData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *AuthData) Reset() {
	*x = AuthData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthData) ProtoMessage() {}

func (x *AuthData) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthData.ProtoReflect.Descriptor instead.
func (*AuthData) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{1}
}

func (x *AuthData) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type TweetData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID   string                 `protobuf:"bytes,1,opt,name=UserID,proto3" json:"UserID,omitempty"`
	Username string                 `protobuf:"bytes,2,opt,name=Username,proto3" json:"Username,omitempty"`
	TweetID  string                 `protobuf:"bytes,3,opt,name=TweetID,proto3" json:"TweetID,omitempty"`
	Body     string                 `protobuf:"bytes,4,opt,name=Body,proto3" json:"Body,omitempty"`
	Created  *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=created,proto3" json:"created,omitempty"`
}

func (x *TweetData) Reset() {
	*x = TweetData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TweetData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TweetData) ProtoMessage() {}

func (x *TweetData) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TweetData.ProtoReflect.Descriptor instead.
func (*TweetData) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{2}
}

func (x *TweetData) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *TweetData) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *TweetData) GetTweetID() string {
	if x != nil {
		return x.TweetID
	}
	return ""
}

func (x *TweetData) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

func (x *TweetData) GetCreated() *timestamppb.Timestamp {
	if x != nil {
		return x.Created
	}
	return nil
}

type UserData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AuthData  *AuthData    `protobuf:"bytes,1,opt,name=authData,proto3" json:"authData,omitempty"`
	TweetData []*TweetData `protobuf:"bytes,2,rep,name=tweetData,proto3" json:"tweetData,omitempty"`
}

func (x *UserData) Reset() {
	*x = UserData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserData) ProtoMessage() {}

func (x *UserData) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserData.ProtoReflect.Descriptor instead.
func (*UserData) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{3}
}

func (x *UserData) GetAuthData() *AuthData {
	if x != nil {
		return x.AuthData
	}
	return nil
}

func (x *UserData) GetTweetData() []*TweetData {
	if x != nil {
		return x.TweetData
	}
	return nil
}

type CreateUserReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *CreateUserReq) Reset() {
	*x = CreateUserReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserReq) ProtoMessage() {}

func (x *CreateUserReq) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserReq.ProtoReflect.Descriptor instead.
func (*CreateUserReq) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{4}
}

func (x *CreateUserReq) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type CreateUserRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *CreateUserRes) Reset() {
	*x = CreateUserRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserRes) ProtoMessage() {}

func (x *CreateUserRes) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserRes.ProtoReflect.Descriptor instead.
func (*CreateUserRes) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{5}
}

func (x *CreateUserRes) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type GetUserReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID string `protobuf:"bytes,1,opt,name=UserID,proto3" json:"UserID,omitempty"`
}

func (x *GetUserReq) Reset() {
	*x = GetUserReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserReq) ProtoMessage() {}

func (x *GetUserReq) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserReq.ProtoReflect.Descriptor instead.
func (*GetUserReq) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{6}
}

func (x *GetUserReq) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

type GetUserRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *GetUserRes) Reset() {
	*x = GetUserRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserRes) ProtoMessage() {}

func (x *GetUserRes) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserRes.ProtoReflect.Descriptor instead.
func (*GetUserRes) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{7}
}

func (x *GetUserRes) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type GetAllUsersReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAllUsersReq) Reset() {
	*x = GetAllUsersReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllUsersReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllUsersReq) ProtoMessage() {}

func (x *GetAllUsersReq) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllUsersReq.ProtoReflect.Descriptor instead.
func (*GetAllUsersReq) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{8}
}

type GetAllUsersRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User []*User `protobuf:"bytes,1,rep,name=user,proto3" json:"user,omitempty"`
}

func (x *GetAllUsersRes) Reset() {
	*x = GetAllUsersRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllUsersRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllUsersRes) ProtoMessage() {}

func (x *GetAllUsersRes) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllUsersRes.ProtoReflect.Descriptor instead.
func (*GetAllUsersRes) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{9}
}

func (x *GetAllUsersRes) GetUser() []*User {
	if x != nil {
		return x.User
	}
	return nil
}

type UpdateUserReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *UpdateUserReq) Reset() {
	*x = UpdateUserReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateUserReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUserReq) ProtoMessage() {}

func (x *UpdateUserReq) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUserReq.ProtoReflect.Descriptor instead.
func (*UpdateUserReq) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{10}
}

func (x *UpdateUserReq) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type UpdateUserRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *UpdateUserRes) Reset() {
	*x = UpdateUserRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateUserRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUserRes) ProtoMessage() {}

func (x *UpdateUserRes) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUserRes.ProtoReflect.Descriptor instead.
func (*UpdateUserRes) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{11}
}

func (x *UpdateUserRes) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type DeleteUserReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID string `protobuf:"bytes,1,opt,name=UserID,proto3" json:"UserID,omitempty"`
}

func (x *DeleteUserReq) Reset() {
	*x = DeleteUserReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteUserReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteUserReq) ProtoMessage() {}

func (x *DeleteUserReq) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteUserReq.ProtoReflect.Descriptor instead.
func (*DeleteUserReq) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{12}
}

func (x *DeleteUserReq) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

type DeleteUserRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *DeleteUserRes) Reset() {
	*x = DeleteUserRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteUserRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteUserRes) ProtoMessage() {}

func (x *DeleteUserRes) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteUserRes.ProtoReflect.Descriptor instead.
func (*DeleteUserRes) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{13}
}

func (x *DeleteUserRes) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type GetAllUserDataReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID string `protobuf:"bytes,1,opt,name=UserID,proto3" json:"UserID,omitempty"`
}

func (x *GetAllUserDataReq) Reset() {
	*x = GetAllUserDataReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[14]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllUserDataReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllUserDataReq) ProtoMessage() {}

func (x *GetAllUserDataReq) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[14]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllUserDataReq.ProtoReflect.Descriptor instead.
func (*GetAllUserDataReq) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{14}
}

func (x *GetAllUserDataReq) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

type GetAllUserDataRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Userdata *UserData `protobuf:"bytes,1,opt,name=userdata,proto3" json:"userdata,omitempty"`
}

func (x *GetAllUserDataRes) Reset() {
	*x = GetAllUserDataRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[15]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllUserDataRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllUserDataRes) ProtoMessage() {}

func (x *GetAllUserDataRes) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[15]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllUserDataRes.ProtoReflect.Descriptor instead.
func (*GetAllUserDataRes) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{15}
}

func (x *GetAllUserDataRes) GetUserdata() *UserData {
	if x != nil {
		return x.Userdata
	}
	return nil
}

var File_user_proto protoreflect.FileDescriptor

var file_user_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6c, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x22, 0x20, 0x0a, 0x08, 0x41, 0x75, 0x74, 0x68, 0x44, 0x61, 0x74, 0x61, 0x12, 0x14,
	0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x22, 0xa3, 0x01, 0x0a, 0x09, 0x54, 0x77, 0x65, 0x65, 0x74, 0x44, 0x61,
	0x74, 0x61, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x55, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x55, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x54, 0x77, 0x65, 0x65, 0x74, 0x49,
	0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x54, 0x77, 0x65, 0x65, 0x74, 0x49, 0x44,
	0x12, 0x12, 0x0a, 0x04, 0x42, 0x6f, 0x64, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x42, 0x6f, 0x64, 0x79, 0x12, 0x34, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x22, 0x67, 0x0a, 0x08, 0x55, 0x73,
	0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x12, 0x2b, 0x0a, 0x08, 0x61, 0x75, 0x74, 0x68, 0x44, 0x61,
	0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x41, 0x75, 0x74, 0x68, 0x44, 0x61, 0x74, 0x61, 0x52, 0x08, 0x61, 0x75, 0x74, 0x68, 0x44,
	0x61, 0x74, 0x61, 0x12, 0x2e, 0x0a, 0x09, 0x74, 0x77, 0x65, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54,
	0x77, 0x65, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x09, 0x74, 0x77, 0x65, 0x65, 0x74, 0x44,
	0x61, 0x74, 0x61, 0x22, 0x30, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x12, 0x1f, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x30, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x12, 0x1f, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x24, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x22, 0x2d, 0x0a,
	0x0a, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x12, 0x1f, 0x0a, 0x04, 0x75,
	0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x10, 0x0a, 0x0e,
	0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x22, 0x31,
	0x0a, 0x0e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73,
	0x12, 0x1f, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65,
	0x72, 0x22, 0x30, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x12, 0x1f, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75,
	0x73, 0x65, 0x72, 0x22, 0x30, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x12, 0x1f, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x27, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x22, 0x30,
	0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x12,
	0x1f, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72,
	0x22, 0x2b, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x44, 0x61,
	0x74, 0x61, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x22, 0x40, 0x0a,
	0x11, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x52,
	0x65, 0x73, 0x12, 0x2b, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x44, 0x61, 0x74, 0x61, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x64, 0x61, 0x74, 0x61, 0x32,
	0xef, 0x02, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x38, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x14, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x1a, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x12, 0x2f, 0x0a, 0x07, 0x47, 0x65, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x12, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x12, 0x3b, 0x0a, 0x0b, 0x47, 0x65,
	0x74, 0x41, 0x4c, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71,
	0x1a, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x55,
	0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x12, 0x38, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x14, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x12, 0x38, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12,
	0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x12, 0x44, 0x0a, 0x0e, 0x47,
	0x65, 0x74, 0x41, 0x6c, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x12, 0x18, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x55, 0x73, 0x65, 0x72,
	0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x1a, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65,
	0x73, 0x42, 0x03, 0x5a, 0x01, 0x2e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_user_proto_rawDescOnce sync.Once
	file_user_proto_rawDescData = file_user_proto_rawDesc
)

func file_user_proto_rawDescGZIP() []byte {
	file_user_proto_rawDescOnce.Do(func() {
		file_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_proto_rawDescData)
	})
	return file_user_proto_rawDescData
}

var file_user_proto_msgTypes = make([]protoimpl.MessageInfo, 16)
var file_user_proto_goTypes = []interface{}{
	(*User)(nil),                  // 0: proto.User
	(*AuthData)(nil),              // 1: proto.AuthData
	(*TweetData)(nil),             // 2: proto.TweetData
	(*UserData)(nil),              // 3: proto.UserData
	(*CreateUserReq)(nil),         // 4: proto.CreateUserReq
	(*CreateUserRes)(nil),         // 5: proto.CreateUserRes
	(*GetUserReq)(nil),            // 6: proto.GetUserReq
	(*GetUserRes)(nil),            // 7: proto.GetUserRes
	(*GetAllUsersReq)(nil),        // 8: proto.GetAllUsersReq
	(*GetAllUsersRes)(nil),        // 9: proto.GetAllUsersRes
	(*UpdateUserReq)(nil),         // 10: proto.UpdateUserReq
	(*UpdateUserRes)(nil),         // 11: proto.UpdateUserRes
	(*DeleteUserReq)(nil),         // 12: proto.DeleteUserReq
	(*DeleteUserRes)(nil),         // 13: proto.DeleteUserRes
	(*GetAllUserDataReq)(nil),     // 14: proto.GetAllUserDataReq
	(*GetAllUserDataRes)(nil),     // 15: proto.GetAllUserDataRes
	(*timestamppb.Timestamp)(nil), // 16: google.protobuf.Timestamp
}
var file_user_proto_depIdxs = []int32{
	16, // 0: proto.TweetData.created:type_name -> google.protobuf.Timestamp
	1,  // 1: proto.UserData.authData:type_name -> proto.AuthData
	2,  // 2: proto.UserData.tweetData:type_name -> proto.TweetData
	0,  // 3: proto.CreateUserReq.user:type_name -> proto.User
	0,  // 4: proto.CreateUserRes.user:type_name -> proto.User
	0,  // 5: proto.GetUserRes.user:type_name -> proto.User
	0,  // 6: proto.GetAllUsersRes.user:type_name -> proto.User
	0,  // 7: proto.UpdateUserReq.user:type_name -> proto.User
	0,  // 8: proto.UpdateUserRes.user:type_name -> proto.User
	0,  // 9: proto.DeleteUserRes.user:type_name -> proto.User
	3,  // 10: proto.GetAllUserDataRes.userdata:type_name -> proto.UserData
	4,  // 11: proto.UserService.CreateUser:input_type -> proto.CreateUserReq
	6,  // 12: proto.UserService.GetUser:input_type -> proto.GetUserReq
	8,  // 13: proto.UserService.GetALlUsers:input_type -> proto.GetAllUsersReq
	10, // 14: proto.UserService.UpdateUser:input_type -> proto.UpdateUserReq
	12, // 15: proto.UserService.DeleteUser:input_type -> proto.DeleteUserReq
	14, // 16: proto.UserService.GetAllUserData:input_type -> proto.GetAllUserDataReq
	5,  // 17: proto.UserService.CreateUser:output_type -> proto.CreateUserRes
	7,  // 18: proto.UserService.GetUser:output_type -> proto.GetUserRes
	9,  // 19: proto.UserService.GetALlUsers:output_type -> proto.GetAllUsersRes
	11, // 20: proto.UserService.UpdateUser:output_type -> proto.UpdateUserRes
	13, // 21: proto.UserService.DeleteUser:output_type -> proto.DeleteUserRes
	15, // 22: proto.UserService.GetAllUserData:output_type -> proto.GetAllUserDataRes
	17, // [17:23] is the sub-list for method output_type
	11, // [11:17] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_user_proto_init() }
func file_user_proto_init() {
	if File_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthData); i {
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
		file_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TweetData); i {
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
		file_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserData); i {
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
		file_user_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUserReq); i {
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
		file_user_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUserRes); i {
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
		file_user_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserReq); i {
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
		file_user_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserRes); i {
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
		file_user_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllUsersReq); i {
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
		file_user_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllUsersRes); i {
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
		file_user_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateUserReq); i {
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
		file_user_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateUserRes); i {
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
		file_user_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteUserReq); i {
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
		file_user_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteUserRes); i {
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
		file_user_proto_msgTypes[14].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllUserDataReq); i {
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
		file_user_proto_msgTypes[15].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllUserDataRes); i {
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
			RawDescriptor: file_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   16,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_user_proto_goTypes,
		DependencyIndexes: file_user_proto_depIdxs,
		MessageInfos:      file_user_proto_msgTypes,
	}.Build()
	File_user_proto = out.File
	file_user_proto_rawDesc = nil
	file_user_proto_goTypes = nil
	file_user_proto_depIdxs = nil
}
