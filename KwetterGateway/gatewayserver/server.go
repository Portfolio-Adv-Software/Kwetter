package gatewayserver

import (
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
	"time"
)

type AuthServiceServer struct {
	pb.UnimplementedAuthServiceServer
	AuthClient pb.AuthServiceClient
}

func (a AuthServiceServer) Register(ctx context.Context, req *pb.RegisterReq) (*pb.RegisterRes, error) {
	return a.AuthClient.Register(ctx, req)
}

func (a AuthServiceServer) Login(ctx context.Context, req *pb.LoginReq) (*pb.LoginRes, error) {
	log.Println(req.String())
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

type Server struct {
	authConn  *grpc.ClientConn
	userConn  *grpc.ClientConn
	trendConn *grpc.ClientConn
	tweetConn *grpc.ClientConn

	authClient  pb.AuthServiceClient
	userClient  pb.UserServiceClient
	trendClient pb.TrendServiceClient
	tweetClient pb.TweetServiceClient

	mux        *runtime.ServeMux
	grpcServer *grpc.Server
}

func InitMux(wg *sync.WaitGroup, config *ServiceConfig) {
	defer wg.Done()
	ctx := context.Background()
	server := &Server{mux: runtime.NewServeMux()}
	server.setupGRPCConnections(config)
	server.registerLocalEndpoints(ctx)
	srv := &http.Server{
		Addr:    ":8080",
		Handler: server.mux,
	}
	//registerEndpoints(ctx, mux, config)
	//handler := loggingMiddleware(mux)

	var srvWg sync.WaitGroup
	srvWg.Add(1)
	log.Println("starting gateway server on port 8080...")
	go func() {
		err := srv.ListenAndServe()
		if err != nil && err.Error() != "http: Server closed" {
			log.Fatalf("failed to start gateway server: %v", err)
		}
		srvWg.Done()
	}()
	log.Println("gateway server successfully started on port 8080...")
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	log.Println("\nShutting down gateway server...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Failed to gracefully shutdown gateway server: %v", err)
	}
	log.Println("gateway successfully shutdown.")
	srvWg.Wait()
}

//loggingMiddleware gives logs in terminal about http requests.
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

func (s *Server) setupGRPCConnections(config *ServiceConfig) {
	var err error
	s.authConn, err = grpc.Dial(config.AuthServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to dial authservice: %v", err)
	}
	s.authClient = pb.NewAuthServiceClient(s.authConn)

	s.userConn, err = grpc.Dial(config.UserServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to dial userservice: %v", err)
	}
	s.userClient = pb.NewUserServiceClient(s.userConn)

	s.trendConn, err = grpc.Dial(config.TrendServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to dial trendservice: %v", err)
	}
	s.trendClient = pb.NewTrendServiceClient(s.trendConn)

	s.tweetConn, err = grpc.Dial(config.TweetServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to dial tweetservice: %v", err)
	}
	s.tweetClient = pb.NewTweetServiceClient(s.tweetConn)
}

// registers mux to local gatewayserver functions which forward to endpoints.
func (s *Server) registerLocalEndpoints(ctx context.Context) {
	// Register the AuthServiceServer
	authServer := &AuthServiceServer{AuthClient: s.authClient}
	pb.RegisterAuthServiceHandlerServer(ctx, s.mux, authServer)

	// Register the UserServiceServer
	userServer := &UserServiceServer{UserClient: s.userClient}
	pb.RegisterUserServiceHandlerServer(ctx, s.mux, userServer)

	// Register the TrendServiceServer
	trendServer := &TrendServiceServer{TrendClient: s.trendClient}
	pb.RegisterTrendServiceHandlerServer(ctx, s.mux, trendServer)

	// Register the TweetServiceServer
	tweetServer := &TweetServiceServer{TweetClient: s.tweetClient}
	pb.RegisterTweetServiceHandlerServer(ctx, s.mux, tweetServer)
}

func (s *Server) StartGRPCServer(wg *sync.WaitGroup) {
	defer wg.Done()
	s.grpcServer = grpc.NewServer()
	pb.RegisterAuthServiceServer(s.grpcServer, &AuthServiceServer{AuthClient: s.authClient})
	pb.RegisterUserServiceServer(s.grpcServer, &UserServiceServer{UserClient: s.userClient})
	pb.RegisterTrendServiceServer(s.grpcServer, &TrendServiceServer{TrendClient: s.trendClient})
	pb.RegisterTweetServiceServer(s.grpcServer, &TweetServiceServer{TweetClient: s.tweetClient})
	reflection.Register(s.grpcServer)

	listener, err := net.Listen("tcp", ":50055")
	if err != nil {
		log.Fatalf("Unable to listen on port :50055: %v", err)
	}

	var srvWg sync.WaitGroup
	srvWg.Add(1)
	log.Println("Starting gRPC server on port 50055...")
	go func() {
		if err := s.grpcServer.Serve(listener); err != nil {
			log.Fatalf("Failed to serve gRPC: %v", err)
		}
		srvWg.Done()
	}()
	log.Println("gRPC server successfully started on port 50055...")

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	log.Println("\nStopping the gRPC server...")
	s.grpcServer.Stop()
	listener.Close()
	log.Println("gRPC successfully shutdown.")
	srvWg.Wait()
}
