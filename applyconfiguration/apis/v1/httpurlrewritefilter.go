/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	apisv1 "sigs.k8s.io/gateway-api/apis/v1"
)

// HTTPURLRewriteFilterApplyConfiguration represents a declarative configuration of the HTTPURLRewriteFilter type for use
// with apply.
type HTTPURLRewriteFilterApplyConfiguration struct {
	Hostname *apisv1.PreciseHostname             `json:"hostname,omitempty"`
	Path     *HTTPPathModifierApplyConfiguration `json:"path,omitempty"`
}

// HTTPURLRewriteFilterApplyConfiguration constructs a declarative configuration of the HTTPURLRewriteFilter type for use with
// apply.
func HTTPURLRewriteFilter() *HTTPURLRewriteFilterApplyConfiguration {
	return &HTTPURLRewriteFilterApplyConfiguration{}
}

// WithHostname sets the Hostname field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Hostname field is set to the value of the last call.
func (b *HTTPURLRewriteFilterApplyConfiguration) WithHostname(value apisv1.PreciseHostname) *HTTPURLRewriteFilterApplyConfiguration {
	b.Hostname = &value
	return b
}

// WithPath sets the Path field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Path field is set to the value of the last call.
func (b *HTTPURLRewriteFilterApplyConfiguration) WithPath(value *HTTPPathModifierApplyConfiguration) *HTTPURLRewriteFilterApplyConfiguration {
	b.Path = value
	return b
}
