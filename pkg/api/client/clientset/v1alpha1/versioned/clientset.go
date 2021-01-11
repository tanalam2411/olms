package versioned

import (
	"fmt"
	olmssv1alpha1 "github.com/tanalam2411/olms/pkg/api/client/clientset/v1alpha1/versioned/typed/olmss/v1alpha1"
	"k8s.io/client-go/discovery"
	rest "k8s.io/client-go/rest"
	flowcontrol "k8s.io/client-go/util/flowcontrol"
)

type Interface interface{
	Discovery() discovery.DiscoveryInterface
	OLMSsV1alpha1() olmssv1alpha1.OLMSV1alpha1Interface
}


type Clientset struct {
	*discovery.DiscoveryClient
	olmssV1alpha1 *olmssv1alpha1.OLMSsV1alpha1Client
}

func (c *Clientset) OLMSsV1alpha1() olmssv1alpha1.OLMSV1alpha1Interface {
	return c.olmssV1alpha1
}


func NewForConfig(c *rest.Config) (*Clientset, error) {
	configShallowCopy := *c
	if configShallowCopy.RateLimiter == nil && configShallowCopy.QPS > 0 {
		if configShallowCopy.Burst <= 0 {
			return nil, fmt.Errorf("burst is required to be greater than 0 when RateLimiter is not set and QPS is set to greater than 0")
		}
		configShallowCopy.RateLimiter = flowcontrol.NewTokenBucketRateLimiter(configShallowCopy.QPS, configShallowCopy.Burst)
	}
	var cs Clientset
	var err error
	cs.olmssV1alpha1, err = olmssv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}

	cs.DiscoveryClient, err = discovery.NewDiscoveryClientForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	return &cs, nil
}


func NewForConfigOrDie(c *rest.Config) *Clientset {
	var cs Clientset
	cs.olmssV1alpha1 = olmssv1alpha1.NewForConfigOrDie(c)

	cs.DiscoveryClient = discovery.NewDiscoveryClientForConfigOrDie(c)
	return &cs
}


func New(c rest.Interface) *Clientset {
	var cs Clientset
	cs.olmssV1alpha1 = olmssv1alpha1.New(c)
	
	cs.DiscoveryClient = discovery.NewDiscoveryClient(c)
	return &cs
}