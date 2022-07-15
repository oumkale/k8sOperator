package podtatohead

import (
	"context"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	"git.com.info/users/oumk/repos/podtatohead/api/v1alpha1"
	crdv1alpha1 "git.com.info/users/oumk/repos/podtatohead/api/v1alpha1"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

type Versions struct {
	version    string
	newVersion string
}

// PodtatoheadReconciler reconciles a PodTatoHead object
type PodtatoheadReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=crd.demo.com,resources=podtatoheads,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=crd.demo.com,resources=podtatoheads/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=crd.demo.com,resources=podtatoheads/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the PodTatoHead object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.10.0/pkg/reconcile
func (r *PodtatoheadReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {

	reqLogger := log.FromContext(ctx)
	reqLogger.Info("Reconciling PodTatoHead ...")

	// TODO(user): your logic here

	result := ctrl.Result{}

	// Get the Artifactory instance
	podtatohead := &v1alpha1.PodTatoHead{}
	err := r.Get(context.TODO(), req.NamespacedName, podtatohead)
	if err != nil {
		if errors.IsNotFound(err) {
			return result, nil
		}
		return result, err
	}
	if err := r.createResources(podtatohead, req); err != nil {
		reqLogger.Error(err, "Failed to create the resource")
		return reconcile.Result{}, err
	}

	if err := r.manageResources(podtatohead, r.Client); err != nil {
		reqLogger.Error(err, "Failed to manage resource required for the PodTatoHead CR")
		return reconcile.Result{}, err
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *PodtatoheadReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&crdv1alpha1.PodTatoHead{}).
		Owns(&corev1.Service{}).
		Owns(&appsv1.Deployment{}).
		Owns(&corev1.PersistentVolumeClaim{}).
		Owns(&corev1.ServiceAccount{}).
		Owns(&corev1.ConfigMap{}).
		Owns(&corev1.Secret{}).
		Complete(r)
}
