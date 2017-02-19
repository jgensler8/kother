package ignition

import (
	"github.com/jgensler8/kother/pkg/ignition/units"
	"github.com/jgensler8/kother/pkg/ignition/files"
	"github.com/coreos/ignition/config/types"
	"github.com/coreos/go-systemd/unit"
	"github.com/jgensler8/kother/pkg/spec"
)

func DefaultIgnition(c *spec.Component, s *spec.Spec) (_ *types.Config, err error) {
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
	return &types.Config{
		Ignition: types.Ignition{
			Version: types.IgnitionVersion(types.MaxVersion),
		},
		Systemd: types.Systemd{
			Units: []types.SystemdUnit{
				*ku,
				*do,
			},
		},
		Storage: types.Storage{
			Files: []types.File{
				*fi,
			},
		},
	}, err
}

func DefaultKubeletUnit(s *spec.Spec) (*types.SystemdUnit, error){
	u := units.DefaultUnit(s)
	us, err := u.String()
	return &types.SystemdUnit{
		Name: types.SystemdUnitName("kubelet.service"),
		Enable: true,
		Contents: us,
	}, err
}

func DefaultDockerDropIn() (*types.SystemdUnit, error){
	u := units.UnitFileBuilder.AddUnitOption(unit.NewUnitOption("Service", "Environment", "DOCKER_OPTS=\"${DOCKER_OPTS} --log-opt log-limit=50m\"")).Build()
	s, err := u.String()
	return &types.SystemdUnit{
		Name: "docker.service",
		DropIns: []types.SystemdUnitDropIn{
			types.SystemdUnitDropIn{
				Name: "docker-10.conf",
				Contents: s,
			},
		},
	}, err
}