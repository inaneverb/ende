package cmd

import (
	"github.com/urfave/cli/v2"
)

type FBoolAct = func(*cli.Context, bool) error

func FBool(name, alias, cat, usage string, val bool, act FBoolAct) *cli.BoolFlag {

	var defaultText = "true"
	if !val {
		defaultText = "false"
	}

	var f = cli.BoolFlag{
		Name: name, Category: cat, DefaultText: defaultText,
		HasBeenSet: val, Value: val,
		Aliases: []string{"s"}, Action: act,
		Usage: usage,
	}

	if alias != "" {
		f.Aliases = []string{alias}
	}

	return &f
}
