package k8s

import (
	"context"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	corev1 "k8s.io/api/core/v1"
	kerrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (c *K8sClient) EnsurePersistentVolume(ctx context.Context, persistentVolume *corev1.PersistentVolume) (*corev1.PersistentVolume, error) {
	object, err := c.GetPersistentVolume(ctx, persistentVolume.Name, metav1.GetOptions{})
	if err != nil {
		if kerrors.IsNotFound(err) {
			object, err = c.CreatePersistentVolume(ctx, persistentVolume, metav1.CreateOptions{})
			if err != nil {
				logrus.Errorf("failed to call k8s CreatePersistentVolume: %v", err)
				return nil, errors.Wrapf(err, "failed to call k8s CreatePersistentVolume")
			}
		}
		return nil, errors.Wrapf(err, "failed to call k8s GetPersistentVolume")
	}
	return object, nil
}

func (c *K8sClient) CreatePersistentVolume(ctx context.Context, persistentVolume *corev1.PersistentVolume, opts metav1.CreateOptions) (*corev1.PersistentVolume, error) {
	return c.clientSet.CoreV1().PersistentVolumes().Create(ctx, persistentVolume, opts)
}

func (c *K8sClient) UpdatePersistentVolume(ctx context.Context, persistentVolume *corev1.PersistentVolume, opts metav1.UpdateOptions) (*corev1.PersistentVolume, error) {
	return c.clientSet.CoreV1().PersistentVolumes().Update(ctx, persistentVolume, opts)
}

func (c *K8sClient) DeletePersistentVolume(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.clientSet.CoreV1().PersistentVolumes().Delete(ctx, name, opts)
}

func (c *K8sClient) GetPersistentVolume(ctx context.Context, name string, opts metav1.GetOptions) (*corev1.PersistentVolume, error) {
	return c.clientSet.CoreV1().PersistentVolumes().Get(ctx, name, opts)
}

func (c *K8sClient) ListPersistentVolume(ctx context.Context, opts metav1.ListOptions) (*corev1.PersistentVolumeList, error) {
	return c.clientSet.CoreV1().PersistentVolumes().List(ctx, opts)
}
