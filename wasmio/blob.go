package wasmio

import (
	"io"
	"syscall/js"
)

type Blob struct {
	b js.Value
}

func NewBlobReader(b js.Value) io.Reader {
	return &Blob{
		b: b,
	}
}

var _ io.Reader = (*Blob)(nil)

// This doesn't work on go 1.12.6
func (b *Blob) Read(p []byte) (int, error) {
	promise := b.b.Call("read", js.TypedArrayOf(p))

	done := make(chan struct{})
	var (
		cb    js.Func
		nread int
		eof   bool
	)
	cb = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		go func() {
			readResult := args[0]
			nread = readResult.Get("nread").Int()
			eof = readResult.Get("eof").Bool()
			close(done)
		}()
		return nil
	})
	promise.Call("then", cb)
	<-done
	cb.Release()

	var err error
	if eof {
		err = io.EOF
	}
	return nread, err
}
