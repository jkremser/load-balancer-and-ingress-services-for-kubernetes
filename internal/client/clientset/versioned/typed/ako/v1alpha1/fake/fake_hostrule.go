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
	v1alpha1 "github.com/avinetworks/ako/internal/apis/ako/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeHostRules implements HostRuleInterface
type FakeHostRules struct {
	Fake *FakeAkoV1alpha1
	ns   string
}

var hostrulesResource = schema.GroupVersionResource{Group: "ako.vmware.com", Version: "v1alpha1", Resource: "hostrules"}

var hostrulesKind = schema.GroupVersionKind{Group: "ako.vmware.com", Version: "v1alpha1", Kind: "HostRule"}

// Get takes name of the hostRule, and returns the corresponding hostRule object, and an error if there is any.
func (c *FakeHostRules) Get(name string, options v1.GetOptions) (result *v1alpha1.HostRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(hostrulesResource, c.ns, name), &v1alpha1.HostRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HostRule), err
}

// List takes label and field selectors, and returns the list of HostRules that match those selectors.
func (c *FakeHostRules) List(opts v1.ListOptions) (result *v1alpha1.HostRuleList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(hostrulesResource, hostrulesKind, c.ns, opts), &v1alpha1.HostRuleList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.HostRuleList{ListMeta: obj.(*v1alpha1.HostRuleList).ListMeta}
	for _, item := range obj.(*v1alpha1.HostRuleList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested hostRules.
func (c *FakeHostRules) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(hostrulesResource, c.ns, opts))

}

// Create takes the representation of a hostRule and creates it.  Returns the server's representation of the hostRule, and an error, if there is any.
func (c *FakeHostRules) Create(hostRule *v1alpha1.HostRule) (result *v1alpha1.HostRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(hostrulesResource, c.ns, hostRule), &v1alpha1.HostRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HostRule), err
}

// Update takes the representation of a hostRule and updates it. Returns the server's representation of the hostRule, and an error, if there is any.
func (c *FakeHostRules) Update(hostRule *v1alpha1.HostRule) (result *v1alpha1.HostRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(hostrulesResource, c.ns, hostRule), &v1alpha1.HostRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HostRule), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeHostRules) UpdateStatus(hostRule *v1alpha1.HostRule) (*v1alpha1.HostRule, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(hostrulesResource, "status", c.ns, hostRule), &v1alpha1.HostRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HostRule), err
}

// Delete takes name of the hostRule and deletes it. Returns an error if one occurs.
func (c *FakeHostRules) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(hostrulesResource, c.ns, name), &v1alpha1.HostRule{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeHostRules) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(hostrulesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.HostRuleList{})
	return err
}

// Patch applies the patch and returns the patched hostRule.
func (c *FakeHostRules) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.HostRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(hostrulesResource, c.ns, name, pt, data, subresources...), &v1alpha1.HostRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HostRule), err
}
