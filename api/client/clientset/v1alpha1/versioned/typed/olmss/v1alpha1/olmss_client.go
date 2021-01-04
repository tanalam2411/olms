package v1alpha1

import (
	"k8s.io/api/node/v1alpha1"
	"k8s.io/client-go/kubernetes/scheme"
	rest "k8s.io/client-go/rest"
)

type OLMSV1alpha1Interface interface {
	RESTClient() rest.Interface
	OLMSsGetter
}

type OLMSsV1alpha1Client struct {
	restClient rest.Interface
}

func (c *OLMSsV1alpha1Client) OLMSs(namespace string) OLMSInterface {
	return newOLMSs(c, namespace)
}

func NewForConfig(c *rest.Config) (*OLMSsV1alpha1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &OLMSsV1alpha1Client{client}, nil
}

func NewForConfigOrDie(c *rest.Config) *OLMSsV1alpha1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new OLMSsV1alpha1Client for the given RESTClient
func New(c rest.Interface) *OLMSsV1alpha1Client {
	return &OLMSsV1alpha1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1alpha1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

func (c *OLMSsV1alpha1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
