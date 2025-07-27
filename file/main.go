package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	/*
	
	file, err := os.Create("example.txt")
	if err != nil {
		fmt.Println("Error while creating file: ", err)
		return
	}
	defer file.Close()

	content := "Hi laado"
	byte, errors := io.WriteString(file, content+"\n")
	fmt.Println("byte written: ", byte)
	if errors != nil {
		fmt.Println("Error while writing file: ", errors)
		return
	}

	fmt.Println("sucessfully created file")

	*/


	

	// reading a file using buffer
	
	file, err := os.Open("example.txt")
	if err != nil {
		fmt.Println("Error while opening file: ", err)
		return
	}
	defer file.Close()

	// Create a buffer to read the file content
	buffer := make([]byte, 1024)    // create a buffer of 1KB , buffer is a slice of bytes 

	// Read the file content into the buffer
	for {
		n, err := file.Read(buffer)
		if err == io.EOF {    // End of file reached
			break
		}
		if err != nil {
			fmt.Println("Error while reading file", err)
			return
		}

		// Process the read content
		fmt.Println(string(buffer[:n]))      
	}
	

	


	/*

	//Read the entire file into a byte slice

	content, err := os.ReadFile("example.txt")
	if err != nil {
		fmt.Println("Error while reading file ", err)
		return
	}
	fmt.Println(string(content))

	*/


}
