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

func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Sedang Belajar Goroutine"
}

func TestCreateChannelAsParameter(t *testing.T)  {
	//membuat channel
	channel := make(chan string)
	//Menutup channel
	defer close(channel)

	go GiveMeResponse(channel)

	data := <- channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)
}

func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 5) // Kode ini berarti membuat channel dengan buffer 5, Jadi setiap data yang dikirim ke channel dan belum diambil akan ditempatkan di antrian buffer, jadi selama slot buffer masih ada, jika menambahkan data saat ada data yang belum diambil, tidak akan menyebabkan deadlock hingga slot buffernya habis. Karena datanya akan ditempatkan di buffer hingga ada yang mengambil
	defer close(channel)

	go func ()  {
		channel <- "Data Farhan"
		channel <- "Data Husyen"
		channel <- "Data Ramadhan"
	}()

	go func ()  {
		fmt.Println(<-channel)
		fmt.Println(<-channel)
		fmt.Println(<-channel)
	}()

	time.Sleep(2 * time.Second)
	fmt.Println("Selesai Mengirim Data")
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