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
// ScaleOptions struct for ScaleOptions
type ScaleOptions struct {
	Enabled bool `json:"enabled"`
	DesiredCpu float64 `json:"desiredCpu,omitempty"`
	DesiredMem float64 `json:"desiredMem,omitempty"`
	DesiredGpu int32 `json:"desiredGpu,omitempty"`
	OnDemandPct int32 `json:"onDemandPct,omitempty"`
	Excludes []string `json:"excludes,omitempty"`
	KeepDesiredCapacity bool `json:"keepDesiredCapacity,omitempty"`
}
