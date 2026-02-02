package main

import "fmt"

func main2() {
    // ============================================
    // PARAMETERS vs ARGUMENTS
    // ============================================
    
    // Parameters: Variables in function definition
    // Arguments: Actual values passed when calling
    
    printInfo("Go", 2009)  // "Go" and 2009 are arguments
    
    // ============================================
    // MULTIPLE PARAMETERS
    // ============================================
    
    displayPerson("TOM", 25, "Berlin")
    
    // ============================================
    // SAME TYPE PARAMETERS (Shorthand)
    // ============================================
    
    fmt.Println("Sum:", sum(10, 20, 30))
    
    // ============================================
    // MIXED TYPE PARAMETERS
    // ============================================
    
    introduce("tom", 30, 1.75)
}

// ============================================
// FUNCTION WITH PARAMETERS (Full syntax)
// ============================================

func printInfo(language string, year int) {
    // language and year are parameters
    fmt.Printf("%s was created in %d\n", language, year)
}

// ============================================
// MULTIPLE PARAMETERS
// ============================================

func displayPerson(name string, age int, city string) {
    fmt.Printf("Name: %s, Age: %d, City: %s\n", name, age, city)
}

// ============================================
// SAME TYPE PARAMETERS (Shorthand)
// ============================================

// Instead of: func sum(a int, b int, c int)
func sum(a, b, c int) int {
    return a + b + c
}

// ============================================
// MIXED TYPE PARAMETERS
// ============================================

func introduce(name string, age int, height float64) {
    fmt.Printf("%s is %d years old and %.2fm tall\n", name, age, height)
}

// ============================================
// REQUIRED ORDER
// ============================================

// Parameters must be passed in exact order!
func createUser(username string, age int, isActive bool) {
    fmt.Printf("User: %s, Age: %d, Active: %v\n", username, age, isActive)
}

// CORRECT:
// createUser("alice", 25, true)

// WRONG (compilation error):
// createUser(25, "alice", true)  // Types don't match!
