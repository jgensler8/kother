package cmd

import (
	"github.com/urfave/cli"
)

func GetApp() (* cli.App){
	app := cli.NewApp()
	app.Version = "0.1.0"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "work-dir",
			Usage: "Working Directory for the command. This will be used either to Read or Create files in. Directory will *not* be created.",
		},
		cli.StringFlag{
			Name:  "log level. TODO.",
			Usage: "-log=1 OR -log=2 OR -log=999",
		},
	}
	app.Commands = []cli.Command{
		{
			Name:    "init",
			Usage:   "Initialize a directory to the kother spec. This generates Manifests and Configuration files.",
			Action:  func(c *cli.Context) error {
				return Init(c)
			},
		},
		{
			Name:    "validate",
			Usage:   "Validate a directory with the kother spec.",
			Action:  func(c *cli.Context) error {
				return Validate(c)
			},
			Flags:  []cli.Flag{
				cli.StringFlag{
					Name:  "config",
					Usage: "Load configuration from `FILE`",
				},
			},
		},
		{
			Name:        "deploy",
			Usage:       "Validate and Deploy a directory. Note that this incldue both Create and Update. (Idempotent)",
			Action: func(c *cli.Context) error {
				return Deploy(c)
			},
			Flags:  []cli.Flag{
				cli.StringFlag{
					Name:  "commit-hash",
					Usage: "Default: \"\"; For use with feature branches. All resources will be tagged with commithash=${your-commit-hash}.",
				},
				cli.StringFlag{
					Name:  "commit-tag",
					Usage: "Default: \"\"; For use with feature branches. The Stack will be named appened with this.",
				},
				cli.StringFlag{
					Name:  "print-cloudformation",
					Usage: "Default: \"false\"; Writes CloudFormation to the STDOUT.",
				},
				cli.StringFlag{
					Name:  "generate-cloudformation",
					Usage: "Default: \"false\"; Writes CloudFormation to the working directory.",
				},
				cli.StringFlag{
					Name:  "submit-cloudformation",
					Usage: "Default: \"true\"; Submits CloudFormation to AWS.",
				},
				cli.StringFlag{
					Name:  "tags",
					Usage: "Default: \"\"; Comma Separated set of Tags to attach to all AWS Resources.",
				},
			},
		},
	}
	return app
}