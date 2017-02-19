package units

import (
	"github.com/coreos/go-systemd/unit"
	"fmt"
	"github.com/jgensler8/kother/pkg/spec"
)

func DefaultUnit(s *spec.Spec) (UnitFile) {
	image := fmt.Sprintf("%s:%s", s.Hyperkube.Image.Name, s.Hyperkube.Image.Tag)
	name := "kubelet"
	return UnitFileBuilder.
		AddUnitOption(unit.NewUnitOption("Service", "ExecStartPre", fmt.Sprintf("/usr/bin/docker pull %s", image))).
		AddUnitOption(unit.NewUnitOption("Service", "ExecStartPre", fmt.Sprintf("-/usr/bin/docker rm %s", name))).
		AddUnitOption(unit.NewUnitOption("Service", "ExecStart", fmt.Sprintf("/usr/bin/docker run -it --net=host --name %s --hostname %s.$(hostname) %s /kubelet --api-servers=https://%s:443 --register-schedulable=false --container-runtime=docker --allow-privileged=true --pod-manifest-path=%s --cluster_dns=%s --cluster_domain=%s", name, name, image, s.Config.DNS.APIServerDNS, "/etc/kubernetes/pods", s.Config.CIDR.Cluster, s.Config.DNS.RootDomain))).
		AddUnitOption(unit.NewUnitOption("Service", "ExecStop", fmt.Sprintf("-/usr/bin/docker stop %s", name))).
		AddUnitOption(unit.NewUnitOption("Service", "ExecStop", fmt.Sprintf("-/usr/bin/docker rm %s", name))).
		AddUnitOption(unit.NewUnitOption("Service", "Restart", "always")).
		AddUnitOption(unit.NewUnitOption("Service", "RestartSec", "1")).
		AddUnitOption(unit.NewUnitOption("Install", "WantedBy", "multi-user.target")).
		Build()
}