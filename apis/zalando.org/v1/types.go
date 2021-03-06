// +k8s:deepcopy-gen=package

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:deepcopy-gen=true
type RouteGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RouteGroupSpec   `json:"spec"`
	Status RouteGroupStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:deepcopy-gen=true
type RouteGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []RouteGroup `json:"items"`
}

// +k8s:deepcopy-gen=true
type RouteGroupSpec struct {
	Hosts           []string                     `json:"hosts,omitempty"`
	Backends        []RouteGroupBackend          `json:"backends"`
	DefaultBackends []RouteGroupBackendReference `json:"defaultBackends,omitempty"`
	Routes          []RouteGroupRouteSpec        `json:"routes,omitempty"`
}

// +k8s:deepcopy-gen=true
type RouteGroupBackend struct {
	// Name is the BackendName that can be referenced as RouteGroupBackendReference
	Name string `json:"name"`
	// Type is one of "service|shunt|loopback|dynamic|lb|network"
	Type string `json:"type"`
	// Address is required for Type network
	Address string `json:"address,omitempty"`
	// Algorithm is required for Type lb
	Algorithm string `json:"algorithm,omitempty"`
	// Endpoints is required for Type lb
	Endpoints []string `json:"endpoints,omitempty"`
	// ServiceName is required for Type service
	ServiceName string `json:"serviceName,omitempty"`
	// ServicePort is required for Type service
	ServicePort int `json:"servicePort,omitempty"`
}

// +k8s:deepcopy-gen=true
type RouteGroupBackendReference struct {
	// BackendName references the skipperBackend by name
	BackendName string `json:"backendName"`
	// Weight defines the traffic weight, if there are 2 or more
	// default backends
	Weight int `json:"weight"`
}

// +k8s:deepcopy-gen=true
type RouteGroupRouteSpec struct {
	// Path specifies Path predicate, only one of Path or PathSubtree is allowed
	Path string `json:"path,omitempty"`

	// PathSubtree specifies PathSubtree predicate, only one of Path or PathSubtree is allowed
	PathSubtree string `json:"pathSubtree,omitempty"`

	// PathRegexp can be added additionally
	PathRegexp string `json:"pathRegexp,omitempty"`

	// RouteGroupBackendReference specifies the list of backendReference that should
	// be applied to override the defaultBackends
	Backends []RouteGroupBackendReference `json:"backends,omitempty"`

	// Filters specifies the list of filters applied to the routeSpec
	Filters []string `json:"filters,omitempty"`

	// Predicates specifies the list of predicates applied to the routeSpec
	Predicates []string `json:"predicates,omitempty"`

	// Methods defines valid HTTP methods for the specified routeSpec
	Methods []string `json:"methods,omitempty"`
}

// +k8s:deepcopy-gen=true
type RouteGroupStatus struct {
	// LoadBalancer is similar to ingress status, such that
	// external-dns has the same style as in ingress
	LoadBalancer RouteGroupLoadBalancerStatus `json:"loadBalancer"`
}

// +k8s:deepcopy-gen=true
type RouteGroupLoadBalancerStatus struct {
	// RouteGroup is similar to Ingress in ingress status.LoadBalancer.
	RouteGroup []RouteGroupLoadBalancer `json:"routegroup"`
}

// +k8s:deepcopy-gen=true
type RouteGroupLoadBalancer struct {
	// IP is the IP address of the load balancer and is empty if Hostname is set
	IP string `json:"ip,omitempty"`
	// Hostname is the hostname of the load balancer and is empty if IP is set
	Hostname string `json:"hostname,omitempty"`
}
