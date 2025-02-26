package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Table test
func TestTable(t *testing.T) {
	tests := []struct {
		name string
		request string
		expected string
	} {
		{
			name : "Test Hello Farhan",
			request : "Farhan",
			expected : "Hello Farhan",
		}, 
		{
			name : "Test Hello Husyen",
			request : "Husyen",
			expected : "Hello Husyen",
		},
	}

	for _, test  := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}

// function subtest memungkinkan untuk memiliki test didalam unit test
func TestSubTest(t *testing.T) {
	t.Run("Farhan", func(t *testing.T) {
		result := HelloWorld("Farhan")
		require.Equal(t, "Hello Farhan", result)
	})
	t.Run("HR", func(t *testing.T) {
		result := HelloWorld("HR")
		require.Equal(t, "Hello HR", result)
	})
}

//fungsi TestMain bisa dimanfaatkan agar bisa mirip dengan before dan after test
func TestMain(m *testing.M) {
	//before
	fmt.Println("Selamat Datang di Unit Test")
	m.Run()
	//after
	fmt.Println("Good bye Unit Test")
}

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