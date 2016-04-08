package client

import (
	"fmt"
	"text/tabwriter"

	Cli "github.com/daolicloud/daolictl/cli"
)

// CmdGroup is the parent subcommand for all group commands
//
// Usage: daolictl group <COMMAND> [OPTIONS]
func (cli *DaoliCli) CmdGroup(args ...string) error {
	cmd := Cli.Subcmd("group", []string{"COMMAND [OPTIONS]"}, groupUsage(), false)
	err := ParseFlags(cmd, args, true)
	cmd.Usage()
	return err
}

// CmdGroupList lists all the groups managed by daolictl
//
// Usage: daolictl group list
func (cli *DaoliCli) CmdGroupList(args ...string) error {
	cmd := Cli.Subcmd("group list", []string{"COMMAND"}, "Lists groups", true)
	if err := ParseFlags(cmd, args, true); err != nil {
		return err
	}

	groups, err := cli.client.GroupList()
	if err != nil {
		return err
	}

	wr := tabwriter.NewWriter(cli.out, 0, 0, 0, ' ', 0)
	for _, group := range groups {
		fmt.Fprintf(wr, "%s\n", group)
	}
	wr.Flush()
	return nil
}

// Usage: daolictl group show <GROUP-NAME>
func (cli *DaoliCli) CmdGroupShow(args ...string) error {
	cmd := Cli.Subcmd("group show", []string{"GROUP"}, "Displays detailed information on the group", false)
	if err := ParseFlags(cmd, args, true); err != nil {
		return err
	}

	if len(cmd.Args()) <= 0 {
		cmd.Usage()
		return nil
	}

	members, err := cli.client.GroupShow(cmd.Arg(0))
	if err != nil {
		return err
	}

	for _, member := range members {
		fmt.Fprintf(cli.out, "%s\n", member)
	}
	return nil
}

// CmdGroupCreate creates a new group with a given name
//
// Usage: daolictl group create <GROUP-NAME>
func (cli *DaoliCli) CmdGroupCreate(args ...string) error {
	cmd := Cli.Subcmd("group create", []string{"GROUP"}, "Creates a new group with a name specified by the user", false)
	if err := ParseFlags(cmd, args, true); err != nil {
		return err
	}

	if len(cmd.Args()) <= 0 {
		cmd.Usage()
		return nil
	}

	group, err := cli.client.GroupCreate(cmd.Arg(0))
	if err != nil {
		return err
	}
	fmt.Fprintf(cli.out, "%s\n", group)
	return nil
}

// CmdGroupDelete creates a new group with a given name
//
// Usage: daolictl group delete <GROUP-NAME>
func (cli *DaoliCli) CmdGroupDelete(args ...string) error {
	cmd := Cli.Subcmd("group delete", []string{"GROUP [GROUP...]"}, "Deletes one or more groups", false)
	if err := ParseFlags(cmd, args, true); err != nil {
		return err
	}

	if len(cmd.Args()) <= 0 {
		cmd.Usage()
		return nil
	}

	var err error
	for _, group := range cmd.Args() {
		if e := cli.client.GroupDelete(group); e != nil {
			err = e
			continue
		}
	}

	return err
}

func groupUsage() string {
	groupCommands := map[string]string{
		"list":   "List all groups",
		"show":   "List members on the group",
		"create": "Create a group",
		"delete": "Delete a group",
	}

	help := "Commands:\n"

	for cmd, description := range groupCommands {
		help += fmt.Sprintf("  %-25.25s%s\n", cmd, description)
	}

	help += fmt.Sprintf("\nRun 'daolictl group COMMAND --help' for more information on a command.")
	return help
}
