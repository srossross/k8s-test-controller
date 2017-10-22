package run

import (
	corev1 "k8s.io/api/core/v1"

	v1alpha1 "github.com/srossross/k8s-test-controller/pkg/apis/tester/v1alpha1"
	controller "github.com/srossross/k8s-test-controller/pkg/controller"
)

// Interface to be passed into work function
type Interface interface {
	PodStateChange(ctrl controller.Interface, testRun *v1alpha1.TestRun, pod *corev1.Pod) error
	UpdateTestRun(ctrl controller.Interface, testRun *v1alpha1.TestRun) error
}

type runner struct {
}

func New() Interface {
	return &runner{}
}
