package main

import (
	"fmt"
)

func main() {
	i, j := 0, 1
	n := 30
	for counter := 0; counter < n; counter++ {
		i, j = j, i+j // tuple assignment makes it simple
		fmt.Printf("%d\t", i)
	}
	fmt.Println()
}
