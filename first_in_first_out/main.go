package main

import (
	"errors"
	"fmt"
)

// Queue represents a FIFO data structure
type Queue struct {
	items []int
}

// Enqueue adds an item to the end of the queue
func (q *Queue) Enqueue(item int) {
	q.items = append(q.items, item)
}

// Dequeue removes and returns the item from the front of the queue
func (q *Queue) Dequeue() (int, error) {
	if len(q.items) == 0 {
		return 0, errors.New("queue is empty")
	}
	// Get the first item
	item := q.items[0]
	// Remove the first item
	q.items = q.items[1:]
	return item, nil
}

// Peek returns the item at the front of the queue without removing it
func (q *Queue) Peek() (int, error) {
	if len(q.items) == 0 {
		return 0, errors.New("queue is empty")
	}
	return q.items[0], nil
}

// IsEmpty checks if the queue is empty
func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}

// Size returns the number of items in the queue
func (q *Queue) Size() int {
	return len(q.items)
}

func main() {
	queue := &Queue{}

	// Enqueue items
	queue.Enqueue(10)
	queue.Enqueue(20)
	queue.Enqueue(30)
	fmt.Println("Queue after enqueuing 10, 20, 30:", queue.items)

	// Peek at the front item
	front, err := queue.Peek()
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Front item:", front)
	}

	// Dequeue items
	for !queue.IsEmpty() {
		item, err := queue.Dequeue()
		if err != nil {
			fmt.Println("Error:", err)
			break
		}
		fmt.Println("Dequeued item:", item)
	}

	// Attempt to dequeue from an empty queue
	_, err = queue.Dequeue()
	if err != nil {
		fmt.Println("Error:", err)
	}
}
