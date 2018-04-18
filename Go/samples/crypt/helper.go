package crypt

import (
	"fmt"
	"math/rand"
)

var i, t int

//sriting this in a package to see how the usage of exported
//initializing variabes using the init function
func init() {
	i = 100
	t = 60
}

//Compute () - Dummy function exported
func Compute() int64 {
	secureRand := rand.Int63()
	fmt.Println((int64(i*t) * secureRand))
	//return i * t * secureRand // Remember - there is no auto boxing !!
	return (int64(i*t) * secureRand) % int64(56)
}
