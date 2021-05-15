package main

import (
	"fmt"
)

/*
Dynamic Programming is an algorithmic paradigm that solves a given complex problem by breaking it into subproblems
and stores the results of subproblems to avoid computing the same results again. Following are the two main properties
of a problem that suggests that the given problem can be solved using Dynamic programming.
1) Overlapping Subproblems

fibonacci series -> fib(n) = fib(n-1) + fib(n-2)


                         fib(5)
                     /             \
               fib(4)                fib(3)
             /      \                /     \
         fib(3)      fib(2)         fib(2)    fib(1)
        /     \        /    \       /    \
  fib(2)   fib(1)  fib(1) fib(0) fib(1) fib(0)
  /    \
fib(1) fib(0)

we can see from above that fib(3), fb(2), fib(1). fib(0) are computed many times.
Here instead of computing the same operation again we store the computed value.

here we can create a lookup table to store computed value

 Memorization(top down)
 The memoized program for a problem is similar to the recursive version with
 a small modification that it looks into a lookup table before computing solutions

 Tabulation (Bottom Up)
 The tabulated program for a given problem builds a table in bottom up fashion and returns the last entry from table
*/

func fibTopDown(n int, arr []int) int {
	fmt.Println(n, arr)
	if n == 0 || n == 1 {
		arr[n] = n
		return n
	}
	if arr[n] == 0 {
		arr[n] = fibTopDown(n-1, arr) + fibTopDown(n-2, arr)
	}
	fmt.Println("returning:", arr[n])
	return arr[n]
}

func fibBottomUp(n int, arr []int) int {
	if n == 0 || n == 1 {
		return n
	}
	arr[1] = 1
	i := 2
	for i <= n {
		fmt.Println(arr)
		arr[i] = arr[i-1] + arr[i-2]
		i += 1
	}
	return arr[n]
}

func main() {
	fmt.Println("Fibonacci series using dynamic programming Memorization(top Down)")
	n := 10
	var arr = make([]int, n+1)
	fmt.Println(fibTopDown(n, arr))
	fmt.Println(arr)

	fmt.Println("Fibonacci series using dynamic programming Memorization(top Down)")
	var arr1 = make([]int, n+1)
	fmt.Println(fibBottomUp(n, arr1))
	fmt.Println(arr1)

}
