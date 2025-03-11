package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func changeCountryToJapan(address *Address) {
	address.Country = "Japan"
}

func main()  {
	address := Address{}

	changeCountryToJapan(&address)
	fmt.Println(address)
}