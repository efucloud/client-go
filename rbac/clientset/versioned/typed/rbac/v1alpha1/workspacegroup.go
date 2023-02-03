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

package v1alpha1

import (
	"context"
	"time"

	v1alpha1 "github.com/efucloud/api/rbac/v1alpha1"
	scheme "github.com/efucloud/client-go/rbac/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// WorkspaceGroupsGetter has a method to return a WorkspaceGroupInterface.
// A group's client should implement this interface.
type WorkspaceGroupsGetter interface {
	WorkspaceGroups() WorkspaceGroupInterface
}

// WorkspaceGroupInterface has methods to work with WorkspaceGroup resources.
type WorkspaceGroupInterface interface {
	Create(ctx context.Context, workspaceGroup *v1alpha1.WorkspaceGroup, opts v1.CreateOptions) (*v1alpha1.WorkspaceGroup, error)
	Update(ctx context.Context, workspaceGroup *v1alpha1.WorkspaceGroup, opts v1.UpdateOptions) (*v1alpha1.WorkspaceGroup, error)
	UpdateStatus(ctx context.Context, workspaceGroup *v1alpha1.WorkspaceGroup, opts v1.UpdateOptions) (*v1alpha1.WorkspaceGroup, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.WorkspaceGroup, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.WorkspaceGroupList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.WorkspaceGroup, err error)
	WorkspaceGroupExpansion
}

// workspaceGroups implements WorkspaceGroupInterface
type workspaceGroups struct {
	client rest.Interface
}

// newWorkspaceGroups returns a WorkspaceGroups
func newWorkspaceGroups(c *RbacV1alpha1Client) *workspaceGroups {
	return &workspaceGroups{
		client: c.RESTClient(),
	}
}

// Get takes name of the workspaceGroup, and returns the corresponding workspaceGroup object, and an error if there is any.
func (c *workspaceGroups) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.WorkspaceGroup, err error) {
	result = &v1alpha1.WorkspaceGroup{}
	err = c.client.Get().
		Resource("workspacegroups").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of WorkspaceGroups that match those selectors.
func (c *workspaceGroups) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.WorkspaceGroupList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.WorkspaceGroupList{}
	err = c.client.Get().
		Resource("workspacegroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested workspaceGroups.
func (c *workspaceGroups) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("workspacegroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a workspaceGroup and creates it.  Returns the server's representation of the workspaceGroup, and an error, if there is any.
func (c *workspaceGroups) Create(ctx context.Context, workspaceGroup *v1alpha1.WorkspaceGroup, opts v1.CreateOptions) (result *v1alpha1.WorkspaceGroup, err error) {
	result = &v1alpha1.WorkspaceGroup{}
	err = c.client.Post().
		Resource("workspacegroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(workspaceGroup).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a workspaceGroup and updates it. Returns the server's representation of the workspaceGroup, and an error, if there is any.
func (c *workspaceGroups) Update(ctx context.Context, workspaceGroup *v1alpha1.WorkspaceGroup, opts v1.UpdateOptions) (result *v1alpha1.WorkspaceGroup, err error) {
	result = &v1alpha1.WorkspaceGroup{}
	err = c.client.Put().
		Resource("workspacegroups").
		Name(workspaceGroup.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(workspaceGroup).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *workspaceGroups) UpdateStatus(ctx context.Context, workspaceGroup *v1alpha1.WorkspaceGroup, opts v1.UpdateOptions) (result *v1alpha1.WorkspaceGroup, err error) {
	result = &v1alpha1.WorkspaceGroup{}
	err = c.client.Put().
		Resource("workspacegroups").
		Name(workspaceGroup.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(workspaceGroup).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the workspaceGroup and deletes it. Returns an error if one occurs.
func (c *workspaceGroups) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("workspacegroups").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *workspaceGroups) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("workspacegroups").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched workspaceGroup.
func (c *workspaceGroups) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.WorkspaceGroup, err error) {
	result = &v1alpha1.WorkspaceGroup{}
	err = c.client.Patch(pt).
		Resource("workspacegroups").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}