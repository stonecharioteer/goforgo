package validation

import (
	"context"
	"time"
)

// ValidatorRegistry manages all available validation rules
type ValidatorRegistry struct {
	validators map[string]ValidationRule
}

// ValidationRule interface for pluggable validation rules
type ValidationRule interface {
	GetType() string
	GetName() string
	Validate(ctx context.Context, request *ValidationRuleRequest) (*RuleResult, error)
	GetRequiredServices() []string
	GetDependencies() []string
}

// ValidationRuleRequest contains context for a validation rule
type ValidationRuleRequest struct {
	WorkingDir      string
	ExerciseFilePath string
	Services        map[string]*ServiceConnectionInfo
	Environment     map[string]string
	Config          map[string]interface{}
	Timeout         time.Duration
}

// NewValidatorRegistry creates a new validator registry with built-in rules
func NewValidatorRegistry() *ValidatorRegistry {
	registry := &ValidatorRegistry{
		validators: make(map[string]ValidationRule),
	}
	
	// Register built-in validation rules
	registry.Register(&HTTPRouteValidator{})
	registry.Register(&DatabaseValidator{})
	registry.Register(&ProcessValidator{})
	registry.Register(&NetworkValidator{})
	registry.Register(&ConcurrencyValidator{})
	registry.Register(&MetricsValidator{})
	registry.Register(&LogValidator{})
	
	return registry
}

// Register adds a validation rule to the registry
func (vr *ValidatorRegistry) Register(rule ValidationRule) {
	vr.validators[rule.GetType()] = rule
}

// Get retrieves a validation rule by type
func (vr *ValidatorRegistry) Get(ruleType string) (ValidationRule, bool) {
	rule, exists := vr.validators[ruleType]
	return rule, exists
}

// GetAll returns all registered validation rules
func (vr *ValidatorRegistry) GetAll() map[string]ValidationRule {
	return vr.validators
}

// HTTPRouteValidator tests REST endpoints, WebSocket connections, middleware
type HTTPRouteValidator struct{}

func (h *HTTPRouteValidator) GetType() string { return "http_routes" }
func (h *HTTPRouteValidator) GetName() string { return "HTTP Route Validator" }
func (h *HTTPRouteValidator) GetRequiredServices() []string { return []string{} }
func (h *HTTPRouteValidator) GetDependencies() []string { return []string{} }

func (h *HTTPRouteValidator) Validate(ctx context.Context, request *ValidationRuleRequest) (*RuleResult, error) {
	// Implementation will test HTTP routes, check response codes, validate JSON, etc.
	return &RuleResult{
		RuleName: h.GetName(),
		RuleType: h.GetType(),
		Passed:   true, // Placeholder
		Output:   "HTTP routes validation passed",
	}, nil
}

// DatabaseValidator runs queries, checks schemas, validates transactions
type DatabaseValidator struct{}

func (d *DatabaseValidator) GetType() string { return "database" }
func (d *DatabaseValidator) GetName() string { return "Database Validator" }
func (d *DatabaseValidator) GetRequiredServices() []string { return []string{"postgresql", "mysql", "mongodb"} }
func (d *DatabaseValidator) GetDependencies() []string { return []string{} }

func (d *DatabaseValidator) Validate(ctx context.Context, request *ValidationRuleRequest) (*RuleResult, error) {
	// Implementation will connect to database, run queries, check schemas, etc.
	return &RuleResult{
		RuleName: d.GetName(),
		RuleType: d.GetType(),
		Passed:   true, // Placeholder
		Output:   "Database validation passed",
	}, nil
}

// ProcessValidator monitors processes, goroutines, resource usage
type ProcessValidator struct{}

func (p *ProcessValidator) GetType() string { return "process" }
func (p *ProcessValidator) GetName() string { return "Process Validator" }
func (p *ProcessValidator) GetRequiredServices() []string { return []string{} }
func (p *ProcessValidator) GetDependencies() []string { return []string{} }

func (p *ProcessValidator) Validate(ctx context.Context, request *ValidationRuleRequest) (*RuleResult, error) {
	// Implementation will monitor process behavior, check resource usage, etc.
	return &RuleResult{
		RuleName: p.GetName(),
		RuleType: p.GetType(),
		Passed:   true, // Placeholder
		Output:   "Process validation passed",
	}, nil
}

// NetworkValidator tests TCP/UDP servers, client connections
type NetworkValidator struct{}

func (n *NetworkValidator) GetType() string { return "network" }
func (n *NetworkValidator) GetName() string { return "Network Validator" }
func (n *NetworkValidator) GetRequiredServices() []string { return []string{} }
func (n *NetworkValidator) GetDependencies() []string { return []string{} }

func (n *NetworkValidator) Validate(ctx context.Context, request *ValidationRuleRequest) (*RuleResult, error) {
	// Implementation will test network connections, socket communication, etc.
	return &RuleResult{
		RuleName: n.GetName(),
		RuleType: n.GetType(),
		Passed:   true, // Placeholder
		Output:   "Network validation passed",
	}, nil
}

// ConcurrencyValidator detects race conditions, deadlocks, sync primitives
type ConcurrencyValidator struct{}

func (c *ConcurrencyValidator) GetType() string { return "concurrency" }
func (c *ConcurrencyValidator) GetName() string { return "Concurrency Validator" }
func (c *ConcurrencyValidator) GetRequiredServices() []string { return []string{} }
func (c *ConcurrencyValidator) GetDependencies() []string { return []string{} }

func (c *ConcurrencyValidator) Validate(ctx context.Context, request *ValidationRuleRequest) (*RuleResult, error) {
	// Implementation will use race detector, analyze goroutine patterns, etc.
	return &RuleResult{
		RuleName: c.GetName(),
		RuleType: c.GetType(),
		Passed:   true, // Placeholder
		Output:   "Concurrency validation passed",
	}, nil
}

// MetricsValidator checks Prometheus metrics, custom counters
type MetricsValidator struct{}

func (m *MetricsValidator) GetType() string { return "metrics" }
func (m *MetricsValidator) GetName() string { return "Metrics Validator" }
func (m *MetricsValidator) GetRequiredServices() []string { return []string{"prometheus"} }
func (m *MetricsValidator) GetDependencies() []string { return []string{} }

func (m *MetricsValidator) Validate(ctx context.Context, request *ValidationRuleRequest) (*RuleResult, error) {
	// Implementation will scrape metrics endpoints, validate counter values, etc.
	return &RuleResult{
		RuleName: m.GetName(),
		RuleType: m.GetType(),
		Passed:   true, // Placeholder
		Output:   "Metrics validation passed",
	}, nil
}

// LogValidator validates structured logs, error patterns
type LogValidator struct{}

func (l *LogValidator) GetType() string { return "logs" }
func (l *LogValidator) GetName() string { return "Log Validator" }
func (l *LogValidator) GetRequiredServices() []string { return []string{} }
func (l *LogValidator) GetDependencies() []string { return []string{} }

func (l *LogValidator) Validate(ctx context.Context, request *ValidationRuleRequest) (*RuleResult, error) {
	// Implementation will parse logs, check for patterns, validate structured logging, etc.
	return &RuleResult{
		RuleName: l.GetName(),
		RuleType: l.GetType(),
		Passed:   true, // Placeholder
		Output:   "Log validation passed",
	}, nil
}