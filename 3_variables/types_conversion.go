package main

import (
    "fmt"
    "strconv"
)

func main8() {
    // ============================================
    // NUMERIC CONVERSIONS (explicit only!)
    // ============================================
    
    var i int = 42
    var f float64 = float64(i)  // int → float64
    var u uint = uint(f)         // float64 → uint
    
    fmt.Printf("int: %d → float64: %.2f → uint: %d\n", i, f, u)
    
    // Precision loss warning!
    var bigFloat float64 = 3.14159
    var smallInt int = int(bigFloat) // Truncates to 3!
    fmt.Printf("\nPrecision loss: %.5f → %d (decimal lost!)\n", bigFloat, smallInt)
    
    // ============================================
    // STRING CONVERSIONS
    // ============================================
    
    // Number → String
    num := 123
    str := strconv.Itoa(num) // Integer to ASCII
    fmt.Printf("\nNumber to string: %d → \"%s\" (type: %T)\n", num, str, str)
    
    // String → Number
    str2 := "456"
    num2, err := strconv.Atoi(str2) // ASCII to Integer
    if err != nil {
        fmt.Println("Conversion error:", err)
    } else {
        fmt.Printf("String to number: \"%s\" → %d (type: %T)\n", str2, num2, num2)
    }
    
    // String → Float
    str3 := "3.14"
    float, err := strconv.ParseFloat(str3, 64) // 64-bit precision
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Printf("String to float: \"%s\" → %.2f\n", str3, float)
    }
    
    // Float → String
    pi := 3.14159265
    piStr := fmt.Sprintf("%.2f", pi) // Format to 2 decimals
    fmt.Printf("Float to string: %f → \"%s\"\n", pi, piStr)
    
    // ============================================
    // RUNE & BYTE CONVERSIONS
    // ============================================
    
    char := 'A'
    charCode := int(char) // rune → int
    fmt.Printf("\nChar '%c' has ASCII code: %d\n", char, charCode)
    
    code := 66
    charFromCode := rune(code) // int → rune
    fmt.Printf("ASCII %d represents: '%c'\n", code, charFromCode)

	// Byte slice to string
	byteSlice := []byte{72, 101, 108, 108, 111}
	strFromBytes := string(byteSlice)
	fmt.Printf("\nByte slice %v to string: \"%s\"\n", byteSlice, strFromBytes)

	// String to byte slice
	originalStr := "World"
	bytesFromStr := []byte(originalStr)
	fmt.Printf("String \"%s\" to byte slice: %v\n", originalStr, bytesFromStr)

	// Note: Go does not perform implicit type conversions. You must explicitly convert types when needed.


	//var a int = 10
	 // var b float64 = a     //  COMPILE ERROR!
	 // var b float64 = float64(a) //  Correct
}
