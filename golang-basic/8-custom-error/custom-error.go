package main

import "fmt"

type validationError struct {
	Message string
}

func (v *validationError) Error() string {
	return v.Message
}

type notFoundError struct {
	Message string
}


func (n *notFoundError) Error() string {
	return n.Message
}

func SaveData(id string, data any) error {
	if id == "" {
		return &validationError{"ID tidak boleh kosong"}
	}
	if id != "farhan" {
		return &notFoundError{"Id tidak ditemukan"}
	}

	return nil
} 


func main()  {
		err := SaveData("", nil)
		if err != nil {
			if validationErr, ok := err.(*validationError); ok {
				fmt.Println("Validation error:", validationErr.Error())
			} else if notFoundErr, ok := err.(*notFoundError); ok {
				fmt.Println("404 not found", notFoundErr.Error())
			} else {
				fmt.Println("Unknown Error")
			}
		} else {
			fmt.Println("200: Success, Hallo Farhan")
		}

}