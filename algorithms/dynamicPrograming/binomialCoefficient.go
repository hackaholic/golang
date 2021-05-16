package main

import "fmt"

/* write a program to calculate binomialCofficient given n and r
n >= 1, 0=< r <= n
C(n, 0) = 1
C(n, n) = 1

C(n,r) = C(n-1, r) + C(n-1, r-1)

*/

func binomialCoefficient(n, r int) int {
	if n < 1 {
		return 0
	}
	if n == r || r == 0 {
		return 1
	}
	return binomialCoefficient(n-1, r) + binomialCoefficient(n-1, r-1)
}

func binomialCoefficientDynamic(n, r int, arr [][]int) int {
	if n < r {
		return 0
	}
	for i := 1; i <= n; i++ {
		for j := 0; j <= i && j <= r; j++ {
			if j == 0 {
				arr[i][j] = 1
			} else if i == j {
				arr[i][j] = 1
			} else {
				arr[i][j] = arr[i-1][j] + arr[i-1][j-1]
			}
		}
	}
	return arr[n][r]
}

func main() {
	n, r := 50, 10
	var arr = make([][]int, n+1)
	for i := 0; i <= n; i++ {
		arr[i] = make([]int, r+1)
	}
	fmt.Println("Evaluating Binomial Coefficient ")
	// without dynamic programming
	//fmt.Println("C", n, r, binomialCoefficient(n, r))
	fmt.Println("Using Dynamic Programming")
	fmt.Println("C", n, r, binomialCoefficientDynamic(n, r, arr))

	//fmt.Println(arr)
}
