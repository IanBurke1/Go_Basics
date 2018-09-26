package main

import "fmt"

func main() {
	// Define map
	//map[key type] value type
	// key = first names and value = email
	// emails := make(map[string]string)

	// Declare map and add key values
	emails := map[string]string{"Bob": "bob@gmail.com", "Joe": "Joe@gmail.com"}

	// Assign key
	// emails["Bob"] = "bob@gmail.com"
	// emails["Joe"] = "Joe@gmail.com"
	// emails["Mike"] = "mike@gmail.com"

	fmt.Println(emails)
	fmt.Println(len(emails)) //length of map
	fmt.Println(emails["Bob"])

	// Delete from map
	delete(emails, "Bob")
	fmt.Println(emails)
}
