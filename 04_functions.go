package main

import "fmt"

// Basic functions that takes in arguments and declares a return type just like Java
func greeting(name string) string {
	return "Hello " + name
}

func getSum(num1 int, num2 int) int {
	return num1 + num2
}

func main() {
	// Calling the functions and passing in data
	fmt.Println(greeting("Ian"))
	fmt.Println(getSum(3, 4))
}
