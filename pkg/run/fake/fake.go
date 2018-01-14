package run

import (
	corev1 "k8s.io/api/core/v1"

	v1alpha1 "github.com/srossross/k8s-test-controller/pkg/apis/tester/v1alpha1"
	controller "github.com/srossross/k8s-test-controller/pkg/controller"
	run "github.com/srossross/k8s-test-controller/pkg/run"
)

type fakeRunner struct {
}

func (f *fakeRunner) PodStateChange(ctrl controller.Interface, testRun *v1alpha1.TestRun, pod *corev1.Pod) error {
	return nil
}

func (f *fakeRunner) UpdateTestRun(ctrl controller.Interface, testRun *v1alpha1.TestRun) error {
	return nil
}

// New creates a fake runner
func New() run.Interface {
	return &fakeRunner{}
}
