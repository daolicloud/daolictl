package client

import (
	"fmt"
	"strings"
	"text/tabwriter"

	Cli "github.com/daolicloud/daolictl/cli"
)

// Usage: daolictl policy
func (cli *DaoliCli) CmdPolicy(args ...string) error {
	cmd := Cli.Subcmd("policy", []string{"COMMAND [OPTIONS]"}, policyUsage(), false)
	err := ParseFlags(cmd, args, true)
	cmd.Usage()
	return err
}

//Usage: daolictl policy create <PEER>
func (cli *DaoliCli) CmdPolicyList(args ...string) error {
	cmd := Cli.Subcmd("policy list", nil, "Lists policies", true)
	if err := ParseFlags(cmd, args, true); err != nil {
		return err
	}

	polis, err := cli.client.PolicyList()
	if err != nil {
		return err
	}

	wr := tabwriter.NewWriter(cli.out, 20, 1, 3, ' ', 0)
	for _, p := range polis {
		parts := strings.Split(p, ":")
		if len(parts) == 2 {
			fmt.Fprintf(wr, "%s\t%s\t", parts[0], parts[1])
			fmt.Fprintf(wr, "\n")
		}
	}
	wr.Flush()
	return nil
}

// Usage: daolictl policy create <CONTAINER:CONTAINER>
func (cli *DaoliCli) CmdPolicyCreate(args ...string) error {
	cmd := Cli.Subcmd("policy create", []string{"CONTAINER:CONTAINER"}, "Creates a policy with container peer", false)
	if err := ParseFlags(cmd, args, true); err != nil {
		return err
	}

	if len(cmd.Args()) <= 0 {
		cmd.Usage()
		return nil
	}

	err := cli.client.PolicyCreate(cmd.Arg(0))
	if err != nil {
		return err
	}
	fmt.Fprintf(cli.out, "%s\n", cmd.Arg(0))
	return nil
}

// Usage: daolictl policy delete <CONTAINER:CONTAINER>
func (cli *DaoliCli) CmdPolicyDelete(args ...string) error {
	cmd := Cli.Subcmd("policy delete", []string{"CONTAINER:CONTAINER"}, "Delete a policy with container peer", false)
	if err := ParseFlags(cmd, args, true); err != nil {
		return err
	}

	if len(cmd.Args()) <= 0 {
		cmd.Usage()
		return nil
	}

	err := cli.client.PolicyDelete(cmd.Arg(0))
	if err != nil {
		return err
	}
	fmt.Fprintf(cli.out, "%s\n", cmd.Arg(0))
	return nil
}

func policyUsage() string {
	policyCommands := map[string]string{
		"list":   "List all policy",
		"create": "Create a rule",
		"delete": "Delete a rule",
	}

	help := "Commands:\n"

	for cmd, description := range policyCommands {
		help += fmt.Sprintf("  %-25.25s%s\n", cmd, description)
	}

	help += fmt.Sprintf("\nRun 'daolictl policy COMMAND --help' for more information on a command.")
	return help
}
