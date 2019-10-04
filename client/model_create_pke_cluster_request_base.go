/*
 * Pipeline API
 *
 * Pipeline is a feature rich application platform, built for containers on top of Kubernetes to automate the DevOps experience, continuous application development and the lifecycle of deployments. 
 *
 * API version: latest
 * Contact: info@banzaicloud.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client
// CreatePkeClusterRequestBase struct for CreatePkeClusterRequestBase
type CreatePkeClusterRequestBase struct {
	Name string `json:"name"`
	Features []Feature `json:"features,omitempty"`
	SecretId string `json:"secretId,omitempty"`
	SecretName string `json:"secretName,omitempty"`
	SshSecretId string `json:"sshSecretId,omitempty"`
	ScaleOptions ScaleOptions `json:"scaleOptions,omitempty"`
	Type string `json:"type"`
	Kubernetes CreatePkeClusterKubernetes `json:"kubernetes"`
}
