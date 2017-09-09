package main

import "crypto/sha256"
import "crypto/sha512"
import "fmt"

func main() {
    c1 := sha256.Sum256([]byte("x"))
    c2 := sha256.Sum256([]byte("X"))

    fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)

    hash3 := sha512.Sum512([]byte("Quick brown fox"))

    fmt.Printf("%x\n", hash3)

}
