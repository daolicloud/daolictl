package cli

import "flag"

type Options struct {
	CAFile   string
	CertFile string
	KeyFile  string

	// client-only option
	InsecureSkipVerify bool
}

// CommonFlags represents flags that are common to the client.
type CommonFlags struct {
	FlagSet   *flag.FlagSet
	PostParse func()

	Debug      bool
	Hosts      []string
	LogLevel   string
	TLS        bool
	TLSVerify  bool
	TLSOptions *Options
}

// Command is the struct containing the command name and description
type Command struct {
	Name        string
	Description string
}

var globalCommands = []Command{
	{"firewall", "Firewall management"},
	{"group", "Group management"},
	{"member", "Group member management"},
	//{"policy", "Policy management"},
	{"cut", "Do container isolation"},
	{"cutls", "List all cuts"},
	{"uncut", "Undo container isolation"},
	{"version", "Show the daolinet version information"},
}

// GlobalCommands stores all the daolictl command
var GlobalCommands = make(map[string]Command)

func init() {
	for _, cmd := range globalCommands {
		GlobalCommands[cmd.Name] = cmd
	}
}
