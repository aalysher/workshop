package main

import (
	"fmt"
	"runtime"
	"time"
)

func main_() {
	ch := make(chan int)

	ch <- 1
}

func main() {
	fmt.Println(runtime.NumGoroutine())
	go func() {
		for i := 0; i < 5; i++ {
			time.Sleep(time.Second)
			fmt.Println("iter: ", i)
		}
		fmt.Println(runtime.NumGoroutine())
	}()
	select {}
	fmt.Println("finish")
}
