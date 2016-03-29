package client

import (
	"fmt"
	Cli "github.com/daolicloud/daolictl/cli"
)

// Usage: daolictl member
func (cli *DaoliCli) CmdMember(args ...string) error {
	cmd := Cli.Subcmd("member", []string{"COMMAND [OPTIONS]"}, memberUsage(), false)
	err := ParseFlags(cmd, args, true)
	cmd.Usage()
	return err
}

// Usage: daolictl member add <NETWORK-NAME> --group <GROUP-NAME>
func (cli *DaoliCli) CmdMemberAdd(args ...string) error {
	cmd := Cli.Subcmd("member add", []string{"NETWORK"}, "Add a network to the group", false)
	group := cmd.String("group", "", "Group Name")
	if err := ParseFlags(cmd, args, true); err != nil {
		return err
	}

	if len(cmd.Args()) <= 0 {
		cmd.Usage()
		return nil
	}

	if *group == "" {
		cmd.Usage()
		return nil
	}

	err := cli.client.MemberAdd(cmd.Arg(0), *group)
	if err != nil {
		return err
	}
	fmt.Fprintf(cli.out, "%s\n", cmd.Arg(0))
	return nil
}

// Usage: daolictl member rm <NETWORK-NAME> --group <GROUP-NAME>
func (cli *DaoliCli) CmdMemberRm(args ...string) error {
	cmd := Cli.Subcmd("member rm", []string{"NETWORK"}, "Remove a network from the group", false)
	group := cmd.String("group", "", "Group Name")
	if err := ParseFlags(cmd, args, true); err != nil {
		return err
	}

	if len(cmd.Args()) <= 0 {
		cmd.Usage()
		return nil
	}

	if *group == "" {
		cmd.Usage()
		return nil
	}

	err := cli.client.MemberRemove(cmd.Arg(0), *group)
	if err != nil {
		return err
	}
	fmt.Fprintf(cli.out, "%s\n", cmd.Arg(0))
	return nil

}

func memberUsage() string {
	memberCommands := map[string]string{
		"add": "Add a member to group",
		"rm":  "Remove a group member",
	}

	help := "Commands:\n"

	for cmd, description := range memberCommands {
		help += fmt.Sprintf("  %-25.25s%s\n", cmd, description)
	}

	help += fmt.Sprintf("\nRun 'daolictl member COMMAND --help' for more information on a command.")
	return help
}
