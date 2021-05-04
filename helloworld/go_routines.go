package main

import (
	"fmt"
	"time"
)

func goRoutine(s string) {
	println("Go Routine Test")
	for i := 0; i < 5; i++ {
		fmt.Println(s)
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	go goRoutine("Hello World")
	println("I am Called after goRoutine")
	// sleep for 5 sec
	time.Sleep(5 * time.Second) // time.Second gives 1 sec in microsecond 1000000000 -> 1 sec
}
