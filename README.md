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

    * ConstSpec and iota could be used to achieve enum-like behavior - https://play.golang.org/p/dPLyGUKu-hJ

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

# Maps

- Go's implementation of Maps is a HashMap under the hood, so best case is O(1)
- https://dave.cheney.net/2018/05/29/how-the-go-runtime-implements-maps-efficiently-without-generics

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
      * Good material - https://www.ardanlabs.com/blog/2014/12/using-pointers-in-go.html. Excerpt from it:
        * Think of every struct as having a nature. If the nature of the struct is something that should not be changed, like a time, a color or a coordinate, then implement the struct as a primitive data value. If the nature of the struct is something that can be changed, even if it never is in your program, it is not a primitive data value and should be implemented to be shared with a pointer. Don’t create structs that have a duality of nature.
      * Even better material
        * Benchmark performance and then make an informed decision

# Methods

    * Object oriented approach vs Go approach
    * Functions with `receiver` argument
    * Defined on types
    * Available on any instance of the type on which the method is defined on
    * Allows you to write custom types to abstract/group related logic
    * Pointer receivers vs Value receivers

# Custom Packages and Go Modules

    * Convention / Common practices
      * all files belonging to a package reside in a directory
      * name the folder same as the package name
    * Import statements
    * Exported names
    * Go modules is Go's dependency management system
    * Good material - https://blog.golang.org/using-go-modules

# Interfaces

    * Defines requirements explicitly
    * Code encapsulation and reuse
    * Implicit conformance
    * Empty interfaces
    * A variable of interface type is stored as a pair of (type, value)
    * Good material - https://www.calhoun.io/how-do-interfaces-work-in-go/

# Reflect

    * Reflect is basically used to "examine" a variable's (of any type) fields and values
    * Some examples of examination-
      * how many fields does this variable have?
      * how many methods does the type of this variable have?
      * does this variable of type "X" implement interface of type "Y"?
    * One of the downsides as I see it-
      * losing type information
      * a lot of methods in the Reflect package could "panic" (at runtime), so you need to handle errors well most likely using "CanXXX" methods of the reflect package
    * Values can be "set" on the variable you're reflecting as well. But only if it's "settable". IOW, only if you are reflecting on a variable which was passed by reference
    * Good material - https://blog.golang.org/laws-of-reflection

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
