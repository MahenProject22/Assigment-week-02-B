package main

import "fmt"

func linearSearch(numbers []int, target int) int {
	for i := 0; i < len(numbers); i++ {
		if numbers[i] == target {
			return i
		}
	}
	return -1
}

func main() {
	numbers := []int{39, 78, 85, 53, 32}

	// target exists
	result1 := linearSearch(numbers, 85)
	fmt.Println("Search 85, found at:", result1)

	// target does not exist
	result2 := linearSearch(numbers, 100)
	fmt.Println("Search 100, result:", result2)
}
