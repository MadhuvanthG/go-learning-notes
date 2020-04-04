package main

import "fmt"

func demoMaps() {
	// Maps is a collection of key value pairs
	// Keys should of the same type and similarly, values should be of the same type
	// Unlike structs where values can be of different types

	colors := map[string]string{"red": "#ff0000", "white": "#ffffff"}

	fmt.Println(colors)
}
