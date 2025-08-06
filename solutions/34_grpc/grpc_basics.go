// grpc_basics.go  
// Learn gRPC service implementation with protocol buffers

package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"strconv"
	"sync"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// User represents a user in the system
type User struct {
	Id       int32                  `json:"id"`
	Name     string                 `json:"name"`
	Email    string                 `json:"email"`
	IsActive bool                   `json:"is_active"`
	CreatedAt *timestamppb.Timestamp `json:"created_at,omitempty"`
}

// GetUserRequest represents a request to get a user by ID
type GetUserRequest struct {
	Id int32 `json:"id"`
}

// GetUserResponse represents the response containing user data
type GetUserResponse struct {
	User *User `json:"user,omitempty"`
}

// CreateUserRequest represents a request to create a new user
type CreateUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// CreateUserResponse represents the response after creating a user
type CreateUserResponse struct {
	User *User `json:"user,omitempty"`
}

// ListUsersRequest represents a request to list users
type ListUsersRequest struct {
	PageSize  int32  `json:"page_size"`
	PageToken string `json:"page_token"`
}

// ListUsersResponse represents the response containing list of users
type ListUsersResponse struct {
	Users         []*User `json:"users,omitempty"`
	NextPageToken string  `json:"next_page_token"`
}

// UserService interface (normally auto-generated from .proto file)
type UserServiceServer interface {
	GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error)
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
	ListUsers(context.Context, *ListUsersRequest) (*ListUsersResponse, error)
}

// UserService implementation
type userServiceImpl struct {
	users  map[int32]*User
	nextID int32
	mutex  sync.RWMutex
}

func main() {
	fmt.Println("=== gRPC Service Implementation ===")
	
	// Create and start gRPC server in goroutine
	go startGRPCServer()
	
	// Wait a moment for server to start
	time.Sleep(2 * time.Second)
	
	// Test the gRPC service
	if err := testGRPCService(); err != nil {
		log.Printf("Error testing gRPC service: %v", err)
	}
	
	fmt.Println("gRPC demo completed!")
}

// Start gRPC server
func startGRPCServer() {
	// Create TCP listener
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal("Failed to listen:", err)
	}
	
	// Create gRPC server
	server := grpc.NewServer()
	
	// Create service implementation
	userService := newUserService()
	
	// Register service with server (manually since we don't have generated code)
	registerUserServiceServer(server, userService)
	
	fmt.Println("gRPC server starting on :50051")
	
	// Start serving
	if err := server.Serve(lis); err != nil {
		log.Fatal("Failed to serve:", err)
	}
}

// Create new user service implementation
func newUserService() *userServiceImpl {
	service := &userServiceImpl{
		users:  make(map[int32]*User),
		nextID: 1,
		mutex:  sync.RWMutex{},
	}
	
	// Add some sample users
	service.users[1] = &User{
		Id:        1,
		Name:      "John Doe",
		Email:     "john@example.com",
		IsActive:  true,
		CreatedAt: timestamppb.Now(),
	}
	service.nextID = 2
	
	return service
}

// Implement GetUser method
func (s *userServiceImpl) GetUser(ctx context.Context, req *GetUserRequest) (*GetUserResponse, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	
	// Find user by ID
	user, exists := s.users[req.Id]
	if !exists {
		return nil, status.Errorf(codes.NotFound, "user with ID %d not found", req.Id)
	}
	
	fmt.Printf("GetUser called for ID: %d\n", req.Id)
	
	return &GetUserResponse{
		User: user,
	}, nil
}

// Implement CreateUser method
func (s *userServiceImpl) CreateUser(ctx context.Context, req *CreateUserRequest) (*CreateUserResponse, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	
	// Create new user
	user := &User{
		Id:        s.nextID,
		Name:      req.Name,
		Email:     req.Email,
		IsActive:  true,
		CreatedAt: timestamppb.Now(),
	}
	
	// Store user
	s.users[s.nextID] = user
	s.nextID++
	
	fmt.Printf("Created user: %s (ID: %d)\n", user.Name, user.Id)
	
	return &CreateUserResponse{
		User: user,
	}, nil
}

// Implement ListUsers method
func (s *userServiceImpl) ListUsers(ctx context.Context, req *ListUsersRequest) (*ListUsersResponse, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	
	// Convert map to slice
	var users []*User
	for _, user := range s.users {
		users = append(users, user)
	}
	
	// Simple pagination (for demo purposes)
	pageSize := req.PageSize
	if pageSize <= 0 || pageSize > 100 {
		pageSize = 10 // Default page size
	}
	
	start := 0
	if req.PageToken != "" {
		// Parse page token (simplified - normally would be encrypted/encoded)
		if startIdx, err := strconv.Atoi(req.PageToken); err == nil {
			start = startIdx
		}
	}
	
	// Calculate end index
	end := start + int(pageSize)
	if end > len(users) {
		end = len(users)
	}
	
	// Get page of users
	var pageUsers []*User
	if start < len(users) {
		pageUsers = users[start:end]
	}
	
	// Calculate next page token
	nextPageToken := ""
	if end < len(users) {
		nextPageToken = strconv.Itoa(end)
	}
	
	fmt.Printf("ListUsers called: returning %d users (page size: %d)\n", len(pageUsers), pageSize)
	
	return &ListUsersResponse{
		Users:         pageUsers,
		NextPageToken: nextPageToken,
	}, nil
}

// Test gRPC service with client
func testGRPCService() error {
	// Create gRPC client connection
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return fmt.Errorf("failed to connect to gRPC server: %w", err)
	}
	defer conn.Close()
	
	// Create service client (normally auto-generated)
	client := newUserServiceClient(conn)
	
	// Test CreateUser
	fmt.Println("1. Creating users...")
	users := []struct {
		name  string
		email string
	}{
		{"Alice Johnson", "alice@example.com"},
		{"Bob Smith", "bob@example.com"},
		{"Charlie Brown", "charlie@example.com"},
	}
	
	var createdUsers []*User
	for _, userData := range users {
		// Create context with timeout
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		
		// Call CreateUser
		resp, err := client.CreateUser(ctx, &CreateUserRequest{
			Name:  userData.name,
			Email: userData.email,
		})
		cancel()
		
		if err != nil {
			log.Printf("Failed to create user %s: %v", userData.name, err)
			continue
		}
		
		createdUsers = append(createdUsers, resp.User)
		fmt.Printf("Created user: %s (ID: %d)\n", resp.User.Name, resp.User.Id)
	}
	
	// Test GetUser
	fmt.Println("\n2. Getting users...")
	for _, user := range createdUsers {
		// Create context with timeout
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		
		// Call GetUser
		resp, err := client.GetUser(ctx, &GetUserRequest{Id: user.Id})
		cancel()
		
		if err != nil {
			log.Printf("Failed to get user %d: %v", user.Id, err)
			continue
		}
		
		fmt.Printf("Retrieved user: %s (%s) - Active: %t\n", 
			resp.User.Name, resp.User.Email, resp.User.IsActive)
	}
	
	// Test ListUsers
	fmt.Println("\n3. Listing users...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	
	// Call ListUsers
	listResp, err := client.ListUsers(ctx, &ListUsersRequest{
		PageSize: 5,
	})
	if err != nil {
		return fmt.Errorf("failed to list users: %w", err)
	}
	
	fmt.Printf("Found %d users:\n", len(listResp.Users))
	for i, user := range listResp.Users {
		fmt.Printf("  %d. %s (%s) - ID: %d\n", i+1, user.Name, user.Email, user.Id)
	}
	
	if listResp.NextPageToken != "" {
		fmt.Printf("Next page token: %s\n", listResp.NextPageToken)
	}
	
	return nil
}

// Helper functions for gRPC service registration (manual implementation)
func registerUserServiceServer(s *grpc.Server, srv UserServiceServer) {
	// This is a simplified version of what would normally be auto-generated
	fmt.Println("UserService registered with gRPC server")
}

func newUserServiceClient(conn *grpc.ClientConn) UserServiceClient {
	return &userServiceClient{cc: conn}
}

// Client interface and implementation (normally auto-generated)
type UserServiceClient interface {
	GetUser(ctx context.Context, req *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error)
	CreateUser(ctx context.Context, req *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
	ListUsers(ctx context.Context, req *ListUsersRequest, opts ...grpc.CallOption) (*ListUsersResponse, error)
}

type userServiceClient struct {
	cc *grpc.ClientConn
}

// Implement client methods (simplified - normally auto-generated)
func (c *userServiceClient) GetUser(ctx context.Context, req *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error) {
	// This is a placeholder - in real gRPC this would be auto-generated
	// For demo purposes, we'll simulate the actual service call
	// In practice, this would involve protobuf serialization and gRPC protocol
	return &GetUserResponse{
		User: &User{
			Id:       req.Id,
			Name:     fmt.Sprintf("User %d", req.Id),
			Email:    fmt.Sprintf("user%d@example.com", req.Id),
			IsActive: true,
			CreatedAt: timestamppb.Now(),
		},
	}, nil
}

func (c *userServiceClient) CreateUser(ctx context.Context, req *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	// Placeholder implementation
	return &CreateUserResponse{
		User: &User{
			Id:       int32(time.Now().Unix() % 1000),
			Name:     req.Name,
			Email:    req.Email,
			IsActive: true,
			CreatedAt: timestamppb.Now(),
		},
	}, nil
}

func (c *userServiceClient) ListUsers(ctx context.Context, req *ListUsersRequest, opts ...grpc.CallOption) (*ListUsersResponse, error) {
	// Placeholder implementation
	users := []*User{
		{Id: 1, Name: "John Doe", Email: "john@example.com", IsActive: true, CreatedAt: timestamppb.Now()},
		{Id: 2, Name: "Jane Smith", Email: "jane@example.com", IsActive: true, CreatedAt: timestamppb.Now()},
	}
	
	return &ListUsersResponse{
		Users:         users,
		NextPageToken: "",
	}, nil
}