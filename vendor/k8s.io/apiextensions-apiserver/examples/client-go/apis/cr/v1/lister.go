/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// This file was automatically generated by lister-gen

package v1

import (
	v1 "k8s.io/api/core/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	// api "k8s.io/client-go/pkg/api"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// E2ETestRunnerLister helps list E2ETestRunners.
type E2ETestRunnerLister interface {
	// List lists all E2ETestRunners in the indexer.
	List(selector labels.Selector) (ret []*E2ETestRunner, err error)
	// E2ETestRunners returns an object that can list and get E2ETestRunners.
	Get(name string) (*E2ETestRunner, error)
	E2ETestRunners(namespace string) E2ETestRunnerNamespaceLister
}

// e2eTestLister implements the E2ETestRunnerLister interface.
type e2eTestLister struct {
	indexer cache.Indexer
}

// NewE2ETestRunnerLister returns a new E2ETestRunnerLister.
func NewE2ETestRunnerLister(indexer cache.Indexer) E2ETestRunnerLister {
	return &e2eTestLister{indexer: indexer}
}

// List lists all E2ETestRunners in the indexer.
func (s *e2eTestLister) List(selector labels.Selector) (ret []*E2ETestRunner, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*E2ETestRunner))
	})
	return ret, err
}

// Get retrieves the Node from the index for a given name.
func (s *e2eTestLister) Get(name string) (*E2ETestRunner, error) {
	key := &E2ETestRunner{ObjectMeta: meta_v1.ObjectMeta{Name: name}}
	obj, exists, err := s.indexer.Get(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("testrunner"), name)
	}
	return obj.(*E2ETestRunner), nil
}


// E2ETestRunners returns an object that can list and get E2ETestRunners.
func (s *e2eTestLister) E2ETestRunners(namespace string) E2ETestRunnerNamespaceLister {
	return e2eTestNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// E2ETestRunnerNamespaceLister helps list and get E2ETestRunners.
type E2ETestRunnerNamespaceLister interface {
	// List lists all E2ETestRunners in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*E2ETestRunner, err error)
	// Get retrieves the E2ETestRunner from the indexer for a given namespace and name.
	Get(name string) (*E2ETestRunner, error)
}

// e2eTestNamespaceLister implements the E2ETestRunnerNamespaceLister
// interface.
type e2eTestNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all E2ETestRunners in the indexer for a given namespace.
func (s e2eTestNamespaceLister) List(selector labels.Selector) (ret []*E2ETestRunner, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*E2ETestRunner))
	})
	return ret, err
}

// Get retrieves the E2ETestRunner from the indexer for a given namespace and name.
func (s e2eTestNamespaceLister) Get(name string) (*E2ETestRunner, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("pod"), name)
	}
	return obj.(*E2ETestRunner), nil
}
