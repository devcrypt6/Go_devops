package main

import (
	"fmt"
)

func main2() {
	// ============================================
	// SIGNED INTEGERS (can be negative)
	// ============================================

	var age int8 = 25                      // -128 to 127 (1 byte)
	var population int16 = 30000           // -32,768 to 32,767 (2 bytes)
	var distance int32 = 1000000           // -2.1 billion to 2.1 billion (4 bytes)
	var bigNum int64 = 9223372036854775807 // Huge! (8 bytes)

	//int8 is small , int16 is medium, int32 is large, int64 is huge

	// Platform-dependent (32-bit or 64-bit based on system)
	// var counter int = 100 // Most commonly used!

	fmt.Printf("int8:  %d (size: %d bytes)\n", age, 1)
	fmt.Printf("int16: %d (size: %d bytes)\n", population, 2)
	fmt.Printf("int32: %d (size: %d bytes)\n", distance, 4)
	fmt.Printf("int64: %d (size: %d bytes)\n", bigNum, 8)

	// ============================================
	// UNSIGNED INTEGERS (only positive)
	// ============================================

	var smallUint uint8 = 255                  // 0 to 255 (also called "byte")
	var mediumUint uint16 = 65535              // 0 to 65,535
	var largeUint uint32 = 4294967295          // 0 to 4.3 billion
	// var hugeUint uint64 = 18446744073709551615 // Massive!

	// var pointer uint = 100 // Platform-dependent

	fmt.Println("\nUnsigned integers:")
	fmt.Printf("uint8:  %d (max: 255)\n", smallUint)
	fmt.Printf("uint16: %d (max: 65,535)\n", mediumUint)
	fmt.Printf("uint32: %d\n", largeUint)

	// ============================================
	// SPECIAL TYPES
	// ============================================

	var charCode byte = 65     // Alias for uint8 (used for ASCII)
	var unicodeChar rune = 'A' // Alias for int32 (used for Unicode)

	fmt.Println("\nSpecial types:")
	fmt.Printf("byte (uint8): %d = '%c'\n", charCode, charCode)
	fmt.Printf("rune (int32): %d = '%c'\n", unicodeChar, unicodeChar)
	fmt.Printf("Unicode rune: %c = %d\n", '🚀', '🚀')
}

//how to run this code:
// go run 3_variables/numeric_types.go 	
// but there are two main packages in this project, so you need to specify the path to the file
// so run the main2 function, you need to change the name of the main function to main2