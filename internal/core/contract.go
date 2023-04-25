package core

import (
	"io"

	"github.com/urfave/cli/v2"
)

type R = io.Reader
type W = io.Writer

type EnDe interface {
	Encode(cCtx *cli.Context) error
	Decode(cCtx *cli.Context) error
}
