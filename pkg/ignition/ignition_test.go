package ignition_test

import (
	"testing"
	"github.com/jgensler8/kother/pkg/ignition"
	"k8s.io/client-go/pkg/api/v1"
	"github.com/jgensler8/kother/pkg/spec"
)

var (
	Pod = v1.Pod{
		ObjectMeta: v1.ObjectMeta{
			Name: "apiserver",
		},
		Spec: v1.PodSpec{
			Containers: []v1.Container{
				v1.Container{
					Name: "kube-apiserver",
					Ports: []v1.ContainerPort{
						v1.ContainerPort{
							ContainerPort: 443,
						},
						v1.ContainerPort{
							ContainerPort: 8080,
						},
					},
				},
			},
		},
	}
	Spec = spec.Spec{
		Hyperkube: spec.Hyperkube{
			Component: spec.Component{
				Image: spec.Image{
					Name: "quay.io/coreos/hyperkube",
					Tag: "latest",
				},
			},
		},
		Config: spec.Config{
			DNS: spec.DNS{
				RootDomain: "vagrant.local",
				APIServerDNS: "apiserver.vagrant.local",
			},
			CIDR: spec.CIDR{
				Cluster: "1.2.3.4/20",
			},
		},
		Etcd: spec.Etcd{
			Component: spec.Component{
				Pod: &Pod,
				Replicas: 2,
			},
		},
		Vault: spec.Vault{
			Component: spec.Component{
				Pod: &Pod,
				Replicas: 2,
			},
		},
		APIServer: spec.APIServer{
			Component: spec.Component{
				Pod: &Pod,
				Replicas: 2,
			},
		},
		Scheduler: spec.Scheduler{
			Component: spec.Component{
				Pod: &Pod,
				Replicas: 2,
			},
		},
		ControllerManager: spec.ControllerManager{
			Component: spec.Component{
				Pod: &Pod,
				Replicas: 2,
			},
		},
	}
)

func TestDefaultIgnition(t *testing.T) {
	_, err := ignition.DefaultIgnition(&Spec.APIServer.Component, &Spec)
	if err != nil {
		t.Fatalf("%v", err)
	}
	//t.Logf("%v", x)
}
