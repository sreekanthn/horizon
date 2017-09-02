package main

import "fmt"
import "os"

func main(){

    s:=""
    sep:=""
    for i, arg := range os.Args[1:] {
      fmt.Println(i)
      s += sep + arg
      sep=" "
    }
    fmt.Println(os.Args[0])
    fmt.Println(s)

}
