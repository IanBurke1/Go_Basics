package main

import "fmt"

func main() {
	x := 5
	y := 10

	// and -> &&
	// or -> ||

	// If else
	if x <= y {
		fmt.Println("%d is less than or equal to %d\n", x, y)
	} else {
		fmt.Println("%d is less than %d\n", y, x)
	}

	// else if
	color := "red"

	if color == "red" {
		fmt.Println("color is red")
	} else if color == "blue" {
		fmt.Println("color is blue")
	} else {
		fmt.Println("color is NOT red or blue")
	}

	// Switch
	switch color {
	case "red":
		fmt.Println("color is red")
	case "blue":
		fmt.Println("color is blue")
	default:
		fmt.Println("color is NOT red or blue")
	}
}
