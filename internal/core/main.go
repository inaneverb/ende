package core

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/urfave/cli/v2"

	"github.com/inaneverb/ende/internal/pkg/bw"
)

type R = io.Reader
type W = io.Writer

type EnDe interface {
	Encode(cCtx *cli.Context) error
	Decode(cCtx *cli.Context) error
}

// do encapsulate the main preparation job. It parses given arguments
// from cli.Context, prepares io.Reader, io.Writer to work with input and output
// and calls 'cb' providing these streams. It uses 'act' to wrap any occurred
// error and writes all occurred errors to the stderr. Always returns nil.
//
// The parsing arguments process includes:
// - Requiring 1 or 2 arguments,
// - Treating '--' from 1st argument as stdin,
// - Analysing 1st argument, to determine, whether its filepath or RAW data,
// - Ensuring, that 2nd argument is filepath (if provided).
func do(cCtx *cli.Context, act string, cb func(r R, w W) (R, W, error)) error {

	var r R = os.Stdin
	var w W = bw.NewBufferedWriter(os.Stdout)

	var wFile bool

	var r1 io.Reader
	var w1 io.Writer

	var f *os.File
	var err error
	var retErrs []error

	var args = cCtx.Args().Slice()
	if n := len(args); n < 1 || n > 2 {
		err = fmt.Errorf("%s: incorrect number of arguments: %d", act, n)
		retErrs = append(retErrs, err)
		goto EXIT
	}

	if len(args) > 0 && args[0] != "--" {
		f, err = stat(args[0])
		switch {
		case errors.Is(err, os.ErrInvalid):
			r = bytes.NewBuffer([]byte(args[0]))

		case err != nil:
			err = fmt.Errorf("%s: failed to open input: %w", act, err)
			retErrs = append(retErrs, err)
			goto EXIT

		default:
			r = f
		}
	}

	if len(args) > 1 {
		if f, err = stat(args[1]); err != nil {
			err = fmt.Errorf("%s: failed to open output: %w", act, err)
			retErrs = append(retErrs, err)
			goto EXIT
		} else {
			w = f
			wFile = true
		}
	}

	if r1, w1, err = cb(r, w); err != nil {
		retErrs = append(retErrs, fmt.Errorf("%s: %w", act, err))
		goto EXIT
	} else {
		_, _ = fmt.Fprintf(w, "\n")
	}

EXIT:

	for _, elem := range []struct {
		Obj  any
		Cond bool
		What string
	}{
		{r, true, "reader"},
		{w, wFile || len(retErrs) == 0, "writer"},
		{r1, r1 != nil && r1 != r, "returned reader"},
		{w1, w1 != nil && w1 != w, "returned writer"},
	} {
		if elem.Cond {
			if err = cl(elem.Obj); err != nil {
				err = fmt.Errorf("%s: failed to close %s: %w", act, elem.What, err)
				retErrs = append(retErrs, err)
			}
		}
	}

	for i, n := 0, len(retErrs); i < n; i++ {
		_, _ = fmt.Fprintf(os.Stderr, "error: %s\n", retErrs[i])
	}

	return nil
}

// stat tries to stat and open file with given 's' filepath.
// It returns os.ErrInvalid if it's not a filepath at all.
func stat(s string) (*os.File, error) {

	var fi, err = os.Stat(s)
	switch {
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
