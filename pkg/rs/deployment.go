package rs

import (
	"git.com.info/users/oumk/repos/podtatohead/api/v1alpha1"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
)

//Deployment returns the Deployment object for the PodTatoHead
func Deployment(podtatohead *v1alpha1.PodTatoHead, name, image, namespace string, env []corev1.EnvVar, scheme *runtime.Scheme) *appsv1.Deployment {

	if image == "" {
		image = "kaleoum/podtato-main:v0.1.1"
	}

	dep := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "podtato-" + name,
			Namespace: namespace,
			Labels: map[string]string{
				"app":  "podtato-head",
				"name": "podtato-" + name,
			},
		},
		Spec: appsv1.DeploymentSpec{
			//Replicas: &stsObject.replicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app":  "podtato-head",
					"name": "podtato-" + name,
				},
			},

			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app":  "podtato-head",
						"name": "podtato-" + name,
					},
				},
				Spec: corev1.PodSpec{

					Containers: []corev1.Container{
						{
							Image:           image,
							Name:            name,
							ImagePullPolicy: "IfNotPresent",
							Ports: []corev1.ContainerPort{
								{
									ContainerPort: 9000,
								},
							},
							Env: env,
						},
					},
				},
			},
		},
	}
	controllerutil.SetControllerReference(podtatohead, dep, scheme)
	return dep
}
