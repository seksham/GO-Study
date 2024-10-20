package main

import (
	"errors"
	"fmt"
)

type Stack struct {
	items []any
}

func (s Stack) Size() int {
	return len(s.items)
}

func (s *Stack) Pop()(any, error){
	if len(s.items) == 0{
		return nil, errors.New("Stack is empty")
	}
	elem := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return elem, nil
}

func (s Stack) Peek()(any, error){
	if len(s.items) == 0{
		return nil, errors.New("Stack is empty")
	}
	return s.items[len(s.items)-1], nil
}

func (s *Stack) Push(items ...any)(int){
	s.items = append(s.items, items...)
	return len(items)
}

func (s Stack) String() string{
	return fmt.Sprintf("Stack: %v", s.items)
}

func (s Stack) IsEmpty() bool{
	return len(s.items) == 0
}

func main() {
	s := Stack{}
	fmt.Println(s)
	fmt.Println("Size: ", s.Size())
	s.Push(1,2,3)
	fmt.Println(s)
	fmt.Println("Size: ", s.Size())
	n, err := s.Pop()
	if err != nil{
		fmt.Println(err)
	} else {
		fmt.Println("Popped: ", n)
	}
	fmt.Println("Size: ", s.Size())
	n, err = s.Pop()
	if err != nil{
		fmt.Println(err)
	} else {
		fmt.Println("Popped: ", n)
	}
	fmt.Println("Size: ", s.Size())
	n, err = s.Pop()
	if err != nil{
		fmt.Println(err)
	} else {
		fmt.Println("Popped: ", n)
	}
	fmt.Println("Size: ", s.Size())
	fmt.Println(s)
	n, err = s.Pop()
	if err != nil{
		fmt.Println(err)
	} else {
		fmt.Println("Popped: ", n)
	}
	fmt.Println("Size: ", s.Size())
	fmt.Println(s)
	n, err = s.Peek()
	if err != nil{
		fmt.Println(err)
	} else {
		fmt.Println("Peeked: ", n)
	}
	fmt.Println("Size: ", s.Size())
	fmt.Println(s)
	s.Push(1, "a", "b", 2)
	fmt.Println(s)
	fmt.Println("Size: ", s.Size())
	n, err = s.Peek()
	if err != nil{
		fmt.Println(err)
	} else {
		fmt.Println("Peeked: ", n)
	}
	fmt.Println("Size: ", s.Size())
	n, err = s.Pop()
	if err != nil{
		fmt.Println(err)
	} else {
		fmt.Println("Popped: ", n)
	}
	fmt.Println("Size: ", s.Size())
	fmt.Println(s)
	n, err = s.Pop()
	if err != nil{
		fmt.Println(err)
	} else {
		fmt.Println("Popped: ", n)
	}
	fmt.Println("Size: ", s.Size())
	fmt.Println(s)
}
