package main

import "fmt"

type Queue struct {
	items []any
}

func (q Queue) Front() (any, error) {
	if len(q.items) == 0 {
		var zero any
		return zero, fmt.Errorf("Queue is empty")
	}
	return q.items[0], nil
}

// [4]
func (q *Queue) Enqueue(item any) {
	q.items = append(q.items, item)
}

func (q *Queue) Dequeue() (any, error) {
    if q.IsEmpty() {
        return nil, fmt.Errorf("Queue is empty. Nothing to dequeue")
    }
    item := q.items[0]
    q.items = q.items[1:]
    return item, nil
}

func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}

func (q *Queue) Size() int {
	return len(q.items)
}
func (q Queue) String() string {
	return fmt.Sprintf("Queue: %v", q.items)
}

func main() {
	queue := Queue{}
	queue.Enqueue(1)
	queue.Enqueue("b")
	queue.Enqueue(3)
	fmt.Println("Size of queue: ", queue.Size())
	front, err := queue.Front()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Front of queue: ", front)
	item, err := queue.Dequeue()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Dequeue: ", item)
	fmt.Println(queue)
	front, err = queue.Front()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Front of queue: ", front)
	}
	item, err = queue.Dequeue()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Dequeue: ", item)
	fmt.Println(queue)
	front, err = queue.Front()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Front of queue: ", front)
	}
	fmt.Println("Is queue empty? ", queue.IsEmpty())
	item, err = queue.Dequeue()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Dequeue: ", item)
	front, err = queue.Front()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Front of queue: ", front)
	}
	fmt.Println(queue)
	fmt.Println("Is queue empty? ", queue.IsEmpty())
	fmt.Println("Size of queue: ", queue.Size())

}
