package main

import (
	"fmt"
	"math"
)

func main3() {
	// ============================================
	// FLOATING-POINT NUMBERS
	// ============================================

	var price32 float32 = 19.99            // 6-7 decimal precision
	var precise float64 = 3.14159265358979 // 15-16 decimal precision

	fmt.Printf("float32: %.2f (precision: ~7 digits)\n", price32)
	fmt.Printf("float64: %.15f (precision: ~15 digits)\n", precise)

	// Scientific notation
	var lightSpeed float64 = 2.998e8 // 299,800,000 m/s
	fmt.Printf("Speed of light: %.2e m/s\n", lightSpeed)

	// Special values
	var infinity = math.Inf(1)  // Positive infinity
	var notANumber = math.NaN() // Not a Number

	fmt.Println("\nSpecial float values:")
	fmt.Printf("Infinity: %v, IsInf: %v\n", infinity, math.IsInf(infinity, 1))
	fmt.Printf("NaN: %v, IsNaN: %v\n", notANumber, math.IsNaN(notANumber))

	// ============================================
	// COMPLEX NUMBERS (for scientific computing)
	// ============================================

	var z1 complex64 = 3 + 4i
	var z2 complex128 = complex(5, 12) // Alternative syntax

	fmt.Println("\nComplex numbers:")
	fmt.Printf("z1 = %v\n", z1)
	fmt.Printf("z2 = %v\n", z2)
	fmt.Printf("Real part of z1: %v\n", real(z1))
	fmt.Printf("Imaginary part of z1: %v\n", imag(z1))

	// Complex multiplication
	result := z1 * complex64(z2)
	fmt.Printf("z1 * z2 = %v\n", result)

	//When to use float32 vs float64:

	// float64 - Default for most calculations, scientific computing, finance

	// float32 - Memory-constrained systems, graphics, ML models
}
