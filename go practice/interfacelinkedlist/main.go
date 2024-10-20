package main

import (
	"errors"
	"fmt"
)

type Node struct {
	data any
	next *Node
}

type LinkedList struct {
	head *Node
}

// Insert adds a new node with the given data to the end of the list
func (l *LinkedList) Insert(item any) {
	newNode := new(Node)
	newNode.data = item
	if l.head == nil {
		l.head = newNode
	} else {
		current := l.head
		for current.next != nil {
			current = current.next
		}
		current.next = newNode
	}
}

// Delete removes the first occurrence of a node with the specified data
func (l *LinkedList) Delete(data any) error {
	if l.head == nil {
		return errors.New("Data can not be deleted from empty list")
	}
	if l.head.data == data {
		l.head = l.head.next
		fmt.Println(data, "deleted")
		return nil
	}
	current := l.head
	for current.next != nil {
		if current.next.data == data {
			current.next = current.next.next
			fmt.Println(data, "deleted")
			return nil
		}
		current = current.next
	}
	return errors.New("Data not found to be deleted")

}

// Search looks for a node with the given data and returns true if found, false otherwise
func (l *LinkedList) Search(data any) bool {
	if l.head == nil {
		return false
	}
	current := l.head
	for current != nil {
		if current.data == data {
			return true
		}
		current = current.next
	}
	return false
}

// Display prints all elements in the list
func (l LinkedList) Display() {
	if l.head == nil {
		fmt.Println(l.head)
		return
	}
	current := l.head
	for current.next != nil {
		fmt.Printf("%v ->", current.data)
		current = current.next
	}
	fmt.Println(current.data)
}

// IsEmpty returns true if the list has no elements, false otherwise
func (l *LinkedList) IsEmpty() bool {
	return l.head == nil
}

// Length returns the number of nodes in the list
func (l *LinkedList) Length() int {
	if l.head == nil {
		return 0
	}
	count := 0
	current := l.head
	for current.next != nil {
		count++
		current = current.next
	}
	return count+1
}

// Clear removes all nodes from the list
func (l *LinkedList) Clear() {
	l.head = nil
}

// GetFirst returns the data of the first node in the list
func (l *LinkedList) GetFirst() any {
	if l.head == nil {
		return nil
	}
	return l.head.data
}

// GetLast returns the data of the last node in the list
func (l *LinkedList) GetLast() any {
	if l.head == nil {
		return nil
	}
	current := l.head
	for current.next != nil {
		current = current.next
	}
	return current.data
}

//nil
//1
// 1 2
// 1 2 3

// InsertAt adds a new node with the given data at the specified position
func (l *LinkedList) InsertAt(data any, position int) error {
	if position < 0 || position > l.Length() {
		return errors.New("Invalid position")
	}
	newNode := new(Node)
	newNode.data = data
	//Only head present
	if position == 0 {
		newNode.next = l.head
		l.head = newNode
		fmt.Printf("%v Inserted at position: %v\n", data, position)
		return nil
	}
	current := l.head
	for i := 0; i < position-1; i++ {
		current = current.next
	}
	newNode.next = current.next
	current.next = newNode
	fmt.Printf("%v Inserted at position: %v\n", data, position)
	return nil
}

// DeleteAt removes the node at the specified position
func (l *LinkedList) DeleteAt() {

}

// Reverse changes the order of the list so that the last element becomes the first, and so on
func (l *LinkedList) Reverse() {

}

func main() {
	list := LinkedList{}
	fmt.Printf("list: %+v\n", list)
	list.Display()
	err := list.Delete(2)
	if err != nil {
		fmt.Println(err)
	}
	list.Insert(1)
	list.Display()
	err = list.Delete(1)
	if err != nil {
		fmt.Println(err)
	}
	list.Insert(2)
	list.Display()
	list.Insert(3)
	list.Display()
	list.Insert(4)
	list.Display()
	list.Insert(5)
	err = list.Delete(3)
	if err != nil {
		fmt.Println(err)
	}
	list.Display()
	err = list.Delete(5)
	if err != nil {
		fmt.Println(err)
	}
	list.Display()
	err = list.Delete(1)
	if err != nil {
		fmt.Println(err)
	}
	err = list.InsertAt(9,0)
	if err!= nil{
		fmt.Println(err)
	}
	list.Display()
	fmt.Println(list.Length())
	err = list.InsertAt(99,3)
	if err!= nil{
		fmt.Println(err)
	}
	list.Display()
	err = list.InsertAt(99,5)
	if err!= nil{
		fmt.Println(err)
	}
	list.Display()
	err = list.InsertAt(999,1)
	if err!= nil{
		fmt.Println(err)
	}
	list.Display()



}
