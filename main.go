package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	// Specify the file path
	filePath := "example.txt"

	// Read the file contents
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error reading file: %s\n", err)
		return
	}

	// Print the file contents
	fmt.Println("File contents:")
	fmt.Println(string(content))
}
