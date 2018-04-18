package main

import (
	"fmt"
	"os"
	"strings"
)

// print the base name of the File given Stripping out other details from the full path

func main() {
	//fmt.Println(len(os.Args))
	if len(os.Args) <= 1 {
		fmt.Println("Usage : basename <filepath1> <filepath2> ....")
		os.Exit(1)
	}
	for _, name := range os.Args[1:] {
		trimName := findBaseName(name)
		fmt.Printf("%s\n", trimName)
	}
}

func findBaseName(fileName string) string {
	//find index of the last slash in the path
	slashIndex := strings.LastIndex(fileName, "/")
	if slashIndex != -1 {
		fileName = fileName[slashIndex+1:]
	}

	// check if have any dots in the file name for the file extension
	dotIndex := strings.LastIndex(fileName, ".")
	if dotIndex >= 0 {
		fileName = fileName[:dotIndex]
	}
	return fileName
}
