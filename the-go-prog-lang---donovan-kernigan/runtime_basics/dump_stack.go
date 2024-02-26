package runtimebasics

import (
	"os"
	"runtime"
)

func PrintStackDump() {
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	os.Stdout.Write(buf[:n])
}
