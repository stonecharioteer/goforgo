package validation

import "time"

// EnhancedExerciseValidation extends the basic validation with universal validation support
type EnhancedExerciseValidation struct {
	// Legacy fields for backward compatibility
	Mode           string   `toml:"mode"`           // "build", "test", "run", "static", "universal"
	Timeout        string   `toml:"timeout"`        // e.g., "30s"
	ExpectedOutput string   `toml:"expected_output,omitempty"`
	StaticCheck    string   `toml:"static_check,omitempty"`
	RequiredFiles  []string `toml:"required_files,omitempty"`
	
	// New universal validation fields
	Services       []ServiceSpec         `toml:"services,omitempty"`       // Required services
	Rules          []ValidationRuleSpec  `toml:"rules,omitempty"`          // Validation rules to execute
	Environment    map[string]string     `toml:"environment,omitempty"`    // Environment variables
	WorkingDir     string                `toml:"working_dir,omitempty"`    // Working directory override
	SetupScript    string                `toml:"setup_script,omitempty"`   // Setup script to run before validation
	TeardownScript string                `toml:"teardown_script,omitempty"` // Cleanup script to run after validation
	Parallel       bool                  `toml:"parallel,omitempty"`       // Can run validation rules in parallel
}

// Example TOML configurations for different exercise types:

// HTTP Server Exercise Example:
/*
[validation]
mode = "universal"
timeout = "60s"

[[validation.services]]
type = "postgresql"
name = "main_db"
version = "15"
config = { POSTGRES_DB = "testdb", POSTGRES_USER = "testuser", POSTGRES_PASSWORD = "testpass" }
fixtures = ["schema.sql", "seed_data.sql"]

[[validation.rules]]
type = "http_routes"
name = "api_endpoints"
config = { 
    base_url = "http://localhost:8080",
    routes = [
        { method = "GET", path = "/api/users", expect_status = 200, expect_json = true },
        { method = "POST", path = "/api/users", body = "test_user.json", expect_status = 201 },
        { method = "GET", path = "/health", expect_status = 200, expect_body = "OK" }
    ]
}

[[validation.rules]]
type = "database"
name = "data_persistence"
config = {
    service = "main_db",
    queries = [
        { query = "SELECT COUNT(*) FROM users", expect_result = "3" },
        { query = "SELECT name FROM users WHERE id = 1", expect_result = "Alice" }
    ]
}
depends_on = ["api_endpoints"]

[validation.environment]
DB_HOST = "${services.main_db.host}"
DB_PORT = "${services.main_db.port}"
DB_NAME = "${services.main_db.database}"
*/

// Microservices Circuit Breaker Exercise Example:
/*
[validation]
mode = "universal"
timeout = "90s"

[[validation.services]]
type = "redis"
name = "cache"
version = "7"
config = { maxmemory = "128mb" }

[[validation.services]]
type = "http_mock"
name = "upstream_service"
version = "latest"
config = { 
    port = 9090,
    failure_rate = 0.5,
    failure_after = "10s"
}

[[validation.rules]]
type = "process"
name = "circuit_breaker_states"
config = {
    monitor_duration = "30s",
    expected_states = ["CLOSED", "OPEN", "HALF_OPEN"],
    state_transitions = true
}

[[validation.rules]]
type = "metrics"
name = "circuit_breaker_metrics"
config = {
    endpoint = "http://localhost:8080/metrics",
    metrics = [
        { name = "circuit_breaker_requests_total", type = "counter" },
        { name = "circuit_breaker_state", type = "gauge" },
        { name = "circuit_breaker_failures_total", type = "counter" }
    ]
}

[[validation.rules]]
type = "concurrency"
name = "thread_safety"
config = {
    concurrent_requests = 100,
    duration = "15s",
    race_detection = true
}
*/

// gRPC Streaming Exercise Example:
/*
[validation]
mode = "universal" 
timeout = "45s"
setup_script = "compile_proto.sh"

[[validation.rules]]
type = "grpc_service"
name = "bidirectional_streaming"
config = {
    service = "ChatService",
    method = "Chat",
    endpoint = "localhost:50051",
    test_scenarios = [
        {
            type = "bidirectional_stream",
            send_messages = ["Hello", "World", "Test"],
            expect_responses = 3,
            expect_order = true
        }
    ]
}

[[validation.rules]]
type = "grpc_interceptors"
name = "logging_interceptor"
config = {
    interceptor_type = "unary",
    expected_logs = ["Request received", "Request processed", "Response sent"]
}
*/

// Database Connection Pool Exercise Example:
/*
[validation]
mode = "universal"
timeout = "120s"

[[validation.services]]
type = "postgresql"
name = "main_db"
version = "15"
config = { 
    POSTGRES_DB = "pooltest", 
    POSTGRES_USER = "testuser", 
    POSTGRES_PASSWORD = "testpass",
    max_connections = "20"
}
fixtures = ["pool_test_schema.sql"]

[[validation.rules]]
type = "database"
name = "connection_pooling"
config = {
    service = "main_db",
    concurrent_connections = 50,
    test_duration = "30s",
    expected_max_pool_size = 10,
    connection_reuse = true
}

[[validation.rules]]
type = "process"
name = "resource_usage"
config = {
    monitor_duration = "30s",
    max_memory_mb = 100,
    max_cpu_percent = 80,
    max_goroutines = 100
}
parallel = true
*/

// ConfigParser handles parsing enhanced TOML configurations
type ConfigParser struct{}

// ParseEnhancedValidation parses the new validation configuration format
func (cp *ConfigParser) ParseEnhancedValidation(data map[string]interface{}) (*EnhancedExerciseValidation, error) {
	// Implementation will parse the enhanced TOML format and return structured config
	validation := &EnhancedExerciseValidation{}
	
	// Parse legacy fields for backward compatibility
	if mode, ok := data["mode"].(string); ok {
		validation.Mode = mode
	}
	
	if timeout, ok := data["timeout"].(string); ok {
		validation.Timeout = timeout
	}
	
	// Parse new fields
	if services, ok := data["services"].([]interface{}); ok {
		for _, service := range services {
			if serviceMap, ok := service.(map[string]interface{}); ok {
				spec, err := cp.parseServiceSpec(serviceMap)
				if err != nil {
					return nil, err
				}
				validation.Services = append(validation.Services, *spec)
			}
		}
	}
	
	if rules, ok := data["rules"].([]interface{}); ok {
		for _, rule := range rules {
			if ruleMap, ok := rule.(map[string]interface{}); ok {
				spec, err := cp.parseRuleSpec(ruleMap)
				if err != nil {
					return nil, err
				}
				validation.Rules = append(validation.Rules, *spec)
			}
		}
	}
	
	return validation, nil
}

func (cp *ConfigParser) parseServiceSpec(data map[string]interface{}) (*ServiceSpec, error) {
	spec := &ServiceSpec{}
	
	if serviceType, ok := data["type"].(string); ok {
		spec.Type = serviceType
	}
	
	if name, ok := data["name"].(string); ok {
		spec.Name = name
	}
	
	if version, ok := data["version"].(string); ok {
		spec.Version = version
	}
	
	if config, ok := data["config"].(map[string]interface{}); ok {
		spec.Config = config
	}
	
	return spec, nil
}

func (cp *ConfigParser) parseRuleSpec(data map[string]interface{}) (*ValidationRuleSpec, error) {
	spec := &ValidationRuleSpec{}
	
	if ruleType, ok := data["type"].(string); ok {
		spec.Type = ruleType
	}
	
	if name, ok := data["name"].(string); ok {
		spec.Name = name
	}
	
	if config, ok := data["config"].(map[string]interface{}); ok {
		spec.Config = config
	}
	
	if dependsOn, ok := data["depends_on"].([]interface{}); ok {
		for _, dep := range dependsOn {
			if depStr, ok := dep.(string); ok {
				spec.DependsOn = append(spec.DependsOn, depStr)
			}
		}
	}
	
	if parallel, ok := data["parallel"].(bool); ok {
		spec.Parallel = parallel
	}
	
	return spec, nil
}

// GetDefaultTimeout returns default timeout for validation mode
func GetDefaultTimeout(mode string) time.Duration {
	switch mode {
	case "universal":
		return 120 * time.Second // Universal validation can take longer
	case "build":
		return 30 * time.Second
	case "test":
		return 60 * time.Second
	case "run":
		return 30 * time.Second
	default:
		return 30 * time.Second
	}
}