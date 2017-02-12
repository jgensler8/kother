package ignition

import (
	"github.com/jgensler8/kother/pkg/ignition/units"
	"github.com/jgensler8/kother/pkg/ignition/files"
	"github.com/coreos/ignition/config/types"
	"github.com/coreos/go-systemd/unit"
	"k8s.io/client-go/pkg/api/v1"
)

func DefaultIgnition(k units.KubeletUnit, p v1.Pod) (_ *types.Config, err error) {
	ku, err := DefaultKubeletUnit(k)
	if err != nil {
		return nil, err
	}
	do, err := DefaultDockerDropIn()
	if err != nil {
		return nil, err
	}
	fi, err := files.ManifestToFile(p)
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

func DefaultKubeletUnit(k units.KubeletUnit) (*types.SystemdUnit, error){
	u := units.DefaultUnit(k)
	s, err := u.String()
	return &types.SystemdUnit{
		Name: types.SystemdUnitName(k.Name),
		Enable: true,
		Contents: s,
	}, err
}

func DefaultDockerDropIn() (*types.SystemdUnit, error){
	u := units.UnitFileBuilder.AddUnitOption(unit.NewUnitOption("Service", "Environment", "DOCKER_OPTS=\"${DOCKER_OPTS} --log-opt log-limit=50m\"")).Build()
	s, err := u.String()
	return &types.SystemdUnit{
		DropIns: []types.SystemdUnitDropIn{
			types.SystemdUnitDropIn{
				Name: "docker.service",
				Contents: s,
			},
		},
	}, err
}