package main

import "fmt"

type Address struct {
	City, Province, Country string
}

// Secara default yang dipassing ke dalam method/function dalam bahasa pemrogramman
// go adalah valuenya bukan reference, jadi data aslinya seperti pada contoh dibawah ini
// tidak terpengaruh
func main() {
	// address2 := address1
	// address2.City = "Jabar"
	// fmt.Println(address1)
	// fmt.Println(address2)
	
	// Pointer adalah kemampuan membuat reference ke lokasi data di memory yang sama, 
	// tanpa menduplikasi data yang sudah ada
	//  dengan kemampuan pointer, kita bisa membuat pass by reference
	
	var address1 Address= Address{"Subang", "Jawa Barat", "Indonesia"}
	var address2 *Address = &address1 //dengan menggunakan pointer address1 akan berubah karena menunjuk ke data yang sama (bisa ditulis dengan address2 := &address1 saja untuk menyingkatkan)
	address2.City = "Bogor"
	fmt.Println(address1) //ikut berubah
	fmt.Println(address2)
}