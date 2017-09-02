package main

import (
    "bufio"
    "fmt"
    "os"
)

func main(){
  counts:=make(map[string]int)
  input:=bufio.NewScanner(os.Stdin)
  for input.Scan(){
    line:=input.Text()
    counts[line]++
  }

  fmt.Println("Done function")

  // now check the map to see if there are dupliacated lines
  for line, n := range counts {
    if n >1 {
      fmt.Printf("Lines Containing duplicate : %d\t%s\n", n, line)
    }
  }
}
