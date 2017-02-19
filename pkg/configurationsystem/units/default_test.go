package units_test

import (
	"testing"
	"github.com/jgensler8/kother/pkg/configurationsystem/units"
	"io/ioutil"
	"strings"
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

func TestDefaultUnit(t *testing.T) {
	u := units.DefaultUnit(&Spec)
	if len(u.Lines) != 8 {
		t.Fail()
	}
}

func TestDefaultUnit_Serialize(t *testing.T) {
	u := units.DefaultUnit(&Spec)
	r := u.Serialize()
	doc, err := ioutil.ReadAll(r)
	if err != nil {
		t.Fatalf("%v", err)
	}
	if strings.Count(string(doc), "\n") != 11 {
		t.Logf("Lines: %d", strings.Count(string(doc), "\n"))
		t.Logf("%v", string(doc))
		t.Fail()
	}
}