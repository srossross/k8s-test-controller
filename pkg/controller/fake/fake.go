package fake

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/labels"
	typedv1 "k8s.io/client-go/kubernetes/typed/core/v1"
	cache "k8s.io/client-go/tools/cache"

	v1alpha1 "github.com/srossross/k8s-test-controller/pkg/apis/tester/v1alpha1"
	srossrossv1alpha1 "github.com/srossross/k8s-test-controller/pkg/client/typed/srossross/v1alpha1"
	controller "github.com/srossross/k8s-test-controller/pkg/controller"
	listerV1alpha1 "github.com/srossross/k8s-test-controller/pkg/listers/srossross/v1alpha1"
)

type fakeController struct {
}

func (f *fakeController) TestTemplateLister() listerV1alpha1.TestTemplateLister {
	return nil
}

func (f *fakeController) TestRunLister() listerV1alpha1.TestRunLister {
	return nil
}

func (f *fakeController) SrossrossV1alpha1() srossrossv1alpha1.SrossrossV1alpha1Interface {
	return nil
}

func (f *fakeController) PodInformer() cache.SharedIndexInformer {
	return nil
}
func (f *fakeController) CoreV1() typedv1.CoreV1Interface {
	return nil
}

func (f *fakeController) GetTestRunFromKey(key string) (*v1alpha1.TestRun, error) {
	return nil, nil
}
func (f *fakeController) GetPodAndTestRunFromKey(key string) (*v1alpha1.TestRun, *corev1.Pod, error) {
	return nil, nil, nil
}
func (f *fakeController) TestRunnerRemovePodsForDeletedTest(key string) error {
	return nil
}

func (f *fakeController) CreatePod(Namespace string, pod *corev1.Pod) (*corev1.Pod, error) {
	return nil, nil
}
func (f *fakeController) ListPods(Namespace string, selector labels.Selector) ([]*corev1.Pod, error) {
	return nil, nil
}
func (f *fakeController) GetPod(Namespace string, Name string) (*corev1.Pod, error) {
	return nil, nil
}

// NewController this is a fake
func NewController() controller.Interface {
	return &fakeController{}
}
