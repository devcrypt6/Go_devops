package main

import "fmt"

func main6() {
    a := 60  // Binary: 0011 1100
    b := 13  // Binary: 0000 1101
    
    fmt.Printf("a = %d (binary: %08b)\n", a, a)
    fmt.Printf("b = %d (binary: %08b)\n\n", b, b)
    
    // Bitwise AND (&)
    result := a & b // 0000 1100 = 12
    fmt.Printf("a & b  = %d (binary: %08b) [AND]\n", result, result)
    
    // Bitwise OR (|)
    result = a | b // 0011 1101 = 61
    fmt.Printf("a | b  = %d (binary: %08b) [OR]\n", result, result)
    
    // Bitwise XOR (^)
    result = a ^ b // 0011 0001 = 49
    fmt.Printf("a ^ b  = %d (binary: %08b) [XOR]\n", result, result)
    
    // Left shift (<<)
    result = a << 2 // 1111 0000 = 240 (multiply by 4)
    fmt.Printf("a << 2 = %d (binary: %08b) [Left shift = multiply by 2^n]\n", result, result)
    
    // Right shift (>>)
    result = a >> 2 // 0000 1111 = 15 (divide by 4)
    fmt.Printf("a >> 2 = %d (binary: %08b) [Right shift = divide by 2^n]\n", result, result)
    
    // Practical use: Check if number is even/odd
    num := 17
    if num & 1 == 0 {
        fmt.Printf("\n%d is even\n", num)
    } else {
        fmt.Printf("\n%d is odd\n", num)
    }


	//Real-World Uses of Bitwise Operators:

	// Permissions: chmod 755 in Unix

	// Flags: Enable/disable features efficiently

	// Cryptography: XOR encryption

	// Network protocols: IP address manipulation

	// Performance: Faster than multiplication/division by powers of 2
}
