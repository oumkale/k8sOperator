package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// PodTatoHeadSpec defines the desired state of PodTatoHead
type PodTatoHeadSpec struct {
	Name         string          `json:"name,omitempty"`
	Namespace    string          `json:"namespace,omitempty"`
	Version      string          `json:"version,omitempty"`
	ReplicaCount int32           `json:"replicaCount,omitempty"`
	Main         MainInfo        `json:"main,omitempty"`
	Hats         HatsInfo        `json:"hats,omitempty"`
	LeftLeg      LeftLegInfo     `json:"leftLeg,omitempty"`
	LeftArm      LeftArmInfo     `json:"leftArm,omitempty"`
	RightLeg     RightLegInfo    `json:"rightLeg,omitempty"`
	RightArm     RightArmInfo    `json:"rightArm,omitempty"`
	ENV          []corev1.EnvVar `json:"env,omitempty"`
	Port         int             `json:"port,omitempty"`
}

// MainInfo defines the observed state of CopySystemYaml
type MainInfo struct {
	Image       string `json:"image,omitempty"`
	Name        string `json:"name,omitempty"`
	Port        int32  `json:"port,omitempty"`
	ServiceType string `json:"serviceType,omitempty"`
}

// LeftLegInfo defines the observed state of WaitForDb
type LeftLegInfo struct {
	Image       string `json:"image,omitempty"`
	Name        string `json:"name,omitempty"`
	Port        int32  `json:"port,omitempty"`
	ServiceType string `json:"serviceType,omitempty"`
}

// HatsInfo defines the observed state of WaitForDb
type HatsInfo struct {
	Image       string `json:"image,omitempty"`
	Name        string `json:"name,omitempty"`
	Port        int32  `json:"port,omitempty"`
	ServiceType string `json:"serviceType,omitempty"`
}

// LeftArmInfo defines the observed state of WaitForDb
type LeftArmInfo struct {
	Image       string `json:"image,omitempty"`
	Name        string `json:"name,omitempty"`
	Port        int32  `json:"port,omitempty"`
	ServiceType string `json:"serviceType,omitempty"`
}

// RightLegInfo defines the observed state of Marshaller
type RightLegInfo struct {
	Image       string `json:"image,omitempty"`
	Name        string `json:"name,omitempty"`
	Port        int32  `json:"port,omitempty"`
	ServiceType string `json:"serviceType,omitempty"`
}

type RightArmInfo struct {
	Image       string `json:"image,omitempty"`
	Name        string `json:"name,omitempty"`
	Port        int32  `json:"port,omitempty"`
	ServiceType string `json:"serviceType,omitempty"`
}

// PodTatoHeadStatus defines the observed state of PodTatoHead
type PodTatoHeadStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// PodTatoHead is the Schema for the PodTatoHeads API
type PodTatoHead struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PodTatoHeadSpec   `json:"spec,omitempty"`
	Status PodTatoHeadStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// PodTatoHeadList contains a list of PodTatoHead
type PodTatoHeadList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PodTatoHead `json:"items"`
}

// ConfigMap is an simpler implementation of corev1.ConfigMaps, needed for experiments
type ConfigMap struct {
	Data      map[string]string `json:"data,omitempty"`
	Name      string            `json:"name"`
	MountPath string            `json:"mountPath"`
}

// Secret is an simpler implementation of corev1.Secret
type Secret struct {
	Name      string `json:"name"`
	MountPath string `json:"mountPath"`
}

func init() {
	SchemeBuilder.Register(&PodTatoHead{}, &PodTatoHeadList{})
}
