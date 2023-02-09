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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/efucloud/client-go/cluster/clientset/versioned/typed/cluster/v1alpha1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeClusterV1alpha1 struct {
	*testing.Fake
}

func (c *FakeClusterV1alpha1) Clusters() v1alpha1.ClusterInterface {
	return &FakeClusters{c}
}

func (c *FakeClusterV1alpha1) ClusterProfiles() v1alpha1.ClusterProfileInterface {
	return &FakeClusterProfiles{c}
}

func (c *FakeClusterV1alpha1) ClusterWorkspaces() v1alpha1.ClusterWorkspaceInterface {
	return &FakeClusterWorkspaces{c}
}

func (c *FakeClusterV1alpha1) DeployTemplates() v1alpha1.DeployTemplateInterface {
	return &FakeDeployTemplates{c}
}

func (c *FakeClusterV1alpha1) Workspaces() v1alpha1.WorkspaceInterface {
	return &FakeWorkspaces{c}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeClusterV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
