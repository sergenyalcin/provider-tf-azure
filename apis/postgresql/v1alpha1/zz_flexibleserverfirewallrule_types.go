/*
Copyright 2020 The Crossplane Authors.

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

// Code generated by terrajet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type FlexibleServerFirewallRuleObservation struct {
}

type FlexibleServerFirewallRuleParameters struct {

	// +kubebuilder:validation:Required
	EndIPAddress *string `json:"endIpAddress" tf:"end_ip_address,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +crossplane:generate:reference:type=FlexibleServer
	// +kubebuilder:validation:Optional
	ServerID *string `json:"serverId,omitempty" tf:"server_id,omitempty"`

	// +kubebuilder:validation:Optional
	ServerIDRef *v1.Reference `json:"serverIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ServerIDSelector *v1.Selector `json:"serverIdSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Required
	StartIPAddress *string `json:"startIpAddress" tf:"start_ip_address,omitempty"`
}

// FlexibleServerFirewallRuleSpec defines the desired state of FlexibleServerFirewallRule
type FlexibleServerFirewallRuleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FlexibleServerFirewallRuleParameters `json:"forProvider"`
}

// FlexibleServerFirewallRuleStatus defines the observed state of FlexibleServerFirewallRule.
type FlexibleServerFirewallRuleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FlexibleServerFirewallRuleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// FlexibleServerFirewallRule is the Schema for the FlexibleServerFirewallRules API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfazure}
type FlexibleServerFirewallRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FlexibleServerFirewallRuleSpec   `json:"spec"`
	Status            FlexibleServerFirewallRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FlexibleServerFirewallRuleList contains a list of FlexibleServerFirewallRules
type FlexibleServerFirewallRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FlexibleServerFirewallRule `json:"items"`
}

// Repository type metadata.
var (
	FlexibleServerFirewallRule_Kind             = "FlexibleServerFirewallRule"
	FlexibleServerFirewallRule_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: FlexibleServerFirewallRule_Kind}.String()
	FlexibleServerFirewallRule_KindAPIVersion   = FlexibleServerFirewallRule_Kind + "." + CRDGroupVersion.String()
	FlexibleServerFirewallRule_GroupVersionKind = CRDGroupVersion.WithKind(FlexibleServerFirewallRule_Kind)
)

func init() {
	SchemeBuilder.Register(&FlexibleServerFirewallRule{}, &FlexibleServerFirewallRuleList{})
}