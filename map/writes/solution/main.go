package main

import (
	"fmt"
	"sync"
)

func main() {
	writes := 1000
	storage := make(map[int]int, writes)

	wg := sync.WaitGroup{}
	mu := sync.Mutex{}

	wg.Add(writes)
	for i := 0; i < writes; i++ {
		go func() {
			defer wg.Done()
			mu.Lock()
			defer mu.Unlock()
			storage[i] = i
		}()
	}

	wg.Wait()
	fmt.Println(storage)
}
