package main

import (
	"fmt"
)

func main9() {
	// ============================================
	// DEFER: DELAY EXECUTION UNTIL FUNCTION EXITS
	// ============================================

	fmt.Println("Start")
	defer fmt.Println("This runs last!") // Deferred [web:112]
	fmt.Println("Middle")
	fmt.Println("End")
	// Output: Start, Middle, End, This runs last!

	// ============================================
	// MULTIPLE DEFERS: LIFO (Last In, First Out)
	// ============================================

	fmt.Println("\nMultiple defers:")
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
	// Output: 3, 2, 1 (reverse order!)

	// ============================================
	// DEFER USE CASE: CLEANUP
	// ============================================

	demoFileOperation()

	// ============================================
	// DEFER WITH CLOSURES
	// ============================================

	demoClosureDefer()

	// ============================================
	// PANIC & RECOVER
	// ============================================

	fmt.Println("\n=== Panic & Recover Demo ===")
	safeDivide1(10, 2) // Normal operation
	safeDivide1(10, 0) // Will panic, but recovered
	fmt.Println("Program continues after recovery!")

	// ============================================
	// RECOVER IN PRACTICE
	// ============================================

	result := robustOperation()
	fmt.Println("Result:", result)
}

// ============================================
// DEFER FOR RESOURCE CLEANUP
// ============================================

func demoFileOperation() {
	fmt.Println("\n=== File Operation Demo ===")

	// Simulate opening file
	file, err := openFile("data.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Defer cleanup - runs even if errors occur [web:112]
	defer closeFile(file)

	// Process file...
	fmt.Println("Processing file...")

	// defer closeFile() will run when function exits
}

func openFile(name string) (string, error) {
	fmt.Println("Opening file:", name)
	return name, nil
}

func closeFile(name string) {
	fmt.Println("Closing file:", name)
}

// ============================================
// DEFER WITH CLOSURES (Values captured!)
// ============================================

func demoClosureDefer() {
	fmt.Println("\n=== Defer Closure Demo ===")

	x := 10
	defer fmt.Println("Deferred x:", x) // Captures current value (10)

	x = 20
	fmt.Println("Current x:", x)
	// When function exits: "Deferred x: 10" (not 20!)

	// To capture updated value, use closure:
	defer func() {
		fmt.Println("Deferred closure x:", x) // Captures reference!
	}()
}

// ============================================
// PANIC: CRITICAL ERROR
// ============================================

func riskyOperation() {
	fmt.Println("Doing something risky...")
	panic("Something went wrong!") // Immediately stops execution [web:112]
	fmt.Println("This won't print")
}

// ============================================
// RECOVER: CATCH PANIC [web:112][web:115]
// ============================================

func safeDivide1(a, b int) {
	// Defer must be set BEFORE panic occurs
	defer func() {
		if r := recover(); r != nil {
			// recover() returns panic value [web:115]
			fmt.Printf("Recovered from panic: %v\n", r)
		}
	}()

	if b == 0 {
		panic("division by zero")
	}

	fmt.Printf("%d / %d = %d\n", a, b, a/b)
}

// ============================================
// PRODUCTION PATTERN: GRACEFUL ERROR HANDLING
// ============================================

func robustOperation() (result int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Operation failed, recovering gracefully...")
			result = -1 // Set default value on panic
		}
	}()

	// Simulate dangerous operation
	result = 100
	// panic("unexpected error") // Uncomment to test recovery

	return result
}

// ============================================
// REAL-WORLD: HTTP SERVER ERROR HANDLING
// ============================================

func handleRequest() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Request handler panicked:", r)
			// Log error, send 500 response, etc.
		}
	}()

	// Process request...
	// If panic occurs, it's caught and logged
}
