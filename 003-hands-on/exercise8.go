package main

import (
	"fmt"
)

const (
	a     = 42.01
	b int = 42
)

func main() {
	fmt.Println(a, "\n", b)
	fmt.Printf("%T\n%T\n", a, b)
}
