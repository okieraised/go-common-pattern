package main

import "fmt"

type Deque struct {
	items []int
}

func (d *Deque) PushFront(item int) {
	d.items = append([]int{item}, d.items...)
}

func (d *Deque) PushBack(item int) {
	d.items = append(d.items, item)
}

func (d *Deque) PopFront() (int, bool) {
	if len(d.items) == 0 {
		return 0, false
	}
	item := d.items[0]
	d.items = d.items[1:]
	return item, true
}

func (d *Deque) PopBack() (int, bool) {
	if len(d.items) == 0 {
		return 0, false
	}
	item := d.items[len(d.items)-1]
	d.items = d.items[:len(d.items)-1]
	return item, true
}

func (d *Deque) IsEmpty() bool {
	return len(d.items) == 0
}

func (d *Deque) Size() int {
	return len(d.items)
}

func main() {
	dq := &Deque{}
	dq.PushBack(1)
	dq.PushFront(0)
	dq.PushBack(2)

	fmt.Println(dq.items) // Output: [0 1 2]

	front, _ := dq.PopFront()
	fmt.Println(front)    // Output: 0
	fmt.Println(dq.items) // Output: [1 2]

	back, _ := dq.PopBack()
	fmt.Println(back)     // Output: 2
	fmt.Println(dq.items) // Output: [1]
}
