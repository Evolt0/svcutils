package k8s

import (
	"context"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (c *K8sClient) CreatePod(ctx context.Context, namespace string, pod *corev1.Pod, opts metav1.CreateOptions) (*corev1.Pod, error) {
	return c.clientSet.CoreV1().Pods(namespace).Create(ctx, pod, opts)
}

func (c *K8sClient) UpdatePod(ctx context.Context, namespace string, pod *corev1.Pod, opts metav1.UpdateOptions) (*corev1.Pod, error) {
	return c.clientSet.CoreV1().Pods(namespace).Update(ctx, pod, opts)
}

func (c *K8sClient) DeletePod(ctx context.Context, namespace string, name string, opts metav1.DeleteOptions) error {
	return c.clientSet.CoreV1().Pods(namespace).Delete(ctx, name, opts)
}

func (c *K8sClient) GetPod(ctx context.Context, namespace string, name string, opts metav1.GetOptions) (*corev1.Pod, error) {
	return c.clientSet.CoreV1().Pods(namespace).Get(ctx, name, opts)
}

func (c *K8sClient) ListPod(ctx context.Context, namespace string, opts metav1.ListOptions) (*corev1.PodList, error) {
	return c.clientSet.CoreV1().Pods(namespace).List(ctx, opts)
}
