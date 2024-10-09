package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	fmt.Println("Web request")

	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("Error occure while doing web request", err)
	}

	fmt.Println("Response status: ", res.Status)
	fmt.Println("Response status code: ", res.StatusCode)

	data, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error occure while reading response body", err)
	}

	fmt.Println("Response body: ", string(data))

	defer res.Body.Close()
}
