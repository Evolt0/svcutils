package k8s

import (
	"context"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"k8s.io/api/extensions/v1beta1"
	kerrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (c *K8sClient) EnsureIngress(ctx context.Context, namespace string, ingress *v1beta1.Ingress) (*v1beta1.Ingress, error) {
	object, err := c.GetIngress(ctx, namespace, ingress.Name, metav1.GetOptions{})
	if err != nil {
		if kerrors.IsNotFound(err) {
			object, err = c.CreateIngress(ctx, namespace, ingress, metav1.CreateOptions{})
			if err != nil {
				logrus.Errorf("failed to call k8s CreateIngress: %v", err)
				return nil, errors.Wrapf(err, "failed to call k8s CreateIngress")
			}
		}
		return nil, errors.Wrapf(err, "failed to call k8s GetIngress")
	}
	return object, nil
}

func (c *K8sClient) CreateIngress(ctx context.Context, namespace string, ingress *v1beta1.Ingress, opts metav1.CreateOptions) (*v1beta1.Ingress, error) {
	return c.clientSet.ExtensionsV1beta1().Ingresses(namespace).Create(ctx, ingress, opts)
}

func (c *K8sClient) UpdateIngress(ctx context.Context, namespace string, ingress *v1beta1.Ingress, opts metav1.UpdateOptions) (*v1beta1.Ingress, error) {
	return c.clientSet.ExtensionsV1beta1().Ingresses(namespace).Update(ctx, ingress, opts)
}

func (c *K8sClient) DeleteIngress(ctx context.Context, namespace string, name string, opts metav1.DeleteOptions) error {
	return c.clientSet.ExtensionsV1beta1().Ingresses(namespace).Delete(ctx, name, opts)
}

func (c *K8sClient) GetIngress(ctx context.Context, namespace string, name string, opts metav1.GetOptions) (*v1beta1.Ingress, error) {
	return c.clientSet.ExtensionsV1beta1().Ingresses(namespace).Get(ctx, name, opts)
}

func (c *K8sClient) ListIngress(ctx context.Context, namespace string, opts metav1.ListOptions) (*v1beta1.IngressList, error) {
	return c.clientSet.ExtensionsV1beta1().Ingresses(namespace).List(ctx, opts)
}
