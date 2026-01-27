package main

import "fmt" // importing the fmt package for formatted I/O

func main() {
	var name string
	fmt.Println("Hello World from Go!")
	fmt.Scan(&name)
	fmt.Printf("Hello, %s\n", name)
}

//What does package mean?
// A package in Go is a way to organize and reuse code. It is similar to a library or module in other programming languages.
// a package is a way to group functions (where main is a function) together. and the main package is a special package that tells the Go compiler that this is the starting point of the program.
// Main contains lot of functions that are used to run the program.
//Every Go program starts with a main package and a main function.

// importing modules same as other languages
// import "fmt" is importing the fmt package, which contains functions for formatting text, including printing to the console.
// This package is one of the standard libraries in Go and is commonly used for input and output operations.

// Two ways to execute a Go program:
// 1. go run main.go : This command compiles and runs the Go program in a single step. It is useful for quick testing and development.
// 2. go build main.go : This command compiles the Go program and creates an executable file. You can then run the executable file separately. This is useful for deploying the program.
// ./main (on Unix-based systems) or main.exe (on Windows) is the name of the executable file created by the go build command.

//  When the program is executed, the code inside the main function is run first.
// fmt.Println("Hello World from Go!") is a function call that prints the string "Hello World from Go!" to the console, followed by a newline character.

// call code in an external package
// When you need your code to do something that might have been implemented by someone else, you can look for a package that has functions you can use in your code.
// You can import that package using the import statement, and then call the functions from that package in your code.

// pkg.go.dev is the official Go package documentation website where you can find information about various Go packages, including standard library packages and third-party packages. It provides documentation, examples, and usage instructions for each package, making it easier for developers to understand and utilize them in their projects.

// fmt.Println is a function from the fmt package that prints the specified string to the console.

// So Baiscally every GO program is made up of packages and functions
// Programs Start running in package main

// Tools
// Go lang
//gopls - Language server
// go-outline - Document symbols
// dlv - Debugger
// staticcheck - Linter
//
