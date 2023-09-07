/*
Copyright 2022 The Crossplane Authors.

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

package v1alpha1

import (
	"reflect"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

// User A representation of a user.
type User struct {

	// The email of the user or name of the group.
	// Example: administrator@vmware.com
	// Required: true
	Email *string `json:"email"`

	// Type of the principal. Currently supported 'user' (default) and 'group'.
	// Example: user
	Type string `json:"type,omitempty"`
}

// Constraint A constraint that is conveyed to the policy engine.
type Constraint struct {

	// An expression of the form "[!]tag-key[:[tag-value]]", used to indicate a constraint match on keys and values of tags.
	//
	// Example: ha:strong
	// Required: true
	Expression *string `json:"expression"`

	// Indicates whether this constraint should be strictly enforced or not.
	// Required: true
	Mandatory *bool `json:"mandatory"`
}

// ZoneAssignmentSpecification A zone assignment configuration
type ZoneAssignmentSpecification struct {

	// The maximum amount of cpus that can be used by this cloud zone. Default is 0 (unlimited cpu).
	// Example: 2048
	CPULimit int64 `json:"cpuLimit,omitempty"`

	// The maximum number of instances that can be provisioned in this cloud zone. Default is 0 (unlimited instances).
	// Example: 50
	MaxNumberInstances int64 `json:"maxNumberInstances,omitempty"`

	// The maximum amount of memory that can be used by this cloud zone. Default is 0 (unlimited memory).
	// Example: 2048
	MemoryLimitMB int64 `json:"memoryLimitMB,omitempty"`

	// The priority of this zone in the current project. Lower numbers mean higher priority. Default is 0 (highest)
	// Example: 1
	Priority int32 `json:"priority,omitempty"`

	// Defines an upper limit on storage that can be requested from a cloud zone which is part of this project. Default is 0 (unlimited storage). Please note that this feature is supported only for vSphere cloud zones. Not valid for other cloud zone types.
	// Example: 20
	StorageLimitGB int64 `json:"storageLimitGB,omitempty"`

	// The Cloud Zone Id
	// Example: 77ee1
	ZoneID string `json:"zoneId,omitempty"`
}

// ProjectTestParameters are the configurable fields of a Project.
type ProjectTestParameters struct {
	// A human-friendly name used as an identifier in APIs that support this option.
	// Required: true
	// +immutable
	Name *string `json:"name"`

	// They are not required.
	Administrators               []*User                        `json:"administrators,omitempty"`
	Viewers                      []*User                        `json:"viewers,omitempty"`
	Members                      []*User                        `json:"members,omitempty"`
	MachineNamingTemplate        string                         `json:"machineNamingTemplate,omitempty"`
	PlacementPolicy              string                         `json:"placementPolicy,omitempty"`
	OperationTimeout             *int64                         `json:"operationTimeout,omitempty"`
	Description                  string                         `json:"description,omitempty"`
	SharedResources              *bool                          `json:"sharedResources,omitempty"`
	Constraints                  map[string][]Constraint        `json:"constraints,omitempty"`
	CustomProperties             map[string]string              `json:"customProperties,omitempty"`
	ZoneAssignmentConfigurations []*ZoneAssignmentSpecification `json:"zoneAssignmentConfigurations,omitempty"`
}

// ProjectTestObservation are the observable fields of a Project.
type ProjectTestObservation struct {
	Name                         *string                        `json:"name,omitempty"`
	ID                           string                         `json:"id"`
	Administrators               []*User                        `json:"administrators,omitempty"`
	Viewers                      []*User                        `json:"viewers,omitempty"`
	Members                      []*User                        `json:"members,omitempty"`
	MachineNamingTemplate        string                         `json:"machineNamingTemplate,omitempty"`
	PlacementPolicy              string                         `json:"placementPolicy,omitempty"`
	OperationTimeout             *int64                         `json:"operationTimeout,omitempty"`
	Description                  string                         `json:"description,omitempty"`
	SharedResources              bool                           `json:"sharedResources,omitempty"`
	Constraints                  map[string][]Constraint        `json:"constraints,omitempty"`
	CustomProperties             map[string]string              `json:"customProperties,omitempty"`
	ZoneAssignmentConfigurations []*ZoneAssignmentSpecification `json:"zoneAssignmentConfigurations,omitempty"`
}

// A ProjectTestSpec defines the desired state of a ProjectTest.
type ProjectTestSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       ProjectTestParameters `json:"forProvider"`
}

// A ProjectTestStatus represents the observed state of a ProjectTest.
type ProjectTestStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          ProjectTestObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// A ProjectTest is an example API type.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,ankatest}
type ProjectTest struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ProjectTestSpec   `json:"spec"`
	Status ProjectTestStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ProjectTestList contains a list of ProjectTest
type ProjectTestList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ProjectTest `json:"items"`
}

// ProjectTest type metadata.
var (
	ProjectTestKind             = reflect.TypeOf(ProjectTest{}).Name()
	ProjectTestGroupKind        = schema.GroupKind{Group: Group, Kind: ProjectTestKind}.String()
	ProjectTestKindAPIVersion   = ProjectTestKind + "." + SchemeGroupVersion.String()
	ProjectTestGroupVersionKind = SchemeGroupVersion.WithKind(ProjectTestKind)
)

func init() {
	SchemeBuilder.Register(&ProjectTest{}, &ProjectTestList{})
}
