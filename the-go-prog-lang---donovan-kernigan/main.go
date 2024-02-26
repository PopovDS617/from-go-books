package main

import (
	"fmt"
	"tgpldk/oop"
)

func main() {
	d := oop.Data

	// oop.Data.internalData - cannot use it because it is hidden inside oop package

	sd := d.GetHash()

	fmt.Println(sd)

}
