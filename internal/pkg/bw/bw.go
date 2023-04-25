package bw

import (
	"io"
)

type bw struct {
	origin  io.Writer
	buf     []byte
	drained bool
}

func NewBufferedWriter(origin io.Writer) io.WriteCloser {
	return &bw{origin, make([]byte, 0, 1024*1024), false}
}

func (b *bw) Write(p []byte) (int, error) {

	switch {
	case len(p)+len(b.buf) > cap(b.buf):
		var n, err = b.ensureFlush()
		if err != nil {
			return n, err
		}
		fallthrough

	case b.drained:
		return b.origin.Write(p)

	default:
		b.buf = append(b.buf, p...)
		return len(p), nil
	}
}

func (b *bw) Close() error {
	var _, err = b.ensureFlush()
	return err
}

func (b *bw) ensureFlush() (int, error) {
	b.drained = true
	if len(b.buf) > 0 {
		return b.origin.Write(b.buf)
	}
	return 0, nil
}
