// http_client_advanced.go - SOLUTION
// Learn advanced HTTP client features: custom transports, retries, timeouts, and middleware

package main

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"math"
	"net/http"
	"sync"
	"time"
)

// Custom HTTP client with advanced configuration
type AdvancedHTTPClient struct {
	client    *http.Client
	maxRetries int
	userAgent  string
}

// HTTP client builder pattern
type ClientBuilder struct {
	timeout       time.Duration
	maxRetries    int
	userAgent     string
	transport     http.RoundTripper
}

// Create new client builder
func NewClientBuilder() *ClientBuilder {
	return &ClientBuilder{
		timeout:    30 * time.Second,
		maxRetries: 3,
		userAgent:  "AdvancedHTTPClient/1.0",
		transport:  http.DefaultTransport,
	}
}

// Builder methods for configuration
func (cb *ClientBuilder) WithTimeout(timeout time.Duration) *ClientBuilder {
	cb.timeout = timeout
	return cb
}

func (cb *ClientBuilder) WithRetries(maxRetries int) *ClientBuilder {
	cb.maxRetries = maxRetries
	return cb
}

func (cb *ClientBuilder) WithUserAgent(userAgent string) *ClientBuilder {
	cb.userAgent = userAgent
	return cb
}

func (cb *ClientBuilder) WithCustomTransport(transport http.RoundTripper) *ClientBuilder {
	cb.transport = transport
	return cb
}

func (cb *ClientBuilder) Build() *AdvancedHTTPClient {
	return &AdvancedHTTPClient{
		client: &http.Client{
			Transport: cb.transport,
			Timeout:   cb.timeout,
		},
		maxRetries: cb.maxRetries,
		userAgent:  cb.userAgent,
	}
}

// HTTP client methods
func (c *AdvancedHTTPClient) Get(url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", c.userAgent)
	return c.DoWithRetry(req)
}

func (c *AdvancedHTTPClient) Post(url string, contentType string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", contentType)
	req.Header.Set("User-Agent", c.userAgent)
	return c.DoWithRetry(req)
}

func (c *AdvancedHTTPClient) DoWithRetry(req *http.Request) (*http.Response, error) {
	var lastErr error
	
	for attempt := 0; attempt <= c.maxRetries; attempt++ {
		if attempt > 0 {
			waitTime := time.Duration(math.Pow(2, float64(attempt-1))) * time.Second
			time.Sleep(waitTime)
		}
		
		resp, err := c.client.Do(req)
		if err == nil && resp.StatusCode < 500 {
			return resp, nil
		}
		
		if resp != nil {
			resp.Body.Close()
		}
		lastErr = err
	}
	
	return nil, fmt.Errorf("request failed after %d retries: %v", c.maxRetries, lastErr)
}

// Custom transport for logging requests
type LoggingTransport struct {
	Transport http.RoundTripper
}

// Implement RoundTripper interface
func (lt *LoggingTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	start := time.Now()
	fmt.Printf("🔗 %s %s\n", req.Method, req.URL.String())
	
	resp, err := lt.Transport.RoundTrip(req)
	elapsed := time.Since(start)
	
	if err != nil {
		fmt.Printf("❌ Request failed in %v: %v\n", elapsed, err)
	} else {
		fmt.Printf("✅ Response %s in %v\n", resp.Status, elapsed)
	}
	
	return resp, err
}

// Custom transport for adding authentication
type AuthTransport struct {
	Transport http.RoundTripper
	Token     string
}

// Implement RoundTripper interface for auth
func (at *AuthTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("Authorization", "Bearer "+at.Token)
	return at.Transport.RoundTrip(req)
}

// Rate limiting transport
type RateLimitTransport struct {
	Transport http.RoundTripper
	limiter   chan struct{}
}

func NewRateLimitTransport(transport http.RoundTripper, requestsPerSecond int) *RateLimitTransport {
	limiter := make(chan struct{}, requestsPerSecond)
	
	// Fill the limiter
	for i := 0; i < requestsPerSecond; i++ {
		limiter <- struct{}{}
	}
	
	// Refill the limiter every second
	go func() {
		ticker := time.NewTicker(time.Second)
		defer ticker.Stop()
		
		for range ticker.C {
			select {
			case limiter <- struct{}{}:
			default:
			}
		}
	}()
	
	return &RateLimitTransport{
		Transport: transport,
		limiter:   limiter,
	}
}

// Implement RoundTripper interface for rate limiting
func (rt *RateLimitTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	<-rt.limiter // Wait for rate limit token
	return rt.Transport.RoundTrip(req)
}

// Circuit breaker pattern implementation
type CircuitBreaker struct {
	mu           sync.Mutex
	state        int
	failures     int
	threshold    int
	timeout      time.Duration
	lastFailTime time.Time
}

// Circuit breaker states
const (
	StateClosed = iota
	StateOpen
	StateHalfOpen
)

// Circuit breaker methods
func NewCircuitBreaker(threshold int, timeout time.Duration) *CircuitBreaker {
	return &CircuitBreaker{
		state:     StateClosed,
		threshold: threshold,
		timeout:   timeout,
	}
}

func (cb *CircuitBreaker) Call(fn func() error) error {
	cb.mu.Lock()
	defer cb.mu.Unlock()
	
	// Check if we should move from Open to Half-Open
	if cb.state == StateOpen && time.Since(cb.lastFailTime) > cb.timeout {
		cb.state = StateHalfOpen
	}
	
	// Reject calls in Open state
	if cb.state == StateOpen {
		return errors.New("circuit breaker is open")
	}
	
	// Execute the function
	err := fn()
	
	if err != nil {
		cb.recordFailure()
		return err
	}
	
	cb.recordSuccess()
	return nil
}

func (cb *CircuitBreaker) recordSuccess() {
	cb.failures = 0
	cb.state = StateClosed
}

func (cb *CircuitBreaker) recordFailure() {
	cb.failures++
	cb.lastFailTime = time.Now()
	
	if cb.failures >= cb.threshold {
		cb.state = StateOpen
	}
}

// Retry policy implementation
type RetryPolicy struct {
	maxRetries int
	baseDelay  time.Duration
}

func NewRetryPolicy(maxRetries int, baseDelay time.Duration) *RetryPolicy {
	return &RetryPolicy{
		maxRetries: maxRetries,
		baseDelay:  baseDelay,
	}
}

func (rp *RetryPolicy) Execute(fn func() error) error {
	var lastErr error
	
	for attempt := 0; attempt <= rp.maxRetries; attempt++ {
		if attempt > 0 {
			delay := time.Duration(math.Pow(2, float64(attempt-1))) * rp.baseDelay
			time.Sleep(delay)
		}
		
		err := fn()
		if err == nil {
			return nil
		}
		
		if !rp.shouldRetry(err, attempt) {
			return err
		}
		
		lastErr = err
	}
	
	return fmt.Errorf("failed after %d retries: %v", rp.maxRetries, lastErr)
}

func (rp *RetryPolicy) shouldRetry(err error, attempt int) bool {
	// Don't retry if we've exhausted attempts
	if attempt >= rp.maxRetries {
		return false
	}
	
	// Add logic to determine if error is retryable
	// For this example, we'll retry most errors
	return true
}

// HTTP request/response middleware
type Middleware func(http.RoundTripper) http.RoundTripper

// Create middleware chain
func ChainMiddleware(transport http.RoundTripper, middlewares ...Middleware) http.RoundTripper {
	for i := len(middlewares) - 1; i >= 0; i-- {
		transport = middlewares[i](transport)
	}
	return transport
}

// Metrics middleware
func MetricsMiddleware() Middleware {
	return func(next http.RoundTripper) http.RoundTripper {
		return &metricsTransport{Transport: next}
	}
}

type metricsTransport struct {
	Transport http.RoundTripper
}

func (mt *metricsTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	start := time.Now()
	resp, err := mt.Transport.RoundTrip(req)
	elapsed := time.Since(start)
	
	fmt.Printf("📊 Metrics - Method: %s, URL: %s, Duration: %v\n", 
		req.Method, req.URL.String(), elapsed)
	
	return resp, err
}

// Timeout middleware
func TimeoutMiddleware(timeout time.Duration) Middleware {
	return func(next http.RoundTripper) http.RoundTripper {
		return &timeoutTransport{Transport: next, timeout: timeout}
	}
}

type timeoutTransport struct {
	Transport http.RoundTripper
	timeout   time.Duration
}

func (tt *timeoutTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	ctx, cancel := context.WithTimeout(req.Context(), tt.timeout)
	defer cancel()
	
	req = req.WithContext(ctx)
	return tt.Transport.RoundTrip(req)
}

func main() {
	fmt.Println("=== Advanced HTTP Client ===")
	
	fmt.Println("\n=== Basic HTTP Client with Configuration ===")
	
	// Create advanced HTTP client with builder
	client := NewClientBuilder().
		WithTimeout(10 * time.Second).
		WithRetries(2).
		WithUserAgent("GoForGo-Advanced/1.0").
		Build()
	
	// Test simple GET request
	testURL := "https://httpbin.org/get"
	fmt.Printf("Making GET request to: %s\n", testURL)
	
	resp, err := client.Get(testURL)
	if err != nil {
		fmt.Printf("❌ Request failed: %v\n", err)
	} else {
		fmt.Printf("✅ Request successful: %s\n", resp.Status)
		resp.Body.Close()
	}
	
	fmt.Println("\n=== Custom Transport - Logging ===")
	
	// Create HTTP client with logging transport
	baseTransport := http.DefaultTransport
	loggingTransport := &LoggingTransport{Transport: baseTransport}
	
	loggingClient := &http.Client{
		Transport: loggingTransport,
		Timeout:   10 * time.Second,
	}
	
	// Make request with logging
	fmt.Println("Making request with logging transport:")
	resp, err = loggingClient.Get(testURL)
	if err != nil {
		fmt.Printf("Request failed: %v\n", err)
	} else {
		fmt.Printf("Request completed: %s\n", resp.Status)
		resp.Body.Close()
	}
	
	fmt.Println("\n=== Authentication Transport ===")
	
	// Create client with authentication
	authTransport := &AuthTransport{
		Transport: http.DefaultTransport,
		Token:     "test-token-12345",
	}
	authClient := &http.Client{
		Transport: authTransport,
		Timeout:   10 * time.Second,
	}
	
	// Test authenticated request
	authURL := "https://httpbin.org/bearer"
	fmt.Printf("Making authenticated request to: %s\n", authURL)
	
	resp, err = authClient.Get(authURL)
	if err != nil {
		fmt.Printf("❌ Authenticated request failed: %v\n", err)
	} else {
		fmt.Printf("✅ Authenticated request successful: %s\n", resp.Status)
		resp.Body.Close()
	}
	
	fmt.Println("\n=== Rate Limiting ===")
	
	// Create rate limited client
	rateLimitTransport := NewRateLimitTransport(http.DefaultTransport, 1)
	rateLimitClient := &http.Client{
		Transport: rateLimitTransport,
		Timeout:   10 * time.Second,
	}
	
	// Make multiple requests to test rate limiting
	fmt.Println("Testing rate limiting (3 requests):")
	for i := 0; i < 3; i++ {
		start := time.Now()
		resp, err = rateLimitClient.Get(testURL)
		elapsed := time.Since(start)
		
		if err != nil {
			fmt.Printf("  Request %d failed: %v\n", i+1, err)
		} else {
			fmt.Printf("  Request %d completed in %v: %s\n", i+1, elapsed, resp.Status)
			resp.Body.Close()
		}
	}
	
	fmt.Println("\n=== Circuit Breaker Pattern ===")
	
	// Test circuit breaker
	circuitBreaker := NewCircuitBreaker(3, 2*time.Second)
	
	// Simulate failing service
	failingService := func() error {
		return errors.New("service unavailable")
	}
	
	fmt.Println("Testing circuit breaker with failing service:")
	for i := 0; i < 8; i++ {
		err := circuitBreaker.Call(failingService)
		if err != nil {
			fmt.Printf("  Call %d: %v\n", i+1, err)
		} else {
			fmt.Printf("  Call %d: Success\n", i+1)
		}
	}
	
	fmt.Println("\n=== Retry Policy ===")
	
	// Test retry policy
	retryPolicy := NewRetryPolicy(3, 100*time.Millisecond)
	
	// Service that fails first few times
	attemptCount := 0
	unreliableService := func() error {
		attemptCount++
		if attemptCount < 3 {
			return errors.New("temporary failure")
		}
		return nil
	}
	
	fmt.Println("Testing retry policy with unreliable service:")
	err = retryPolicy.Execute(unreliableService)
	if err != nil {
		fmt.Printf("❌ Service failed after retries: %v\n", err)
	} else {
		fmt.Printf("✅ Service succeeded after %d attempts\n", attemptCount)
	}
	
	fmt.Println("\n=== Middleware Chain ===")
	
	// Create middleware chain
	baseTransport = http.DefaultTransport
	middlewareChain := ChainMiddleware(baseTransport,
		MetricsMiddleware(),
		TimeoutMiddleware(5*time.Second),
	)
	
	middlewareClient := &http.Client{
		Transport: middlewareChain,
		Timeout:   15 * time.Second,
	}
	
	// Test request with middleware chain
	fmt.Println("Making request with middleware chain:")
	resp, err = middlewareClient.Get(testURL)
	if err != nil {
		fmt.Printf("❌ Middleware request failed: %v\n", err)
	} else {
		fmt.Printf("✅ Middleware request successful: %s\n", resp.Status)
		resp.Body.Close()
	}
	
	fmt.Println("\n=== Context and Cancellation ===")
	
	// Test request cancellation
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()
	
	// Create request with context
	req, err := http.NewRequestWithContext(ctx, "GET", "https://httpbin.org/delay/1", nil)
	if err != nil {
		fmt.Printf("Error creating request: %v\n", err)
		return
	}
	
	fmt.Println("Testing request cancellation:")
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err = client.Do(req)
	if err != nil {
		fmt.Printf("✅ Request correctly cancelled: %v\n", err)
	} else {
		fmt.Printf("❌ Request should have been cancelled\n")
		resp.Body.Close()
	}
	
	fmt.Println("\n=== Best Practices Summary ===")
	fmt.Println("✅ Configure appropriate timeouts")
	fmt.Println("✅ Implement retry logic with exponential backoff")
	fmt.Println("✅ Use circuit breakers for failing services")
	fmt.Println("✅ Add rate limiting to prevent overwhelming servers")
	fmt.Println("✅ Use middleware for cross-cutting concerns")
	fmt.Println("✅ Implement proper authentication")
	fmt.Println("✅ Log requests for debugging and monitoring")
	fmt.Println("✅ Use context for cancellation and timeouts")
	fmt.Println("✅ Handle different types of errors appropriately")
	fmt.Println("✅ Pool connections for better performance")
}