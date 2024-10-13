package main

import (
	"fmt"
	"io"
	"os"
)


// Checking if a file exists
func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}

// Creating a file
func createFile(filename string, content string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(content)
	return err
}

// Reading a file
func readFile(filename string) (string, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

// Updating a file (appending)
func appendToFile(filename string, content string) error {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(content)
	return err
}

// Renaming a file
func renameFile(oldName, newName string) error {
	return os.Rename(oldName, newName)
}

// Deleting a file
func deleteFile(filename string) error {
	return os.Remove(filename)
}

func main() {
	fileName := os.Args[1]
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	io.Copy(os.Stdout, file)

	// Example usage of new functions
	newFileName := "example.txt"
	
	if !fileExists(newFileName) {
		err = createFile(newFileName, "Hello, World!")
		if err != nil {
			fmt.Println("Error creating file:", err)
		}
	}

	content, err := readFile(newFileName)
	if err != nil {
		fmt.Println("Error reading file:", err)
	} else {
		fmt.Println("File content:", content)
	}

	err = appendToFile(newFileName, "\nAppended text")
	if err != nil {
		fmt.Println("Error appending to file:", err)
	}

	err = renameFile(newFileName, "renamed_example.txt")
	if err != nil {
		fmt.Println("Error renaming file:", err)
	}

	err = deleteFile("renamed_example.txt")
	if err != nil {
		fmt.Println("Error deleting file:", err)
	}
}

