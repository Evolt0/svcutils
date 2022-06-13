package k8s

import (
	"context"
	"fmt"
	"github.com/Evolt0/svcutils/utils"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	kerrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"time"
)

func (c *K8sClient) EnsureDeployment(ctx context.Context, namespace string, deployment *appsv1.Deployment) (*appsv1.Deployment, error) {
	object, err := c.GetDeployment(ctx, namespace, deployment.Name, metav1.GetOptions{})
	if err != nil {
		if !kerrors.IsNotFound(err) {
			return nil, errors.Wrapf(err, "failed to call k8s GetDeployment")
		}
		object, err = c.CreateDeployment(ctx, namespace, deployment, metav1.CreateOptions{})
		if err != nil {
			return nil, errors.Wrapf(err, "failed to call k8s CreateDeployment")
		}
	}
	// 查看deployment的启动状态
	err = utils.RetryFunc(func() error {
		object, err = c.GetDeployment(ctx, namespace, deployment.Name, metav1.GetOptions{})
		if err != nil {
			return err
		}
		if object.Status.ReadyReplicas == object.Status.Replicas {
			return nil
		}
		logrus.Warnf("%s ready replicas is not expected, wait time to watch", deployment.Name)
		return errors.New(fmt.Sprintf("invalid replicas status: %v", object.Status))
	}, 10, 3*time.Second)
	if err != nil {
		return nil, err
	}

	// 查看pod的启动状态
	err = utils.RetryFunc(func() error {
		pods, err := c.ListPod(ctx, namespace, metav1.ListOptions{LabelSelector: fmt.Sprintf("app=%s", deployment.Spec.Template.Labels["app"])})
		if err != nil {
			return err
		}
		for _, pod := range pods.Items {
			if pod.Status.Phase != corev1.PodRunning {
				logrus.Warnf("%s pod is not running, wait time to watch", pod.Name)
				return errors.New(fmt.Sprintf("invalid pod status: %v", pod.Status))
			}
		}
		return nil
	}, 10, 3*time.Second)
	if err != nil {
		return nil, err
	}

	return object, nil
}

func (c *K8sClient) CreateDeployment(ctx context.Context, namespace string, deployment *appsv1.Deployment, opts metav1.CreateOptions) (*appsv1.Deployment, error) {
	return c.clientSet.AppsV1().Deployments(namespace).Create(ctx, deployment, opts)
}

func (c *K8sClient) UpdateDeployment(ctx context.Context, namespace string, deployment *appsv1.Deployment, opts metav1.UpdateOptions) (*appsv1.Deployment, error) {
	return c.clientSet.AppsV1().Deployments(namespace).Update(ctx, deployment, opts)
}

func (c *K8sClient) DeleteDeployment(ctx context.Context, namespace string, name string, opts metav1.DeleteOptions) error {
	return c.clientSet.AppsV1().Deployments(namespace).Delete(ctx, name, opts)
}

func (c *K8sClient) GetDeployment(ctx context.Context, namespace string, name string, opts metav1.GetOptions) (*appsv1.Deployment, error) {
	return c.clientSet.AppsV1().Deployments(namespace).Get(ctx, name, opts)
}

func (c *K8sClient) ListDeployment(ctx context.Context, namespace string, opts metav1.ListOptions) (*appsv1.DeploymentList, error) {
	return c.clientSet.AppsV1().Deployments(namespace).List(ctx, opts)
}
