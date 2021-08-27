/*


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

// PagerSpec defines the desired state of Pager
type PagerSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// message_id is the target resource's id to record
	MessageID string `json:"message_id,omitempty"`
	// message_name is the target resource's name to record
	MessageName string `json:"message_name,omitempty"`
	// message_type is the target resource's type to record
	MessageType string `json:"message_type,omitempty"`
	// message_content is the target resource's type to record
	MessageContent string `json:"message_content,omitempty"`
}

// PagerStatus defines the observed state of Pager
type PagerStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +genclient
// +kubebuilder:object:root=true

// Pager is the Schema for the pagers API
type Pager struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PagerSpec   `json:"spec,omitempty"`
	Status PagerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PagerList contains a list of Pager
type PagerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Pager `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Pager{}, &PagerList{})
}
