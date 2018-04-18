package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		// check if the url has an http:// prefix
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}
		fetchText, err := http.Get(url)
		if err != nil {
			fmt.Printf("Unable to fetch the URL : %v", err)
			os.Exit(1)
		}
		body, err := ioutil.ReadAll(fetchText.Body)
		if err != nil {
			fmt.Printf("Error reading the body : %v", err)
			os.Exit(1)
		}
		fmt.Printf("%s", body)
	}
}
