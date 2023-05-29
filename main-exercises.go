package main

import "strings"

// solutions for exercises in https://go.dev/tour
func Pic(dx, dy int) [][]uint8 {
	images := make([][]uint8, dy)

	for i := 0; i < dy; i++ {
		thisImage := make([]uint8, dx)
		for j := 0; j < dx; j++ {
			value := 2 * 2
			thisImage = append(thisImage, uint8(value))
		}

		images = append(images, thisImage)
	}

	return images
}

func WordCount(s string) map[string]int {
	result := make(map[string]int)
	split := strings.Fields(s)

	for _, word := range split {
		result[word] = result[word] + 1
	}

	return result
}

func fibonacci() func() int {
	first, second := 0, 1
	return func() int {
		old := first
		first, second = second, first+second
		return old
	}
}
