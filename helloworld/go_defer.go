// A defer statement defers the execution of a function until the surrounding function returns.

package main

import (
	"fmt"
)

func test() {
	// Deferred function calls are pushed onto a stack. When a function returns,
	// its deferred calls are executed in last-in-first-out order.
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
}
func main() {
	defer fmt.Println("Main executed successfully")
	fmt.Println("World World")
	test()
}
