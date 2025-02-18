package main

import "fmt"

func merge(arr []int, left, mid, right int) {
	n1 := mid - left + 1
	n2 := right - mid

	// Create temp arrays
	L := make([]int, n1)
	R := make([]int, n2)

	for i := 0; i < n1; i++ {
		L[i] = arr[left+i]
	}

	for j := 0; j < n2; j++ {
		R[j] = arr[mid+1+j]
	}

	i := 0    // Initial index of first subarray
	j := 0    // Initial index of second subarray
	k := left // Initial index of merged subarray

	// Merge the temp arrays back into arr[left..right]
	for i < n1 && j < n2 {
		if L[i] < R[j] {
			arr[k] = L[i]
			i += 1
		} else {
			arr[k] = R[j]
			j += 1
		}
		k += 1
	}

	// Copy the remaining elements of L[], if there are any
	for i < n1 {
		arr[k] = L[i]
		i += 1
		k += 1
	}

	// Copy the remaining elements of R[], if there are any
	for j < n2 {
		arr[k] = R[j]
		j += 1
		k += 1
	}
}

func mergeSort(arr []int, left, right int) {
	if left < right {
		mid := left + (right-left)/2

		mergeSort(arr, left, mid)
		mergeSort(arr, mid+1, right)
		merge(arr, left, mid, right)
	}
}

func main() {
	arr := []int{10, 7, 8, 9, 1, 5, 20}
	mergeSort(arr, 0, len(arr)-1)

	fmt.Println("arr", arr)
}
