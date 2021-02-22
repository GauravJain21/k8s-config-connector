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

	v1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/sql/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeSQLDatabases implements SQLDatabaseInterface
type FakeSQLDatabases struct {
	Fake *FakeSqlV1beta1
	ns   string
}

var sqldatabasesResource = schema.GroupVersionResource{Group: "sql.cnrm.cloud.google.com", Version: "v1beta1", Resource: "sqldatabases"}

var sqldatabasesKind = schema.GroupVersionKind{Group: "sql.cnrm.cloud.google.com", Version: "v1beta1", Kind: "SQLDatabase"}

// Get takes name of the sQLDatabase, and returns the corresponding sQLDatabase object, and an error if there is any.
func (c *FakeSQLDatabases) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.SQLDatabase, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(sqldatabasesResource, c.ns, name), &v1beta1.SQLDatabase{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.SQLDatabase), err
}

// List takes label and field selectors, and returns the list of SQLDatabases that match those selectors.
func (c *FakeSQLDatabases) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.SQLDatabaseList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(sqldatabasesResource, sqldatabasesKind, c.ns, opts), &v1beta1.SQLDatabaseList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.SQLDatabaseList{ListMeta: obj.(*v1beta1.SQLDatabaseList).ListMeta}
	for _, item := range obj.(*v1beta1.SQLDatabaseList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested sQLDatabases.
func (c *FakeSQLDatabases) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(sqldatabasesResource, c.ns, opts))

}

// Create takes the representation of a sQLDatabase and creates it.  Returns the server's representation of the sQLDatabase, and an error, if there is any.
func (c *FakeSQLDatabases) Create(ctx context.Context, sQLDatabase *v1beta1.SQLDatabase, opts v1.CreateOptions) (result *v1beta1.SQLDatabase, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(sqldatabasesResource, c.ns, sQLDatabase), &v1beta1.SQLDatabase{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.SQLDatabase), err
}

// Update takes the representation of a sQLDatabase and updates it. Returns the server's representation of the sQLDatabase, and an error, if there is any.
func (c *FakeSQLDatabases) Update(ctx context.Context, sQLDatabase *v1beta1.SQLDatabase, opts v1.UpdateOptions) (result *v1beta1.SQLDatabase, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(sqldatabasesResource, c.ns, sQLDatabase), &v1beta1.SQLDatabase{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.SQLDatabase), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeSQLDatabases) UpdateStatus(ctx context.Context, sQLDatabase *v1beta1.SQLDatabase, opts v1.UpdateOptions) (*v1beta1.SQLDatabase, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(sqldatabasesResource, "status", c.ns, sQLDatabase), &v1beta1.SQLDatabase{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.SQLDatabase), err
}

// Delete takes name of the sQLDatabase and deletes it. Returns an error if one occurs.
func (c *FakeSQLDatabases) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(sqldatabasesResource, c.ns, name), &v1beta1.SQLDatabase{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSQLDatabases) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(sqldatabasesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.SQLDatabaseList{})
	return err
}

// Patch applies the patch and returns the patched sQLDatabase.
func (c *FakeSQLDatabases) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.SQLDatabase, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(sqldatabasesResource, c.ns, name, pt, data, subresources...), &v1beta1.SQLDatabase{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.SQLDatabase), err
}
