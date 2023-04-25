package rw

import (
	"io"
)

type tsr struct {
	origin io.Reader
	buf    []byte
	wasN   bool
}

func NewTrimSpacesReader(origin io.Reader) io.ReadCloser {
	return &tsr{origin, nil, false}
}

func (r *tsr) Read(p []byte) (int, error) {

	var n int

	switch {
	case r.buf == nil:
		r.buf = make([]byte, len(p)-1)

	case len(r.buf) > 0:
		n = len(r.buf)
		copy(p, r.buf)
	}

	r.buf = r.buf[:cap(r.buf)]

	var m, err = r.origin.Read(r.buf)
	if err != nil {
		return n, err
	}

	if n > 0 && r.wasN {
		p[len(p)-1] = '\n'
	}

	r.buf = r.buf[:m]

	if r.buf[m-1] == '\n' {
		r.wasN = true
		r.buf = r.buf[:m-1]
	}

	return n, err
}

func (r *tsr) Close() error { return nil }
