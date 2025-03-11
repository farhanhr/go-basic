package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Maried() {
	man.Name = "Mr. " + man.Name
}

func main()  {

	// Walaupun method akan menempel di struct, tapi sebenarnya data struct yang diakses di dalam method adalah pass by value
	// Sangat direkomendasikan menggunakan pointer di method, sehingga tidak boros memory karena harus selalu diduplikasi ketika memanggil method

	farhan := Man{"Farhan"}
	farhan.Maried()
	fmt.Println(farhan)
}