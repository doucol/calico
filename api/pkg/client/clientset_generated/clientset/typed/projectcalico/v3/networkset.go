// Copyright (c) 2025 Tigera, Inc. All rights reserved.

// Code generated by client-gen. DO NOT EDIT.

package v3

import (
	context "context"

	projectcalicov3 "github.com/projectcalico/api/pkg/apis/projectcalico/v3"
	scheme "github.com/projectcalico/api/pkg/client/clientset_generated/clientset/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// NetworkSetsGetter has a method to return a NetworkSetInterface.
// A group's client should implement this interface.
type NetworkSetsGetter interface {
	NetworkSets(namespace string) NetworkSetInterface
}

// NetworkSetInterface has methods to work with NetworkSet resources.
type NetworkSetInterface interface {
	Create(ctx context.Context, networkSet *projectcalicov3.NetworkSet, opts v1.CreateOptions) (*projectcalicov3.NetworkSet, error)
	Update(ctx context.Context, networkSet *projectcalicov3.NetworkSet, opts v1.UpdateOptions) (*projectcalicov3.NetworkSet, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*projectcalicov3.NetworkSet, error)
	List(ctx context.Context, opts v1.ListOptions) (*projectcalicov3.NetworkSetList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *projectcalicov3.NetworkSet, err error)
	NetworkSetExpansion
}

// networkSets implements NetworkSetInterface
type networkSets struct {
	*gentype.ClientWithList[*projectcalicov3.NetworkSet, *projectcalicov3.NetworkSetList]
}

// newNetworkSets returns a NetworkSets
func newNetworkSets(c *ProjectcalicoV3Client, namespace string) *networkSets {
	return &networkSets{
		gentype.NewClientWithList[*projectcalicov3.NetworkSet, *projectcalicov3.NetworkSetList](
			"networksets",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *projectcalicov3.NetworkSet { return &projectcalicov3.NetworkSet{} },
			func() *projectcalicov3.NetworkSetList { return &projectcalicov3.NetworkSetList{} },
		),
	}
}
