// GoForGo Solution: Kubernetes Basic Client
// Complete implementation of Kubernetes API client operations

package main

import (
	"context"
	"fmt"
	"log"
	"path/filepath"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

// initKubernetesClient initializes the Kubernetes client
func initKubernetesClient() (kubernetes.Interface, error) {
	// Build kubeconfig path
	var kubeconfig string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = filepath.Join(home, ".kube", "config")
	}

	// Build config from kubeconfig file
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		return nil, fmt.Errorf("failed to build config: %w", err)
	}

	// Create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, fmt.Errorf("failed to create clientset: %w", err)
	}

	log.Println("Successfully initialized Kubernetes client")
	return clientset, nil
}

// listNamespaces lists all namespaces in the cluster
func listNamespaces(clientset kubernetes.Interface) error {
	fmt.Println("\n=== Namespaces ===")
	
	namespaces, err := clientset.CoreV1().Namespaces().List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return fmt.Errorf("failed to list namespaces: %w", err)
	}

	fmt.Printf("Found %d namespaces:\n", len(namespaces.Items))
	for _, ns := range namespaces.Items {
		fmt.Printf("  - Name: %s, Created: %s, Status: %s\n",
			ns.Name,
			ns.CreationTimestamp.Format("2006-01-02 15:04:05"),
			ns.Status.Phase,
		)
	}

	return nil
}

// listPods lists all pods in the specified namespace
func listPods(clientset kubernetes.Interface, namespace string) error {
	fmt.Printf("\n=== Pods in namespace '%s' ===\n", namespace)
	
	pods, err := clientset.CoreV1().Pods(namespace).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return fmt.Errorf("failed to list pods in namespace %s: %w", namespace, err)
	}

	if len(pods.Items) == 0 {
		fmt.Printf("No pods found in namespace '%s'\n", namespace)
		return nil
	}

	fmt.Printf("Found %d pods:\n", len(pods.Items))
	for _, pod := range pods.Items {
		fmt.Printf("  - Name: %s\n", pod.Name)
		fmt.Printf("    Status: %s\n", pod.Status.Phase)
		fmt.Printf("    Node: %s\n", pod.Spec.NodeName)
		fmt.Printf("    Containers: %d\n", len(pod.Spec.Containers))
		fmt.Printf("    Ready: %s\n", getPodReadyStatus(&pod))
		fmt.Println()
	}

	return nil
}

// getPodReadyStatus returns ready status string for a pod
func getPodReadyStatus(pod *metav1.Object) string {
	// This is a simplified version - in real implementation you'd check container statuses
	return "Unknown"
}

// describePod gets detailed information about a specific pod
func describePod(clientset kubernetes.Interface, namespace, podName string) error {
	fmt.Printf("\n=== Pod Details: %s/%s ===\n", namespace, podName)
	
	pod, err := clientset.CoreV1().Pods(namespace).Get(context.Background(), podName, metav1.GetOptions{})
	if err != nil {
		return fmt.Errorf("failed to get pod %s/%s: %w", namespace, podName, err)
	}

	fmt.Printf("Name: %s\n", pod.Name)
	fmt.Printf("Namespace: %s\n", pod.Namespace)
	fmt.Printf("Node: %s\n", pod.Spec.NodeName)
	fmt.Printf("Status: %s\n", pod.Status.Phase)
	fmt.Printf("Created: %s\n", pod.CreationTimestamp.Format("2006-01-02 15:04:05"))
	
	fmt.Printf("Labels:\n")
	for key, value := range pod.Labels {
		fmt.Printf("  %s: %s\n", key, value)
	}

	fmt.Printf("Containers (%d):\n", len(pod.Spec.Containers))
	for _, container := range pod.Spec.Containers {
		fmt.Printf("  - Name: %s\n", container.Name)
		fmt.Printf("    Image: %s\n", container.Image)
		fmt.Printf("    Ports: %v\n", container.Ports)
		
		if container.Resources.Requests != nil {
			fmt.Printf("    Resource Requests: %v\n", container.Resources.Requests)
		}
		if container.Resources.Limits != nil {
			fmt.Printf("    Resource Limits: %v\n", container.Resources.Limits)
		}
	}

	return nil
}

// listServices lists all services in the specified namespace
func listServices(clientset kubernetes.Interface, namespace string) error {
	fmt.Printf("\n=== Services in namespace '%s' ===\n", namespace)
	
	services, err := clientset.CoreV1().Services(namespace).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return fmt.Errorf("failed to list services in namespace %s: %w", namespace, err)
	}

	if len(services.Items) == 0 {
		fmt.Printf("No services found in namespace '%s'\n", namespace)
		return nil
	}

	fmt.Printf("Found %d services:\n", len(services.Items))
	for _, svc := range services.Items {
		fmt.Printf("  - Name: %s\n", svc.Name)
		fmt.Printf("    Type: %s\n", svc.Spec.Type)
		fmt.Printf("    Cluster IP: %s\n", svc.Spec.ClusterIP)
		
		if len(svc.Spec.Ports) > 0 {
			fmt.Printf("    Ports: ")
			for i, port := range svc.Spec.Ports {
				if i > 0 {
					fmt.Printf(", ")
				}
				fmt.Printf("%d/%s", port.Port, port.Protocol)
			}
			fmt.Println()
		}
		
		if len(svc.Status.LoadBalancer.Ingress) > 0 {
			fmt.Printf("    External IP: %s\n", svc.Status.LoadBalancer.Ingress[0].IP)
		}
		fmt.Println()
	}

	return nil
}

// getClusterInfo retrieves and displays cluster information
func getClusterInfo(clientset kubernetes.Interface) error {
	fmt.Println("\n=== Cluster Information ===")
	
	// Get server version
	version, err := clientset.Discovery().ServerVersion()
	if err != nil {
		return fmt.Errorf("failed to get server version: %w", err)
	}

	fmt.Printf("Kubernetes Version: %s\n", version.GitVersion)
	fmt.Printf("Go Version: %s\n", version.GoVersion)
	fmt.Printf("Platform: %s\n", version.Platform)

	// Get API resources (simplified)
	groups, err := clientset.Discovery().ServerGroups()
	if err != nil {
		return fmt.Errorf("failed to get server groups: %w", err)
	}

	fmt.Printf("API Groups: %d\n", len(groups.Groups))
	for _, group := range groups.Groups[:min(5, len(groups.Groups))] {
		fmt.Printf("  - %s (versions: %v)\n", group.Name, group.Versions)
	}

	if len(groups.Groups) > 5 {
		fmt.Printf("  ... and %d more groups\n", len(groups.Groups)-5)
	}

	return nil
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	// Initialize the Kubernetes client
	clientset, err := initKubernetesClient()
	if err != nil {
		log.Fatal("Failed to initialize Kubernetes client:", err)
	}

	// Get and display cluster information
	if err := getClusterInfo(clientset); err != nil {
		log.Printf("Failed to get cluster info: %v", err)
	}

	// List all namespaces
	if err := listNamespaces(clientset); err != nil {
		log.Printf("Failed to list namespaces: %v", err)
	}

	// List pods in the "default" namespace
	if err := listPods(clientset, "default"); err != nil {
		log.Printf("Failed to list pods in default namespace: %v", err)
	}

	// List pods in the "kube-system" namespace
	if err := listPods(clientset, "kube-system"); err != nil {
		log.Printf("Failed to list pods in kube-system namespace: %v", err)
	}

	// Try to describe the first pod in default namespace
	pods, err := clientset.CoreV1().Pods("default").List(context.Background(), metav1.ListOptions{Limit: 1})
	if err == nil && len(pods.Items) > 0 {
		if err := describePod(clientset, "default", pods.Items[0].Name); err != nil {
			log.Printf("Failed to describe pod: %v", err)
		}
	}

	// List services in the "default" namespace
	if err := listServices(clientset, "default"); err != nil {
		log.Printf("Failed to list services in default namespace: %v", err)
	}

	fmt.Println("Kubernetes basic client operations completed!")
}