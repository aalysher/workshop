package main

import (
	"fmt"
	"sync"
)

func main() {
	urls := make([]string, 1000000000)
	codes := make(map[int]int)

	mu := sync.Mutex{}
	wg := sync.WaitGroup{}

	for _, url := range urls {
		wg.Add(1)
		go func() {
			defer wg.Done()
			code := sendRequest(url)
			mu.Lock()
			codes[code]++
			mu.Unlock()
		}()
	}

	wg.Wait()
	fmt.Println(codes)
}

// with worker pool pattern
func _() {
	var (
		urls  = make([]string, 1000000000)
		codes = make(map[int]int)

		wg = sync.WaitGroup{}
		mu = sync.Mutex{}

		ch = make(chan string)
	)

	go func() {
		for _, url := range urls {
			ch <- url
		}
		close(ch)
	}()

	wg.Add(5)
	for range 5 {
		go func() {
			defer wg.Done()
			for url := range ch {
				code := sendRequest(url)

				mu.Lock()
				codes[code]++
				mu.Unlock()
			}
		}()
	}

	wg.Wait()
	fmt.Println(codes)
}

// Semaphore
func _() {
	var (
		urls  = make([]string, 1000000000)
		codes = make(map[int]int)
		sema  = make(chan struct{}, 5)
		mu    = sync.Mutex{}
	)

	for _, url := range urls {
		sema <- struct{}{}
		go func() {
			defer func() {
				<-sema
			}()
			code := sendRequest(url)

			mu.Lock()
			codes[code]++
			mu.Unlock()
		}()
	}

}

func sendRequest(url string) (code int) {
	panic("implement me")
}
