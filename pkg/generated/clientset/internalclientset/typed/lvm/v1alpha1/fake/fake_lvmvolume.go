/*
Copyright 2021 The OpenEBS Authors

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
	v1alpha1 "github.com/openebs/lvm-localpv/pkg/apis/openebs.io/lvm/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeLVMVolumes implements LVMVolumeInterface
type FakeLVMVolumes struct {
	Fake *FakeLocalV1alpha1
	ns   string
}

var lvmvolumesResource = schema.GroupVersionResource{Group: "local.openebs.io", Version: "v1alpha1", Resource: "lvmvolumes"}

var lvmvolumesKind = schema.GroupVersionKind{Group: "local.openebs.io", Version: "v1alpha1", Kind: "LVMVolume"}

// Get takes name of the lVMVolume, and returns the corresponding lVMVolume object, and an error if there is any.
func (c *FakeLVMVolumes) Get(name string, options v1.GetOptions) (result *v1alpha1.LVMVolume, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(lvmvolumesResource, c.ns, name), &v1alpha1.LVMVolume{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LVMVolume), err
}

// List takes label and field selectors, and returns the list of LVMVolumes that match those selectors.
func (c *FakeLVMVolumes) List(opts v1.ListOptions) (result *v1alpha1.LVMVolumeList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(lvmvolumesResource, lvmvolumesKind, c.ns, opts), &v1alpha1.LVMVolumeList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.LVMVolumeList{ListMeta: obj.(*v1alpha1.LVMVolumeList).ListMeta}
	for _, item := range obj.(*v1alpha1.LVMVolumeList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested lVMVolumes.
func (c *FakeLVMVolumes) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(lvmvolumesResource, c.ns, opts))

}

// Create takes the representation of a lVMVolume and creates it.  Returns the server's representation of the lVMVolume, and an error, if there is any.
func (c *FakeLVMVolumes) Create(lVMVolume *v1alpha1.LVMVolume) (result *v1alpha1.LVMVolume, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(lvmvolumesResource, c.ns, lVMVolume), &v1alpha1.LVMVolume{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LVMVolume), err
}

// Update takes the representation of a lVMVolume and updates it. Returns the server's representation of the lVMVolume, and an error, if there is any.
func (c *FakeLVMVolumes) Update(lVMVolume *v1alpha1.LVMVolume) (result *v1alpha1.LVMVolume, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(lvmvolumesResource, c.ns, lVMVolume), &v1alpha1.LVMVolume{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LVMVolume), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeLVMVolumes) UpdateStatus(lVMVolume *v1alpha1.LVMVolume) (*v1alpha1.LVMVolume, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(lvmvolumesResource, "status", c.ns, lVMVolume), &v1alpha1.LVMVolume{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LVMVolume), err
}

// Delete takes name of the lVMVolume and deletes it. Returns an error if one occurs.
func (c *FakeLVMVolumes) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(lvmvolumesResource, c.ns, name), &v1alpha1.LVMVolume{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeLVMVolumes) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(lvmvolumesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.LVMVolumeList{})
	return err
}

// Patch applies the patch and returns the patched lVMVolume.
func (c *FakeLVMVolumes) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.LVMVolume, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(lvmvolumesResource, c.ns, name, pt, data, subresources...), &v1alpha1.LVMVolume{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LVMVolume), err
}
