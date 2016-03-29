package client

import (
	Cli "github.com/daolicloud/daolictl/cli"
)

// CmdVersion shows Daolictl version information
// Usage: daolictl version
func (cli *DaoliCli) CmdVersion(args ...string) (err error) {
	cmd := Cli.Subcmd("version", nil, Cli.GlobalCommands["version"].Description, true)
	ParseFlags(cmd, args, true)
	cli.out.Write([]byte(DefaultVersion))
	cli.out.Write([]byte{'\n'})
	return err
}
