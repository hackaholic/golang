package main

import (
	"fmt"
)

// Go has pointers. A pointer holds the memory address of a value.
func main() {
	var i int = 10
	var p *int = &i // P is a pointer pointing to type int value
	fmt.Println("Address of i ->", p)
	// The * operator denotes the pointer's underlying value.
	fmt.Println("Accessing value of i using pointer ->", *p)
	// Unlike C, Go has no pointer arithmetic
}
