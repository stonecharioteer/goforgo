// GoForGo Solution: Kubernetes Controllers
// Complete implementation of a Pod controller that watches and reacts to Pod changes

package main

import (
	"context"
	"fmt"
	"log"
	"path/filepath"
	"time"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"k8s.io/client-go/util/workqueue"
)

// PodController watches Pod resources and reacts to changes
type PodController struct {
	clientset kubernetes.Interface
	indexer   cache.Indexer
	queue     workqueue.RateLimitingInterface
	informer  cache.Controller
}

// initKubernetesClient initializes the Kubernetes client
func initKubernetesClient() (kubernetes.Interface, error) {
	var kubeconfig string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = filepath.Join(home, ".kube", "config")
	}

	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		return nil, fmt.Errorf("failed to build config: %w", err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, fmt.Errorf("failed to create clientset: %w", err)
	}

	log.Println("Successfully initialized Kubernetes client for controller")
	return clientset, nil
}

// NewPodController creates a new Pod controller
func NewPodController(clientset kubernetes.Interface) *PodController {
	// Create a ListWatcher for Pods in all namespaces
	podListWatcher := &cache.ListWatch{
		ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
			return clientset.CoreV1().Pods(metav1.NamespaceAll).List(context.Background(), options)
		},
		WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
			return clientset.CoreV1().Pods(metav1.NamespaceAll).Watch(context.Background(), options)
		},
	}

	// Create a rate limiting work queue
	queue := workqueue.NewRateLimitingQueue(workqueue.DefaultControllerRateLimiter())

	// Create the controller
	controller := &PodController{
		clientset: clientset,
		queue:     queue,
	}

	// Create an informer
	indexer, informer := cache.NewIndexerInformer(
		podListWatcher,
		&corev1.Pod{},
		time.Second*30, // resync period
		cache.ResourceEventHandlerFuncs{
			AddFunc:    controller.handleAdd,
			UpdateFunc: controller.handleUpdate,
			DeleteFunc: controller.handleDelete,
		},
		cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc},
	)

	controller.indexer = indexer
	controller.informer = informer

	log.Println("Created Pod controller with informer and work queue")
	return controller
}

// handleAdd is called when a pod is created
func (c *PodController) handleAdd(obj interface{}) {
	pod := obj.(*corev1.Pod)
	log.Printf("Pod ADDED: %s/%s", pod.Namespace, pod.Name)
	
	key, err := keyFunc(obj)
	if err != nil {
		log.Printf("Error getting key for added pod: %v", err)
		return
	}
	c.queue.Add(key)
}

// handleUpdate is called when a pod is updated
func (c *PodController) handleUpdate(oldObj, newObj interface{}) {
	oldPod := oldObj.(*corev1.Pod)
	newPod := newObj.(*corev1.Pod)
	
	// Only process significant updates (e.g., status changes)
	if oldPod.Status.Phase != newPod.Status.Phase {
		log.Printf("Pod UPDATED: %s/%s - Status: %s -> %s", 
			newPod.Namespace, newPod.Name, oldPod.Status.Phase, newPod.Status.Phase)
		
		key, err := keyFunc(newObj)
		if err != nil {
			log.Printf("Error getting key for updated pod: %v", err)
			return
		}
		c.queue.Add(key)
	}
}

// handleDelete is called when a pod is deleted
func (c *PodController) handleDelete(obj interface{}) {
	pod, ok := obj.(*corev1.Pod)
	if !ok {
		// Handle DeletedFinalStateUnknown
		tombstone, ok := obj.(cache.DeletedFinalStateUnknown)
		if !ok {
			log.Printf("Error decoding object, invalid type")
			return
		}
		pod, ok = tombstone.Obj.(*corev1.Pod)
		if !ok {
			log.Printf("Error decoding object tombstone, invalid type")
			return
		}
	}
	
	log.Printf("Pod DELETED: %s/%s", pod.Namespace, pod.Name)
	
	key, err := keyFunc(obj)
	if err != nil {
		log.Printf("Error getting key for deleted pod: %v", err)
		return
	}
	c.queue.Add(key)
}

// keyFunc extracts a unique key from a Kubernetes object
func keyFunc(obj interface{}) (string, error) {
	return cache.MetaNamespaceKeyFunc(obj)
}

// processItem processes a single item from the work queue
func (c *PodController) processItem() bool {
	key, quit := c.queue.Get()
	if quit {
		return false
	}
	defer c.queue.Done(key)

	err := c.syncPod(key.(string))
	if err == nil {
		// No error, tell the queue to stop tracking history for this key
		c.queue.Forget(key)
		return true
	}

	// Handle error
	log.Printf("Error processing pod %s: %v", key, err)
	
	// Re-queue the item with rate limiting
	c.queue.AddRateLimited(key)
	return true
}

// syncPod handles the actual pod processing logic
func (c *PodController) syncPod(key string) error {
	namespace, name, err := cache.SplitMetaNamespaceKey(key)
	if err != nil {
		return fmt.Errorf("invalid resource key: %s", key)
	}

	// Get the pod from our local cache
	obj, exists, err := c.indexer.GetByKey(key)
	if err != nil {
		return fmt.Errorf("error fetching pod %s from store: %v", key, err)
	}

	if !exists {
		log.Printf("Pod %s no longer exists, handling deletion", key)
		// Pod was deleted, perform cleanup if needed
		return nil
	}

	pod := obj.(*corev1.Pod)
	
	// Implement custom controller logic here
	// For this example, we'll add a label if it doesn't exist
	if pod.Labels == nil {
		pod.Labels = make(map[string]string)
	}

	const managedLabel = "controller.example.com/managed"
	if _, exists := pod.Labels[managedLabel]; !exists {
		log.Printf("Adding managed label to pod %s/%s", namespace, name)
		
		// Create a copy of the pod to update
		updatedPod := pod.DeepCopy()
		updatedPod.Labels[managedLabel] = "true"
		updatedPod.Labels["controller.example.com/last-seen"] = time.Now().Format(time.RFC3339)
		
		// Update the pod
		_, err := c.clientset.CoreV1().Pods(namespace).Update(
			context.Background(), updatedPod, metav1.UpdateOptions{})
		if err != nil {
			return fmt.Errorf("failed to update pod %s/%s: %v", namespace, name, err)
		}
		
		log.Printf("Successfully updated pod %s/%s with managed labels", namespace, name)
	}

	// Log pod status for demonstration
	log.Printf("Processing pod %s/%s - Status: %s, Node: %s", 
		namespace, name, pod.Status.Phase, pod.Spec.NodeName)

	return nil
}

// Run starts the controller
func (c *PodController) Run(stopCh <-chan struct{}) {
	defer utilruntime.HandleCrash()
	defer c.queue.ShutDown()

	log.Println("Starting Pod controller")

	// Start the informer
	go c.informer.Run(stopCh)

	// Wait for cache to sync
	log.Println("Waiting for cache to sync...")
	if !cache.WaitForCacheSync(stopCh, c.informer.HasSynced) {
		log.Fatal("Timed out waiting for cache to sync")
	}
	log.Println("Cache synced successfully")

	// Start worker goroutines
	log.Println("Starting worker...")
	go wait.Until(c.runWorker, time.Second, stopCh)

	log.Println("Pod controller started successfully")
	<-stopCh
	log.Println("Shutting down Pod controller")
}

// runWorker is a worker goroutine that processes items from the queue
func (c *PodController) runWorker() {
	for c.processItem() {
	}
}

func main() {
	// Initialize Kubernetes client
	clientset, err := initKubernetesClient()
	if err != nil {
		log.Fatal("Failed to initialize Kubernetes client:", err)
	}

	// Create the PodController
	controller := NewPodController(clientset)

	// Create a stop channel for graceful shutdown
	stopCh := make(chan struct{})

	// Start the controller in a goroutine
	go controller.Run(stopCh)

	// Run for 2 minutes to observe pod changes
	log.Println("Controller running... watching for pod changes for 2 minutes")
	
	// You can create/delete pods in another terminal to see the controller in action
	log.Println("Try creating/deleting pods in another terminal: kubectl run test-pod --image=nginx")
	
	time.Sleep(2 * time.Minute)

	// Signal the controller to stop and wait for it to shut down
	log.Println("Stopping controller...")
	close(stopCh)
	
	// Give it a moment to shut down gracefully
	time.Sleep(5 * time.Second)

	fmt.Println("Kubernetes controller operations completed!")
}