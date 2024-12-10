package main

import "fmt"

type Item struct {
	value int
	next  *Item
}

type Stack struct {
	head *Item
}

func NewStack() Stack {
	return Stack{}
}

func (s *Stack) Push(value int) {
	s.head = &Item{value: value, next: s.head}
}

func (s *Stack) Pop() int {
	if s.head == nil {
		return -1
	}

	value := s.head.value
	s.head = s.head.next
	return value
}

func main() {
	stack := NewStack()
	stack.Push(10)
	stack.Push(20)
	stack.Push(30)

	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
}
