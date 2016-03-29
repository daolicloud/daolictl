package cli

import "flag"

// ClientFlags represents flags for the daolictl client.
type ClientFlags struct {
	FlagSet *flag.FlagSet
	Common *CommonFlags
	PostParse func()
}
