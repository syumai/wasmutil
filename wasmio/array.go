package wasmio

import (
	"io"
	"syscall/js"
)

type Array struct {
	b js.Value
}

func NewArrayReader(b js.Value) io.Reader {
	return &Array{
		b: b,
	}
}

var _ io.Reader = (*Array)(nil)

func (a *Array) Read(p []byte) (int, error) {
	readResult := a.b.Call("read", js.TypedArrayOf(p))
	n := readResult.Get("nread").Int()
	eof := readResult.Get("eof").Bool()

	var err error
	if eof {
		err = io.EOF
	}
	return n, err
}
