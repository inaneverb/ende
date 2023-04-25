package core

import (
	"encoding/base64"
	"io"

	"github.com/urfave/cli/v2"
)

// Base64 implements EnDe and is the main entry point to perform base64
// encode/decode operations.
type Base64 struct {
	stdEnc bool
	urlEnc bool
	rawEnc bool
}

func NewBase64() *Base64 { return &Base64{} }

func (b *Base64) Encode(cCtx *cli.Context) error {
	return do(cCtx, "base64.encode", func(r R, w W) (R, W, error) {
		w = base64.NewEncoder(b.getEnc(), w)
		var _, err = io.Copy(w, r)
		return r, w, err
	})
}

func (b *Base64) Decode(cCtx *cli.Context) error {
	return do(cCtx, "base64.decode", func(r R, w W) (R, W, error) {
		r = base64.NewDecoder(b.getEnc(), r)
		var _, err = io.Copy(w, r)
		return r, w, err
	})
}

func (b *Base64) SetStdEnc(_ *cli.Context, v bool) error { b.stdEnc = v; return nil }
func (b *Base64) SetUrlEnc(_ *cli.Context, v bool) error { b.urlEnc = v; return nil }
func (b *Base64) SetRawEnc(_ *cli.Context, v bool) error { b.rawEnc = v; return nil }

func (b *Base64) getEnc() *base64.Encoding {
	switch {
	case b.stdEnc && b.rawEnc:
		return base64.RawStdEncoding
	case b.urlEnc && b.rawEnc:
		return base64.RawURLEncoding
	case b.urlEnc:
		return base64.URLEncoding
	default:
		return base64.StdEncoding
	}
}
