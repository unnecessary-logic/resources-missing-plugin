package scan

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type Result struct {
	Namespace string
	Pod       string
	Container []string
}

var required = []corev1.ResourceName{
	corev1.ResourceCPU,
	corev1.ResourceMemory,
}

func missingResourcesHelper(r corev1.ResourceRequirements) bool {
	for _, name := range required {
		if _, ok := r.Requests[name]; !ok {
			return true
		}
		if _, ok := r.Limits[name]; !ok {
			return true
		}
	}
	return false
}

func FindMissingResources(pod *corev1.Pod) []Result {
	var res []Result
	var names []string
	check := func(cont []corev1.Container) {
		for i := range cont {
			container := &cont[i]
			if missingResourcesHelper(container.Resources) {
				names = append(names, container.Name)
			}
		}
	}

	check(pod.Spec.Containers)
	check(pod.Spec.InitContainers)

	if len(names) > 0 {
		res = append(res, Result{
			Namespace: pod.Namespace,
			Pod:       pod.Name,
			Container: names,
		})
	}
	return res
}

func RunScan(ctx context.Context, client kubernetes.Interface, namespace string) ([]Result, error) {
	scan, err := client.CoreV1().Pods(namespace).List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	var res []Result
	for i := range scan.Items {
		res = append(res, FindMissingResources(&scan.Items[i])...)
	}

	return res, nil
}
