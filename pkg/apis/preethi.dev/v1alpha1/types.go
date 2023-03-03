/*
Copyright 2023.

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

// MyKindSpec defines the desired state of MyKind
type ExperimentSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Title                 string
	Description           string
	Version               string
	Tags                  []string
	Contributions         map[string]string
	Experimentlabel       string
	Secrets               interface{}
	Controls              []interface{}
	Configuration         map[string]interface{}
	Steadystatehypothesis SteadyStateSpec
	Method                []MethodSpec
	Rollback              []interface{}
}

/*type ConfigurationSpec struct {
    LabelSelector string
    NamePattern string
    Namespace string
    Qty int
    PodDeletion struct {
        Rand bool
    }
    Hypothesis struct {
        Strategy string
    }
    RollbackStrategy string
}*/

type SteadyStateSpec struct {
	Title  string
	Probes []Probe
}
type Probe struct {
	Name       string
	Type       string
	Tolerance  bool
	Background bool
	Provider   struct {
		Type      string
		Module    string
		Func      string
		Arguments map[string]interface{}
	}
}
type MethodSpec struct {
	Name       string
	Type       string
	Background bool
	Provider   struct {
		Type      string
		Module    string
		Func      string
		Arguments map[string]interface{}
	}
	Pauses struct {
		Before int
		After  int
	}
}

// MyKindStatus defines the observed state of MyKind
/*type ExperimentStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file


}*/

//+genclient
//+k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type Experiment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec ExperimentSpec `json:"spec,omitempty"`
	//Status ExperimentStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type MyExperimentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Experiment `json:"items"`
}
