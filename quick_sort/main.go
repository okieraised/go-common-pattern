package main

import "fmt"

func partition(arr []int, low, high int) int {
	// Choose the pivot
	pivot := arr[high]

	// Index of smaller element and indicates
	// the right position of pivot found so far
	i := low - 1

	// Traverse arr[low:high] and move all smaller
	// elements to the left side. Elements from low to
	// i are smaller after every iteration
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			i += 1
			arr = swap(arr, i, j)
		}
	}
	// Move pivot after smaller elements and return its position
	arr = swap(arr, i+1, high)
	return i + 1
}

func swap(arr []int, i, j int) []int {
	arr[i], arr[j] = arr[j], arr[i]
	return arr
}

func quickSort(arr []int, low, high int) {
	if low < high {
		// pi is the partition return index of pivot
		pi := partition(arr, low, high)

		// Recursion calls for smaller elements and greater or equals elements
		quickSort(arr, low, pi-1)
		quickSort(arr, pi+1, high)
	}
}

func main() {
	arr := []int{10, 7, 8, 9, 1, 5, 5, 9, 20}
	n := len(arr)

	quickSort(arr, 0, n-1)

	fmt.Println("arr", arr)
}
