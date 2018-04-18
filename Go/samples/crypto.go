package main

import (
	"crypto/sha256"
	"fmt"

	"github.com/sreekanthn/horizon/Go/samples/crypt"
)

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	diffBytes := compare(&c1, &c2)
	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
	fmt.Printf("The number of different bytes %d \n", diffBytes)
	fmt.Println(crypt.Compute())
}

func compare(ptr1 *[32]byte, ptr2 *[32]byte) int {
	diffBytes := 0
	for i := range ptr1 {
		if !(ptr1[i] == ptr2[i]) {
			diffBytes++
		}
	}
	return diffBytes
}
