package validation

import (
	"context"
	"fmt"
	"log"
	"sort"
	"sync"
)

// NewResourceManager creates a new resource manager
func NewResourceManager() *ResourceManager {
	return &ResourceManager{
		activeServices: make(map[string]Service),
		activeNetworks: make(map[string]ContainerNetwork),
		cleanupTasks:   make([]CleanupTask, 0),
	}
}

// RegisterService registers a service for resource management
func (rm *ResourceManager) RegisterService(name string, service Service) {
	rm.activeServices[name] = service
	
	// Add cleanup task for this service
	rm.AddCleanupTask(CleanupTask{
		Name:     fmt.Sprintf("stop_service_%s", name),
		Priority: 100, // Services should be stopped before networks
		Execute: func(ctx context.Context) error {
			log.Printf("Cleanup: Stopping service %s", name)
			return service.Stop(ctx)
		},
	})
}

// RegisterNetwork registers a network for resource management
func (rm *ResourceManager) RegisterNetwork(name string, network ContainerNetwork) {
	rm.activeNetworks[name] = network
	
	// Add cleanup task for this network
	rm.AddCleanupTask(CleanupTask{
		Name:     fmt.Sprintf("remove_network_%s", name),
		Priority: 200, // Networks should be removed after services
		Execute: func(ctx context.Context) error {
			log.Printf("Cleanup: Removing network %s", name)
			return network.Remove(ctx, name)
		},
	})
}

// AddCleanupTask adds a custom cleanup task
func (rm *ResourceManager) AddCleanupTask(task CleanupTask) {
	rm.cleanupTasks = append(rm.cleanupTasks, task)
}

// Cleanup performs all registered cleanup tasks in priority order
func (rm *ResourceManager) Cleanup(ctx context.Context) error {
	log.Printf("🧹 Starting resource cleanup...")
	
	// Sort cleanup tasks by priority (lower numbers = higher priority)
	sort.Slice(rm.cleanupTasks, func(i, j int) bool {
		return rm.cleanupTasks[i].Priority < rm.cleanupTasks[j].Priority
	})
	
	var errors []error
	var wg sync.WaitGroup
	errorsChan := make(chan error, len(rm.cleanupTasks))
	
	// Execute cleanup tasks in priority groups
	priorityGroups := rm.groupTasksByPriority()
	
	for priority, tasks := range priorityGroups {
		log.Printf("Cleanup: Executing priority %d tasks (%d tasks)", priority, len(tasks))
		
		// Execute all tasks in this priority group concurrently
		for _, task := range tasks {
			wg.Add(1)
			go func(t CleanupTask) {
				defer wg.Done()
				log.Printf("  Executing cleanup task: %s", t.Name)
				if err := t.Execute(ctx); err != nil {
					log.Printf("  ❌ Cleanup task %s failed: %v", t.Name, err)
					errorsChan <- fmt.Errorf("task %s failed: %w", t.Name, err)
				} else {
					log.Printf("  ✅ Cleanup task %s completed", t.Name)
				}
			}(task)
		}
		
		// Wait for all tasks in this priority group to complete before moving to next
		wg.Wait()
	}
	
	close(errorsChan)
	
	// Collect any errors
	for err := range errorsChan {
		errors = append(errors, err)
	}
	
	// Clear the cleanup tasks and resources
	rm.cleanupTasks = make([]CleanupTask, 0)
	rm.activeServices = make(map[string]Service)
	rm.activeNetworks = make(map[string]ContainerNetwork)
	
	if len(errors) > 0 {
		log.Printf("❌ Cleanup completed with %d errors", len(errors))
		return fmt.Errorf("cleanup failed with %d errors: %v", len(errors), errors)
	}
	
	log.Printf("✅ Resource cleanup completed successfully")
	return nil
}

// groupTasksByPriority groups cleanup tasks by their priority level
func (rm *ResourceManager) groupTasksByPriority() map[int][]CleanupTask {
	groups := make(map[int][]CleanupTask)
	
	for _, task := range rm.cleanupTasks {
		groups[task.Priority] = append(groups[task.Priority], task)
	}
	
	return groups
}

// GetActiveServices returns all currently active services
func (rm *ResourceManager) GetActiveServices() map[string]Service {
	return rm.activeServices
}

// GetActiveNetworks returns all currently active networks
func (rm *ResourceManager) GetActiveNetworks() map[string]ContainerNetwork {
	return rm.activeNetworks
}

// GetCleanupTaskCount returns the number of pending cleanup tasks
func (rm *ResourceManager) GetCleanupTaskCount() int {
	return len(rm.cleanupTasks)
}

// ForceCleanup performs immediate cleanup of all resources without waiting
func (rm *ResourceManager) ForceCleanup(ctx context.Context) error {
	log.Printf("🚨 Force cleanup initiated")
	
	var errors []error
	
	// Stop all services immediately
	for name, service := range rm.activeServices {
		log.Printf("Force stopping service: %s", name)
		if err := service.Stop(ctx); err != nil {
			errors = append(errors, fmt.Errorf("failed to force stop service %s: %w", name, err))
		}
	}
	
	// Remove all networks immediately
	for name, network := range rm.activeNetworks {
		log.Printf("Force removing network: %s", name)
		if err := network.Remove(ctx, name); err != nil {
			errors = append(errors, fmt.Errorf("failed to force remove network %s: %w", name, err))
		}
	}
	
	// Clear resources
	rm.activeServices = make(map[string]Service)
	rm.activeNetworks = make(map[string]ContainerNetwork)
	rm.cleanupTasks = make([]CleanupTask, 0)
	
	if len(errors) > 0 {
		return fmt.Errorf("force cleanup encountered %d errors: %v", len(errors), errors)
	}
	
	log.Printf("✅ Force cleanup completed")
	return nil
}

// GetResourceSummary returns a summary of active resources
func (rm *ResourceManager) GetResourceSummary() map[string]int {
	return map[string]int{
		"services":      len(rm.activeServices),
		"networks":      len(rm.activeNetworks),
		"cleanup_tasks": len(rm.cleanupTasks),
	}
}

// Simple implementation of ContainerNetwork interface for now
type SimpleContainerNetwork struct {
	name string
}

func (scn *SimpleContainerNetwork) Create(ctx context.Context, name string) error {
	log.Printf("Creating network: %s", name)
	scn.name = name
	// TODO: Implement actual Docker network creation
	return nil
}

func (scn *SimpleContainerNetwork) Remove(ctx context.Context, name string) error {
	log.Printf("Removing network: %s", name)
	// TODO: Implement actual Docker network removal
	return nil
}

func (scn *SimpleContainerNetwork) GetName() string {
	return scn.name
}