package cloudconfig

import (
	"github.com/jgensler8/kother/pkg/cloudconfig/files"
	"github.com/jgensler8/kother/pkg/spec"
	"github.com/coreos/coreos-cloudinit/config"
	"github.com/jgensler8/kother/pkg/configurationsystem/units"
)

func DefaultKubeletUnit(s *spec.Spec) (*config.Unit, error){
	u := units.DefaultUnit(s)
	us, err := u.String()
	return &config.Unit{
		Name: "kubelet.service",
		Enable: true,
		Command: "start",
		Content: us,
	}, err
}

func DefaultDockerDropIn() (*config.Unit, error){
	return &config.Unit{
		Name: "docker.service",
		DropIns: []config.UnitDropIn{
			config.UnitDropIn{
				Name: "docker-10.conf",
				Content: "[Serivce]\nEnvironment=DOCKER_OPTS=\"${DOCKER_OPTS} --log-opt log-limit=50m\"",
			},
		},
	}, nil
}

func DefaultCloudConfig(c *spec.Component, s *spec.Spec) (_ *config.CloudConfig, err error) {
	ku, err := DefaultKubeletUnit(s)
	if err != nil {
		return nil, err
	}
	do, err := DefaultDockerDropIn()
	if err != nil {
		return nil, err
	}
	fi, err := files.ManifestToFile(c)
	if err != nil {
		return nil, err
	}
	cc := config.CloudConfig{
		WriteFiles: []config.File{
			*fi,
		},
		CoreOS: config.CoreOS{
			Units: []config.Unit{
				*ku,
				*do,
			},
		},
	}
	return &cc, err
}
