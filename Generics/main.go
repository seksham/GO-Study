package main

import "fmt"

func testSwitchCases(a int){
	switch a{
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	default:
		fmt.Println("Other")
	}
}

func testSwitchCasesWithInterface(a interface{}){
	switch types := a.(type){
	case string:
		fmt.Println(types, "is string")
	case int:
		fmt.Println(types, "is int")
	default:
		fmt.Println(types, "is other")
	}
}
//You could use more specific constraints, e.g., [T ~int | ~string] to allow only int-like or string-like types.
// any is a type alias for interface{}.

func testSwitchWithGenerics[T any](a T) {
    if str, ok := any(a).(string); ok {
        fmt.Println(str, "is string")
    } else if num, ok := any(a).(int); ok {
        fmt.Println(num, "is int")
    } else {
        fmt.Println(a, "is other")
    }
}

func testSwitchWithGenerics2[T any](a T) {
    switch v := any(a).(type) {
    case string:
        fmt.Println(v, "is string")
    case int:
        fmt.Println(v, "is int")
	default:
		fmt.Println(v, "is other")
	}
}

// New generic functions to explain generics further

// Generic function to print any slice
func PrintSlice[T any](s []T) {
    fmt.Print("Slice contents: ")
    for _, v := range s {
        fmt.Printf("%v ", v)
    }
    fmt.Println()
}

// Generic function to find the maximum value in a slice
func Max[T int | float64](s []T) T {
    if len(s) == 0 {
        var zero T
        return zero
    }
    max := s[0]
    for _, v := range s[1:] {
        if v > max {
            max = v
        }
    }
    return max
}

// Generic Stack implementation
type Stack[T any] struct {
    items []T
}

func (s *Stack[T]) Push(item T) {
    s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() (T, bool) {
    if len(s.items) == 0 {
        var zero T
        return zero, false
    }
    item := s.items[len(s.items)-1]
    s.items = s.items[:len(s.items)-1]
    return item, true
}

func main() {
	//below will cause a runtime error because the map is not initialized
	// var m map[string]int
	// m["a"] = 1
	// fmt.Println(m)

	//below will not cause a runtime error because the map is initialized with a literal. Literal is a value that is known at compile time. Example: var a = 1	
	var m = map[string]int{}
	m["a"] = 1
	fmt.Println(m)

	//A sample code using switch with channel
	c := make(chan int)
	go func() {
		c <- 1
	}()
	select {
	case <-c:
		fmt.Println("Received from channel")
	}

	myMap := map[string]int{"a": 1, "b": 2, "c": 3}

	if value, exists := myMap["key"]; exists {
		fmt.Println("Value found:", value)
	} else {
		fmt.Println("Key not found in map")
	}

	testSwitchCases(3)
	testSwitchCasesWithInterface(true)
	testSwitchWithGenerics(1)
	testSwitchWithGenerics2(true)
	a := [3]int{1,2,3}
	func(a *[3]int){
		a[0]= 4
	}(&a)
	
	fmt.Println(a)

    // Demonstrating new generic functions
    intSlice := []int{1, 5, 3, 7, 2}
    floatSlice := []float64{1.1, 5.5, 3.3, 7.7, 2.2}
    stringSlice := []string{"apple", "banana", "cherry"}

    PrintSlice(intSlice)
    PrintSlice(floatSlice)
    PrintSlice(stringSlice)

    fmt.Printf("Max of intSlice: %d\n", Max(intSlice))
    fmt.Printf("Max of floatSlice: %.2f\n", Max(floatSlice))

    // Using generic Stack
    intStack := &Stack[int]{}
    intStack.Push(1)
    intStack.Push(2)
    intStack.Push(3)

    if val, ok := intStack.Pop(); ok {
        fmt.Printf("Popped from intStack: %d\n", val)
    }

    stringStack := &Stack[string]{}
    stringStack.Push("Hello")
    stringStack.Push("World")

    if val, ok := stringStack.Pop(); ok {
        fmt.Printf("Popped from stringStack: %s\n", val)
    }
}
