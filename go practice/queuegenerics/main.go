package main

import (
	"errors"
	"fmt"
)

type Queue[T any] struct {
	items []T
}

func (q *Queue[T]) Dequeue() (T, error) {
	if len(q.items) == 0 {
		var zero T
		return zero, errors.New("Queue is empty, nothing to dequeue")
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item, nil
}

func (q *Queue[T]) Enqueue(item T) {
	q.items = append(q.items, item)
}

func (q *Queue[T]) Front() (T, error) {
	if len(q.items) == 0 {
		var zero T
		return zero, errors.New("Queue is empty, nothing to return")
	}
	return q.items[0], nil
}

func (q *Queue[T]) IsEmpty() bool {
	return len(q.items) == 0
}

func (q *Queue[T]) Size() int {
	return len(q.items)
}

func (q *Queue[T]) String() string {
	return fmt.Sprintf("Queue: %v", q.items)
}

func main() {
	// Create a new queue of integers
	q := new(Queue[int])
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	fmt.Println(q)
	item, err := q.Dequeue()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Dequeued: ", item)
	fmt.Println(q)
	front, err := q.Front()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Front of queue: ", front)
	fmt.Println("Is queue empty? ", q.IsEmpty())
	item, err = q.Dequeue()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Dequeued: ", item)
	front, err = q.Front()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Front of queue: ", front)
	fmt.Println(q)
	fmt.Println("Is queue empty? ", q.IsEmpty())
	fmt.Println("Size of queue: ", q.Size())

	// Create a new queue of strings
	q2 := new(Queue[string])
	q2.Enqueue("a")
	q2.Enqueue("b")
	q2.Enqueue("c")
	fmt.Println(q2)
	item2, err := q2.Dequeue()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Dequeued: ", item2)
	front2, err := q2.Front()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Front of queue: ", front2)
	fmt.Println(q2)
	fmt.Println("Is queue empty? ", q2.IsEmpty())
	fmt.Println("Size of queue: ", q2.Size())
	item2, err = q2.Dequeue()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Dequeued: ", item2)
	fmt.Println(q2)
	front2, err = q2.Front()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Front of queue: ", front2)
	fmt.Println(q2)
	fmt.Println("Is queue empty? ", q2.IsEmpty())
	fmt.Println("Size of queue: ", q2.Size())
	item2, err = q2.Dequeue()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Dequeued: ", item2)
	fmt.Println(q2)
	front2, err = q2.Front()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Front of queue: ", front2)
	fmt.Println(q2)
	fmt.Println("Is queue empty? ", q2.IsEmpty())
	
}
