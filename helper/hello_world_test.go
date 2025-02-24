package helper

import "testing"

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Farhan")
	if result != "Hello Farhan" {
		// unit test failed
		panic("Result is not as expected")
	}
}