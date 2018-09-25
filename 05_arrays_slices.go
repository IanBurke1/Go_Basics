package main

import "fmt"

func main() {
	// Arrays
	// Declaring an array of size 2 and type string
	// var fruitArr [2]string

	// Assign values to array index
	// fruitArr[0] = "Apple"
	// fruitArr[1] = "Orange"

	//Declare and Assign - short way of assigning
	fruitArr := [2]string{"Apple", "Orange"}

	//print out the array
	fmt.Println(fruitArr)
	//print out element at index 1 in array
	fmt.Println(fruitArr[1])

	//Slices
	fruitSlice := []string{"Apple", "Orange", "Grape"}

	fmt.Println(fruitSlice)
	fmt.Println(len(fruitSlice)) //length of array
	fmt.Println(fruitSlice[1:2]) //starts at 1 and stops before 2. prints out 1 and everything else before 2

}
