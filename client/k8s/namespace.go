package k8s

import (
	"context"
	"fmt"
	"github.com/pkg/errors"

	corev1 "k8s.io/api/core/v1"
	kerrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (c *K8sClient) EnsureNamespace(
	ctx context.Context,
	namespace *corev1.Namespace,
) (*corev1.Namespace, error) {
	object, err := c.GetNamespace(ctx, namespace.Name, metav1.GetOptions{})
	if err != nil {
		if !kerrors.IsNotFound(err) {
			return nil, errors.Wrapf(err, "failed to call k8s GetNamespace")
		}
		object, err = c.CreateNamespace(ctx, namespace, metav1.CreateOptions{})
		if err != nil {
			return nil, errors.Wrapf(err, "failed to call k8s CreateNamespace")
		}
	}
	switch object.Status.Phase {
	case corev1.NamespaceActive:
		return object, nil
	default:
		return nil, fmt.Errorf("invalid phase(%s)", object.Status.Phase)
	}
}

func (c *K8sClient) CreateNamespace(
	ctx context.Context,
	namespace *corev1.Namespace,
	opts metav1.CreateOptions,
) (*corev1.Namespace, error) {
	return c.clientSet.CoreV1().Namespaces().Create(ctx, namespace, opts)
}

func (c *K8sClient) UpdateNamespace(
	ctx context.Context,
	namespace *corev1.Namespace,
	opts metav1.UpdateOptions,
) (*corev1.Namespace, error) {
	return c.clientSet.CoreV1().Namespaces().Update(ctx, namespace, opts)
}

func (c *K8sClient) DeleteNamespace(
	ctx context.Context,
	name string,
	opts metav1.DeleteOptions,
) error {
	return c.clientSet.CoreV1().Namespaces().Delete(ctx, name, opts)
}

func (c *K8sClient) GetNamespace(
	ctx context.Context,
	name string,
	opts metav1.GetOptions,
) (*corev1.Namespace, error) {
	return c.clientSet.CoreV1().Namespaces().Get(ctx, name, opts)
}

func (c *K8sClient) ListNamespace(
	ctx context.Context,
	opts metav1.ListOptions,
) (*corev1.NamespaceList, error) {
	return c.clientSet.CoreV1().Namespaces().List(ctx, opts)
}
