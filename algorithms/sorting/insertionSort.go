package main

import (
	"fmt"
)

func insertionSort(arr []int) []int {
	var i int = 1
	for i < len(arr) {
		key := arr[i]
		k := i - 1
		fmt.Println(i, k)
		for k > 0 && arr[k] > key {
			fmt.Println(k)
			arr[k+1] = arr[k]
			k -= 1
		}
		arr[k+1] = key
		i += 1
	}
	return arr
}

func main() {
	var arr = []int{-1, 0, 2, 4, 6, 1, 8, 6, 0, 4, 5}
	fmt.Println(insertionSort(arr))
}
