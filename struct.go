package main

import (
	"fmt"
)

type Person struct {
	name   string
	age    int
	job    string
	salary int
}

type student struct {
	int
	string
	float64
}

func main() {
	var pers1 Person
	var pers2 Person

	// Pers1 specification
	pers1.name = "Hege"
	pers1.age = 45
	pers1.job = "Teacher"
	pers1.salary = 6000

	// Pers2 specification
	pers2.name = "Cecilie"
	pers2.age = 24
	pers2.job = "Marketing"
	pers2.salary = 4500

	printPerson(pers1)
	fmt.Println()
	printPerson(pers2)

	value := student{123, "Bud", 8900.23}

	// Display the values of the fields
	fmt.Println()
	fmt.Println("Enrollment number : ", value.int)
	fmt.Println("Student name : ", value.string)
	fmt.Println("Package price : ", value.float64)
}

func printPerson(pers Person) {
	fmt.Println("Name: ", pers.name)
	fmt.Println("Age: ", pers.age)
	fmt.Println("Job: ", pers.job)
	fmt.Println("Salary: ", pers.salary)
}
