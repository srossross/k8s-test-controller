/*

MIT License

Copyright (c) 2017 Sean Ross-Ross

See License in the root of this repo.

*/
package fake

import (
	v1alpha1 "github.com/srossross/k8s-test-controller/pkg/apis/tester/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeTestTemplates implements TestTemplateInterface
type FakeTestTemplates struct {
	Fake *FakeSrossrossV1alpha1
	ns   string
}

var testtemplatesResource = schema.GroupVersionResource{Group: "srossross.github.io", Version: "v1alpha1", Resource: "testtemplates"}

var testtemplatesKind = schema.GroupVersionKind{Group: "srossross.github.io", Version: "v1alpha1", Kind: "TestTemplate"}

func (c *FakeTestTemplates) Create(testTemplate *v1alpha1.TestTemplate) (result *v1alpha1.TestTemplate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(testtemplatesResource, c.ns, testTemplate), &v1alpha1.TestTemplate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TestTemplate), err
}

func (c *FakeTestTemplates) Update(testTemplate *v1alpha1.TestTemplate) (result *v1alpha1.TestTemplate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(testtemplatesResource, c.ns, testTemplate), &v1alpha1.TestTemplate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TestTemplate), err
}

func (c *FakeTestTemplates) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(testtemplatesResource, c.ns, name), &v1alpha1.TestTemplate{})

	return err
}

func (c *FakeTestTemplates) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(testtemplatesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.TestTemplateList{})
	return err
}

func (c *FakeTestTemplates) Get(name string, options v1.GetOptions) (result *v1alpha1.TestTemplate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(testtemplatesResource, c.ns, name), &v1alpha1.TestTemplate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TestTemplate), err
}

func (c *FakeTestTemplates) List(opts v1.ListOptions) (result *v1alpha1.TestTemplateList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(testtemplatesResource, testtemplatesKind, c.ns, opts), &v1alpha1.TestTemplateList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.TestTemplateList{}
	for _, item := range obj.(*v1alpha1.TestTemplateList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested testTemplates.
func (c *FakeTestTemplates) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(testtemplatesResource, c.ns, opts))

}

// Patch applies the patch and returns the patched testTemplate.
func (c *FakeTestTemplates) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.TestTemplate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(testtemplatesResource, c.ns, name, data, subresources...), &v1alpha1.TestTemplate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TestTemplate), err
}
