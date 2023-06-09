package config

type ServiceConfig struct {
	AuthServiceAddr  string
	UserServiceAddr  string
	TrendServiceAddr string
	TweetServiceAddr string
}

var Config ServiceConfig

func SetConfig() {
	Config.AuthServiceAddr = "authservice:50053"
	Config.UserServiceAddr = "accountservice:50054"
	Config.TrendServiceAddr = "trendservice:50052"
	Config.TweetServiceAddr = "tweetservice:50051"
}
