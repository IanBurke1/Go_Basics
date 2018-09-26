package main

import "fmt"

func main()  {
	// for loop - Long method
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++ 
	}

	// for loop - short method
	for i := 1; i <= 10; i++ {
		fmt.Println("Number %d\n" i)
	}

	// FizzBuzz
	for i := 1; i <= 100; 1++ {
		if i % 15 == 0 {
			fmt.Println("FizzBuzz")
		} else if i % 3 == 0 {
			fmt.Println("Fizz")
		} else if i % 5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}