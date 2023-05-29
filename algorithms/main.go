package main

import (
	"fmt"
	"go-learning-notes/algorithms/linkedlist"
)

func main() {
	// runFindItinerary()

	// runBinarySearch()

	// linked list testing
	ll := linkedlist.New()
	ll.Add("hello")
	ll.Add("world")
	ll.Add("linked")
	ll.Add("list")

	ll.Print()

	fmt.Printf("is java there in linked list: %v \n", ll.Contains("java"))
	ll.PollFirst()
	ll.Print()

	ll.Add("hello")
	ll.Print()

	ll.Remove("linked")
	ll.Print()
}
