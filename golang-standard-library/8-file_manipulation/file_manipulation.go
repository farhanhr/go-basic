package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func createNewFile(name string, message string) error {
	file, err := os.OpenFile(name, os.O_CREATE|os.O_WRONLY, 0666)
	if err!= nil {
		return err
	}

	defer file.Close()

	file.WriteString(message)
	return nil
}

func readFile(name string) (string, error) {
	file, err := os.OpenFile(name, os.O_RDONLY, 0666)
	if err!= nil {
		return "", err
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	var message string
	for {
		line, _, err := reader.ReadLine()
		message += string(line)	+ "\n"
		if err == io.EOF {
			break
		}
	}

	return message, nil
}

func main() {
	// // write ke file
	// createNewFile("sample.log", "this is sample log that will created in this folder")

	result, _ := readFile("sample.log")

	fmt.Println(result)

}