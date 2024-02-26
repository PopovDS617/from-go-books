package interfaces

import (
	"bytes"
	"fmt"
	"io"
)

var debug = false

func NullishInterfaces() {
	// var buf *bytes.Buffer  => will panic
	var buf io.Writer

	if debug {
		buf = new(bytes.Buffer)
	}
	f(buf)

	if debug {
		fmt.Println(buf)
	}

}
func f(out io.Writer) {
	if out != nil {
		out.Write([]byte("done"))
	}

}
