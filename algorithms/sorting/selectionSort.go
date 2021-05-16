package main

import (
	"fmt"
)

func selectionSort(arr []int) []int {
	var i, j, min int
	for i < len(arr) {
		j, min = i, i
		for j < len(arr) {
			if arr[min] > arr[j] {
				min = j
			}
			j += 1
		}
		if i != min {
			arr[i], arr[min] = arr[min], arr[i]
		}
		i += 1
	}
	return arr
}

func main() {
	var arr = []int{-1, 0, 2, 4, 6, 1, 8, 6, 0, 4, 5}
	fmt.Println(selectionSort(arr))
}
