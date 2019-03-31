package main

import (
	"fmt"
)

type lala int

var x lala

func main() {

	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)

}
