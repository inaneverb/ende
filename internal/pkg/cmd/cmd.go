package cmd

import (
	"os"
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
		Description: strings.TrimSpace(desc), Action: wrapAct(act),
	}

	if alias != "" {
		cmd.Aliases = []string{alias}
	}
	if len(sub) > 0 {
		cmd.Subcommands = sub
	}

	return &cmd
}

// wrapAct wraps given cli.ActionFunc, returning a new one, that calls provided,
// and writes a returned error to the stderr (if any), returning nil error.
func wrapAct(cb cli.ActionFunc) cli.ActionFunc {
	if cb == nil {
		return nil
	}
	return func(cCtx *cli.Context) error {
		if err := cb(cCtx); err != nil {
			_, _ = os.Stderr.WriteString("error: " + err.Error())
		}
		return nil
	}
}
