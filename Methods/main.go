package main

import (
	"fmt"
	"math"
)

// a. Declaration
type Rectangle struct {
	Width  float64
	Height float64
}

// Method with value receiver
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// b. Pointer receiver
func (r *Rectangle) Scale(factor float64) {
	r.Width *= factor
	r.Height *= factor
}

// c. Methods on different types
// i. Method on basic data type
type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

// ii. Method on composite type (slice)
type IntSlice []int

func (is IntSlice) Sum() int {
	total := 0
	for _, v := range is {
		total += v
	}
	return total
}

// iii. Method on function type
type MathFunc func(float64) float64

func (f MathFunc) Apply(x float64) float64 {
	return f(x)
}

// d. Method of embedded types
type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Embedding Circle in a new type
type ColoredCircle struct {
	Circle
	Color string
}

// i. Method resolution process
// ii. Polymorphism
func PrintArea(s Shape) {
	fmt.Printf("Area: %f\n", s.Area())
}

func main() {
	// Using Rectangle methods
	r := Rectangle{Width: 10, Height: 5}
	fmt.Printf("Rectangle Area: %f\n", r.Area())
	r.Scale(2)
	fmt.Printf("Scaled Rectangle: %+v\n", r)

	// Using MyFloat method
	f := MyFloat(-4.5)
	fmt.Printf("Absolute value: %f\n", f.Abs())

	// Using IntSlice method
	is := IntSlice{1, 2, 3, 4, 5}
	fmt.Printf("Sum of slice: %d\n", is.Sum())

	// Using MathFunc method
	square := MathFunc(func(x float64) float64 { return x * x })
	fmt.Printf("Square of 3: %f\n", square.Apply(3))

	// Using embedded type and demonstrating polymorphism
	c := Circle{Radius: 5}
	cc := ColoredCircle{Circle: Circle{Radius: 7}, Color: "Red"}

	PrintArea(c)
	PrintArea(cc) // ColoredCircle inherits Area() method from Circle
}
