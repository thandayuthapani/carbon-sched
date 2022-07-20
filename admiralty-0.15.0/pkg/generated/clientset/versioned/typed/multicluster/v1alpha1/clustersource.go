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

// ClusterSourcesGetter has a method to return a ClusterSourceInterface.
// A group's client should implement this interface.
type ClusterSourcesGetter interface {
	ClusterSources() ClusterSourceInterface
}

// ClusterSourceInterface has methods to work with ClusterSource resources.
type ClusterSourceInterface interface {
	Create(ctx context.Context, clusterSource *v1alpha1.ClusterSource, opts v1.CreateOptions) (*v1alpha1.ClusterSource, error)
	Update(ctx context.Context, clusterSource *v1alpha1.ClusterSource, opts v1.UpdateOptions) (*v1alpha1.ClusterSource, error)
	UpdateStatus(ctx context.Context, clusterSource *v1alpha1.ClusterSource, opts v1.UpdateOptions) (*v1alpha1.ClusterSource, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.ClusterSource, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.ClusterSourceList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ClusterSource, err error)
	ClusterSourceExpansion
}

// clusterSources implements ClusterSourceInterface
type clusterSources struct {
	client rest.Interface
}

// newClusterSources returns a ClusterSources
func newClusterSources(c *MulticlusterV1alpha1Client) *clusterSources {
	return &clusterSources{
		client: c.RESTClient(),
	}
}

// Get takes name of the clusterSource, and returns the corresponding clusterSource object, and an error if there is any.
func (c *clusterSources) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ClusterSource, err error) {
	result = &v1alpha1.ClusterSource{}
	err = c.client.Get().
		Resource("clustersources").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ClusterSources that match those selectors.
func (c *clusterSources) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ClusterSourceList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ClusterSourceList{}
	err = c.client.Get().
		Resource("clustersources").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested clusterSources.
func (c *clusterSources) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("clustersources").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a clusterSource and creates it.  Returns the server's representation of the clusterSource, and an error, if there is any.
func (c *clusterSources) Create(ctx context.Context, clusterSource *v1alpha1.ClusterSource, opts v1.CreateOptions) (result *v1alpha1.ClusterSource, err error) {
	result = &v1alpha1.ClusterSource{}
	err = c.client.Post().
		Resource("clustersources").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(clusterSource).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a clusterSource and updates it. Returns the server's representation of the clusterSource, and an error, if there is any.
func (c *clusterSources) Update(ctx context.Context, clusterSource *v1alpha1.ClusterSource, opts v1.UpdateOptions) (result *v1alpha1.ClusterSource, err error) {
	result = &v1alpha1.ClusterSource{}
	err = c.client.Put().
		Resource("clustersources").
		Name(clusterSource.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(clusterSource).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *clusterSources) UpdateStatus(ctx context.Context, clusterSource *v1alpha1.ClusterSource, opts v1.UpdateOptions) (result *v1alpha1.ClusterSource, err error) {
	result = &v1alpha1.ClusterSource{}
	err = c.client.Put().
		Resource("clustersources").
		Name(clusterSource.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(clusterSource).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the clusterSource and deletes it. Returns an error if one occurs.
func (c *clusterSources) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("clustersources").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *clusterSources) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("clustersources").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched clusterSource.
func (c *clusterSources) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ClusterSource, err error) {
	result = &v1alpha1.ClusterSource{}
	err = c.client.Patch(pt).
		Resource("clustersources").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
