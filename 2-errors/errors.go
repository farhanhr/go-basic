package main

import (
	"errors"
	"fmt"
)

var (
	ErrValidation = errors.New("validation error")
	ErrNotFound = errors.New("not found error")
)

func GetById(id string) error {
	if id == "" {
		return ErrValidation
	}

	if id != "farhan" {
		return ErrNotFound
	}
	
	return nil
}

func main() {
	err := GetById("farhan")

	if err != nil {
		if errors.Is(err, ErrValidation) {
			fmt.Println("Validation error")
		} else if errors.Is(err, ErrNotFound) {
			fmt.Println("Not found error")
		} else {
			fmt.Println("Unknown error")
		}
	}

}