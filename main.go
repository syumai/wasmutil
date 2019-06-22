package main

import (
	"bytes"
	"fmt"
	"io"
	"syscall/js"

	"github.com/syumai/wasmutil/wasmio"
)

func readAll(ary js.Value) (string, error) {
	br := wasmio.NewArrayReader(ary)
	var buf bytes.Buffer
	if _, err := io.Copy(&buf, br); err != nil {
		return "", err
	}
	return buf.String(), nil
}

func main() {
	window := js.Global()
	window.Set("readAll", js.FuncOf(func(_ js.Value, args []js.Value) interface{} {
		result, err := readAll(args[0])
		if err != nil {
			panic(err)
		}
		fmt.Println(result)
		return nil
	}))
	select {}
}
