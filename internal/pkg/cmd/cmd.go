package cmd

import (
	"strings"

	"github.com/urfave/cli/v2"
)

// Cmd is a constructor for cli.Command. It creates a new cli.Command object
// and fills it with given 'name', 'alias', 'usage', 'usageText', 'desc'.
// It also sets 'flags', 'sub' if they're not nil. Also wraps given 'act'
// to write returned error to the stderr.
func Cmd(
	name, alias, usage, usageText, desc string,
	flags []cli.Flag, act cli.ActionFunc, sub []*cli.Command) *cli.Command {

	var cmd = cli.Command{
		Name: name, Usage: usage, UsageText: usageText, Flags: flags,
		Description: strings.TrimSpace(desc), Action: act,
		HideHelpCommand: true,
	}

	if alias != "" {
		cmd.Aliases = []string{alias}
	}
	if len(sub) > 0 {
		cmd.Subcommands = sub
	}

	return &cmd
}
