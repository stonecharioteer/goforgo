// distributed_tracing.go
// Learn distributed tracing for microservice observability

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
)

// Trace context structure
type TraceContext struct {
	TraceID      string             `json:"trace_id"`
	SpanID       string             `json:"span_id"`
	ParentSpanID string             `json:"parent_span_id,omitempty"`
	Operation    string             `json:"operation"`
	StartTime    time.Time          `json:"start_time"`
	EndTime      time.Time          `json:"end_time,omitempty"`
	Duration     time.Duration      `json:"duration,omitempty"`
	Status       string             `json:"status"`
	Tags         map[string]string  `json:"tags,omitempty"`
	Spans        map[string]*Span   `json:"spans,omitempty"`
}

// Span structure for individual operations
type Span struct {
	SpanID       string            `json:"span_id"`
	ParentSpanID string            `json:"parent_span_id,omitempty"`
	Operation    string            `json:"operation"`
	StartTime    time.Time         `json:"start_time"`
	EndTime      time.Time         `json:"end_time,omitempty"`
	Duration     time.Duration     `json:"duration,omitempty"`
	Status       string            `json:"status"`
	Tags         map[string]string `json:"tags,omitempty"`
	Logs         []string          `json:"logs,omitempty"`
}

// Trace collector structure
type TraceCollector struct {
	traces map[string]*TraceContext
	mutex  sync.RWMutex
}

// Global trace collector instance
var (
	collector *TraceCollector
)

func main() {
	fmt.Println("=== Distributed Tracing ===")
	
	// Initialize global trace collector
	collector = NewTraceCollector()
	
	// Setup HTTP routes
	http.HandleFunc("/order", handleOrderRequest)
	http.HandleFunc("/traces", handleTraces)
	http.HandleFunc("/traces/", handleSingleTrace)
	
	fmt.Println("Distributed Tracing Demo starting on :8080")
	fmt.Println("Available endpoints:")
	fmt.Println("  POST /order - Process order (triggers multiple service calls)")
	fmt.Println("  GET  /traces - List all traces")
	fmt.Println("  GET  /traces/{traceId} - Get specific trace details")
	
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// Create new trace collector
func NewTraceCollector() *TraceCollector {
	return &TraceCollector{
		traces: make(map[string]*TraceContext),
		mutex:  sync.RWMutex{},
	}
}

// Start new trace
func (tc *TraceCollector) StartTrace(operation string) *TraceContext {
	tc.mutex.Lock()
	defer tc.mutex.Unlock()
	
	traceID := generateID()
	spanID := generateID()
	
	trace := &TraceContext{
		TraceID:   traceID,
		SpanID:    spanID,
		Operation: operation,
		StartTime: time.Now(),
		Status:    "in_progress",
		Tags:      make(map[string]string),
		Spans:     make(map[string]*Span),
	}
	
	// Store trace in collector
	tc.traces[traceID] = trace
	
	log.Printf("Started trace %s for operation: %s", traceID, operation)
	return trace
}

// Start new span within trace
func (tc *TraceCollector) StartSpan(traceID, parentSpanID, operation string) *Span {
	tc.mutex.Lock()
	defer tc.mutex.Unlock()
	
	spanID := generateID()
	
	span := &Span{
		SpanID:       spanID,
		ParentSpanID: parentSpanID,
		Operation:    operation,
		StartTime:    time.Now(),
		Status:       "in_progress",
		Tags:         make(map[string]string),
		Logs:         make([]string, 0),
	}
	
	// Add span to trace
	trace, exists := tc.traces[traceID]
	if exists {
		if trace.Spans == nil {
			trace.Spans = make(map[string]*Span)
		}
		trace.Spans[spanID] = span
	}
	
	log.Printf("Started span %s for operation: %s (trace: %s)", spanID, operation, traceID)
	return span
}

// Finish span
func (tc *TraceCollector) FinishSpan(traceID, spanID string, status string, tags map[string]string) {
	tc.mutex.Lock()
	defer tc.mutex.Unlock()
	
	trace, exists := tc.traces[traceID]
	if !exists {
		return
	}
	
	span, exists := trace.Spans[spanID]
	if !exists {
		return
	}
	
	// Update span with completion details
	span.EndTime = time.Now()
	span.Duration = span.EndTime.Sub(span.StartTime)
	span.Status = status
	span.Tags = addTags(span.Tags, tags)
	
	log.Printf("Finished span %s with status: %s (duration: %v)", spanID, status, span.Duration)
}

// Finish trace
func (tc *TraceCollector) FinishTrace(traceID string, status string) {
	tc.mutex.Lock()
	defer tc.mutex.Unlock()
	
	trace, exists := tc.traces[traceID]
	if !exists {
		return
	}
	
	// Update trace with completion details
	trace.EndTime = time.Now()
	trace.Duration = trace.EndTime.Sub(trace.StartTime)
	trace.Status = status
	
	log.Printf("Finished trace %s with status: %s (duration: %v)", traceID, status, trace.Duration)
}

// Get all traces
func (tc *TraceCollector) GetTraces() map[string]*TraceContext {
	tc.mutex.RLock()
	defer tc.mutex.RUnlock()
	
	// Return copy of traces
	result := make(map[string]*TraceContext)
	for id, trace := range tc.traces {
		result[id] = trace
	}
	
	return result
}

// Get specific trace
func (tc *TraceCollector) GetTrace(traceID string) *TraceContext {
	tc.mutex.RLock()
	defer tc.mutex.RUnlock()
	
	trace, exists := tc.traces[traceID]
	if !exists {
		return nil
	}
	
	return trace
}

// Generate unique ID
func generateID() string {
	return strconv.FormatInt(time.Now().UnixNano(), 16) + strconv.Itoa(rand.Intn(1000))
}

// Add tags to span/trace
func addTags(existing map[string]string, new map[string]string) map[string]string {
	if existing == nil {
		existing = make(map[string]string)
	}
	
	for k, v := range new {
		existing[k] = v
	}
	
	return existing
}

// Order processing handler with distributed tracing
func handleOrderRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	
	// Start main trace for order processing
	trace := collector.StartTrace("process_order")
	defer func() {
		collector.FinishTrace(trace.TraceID, "success")
	}()
	
	// Simulate order processing workflow with multiple services
	orderID := fmt.Sprintf("order_%d", rand.Intn(10000))
	
	// Call each service with tracing
	if !validateOrder(trace.TraceID, trace.SpanID, orderID) {
		collector.FinishTrace(trace.TraceID, "error")
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Order validation failed")
		return
	}
	
	if !reserveInventory(trace.TraceID, trace.SpanID, orderID) {
		collector.FinishTrace(trace.TraceID, "error")
		w.WriteHeader(http.StatusConflict)
		fmt.Fprint(w, "Inventory reservation failed")
		return
	}
	
	if !processPayment(trace.TraceID, trace.SpanID, orderID) {
		collector.FinishTrace(trace.TraceID, "error")
		w.WriteHeader(http.StatusPaymentRequired)
		fmt.Fprint(w, "Payment processing failed")
		return
	}
	
	if !shipOrder(trace.TraceID, trace.SpanID, orderID) {
		collector.FinishTrace(trace.TraceID, "error")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Shipping failed")
		return
	}
	
	// Return success response
	w.Header().Set("Content-Type", "application/json")
	response := map[string]interface{}{
		"order_id": orderID,
		"status":   "processed",
		"trace_id": trace.TraceID,
	}
	json.NewEncoder(w).Encode(response)
}

// Validate order service call
func validateOrder(traceID, parentSpanID, orderID string) bool {
	span := collector.StartSpan(traceID, parentSpanID, "validate_order")
	defer func() {
		if r := recover(); r != nil {
			collector.FinishSpan(traceID, span.SpanID, "error", map[string]string{
				"error": fmt.Sprintf("%v", r),
			})
		}
	}()
	
	// Simulate validation logic
	time.Sleep(time.Duration(rand.Intn(50)) * time.Millisecond)
	
	// Add tags to span
	tags := map[string]string{
		"order_id": orderID,
		"service":  "validation",
	}
	
	// Simulate 10% failure rate
	if rand.Float32() < 0.1 {
		tags["error"] = "validation_failed"
		collector.FinishSpan(traceID, span.SpanID, "error", tags)
		return false
	}
	
	tags["result"] = "valid"
	collector.FinishSpan(traceID, span.SpanID, "success", tags)
	return true
}

// Reserve inventory service call
func reserveInventory(traceID, parentSpanID, orderID string) bool {
	span := collector.StartSpan(traceID, parentSpanID, "reserve_inventory")
	defer func() {
		if r := recover(); r != nil {
			collector.FinishSpan(traceID, span.SpanID, "error", map[string]string{
				"error": fmt.Sprintf("%v", r),
			})
		}
	}()
	
	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	
	tags := map[string]string{
		"order_id": orderID,
		"service":  "inventory",
	}
	
	// Simulate 15% failure rate
	if rand.Float32() < 0.15 {
		tags["error"] = "insufficient_inventory"
		collector.FinishSpan(traceID, span.SpanID, "error", tags)
		return false
	}
	
	tags["result"] = "reserved"
	collector.FinishSpan(traceID, span.SpanID, "success", tags)
	return true
}

// Process payment service call
func processPayment(traceID, parentSpanID, orderID string) bool {
	span := collector.StartSpan(traceID, parentSpanID, "process_payment")
	defer func() {
		if r := recover(); r != nil {
			collector.FinishSpan(traceID, span.SpanID, "error", map[string]string{
				"error": fmt.Sprintf("%v", r),
			})
		}
	}()
	
	time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
	
	tags := map[string]string{
		"order_id": orderID,
		"service":  "payment",
	}
	
	// Simulate 20% failure rate
	if rand.Float32() < 0.2 {
		tags["error"] = "payment_declined"
		collector.FinishSpan(traceID, span.SpanID, "error", tags)
		return false
	}
	
	tags["result"] = "processed"
	collector.FinishSpan(traceID, span.SpanID, "success", tags)
	return true
}

// Ship order service call
func shipOrder(traceID, parentSpanID, orderID string) bool {
	span := collector.StartSpan(traceID, parentSpanID, "ship_order")
	defer func() {
		if r := recover(); r != nil {
			collector.FinishSpan(traceID, span.SpanID, "error", map[string]string{
				"error": fmt.Sprintf("%v", r),
			})
		}
	}()
	
	time.Sleep(time.Duration(rand.Intn(150)) * time.Millisecond)
	
	tags := map[string]string{
		"order_id": orderID,
		"service":  "shipping",
	}
	
	// Simulate 5% failure rate
	if rand.Float32() < 0.05 {
		tags["error"] = "shipping_unavailable"
		collector.FinishSpan(traceID, span.SpanID, "error", tags)
		return false
	}
	
	tags["result"] = "shipped"
	collector.FinishSpan(traceID, span.SpanID, "success", tags)
	return true
}

// Get all traces handler
func handleTraces(w http.ResponseWriter, r *http.Request) {
	traces := collector.GetTraces()
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(traces)
}

// Get single trace handler
func handleSingleTrace(w http.ResponseWriter, r *http.Request) {
	traceID := strings.TrimPrefix(r.URL.Path, "/traces/")
	if traceID == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Trace ID required")
		return
	}
	
	trace := collector.GetTrace(traceID)
	if trace == nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "Trace not found")
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(trace)
}