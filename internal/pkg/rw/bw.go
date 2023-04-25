package rw

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

func (w *bw) Write(p []byte) (int, error) {

	switch {
	case len(p)+len(w.buf) > cap(w.buf):
		var n, err = w.ensureFlush()
		if err != nil {
			return n, err
		}
		fallthrough

	case w.drained:
		return w.origin.Write(p)

	default:
		w.buf = append(w.buf, p...)
		return len(p), nil
	}
}

func (w *bw) Close() error {
	var _, err = w.ensureFlush()
	return err
}

func (w *bw) ensureFlush() (int, error) {
	w.drained = true
	if len(w.buf) > 0 {
		return w.origin.Write(w.buf)
	}
	return 0, nil
}
