package main

import "fmt"

func main() {
	numbers := []int{39, 78, 85, 53, 32}
	target := 85
	for i := 0; i < len(numbers); i++ {
		fmt.Println("Checking index", i, "value", numbers[i])
		if numbers[i] == target {
			fmt.Println("Target found at index", i)
		}
	}
}
