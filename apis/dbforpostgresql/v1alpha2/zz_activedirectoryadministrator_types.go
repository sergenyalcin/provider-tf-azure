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

package v1alpha2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type ActiveDirectoryAdministratorObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type ActiveDirectoryAdministratorParameters struct {

	// +kubebuilder:validation:Required
	ObjectID *string `json:"objectId" tf:"object_id,omitempty"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-azure/apis/azure/v1alpha2.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// +crossplane:generate:reference:type=Server
	// +kubebuilder:validation:Optional
	ServerName *string `json:"serverName,omitempty" tf:"server_name,omitempty"`

	// +kubebuilder:validation:Optional
	ServerNameRef *v1.Reference `json:"serverNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ServerNameSelector *v1.Selector `json:"serverNameSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Required
	TenantID *string `json:"tenantId" tf:"tenant_id,omitempty"`
}

// ActiveDirectoryAdministratorSpec defines the desired state of ActiveDirectoryAdministrator
type ActiveDirectoryAdministratorSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ActiveDirectoryAdministratorParameters `json:"forProvider"`
}

// ActiveDirectoryAdministratorStatus defines the observed state of ActiveDirectoryAdministrator.
type ActiveDirectoryAdministratorStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ActiveDirectoryAdministratorObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ActiveDirectoryAdministrator is the Schema for the ActiveDirectoryAdministrators API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type ActiveDirectoryAdministrator struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ActiveDirectoryAdministratorSpec   `json:"spec"`
	Status            ActiveDirectoryAdministratorStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ActiveDirectoryAdministratorList contains a list of ActiveDirectoryAdministrators
type ActiveDirectoryAdministratorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ActiveDirectoryAdministrator `json:"items"`
}

// Repository type metadata.
var (
	ActiveDirectoryAdministrator_Kind             = "ActiveDirectoryAdministrator"
	ActiveDirectoryAdministrator_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ActiveDirectoryAdministrator_Kind}.String()
	ActiveDirectoryAdministrator_KindAPIVersion   = ActiveDirectoryAdministrator_Kind + "." + CRDGroupVersion.String()
	ActiveDirectoryAdministrator_GroupVersionKind = CRDGroupVersion.WithKind(ActiveDirectoryAdministrator_Kind)
)

func init() {
	SchemeBuilder.Register(&ActiveDirectoryAdministrator{}, &ActiveDirectoryAdministratorList{})
}
