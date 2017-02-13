package cmd

import (
	"github.com/urfave/cli"
	"github.com/jgensler8/kother/pkg/spec"
	"github.com/jgensler8/kother/pkg/validate"
)

func Validate(c *cli.Context) (s *spec.Spec, err error) {
	wd := c.GlobalString(flagWorkDir)
	s, err = validate.Validate(&wd)
	return
}