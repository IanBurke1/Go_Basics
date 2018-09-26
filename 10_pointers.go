package main

import "fmt"

func main() {
	// Pointers allows us to point to memory address/location of a value in a variable
	a := 5
	b := &a //Assigning b to a pointer of a

	fmt.Println(a, b) // Expected output: 5 0xc..(memory address)
	// find the type of b..
	fmt.Println("%T\n", b) // *int represents a pointer

	// Use * to read val from address
	fmt.Println(*b)  // Expected output: 5
	fmt.Println(*&a) // Expected output: 5

	// Change val with pointer
	*b = 10
	fmt.Println(a) // a = 10

	// Everything in Go is passed by value
}
