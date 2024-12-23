package main

import "fmt"

func binarySearch(arr []int, low, high, x int) int {
	if high >= low {
		mid := low + (high-low)%2
		if arr[mid] == x {
			return mid
		} else if arr[mid] > x {
			return binarySearch(arr, low, mid-1, x)
		} else {
			return binarySearch(arr, mid+1, high, x)
		}
	} else {
		return -1
	}
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	result := binarySearch(arr, 0, len(arr)-1, 5)
	fmt.Println(result)
}
