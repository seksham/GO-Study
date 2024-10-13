package main

import "fmt"

// Basic struct definition
type contactInfo struct {
	address string
	zip     int
}

// Struct embedding for composition
type person struct {
	name string
	age  int
	contactInfo // Embedded struct
}

// Struct method
func (p person) print() {
	fmt.Printf("%+v\n", p)
}

// Pointer receiver method
func (p *person) updateName(n string) {
	// (*p).name = n //This also works
	p.name = n
}

func main() {
	// Struct initialization
	p := person{
		"saksham",
		18,
		contactInfo{
			"pune",
			411057,
		},
	}
	// Struct initialization with field names
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
	// Pointer to struct
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

	// Anonymous struct
	point := struct {
		X, Y int
	}{10, 20}
	fmt.Printf("Anonymous struct: %+v\n", point)

	// Struct with tags
	type User struct {
		Name  string `json:"name" validate:"required"`
		Email string `json:"email" validate:"required,email"`
	}
	user := User{Name: "John Doe", Email: "john@example.com"}
	fmt.Printf("User: %+v\n", user)

	// Struct initialization with new
	newPerson := new(person)
	newPerson.name = "Alice"
	newPerson.age = 25
	fmt.Printf("New person: %+v\n", newPerson)
}
