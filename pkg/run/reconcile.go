package run

import (
	"fmt"
	"log"
	"time"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/util/wait"

	v1alpha1 "github.com/srossross/k8s-test-controller/pkg/apis/tester/v1alpha1"
	controller "github.com/srossross/k8s-test-controller/pkg/controller"
)

func getTestsForTestRun(ctrl controller.Interface, testRun *v1alpha1.TestRun) ([]*v1alpha1.TestTemplate, error) {
	selector, err := metav1.LabelSelectorAsSelector(testRun.Spec.Selector)
	if selector.String() == "" {
		selector = labels.Everything()
	}

	if err != nil {
		return nil, err
	}

	return ctrl.TestTemplateLister().TestTemplates(testRun.Namespace).List(selector)
}

func initializeStatus(ctrl controller.Interface, testRun *v1alpha1.TestRun) error {
	tests, err := getTestsForTestRun(ctrl, testRun)

	if err != nil {
		return fmt.Errorf("error getting list of tests: %s", err.Error())
	}

	log.Printf("testRun.Status.Status is '%v'", testRun.Status.Status)
	testRunCopy := testRun.DeepCopy()

	testRecords := []v1alpha1.TestRunRecord{}
	for _, test := range tests {
		testRecords = append(testRecords, v1alpha1.TestRunRecord{
			TestName:  test.Name,
			PodRef:    nil,
			StartTime: nil,
			EndTime:   nil,
			Result:    "N/A",
		})
	}

	testRunCopy.Status.Status = v1alpha1.TestRunRunning
	testRunCopy.Status.Records = testRecords

	log.Printf("Initialize '%v/%v'", testRun.Namespace, testRun.Name)
	if _, err = ctrl.SrossrossV1alpha1().TestRuns(testRun.Namespace).Update(testRunCopy); err != nil {
		return err
	}

	return wait.Poll(500*time.Millisecond, 60*time.Second, func() (bool, error) {
		t, err := ctrl.TestRunLister().TestRuns(testRun.Namespace).Get(testRun.Name)
		if err != nil {
			return true, err
		}
		if t.Status.Records != nil {
			return true, nil
		}
		return false, nil
	})

}

type runStats struct {
	CompletedCount int
	FailCount      int
}

func testFinished(ctrl controller.Interface, testRun *v1alpha1.TestRun, i int, Result string) error {
	testRun = testRun.DeepCopy()

	testRun.Status.Records[i].EndTime = &metav1.Time{Time: time.Now()}
	if testRun.Status.Records[i].StartTime == nil {
		testRun.Status.Records[i].StartTime = testRun.Status.Records[i].EndTime
	}
	testRun.Status.Records[i].Result = Result

	if _, err := ctrl.SrossrossV1alpha1().TestRuns(testRun.Namespace).Update(testRun); err != nil {
		return err
	}

	return wait.Poll(500*time.Millisecond, 60*time.Second, func() (bool, error) {
		t, err := ctrl.TestRunLister().TestRuns(testRun.Namespace).Get(testRun.Name)
		if err != nil {
			return true, err
		}
		if t.Status.Records[i].EndTime != nil {
			return true, nil
		}
		return false, nil
	})

}

func testStarted(ctrl controller.Interface, testRun *v1alpha1.TestRun, i int, pod *corev1.Pod) error {
	testRun = testRun.DeepCopy()

	testRun.Status.Records[i].StartTime = &metav1.Time{Time: time.Now()}
	testRun.Status.Records[i].PodRef = &corev1.ObjectReference{
		Kind:            "Pod",
		Namespace:       pod.Namespace,
		Name:            pod.Name,
		APIVersion:      pod.APIVersion,
		ResourceVersion: pod.ResourceVersion,
	}

	if _, err := ctrl.SrossrossV1alpha1().TestRuns(testRun.Namespace).Update(testRun); err != nil {
		return err
	}

	return wait.Poll(500*time.Millisecond, 60*time.Second, func() (bool, error) {
		t, err := ctrl.TestRunLister().TestRuns(testRun.Namespace).Get(testRun.Name)
		if err != nil {
			return true, err
		}
		if t.Status.Records[i].StartTime != nil {
			return true, nil
		}
		return false, nil
	})
}

func runNextNTests(ctrl controller.Interface, testRun *v1alpha1.TestRun, tests []*v1alpha1.TestTemplate, JobsSlots int) error {
	podMap, err := createPodMap(ctrl, testRun)
	if err != nil {
		return err
	}

	testMap := createTestMap(tests)

	for i, record := range testRun.Status.Records {
		if JobsSlots <= 0 {
			log.Printf("  | No more jobs allowed (maxjobs: %v). Will wait for next event", getJobSlots(testRun))
			return nil
		}

		if record.EndTime != nil {
			// The Test has ended
			continue
		}
		if record.StartTime == nil {
			// The Test not started
			test, ok := testMap[record.TestName]
			if !ok {
				err = testFinished(ctrl, testRun, i, "TestRemoved")
				return err
			}

			var pod *corev1.Pod
			pod, err = CreateTestPod(ctrl, testRun, test)

			if err != nil {
				return testFinished(ctrl, testRun, i, "PodStartError")
			}
			return testStarted(ctrl, testRun, i, pod)
		}

		// record.StartTime is non nil

		if pod, ok := podMap[record.TestName]; ok {
			log.Printf("  |         - Pod '%v' exists - Status: %v", pod.Name, pod.Status.Phase)
			switch pod.Status.Phase {
			case corev1.PodSucceeded:
				return testFinished(ctrl, testRun, i, string(pod.Status.Phase))
			case corev1.PodFailed:
				return testFinished(ctrl, testRun, i, string(pod.Status.Phase))
			case corev1.PodUnknown:
				return testFinished(ctrl, testRun, i, string(pod.Status.Phase))
			// These are running and taking up a job slot!
			case corev1.PodPending:
				JobsSlots--
				continue
			case corev1.PodRunning:
				JobsSlots--
				continue
			}
		}
	}
	// Only get here if all jobs finish
	return nil
}

func createPodMap(ctrl controller.Interface, testRun *v1alpha1.TestRun) (map[string]*corev1.Pod, error) {
	pods, err := ctrl.ListPods(testRun.Namespace, labels.Everything())
	podMap := make(map[string]*corev1.Pod)
	if err != nil {
		return podMap, fmt.Errorf("Error getting list of pods: %s", err.Error())
	}

	pods = controller.TestRunFilter(pods, testRun.Name)

	log.Printf("  | Total Pod Count: %v", len(pods))

	for _, pod := range pods {
		// log.Printf("  |  Pod: %v", pod.Name)
		podMap[pod.Labels["test-name"]] = pod
	}

	return podMap, nil

}

func createTestMap(tests []*v1alpha1.TestTemplate) map[string]*v1alpha1.TestTemplate {
	testMap := make(map[string]*v1alpha1.TestTemplate)

	for _, test := range tests {
		testMap[test.Name] = test
	}
	return testMap
}

func getJobSlots(testRun *v1alpha1.TestRun) int {
	// FIXME: should be a default in the schema ...
	if testRun.Spec.MaxJobs > 0 {
		return testRun.Spec.MaxJobs
	}
	return 1
}

func testRunComplete(ctrl controller.Interface, testRun *v1alpha1.TestRun, stats runStats) error {
	Message := fmt.Sprintf("Ran %v tests, %v failures", stats.CompletedCount, stats.FailCount)
	var Reason string
	testRun = testRun.DeepCopy()

	testRun.Status.Status = v1alpha1.TestRunComplete
	testRun.Status.Success = stats.FailCount == 0
	testRun.Status.Message = Message

	log.Printf("Saving '%v/%v'", testRun.Namespace, testRun.Name)
	if _, err := ctrl.SrossrossV1alpha1().TestRuns(testRun.Namespace).Update(testRun); err != nil {
		return err
	}
	log.Printf("We are done here %v tests completed", stats.CompletedCount)

	switch stats.FailCount == 0 {
	case true:
		Reason = "TestRunSuccess"
	case false:
		Reason = "TestRunFail"
	}
	return CreateTestRunEvent(ctrl, testRun, "", Reason, Message)
}

// UpdateTestRun will Reconcile a single test run
func (r *runner) UpdateTestRun(ctrl controller.Interface, testRun *v1alpha1.TestRun) error {

	if testRun.Status.Status == v1alpha1.TestRunComplete {
		log.Printf("  | '%v/%v' is already Complete - Skipping", testRun.Namespace, testRun.Name)
		return nil
	}

	if testRun.Status.Status == "" {
		err := initializeStatus(ctrl, testRun)
		if err != nil {
			return err
		}
	}

	stats := runStats{0, 0}
	for _, record := range testRun.Status.Records {
		if record.EndTime != nil {
			stats.CompletedCount++
			if record.Result != string(corev1.PodSucceeded) {
				stats.FailCount++
			}
		}
	}
	if stats.CompletedCount == len(testRun.Status.Records) {
		return testRunComplete(ctrl, testRun, stats)
	}

	log.Printf("Running '%v/%v'", testRun.Namespace, testRun.Name)

	log.Printf("  | %v/%v", testRun.Namespace, testRun.Name)

	tests, err := getTestsForTestRun(ctrl, testRun)

	if err != nil {
		return fmt.Errorf("error getting list of tests: %s", err.Error())
	}
	log.Printf("  | Test Count: %v", len(tests))

	JobsSlots := getJobSlots(testRun)

	return runNextNTests(ctrl, testRun, tests, JobsSlots)
}
