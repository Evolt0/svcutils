package k8s

import (
	"context"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	corev1 "k8s.io/api/core/v1"
	kerrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (c *K8sClient) EnsureSecret(ctx context.Context, namespace string, secret *corev1.Secret) (*corev1.Secret, error) {
	object, err := c.GetSecret(ctx, namespace, secret.Name, metav1.GetOptions{})
	if err != nil {
		if kerrors.IsNotFound(err) {
			object, err = c.CreateSecret(ctx, namespace, secret, metav1.CreateOptions{})
			if err != nil {
				logrus.Errorf("failed to call CreateSecret: %v", err)
				return nil, errors.Wrapf(err, "failed to call CreateSecret")
			}
		}
		return nil, errors.Wrapf(err, "failed to call GetSecret")
	}
	return object, nil
}

func (c *K8sClient) CreateSecret(ctx context.Context, namespace string, secret *corev1.Secret, opts metav1.CreateOptions) (*corev1.Secret, error) {
	return c.clientSet.CoreV1().Secrets(namespace).Create(ctx, secret, opts)
}

func (c *K8sClient) UpdateSecret(ctx context.Context, namespace string, secret *corev1.Secret, opts metav1.UpdateOptions) (*corev1.Secret, error) {
	return c.clientSet.CoreV1().Secrets(namespace).Update(ctx, secret, opts)
}

func (c *K8sClient) DeleteSecret(ctx context.Context, namespace string, name string, opts metav1.DeleteOptions) error {
	return c.clientSet.CoreV1().Secrets(namespace).Delete(ctx, name, opts)
}

func (c *K8sClient) GetSecret(ctx context.Context, namespace string, name string, opts metav1.GetOptions) (*corev1.Secret, error) {
	return c.clientSet.CoreV1().Secrets(namespace).Get(ctx, name, opts)
}

func (c *K8sClient) ListSecret(ctx context.Context, namespace string, opts metav1.ListOptions) (*corev1.SecretList, error) {
	return c.clientSet.CoreV1().Secrets(namespace).List(ctx, opts)
}
