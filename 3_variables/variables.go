package main

import ( 
	  "fmt"
	  "math/cmplx"
)

var packageLevelVar string = "I am a package level variable"


var (	// Declaring multiple package-level variables
	IsthatGo  = true
	MaxInt	= 1<<63 - 1
	z complex128 = cmplx.Sqrt(-5 + 12i)
	euler    = 2.718
	complexN = cmplx.Sqrt(-5 + 12i)
)


func main() {
	 
	var name string = "hello"

	var sur_name = "world" // type inference

	fmt.Println("Variable value:", name)

	fmt.Println("Variable value with type inference:", sur_name)
	 // if you declare a variable without type, Go will infer the type based on the assigned value.

	// Short variable declaration
	age := 30 // type inferred as int
	fmt.Println("Age:", age)
	// outside a function, we cannot use := for variable declaration. It can only be used inside functions.
	
	// Multiple variable declaration
	var x, y int = 10, 20
	fmt.Println("x:", x, "y:", y)

	// int has so many types in GO like int8, int16, int32, int64 but internal CPU architecture decides which int to use when we use int keyword.

	var str string
	// if we want to assign the variable later we can do it like this
	str = "GoLang"
	fmt.Println("String value:", str)
	// in this case, short variable declaration (:=) cannot be used because it requires an initial value at the time of declaration.

	// Default zero values
	var isActive bool
	var score float64
	fmt.Println("Default bool value:", isActive) // false
	fmt.Println("Default float64 value:", score) // 0

	//floates 
	var pi float32 = 3.14
	// var pi = 3.14 // type inference as float64
	// pi:= 3.14 // short variable declaration as float64
	fmt.Println("Float32 value:", pi)

	// Scope of variables
	{
		var innerVar string = "I am inside a block"
		fmt.Println(innerVar)
	}
	// fmt.Println(innerVar) // This will give an error because innerVar is not accessible outside the block

	// but outside main function we can declare package level variables that can be accessed by all functions in the same package.
	fmt.Println(packageLevelVar)


	fmt.Printf("Type: %T Value: %v\n", IsthatGo, IsthatGo)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
	fmt.Printf("Type: %T Value: %v\n", euler, euler)
	fmt.Printf("Type: %T Value: %v\n", complexN, complexN)


	//Type conversion
	var intVar int = 42
	var floatVar float64 = float64(intVar) // converting int to float64
	fmt.Println("Converted float value:", floatVar)

	var anotherInt int = int(floatVar) // converting float64 back to int
	fmt.Println("Converted int value:", anotherInt)

	// Note: Go does not perform implicit type conversions. You must explicitly convert types when needed.
	// strings type conversion
	var strNum string = "123"
	// var intFromStr int = string(strNum[0]) // This will not convert string to int, it will give ASCII value of '1'
	var intFromStr int
	fmt.Println("String number:", strNum)
	fmt.Sscanf(strNum, "%d", &intFromStr)

}







// In GO, variables are used to store data values. A variable is a named storage location in memory that holds a value of a specific type.
// To declare a variable in Go, you use the var keyword followed by the variable name and its type. You can also assign an initial value to the variable using the = operator.
// In the example above, we declare a variable named name of type string and assign it the value "hello".

// if the declared variable is not used anywhere in the code, Go will throw a compilation error. This is because Go enforces the use of declared variables to help catch potential bugs and improve code quality.

// Go is a statically typed language, which means that the type of a variable must be known at compile time. Once a variable is declared with a specific type, it cannot hold values of other types.
// Go also supports type inference, which means that if you assign a value to a variable without specifying its type, Go will automatically infer the type based on the assigned value.
// For example, you can declare and initialize a variable like this:  name := "hello" . Here, Go infers that name is of type string based on the assigned value "hello".

// Variables in Go can be of various types, including basic types like int, float64, string, and bool, as well as more complex types like arrays, slices, maps, and structs.
// You can also declare multiple variables in a single line, like this: var x, y int = 10, 20

// Variables can also be declared at the package level (outside of any function) or at the function level (inside a function). Package-level variables have a broader scope and can be accessed by any function within the same package, while function-level variables are only accessible within the function where they are declared.

// In addition to the var keyword, Go also provides the short variable declaration syntax using :=, which can only be used inside functions. This syntax allows you to declare and initialize a variable in one step without explicitly specifying its type.
// For example:  count := 42  Here, Go infers that count is of type int based on the assigned value 42.

// Overall, variables are a fundamental concept in Go programming, and understanding how to declare, initialize, and use them is essential for writing effective Go code.

// Note: In Go, every variable must be used; otherwise, the compiler will raise an error for unused variables. This is a design choice in Go to help developers avoid potential bugs and maintain clean code.



// Basic data types in Go:
// 1. int: Represents integer values (e.g., 42, -7). int, int8, int16, int32, and int64 are different sizes of integers.
//         uint, uint8, uint16, uint32, and uint64 are unsigned integers (non-negative).
// 2. float32 and float64: Represent floating-point numbers (e.g., 3.14, -0.001).
// 3. string: Represents a sequence of characters (e.g., "hello", "GoLang").
// 4. bool: Represents boolean values (true or false).

// Each data type has its own set of operations and behaviors. For example, you can perform arithmetic operations on int and float types, concatenate strings, and use logical operators with bool types.

// When a variable is declared but not initialized, it gets a default zero value based on its type:
// - int: 0
// - float32/float64: 0.0
// - string: "" (empty string)
// - bool: false

// Understanding these basic data types and their behaviors is crucial for effective programming in Go.

