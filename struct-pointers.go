package main

import "fmt"

// A struct is a group of fields that describe something
// Similar to object in JS
type product struct {
	code int
	name string
}

type cart []product

func demoStructAndPointers() {
	user1Cart := cart{product{23, "kitkat"}, product{25, "wristwatch"}}

	// Assigning value to a field in struct
	// We update the code of first product directly
	user1Cart[0].code = 27

	// To update code of second product, let's use a receiver function `updateCode`
	user1Cart[1].updateCode(30)

	// Now, the new code 30 will NOT be updated when you print out
	// Because Go passes by value by default
	fmt.Printf("%v", user1Cart)

	// Pointers is the answer, if you want to update code on the product it's called on
	// Note that the only difference is which function we're calling (`updateCode` vs `updateProductCode`)
	// `updateProductCode` receives a pointer to a product and hence always operates on the product the function is called on
	user1Cart[1].updateProductCode(35)

	fmt.Printf("%v", user1Cart)
}

// Receiver functions on struct
// a function to update code of any given product
func (p product) updateCode(newCode int) {
	p.code = newCode
}

// Receives pointer to a product
// So always operates on the product the function is called on
func (p *product) updateProductCode(newCode int) {
	p.code = newCode
}
