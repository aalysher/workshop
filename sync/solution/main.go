package main

import (
	"fmt"
	"sync"
)

// print square
func main() {
	counter := 20
	wg := sync.WaitGroup{}
	for i := 0; i < counter; i++ {
		wg.Add(1)
		go func(in int) {
			defer wg.Done()
			fmt.Println(in * in)
		}(i)
	}

	wg.Wait()
}
