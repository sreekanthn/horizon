package main

import (
    "fmt"
    "io"
    "io/ioutil"
    "net/http"
    "os"

)

func main(){
  ch:=make(chan string)
  for _, url:=range os.Args[1:]{
    go fetch (url, ch)
  }

  for range os.Args[1:]{
    fmt.Println (<-ch)
  }
}

func fetch(url string, ch chan<-string){
  resp, err := http.Get(url)
  if err!=nil{
    ch <- fmt.Sprint(err)
    return
  }

  noofbytes, err := io.Copy(ioutil.Discard, resp.Body)
  resp.Body.Close()
  if err!=nil{
    ch <- fmt.Sprintf("Error occured while reading from %s :: %v " , url, err)
    return
  }

  ch <- fmt.Sprintf("The number of bytes received from url %s is  %20d", url, noofbytes)
}
