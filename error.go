package main

import (
	"errors"
	"fmt"
)

func Exception(message string) error {
	return errors.New(message)
}

func riskyFunction(message string) {
	panic(message)
}

func handlePanic() {
	if r := recover(); r != nil {
		fmt.Println("Recovered from panic:", r)
	}
}

func main() {
	err := Exception("No data found")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("ok")
	}

	defer handlePanic()

	riskyFunction("Keu kisu bole na")

	fmt.Println("This will not print")
}
