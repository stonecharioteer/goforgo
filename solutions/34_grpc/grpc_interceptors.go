// grpc_interceptors.go
// Learn gRPC interceptors for logging, authentication, metrics, and middleware patterns

package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// Request and response types

// EchoRequest represents an echo request
type EchoRequest struct {
	Message  string `json:"message"`
	Username string `json:"username"`
}

// EchoResponse represents an echo response
type EchoResponse struct {
	Message     string `json:"message"`
	ProcessedBy string `json:"processed_by"`
	ProcessedAt string `json:"processed_at"`
}

// SecureRequest represents a request that requires authentication
type SecureRequest struct {
	Data      string `json:"data"`
	RequestId string `json:"request_id"`
}

// SecureResponse represents a secure response
type SecureResponse struct {
	Result         string        `json:"result"`
	Authenticated  bool          `json:"authenticated"`
	ProcessingTime time.Duration `json:"processing_time"`
}

// Service interface
type InterceptorServiceServer interface {
	Echo(context.Context, *EchoRequest) (*EchoResponse, error)
	SecureOperation(context.Context, *SecureRequest) (*SecureResponse, error)
}

// Service implementation
type interceptorServiceImpl struct{}

func main() {
	fmt.Println("=== gRPC Interceptors & Middleware ===")
	
	// Start server with interceptors
	go startServerWithInterceptors()
	
	// Wait for server startup
	time.Sleep(2 * time.Second)
	
	// Test interceptors
	fmt.Println("Testing gRPC interceptors...")
	
	fmt.Println("1. Testing basic echo with logging interceptor...")
	if err := testEchoWithLogging(); err != nil {
		log.Printf("Echo test error: %v", err)
	}
	
	fmt.Println("\n2. Testing authentication interceptor...")
	if err := testAuthenticationInterceptor(); err != nil {
		log.Printf("Auth test error: %v", err)
	}
	
	fmt.Println("\n3. Testing metrics interceptor...")
	if err := testMetricsInterceptor(); err != nil {
		log.Printf("Metrics test error: %v", err)
	}
	
	fmt.Println("\n4. Testing client interceptors...")
	if err := testClientInterceptors(); err != nil {
		log.Printf("Client interceptor test error: %v", err)
	}
	
	fmt.Println("\nInterceptor demo completed!")
}

// Start server with interceptors
func startServerWithInterceptors() {
	// Create listener
	lis, err := net.Listen("tcp", ":50053")
	if err != nil {
		log.Fatal("Failed to listen:", err)
	}
	
	// Create server with interceptors
	server := grpc.NewServer(
		createUnaryInterceptorChain(),
		createStreamInterceptorChain(),
	)
	
	// Register service
	service := newInterceptorService()
	registerInterceptorServiceServer(server, service)
	
	fmt.Println("Interceptor gRPC server starting on :50053")
	
	// Start serving
	if err := server.Serve(lis); err != nil {
		log.Fatal("Failed to serve:", err)
	}
}

// Create interceptor service
func newInterceptorService() *interceptorServiceImpl {
	return &interceptorServiceImpl{}
}

// Implement Echo method
func (s *interceptorServiceImpl) Echo(ctx context.Context, req *EchoRequest) (*EchoResponse, error) {
	// Process echo request
	response := &EchoResponse{
		Message:     fmt.Sprintf("Echo: %s (from %s)", req.Message, req.Username),
		ProcessedBy: "interceptor-service",
		ProcessedAt: time.Now().Format(time.RFC3339),
	}
	
	fmt.Printf("Processing echo request: %s\n", req.Message)
	return response, nil
}

// Implement SecureOperation method
func (s *interceptorServiceImpl) SecureOperation(ctx context.Context, req *SecureRequest) (*SecureResponse, error) {
	startTime := time.Now()
	
	// Extract user from context (set by auth interceptor)
	username := getFromContext(ctx, "username")
	
	// Process secure request
	result := fmt.Sprintf("Processed secure data: %s (by %s)", req.Data, username)
	
	response := &SecureResponse{
		Result:         result,
		Authenticated:  username != "",
		ProcessingTime: time.Since(startTime),
	}
	
	fmt.Printf("Processing secure operation for user: %s\n", username)
	return response, nil
}

// Logging interceptor
func loggingUnaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	start := time.Now()
	
	fmt.Printf("[LOGGING] Request started: %s\n", info.FullMethod)
	
	// Call handler
	resp, err := handler(ctx, req)
	
	// Calculate duration
	duration := time.Since(start)
	
	// Log completion
	if err != nil {
		fmt.Printf("[LOGGING] Request failed: %s (duration: %v, error: %v)\n", 
			info.FullMethod, duration, err)
	} else {
		fmt.Printf("[LOGGING] Request completed: %s (duration: %v)\n", 
			info.FullMethod, duration)
	}
	
	return resp, err
}

// Authentication interceptor
func authUnaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	// Skip auth for non-secure methods
	if !strings.Contains(info.FullMethod, "SecureOperation") {
		return handler(ctx, req)
	}
	
	// Extract metadata from context
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "missing metadata")
	}
	
	// Check authorization header
	authHeaders := md.Get("authorization")
	if len(authHeaders) == 0 {
		return nil, status.Errorf(codes.Unauthenticated, "missing authorization header")
	}
	
	// Validate token (simplified)
	token := authHeaders[0]
	if !strings.HasPrefix(token, "Bearer ") {
		return nil, status.Errorf(codes.Unauthenticated, "invalid authorization format")
	}
	
	// Extract and validate token
	actualToken := strings.TrimPrefix(token, "Bearer ")
	username := validateToken(actualToken)
	
	if username == "" {
		return nil, status.Errorf(codes.Unauthenticated, "invalid token")
	}
	
	// Add username to context
	ctx = addToContext(ctx, "username", username)
	
	fmt.Printf("[AUTH] Authenticated user: %s\n", username)
	
	return handler(ctx, req)
}

// Metrics interceptor
func metricsUnaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	start := time.Now()
	
	// Call handler
	resp, err := handler(ctx, req)
	
	// Record metrics
	duration := time.Since(start)
	method := info.FullMethod
	
	recordRequestCount(method)
	recordResponseTime(method, duration)
	
	if err != nil {
		recordErrorCount(method)
	}
	
	fmt.Printf("[METRICS] %s: duration=%v, error=%v\n", method, duration, err != nil)
	
	return resp, err
}

// Create unary interceptor chain
func createUnaryInterceptorChain() grpc.ServerOption {
	return grpc.ChainUnaryInterceptor(
		loggingUnaryInterceptor,
		metricsUnaryInterceptor,
		authUnaryInterceptor,
	)
}

// Create stream interceptor chain (simplified)
func createStreamInterceptorChain() grpc.ServerOption {
	return grpc.ChainStreamInterceptor(
		loggingStreamInterceptor,
	)
}

// Simple stream logging interceptor
func loggingStreamInterceptor(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	fmt.Printf("[STREAM LOGGING] Stream started: %s\n", info.FullMethod)
	
	err := handler(srv, stream)
	
	if err != nil {
		fmt.Printf("[STREAM LOGGING] Stream failed: %s (error: %v)\n", info.FullMethod, err)
	} else {
		fmt.Printf("[STREAM LOGGING] Stream completed: %s\n", info.FullMethod)
	}
	
	return err
}

// Validate token helper
func validateToken(token string) string {
	// Simple token validation (in production, use proper JWT/OAuth)
	validTokens := map[string]string{
		"admin123":  "admin",
		"user456":   "john_doe",
		"guest789":  "guest_user",
	}
	
	return validTokens[token]
}

// Test functions

// Test echo with logging
func testEchoWithLogging() error {
	// Create client connection
	conn, err := grpc.Dial("localhost:50053", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}
	defer conn.Close()
	
	// Create client
	client := newInterceptorServiceClient(conn)
	
	// Call Echo method
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	
	resp, err := client.Echo(ctx, &EchoRequest{
		Message:  "Hello from client!",
		Username: "test_user",
	})
	if err != nil {
		return err
	}
	
	fmt.Printf("Echo response: %s (processed by: %s)\n", resp.Message, resp.ProcessedBy)
	return nil
}

// Test authentication interceptor
func testAuthenticationInterceptor() error {
	// Create client connection
	conn, err := grpc.Dial("localhost:50053", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}
	defer conn.Close()
	
	// Create client
	client := newInterceptorServiceClient(conn)
	
	// Test without authentication (should fail)
	ctx1, cancel1 := context.WithTimeout(context.Background(), 5*time.Second)
	_, err = client.SecureOperation(ctx1, &SecureRequest{
		Data:      "sensitive data",
		RequestId: "req-1",
	})
	cancel1()
	
	if err != nil {
		fmt.Printf("Expected auth error: %v\n", err)
	}
	
	// Test with valid authentication
	ctx2, cancel2 := context.WithTimeout(context.Background(), 5*time.Second)
	ctx2 = addAuthMetadata(ctx2, "admin123")
	defer cancel2()
	
	resp, err := client.SecureOperation(ctx2, &SecureRequest{
		Data:      "sensitive data",
		RequestId: "req-2",
	})
	if err != nil {
		return err
	}
	
	fmt.Printf("Secure response: %s (authenticated: %t)\n", resp.Result, resp.Authenticated)
	return nil
}

// Test metrics interceptor
func testMetricsInterceptor() error {
	// Create client and make multiple calls to generate metrics
	conn, err := grpc.Dial("localhost:50053", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}
	defer conn.Close()
	
	client := newInterceptorServiceClient(conn)
	
	// Make multiple requests to generate metrics
	for i := 0; i < 5; i++ {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		
		_, err := client.Echo(ctx, &EchoRequest{
			Message:  fmt.Sprintf("Metrics test message %d", i+1),
			Username: "metrics_user",
		})
		cancel()
		
		if err != nil {
			log.Printf("Echo error (expected for metrics): %v", err)
		}
		
		time.Sleep(500 * time.Millisecond)
	}
	
	fmt.Println("Metrics test completed - check server logs for metric recordings")
	return nil
}

// Test client interceptors
func testClientInterceptors() error {
	// Create connection with client interceptors
	conn, err := grpc.Dial("localhost:50053",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithChainUnaryInterceptor(clientLoggingInterceptor),
	)
	if err != nil {
		return err
	}
	defer conn.Close()
	
	client := newInterceptorServiceClient(conn)
	
	// Make request (will go through client interceptors)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	
	resp, err := client.Echo(ctx, &EchoRequest{
		Message:  "Client interceptor test",
		Username: "interceptor_user",
	})
	if err != nil {
		return err
	}
	
	fmt.Printf("Client interceptor response: %s\n", resp.Message)
	return nil
}

// Client interceptor for adding request ID
func clientLoggingInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	start := time.Now()
	
	fmt.Printf("[CLIENT] Request started: %s\n", method)
	
	// Add request ID to metadata
	requestID := generateRequestID()
	ctx = metadata.AppendToOutgoingContext(ctx, "request-id", requestID)
	
	// Call method
	err := invoker(ctx, method, req, reply, cc, opts...)
	
	// Log completion
	duration := time.Since(start)
	if err != nil {
		fmt.Printf("[CLIENT] Request failed: %s (duration: %v, error: %v)\n", method, duration, err)
	} else {
		fmt.Printf("[CLIENT] Request completed: %s (duration: %v)\n", method, duration)
	}
	
	return err
}

// Helper functions

func generateRequestID() string {
	return fmt.Sprintf("req_%d", time.Now().UnixNano())
}

func addToContext(ctx context.Context, key, value string) context.Context {
	return context.WithValue(ctx, key, value)
}

func getFromContext(ctx context.Context, key string) string {
	if value := ctx.Value(key); value != nil {
		if str, ok := value.(string); ok {
			return str
		}
	}
	return ""
}

func addAuthMetadata(ctx context.Context, token string) context.Context {
	return metadata.AppendToOutgoingContext(ctx, "authorization", "Bearer "+token)
}

// Metric recording functions (simplified)
func recordRequestCount(method string) {
	fmt.Printf("[METRIC] Request count incremented for %s\n", method)
}

func recordResponseTime(method string, duration time.Duration) {
	fmt.Printf("[METRIC] Response time recorded for %s: %v\n", method, duration)
}

func recordErrorCount(method string) {
	fmt.Printf("[METRIC] Error count incremented for %s\n", method)
}

// Client and service interfaces (simplified)
type InterceptorServiceClient interface {
	Echo(ctx context.Context, req *EchoRequest, opts ...grpc.CallOption) (*EchoResponse, error)
	SecureOperation(ctx context.Context, req *SecureRequest, opts ...grpc.CallOption) (*SecureResponse, error)
}

type interceptorServiceClient struct {
	cc *grpc.ClientConn
}

func newInterceptorServiceClient(conn *grpc.ClientConn) InterceptorServiceClient {
	return &interceptorServiceClient{cc: conn}
}

// Simplified client implementations (normally auto-generated)
func (c *interceptorServiceClient) Echo(ctx context.Context, req *EchoRequest, opts ...grpc.CallOption) (*EchoResponse, error) {
	// Placeholder implementation
	return &EchoResponse{
		Message:     fmt.Sprintf("Echo: %s", req.Message),
		ProcessedBy: "client-stub",
		ProcessedAt: time.Now().Format(time.RFC3339),
	}, nil
}

func (c *interceptorServiceClient) SecureOperation(ctx context.Context, req *SecureRequest, opts ...grpc.CallOption) (*SecureResponse, error) {
	// Check for auth metadata
	md, ok := metadata.FromOutgoingContext(ctx)
	authenticated := ok && len(md.Get("authorization")) > 0
	
	if !authenticated {
		return nil, status.Errorf(codes.Unauthenticated, "missing authorization")
	}
	
	return &SecureResponse{
		Result:         fmt.Sprintf("Processed: %s", req.Data),
		Authenticated:  authenticated,
		ProcessingTime: time.Millisecond * 100,
	}, nil
}

func registerInterceptorServiceServer(s *grpc.Server, srv InterceptorServiceServer) {
	fmt.Println("InterceptorService registered with gRPC server")
}