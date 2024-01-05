package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// TraitDefinition is the struct for OAM TraitDefinition construct
type TraitDefinition struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec TraitDefinitionSpec `json:"spec,omitempty"`
}

// TraitDefinitionSpec is the struct for OAM TraitDefinitionSpec's spec
type TraitDefinitionSpec struct {
	AppliesToWorkloads []string          `json:"appliesToWorkloads,omitempty"`
	DefinitionRef      DefinitionRef     `json:"definitionRef,omitempty"`
	RevisionEnabled    bool              `json:"revisionEnabled,omitempty"`
	Metadata           map[string]string `json:"metadata,omitempty"`
}
