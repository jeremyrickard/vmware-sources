/*
Copyright 2020 The Knative Authors

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
	"time"

	v1alpha1 "github.com/mattmoor/vmware-sources/pkg/apis/sources/v1alpha1"
	scheme "github.com/mattmoor/vmware-sources/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// VSphereSourcesGetter has a method to return a VSphereSourceInterface.
// A group's client should implement this interface.
type VSphereSourcesGetter interface {
	VSphereSources(namespace string) VSphereSourceInterface
}

// VSphereSourceInterface has methods to work with VSphereSource resources.
type VSphereSourceInterface interface {
	Create(*v1alpha1.VSphereSource) (*v1alpha1.VSphereSource, error)
	Update(*v1alpha1.VSphereSource) (*v1alpha1.VSphereSource, error)
	UpdateStatus(*v1alpha1.VSphereSource) (*v1alpha1.VSphereSource, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.VSphereSource, error)
	List(opts v1.ListOptions) (*v1alpha1.VSphereSourceList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.VSphereSource, err error)
	VSphereSourceExpansion
}

// vSphereSources implements VSphereSourceInterface
type vSphereSources struct {
	client rest.Interface
	ns     string
}

// newVSphereSources returns a VSphereSources
func newVSphereSources(c *SamplesV1alpha1Client, namespace string) *vSphereSources {
	return &vSphereSources{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the vSphereSource, and returns the corresponding vSphereSource object, and an error if there is any.
func (c *vSphereSources) Get(name string, options v1.GetOptions) (result *v1alpha1.VSphereSource, err error) {
	result = &v1alpha1.VSphereSource{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("vspheresources").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of VSphereSources that match those selectors.
func (c *vSphereSources) List(opts v1.ListOptions) (result *v1alpha1.VSphereSourceList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.VSphereSourceList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("vspheresources").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested vSphereSources.
func (c *vSphereSources) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("vspheresources").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a vSphereSource and creates it.  Returns the server's representation of the vSphereSource, and an error, if there is any.
func (c *vSphereSources) Create(vSphereSource *v1alpha1.VSphereSource) (result *v1alpha1.VSphereSource, err error) {
	result = &v1alpha1.VSphereSource{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("vspheresources").
		Body(vSphereSource).
		Do().
		Into(result)
	return
}

// Update takes the representation of a vSphereSource and updates it. Returns the server's representation of the vSphereSource, and an error, if there is any.
func (c *vSphereSources) Update(vSphereSource *v1alpha1.VSphereSource) (result *v1alpha1.VSphereSource, err error) {
	result = &v1alpha1.VSphereSource{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("vspheresources").
		Name(vSphereSource.Name).
		Body(vSphereSource).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *vSphereSources) UpdateStatus(vSphereSource *v1alpha1.VSphereSource) (result *v1alpha1.VSphereSource, err error) {
	result = &v1alpha1.VSphereSource{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("vspheresources").
		Name(vSphereSource.Name).
		SubResource("status").
		Body(vSphereSource).
		Do().
		Into(result)
	return
}

// Delete takes name of the vSphereSource and deletes it. Returns an error if one occurs.
func (c *vSphereSources) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("vspheresources").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *vSphereSources) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("vspheresources").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched vSphereSource.
func (c *vSphereSources) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.VSphereSource, err error) {
	result = &v1alpha1.VSphereSource{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("vspheresources").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
