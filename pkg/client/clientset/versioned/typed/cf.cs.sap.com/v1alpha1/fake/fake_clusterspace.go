/*
SPDX-FileCopyrightText: 2023 SAP SE or an SAP affiliate company and cf-service-operator contributors
SPDX-License-Identifier: Apache-2.0
*/

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/sap/cf-service-operator/api/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeClusterSpaces implements ClusterSpaceInterface
type FakeClusterSpaces struct {
	Fake *FakeCfV1alpha1
}

var clusterspacesResource = schema.GroupVersionResource{Group: "cf.cs.sap.com", Version: "v1alpha1", Resource: "clusterspaces"}

var clusterspacesKind = schema.GroupVersionKind{Group: "cf.cs.sap.com", Version: "v1alpha1", Kind: "ClusterSpace"}

// Get takes name of the clusterSpace, and returns the corresponding clusterSpace object, and an error if there is any.
func (c *FakeClusterSpaces) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ClusterSpace, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(clusterspacesResource, name), &v1alpha1.ClusterSpace{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterSpace), err
}

// List takes label and field selectors, and returns the list of ClusterSpaces that match those selectors.
func (c *FakeClusterSpaces) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ClusterSpaceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(clusterspacesResource, clusterspacesKind, opts), &v1alpha1.ClusterSpaceList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ClusterSpaceList{ListMeta: obj.(*v1alpha1.ClusterSpaceList).ListMeta}
	for _, item := range obj.(*v1alpha1.ClusterSpaceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested clusterSpaces.
func (c *FakeClusterSpaces) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(clusterspacesResource, opts))
}

// Create takes the representation of a clusterSpace and creates it.  Returns the server's representation of the clusterSpace, and an error, if there is any.
func (c *FakeClusterSpaces) Create(ctx context.Context, clusterSpace *v1alpha1.ClusterSpace, opts v1.CreateOptions) (result *v1alpha1.ClusterSpace, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(clusterspacesResource, clusterSpace), &v1alpha1.ClusterSpace{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterSpace), err
}

// Update takes the representation of a clusterSpace and updates it. Returns the server's representation of the clusterSpace, and an error, if there is any.
func (c *FakeClusterSpaces) Update(ctx context.Context, clusterSpace *v1alpha1.ClusterSpace, opts v1.UpdateOptions) (result *v1alpha1.ClusterSpace, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(clusterspacesResource, clusterSpace), &v1alpha1.ClusterSpace{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterSpace), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeClusterSpaces) UpdateStatus(ctx context.Context, clusterSpace *v1alpha1.ClusterSpace, opts v1.UpdateOptions) (*v1alpha1.ClusterSpace, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(clusterspacesResource, "status", clusterSpace), &v1alpha1.ClusterSpace{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterSpace), err
}

// Delete takes name of the clusterSpace and deletes it. Returns an error if one occurs.
func (c *FakeClusterSpaces) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(clusterspacesResource, name, opts), &v1alpha1.ClusterSpace{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeClusterSpaces) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(clusterspacesResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ClusterSpaceList{})
	return err
}

// Patch applies the patch and returns the patched clusterSpace.
func (c *FakeClusterSpaces) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ClusterSpace, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(clusterspacesResource, name, pt, data, subresources...), &v1alpha1.ClusterSpace{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterSpace), err
}
