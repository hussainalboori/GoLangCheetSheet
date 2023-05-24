package main

import (
	"fmt"
	"os"
)

func main() {
	// Retrieving the current working directory
	wd, err := os.Getwd()
	if err != nil {
		// handle error
	}
	fmt.Println("Current working directory:", wd)

	// Changing the current working directory
	err = os.Chdir("/path/to/directory")
	if err != nil {
		// handle error
	}

	// Creating a new directory
	err = os.Mkdir("/path/to/directory", os.ModePerm)
	if err != nil {
		// handle error
	}

	// Creating a new directory (including parent directories)
	err = os.MkdirAll("/path/to/directory", os.ModePerm)
	if err != nil {
		// handle error
	}

	// Removing a directory
	err = os.Remove("/path/to/directory")
	if err != nil {
		// handle error
	}

	// Removing a directory (including its contents)
	err = os.RemoveAll("/path/to/directory")
	if err != nil {
		// handle error
	}

	// Opening a file for reading
	file, err := os.Open("/path/to/file")
	if err != nil {
		// handle error
	}
	defer file.Close()

	// Opening a file for writing
	file, err = os.Create("/path/to/file")
	if err != nil {
		// handle error
	}
	defer file.Close()

	// Reading the contents of a file
	data := make([]byte, 1024)
	n, err := file.Read(data)
	if err != nil {
		// handle error
	}
	content := string(data[:n])
	fmt.Println("File content:", content)

	// Writing data to a file
	data = []byte("Hello, world!")
	_, err = file.Write(data)
	if err != nil {
		// handle error
	}

	// Checking if a file or directory exists
	if _, err := os.Stat("/path/to/file_or_directory"); os.IsNotExist(err) {
		fmt.Println("File or directory does not exist")
	} else {
		fmt.Println("File or directory exists")
	}

	// Renaming a file or directory
	err = os.Rename("/path/to/old_name", "/path/to/new_name")
	if err != nil {
		// handle error
	}

	// Getting file information
	fileInfo, err := os.Stat("/path/to/file")
	if err != nil {
		// handle error
	}
	fmt.Println("File name:", fileInfo.Name())
	fmt.Println("Size:", fileInfo.Size())
	fmt.Println("Is directory?", fileInfo.IsDir())
	fmt.Println("Permission mode:", fileInfo.Mode())

	// Listing files in a directory
	files, err := os.ReadDir("/path/to/directory")
	if err != nil {
		// handle error
	}
	for _, file := range files {
		fmt.Println(file.Name())
	}
}
