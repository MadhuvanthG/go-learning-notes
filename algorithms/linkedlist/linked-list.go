package linkedlist

import (
	"fmt"
	"strings"
)

// Implementation of a SINGLE LINKED LIST
type linkedList struct {
	start *node
	end   *node
}

type node struct {
	value string
	next  *node
}

func New() linkedList {
	return linkedList{}
}

// pacakge visibility - to assess
// 1. should func names be capitalized
// 2. "linkedList" is unexported, how can the caller then work with it? can it even declare a value to be of this type?
func (l *linkedList) Add(element string) {
	newNode := node{
		value: element,
	}

	if l.start == nil && l.end == nil {
		// this is our first node
		l.start = &newNode
		l.end = &newNode
		return
	}

	l.end.next = &newNode
	l.end = &newNode
}

func (l *linkedList) Remove(element string) bool {
	if l.start.value == element {
		l.start = l.start.next
		return true
	}

	previousNode := l.start
	currentNode := l.start
	for currentNode.next != nil {
		if currentNode.value == element {
			previousNode.next = currentNode.next
			return true
		}
		previousNode = currentNode
		currentNode = currentNode.next
	}

	return false
}

func (l *linkedList) Contains(element string) bool {
	currentNode := l.start
	for currentNode.next != nil {
		if currentNode.value == element {
			return true
		}
		currentNode = currentNode.next
	}

	return false
}

func (l *linkedList) PollFirst() string {
	if l.start == nil {
		return ""
	}

	target := l.start.value
	l.start = l.start.next

	return target
}

/*
* not possible to implement this without a doubly linked list
func (l *linkedList) pollLast() string {
	if l.end == nil {
		return ""
	}



}
*/

func (l *linkedList) Print() {
	builder := strings.Builder{}

	currentNode := l.start
	for currentNode.next != nil {
		builder.WriteString(fmt.Sprintf("%s -> ", currentNode.value))
		currentNode = currentNode.next
	}
	// append the last node
	builder.WriteString(fmt.Sprintf("%s -> ", currentNode.value))

	fmt.Printf("Linked list: %s \n", (&builder).String())
}
