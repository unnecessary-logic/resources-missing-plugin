package scan

import (
	"context"
	"reflect"
	"testing"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes/fake"
)

func podhelper(ns, name string, hasResourcesMem bool, hasResourcesCpu bool, hasResourcesMemInit bool, hasResourcesCpuInit bool) *corev1.Pod {
	container := corev1.Container{Name: "app"}
	initContainer := corev1.Container{Name: "initApp"}

	container.Resources = corev1.ResourceRequirements{
		Requests: corev1.ResourceList{},
		Limits:   corev1.ResourceList{},
	}

	initContainer.Resources = corev1.ResourceRequirements{
		Requests: corev1.ResourceList{},
		Limits:   corev1.ResourceList{},
	}

	if hasResourcesMem {
		container.Resources.Requests[corev1.ResourceMemory] = resource.MustParse("100Mi")
		container.Resources.Limits[corev1.ResourceMemory] = resource.MustParse("200Mi")
	}
	if hasResourcesCpu {
		container.Resources.Requests[corev1.ResourceCPU] = resource.MustParse("0.1")
		container.Resources.Limits[corev1.ResourceCPU] = resource.MustParse("0.2")
	}

	if hasResourcesMemInit {
		initContainer.Resources.Requests[corev1.ResourceMemory] = resource.MustParse("100Mi")
		initContainer.Resources.Limits[corev1.ResourceMemory] = resource.MustParse("200Mi")
	}
	if hasResourcesCpuInit {
		initContainer.Resources.Requests[corev1.ResourceCPU] = resource.MustParse("0.1")
		initContainer.Resources.Limits[corev1.ResourceCPU] = resource.MustParse("0.2")
	}

	return &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{Namespace: ns, Name: name},
		Spec: corev1.PodSpec{
			Containers:     []corev1.Container{container},
			InitContainers: []corev1.Container{initContainer}},
	}
}

func TestRunScan(t *testing.T) {
	tests := []struct {
		name      string
		namespace string
		pods      []runtime.Object
		wantPods  []string
	}{
		{
			name:      "all namespaces",
			namespace: "",
			pods: []runtime.Object{
				podhelper("a", "a", false, false, false, false),
				podhelper("b", "b", false, false, false, false),
				podhelper("c", "c", true, true, true, true),
				podhelper("d", "d", false, true, true, false),
			},
			wantPods: []string{"a", "b", "d"},
		},
		{
			name:      "select namespace",
			namespace: "a",
			pods: []runtime.Object{
				podhelper("a", "a", false, false, false, false),
				podhelper("b", "b", false, false, false, false),
			},
			wantPods: []string{"a"},
		},
		{
			name:      "no resources found all",
			namespace: "",
			pods: []runtime.Object{
				podhelper("a", "a", true, true, true, true),
				podhelper("b", "b", true, true, true, true),
			},
		},
		{
			name:      "no resources found select",
			namespace: "a",
			pods: []runtime.Object{
				podhelper("a", "a", true, true, true, true),
				podhelper("b", "b", false, true, true, false),
			},
		},
		{
			name:      "initContainers found",
			namespace: "a",
			pods: []runtime.Object{
				podhelper("a", "a", true, true, true, false),
				podhelper("b", "b", false, true, true, false),
			},
			wantPods: []string{"a"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := fake.NewSimpleClientset(tt.pods...)

			got, err := RunScan(context.Background(), client, tt.namespace)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			var names []string
			for _, r := range got {
				names = append(names, r.Pod)
			}
			if !reflect.DeepEqual(names, tt.wantPods) {
				t.Errorf("got %v, want %v", names, tt.wantPods)
			}
		})
	}

}
