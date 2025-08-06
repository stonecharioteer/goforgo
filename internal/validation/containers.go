package validation

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	_ "github.com/lib/pq"
)

// PostgreSQLContainer wraps a testcontainers PostgreSQL instance
type PostgreSQLContainer struct {
	container  testcontainers.Container
	name       string
	version    string
	config     map[string]interface{}
	connection *ServiceConnectionInfo
}

// NewPostgreSQLContainer creates a new PostgreSQL container service
func NewPostgreSQLContainer(name, version string, config map[string]interface{}) *PostgreSQLContainer {
	return &PostgreSQLContainer{
		name:    name,
		version: version,
		config:  config,
	}
}

// Start starts the PostgreSQL container
func (p *PostgreSQLContainer) Start(ctx context.Context) error {
	log.Printf("🐘 Starting PostgreSQL container: %s (version: %s)", p.name, p.version)

	// Set default values
	database := "testdb"
	username := "testuser"
	password := "testpass"

	// Override from config if provided
	if db, ok := p.config["POSTGRES_DB"].(string); ok {
		database = db
	}
	if user, ok := p.config["POSTGRES_USER"].(string); ok {
		username = user
	}
	if pass, ok := p.config["POSTGRES_PASSWORD"].(string); ok {
		password = pass
	}

	// Create container request
	req := testcontainers.ContainerRequest{
		Image:        fmt.Sprintf("postgres:%s", p.version),
		ExposedPorts: []string{"5432/tcp"},
		Env: map[string]string{
			"POSTGRES_DB":       database,
			"POSTGRES_USER":     username,
			"POSTGRES_PASSWORD": password,
		},
		WaitingFor: wait.ForAll(
			wait.ForLog("database system is ready to accept connections").WithOccurrence(2),
			wait.ForExposedPort(),
		).WithDeadline(60 * time.Second),
		Name: p.name,
	}

	// Start the container
	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		return fmt.Errorf("failed to start PostgreSQL container: %w", err)
	}

	p.container = container

	// Get connection details
	host, err := container.Host(ctx)
	if err != nil {
		return fmt.Errorf("failed to get container host: %w", err)
	}

	port, err := container.MappedPort(ctx, "5432")
	if err != nil {
		return fmt.Errorf("failed to get container port: %w", err)
	}

	// Create connection info
	p.connection = &ServiceConnectionInfo{
		Host:     host,
		Port:     port.Int(),
		Database: database,
		Username: username,
		Password: password,
		URL:      fmt.Sprintf("postgresql://%s:%s@%s:%d/%s", username, password, host, port.Int(), database),
		Env: map[string]string{
			"DB_HOST":     host,
			"DB_PORT":     port.Port(),
			"DB_NAME":     database,
			"DB_USER":     username,
			"DB_PASSWORD": password,
			"DATABASE_URL": fmt.Sprintf("postgresql://%s:%s@%s:%d/%s", username, password, host, port.Int(), database),
		},
	}

	log.Printf("✅ PostgreSQL container started at %s:%d", host, port.Int())
	return nil
}

// Stop stops the PostgreSQL container
func (p *PostgreSQLContainer) Stop(ctx context.Context) error {
	if p.container == nil {
		return nil
	}

	log.Printf("🛑 Stopping PostgreSQL container: %s", p.name)
	return p.container.Terminate(ctx)
}

// IsReady checks if PostgreSQL is ready to accept connections
func (p *PostgreSQLContainer) IsReady(ctx context.Context) (bool, error) {
	if p.connection == nil {
		return false, fmt.Errorf("connection info not available")
	}

	// Try to connect to the database
	db, err := sql.Open("postgres", p.connection.URL)
	if err != nil {
		return false, err
	}
	defer db.Close()

	// Test the connection
	if err := db.PingContext(ctx); err != nil {
		return false, err
	}

	return true, nil
}

// GetConnectionInfo returns the connection information
func (p *PostgreSQLContainer) GetConnectionInfo() *ServiceConnectionInfo {
	return p.connection
}

// GetServiceType returns the service type
func (p *PostgreSQLContainer) GetServiceType() string {
	return "postgresql"
}

// GetServiceName returns the service name
func (p *PostgreSQLContainer) GetServiceName() string {
	return p.name
}

// RedisContainer wraps a testcontainers Redis instance
type RedisContainer struct {
	container  testcontainers.Container
	name       string
	version    string
	config     map[string]interface{}
	connection *ServiceConnectionInfo
}

// NewRedisContainer creates a new Redis container service
func NewRedisContainer(name, version string, config map[string]interface{}) *RedisContainer {
	return &RedisContainer{
		name:    name,
		version: version,
		config:  config,
	}
}

// Start starts the Redis container
func (r *RedisContainer) Start(ctx context.Context) error {
	log.Printf("🔴 Starting Redis container: %s (version: %s)", r.name, r.version)

	// Create container request
	req := testcontainers.ContainerRequest{
		Image:        fmt.Sprintf("redis:%s", r.version),
		ExposedPorts: []string{"6379/tcp"},
		WaitingFor:   wait.ForLog("Ready to accept connections"),
		Name:         r.name,
	}

	// Add custom config if provided
	if maxmemory, ok := r.config["maxmemory"].(string); ok {
		req.Cmd = []string{"redis-server", "--maxmemory", maxmemory}
	}

	// Start the container
	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		return fmt.Errorf("failed to start Redis container: %w", err)
	}

	r.container = container

	// Get connection details
	host, err := container.Host(ctx)
	if err != nil {
		return fmt.Errorf("failed to get container host: %w", err)
	}

	port, err := container.MappedPort(ctx, "6379")
	if err != nil {
		return fmt.Errorf("failed to get container port: %w", err)
	}

	// Create connection info
	r.connection = &ServiceConnectionInfo{
		Host: host,
		Port: port.Int(),
		URL:  fmt.Sprintf("redis://%s:%d", host, port.Int()),
		Env: map[string]string{
			"REDIS_HOST": host,
			"REDIS_PORT": port.Port(),
			"REDIS_URL":  fmt.Sprintf("redis://%s:%d", host, port.Int()),
		},
	}

	log.Printf("✅ Redis container started at %s:%d", host, port.Int())
	return nil
}

// Stop stops the Redis container
func (r *RedisContainer) Stop(ctx context.Context) error {
	if r.container == nil {
		return nil
	}

	log.Printf("🛑 Stopping Redis container: %s", r.name)
	return r.container.Terminate(ctx)
}

// IsReady checks if Redis is ready
func (r *RedisContainer) IsReady(ctx context.Context) (bool, error) {
	if r.connection == nil {
		return false, fmt.Errorf("connection info not available")
	}

	// For now, assume Redis is ready if the container started
	// TODO: Implement actual Redis ping check
	return true, nil
}

// GetConnectionInfo returns the connection information
func (r *RedisContainer) GetConnectionInfo() *ServiceConnectionInfo {
	return r.connection
}

// GetServiceType returns the service type
func (r *RedisContainer) GetServiceType() string {
	return "redis"
}

// GetServiceName returns the service name
func (r *RedisContainer) GetServiceName() string {
	return r.name
}