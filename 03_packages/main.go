package main

// Imports can be written in the following way:
// fmt stands for format
// math imports all math classes
// importing strutil package which contains reverse str function
import (
	"fmt"
	"math"

	"github.com/ianburkeixiv/go_basics/03_packages/strutil"
)

func main() {
	fmt.Println(math.Floor(2.7)) //round to last whole number i.e. 2
	fmt.Println(math.Ceil(2.7))  //round to next whole number i.e. 3
	fmt.Println(math.Sqrt(16))   //square root

	// using our imported strutil Reverse function that we created
	fmt.Println(strutil.Reverse("hello"))

}
