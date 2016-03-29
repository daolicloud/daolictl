package client

import (
	"fmt"
	"strconv"
	"strings"
	"text/tabwriter"

	Cli "github.com/daolicloud/daolictl/cli"
)

// Usage: daolictl firewall <COMMAND> [OPTIONS]
func (cli *DaoliCli) CmdFirewall(args ...string) error {
	cmd := Cli.Subcmd("firewall", []string{"COMMAND [OPTIONS]"}, firewallUsage(), false)
	err := ParseFlags(cmd, args, true)
	cmd.Usage()
	return err
}

// Usage: daolictl firewall list
func (cli *DaoliCli) CmdFirewallList(args ...string) error {
	cmd := Cli.Subcmd("firewall list", nil, "Lists firewalls", true)
	if err := ParseFlags(cmd, args, true); err != nil {
		return err
	}

	firewalls, err := cli.client.FirewallList()
	if err != nil {
		return err
	}

	wr := tabwriter.NewWriter(cli.out, 20, 1, 3, ' ', 0)
	fmt.Fprintln(wr, "FIREWALL NAME\tCONTAINER\tGATEWAY IP\tGATEWAY PORT\tSERVICE PORT")
	for _, firewall := range firewalls {
		fmt.Fprintf(wr, "%s\t%s\t%s\t%d\t%d\t",
			firewall.Name,
			strings.TrimLeft(firewall.Container, "/"),
			firewall.GatewayIP,
			firewall.GatewayPort,
			firewall.ServicePort,
		)
		fmt.Fprint(wr, "\n")
	}
	wr.Flush()
	return nil
}

// Usage: daolictl firewall show <CONTAINER>
func (cli *DaoliCli) CmdFirewallShow(args ...string) error {
	cmd := Cli.Subcmd("firewall show", []string{"CONTAINER"}, "Displays detailed information on the container", false)
	if err := ParseFlags(cmd, args, true); err != nil {
		return err
	}

	if len(cmd.Args()) <= 0 {
		cmd.Usage()
		return nil
	}

	firewalls, err := cli.client.FirewallShow(cmd.Arg(0))
	if err != nil {
		return err
	}

	wr := tabwriter.NewWriter(cli.out, 20, 1, 3, ' ', 0)
	fmt.Fprintln(wr, "FIREWALL NAME\tGATEWAY IP\tGATEWAY PORT\tSERVICE PORT")
	for _, firewall := range firewalls {
		fmt.Fprintf(wr, "%s\t%s\t%d\t%d\t",
			firewall.Name,
			firewall.GatewayIP,
			firewall.GatewayPort,
			firewall.ServicePort,
		)
		fmt.Fprint(wr, "\n")
	}
	wr.Flush()
	return nil
}

// Usage: daolictl firewall create <NAME> --rule <GATEWAYPORT:SERVICEPORT> --container <CONTAINER>
func (cli *DaoliCli) CmdFirewallCreate(args ...string) error {
	cmd := Cli.Subcmd("firewall create", []string{"[OPTIONS] FIREWALL-NAME"},
		"Creates a firewall rule with a given name", false)
	rule := cmd.String("rule", "", "Add firewall mapping")
	container := cmd.String("container", "", "Add firewall target")

	if err := ParseFlags(cmd, args, true); err != nil {
		return err
	}

	parts := strings.Split(*rule, ":")
	if len(parts) != 2 {
		fmt.Fprintf(cli.err, "--rule %s format is <GATEWAYPORT:SERVICEPORT>", *rule)
		cmd.Usage()
		return nil
	}

	if len(cmd.Args()) <= 0 {
		cmd.Usage()
		return nil
	}

	gport, err := strconv.Atoi(parts[0])
	if err != nil {
		return fmt.Errorf("GATEWAYPORT (%s) must be available number", parts[0])
	}

	sport, err := strconv.Atoi(parts[1])
	if err != nil {
		return fmt.Errorf("SERVICEPORT (%s) must be available number", parts[1])
	}

	fw := Firewall{
		Name:        cmd.Arg(0),
		Container:   *container,
		GatewayPort: gport,
		ServicePort: sport,
	}

	firewall, err := cli.client.FirewallCreate(fw)
	if err != nil {
		return err
	}

	wr := tabwriter.NewWriter(cli.out, 20, 1, 3, ' ', 0)
	fmt.Fprintln(wr, "FIREWALL NAME\tGATEWAY IP\tGATEWAY PORT\tSERVICE PORT")
	fmt.Fprintf(wr, "%s\t%s\t%d\t%d\t",
		firewall.Name,
		firewall.GatewayIP,
		firewall.GatewayPort,
		firewall.ServicePort,
	)
	fmt.Fprintln(wr)
	wr.Flush()
	return nil
}

// Usage: daolictl firewall delete <FIREWALL-NAME>
func (cli *DaoliCli) CmdFirewallDelete(args ...string) error {
	cmd := Cli.Subcmd("firewall delete", []string{"FIREWALL-NAME"}, "Deletes one or more firewalls", false)
	if err := ParseFlags(cmd, args, true); err != nil {
		return err
	}

	if len(cmd.Args()) <= 0 {
		cmd.Usage()
		return nil
	}

	var err error
	for _, name := range cmd.Args() {
		if e := cli.client.FirewallDelete(name); e != nil {
			err = e
			continue
		}
	}

	return err
}

func firewallUsage() string {
	firewallCommands := map[string]string{
		"list":   "List all firewall rules",
		"show":   "Display detailed information on the container",
		"create": "Create a firewall rule",
		"delete": "Delete a firewall rule",
	}

	help := "Commands:\n"

	for cmd, description := range firewallCommands {
		help += fmt.Sprintf("  %-25.25s%s\n", cmd, description)
	}

	help += fmt.Sprintf("\nRun 'daolictl firewall COMMAND --help' for more information on a command.")
	return help
}
