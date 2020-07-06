# What is Go?
    * Conceived and Developed at Google
    * Compiled
    * Statically and strongly typed
# Why Go - including but not limited to:
    * Very strong native library support
    * Highly opinionated
    * Native tooling support is strong
    * Needs a lot more research and experience
    * More resources - https://github.com/golang/go/wiki/WhyGo
# Packages - Introduction
    * Packages is a collection of source files that perform a specific function
    * All applications have a `main` package with a function `main` which is the entry point
    * For better readability, reusability and maintainability
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
    * Creating struct literals
    * Updating fields using dot notation
    * Struct is a value type
    * Embedding structs
# Pointers
    * What are value types
    * What are reference types
    * When to use Pointers
      * There is no blueprint; down to the needs of your program
      * Some use-cases might be:
        * Pass mutable values around
        * Working with really large structs; what is really large is subjective, so benchmarking is a better tool to make a decision
      * Understand the effect on Garbage Collection when you use pointers
        * https://www.ardanlabs.com/blog/2017/05/language-mechanics-on-stacks-and-pointers.html
      * Good material - https://www.ardanlabs.com/blog/2014/12/using-pointers-in-go.html
      * Even better material
        * Benchmark performance and then make an informed decision
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
    * Can you build a custom package as an executable?
# Interfaces
    * Defines requirements explicitly
    * Helps reuse code
    * Good material - https://www.calhoun.io/how-do-interfaces-work-in-go/
# Concurrency - Channels and Go routines
    * Material to watch - Rob pike's talk - https://www.youtube.com/watch?v=f6kdp27TYZs
# Error handling
    * Defer and stacking defers
    * Panic
    * Recover
# gRPC
# http package
# Testing in Go
# Database interactions