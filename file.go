package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"time"
)

func createFile(filename string) *os.File {
	file, err := os.Create(filename)

	if err != nil {
		fmt.Println(err)

		return nil
	}

	fmt.Println("Successfully created file: ", file.Name())

	return file
}

func writeIntoFile(file *os.File, text string) {

	fmt.Println()
	fmt.Println("Writing text line into the file: ", text)

	size, errors := io.WriteString(file, text)

	if errors != nil {
		fmt.Println("Error while write to the file: ", errors)
	}

	fmt.Println("Successfully written to file and the byte size is: ", size)
}

func readFormFile(filename string) {
	fmt.Println()
	fmt.Println("Reading data form file: ", filename)

	file, errors := os.Open(filename)

	if errors != nil {
		fmt.Println("Error While reading file: ", errors)
	}

	fmt.Println()
	fmt.Println("Start reading data from file.")

	buffer := make([]byte, 1024)

	for {
		n, nError := file.Read(buffer)

		if nError == io.EOF {
			break
		} else if nError != nil {
			fmt.Println("Error While reading file: ", nError)
			return
		}

		fmt.Println(string(buffer[:n]))
	}

	defer file.Close()
}

func main() {
	start := time.Now()

	filename := "example.txt"

	file := createFile(filename)
	if file != nil {
		writeIntoFile(file, "Bangladesh is our homeland. ")
		writeIntoFile(file, "It is a beautiful country. ")
	}

	readFormFile(filename)
	fmt.Printf("Execution time: %s\n", time.Since(start))

	file.Close()

	content, cError := ioutil.ReadFile(filename)

	if cError != nil {
		fmt.Println("Error While reading file via ioutil: ", cError)
	}

	fmt.Println()
	fmt.Println("Reading file via ioutil")
	fmt.Println(string(content))

	content2, cError2 := os.ReadFile(filename)

	if cError2 != nil {
		fmt.Println("Error While reading file via OS: ", cError2)
	}

	fmt.Println()
	fmt.Println("Reading file via OS")
	fmt.Println(string(content2))
}
