//lint:file-ignore ST1005 Generated code

// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"time"

	v1 "github.com/szuecs/routegroup-client/apis/zalando.org/v1"
	scheme "github.com/szuecs/routegroup-client/client/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// RouteGroupsGetter has a method to return a RouteGroupInterface.
// A group's client should implement this interface.
type RouteGroupsGetter interface {
	RouteGroups(namespace string) RouteGroupInterface
}

// RouteGroupInterface has methods to work with RouteGroup resources.
type RouteGroupInterface interface {
	Create(*v1.RouteGroup) (*v1.RouteGroup, error)
	Update(*v1.RouteGroup) (*v1.RouteGroup, error)
	UpdateStatus(*v1.RouteGroup) (*v1.RouteGroup, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error
	Get(name string, options metav1.GetOptions) (*v1.RouteGroup, error)
	List(opts metav1.ListOptions) (*v1.RouteGroupList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.RouteGroup, err error)
	RouteGroupExpansion
}

// routeGroups implements RouteGroupInterface
type routeGroups struct {
	client rest.Interface
	ns     string
}

// newRouteGroups returns a RouteGroups
func newRouteGroups(c *ZalandoV1Client, namespace string) *routeGroups {
	return &routeGroups{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the routeGroup, and returns the corresponding routeGroup object, and an error if there is any.
func (c *routeGroups) Get(name string, options metav1.GetOptions) (result *v1.RouteGroup, err error) {
	result = &v1.RouteGroup{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("routegroups").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of RouteGroups that match those selectors.
func (c *routeGroups) List(opts metav1.ListOptions) (result *v1.RouteGroupList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.RouteGroupList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("routegroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested routeGroups.
func (c *routeGroups) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("routegroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a routeGroup and creates it.  Returns the server's representation of the routeGroup, and an error, if there is any.
func (c *routeGroups) Create(routeGroup *v1.RouteGroup) (result *v1.RouteGroup, err error) {
	result = &v1.RouteGroup{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("routegroups").
		Body(routeGroup).
		Do().
		Into(result)
	return
}

// Update takes the representation of a routeGroup and updates it. Returns the server's representation of the routeGroup, and an error, if there is any.
func (c *routeGroups) Update(routeGroup *v1.RouteGroup) (result *v1.RouteGroup, err error) {
	result = &v1.RouteGroup{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("routegroups").
		Name(routeGroup.Name).
		Body(routeGroup).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *routeGroups) UpdateStatus(routeGroup *v1.RouteGroup) (result *v1.RouteGroup, err error) {
	result = &v1.RouteGroup{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("routegroups").
		Name(routeGroup.Name).
		SubResource("status").
		Body(routeGroup).
		Do().
		Into(result)
	return
}

// Delete takes name of the routeGroup and deletes it. Returns an error if one occurs.
func (c *routeGroups) Delete(name string, options *metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("routegroups").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *routeGroups) DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("routegroups").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched routeGroup.
func (c *routeGroups) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.RouteGroup, err error) {
	result = &v1.RouteGroup{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("routegroups").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
