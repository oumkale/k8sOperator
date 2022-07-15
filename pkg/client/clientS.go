package client

import (
	"git.com.info/users/oumk/repos/podtatohead/api/v1alpha1"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
)

type ExampleV1Alpha1Interface interface {
	PodTatoHead(namespace string) PodTatoHeadInterface
}

type PodTatoHeadClient struct {
	restClient rest.Interface
}

func NewForConfig(c *rest.Config) (*PodTatoHeadClient, error) {
	config := *c
	config.ContentConfig.GroupVersion = &v1alpha1.GroupVersion
	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()
	config.UserAgent = rest.DefaultKubernetesUserAgent()

	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}

	return &PodTatoHeadClient{restClient: client}, nil
}

func (c *PodTatoHeadClient) PodTatoHead(namespace string) PodTatoHeadInterface {
	return &PodTatoHeadReconciler{
		restClient: c.restClient,
	}
}
