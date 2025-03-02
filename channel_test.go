package golang_goroutine

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T)  {
	//membuat channel
	channel := make(chan string)
	//Menutup channel
	defer close(channel)

	go func() {
		time.Sleep(2 * time.Second)
		// Mengirim data ke channel
		channel <- "Data Farhan"  
		fmt.Println("Selesai Mengiri Data ke channel") // Jika data yang dikirim ke channel belum ada yang mengambil, maka data akan tetap berada di channel dan memblok channel hingga ada yang mengambil datanya. Sehingga baris kode ini tidak akan dieksekusi jika tidak ada yang mengambil data dari channel
	}()

	data := <- channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)
}

// func TestCreateChannel(t *testing.T)  {
// 	//membuat channel
// 	channel := make(chan string)
// 	// Mengirim data ke channel
// 	channel <- "Farhan"
// 	// Menerima data dari channel
// 	data:= <-channel

// 	fmt.Println(data)
// 	// Menutup channel
// 	close(channel)
// }