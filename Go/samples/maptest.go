package main

import (
	"fmt"
	"sort"
	"time"
)

func main() {
	agemap := make(map[string]int)
	agemap["sreekanth"] = 39
	agemap["suja"] = 38
	agemap["sruthy"] = 12

	t, _ := time.Parse(time.UnixDate, "2018 02 14 00 00")
	if time.Now().Local().After(t) {
		agemap["sruthy"]++
	}

	fmt.Printf("Age of Sruthy %d \n", agemap["sruthy"])

	var names []string
	for name := range agemap {
		names = append(names, name)
	}

	sort.Strings(names)

	for _, name := range names {
		fmt.Println(name)
	}
}
