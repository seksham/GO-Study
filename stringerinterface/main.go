package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%s is %d years old", p.Name, p.Age)
}

func main() {
	alice := Person{Name: "Alice", Age: 30}
	bob := Person{Name: "Bob", Age: 25}

	fmt.Println(alice) // Output: Alice is 30 years old
	fmt.Println(bob)   // Output: Bob is 25 years old
	//Initialize a slice of Stringer interface with concrete types. The concrete types implement the String() method.
	people := []fmt.Stringer{alice, bob}
	for _, person := range people {
		fmt.Println(person)
	}
}
