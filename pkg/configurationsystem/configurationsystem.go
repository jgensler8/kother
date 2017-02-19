package configurationsystem

import (
	"fmt"
	"encoding/json"
	"github.com/jgensler8/kother/pkg/spec"
	"github.com/jgensler8/kother/pkg/ignition"
	"github.com/jgensler8/kother/pkg/cloudconfig"
)

var (
	ConfigurationSystem_Ignition string = "ignition"
	ConfigurationSystem_CloudConfig string = "cloud-config"
)

func GetUserData(c *spec.Component, s *spec.Spec) (_ string, err error) {
	switch s.Context.ConfigurationSystem {
	case ConfigurationSystem_Ignition:
		q, err := ignition.DefaultIgnition(c, s)
		if err != nil {
			fmt.Printf("Couldn't create Default Ignition template")
			return "", err
		}
		var b []byte
		b, err = json.Marshal(q)
		return string(b), err
	case ConfigurationSystem_CloudConfig:
		q, err := cloudconfig.DefaultCloudConfig(c, s)
		if err != nil {
			fmt.Printf("Couldn't create Default CloudConfig template")
			return "", err
		}
		return q.String(), err
	default:
		return "", fmt.Errorf("Couldn't find configuration system...")
	}
}