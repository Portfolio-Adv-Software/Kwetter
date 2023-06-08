package config

type ServiceConfig struct {
	AuthServiceAddr  string
	UserServiceAddr  string
	TrendServiceAddr string
	TweetServiceAddr string
}

var Config ServiceConfig

func SetConfig() {
	Config.AuthServiceAddr = "auth-service:50053"
	Config.UserServiceAddr = "account-service:50054"
	Config.TrendServiceAddr = "trend-service:50052"
	Config.TweetServiceAddr = "tweet-service:50051"
}
