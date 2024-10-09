package main

import (
	"encoding/json"
	"fmt"
)

type Profile struct {
	Address string `json:"address"`
	Phone   string `json:"phone"`
}

type Person struct {
	Name    string    `json:"name"`
	Email   string    `json:"email"`
	Age     int       `json:"age"`
	Profile []Profile `json:"profiles"`
}

func main() {
	fmt.Println("Learning Json")

	profile1 := Profile{Address: "Dhaka", Phone: "017545544"}
	profile2 := Profile{Address: "Barishal", Phone: "888888888"}

	person := Person{Name: "Ruhul", Email: "rihul.ruhul@gmail.com", Age: 25}
	person.Profile = []Profile{profile1, profile2}

	// Encoding to json
	jsonString, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error Occur while encoding to json", err)

		return
	}
	fmt.Println("Encoding JSON string", string(jsonString))

	// Decoding Json DATA
	var decodedData Person
	err = json.Unmarshal(jsonString, &decodedData)

	if err != nil {
		fmt.Println("Error Occur while decoding to json", err)
	}

	fmt.Println("Decoded JSON string", decodedData)

	// Decode static json string
	var decodedData2 Person
	staticString := `{"name":"Ruhul","email":"rihul.ruhul@gmail.com","age":25,"profiles":[{"address":"Dhaka","phone":"017545544"},{"address":"Barishal","phone":"888888888"}]}`

	err = json.Unmarshal([]byte(staticString), &decodedData2)
	if err != nil {
		return
	}

	fmt.Println("Decoded Static JSON string", decodedData2)

}
