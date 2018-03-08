package v1

import (
	scheme "github.com/kubermatic/kubermatic/api/pkg/crd/client/clientset/versioned/scheme"
	v1 "github.com/kubermatic/kubermatic/api/pkg/crd/prometheus/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// PrometheusesGetter has a method to return a PrometheusInterface.
// A group's client should implement this interface.
type PrometheusesGetter interface {
	Prometheuses(namespace string) PrometheusInterface
}

// PrometheusInterface has methods to work with Prometheus resources.
type PrometheusInterface interface {
	Create(*v1.Prometheus) (*v1.Prometheus, error)
	Update(*v1.Prometheus) (*v1.Prometheus, error)
	UpdateStatus(*v1.Prometheus) (*v1.Prometheus, error)
	Delete(name string, options *meta_v1.DeleteOptions) error
	DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error
	Get(name string, options meta_v1.GetOptions) (*v1.Prometheus, error)
	List(opts meta_v1.ListOptions) (*v1.PrometheusList, error)
	Watch(opts meta_v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.Prometheus, err error)
	PrometheusExpansion
}

// prometheuses implements PrometheusInterface
type prometheuses struct {
	client rest.Interface
	ns     string
}

// newPrometheuses returns a Prometheuses
func newPrometheuses(c *MonitoringV1Client, namespace string) *prometheuses {
	return &prometheuses{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the prometheus, and returns the corresponding prometheus object, and an error if there is any.
func (c *prometheuses) Get(name string, options meta_v1.GetOptions) (result *v1.Prometheus, err error) {
	result = &v1.Prometheus{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("prometheuses").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Prometheuses that match those selectors.
func (c *prometheuses) List(opts meta_v1.ListOptions) (result *v1.PrometheusList, err error) {
	result = &v1.PrometheusList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("prometheuses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested prometheuses.
func (c *prometheuses) Watch(opts meta_v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("prometheuses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a prometheus and creates it.  Returns the server's representation of the prometheus, and an error, if there is any.
func (c *prometheuses) Create(prometheus *v1.Prometheus) (result *v1.Prometheus, err error) {
	result = &v1.Prometheus{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("prometheuses").
		Body(prometheus).
		Do().
		Into(result)
	return
}

// Update takes the representation of a prometheus and updates it. Returns the server's representation of the prometheus, and an error, if there is any.
func (c *prometheuses) Update(prometheus *v1.Prometheus) (result *v1.Prometheus, err error) {
	result = &v1.Prometheus{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("prometheuses").
		Name(prometheus.Name).
		Body(prometheus).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *prometheuses) UpdateStatus(prometheus *v1.Prometheus) (result *v1.Prometheus, err error) {
	result = &v1.Prometheus{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("prometheuses").
		Name(prometheus.Name).
		SubResource("status").
		Body(prometheus).
		Do().
		Into(result)
	return
}

// Delete takes name of the prometheus and deletes it. Returns an error if one occurs.
func (c *prometheuses) Delete(name string, options *meta_v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("prometheuses").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *prometheuses) DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("prometheuses").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched prometheus.
func (c *prometheuses) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.Prometheus, err error) {
	result = &v1.Prometheus{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("prometheuses").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
