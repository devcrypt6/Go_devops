package main

import "fmt"

func main() {
	const name string = "Golang";
	// name =" Go" // This will give an error because constants cannot be changed
	// const pi:= 3.14 // type inference , but this syntax is invalid for constants
	const pi float64 = 3.14
	fmt.Println("Constant value:", name)

	const(
		port = 8080
		host = "localhost"
	)
	fmt.Println("Server running at", host, "on port", port)	
	
	main2()
}

// In Go, a constant is a variable whose value cannot be changed once it is assigned. 
// Constants are declared using the const keyword.
// In this example, we declare a constant named name of type string and assign it the value "Golang".
// If we try to change the value of the constant later in the code, it will result in a compilation error.
// Constants are useful for values that should remain unchanged throughout the program, such as mathematical constants or configuration values.

//constants can be declare at package level or inside functions.
// Constants can be of types: boolean, numeric (integer, float, complex), and string.
// Constants cannot be declared using the := syntax; they must be declared using the const keyword.
// Constants can also be grouped together using parentheses, similar to variable declarations.
// Example:
// const (
//     Pi = 3.14
//     E  = 2.71
// )
// This declares two constants, Pi and E, in a single block.


// Numeric Constants 
const (
	// Integer constant
    Big = 1 << 100 //  << is the left shift operator, shifting 1 by 100 bits to the left
	// 1<<n is equivalent to 2^n
	Small = Big >> 99 // Right shift Big by 99 bits to get 2

	//Basically x << n  is equal to x * 2^n
	// and x >> n is equal to x / 2^n
)

func needInt(x int) int {
	return x * 10 + 1
}

func needFloat(x float64) float64 {
	return x * 0.1
}

func main2() {
	fmt.Println(needInt(Small))  // OK: Small is 2
	fmt.Println(needFloat(Small)) // OK: Small is 2
	fmt.Println(needFloat(Big))   // OK: Big is too large for an int, but fine for a float64
	// fmt.Println(needInt(Big))    // Error: Big is too large for an int
}

// In Go, numeric constants are untyped by default, which means they can be used in different contexts without explicit type conversion.
// In the example above, we define two constants, Big and Small, using bitwise operations.
// The needInt function takes an integer argument and performs some arithmetic operations, while the needFloat function takes a float64 argument.
// When we call needInt with Small, it works fine because Small can fit within the range of an int.
// Similarly, calling needFloat with Small also works fine.
// However, when we try to call needInt with Big, it results in a compilation error because Big is too large to fit within the range of an int.
// On the other hand, calling needFloat with Big works fine because float64 can handle larger values.
// This demonstrates how untyped numeric constants in Go can be flexible and adapt to different types based on the context in which they are used.