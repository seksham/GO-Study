package main

import (
	"errors"
	"fmt"
)

type Allowed interface {
	int | string
}

type Stack[T Allowed] struct {
	items []T
}

func (s Stack[T]) Size() int {
	return len(s.items)
}

func (s *Stack[T]) Push(elements ...T) int {
	s.items = append(s.items, elements...)
	return len(elements)
}

func (s Stack[T]) String() string {
	return fmt.Sprintf("Stack: %v", s.items)
}

func (s *Stack[T]) Pop() (T, error) {
	if len(s.items) == 0 {
		var zero T
		return zero, errors.New("stack is empty")
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item, nil
}

func (s *Stack[T]) Peek() (T, error) {
	if len(s.items) == 0 {
		var zero T
		return zero, errors.New("stack is empty")
	}
	return s.items[len(s.items)-1], nil
}
func (s *Stack[T]) Clear() {
	s.items = []T{}
}

func (s Stack[T]) IsEmpty() bool {
	return len(s.items) == 0
}

func main() {
	// Test Stack[int]
	fmt.Println("Testing Stack[int]:")
	intStack := new(Stack[int])
	fmt.Printf("New stack: %v\n", intStack)
	fmt.Printf("Is empty: %v\n", intStack.IsEmpty())
	fmt.Printf("Size: %d\n", intStack.Size())

	fmt.Println("\nPushing elements:")
	intStack.Push(1, 2, 3, 4, 5)
	fmt.Println("Is empty:", intStack.IsEmpty())
	fmt.Printf("After push: %v\n", intStack)
	fmt.Printf("Size: %d\n", intStack.Size())

	fmt.Println("\nPeeking:")
	if val, err := intStack.Peek(); err == nil {
		fmt.Printf("Peeked: %d\n", val)
	} else {
		fmt.Printf("Peek error: %v\n", err)
	}

	fmt.Println("\nPopping elements:")
	for {
		if val, err := intStack.Pop(); err == nil {
			fmt.Printf("Popped: %d\n", val)
		} else {
			fmt.Printf("Pop error: %v\n", err)
			break
		}
	}

	fmt.Printf("\nAfter popping all: %v\n", intStack)
	fmt.Printf("Is empty: %v\n", intStack.Size() == 0)

	// Test Stack[string]
	fmt.Println("\nTesting Stack[string]:")
	strStack := new(Stack[string])
	fmt.Printf("New stack: %v\n", strStack)
	fmt.Printf("Is empty: %v\n", strStack.IsEmpty())

	fmt.Println("\nPushing elements:")
	strStack.Push("Hello", "World", "Go")
	fmt.Printf("After push: %v\n", strStack)
	fmt.Printf("Size: %d\n", strStack.Size())

	fmt.Println("\nPeeking:")
	if val, err := strStack.Peek(); err == nil {
		fmt.Printf("Peeked: %s\n", val)
	} else {
		fmt.Printf("Peek error: %v\n", err)
	}

	fmt.Println("\nPopping elements:")
	for i := 0; i < 4; i++ {
		if val, err := strStack.Pop(); err == nil {
			fmt.Printf("Popped: %s\n", val)
		} else {
			fmt.Printf("Pop error: %v\n", err)
		}
	}

	fmt.Println("\nClearing the stack:")
	strStack.Push("A", "B", "C")
	fmt.Printf("Before clear: %v\n", strStack)
	strStack.Clear()
	fmt.Printf("After clear: %v\n", strStack)
	fmt.Printf("Is empty: %v\n", strStack.IsEmpty())
}
