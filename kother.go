package main

import (
	"github.com/jgensler8/kother/cmd"
	"os"
)

func main() {
	a := cmd.GetApp()
	a.Run(os.Args)
}