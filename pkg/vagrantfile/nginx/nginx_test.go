package nginx_test

import (
	"testing"
	"github.com/jgensler8/kother/pkg/vagrantfile/nginx"
	"github.com/jgensler8/kother/pkg/spec"
	"k8s.io/client-go/pkg/api/v1"
)

var (
	s = spec.Spec{
		Config: spec.Config{
			DNS: spec.DNS{
				RootDomain: "vagrant.local",
			},
		},
	}
	c = spec.Component{
		Pod: &v1.Pod{
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
		},
		Replicas: 2,
	}
)

func TestGetNGINXConfig(t *testing.T) {
	_, err := nginx.GetNGINXConfig(&c, &s)
	if err != nil {
		t.Logf("%v", err)
		t.Fail()
	}
}