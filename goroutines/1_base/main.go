package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		fmt.Println("Hello, world!")
	}()

	go printHello()

	var p printer
	p.printHello()

	time.Sleep(time.Second)
}

func printHello() {
	fmt.Println("Hello, world from function")
}

type printer struct{}

func (p *printer) printHello() {
	fmt.Println("Hello, world from method")
}
