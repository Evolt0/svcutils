package k8s

import (
	"context"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/api/extensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

type Client interface {

	// Deployment
	EnsureDeployment(ctx context.Context, namespace string, deployment *appsv1.Deployment) (*appsv1.Deployment, error)
	CreateDeployment(ctx context.Context, namespace string, deployment *appsv1.Deployment, opts metav1.CreateOptions) (*appsv1.Deployment, error)
	UpdateDeployment(ctx context.Context, namespace string, deployment *appsv1.Deployment, opts metav1.UpdateOptions) (*appsv1.Deployment, error)
	DeleteDeployment(ctx context.Context, namespace string, name string, opts metav1.DeleteOptions) error
	GetDeployment(ctx context.Context, namespace string, name string, opts metav1.GetOptions) (*appsv1.Deployment, error)
	ListDeployment(ctx context.Context, namespace string, opts metav1.ListOptions) (*appsv1.DeploymentList, error)

	// POD
	CreatePod(ctx context.Context, namespace string, pod *corev1.Pod, opts metav1.CreateOptions) (*corev1.Pod, error)
	UpdatePod(ctx context.Context, namespace string, pod *corev1.Pod, opts metav1.UpdateOptions) (*corev1.Pod, error)
	DeletePod(ctx context.Context, namespace string, name string, opts metav1.DeleteOptions) error
	GetPod(ctx context.Context, namespace string, name string, opts metav1.GetOptions) (*corev1.Pod, error)
	ListPod(ctx context.Context, namespace string, opts metav1.ListOptions) (*corev1.PodList, error)

	// Service
	EnsureService(ctx context.Context, namespace string, service *corev1.Service) (*corev1.Service, error)
	CreateService(ctx context.Context, namespace string, service *corev1.Service, opts metav1.CreateOptions) (*corev1.Service, error)
	UpdateService(ctx context.Context, namespace string, service *corev1.Service, opts metav1.UpdateOptions) (*corev1.Service, error)
	DeleteService(ctx context.Context, namespace string, name string, opts metav1.DeleteOptions) error
	GetService(ctx context.Context, namespace string, name string, opts metav1.GetOptions) (*corev1.Service, error)
	ListService(ctx context.Context, namespace string, opts metav1.ListOptions) (*corev1.ServiceList, error)

	// Ingress
	EnsureIngress(ctx context.Context, namespace string, ingress *v1beta1.Ingress) (*v1beta1.Ingress, error)
	CreateIngress(ctx context.Context, namespace string, ingress *v1beta1.Ingress, opts metav1.CreateOptions) (*v1beta1.Ingress, error)
	UpdateIngress(ctx context.Context, namespace string, ingress *v1beta1.Ingress, opts metav1.UpdateOptions) (*v1beta1.Ingress, error)
	DeleteIngress(ctx context.Context, namespace string, name string, opts metav1.DeleteOptions) error
	GetIngress(ctx context.Context, namespace string, name string, opts metav1.GetOptions) (*v1beta1.Ingress, error)
	ListIngress(ctx context.Context, namespace string, opts metav1.ListOptions) (*v1beta1.IngressList, error)

	// ConfigMap
	EnsureConfigMap(ctx context.Context, namespace string, configMap *corev1.ConfigMap) (*corev1.ConfigMap, error)
	CreateConfigMap(ctx context.Context, namespace string, configMap *corev1.ConfigMap, opts metav1.CreateOptions) (*corev1.ConfigMap, error)
	UpdateConfigMap(ctx context.Context, namespace string, configMap *corev1.ConfigMap, opts metav1.UpdateOptions) (*corev1.ConfigMap, error)
	DeleteConfigMap(ctx context.Context, namespace string, name string, opts metav1.DeleteOptions) error
	GetConfigMap(ctx context.Context, namespace string, name string, opts metav1.GetOptions) (*corev1.ConfigMap, error)
	ListConfigMap(ctx context.Context, namespace string, opts metav1.ListOptions) (*corev1.ConfigMapList, error)
	// Namespace
	EnsureNamespace(ctx context.Context, namespace *corev1.Namespace) (*corev1.Namespace, error)
	CreateNamespace(ctx context.Context, namespace *corev1.Namespace, opts metav1.CreateOptions) (*corev1.Namespace, error)
	UpdateNamespace(ctx context.Context, namespace *corev1.Namespace, opts metav1.UpdateOptions) (*corev1.Namespace, error)
	DeleteNamespace(ctx context.Context, name string, opts metav1.DeleteOptions) error
	GetNamespace(ctx context.Context, name string, opts metav1.GetOptions) (*corev1.Namespace, error)
	ListNamespace(ctx context.Context, opts metav1.ListOptions) (*corev1.NamespaceList, error)

	// PersistentVolume
	EnsurePersistentVolume(ctx context.Context, persistentVolume *corev1.PersistentVolume) (*corev1.PersistentVolume, error)
	CreatePersistentVolume(ctx context.Context, persistentVolume *corev1.PersistentVolume, opts metav1.CreateOptions) (*corev1.PersistentVolume, error)
	UpdatePersistentVolume(ctx context.Context, persistentVolume *corev1.PersistentVolume, opts metav1.UpdateOptions) (*corev1.PersistentVolume, error)
	DeletePersistentVolume(ctx context.Context, name string, opts metav1.DeleteOptions) error
	GetPersistentVolume(ctx context.Context, name string, opts metav1.GetOptions) (*corev1.PersistentVolume, error)
	ListPersistentVolume(ctx context.Context, opts metav1.ListOptions) (*corev1.PersistentVolumeList, error)

	// PersistentVolumeClaim
	EnsurePersistentVolumeClaim(ctx context.Context, namespace string, persistentVolumeClaim *corev1.PersistentVolumeClaim) (*corev1.PersistentVolumeClaim, error)
	CreatePersistentVolumeClaim(ctx context.Context, namespace string, persistentVolumeClaim *corev1.PersistentVolumeClaim, opts metav1.CreateOptions) (*corev1.PersistentVolumeClaim, error)
	UpdatePersistentVolumeClaim(ctx context.Context, namespace string, persistentVolumeClaim *corev1.PersistentVolumeClaim, opts metav1.UpdateOptions) (*corev1.PersistentVolumeClaim, error)
	DeletePersistentVolumeClaim(ctx context.Context, namespace string, name string, opts metav1.DeleteOptions) error
	GetPersistentVolumeClaim(ctx context.Context, namespace string, name string, opts metav1.GetOptions) (*corev1.PersistentVolumeClaim, error)
	ListPersistentVolumeClaim(ctx context.Context, namespace string, opts metav1.ListOptions) (*corev1.PersistentVolumeClaimList, error)

	// Secret
	EnsureSecret(ctx context.Context, namespace string, secret *corev1.Secret) (*corev1.Secret, error)
	CreateSecret(ctx context.Context, namespace string, secret *corev1.Secret, opts metav1.CreateOptions) (*corev1.Secret, error)
	UpdateSecret(ctx context.Context, namespace string, secret *corev1.Secret, opts metav1.UpdateOptions) (*corev1.Secret, error)
	DeleteSecret(ctx context.Context, namespace string, name string, opts metav1.DeleteOptions) error
	GetSecret(ctx context.Context, namespace string, name string, opts metav1.GetOptions) (*corev1.Secret, error)
	ListSecret(ctx context.Context, namespace string, opts metav1.ListOptions) (*corev1.SecretList, error)
}

type K8sClient struct {
	InCluster  bool
	ConfigPath string
	ConfigData []byte
	clientSet  *kubernetes.Clientset
}

func NewClient(opts ...Option) (*K8sClient, error) {
	c := &K8sClient{}
	for _, opt := range opts {
		opt(c)
	}
	if err := c.setup(); err != nil {
		return nil, err
	}
	return c, nil
}

func (c *K8sClient) Init() {
	if err := c.setup(); err != nil {
		logrus.Errorf("failed to setup: %v", err)
	}
}

func (c *K8sClient) setup() error {
	var (
		config *rest.Config
		err    error
	)
	if c.InCluster {
		config, err = rest.InClusterConfig()
	} else {
		if len(c.ConfigData) != 0 {
			config, err = clientcmd.RESTConfigFromKubeConfig(c.ConfigData)
		} else {
			config, err = clientcmd.BuildConfigFromFlags("", c.ConfigPath)
		}
	}

	if err != nil {
		return errors.Wrapf(err, "failed to build config")
	}
	base, err := kubernetes.NewForConfig(config)
	if err != nil {
		return errors.Wrapf(err, "failed to new clientSet")
	}
	c.clientSet = base
	return nil
}
