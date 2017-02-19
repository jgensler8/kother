package cmd

import (
	"github.com/urfave/cli"
	"fmt"
	"github.com/jgensler8/kother/pkg/vagrantfile"
	"github.com/jgensler8/kother/pkg/configurationsystem"
)

func DeployAWS(c *cli.Context) (err error) {
	err = validateConfigurationSystem(c)
	if err != nil {
		return
	}
	s, err := Validate(c)
	if (err != nil) {
		return err
	}
	fmt.Printf("%v", c.GlobalString("commit-hash"))
	fmt.Printf("%v", s)
	return nil
}

func DeployVagrantfile(c *cli.Context) (err error) {
	fmt.Printf("%v", c.GlobalString("commit-hash"))
	err = validateConfigurationSystem(c)
	if err != nil {
		return
	}
	s, err := Validate(c)
	if (err != nil) {
		return err
	}
	v, err := vagrantfile.SpecToVagrantfile(s)
	if err != nil {
		fmt.Printf("Couldn't turn Spec into Vagrantfile")
		return
	}

	fmt.Printf("%v", *v.Contents)

	return nil
}

func validateConfigurationSystem(c *cli.Context) (err error) {
	configurationSystem := c.GlobalString(flagConfigurationSystem)
	if configurationSystem != configurationsystem.ConfigurationSystem_CloudConfig &&
		configurationSystem != configurationsystem.ConfigurationSystem_Ignition {
		return fmt.Errorf("flag (%s) was provided with an invalid option (%s). Should be either '%s' or '%s'",
			flagConfigurationSystem,
			configurationSystem,
			configurationsystem.ConfigurationSystem_CloudConfig,
			configurationsystem.ConfigurationSystem_Ignition)
	} else {
		return nil
	}
}