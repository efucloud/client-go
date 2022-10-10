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
	"context"

	v1alpha1 "github.com/efucloud/api/rbac/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeKubeUserAPIKeys implements KubeUserAPIKeyInterface
type FakeKubeUserAPIKeys struct {
	Fake *FakeRbacV1alpha1
}

var kubeuserapikeysResource = schema.GroupVersionResource{Group: "rbac.efucloud.com", Version: "v1alpha1", Resource: "kubeuserapikeys"}

var kubeuserapikeysKind = schema.GroupVersionKind{Group: "rbac.efucloud.com", Version: "v1alpha1", Kind: "KubeUserAPIKey"}

// Get takes name of the kubeUserAPIKey, and returns the corresponding kubeUserAPIKey object, and an error if there is any.
func (c *FakeKubeUserAPIKeys) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.KubeUserAPIKey, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(kubeuserapikeysResource, name), &v1alpha1.KubeUserAPIKey{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KubeUserAPIKey), err
}

// List takes label and field selectors, and returns the list of KubeUserAPIKeys that match those selectors.
func (c *FakeKubeUserAPIKeys) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.KubeUserAPIKeyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(kubeuserapikeysResource, kubeuserapikeysKind, opts), &v1alpha1.KubeUserAPIKeyList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.KubeUserAPIKeyList{ListMeta: obj.(*v1alpha1.KubeUserAPIKeyList).ListMeta}
	for _, item := range obj.(*v1alpha1.KubeUserAPIKeyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested kubeUserAPIKeys.
func (c *FakeKubeUserAPIKeys) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(kubeuserapikeysResource, opts))
}

// Create takes the representation of a kubeUserAPIKey and creates it.  Returns the server's representation of the kubeUserAPIKey, and an error, if there is any.
func (c *FakeKubeUserAPIKeys) Create(ctx context.Context, kubeUserAPIKey *v1alpha1.KubeUserAPIKey, opts v1.CreateOptions) (result *v1alpha1.KubeUserAPIKey, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(kubeuserapikeysResource, kubeUserAPIKey), &v1alpha1.KubeUserAPIKey{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KubeUserAPIKey), err
}

// Update takes the representation of a kubeUserAPIKey and updates it. Returns the server's representation of the kubeUserAPIKey, and an error, if there is any.
func (c *FakeKubeUserAPIKeys) Update(ctx context.Context, kubeUserAPIKey *v1alpha1.KubeUserAPIKey, opts v1.UpdateOptions) (result *v1alpha1.KubeUserAPIKey, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(kubeuserapikeysResource, kubeUserAPIKey), &v1alpha1.KubeUserAPIKey{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KubeUserAPIKey), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeKubeUserAPIKeys) UpdateStatus(ctx context.Context, kubeUserAPIKey *v1alpha1.KubeUserAPIKey, opts v1.UpdateOptions) (*v1alpha1.KubeUserAPIKey, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(kubeuserapikeysResource, "status", kubeUserAPIKey), &v1alpha1.KubeUserAPIKey{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KubeUserAPIKey), err
}

// Delete takes name of the kubeUserAPIKey and deletes it. Returns an error if one occurs.
func (c *FakeKubeUserAPIKeys) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(kubeuserapikeysResource, name, opts), &v1alpha1.KubeUserAPIKey{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeKubeUserAPIKeys) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(kubeuserapikeysResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.KubeUserAPIKeyList{})
	return err
}

// Patch applies the patch and returns the patched kubeUserAPIKey.
func (c *FakeKubeUserAPIKeys) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.KubeUserAPIKey, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(kubeuserapikeysResource, name, pt, data, subresources...), &v1alpha1.KubeUserAPIKey{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KubeUserAPIKey), err
}
