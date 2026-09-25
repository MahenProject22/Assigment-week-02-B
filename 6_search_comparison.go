package main

import "fmt"

func linearSearch(arr []int, target int) (int, int) {
	comparisons := 0

	for i := 0; i < len(arr); i++ {
		comparisons++

		if arr[i] == target {
			return i, comparisons
		}
	}

	return -1, comparisons
}

func binarySearch(arr []int, target int) (int, int) {
	left := 0
	right := len(arr) - 1
	comparisons := 0

	for left <= right {
		mid := (left + right) / 2
		comparisons++

		if arr[mid] == target {
			return mid, comparisons
		} else if target < arr[mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return -1, comparisons
}

func main() {
	studentIDs := []int{39, 78, 85, 53, 32}

	target1 := 85

	linearIndex, linearComparisons := linearSearch(studentIDs, target1)
	binaryIndex, binaryComparisons := binarySearch(studentIDs, target1)

	fmt.Println("Searching for:", target1)

	fmt.Println("Linear Search:")
	fmt.Println("Index:", linearIndex)
	fmt.Println("Comparisons:", linearComparisons)

	fmt.Println("Binary Search:")
	fmt.Println("Index:", binaryIndex)
	fmt.Println("Comparisons:", binaryComparisons)

	fmt.Println()

	target2 := 85

	linearIndex, linearComparisons = linearSearch(studentIDs, target2)
	binaryIndex, binaryComparisons = binarySearch(studentIDs, target2)

	fmt.Println("Searching for:", target2)

	fmt.Println("Linear Search:")
	fmt.Println("Index:", linearIndex)
	fmt.Println("Comparisons:", linearComparisons)

	fmt.Println("Binary Search:")
	fmt.Println("Index:", binaryIndex)
	fmt.Println("Comparisons:", binaryComparisons)
}
