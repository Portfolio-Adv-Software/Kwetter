package gatewayserver

import (
	"bytes"
	"fmt"
	"github.com/Portfolio-Adv-Software/Kwetter/KwetterGateway/config"
	pb "github.com/Portfolio-Adv-Software/Kwetter/KwetterGateway/proto"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"sync"
)

type UserDataServiceServer struct {
	pb.UnimplementedUserDataServiceServer
	AuthClient  pb.AuthServiceClient
	UserClient  pb.UserServiceClient
	TweetClient pb.TweetServiceClient
}

func (u UserDataServiceServer) GetAllUserData(ctx context.Context, req *pb.GetAllUserDataReq) (*pb.GetAllUserDataRes, error) {
	userId := req.GetUserId()
	authReq := &pb.GetDataReq{UserId: userId}
	userReq := &pb.GetUserReq{UserID: userId}
	tweetReq := &pb.ReturnAllReq{UserId: userId}
	authRes, err := u.AuthClient.GetData(ctx, authReq)
	if err != nil {
		return nil, err
	}
	userRes, err := u.UserClient.GetUser(ctx, userReq)
	if err != nil {
		return nil, err
	}
	tweetRes, err := u.TweetClient.ReturnAll(ctx, tweetReq)
	if err != nil {
		return nil, err
	}
	res := &pb.AllUserData{
		User:      userRes.GetUser(),
		AuthData:  authRes.GetAuthData(),
		TweetData: tweetRes.GetTweetData(),
	}
	return &pb.GetAllUserDataRes{AllUserData: res}, nil
}

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

func (u UserServiceServer) UpdateUser(ctx context.Context, req *pb.UpdateUserReq) (*pb.UpdateUserRes, error) {
	return u.UserClient.UpdateUser(ctx, req)
}

func (u UserServiceServer) DeleteUser(ctx context.Context, req *pb.DeleteUserReq) (*pb.DeleteUserRes, error) {
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

func (t TweetServiceServer) ReturnTweet(ctx context.Context, req *pb.ReturnTweetReq) (*pb.ReturnTweetRes, error) {
	return t.TweetClient.ReturnTweet(ctx, req)
}

func (t TweetServiceServer) PostTweet(ctx context.Context, req *pb.PostTweetReq) (*pb.PostTweetRes, error) {
	return t.TweetClient.PostTweet(ctx, req)
}

func InitGRPC(wg *sync.WaitGroup, mux *runtime.ServeMux) {
	defer wg.Done()
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	fmt.Println("Starting server on port: 50055")

	var opts []grpc.ServerOption
	opts = append(opts, grpc.UnaryInterceptor(authInterceptor))
	s := grpc.NewServer(opts...)
	ctx := context.Background()
	registerUserDataService(s, ctx, mux)
	registerAuthService(s, ctx, mux, config.Config.AuthServiceAddr)
	registerUserService(s, ctx, mux, config.Config.UserServiceAddr)
	registerTrendService(s, ctx, mux, config.Config.TrendServiceAddr)
	registerTweetService(s, ctx, mux, config.Config.TweetServiceAddr)

	reflection.Register(s)
	listener, err := net.Listen("tcp", ":50055")
	if err != nil {
		log.Fatalf("Unable to listen on port :50055: %v", err)
	}

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

func InitMux(wg *sync.WaitGroup, mux *runtime.ServeMux) {
	defer wg.Done()
	//handler := loggingMiddleware(mux)
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatalf("failed to start gateway server: %v", err)
	}
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Received request: %s %s", r.Method, r.URL.Path)
		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Printf("Failed to read request body: %v", err)
		} else {
			// Log the request body
			log.Printf("Request body: %s", string(body))
		}
		r.Body = io.NopCloser(bytes.NewBuffer(body))
		next.ServeHTTP(w, r)
		//log.Printf("Sent response: %d", w.(http.ResponseWriter))
	})
}

func registerUserDataService(s *grpc.Server, ctx context.Context, mux *runtime.ServeMux) {
	authConn, err := grpc.Dial(config.Config.AuthServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to dial authservice: %v", err)
	}
	userConn, err := grpc.Dial(config.Config.UserServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to dial userservice: %v", err)
	}
	tweetConn, err := grpc.Dial(config.Config.TweetServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to dial tweetservice: %v", err)
	}
	authClient := pb.NewAuthServiceClient(authConn)
	userClient := pb.NewUserServiceClient(userConn)
	tweetClient := pb.NewTweetServiceClient(tweetConn)
	userDataServer := &UserDataServiceServer{
		AuthClient:  authClient,
		UserClient:  userClient,
		TweetClient: tweetClient,
	}
	pb.RegisterUserDataServiceServer(s, userDataServer)
	err = pb.RegisterUserDataServiceHandlerServer(ctx, mux, userDataServer)
	if err != nil {
		log.Fatalf("failed to register UserDataService handler: %v", err)
	}
}
func registerAuthService(s *grpc.Server, ctx context.Context, mux *runtime.ServeMux, AuthServiceAddr string) {
	authConn, err := grpc.Dial(AuthServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to dial authservice: %v", err)
	}
	authClient := pb.NewAuthServiceClient(authConn)
	authServer := &AuthServiceServer{AuthClient: authClient}
	pb.RegisterAuthServiceServer(s, authServer)
	err = pb.RegisterAuthServiceHandlerServer(ctx, mux, authServer)
	if err != nil {
		log.Fatalf("failed to register authservice handler: %v", err)
	}
}
func registerUserService(s *grpc.Server, ctx context.Context, mux *runtime.ServeMux, UserServiceAddr string) {
	userConn, err := grpc.Dial(UserServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to dial userservice: %v", err)
	}
	userClient := pb.NewUserServiceClient(userConn)
	userServer := &UserServiceServer{UserClient: userClient}
	pb.RegisterUserServiceServer(s, userServer)
	err = pb.RegisterUserServiceHandlerServer(ctx, mux, userServer)
	if err != nil {
		log.Fatalf("failed to register userservice handler: %v", err)
	}
}
func registerTrendService(s *grpc.Server, ctx context.Context, mux *runtime.ServeMux, TrendServiceAddr string) {
	trendConn, err := grpc.Dial(TrendServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to dial trendservice: %v", err)
	}
	trendClient := pb.NewTrendServiceClient(trendConn)
	trendServer := &TrendServiceServer{TrendClient: trendClient}
	pb.RegisterTrendServiceServer(s, trendServer)
	err = pb.RegisterTrendServiceHandlerServer(ctx, mux, trendServer)
	if err != nil {
		log.Fatalf("failed to register trendservice handler: %v", err)
	}
}
func registerTweetService(s *grpc.Server, ctx context.Context, mux *runtime.ServeMux, TweetServiceAddr string) {
	tweetConn, err := grpc.Dial(TweetServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to dial tweetservice: %v", err)
	}
	tweetClient := pb.NewTweetServiceClient(tweetConn)
	tweetServer := &TweetServiceServer{TweetClient: tweetClient}
	pb.RegisterTweetServiceServer(s, tweetServer)
	err = pb.RegisterTweetServiceHandlerServer(ctx, mux, tweetServer)
	if err != nil {
		log.Fatalf("failed to register tweetservice handler: %v", err)
	}
}

func authInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	// Pre-processing logic before invoking the RPC method
	log.Println("Interceptor: Before invoking the RPC method")

	excludedMethods := []string{
		"/proto.AuthService/Register",
		"/proto.AuthService/Login",
		"/proto.AuthService/Validate",
	}
	log.Println(info.FullMethod)
	for _, method := range excludedMethods {
		if info.FullMethod == method {
			return handler(ctx, req)
		}
	}

	token, err := extractToken(ctx)
	if err != nil || token.GetToken() == "" {
		return nil, status.Errorf(codes.Unauthenticated, "Invalid token")
	}
	validateReq := &pb.ValidateReq{Token: token.GetToken()}

	AuthServiceAddr := config.Config.AuthServiceAddr
	authConn, err := grpc.Dial(AuthServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to dial authservice: %v", err)
	}
	authClient := pb.NewAuthServiceClient(authConn)
	server := &AuthServiceServer{AuthClient: authClient}
	validateRes, err := server.AuthClient.Validate(ctx, validateReq)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "failed to validate token: %v", err)
	}
	log.Printf("Validation response: %v", validateRes)
	log.Printf(validateReq.Token)

	if info.FullMethod == "/proto.TweetService/PostTweet" {
		tweetReq, ok := req.(*pb.PostTweetReq)
		if !ok {
			log.Println("Invalid request type for PostTweet")
		}
		tweetValue := tweetReq.GetTweet()
		validatedTweet := &pb.Tweet{
			UserID:   validateRes.GetUserid(),
			Username: tweetValue.GetUsername(),
			Body:     tweetValue.GetBody(),
		}
		req = &pb.PostTweetReq{Tweet: validatedTweet}
	}

	// Invoke the RPC method
	resp, err := handler(ctx, req)
	if err != nil {
		log.Println(err)
		return nil, status.Error(codes.Unknown, "Token validation failed")
	}
	return resp, err
}

func extractToken(ctx context.Context) (token *pb.ValidateReq, err error) {
	log.Println("extracting token")
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.Internal, "failed to extract metadata")
	}
	authHeaders, ok := md["authorization"]
	if !ok || len(authHeaders) == 0 {
		return nil, status.Errorf(codes.Unauthenticated, "missing authorization header")
	}
	tokenString := strings.TrimPrefix(authHeaders[0], "Bearer ")
	if tokenString == authHeaders[0] {
		return nil, status.Error(codes.Unauthenticated, "Invalid token format")
	}
	token = &pb.ValidateReq{Token: tokenString}
	log.Println(tokenString)
	return token, nil
}
