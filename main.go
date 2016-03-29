package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/daolicloud/daolictl/cli"
	"github.com/daolicloud/daolictl/client"
	"github.com/daolicloud/daolictl/utils"
)

func main() {
	stdin, stdout, stderr := utils.StdStreams()

	logrus.SetOutput(stderr)
	flag.Usage = func() {
		fmt.Fprint(stdout, "Usage: daolictl [OPTIONS] COMMAND [arg...]\n       daolictl [ --help | -v | --version ]\n\n")
		fmt.Fprint(stdout, "A tool for docker network.\n\nOptions:\n")
		flag.CommandLine.SetOutput(stdout)
		flag.PrintDefaults()

		help := "\nCommands:\n"

		for _, cmd := range globalCommands {
			help += fmt.Sprintf("    %-10.10s%s\n", cmd.Name, cmd.Description)
		}

		help += "\nRun 'daolictl COMMAND --help' for more information on a command."
		fmt.Fprintf(stdout, "%s\n", help)
	}

	flag.Parse()

	if *flVersion {
		fmt.Printf("Daolictl version %s\n", client.DefaultVersion)
		return
	}

	if *flHelp {
		// if global flag --help is present, regardless of what other options and commands there are,
		// just print the usage.
		flag.Usage()
		return
	}

	clientCli := client.NewDaoliCli(stdin, stdout, stderr, clientFlags)
	c := cli.New(clientCli)
	if err := c.Run(flag.Args()...); err != nil {
		fmt.Fprintln(stderr, err)
		os.Exit(1)
	}
}
