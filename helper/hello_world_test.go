package helper

import "testing"

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Farhan")
	if result != "Hello Farhan" {
		// unit test failed
		t.Error("Result must be Hello Farhan")
	}
}