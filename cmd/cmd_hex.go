package cmd

import (
	"github.com/urfave/cli/v2"

	"github.com/inaneverb/ende/internal/core"
	"github.com/inaneverb/ende/internal/pkg/cmd"
)

func Hex() *cli.Command {

	const NAME = "hex"
	var x = core.NewHex()

	var e = cmd.Cmd(
		"encode", "e", gUsageEnc(NAME), UsageText(NAME, "encode"), gDescEnc(NAME),
		nil, x.Encode, nil)
	var d = cmd.Cmd(
		"decode", "d", gUsageDec(NAME), UsageText(NAME, "decode"), gDescDec(NAME),
		nil, x.Decode, nil)

	return cmd.Cmd(
		"hex", "h", gUsage(NAME), UsageText(NAME, ""), gDesc(NAME),
		nil, nil, []*cli.Command{e, d})

}
