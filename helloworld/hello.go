package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	/*  Note: The environment in which these programs are executed is deterministic,
	     so each time you run the example program rand.Intn will return the same number.
	    (To see a different number, seed the number generator
	*/
	rand.Seed(time.Now().UTC().UnixNano())
	fmt.Println("Hello World")
	fmt.Println("Current time:", time.Now())
	fmt.Println("1st Random number is", rand.Intn(100))
	fmt.Println("2st Random number is", rand.Intn(100))
	fmt.Println("Value of Pi:", math.Pi)
	fmt.Printf("Square Root of 7: %f\n", math.Sqrt(7))
}
