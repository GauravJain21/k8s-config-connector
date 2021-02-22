// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// Code generated by main. DO NOT EDIT.

package fake

import (
	"context"

	v1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/compute/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeComputeVPNTunnels implements ComputeVPNTunnelInterface
type FakeComputeVPNTunnels struct {
	Fake *FakeComputeV1beta1
	ns   string
}

var computevpntunnelsResource = schema.GroupVersionResource{Group: "compute.cnrm.cloud.google.com", Version: "v1beta1", Resource: "computevpntunnels"}

var computevpntunnelsKind = schema.GroupVersionKind{Group: "compute.cnrm.cloud.google.com", Version: "v1beta1", Kind: "ComputeVPNTunnel"}

// Get takes name of the computeVPNTunnel, and returns the corresponding computeVPNTunnel object, and an error if there is any.
func (c *FakeComputeVPNTunnels) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.ComputeVPNTunnel, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(computevpntunnelsResource, c.ns, name), &v1beta1.ComputeVPNTunnel{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ComputeVPNTunnel), err
}

// List takes label and field selectors, and returns the list of ComputeVPNTunnels that match those selectors.
func (c *FakeComputeVPNTunnels) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.ComputeVPNTunnelList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(computevpntunnelsResource, computevpntunnelsKind, c.ns, opts), &v1beta1.ComputeVPNTunnelList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.ComputeVPNTunnelList{ListMeta: obj.(*v1beta1.ComputeVPNTunnelList).ListMeta}
	for _, item := range obj.(*v1beta1.ComputeVPNTunnelList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested computeVPNTunnels.
func (c *FakeComputeVPNTunnels) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(computevpntunnelsResource, c.ns, opts))

}

// Create takes the representation of a computeVPNTunnel and creates it.  Returns the server's representation of the computeVPNTunnel, and an error, if there is any.
func (c *FakeComputeVPNTunnels) Create(ctx context.Context, computeVPNTunnel *v1beta1.ComputeVPNTunnel, opts v1.CreateOptions) (result *v1beta1.ComputeVPNTunnel, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(computevpntunnelsResource, c.ns, computeVPNTunnel), &v1beta1.ComputeVPNTunnel{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ComputeVPNTunnel), err
}

// Update takes the representation of a computeVPNTunnel and updates it. Returns the server's representation of the computeVPNTunnel, and an error, if there is any.
func (c *FakeComputeVPNTunnels) Update(ctx context.Context, computeVPNTunnel *v1beta1.ComputeVPNTunnel, opts v1.UpdateOptions) (result *v1beta1.ComputeVPNTunnel, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(computevpntunnelsResource, c.ns, computeVPNTunnel), &v1beta1.ComputeVPNTunnel{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ComputeVPNTunnel), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeComputeVPNTunnels) UpdateStatus(ctx context.Context, computeVPNTunnel *v1beta1.ComputeVPNTunnel, opts v1.UpdateOptions) (*v1beta1.ComputeVPNTunnel, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(computevpntunnelsResource, "status", c.ns, computeVPNTunnel), &v1beta1.ComputeVPNTunnel{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ComputeVPNTunnel), err
}

// Delete takes name of the computeVPNTunnel and deletes it. Returns an error if one occurs.
func (c *FakeComputeVPNTunnels) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(computevpntunnelsResource, c.ns, name), &v1beta1.ComputeVPNTunnel{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeComputeVPNTunnels) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(computevpntunnelsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.ComputeVPNTunnelList{})
	return err
}

// Patch applies the patch and returns the patched computeVPNTunnel.
func (c *FakeComputeVPNTunnels) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.ComputeVPNTunnel, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(computevpntunnelsResource, c.ns, name, pt, data, subresources...), &v1beta1.ComputeVPNTunnel{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ComputeVPNTunnel), err
}
