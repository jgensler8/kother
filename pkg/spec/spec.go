package spec

import "k8s.io/client-go/pkg/api/v1"

type Spec struct {
	Pods []*v1.Pod
	Config Config
	ExtraVars interface{}
}

type Config struct {
	HyperkubeImage string
	HyperkubeTag string
	EtcdImage string
	EtcdTag string
	VaultImage string
	VaultTag string
	CIDR
	DNS
	Tags []Tag
}

type CIDR struct {
	Cluster string
	Pod string
	Service string
}

type DNS struct {
	RootDomain string
	ClusterDNSService string
	KubernetesService string
	// These are part of the spec and are constructed from the RootDomain above
	VaultDNS string
	EtcdDNS string
	APIServerDNS string
	SchedulerDNS string
	ControllerManagerDNS string
}

type Tag struct {
	Key string
	Value string
}