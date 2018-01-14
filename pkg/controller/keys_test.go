package controller

import (
	"testing"

	assert "github.com/stretchr/testify/assert"

	errors "k8s.io/apimachinery/pkg/api/errors"
	// metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	// v1alpha1 "github.com/srossross/k8s-test-controller/pkg/apis/tester/v1alpha1"
)

func TestSplitOnce(t *testing.T) {
	// t.Log("Oh noes - something is false")
	// t.Fail()
	a, b := splitOnce("key:value", ":")
	if a != "key" {
		t.Fatalf("a should equal 'key' (got %v)", a)
	}
	if b != "value" {
		t.Fatalf("a should equal 'value' (got %v)", a)
	}

	return
}

func TestGetTestRunFromKeyWhenEmpty(t *testing.T) {
	assert := assert.New(t)

	ctrl := createFakeController()

	tr, err := ctrl.GetTestRunFromKey("ns/key")

	assert.NotNil(err)
	assert.True(errors.IsNotFound(err))
	assert.Nil(tr)
}

// func TestGetTestRunFromKey(t *testing.T) {
// 	assert := assert.New(t)
//
// 	ctrl := createFakeController(
// 		&v1alpha1.TestRun{
// 			ObjectMeta: metav1.ObjectMeta{
// 				Namespace: "foo",
// 				Name:      "run-42",
// 			},
// 		},
// 	)
//
// 	tr, err := ctrl.GetTestRunFromKey("foo/run-42")
//
// 	assert.Nil(err)
// 	assert.NotNil(tr)
// }
