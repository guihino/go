package main

import (
	"fmt"
)

func main() {
	x := 42
	if x == 43 {
		fmt.Println("the x is not 42")
	} else {
		fmt.Println("the x is 42")
	}
}
