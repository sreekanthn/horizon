package main

import (
    "fmt"
    "io"
    "net/http"
    "os"
)

func main(){
  for _, url := range os.Args[1:]{
    resp, err := http.Get(url)
    if err!=nil{
      fmt.Println("Error Occured !! %v", err)
      os.Exit(1)
    }
    respbody, err := io.Copy(os.Stdout, resp.Body)
    resp.Body.Close()
    if err!=nil{
      fmt.Println("Error Occured while reading url %s !! %v\n", url, err)
      os.Exit(1)
    }
    fmt.Printf("%s", respbody)
  }
}
