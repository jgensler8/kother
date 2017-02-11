package units

import (
	"github.com/coreos/go-systemd/unit"
	"fmt"
)

func DefaultUnit(u KubeletUnit) (UnitFile) {
	image := fmt.Sprintf("%s:%s", u.HyperkubeImage, u.HyperkubeTag)
	return UnitFileBuilder.
		AddUnitOption(unit.NewUnitOption("Service", "ExecStartPre", fmt.Sprintf("/usr/bin/docker pull %s", image))).
		AddUnitOption(unit.NewUnitOption("Service", "ExecStartPre", fmt.Sprintf("-/usr/bin/docker rm %s", u.Name))).
		AddUnitOption(unit.NewUnitOption("Service", "ExecStart", fmt.Sprintf(`/usr/bin/docker \
run -it --net=host \
--name %s \
--hostname %s.$(hostname) \
%s \
/kubelet \
--api-servers=https://%s:443 \
--register-schedulable=false \
--container-runtime=docker \
--allow-privileged=true \
--pod-manifest-path=%s \
--cluster_dns=%s \
--cluster_domain=%s `, u.Name, u.Name, image, u.APIServerDNS, u.ManifestPath, u.ClusterDNS, u.ClusterDomain))).
		AddUnitOption(unit.NewUnitOption("Service", "ExecStop", fmt.Sprintf("-/usr/bin/docker stop %s", u.Name))).
		AddUnitOption(unit.NewUnitOption("Service", "ExecStop", fmt.Sprintf("-/usr/bin/docker rm %s", u.Name))).
		AddUnitOption(unit.NewUnitOption("Service", "Restart", "always")).
		AddUnitOption(unit.NewUnitOption("Service", "RestartSec", "1")).
		AddUnitOption(unit.NewUnitOption("Install", "WantedBy", "multi-user.target")).
		Build()
}