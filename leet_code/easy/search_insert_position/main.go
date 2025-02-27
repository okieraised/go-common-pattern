package main

import "fmt"

// Given a sorted array of distinct integers and a target value, return the index if the target is found. If not, return the index where it would be if it were inserted in order.
//
// You must write an algorithm with O(log n) runtime complexity.
//
// Example 1:
//
// Input: nums = [1,3,5,6], target = 5
// Output: 2
//
// Example 2:
//
// Input: nums = [1,3,5,6], target = 2
// Output: 1
//
// Example 3:
//
// Input: nums = [1,3,5,6], target = 7
// Output: 4
//
// Constraints:
//
//	1 <= nums.length <= 104
//	-104 <= nums[i] <= 104
//	nums contains distinct values sorted in ascending order.
//	-104 <= target <= 104
func main() {
	x := []int{1, 3, 5, 6}
	fmt.Println(searchInsert(x, 5)) // -> 2

	fmt.Println(searchInsert(x, 2)) // -> 1

	fmt.Println(searchInsert(x, 7)) // -> 4

	x = []int{1}
	fmt.Println(searchInsert(x, 0)) // -> 0

	x = []int{1, 3, 5}
	fmt.Println(searchInsert(x, 3)) // -> 1
	fmt.Println(searchInsert(x, 4)) // -> 2

	x = []int{3, 5, 7, 9, 10}
	fmt.Println(searchInsert(x, 8)) // -> 3
}

func searchInsert(nums []int, target int) int {
	l, r := 0, len(nums)
	for l < r {
		mid := (l + r) >> 1
		if nums[mid] >= target {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}
