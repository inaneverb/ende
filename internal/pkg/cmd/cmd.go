package cmd

import (
	"github.com/urfave/cli/v2"
)

func Cmd(
	name, alias, usage, desc string,
	flags []cli.Flag, act cli.ActionFunc, sub []*cli.Command) *cli.Command {

	var cmd = cli.Command{
		Name: name, Usage: usage, Description: desc, Action: act, Flags: flags,
	}

	if alias != "" {
		cmd.Aliases = []string{alias}
	}
	if len(sub) > 0 {
		cmd.Subcommands = sub
	}

	return &cmd
}
