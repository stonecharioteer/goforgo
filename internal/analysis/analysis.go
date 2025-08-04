package analysis

import "fmt"

// StaticCheck defines the interface for all static analysis checks.
type StaticCheck interface {
	Name() string
	Description() string
	Execute(filePath string) (bool, string, error)
}

// registry holds all registered static checks.
var registry = make(map[string]StaticCheck)

// Register adds a new static check to the registry.
func Register(check StaticCheck) {
	if _, exists := registry[check.Name()]; exists {
		panic(fmt.Sprintf("static check already registered: %s", check.Name()))
	}
	registry[check.Name()] = check
}

// GetCheck retrieves a registered static check by name.
func GetCheck(name string) (StaticCheck, bool) {
	check, exists := registry[name]
	return check, exists
}
