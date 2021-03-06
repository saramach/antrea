// Copyright 2021 Antrea Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by client-gen. DO NOT EDIT.

package v1beta2

import (
	"context"

	v1beta2 "github.com/vmware-tanzu/antrea/pkg/apis/controlplane/v1beta2"
	scheme "github.com/vmware-tanzu/antrea/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	rest "k8s.io/client-go/rest"
)

// ClusterGroupMembersGetter has a method to return a ClusterGroupMembersInterface.
// A group's client should implement this interface.
type ClusterGroupMembersGetter interface {
	ClusterGroupMembers() ClusterGroupMembersInterface
}

// ClusterGroupMembersInterface has methods to work with ClusterGroupMembers resources.
type ClusterGroupMembersInterface interface {
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta2.ClusterGroupMembers, error)
	ClusterGroupMembersExpansion
}

// clusterGroupMembers implements ClusterGroupMembersInterface
type clusterGroupMembers struct {
	client rest.Interface
}

// newClusterGroupMembers returns a ClusterGroupMembers
func newClusterGroupMembers(c *ControlplaneV1beta2Client) *clusterGroupMembers {
	return &clusterGroupMembers{
		client: c.RESTClient(),
	}
}

// Get takes name of the clusterGroupMembers, and returns the corresponding clusterGroupMembers object, and an error if there is any.
func (c *clusterGroupMembers) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta2.ClusterGroupMembers, err error) {
	result = &v1beta2.ClusterGroupMembers{}
	err = c.client.Get().
		Resource("clustergroupmembers").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}
