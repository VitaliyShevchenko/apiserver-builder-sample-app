


package actoradmission

import (
	"context"
	aggregatedadmission "gopath/src/pet-project/plugin/admission"
	aggregatedinformerfactory "gopath/src/pet-project/pkg/client/informers_generated/externalversions"
	aggregatedclientset "gopath/src/pet-project/pkg/client/clientset_generated/clientset"
	genericadmissioninitializer "k8s.io/apiserver/pkg/admission/initializer"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/apiserver/pkg/admission"
)

var _ admission.Interface 											= &actorPlugin{}
var _ admission.MutationInterface 									= &actorPlugin{}
var _ admission.ValidationInterface 								= &actorPlugin{}
var _ genericadmissioninitializer.WantsExternalKubeInformerFactory 	= &actorPlugin{}
var _ genericadmissioninitializer.WantsExternalKubeClientSet 		= &actorPlugin{}
var _ aggregatedadmission.WantsAggregatedResourceInformerFactory 	= &actorPlugin{}
var _ aggregatedadmission.WantsAggregatedResourceClientSet 			= &actorPlugin{}

func NewActorPlugin() *actorPlugin {
	return &actorPlugin{
		Handler: admission.NewHandler(admission.Create, admission.Update),
	}
}

type actorPlugin struct {
	*admission.Handler
}

func (p *actorPlugin) ValidateInitialization() error {
	return nil
}

func (p *actorPlugin) Admit(ctx context.Context, a admission.Attributes, o admission.ObjectInterfaces) error {
	return nil
}

func (p *actorPlugin) Validate(ctx context.Context, a admission.Attributes, o admission.ObjectInterfaces) error {
	return nil
}

func (p *actorPlugin) SetAggregatedResourceInformerFactory(aggregatedinformerfactory.SharedInformerFactory) {}

func (p *actorPlugin) SetAggregatedResourceClientSet(aggregatedclientset.Interface) {}

func (p *actorPlugin) SetExternalKubeInformerFactory(informers.SharedInformerFactory) {}

func (p *actorPlugin) SetExternalKubeClientSet(kubernetes.Interface) {}
