package main

import "fmt"

// ============================================
// PACKAGE-LEVEL VARIABLES (Accessible to all functions)
// ============================================

var globalCounter int = 0

func main5() {
	// ============================================
	// LOCAL SCOPE
	// ============================================

	x := 10 // Local to main()
	fmt.Println("x in main:", x)

	demo()
	// x is not accessible here from demo()

	// ============================================
	// SHADOWING (Be careful!)
	// ============================================

	y := 5
	fmt.Println("Outer y:", y)

	if true {
		y := 10                    // Different variable!
		fmt.Println("Inner y:", y) // 10
	}

	fmt.Println("Outer y again:", y) // Still 5

	// ============================================
	// FUNCTION-LEVEL SCOPE
	// ============================================

	incrementCounter()
	incrementCounter()
	fmt.Println("Global counter:", globalCounter)
}

// ============================================
// FUNCTION WITH LOCAL VARIABLES
// ============================================

func demo() {
	x := 20 // Different x from main()
	fmt.Println("x in demo:", x)
}

// ============================================
// ACCESSING PACKAGE-LEVEL VARIABLES
// ============================================

func incrementCounter() {
	globalCounter++ // Can access global variable
}

// ============================================
// PARAMETERS SHADOW GLOBAL VARIABLES
// ============================================

var name string = "Global"

func printName(name string) { // Parameter shadows global
	fmt.Println("Name:", name) // Uses parameter, not global
}

// ============================================
// BLOCK SCOPE
// ============================================

func scopeExample() {
	x := 1

	if x > 0 {
		y := 2 // y only exists in this if block
		fmt.Println(x, y)
	}

	// fmt.Println(y)  // ERROR: y not defined here

	for i := 0; i < 3; i++ {
		z := i * 2 // z only exists in this loop
		fmt.Println(z)
	}

	// fmt.Println(z)  // ERROR: z not defined here
}
