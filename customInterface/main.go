package main

import (
	"fmt"
	"io"
)

// Reader is our custom interface
type Reader interface {
	Read(p []byte) (n int, err error)
}

// StringReader implements the Reader interface
type StringReader struct {
	s string
	i int
}

// Read implements the Reader interface for StringReader
func (sr *StringReader) Read(p []byte) (n int, err error) {
	if sr.i >= len(sr.s) {
		return 0, io.EOF
	}
	n = copy(p, sr.s[sr.i:])
	sr.i += n
	return n, nil
}

// Response is our custom response struct
type Response struct {
	Body Reader
}

func main() {
	// Create a new Response with a StringReader as its Body
	resp := Response{
		Body: &StringReader{s: "Hello, World!"},
	}

	// Read from the Body in chunks
	buffer := make([]byte, 5)
	for {
		n, err := resp.Body.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Printf("Read %d bytes: %s\n", n, buffer[:n])
	}
}
