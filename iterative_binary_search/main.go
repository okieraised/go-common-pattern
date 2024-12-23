package main

import "fmt"

func binarySearch(arr []int, low, high, x int) int {
	for low <= high {
		mid := low + (high-low)%2
		if arr[mid] == x {
			return mid
		} else if arr[mid] < x {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	result := binarySearch(arr, 0, len(arr)-1, 6)
	fmt.Println(result)
}
