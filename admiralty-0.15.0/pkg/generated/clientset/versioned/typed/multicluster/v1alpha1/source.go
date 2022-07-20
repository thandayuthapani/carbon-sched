/*
 * Copyright 2020 The Multicluster-Scheduler Authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	"time"

	v1alpha1 "admiralty.io/multicluster-scheduler/pkg/apis/multicluster/v1alpha1"
	scheme "admiralty.io/multicluster-scheduler/pkg/generated/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// SourcesGetter has a method to return a SourceInterface.
// A group's client should implement this interface.
type SourcesGetter interface {
	Sources(namespace string) SourceInterface
}

// SourceInterface has methods to work with Source resources.
type SourceInterface interface {
	Create(ctx context.Context, source *v1alpha1.Source, opts v1.CreateOptions) (*v1alpha1.Source, error)
	Update(ctx context.Context, source *v1alpha1.Source, opts v1.UpdateOptions) (*v1alpha1.Source, error)
	UpdateStatus(ctx context.Context, source *v1alpha1.Source, opts v1.UpdateOptions) (*v1alpha1.Source, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.Source, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.SourceList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Source, err error)
	SourceExpansion
}

// sources implements SourceInterface
type sources struct {
	client rest.Interface
	ns     string
}

// newSources returns a Sources
func newSources(c *MulticlusterV1alpha1Client, namespace string) *sources {
	return &sources{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the source, and returns the corresponding source object, and an error if there is any.
func (c *sources) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Source, err error) {
	result = &v1alpha1.Source{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("sources").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Sources that match those selectors.
func (c *sources) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.SourceList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.SourceList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("sources").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested sources.
func (c *sources) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("sources").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a source and creates it.  Returns the server's representation of the source, and an error, if there is any.
func (c *sources) Create(ctx context.Context, source *v1alpha1.Source, opts v1.CreateOptions) (result *v1alpha1.Source, err error) {
	result = &v1alpha1.Source{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("sources").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(source).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a source and updates it. Returns the server's representation of the source, and an error, if there is any.
func (c *sources) Update(ctx context.Context, source *v1alpha1.Source, opts v1.UpdateOptions) (result *v1alpha1.Source, err error) {
	result = &v1alpha1.Source{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("sources").
		Name(source.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(source).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *sources) UpdateStatus(ctx context.Context, source *v1alpha1.Source, opts v1.UpdateOptions) (result *v1alpha1.Source, err error) {
	result = &v1alpha1.Source{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("sources").
		Name(source.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(source).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the source and deletes it. Returns an error if one occurs.
func (c *sources) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("sources").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *sources) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("sources").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched source.
func (c *sources) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Source, err error) {
	result = &v1alpha1.Source{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("sources").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
