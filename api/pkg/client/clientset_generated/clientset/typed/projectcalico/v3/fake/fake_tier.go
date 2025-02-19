// Copyright (c) 2025 Tigera, Inc. All rights reserved.

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v3 "github.com/projectcalico/api/pkg/apis/projectcalico/v3"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeTiers implements TierInterface
type FakeTiers struct {
	Fake *FakeProjectcalicoV3
}

var tiersResource = v3.SchemeGroupVersion.WithResource("tiers")

var tiersKind = v3.SchemeGroupVersion.WithKind("Tier")

// Get takes name of the tier, and returns the corresponding tier object, and an error if there is any.
func (c *FakeTiers) Get(ctx context.Context, name string, options v1.GetOptions) (result *v3.Tier, err error) {
	emptyResult := &v3.Tier{}
	obj, err := c.Fake.
		Invokes(testing.NewRootGetActionWithOptions(tiersResource, name, options), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v3.Tier), err
}

// List takes label and field selectors, and returns the list of Tiers that match those selectors.
func (c *FakeTiers) List(ctx context.Context, opts v1.ListOptions) (result *v3.TierList, err error) {
	emptyResult := &v3.TierList{}
	obj, err := c.Fake.
		Invokes(testing.NewRootListActionWithOptions(tiersResource, tiersKind, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v3.TierList{ListMeta: obj.(*v3.TierList).ListMeta}
	for _, item := range obj.(*v3.TierList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested tiers.
func (c *FakeTiers) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchActionWithOptions(tiersResource, opts))
}

// Create takes the representation of a tier and creates it.  Returns the server's representation of the tier, and an error, if there is any.
func (c *FakeTiers) Create(ctx context.Context, tier *v3.Tier, opts v1.CreateOptions) (result *v3.Tier, err error) {
	emptyResult := &v3.Tier{}
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateActionWithOptions(tiersResource, tier, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v3.Tier), err
}

// Update takes the representation of a tier and updates it. Returns the server's representation of the tier, and an error, if there is any.
func (c *FakeTiers) Update(ctx context.Context, tier *v3.Tier, opts v1.UpdateOptions) (result *v3.Tier, err error) {
	emptyResult := &v3.Tier{}
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateActionWithOptions(tiersResource, tier, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v3.Tier), err
}

// Delete takes name of the tier and deletes it. Returns an error if one occurs.
func (c *FakeTiers) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(tiersResource, name, opts), &v3.Tier{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeTiers) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionActionWithOptions(tiersResource, opts, listOpts)

	_, err := c.Fake.Invokes(action, &v3.TierList{})
	return err
}

// Patch applies the patch and returns the patched tier.
func (c *FakeTiers) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v3.Tier, err error) {
	emptyResult := &v3.Tier{}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceActionWithOptions(tiersResource, name, pt, data, opts, subresources...), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v3.Tier), err
}
