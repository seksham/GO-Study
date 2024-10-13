package main

import "fmt"

type triangle struct {
	base   float64
	height float64
}

type square struct {
	side float64
}

type shape interface {
	getArea() float64
}

func (t triangle) getArea() float64 {
	return t.base * t.height / 2
}

func (s square) getArea() float64 {
	return s.side * s.side
}

func printArea(s shape) {
	fmt.Println("area is", s.getArea())
}

func main() {
	sq := square{2.5}
	printArea(sq)
	t := triangle{
		base:   10,
		height: 10,
	}
	printArea(t)
}
