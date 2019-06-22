package blob

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

func (b *Blob) Read(p []byte) (int, error) {
	promise := b.b.Call("read", js.TypedArrayOf(p))

	var (
		cb    js.Func
		nread int
		eof   bool
	)
	done := make(chan struct{})
	cb = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		go func() {
			readResult := args[0]
			nread = readResult.Get("nread").Int()
			eof = readResult.Get("eof").Bool()
			close(done)
			cb.Release()
		}()
		return nil
	})
	promise.Call("then", cb)
	<-done

	var err error
	if eof {
		err = io.EOF
	}
	return nread, err
}
