package main

import (
	"errors"
	"fmt"
	"io"
	"os"
)

// Custom error type
type MyError struct {
	message string
}

func (e *MyError) Error() string {
	return e.message
}

func main() {
	// Error handling
	result, err := divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}

	// Custom error
	err = customErrorFunction()
	if err != nil {
		fmt.Println("Custom error:", err)
	}

	originalErr := errors.New("database connection failed")
	wrappedErr := fmt.Errorf("failed to fetch user data: %w", originalErr)

	// Later in the code
	fmt.Println(wrappedErr)
	// Output: failed to fetch user data: database connection failed

	// You can unwrap the error
	unwrapped := errors.Unwrap(wrappedErr)
	fmt.Println(unwrapped == originalErr) // true

	// errors.Is example
	if errors.Is(wrappedErr, originalErr) {
		fmt.Println("wrappedErr contains originalErr")
	}

	// errors.As example
	var myErr *MyError
	err = customErrorFunction()
	if errors.As(err, &myErr) {
		fmt.Printf("Error is of type MyError: %v\n", myErr)
	}

	// Handling multiple error types
	err = multipleErrorTypes()
	switch {
	case errors.Is(err, os.ErrNotExist):
		fmt.Println("File does not exist:", err)
	case errors.Is(err, os.ErrPermission):
		fmt.Println("Permission denied")
	case err != nil:
		fmt.Println("Unknown error:", err)
	}

	// Panic and defer
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	panicFunction()

	// This line won't be executed due to the panic
	fmt.Println("This line won't be printed")
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

func customErrorFunction() error {
	return &MyError{message: "This is a custom error"}
}

func panicFunction() {
	panic("This is a panic situation")
}

func multipleErrorTypes() error {
	_, err := os.Open("non-existent-file.txt")
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	return nil
}

// Example of error wrapping with custom error types
type DatabaseError struct {
	Err error
}

func (e *DatabaseError) Error() string {
	return fmt.Sprintf("database error: %v", e.Err)
}

func (e *DatabaseError) Unwrap() error {
	return e.Err
}

func fetchUserData() error {
	err := errors.New("connection timeout")
	return &DatabaseError{Err: err}
}

// Example of using io.EOF
func readUntilEOF(r io.Reader) error {
	buf := make([]byte, 1024)
	for {
		_, err := r.Read(buf)
		if err == io.EOF {
			return nil // End of file, not an error
		}
		if err != nil {
			return fmt.Errorf("read error: %w", err)
		}
		// Process buf...
	}
}
