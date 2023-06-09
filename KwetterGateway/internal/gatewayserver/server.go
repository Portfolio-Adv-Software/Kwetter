package gatewayserver

import (
	"bytes"
	"fmt"
	"github.com/Portfolio-Adv-Software/Kwetter/KwetterGateway/internal/config"
	"github.com/Portfolio-Adv-Software/Kwetter/KwetterGateway/internal/proto"
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
	__.UnimplementedUserDataServiceServer
	AuthClient  __.AuthServiceClient
	UserClient  __.UserServiceClient
	TweetClient __.TweetServiceClient
}

func (u UserDataServiceServer) GetAllUserData(ctx context.Context, req *__.GetAllUserDataReq) (*__.GetAllUserDataRes, error) {
	userId := req.GetUserId()
	authReq := &__.GetDataReq{UserId: userId}
	userReq := &__.GetUserReq{UserID: userId}
	tweetReq := &__.ReturnAllReq{UserId: userId}
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
	res := &__.AllUserData{
		User:      userRes.GetUser(),
		AuthData:  authRes.GetAuthData(),
		TweetData: tweetRes.GetTweetData(),
	}
	return &__.GetAllUserDataRes{AllUserData: res}, nil
}

type AuthServiceServer struct {
	__.UnimplementedAuthServiceServer
	AuthClient __.AuthServiceClient
}

func (a AuthServiceServer) Register(ctx context.Context, req *__.RegisterReq) (*__.RegisterRes, error) {
	return a.AuthClient.Register(ctx, req)
	log.Println("Received Register request")
	// You can add more log statements to display the request details if needed
	log.Printf("Register request: %+v", req)

	// Delegate the request to the underlying AuthClient
	res, err := a.AuthClient.Register(ctx, req)

	if err != nil {
		log.Printf("Register error: %v", err)
	} else {
		log.Println("Register response:")
		// You can add more log statements to display the response details if needed
		log.Printf("%+v", res)
	}

	return res, err
}

func (a AuthServiceServer) Login(ctx context.Context, req *__.LoginReq) (*__.LoginRes, error) {
	return a.AuthClient.Login(ctx, req)
}
func (a AuthServiceServer) Validate(ctx context.Context, req *__.ValidateReq) (*__.ValidateRes, error) {
	return a.AuthClient.Validate(ctx, req)
}

type UserServiceServer struct {
	__.UnimplementedUserServiceServer
	UserClient __.UserServiceClient
}

func (u UserServiceServer) UpdateUser(ctx context.Context, req *__.UpdateUserReq) (*__.UpdateUserRes, error) {
	return u.UserClient.UpdateUser(ctx, req)
}
func (u UserServiceServer) DeleteUser(ctx context.Context, req *__.DeleteUserReq) (*__.DeleteUserRes, error) {
	return u.UserClient.DeleteUser(ctx, req)
}

type TrendServiceServer struct {
	__.UnimplementedTrendServiceServer
	TrendClient __.TrendServiceClient
}

func (t TrendServiceServer) GetTrend(ctx context.Context, req *__.GetTrendReq) (*__.GetTrendRes, error) {
	return t.TrendClient.GetTrend(ctx, req)
}

type TweetServiceServer struct {
	__.UnimplementedTweetServiceServer
	TweetClient __.TweetServiceClient
}

func (t TweetServiceServer) ReturnTweet(ctx context.Context, req *__.ReturnTweetReq) (*__.ReturnTweetRes, error) {
	return t.TweetClient.ReturnTweet(ctx, req)
}
func (t TweetServiceServer) PostTweet(ctx context.Context, req *__.PostTweetReq) (*__.PostTweetRes, error) {
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
	authClient := __.NewAuthServiceClient(authConn)
	userClient := __.NewUserServiceClient(userConn)
	tweetClient := __.NewTweetServiceClient(tweetConn)
	userDataServer := &UserDataServiceServer{
		AuthClient:  authClient,
		UserClient:  userClient,
		TweetClient: tweetClient,
	}

	__.RegisterUserDataServiceServer(s, userDataServer)
	err = __.RegisterUserDataServiceHandlerServer(ctx, mux, userDataServer)
	if err != nil {
		log.Fatalf("failed to register UserDataService handler: %v", err)
	}
}
func registerAuthService(s *grpc.Server, ctx context.Context, mux *runtime.ServeMux, AuthServiceAddr string) {
	authConn, err := grpc.Dial(AuthServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to dial authservice: %v", err)
	}
	authClient := __.NewAuthServiceClient(authConn)
	authServer := &AuthServiceServer{AuthClient: authClient}
	__.RegisterAuthServiceServer(s, authServer)
	err = __.RegisterAuthServiceHandlerServer(ctx, mux, authServer)
	if err != nil {
		log.Fatalf("failed to register authservice handler: %v", err)
	}
}
func registerUserService(s *grpc.Server, ctx context.Context, mux *runtime.ServeMux, UserServiceAddr string) {
	userConn, err := grpc.Dial(UserServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to dial userservice: %v", err)
	}
	userClient := __.NewUserServiceClient(userConn)
	userServer := &UserServiceServer{UserClient: userClient}
	__.RegisterUserServiceServer(s, userServer)
	err = __.RegisterUserServiceHandlerServer(ctx, mux, userServer)
	if err != nil {
		log.Fatalf("failed to register userservice handler: %v", err)
	}
}
func registerTrendService(s *grpc.Server, ctx context.Context, mux *runtime.ServeMux, TrendServiceAddr string) {
	trendConn, err := grpc.Dial(TrendServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to dial trendservice: %v", err)
	}
	trendClient := __.NewTrendServiceClient(trendConn)
	trendServer := &TrendServiceServer{TrendClient: trendClient}
	__.RegisterTrendServiceServer(s, trendServer)
	err = __.RegisterTrendServiceHandlerServer(ctx, mux, trendServer)
	if err != nil {
		log.Fatalf("failed to register trendservice handler: %v", err)
	}
}
func registerTweetService(s *grpc.Server, ctx context.Context, mux *runtime.ServeMux, TweetServiceAddr string) {
	tweetConn, err := grpc.Dial(TweetServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to dial tweetservice: %v", err)
	}
	tweetClient := __.NewTweetServiceClient(tweetConn)
	tweetServer := &TweetServiceServer{TweetClient: tweetClient}
	__.RegisterTweetServiceServer(s, tweetServer)
	err = __.RegisterTweetServiceHandlerServer(ctx, mux, tweetServer)
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
		return nil, err
	}
	validateReq := &__.ValidateReq{Token: token.GetToken()}
	AuthServiceAddr := config.Config.AuthServiceAddr
	authConn, err := grpc.Dial(AuthServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to dial authservice: %v", err)
	}
	authClient := __.NewAuthServiceClient(authConn)
	server := &AuthServiceServer{AuthClient: authClient}
	validateRes, err := server.AuthClient.Validate(ctx, validateReq)
	if err != nil {
		return nil, err
	}
	log.Printf("Validation response: %v", validateRes)
	log.Printf(validateReq.Token)
	if info.FullMethod == "/proto.TweetService/PostTweet" {
		tweetReq, ok := req.(*__.PostTweetReq)
		if !ok {
			log.Println("Invalid request type for PostTweet")
		}
		tweetValue := tweetReq.GetTweet()
		if tweetValue.GetUserID() != validateRes.GetUserid() {
			return nil, status.Errorf(codes.Unauthenticated, "UserID mismatch")
		}
		validatedTweet := &__.Tweet{
			UserID:   validateRes.GetUserid(),
			Username: tweetValue.GetUsername(),
			Body:     tweetValue.GetBody(),
		}
		req = &__.PostTweetReq{Tweet: validatedTweet}
	}
	// Invoke the RPC method
	resp, err := handler(ctx, req)
	if err != nil {
		log.Println(err)
		return nil, status.Error(codes.Unknown, "Something went wrong")
	}
	return resp, err
}
func extractToken(ctx context.Context) (token *__.ValidateReq, err error) {
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
	token = &__.ValidateReq{Token: tokenString}
	log.Println(tokenString)
	return token, nil
}
