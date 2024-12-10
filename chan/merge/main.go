package main

import (
	"fmt"
	"sync"
)

func merge(chs ...<-chan int) <-chan int {
	ch := make(chan int)
	wg := sync.WaitGroup{}

	for _, c := range chs {
		wg.Add(1)
		go func(c <-chan int) {
			defer wg.Done()
			for v := range c {
				ch <- v
			}
		}(c)
	}
	go func() {
		wg.Wait()
		close(ch)
	}()

	return ch
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			ch1 <- i
		}
		close(ch1)
	}()

	go func() {
		for i := 10; i < 20; i++ {
			ch2 <- i
		}
		close(ch2)
	}()

	mergedCh := merge(ch1, ch2)

	for v := range mergedCh {
		fmt.Println(v)
	}
}
