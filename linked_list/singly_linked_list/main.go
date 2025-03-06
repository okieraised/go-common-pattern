package main

import (
	"errors"
	"fmt"
)

// Node represents a single node in the linked list
type Node struct {
	data int
	next *Node
}

// LinkedList represents a singly linked list
type LinkedList struct {
	head *Node
	size int
}

// Add adds a new node with the given value at the end of the linked list
func (ll *LinkedList) Add(value int) {
	newNode := &Node{data: value}
	if ll.head == nil {
		ll.head = newNode
	} else {
		current := ll.head
		for current.next != nil {
			current = current.next
		}
		current.next = newNode
	}
	ll.size++
}

// Remove removes the first node with the given value from the linked list
func (ll *LinkedList) Remove(value int) error {
	if ll.head == nil {
		return errors.New("list is empty")
	}

	if ll.head.data == value {
		ll.head = ll.head.next
		ll.size--
		return nil
	}

	current := ll.head
	for current.next != nil && current.next.data != value {
		current = current.next
	}

	if current.next == nil {
		return errors.New("value not found in the list")
	}

	current.next = current.next.next
	ll.size--
	return nil
}

// Search finds a node with the given value in the linked list
func (ll *LinkedList) Search(value int) (*Node, error) {
	current := ll.head
	for current != nil {
		if current.data == value {
			return current, nil
		}
		current = current.next
	}
	return nil, errors.New("value not found in the list")
}

// Display prints all the elements in the linked list
func (ll *LinkedList) Display() {
	if ll.head == nil {
		fmt.Println("List is empty")
		return
	}

	current := ll.head
	for current != nil {
		fmt.Printf("%d -> ", current.data)
		current = current.next
	}
	fmt.Println("nil")
}

// Size returns the number of nodes in the linked list
func (ll *LinkedList) Size() int {
	return ll.size
}

func main() {
	ll := &LinkedList{}

	// Add nodes to the list
	ll.Add(10)
	ll.Add(20)
	ll.Add(30)
	fmt.Println("Linked List after adding 10, 20, 30:")
	ll.Display()

	// Search for a value
	node, err := ll.Search(20)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Found node with value: %d\n", node.data)
	}

	// Remove a value
	err = ll.Remove(20)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Linked List after removing 20:")
		ll.Display()
	}

	// Display size of the list
	fmt.Printf("Size of the list: %d\n", ll.Size())

	// Remove all nodes
	_ = ll.Remove(10)
	_ = ll.Remove(30)
	fmt.Println("Linked List after removing all nodes:")
	ll.Display()
}
