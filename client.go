package main

import (
	"flag"

	"github.com/daolicloud/daolictl/cli"
	"github.com/daolicloud/daolictl/utils"
)

var clientFlags = &cli.ClientFlags{FlagSet: new(flag.FlagSet), Common: commonFlags}

func init() {
	clientFlags.PostParse = func() {
		clientFlags.Common.PostParse()

		if clientFlags.Common.Debug {
			utils.EnableDebug()
		}
	}
}
