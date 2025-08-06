// GoForGo Solution: Kubernetes Deployment Automation
// Complete implementation of deployment automation with rolling updates and resource management

package main

import (
	"context"
	"fmt"
	"log"
	"path/filepath"
	"time"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

// DeploymentAutomator handles automated deployments and operations
type DeploymentAutomator struct {
	clientset kubernetes.Interface
	namespace string
}

// ApplicationConfig defines the configuration for an application deployment
type ApplicationConfig struct {
	Name         string
	Image        string
	Replicas     int32
	Port         int32
	Env          map[string]string
	ConfigData   map[string]string
	HealthCheck  bool
	ResourceReqs ResourceRequirements
}

// ResourceRequirements defines CPU and memory requirements
type ResourceRequirements struct {
	CPURequest    string
	MemoryRequest string
	CPULimit      string
	MemoryLimit   string
}

// DeploymentStatus represents the current status of a deployment
type DeploymentStatus struct {
	Name              string
	Replicas          int32
	ReadyReplicas     int32
	UpdatedReplicas   int32
	AvailableReplicas int32
	Conditions        []appsv1.DeploymentCondition
}

// NewDeploymentAutomator creates a new deployment automator
func NewDeploymentAutomator(clientset kubernetes.Interface, namespace string) *DeploymentAutomator {
	return &DeploymentAutomator{
		clientset: clientset,
		namespace: namespace,
	}
}

// deployApplication creates a complete application stack
func (da *DeploymentAutomator) deployApplication(ctx context.Context, appConfig ApplicationConfig) error {
	log.Printf("Deploying application: %s", appConfig.Name)

	// Create ConfigMap if config data provided
	if len(appConfig.ConfigData) > 0 {
		if err := da.createConfigMap(ctx, appConfig.Name+"-config", appConfig.ConfigData); err != nil {
			return fmt.Errorf("failed to create ConfigMap: %w", err)
		}
	}

	// Create Deployment
	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      appConfig.Name,
			Namespace: da.namespace,
			Labels: map[string]string{
				"app": appConfig.Name,
			},
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: &appConfig.Replicas,
			Strategy: appsv1.DeploymentStrategy{
				Type: appsv1.RollingUpdateDeploymentStrategyType,
				RollingUpdate: &appsv1.RollingUpdateDeployment{
					MaxUnavailable: &intstr.IntOrString{Type: intstr.String, StrVal: "25%"},
					MaxSurge:       &intstr.IntOrString{Type: intstr.String, StrVal: "25%"},
				},
			},
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": appConfig.Name,
				},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": appConfig.Name,
					},
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  appConfig.Name,
							Image: appConfig.Image,
							Ports: []corev1.ContainerPort{
								{
									ContainerPort: appConfig.Port,
									Protocol:      corev1.ProtocolTCP,
								},
							},
							Env: da.buildEnvVars(appConfig.Env),
							Resources: corev1.ResourceRequirements{
								Requests: corev1.ResourceList{
									corev1.ResourceCPU:    resource.MustParse(appConfig.ResourceReqs.CPURequest),
									corev1.ResourceMemory: resource.MustParse(appConfig.ResourceReqs.MemoryRequest),
								},
								Limits: corev1.ResourceList{
									corev1.ResourceCPU:    resource.MustParse(appConfig.ResourceReqs.CPULimit),
									corev1.ResourceMemory: resource.MustParse(appConfig.ResourceReqs.MemoryLimit),
								},
							},
						},
					},
				},
			},
		},
	}

	// Add health check if enabled
	if appConfig.HealthCheck {
		deployment.Spec.Template.Spec.Containers[0].ReadinessProbe = &corev1.Probe{
			ProbeHandler: corev1.ProbeHandler{
				HTTPGet: &corev1.HTTPGetAction{
					Path: "/",
					Port: intstr.FromInt(int(appConfig.Port)),
				},
			},
			InitialDelaySeconds: 10,
			PeriodSeconds:       5,
		}
		deployment.Spec.Template.Spec.Containers[0].LivenessProbe = &corev1.Probe{
			ProbeHandler: corev1.ProbeHandler{
				HTTPGet: &corev1.HTTPGetAction{
					Path: "/",
					Port: intstr.FromInt(int(appConfig.Port)),
				},
			},
			InitialDelaySeconds: 30,
			PeriodSeconds:       10,
		}
	}

	// Add ConfigMap volume if config data exists
	if len(appConfig.ConfigData) > 0 {
		deployment.Spec.Template.Spec.Volumes = []corev1.Volume{
			{
				Name: "config-volume",
				VolumeSource: corev1.VolumeSource{
					ConfigMap: &corev1.ConfigMapVolumeSource{
						LocalObjectReference: corev1.LocalObjectReference{
							Name: appConfig.Name + "-config",
						},
					},
				},
			},
		}
		deployment.Spec.Template.Spec.Containers[0].VolumeMounts = []corev1.VolumeMount{
			{
				Name:      "config-volume",
				MountPath: "/etc/config",
			},
		}
	}

	// Create the Deployment
	_, err := da.clientset.AppsV1().Deployments(da.namespace).Create(ctx, deployment, metav1.CreateOptions{})
	if err != nil {
		return fmt.Errorf("failed to create Deployment: %w", err)
	}

	// Create Service
	if err := da.createService(ctx, appConfig.Name, appConfig.Name, appConfig.Port, corev1.ServiceTypeClusterIP); err != nil {
		return fmt.Errorf("failed to create Service: %w", err)
	}

	log.Printf("Successfully deployed application: %s", appConfig.Name)
	return nil
}

// buildEnvVars converts environment variable map to Kubernetes EnvVar slice
func (da *DeploymentAutomator) buildEnvVars(envMap map[string]string) []corev1.EnvVar {
	var envVars []corev1.EnvVar
	for key, value := range envMap {
		envVars = append(envVars, corev1.EnvVar{
			Name:  key,
			Value: value,
		})
	}
	return envVars
}

// rollingUpdate performs a rolling update with specified parameters
func (da *DeploymentAutomator) rollingUpdate(ctx context.Context, deploymentName, newImage string, maxUnavailable, maxSurge int32) error {
	log.Printf("Starting rolling update for %s to image %s", deploymentName, newImage)

	// Get current deployment
	deployment, err := da.clientset.AppsV1().Deployments(da.namespace).Get(ctx, deploymentName, metav1.GetOptions{})
	if err != nil {
		return fmt.Errorf("failed to get deployment: %w", err)
	}

	// Update image
	deployment.Spec.Template.Spec.Containers[0].Image = newImage

	// Update rolling update strategy
	deployment.Spec.Strategy.RollingUpdate.MaxUnavailable = &intstr.IntOrString{Type: intstr.Int, IntVal: maxUnavailable}
	deployment.Spec.Strategy.RollingUpdate.MaxSurge = &intstr.IntOrString{Type: intstr.Int, IntVal: maxSurge}

	// Apply update
	_, err = da.clientset.AppsV1().Deployments(da.namespace).Update(ctx, deployment, metav1.UpdateOptions{})
	if err != nil {
		return fmt.Errorf("failed to update deployment: %w", err)
	}

	log.Printf("Rolling update initiated for %s", deploymentName)
	return nil
}

// waitForDeploymentReady waits until deployment is fully rolled out
func (da *DeploymentAutomator) waitForDeploymentReady(ctx context.Context, deploymentName string, timeoutSeconds int64) error {
	log.Printf("Waiting for deployment %s to be ready (timeout: %ds)", deploymentName, timeoutSeconds)

	return wait.PollImmediate(time.Second*5, time.Duration(timeoutSeconds)*time.Second, func() (bool, error) {
		deployment, err := da.clientset.AppsV1().Deployments(da.namespace).Get(ctx, deploymentName, metav1.GetOptions{})
		if err != nil {
			return false, err
		}

		// Check if deployment is ready
		if deployment.Status.ReadyReplicas == *deployment.Spec.Replicas &&
			deployment.Status.UpdatedReplicas == *deployment.Spec.Replicas &&
			deployment.Status.AvailableReplicas == *deployment.Spec.Replicas {
			log.Printf("Deployment %s is ready", deploymentName)
			return true, nil
		}

		log.Printf("Deployment %s status: %d/%d ready, %d updated, %d available",
			deploymentName,
			deployment.Status.ReadyReplicas,
			*deployment.Spec.Replicas,
			deployment.Status.UpdatedReplicas,
			deployment.Status.AvailableReplicas,
		)
		return false, nil
	})
}

// scaleDeployment scales deployment to specified replica count
func (da *DeploymentAutomator) scaleDeployment(ctx context.Context, deploymentName string, replicas int32) error {
	log.Printf("Scaling deployment %s to %d replicas", deploymentName, replicas)

	deployment, err := da.clientset.AppsV1().Deployments(da.namespace).Get(ctx, deploymentName, metav1.GetOptions{})
	if err != nil {
		return fmt.Errorf("failed to get deployment: %w", err)
	}

	deployment.Spec.Replicas = &replicas

	_, err = da.clientset.AppsV1().Deployments(da.namespace).Update(ctx, deployment, metav1.UpdateOptions{})
	if err != nil {
		return fmt.Errorf("failed to scale deployment: %w", err)
	}

	log.Printf("Successfully scaled deployment %s to %d replicas", deploymentName, replicas)
	return nil
}

// getDeploymentStatus returns current status of a deployment
func (da *DeploymentAutomator) getDeploymentStatus(ctx context.Context, deploymentName string) (*DeploymentStatus, error) {
	deployment, err := da.clientset.AppsV1().Deployments(da.namespace).Get(ctx, deploymentName, metav1.GetOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to get deployment: %w", err)
	}

	status := &DeploymentStatus{
		Name:              deployment.Name,
		Replicas:          *deployment.Spec.Replicas,
		ReadyReplicas:     deployment.Status.ReadyReplicas,
		UpdatedReplicas:   deployment.Status.UpdatedReplicas,
		AvailableReplicas: deployment.Status.AvailableReplicas,
		Conditions:        deployment.Status.Conditions,
	}

	return status, nil
}

// rollbackDeployment rolls back to previous revision
func (da *DeploymentAutomator) rollbackDeployment(ctx context.Context, deploymentName string) error {
	log.Printf("Rolling back deployment %s to previous revision", deploymentName)

	// Get current deployment
	deployment, err := da.clientset.AppsV1().Deployments(da.namespace).Get(ctx, deploymentName, metav1.GetOptions{})
	if err != nil {
		return fmt.Errorf("failed to get deployment: %w", err)
	}

	// Trigger rollback by adding rollback annotation
	if deployment.Annotations == nil {
		deployment.Annotations = make(map[string]string)
	}
	deployment.Annotations["deployment.kubernetes.io/revision"] = ""

	_, err = da.clientset.AppsV1().Deployments(da.namespace).Update(ctx, deployment, metav1.UpdateOptions{})
	if err != nil {
		return fmt.Errorf("failed to rollback deployment: %w", err)
	}

	log.Printf("Successfully initiated rollback for deployment %s", deploymentName)
	return nil
}

// createConfigMap creates a ConfigMap for application configuration
func (da *DeploymentAutomator) createConfigMap(ctx context.Context, name string, data map[string]string) error {
	log.Printf("Creating ConfigMap: %s", name)

	configMap := &corev1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: da.namespace,
		},
		Data: data,
	}

	_, err := da.clientset.CoreV1().ConfigMaps(da.namespace).Create(ctx, configMap, metav1.CreateOptions{})
	if err != nil && !errors.IsAlreadyExists(err) {
		return fmt.Errorf("failed to create ConfigMap: %w", err)
	}

	log.Printf("Successfully created ConfigMap: %s", name)
	return nil
}

// createService creates a Service to expose the application
func (da *DeploymentAutomator) createService(ctx context.Context, serviceName, appName string, port int32, serviceType corev1.ServiceType) error {
	log.Printf("Creating Service: %s", serviceName)

	service := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      serviceName,
			Namespace: da.namespace,
			Labels: map[string]string{
				"app": appName,
			},
		},
		Spec: corev1.ServiceSpec{
			Selector: map[string]string{
				"app": appName,
			},
			Ports: []corev1.ServicePort{
				{
					Name:       "http",
					Port:       port,
					TargetPort: intstr.FromInt(int(port)),
					Protocol:   corev1.ProtocolTCP,
				},
			},
			Type: serviceType,
		},
	}

	_, err := da.clientset.CoreV1().Services(da.namespace).Create(ctx, service, metav1.CreateOptions{})
	if err != nil && !errors.IsAlreadyExists(err) {
		return fmt.Errorf("failed to create Service: %w", err)
	}

	log.Printf("Successfully created Service: %s", serviceName)
	return nil
}

// cleanupApplication removes all resources associated with an application
func (da *DeploymentAutomator) cleanupApplication(ctx context.Context, appName string) error {
	log.Printf("Cleaning up application: %s", appName)

	// Delete Deployment
	err := da.clientset.AppsV1().Deployments(da.namespace).Delete(ctx, appName, metav1.DeleteOptions{})
	if err != nil && !errors.IsNotFound(err) {
		return fmt.Errorf("failed to delete Deployment: %w", err)
	}

	// Delete Service
	err = da.clientset.CoreV1().Services(da.namespace).Delete(ctx, appName, metav1.DeleteOptions{})
	if err != nil && !errors.IsNotFound(err) {
		return fmt.Errorf("failed to delete Service: %w", err)
	}

	// Delete ConfigMap
	err = da.clientset.CoreV1().ConfigMaps(da.namespace).Delete(ctx, appName+"-config", metav1.DeleteOptions{})
	if err != nil && !errors.IsNotFound(err) {
		log.Printf("Note: ConfigMap cleanup failed (may not exist): %v", err)
	}

	log.Printf("Successfully cleaned up application: %s", appName)
	return nil
}

func main() {
	// Initialize Kubernetes client
	var kubeconfig string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = filepath.Join(home, ".kube", "config")
	}

	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		log.Fatal("Failed to build config:", err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal("Failed to create clientset:", err)
	}

	// Create DeploymentAutomator for "default" namespace
	automator := NewDeploymentAutomator(clientset, "default")

	ctx := context.Background()

	// Define application configuration
	appConfig := ApplicationConfig{
		Name:     "web-app",
		Image:    "nginx:1.21",
		Replicas: 3,
		Port:     80,
		Env: map[string]string{
			"ENV":   "production",
			"DEBUG": "false",
		},
		ConfigData: map[string]string{
			"app.conf":   "server_name example.com;",
			"nginx.conf": "worker_processes auto;",
		},
		HealthCheck: true,
		ResourceReqs: ResourceRequirements{
			CPURequest:    "100m",
			MemoryRequest: "128Mi",
			CPULimit:      "500m",
			MemoryLimit:   "512Mi",
		},
	}

	// Deploy the application
	if err := automator.deployApplication(ctx, appConfig); err != nil {
		log.Fatalf("Failed to deploy application: %v", err)
	}

	// Wait for deployment to be ready (timeout: 300 seconds)
	if err := automator.waitForDeploymentReady(ctx, appConfig.Name, 300); err != nil {
		log.Printf("Warning: Deployment may not be fully ready: %v", err)
	}

	// Get and display deployment status
	status, err := automator.getDeploymentStatus(ctx, appConfig.Name)
	if err != nil {
		log.Printf("Failed to get deployment status: %v", err)
	} else {
		fmt.Printf("\n=== Deployment Status ===\n")
		fmt.Printf("Name: %s\n", status.Name)
		fmt.Printf("Replicas: %d\n", status.Replicas)
		fmt.Printf("Ready: %d\n", status.ReadyReplicas)
		fmt.Printf("Updated: %d\n", status.UpdatedReplicas)
		fmt.Printf("Available: %d\n", status.AvailableReplicas)
	}

	// Perform rolling update to nginx:1.22
	if err := automator.rollingUpdate(ctx, appConfig.Name, "nginx:1.22", 1, 1); err != nil {
		log.Printf("Failed to perform rolling update: %v", err)
	}

	// Wait for rolling update to complete
	time.Sleep(10 * time.Second) // Give it a moment to start
	if err := automator.waitForDeploymentReady(ctx, appConfig.Name, 300); err != nil {
		log.Printf("Rolling update may not have completed: %v", err)
	}

	// Scale the deployment to 5 replicas
	if err := automator.scaleDeployment(ctx, appConfig.Name, 5); err != nil {
		log.Printf("Failed to scale deployment: %v", err)
	}

	// Wait for scaling to complete
	time.Sleep(5 * time.Second)
	if err := automator.waitForDeploymentReady(ctx, appConfig.Name, 300); err != nil {
		log.Printf("Scaling may not have completed: %v", err)
	}

	// Get final deployment status
	finalStatus, err := automator.getDeploymentStatus(ctx, appConfig.Name)
	if err != nil {
		log.Printf("Failed to get final deployment status: %v", err)
	} else {
		fmt.Printf("\n=== Final Deployment Status ===\n")
		fmt.Printf("Name: %s\n", finalStatus.Name)
		fmt.Printf("Replicas: %d\n", finalStatus.Replicas)
		fmt.Printf("Ready: %d\n", finalStatus.ReadyReplicas)
		fmt.Printf("Updated: %d\n", finalStatus.UpdatedReplicas)
		fmt.Printf("Available: %d\n", finalStatus.AvailableReplicas)
	}

	// Demonstrate rollback (rollback to previous version)
	log.Println("Demonstrating rollback...")
	if err := automator.rollbackDeployment(ctx, appConfig.Name); err != nil {
		log.Printf("Failed to rollback deployment: %v", err)
	}

	// Wait a bit for rollback
	time.Sleep(10 * time.Second)

	// Cleanup application resources
	if err := automator.cleanupApplication(ctx, appConfig.Name); err != nil {
		log.Printf("Failed to cleanup application: %v", err)
	}

	fmt.Println("Kubernetes deployment automation completed!")
}