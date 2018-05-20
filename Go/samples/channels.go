package main

import (
	"fmt"
)

// Note the way chanels are defined
// out is defined as chan<- indicating that this channel is a send only Channel (Can send values to this Channel)
// in id defined as <-chan indicating that this channel is a Receive Only channel
// (Can be read from)
// allowing it to do the range operations on the same

func counter(out chan<- int) {
	for x := 0; x <= 100; x++ {
		out <- x
	}
	close(out)
}

func squarer(out chan<- int, in <-chan int) {
	// Read from the in channel and provide the square to the out channel
	for val := range in {
		out <- val * val
	}
	close(out)

}

func printer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}

}
func main() {
	numbers := make(chan int)
	squares := make(chan int)

	// start subroutines
	go counter(numbers)
	go squarer(squares, numbers)
	printer(squares)
}
