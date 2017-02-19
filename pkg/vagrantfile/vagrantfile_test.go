package vagrantfile_test

import (
	"github.com/jgensler8/kother/pkg/spec"
	"k8s.io/client-go/pkg/api/v1"
	"testing"
	"github.com/jgensler8/kother/pkg/vagrantfile"
	"github.com/jgensler8/kother/pkg/configurationsystem"
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
		Context: spec.CLIContext{
			ConfigurationSystem: configurationsystem.ConfigurationSystem_CloudConfig,
		},
		Config: spec.Config{
			DNS: spec.DNS{
				RootDomain: "vagrant.local",
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

func TestVagrant(t *testing.T) {
	_, err := vagrantfile.SpecToVagrantfile(&Spec)
	if err != nil {
		t.Logf("%v", err)
		t.Fail()
	}
	//t.Logf("%v", *v.Contents)
}