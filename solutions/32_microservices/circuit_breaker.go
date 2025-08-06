// circuit_breaker.go
// Learn circuit breaker pattern for fault tolerance in microservices

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"sync"
	"time"
)

// Circuit breaker states
const (
	StateClosed   = iota
	StateOpen     = iota
	StateHalfOpen = iota
)

// Circuit breaker structure
type CircuitBreaker struct {
	name            string
	maxFailures     int
	timeout         time.Duration
	state           int
	failures        int
	lastFailureTime time.Time
	mutex           sync.RWMutex
}

// Circuit breaker error types
var (
	ErrCircuitOpen      = errors.New("circuit breaker is open")
	ErrTooManyRequests  = errors.New("too many requests")
)

// Global circuit breakers for demo
var (
	userServiceCB    *CircuitBreaker
	orderServiceCB   *CircuitBreaker
	paymentServiceCB *CircuitBreaker
)

func main() {
	fmt.Println("=== Circuit Breaker Pattern ===")
	
	// Create circuit breakers for different services
	userServiceCB = NewCircuitBreaker("user-service", 3, 30*time.Second)
	orderServiceCB = NewCircuitBreaker("order-service", 2, 20*time.Second)
	paymentServiceCB = NewCircuitBreaker("payment-service", 5, 60*time.Second)
	
	// Setup HTTP routes
	http.HandleFunc("/user/", func(w http.ResponseWriter, r *http.Request) {
		handleUserService(userServiceCB, w, r)
	})
	http.HandleFunc("/order/", func(w http.ResponseWriter, r *http.Request) {
		handleOrderService(orderServiceCB, w, r)
	})
	http.HandleFunc("/payment/", func(w http.ResponseWriter, r *http.Request) {
		handlePaymentService(paymentServiceCB, w, r)
	})
	http.HandleFunc("/status", statusHandler)
	
	fmt.Println("Circuit Breaker Demo starting on :8080")
	fmt.Println("Available endpoints:")
	fmt.Println("  GET /user/{id} - User service (fails ~30%)")
	fmt.Println("  GET /order/{id} - Order service (fails ~50%)")
	fmt.Println("  GET /payment/{id} - Payment service (fails ~20%)")
	fmt.Println("  GET /status - Circuit breaker status")
	fmt.Println("\nTry making multiple requests to see circuit breaker in action!")
	
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// Create new circuit breaker
func NewCircuitBreaker(name string, maxFailures int, timeout time.Duration) *CircuitBreaker {
	return &CircuitBreaker{
		name:        name,
		maxFailures: maxFailures,
		timeout:     timeout,
		state:       StateClosed,
		failures:    0,
		mutex:       sync.RWMutex{},
	}
}

// Execute function with circuit breaker protection
func (cb *CircuitBreaker) Execute(fn func() (interface{}, error)) (interface{}, error) {
	cb.mutex.Lock()
	defer cb.mutex.Unlock()
	
	// Check current state
	switch cb.state {
	case StateOpen:
		// Check if timeout has passed
		if time.Since(cb.lastFailureTime) > cb.timeout {
			cb.state = StateHalfOpen
			log.Printf("Circuit breaker %s: OPEN -> HALF-OPEN", cb.name)
		} else {
			return nil, ErrCircuitOpen
		}
	case StateHalfOpen:
		// Allow one request in half-open state
		// fall through to execute
	case StateClosed:
		// Normal operation
		// fall through to execute
	}
	
	// Execute the function
	cb.mutex.Unlock()
	result, err := fn()
	cb.mutex.Lock()
	
	if err != nil {
		cb.onFailure()
		return nil, err
	}
	
	cb.onSuccess()
	return result, nil
}

// Handle successful execution
func (cb *CircuitBreaker) onSuccess() {
	cb.failures = 0
	prevState := cb.state
	cb.state = StateClosed
	if prevState != StateClosed {
		log.Printf("Circuit breaker %s: %s -> CLOSED", cb.name, cb.stateString(prevState))
	}
}

// Handle failed execution
func (cb *CircuitBreaker) onFailure() {
	cb.failures++
	cb.lastFailureTime = time.Now()
	
	if cb.failures >= cb.maxFailures && cb.state == StateClosed {
		cb.state = StateOpen
		log.Printf("Circuit breaker %s: CLOSED -> OPEN (failures: %d)", cb.name, cb.failures)
	}
}

// Get current state (thread-safe)
func (cb *CircuitBreaker) GetState() string {
	cb.mutex.RLock()
	defer cb.mutex.RUnlock()
	
	return cb.stateString(cb.state)
}

// Get current failure count (thread-safe)
func (cb *CircuitBreaker) GetFailures() int {
	cb.mutex.RLock()
	defer cb.mutex.RUnlock()
	return cb.failures
}

func (cb *CircuitBreaker) stateString(state int) string {
	switch state {
	case StateClosed:
		return "CLOSED"
	case StateOpen:
		return "OPEN"
	case StateHalfOpen:
		return "HALF-OPEN"
	default:
		return "UNKNOWN"
	}
}

// User service handler with circuit breaker
func handleUserService(cb *CircuitBreaker, w http.ResponseWriter, r *http.Request) {
	userID := strings.TrimPrefix(r.URL.Path, "/user/")
	
	result, err := cb.Execute(func() (interface{}, error) {
		return simulateUserService(userID)
	})
	
	if err != nil {
		if err == ErrCircuitOpen {
			w.WriteHeader(http.StatusServiceUnavailable)
			fmt.Fprint(w, "User service temporarily unavailable (circuit breaker open)")
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "User service error: %v", err)
		}
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

// Order service handler with circuit breaker
func handleOrderService(cb *CircuitBreaker, w http.ResponseWriter, r *http.Request) {
	orderID := strings.TrimPrefix(r.URL.Path, "/order/")
	
	result, err := cb.Execute(func() (interface{}, error) {
		return simulateOrderService(orderID)
	})
	
	if err != nil {
		if err == ErrCircuitOpen {
			w.WriteHeader(http.StatusServiceUnavailable)
			fmt.Fprint(w, "Order service temporarily unavailable (circuit breaker open)")
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Order service error: %v", err)
		}
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

// Payment service handler with circuit breaker
func handlePaymentService(cb *CircuitBreaker, w http.ResponseWriter, r *http.Request) {
	paymentID := strings.TrimPrefix(r.URL.Path, "/payment/")
	
	result, err := cb.Execute(func() (interface{}, error) {
		return simulatePaymentService(paymentID)
	})
	
	if err != nil {
		if err == ErrCircuitOpen {
			w.WriteHeader(http.StatusServiceUnavailable)
			fmt.Fprint(w, "Payment service temporarily unavailable (circuit breaker open)")
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Payment service error: %v", err)
		}
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

// Simulate user service (fails ~30% of the time)
func simulateUserService(userID string) (map[string]interface{}, error) {
	// Simulate network delay
	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	
	if rand.Float32() < 0.3 {
		return nil, errors.New("user service unavailable")
	}
	
	return map[string]interface{}{
		"user_id": userID,
		"name":    fmt.Sprintf("User%s", userID),
		"email":   fmt.Sprintf("user%s@example.com", userID),
		"active":  true,
	}, nil
}

// Simulate order service (fails ~50% of the time)
func simulateOrderService(orderID string) (map[string]interface{}, error) {
	time.Sleep(time.Duration(rand.Intn(150)) * time.Millisecond)
	
	if rand.Float32() < 0.5 {
		return nil, errors.New("order service timeout")
	}
	
	return map[string]interface{}{
		"order_id": orderID,
		"status":   "confirmed",
		"total":    rand.Float64() * 1000,
		"items":    rand.Intn(10) + 1,
	}, nil
}

// Simulate payment service (fails ~20% of the time)
func simulatePaymentService(paymentID string) (map[string]interface{}, error) {
	time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
	
	if rand.Float32() < 0.2 {
		return nil, errors.New("payment service error")
	}
	
	return map[string]interface{}{
		"payment_id": paymentID,
		"status":     "processed",
		"amount":     rand.Float64() * 500,
		"method":     "credit_card",
	}, nil
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
	status := map[string]interface{}{
		"timestamp": time.Now().Format(time.RFC3339),
		"circuit_breakers": map[string]interface{}{
			"user-service": map[string]interface{}{
				"state":    userServiceCB.GetState(),
				"failures": userServiceCB.GetFailures(),
			},
			"order-service": map[string]interface{}{
				"state":    orderServiceCB.GetState(),
				"failures": orderServiceCB.GetFailures(),
			},
			"payment-service": map[string]interface{}{
				"state":    paymentServiceCB.GetState(),
				"failures": paymentServiceCB.GetFailures(),
			},
		},
	}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(status)
}