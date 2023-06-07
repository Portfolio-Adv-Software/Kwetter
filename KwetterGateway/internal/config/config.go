package config

type ServiceConfig struct {
	AuthServiceAddr  string
	UserServiceAddr  string
	TrendServiceAddr string
	TweetServiceAddr string
}

var Config ServiceConfig

func SetConfig() {
	Config.AuthServiceAddr = "localhost:50053"
	Config.UserServiceAddr = "localhost:50054"
	Config.TrendServiceAddr = "localhost:50052"
	Config.TweetServiceAddr = "localhost:50051"
}
