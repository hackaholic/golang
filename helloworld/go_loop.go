package main

import (
	"fmt"
)

func main() {
	sum, n := 0, 10
	// for loop
	for i := 1; i <= n; i++ {
		fmt.Printf("i -> %d\n", i)
		sum += i
	}
	fmt.Printf("\nsum -> %d\n", sum)

	/*
		    For is Go's "while"
			At that point you can drop the semicolons: C's while is spelled for in Go.
			for ever loop
			for {
				statements
			}
	*/
	sum = 1
	for sum < 100 {
		sum += sum
	}
	fmt.Println(sum)

}
