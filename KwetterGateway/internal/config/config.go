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
	Config.UserServiceAddr = ":50054"
	Config.TrendServiceAddr = ":50052"
	Config.TweetServiceAddr = ":50051"
}
