package models

import "google.golang.org/protobuf/types/known/timestamppb"

type Tweet struct {
	UserID   string                 `bson:"_userid,omitempty"`
	Username string                 `bson:"username"`
	TweetID  string                 `bson:"tweetid"`
	Body     string                 `bson:"body"`
	Created  *timestamppb.Timestamp `bson:"created"`
}
