package validation

import (
	"context"
	"fmt"
	"log"
	"sync"
)

// NewServiceRegistry creates a new service registry with default configuration
func NewServiceRegistry() *ServiceRegistry {
	config := &ServiceRegistryConfig{
		NetworkName:       "goforgo-validation",
		DefaultPullPolicy: "IfNotPresent",
		LogLevel:          "info",
	}

	return &ServiceRegistry{
		services: make(map[string]Service),
		config:   config,
	}
}

// CreateService creates a service based on the provided specification
func (sr *ServiceRegistry) CreateService(ctx context.Context, spec ServiceSpec) (Service, error) {
	log.Printf("Creating service: %s (%s)", spec.Name, spec.Type)

	switch spec.Type {
	case "postgresql":
		return sr.createPostgreSQLService(ctx, spec)
	case "redis":
		return sr.createRedisService(ctx, spec)
	case "mongodb":
		return sr.createMongoDBService(ctx, spec)
	case "rabbitmq":
		return sr.createRabbitMQService(ctx, spec)
	case "http_mock":
		return sr.createHTTPMockService(ctx, spec)
	default:
		return nil, fmt.Errorf("unsupported service type: %s", spec.Type)
	}
}

// GetService retrieves a running service by name
func (sr *ServiceRegistry) GetService(name string) (Service, bool) {
	service, exists := sr.services[name]
	return service, exists
}

// RegisterService registers a service in the registry
func (sr *ServiceRegistry) RegisterService(name string, service Service) {
	sr.services[name] = service
}

// StopAllServices stops all registered services
func (sr *ServiceRegistry) StopAllServices(ctx context.Context) error {
	var wg sync.WaitGroup
	errors := make(chan error, len(sr.services))

	for name, service := range sr.services {
		wg.Add(1)
		go func(serviceName string, svc Service) {
			defer wg.Done()
			log.Printf("Stopping service: %s", serviceName)
			if err := svc.Stop(ctx); err != nil {
				errors <- fmt.Errorf("failed to stop service %s: %w", serviceName, err)
			}
		}(name, service)
	}

	wg.Wait()
	close(errors)

	// Collect any errors
	var errorList []error
	for err := range errors {
		errorList = append(errorList, err)
	}

	if len(errorList) > 0 {
		return fmt.Errorf("failed to stop some services: %v", errorList)
	}

	return nil
}

// Service factory methods (these will use testcontainers-go in the actual implementation)

func (sr *ServiceRegistry) createPostgreSQLService(ctx context.Context, spec ServiceSpec) (Service, error) {
	log.Printf("Creating PostgreSQL service with spec: %+v", spec)

	// Use the actual testcontainers implementation
	service := NewPostgreSQLContainer(spec.Name, spec.Version, spec.Config)

	sr.RegisterService(spec.Name, service)
	return service, nil
}

func (sr *ServiceRegistry) createRedisService(ctx context.Context, spec ServiceSpec) (Service, error) {
	log.Printf("Creating Redis service with spec: %+v", spec)

	// Use the actual testcontainers implementation
	service := NewRedisContainer(spec.Name, spec.Version, spec.Config)

	sr.RegisterService(spec.Name, service)
	return service, nil
}

func (sr *ServiceRegistry) createMongoDBService(ctx context.Context, spec ServiceSpec) (Service, error) {
	log.Printf("Creating MongoDB service with spec: %+v", spec)

	// TODO: Implement using testcontainers-go MongoDB container
	service := &MongoDBService{
		name:    spec.Name,
		version: spec.Version,
		config:  spec.Config,
	}

	sr.RegisterService(spec.Name, service)
	return service, nil
}

func (sr *ServiceRegistry) createRabbitMQService(ctx context.Context, spec ServiceSpec) (Service, error) {
	log.Printf("Creating RabbitMQ service with spec: %+v", spec)

	// TODO: Implement using testcontainers-go RabbitMQ container
	service := &RabbitMQService{
		name:    spec.Name,
		version: spec.Version,
		config:  spec.Config,
	}

	sr.RegisterService(spec.Name, service)
	return service, nil
}

func (sr *ServiceRegistry) createHTTPMockService(ctx context.Context, spec ServiceSpec) (Service, error) {
	log.Printf("Creating HTTP Mock service with spec: %+v", spec)

	// TODO: Implement using testcontainers-go generic container with mock server
	service := &HTTPMockService{
		name:    spec.Name,
		version: spec.Version,
		config:  spec.Config,
	}

	sr.RegisterService(spec.Name, service)
	return service, nil
}

// MongoDBService represents a MongoDB database service (not yet implemented)
type MongoDBService struct {
	name    string
	version string
	config  map[string]interface{}
}

func (m *MongoDBService) Start(ctx context.Context) error {
	return fmt.Errorf("service %q: mongodb not yet implemented", m.name)
}

func (m *MongoDBService) Stop(ctx context.Context) error {
	log.Printf("Stopping MongoDB service: %s", m.name)
	// TODO: Stop testcontainer
	return nil
}

func (m *MongoDBService) IsReady(ctx context.Context) (bool, error) {
	return false, nil
}

func (m *MongoDBService) GetConnectionInfo() *ServiceConnectionInfo {
	// TODO: Extract real connection info from testcontainer
	return &ServiceConnectionInfo{
		Host:     "localhost",
		Port:     27017,
		Database: "testdb",
		URL:      "mongodb://localhost:27017/testdb",
		Env: map[string]string{
			"MONGO_HOST": "localhost",
			"MONGO_PORT": "27017",
			"MONGO_DB":   "testdb",
			"MONGO_URL":  "mongodb://localhost:27017/testdb",
		},
	}
}

func (m *MongoDBService) GetServiceType() string { return "mongodb" }
func (m *MongoDBService) GetServiceName() string { return m.name }

// RabbitMQService represents a RabbitMQ message queue service
type RabbitMQService struct {
	name    string
	version string
	config  map[string]interface{}
}

func (r *RabbitMQService) Start(ctx context.Context) error {
	return fmt.Errorf("service %q: rabbitmq not yet implemented", r.name)
}

func (r *RabbitMQService) Stop(ctx context.Context) error {
	log.Printf("Stopping RabbitMQ service: %s", r.name)
	// TODO: Stop testcontainer
	return nil
}

func (r *RabbitMQService) IsReady(ctx context.Context) (bool, error) {
	return false, nil
}

func (r *RabbitMQService) GetConnectionInfo() *ServiceConnectionInfo {
	// TODO: Extract real connection info from testcontainer
	return &ServiceConnectionInfo{
		Host:     "localhost",
		Port:     5672,
		Username: "guest",
		Password: "guest",
		URL:      "amqp://guest:guest@localhost:5672/",
		Env: map[string]string{
			"RABBITMQ_HOST": "localhost",
			"RABBITMQ_PORT": "5672",
			"RABBITMQ_USER": "guest",
			"RABBITMQ_PASS": "guest",
			"RABBITMQ_URL":  "amqp://guest:guest@localhost:5672/",
		},
	}
}

func (r *RabbitMQService) GetServiceType() string { return "rabbitmq" }
func (r *RabbitMQService) GetServiceName() string { return r.name }

// HTTPMockService represents a mock HTTP service for testing
type HTTPMockService struct {
	name    string
	version string
	config  map[string]interface{}
}

func (h *HTTPMockService) Start(ctx context.Context) error {
	return fmt.Errorf("service %q: http_mock not yet implemented", h.name)
}

func (h *HTTPMockService) Stop(ctx context.Context) error {
	log.Printf("Stopping HTTP Mock service: %s", h.name)
	// TODO: Stop testcontainer
	return nil
}

func (h *HTTPMockService) IsReady(ctx context.Context) (bool, error) {
	return false, nil
}

func (h *HTTPMockService) GetConnectionInfo() *ServiceConnectionInfo {
	// TODO: Extract real connection info from testcontainer
	port := 8080
	if configPort, ok := h.config["port"].(int); ok {
		port = configPort
	}

	return &ServiceConnectionInfo{
		Host: "localhost",
		Port: port,
		URL:  fmt.Sprintf("http://localhost:%d", port),
		Env: map[string]string{
			"HTTP_MOCK_HOST": "localhost",
			"HTTP_MOCK_PORT": fmt.Sprintf("%d", port),
			"HTTP_MOCK_URL":  fmt.Sprintf("http://localhost:%d", port),
		},
	}
}

func (h *HTTPMockService) GetServiceType() string { return "http_mock" }
func (h *HTTPMockService) GetServiceName() string { return h.name }
