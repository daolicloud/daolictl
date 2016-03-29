package main

import (
	"flag"
	"sort"

	"github.com/daolicloud/daolictl/cli"
)

var (
	flHelp    = flag.Bool("h, --help", false, "Print usage")
	flVersion = flag.Bool("v, --version", false, "Print version information and quit")
)

type byName []cli.Command

func (a byName) Len() int           { return len(a) }
func (a byName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byName) Less(i, j int) bool { return a[i].Name < a[j].Name }

var globalCommands []cli.Command

func init() {
	for _, cmd := range cli.GlobalCommands {
		globalCommands = append(globalCommands, cmd)
	}
	sort.Sort(byName(globalCommands))
}
