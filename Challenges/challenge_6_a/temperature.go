package main

import "fmt"

// Convert Celsius to Fahrenheit
func celsiusToFahrenheit(celsius float64) float64 {

    // Formula: F = C × 9/5 + 32
	fahrenheit := celsius*9/5 + 32
	return fahrenheit
}

// Convert Fahrenheit to Celsius
func fahrenheitToCelsius(fahrenheit float64) float64 {

    // Formula: C = (F - 32) × 5/9
	celsius := (fahrenheit - 32) * 5 / 9
	return celsius
}

// Convert Celsius to Kelvin
func celsiusToKelvin(celsius float64) float64 {

    // Formula: K = C + 273.15
	kelvin := celsius + 273.15
	return kelvin
}	

// Convert Kelvin to Celsius
func kelvinToCelsius(kelvin float64) float64 {

	// Formula: C = K - 273.15
	celsius := kelvin - 273.15
	return celsius
}

// Convert temperature between any two units
// Returns converted value and error if invalid unit
func convertTemperature(value float64, fromUnit, toUnit string) (float64, error) {

    // Support: "C", "F", "K"
    // Example: convertTemperature(100, "C", "F") → 212, nil
	var celsiusValue float64
	switch fromUnit {
	case "C":
		celsiusValue = value
	case "F":
		celsiusValue = fahrenheitToCelsius(value)
	case "K":
		celsiusValue = kelvinToCelsius(value)
	default:
		return 0, fmt.Errorf("invalid fromUnit: %s", fromUnit)
	}

	var result float64
	switch toUnit {
	case "C":
		result = celsiusValue
	case "F":
		result = celsiusToFahrenheit(celsiusValue)
	case "K":
		result = celsiusToKelvin(celsiusValue)
	default:
		return 0, fmt.Errorf("invalid toUnit: %s", toUnit)
	}

	return result, nil
}

func main() {
    fmt.Println("=== Temperature Converter ===")
    
    // Test conversions
    fmt.Printf("25°C = %.2f°F\n", celsiusToFahrenheit(25))
    fmt.Printf("77°F = %.2f°C\n", fahrenheitToCelsius(77))
    fmt.Printf("0°C = %.2fK\n", celsiusToKelvin(0))
    
    // Test generic converter
    if result, err := convertTemperature(100, "C", "F"); err == nil {
        fmt.Printf("100°C = %.2f°F\n", result)
    }
}
