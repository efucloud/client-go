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

	v1alpha1 "github.com/efucloud/api/monitor/v1alpha1"
	scheme "github.com/efucloud/client-go/monitor/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ProjectArchitecturesGetter has a method to return a ProjectArchitectureInterface.
// A group's client should implement this interface.
type ProjectArchitecturesGetter interface {
	ProjectArchitectures(namespace string) ProjectArchitectureInterface
}

// ProjectArchitectureInterface has methods to work with ProjectArchitecture resources.
type ProjectArchitectureInterface interface {
	Create(ctx context.Context, projectArchitecture *v1alpha1.ProjectArchitecture, opts v1.CreateOptions) (*v1alpha1.ProjectArchitecture, error)
	Update(ctx context.Context, projectArchitecture *v1alpha1.ProjectArchitecture, opts v1.UpdateOptions) (*v1alpha1.ProjectArchitecture, error)
	UpdateStatus(ctx context.Context, projectArchitecture *v1alpha1.ProjectArchitecture, opts v1.UpdateOptions) (*v1alpha1.ProjectArchitecture, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.ProjectArchitecture, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.ProjectArchitectureList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ProjectArchitecture, err error)
	ProjectArchitectureExpansion
}

// projectArchitectures implements ProjectArchitectureInterface
type projectArchitectures struct {
	client rest.Interface
	ns     string
}

// newProjectArchitectures returns a ProjectArchitectures
func newProjectArchitectures(c *MonitorV1alpha1Client, namespace string) *projectArchitectures {
	return &projectArchitectures{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the projectArchitecture, and returns the corresponding projectArchitecture object, and an error if there is any.
func (c *projectArchitectures) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ProjectArchitecture, err error) {
	result = &v1alpha1.ProjectArchitecture{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("projectarchitectures").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ProjectArchitectures that match those selectors.
func (c *projectArchitectures) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ProjectArchitectureList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ProjectArchitectureList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("projectarchitectures").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested projectArchitectures.
func (c *projectArchitectures) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("projectarchitectures").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a projectArchitecture and creates it.  Returns the server's representation of the projectArchitecture, and an error, if there is any.
func (c *projectArchitectures) Create(ctx context.Context, projectArchitecture *v1alpha1.ProjectArchitecture, opts v1.CreateOptions) (result *v1alpha1.ProjectArchitecture, err error) {
	result = &v1alpha1.ProjectArchitecture{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("projectarchitectures").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(projectArchitecture).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a projectArchitecture and updates it. Returns the server's representation of the projectArchitecture, and an error, if there is any.
func (c *projectArchitectures) Update(ctx context.Context, projectArchitecture *v1alpha1.ProjectArchitecture, opts v1.UpdateOptions) (result *v1alpha1.ProjectArchitecture, err error) {
	result = &v1alpha1.ProjectArchitecture{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("projectarchitectures").
		Name(projectArchitecture.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(projectArchitecture).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *projectArchitectures) UpdateStatus(ctx context.Context, projectArchitecture *v1alpha1.ProjectArchitecture, opts v1.UpdateOptions) (result *v1alpha1.ProjectArchitecture, err error) {
	result = &v1alpha1.ProjectArchitecture{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("projectarchitectures").
		Name(projectArchitecture.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(projectArchitecture).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the projectArchitecture and deletes it. Returns an error if one occurs.
func (c *projectArchitectures) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("projectarchitectures").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *projectArchitectures) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("projectarchitectures").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched projectArchitecture.
func (c *projectArchitectures) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ProjectArchitecture, err error) {
	result = &v1alpha1.ProjectArchitecture{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("projectarchitectures").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
