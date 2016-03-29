package client

import (
	"flag"
	"os"
)

func ParseFlags(fs *flag.FlagSet, args []string, withHelp bool) error {
	var help *bool
	if withHelp {
		help = fs.Bool("help", false, "Print usage")
	}
	if err := fs.Parse(args); err != nil {
		return err
	}
	if help != nil && *help {
		fs.SetOutput(os.Stdout)
		fs.Usage()
		os.Exit(0)
	}
	return nil
}
