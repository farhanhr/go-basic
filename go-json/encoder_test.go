package gojson

import (
	"encoding/json"
	"os"
	"testing"
)

func TestEncoder(t *testing.T) {
	writer, _ := os.Create("Person.json")
	encoder := json.NewEncoder(writer)

	customer := Customer{
		FirstName: "Farhan",
		MiddleName: "Husyen",
		LastName: "Ramadhan",
	}

	encoder.Encode(customer)
	
}