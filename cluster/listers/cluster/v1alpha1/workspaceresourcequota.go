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
	v1alpha1 "github.com/efucloud/api/cluster/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// WorkspaceResourceQuotaLister helps list WorkspaceResourceQuotas.
// All objects returned here must be treated as read-only.
type WorkspaceResourceQuotaLister interface {
	// List lists all WorkspaceResourceQuotas in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.WorkspaceResourceQuota, err error)
	// Get retrieves the WorkspaceResourceQuota from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.WorkspaceResourceQuota, error)
	WorkspaceResourceQuotaListerExpansion
}

// workspaceResourceQuotaLister implements the WorkspaceResourceQuotaLister interface.
type workspaceResourceQuotaLister struct {
	indexer cache.Indexer
}

// NewWorkspaceResourceQuotaLister returns a new WorkspaceResourceQuotaLister.
func NewWorkspaceResourceQuotaLister(indexer cache.Indexer) WorkspaceResourceQuotaLister {
	return &workspaceResourceQuotaLister{indexer: indexer}
}

// List lists all WorkspaceResourceQuotas in the indexer.
func (s *workspaceResourceQuotaLister) List(selector labels.Selector) (ret []*v1alpha1.WorkspaceResourceQuota, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.WorkspaceResourceQuota))
	})
	return ret, err
}

// Get retrieves the WorkspaceResourceQuota from the index for a given name.
func (s *workspaceResourceQuotaLister) Get(name string) (*v1alpha1.WorkspaceResourceQuota, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("workspaceresourcequota"), name)
	}
	return obj.(*v1alpha1.WorkspaceResourceQuota), nil
}
