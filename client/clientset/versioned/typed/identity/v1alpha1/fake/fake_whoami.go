/*
Copyright The Kubeshield Authors.

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

	v1alpha1 "kubeshield.dev/identity-server/apis/identity/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	testing "k8s.io/client-go/testing"
)

// FakeWhoAmIs implements WhoAmIInterface
type FakeWhoAmIs struct {
	Fake *FakeIdentityV1alpha1
}

var whoamisResource = schema.GroupVersionResource{Group: "identity.kubeshield.io", Version: "v1alpha1", Resource: "whoamis"}

var whoamisKind = schema.GroupVersionKind{Group: "identity.kubeshield.io", Version: "v1alpha1", Kind: "WhoAmI"}

// Create takes the representation of a whoAmI and creates it.  Returns the server's representation of the whoAmI, and an error, if there is any.
func (c *FakeWhoAmIs) Create(ctx context.Context, whoAmI *v1alpha1.WhoAmI, opts v1.CreateOptions) (result *v1alpha1.WhoAmI, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(whoamisResource, whoAmI), &v1alpha1.WhoAmI{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.WhoAmI), err
}
