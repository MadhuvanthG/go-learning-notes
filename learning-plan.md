# Packages, Import statements, Exported names
# Variable declarations and zero values
# Functions
    * Can omit types if two consecutive params share the same type
    * No default values and that's intentional by design
    * Can return multiple values and that's a common pattern used in functions that can throw an error
    * `defer` and stacking defers
# Arrays and Slices
    * What are slices? - Reference to underlying array
    * Creating a slice by slicing an array or slice using half-open range
    * Growing slices using append, copy
# Structs
# Pointers (value vs reference types)
# Methods
    * Object oriented approach vs Go approach
    * Functions with `receiver` argument
    * Defined on types
    * Available on any instance of the type on which the method is defined on
    * Allows you to write custom types to abstract group related logic
    * Pointer receivers vs Value receivers
    * see udemy course material
# Interfaces
    * Good material - https://www.calhoun.io/how-do-interfaces-work-in-go/
# Error handling - <DIG DEEP>
    * Defer and stacking defers
    * Panic
    * Recover
# Channels and Go routines - <DIG DEEP>
# Modules - <DIG DEEP>
# gRPC - <DIG DEEP>
# Testing in Go
# http package - <DIG DEEP>
# Database interactions