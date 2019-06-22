package main

import (
	"bytes"
	"fmt"
	"io"
	"syscall/js"

	"github.com/syumai/wasmutil/blob"
)

func main() {
	window := js.Global()
	window.Set("readAll", js.FuncOf(func(_ js.Value, args []js.Value) interface{} {
		br := blob.NewBlobReader(args[0])
		var buf bytes.Buffer
		if _, err := io.Copy(&buf, br); err != nil {
			panic(err)
		}
		fmt.Println(buf.String())
		return nil
	}))
	select {}
}
