// GoForGo Exercise: Kubernetes Basic Client
// Learn how to interact with Kubernetes API using the Go client library

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

// TODO: Create a function to initialize Kubernetes client
// Parameters: none (use default kubeconfig)
// Returns: kubernetes.Interface, error
// Use clientcmd.BuildConfigFromFlags to load config from ~/.kube/config
func initKubernetesClient() (kubernetes.Interface, error) {
	// Your initKubernetesClient implementation here
	return nil, nil
}

// TODO: Create a function to list all namespaces
// Parameters: clientset kubernetes.Interface
// Returns: error
// Use clientset.CoreV1().Namespaces().List() to get namespaces
// Print namespace names and creation timestamps
func listNamespaces(clientset kubernetes.Interface) error {
	// Your listNamespaces implementation here
	return nil
}

// TODO: Create a function to list pods in a specific namespace
// Parameters: clientset kubernetes.Interface, namespace string
// Returns: error
// Use clientset.CoreV1().Pods(namespace).List() to get pods
// Print pod names, status, and node assignments
func listPods(clientset kubernetes.Interface, namespace string) error {
	// Your listPods implementation here
	return nil
}

// TODO: Create a function to get detailed information about a specific pod
// Parameters: clientset kubernetes.Interface, namespace, podName string
// Returns: error
// Use clientset.CoreV1().Pods(namespace).Get() to get pod details
// Print container information, resource requests/limits, and labels
func describePod(clientset kubernetes.Interface, namespace, podName string) error {
	// Your describePod implementation here
	return nil
}

// TODO: Create a function to list all services in a namespace
// Parameters: clientset kubernetes.Interface, namespace string
// Returns: error
// Use clientset.CoreV1().Services(namespace).List() to get services
// Print service names, types, cluster IPs, and ports
func listServices(clientset kubernetes.Interface, namespace string) error {
	// Your listServices implementation here
	return nil
}

// TODO: Create a function to check cluster information
// Parameters: clientset kubernetes.Interface
// Returns: error
// Use clientset.Discovery().ServerVersion() to get cluster version
// List all available API resources using Discovery client
func getClusterInfo(clientset kubernetes.Interface) error {
	// Your getClusterInfo implementation here
	return nil
}

func main() {
	// TODO: Initialize the Kubernetes client

	// TODO: Get and display cluster information

	// TODO: List all namespaces

	// TODO: List pods in the "default" namespace

	// TODO: List pods in the "kube-system" namespace

	// TODO: If there are pods in default namespace, describe the first one

	// TODO: List services in the "default" namespace

	fmt.Println("Kubernetes basic client operations completed!")
}