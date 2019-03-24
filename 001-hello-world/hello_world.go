package main

import "fmt"

func main() {
	fmt.Println("Hello world")

	foo()

	fmt.Println("something else")

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	bar()
}

func bar() {
	fmt.Println("and then we exited")
}

func foo() {
	fmt.Println("I'm in foo")
}
