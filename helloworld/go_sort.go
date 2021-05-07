package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type byLen []string

// implementing sort.Interface to sort by length
func (s byLen) Len() int {
	return len(s)
}
func (s byLen) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}
func (s byLen) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func stringExist(arr []string, s string, pos int) bool {
	if pos < len(arr) && s == arr[pos] {
		return true
	}
	return false
}

func main() {
	arr_pet := []string{"Odin-Sama", "coco", "mini", "kk", "dora", "oreo", "simba"}
	arr_int := []int{1, 5, 7, 8, 3, 4}

	fmt.Println("Unsorted array of string:\n", arr_pet)

	// sort.Strings(arr []string) -> sort array of type string
	sort.Strings(arr_pet)
	fmt.Println("After Sort:\n", arr_pet)

	fmt.Println("Unsorted array of Int:\n", arr_pet)
	// sort.Ints(x []int)
	sort.Ints(arr_int)
	fmt.Println("After Sort:\n", arr_int)

	// sort.SearchStrings(a []string, x string) int , if x: return pos else return insert position of x
	// sort.SearchInts(a []int, x int) int
	println("Searching mini:", sort.SearchStrings(arr_pet, "mini"), stringExist(arr_pet, "mini", sort.SearchStrings(arr_pet, "mini")))
	println("Searching coco:", sort.SearchStrings(arr_pet, "coco"), stringExist(arr_pet, "coco", sort.SearchStrings(arr_pet, "coco")))
	println("Searching mini1:", sort.SearchStrings(arr_pet, "mini1"), stringExist(arr_pet, "mini1", sort.SearchStrings(arr_pet, "mini1")))
	println("Searching zzz:", sort.SearchStrings(arr_pet, "zzz"), stringExist(arr_pet, "zzz", sort.SearchStrings(arr_pet, "zzz")))

	// Custom Sort and search
	fmt.Println("Custom Sort and search")

	// sort by length of string
	sortByLength := func(i, j int) bool {
		return len(arr_pet[i]) < len(arr_pet[j])
	}

	//sort.Sort(byLen(arr_pet))
	//sort.Slice(arr_pet, func(i, j int) bool {
	//	return len(arr_pet[i]) < len(arr_pet[j])
	//})
	sort.Slice(arr_pet, sortByLength)
	fmt.Println(arr_pet)

	//Shuffling arr_pet
	rand.Seed(time.Now().UTC().UnixNano())
	// rand.Shuffle(n int, swap func(i, j int))
	rand.Shuffle(len(arr_pet), func(i, j int) { arr_pet[i], arr_pet[j] = arr_pet[j], arr_pet[i] })
	fmt.Println("Shuffled arry_pet:", arr_pet)

	type PetInfo struct {
		name string
		age  float32
	}

	petinfo := []PetInfo{
		{"coco", 2.5},
		{"mini", 3},
		{"dora", 1},
	}

	fmt.Println("Arrya with struct type -> petinfo\n", petinfo)

	// sort on age
	sort.Slice(petinfo, func(i, j int) bool {
		return petinfo[i].age < petinfo[j].age
	})

	fmt.Println("After sort on age petinfo\n", petinfo)

	fmt.Println(sort.Search(len(petinfo), func(i int) bool { return petinfo[i].name == "mini1" }))

}
