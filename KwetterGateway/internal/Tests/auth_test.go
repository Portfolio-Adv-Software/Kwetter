package Tests

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/Portfolio-Adv-Software/Kwetter/KwetterGateway/internal/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc"
	"net"
)

type AuthServiceServerMock struct {
	__.UnimplementedAuthServiceServer
}

type LoginMock struct {
	Email    string
	Password string
}

func (m *AuthServiceServerMock) Register(ctx context.Context, req *__.RegisterReq) (*__.RegisterRes, error) {
	if req.Email == "existing@example.com" {
		return nil, status.Errorf(codes.AlreadyExists, "Email is already registered")
	}

	// Simulate successful registration
	insertedUserID := primitive.NewObjectID()
	fmt.Println("Created user with id: " + insertedUserID.Hex())

	return &__.RegisterRes{
		Status: "Registration successful",
	}, nil
}

func (m *AuthServiceServerMock) Login(ctx context.Context, req *__.LoginReq) (*__.LoginRes, error) {
	if req.GetEmail() == "nonexisting@example.com" {
		return nil, status.Errorf(codes.NotFound, "User with that email not found")
	}

	// Simulate successful login
	user := LoginMock{
		Email:    req.Email,
		Password: req.Password,
	}
	fmt.Println(user.Email + " logged in.")

	token := "generated_token"

	return &__.LoginRes{
		Status:      "Login successful",
		AccessToken: token,
	}, nil
}

func TestRegister(t *testing.T) {
	// Create a new test gRPC server
	testServer := grpc.NewServer()

	// Create an instance of the AuthServiceServerMock
	authServer := &AuthServiceServerMock{}

	// Register the AuthServiceServerMock as the gRPC server implementation
	__.RegisterAuthServiceServer(testServer, authServer)

	// Start the test server in a separate goroutine
	lis, err := net.Listen("tcp", ":0")
	if err != nil {
		t.Fatalf("failed to start test server: %v", err)
	}
	go func() {
		if err := testServer.Serve(lis); err != nil {
			t.Fatalf("failed to start test server: %v", err)
		}
	}()

	// Create a new gRPC connection to the test server
	conn, err := grpc.Dial(lis.Addr().String(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("failed to connect to test server: %v", err)
	}
	defer conn.Close()

	// Create an instance of the AuthServiceClient using the test server connection
	authClient := __.NewAuthServiceClient(conn)

	// Create a new Gin router
	router := gin.Default()

	// Register the register route
	router.POST("/register", func(c *gin.Context) {
		// Convert the HTTP request body to a RegisterReq protobuf message
		var req __.RegisterReq
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Call the Register handler using the authClient
		res, err := authClient.Register(c, &req)
		if err != nil {
			// Handle the error returned by the Register function
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Return the RegisterRes protobuf message as JSON response
		c.JSON(http.StatusOK, res)
	})

	t.Run("successful registration", func(t *testing.T) {
		// Create a test request body with unique registration details
		body := &__.RegisterReq{
			Email:          "test@example.com",
			Password:       "password",
			DataPermission: true,
		}
		jsonBody, _ := json.Marshal(body)

		// Create a test HTTP request with the JSON body
		req, _ := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(jsonBody))
		req.Header.Set("Content-Type", "application/json")

		// Create a test HTTP response recorder
		res := httptest.NewRecorder()

		// Perform the request
		router.ServeHTTP(res, req)

		// Assert that the response has the expected status code
		assert.Equal(t, http.StatusOK, res.Code)

		// Parse the response body into a RegisterRes protobuf message
		var registerRes __.RegisterRes
		if err := json.Unmarshal(res.Body.Bytes(), &registerRes); err != nil {
			t.Fatalf("failed to unmarshal response: %v", err)
		}

		// Assert the status field of the RegisterRes message
		assert.Equal(t, "Registration successful", registerRes.Status)
	})

	t.Run("duplicate registration", func(t *testing.T) {
		// Create a test request body with duplicate registration details
		body := &__.RegisterReq{
			Email:          "existing@example.com",
			Password:       "password",
			DataPermission: true,
		}
		jsonBody, _ := json.Marshal(body)

		// Create a test HTTP request with the JSON body
		req, _ := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(jsonBody))
		req.Header.Set("Content-Type", "application/json")

		// Create a test HTTP response recorder
		res := httptest.NewRecorder()

		// Perform the request
		router.ServeHTTP(res, req)

		// Assert that the response has the expected status code
		assert.Equal(t, http.StatusInternalServerError, res.Code)

		// Parse the response body into a JSON object to access the "error" field
		var errRes struct {
			Error string `json:"error"`
		}
		if err := json.Unmarshal(res.Body.Bytes(), &errRes); err != nil {
			t.Fatalf("failed to unmarshal error response: %v", err)
		}

		// Assert the error message received due to duplicate registration
		assert.Contains(t, errRes.Error, "Email is already registered")
	})
}

func TestLogin(t *testing.T) {
	// Create a new test gRPC server
	testServer := grpc.NewServer()

	// Create an instance of the AuthServiceServerMock
	authServer := &AuthServiceServerMock{}

	// Register the AuthServiceServerMock as the gRPC server implementation
	__.RegisterAuthServiceServer(testServer, authServer)

	// Start the test server in a separate goroutine
	lis, err := net.Listen("tcp", ":0")
	if err != nil {
		t.Fatalf("failed to start test server: %v", err)
	}
	go func() {
		if err := testServer.Serve(lis); err != nil {
			t.Fatalf("failed to start test server: %v", err)
		}
	}()

	// Create a new gRPC connection to the test server
	conn, err := grpc.Dial(lis.Addr().String(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("failed to connect to test server: %v", err)
	}
	defer conn.Close()

	// Create an instance of the AuthServiceClient using the test server connection
	authClient := __.NewAuthServiceClient(conn)

	// Create a new Gin router
	router := gin.Default()

	// Register the login route
	router.POST("/login", func(c *gin.Context) {
		// Convert the HTTP request body to a LoginReq protobuf message
		var req __.LoginReq
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Call the Login handler
		res, err := authClient.Login(c, &req)
		if err != nil {
			// Handle the error returned by the Login function
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Return the LoginRes protobuf message as JSON response
		c.JSON(http.StatusOK, res)
	})

	t.Run("successful login", func(t *testing.T) {
		// Create a test request body with valid login details
		body := &__.LoginReq{
			Email:    "user@example.com",
			Password: "password",
		}
		jsonBody, _ := json.Marshal(body)

		// Create a test HTTP request with the JSON body
		req, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(jsonBody))
		req.Header.Set("Content-Type", "application/json")

		// Create a test HTTP response recorder
		res := httptest.NewRecorder()

		// Perform the request
		router.ServeHTTP(res, req)

		// Assert that the response has the expected status code
		assert.Equal(t, http.StatusOK, res.Code)

		// Parse the response body into a LoginRes protobuf message
		var loginRes __.LoginRes
		if err := json.Unmarshal(res.Body.Bytes(), &loginRes); err != nil {
			t.Fatalf("failed to unmarshal response: %v", err)
		}

		// Assert the status field of the LoginRes message
		assert.Equal(t, "Login successful", loginRes.Status)
	})

	t.Run("invalid login", func(t *testing.T) {
		// Create a test request body with invalid login details
		body := &__.LoginReq{
			Email:    "nonexisting@example.com",
			Password: "password",
		}
		jsonBody, _ := json.Marshal(body)

		// Create a test HTTP request with the JSON body
		req, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(jsonBody))
		req.Header.Set("Content-Type", "application/json")

		// Create a test HTTP response recorder
		res := httptest.NewRecorder()

		// Perform the request
		router.ServeHTTP(res, req)

		// Assert that the response has the expected status code
		assert.Equal(t, http.StatusInternalServerError, res.Code)

		// Parse the response body into a JSON object to access the "error" field
		var errRes struct {
			Error string `json:"error"`
		}
		if err := json.Unmarshal(res.Body.Bytes(), &errRes); err != nil {
			t.Fatalf("failed to unmarshal error response: %v", err)
		}

		// Assert the error message received due to invalid login
		assert.Contains(t, errRes.Error, "User with that email not found")
	})
}
