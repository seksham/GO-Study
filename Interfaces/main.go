package main

import (
	"fmt"
	"sort"
)

type englishBot struct{}
type spanishBot struct{}
type bot interface {
	getGreeting() string
}

func (englishBot) getGreeting() string {
	return "Hello"
}

func (spanishBot) getGreeting() string {
	return "Hola"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

// a. Introduction and declaration
// We've already declared the 'bot' interface above

// b. Implementing an interface
// i. Implicit Interface (already demonstrated with englishBot and spanishBot)

// ii. Polymorphism: Interface as a contract
type germanBot struct{}

func (germanBot) getGreeting() string {
	return "Hallo"
}

// c. Empty interface. Same as any which is alias for interface{}
func printAnything(a interface{}) {
	fmt.Printf("Value: %v, Type: %T\n", a, a)
}

// d. Method set
type reader interface {
	read() string
}

type writer interface {
	write(string)
}

type readWriter interface {
	reader
	writer
}

// e. Frequently used built-in/stdlib interfaces
// i. Stringer interface
type person struct {
	name string
	age  int
}

func (p person) String() string {
	return fmt.Sprintf("%s (%d years)", p.name, p.age)
}

// ii. Interface interface of sort package
type intSlice []int

func (s intSlice) Len() int           { return len(s) }
func (s intSlice) Less(i, j int) bool { return s[i] < s[j] }
func (s intSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

// f. Interface with struct
// i. Interface variable as struct field
type logger interface {
	log(string)
}

type consoleLogger struct{}

func (consoleLogger) log(message string) {
	fmt.Println("Log:", message)
}

// This is composition, not embedding. The service struct has a logger field,
// which means it "has-a" logger, but the logger is not part of the service's type.
// Composition allows the service to use the logger's functionality without inheriting its type.
type service struct {
	logger logger // logger is a field of type logger
}

// ii. Embedding interface into the struct.
// This is embedding, not composition. The embeddedService struct "is-a" logger,
// because it includes the logger interface as part of its type.
type embeddedService struct {
	logger
}

func main() {
	// Using the existing bots

	var e englishBot
	s := spanishBot{}
	g := germanBot{}
	printGreeting(s)
	printGreeting(e)
	printGreeting(g)

	// Using empty interface
	printAnything(42)
	printAnything("Hello")
	printAnything(true)

	// Using Stringer interface
	p := person{name: "Alice", age: 30}
	fmt.Println(p)

	// Using sort.Interface
	numbers := intSlice{3, 1, 4, 1, 5, 9, 2, 6, 5, 3}
	fmt.Println("Before sorting:", numbers)
	sort.Sort(numbers)
	fmt.Println("After sorting:", numbers)

	// Using interface as struct field
	logger := consoleLogger{}
	svc := service{logger: logger}
	svc.logger.log("This is a log message")

	// Using embedded interface
	embeddedSvc := embeddedService{logger: logger}
	embeddedSvc.log("This is an embedded log message")
}
