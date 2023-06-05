// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.23.0
// source: tweet.proto

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

type Tweet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID   string                 `protobuf:"bytes,1,opt,name=UserID,proto3" json:"UserID,omitempty"`
	Username string                 `protobuf:"bytes,2,opt,name=Username,proto3" json:"Username,omitempty"`
	TweetID  string                 `protobuf:"bytes,3,opt,name=TweetID,proto3" json:"TweetID,omitempty"`
	Body     string                 `protobuf:"bytes,4,opt,name=Body,proto3" json:"Body,omitempty"`
	Created  *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=created,proto3" json:"created,omitempty"`
	Trend    []string               `protobuf:"bytes,6,rep,name=Trend,proto3" json:"Trend,omitempty"`
}

func (x *Tweet) Reset() {
	*x = Tweet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tweet_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Tweet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tweet) ProtoMessage() {}

func (x *Tweet) ProtoReflect() protoreflect.Message {
	mi := &file_tweet_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tweet.ProtoReflect.Descriptor instead.
func (*Tweet) Descriptor() ([]byte, []int) {
	return file_tweet_proto_rawDescGZIP(), []int{0}
}

func (x *Tweet) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *Tweet) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *Tweet) GetTweetID() string {
	if x != nil {
		return x.TweetID
	}
	return ""
}

func (x *Tweet) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

func (x *Tweet) GetCreated() *timestamppb.Timestamp {
	if x != nil {
		return x.Created
	}
	return nil
}

func (x *Tweet) GetTrend() []string {
	if x != nil {
		return x.Trend
	}
	return nil
}

// posttrend
type PostTrendReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tweet *Tweet `protobuf:"bytes,1,opt,name=tweet,proto3" json:"tweet,omitempty"`
}

func (x *PostTrendReq) Reset() {
	*x = PostTrendReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tweet_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostTrendReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostTrendReq) ProtoMessage() {}

func (x *PostTrendReq) ProtoReflect() protoreflect.Message {
	mi := &file_tweet_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostTrendReq.ProtoReflect.Descriptor instead.
func (*PostTrendReq) Descriptor() ([]byte, []int) {
	return file_tweet_proto_rawDescGZIP(), []int{1}
}

func (x *PostTrendReq) GetTweet() *Tweet {
	if x != nil {
		return x.Tweet
	}
	return nil
}

type PostTrendRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tweet *Tweet `protobuf:"bytes,1,opt,name=tweet,proto3" json:"tweet,omitempty"`
}

func (x *PostTrendRes) Reset() {
	*x = PostTrendRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tweet_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostTrendRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostTrendRes) ProtoMessage() {}

func (x *PostTrendRes) ProtoReflect() protoreflect.Message {
	mi := &file_tweet_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostTrendRes.ProtoReflect.Descriptor instead.
func (*PostTrendRes) Descriptor() ([]byte, []int) {
	return file_tweet_proto_rawDescGZIP(), []int{2}
}

func (x *PostTrendRes) GetTweet() *Tweet {
	if x != nil {
		return x.Tweet
	}
	return nil
}

// gettrend
type GetTrendReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Trends string `protobuf:"bytes,1,opt,name=Trends,proto3" json:"Trends,omitempty"`
}

func (x *GetTrendReq) Reset() {
	*x = GetTrendReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tweet_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTrendReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTrendReq) ProtoMessage() {}

func (x *GetTrendReq) ProtoReflect() protoreflect.Message {
	mi := &file_tweet_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTrendReq.ProtoReflect.Descriptor instead.
func (*GetTrendReq) Descriptor() ([]byte, []int) {
	return file_tweet_proto_rawDescGZIP(), []int{3}
}

func (x *GetTrendReq) GetTrends() string {
	if x != nil {
		return x.Trends
	}
	return ""
}

type GetTrendRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tweets []*Tweet `protobuf:"bytes,1,rep,name=tweets,proto3" json:"tweets,omitempty"`
}

func (x *GetTrendRes) Reset() {
	*x = GetTrendRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tweet_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTrendRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTrendRes) ProtoMessage() {}

func (x *GetTrendRes) ProtoReflect() protoreflect.Message {
	mi := &file_tweet_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTrendRes.ProtoReflect.Descriptor instead.
func (*GetTrendRes) Descriptor() ([]byte, []int) {
	return file_tweet_proto_rawDescGZIP(), []int{4}
}

func (x *GetTrendRes) GetTweets() []*Tweet {
	if x != nil {
		return x.Tweets
	}
	return nil
}

// deletedata
type DeleteDataReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Userid        string `protobuf:"bytes,1,opt,name=userid,proto3" json:"userid,omitempty"`
	CorrelationId string `protobuf:"bytes,2,opt,name=CorrelationId,proto3" json:"CorrelationId,omitempty"`
}

func (x *DeleteDataReq) Reset() {
	*x = DeleteDataReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tweet_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteDataReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteDataReq) ProtoMessage() {}

func (x *DeleteDataReq) ProtoReflect() protoreflect.Message {
	mi := &file_tweet_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteDataReq.ProtoReflect.Descriptor instead.
func (*DeleteDataReq) Descriptor() ([]byte, []int) {
	return file_tweet_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteDataReq) GetUserid() string {
	if x != nil {
		return x.Userid
	}
	return ""
}

func (x *DeleteDataReq) GetCorrelationId() string {
	if x != nil {
		return x.CorrelationId
	}
	return ""
}

type DeleteDataRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *DeleteDataRes) Reset() {
	*x = DeleteDataRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tweet_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteDataRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteDataRes) ProtoMessage() {}

func (x *DeleteDataRes) ProtoReflect() protoreflect.Message {
	mi := &file_tweet_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteDataRes.ProtoReflect.Descriptor instead.
func (*DeleteDataRes) Descriptor() ([]byte, []int) {
	return file_tweet_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteDataRes) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

var File_tweet_proto protoreflect.FileDescriptor

var file_tweet_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x74, 0x77, 0x65, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb5, 0x01, 0x0a, 0x05, 0x54, 0x77, 0x65, 0x65, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x54, 0x77, 0x65, 0x65, 0x74, 0x49, 0x44, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x54, 0x77, 0x65, 0x65, 0x74, 0x49, 0x44, 0x12, 0x12, 0x0a,
	0x04, 0x42, 0x6f, 0x64, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x42, 0x6f, 0x64,
	0x79, 0x12, 0x34, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x72, 0x65, 0x6e, 0x64,
	0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x54, 0x72, 0x65, 0x6e, 0x64, 0x22, 0x32, 0x0a,
	0x0c, 0x50, 0x6f, 0x73, 0x74, 0x54, 0x72, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x12, 0x22, 0x0a,
	0x05, 0x74, 0x77, 0x65, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x77, 0x65, 0x65, 0x74, 0x52, 0x05, 0x74, 0x77, 0x65, 0x65,
	0x74, 0x22, 0x32, 0x0a, 0x0c, 0x50, 0x6f, 0x73, 0x74, 0x54, 0x72, 0x65, 0x6e, 0x64, 0x52, 0x65,
	0x73, 0x12, 0x22, 0x0a, 0x05, 0x74, 0x77, 0x65, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x77, 0x65, 0x65, 0x74, 0x52, 0x05,
	0x74, 0x77, 0x65, 0x65, 0x74, 0x22, 0x25, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x54, 0x72, 0x65, 0x6e,
	0x64, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x54, 0x72, 0x65, 0x6e, 0x64, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x54, 0x72, 0x65, 0x6e, 0x64, 0x73, 0x22, 0x33, 0x0a, 0x0b,
	0x47, 0x65, 0x74, 0x54, 0x72, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x12, 0x24, 0x0a, 0x06, 0x74,
	0x77, 0x65, 0x65, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x77, 0x65, 0x65, 0x74, 0x52, 0x06, 0x74, 0x77, 0x65, 0x65, 0x74,
	0x73, 0x22, 0x4d, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52,
	0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x69, 0x64, 0x12, 0x24, 0x0a, 0x0d, 0x43, 0x6f,
	0x72, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x43, 0x6f, 0x72, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64,
	0x22, 0x27, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65,
	0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x32, 0xb3, 0x01, 0x0a, 0x0c, 0x54, 0x72,
	0x65, 0x6e, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x35, 0x0a, 0x09, 0x50, 0x6f,
	0x73, 0x74, 0x54, 0x72, 0x65, 0x6e, 0x64, 0x12, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x50, 0x6f, 0x73, 0x74, 0x54, 0x72, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x13, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x54, 0x72, 0x65, 0x6e, 0x64, 0x52, 0x65,
	0x73, 0x12, 0x32, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x54, 0x72, 0x65, 0x6e, 0x64, 0x12, 0x12, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x72, 0x65, 0x6e, 0x64, 0x52, 0x65,
	0x71, 0x1a, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x72, 0x65,
	0x6e, 0x64, 0x52, 0x65, 0x73, 0x12, 0x38, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44,
	0x61, 0x74, 0x61, 0x12, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x1a, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x42,
	0x03, 0x5a, 0x01, 0x2e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tweet_proto_rawDescOnce sync.Once
	file_tweet_proto_rawDescData = file_tweet_proto_rawDesc
)

func file_tweet_proto_rawDescGZIP() []byte {
	file_tweet_proto_rawDescOnce.Do(func() {
		file_tweet_proto_rawDescData = protoimpl.X.CompressGZIP(file_tweet_proto_rawDescData)
	})
	return file_tweet_proto_rawDescData
}

var file_tweet_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_tweet_proto_goTypes = []interface{}{
	(*Tweet)(nil),                 // 0: proto.Tweet
	(*PostTrendReq)(nil),          // 1: proto.PostTrendReq
	(*PostTrendRes)(nil),          // 2: proto.PostTrendRes
	(*GetTrendReq)(nil),           // 3: proto.GetTrendReq
	(*GetTrendRes)(nil),           // 4: proto.GetTrendRes
	(*DeleteDataReq)(nil),         // 5: proto.DeleteDataReq
	(*DeleteDataRes)(nil),         // 6: proto.DeleteDataRes
	(*timestamppb.Timestamp)(nil), // 7: google.protobuf.Timestamp
}
var file_tweet_proto_depIdxs = []int32{
	7, // 0: proto.Tweet.created:type_name -> google.protobuf.Timestamp
	0, // 1: proto.PostTrendReq.tweet:type_name -> proto.Tweet
	0, // 2: proto.PostTrendRes.tweet:type_name -> proto.Tweet
	0, // 3: proto.GetTrendRes.tweets:type_name -> proto.Tweet
	1, // 4: proto.TrendService.PostTrend:input_type -> proto.PostTrendReq
	3, // 5: proto.TrendService.GetTrend:input_type -> proto.GetTrendReq
	5, // 6: proto.TrendService.DeleteData:input_type -> proto.DeleteDataReq
	2, // 7: proto.TrendService.PostTrend:output_type -> proto.PostTrendRes
	4, // 8: proto.TrendService.GetTrend:output_type -> proto.GetTrendRes
	6, // 9: proto.TrendService.DeleteData:output_type -> proto.DeleteDataRes
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_tweet_proto_init() }
func file_tweet_proto_init() {
	if File_tweet_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tweet_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Tweet); i {
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
		file_tweet_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostTrendReq); i {
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
		file_tweet_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostTrendRes); i {
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
		file_tweet_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTrendReq); i {
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
		file_tweet_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTrendRes); i {
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
		file_tweet_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteDataReq); i {
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
		file_tweet_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteDataRes); i {
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
			RawDescriptor: file_tweet_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_tweet_proto_goTypes,
		DependencyIndexes: file_tweet_proto_depIdxs,
		MessageInfos:      file_tweet_proto_msgTypes,
	}.Build()
	File_tweet_proto = out.File
	file_tweet_proto_rawDesc = nil
	file_tweet_proto_goTypes = nil
	file_tweet_proto_depIdxs = nil
}
