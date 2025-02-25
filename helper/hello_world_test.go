package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)



//Dengan menggunakan assert, jika pengecekan gagal maka assert akan memanggil FailNow()
//berarti unit test akan berlanjut dieksekusi
func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Farhan")
	assert.Equal(t, "Hello Farhan", result, "Result Must be 'Hello Farhan'")
}
func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Farhan")
	if result != "Hello Farhan" {
		// unit test failed
		t.Error("Result must be Hello Farhan")
	} 	
}