package main

import "fmt"

func main() {
	// a. Declaration
	// i. Using var keyword
	var colors1 = map[string]string{}

	// ii. Using make built-in function
	colors2 := make(map[string]string)

	// b. Initialization
	colors3 := map[string]string{
		"red":   "#ff0000",
		"blue":  "#0000ff",
		"green": "#00ff00",
	}

	// c. Put, get and delete elements
	colors2["red"] = "#ff0000"
	colors2["blue"] = "#0000ff"
	fmt.Println("After adding elements:", colors2)

	redValue := colors2["red"]
	fmt.Println("Red value:", redValue)

	delete(colors2, "blue")
	fmt.Println("After deleting blue:", colors2)

	// d. Two returns of get element
	yellowValue, exists := colors3["yellow"]
	if exists {
		fmt.Println("Yellow value:", yellowValue)
	} else {
		fmt.Println("Yellow doesn't exist in the map")
	}

	// e. Iterations
	// i. Using range to iterate over keys
	fmt.Println("Iterating over keys:")
	for color := range colors3 {
		fmt.Println(color, ":", colors3[color])
	}

	// ii. By range clause (unchanged)
	fmt.Println("Iterating over key-value pairs:")
	for color, hex := range colors3 {
		fmt.Println("Color:", color, "Hex:", hex)
	}

	fmt.Println("colors1 (declared with var):", colors1)
	fmt.Println("colors2 (created with make):", colors2)
	fmt.Println("colors3 (initialized):", colors3)
}
