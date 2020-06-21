package main

import "fmt"

// A struct is a group of fields that describe something
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
	fmt.Printf("%v", user1Cart)

	// To update code of second product, let's use a function `updateCode`
	updateCode(user1Cart[1], 30)
	fmt.Printf("%v", user1Cart)

	// However, the updated code doesn't reflect in the user's cart
	// Because, struct passes by "value"
	// What if you want the updated code to reflect in the cart

	// Welcome to pointers!
	// Create a reference to second product of user's cart using "&"
	secondProduct := &user1Cart[1]

	// pass the struct by reference it to the function
	updateCodeByReference(secondProduct, 30)
	fmt.Printf("%v", user1Cart)
}

func updateCode(product product, newCode int) {
	product.code = newCode
}

// Accepts a pointer (see "*") to type "product"
func updateCodeByReference(product *product, newCode int) {
	product.code = newCode
}
