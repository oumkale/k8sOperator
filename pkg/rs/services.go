package rs

import (
	"git.com.info/users/oumk/repos/podtatohead/api/v1alpha1"
	"k8s.io/apimachinery/pkg/util/intstr"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
)

// Returns the service object for the PodTatoHead
func Service(podtatohead *v1alpha1.PodTatoHead, name string, port int32, namespace string, scheme *runtime.Scheme) *corev1.Service {

	val := corev1.ServiceTypeClusterIP
	if name == "main" {
		val = corev1.ServiceTypeLoadBalancer
	}
	ser := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "podtato-" + name,
			Namespace: namespace,
			Labels: map[string]string{
				"app":  "podtato-head",
				"name": "podtato-" + name,
			},
		},
		//Status: corev1.ServiceStatus{},
		Spec: corev1.ServiceSpec{
			Selector: map[string]string{
				"app":  "podtato-head",
				"name": "podtato-" + name,
			},
			Type: val,
			Ports: []corev1.ServicePort{
				{
					Name: "http",
					TargetPort: intstr.IntOrString{
						Type:   intstr.Int,
						IntVal: 9000,
					},
					Port:     port,
					Protocol: "TCP",
				},
			},
		},
	}
	// Set PodtatoHead as the owner and controller
	controllerutil.SetControllerReference(podtatohead, ser, scheme)
	return ser
}
