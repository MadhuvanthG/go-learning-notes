package main

import "fmt"

func demoBasicTypes() {
	// Demo how to declare variables and work with it
	// demoVariables()
	// demoTypes()
	demoDatastructures() // Arrays, slices etc..
	defineGlasses()      // Creating and using custom types
}

func demoVariables() {
	// variable declaration
	var productName string

	// variable assignment
	productName = "Eyeglasses"

	// ":=" is a short-hand to declare a new variable without the "var" keyword
	productType := "Optics"

	// For reassignments, just do a "=" and not a ":="
	productType = "Fashion"

	fmt.Println(productName, productType)
}

func demoTypes() {
}

func demoDatastructures() {
	// Array vs Slice
	// Array is a fixed number of elements, can't be resized - specify the length of the array
	var productsArray = [3]string{"watch", "eyeglasses", "phone"}

	// Slice
	// Slice- Associated with an underlying Array that can grow or shrink
	// Reference to a section of an array
	productsSlice := productsArray[:]
	productsSlice[1] = "Ray-Ban eyeglasses"
	fmt.Printf("Array %v", productsArray)
	fmt.Printf("Slice %v", productsSlice)

	// Growing slices safely
	// options - append, copy
	// Appends an element and returns a new slice
	productsSlice = append(productsSlice, "newWatch")

	// New slice returned by append above is pointing to a different array
	// Because we are appending 4th element but the existing underlying array's (i.e. productsArray) capacity is only 3
	// In other words, the new element we appended did not affect the productsArray
	fmt.Printf("Array %v", productsArray)
	fmt.Printf("Slice %v", productsSlice)

	// Half open range based retrieval slice[startIndexIncluding:upToNotIncluding]
	// Prints the first two elements at indices 0 and 1
	// Both startIndexIncluding and upToNotIncluding are optional
	fmt.Println(productsSlice[0:2])

	// For more on relationship between Arrays and Slices
	// https://blog.golang.org/go-slices-usage-and-internals

	// Byte slices

	// Maps
}

// Custom types and method sets
// A new type products that is a slice of strings
type products []string

// Receiver functions
// These are functions that will be available on any instance (or variable) of a specific type (ex. 'products')
// (p products) refers to the receiver of the getNames() method
// where p is a common naming convention to refer to a product instance (or variable) that this function is called upon
func (p products) getNames() []string {
	return p
}

func defineGlasses() {
	sunglasses := products{"Rayban X12"}

	fmt.Println(sunglasses.getNames())
}
