// GoForGo Solution: Kubernetes Custom Resource Definitions (CRDs)
// Complete implementation of CRD creation and management

package main

import (
	"context"
	"fmt"
	"log"
	"path/filepath"
	"time"

	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	apiextensionsclientset "k8s.io/apiextensions-apiserver/pkg/clientset/clientset"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

// initClients initializes all required Kubernetes clients
func initClients() (kubernetes.Interface, apiextensionsclientset.Interface, dynamic.Interface, error) {
	// Build kubeconfig path
	var kubeconfig string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = filepath.Join(home, ".kube", "config")
	}

	// Build config
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("failed to build config: %w", err)
	}

	// Create clients
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("failed to create kubernetes client: %w", err)
	}

	apiExtClient, err := apiextensionsclientset.NewForConfig(config)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("failed to create apiextensions client: %w", err)
	}

	dynamicClient, err := dynamic.NewForConfig(config)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("failed to create dynamic client: %w", err)
	}

	log.Println("Successfully initialized all Kubernetes clients")
	return clientset, apiExtClient, dynamicClient, nil
}

// createWebAppCRD defines a CRD for WebApp resources
func createWebAppCRD() *apiextensionsv1.CustomResourceDefinition {
	return &apiextensionsv1.CustomResourceDefinition{
		ObjectMeta: metav1.ObjectMeta{
			Name: "webapps.apps.example.com",
		},
		Spec: apiextensionsv1.CustomResourceDefinitionSpec{
			Group: "apps.example.com",
			Versions: []apiextensionsv1.CustomResourceDefinitionVersion{
				{
					Name:    "v1",
					Served:  true,
					Storage: true,
					Schema: &apiextensionsv1.CustomResourceValidation{
						OpenAPIV3Schema: &apiextensionsv1.JSONSchemaProps{
							Type: "object",
							Properties: map[string]apiextensionsv1.JSONSchemaProps{
								"spec": {
									Type: "object",
									Properties: map[string]apiextensionsv1.JSONSchemaProps{
										"replicas": {
											Type:    "integer",
											Minimum: &[]float64{1}[0],
										},
										"image": {
											Type: "string",
										},
										"port": {
											Type:    "integer",
											Minimum: &[]float64{1}[0],
											Maximum: &[]float64{65535}[0],
										},
									},
									Required: []string{"image", "replicas", "port"},
								},
							},
						},
					},
				},
			},
			Scope: apiextensionsv1.NamespaceScoped,
			Names: apiextensionsv1.CustomResourceDefinitionNames{
				Plural:   "webapps",
				Singular: "webapp",
				Kind:     "WebApp",
			},
		},
	}
}

// installCRD installs the CRD in the cluster
func installCRD(apiExtClient apiextensionsclientset.Interface, crd *apiextensionsv1.CustomResourceDefinition) error {
	log.Printf("Installing CRD: %s", crd.Name)
	
	_, err := apiExtClient.ApiextensionsV1().CustomResourceDefinitions().Create(
		context.Background(), crd, metav1.CreateOptions{})
	if err != nil {
		return fmt.Errorf("failed to create CRD: %w", err)
	}

	log.Println("CRD installed successfully")
	return nil
}

// waitForCRDEstablished waits for the CRD to be established
func waitForCRDEstablished(apiExtClient apiextensionsclientset.Interface, crdName string) error {
	log.Printf("Waiting for CRD %s to be established...", crdName)
	
	return wait.PollImmediate(time.Second, 30*time.Second, func() (bool, error) {
		crd, err := apiExtClient.ApiextensionsV1().CustomResourceDefinitions().Get(
			context.Background(), crdName, metav1.GetOptions{})
		if err != nil {
			return false, err
		}

		for _, condition := range crd.Status.Conditions {
			if condition.Type == apiextensionsv1.Established &&
				condition.Status == apiextensionsv1.ConditionTrue {
				log.Println("CRD is now established")
				return true, nil
			}
		}
		return false, nil
	})
}

// createWebAppInstance creates a WebApp custom resource instance
func createWebAppInstance(dynamicClient dynamic.Interface, namespace, name, image string, replicas, port int32) error {
	log.Printf("Creating WebApp instance: %s", name)

	// Define the GVR (Group Version Resource)
	gvr := schema.GroupVersionResource{
		Group:    "apps.example.com",
		Version:  "v1",
		Resource: "webapps",
	}

	// Create the unstructured object
	webapp := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "apps.example.com/v1",
			"kind":       "WebApp",
			"metadata": map[string]interface{}{
				"name":      name,
				"namespace": namespace,
			},
			"spec": map[string]interface{}{
				"replicas": replicas,
				"image":    image,
				"port":     port,
			},
		},
	}

	_, err := dynamicClient.Resource(gvr).Namespace(namespace).Create(
		context.Background(), webapp, metav1.CreateOptions{})
	if err != nil {
		return fmt.Errorf("failed to create WebApp %s: %w", name, err)
	}

	log.Printf("WebApp %s created successfully", name)
	return nil
}

// listWebAppInstances lists all WebApp instances in a namespace
func listWebAppInstances(dynamicClient dynamic.Interface, namespace string) error {
	log.Printf("Listing WebApp instances in namespace: %s", namespace)

	gvr := schema.GroupVersionResource{
		Group:    "apps.example.com",
		Version:  "v1",
		Resource: "webapps",
	}

	list, err := dynamicClient.Resource(gvr).Namespace(namespace).List(
		context.Background(), metav1.ListOptions{})
	if err != nil {
		return fmt.Errorf("failed to list WebApps: %w", err)
	}

	fmt.Printf("\nFound %d WebApp instances:\n", len(list.Items))
	for _, item := range list.Items {
		name := item.GetName()
		spec, _ := item.Object["spec"].(map[string]interface{})
		
		fmt.Printf("- Name: %s\n", name)
		if spec != nil {
			fmt.Printf("  Image: %v\n", spec["image"])
			fmt.Printf("  Replicas: %v\n", spec["replicas"])
			fmt.Printf("  Port: %v\n", spec["port"])
		}
		fmt.Println()
	}

	return nil
}

// updateWebAppInstance updates a WebApp instance's replicas
func updateWebAppInstance(dynamicClient dynamic.Interface, namespace, name string, newReplicas int32) error {
	log.Printf("Updating WebApp %s replicas to %d", name, newReplicas)

	gvr := schema.GroupVersionResource{
		Group:    "apps.example.com",
		Version:  "v1",
		Resource: "webapps",
	}

	// Get the existing resource
	webapp, err := dynamicClient.Resource(gvr).Namespace(namespace).Get(
		context.Background(), name, metav1.GetOptions{})
	if err != nil {
		return fmt.Errorf("failed to get WebApp %s: %w", name, err)
	}

	// Update the replicas field
	spec, ok := webapp.Object["spec"].(map[string]interface{})
	if !ok {
		return fmt.Errorf("invalid spec in WebApp %s", name)
	}
	spec["replicas"] = newReplicas

	// Update the resource
	_, err = dynamicClient.Resource(gvr).Namespace(namespace).Update(
		context.Background(), webapp, metav1.UpdateOptions{})
	if err != nil {
		return fmt.Errorf("failed to update WebApp %s: %w", name, err)
	}

	log.Printf("WebApp %s updated successfully", name)
	return nil
}

// deleteWebAppInstance deletes a WebApp instance
func deleteWebAppInstance(dynamicClient dynamic.Interface, namespace, name string) error {
	log.Printf("Deleting WebApp instance: %s", name)

	gvr := schema.GroupVersionResource{
		Group:    "apps.example.com",
		Version:  "v1",
		Resource: "webapps",
	}

	err := dynamicClient.Resource(gvr).Namespace(namespace).Delete(
		context.Background(), name, metav1.DeleteOptions{})
	if err != nil {
		return fmt.Errorf("failed to delete WebApp %s: %w", name, err)
	}

	log.Printf("WebApp %s deleted successfully", name)
	return nil
}

// cleanupCRD deletes the CRD from the cluster
func cleanupCRD(apiExtClient apiextensionsclientset.Interface, crdName string) error {
	log.Printf("Cleaning up CRD: %s", crdName)

	err := apiExtClient.ApiextensionsV1().CustomResourceDefinitions().Delete(
		context.Background(), crdName, metav1.DeleteOptions{})
	if err != nil {
		return fmt.Errorf("failed to delete CRD %s: %w", crdName, err)
	}

	log.Println("CRD cleaned up successfully")
	return nil
}

func main() {
	// Initialize all required clients
	_, apiExtClient, dynamicClient, err := initClients()
	if err != nil {
		log.Fatal("Failed to initialize clients:", err)
	}

	// Create the WebApp CRD definition
	crd := createWebAppCRD()

	// Install the CRD in the cluster
	if err := installCRD(apiExtClient, crd); err != nil {
		log.Fatal("Failed to install CRD:", err)
	}

	// Wait for the CRD to be established
	if err := waitForCRDEstablished(apiExtClient, crd.Name); err != nil {
		log.Fatal("Failed to wait for CRD establishment:", err)
	}

	namespace := "default"

	// Create some WebApp instances
	if err := createWebAppInstance(dynamicClient, namespace, "frontend-app", "nginx:latest", 3, 80); err != nil {
		log.Printf("Failed to create frontend-app: %v", err)
	}

	if err := createWebAppInstance(dynamicClient, namespace, "backend-app", "node:alpine", 2, 3000); err != nil {
		log.Printf("Failed to create backend-app: %v", err)
	}

	// List all WebApp instances
	if err := listWebAppInstances(dynamicClient, namespace); err != nil {
		log.Printf("Failed to list WebApps: %v", err)
	}

	// Update the frontend-app to have 5 replicas
	if err := updateWebAppInstance(dynamicClient, namespace, "frontend-app", 5); err != nil {
		log.Printf("Failed to update frontend-app: %v", err)
	}

	// List instances again to show the update
	fmt.Println("=== After Update ===")
	if err := listWebAppInstances(dynamicClient, namespace); err != nil {
		log.Printf("Failed to list WebApps after update: %v", err)
	}

	// Delete the backend-app instance
	if err := deleteWebAppInstance(dynamicClient, namespace, "backend-app"); err != nil {
		log.Printf("Failed to delete backend-app: %v", err)
	}

	// Final list to show remaining instances
	fmt.Println("=== After Deletion ===")
	if err := listWebAppInstances(dynamicClient, namespace); err != nil {
		log.Printf("Failed to list WebApps after deletion: %v", err)
	}

	// Cleanup: delete the CRD (this will also delete all instances)
	if err := cleanupCRD(apiExtClient, crd.Name); err != nil {
		log.Printf("Failed to cleanup CRD: %v", err)
	}

	fmt.Println("Kubernetes CRD operations completed!")
}