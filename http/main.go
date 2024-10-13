package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:", len(bs))
	return len(bs), nil
}

func main() {
	resp, error := http.Get("https://google.com")
	if error != nil {
		fmt.Println("Error occured")
		os.Exit(1)
	}
	// fmt.Println(resp)
	// bs := make([]byte, 99999)
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))

	// defer resp.Body.Close()
	// body, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	fmt.Println("Error reading body:", err)
	// 	os.Exit(1)
	// }

	// fmt.Println(string(body))
	// io.Copy(os.Stdout, resp.Body)
	io.Copy(logWriter{}, resp.Body)

}
