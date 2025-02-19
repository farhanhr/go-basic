package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main()  {
	
// // Saat kita mengubah variable pointer, maka yang berubah hanya variable tersebut.
// // Semua variable yang mengacu ke data yang sama tidak akan berubah
// Jika kita ingin mengubah seluruh variable yang mengacu ke data tersebut, kita bisa menggunakan operator *

var address1 Address= Address{"Subang", "Jawa Barat", "Indonesia"}
	var address2 *Address = &address1 
	address2.City = "Bogor"
	fmt.Println(address1) 
	fmt.Println(address2)
	
	address2 = &Address{"Jakpus", "Jakarta", "Indonesia"} 
	fmt.Println(address1) // Tidak berubah 
	fmt.Println(address2) // berubah sesuai inisialisasi ulang
	
}