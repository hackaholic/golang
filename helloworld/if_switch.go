package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i, "is even")
		} else {
			fmt.Println(i, "is odd")
		}
	}

	// Switch cases evaluate cases from top to bottom, stopping when a case succeeds
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("you are using OS X.")
	case "linux":
		fmt.Println("You are using Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}

	// Switch without a condition is the same as switch true.
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}
