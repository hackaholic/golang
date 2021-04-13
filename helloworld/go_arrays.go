package main

import (
	"fmt"
)

func main() {
	/*
			Arrays
		    The type [n]T is an array of n values of type T.
	*/
	var arr = [2]string{"Hello", "world"}
	fmt.Println(arr)
	fmt.Println("Accessing valuing using index 0 ->", arr[0])
	fmt.Println("Accessing valuing using index 1 ->", arr[1])

	var arr_n [10]int // all value are initialized with zero becauese of int, "" for string
	fmt.Println("arr_n initialized value", arr_n)
	fmt.Println("Length of arr_n ->", len(arr_n))

	for i := 0; i < len(arr_n); i++ {
		arr_n[i] = 2 * i
	}

	fmt.Println("arr_n ->", arr_n)

	// array slices array[low:high] , high will be excluded
	// A slice does not store any data, it just describes a section of an underlying array.
	// Changing the elements of a slice modifies the corresponding elements of its underlying array.
	fmt.Println("arr_n[1:5] ->", arr_n[1:5])
	fmt.Println("arr_n[4:]", arr_n[4:])

	// looping over arry using range
	// When ranging over a slice, two values are returned for each iteration.
	// The first is the index, and the second is a copy of the element at that index.
	fmt.Println("Looping over array arr_pet")
	var arr_pet = []string{"coco", "Dora", "Mini", "Oreo", "Nova"}
	for i, v := range arr_pet {
		fmt.Println(i, v)
	}

	fmt.Println("Len of arr_pet:", len(arr_pet), "Capacity ->", cap(arr_pet))

	/*
			  you can skip the index or value by assigning to _
			  for i, _ := range arr_pet
		      for _, value := range
			  If you only want the index, you can omit the second variable.
		      for i := range pow
	*/

	// aapend vale to array
	fmt.Println("Appending Simba to arr_pet")
	arr_pet = append(arr_pet, "Simba")
	fmt.Println(arr_pet)
}
