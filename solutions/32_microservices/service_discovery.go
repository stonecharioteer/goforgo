// service_discovery.go
// Learn service discovery patterns for microservices

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"sync"
	"time"
)

// Service instance structure
type ServiceInstance struct {
	ID            string    `json:"id"`
	Name          string    `json:"name"`
	Host          string    `json:"host"`
	Port          int       `json:"port"`
	Health        bool      `json:"health"`
	LastHeartbeat time.Time `json:"last_heartbeat"`
}

// Service registry structure
type ServiceRegistry struct {
	services map[string][]ServiceInstance
	mutex    sync.RWMutex
}

func main() {
	fmt.Println("=== Microservice Discovery ===")
	
	// Create service registry
	registry := newServiceRegistry()
	
	// Start health check routine
	go registry.startHealthChecker()
	
	// Setup HTTP routes
	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		registerHandler(registry, w, r)
	})
	http.HandleFunc("/discover/", func(w http.ResponseWriter, r *http.Request) {
		discoverHandler(registry, w, r)
	})
	http.HandleFunc("/services", func(w http.ResponseWriter, r *http.Request) {
		listServicesHandler(registry, w, r)
	})
	http.HandleFunc("/health", healthHandler)
	
	// Start multiple service instances for demo
	go startDemoService("user-service", 8081)
	go startDemoService("order-service", 8082)
	go startDemoService("payment-service", 8083)
	
	fmt.Println("Service Registry starting on :8080")
	fmt.Println("Demo services starting on :8081, :8082, :8083")
	fmt.Println("Available endpoints:")
	fmt.Println("  POST /register - Register a service")
	fmt.Println("  GET  /discover/{service} - Discover service instances")
	fmt.Println("  GET  /services - List all services")
	
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// Create new service registry
func newServiceRegistry() *ServiceRegistry {
	return &ServiceRegistry{
		services: make(map[string][]ServiceInstance),
		mutex:    sync.RWMutex{},
	}
}

// Register service instance
func (sr *ServiceRegistry) Register(instance ServiceInstance) {
	sr.mutex.Lock()
	defer sr.mutex.Unlock()
	
	// Initialize service map if needed
	if sr.services == nil {
		sr.services = make(map[string][]ServiceInstance)
	}
	
	// Initialize service slice if needed
	if _, exists := sr.services[instance.Name]; !exists {
		sr.services[instance.Name] = make([]ServiceInstance, 0)
	}
	
	// Update existing instance or add new one
	found := false
	for i := range sr.services[instance.Name] {
		if sr.services[instance.Name][i].ID == instance.ID {
			sr.services[instance.Name][i] = instance
			found = true
			break
		}
	}
	
	if !found {
		sr.services[instance.Name] = append(sr.services[instance.Name], instance)
	}
	
	log.Printf("Registered service: %s (%s:%d)", instance.Name, instance.Host, instance.Port)
}

// Discover healthy service instances
func (sr *ServiceRegistry) Discover(serviceName string) []ServiceInstance {
	sr.mutex.RLock()
	defer sr.mutex.RUnlock()
	
	instances, exists := sr.services[serviceName]
	if !exists {
		return nil
	}
	
	// Filter healthy instances
	var healthy []ServiceInstance
	for _, instance := range instances {
		if instance.Health {
			healthy = append(healthy, instance)
		}
	}
	
	return healthy
}

// Get all services
func (sr *ServiceRegistry) GetAllServices() map[string][]ServiceInstance {
	sr.mutex.RLock()
	defer sr.mutex.RUnlock()
	
	// Create copy of services map
	result := make(map[string][]ServiceInstance)
	for name, instances := range sr.services {
		result[name] = make([]ServiceInstance, len(instances))
		copy(result[name], instances)
	}
	
	return result
}

// Health check routine
func (sr *ServiceRegistry) startHealthChecker() {
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()
	
	for range ticker.C {
		sr.mutex.Lock()
		
		for serviceName, instances := range sr.services {
			for i := range instances {
				// Check if instance is stale
				if time.Since(instances[i].LastHeartbeat) > 1*time.Minute {
					sr.services[serviceName][i].Health = false
					log.Printf("Marked unhealthy: %s (%s:%d)", 
						instances[i].Name, instances[i].Host, instances[i].Port)
				}
			}
		}
		
		sr.mutex.Unlock()
	}
}

// Register handler
func registerHandler(registry *ServiceRegistry, w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	
	// Parse service instance from request
	var instance ServiceInstance
	if err := json.NewDecoder(r.Body).Decode(&instance); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Invalid JSON format")
		return
	}
	
	// Set registration timestamp
	instance.LastHeartbeat = time.Now()
	instance.Health = true
	
	// Register the service
	registry.Register(instance)
	
	// Send success response
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, `{"status": "registered"}`)
}

// Discover handler
func discoverHandler(registry *ServiceRegistry, w http.ResponseWriter, r *http.Request) {
	// Extract service name from URL path
	serviceName := strings.TrimPrefix(r.URL.Path, "/discover/")
	
	if serviceName == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Service name required")
		return
	}
	
	// Discover service instances
	instances := registry.Discover(serviceName)
	
	// Return instances as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(instances)
}

// List services handler
func listServicesHandler(registry *ServiceRegistry, w http.ResponseWriter, r *http.Request) {
	// Get all services
	services := registry.GetAllServices()
	
	// Return services as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(services)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, `{"status": "healthy", "service": "registry"}`)
}

// Demo service to simulate microservice instances
func startDemoService(serviceName string, port int) {
	time.Sleep(2 * time.Second) // Wait for registry to start
	
	// Register with service registry
	instance := ServiceInstance{
		ID:            fmt.Sprintf("%s-%d", serviceName, port),
		Name:          serviceName,
		Host:          "localhost",
		Port:          port,
		Health:        true,
		LastHeartbeat: time.Now(),
	}
	
	// Register instance
	go func() {
		for {
			instanceJSON, _ := json.Marshal(instance)
			resp, err := http.Post("http://localhost:8080/register", "application/json", 
				bytes.NewReader(instanceJSON))
			if err == nil {
				resp.Body.Close()
			}
			time.Sleep(30 * time.Second) // Re-register every 30 seconds
		}
	}()
	
	// Start demo service
	mux := http.NewServeMux()
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"status": "healthy", "service": "%s", "port": %d}`, serviceName, port)
	})
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"message": "Hello from %s", "port": %d, "random": %d}`, 
			serviceName, port, rand.Intn(1000))
	})
	
	log.Printf("Demo service %s starting on :%d", serviceName, port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), mux))
}