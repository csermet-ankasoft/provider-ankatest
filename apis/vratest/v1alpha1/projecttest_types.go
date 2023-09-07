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

// ProjectTestParameters are the configurable fields of a ProjectTest.
type ProjectTestParameters struct {
	ConfigurableField string `json:"configurableField"`
}

// ProjectTestObservation are the observable fields of a ProjectTest.
type ProjectTestObservation struct {
	ObservableField string `json:"observableField,omitempty"`
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
