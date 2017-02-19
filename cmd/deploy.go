package cmd

import (
	"github.com/urfave/cli"
	"fmt"
	"github.com/jgensler8/kother/pkg/vagrantfile"
)

func DeployAWS(c *cli.Context) (err error) {
	s, err := Validate(c)
	if (err != nil) {
		return err
	}
	fmt.Printf("%v", c.GlobalString("commit-hash"))
	fmt.Printf("%v", s)
	return nil
}

func DeployVagrantfile(c *cli.Context) (err error) {
	s, err := Validate(c)
	if (err != nil) {
		return err
	}
	v, err := vagrantfile.SpecToVagrantfile(s)
	if err != nil {
		fmt.Printf("Couldn't turn Spec into Vagrantfile")
		return
	}
	//fmt.Printf("%v", s)
	fmt.Printf("%v", *v.Contents)

	return nil
}