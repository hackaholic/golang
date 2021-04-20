/*
 Go does not have classes. However, you can define methods on types.
A method is a function with a special receiver argument.
*/

package main

import (
	"fmt"
	"math"
)

type Circle struct {
	r float64
}

func (c Circle) getArea() float64 {
	return math.Pi * (2 * c.r)
}

/*
	You can declare a method on non-struct types, too.
	In this example we see a numeric type MyFloat with an Abs method.
	You can only declare a method with a receiver whose type is defined in the same package as the method.
	You cannot declare a method with a receiver whose type
	is defined in another package (which includes the built-in types such as int).
*/
type myInt int

func (i myInt) abs() int {
	if i < 0 {
		return int(-i)
	}
	return int(i)
}

func main() {
	c := Circle{2}
	fmt.Println("Area of circle with raduis 2 ->", c.getArea())

	i := myInt(-67)
	fmt.Println("Absolute value of -67 ->", i.abs())
}
