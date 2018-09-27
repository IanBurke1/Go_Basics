package main

import (
	"fmt"
	"strconv"
)

// Define person struct
type Person struct {
	firstName string
	lastName  string
	city      string
	gender    string
	age       int
}

// Greeting method (value reciever)
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age) //change int to string
}

// hasBirthday method (pointer reciever)
func (p *Person) hasBirthday() {
	p.age++
}

// getMarried (pointer reciever)
func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "m" {
		return
	} else {
		p.lastName = spouseLastName
	}
}

func main() {
	// Init person using struct
	person1 := Person{firstName: "Slim", lastName: "Shady", city: "Detroit", gender: "m", age: 25}
	// Alternative
	//person1 := Person{"Derrick", "Zoolander", "Boston", "m", 25}
	person2 := Person{"Britney", "Spears", "Boston", "f", 25}
	// fmt.Println(person1)
	// fmt.Println(person1.firstName) //get firstName
	// person1.age++                  //change the age
	person1.hasBirthday()

	//change lastname if got married
	person1.getMarried("Williams")
	person2.getMarried("Williams")

	fmt.Println(person1.greet())
	fmt.Println(person2.greet())

}
