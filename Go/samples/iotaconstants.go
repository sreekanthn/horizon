package main

import (
	"fmt"
)

const (
	_ = 1 << (10 * iota)
	KiB
	MiB
	GiB
	TiB
	PiB
	EiB
	ZiB
	YiB
)

func main() {
	// It is interesting that you cannot print the value of ZiB directly as it overfolows an int

	fmt.Println(YiB / ZiB)
}
