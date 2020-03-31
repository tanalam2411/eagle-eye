package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ReachabilityTestSpec defines the desired state of ReachabilityTest
type ReachabilityTestSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html

	Layer        string                `json:"layer"`
	Source       SourceEndpoint        `json:"source"`
	Destinations []DestinationEndpoint `json:"destinations"`
	Frequency    string                `json:"frequency,omitempty"`
}

type SourceEndpoint struct {
	Name           string                `json:"name,omitempty"`
	Namespace      string                `json:"namespace"`
	Container      string                `json:"container"`
	SourceSelector *metav1.LabelSelector `json:"sourceSelector,omitempty"`
}

type EndpointKind string

const (
	IP       EndpointKind = "ip"
	Pod      EndpointKind = "pod"
	Service  EndpointKind = "service"
	Selector EndpointKind = "selector"
)

type DestinationEndpoint struct {
	Kind      EndpointKind `json:"kind"`
	Name      string       `json:"name,omitempty"`
	Namespace string       `json:"namespace,omitempty"`
	IP        string       `json:"ip,omitempty"`
	Port      string       `json:"port,omitempty"`
}

// ReachabilityTestStatus defines the observed state of ReachabilityTest
type ReachabilityTestStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html

	TestStatus *runtime.RawExtension `json:"testStatus,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ReachabilityTest is the Schema for the reachabilitytests API
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=reachabilitytests,scope=Cluster
// +kubebuilder:storageversion
type ReachabilityTest struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ReachabilityTestSpec   `json:"spec,omitempty"`
	Status ReachabilityTestStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ReachabilityTestList contains a list of ReachabilityTest
type ReachabilityTestList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ReachabilityTest `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ReachabilityTest{}, &ReachabilityTestList{})
}
