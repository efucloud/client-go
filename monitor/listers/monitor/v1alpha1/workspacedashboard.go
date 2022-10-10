/*
Copyright 2022 The efucloud.com Authors.

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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/efucloud/api/monitor/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// WorkspaceDashboardLister helps list WorkspaceDashboards.
// All objects returned here must be treated as read-only.
type WorkspaceDashboardLister interface {
	// List lists all WorkspaceDashboards in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.WorkspaceDashboard, err error)
	// Get retrieves the WorkspaceDashboard from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.WorkspaceDashboard, error)
	WorkspaceDashboardListerExpansion
}

// workspaceDashboardLister implements the WorkspaceDashboardLister interface.
type workspaceDashboardLister struct {
	indexer cache.Indexer
}

// NewWorkspaceDashboardLister returns a new WorkspaceDashboardLister.
func NewWorkspaceDashboardLister(indexer cache.Indexer) WorkspaceDashboardLister {
	return &workspaceDashboardLister{indexer: indexer}
}

// List lists all WorkspaceDashboards in the indexer.
func (s *workspaceDashboardLister) List(selector labels.Selector) (ret []*v1alpha1.WorkspaceDashboard, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.WorkspaceDashboard))
	})
	return ret, err
}

// Get retrieves the WorkspaceDashboard from the index for a given name.
func (s *workspaceDashboardLister) Get(name string) (*v1alpha1.WorkspaceDashboard, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("workspacedashboard"), name)
	}
	return obj.(*v1alpha1.WorkspaceDashboard), nil
}
