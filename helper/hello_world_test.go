package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

//Dengan menggunakan require, jika pengecekan gagal maka assert akan memanggil FailNow()
// berarti eksekusi unit test tidak akan dilanjutkan
func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Farhan")
	require.Equal(t, "Hello Farhan", result, "Result Must be 'Hello Farhan'")
	fmt.Println("This line will not execute if test is failed")
}

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