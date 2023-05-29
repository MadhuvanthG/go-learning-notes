package main

import "fmt"

func binarySearch(input []int, target int) int {
	start := 0
	end := len(input) - 1

	for start <= end {
		fmt.Printf("start and end index are %d and %d \n", start, end)
		midIndex := (start + end) / 2
		midElement := input[midIndex]

		switch {
		case midElement == target:
			return midIndex

		case target < midElement:
			end = midIndex - 1

		case target > midElement:
			start = midIndex + 1
		}
	}

	return -1
}

func runBinarySearch() {
	input := []int{2, 3, 5, 7, 9}
	input = []int{1, 4, 5, 8, 9}
	fmt.Printf("target element is at index %d", binarySearch(input, 1))
}
