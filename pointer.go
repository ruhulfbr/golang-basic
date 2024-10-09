package main

import "fmt"

type Employee struct {
	name  string
	empid int
}

func changeValByRef(emp *Employee) {
	emp.empid = 69857
	emp.name = "New Name"
}

// Main Function
func main() {
	emp := Employee{"ABC", 19078}
	pts := &emp

	fmt.Println(pts)
	fmt.Println(pts.name)

	// Call a function with reference
	changeValByRef(pts)

	fmt.Println((*pts).name)
}
