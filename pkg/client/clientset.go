package client

import (
	"context"
	"fmt"

	"git.com.info/users/oumk/repos/podtatohead/api/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
)

// PodTatoHeadReconciler reconciles a PodTatoHead object
type PodTatoHeadReconciler struct {
	Scheme     *runtime.Scheme
	restClient rest.Interface
}

var ctx = context.Background()

// PodTatoHeadInterface
type PodTatoHeadInterface interface {
	List(opts metav1.ListOptions) (*v1alpha1.PodTatoHeadList, error)
	Get(name string, options metav1.GetOptions) (*v1alpha1.PodTatoHead, error)
	Create(*v1alpha1.PodTatoHead) (*v1alpha1.PodTatoHead, error)
	Update(*v1alpha1.PodTatoHead) (*v1alpha1.PodTatoHead, error)
	//Watch(opts metav1.ListOptions) (watch.Interface, error)
	// ...
}

func (c *PodTatoHeadReconciler) List(opts metav1.ListOptions) (*v1alpha1.PodTatoHeadList, error) {
	result := v1alpha1.PodTatoHeadList{}
	err := c.restClient.Get().Namespace("oum").Resource("PodTatoHead").VersionedParams(&opts, scheme.ParameterCodec).Do(context.Background()).Into(&result)
	return &result, err
}

//Get return PodTatoHead object
func (c *PodTatoHeadReconciler) Get(name string, opts metav1.GetOptions) (*v1alpha1.PodTatoHead, error) {
	result := v1alpha1.PodTatoHead{}
	err := c.restClient.
		Get().
		Namespace("oum").
		Resource("podtatoheads").
		Name("podtatohead").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do(ctx).
		Into(&result)
	fmt.Println("Result:", err)
	return &result, err
}

//Create creates PodTatoHead object
func (c *PodTatoHeadReconciler) Create(project *v1alpha1.PodTatoHead) (*v1alpha1.PodTatoHead, error) {
	result := v1alpha1.PodTatoHead{}

	err := c.restClient.Post().
		Namespace("oum").
		Resource("podtatoheads").
		Body(project).
		Do(ctx).
		Into(&result)

	return &result, err
}

func (c *PodTatoHeadReconciler) Update(pr *v1alpha1.PodTatoHead) (*v1alpha1.PodTatoHead, error) {
	result := v1alpha1.PodTatoHead{}
	//pr.Status.Status = "Running"
	err := c.restClient.Put().Namespace("oum").Resource("podtatoheads").Name("podtatohead").Body(pr).Do(ctx).Into(&result)
	if err != nil {
		return pr, err
	}
	return &result, nil
}
