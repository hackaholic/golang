package main

import (
	"fmt"
	"strings"
)

func add(x int, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func sum_n(n int) (x, y int) {
	/*
	   A return statement without arguments returns the named return values. This is known as a "naked" return.
	*/
	x = 10 + n
	y = 5 + n
	return
}

func main() {
	fmt.Printf("Sum of 2+3=%d\n", add(2, 3))
	a, b := swap("Hello", "World")
	println(a, b)
	println(sum_n(5))

	/* Functions are values too. They can be passed around just like other values.
	Function values may be used as function arguments and return values.
	*/

	split_string := func(str, sep string) []string {
		return strings.Split(str, sep)
	}
	fmt.Println(split_string("Hello world I am golang!", " "))
}
