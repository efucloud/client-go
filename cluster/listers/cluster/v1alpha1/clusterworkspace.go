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

// ClusterWorkspaceLister helps list ClusterWorkspaces.
// All objects returned here must be treated as read-only.
type ClusterWorkspaceLister interface {
	// List lists all ClusterWorkspaces in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ClusterWorkspace, err error)
	// Get retrieves the ClusterWorkspace from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ClusterWorkspace, error)
	ClusterWorkspaceListerExpansion
}

// clusterWorkspaceLister implements the ClusterWorkspaceLister interface.
type clusterWorkspaceLister struct {
	indexer cache.Indexer
}

// NewClusterWorkspaceLister returns a new ClusterWorkspaceLister.
func NewClusterWorkspaceLister(indexer cache.Indexer) ClusterWorkspaceLister {
	return &clusterWorkspaceLister{indexer: indexer}
}

// List lists all ClusterWorkspaces in the indexer.
func (s *clusterWorkspaceLister) List(selector labels.Selector) (ret []*v1alpha1.ClusterWorkspace, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ClusterWorkspace))
	})
	return ret, err
}

// Get retrieves the ClusterWorkspace from the index for a given name.
func (s *clusterWorkspaceLister) Get(name string) (*v1alpha1.ClusterWorkspace, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("clusterworkspace"), name)
	}
	return obj.(*v1alpha1.ClusterWorkspace), nil
}
