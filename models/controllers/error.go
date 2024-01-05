package controllers

import (
	"github.com/khulnasoft/meshkit/errors"
)

var (
	ErrGetControllerStatusCode         = "11080"
	ErrDeployControllerCode            = "11081"
	ErrGetControllerPublicEndpointCode = "11082"
)

func ErrGetControllerStatus(err error) error {
	return errors.New(ErrGetControllerStatusCode, errors.Alert, []string{"Error getting the status of the meshplay controller"}, []string{err.Error()}, []string{"Controller may not be healthy or not deployed"}, []string{"Make sure the controller is deployed and healthy"})
}

func ErrDeployController(err error) error {
	return errors.New(ErrDeployControllerCode, errors.Alert, []string{"Error deploying Meshplay Operator"}, []string{err.Error()}, []string{"Meshplay Server could not connect to the Kubernetes cluster. Meshplay Operator  was not deployed", "Insufficient file permission to read kubeconfig"}, []string{"Verify that the available kubeconfig is accessible by Meshplay Server - verify sufficient file permissions (only needs read permission)"})
}

func ErrGetControllerPublicEndpoint(err error) error {
	return errors.New(ErrGetControllerPublicEndpointCode, errors.Alert, []string{"Could not get the public endpoint of the controller"}, []string{err.Error()}, []string{"Client configuration may not be valid"}, []string{"Make sure the client configuration is valid"})
}
