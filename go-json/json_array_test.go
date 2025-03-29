package gojson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJsonArray(t *testing.T) {

	customer := Customer{
		FirstName: "Farhan",
		MiddleName: "Husyen",
		LastName: "Ramadhan",
		Hobbies: []string{"Read Novel", "Play game", "Sleep"},
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
	
}

func TestJsonArrayDecode(t *testing.T) {
	jsonString := `{"FirstName":"Farhan","MiddleName":"Husyen","LastName":"Ramadhan","Hobbies":["Read Novel","Play game","Sleep"]}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}
	
	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.Hobbies)
}

func TestJsonArrayComplex(t *testing.T) {
	customer := Customer{
		FirstName: "Farhan",
		Adresses: []Addresses{
			{
			Street: "Jalan Merdeka No. 45",
			Country: "Indonesia",
			PostalCode: "12345",
			},
			{
			Street: "Main Street 101",
			Country: "USA",
			PostalCode: "67890",
			},
		},
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestJsonArrayComplexDecode(t *testing.T) {
	jsonString := `{"FirstName":"Farhan","MiddleName":"","LastName":"","Hobbies":null,"Adresses":[{"Street":"Jalan Merdeka No. 45","Country":"Indonesia","PostalCode":"12345"},{"Street":"Main Street 101","Country":"USA","PostalCode":"67890"}]}`
	
	jsonBytes := []byte(jsonString)

	customer := &Customer{}
	
	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.Adresses)
}