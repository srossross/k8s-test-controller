package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var (
	// TestRunComplete will set the status to complete
	TestRunComplete = "Complete"
	TestRunRunning  = "Running"
)

// +genclient=true
// +genclient=nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +resource:path=testruns

type TestRun struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   TestRunSpec   `json:"spec,omitempty"`
	Status TestRunStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type TestRunList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []TestRun `json:"items"`
}

type TestRunSpec struct {
	// Label selector for pods. Existing ReplicaSets whose pods are
	// selected by this will be the ones affected by this deployment.
	// +optional
	Selector *metav1.LabelSelector `json:"selector,omitempty"`

	// The Maximum number of pods to run symultaniously
	MaxJobs int `json:"max-jobs"`

	// The Maximum number of failures before stoping the test run
	// and mark it as a failure
	MaxFail int `json:"maxfail"`
}

// TestRunRecord is a refrence to a pod run
type TestRunRecord struct {

	// the name of the test to run
	TestName string `json:"testname"`
	// The pod that this run
	PodRef *corev1.ObjectReference `json:"podref"`
	// When the pod was started
	StartTime *metav1.Time `json:"starttime"`
	// When the pod was started
	EndTime *metav1.Time `json:"endtime"`

	// When the pod was started
	Result string `json:"result"`
}

type TestRunStatus struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Success bool   `json:"success"`

	Records []TestRunRecord `json:"records"`
}

// +genclient=true
// +genclient=nonNamespaced
// +genclient=noStatus
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +resource:path=testtemplates

type TestTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec TestTemplateSpec `json:"spec,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type TestTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []TestTemplate `json:"items"`
}

type TestTemplateSpec struct {

	// Description of what the test is about
	Description string `json:"description"`
	// Test run weight. the pods will be run in sorted order of (Weight, Name)
	Weight   int                    `json:"weight"`
	Template corev1.PodTemplateSpec `json:"template"`
}
