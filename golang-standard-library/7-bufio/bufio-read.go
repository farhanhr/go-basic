package main

import (
	"bufio"
	"io"
	"strings"
)

func main()  {
	
	input := strings.NewReader("The line will be read\n by the bufio reader")

	reader := bufio.NewReader(input)

	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}

		print(string(line))
	}
}