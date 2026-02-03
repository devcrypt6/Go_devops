package main

import "fmt"

func main10() {
	// ============================================
	// FACTORIAL (Classic recursion)
	// ============================================

	fmt.Println("Factorial examples:")
	for i := 0; i <= 5; i++ {
		fmt.Printf("%d! = %d\n", i, factorial(i))
	}

	// ============================================
	// FIBONACCI (Multiple recursive calls)
	// ============================================

	fmt.Println("\nFibonacci sequence:")
	for i := 0; i <= 10; i++ {
		fmt.Printf("F(%d) = %d\n", i, fibonacci(i))
	}

	// ============================================
	// POWER (Divide and conquer)
	// ============================================

	fmt.Printf("\n2^10 = %d\n", power(2, 10))
	fmt.Printf("3^4 = %d\n", power(3, 4))

	// ============================================
	// SUM OF DIGITS (Recursive breakdown)
	// ============================================

	fmt.Printf("\nSum of digits(12345) = %d\n", sumDigits(12345))

	// ============================================
	// ARRAY SUM (Slice recursion)
	// ============================================

	numbers := []int{1, 2, 3, 4, 5}
	fmt.Printf("\nSum of %v = %d\n", numbers, sumSlice(numbers))

	// ============================================
	// PALINDROME CHECK
	// ============================================

	fmt.Println("\nPalindrome checks:")
	fmt.Printf("'racecar' is palindrome: %v\n", isPalindrome("racecar"))
	fmt.Printf("'hello' is palindrome: %v\n", isPalindrome("hello"))
}

// ============================================
// FACTORIAL: n! = n × (n-1)!
// ============================================

func factorial(n int) int {
	// Base case [web:125]
	if n == 0 {
		return 1
	}

	// Recursive case [web:122][web:125]
	return n * factorial(n-1)
}

// Execution trace for factorial(5):
// factorial(5) = 5 * factorial(4)
// factorial(4) = 4 * factorial(3)
// factorial(3) = 3 * factorial(2)
// factorial(2) = 2 * factorial(1)
// factorial(1) = 1 * factorial(0)
// factorial(0) = 1 (base case)
// Result: 5 × 4 × 3 × 2 × 1 = 120

// ============================================
// FIBONACCI: F(n) = F(n-1) + F(n-2)
// ============================================

func fibonacci(n int) int {
	// Base cases [web:122]
	if n <= 1 {
		return n
	}

	// Recursive case (two recursive calls)
	return fibonacci(n-1) + fibonacci(n-2)
}

// ============================================
// OPTIMIZED FIBONACCI (with memoization)
// ============================================

var fibCache = make(map[int]int)

func fibonacciMemo(n int) int {
	// Check cache first
	if val, ok := fibCache[n]; ok {
		return val
	}

	// Base cases
	if n <= 1 {
		return n
	}

	// Calculate and cache
	result := fibonacciMemo(n-1) + fibonacciMemo(n-2)
	fibCache[n] = result
	return result
}

// ============================================
// POWER: a^b
// ============================================

func power(base, exponent int) int {
	// Base case
	if exponent == 0 {
		return 1
	}

	// Optimization: divide and conquer
	if exponent%2 == 0 {
		half := power(base, exponent/2)
		return half * half
	}

	return base * power(base, exponent-1)
}

// ============================================
// SUM OF DIGITS
// ============================================

func sumDigits(n int) int {
	// Base case
	if n == 0 {
		return 0
	}

	// Recursive case: last digit + sum of remaining
	return (n % 10) + sumDigits(n/10)
}

// ============================================
// SUM OF SLICE (Recursive)
// ============================================

func sumSlice(numbers []int) int {
	// Base case: empty slice
	if len(numbers) == 0 {
		return 0
	}

	// Recursive case: first element + sum of rest
	return numbers[0] + sumSlice(numbers[1:])
}

// ============================================
// PALINDROME CHECK
// ============================================

func isPalindrome(s string) bool {
	// Base cases
	if len(s) <= 1 {
		return true
	}

	// Check first and last characters
	if s[0] != s[len(s)-1] {
		return false
	}

	// Recursively check middle
	return isPalindrome(s[1 : len(s)-1])
}

// ============================================
// GCD (Greatest Common Divisor) - Euclidean algorithm
// ============================================

func gcd(a, b int) int {
	// Base case
	if b == 0 {
		return a
	}

	// Recursive case
	return gcd(b, a%b)
}
