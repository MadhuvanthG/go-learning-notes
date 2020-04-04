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
	// Array is a fixed number of elements - specify the length of the array
	var productsArray = [3]string{"watch", "eyeglasses", "phone"}
	productsArray[1] = "new watch"

	// Slice
	// Slice- Associated with an underlying Array that can grow or shrink
	var productsSlice = []string{"watch", "eyeglasses", "phone"}

	// Appends an element and returns a new slice
	// You'll have to assign the returned value to be able to refer to the new element
	productsSlice = append(productsSlice, "newWatch")

	// Range based retrieval slice[startIndexIncluding:upToNotIncluding]
	// Prints the first two elements
	fmt.Println(productsSlice[0:2])

	// Slices refer to the same array under the hood
	// So modifying the elements (not the slice itself) of a re-slice (newProductsSlice) modifies the elements of the original slice (productsSlice)
	newProductsSlice := productsSlice[0:2]
	newProductsSlice[1] = "eyeglassesFromSlice"

	// Both startIndexIncluding and upToNotIncluding are optional
	// this prints all the elements in the slice
	fmt.Println(productsSlice[0:])

	// this prints elements at 0, 1, 2
	fmt.Println(productsSlice[:3])

	// length and capacity of the productsSlice
	fmt.Println(len(productsSlice))
	fmt.Println(cap(productsSlice))

	// For more on Arrays vs Slices
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
