package interfaces

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func DynamicInterfaces() {

	data := []byte("hello world")

	// 1
	var w io.Writer
	w.Write(data) // will panic

	// 2
	w = os.Stdout
	res, err := w.Write(data)
	fmt.Println(res, err)
	fmt.Printf("%T\n", w) // *os.File

	// 3
	w = new(bytes.Buffer)
	res, err = w.Write(data)
	fmt.Println(res, err)
	fmt.Printf("%T\n", w) // *bytes.Buffer

	// 4
	w = nil
	w.Write(data) // will panic

}
