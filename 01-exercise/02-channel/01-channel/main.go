package main

import "fmt"

func main() {
	temp := make(chan int)
	go func(a, b int) {
		c := a + b
		temp <- c
	}(1, 2)
	c := <-temp
	fmt.Printf("computed value %v\n", c)
}
