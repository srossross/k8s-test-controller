package loop

import (
	"testing"

	fakecontroller "github.com/srossross/k8s-test-controller/pkg/controller/fake"
	fakerun "github.com/srossross/k8s-test-controller/pkg/run/fake"
	"github.com/stretchr/testify/assert"
)

func TestTake_BadKey(t *testing.T) {
	assert := assert.New(t)
	ctrl := fakecontroller.NewController()
	runner := fakerun.New()
	err := take(ctrl, runner, "BadKey:namespace/name")
	assert.NotNil(err)
}

// func TestTake_ReconsilePodStatus(t *testing.T) {
// 	assert := assert.New(t)
// 	ctrl := fakecontroller.NewController()
// 	runner := fakerun.New()
// 	err := take(ctrl, runner, "Pod:namespace/name")
// 	assert.Nil(err)
// }
