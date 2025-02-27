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

// FakeSpaces implements SpaceInterface
type FakeSpaces struct {
	Fake *FakeCfV1alpha1
	ns   string
}

var spacesResource = schema.GroupVersionResource{Group: "cf.cs.sap.com", Version: "v1alpha1", Resource: "spaces"}

var spacesKind = schema.GroupVersionKind{Group: "cf.cs.sap.com", Version: "v1alpha1", Kind: "Space"}

// Get takes name of the space, and returns the corresponding space object, and an error if there is any.
func (c *FakeSpaces) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Space, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(spacesResource, c.ns, name), &v1alpha1.Space{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Space), err
}

// List takes label and field selectors, and returns the list of Spaces that match those selectors.
func (c *FakeSpaces) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.SpaceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(spacesResource, spacesKind, c.ns, opts), &v1alpha1.SpaceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.SpaceList{ListMeta: obj.(*v1alpha1.SpaceList).ListMeta}
	for _, item := range obj.(*v1alpha1.SpaceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested spaces.
func (c *FakeSpaces) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(spacesResource, c.ns, opts))

}

// Create takes the representation of a space and creates it.  Returns the server's representation of the space, and an error, if there is any.
func (c *FakeSpaces) Create(ctx context.Context, space *v1alpha1.Space, opts v1.CreateOptions) (result *v1alpha1.Space, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(spacesResource, c.ns, space), &v1alpha1.Space{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Space), err
}

// Update takes the representation of a space and updates it. Returns the server's representation of the space, and an error, if there is any.
func (c *FakeSpaces) Update(ctx context.Context, space *v1alpha1.Space, opts v1.UpdateOptions) (result *v1alpha1.Space, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(spacesResource, c.ns, space), &v1alpha1.Space{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Space), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeSpaces) UpdateStatus(ctx context.Context, space *v1alpha1.Space, opts v1.UpdateOptions) (*v1alpha1.Space, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(spacesResource, "status", c.ns, space), &v1alpha1.Space{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Space), err
}

// Delete takes name of the space and deletes it. Returns an error if one occurs.
func (c *FakeSpaces) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(spacesResource, c.ns, name, opts), &v1alpha1.Space{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSpaces) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(spacesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.SpaceList{})
	return err
}

// Patch applies the patch and returns the patched space.
func (c *FakeSpaces) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Space, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(spacesResource, c.ns, name, pt, data, subresources...), &v1alpha1.Space{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Space), err
}
