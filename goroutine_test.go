package golang_goroutine

import (
	"fmt"
	"testing"
	"time"
)

func RunHelloWorld() {
	fmt.Println("Hello World")
}

func DisplayNumber(number int) {
	fmt.Println("Display", number)
}

func TestManyGoroutine(t *testing.T) {
	for i := 0; i<100000; i++ {
		go DisplayNumber(i)
	}
	time.Sleep(5 * time.Second)
}

func TestCreateGoroutine(t *testing.T) {
    go RunHelloWorld()
    fmt.Println("Akan diprint duluan")

    time.Sleep(1 * time.Second)
}