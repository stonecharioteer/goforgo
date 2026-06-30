//go:build integration

package validation

import (
	"context"
	"testing"
	"time"
)

func TestServiceRegistry_CreatePostgreSQLService(t *testing.T) {
	registry := NewServiceRegistry()

	spec := ServiceSpec{
		Type:    "postgresql",
		Name:    "test_postgres",
		Version: "15",
		Config: map[string]interface{}{
			"POSTGRES_DB":       "testdb",
			"POSTGRES_USER":     "testuser",
			"POSTGRES_PASSWORD": "testpass",
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()

	service, err := registry.CreateService(ctx, spec)
	if err != nil {
		t.Fatalf("Failed to create PostgreSQL service: %v", err)
	}

	// Start the service
	if err := service.Start(ctx); err != nil {
		t.Fatalf("Failed to start PostgreSQL service: %v", err)
	}

	// Check if ready
	ready, err := service.IsReady(ctx)
	if err != nil {
		t.Fatalf("Failed to check service readiness: %v", err)
	}

	if !ready {
		t.Fatal("PostgreSQL service should be ready")
	}

	// Get connection info
	connInfo := service.GetConnectionInfo()
	if connInfo == nil {
		t.Fatal("Expected connection info, got nil")
	}

	t.Logf("PostgreSQL service ready at %s:%d", connInfo.Host, connInfo.Port)
	t.Logf("Connection URL: %s", connInfo.URL)

	// Clean up
	if err := service.Stop(ctx); err != nil {
		t.Errorf("Failed to stop PostgreSQL service: %v", err)
	}
}
