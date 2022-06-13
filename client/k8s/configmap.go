package k8s

import (
	"context"
	"github.com/pkg/errors"
	corev1 "k8s.io/api/core/v1"
	kerrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (c *K8sClient) EnsureConfigMap(ctx context.Context, namespace string, configMap *corev1.ConfigMap) (*corev1.ConfigMap, error) {
	object, err := c.GetConfigMap(ctx, namespace, configMap.Name, metav1.GetOptions{})
	if err != nil {
		if !kerrors.IsNotFound(err) {
			return nil, errors.Wrapf(err, "failed to call k8s GetConfigMap")
		}
		object, err = c.CreateConfigMap(ctx, namespace, configMap, metav1.CreateOptions{})
		if err != nil {
			return nil, errors.Wrapf(err, "failed to call k8s CreateConfigMap")
		}
	}
	return object, nil
}

func (c *K8sClient) CreateConfigMap(ctx context.Context, namespace string, configMap *corev1.ConfigMap, opts metav1.CreateOptions) (*corev1.ConfigMap, error) {
	return c.clientSet.CoreV1().ConfigMaps(namespace).Create(ctx, configMap, opts)
}

func (c *K8sClient) UpdateConfigMap(ctx context.Context, namespace string, configMap *corev1.ConfigMap, opts metav1.UpdateOptions) (*corev1.ConfigMap, error) {
	return c.clientSet.CoreV1().ConfigMaps(namespace).Update(ctx, configMap, opts)
}

func (c *K8sClient) DeleteConfigMap(ctx context.Context, namespace string, name string, opts metav1.DeleteOptions) error {
	return c.clientSet.CoreV1().ConfigMaps(namespace).Delete(ctx, name, opts)
}

func (c *K8sClient) GetConfigMap(ctx context.Context, namespace string, name string, opts metav1.GetOptions) (*corev1.ConfigMap, error) {
	return c.clientSet.CoreV1().ConfigMaps(namespace).Get(ctx, name, opts)
}

func (c *K8sClient) ListConfigMap(ctx context.Context, namespace string, opts metav1.ListOptions) (*corev1.ConfigMapList, error) {
	return c.clientSet.CoreV1().ConfigMaps(namespace).List(ctx, opts)
}
