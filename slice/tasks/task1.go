package main

import "fmt"

func main() {
	slice := make([]string, 3, 4) // len = 3, cap = 4
	fmt.Println(slice)            // val = ["", "", ""]

	appendSlice(slice)
	fmt.Println(slice) // ["", "", ""]

	mutateSlice(slice)
	fmt.Println(slice) // [world, "", ""]
}

func appendSlice(s []string) {
	s = append(s, "hello")
}

func mutateSlice(s []string) {
	s[0] = "world"
}
