package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	present := make(map[string]bool) // map of booleans whowing presence of the string
	files := os.Args[1:]             //Assign the files passed in the argument
	if !(len(files) == 1) {
		dedupText(os.Stdin, present)
	} else {
		for _, f := range files {
			file, err := os.Open(f)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			dedupText(file, present)
			file.Close()
		}
	}

	for line, _ := range present {
		fmt.Printf("%s\n", line)
	}

}

func dedupText(f *os.File, present map[string]bool) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		present[input.Text()] = true
	}
}
