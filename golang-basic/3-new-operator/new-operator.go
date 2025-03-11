// Sebelumnya untuk membuat pointer dengan menggunakan operator &
// Go-Lang juga memiliki function new yang bisa digunakan untuk membuat pointer
// Namun function new hanya mengembalikan pointer ke data kosong, artinya tidak ada data awal

package main
import "fmt"

type Address struct {
	City, Province, Country string
}

func main()  {
	
	var alamat1 *Address = new(Address)
	var alamat2 *Address = alamat1

	alamat2.Country = "Jepang"

	fmt.Println(alamat1)
	fmt.Println(alamat2)
	
	// Hasilnya sama saja dan cukup banyak yang menggunakan kata kunci new dibanding &
}