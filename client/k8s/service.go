package k8s

import (
	"context"
	"github.com/pkg/errors"
	corev1 "k8s.io/api/core/v1"
	kerrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (c *K8sClient) EnsureService(ctx context.Context, namespace string, service *corev1.Service) (*corev1.Service, error) {
	object, err := c.GetService(ctx, namespace, service.Name, metav1.GetOptions{})
	if err != nil {
		if !kerrors.IsNotFound(err) {
			return nil, errors.Wrapf(err, "failed to call k8s Getservice")
		}
		object, err = c.CreateService(ctx, namespace, service, metav1.CreateOptions{})
		if err != nil {
			return nil, errors.Wrapf(err, "failed to call k8s Createservice")
		}
	}
	return object, nil
}

func (c *K8sClient) CreateService(ctx context.Context, namespace string, service *corev1.Service, opts metav1.CreateOptions) (*corev1.Service, error) {
	return c.clientSet.CoreV1().Services(namespace).Create(ctx, service, opts)
}

func (c *K8sClient) UpdateService(ctx context.Context, namespace string, service *corev1.Service, opts metav1.UpdateOptions) (*corev1.Service, error) {
	return c.clientSet.CoreV1().Services(namespace).Update(ctx, service, opts)

}

func (c *K8sClient) DeleteService(ctx context.Context, namespace string, name string, opts metav1.DeleteOptions) error {
	return c.clientSet.CoreV1().Services(namespace).Delete(ctx, name, opts)
}

func (c *K8sClient) GetService(ctx context.Context, namespace string, name string, opts metav1.GetOptions) (*corev1.Service, error) {
	return c.clientSet.CoreV1().Services(namespace).Get(ctx, name, opts)
}

func (c *K8sClient) ListService(ctx context.Context, namespace string, opts metav1.ListOptions) (*corev1.ServiceList, error) {
	return c.clientSet.CoreV1().Services(namespace).List(ctx, opts)
}
