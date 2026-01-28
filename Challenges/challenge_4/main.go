package main

import "fmt"

/*
BMI (Body Mass Index) = weight (kg) / (height (m))^2

Categories:
- Underweight: < 18.5
- Normal: 18.5 - 24.9
- Overweight: 25 - 29.9
- Obese: >= 30
*/

func calculateBMI(weight, height float64) float64 {
    return weight / (height * height)
}

func getBMICategory(bmi float64) string {
    // YOUR CODE HERE
    // Return "Underweight", "Normal", "Overweight", or "Obese"
    if bmi < 18.5 {
        return "Underweight"
    } else if bmi < 25 {
        return "Normal"
    } else if bmi < 30 {
        return "Overweight"
    } else {
        return "Obese"
    }
}

func main() {
    var weight, height float64
    
    fmt.Print("Enter weight (kg): ")
    fmt.Scan(&weight)
    
    fmt.Print("Enter height (m): ")
    fmt.Scan(&height)
    
    // Calculate BMI and display result with category
    // Format: "Your BMI is 22.5 (Normal)"

	bmi := calculateBMI(weight, height)
	category := getBMICategory(bmi)
	fmt.Printf("Your BMI is %.2f (%s)\n", bmi, category)

}
