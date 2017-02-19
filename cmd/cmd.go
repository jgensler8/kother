package cmd

import (
	"github.com/urfave/cli"
)

var (
	flagWorkDir = "work-dir"
	flagLogLevel = "log-level"

	flagSubmit = "submit"

	flagCommitHash = "commit-hash"
	flagCommitTag = "commit-tag"
	flagPrint = "print"
	flagSave = "save"
	flagTags = "tags"
	flagConfigurationSystem = "configuration-system"
)

func GetApp() (* cli.App){
	app := cli.NewApp()
	app.Version = "0.1.0"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  flagWorkDir,
			Usage: "Working Directory for the command. This will be used either to Read or Create files in. Directory will *not* be created.",
			Value: ".",
		},
		cli.IntFlag{
			Name:  flagLogLevel,
			Usage: "TODO -log=1 OR -log=2 OR -log=999",
			Value: 0,
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
				_, err := Validate(c)
				return err
			},
		},
		{
			Name:        "deploy",
			Usage:       "Validate and Deploy a directory. Note that this incldue both Create and Update. (Idempotent)",
			Subcommands: []cli.Command{
				{
					Name:  "aws",
					Usage: "set template to CloudFormation",
					Action: func(c *cli.Context) error {
						return DeployAWS(c)
					},
					Flags:  []cli.Flag{
						cli.BoolFlag{
							Name:  flagSubmit,
							Usage: "Submits CloudFormation to AWS.",
						},
					},
				},
				{
					Name:  "vagrant",
					Usage: "set template to Vagrantfile",
					Action: func(c *cli.Context) error {
						return DeployVagrantfile(c)
					},
				},
			},
			Flags:  []cli.Flag{
				cli.StringFlag{
					Name:  flagCommitHash,
					Usage: "For use with feature branches. All resources will be tagged with commithash=${your-commit-hash}.",
					Value: "",
				},
				cli.StringFlag{
					Name:  flagCommitTag,
					Usage: "For use with feature branches. The template name will be named appened with this.",
					Value: "",
				},
				cli.BoolFlag{
					Name: flagPrint,
					Usage: "Writes template to the STDOUT.",
				},
				cli.BoolFlag{
					Name:  flagSave,
					Usage: "Writes template to the working directory.",
				},
				cli.StringFlag{
					Name:  flagTags,
					Usage: "Comma Separated set of Tags to attach to all AWS Resources.",
					Value: "",
				},
				cli.StringFlag{
					Name:  flagConfigurationSystem,
					Usage: "What configuration system to use with CoreOS (ignition or cloud-config)",
					Value: "cloud-config",
				},
			},
		},
	}
	return app
}