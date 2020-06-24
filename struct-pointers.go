package main

import "fmt"

// A struct is a group of fields that describe something
type product struct {
	code     int
	name     string
	quantity int
	color    string
}

type cart []product

func demoStructAndPointers() {
	// Creating user A's cart
	userACart := cart{
		product{10, "Shoes", 1, "red"},
		product{code: 20, name: "Shirt", quantity: 2},
	}

	updateQuantity(userACart[1], 4)
	fmt.Printf("User A's cart: %v\n", userACart)

	// Pointers
	// Pass by reference?
	updateQuantityByReference(&userACart[1], 4)
	fmt.Printf("User A's cart: %v", userACart)
}

func updateQuantity(p product, newQuantity int) {
	p.quantity = newQuantity
}

func updateQuantityByReference(p *product, newQuantity int) {
	p.quantity = newQuantity
}
