package main

import (
	"fmt"
)

func main() {
	favSport := "surfing"
	switch favSport {
	case "skkiing":
		fmt.Println("go to the mountains!")
	case "swimming":
		fmt.Println("got o the pool!")
	case "surfing":
		fmt.Println("go to hawaii!")
	case "tennis":
		fmt.Println("go to the court!")
	}
}
