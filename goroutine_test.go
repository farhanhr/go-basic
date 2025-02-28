package golang_goroutine

import (
	"fmt"
	"testing"
	"time"
)

func RunHelloWorld() {
	fmt.Println("Hello World")
}

func TestCreateGoroutine(t *testing.T) {
    go RunHelloWorld()
    fmt.Println("Akan diprint duluan")

    time.Sleep(1 * time.Second)
}