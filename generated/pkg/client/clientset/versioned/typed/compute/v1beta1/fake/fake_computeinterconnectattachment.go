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

// FakeComputeInterconnectAttachments implements ComputeInterconnectAttachmentInterface
type FakeComputeInterconnectAttachments struct {
	Fake *FakeComputeV1beta1
	ns   string
}

var computeinterconnectattachmentsResource = schema.GroupVersionResource{Group: "compute.cnrm.cloud.google.com", Version: "v1beta1", Resource: "computeinterconnectattachments"}

var computeinterconnectattachmentsKind = schema.GroupVersionKind{Group: "compute.cnrm.cloud.google.com", Version: "v1beta1", Kind: "ComputeInterconnectAttachment"}

// Get takes name of the computeInterconnectAttachment, and returns the corresponding computeInterconnectAttachment object, and an error if there is any.
func (c *FakeComputeInterconnectAttachments) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.ComputeInterconnectAttachment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(computeinterconnectattachmentsResource, c.ns, name), &v1beta1.ComputeInterconnectAttachment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ComputeInterconnectAttachment), err
}

// List takes label and field selectors, and returns the list of ComputeInterconnectAttachments that match those selectors.
func (c *FakeComputeInterconnectAttachments) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.ComputeInterconnectAttachmentList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(computeinterconnectattachmentsResource, computeinterconnectattachmentsKind, c.ns, opts), &v1beta1.ComputeInterconnectAttachmentList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.ComputeInterconnectAttachmentList{ListMeta: obj.(*v1beta1.ComputeInterconnectAttachmentList).ListMeta}
	for _, item := range obj.(*v1beta1.ComputeInterconnectAttachmentList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested computeInterconnectAttachments.
func (c *FakeComputeInterconnectAttachments) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(computeinterconnectattachmentsResource, c.ns, opts))

}

// Create takes the representation of a computeInterconnectAttachment and creates it.  Returns the server's representation of the computeInterconnectAttachment, and an error, if there is any.
func (c *FakeComputeInterconnectAttachments) Create(ctx context.Context, computeInterconnectAttachment *v1beta1.ComputeInterconnectAttachment, opts v1.CreateOptions) (result *v1beta1.ComputeInterconnectAttachment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(computeinterconnectattachmentsResource, c.ns, computeInterconnectAttachment), &v1beta1.ComputeInterconnectAttachment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ComputeInterconnectAttachment), err
}

// Update takes the representation of a computeInterconnectAttachment and updates it. Returns the server's representation of the computeInterconnectAttachment, and an error, if there is any.
func (c *FakeComputeInterconnectAttachments) Update(ctx context.Context, computeInterconnectAttachment *v1beta1.ComputeInterconnectAttachment, opts v1.UpdateOptions) (result *v1beta1.ComputeInterconnectAttachment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(computeinterconnectattachmentsResource, c.ns, computeInterconnectAttachment), &v1beta1.ComputeInterconnectAttachment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ComputeInterconnectAttachment), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeComputeInterconnectAttachments) UpdateStatus(ctx context.Context, computeInterconnectAttachment *v1beta1.ComputeInterconnectAttachment, opts v1.UpdateOptions) (*v1beta1.ComputeInterconnectAttachment, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(computeinterconnectattachmentsResource, "status", c.ns, computeInterconnectAttachment), &v1beta1.ComputeInterconnectAttachment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ComputeInterconnectAttachment), err
}

// Delete takes name of the computeInterconnectAttachment and deletes it. Returns an error if one occurs.
func (c *FakeComputeInterconnectAttachments) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(computeinterconnectattachmentsResource, c.ns, name), &v1beta1.ComputeInterconnectAttachment{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeComputeInterconnectAttachments) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(computeinterconnectattachmentsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.ComputeInterconnectAttachmentList{})
	return err
}

// Patch applies the patch and returns the patched computeInterconnectAttachment.
func (c *FakeComputeInterconnectAttachments) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.ComputeInterconnectAttachment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(computeinterconnectattachmentsResource, c.ns, name, pt, data, subresources...), &v1beta1.ComputeInterconnectAttachment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ComputeInterconnectAttachment), err
}
