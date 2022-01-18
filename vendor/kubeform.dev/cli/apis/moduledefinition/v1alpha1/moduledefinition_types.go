package v1alpha1

import (
	v1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ModuleRef struct {
	TfMarketplace string
}

type Provider struct {
	Name   string `json:"name"`
	Source string `json:"source"`
}

// ModuleDefinitionSpec defines the desired state of ModuleDefinition
type ModuleDefinitionSpec struct {
	Schema    v1.JSONSchemaProps `json:"schema"`
	ModuleRef ModuleRef          `json:"moduleRef"`
	Provider  Provider           `json:"provider"`
}

// ModuleDefinitionStatus defines the observed state of ModuleDefinition
type ModuleDefinitionStatus struct {
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// ModuleDefinition is the Schema for the moduledefinitions API
type ModuleDefinition struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ModuleDefinitionSpec   `json:"spec,omitempty"`
	Status ModuleDefinitionStatus `json:"status,omitempty"`
}
