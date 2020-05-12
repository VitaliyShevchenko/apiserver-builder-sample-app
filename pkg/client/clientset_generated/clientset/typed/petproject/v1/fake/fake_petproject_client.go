// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1 "gopath/src/pet-project/pkg/client/clientset_generated/clientset/typed/petproject/v1"

	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakePetprojectV1 struct {
	*testing.Fake
}

func (c *FakePetprojectV1) Actors(namespace string) v1.ActorInterface {
	return &FakeActors{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakePetprojectV1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
