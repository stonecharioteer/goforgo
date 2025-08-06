package validation

import (
	"context"
	"time"

	"github.com/stonecharioteer/goforgo/internal/exercise"
)

// TestOrchestrator is the main validation engine that orchestrates all testing
type TestOrchestrator struct {
	serviceRegistry *ServiceRegistry
	validatorRegistry *ValidatorRegistry
	resourceManager *ResourceManager
	config *OrchestratorConfig
}

// OrchestratorConfig holds configuration for the test orchestrator
type OrchestratorConfig struct {
	MaxConcurrentServices int
	DefaultTimeout        time.Duration
	CleanupTimeout        time.Duration
	ContainerNetworkName  string
}

// ServiceRegistry manages lifecycle of supporting services (databases, queues, external APIs)
type ServiceRegistry struct {
	services map[string]Service
	network  ContainerNetwork
	config   *ServiceRegistryConfig
}

// ServiceRegistryConfig holds configuration for the service registry
type ServiceRegistryConfig struct {
	NetworkName       string
	DefaultPullPolicy string
	LogLevel          string
}

// Service represents a managed service (container or external service)
type Service interface {
	Start(ctx context.Context) error
	Stop(ctx context.Context) error
	IsReady(ctx context.Context) (bool, error)
	GetConnectionInfo() *ServiceConnectionInfo
	GetServiceType() string
	GetServiceName() string
}

// ServiceConnectionInfo contains connection details for a service
type ServiceConnectionInfo struct {
	Host     string
	Port     int
	Database string
	Username string
	Password string
	URL      string
	Env      map[string]string
}

// ValidationRequest represents a validation request for an exercise
type ValidationRequest struct {
	Exercise      *exercise.Exercise
	WorkingDir    string
	Services      []ServiceSpec
	Rules         []ValidationRuleSpec
	Environment   map[string]string
	Timeout       time.Duration
}

// ServiceSpec defines requirements for a service
type ServiceSpec struct {
	Type       string                 `toml:"type"`       // "postgresql", "redis", "mongodb", etc.
	Name       string                 `toml:"name"`       // unique name for this service instance
	Version    string                 `toml:"version"`    // container version/tag
	Config     map[string]interface{} `toml:"config"`     // service-specific configuration
	WaitFor    []WaitCondition        `toml:"wait_for"`   // conditions to wait for before considering ready
	Fixtures   []string               `toml:"fixtures"`   // files to load into service
	Persistent bool                   `toml:"persistent"` // whether to reuse across exercises
}

// WaitCondition defines a condition to wait for
type WaitCondition struct {
	Type    string        `toml:"type"`    // "port", "log", "http", "tcp"
	Target  string        `toml:"target"`  // what to check
	Timeout time.Duration `toml:"timeout"` // how long to wait
}

// ValidationRuleSpec defines a validation rule to execute
type ValidationRuleSpec struct {
	Type     string                 `toml:"type"`     // "http_routes", "database", "process", etc.
	Name     string                 `toml:"name"`     // unique name for this rule
	Config   map[string]interface{} `toml:"config"`   // rule-specific configuration
	DependsOn []string              `toml:"depends_on"` // other rules that must pass first
	Parallel bool                   `toml:"parallel"`  // can run in parallel with others
}

// ValidationResult contains the results of validation
type ValidationResult struct {
	Success          bool                        `json:"success"`
	Duration         time.Duration               `json:"duration"`
	ServiceResults   map[string]*ServiceResult   `json:"service_results"`
	ValidationResults map[string]*RuleResult     `json:"validation_results"`
	Environment      map[string]string           `json:"environment"`
	Error            string                      `json:"error,omitempty"`
	Logs             []string                    `json:"logs,omitempty"`
}

// ServiceResult contains results from service management
type ServiceResult struct {
	ServiceName string            `json:"service_name"`
	ServiceType string            `json:"service_type"`
	Started     bool              `json:"started"`
	Ready       bool              `json:"ready"`
	Connection  *ServiceConnectionInfo `json:"connection"`
	Duration    time.Duration     `json:"duration"`
	Error       string            `json:"error,omitempty"`
	Logs        []string          `json:"logs,omitempty"`
}

// RuleResult contains results from validation rules
type RuleResult struct {
	RuleName string        `json:"rule_name"`
	RuleType string        `json:"rule_type"`
	Passed   bool          `json:"passed"`
	Duration time.Duration `json:"duration"`
	Output   string        `json:"output"`
	Error    string        `json:"error,omitempty"`
	Details  interface{}   `json:"details,omitempty"`
}

// ContainerNetwork manages docker network for services
type ContainerNetwork interface {
	Create(ctx context.Context, name string) error
	Remove(ctx context.Context, name string) error
	GetName() string
}

// ResourceManager handles cleanup and resource management across all test scenarios
type ResourceManager struct {
	activeServices map[string]Service
	activeNetworks map[string]ContainerNetwork
	cleanupTasks   []CleanupTask
}

// CleanupTask represents a cleanup operation
type CleanupTask struct {
	Name     string
	Priority int
	Execute  func(ctx context.Context) error
}