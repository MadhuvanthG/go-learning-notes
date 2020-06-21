# What is Go?
    * Developed by Google
    * Compiled
    * Statically typed
    * 
    * 
    * 
    * 
    * 
# Why Go - <DIG DEEP>
    * Very strong and broad native library support
    * Highly opinionated be it-
      * formatting
      * code style
      * compartmentalizing
    * 
    * 
    * More resources - https://github.com/golang/go/wiki/WhyGo
# Packages - Introduction
    * Packages is a collection of source files that perform a specific function
    * All applications have a `main` package with a function `main` which is the entry point
# Variables, Constants and Zero values
# Arrays and Slices
    * What are slices? - Reference to underlying array
    * Creating a slice by slicing an array or slice using half-open range
    * Growing slices using append, copy
      * If the slice grows beyond original array's capacity, it's allocated a new and bigger underlying array
# Functions
    * Can omit types if two consecutive params share the same type
    * No default values and that's intentional by design
    * Can return multiple values and that's a common pattern used in functions that can throw an error
# Structs
    * Creating struct literals using order or field-names
    * Updating fields using dot notation
    * Struct is a value type
    * 
# Pointers
    * What are value types
    * What are reference types
# Methods
    * Object oriented approach vs Go approach
    * Functions with `receiver` argument
    * Defined on types
    * Available on any instance of the type on which the method is defined on
    * Allows you to write custom types to abstract group related logic
    * Pointer receivers vs Value receivers
# Custom Packages and Go Modules
    * Convention / Common practices
      * all files belonging to a package reside in a directory
      * name the folder same as the package name
    * Import statements
    * Exported names
# Interfaces
    * Defines requirements explicitly
    * Helps you encapsulate code
    * Good material - https://www.calhoun.io/how-do-interfaces-work-in-go/
# Concurrency - Channels and Go routines - <DIG DEEP>
# Error handling - <DIG DEEP>
    * Defer and stacking defers
    * Panic
    * Recover
# gRPC - <DIG DEEP>
# Testing in Go
# http package - <DIG DEEP>
# Database interactions