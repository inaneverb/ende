package cmd

import (
	"fmt"

	"github.com/inaneverb/ende/internal/pkg/version"
)

func UsageText(encoder, op string) string {
	const S = version.APP_NAME +
		" %s %s [<flags>] [IN: <data>|<path> [OUT:<path>]]"
	if encoder == "" {
		encoder = "<encoder>"
	}
	if op == "" {
		op = "<op>"
	}
	return fmt.Sprintf(S, encoder, op)
}

func gUsage(name string) string {
	const S = "Perform %s encode/decode operations"
	return fmt.Sprintf(S, name)
}

func gUsageEnc(name string) string {
	const S = "Perform %s encode operation"
	return fmt.Sprintf(S, name)
}

func gUsageDec(name string) string {
	const S = "Perform %s decode operation"
	return fmt.Sprintf(S, name)
}

func gDesc(name string) string {
	const S = `
Encodes or decodes the data using %s encoder/decoder.
You should specify the subcommand operation 
to perform a desired action.

You may use any of these inputs:
	- Stdin (by default)
	- Specify filepath to the input file as 1nd arg
	- Specify data as 1nd arg

You may use any of these outputs:
	- Stdout (by default)
	- Specify filepath where to store output as 2nd arg

If you want to use stream data from stdin, but also specify
the filepath, where you want to store the output data,
just use '--' as the 1st arg.
`
	return fmt.Sprintf(S, name)
}

func gDescEnc(name string) string {
	const S = `
Encodes either stream data from stdin, data source 
of given filepath from 1th argument or the 1st argument 
itself using %s encoder and writes the output to the file, 
path to which is given as 2nd argument, or to the stdout 
otherwise.`
	return fmt.Sprintf(S, name)
}

func gDescDec(name string) string {
	const S = `
Decodes either stream data from stdin, data source 
of given filepath from 1th argument or the 1st argument 
itself using %s decoder and writes the output to the file, 
path to which is given as 2nd argument, or to the stdout 
otherwise.`
	return fmt.Sprintf(S, name)
}
