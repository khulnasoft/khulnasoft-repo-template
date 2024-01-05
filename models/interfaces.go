package models

import "github.com/khulnasoft/meshkit/models/meshmodel/core/v1alpha1"

// anything that can be validated is a Validator
type Validator interface {
	Validate([]byte) error
}

// An entity that is used to expose a particular
// system's capabilities in Meshplay
// A Package should have all the information that we need to generate the components
type Package interface {
	GenerateComponents() ([]v1alpha1.ComponentDefinition, error)
}

// Supports pulling packages from Artifact Hub and other sources like Docker Hub.
// Should envelope Meshplay Application importer - to be implemented
type PackageManager interface {
	GetPackage() (Package, error)
}
