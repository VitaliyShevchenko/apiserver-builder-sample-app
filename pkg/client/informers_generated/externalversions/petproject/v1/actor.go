// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	petprojectv1 "gopath/src/pet-project/pkg/apis/petproject/v1"
	clientset "gopath/src/pet-project/pkg/client/clientset_generated/clientset"
	internalinterfaces "gopath/src/pet-project/pkg/client/informers_generated/externalversions/internalinterfaces"
	v1 "gopath/src/pet-project/pkg/client/listers_generated/petproject/v1"
	time "time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ActorInformer provides access to a shared informer and lister for
// Actors.
type ActorInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.ActorLister
}

type actorInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewActorInformer constructs a new informer for Actor type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewActorInformer(client clientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredActorInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredActorInformer constructs a new informer for Actor type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredActorInformer(client clientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.PetprojectV1().Actors(namespace).List(options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.PetprojectV1().Actors(namespace).Watch(options)
			},
		},
		&petprojectv1.Actor{},
		resyncPeriod,
		indexers,
	)
}

func (f *actorInformer) defaultInformer(client clientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredActorInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *actorInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&petprojectv1.Actor{}, f.defaultInformer)
}

func (f *actorInformer) Lister() v1.ActorLister {
	return v1.NewActorLister(f.Informer().GetIndexer())
}
