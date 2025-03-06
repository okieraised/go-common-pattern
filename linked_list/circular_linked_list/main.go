package main

import (
	"errors"
	"fmt"
)

// Node represents a single node in the circular linked list
type Node struct {
	data int
	next *Node
}

// CircularLinkedList represents a circular linked list
type CircularLinkedList struct {
	tail *Node
	size int
}

// Add adds a new node with the given value to the circular linked list
func (cll *CircularLinkedList) Add(value int) {
	newNode := &Node{data: value}

	if cll.tail == nil {
		// First node in the list
		newNode.next = newNode
		cll.tail = newNode
	} else {
		// Add the new node to the end of the list
		newNode.next = cll.tail.next
		cll.tail.next = newNode
		cll.tail = newNode
	}
	cll.size++
}

// Remove removes the node with the given value from the circular linked list
func (cll *CircularLinkedList) Remove(value int) error {
	if cll.tail == nil {
		return errors.New("list is empty")
	}

	current := cll.tail.next
	prev := cll.tail

	for i := 0; i < cll.size; i++ {
		if current.data == value {
			if cll.size == 1 {
				// Removing the only node in the list
				cll.tail = nil
			} else {
				// Update links to remove the current node
				prev.next = current.next
				if current == cll.tail {
					// If removing the tail, update the tail pointer
					cll.tail = prev
				}
			}
			cll.size--
			return nil
		}
		prev = current
		current = current.next
	}
	return errors.New("value not found in the list")
}

// Display prints all the elements in the circular linked list
func (cll *CircularLinkedList) Display() {
	if cll.tail == nil {
		fmt.Println("List is empty")
		return
	}

	current := cll.tail.next
	for i := 0; i < cll.size; i++ {
		fmt.Printf("%d ", current.data)
		current = current.next
	}
	fmt.Println()
}

// Search finds a node with the given value in the circular linked list
func (cll *CircularLinkedList) Search(value int) (*Node, error) {
	if cll.tail == nil {
		return nil, errors.New("list is empty")
	}

	current := cll.tail.next
	for i := 0; i < cll.size; i++ {
		if current.data == value {
			return current, nil
		}
		current = current.next
	}
	return nil, errors.New("value not found in the list")
}

// Size returns the number of nodes in the circular linked list
func (cll *CircularLinkedList) Size() int {
	return cll.size
}

func main() {
	cll := &CircularLinkedList{}

	// Add nodes to the list
	cll.Add(10)
	cll.Add(20)
	cll.Add(30)
	fmt.Println("Circular Linked List after adding 10, 20, 30:")
	cll.Display()

	// Search for a value
	node, err := cll.Search(20)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Found node with value: %d\n", node.data)
	}

	// Remove a value
	err = cll.Remove(20)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Circular Linked List after removing 20:")
		cll.Display()
	}

	// Display size of the list
	fmt.Printf("Size of the list: %d\n", cll.Size())

	// Remove all nodes
	_ = cll.Remove(10)
	_ = cll.Remove(30)
	fmt.Println("Circular Linked List after removing all nodes:")
	cll.Display()
}
