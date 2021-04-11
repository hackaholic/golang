package main

import (
	"fmt"
)

func main() {
	// Struct is a collection of fields
	type Vertex struct {
		x int
		y int
	}
	var v Vertex = Vertex{1, 2}

	var (
		v1 = Vertex{1, 2}  // has type Vertex
		v2 = Vertex{x: 1}  // Y:0 is implicit
		v3 = Vertex{}      // X:0 and Y:0
		p  = &Vertex{1, 2} // has type *Vertex
	)

	fmt.Println(v)
	fmt.Println(v1, p, v2, v3)

	// Struct fields are accessed using a dot.
	fmt.Println("Value of v.x ->", v.x)
	v.x = 10
	fmt.Println("New value of v.x ->", v.x)

	// Pointers to structs
	var st_p *Vertex = &v
	fmt.Println("Address of v ->", st_p)
	st_p.x = 1e5
	fmt.Println("New value of v.x", v.x)
}
