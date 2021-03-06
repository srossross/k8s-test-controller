package run

import (
	"fmt"
	"log"
	"os"
	"time"

	v1alpha1 "github.com/srossross/k8s-test-controller/pkg/apis/tester/v1alpha1"
	controller "github.com/srossross/k8s-test-controller/pkg/controller"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// CreateTestRunEvent creates an event
func CreateTestRunEvent(
	ctrl controller.Interface,
	testRun *v1alpha1.TestRun,
	testName string,
	Reason string,
	Message string,
) error {

	objectReference := v1.ObjectReference{
		// FIXME: not sure why testRun.Kind is empty
		Kind:      "TestRun",
		Namespace: testRun.Namespace,
		Name:      testRun.Name,
		UID:       testRun.UID,
		// FIXME: not sure why testRun.APIVersion is empty
		APIVersion:      APIVersion,
		ResourceVersion: testRun.ResourceVersion,
	}

	hostname, err := os.Hostname()
	if err != nil {
		hostname = "hostname"
	}
	now := metav1.Time{Time: time.Now()}
	event := &v1.Event{
		TypeMeta: metav1.TypeMeta{},
		ObjectMeta: metav1.ObjectMeta{
			GenerateName: "test-run-event",
			Labels: map[string]string{
				"test-run":  testRun.Name,
				"test-name": testName,
			},
		},
		InvolvedObject: objectReference,

		// Machine Reason
		Reason: Reason,
		// User readable reason
		Message: Message,

		// FIXME: populate with real values
		Source: v1.EventSource{
			Component: "test-controller",
			Host:      hostname,
		},
		FirstTimestamp: now,
		LastTimestamp:  now,
		Count:          1,
		Type:           "Normal",
	}

	_, err = ctrl.CoreV1().Events(testRun.Namespace).Create(event)

	if err != nil {
		log.Printf("Error Creating event while starting test %v", err)
		return err
	}

	return nil

}

// CreateTestRunEventStart will create a k8s event when a test pod is created
func CreateTestRunEventStart(ctrl controller.Interface, testRun *v1alpha1.TestRun, test *v1alpha1.TestTemplate) error {
	return CreateTestRunEvent(
		ctrl, testRun, test.Name,
		"TestStart",
		fmt.Sprintf("Starting test %s", test.Name),
	)
}
