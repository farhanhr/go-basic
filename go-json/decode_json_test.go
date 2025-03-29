package gojson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJsonDecode(t *testing.T) {
	jsonString := `{"FirstName":"Farhan","LastName":"Ramadhan","MiddleName":"Husyen"}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}

	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.MiddleName)
	fmt.Println(customer.LastName)
}