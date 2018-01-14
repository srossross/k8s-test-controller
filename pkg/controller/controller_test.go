package controller

import (
	"testing"
	"time"

	assert "github.com/stretchr/testify/assert"

	runtime "k8s.io/apimachinery/pkg/runtime"
	fakev1 "k8s.io/client-go/kubernetes/typed/core/v1/fake"

	fake "github.com/srossross/k8s-test-controller/pkg/client/fake"
	factory "github.com/srossross/k8s-test-controller/pkg/informers/externalversions"
)

func createFakeController(objects ...runtime.Object) Interface {
	cl := fake.NewSimpleClientset(objects...)
	coreV1Client := &fakev1.FakeCoreV1{}
	sharedFactory := factory.NewSharedInformerFactory(cl, time.Second*30)
	return NewTestController(&sharedFactory, cl, coreV1Client)
}

func TestFakeController(t *testing.T) {
	assert := assert.New(t)

	ctrl := createFakeController()

	assert.NotNil(ctrl)
}
