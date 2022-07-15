package podtatohead

import (
	"context"
	"fmt"

	"git.com.info/users/oumk/repos/podtatohead/api/v1alpha1"
	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

// manageResources will ensure that the resource are with the expected values in the cluster
func (r *PodtatoheadReconciler) manageResources(podtatohead *v1alpha1.PodTatoHead, client client.Client) error {

	fmt.Println("PodTatoHead Version", podtatohead.Spec.Version)
	// get the latest version of podtatohead deployment
	dep := &appsv1.Deployment{}
	err := client.Get(context.TODO(), types.NamespacedName{Name: "podtato-" + podtatohead.Spec.Main.Name, Namespace: podtatohead.Namespace}, dep)
	if err != nil {
		if errors.IsNotFound(err) {
			return nil
		}
		return err
	}

	r.ensureDepSize(podtatohead, dep, client, "podtato-"+podtatohead.Spec.Main.Name, podtatohead.Spec.Main.Image)

	// get the latest version of podtatohead deployment
	dep = &appsv1.Deployment{}
	err = client.Get(context.TODO(), types.NamespacedName{Name: "podtato-" + podtatohead.Spec.Hats.Name, Namespace: podtatohead.Namespace}, dep)
	if err != nil {
		if errors.IsNotFound(err) {
			return nil
		}
		return err
	}
	r.ensureDepSize(podtatohead, dep, client, "podtato-"+podtatohead.Spec.Hats.Name, podtatohead.Spec.Hats.Image)

	// get the latest version of podtatohead deployment
	dep = &appsv1.Deployment{}
	err = client.Get(context.TODO(), types.NamespacedName{Name: "podtato-" + podtatohead.Spec.LeftLeg.Name, Namespace: podtatohead.Namespace}, dep)
	if err != nil {
		if errors.IsNotFound(err) {
			return nil
		}
		return err
	}
	r.ensureDepSize(podtatohead, dep, client, "podtato-"+podtatohead.Spec.LeftLeg.Name, podtatohead.Spec.LeftLeg.Image)

	dep = &appsv1.Deployment{}
	err = client.Get(context.TODO(), types.NamespacedName{Name: "podtato-" + podtatohead.Spec.LeftArm.Name, Namespace: podtatohead.Namespace}, dep)
	if err != nil {
		if errors.IsNotFound(err) {
			return nil
		}
		return err
	}

	r.ensureDepSize(podtatohead, dep, client, "podtato-"+podtatohead.Spec.LeftArm.Name, podtatohead.Spec.LeftArm.Image)

	// get the latest version of podtatohead deployment
	dep = &appsv1.Deployment{}
	err = client.Get(context.TODO(), types.NamespacedName{Name: "podtato-" + podtatohead.Spec.RightLeg.Name, Namespace: podtatohead.Namespace}, dep)
	if err != nil {
		if errors.IsNotFound(err) {
			return nil
		}
		return err
	}
	r.ensureDepSize(podtatohead, dep, client, "podtato-"+podtatohead.Spec.RightLeg.Name, podtatohead.Spec.RightLeg.Image)

	// get the latest version of podtatohead deployment
	dep = &appsv1.Deployment{}
	err = client.Get(context.TODO(), types.NamespacedName{Name: "podtato-" + podtatohead.Spec.RightArm.Name, Namespace: podtatohead.Namespace}, dep)
	if err != nil {
		if errors.IsNotFound(err) {
			return nil
		}
		return err
	}
	r.ensureDepSize(podtatohead, dep, client, "podtato-"+podtatohead.Spec.RightArm.Name, podtatohead.Spec.RightArm.Image)

	return nil
}

// ensureDepSize will ensure that the quanity of instances in the cluster for the Database deployment is the same defined in the CR
func (r *PodtatoheadReconciler) ensureDepSize(podtatohead *v1alpha1.PodTatoHead, dep *v1.Deployment, client client.Client, name, image string) error {
	reqLogger := log.FromContext(context.TODO())

	size := podtatohead.Spec.ReplicaCount
	if *dep.Spec.Replicas != size {
		// Set the number of Replica cas spec in the CR
		dep.Spec.Replicas = &size
		reqLogger.Info("Updating PodTatoHead object to desired state")
		if err := client.Update(context.TODO(), dep); err != nil {
			return err
		}
	}

	if *&dep.Spec.Template.Spec.Containers[0].Image != image {
		// Set the number of Replica cas spec in the CR
		dep.Spec.Replicas = &size
		dep.Spec.Template.Spec.Containers[0].Image = image
		reqLogger.Info("Updating PodTatoHead object to desired state")
		if err := client.Update(context.TODO(), dep); err != nil {
			return err
		}
	}
	return nil
}

//createResources will create the secondary resource which are required in order to make works successfully the primary resource(CR)
func (r *PodtatoheadReconciler) createResources(podtatohead *v1alpha1.PodTatoHead, request reconcile.Request) error {
	reqLogger := log.FromContext(context.TODO())

	// Check if Statefulset for the app exist, if not create one
	if err := r.createAppDeploy(podtatohead); err != nil {
		reqLogger.Error(err, "Failed to create Statefulsets")
		return err
	}

	// Check if service for the app exist, if not create one
	if err := r.createAppService(podtatohead); err != nil {
		reqLogger.Error(err, "Failed to create Service")
		return err
	}

	return nil
}
