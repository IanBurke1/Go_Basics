package main

import "fmt"

func main() {
	// MAIN TYPES:
	// string
	// bool
	// int
	// int  int8  int16  int32  int64 -
	// uint uint8 uint16 uint32 uint64 uintptr - unsigned int is 0+.. no negative numbers
	// byte - alias for uint8
	// rune - alias for int32
	// float32 float64
	// complex64 complex128 - for really large numbers

	//using var
	//var name = "Ian"
	var age = 24
	const isCool = true //const = constant
	//var size float32 = 2.3

	//Shorthand assignment. Auto assign name a string
	name := "Ian"
	size := 2.3

	fmt.Println(name, age, isCool, size)
	//get the type
	fmt.Printf("%T\n", size)
}

// To run the program enter: go run main.go
// To create an executable file enter: go build
// The exe file will be stored in the bin folder.
