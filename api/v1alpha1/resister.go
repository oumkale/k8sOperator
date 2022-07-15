package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

const GroupName = "crd.demo.com"

func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(GroupVersion,
		&PodTatoHead{},
		&PodTatoHeadList{},
	)

	metav1.AddToGroupVersion(scheme, GroupVersion)
	return nil
}
