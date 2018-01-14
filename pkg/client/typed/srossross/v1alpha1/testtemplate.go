/*

MIT License

Copyright (c) 2017 Sean Ross-Ross

See License in the root of this repo.

*/
package v1alpha1

import (
	v1alpha1 "github.com/srossross/k8s-test-controller/pkg/apis/tester/v1alpha1"
	scheme "github.com/srossross/k8s-test-controller/pkg/client/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// TestTemplatesGetter has a method to return a TestTemplateInterface.
// A group's client should implement this interface.
type TestTemplatesGetter interface {
	TestTemplates(namespace string) TestTemplateInterface
}

// TestTemplateInterface has methods to work with TestTemplate resources.
type TestTemplateInterface interface {
	Create(*v1alpha1.TestTemplate) (*v1alpha1.TestTemplate, error)
	Update(*v1alpha1.TestTemplate) (*v1alpha1.TestTemplate, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.TestTemplate, error)
	List(opts v1.ListOptions) (*v1alpha1.TestTemplateList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.TestTemplate, err error)
	TestTemplateExpansion
}

// testTemplates implements TestTemplateInterface
type testTemplates struct {
	client rest.Interface
	ns     string
}

// newTestTemplates returns a TestTemplates
func newTestTemplates(c *SrossrossV1alpha1Client, namespace string) *testTemplates {
	return &testTemplates{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Create takes the representation of a testTemplate and creates it.  Returns the server's representation of the testTemplate, and an error, if there is any.
func (c *testTemplates) Create(testTemplate *v1alpha1.TestTemplate) (result *v1alpha1.TestTemplate, err error) {
	result = &v1alpha1.TestTemplate{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("testtemplates").
		Body(testTemplate).
		Do().
		Into(result)
	return
}

// Update takes the representation of a testTemplate and updates it. Returns the server's representation of the testTemplate, and an error, if there is any.
func (c *testTemplates) Update(testTemplate *v1alpha1.TestTemplate) (result *v1alpha1.TestTemplate, err error) {
	result = &v1alpha1.TestTemplate{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("testtemplates").
		Name(testTemplate.Name).
		Body(testTemplate).
		Do().
		Into(result)
	return
}

// Delete takes name of the testTemplate and deletes it. Returns an error if one occurs.
func (c *testTemplates) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("testtemplates").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *testTemplates) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("testtemplates").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Get takes name of the testTemplate, and returns the corresponding testTemplate object, and an error if there is any.
func (c *testTemplates) Get(name string, options v1.GetOptions) (result *v1alpha1.TestTemplate, err error) {
	result = &v1alpha1.TestTemplate{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("testtemplates").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of TestTemplates that match those selectors.
func (c *testTemplates) List(opts v1.ListOptions) (result *v1alpha1.TestTemplateList, err error) {
	result = &v1alpha1.TestTemplateList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("testtemplates").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested testTemplates.
func (c *testTemplates) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("testtemplates").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Patch applies the patch and returns the patched testTemplate.
func (c *testTemplates) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.TestTemplate, err error) {
	result = &v1alpha1.TestTemplate{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("testtemplates").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
