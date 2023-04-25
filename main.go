package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/urfave/cli/v2"

	"github.com/inaneverb/ende/cmd"
	"github.com/inaneverb/ende/internal/pkg/version"
)

func main() {

	const Usage = "CLI tool to perform encode/decode operations"
	const UsageArgs = version.APP_NAME +
		" <encoder> <op> [IN: <data>|<filepath> [OUT:<filepath>]]"

	const Description = `
This tool allows you to perform encode decode operations.
It allows you to work with stdin, stdout, or with a files.
`

	var app = cli.App{
		Name:      version.APP_NAME,
		Version:   version.VERSION + "; commit: " + version.Commit,
		Usage:     Usage,
		UsageText: UsageArgs,
		//ArgsUsage: UsageArgs,
		Commands: []*cli.Command{
			cmd.Base64(),
		},
		Description:          strings.TrimSpace(Description),
		Flags:                []cli.Flag{},
		EnableBashCompletion: true,
		HideHelp:             false,
		HideVersion:          false,
		CommandNotFound: func(cCtx *cli.Context, cmd string) {
			_, _ = fmt.Fprintf(cCtx.App.Writer, "Thar be no '%q' here.\n", cmd)
		},
	}

	_ = app.Run(os.Args)
}
