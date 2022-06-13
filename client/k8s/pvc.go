package k8s

import (
	"context"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	corev1 "k8s.io/api/core/v1"
	kerrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (c *K8sClient) EnsurePersistentVolumeClaim(ctx context.Context, namespace string, persistentVolumeClaim *corev1.PersistentVolumeClaim) (*corev1.PersistentVolumeClaim, error) {
	object, err := c.GetPersistentVolumeClaim(ctx, namespace, persistentVolumeClaim.Name, metav1.GetOptions{})
	if err != nil {
		if kerrors.IsNotFound(err) {
			object, err = c.CreatePersistentVolumeClaim(ctx, namespace, persistentVolumeClaim, metav1.CreateOptions{})
			if err != nil {
				logrus.Errorf("failed to call k8s CreatePersistentVolumeClaim: %v", err)
				return nil, errors.Wrapf(err, "failed to call k8s CreatePersistentVolumeClaim")
			}
		}
		return nil, errors.Wrapf(err, "failed to call k8s GetPersistentVolumeClaim")
	}
	return object, nil
}

func (c *K8sClient) CreatePersistentVolumeClaim(ctx context.Context, namespace string, persistentVolumeClaim *corev1.PersistentVolumeClaim, opts metav1.CreateOptions) (*corev1.PersistentVolumeClaim, error) {
	return c.clientSet.CoreV1().PersistentVolumeClaims(namespace).Create(ctx, persistentVolumeClaim, opts)
}

func (c *K8sClient) UpdatePersistentVolumeClaim(ctx context.Context, namespace string, persistentVolumeClaim *corev1.PersistentVolumeClaim, opts metav1.UpdateOptions) (*corev1.PersistentVolumeClaim, error) {
	return c.clientSet.CoreV1().PersistentVolumeClaims(namespace).Update(ctx, persistentVolumeClaim, opts)
}

func (c *K8sClient) DeletePersistentVolumeClaim(ctx context.Context, namespace string, name string, opts metav1.DeleteOptions) error {
	return c.clientSet.CoreV1().PersistentVolumeClaims(namespace).Delete(ctx, name, opts)
}

func (c *K8sClient) GetPersistentVolumeClaim(ctx context.Context, namespace string, name string, opts metav1.GetOptions) (*corev1.PersistentVolumeClaim, error) {
	return c.clientSet.CoreV1().PersistentVolumeClaims(namespace).Get(ctx, name, opts)

}

func (c *K8sClient) ListPersistentVolumeClaim(ctx context.Context, namespace string, opts metav1.ListOptions) (*corev1.PersistentVolumeClaimList, error) {
	return c.clientSet.CoreV1().PersistentVolumeClaims(namespace).List(ctx, opts)

}
