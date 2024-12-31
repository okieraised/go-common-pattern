package main

import "sort"

// SortAndFilter sorts the array and filter out elements not equal to the previous element + 1
func SortAndFilter(arr []int) []int {
	sort.Ints(arr)

	var longest []int
	var current []int

	// Iterate through the sorted array
	for i, val := range arr {
		if i == 0 || val == arr[i-1]+1 {
			// Add to the current subarray if it's the first element or consecutive
			current = append(current, val)
		} else if val != arr[i-1] { // Reset current subarray if not consecutive
			if len(current) > len(longest) {
				longest = current
			}
			current = []int{val}
		}
	}

	// Check the last subarray
	if len(current) > len(longest) {
		longest = current
	}

	return longest
}

func MakeConsecutive(nums []int) []int {
	if len(nums) < 2 {
		return nums // If the slice has less than 2 elements, return it as is
	}

	consecutive := []int{}
	for i := 0; i < len(nums)-1; i++ {
		consecutive = append(consecutive, nums[i]) // Add the current element
		if nums[i] < nums[i+1] {
			// Add missing numbers between nums[i] and nums[i+1]
			for j := nums[i] + 1; j < nums[i+1]; j++ {
				consecutive = append(consecutive, j)
			}
		} else if nums[i] > nums[i+1] {
			// Add missing numbers between nums[i] and nums[i+1] (decreasing order)
			for j := nums[i] - 1; j > nums[i+1]; j-- {
				consecutive = append(consecutive, j)
			}
		}
	}
	// Add the last element of the slice
	consecutive = append(consecutive, nums[len(nums)-1])

	return consecutive
}

func IsEquallySpaced(nums []int) bool {
	if len(nums) < 2 {
		return true // A slice with 0 or 1 elements is trivially equally spaced
	}

	// Calculate the spacing between the first two elements
	spacing := nums[1] - nums[0]

	// Check if the spacing is consistent across the entire slice
	for i := 1; i < len(nums)-1; i++ {
		if nums[i+1]-nums[i] != spacing {
			return false
		}
	}

	return true
}
