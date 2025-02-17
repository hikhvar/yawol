package v1beta1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Annotation for settings in svc object
const (
	// ServiceImageID overwrite default imageID
	ServiceImageID = "yawol.stackit.cloud/imageId"
	// ServiceFlavorID overwrite default flavorID
	ServiceFlavorID = "yawol.stackit.cloud/flavorId"
	// AvailabilityZoneID set availability zone for specific service
	ServiceAvailabilityZone = "yawol.stackit.cloud/availabilityZone"
	// ServiceInternalLoadbalancer sets the internal flag in LB objects
	ServiceInternalLoadbalancer = "yawol.stackit.cloud/internalLB"
	// ServiceDebug set in lb object an debug setting
	ServiceDebug = "yawol.stackit.cloud/debug"
	// ServiceDebugSSHKey set an sshkey
	ServiceDebugSSHKey = "yawol.stackit.cloud/debugsshkey"
	// ServiceClassName for filtering services in cloud-controller
	ServiceClassName = "yawol.stackit.cloud/className"
	// ServiceReplicas for setting loadbalancer replicas in cloud-controller
	ServiceReplicas = "yawol.stackit.cloud/replicas"
	// ServiceTCPProxyProtocol enables the HAProxy TCP Proxy Protocol for all TCP connections
	ServiceTCPProxyProtocol = "yawol.stackit.cloud/tcpProxyProtocol"
	// ServiceTCPProxyProtocolPortsFilter enables for the specified ports (comma separated list)
	ServiceTCPProxyProtocolPortsFilter = "yawol.stackit.cloud/tcpProxyProtocolPortsFilter"
	// ServiceExistingFloatingIP enables usage of existing Floating IP
	ServiceExistingFloatingIP = "yawol.stackit.cloud/existingFloatingIP"
)

// +kubebuilder:object:root=true
// +kubebuilder:resource:shortName=lb
// +kubebuilder:subresource:status
// +kubebuilder:subresource:scale:specpath=.spec.replicas,statuspath=.status.replicas
// +kubebuilder:printcolumn:name="DESIRED",type=string,JSONPath=`.spec.replicas`
// +kubebuilder:printcolumn:name="CURRENT",type=string,JSONPath=`.status.replicas`
// +kubebuilder:printcolumn:name="READY",type=string,JSONPath=`.status.readyReplicas`
// +kubebuilder:printcolumn:name="externalIP",type=string,JSONPath=`.status.externalIP`
// +kubebuilder:printcolumn:name="AGE",type=date,JSONPath=`.metadata.creationTimestamp`

// LoadBalancer is the Schema for the YAWOL LoadBalancer API
type LoadBalancer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`

	Spec   LoadBalancerSpec   `json:"spec,omitempty"`
	Status LoadBalancerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LoadBalancerList contains a list of LoadBalancer.
type LoadBalancerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LoadBalancer `json:"items"`
}

// LoadBalancerSpec defines the desired state of LoadBalancer
type LoadBalancerSpec struct {
	// This label selector matches the load balancer sets deriving from the load balancer
	Selector metav1.LabelSelector `json:"selector"`
	// Replicas defines the number of LoadBalancers that should run.
	// +kubebuilder:default:=1
	// +kubebuilder:validation:Minimum:=0
	// +optional
	Replicas int `json:"replicas,omitempty"`
	// ExistingFloatingIP uses a existing Floating IP as FIP
	// +optional
	ExistingFloatingIP *string `json:"existingFloatingIP,omitempty"`
	// Debug are settings for debugging an loadbalancer.
	// +optional
	DebugSettings LoadBalancerDebugSettings `json:"debugSettings,omitempty"`
	// Endpoints defines the Endpoints for the LoadBalancer.
	Endpoints []LoadBalancerEndpoint `json:"endpoints,omitempty"`
	// Ports defines the Ports for the LoadBalancer (copy from service)
	Ports []corev1.ServicePort `json:"ports,omitempty"`
	// Infrastructure defines parameters for the Infrastructure
	Infrastructure LoadBalancerInfrastructure `json:"infrastructure"`
	// Options for additional LoadBalancer settings
	// +optional
	Options LoadBalancerOptions `json:"options,omitempty"`
}

type LoadBalancerOptions struct {
	// InternalLB is a bool for internal LoadBalancer. If set to false a FloatingIP will be assigned to the LB. Defaults to false.
	// +kubebuilder:default:=false
	// +optional
	InternalLB bool `json:"internalLB,omitempty"`
	// LoadBalancerSourceRanges restrict traffic to IP ranges for the LoadBalancer (copy from service)
	// +optional
	LoadBalancerSourceRanges []string `json:"loadBalancerSourceRanges,omitempty"`
	// TCPProxyProtocol enables HAProxy TCP Proxy Protocol
	// +optional
	TCPProxyProtocol bool `json:"tcpProxyProtocol,omitempty"`
	// TCPProxyProtocolPortList enables HAProxy TCP Proxy Protocol for specified ports.
	// If empty it is enabled for all ports. Only has an affect if TCPProxyProtocol is enabled.
	// +optional
	TCPProxyProtocolPortsFilter []int32 `json:"tcpProxyProtocolPortFilter,omitempty"`
}

// LoadBalancerDebugSettings defines debug settings for the LoadBalancer
type LoadBalancerDebugSettings struct {
	// Enabled defines if debugging is enabled
	// +optional
	Enabled bool `json:"enabled"`
	// SshKey is a openstack sshkey name for debugging
	// +optional
	SshkeyName string `json:"sshkeyName,omitempty"`
}

// LoadBalancerEndpoint defines a Endpoint for the LoadBalancer
type LoadBalancerEndpoint struct {
	// Name defines a name for the Endpoint (example: node name).
	Name string `json:"name"`
	// Addresses is a list of addresses for the endpoint, they can contain IPv4 and IPv6 addresses.
	Addresses []string `json:"addresses,omitempty"`
}

// LoadBalancerInfrastructure defines infrastructure defaults for the LoadBalancer
type LoadBalancerInfrastructure struct {
	// FloatingNetID defines a openstack ID for the floatingNet.
	// +optional
	FloatingNetID *string `json:"floatingNetID,omitempty"`
	// NetworkID defines a openstack ID for the network.
	NetworkID string `json:"networkID"`
	// Flavor defines openstack flavor for the LoadBalancer. Uses a default if not defined.
	// +optional
	Flavor *OpenstackFlavorRef `json:"flavor,omitempty"`
	// Image defines openstack image for the LoadBalancer. Uses a default if not defined.
	// +optional
	Image *OpenstackImageRef `json:"image,omitempty"`
	// AvailabilityZone defines the openstack availability zone for the LoadBalancer.
	// +optional
	AvailabilityZone string `json:"availabilityZone"`
	// AuthSecretRef defines a secretRef for the openstack secret.
	AuthSecretRef corev1.SecretReference `json:"authSecretRef"`
}

// OpenstackImageRef defines a reference to a Openstack image.
type OpenstackImageRef struct {
	// ImageID is the image ID used for requesting virtual machines.
	// +optional
	ImageID *string `json:"image_id,omitempty"`
	// ImageName is the name of the image used for requesting virtual machines.
	// ImageName is only used if ImageID is not defined.
	// +optional
	ImageName *string `json:"image_name,omitempty"`
	// ImageSearch is a search string to find the image used for requesting virtual machines.
	// Search will be performed in metadata of the images.
	// ImageSearch is only used if ImageName and ImageID are not defined.
	// +optional
	ImageSearch *string `json:"image_search,omitempty"`
}

// OpenstackFlavorRef defines a reference to a Openstack flavor.
// The Flavor defines the amount of cpu cores and memory as well as the size of the root disk.
type OpenstackFlavorRef struct {
	// FlavorID is the flavor ID used for requesting virtual machines.
	// +optional
	FlavorID *string `json:"flavor_id,omitempty"`
	// FlavorName is the name of the flavor used for requesting virtual machines.
	// FlavorName is only used if FlavorID is not defined.
	// +optional
	FlavorName *string `json:"flavor_name,omitempty"`
	// FlavorSearch is a search string to find the flavor used for requesting virtual machines.
	// Search will be performed in metadata of the flavors.
	// FlavorSearch is only used if FlavorName and FlavorID are not defined.
	// +optional
	FlavorSearch *string `json:"flavor_search,omitempty"`
}

// LoadBalancerRef defines a reference to a LoadBalancer object.
type LoadBalancerRef struct {
	// Name is unique within a namespace to reference a LoadBalancer resource.
	Name string `json:"name"`
	// Namespace defines the space within which the LoadBalancer name must be unique.
	Namespace string `json:"namespace"`
}

// LoadBalancerStatus defines the observed state of LoadBalancer.
type LoadBalancerStatus struct {
	// ReadyReplicas are the current running replicas.
	// +optional
	ReadyReplicas *int `json:"readyReplicas,omitempty"`
	// Replicas displays the running lb replicas under this deployment
	// +optional
	Replicas *int `json:"replicas,omitempty"`
	// ExternalIP is the current externalIP (FIP or private). If not defined, no ExternalIP is bound yet.
	// +optional
	ExternalIP *string `json:"externalIP,omitempty"`
	// FloatingID is the current openstack ID from the FloatingIP.
	// +optional
	FloatingID *string `json:"floatingID,omitempty"`
	// FloatingName is the current openstack name from the FloatingIP.
	// +optional
	FloatingName *string `json:"floatingName,omitempty"`
	// PortID is the current openstack ID from the virtual Port.
	// +optional
	PortID *string `json:"portID,omitempty"`
	// PortName is the current openstack name from the virtual Port.
	// +optional
	PortName *string `json:"portName,omitempty"`
	// SecurityGroupID is the current security group ID mapped to the port
	// +optional
	SecurityGroupID *string `json:"security_group_id,omitempty"`
	// SecurityGroupName is the current security group name mapped to the port
	// +optional
	SecurityGroupName *string `json:"security_group_name,omitempty"`
	// LastOpenstackReconcile contains the timestamp of the last openstack reconciliation.
	// +optional
	LastOpenstackReconcile *metav1.Time `json:"lastOpenstackReconcile,omitempty"`
	// OpenstackReconcileHash contains a hash of openstack related settings to reset the LastOpenstackReconcile timer if needed.
	// +optional
	OpenstackReconcileHash *string `json:"openstackReconcileHash,omitempty"`
}

func init() {
	SchemeBuilder.Register(&LoadBalancer{}, &LoadBalancerList{})
}
