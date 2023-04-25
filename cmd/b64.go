package cmd

import (
	"github.com/urfave/cli/v2"

	"github.com/inaneverb/ende/internal/core"
	"github.com/inaneverb/ende/internal/pkg/cmd"
)

func Base64() *cli.Command {

	const NAME = "base64"
	var x = core.NewBase64()

	var flags = []cli.Flag{
		cmd.FBool("std", "s", "behavior", "use base charset and '+', '/' runes",
			true, x.SetStdEnc),
		cmd.FBool("url", "u", "behavior", "use base charset and '-', '_' runes",
			false, x.SetUrlEnc),
		cmd.FBool("raw", "r", "padding", "disable paddings",
			false, x.SetRawEnc),
	}

	var e = cmd.Cmd(
		"encode", "e", gUsageEnc(NAME), gDescEnc(NAME), flags, x.Encode, nil)
	var d = cmd.Cmd(
		"decode", "d", gUsageDec(NAME), gDescDec(NAME), flags, x.Decode, nil)

	return cmd.Cmd(
		"base64", "b64", gUsage(NAME), gDesc(NAME),
		nil, nil, []*cli.Command{e, d})
}
