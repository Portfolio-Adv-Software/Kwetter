package config

type ServiceConfig struct {
	AuthServiceAddr  string
	UserServiceAddr  string
	TrendServiceAddr string
	TweetServiceAddr string
}

var Config ServiceConfig

func SetConfig() {
	Config.AuthServiceAddr = ":50053"
	Config.UserServiceAddr = "account-service.default:50054"
	Config.TrendServiceAddr = "trend-service.default:50052"
	Config.TweetServiceAddr = "tweet-service.default:50051"
}
