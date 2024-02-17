package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	// what is the output 4 4 4

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(i)
		}()
		wg.Wait()
	}
}
