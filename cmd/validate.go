package cmd

import (
	"github.com/urfave/cli"
	"github.com/jgensler8/kother/pkg/spec"
	"github.com/jgensler8/kother/pkg/validate"
)

func Validate(c *cli.Context) (s *spec.Spec, err error) {
	cc := spec.CLIContext{
		WorkDir: c.GlobalString(flagWorkDir),
		CommitHash: c.GlobalString(flagCommitHash),
		CommitTag: c.GlobalString(flagCommitTag),
		ConfigurationSystem: c.GlobalString(flagConfigurationSystem),
	}
	s, err = validate.Validate(&cc)
	return
}