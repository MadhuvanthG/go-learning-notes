package main

import (
	"fmt"
	"reflect"
)

// A struct representing attributes of a single item in a database
type DatabaseItem struct {
	Id      int
	Created string
	Updated string
	Key     string
}

// A method to loop over the database item and prepare to update the values to the table
// Preparation could mean, building an expression with key name and value, like in case of DynamoDB
func (di DatabaseItem) prepareItemForUpdate() {
	// 1. Loop over each item in the "DatabaseItem" struct
	// to build an expression that is compatible with your database queries

	// How to loop over a struct's fields?
	// Treat struct as an "empty" interface
	// Reflect on its type and value to "examine" its fields
	// Reflect is basically used to "examine" a variable's (of any type) fields and values
	// Some examples of examination-
	// how many fields does this variable have?
	// how many methods does the type of this variable have?
	// does this variable of type "X" implement interface of type "Y"?
	// ...

	diReflectionValue := reflect.ValueOf(di)
	diReflectionType := reflect.TypeOf(di)

	// Loop over my struct and print all of its field names and values
	for i := 0; i < diReflectionType.NumField(); i++ {
		key := diReflectionType.Field(i).Name
		value := diReflectionValue.FieldByName(key).Interface()

		fmt.Println(key, value)
	}
}
