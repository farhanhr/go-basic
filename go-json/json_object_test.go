package gojson

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Addresses struct {
	Street string
	Country string
	PostalCode string
}

type Customer struct {
	FirstName string
	MiddleName string
	LastName string
	Hobbies []string
	Adresses []Addresses
}

func TestJsonObject(t *testing.T) {
	customer := Customer{
		FirstName: "Farhan",
		MiddleName: "Husyen",
		LastName: "Ramadhan",
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}