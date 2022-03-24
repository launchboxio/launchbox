/*
Copyright 2022.

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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// RevisionSpec defines the desired state of Revision

type Autoscaling struct {
	Enabled                 bool  `json:"enabled,omitempty"`
	MaxSize                 int32 `json:"maxSize,omitempty"`
	MinSize                 int32 `json:"minSize,omitempty"`
	TargetCpuUtilization    int32 `json:"targetCpuUtilization,omitempty"`
	TargetMemoryUtilization int32 `json:"targetMemoryUtilization,omitempty"`
	TargetRequestsPerSecond int32 `json:"targetRequestsPerSecond,omitempty"`
}

type RevisionSpec struct {
	CommitSha   string      `json:"commitSha,omitempty"`
	ProjectName string      `json:"projectName"`
	ConfigMap   string      `json:"configMap"`
	Image       string      `json:"image"`
	Ports       []Port      `json:"ports"`
	Autoscaling Autoscaling `json:"autoscaling,omitempty"`
}

// RevisionStatus defines the observed state of Revision
type RevisionStatus struct {
	Service           string `json:"service"`
	Deployment        string `json:"deployment"`
	State             string `json:"state"`
	TrafficPercentage uint   `json:"trafficPercentage,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Revision is the Schema for the revisions API
type Revision struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RevisionSpec   `json:"spec,omitempty"`
	Status RevisionStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// RevisionList contains a list of Revision
type RevisionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Revision `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Revision{}, &RevisionList{})
}
