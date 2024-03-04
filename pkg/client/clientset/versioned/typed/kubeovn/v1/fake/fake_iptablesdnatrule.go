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

package fake

import (
	"context"

	v1 "github.com/kubeovn/kube-ovn/pkg/apis/kubeovn/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeIptablesDnatRules implements IptablesDnatRuleInterface
type FakeIptablesDnatRules struct {
	Fake *FakeKubeovnV1
}

var iptablesdnatrulesResource = v1.SchemeGroupVersion.WithResource("iptables-dnat-rules")

var iptablesdnatrulesKind = v1.SchemeGroupVersion.WithKind("IptablesDnatRule")

// Get takes name of the iptablesDnatRule, and returns the corresponding iptablesDnatRule object, and an error if there is any.
func (c *FakeIptablesDnatRules) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.IptablesDnatRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(iptablesdnatrulesResource, name), &v1.IptablesDnatRule{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.IptablesDnatRule), err
}

// List takes label and field selectors, and returns the list of IptablesDnatRules that match those selectors.
func (c *FakeIptablesDnatRules) List(ctx context.Context, opts metav1.ListOptions) (result *v1.IptablesDnatRuleList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(iptablesdnatrulesResource, iptablesdnatrulesKind, opts), &v1.IptablesDnatRuleList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.IptablesDnatRuleList{ListMeta: obj.(*v1.IptablesDnatRuleList).ListMeta}
	for _, item := range obj.(*v1.IptablesDnatRuleList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested iptablesDnatRules.
func (c *FakeIptablesDnatRules) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(iptablesdnatrulesResource, opts))
}

// Create takes the representation of a iptablesDnatRule and creates it.  Returns the server's representation of the iptablesDnatRule, and an error, if there is any.
func (c *FakeIptablesDnatRules) Create(ctx context.Context, iptablesDnatRule *v1.IptablesDnatRule, opts metav1.CreateOptions) (result *v1.IptablesDnatRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(iptablesdnatrulesResource, iptablesDnatRule), &v1.IptablesDnatRule{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.IptablesDnatRule), err
}

// Update takes the representation of a iptablesDnatRule and updates it. Returns the server's representation of the iptablesDnatRule, and an error, if there is any.
func (c *FakeIptablesDnatRules) Update(ctx context.Context, iptablesDnatRule *v1.IptablesDnatRule, opts metav1.UpdateOptions) (result *v1.IptablesDnatRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(iptablesdnatrulesResource, iptablesDnatRule), &v1.IptablesDnatRule{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.IptablesDnatRule), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeIptablesDnatRules) UpdateStatus(ctx context.Context, iptablesDnatRule *v1.IptablesDnatRule, opts metav1.UpdateOptions) (*v1.IptablesDnatRule, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(iptablesdnatrulesResource, "status", iptablesDnatRule), &v1.IptablesDnatRule{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.IptablesDnatRule), err
}

// Delete takes name of the iptablesDnatRule and deletes it. Returns an error if one occurs.
func (c *FakeIptablesDnatRules) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(iptablesdnatrulesResource, name, opts), &v1.IptablesDnatRule{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeIptablesDnatRules) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(iptablesdnatrulesResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1.IptablesDnatRuleList{})
	return err
}

// Patch applies the patch and returns the patched iptablesDnatRule.
func (c *FakeIptablesDnatRules) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.IptablesDnatRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(iptablesdnatrulesResource, name, pt, data, subresources...), &v1.IptablesDnatRule{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.IptablesDnatRule), err
}