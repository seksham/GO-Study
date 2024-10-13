package main

import "fmt"

type contactInfo struct {
	address string
	zip     int
}

type person struct {
	name string
	age  int
	contactInfo
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

func (p *person) updateName(n string){
	// (*p).name = n //This also works
	p.name = n 
}

func main() {
	p := person{
		"saksham",
		18,
		contactInfo{
			"pune",
			411057,
		},
	}
	p2 := person{
		age:  33,
		name: "anuj",
		contactInfo: contactInfo{
			"pune",
			411057,
		},
	}
	fmt.Printf("%+v\n", p)
	p2.print()
	p3 := &p2
	p3.updateName("alex1")
	p2.print()
	p2.updateName("alex2") //without & also works the same way
	p2.print()
	// New type of primitive types
	type Celsius float64
	type Fahrenheit float64

	// Declaration of struct
	type Rectangle struct {
		Width  float64
		Height float64
	}

	// Initialization of struct
	r1 := Rectangle{Width: 10, Height: 5}
	r2 := Rectangle{10, 5} // Order matters when not using field names

	// Setting and getting struct value
	r1.Width = 15
	fmt.Printf("Rectangle 1 width: %f\n", r1.Width)

	// Pointer to struct
	r3 := &Rectangle{Width: 20, Height: 10}
	r3.Height = 15 // Automatically dereferenced, same as (*r3).Height = 15
	fmt.Printf("Rectangle 3: %+v\n", r3)

	// Struct equality
	fmt.Printf("r1 == r2: %v\n", r1 == r2)
	fmt.Printf("r1 == *r3: %v\n", r1 == *r3)

	// Using the new type
	var temp Celsius = 100
	fmt.Printf("Temperature in Celsius: %.2f°C\n", temp)

	// Function to convert Celsius to Fahrenheit
	celsiusToFahrenheit := func(c Celsius) Fahrenheit {
		return Fahrenheit(c*9/5 + 32)
	}

	fmt.Printf("Temperature in Fahrenheit: %.2f°F\n", celsiusToFahrenheit(temp))
}
