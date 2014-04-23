// Interfaces are satisfied implicitly
// メソッドを実装する事で暗黙的にインターフェースを継承出来るようになる
// Scalaのトレイトに相当

package main

import (
	"fmt"
	"os"
	// "io"
)

type Reader interface {
	Read(b []byte) (n int, err error)
}

type Writer interface {
	Write(b []byte) (n int, err error)
}

type ReadWrite interface {
	Reader
	Writer
}

func main() {
	// var w io.Writer
	var w Writer

	// os.Stdout implements Writer
	w = os.Stdout

	fmt.Fprintf(w, "hello, writer\n")
}
