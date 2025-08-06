// GoForGo Solution: Kubernetes Operators
// Complete implementation of a WebApp operator that manages custom resources

package main

import (
	"context"
	"fmt"
	"log"
	"path/filepath"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

// WebAppOperator manages the lifecycle of WebApp custom resources
type WebAppOperator struct {
	kubeClient    kubernetes.Interface
	dynamicClient dynamic.Interface
	namespace     string
	webAppGVR     schema.GroupVersionResource
}

// NewWebAppOperator creates a new WebApp operator
func NewWebAppOperator(kubeClient kubernetes.Interface, dynamicClient dynamic.Interface, namespace string) *WebAppOperator {
	return &WebAppOperator{
		kubeClient:    kubeClient,
		dynamicClient: dynamicClient,
		namespace:     namespace,
		webAppGVR: schema.GroupVersionResource{
			Group:    "apps.example.com",
			Version:  "v1",
			Resource: "webapps",
		},
	}
}

// reconcileWebApp ensures the desired state matches the actual state
func (o *WebAppOperator) reconcileWebApp(ctx context.Context, webAppName string) error {
	log.Printf("Reconciling WebApp: %s", webAppName)

	// Get the WebApp custom resource
	webApp, err := o.getWebApp(ctx, webAppName)
	if err != nil {
		if errors.IsNotFound(err) {
			log.Printf("WebApp %s not found, skipping reconciliation", webAppName)
			return nil
		}
		return fmt.Errorf("failed to get WebApp %s: %w", webAppName, err)
	}

	// Update status to "Reconciling"
	if err := o.updateWebAppStatus(ctx, webApp, "Reconciling", "Operator is processing the WebApp"); err != nil {
		log.Printf("Failed to update status: %v", err)
	}

	// Ensure Deployment exists and is up to date
	if err := o.ensureDeployment(ctx, webApp); err != nil {
		o.updateWebAppStatus(ctx, webApp, "Failed", fmt.Sprintf("Failed to ensure Deployment: %v", err))
		return fmt.Errorf("failed to ensure Deployment: %w", err)
	}

	// Ensure Service exists and is up to date
	if err := o.ensureService(ctx, webApp); err != nil {
		o.updateWebAppStatus(ctx, webApp, "Failed", fmt.Sprintf("Failed to ensure Service: %v", err))
		return fmt.Errorf("failed to ensure Service: %w", err)
	}

	// Update status to "Ready"
	if err := o.updateWebAppStatus(ctx, webApp, "Ready", "WebApp is running successfully"); err != nil {
		log.Printf("Failed to update final status: %v", err)
	}

	log.Printf("Successfully reconciled WebApp: %s", webAppName)
	return nil
}

// getWebApp retrieves a WebApp custom resource by name
func (o *WebAppOperator) getWebApp(ctx context.Context, name string) (*unstructured.Unstructured, error) {
	return o.dynamicClient.Resource(o.webAppGVR).Namespace(o.namespace).Get(ctx, name, metav1.GetOptions{})
}

// ensureDeployment creates or updates a Deployment based on WebApp spec
func (o *WebAppOperator) ensureDeployment(ctx context.Context, webApp *unstructured.Unstructured) error {
	deploymentName := webApp.GetName()
	
	// Create the desired Deployment spec
	deployment, err := o.createDeploymentSpec(webApp)
	if err != nil {
		return fmt.Errorf("failed to create Deployment spec: %w", err)
	}

	// Check if Deployment already exists
	existingDeployment, err := o.kubeClient.AppsV1().Deployments(o.namespace).Get(ctx, deploymentName, metav1.GetOptions{})
	if err != nil {
		if errors.IsNotFound(err) {
			// Create new Deployment
			_, err := o.kubeClient.AppsV1().Deployments(o.namespace).Create(ctx, deployment, metav1.CreateOptions{})
			if err != nil {
				return fmt.Errorf("failed to create Deployment: %w", err)
			}
			log.Printf("Created Deployment: %s", deploymentName)
			return nil
		}
		return fmt.Errorf("failed to get existing Deployment: %w", err)
	}

	// Update existing Deployment
	deployment.ResourceVersion = existingDeployment.ResourceVersion
	_, err = o.kubeClient.AppsV1().Deployments(o.namespace).Update(ctx, deployment, metav1.UpdateOptions{})
	if err != nil {
		return fmt.Errorf("failed to update Deployment: %w", err)
	}
	log.Printf("Updated Deployment: %s", deploymentName)
	return nil
}

// ensureService creates or updates a Service based on WebApp spec
func (o *WebAppOperator) ensureService(ctx context.Context, webApp *unstructured.Unstructured) error {
	serviceName := webApp.GetName()
	
	// Create the desired Service spec
	service, err := o.createServiceSpec(webApp)
	if err != nil {
		return fmt.Errorf("failed to create Service spec: %w", err)
	}

	// Check if Service already exists
	existingService, err := o.kubeClient.CoreV1().Services(o.namespace).Get(ctx, serviceName, metav1.GetOptions{})
	if err != nil {
		if errors.IsNotFound(err) {
			// Create new Service
			_, err := o.kubeClient.CoreV1().Services(o.namespace).Create(ctx, service, metav1.CreateOptions{})
			if err != nil {
				return fmt.Errorf("failed to create Service: %w", err)
			}
			log.Printf("Created Service: %s", serviceName)
			return nil
		}
		return fmt.Errorf("failed to get existing Service: %w", err)
	}

	// Update existing Service
	service.ResourceVersion = existingService.ResourceVersion
	service.Spec.ClusterIP = existingService.Spec.ClusterIP // Preserve ClusterIP
	_, err = o.kubeClient.CoreV1().Services(o.namespace).Update(ctx, service, metav1.UpdateOptions{})
	if err != nil {
		return fmt.Errorf("failed to update Service: %w", err)
	}
	log.Printf("Updated Service: %s", serviceName)
	return nil
}

// createDeploymentSpec builds a Deployment spec from WebApp
func (o *WebAppOperator) createDeploymentSpec(webApp *unstructured.Unstructured) (*appsv1.Deployment, error) {
	name := webApp.GetName()
	
	// Extract spec values
	image, err := getStringFromSpec(webApp, "spec", "image")
	if err != nil {
		return nil, fmt.Errorf("failed to get image: %w", err)
	}
	
	replicas, err := getIntFromSpec(webApp, "spec", "replicas")
	if err != nil {
		return nil, fmt.Errorf("failed to get replicas: %w", err)
	}
	
	port, err := getIntFromSpec(webApp, "spec", "port")
	if err != nil {
		return nil, fmt.Errorf("failed to get port: %w", err)
	}

	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: o.namespace,
			Labels: map[string]string{
				"app":                          name,
				"operator.example.com/managed": "true",
			},
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: &replicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": name,
				},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": name,
					},
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  "app",
							Image: image,
							Ports: []corev1.ContainerPort{
								{
									ContainerPort: port,
									Protocol:      corev1.ProtocolTCP,
								},
							},
							Resources: corev1.ResourceRequirements{
								Requests: corev1.ResourceList{
									corev1.ResourceCPU:    *resource.NewMilliQuantity(100, resource.DecimalSI),
									corev1.ResourceMemory: *resource.NewQuantity(128*1024*1024, resource.BinarySI),
								},
							},
						},
					},
				},
			},
		},
	}

	return deployment, nil
}

// createServiceSpec builds a Service spec from WebApp
func (o *WebAppOperator) createServiceSpec(webApp *unstructured.Unstructured) (*corev1.Service, error) {
	name := webApp.GetName()
	
	port, err := getIntFromSpec(webApp, "spec", "port")
	if err != nil {
		return nil, fmt.Errorf("failed to get port: %w", err)
	}

	service := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: o.namespace,
			Labels: map[string]string{
				"app":                          name,
				"operator.example.com/managed": "true",
			},
		},
		Spec: corev1.ServiceSpec{
			Selector: map[string]string{
				"app": name,
			},
			Ports: []corev1.ServicePort{
				{
					Name:       "http",
					Port:       port,
					TargetPort: intstr.FromInt(int(port)),
					Protocol:   corev1.ProtocolTCP,
				},
			},
			Type: corev1.ServiceTypeClusterIP,
		},
	}

	return service, nil
}

// updateWebAppStatus updates the status field of a WebApp custom resource
func (o *WebAppOperator) updateWebAppStatus(ctx context.Context, webApp *unstructured.Unstructured, phase, message string) error {
	// Set status fields
	status := map[string]interface{}{
		"phase":   phase,
		"message": message,
	}
	
	if err := unstructured.SetNestedMap(webApp.Object, status, "status"); err != nil {
		return fmt.Errorf("failed to set status: %w", err)
	}

	_, err := o.dynamicClient.Resource(o.webAppGVR).Namespace(o.namespace).UpdateStatus(ctx, webApp, metav1.UpdateOptions{})
	if err != nil {
		return fmt.Errorf("failed to update status: %w", err)
	}

	return nil
}

// getStringFromSpec safely extracts string values from WebApp spec
func getStringFromSpec(webApp *unstructured.Unstructured, fieldPath ...string) (string, error) {
	value, found, err := unstructured.NestedString(webApp.Object, fieldPath...)
	if err != nil {
		return "", err
	}
	if !found {
		return "", fmt.Errorf("field %v not found", fieldPath)
	}
	return value, nil
}

// getIntFromSpec safely extracts integer values from WebApp spec
func getIntFromSpec(webApp *unstructured.Unstructured, fieldPath ...string) (int32, error) {
	value, found, err := unstructured.NestedInt64(webApp.Object, fieldPath...)
	if err != nil {
		return 0, err
	}
	if !found {
		return 0, fmt.Errorf("field %v not found", fieldPath)
	}
	return int32(value), nil
}

func main() {
	// Initialize clients
	var kubeconfig string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = filepath.Join(home, ".kube", "config")
	}

	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		log.Fatal("Failed to build config:", err)
	}

	kubeClient, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal("Failed to create Kubernetes client:", err)
	}

	dynamicClient, err := dynamic.NewForConfig(config)
	if err != nil {
		log.Fatal("Failed to create dynamic client:", err)
	}

	// Create the WebApp operator
	operator := NewWebAppOperator(kubeClient, dynamicClient, "default")

	ctx := context.Background()

	// Create a sample WebApp custom resource for testing
	webAppGVR := schema.GroupVersionResource{
		Group:    "apps.example.com",
		Version:  "v1",
		Resource: "webapps",
	}

	sampleWebApp := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "apps.example.com/v1",
			"kind":       "WebApp",
			"metadata": map[string]interface{}{
				"name":      "demo-webapp",
				"namespace": "default",
			},
			"spec": map[string]interface{}{
				"replicas": 2,
				"image":    "nginx:latest",
				"port":     80,
			},
		},
	}

	// Create the WebApp resource
	_, err = dynamicClient.Resource(webAppGVR).Namespace("default").Create(ctx, sampleWebApp, metav1.CreateOptions{})
	if err != nil && !errors.IsAlreadyExists(err) {
		log.Printf("Failed to create sample WebApp (may already exist): %v", err)
	} else if err == nil {
		log.Println("Created sample WebApp: demo-webapp")
	}

	// Reconcile the WebApp
	if err := operator.reconcileWebApp(ctx, "demo-webapp"); err != nil {
		log.Printf("Failed to reconcile WebApp: %v", err)
	}

	// Check if resources were created successfully
	deployments, err := kubeClient.AppsV1().Deployments("default").List(ctx, metav1.ListOptions{
		LabelSelector: "operator.example.com/managed=true",
	})
	if err != nil {
		log.Printf("Failed to list deployments: %v", err)
	} else {
		log.Printf("Found %d managed deployments", len(deployments.Items))
		for _, deployment := range deployments.Items {
			log.Printf("  Deployment: %s (replicas: %d)", deployment.Name, *deployment.Spec.Replicas)
		}
	}

	services, err := kubeClient.CoreV1().Services("default").List(ctx, metav1.ListOptions{
		LabelSelector: "operator.example.com/managed=true",
	})
	if err != nil {
		log.Printf("Failed to list services: %v", err)
	} else {
		log.Printf("Found %d managed services", len(services.Items))
		for _, service := range services.Items {
			log.Printf("  Service: %s (ports: %v)", service.Name, service.Spec.Ports)
		}
	}

	// Update the WebApp (change replicas to 3)
	updatedWebApp, err := dynamicClient.Resource(webAppGVR).Namespace("default").Get(ctx, "demo-webapp", metav1.GetOptions{})
	if err != nil {
		log.Printf("Failed to get WebApp for update: %v", err)
	} else {
		if err := unstructured.SetNestedField(updatedWebApp.Object, int64(3), "spec", "replicas"); err != nil {
			log.Printf("Failed to set replicas field: %v", err)
		} else {
			_, err = dynamicClient.Resource(webAppGVR).Namespace("default").Update(ctx, updatedWebApp, metav1.UpdateOptions{})
			if err != nil {
				log.Printf("Failed to update WebApp: %v", err)
			} else {
				log.Println("Updated WebApp to 3 replicas")

				// Reconcile again
				if err := operator.reconcileWebApp(ctx, "demo-webapp"); err != nil {
					log.Printf("Failed to reconcile updated WebApp: %v", err)
				}
			}
		}
	}

	// Check final state
	deployment, err := kubeClient.AppsV1().Deployments("default").Get(ctx, "demo-webapp", metav1.GetOptions{})
	if err != nil {
		log.Printf("Failed to get final deployment state: %v", err)
	} else {
		log.Printf("Final deployment state: %s has %d replicas", deployment.Name, *deployment.Spec.Replicas)
	}

	fmt.Println("Kubernetes operator operations completed!")
}