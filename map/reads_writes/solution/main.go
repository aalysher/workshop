package main

import (
	"fmt"
	"sync"
)

func main() {
	ops := 1000
	storage := make(map[int]int, ops)

	wg := sync.WaitGroup{}
	mu := sync.RWMutex{}

	wg.Add(ops)

	for i := 0; i < ops; i++ {
		go func() {
			defer wg.Done()

			mu.Lock()
			defer mu.Unlock()
			storage[i] = i
		}()
	}
	wg.Add(ops)
	for i := 0; i < ops; i++ {
		go func() {
			defer wg.Done()

			mu.RLock()
			defer mu.RUnlock()
			_, _ = storage[i]
		}()
	}

	wg.Wait()
	fmt.Println(storage)
}
