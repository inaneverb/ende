package core

import (
	"encoding/hex"
	"io"

	"github.com/urfave/cli/v2"
)

type Hex struct{}

func NewHex() *Hex { return &Hex{} }

func (h *Hex) Encode(cCtx *cli.Context) error {
	return do(cCtx, "hex.encode", func(r R, w W) (R, W, error) {
		var _, err = io.Copy(hex.NewEncoder(w), r)
		return r, w, err
	})
}

func (h *Hex) Decode(cCtx *cli.Context) error {
	return do(cCtx, "hex.decode", func(r R, w W) (R, W, error) {
		var _, err = io.Copy(w, hex.NewDecoder(r))
		return r, w, err
	})
}
