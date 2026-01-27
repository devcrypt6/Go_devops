package main


import (
	"fmt" // importing the fmt package for formatted I/O
	"math" // importing the math package for mathematical functions
	"math/rand" // importing the rand package for random number generation
)


func main() {
	fmt.Println("My favorite number is", math.Sqrt(16))
	fmt.Println("A random number:", rand.Intn(100))
	// fmt.Println( math.pi) // This will give an error because 'pi' should be 'Pi'
	fmt.Println( math.Pi)
}


// Here math is another package from the Go standard library that provides basic constants and mathematical functions.
// Package rand implements pseudo-random number generators.
// In this example, we are using the Sqrt function from the math package to calculate the square root of 16,


// and the Intn function from the rand package to generate a random integer between 0 and 99.
// This demonstrates how you can use multiple packages in a single Go program to perform different tasks.
// Instead of writing all the code from scratch, you can leverage existing packages to make your development process faster and easier.

//also instead of importing each package separately we can group them using parentheses as shown above. import ( "fmt" "math" "math/rand" )
// Note: math.pi will give an error because Go is case-sensitive and the correct constant name is math.Pi with an uppercase 'P'.
// So always remember to check the exact names of functions and constants in the package documentation.
// when importing external packages , a name is exported if it begins with a capital letter.
