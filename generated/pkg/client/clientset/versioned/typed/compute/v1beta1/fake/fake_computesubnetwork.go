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

// FakeComputeSubnetworks implements ComputeSubnetworkInterface
type FakeComputeSubnetworks struct {
	Fake *FakeComputeV1beta1
	ns   string
}

var computesubnetworksResource = schema.GroupVersionResource{Group: "compute.cnrm.cloud.google.com", Version: "v1beta1", Resource: "computesubnetworks"}

var computesubnetworksKind = schema.GroupVersionKind{Group: "compute.cnrm.cloud.google.com", Version: "v1beta1", Kind: "ComputeSubnetwork"}

// Get takes name of the computeSubnetwork, and returns the corresponding computeSubnetwork object, and an error if there is any.
func (c *FakeComputeSubnetworks) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.ComputeSubnetwork, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(computesubnetworksResource, c.ns, name), &v1beta1.ComputeSubnetwork{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ComputeSubnetwork), err
}

// List takes label and field selectors, and returns the list of ComputeSubnetworks that match those selectors.
func (c *FakeComputeSubnetworks) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.ComputeSubnetworkList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(computesubnetworksResource, computesubnetworksKind, c.ns, opts), &v1beta1.ComputeSubnetworkList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.ComputeSubnetworkList{ListMeta: obj.(*v1beta1.ComputeSubnetworkList).ListMeta}
	for _, item := range obj.(*v1beta1.ComputeSubnetworkList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested computeSubnetworks.
func (c *FakeComputeSubnetworks) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(computesubnetworksResource, c.ns, opts))

}

// Create takes the representation of a computeSubnetwork and creates it.  Returns the server's representation of the computeSubnetwork, and an error, if there is any.
func (c *FakeComputeSubnetworks) Create(ctx context.Context, computeSubnetwork *v1beta1.ComputeSubnetwork, opts v1.CreateOptions) (result *v1beta1.ComputeSubnetwork, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(computesubnetworksResource, c.ns, computeSubnetwork), &v1beta1.ComputeSubnetwork{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ComputeSubnetwork), err
}

// Update takes the representation of a computeSubnetwork and updates it. Returns the server's representation of the computeSubnetwork, and an error, if there is any.
func (c *FakeComputeSubnetworks) Update(ctx context.Context, computeSubnetwork *v1beta1.ComputeSubnetwork, opts v1.UpdateOptions) (result *v1beta1.ComputeSubnetwork, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(computesubnetworksResource, c.ns, computeSubnetwork), &v1beta1.ComputeSubnetwork{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ComputeSubnetwork), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeComputeSubnetworks) UpdateStatus(ctx context.Context, computeSubnetwork *v1beta1.ComputeSubnetwork, opts v1.UpdateOptions) (*v1beta1.ComputeSubnetwork, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(computesubnetworksResource, "status", c.ns, computeSubnetwork), &v1beta1.ComputeSubnetwork{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ComputeSubnetwork), err
}

// Delete takes name of the computeSubnetwork and deletes it. Returns an error if one occurs.
func (c *FakeComputeSubnetworks) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(computesubnetworksResource, c.ns, name), &v1beta1.ComputeSubnetwork{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeComputeSubnetworks) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(computesubnetworksResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.ComputeSubnetworkList{})
	return err
}

// Patch applies the patch and returns the patched computeSubnetwork.
func (c *FakeComputeSubnetworks) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.ComputeSubnetwork, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(computesubnetworksResource, c.ns, name, pt, data, subresources...), &v1beta1.ComputeSubnetwork{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ComputeSubnetwork), err
}
