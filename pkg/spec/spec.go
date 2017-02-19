package spec

import "k8s.io/client-go/pkg/api/v1"

type Spec struct {
	//Pods []*v1.Pod
	Hyperkube `json:"Hyperkube,omitempty"`
	APIServer `json:"APIServer,omitempty"`
	ControllerManager `json:"ControllerManager,omitempty"`
	Scheduler `json:"Scheduler,omitempty"`
	Kubelet `json:"Kubelet,omitempty"`
	Etcd `json:"Etcd,omitempty"`
	Vault `json:"Vault,omitempty"`
	Config Config `json:"Config,omitempty"`
	ExtraVars interface{} `json:"ExtraVars,omitempty"`
}

type Config struct {
	CIDR `json:"CIDR,omitempty"`
	DNS `json:"DNS,omitempty"`
	Tags []Tag `json:"Tags,omitempty"`
}

type CIDR struct {
	Cluster string `json:"Cluster,omitempty"`
	Pod string `json:"Pod,omitempty"`
	Service string `json:"Service,omitempty"`
}

type DNS struct {
	RootDomain string `json:"RootDomain,omitempty"`
	ClusterDNSService string `json:"ClusterDNSService,omitempty"`
	KubernetesService string `json:"KubernetesService,omitempty"`
	// These are part of the spec and are constructed from the RootDomain above
	VaultDNS string `json:"VaultDNS,omitempty"`
	EtcdDNS string `json:"EtcdDNS,omitempty"`
	APIServerDNS string `json:"APIServerDNS,omitempty"`
	SchedulerDNS string `json:"SchedulerDNS,omitempty"`
	ControllerManagerDNS string `json:"ControllerManagerDNS,omitempty"`
}

type Tag struct {
	Key string `json:"Key,omitempty"`
	Value string `json:"Value,omitempty"`
}

type Hyperkube struct {
	Component
}
type APIServer struct {
	Component
}
type ControllerManager struct {
	Component
}
type Scheduler struct {
	Component
}
type Proxy struct {
	Component
}
type Kubelet struct {
	Component
}
type Etcd struct {
	Component
}
type Vault struct {
	Component
}

type Component struct {
	Image
	Pod *v1.Pod `json:"Pod,omitempty"`
	Replicas int `json:"Replicas,omitempty"`
}

type Image struct {
	Name string `json:"Image,omitempty"`
	Tag string `json:"Tag,omitempty"`
}