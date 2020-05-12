


package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Actor
// +k8s:openapi-gen=true
// +resource:path=actors,strategy=ActorStrategy
type Actor struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ActorSpec   `json:"spec,omitempty"`
	Status ActorStatus `json:"status,omitempty"`
}

// ActorSpec defines the desired state of Actor
type ActorSpec struct {
	Name string
	Surname string
}

// ActorStatus defines the observed state of Actor
type ActorStatus struct {
}
