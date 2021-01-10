package v1alpha1

import (
	"context"
	"github.com/tanalam2411/olms/pkg/api/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"
	scheme "k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"time"
)

type OLMSsGetter interface {
	OLMSs(namespace string) OLMSInterface
}

type OLMSInterface interface {
	Create(ctx context.Context, olms *v1alpha1.OLMS, opts metav1.CreateOptions) (*v1alpha1.OLMS, error)
	Update(ctx context.Context, olms *v1alpha1.OLMS, opts metav1.UpdateOptions) (*v1alpha1.OLMS, error)
	UpdateStatus(ctx context.Context, olms *v1alpha1.OLMS, opts metav1.UpdateOptions) (*v1alpha1.OLMS, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1alpha1.OLMS, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1alpha1.OLMSList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1alpha1.OLMS, err error)

	OLMSExpansion
}

// olmss implements OLMS
type olmss struct {
	client rest.Interface
	ns     string
}

func newOLMSs(c *OLMSsV1alpha1Client, namespace string) *olmss {
	return &olmss{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

func (c *olmss) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1alpha1.OLMS, err error) {
	result = &v1alpha1.OLMS{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("olms").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

func (c *olmss) List(ctx context.Context, opts metav1.ListOptions) (result *v1alpha1.OLMSList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.OLMSList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("olms").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

func (c *olmss) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("olms").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

func (c *olmss) Create(ctx context.Context, olms *v1alpha1.OLMS, opts metav1.CreateOptions) (result *v1alpha1.OLMS, err error) {
	result = &v1alpha1.OLMS{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("olms").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(olms).
		Do(ctx).
		Into(result)
	return
}

func (c *olmss) Update(ctx context.Context, olms *v1alpha1.OLMS, opts metav1.UpdateOptions) (result *v1alpha1.OLMS, err error) {
	result = &v1alpha1.OLMS{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("olms").
		Name(olms.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(olms).
		Do(ctx).
		Into(result)
	return
}

func (c *olmss) UpdateStatus(ctx context.Context, olms *v1alpha1.OLMS, opts metav1.UpdateOptions) (result *v1alpha1.OLMS, err error) {
	result = &v1alpha1.OLMS{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("olms").
		Name(olms.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(olms).
		Do(ctx).
		Into(result)
	return
}

func (c *olmss) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("olms").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

func (c *olmss) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("olms").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

func (c *olmss) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1alpha1.OLMS, err error) {
	result = &v1alpha1.OLMS{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("olms").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
