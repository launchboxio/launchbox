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

type Metrics struct {
	Enabled bool `json:"enabled"`
}

type Logs struct {
	Enabled bool `json:"enabled"`
}

type Ingress struct {
	Enabled bool `json:"enabled"`
}

type Port struct {
	Protocol   string `json:"protocol,omitempty"`
	Port       int32  `json:"port"`
	Name       string `json:"name"`
	TargetPort string `json:"targetPort,omitempty"`
}

type ActiveRevisionStatus struct {
	RevisionId        string `json:"revisionId"`
	Status            string `json:"status,omitempty"`
	Replicas          uint32 `json:"replicas,omitempty"`
	TrafficPercentage uint   `json:"trafficPercentage,omitempty"`
}

// ProjectSpec defines the desired state of Project
type ProjectSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of Project. Edit project_types.go to remove/update
	Branch        string  `json:"branch,omitempty"`
	Repo          string  `json:"repo,omitempty"`
	Name          string  `json:"name,omitempty"`
	ProjectId     uint    `json:"projectId"`
	ApplicationId uint    `json:"applicationId"`
	Metrics       Metrics `json:"metrics,omitempty"`
	Logs          Logs    `json:"logs,omitempty"`
	Ports         []Port  `json:"ports,omitempty"`
	Ingress       Ingress `json:"ingress,omitempty"`
}

// ProjectStatus defines the observed state of Project
type ProjectStatus struct {
	ServiceAccount  string                 `json:"serviceAccount"`
	RootService     string                 `json:"rootService,omitempty"`
	PrimaryRevision string                 `json:"primaryRevision,omitempty"`
	ActiveRevisions []ActiveRevisionStatus `json:"activeRevisions,omitempty"`
	State           string                 `json:"state,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Project is the Schema for the projects API
type Project struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ProjectSpec   `json:"spec,omitempty"`
	Status ProjectStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ProjectList contains a list of Project
type ProjectList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Project `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Project{}, &ProjectList{})
}
