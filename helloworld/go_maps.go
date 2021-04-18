package main

import (
	"fmt"
	"strings"
)

type Pet struct {
	name string
	age  float64
}

func WordCount(s string) map[string]int {
	wordcount := make(map[string]int)
	for _, v := range strings.Split(s, " ") {
		wordcount[v] += 1
	}
	return wordcount
}

func main() {
	/*
		 A map maps keys to values.
		The zero value of a map is nil. A nil map has no keys, nor can keys be added.
		The make function returns a map of the given type, initialized and ready for use.
	*/

	dict := make(map[int]Pet)

	dict[1] = Pet{"Coco", 2}
	dict[2] = Pet{"Mini", 3}
	dict[3] = Pet{"Dora", 2.5}
	dict[4] = Pet{"Oreo", 2.8}
	dict[5] = Pet{"Nova", 0.6}
	fmt.Println(dict)

	// Map literals are like struct literals, but the keys are required.
	dict_data := map[int]Pet{
		1: Pet{"Coco", 2},
		2: Pet{"Mini", 3},
		3: Pet{"Dora", 2.5},
		4: Pet{"Oreo", 2.8},
		5: Pet{"Nova", 0.6},
	}
	fmt.Println(dict_data)

	// accessing a value of a key
	fmt.Println("accessing a value of a key")
	fmt.Println("dict[2] ->", dict[2])
	fmt.Println("dict[4] ->", dict[4])

	// delete an element   delete(map, key)
	fmt.Println("Deleting key 2 ..")
	delete(dict, 2)
	fmt.Println(dict)

	// Test that a key is present or not ,
	_, flag := dict[2] // If key is in dict, flag is true. If not, flag is false.
	fmt.Println("key 2 is present in map dict?:", flag)

	fmt.Println("Word Count")
	fmt.Println("I am learning Go!", WordCount("I am learning Go!"))
	fmt.Println("I ate a donut. Then I ate another donut.", WordCount("I ate a donut. Then I ate another donut."))
	fmt.Println("A man a plan a canal panama.", WordCount("A man a plan a canal panama."))
	fmt.Println("coco dora coco mini dora oreo mini dora", WordCount("coco dora coco mini dora oreo mini dora"))

}
