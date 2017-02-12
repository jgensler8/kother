package cmd

import "github.com/urfave/cli"

func Deploy(c *cli.Context) (err error) {
	err = Validate(c)
	if (err != nil) {
		return err
	}
	return nil
}