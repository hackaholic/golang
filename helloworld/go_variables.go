package main

import (
	"fmt"
)

var a, b bool

func main() {
	var x, y, z = "hello", "world", false
	var i int = 10
	fmt.Println(a, b, x, y, z, i)

	/*
		Inside a function, the := short assignment statement can be used in place of a var declaration with implicit type.
	*/
	j := 10
	fmt.Println(j)

	/*
			 Variables declared without an explicit initial value are given their zero value.
		    The zero value is
		                  0 for numeric types,
		                  false for the boolean type, and
		                 "" (the empty string) for strings.
	*/

	var k int      // 0
	var str string // ""
	var flag bool  // false
	fmt.Println("Variables declared without an explicit initial value are given their zero value.")
	fmt.Printf("k -> %d, str -> %s, flag -> %b\n", k, str, flag)

	// type conversion T(v)
	n := float64(i)
	fmt.Printf("Type of n is %T\n", n)

	/*
			 Constants are declared like variables, but with the const keyword.
		     Constants can be character, string, boolean, or numeric values.
		     Constants cannot be declared using the := syntax.
	*/
	const word = "Hello Golang"
	fmt.Printf(" word of type: %T, value: %s\n", word, word)
	// word = "Hello World" this will fail
}
