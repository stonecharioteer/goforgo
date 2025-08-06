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

// TODO: Service instance structure
type ServiceInstance struct {
	/* define fields: ID, Name, Host, Port, Health, LastHeartbeat */
}

// TODO: Service registry structure
type ServiceRegistry struct {
	/* define fields: services map, mutex for thread safety */
}

func main() {
	fmt.Println("=== Microservice Discovery ===")
	
	// TODO: Create service registry
	registry := /* create new service registry */
	
	// TODO: Start health check routine
	/* start health checker in goroutine */
	
	// Setup HTTP routes
	http.HandleFunc("/register", /* create register handler with registry */)
	http.HandleFunc("/discover/", /* create discover handler with registry */)
	http.HandleFunc("/services", /* create list services handler with registry */)
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

// TODO: Create new service registry
func newServiceRegistry() *ServiceRegistry {
	return &ServiceRegistry{
		/* initialize fields */
	}
}

// TODO: Register service instance
func (sr *ServiceRegistry) Register(instance ServiceInstance) {
	/* lock mutex */
	/* defer unlock mutex */
	
	// TODO: Initialize service map if needed
	if /* check if services map is nil */ {
		/* initialize services map */
	}
	
	// TODO: Initialize service slice if needed
	if /* check if service name doesn't exist in services map */ {
		/* initialize empty slice for this service */
	}
	
	// TODO: Update existing instance or add new one
	found := false
	for i := range /* iterate over service instances */ {
		if /* check if instance ID matches */ {
			/* update existing instance */
			found = true
			break
		}
	}
	
	if /* check if not found */ {
		/* append new instance to service slice */
	}
	
	/* log service registration */
}

// TODO: Discover healthy service instances
func (sr *ServiceRegistry) Discover(serviceName string) []ServiceInstance {
	/* lock mutex */
	/* defer unlock mutex */
	
	instances, exists := /* get instances from services map */
	if /* check if service doesn't exist */ {
		return nil
	}
	
	// TODO: Filter healthy instances
	var healthy []ServiceInstance
	for _, instance := range instances {
		if /* check if instance is healthy */ {
			/* append to healthy slice */
		}
	}
	
	return healthy
}

// TODO: Get all services
func (sr *ServiceRegistry) GetAllServices() map[string][]ServiceInstance {
	/* lock mutex */
	/* defer unlock mutex */
	
	// TODO: Create copy of services map
	result := make(map[string][]ServiceInstance)
	for name, instances := range /* iterate over services */ {
		/* copy instances to result */
	}
	
	return result
}

// TODO: Health check routine
func (sr *ServiceRegistry) startHealthChecker() {
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()
	
	for range ticker.C {
		/* lock mutex */
		
		for serviceName, instances := range /* iterate over services */ {
			for i := range instances {
				// TODO: Check if instance is stale
				if /* check if last heartbeat is older than 1 minute */ {
					/* mark instance as unhealthy */
					/* log unhealthy instance */
				}
			}
		}
		
		/* unlock mutex */
	}
}

// TODO: Register handler
func registerHandler(registry *ServiceRegistry, w http.ResponseWriter, r *http.Request) {
	if /* check if method is not POST */ {
		/* write method not allowed status */
		return
	}
	
	// TODO: Parse service instance from request
	var instance ServiceInstance
	if err := /* decode JSON from request body */; err != nil {
		/* write bad request status */
		/* write error message */
		return
	}
	
	// TODO: Set registration timestamp
	/* set LastHeartbeat to current time */
	/* set Health to true */
	
	// TODO: Register the service
	/* call registry Register method */
	
	// TODO: Send success response
	/* write OK status */
	/* write success message as JSON */
}

// TODO: Discover handler
func discoverHandler(registry *ServiceRegistry, w http.ResponseWriter, r *http.Request) {
	// TODO: Extract service name from URL path
	serviceName := /* extract service name from r.URL.Path after /discover/ */
	
	if /* check if service name is empty */ {
		/* write bad request status */
		/* write error message */
		return
	}
	
	// TODO: Discover service instances
	instances := /* call registry Discover method */
	
	// TODO: Return instances as JSON
	/* set content type to application/json */
	/* encode instances as JSON and write response */
}

// TODO: List services handler
func listServicesHandler(registry *ServiceRegistry, w http.ResponseWriter, r *http.Request) {
	// TODO: Get all services
	services := /* call registry GetAllServices method */
	
	// TODO: Return services as JSON
	/* set content type to application/json */
	/* encode services as JSON and write response */
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