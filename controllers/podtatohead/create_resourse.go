package podtatohead

import (
	"context"

	"git.com.info/users/oumk/repos/podtatohead/api/v1alpha1"
	"git.com.info/users/oumk/repos/podtatohead/pkg/rs"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
)

// Check if Service for the app exist, if not create one
func (r *PodtatoheadReconciler) createAppService(podtatohead *v1alpha1.PodTatoHead) error {
	service := &corev1.Service{}
	err := r.Client.Get(context.TODO(), types.NamespacedName{Name: "podtato-" + podtatohead.Spec.Main.Name, Namespace: podtatohead.Namespace}, service)
	if err != nil {
		if errors.IsNotFound(err) {
			if err := r.Client.Create(context.TODO(), rs.Service(podtatohead, podtatohead.Spec.Main.Name, podtatohead.Spec.Main.Port, podtatohead.Namespace, r.Scheme)); err != nil {
				return err
			}
		}
	}

	service = &corev1.Service{}
	err = r.Client.Get(context.TODO(), types.NamespacedName{Name: "podtato-" + podtatohead.Spec.Hats.Name, Namespace: podtatohead.Namespace}, service)
	if err != nil {
		if errors.IsNotFound(err) {
			if err := r.Client.Create(context.TODO(), rs.Service(podtatohead, podtatohead.Spec.Hats.Name, podtatohead.Spec.Hats.Port, podtatohead.Namespace, r.Scheme)); err != nil {
				return err
			}
		}
	}
	service = &corev1.Service{}
	err = r.Client.Get(context.TODO(), types.NamespacedName{Name: "podtato-" + podtatohead.Spec.LeftLeg.Name, Namespace: podtatohead.Namespace}, service)
	if err != nil {
		if errors.IsNotFound(err) {
			if err := r.Client.Create(context.TODO(), rs.Service(podtatohead, podtatohead.Spec.LeftLeg.Name, podtatohead.Spec.LeftLeg.Port, podtatohead.Namespace, r.Scheme)); err != nil {
				return err
			}
		}
	}

	service = &corev1.Service{}
	err = r.Client.Get(context.TODO(), types.NamespacedName{Name: "podtato-" + podtatohead.Spec.LeftArm.Name, Namespace: podtatohead.Namespace}, service)
	if err != nil {
		if errors.IsNotFound(err) {
			if err := r.Client.Create(context.TODO(), rs.Service(podtatohead, podtatohead.Spec.LeftArm.Name, podtatohead.Spec.LeftArm.Port, podtatohead.Namespace, r.Scheme)); err != nil {
				return err
			}
		}
	}

	service = &corev1.Service{}
	err = r.Client.Get(context.TODO(), types.NamespacedName{Name: "podtato-" + podtatohead.Spec.RightLeg.Name, Namespace: podtatohead.Namespace}, service)
	if err != nil {
		if errors.IsNotFound(err) {
			if err := r.Client.Create(context.TODO(), rs.Service(podtatohead, podtatohead.Spec.RightLeg.Name, podtatohead.Spec.RightLeg.Port, podtatohead.Namespace, r.Scheme)); err != nil {
				return err
			}
		}
	}
	service = &corev1.Service{}
	err = r.Client.Get(context.TODO(), types.NamespacedName{Name: "podtato-" + podtatohead.Spec.RightArm.Name, Namespace: podtatohead.Namespace}, service)
	if err != nil {
		if errors.IsNotFound(err) {
			if err := r.Client.Create(context.TODO(), rs.Service(podtatohead, podtatohead.Spec.RightArm.Name, podtatohead.Spec.RightArm.Port, podtatohead.Namespace, r.Scheme)); err != nil {
				return err
			}
		}
	}

	// service = &corev1.Service{}
	// err = r.Client.Get(context.TODO(), types.NamespacedName{Name: "podtato-" + podtatohead.Spec.HatsNew.Name, Namespace: podtatohead.Namespace}, service)
	// if err != nil {
	// 	if errors.IsNotFound(err) {
	// 		if err := r.Client.Create(context.TODO(), rs.Service(podtatohead, podtatohead.Spec.HatsNew.Name, podtatohead.Spec.HatsNew.Port, podtatohead.Namespace, r.Scheme)); err != nil {
	// 			return err
	// 		}
	// 	}
	// }

	return nil
}

// Check if Deployment for the app exist, if not create one
func (r *PodtatoheadReconciler) createAppDeploy(podtatohead *v1alpha1.PodTatoHead) error {

	ins := &appsv1.Deployment{}
	err := r.Client.Get(context.TODO(), types.NamespacedName{Name: "podtato-" + podtatohead.Spec.Main.Name, Namespace: podtatohead.Namespace}, ins)
	if err != nil {
		if errors.IsNotFound(err) {
			if err := r.Client.Create(context.TODO(), rs.Deployment(podtatohead, podtatohead.Spec.Main.Name, podtatohead.Spec.Main.Image, podtatohead.Namespace, podtatohead.Spec.ENV, r.Scheme)); err != nil {
				return err
			}
		}
	}

	ins = &appsv1.Deployment{}
	err = r.Client.Get(context.TODO(), types.NamespacedName{Name: "podtato-" + podtatohead.Spec.Hats.Name, Namespace: podtatohead.Namespace}, ins)
	if err != nil {
		if errors.IsNotFound(err) {
			if err := r.Client.Create(context.TODO(), rs.Deployment(podtatohead, podtatohead.Spec.Hats.Name, podtatohead.Spec.Hats.Image, podtatohead.Namespace, podtatohead.Spec.ENV, r.Scheme)); err != nil {
				return err
			}
		}
	}
	ins = &appsv1.Deployment{}
	err = r.Client.Get(context.TODO(), types.NamespacedName{Name: "podtato-" + podtatohead.Spec.LeftLeg.Name, Namespace: podtatohead.Namespace}, ins)
	if err != nil {
		if errors.IsNotFound(err) {
			if err := r.Client.Create(context.TODO(), rs.Deployment(podtatohead, podtatohead.Spec.LeftLeg.Name, podtatohead.Spec.LeftLeg.Image, podtatohead.Namespace, podtatohead.Spec.ENV, r.Scheme)); err != nil {
				return err
			}
		}
	}
	ins = &appsv1.Deployment{}
	err = r.Client.Get(context.TODO(), types.NamespacedName{Name: "podtato-" + podtatohead.Spec.LeftArm.Name, Namespace: podtatohead.Namespace}, ins)
	if err != nil {
		if errors.IsNotFound(err) {
			if err := r.Client.Create(context.TODO(), rs.Deployment(podtatohead, podtatohead.Spec.LeftArm.Name, podtatohead.Spec.LeftArm.Image, podtatohead.Namespace, podtatohead.Spec.ENV, r.Scheme)); err != nil {
				return err
			}
		}
	}
	ins = &appsv1.Deployment{}
	err = r.Client.Get(context.TODO(), types.NamespacedName{Name: "podtato-" + podtatohead.Spec.RightLeg.Name, Namespace: podtatohead.Namespace}, ins)
	if err != nil {
		if errors.IsNotFound(err) {
			if err := r.Client.Create(context.TODO(), rs.Deployment(podtatohead, podtatohead.Spec.RightLeg.Name, podtatohead.Spec.RightLeg.Image, podtatohead.Namespace, podtatohead.Spec.ENV, r.Scheme)); err != nil {
				return err
			}
		}
	}
	ins = &appsv1.Deployment{}
	err = r.Client.Get(context.TODO(), types.NamespacedName{Name: "podtato-" + podtatohead.Spec.RightArm.Name, Namespace: podtatohead.Namespace}, ins)
	if err != nil {
		if errors.IsNotFound(err) {
			if err := r.Client.Create(context.TODO(), rs.Deployment(podtatohead, podtatohead.Spec.RightArm.Name, podtatohead.Spec.RightArm.Image, podtatohead.Namespace, podtatohead.Spec.ENV, r.Scheme)); err != nil {
				return err
			}
		}
	}

	// ins = &appsv1.Deployment{}
	// err = r.Client.Get(context.TODO(), types.NamespacedName{Name: "podtato-" + podtatohead.Spec.HatsNew.Name, Namespace: podtatohead.Namespace}, ins)
	// if err != nil {
	// 	if errors.IsNotFound(err) {
	// 		if err := r.Client.Create(context.TODO(), rs.Deployment(podtatohead, podtatohead.Spec.HatsNew.Name, podtatohead.Spec.HatsNew.Image, podtatohead.Namespace, podtatohead.Spec.ENV, r.Scheme)); err != nil {
	// 			return err
	// 		}
	// 	}
	// }

	return nil
}
