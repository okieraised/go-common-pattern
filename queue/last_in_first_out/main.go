package main

import (
	"errors"
	"fmt"
)

// Stack represents a LIFO data structure
type Stack struct {
	items []int
}

// Push adds an item to the top of the stack
func (s *Stack) Push(item int) {
	s.items = append(s.items, item)
}

// Pop removes and returns the item from the top of the stack
func (s *Stack) Pop() (int, error) {
	if len(s.items) == 0 {
		return 0, errors.New("stack is empty")
	}
	// Get the last item
	lastIndex := len(s.items) - 1
	item := s.items[lastIndex]
	// Remove the last item
	s.items = s.items[:lastIndex]
	return item, nil
}

// Peek returns the item at the top of the stack without removing it
func (s *Stack) Peek() (int, error) {
	if len(s.items) == 0 {
		return 0, errors.New("stack is empty")
	}
	return s.items[len(s.items)-1], nil
}

// IsEmpty checks if the stack is empty
func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

// Size returns the number of items in the stack
func (s *Stack) Size() int {
	return len(s.items)
}

func main() {
	stack := &Stack{}

	// Push items onto the stack
	stack.Push(10)
	stack.Push(20)
	stack.Push(30)

	fmt.Println("Stack after pushing 10, 20, 30:", stack.items)

	// Peek at the top item
	top, err := stack.Peek()
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Top item:", top)
	}

	// Pop items from the stack
	for !stack.IsEmpty() {
		item, err := stack.Pop()
		if err != nil {
			fmt.Println("Error:", err)
			break
		}
		fmt.Println("Popped item:", item)
	}

	// Attempt to pop from an empty stack
	_, err = stack.Pop()
	if err != nil {
		fmt.Println("Error:", err)
	}
}
