// grpc_streaming.go
// Learn gRPC streaming patterns: server streaming, client streaming, and bidirectional streaming

package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net"
	"strconv"
	"sync"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// Message types for streaming

// LogEntry represents a log entry in the system
type LogEntry struct {
	Id        int64     `json:"id"`
	Level     string    `json:"level"`
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
	Service   string    `json:"service"`
}

// LogRequest represents a request to get logs
type LogRequest struct {
	Service   string    `json:"service"`
	Level     string    `json:"level"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
}

// MetricData represents a metric data point
type MetricData struct {
	Name      string            `json:"name"`
	Value     float64           `json:"value"`
	Timestamp time.Time         `json:"timestamp"`
	Labels    map[string]string `json:"labels"`
}

// AggregatedMetric represents aggregated metric result
type AggregatedMetric struct {
	Name    string  `json:"name"`
	Count   int64   `json:"count"`
	Sum     float64 `json:"sum"`
	Average float64 `json:"average"`
	Min     float64 `json:"min"`
	Max     float64 `json:"max"`
}

// ChatMessage represents a chat message
type ChatMessage struct {
	Username  string    `json:"username"`
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
	Room      string    `json:"room"`
}

// Streaming service interface
type StreamingServiceServer interface {
	StreamLogs(req *LogRequest, stream StreamingService_StreamLogsServer) error
	SendMetrics(stream StreamingService_SendMetricsServer) error
	ChatStream(stream StreamingService_ChatStreamServer) error
}

// Stream interfaces (normally auto-generated)
type StreamingService_StreamLogsServer interface {
	Send(*LogEntry) error
}

type StreamingService_SendMetricsServer interface {
	Recv() (*MetricData, error)
	SendAndClose(*map[string]*AggregatedMetric) error
}

type StreamingService_ChatStreamServer interface {
	Send(*ChatMessage) error
	Recv() (*ChatMessage, error)
}

// Streaming service implementation
type streamingServiceImpl struct {
	logs        []LogEntry
	metrics     chan MetricData
	chatClients map[string]StreamingService_ChatStreamServer
	mutex       sync.RWMutex
	logID       int64
}

func main() {
	fmt.Println("=== gRPC Streaming Patterns ===")
	
	// Start gRPC server in goroutine
	go startStreamingServer()
	
	// Wait for server to start
	time.Sleep(2 * time.Second)
	
	fmt.Println("Testing gRPC streaming...")
	
	// Test server streaming
	fmt.Println("1. Testing server streaming (log streaming)...")
	if err := testServerStreaming(); err != nil {
		log.Printf("Server streaming error: %v", err)
	}
	
	// Test client streaming
	fmt.Println("\n2. Testing client streaming (metric aggregation)...")
	if err := testClientStreaming(); err != nil {
		log.Printf("Client streaming error: %v", err)
	}
	
	// Test bidirectional streaming
	fmt.Println("\n3. Testing bidirectional streaming (chat)...")
	if err := testBidirectionalStreaming(); err != nil {
		log.Printf("Bidirectional streaming error: %v", err)
	}
	
	fmt.Println("\nStreaming demo completed!")
}

// Start streaming gRPC server
func startStreamingServer() {
	// Create TCP listener on port 50052
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatal("Failed to listen:", err)
	}
	
	// Create gRPC server
	server := grpc.NewServer()
	
	// Create streaming service
	streamingService := newStreamingService()
	
	// Register service
	registerStreamingServiceServer(server, streamingService)
	
	fmt.Println("Streaming gRPC server starting on :50052")
	
	// Start serving
	if err := server.Serve(lis); err != nil {
		log.Fatal("Failed to serve:", err)
	}
}

// Create new streaming service
func newStreamingService() *streamingServiceImpl {
	service := &streamingServiceImpl{
		logs:        make([]LogEntry, 0),
		metrics:     make(chan MetricData, 100),
		chatClients: make(map[string]StreamingService_ChatStreamServer),
		mutex:       sync.RWMutex{},
		logID:       1,
	}
	
	// Generate sample logs
	service.generateInitialLogs()
	
	// Start background log generator
	go service.generateLogs()
	
	return service
}

// Generate initial sample logs
func (s *streamingServiceImpl) generateInitialLogs() {
	services := []string{"auth-service", "user-service", "order-service"}
	levels := []string{"INFO", "WARN", "ERROR", "DEBUG"}
	messages := []string{"Request processed", "Database connected", "Cache miss", "Validation failed"}
	
	for i := 0; i < 20; i++ {
		s.logs = append(s.logs, LogEntry{
			Id:        s.logID,
			Level:     levels[rand.Intn(len(levels))],
			Message:   messages[rand.Intn(len(messages))],
			Timestamp: time.Now().Add(-time.Duration(i) * time.Minute),
			Service:   services[rand.Intn(len(services))],
		})
		s.logID++
	}
}

// Generate logs continuously (background task)
func (s *streamingServiceImpl) generateLogs() {
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()
	
	services := []string{"auth-service", "user-service", "order-service"}
	levels := []string{"INFO", "WARN", "ERROR", "DEBUG"}
	messages := []string{"Request processed", "Database connected", "Cache miss", "Validation failed"}
	
	for range ticker.C {
		// Generate random log entry
		logEntry := LogEntry{
			Id:        s.logID,
			Level:     levels[rand.Intn(len(levels))],
			Message:   messages[rand.Intn(len(messages))],
			Timestamp: time.Now(),
			Service:   services[rand.Intn(len(services))],
		}
		
		s.mutex.Lock()
		s.logs = append(s.logs, logEntry)
		s.logID++
		s.mutex.Unlock()
		
		// Keep only recent logs (limit to 100)
		s.mutex.Lock()
		if len(s.logs) > 100 {
			s.logs = s.logs[len(s.logs)-100:]
		}
		s.mutex.Unlock()
	}
}

// Implement server streaming (StreamLogs)
func (s *streamingServiceImpl) StreamLogs(req *LogRequest, stream StreamingService_StreamLogsServer) error {
	fmt.Printf("StreamLogs started for service: %s, level: %s\n", req.Service, req.Level)
	
	// Send existing logs that match filter
	s.mutex.RLock()
	var matchingLogs []LogEntry
	for _, log := range s.logs {
		if (req.Service == "" || log.Service == req.Service) &&
			(req.Level == "" || log.Level == req.Level) {
			matchingLogs = append(matchingLogs, log)
		}
	}
	s.mutex.RUnlock()
	
	for _, logEntry := range matchingLogs {
		// Send log entry to stream
		if err := stream.Send(&logEntry); err != nil {
			return fmt.Errorf("failed to send log entry: %w", err)
		}
		
		// Small delay between sends
		time.Sleep(100 * time.Millisecond)
	}
	
	// Continue streaming new logs for 30 seconds
	timeout := time.NewTimer(30 * time.Second)
	defer timeout.Stop()
	
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	
	for {
		select {
		case <-timeout.C:
			fmt.Println("StreamLogs completed after timeout")
			return nil
		case <-ticker.C:
			// Send recent matching logs
			s.mutex.RLock()
			for _, log := range s.logs[len(s.logs)-5:] {
				if (req.Service == "" || log.Service == req.Service) &&
					(req.Level == "" || log.Level == req.Level) {
					if err := stream.Send(&log); err != nil {
						s.mutex.RUnlock()
						return err
					}
				}
			}
			s.mutex.RUnlock()
		}
	}
}

// Implement client streaming (SendMetrics)
func (s *streamingServiceImpl) SendMetrics(stream StreamingService_SendMetricsServer) error {
	fmt.Println("SendMetrics client streaming started")
	
	var receivedMetrics []MetricData
	
	// Receive metrics from client stream
	for {
		metric, err := stream.Recv()
		if err == io.EOF {
			// Client finished sending, calculate aggregation
			break
		}
		if err != nil {
			return fmt.Errorf("failed to receive metric: %w", err)
		}
		
		receivedMetrics = append(receivedMetrics, *metric)
		fmt.Printf("Received metric: %s = %f\n", metric.Name, metric.Value)
	}
	
	// Aggregate metrics by name
	aggregation := aggregateMetrics(receivedMetrics)
	
	// Send aggregated result back to client
	return stream.SendAndClose(&aggregation)
}

// Implement bidirectional streaming (ChatStream)
func (s *streamingServiceImpl) ChatStream(stream StreamingService_ChatStreamServer) error {
	fmt.Println("ChatStream bidirectional streaming started")
	
	// Add client to chat clients map
	clientID := fmt.Sprintf("client_%d", time.Now().UnixNano())
	
	s.mutex.Lock()
	s.chatClients[clientID] = stream
	s.mutex.Unlock()
	
	// Remove client when done
	defer func() {
		s.mutex.Lock()
		delete(s.chatClients, clientID)
		s.mutex.Unlock()
		fmt.Printf("Client %s disconnected from chat\n", clientID)
	}()
	
	// Handle incoming messages from client
	for {
		message, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return fmt.Errorf("failed to receive chat message: %w", err)
		}
		
		fmt.Printf("Received chat message from %s: %s\n", message.Username, message.Message)
		
		// Broadcast message to all connected clients
		s.broadcastMessage(message)
	}
}

// Broadcast message to all connected clients
func (s *streamingServiceImpl) broadcastMessage(message *ChatMessage) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	
	for clientID, stream := range s.chatClients {
		// Send message to each client
		if err := stream.Send(message); err != nil {
			log.Printf("Failed to send message to client %s: %v", clientID, err)
			delete(s.chatClients, clientID)
		}
	}
}

// Aggregate metrics helper function
func aggregateMetrics(metrics []MetricData) map[string]*AggregatedMetric {
	aggregation := make(map[string]*AggregatedMetric)
	
	for _, metric := range metrics {
		name := metric.Name
		if _, exists := aggregation[name]; !exists {
			aggregation[name] = &AggregatedMetric{
				Name: name,
				Min:  metric.Value,
				Max:  metric.Value,
			}
		}
		
		agg := aggregation[name]
		agg.Count++
		agg.Sum += metric.Value
		if metric.Value < agg.Min {
			agg.Min = metric.Value
		}
		if metric.Value > agg.Max {
			agg.Max = metric.Value
		}
	}
	
	// Calculate averages
	for _, agg := range aggregation {
		if agg.Count > 0 {
			agg.Average = agg.Sum / float64(agg.Count)
		}
	}
	
	return aggregation
}

// Test server streaming
func testServerStreaming() error {
	// Create client connection
	conn, err := grpc.Dial("localhost:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}
	defer conn.Close()
	
	// Create streaming client
	client := newStreamingServiceClient(conn)
	
	// Create stream request
	req := &LogRequest{
		Service: "user-service",
		Level:   "",
	}
	
	// Start server stream
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	
	stream, err := client.StreamLogs(ctx, req)
	if err != nil {
		return err
	}
	
	// Receive streaming logs
	count := 0
	for {
		logEntry, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		
		count++
		fmt.Printf("Received log: [%s] %s - %s\n", logEntry.Level, logEntry.Service, logEntry.Message)
		
		// Limit demo to 10 logs
		if count >= 10 {
			break
		}
	}
	
	fmt.Printf("Server streaming completed, received %d logs\n", count)
	return nil
}

// Test client streaming
func testClientStreaming() error {
	// Create client connection
	conn, err := grpc.Dial("localhost:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}
	defer conn.Close()
	
	// Create streaming client
	client := newStreamingServiceClient(conn)
	
	// Start client stream
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	
	stream, err := client.SendMetrics(ctx)
	if err != nil {
		return err
	}
	
	// Send multiple metrics
	metrics := []MetricData{
		{Name: "cpu_usage", Value: 45.2, Timestamp: time.Now(), Labels: map[string]string{"host": "server1"}},
		{Name: "memory_usage", Value: 67.8, Timestamp: time.Now(), Labels: map[string]string{"host": "server1"}},
		{Name: "cpu_usage", Value: 52.1, Timestamp: time.Now(), Labels: map[string]string{"host": "server2"}},
		{Name: "memory_usage", Value: 71.3, Timestamp: time.Now(), Labels: map[string]string{"host": "server2"}},
		{Name: "cpu_usage", Value: 38.9, Timestamp: time.Now(), Labels: map[string]string{"host": "server1"}},
	}
	
	for _, metric := range metrics {
		if err := stream.Send(&metric); err != nil {
			return err
		}
		fmt.Printf("Sent metric: %s = %f\n", metric.Name, metric.Value)
	}
	
	// Close send and receive response
	response, err := stream.CloseAndRecv()
	if err != nil {
		return err
	}
	
	fmt.Println("Aggregation results:")
	for name, agg := range *response {
		fmt.Printf("  %s: count=%d, avg=%.2f, min=%.2f, max=%.2f\n",
			name, agg.Count, agg.Average, agg.Min, agg.Max)
	}
	return nil
}

// Test bidirectional streaming
func testBidirectionalStreaming() error {
	// Create client connection
	conn, err := grpc.Dial("localhost:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}
	defer conn.Close()
	
	// Create streaming client
	client := newStreamingServiceClient(conn)
	
	// Start bidirectional stream
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	
	stream, err := client.ChatStream(ctx)
	if err != nil {
		return err
	}
	
	// Start goroutine to receive messages
	go func() {
		for {
			message, err := stream.Recv()
			if err == io.EOF || err != nil {
				return
			}
			fmt.Printf("Chat received: [%s] %s\n", message.Username, message.Message)
		}
	}()
	
	// Send chat messages
	messages := []string{
		"Hello everyone!",
		"How is everyone doing?",
		"This is a test of bidirectional streaming",
	}
	
	for i, msgText := range messages {
		message := &ChatMessage{
			Username:  fmt.Sprintf("TestUser%d", i+1),
			Message:   msgText,
			Timestamp: time.Now(),
			Room:      "general",
		}
		
		if err := stream.Send(message); err != nil {
			return err
		}
		
		fmt.Printf("Chat sent: [%s] %s\n", message.Username, message.Message)
		time.Sleep(2 * time.Second)
	}
	
	// Close send
	stream.CloseSend()
	time.Sleep(2 * time.Second)
	
	return nil
}

// Client interface and implementation (simplified)
type StreamingServiceClient interface {
	StreamLogs(ctx context.Context, req *LogRequest, opts ...grpc.CallOption) (StreamingService_StreamLogsClient, error)
	SendMetrics(ctx context.Context, opts ...grpc.CallOption) (StreamingService_SendMetricsClient, error)
	ChatStream(ctx context.Context, opts ...grpc.CallOption) (StreamingService_ChatStreamClient, error)
}

type StreamingService_StreamLogsClient interface {
	Recv() (*LogEntry, error)
}

type StreamingService_SendMetricsClient interface {
	Send(*MetricData) error
	CloseAndRecv() (*map[string]*AggregatedMetric, error)
}

type StreamingService_ChatStreamClient interface {
	Send(*ChatMessage) error
	Recv() (*ChatMessage, error)
	CloseSend() error
}

type streamingServiceClient struct {
	cc *grpc.ClientConn
}

func newStreamingServiceClient(conn *grpc.ClientConn) StreamingServiceClient {
	return &streamingServiceClient{cc: conn}
}

// Simplified client implementations (normally auto-generated)
func (c *streamingServiceClient) StreamLogs(ctx context.Context, req *LogRequest, opts ...grpc.CallOption) (StreamingService_StreamLogsClient, error) {
	return &streamLogsClient{}, nil
}

func (c *streamingServiceClient) SendMetrics(ctx context.Context, opts ...grpc.CallOption) (StreamingService_SendMetricsClient, error) {
	return &sendMetricsClient{}, nil
}

func (c *streamingServiceClient) ChatStream(ctx context.Context, opts ...grpc.CallOption) (StreamingService_ChatStreamClient, error) {
	return &chatStreamClient{}, nil
}

// Simplified stream client implementations
type streamLogsClient struct{}
type sendMetricsClient struct{}
type chatStreamClient struct{}

func (s *streamLogsClient) Recv() (*LogEntry, error) { return &LogEntry{}, io.EOF }
func (s *sendMetricsClient) Send(*MetricData) error  { return nil }
func (s *sendMetricsClient) CloseAndRecv() (*map[string]*AggregatedMetric, error) {
	return &map[string]*AggregatedMetric{}, nil
}
func (s *chatStreamClient) Send(*ChatMessage) error { return nil }
func (s *chatStreamClient) Recv() (*ChatMessage, error) { return &ChatMessage{}, io.EOF }
func (s *chatStreamClient) CloseSend() error        { return nil }

// Service registration helper
func registerStreamingServiceServer(s *grpc.Server, srv StreamingServiceServer) {
	fmt.Println("StreamingService registered with gRPC server")
}