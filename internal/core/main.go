package core

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/urfave/cli/v2"
)

// do encapsulate the main preparation job. It parses given arguments
// from cli.Context, prepares io.Reader, io.Writer to work with input and output
// and calls 'cb' providing these streams. It uses 'act' to wrap any occurred
// error.
//
// The parsing arguments process includes:
// - Requiring 1 or 2 arguments,
// - Treating '--' from 1st argument as stdin,
// - Analysing 1st argument, to determine, whether its filepath or RAW data,
// - Ensuring, that 2nd argument is filepath (if provided).
func do(cCtx *cli.Context, act string, cb func(r R, w W) (R, W, error)) error {

	var r R = os.Stdin
	var w W = os.Stdout

	var args = cCtx.Args().Slice()
	if n := len(args); n < 1 || n > 2 {
		return fmt.Errorf("%s: incorrect number of arguments: %d", act, n)
	}

	var f *os.File
	var err error

	if len(args) > 0 && args[0] != "--" {
		switch f, err = stat(args[0]); {
		case errors.Is(err, os.ErrInvalid):
			r = bytes.NewBuffer([]byte(args[0]))
		case err != nil:
			return fmt.Errorf("%s: failed to open input: %w", act, err)
		default:
			r = f
		}
	}

	if len(args) > 1 {
		if f, err = stat(args[1]); err != nil {
			return fmt.Errorf("%s: failed to open output: %w", act, err)
		} else {
			w = f
		}
	}

	var r1 io.Reader
	var w1 io.Writer

	if r1, w1, err = cb(r, w); err != nil {
		return fmt.Errorf("%s: %w", act, err)
	}

	if r1 != r && r1 != nil {
		if err = cl(r1); err != nil {
			return fmt.Errorf("%s: failed to close new reader: %w", act, err)
		}
	}

	if w1 != w && w1 != nil {
		if err = cl(w1); err != nil {
			return fmt.Errorf("%s: failed to close new writer: %w", act, err)
		}
	}

	return nil
}

// stat tries to stat and open file with given 's' filepath.
// It returns os.ErrInvalid if it's not a filepath at all.
func stat(s string) (*os.File, error) {

	var fi, err = os.Stat(s)
	switch {
	case errors.Is(err, os.ErrNotExist):
		return nil, err
	case errors.Is(err, os.ErrPermission):
		return nil, err
	case err != nil:
		return nil, os.ErrInvalid
	}

	if fi.IsDir() {
		return nil, fmt.Errorf("it is dir")
	}

	if fi.Mode()&0400 == 0400 {
		return nil, os.ErrPermission
	}

	return os.Open(s)
}

// cl wraps 'v' trying to cast it to the io.Closer, and calling Close()
// if cast was successful. Returns an error, that returned from Close().
func cl(v any) error {
	if c, _ := v.(io.Closer); c != nil {
		return c.Close()
	}
	return nil
}
