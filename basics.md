# Packages and Import statements

* Two types of packages - Executable (aka `main`) and Reusable (aka `user-defined`)
* Executable packages run using `go build`; outputs a runnable/executable file 
* Reusable packages are used when you create libraries etc..
* Every file should belong to a package in Go. This is mentioned at the top of every file
  `package <package-name>` where package-name can be either `main` or an `user-defined` one
* You import packages that you want using `import` statement
* There are built-in packages like `fmt`, `math`, `crypto`, `io` etc..
  * Official docs of built-in/standard packages - https://golang.org/pkg/
  
Here's an eaxample of how an application's main package will look like that imports a variety of packages so it could use some functions from them
    - standard package `fmt`
    - user-defined packages `calculator` and `uploader`

![Appplication's main package](/images/application-main-package.png.png)

# Zero values

These are the default value that gets assigned to primitive type variables

![Zero values](/images/zero-values.png)

For ex., a declaration like this without any value assignment will get assigned an empty string `""`
`var name string`

# Value and Reference types

What types behave as value vs reference types when passed around-

![value-reference-types](/images/value-reference-types.png)

# Difference between Maps and Struct

Basic differences between Maps and Struct

![Maps-vs-Struct](/images/maps-vs-struct.png)