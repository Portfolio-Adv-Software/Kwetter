package gatewayserver

import (
	"fmt"
	pb "github.com/Portfolio-Adv-Software/Kwetter/KwetterGateway/proto"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync"
)

type AuthServiceServer struct {
	pb.UnimplementedAuthServiceServer
	AuthClient pb.AuthServiceClient
}

func (a AuthServiceServer) Register(ctx context.Context, req *pb.RegisterReq) (*pb.RegisterRes, error) {
	return a.AuthClient.Register(ctx, req)
}

func (a AuthServiceServer) Login(ctx context.Context, req *pb.LoginReq) (*pb.LoginRes, error) {
	return a.AuthClient.Login(ctx, req)
}

func (a AuthServiceServer) Validate(ctx context.Context, req *pb.ValidateReq) (*pb.ValidateRes, error) {
	return a.AuthClient.Validate(ctx, req)
}

type UserServiceServer struct {
	pb.UnimplementedUserServiceServer
	UserClient pb.UserServiceClient
}

func (u UserServiceServer) GetUser(ctx context.Context, req *pb.GetUserReq) (*pb.GetUserRes, error) {
	return u.UserClient.GetUser(ctx, req)
}

func (u UserServiceServer) GetAllUsers(ctx context.Context, req *pb.GetAllUsersReq) (*pb.GetAllUsersRes, error) {
	return u.UserClient.GetAllUsers(ctx, req)
}

func (u UserServiceServer) UpdateUser(ctx context.Context, req *pb.UpdateUserReq) (*pb.UpdateUserRes, error) {
	return u.UserClient.UpdateUser(ctx, req)
}

func (u UserServiceServer) DeleteUser(ctx context.Context, req *pb.DeleteUserReq) (*pb.UpdateUserRes, error) {
	return u.UserClient.DeleteUser(ctx, req)
}

type TrendServiceServer struct {
	pb.UnimplementedTrendServiceServer
	TrendClient pb.TrendServiceClient
}

func (t TrendServiceServer) GetTrend(ctx context.Context, req *pb.GetTrendReq) (*pb.GetTrendRes, error) {
	return t.TrendClient.GetTrend(ctx, req)
}

type TweetServiceServer struct {
	pb.UnimplementedTweetServiceServer
	TweetClient pb.TweetServiceClient
}

func (t TweetServiceServer) ReturnAll(ctx context.Context, req *pb.ReturnAllReq) (*pb.ReturnAllRes, error) {
	return t.TweetClient.ReturnAll(ctx, req)
}

func (t TweetServiceServer) ReturnTweet(ctx context.Context, req *pb.ReturnTweetReq) (*pb.ReturnTweetRes, error) {
	return t.TweetClient.ReturnTweet(ctx, req)
}

func (t TweetServiceServer) PostTweet(ctx context.Context, req *pb.PostTweetReq) (*pb.PostTweetRes, error) {
	return t.TweetClient.PostTweet(ctx, req)
}

func InitGRPC(wg *sync.WaitGroup, config *ServiceConfig) {
	defer wg.Done()
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	fmt.Println("Starting server on port: 50055")

	listener, err := net.Listen("tcp", ":50055")
	if err != nil {
		log.Fatalf("Unable to listen on port :50055: %v", err)
	}

	var opts []grpc.ServerOption
	s := grpc.NewServer(opts...)
	registerAuthService(s, config)
	registerUserService(s, config)
	registerTrendService(s, config)
	registerTweetService(s, config)
	reflection.Register(s)

	go func() {
		if err := s.Serve(listener); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()
	fmt.Println("Server succesfully started on port :50055")
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)

	// Block main routine until a signal is received
	// As long as user doesn't press CTRL+C a message is not passed and our main routine keeps running
	<-c
	// After receiving CTRL+C Properly stop the server
	fmt.Println("\nStopping the server...")
	s.Stop()
	listener.Close()
	fmt.Println("Done.")
}

func InitMux(wg *sync.WaitGroup, config *ServiceConfig) {
	defer wg.Done()
	ctx := context.Background()
	mux := runtime.NewServeMux()
	registerEndpoints(ctx, mux, config)
	//handler := loggingMiddleware(mux)
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatalf("failed to start gateway server: %v", err)
	}
}

//func loggingMiddleware(next http.Handler) http.Handler {
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		log.Printf("Received request: %s %s", r.Method, r.URL.Path)
//		body, err := io.ReadAll(r.Body)
//		if err != nil {
//			log.Printf("Failed to read request body: %v", err)
//		} else {
//			// Log the request body
//			log.Printf("Request body: %s", string(body))
//		}
//		r.Body = io.NopCloser(bytes.NewBuffer(body))
//		next.ServeHTTP(w, r)
//		log.Printf("Sent response: %d", w.(http.ResponseWriter))
//	})
//}

func registerEndpoints(ctx context.Context, mux *runtime.ServeMux, config *ServiceConfig) {
	err := pb.RegisterAuthServiceHandlerFromEndpoint(ctx, mux, config.AuthServiceAddr, []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())})
	if err != nil {
		log.Fatalf("failed to register AuthService handler: %v", err)
	}
	err = pb.RegisterUserServiceHandlerFromEndpoint(ctx, mux, config.UserServiceAddr, []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())})
	if err != nil {
		log.Fatalf("failed to register UserService handler: %v", err)
	}
	err = pb.RegisterTrendServiceHandlerFromEndpoint(ctx, mux, config.TrendServiceAddr, []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())})
	if err != nil {
		log.Fatalf("failed to register TrendService handler: %v", err)
	}
	err = pb.RegisterTweetServiceHandlerFromEndpoint(ctx, mux, config.TweetServiceAddr, []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())})
	if err != nil {
		log.Fatalf("failed to register TweetService handler: %v", err)
	}
}

func registerAuthService(s *grpc.Server, config *ServiceConfig) {
	authConn, err := grpc.Dial(config.AuthServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to dial authservice: %v", err)
	}
	authClient := pb.NewAuthServiceClient(authConn)
	authServer := &AuthServiceServer{AuthClient: authClient}
	pb.RegisterAuthServiceServer(s, authServer)
}
func registerUserService(s *grpc.Server, config *ServiceConfig) {
	userConn, err := grpc.Dial(config.UserServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to dial userservice: %v", err)
	}
	userClient := pb.NewUserServiceClient(userConn)
	userServer := &UserServiceServer{UserClient: userClient}
	pb.RegisterUserServiceServer(s, userServer)
}
func registerTrendService(s *grpc.Server, config *ServiceConfig) {
	trendConn, err := grpc.Dial(config.TrendServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to dial trendservice: %v", err)
	}
	trendClient := pb.NewTrendServiceClient(trendConn)
	trendServer := &TrendServiceServer{TrendClient: trendClient}
	pb.RegisterTrendServiceServer(s, trendServer)
}
func registerTweetService(s *grpc.Server, config *ServiceConfig) {
	tweetConn, err := grpc.Dial(config.TweetServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to dial tweetservice: %v", err)
	}
	tweetClient := pb.NewTweetServiceClient(tweetConn)
	tweetServer := &TweetServiceServer{TweetClient: tweetClient}
	pb.RegisterTweetServiceServer(s, tweetServer)
}
