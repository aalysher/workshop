package main

import (
	"fmt"
	"time"
)

func main() {
	counter := 20
	for i := 0; i < counter; i++ {
		go func() {
			fmt.Println(i * i)
		}()
	}

	time.Sleep(time.Second)
}
