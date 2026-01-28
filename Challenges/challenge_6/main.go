

// Convert between:

// Decimal ↔ Binary

// Decimal ↔ Hexadecimal

// Decimal ↔ Octal

package main

import (
    "fmt"
    "strconv"
)

func decimalToBinary(n int) string {
    return strconv.FormatInt(int64(n), 2)
}

func decimalToHex(n int) string {
    return strconv.FormatInt(int64(n), 16)
}

func decimalToOctal(n int) string {
    return strconv.FormatInt(int64(n), 8)
}

func binaryToDecimal(binary string) (int, error) {
    result, err := strconv.ParseInt(binary, 2, 64)
    return int(result), err
}

func hexToDecimal(hex string) (int, error) {
    result, err := strconv.ParseInt(hex, 16, 64)
    return int(result), err
}

func main() {
    fmt.Println("=== Number System Converter ===")
    fmt.Println("1. Decimal to Binary")
    fmt.Println("2. Decimal to Hexadecimal")
    fmt.Println("3. Decimal to Octal")
    fmt.Println("4. Binary to Decimal")
    fmt.Println("5. Hex to Decimal")
    
    var choice int
    fmt.Print("\nChoose option: ")
    fmt.Scan(&choice)
    
    // YOUR CODE HERE
    // Implement menu-driven converter
    switch choice {
    case 1:
        var dec int
        fmt.Print("Enter decimal number: ")
        fmt.Scan(&dec)
        fmt.Printf("Binary: %s\n", decimalToBinary(dec))
    case 2:
        var dec int
        fmt.Print("Enter decimal number: ")
        fmt.Scan(&dec)
        fmt.Printf("Hexadecimal: %s\n", decimalToHex(dec))
    case 3:
        var dec int
        fmt.Print("Enter decimal number: ")
        fmt.Scan(&dec)
        fmt.Printf("Octal: %s\n", decimalToOctal(dec))
    case 4:
        var bin string
        fmt.Print("Enter binary number: ")
        fmt.Scan(&bin)
        dec, err := binaryToDecimal(bin)
        if err != nil {
            fmt.Println("Error:", err)
        } else {
            fmt.Printf("Decimal: %d\n", dec)
        }
    case 5:
        var hex string
        fmt.Print("Enter hexadecimal number: ") 
        fmt.Scan(&hex)
        dec, err := hexToDecimal(hex)
        if err != nil {
            fmt.Println("Error:", err)
        } else {
            fmt.Printf("Decimal: %d\n", dec)
        }
    default:
        fmt.Println("Invalid choice")
    }
}
