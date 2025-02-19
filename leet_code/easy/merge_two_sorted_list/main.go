package main

import "fmt"

// You are given the heads of two sorted linked lists list1 and list2.
//
// Merge the two lists into one sorted list. The list should be made by splicing together the nodes of the first two lists.
//
// Return the head of the merged linked list.
//
// Example 1:
//
// Input: list1 = [1,2,4], list2 = [1,3,4]
// Output: [1,1,2,3,4,4]
//
// Example 2:
//
// Input: list1 = [], list2 = []
// Output: []
//
// Example 3:
//
// Input: list1 = [], list2 = [0]
// Output: [0]
//
// Constraints:
//
//	The number of nodes in both lists is in the range [0, 50].
//	-100 <= Node.val <= 100
//	Both list1 and list2 are sorted in non-decreasing order.
func main() {
	l1 := &ListNode{}
	l1.Val = 1
	l1.Next = &ListNode{
		Val: 2,
	}
	l1.Next.Next = &ListNode{
		Val: 4,
	}

	l2 := &ListNode{}
	l2.Val = 1
	l2.Next = &ListNode{
		Val: 3,
	}
	l2.Next.Next = &ListNode{
		Val: 4,
	}

	res := mergeTwoLists(l1, l2)
	current := res.Next
	for current != nil {
		fmt.Printf("%d -> ", current.Val)
		current = current.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

	if list1 == nil && list2 == nil {
		return nil
	}

	if list1 == nil {
		return list2
	}

	if list2 == nil {
		return list1
	}

	current1 := list1
	current2 := list2

	if current1.Val <= current2.Val {
		current1.Next = mergeTwoLists(current1.Next, current2)
		return current1

	} else {
		current2.Next = mergeTwoLists(current1, current2.Next)
		return current2
	}

}
