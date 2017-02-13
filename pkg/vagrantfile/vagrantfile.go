package vagrantfile

import (
	"github.com/jgensler8/kother/pkg/spec"
	"fmt"
	"k8s.io/client-go/pkg/api/v1"
	"github.com/nginxinc/kubernetes-ingress/nginx-controller/nginx"
)

type Vagrantfile struct {
	Contents Block
}

// TODO: Change this to a template built on a more descriptive config (like nginx)
type Block struct {
	Start string
	Lines []Block
	End string
}

func SpecToVagrantfile(s *spec.Spec) (v *Vagrantfile, err error){
	v = DefaultVagrantfile(s.Config.DNS.RootDomain)
	for _, p := range s.Pods {
		v.Contents.Lines = append(v.Contents.Lines, DefaultLoadBalancerBlock(p, s))
	}
	return v, nil
}

func DefaultVagrantfile(tld string) (*Vagrantfile) {
	return &Vagrantfile{
		Block{
			Start: "Vagrant.configure(\"2\") do |config|",
			Lines: []Block{
				Block{ Start: "config.vm.box = \"coreos-stable\"", Lines: []Block{}, End: ""},
				Block{ Start: "config.vm.network \"private_network\", type: \"dhcp\"", Lines: []Block{}, End: "" },
				Block{ Start: "config.landrush.enabled = true", Lines: []Block{}, End: ""},
				Block{ Start: fmt.Sprintf("config.landrush.tld = '%s'", tld), Lines: []Block{}, End: ""},
			},
			End: "end",
		},
	}
}

func DefaultLoadBalancerBlock(p *v1.Pod, s *spec.Spec) (Block){
	portString := GetPortString(p)
	nginxConfig := GetNGINXConfig(p)
	return Block {
		Start: fmt.Sprintf("config.vm.define \"%s-lb\" do |lb|", ) ,
		Lines: []Block{
			Block{ Start: fmt.Sprintf("lb.vm.hostname = \"%s-lb.%s\"", p.ObjectMeta.Name, s.Config.DNS.RootDomain), Lines: []Block{}, End: ""},
			Block{ Start: fmt.Sprintf("lb.provision \"shell\", inline: \"echo '%s' > /tmp/nginx.conf\"", nginxConfig), Lines: []Block{}, End: ""},
			Block{ Start: fmt.Sprintf("lb.vm.provision \"shell\", inline: ( \"docker run -d %s --net=host --privileged --name nginx --restart always -v /tmp/nginx.conf:/etc/nginx/nginx.conf nginx", portString), Lines: []Block{}, End: ""},
		},
		End: "end",
	}
}

func DefaultPodsBlock(p *v1.Pod) (s string) {
	return ""
}

func GetPortString(p *v1.Pod) (s string) {
	for _, c := range p.Spec.Containers {
		for _, p := range c.Ports {
			s = fmt.Sprintf(" %s -p %d:%d ", s, p.ContainerPort, p.ContainerPort)
		}
	}
	return
}

func GetNGINXConfig(p *v1.Pod) (s string) {
	c := nginx.IngressNginxConfig {
		Upstreams: []nginx.Upstream{

		},
		Servers: []nginx.Server{
			nginx.Server{
				Name: p.ObjectMeta.Name,
				Locations: []nginx.Location{
					nginx.Location{
						Upstream:	nginx.Upstream{
									Name: "qwer",
									UpstreamServers: []nginx.UpstreamServer{
										nginx.UpstreamServer{
											Address: "qwer",
											Port: "80",
										},
									},
								},
					},
				},
			},
		},
	}
	return c.Servers[0].Name
}