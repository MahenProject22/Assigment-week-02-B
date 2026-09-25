package main

import "fmt"

func main() {
	numbers := []int{39, 78, 85, 53, 32}
	target := 85
	foundIndex := -1
	for i := 0; i < len(numbers); i++ {
		if numbers[i] == target {
			foundIndex = i
			break
		}
	}
	if foundIndex != -1 {
		fmt.Println("Found at index:", foundIndex)
	} else {
		fmt.Println("Target not found")
	}
}
