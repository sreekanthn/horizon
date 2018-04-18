package main

import "fmt"
import "os"

func main() {

	var s, separator string

	for i := 1; i < len(os.Args); i++ {
		s += separator + os.Args[i]
		separator = " "
	}
	fmt.Println(os.Args[0])
	fmt.Println(s)

}
