package cmd

import (
	"github.com/urfave/cli"
	"github.com/golang/glog"
)

func Init(c *cli.Context) (err error) {
	glog.V(1).Infof("Initializing App")
	return nil
}