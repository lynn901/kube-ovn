/*
Copyright The Kubernetes Authors.

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

package v1

import (
	"context"

	v1 "github.com/kubeovn/kube-ovn/pkg/apis/kubeovn/v1"
	scheme "github.com/kubeovn/kube-ovn/pkg/client/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// SecurityGroupsGetter has a method to return a SecurityGroupInterface.
// A group's client should implement this interface.
type SecurityGroupsGetter interface {
	SecurityGroups() SecurityGroupInterface
}

// SecurityGroupInterface has methods to work with SecurityGroup resources.
type SecurityGroupInterface interface {
	Create(ctx context.Context, securityGroup *v1.SecurityGroup, opts metav1.CreateOptions) (*v1.SecurityGroup, error)
	Update(ctx context.Context, securityGroup *v1.SecurityGroup, opts metav1.UpdateOptions) (*v1.SecurityGroup, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, securityGroup *v1.SecurityGroup, opts metav1.UpdateOptions) (*v1.SecurityGroup, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.SecurityGroup, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.SecurityGroupList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.SecurityGroup, err error)
	SecurityGroupExpansion
}

// securityGroups implements SecurityGroupInterface
type securityGroups struct {
	*gentype.ClientWithList[*v1.SecurityGroup, *v1.SecurityGroupList]
}

// newSecurityGroups returns a SecurityGroups
func newSecurityGroups(c *KubeovnV1Client) *securityGroups {
	return &securityGroups{
		gentype.NewClientWithList[*v1.SecurityGroup, *v1.SecurityGroupList](
			"security-groups",
			c.RESTClient(),
			scheme.ParameterCodec,
			"",
			func() *v1.SecurityGroup { return &v1.SecurityGroup{} },
			func() *v1.SecurityGroupList { return &v1.SecurityGroupList{} }),
	}
}