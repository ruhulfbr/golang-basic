// Go program to illustrate
// how to trim a string
package main

import (
	"fmt"
	"strings"
)

// Main method
func main() {

	// Creating and initializing string
	// Using shorthand declaration
	str1 := "!!Welcome to     GeeksforGeeks !!"
	str2 := "@@This is the tutorial of Golang$$"

	// Displaying strings
	fmt.Println("Strings before trimming:")
	fmt.Println("String 1: ", str1)
	fmt.Println("String 2:", str2)

	// Trimming the given strings
	// Using Trim() function
	res1 := strings.TrimSpace(str1)
	res2 := strings.Trim(str2, "@$")

	trimmed := strings.Join(strings.Fields(str1), " ")

	// Displaying the results
	fmt.Println("\nStrings after trimming:")
	fmt.Println("Result 1: ", res1)
	fmt.Println("Result 2:", res2)
	fmt.Println("trimmed 2:", trimmed)

	y := 458

	// taking a pointer variable using
	// := by assigning it with the
	// address of variable y
	p := &y

	fmt.Println()
	fmt.Println("Value stored in y = ", y)
	fmt.Println("Address of y = ", &y)
	fmt.Println("Value stored in pointer variable p = ", p)

	fmt.Println()

	fmt.Println("Value stored in y(*p) = ", *p)

	*p = 500

	fmt.Println("Value stored in y(*p) after Changing = ", y)

	fmt.Println()

	var x = 100

	fmt.Printf("The value of x before function call is: %d\n", x)

	var pa *int = &x

	ptf(pa)
	fmt.Printf("The value of x after function call is: %d\n", x)
}

func ptf(a *int) {
	*a = 7480
}
