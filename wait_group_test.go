package golang_goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func RunAsynchronous(group *sync.WaitGroup) {
	defer group.Done()

	group.Add(1)

	fmt.Println("Hello Wait Group")
	time.Sleep(1 * time.Second)
}

func TestGroupWait(t *testing.T) {
	group := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go RunAsynchronous(group)
	}

	group.Wait()
	print("Selesai Menunggu")
}