package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	counts := make(map[string]int) // create a map to store the duplicated strings and the counts for each of them
	files := os.Args[1:]           //Assign the files passed in the argument
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		//as there are some files to process let us process them one-by-one
		for _, file := range files {
			f, error := os.Open(file)
			// handle if there is an error; else
			if error != nil {
				fmt.Fprintf(os.Stderr, "%v\n", error)
				continue // with the next file
			}
			countLines(f, counts)
			// close the file
			f.Close()
		}
	}

	for line, num := range counts {
		if num > 1 {
			fmt.Printf("%d\t%s\n", num, line)
		}
	}
}

// helper function to count the number of duplicate lines
func countLines(f *os.File, counts map[string]int) {
	readBuffer := bufio.NewScanner(f)
	for readBuffer.Scan() {
		counts[readBuffer.Text()]++
	}
}
