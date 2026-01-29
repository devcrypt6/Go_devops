package main

import "fmt"

// IfElseDemo demonstrates if-else control flow
func IfElseDemo() {
	// ============================================
	// SIMPLE IF
	// ============================================

	age := 20

	if age >= 18 {
		fmt.Println("You are an adult")
	}

	// ============================================
	// IF-ELSE
	// ============================================

	temperature := 25

	if temperature > 30 {
		fmt.Println("It's hot outside!")
	} else {
		fmt.Println("Weather is pleasant")
	}

	// ============================================
	// IF-ELSE IF-ELSE
	// ============================================

	score := 85

	if score >= 90 {
		fmt.Println("Grade: A (Excellent!)")
	} else if score >= 80 {
		fmt.Println("Grade: B (Good!)")
	} else if score >= 70 {
		fmt.Println("Grade: C (Fair)")
	} else if score >= 60 {
		fmt.Println("Grade: D (Pass)")
	} else {
		fmt.Println("Grade: F (Fail)")
	}

	// ============================================
	// NESTED IF (Use sparingly!)
	// ============================================

	isStudent := true
	hasPassed := true

	if isStudent {
		if hasPassed {
			fmt.Println("Congratulations on passing!")
		} else {
			fmt.Println("Better luck next time")
		}
	}

	// Better: Use logical AND
	if isStudent && hasPassed {
		fmt.Println("Student passed ✓")
	}
}
