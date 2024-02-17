package main

import (
	"fmt"
	"time"
)

func fun(s string) {
	for i := 0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(1 * time.Millisecond)
	}
}

func main() {
	// Direct call
	fun("direct call")

	// goroutine function call
	go fun("goroutine call")

	// goroutine with anonymous function
	go func() {
		fmt.Println("Anonymous function")
	}()

	// goroutine with function value call
	functionValue := fun
	go functionValue("Test")

	// wait for goroutines to end
	time.Sleep(5 * time.Millisecond)
	fmt.Println("done..")
}
