package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// Reading from a file
	data, err := ioutil.ReadFile("file.txt")
	if err != nil {
		// handle error
	}

	// Writing to a file
	data := []byte("Hello, world!")
	err = ioutil.WriteFile("file.txt", data, 0644)
	if err != nil {
		// handle error
	}

	// Reading a directory
	files, err := ioutil.ReadDir("dir")
	if err != nil {
		// handle error
	}
	for _, file := range files {
		fmt.Println(file.Name())
	}

	// Temporarily creating a file
	tempFile, err := ioutil.TempFile("", "example")
	if err != nil {
		// handle error
	}
	defer os.Remove(tempFile.Name()) // Clean up the temporary file

	// Reading from a io.Reader into a byte slice
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		// handle error
	}

	// Writing from a byte slice to a io.Writer
	n, err := ioutil.WriteAll(writer, data)
	if err != nil {
		// handle error
	}
}
