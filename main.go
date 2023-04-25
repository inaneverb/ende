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
	const Description = `
This tool allows you to perform encode decode operations.
It allows you to work with stdin, stdout, or with a files.
`
	var ver = version.VERSION +
		"; commit: " + version.Commit +
		"; date: " + version.Date

	var app = cli.App{
		Name:      version.APP_NAME,
		Version:   ver,
		Usage:     Usage,
		UsageText: cmd.UsageText("", ""),
		Commands: []*cli.Command{
			cmd.Base64(), cmd.Hex(),
		},
		Description: strings.TrimSpace(Description),
		Flags:       []cli.Flag{},
		CommandNotFound: func(cCtx *cli.Context, cmd string) {
			_, _ = fmt.Fprintf(cCtx.App.Writer, "Thar be no '%q' here.\n", cmd)
		},
		HideHelpCommand: true,
	}

	_ = app.Run(os.Args)
}
